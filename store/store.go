package store

import "github.com/drhamilton/enrollment-rules-svc/engine"

// RuleStore fetches the rule set that applies to a given product type.
// The interface exists so a YAML- or DB-backed implementation can replace
// the in-memory store without touching the server or engine.
type RuleStore interface {
	GetRules(productType string) ([]engine.Rule, error)
}
