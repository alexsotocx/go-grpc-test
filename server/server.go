package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/alexsotocx/go-grpc-test/proto"
	"google.golang.org/grpc"
)

var serverName = os.Getenv("SERVER_NAME")

const (
	port = 8000
)

type helloServer struct {
	pb.UnimplementedHelloServer
}

func (s *helloServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received hello request")
	return &pb.GreetResponse{Message: fmt.Sprintf("Hello, from %s", serverName)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))

	log.Printf("Staring server with params %s", serverName)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening...")

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	s := new(helloServer)
	pb.RegisterHelloServer(grpcServer, s)
	grpcServer.Serve(lis)
}
