package main

import (
	"log"
	"net"
	"net/http"
    "flag"

    "github.com/Cytram/grpc-test/pkg/api"
	pb "github.com/Cytram/grpc-test/proto/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
    processingService string
)

const (
	port = ":50051"
    host = "localhost:50052"
)

// starts the Prometheus stats endpoint server
func startPromHTTPServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("prometheus err", port)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go startPromHTTPServer("5001")
    conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

    Grayscale := flag.Bool("grayscale", false, "Whether or not to grayscale image")
    Scale := flag.String("scale", "", "Whether or not to scale image and what size")

    flag.Parse()

	s := grpc.NewServer()
    pb.RegisterImageScalerServer(s, &api.Server{Grayscale: *Grayscale, Scale: *Scale})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
