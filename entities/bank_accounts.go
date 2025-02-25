// This file was auto-generated by Fern from our API Definition.

package entities

import (
	monitegoclient "github.com/team-monite/monite-go-client"
)

type CreateEntityBankAccountRequest struct {
	// The name of the person or business that owns this bank account. Required if the account currency is GBP or USD.
	AccountHolderName *string `json:"account_holder_name,omitempty" url:"-"`
	// The bank account number. Required if the account currency is GBP or USD. UK account numbers typically contain 8 digits. US bank account numbers contain 9 to 12 digits.
	AccountNumber *string `json:"account_number,omitempty" url:"-"`
	// The bank name.
	BankName *string `json:"bank_name,omitempty" url:"-"`
	// The SWIFT/BIC code of the bank.
	Bic *string `json:"bic,omitempty" url:"-"`
	// The country in which the bank account is registered, repsesented as a two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country monitegoclient.AllowedCountries `json:"country" url:"-"`
	// The currency of the bank account, represented as a three-letter ISO [currency code](https://docs.monite.com/docs/currencies).
	Currency monitegoclient.CurrencyEnum `json:"currency" url:"-"`
	// User-defined name of this bank account, such as 'Primary account' or 'Savings account'.
	DisplayName *string `json:"display_name,omitempty" url:"-"`
	// The IBAN of the bank account. Required if the account currency is EUR.
	Iban *string `json:"iban,omitempty" url:"-"`
	// If set to `true` or if this is the first bank account added for the given currency, this account becomes the default one for its currency.
	IsDefaultForCurrency *bool `json:"is_default_for_currency,omitempty" url:"-"`
	// The bank's routing transit number (RTN). Required if the account currency is USD. US routing numbers consist of 9 digits.
	RoutingNumber *string `json:"routing_number,omitempty" url:"-"`
	// The bank's sort code. Required if the account currency is GBP.
	SortCode *string `json:"sort_code,omitempty" url:"-"`
}

type UpdateEntityBankAccountRequest struct {
	// The name of the person or business that owns this bank account. If the account currency is GBP or USD, the holder name cannot be changed to an empty string.
	AccountHolderName *string `json:"account_holder_name,omitempty" url:"-"`
	// User-defined name of this bank account, such as 'Primary account' or 'Savings account'.
	DisplayName *string `json:"display_name,omitempty" url:"-"`
}
