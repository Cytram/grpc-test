package api

import (
	"context"
	"fmt"
    "log"

	api "github.com/Cytram/grpc-test/proto/api"
	pb "github.com/Cytram/grpc-test/proto/process"
    "google.golang.org/grpc"
)

const (
	host = "localhost:50052"
    scale = "1024x768"
    grayscale = "0"
)

// Server is a server implementing the proto API
type Server struct{}

// ScaleImage echoes the image provides in the request
func (s *Server) ScaleImage(ctx context.Context, req *api.ScaleImageRequest) (
	*api.ScaleImageReply, error) {
	// Echo
	fmt.Println("Recieved image...")
    fmt.Printf("%t\n", grayscale)

    conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewProcessorClient(conn)

    ctx = context.Background()
	resp, err := client.ProcessImage(ctx, &pb.ProcessImageRequest{
		Image: &pb.Image{
            Content: req.Image.GetContent(),
            Source: &pb.ImageSource{
                HttpUri: req.Image.Source.GetHttpUri(),
            },
		},
        Scale: scale,
        Grayscale: grayscale,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &api.ScaleImageReply{
		Content: resp.GetContent(),
	}, nil
}
