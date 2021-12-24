// Package main implements a client for DataAnalysis service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
    pb "github.com/hzget/analysisdriver"
)

const (
	address     = "192.168.1.12:50051"
	defaultName = "abc"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

    c := pb.NewDataAnalysisClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AnalyzeByAuthor(ctx, &pb.Author{Name: name})
	if err != nil {
		log.Fatalf("could not analyze: %v", err)
	}

	log.Printf("Analysis result: %d\n", r.GetScore())
}
