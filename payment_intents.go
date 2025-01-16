// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type PaymentIntentsGetRequest struct {
	// Order by
	Order *OrderEnum `json:"-" url:"order,omitempty"`
	// Max is 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// A token, obtained from previous page. Prior over other filters
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// Allowed sort fields
	Sort *PaymentIntentCursorFields `json:"-" url:"sort,omitempty"`
	// ID of a payable or receivable invoice. If provided, returns only payment intents associated with the specified invoice.
	ObjectId *string `json:"-" url:"object_id,omitempty"`
}

type PaymentIntentCursorFields string

const (
	PaymentIntentCursorFieldsId        PaymentIntentCursorFields = "id"
	PaymentIntentCursorFieldsCreatedAt PaymentIntentCursorFields = "created_at"
)

func NewPaymentIntentCursorFieldsFromString(s string) (PaymentIntentCursorFields, error) {
	switch s {
	case "id":
		return PaymentIntentCursorFieldsId, nil
	case "created_at":
		return PaymentIntentCursorFieldsCreatedAt, nil
	}
	var t PaymentIntentCursorFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PaymentIntentCursorFields) Ptr() *PaymentIntentCursorFields {
	return &p
}

type PaymentIntentHistory struct {
	Id              string    `json:"id" url:"id"`
	CreatedAt       time.Time `json:"created_at" url:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" url:"updated_at"`
	PaymentIntentId string    `json:"payment_intent_id" url:"payment_intent_id"`
	Status          string    `json:"status" url:"status"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentIntentHistory) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *PaymentIntentHistory) GetCreatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.CreatedAt
}

func (p *PaymentIntentHistory) GetUpdatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.UpdatedAt
}

func (p *PaymentIntentHistory) GetPaymentIntentId() string {
	if p == nil {
		return ""
	}
	return p.PaymentIntentId
}

func (p *PaymentIntentHistory) GetStatus() string {
	if p == nil {
		return ""
	}
	return p.Status
}

func (p *PaymentIntentHistory) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentIntentHistory) UnmarshalJSON(data []byte) error {
	type embed PaymentIntentHistory
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PaymentIntentHistory(unmarshaler.embed)
	p.CreatedAt = unmarshaler.CreatedAt.Time()
	p.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentIntentHistory) MarshalJSON() ([]byte, error) {
	type embed PaymentIntentHistory
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*p),
		CreatedAt: internal.NewDateTime(p.CreatedAt),
		UpdatedAt: internal.NewDateTime(p.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (p *PaymentIntentHistory) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PaymentIntentHistoryResponse struct {
	// Payment intent history
	Data []*PaymentIntentHistory `json:"data" url:"data"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentIntentHistoryResponse) GetData() []*PaymentIntentHistory {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PaymentIntentHistoryResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentIntentHistoryResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PaymentIntentHistoryResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaymentIntentHistoryResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentIntentHistoryResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PaymentIntentResponse struct {
	Id                    string                         `json:"id" url:"id"`
	UpdatedAt             time.Time                      `json:"updated_at" url:"updated_at"`
	Amount                int                            `json:"amount" url:"amount"`
	ApplicationFeeAmount  *int                           `json:"application_fee_amount,omitempty" url:"application_fee_amount,omitempty"`
	BatchPaymentId        *string                        `json:"batch_payment_id,omitempty" url:"batch_payment_id,omitempty"`
	Currency              string                         `json:"currency" url:"currency"`
	Invoice               *Invoice                       `json:"invoice,omitempty" url:"invoice,omitempty"`
	Object                *PaymentObject                 `json:"object,omitempty" url:"object,omitempty"`
	Payer                 *AccountResponse               `json:"payer,omitempty" url:"payer,omitempty"`
	PaymentLinkId         *string                        `json:"payment_link_id,omitempty" url:"payment_link_id,omitempty"`
	PaymentMethods        []MoniteAllPaymentMethodsTypes `json:"payment_methods" url:"payment_methods"`
	PaymentReference      *string                        `json:"payment_reference,omitempty" url:"payment_reference,omitempty"`
	Provider              *string                        `json:"provider,omitempty" url:"provider,omitempty"`
	Recipient             *RecipientAccountResponse      `json:"recipient" url:"recipient"`
	SelectedPaymentMethod *MoniteAllPaymentMethodsTypes  `json:"selected_payment_method,omitempty" url:"selected_payment_method,omitempty"`
	Status                string                         `json:"status" url:"status"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentIntentResponse) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *PaymentIntentResponse) GetUpdatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.UpdatedAt
}

func (p *PaymentIntentResponse) GetAmount() int {
	if p == nil {
		return 0
	}
	return p.Amount
}

func (p *PaymentIntentResponse) GetApplicationFeeAmount() *int {
	if p == nil {
		return nil
	}
	return p.ApplicationFeeAmount
}

func (p *PaymentIntentResponse) GetBatchPaymentId() *string {
	if p == nil {
		return nil
	}
	return p.BatchPaymentId
}

func (p *PaymentIntentResponse) GetCurrency() string {
	if p == nil {
		return ""
	}
	return p.Currency
}

func (p *PaymentIntentResponse) GetInvoice() *Invoice {
	if p == nil {
		return nil
	}
	return p.Invoice
}

func (p *PaymentIntentResponse) GetObject() *PaymentObject {
	if p == nil {
		return nil
	}
	return p.Object
}

func (p *PaymentIntentResponse) GetPayer() *AccountResponse {
	if p == nil {
		return nil
	}
	return p.Payer
}

func (p *PaymentIntentResponse) GetPaymentLinkId() *string {
	if p == nil {
		return nil
	}
	return p.PaymentLinkId
}

func (p *PaymentIntentResponse) GetPaymentMethods() []MoniteAllPaymentMethodsTypes {
	if p == nil {
		return nil
	}
	return p.PaymentMethods
}

func (p *PaymentIntentResponse) GetPaymentReference() *string {
	if p == nil {
		return nil
	}
	return p.PaymentReference
}

func (p *PaymentIntentResponse) GetProvider() *string {
	if p == nil {
		return nil
	}
	return p.Provider
}

func (p *PaymentIntentResponse) GetRecipient() *RecipientAccountResponse {
	if p == nil {
		return nil
	}
	return p.Recipient
}

func (p *PaymentIntentResponse) GetSelectedPaymentMethod() *MoniteAllPaymentMethodsTypes {
	if p == nil {
		return nil
	}
	return p.SelectedPaymentMethod
}

func (p *PaymentIntentResponse) GetStatus() string {
	if p == nil {
		return ""
	}
	return p.Status
}

func (p *PaymentIntentResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentIntentResponse) UnmarshalJSON(data []byte) error {
	type embed PaymentIntentResponse
	var unmarshaler = struct {
		embed
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PaymentIntentResponse(unmarshaler.embed)
	p.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentIntentResponse) MarshalJSON() ([]byte, error) {
	type embed PaymentIntentResponse
	var marshaler = struct {
		embed
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*p),
		UpdatedAt: internal.NewDateTime(p.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (p *PaymentIntentResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PaymentIntentsListResponse struct {
	Data []*PaymentIntentResponse `json:"data" url:"data"`
	// A token that can be sent in the `pagination_token` query parameter to get the next page of results, or `null` if there is no next page (i.e. you've reached the last page).
	NextPaginationToken *string `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`
	// A token that can be sent in the `pagination_token` query parameter to get the previous page of results, or `null` if there is no previous page (i.e. you've reached the first page).
	PrevPaginationToken *string `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentIntentsListResponse) GetData() []*PaymentIntentResponse {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PaymentIntentsListResponse) GetNextPaginationToken() *string {
	if p == nil {
		return nil
	}
	return p.NextPaginationToken
}

func (p *PaymentIntentsListResponse) GetPrevPaginationToken() *string {
	if p == nil {
		return nil
	}
	return p.PrevPaginationToken
}

func (p *PaymentIntentsListResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentIntentsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PaymentIntentsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaymentIntentsListResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentIntentsListResponse) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type UpdatePaymentIntentPayload struct {
	Amount int `json:"amount" url:"-"`
}
