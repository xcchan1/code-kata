package accounting_providers

import (
	"context"

	"code-kata/accounting_providers/model"
)

func RetrieveBalanceSheet(ctx context.Context, providerID string, params *model.RetrieveBalanceSheetParams) (*model.BalanceSheet, error) {
	provider, err := GetProvider(providerID)
	if err != nil {
		return nil, err
	}
	balanceSheet, err := provider.RetrieveBalanceSheet(ctx, params)
	return balanceSheet, err
}
