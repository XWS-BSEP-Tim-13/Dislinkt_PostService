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
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/tracer"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	otgo "github.com/opentracing/opentracing-go"
	logg "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	config *config.Config
	tracer otgo.Tracer
	closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init()
	otgo.SetGlobalTracer(tracer)

	return &Server{
		config: config,
		tracer: tracer,
		closer: closer,
	}
}

const (
	serverCertFile = "cert/cert.pem"
	serverKeyFile  = "cert/key.pem"
	clientCertFile = "cert/client-cert.pem"
)

func (server *Server) Start() {
	logger := logger.InitLogger("post-service", context.TODO())

	mongoClient := server.initMongoClient()
	postStore := server.initPostStore(mongoClient)
	messageStore := server.initMessageStore(mongoClient)
	imageStore := server.initUploadImageStore()
	eventStore := server.initEventStore(mongoClient)
	postService := server.initPostService(postStore, imageStore, logger, messageStore, eventStore)
	postHandler := server.initPostHandler(postService, logger)

	server.startGrpcServer(postHandler)
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
	store.DeleteAll(context.TODO())
	for _, post := range posts {
		err := store.Insert(context.TODO(), post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	store.DeleteAll()
	for _, message := range messages {
		err := store.Insert(message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUploadImageStore() domain.UploadImageStore {
	imageStore := persistence.NewUploadImageStore(server.config.SecretAccessKey, server.config.AccessKey)
	imageStore.Start(context.TODO())
	return imageStore
}

func (server *Server) initEventStore(client *mongo.Client) domain.EventStore {
	store := persistence.NewEventMongoDBStore(client)
	store.DeleteAll()
	for _, event := range events {
		err := store.Insert(event)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPostService(store domain.PostStore, imageStore domain.UploadImageStore, logger *logger.Logger, messageStore domain.MessageStore, eventStore domain.EventStore) *application.PostService {
	return application.NewPostService(store, imageStore, logger, messageStore, eventStore)
}

func (server *Server) initPostHandler(service *application.PostService, logger *logger.Logger) *api.PostHandler {
	return api.NewPostHandler(service, logger)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler) {
	/*cert, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	pemClientCA, err := ioutil.ReadFile(clientCertFile)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	}*/

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(server.tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(server.tracer)),
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		logg.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
