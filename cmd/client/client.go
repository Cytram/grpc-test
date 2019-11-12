package main

import (
	"context"
	"io/ioutil"
	"log"
    "flag"

	pb "github.com/Cytram/grpc-test/proto/api"
	"google.golang.org/grpc"
)

const (
	host = "78.47.84.126:50051"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewImageScalerClient(conn)

    imageURL := flag.String("url", "", "URL for where to download image")
    flag.Parse()

	image, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		log.Fatal("Couldn't read input image")
	}
	ctx := context.Background()
	resp, err := client.ScaleImage(ctx, &pb.ScaleImageRequest{
		Image: &pb.Image{
			Content: image,
            Source: &pb.ImageSource{
                HttpUri: *imageURL,
            },
		},
	})
	if err != nil {
		log.Fatal("error")
	}
	ioutil.WriteFile("out.jpg", resp.GetContent(), 0644)
}
