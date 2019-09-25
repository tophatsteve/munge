package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/tophatsteve/munge/capitalise"
	"google.golang.org/grpc"
)

var serverPort string

func init() {
	serverPort = os.Getenv("PORT")
}

func main() {

	if serverPort == "" {
		serverPort = "8120"
	}

	log.Printf("Starting capitalise on port %s", serverPort)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":"+serverPort)

	if err != nil {
		panic(err)
	}

	log.Printf("Start grpc server")

	grpcServer := grpc.NewServer()

	log.Printf("Register capitalise server")

	capitalise.RegisterCapitaliseServer(grpcServer, capitalise.NewService())

	go func() {
		grpcServer.Serve(lis)
	}()

	<-stop

	log.Printf("Stopping capitalise service")

	grpcServer.GracefulStop()
}
