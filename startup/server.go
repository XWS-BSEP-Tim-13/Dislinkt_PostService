package startup

import (
	"context"
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/api"
	post "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/persistence"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/logger"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/startup/config"
	logg "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	logger := logger.InitLogger("post-service", context.TODO())

	mongoClient := server.initMongoClient()
	productStore := server.initPostStore(mongoClient)
	imageStore := server.initUploadImageStore()
	productService := server.initPostService(productStore, imageStore, logger)
	productHandler := server.initPostHandler(productService, logger)

	server.startGrpcServer(productHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	store.DeleteAll()
	for _, post := range posts {
		err := store.Insert(post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUploadImageStore() domain.UploadImageStore {
	imageStore := persistence.NewUploadImageStore(server.config.SecretAccessKey, server.config.AccessKey)
	imageStore.Start()
	return imageStore
}

func (server *Server) initPostService(store domain.PostStore, imageStore domain.UploadImageStore, logger *logger.Logger) *application.PostService {
	return application.NewPostService(store, imageStore, logger)
}

func (server *Server) initPostHandler(service *application.PostService, logger *logger.Logger) *api.PostHandler {
	return api.NewPostHandler(service, logger)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		logg.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
