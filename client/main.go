package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/richardimaoka/go-grpc-streaming/proto"
)

var addr string = "localhost:50051"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func pollCommands(c pb.ExecCommandStreamingClient) {
	log.Println("pollCommands was invoked")

	req := &pb.RegisterClient{
		CurrentDirectory:     "",
		TernminalClientToken: "",
		AppPageToken:         "",
	}
	stream, err := c.PollCommands(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Erro while reading from the stream: %v\n", err)
		}

		log.Printf("Command: %s\n", msg.Command)
		cmd := exec.Command("sh", "-c", msg.Command)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			fmt.Printf("ERROR: %v, %s\n", err, stderr.Bytes())
			continue
		}
		fmt.Printf("%s\n", stdout.Bytes())
	}
}

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get wd")
	}
	fmt.Println("wd = ", curDir)

	secretToken := RandomString(40)
	fmt.Println("please copy and paste this :", secretToken)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewExecCommandStreamingClient(conn)
	pollCommands(c)

}
