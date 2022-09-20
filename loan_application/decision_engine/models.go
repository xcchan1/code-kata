package decision_engine

type RequestLoanDecisionParams struct {
	BusinessDetails    *BusinessDetails
	PreAssessmentValue uint
}

type BusinessDetails struct {
	Name              string
	YearEstablished   uint
	ProfitLossSummary *BusinessYearlySummary
}

type BusinessYearlySummary struct {
	Summary []*BusinessSummary
}

type BusinessSummary struct {
	Year         uint
	ProfitOrLoss int64
}

type LoanDecisionOutcome struct {
	Verdict bool
}
