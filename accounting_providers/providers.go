package accounting_providers

import (
	"context"
	"fmt"

	"code-kata/accounting_providers/model"
	"code-kata/accounting_providers/myob"
	"code-kata/accounting_providers/xero"
)

const (
	ProviderMYOB string = "MYOB"
	ProviderXero string = "Xero"
)

type ProviderNotFoundError struct {
	ProviderID string
}

func (e *ProviderNotFoundError) Error() string {
	return fmt.Sprintf("Provider %s not found", e.ProviderID)
}

func AllProviders() []string {
	return []string{
		ProviderMYOB,
		ProviderXero,
	}
}

type Provider interface {
	RetrieveBalanceSheet(ctx context.Context, params *model.RetrieveBalanceSheetParams) (*model.BalanceSheet, error)
}

func newMYOBProvider() Provider {
	return &myob.MYOBProvider{}
}

func newXeroProvider() Provider {
	return &xero.XeroProvider{}
}

var providers = map[string]func() Provider{
	ProviderMYOB: newMYOBProvider,
	ProviderXero: newXeroProvider,
}

func GetProvider(providerID string) (Provider, error) {
	providerFunc, found := providers[providerID]
	if found {
		provider := providerFunc()
		return provider, nil
	} else {
		return nil, &ProviderNotFoundError{ProviderID: providerID}
	}
}
