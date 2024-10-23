package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Pidu2/anker/sender/mail"
	"google.golang.org/grpc"
)

var APP_PORT = os.Getenv("APP_PORT")

func main() {
	//send("subjekkt", "my teggscht", []string{"beatschaerz@hotmail.com"})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", APP_PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", APP_PORT, err)
	}

	s := mail.Server{}

	grpcServer := grpc.NewServer()

	mail.RegisterMailServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v", err, APP_PORT)
	}

}
