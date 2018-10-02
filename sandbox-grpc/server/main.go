package main

import (
	"fmt"
	"log"
	"net"
	"sandbox-grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := api.Server{}

	// create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a gRPC server object
	gprcServer := grpc.NewServer(opts...)

	// attach the Ping service to the server
	api.RegisterPingServer(gprcServer, &s)

	// start the server
	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
