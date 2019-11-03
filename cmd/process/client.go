package main

import (
	"context"
	"io/ioutil"
	"log"

	pb "github.com/Cytram/grpc-test/proto/process"
	"google.golang.org/grpc"
)

const (
	host = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProcessorClient(conn)

	image, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		log.Fatal("Couldn't read input image")
	}
	ctx := context.Background()
	resp, err := client.ProcessImage(ctx, &pb.ProcessImageRequest{
		Image: &pb.Image{
			Content: image,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("out.jpg", resp.GetContent(), 0644)
}
