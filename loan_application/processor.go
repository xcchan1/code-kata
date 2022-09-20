package loan_application

import (
	"context"
	"sort"

	"code-kata/accounting_providers"
	"code-kata/accounting_providers/model"
	"code-kata/loan_application/decision_engine"
)

const (
	loanProfitWindowMonths = 12
)

type LoanApplicationProcessor interface {
	SubmitLoanApplication(ctx context.Context, params *LoanApplicationParams) (*LoanApplicationResult, error)
}

type loanApplicationProcessor struct {
	accountingProvider accounting_providers.Provider
	decisionEngine     decision_engine.Engine
}

func NewLoanApplicationProcessor(accountingProviderID string) (LoanApplicationProcessor, error) {
	provider, err := accounting_providers.GetProvider(accountingProviderID)
	if err != nil {
		return nil, err
	}
	return &loanApplicationProcessor{
		accountingProvider: provider,
		decisionEngine:     decision_engine.NewDecisionEngine(),
	}, nil
}

func (p *loanApplicationProcessor) SubmitLoanApplication(ctx context.Context, params *LoanApplicationParams) (*LoanApplicationResult, error) {
	getBalanceSheetParams := model.RetrieveBalanceSheetParams{
		BusinessName: params.BusinessName,
	}
	balanceSheet, err := p.accountingProvider.RetrieveBalanceSheet(ctx, &getBalanceSheetParams)
	if err != nil {
		return nil, err
	}
	yearlySummary := generateBusinessYearlySummary(ctx, balanceSheet.Details)
	preAssessmentValue := calculatePreAssessmentValue(ctx, params.LoanAmount, balanceSheet.Details)
	loanDecisionParams := decision_engine.RequestLoanDecisionParams{
		BusinessDetails: &decision_engine.BusinessDetails{
			Name:              params.BusinessName,
			YearEstablished:   params.YearEstablished,
			ProfitLossSummary: yearlySummary,
		},
		PreAssessmentValue: preAssessmentValue,
	}
	decisionOutcome := p.decisionEngine.RequestLoanDecision(ctx, &loanDecisionParams)
	result := LoanApplicationResult{
		Outcome: decisionOutcome.Verdict,
	}
	if decisionOutcome.Verdict {
		result.PreAssessmentValue = preAssessmentValue
		result.EligibleLoanAmount = calculateEligibleLoanAmount(ctx, params.LoanAmount, preAssessmentValue)
	}
	return &result, nil
}

// Assumes that the data obtained from accounting providers are sorted in descending order by year-month.
// As an improvement, we could enforce this in the accounting_providers package.
func calculatePreAssessmentValue(ctx context.Context, loanAmount uint64, balanceSheet []*model.BalanceSheetDetail) uint {
	totalAssetValue := uint64(0)
	profitOrLoss := int64(0)
	limit := loanProfitWindowMonths
	if len(balanceSheet) < limit {
		limit = len(balanceSheet)
	}
	for i := 0; i < limit; i++ {
		totalAssetValue += balanceSheet[i].AssetsValue
		profitOrLoss += balanceSheet[i].ProfitOrLoss
	}
	if profitOrLoss > 0 {
		return 60
	}
	avgAssetValue := totalAssetValue / uint64(limit)
	if avgAssetValue > loanAmount {
		return 100
	}
	return 20
}

func generateBusinessYearlySummary(ctx context.Context, balanceSheet []*model.BalanceSheetDetail) *decision_engine.BusinessYearlySummary {
	summaryByYear := map[uint]*decision_engine.BusinessSummary{}
	for _, detail := range balanceSheet {
		summary, ok := summaryByYear[detail.Year]
		if !ok {
			summary = &decision_engine.BusinessSummary{
				Year: detail.Year,
			}
			summaryByYear[detail.Year] = summary
		}
		summary.ProfitOrLoss += detail.ProfitOrLoss
	}

	summaries := make([]*decision_engine.BusinessSummary, 0, len(summaryByYear))
	for _, summary := range summaryByYear {
		summaries = append(summaries, summary)
	}

	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].Year < summaries[j].Year
	})

	return &decision_engine.BusinessYearlySummary{
		Summary: summaries,
	}
}

func calculateEligibleLoanAmount(ctx context.Context, loanAmount uint64, preAssessmentValue uint) uint64 {
	if preAssessmentValue == 100 {
		return loanAmount
	}
	percent := float64(preAssessmentValue) / 100.0
	eligibleLoanAmount := float64(loanAmount) * percent
	return uint64(eligibleLoanAmount)
}
