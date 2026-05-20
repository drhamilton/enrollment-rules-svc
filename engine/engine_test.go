package engine_test

import (
	"testing"

	"github.com/drhamilton/enrollment-rules-svc/engine"
)

func TestEvaluate_MinAge(t *testing.T) {
	rule := engine.Rule{
		Name:          "min_age",
		ConditionType: "min_age",
		Params:        map[string]any{"value": 18},
	}

	d := engine.Evaluate(rule, engine.Applicant{Age: 17})
	if d.Passed {
		t.Errorf("Expected rule to fail for age 17, but it passed")
	}

	d = engine.Evaluate(rule, engine.Applicant{Age: 18})
	if !d.Passed {
		t.Errorf("Expected rule to pass for age 18, but it failed")
	}
}
