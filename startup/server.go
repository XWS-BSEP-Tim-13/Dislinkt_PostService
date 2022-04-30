package startup

import (
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/application"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/domain"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/api"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/infrastructure/persistence"
	"github.com/XWS-BSEP-Tim-13/Dislinkt_PostService/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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
	productStore := server.initProductStore(mongoClient)

	productService := server.initProductService(productStore)

	productHandler := server.initProductHandler(productService)

	server.startGrpcServer(productHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initProductStore(client *mongo.Client) domain.ProductStore {
	store := persistence.NewProductMongoDBStore(client)
	//store.DeleteAll()
	//for _, product := range products {
	//	err := store.Insert(product)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	return store
}

func (server *Server) initProductService(store domain.ProductStore) *application.ProductService {
	return application.NewProductService(store)
}

func (server *Server) initProductHandler(service *application.ProductService) *api.PostHandler {
	return api.NewProductHandler(service)
}

func (server *Server) startGrpcServer(productHandler *api.PostHandler) {
	//listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//grpcServer := grpc.NewServer()
	//post.RegisterPostServiceServer(grpcServer, productHandler)
	//if err := grpcServer.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
}
