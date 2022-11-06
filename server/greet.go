package main

import (
	"context"
	"log"

	pb "github.com/richardimaoka/go-grpc-streaming/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func repeatPull(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "abc"
	}
	close(ch)
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with :%v\n", in)

	ch := make(chan string)
	go repeatPull(ch)
	for {
		s, ok := <-ch
		if !ok {
			break
		}

		stream.Send(&pb.GreetResponse{
			Result: s,
		})

	}
	return nil
}
