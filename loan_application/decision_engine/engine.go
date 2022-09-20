package decision_engine

import "context"

type Engine interface {
	RequestLoanDecision(ctx context.Context, params *RequestLoanDecisionParams) *LoanDecisionOutcome
}

func NewDecisionEngine() Engine {
	return &dummyDecisionEngine{}
}

type dummyDecisionEngine struct {
}

func (e *dummyDecisionEngine) RequestLoanDecision(ctx context.Context, params *RequestLoanDecisionParams) *LoanDecisionOutcome {
	return &LoanDecisionOutcome{
		Verdict: params.PreAssessmentValue > 30,
	}
}
