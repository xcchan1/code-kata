package loan_application

type LoanApplicationParams struct {
	BusinessName    string
	YearEstablished uint
	LoanAmount      uint64
}

type LoanApplicationResult struct {
	Outcome            bool
	PreAssessmentValue uint
	EligibleLoanAmount uint64
}
