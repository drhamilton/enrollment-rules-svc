package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/drhamilton/enrollment-rules-svc/gen/enrollment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "server address")
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRulesServiceClient(conn)
	applicant := &pb.Applicant{
		Id:          "demo-001",
		Age:         32,
		State:       "ME",
		ProductType: "term_life",
		RiskScore:   5.5,
	}

	decision, err := client.EvaluateApplicant(context.Background(), applicant)
	if err != nil {
		log.Fatalf("EvaluateApplicant: %v", err)
	}
	fmt.Printf("eligible=%v reason=%s\n", decision.Eligible, decision.Reason)
}
