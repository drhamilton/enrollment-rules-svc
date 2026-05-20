package engine

import "fmt"

type Applicant struct {
	ID          string
	Age         int
	State       string
	ProductType string
	RiskScore   float32
}

type Rule struct {
	Name          string
	ConditionType string
	Params        map[string]any
}

type Decision struct {
	RuleName string
	Passed   bool
	Reason   string
}

type RuleFunc func(rule Rule, applicant Applicant) Decision

var registry = map[string]RuleFunc{}

func Register(conditionType string, fn RuleFunc) {
	registry[conditionType] = fn
}

func Evaluate(rule Rule, applicant Applicant) Decision {
	fn, ok := registry[rule.ConditionType]
	if !ok {
		return Decision{RuleName: rule.Name, Passed: false, Reason: fmt.Sprintf("unknown condition type %q", rule.ConditionType)}
	}
	return fn(rule, applicant)
}

