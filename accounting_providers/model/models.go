package model

type RetrieveBalanceSheetParams struct {
	BusinessName string
}

type BalanceSheet struct {
	Details []*BalanceSheetDetail
}

type BalanceSheetDetail struct {
	Year         uint
	Month        uint
	ProfitOrLoss int64
	AssetsValue  uint64
}

type GetBusinessYearlySummaryParams struct {
	BusinessName string
}
