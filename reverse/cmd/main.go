package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/tophatsteve/munge/reverse"
	"google.golang.org/grpc"
)

var serverPort string

func init() {
	serverPort = os.Getenv("PORT")
}

func main() {

	if serverPort == "" {
		serverPort = "80"
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":"+serverPort)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reverse.RegisterReverseServer(grpcServer, reverse.NewService())

	go func() {
		grpcServer.Serve(lis)
	}()

	<-stop
	grpcServer.GracefulStop()
}
