// This file was auto-generated by Fern from our API Definition.

package accounting

import (
	monitegoclient "github.com/team-monite/monite-go-client"
)

type LedgerAccountsGetRequest struct {
	// Order by
	Order *monitegoclient.OrderEnum `json:"-" url:"order,omitempty"`
	// Max is 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// A token, obtained from previous page. Prior over other filters
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// Allowed sort fields
	Sort *monitegoclient.LedgerAccountCursorFields `json:"-" url:"sort,omitempty"`
}
