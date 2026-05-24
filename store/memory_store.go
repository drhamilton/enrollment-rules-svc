package store

import "github.com/drhamilton/enrollment-rules-svc/engine"

// MemoryStore is a trivial in-memory RuleStore for the tracer-bullet build.
// Returns the same rule set regardless of productType — product filtering
// lands with the YAML store.
type MemoryStore struct {
	rules []engine.Rule
}

func NewMemoryStore(rules []engine.Rule) *MemoryStore {
	return &MemoryStore{rules: rules}
}

func (s *MemoryStore) GetRules(_ string) ([]engine.Rule, error) {
	return s.rules, nil
}
