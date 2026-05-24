package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/drhamilton/enrollment-rules-svc/engine"
	pb "github.com/drhamilton/enrollment-rules-svc/gen/enrollment/v1"
	"github.com/drhamilton/enrollment-rules-svc/server"
	"github.com/drhamilton/enrollment-rules-svc/store"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 50051, "gRPC server port")
	flag.Parse()

	ruleStore := store.NewMemoryStore([]engine.Rule{
		{Name: "min_age", ConditionType: "min_age", Params: map[string]any{"value": 18}},
		{Name: "max_age", ConditionType: "max_age", Params: map[string]any{"value": 70}},
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRulesServiceServer(s, server.New(ruleStore))

	log.Printf("enrollment-rules-svc listening on :%d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
