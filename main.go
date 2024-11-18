package main

import (
	"log"
	"net"

	"github.com/kenmobility/grpc-project/config"
	"github.com/kenmobility/grpc-project/db"
	"github.com/kenmobility/grpc-project/gapi"
	"github.com/kenmobility/grpc-project/models"
	"github.com/kenmobility/grpc-project/pb"
	"github.com/kenmobility/grpc-project/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// load env variables
	config := config.LoadConfig("")

	db, err := db.ConnectPostgresDb(*config)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	store := repo.NewGormStoreRepository(db)

	// Run migrations
	models.Migrate(db)

	runGrpcServer(*config, store)
}

func runGrpcServer(config config.Config, store repo.StoreRepository) {

	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Enable reflection
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	pb.RegisterUserServiceServer(grpcServer, server)
	pb.RegisterOrderServiceServer(grpcServer, server)

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server: ", err)
	}
}
