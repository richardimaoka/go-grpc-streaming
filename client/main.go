package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

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

	c := pb.NewGreetServiceClient(conn)

	doGreetManyTimes(c)
}
