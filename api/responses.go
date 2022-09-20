package api

type Response struct {
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

type BalanceSheetDetail struct {
	Year         uint   `json:"year"`
	Month        uint   `json:"month"`
	ProfitOrLoss int64  `json:"profit_or_loss"` // TODO change to decimal
	AssetsValue  uint64 `json:"assets_value"`
}

type LoanApplicationResult struct {
	Verdict            bool    `json:"verdict"`
	PreAssessmentValue *uint   `json:"pre_assessment_value,omitempty"`
	EligibleLoanAmount *uint64 `json:"eligible_loan_amount,omitempty"`
}
