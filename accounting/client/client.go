// This file was auto-generated by Fern from our API Definition.

package client

import (
	connections "github.com/team-monite/monite-go-client/accounting/connections"
	ledgeraccounts "github.com/team-monite/monite-go-client/accounting/ledgeraccounts"
	payables "github.com/team-monite/monite-go-client/accounting/payables"
	receivables "github.com/team-monite/monite-go-client/accounting/receivables"
	syncedrecords "github.com/team-monite/monite-go-client/accounting/syncedrecords"
	taxrates "github.com/team-monite/monite-go-client/accounting/taxrates"
	core "github.com/team-monite/monite-go-client/core"
	internal "github.com/team-monite/monite-go-client/internal"
	option "github.com/team-monite/monite-go-client/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	Payables       *payables.Client
	Receivables    *receivables.Client
	Connections    *connections.Client
	SyncedRecords  *syncedrecords.Client
	TaxRates       *taxrates.Client
	LedgerAccounts *ledgeraccounts.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:         options.ToHeader(),
		Payables:       payables.NewClient(opts...),
		Receivables:    receivables.NewClient(opts...),
		Connections:    connections.NewClient(opts...),
		SyncedRecords:  syncedrecords.NewClient(opts...),
		TaxRates:       taxrates.NewClient(opts...),
		LedgerAccounts: ledgeraccounts.NewClient(opts...),
	}
}
