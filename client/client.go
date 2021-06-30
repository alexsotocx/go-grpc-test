package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/alexsotocx/go-grpc-test/proto"
	"google.golang.org/grpc"
)

var (
	serverAddress = os.Getenv("SERVER_ADDRESS")
	serverPort    = os.Getenv("SERVER_PORT")
)

func executeInInterval(client pb.HelloClient) {
	ticker := time.Tick(5 * time.Second)
	for range ticker {
		response, err := client.Greet(context.Background(), &pb.GreetRequest{})
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Response from server %s", response.Message)
		}
	}
}

func main() {
	connectionString := fmt.Sprintf("%s:%s", serverAddress, serverPort)
	log.Printf("Connecting to server on %s", connectionString)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(connectionString, opts...)
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloClient(conn)
	executeInInterval(client)
}
