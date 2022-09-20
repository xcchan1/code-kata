package loan_application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"code-kata/accounting_providers/model"
)

func TestCalculatePreAssessmentValue_Profit(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
	}
	pav := calculatePreAssessmentValue(ctx, 100000, mockData)
	assert.Equal(t, pav, uint(60))
}

func TestCalculatePreAssessmentValue_Profit_MoreThan12Months(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
		{
			Year:         2020,
			Month:        8,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        7,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        6,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        5,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
		{
			Year:         2020,
			Month:        4,
			ProfitOrLoss: 250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        3,
			ProfitOrLoss: 1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        2,
			ProfitOrLoss: 2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        1,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
		{
			Year:         2019,
			Month:        12,
			ProfitOrLoss: -187000000,
			AssetsValue:  0,
		},
	}
	pav := calculatePreAssessmentValue(ctx, 100000, mockData)
	assert.Equal(t, pav, uint(60))
}

func TestCalculatePreAssessmentValue_Loss_AssetValueGreaterThanLoanAmount(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: -250000,
			AssetsValue:  1234,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: -1150,
			AssetsValue:  5789,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: -2500,
			AssetsValue:  22345,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  223452,
		},
	}
	pav := calculatePreAssessmentValue(ctx, 100, mockData)
	assert.Equal(t, pav, uint(100))
}

func TestCalculatePreAssessmentValue_Loss_AssetValueLesserThanLoanAmount(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: -250000,
			AssetsValue:  12,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: -1150,
			AssetsValue:  12,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: -2500,
			AssetsValue:  12,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  12,
		},
	}
	pav := calculatePreAssessmentValue(ctx, 100, mockData)
	assert.Equal(t, pav, uint(20))
}

func TestCalculatePreAssessmentValue_Loss_AssetValueEqualLoanAmount(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: -250000,
			AssetsValue:  100,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: -1150,
			AssetsValue:  100,
		},
		{
			Year:         2020,
			Month:        10,
			ProfitOrLoss: -2500,
			AssetsValue:  100,
		},
		{
			Year:         2020,
			Month:        9,
			ProfitOrLoss: -187000,
			AssetsValue:  100,
		},
	}
	pav := calculatePreAssessmentValue(ctx, 100, mockData)
	assert.Equal(t, pav, uint(20))
}

func TestGenerateBusinessYearlySummary(t *testing.T) {
	ctx := context.Background()
	mockData := []*model.BalanceSheetDetail{
		{
			Year:         2021,
			Month:        2,
			ProfitOrLoss: 1000,
			AssetsValue:  100,
		},
		{
			Year:         2021,
			Month:        1,
			ProfitOrLoss: 1000,
			AssetsValue:  100,
		},
		{
			Year:         2020,
			Month:        12,
			ProfitOrLoss: 2500,
			AssetsValue:  100,
		},
		{
			Year:         2020,
			Month:        11,
			ProfitOrLoss: 2500,
			AssetsValue:  100,
		},
	}
	summary := generateBusinessYearlySummary(ctx, mockData)
	assert.Equal(t, 2, len(summary.Summary))

	assert.Equal(t, uint(2020), summary.Summary[0].Year)
	assert.Equal(t, int64(5000), summary.Summary[0].ProfitOrLoss)

	assert.Equal(t, uint(2021), summary.Summary[1].Year)
	assert.Equal(t, int64(2000), summary.Summary[1].ProfitOrLoss)
}
