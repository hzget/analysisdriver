// Package main implements a client for DataAnalysis service.
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/hzget/analysisdriver"
	"google.golang.org/grpc"
)

const (
	//address     = "192.168.1.12:50051"
	address     = "localhost:50051"
	defaultName = "abc"
	id          = 123
	text        = "let's play tennis"
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

	r2, err := c.AnalyzeByPostId(ctx, &pb.Id{Id: id})
	if err != nil {
		log.Fatalf("could not analyze: %v", err)
	}

	log.Printf("AnalysisByPostId result: %s\n", r2.GetResult())

	r3, err := c.AnalyzePost(ctx, &pb.Text{Text: text})
	if err != nil {
		log.Fatalf("could not analyze: %v", err)
	}

	log.Printf("AnalysisPost result: %s\n", r3.GetResult())
}
