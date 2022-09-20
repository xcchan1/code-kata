package api

import (
	"code-kata/accounting_providers/model"
)

func toRetrieveBalanceSheetResponse(balanceSheet *model.BalanceSheet) Response {
	details := make([]*BalanceSheetDetail, 0, len(balanceSheet.Details))
	for _, detail := range balanceSheet.Details {
		details = append(details, &BalanceSheetDetail{
			Year:         detail.Year,
			Month:        detail.Month,
			ProfitOrLoss: detail.ProfitOrLoss,
			AssetsValue:  detail.AssetsValue,
		})
	}
	return Response{
		Data: details,
	}
}
