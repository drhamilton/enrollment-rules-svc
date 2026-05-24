package server

import (
	"context"

	"github.com/drhamilton/enrollment-rules-svc/engine"
	pb "github.com/drhamilton/enrollment-rules-svc/gen/enrollment/v1"
	"github.com/drhamilton/enrollment-rules-svc/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RulesServer struct {
	pb.UnimplementedRulesServiceServer
	store store.RuleStore
}

func New(s store.RuleStore) *RulesServer {
	return &RulesServer{store: s}
}

func toApplicant(a *pb.Applicant) engine.Applicant {
	return engine.Applicant{
		ID:          a.Id,
		Age:         int(a.Age),
		State:       a.State,
		ProductType: a.ProductType,
		RiskScore:   a.RiskScore,
	}
}

func (s *RulesServer) EvaluateApplicant(_ context.Context, req *pb.Applicant) (*pb.Decision, error) {
	rules, err := s.store.GetRules(req.ProductType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "loading rules: %v", err)
	}
	decisions := engine.EvaluateAll(rules, toApplicant(req))
	eligible, reason := engine.Verdict(decisions)
	return &pb.Decision{
		ApplicantId: req.Id,
		Eligible:    eligible,
		Reason:      reason,
	}, nil
}
