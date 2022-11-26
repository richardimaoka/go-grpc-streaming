package main

import (
	"log"
	"net"

	pb "github.com/richardimaoka/go-grpc-streaming/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.ExecCommandStreamingServer
}

func repeatCommand(ch chan string) {
	for i := 0; i < 2; i++ {
		ch <- "echo abc"
	}
	close(ch)
}

func sendCommands(ch chan string) {
	ch <- "echo abc"
	ch <- "docker pull nginx"
	ch <- "docker inspect nginx"
	close(ch)
}

func (s *Server) PollCommands(in *pb.RegisterClient, stream pb.ExecCommandStreaming_PollCommandsServer) error {
	log.Printf("GreetManyTimes function was invoked with :%v\n", in)

	ch := make(chan string)
	go sendCommands(ch)
	for {
		cmd, ok := <-ch
		if !ok {
			break
		}

		stream.Send(&pb.ExecCommand{
			Command: cmd,
		})

	}
	return nil
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterExecCommandStreamingServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
