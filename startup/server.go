package startup

import (
	"fmt"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/api"
	post "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/grpc/proto"
	mw "github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/middleware"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/persistence"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/startup/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_tags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
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
	mongoClient := server.initMongoClient()
	productStore := server.initPostStore(mongoClient)
	imageStore := server.initUploadImageStore()
	productService := server.initPostService(productStore, imageStore)

	productHandler := server.initPostHandler(productService)

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

func (server *Server) initPostService(store domain.PostStore, imageStore domain.UploadImageStore) *application.PostService {
	return application.NewPostService(store, imageStore)
}

func (server *Server) initPostHandler(service *application.PostService) *api.PostHandler {
	return api.NewPostHandler(service)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		logg.Fatalf("failed to listen: %v", err)
	}

	logger := logg.WithFields(logg.Fields{
		"goapi": "server",
	})

	CommonInterceptors := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(logger),
		mw.AuthInterceptor(),
		grpc_tags.UnaryServerInterceptor(),
	))

	opts := []grpc.ServerOption{
		CommonInterceptors,
	}

	grpcServer := grpc.NewServer(opts...)
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
