package engine

import "fmt"

func init() {
	Register("min_age", evaluateMinAge)
	Register("max_age", evaluateMaxAge)
	Register("resides_in_state", evaluateResidesInState)
}

func evaluateResidesInState(rule Rule, applicant Applicant) Decision {
	state := rule.Params["value"].(string)
	passed := applicant.State == state
	reason := "state requirement met"
	if !passed {
		reason = fmt.Sprintf("applicant resides in %q, required %q", applicant.State, state)
	}
	return Decision{RuleName: rule.Name, Passed: passed, Reason: reason}
}

func evaluateMinAge(rule Rule, applicant Applicant) Decision {
	min := rule.Params["value"].(int)
	passed := applicant.Age >= min
	reason := "age requirement met"
	if !passed {
		reason = fmt.Sprintf("age %d is below minimum %d", applicant.Age, min)
	}
	return Decision{RuleName: rule.Name, Passed: passed, Reason: reason}
}

func evaluateMaxAge(rule Rule, applicant Applicant) Decision {
	max := rule.Params["value"].(int)
	passed := applicant.Age <= max
	reason := "age requirement met"
	if !passed {
		reason = fmt.Sprintf("age %d exceeds maximum %d", applicant.Age, max)
	}
	return Decision{RuleName: rule.Name, Passed: passed, Reason: reason}
}
