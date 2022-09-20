package accounting_providers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"code-kata/accounting_providers/myob"
)

func TestGetProvider(t *testing.T) {
	provider, err := GetProvider(ProviderMYOB)
	assert.NoError(t, err)
	_, ok := provider.(*myob.MYOBProvider)
	assert.True(t, ok)
}

func TestGetProvider_InvalidProvider(t *testing.T) {
	_, err := GetProvider("aaaa")
	assert.Error(t, err)
	pnfErr, ok := err.(*ProviderNotFoundError)
	assert.True(t, ok)
	assert.Equal(t, "aaaa", pnfErr.ProviderID)
}
