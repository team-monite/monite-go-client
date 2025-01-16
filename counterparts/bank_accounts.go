// This file was auto-generated by Fern from our API Definition.

package counterparts

import (
	monitegoclient "github.com/team-monite/monite-go-client"
)

type CreateCounterpartBankAccount struct {
	// The name of the person or business that owns this bank account. Required for US bank accounts to accept ACH payments.
	AccountHolderName *string `json:"account_holder_name,omitempty" url:"-"`
	// The bank account number. Required for US bank accounts to accept ACH payments. US account numbers contain 9 to 12 digits. UK account numbers typically contain 8 digits.
	AccountNumber *string `json:"account_number,omitempty" url:"-"`
	// The BIC/SWIFT code of the bank.
	Bic      *string                         `json:"bic,omitempty" url:"-"`
	Country  monitegoclient.AllowedCountries `json:"country" url:"-"`
	Currency monitegoclient.CurrencyEnum     `json:"currency" url:"-"`
	// The IBAN of the bank account.
	Iban                 *string `json:"iban,omitempty" url:"-"`
	IsDefaultForCurrency *bool   `json:"is_default_for_currency,omitempty" url:"-"`
	Name                 *string `json:"name,omitempty" url:"-"`
	// Metadata for partner needs.
	PartnerMetadata map[string]interface{} `json:"partner_metadata,omitempty" url:"-"`
	// The bank's routing transit number (RTN). Required for US bank accounts to accept ACH payments. US routing numbers consist of 9 digits.
	RoutingNumber *string `json:"routing_number,omitempty" url:"-"`
	// The bank's sort code.
	SortCode *string `json:"sort_code,omitempty" url:"-"`
}

type UpdateCounterpartBankAccount struct {
	// The name of the person or business that owns this bank account. Required for US bank accounts to accept ACH payments.
	AccountHolderName *string `json:"account_holder_name,omitempty" url:"-"`
	// The bank account number. Required for US bank accounts to accept ACH payments. US account numbers contain 9 to 12 digits. UK account numbers typically contain 8 digits.
	AccountNumber *string `json:"account_number,omitempty" url:"-"`
	// The BIC/SWIFT code of the bank.
	Bic      *string                          `json:"bic,omitempty" url:"-"`
	Country  *monitegoclient.AllowedCountries `json:"country,omitempty" url:"-"`
	Currency *monitegoclient.CurrencyEnum     `json:"currency,omitempty" url:"-"`
	// The IBAN of the bank account.
	Iban *string `json:"iban,omitempty" url:"-"`
	Name *string `json:"name,omitempty" url:"-"`
	// Metadata for partner needs.
	PartnerMetadata map[string]interface{} `json:"partner_metadata,omitempty" url:"-"`
	// The bank's routing transit number (RTN). Required for US bank accounts to accept ACH payments. US routing numbers consist of 9 digits.
	RoutingNumber *string `json:"routing_number,omitempty" url:"-"`
	// The bank's sort code.
	SortCode *string `json:"sort_code,omitempty" url:"-"`
}
