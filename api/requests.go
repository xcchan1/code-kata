package api

type RetreiveBalanceSheetRequest struct {
	BusinessName       string `json:"business_name" binding:"required"`
	AccountingProvider string `json:"accounting_provider" binding:"required"`
}

type SubmitLoanApplicationRequest struct {
	BusinessName       string `json:"business_name" binding:"required"`
	YearEstablished    uint   `json:"year_established" binding:"required,min=1800"`
	AccountingProvider string `json:"accounting_provider" binding:"required"`
	LoanAmount         uint64 `json:"loan_amount" binding:"required,min=1"`
}
