package main

import (
	"context"
	"log"

	pb "github.com/richardimaoka/go-grpc-streaming/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Richard",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting :%s\n", res.Result)
}
