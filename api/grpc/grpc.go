package grpc

import (
	"fmt"
	"log"
	"os"

	"github.com/adisnuhic/scrapper_api/config"
	"google.golang.org/grpc"
)

var clientConnection grpc.ClientConnInterface

// Init initialize grpc client
func Init(cfg *config.AppConfig) {
	env := os.Getenv("ENV")
	clientConnection = initGrpcClient(cfg.GRPCConnections[env])
}

func initGrpcClient(cfg config.GRPCConnection) grpc.ClientConnInterface {
	conn, err := grpc.Dial(cfg.Address, cfg.GRPCOptions...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GRPC client initialized successfully!")

	return conn
}

// Connection returns grpc client connection object
func Connection() grpc.ClientConnInterface {
	return clientConnection
}
