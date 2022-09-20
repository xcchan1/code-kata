package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code-kata/accounting_providers/model"
)

func TestToRetrieveBalanceSheetResponse(t *testing.T) {
	balanceSheet := model.BalanceSheet{
		Details: []*model.BalanceSheetDetail{
			{
				Year:         2022,
				Month:        1,
				ProfitOrLoss: 100,
				AssetsValue:  100,
			},
			{
				Year:         2022,
				Month:        2,
				ProfitOrLoss: 1000,
				AssetsValue:  1000,
			},
		},
	}
	resp := toRetrieveBalanceSheetResponse(&balanceSheet)
	details, ok := resp.Data.([]*BalanceSheetDetail)
	assert.True(t, ok)
	assert.Equal(t, len(balanceSheet.Details), len(details))
	for i, detail := range balanceSheet.Details {
		assert.Equal(t, detail.Year, details[i].Year)
		assert.Equal(t, detail.Month, details[i].Month)
		assert.Equal(t, detail.ProfitOrLoss, details[i].ProfitOrLoss)
		assert.Equal(t, detail.AssetsValue, details[i].AssetsValue)
	}
}
