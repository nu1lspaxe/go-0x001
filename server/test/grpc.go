package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:6000"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDigimonClient(conn)

	digimon, err := c.Create(context.Background(), &pb.CreateRequest{
		Name: "Agumon",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r, err := c.QueryStream(context.Background(), &pb.QueryRequest{Id: digimon.GetId()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		msg, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream read failed: %v", err)
		}

		fmt.Println(msg)
	}
}
