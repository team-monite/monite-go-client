// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type PaymentsBatchPaymentRequest struct {
	PayerBankAccountId string                 `json:"payer_bank_account_id" url:"-"`
	PaymentIntents     []*SinglePaymentIntent `json:"payment_intents,omitempty" url:"-"`
	paymentMethod      string
}

func (p *PaymentsBatchPaymentRequest) PaymentMethod() string {
	return p.paymentMethod
}

func (p *PaymentsBatchPaymentRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler PaymentsBatchPaymentRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*p = PaymentsBatchPaymentRequest(body)
	p.paymentMethod = "us_ach"
	return nil
}

func (p *PaymentsBatchPaymentRequest) MarshalJSON() ([]byte, error) {
	type embed PaymentsBatchPaymentRequest
	var marshaler = struct {
		embed
		PaymentMethod string `json:"payment_method"`
	}{
		embed:         embed(*p),
		PaymentMethod: "us_ach",
	}
	return json.Marshal(marshaler)
}

type PaymentIntentPayoutMethod string

const (
	PaymentIntentPayoutMethodBankAccount PaymentIntentPayoutMethod = "bank_account"
	PaymentIntentPayoutMethodPaperCheck  PaymentIntentPayoutMethod = "paper_check"
)

func NewPaymentIntentPayoutMethodFromString(s string) (PaymentIntentPayoutMethod, error) {
	switch s {
	case "bank_account":
		return PaymentIntentPayoutMethodBankAccount, nil
	case "paper_check":
		return PaymentIntentPayoutMethodPaperCheck, nil
	}
	var t PaymentIntentPayoutMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PaymentIntentPayoutMethod) Ptr() *PaymentIntentPayoutMethod {
	return &p
}

type PaymentIntentsRecipient struct {
	Id            string                     `json:"id" url:"id"`
	BankAccountId *string                    `json:"bank_account_id,omitempty" url:"bank_account_id,omitempty"`
	PayoutMethod  *PaymentIntentPayoutMethod `json:"payout_method,omitempty" url:"payout_method,omitempty"`
	type_         string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentIntentsRecipient) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *PaymentIntentsRecipient) GetBankAccountId() *string {
	if p == nil {
		return nil
	}
	return p.BankAccountId
}

func (p *PaymentIntentsRecipient) GetPayoutMethod() *PaymentIntentPayoutMethod {
	if p == nil {
		return nil
	}
	return p.PayoutMethod
}

func (p *PaymentIntentsRecipient) Type() string {
	return p.type_
}

func (p *PaymentIntentsRecipient) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentIntentsRecipient) UnmarshalJSON(data []byte) error {
	type embed PaymentIntentsRecipient
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PaymentIntentsRecipient(unmarshaler.embed)
	if unmarshaler.Type != "counterpart" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "counterpart", unmarshaler.Type)
	}
	p.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentIntentsRecipient) MarshalJSON() ([]byte, error) {
	type embed PaymentIntentsRecipient
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "counterpart",
	}
	return json.Marshal(marshaler)
}

func (p *PaymentIntentsRecipient) String() string {
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

type PaymentObjectPayable struct {
	Id    string `json:"id" url:"id"`
	type_ string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentObjectPayable) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *PaymentObjectPayable) Type() string {
	return p.type_
}

func (p *PaymentObjectPayable) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentObjectPayable) UnmarshalJSON(data []byte) error {
	type embed PaymentObjectPayable
	var unmarshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PaymentObjectPayable(unmarshaler.embed)
	if unmarshaler.Type != "payable" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "payable", unmarshaler.Type)
	}
	p.type_ = unmarshaler.Type
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentObjectPayable) MarshalJSON() ([]byte, error) {
	type embed PaymentObjectPayable
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "payable",
	}
	return json.Marshal(marshaler)
}

func (p *PaymentObjectPayable) String() string {
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

type PaymentsBatchPaymentResponse struct {
	Id                 string                         `json:"id" url:"id"`
	CreatedAt          time.Time                      `json:"created_at" url:"created_at"`
	Error              map[string]interface{}         `json:"error,omitempty" url:"error,omitempty"`
	PayerBankAccountId string                         `json:"payer_bank_account_id" url:"payer_bank_account_id"`
	PaymentIntents     []*SinglePaymentIntentResponse `json:"payment_intents" url:"payment_intents"`
	Status             PaymentsBatchPaymentStatus     `json:"status" url:"status"`
	TotalAmount        *int                           `json:"total_amount,omitempty" url:"total_amount,omitempty"`
	paymentMethod      string

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaymentsBatchPaymentResponse) GetId() string {
	if p == nil {
		return ""
	}
	return p.Id
}

func (p *PaymentsBatchPaymentResponse) GetCreatedAt() time.Time {
	if p == nil {
		return time.Time{}
	}
	return p.CreatedAt
}

func (p *PaymentsBatchPaymentResponse) GetError() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.Error
}

func (p *PaymentsBatchPaymentResponse) GetPayerBankAccountId() string {
	if p == nil {
		return ""
	}
	return p.PayerBankAccountId
}

func (p *PaymentsBatchPaymentResponse) GetPaymentIntents() []*SinglePaymentIntentResponse {
	if p == nil {
		return nil
	}
	return p.PaymentIntents
}

func (p *PaymentsBatchPaymentResponse) GetStatus() PaymentsBatchPaymentStatus {
	if p == nil {
		return ""
	}
	return p.Status
}

func (p *PaymentsBatchPaymentResponse) GetTotalAmount() *int {
	if p == nil {
		return nil
	}
	return p.TotalAmount
}

func (p *PaymentsBatchPaymentResponse) PaymentMethod() string {
	return p.paymentMethod
}

func (p *PaymentsBatchPaymentResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaymentsBatchPaymentResponse) UnmarshalJSON(data []byte) error {
	type embed PaymentsBatchPaymentResponse
	var unmarshaler = struct {
		embed
		CreatedAt     *internal.DateTime `json:"created_at"`
		PaymentMethod string             `json:"payment_method"`
	}{
		embed: embed(*p),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*p = PaymentsBatchPaymentResponse(unmarshaler.embed)
	p.CreatedAt = unmarshaler.CreatedAt.Time()
	if unmarshaler.PaymentMethod != "us_ach" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "us_ach", unmarshaler.PaymentMethod)
	}
	p.paymentMethod = unmarshaler.PaymentMethod
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "payment_method")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaymentsBatchPaymentResponse) MarshalJSON() ([]byte, error) {
	type embed PaymentsBatchPaymentResponse
	var marshaler = struct {
		embed
		CreatedAt     *internal.DateTime `json:"created_at"`
		PaymentMethod string             `json:"payment_method"`
	}{
		embed:         embed(*p),
		CreatedAt:     internal.NewDateTime(p.CreatedAt),
		PaymentMethod: "us_ach",
	}
	return json.Marshal(marshaler)
}

func (p *PaymentsBatchPaymentResponse) String() string {
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

type PaymentsBatchPaymentStatus string

const (
	PaymentsBatchPaymentStatusCreated             PaymentsBatchPaymentStatus = "created"
	PaymentsBatchPaymentStatusProcessing          PaymentsBatchPaymentStatus = "processing"
	PaymentsBatchPaymentStatusPartiallySuccessful PaymentsBatchPaymentStatus = "partially_successful"
	PaymentsBatchPaymentStatusSucceeded           PaymentsBatchPaymentStatus = "succeeded"
	PaymentsBatchPaymentStatusFailed              PaymentsBatchPaymentStatus = "failed"
)

func NewPaymentsBatchPaymentStatusFromString(s string) (PaymentsBatchPaymentStatus, error) {
	switch s {
	case "created":
		return PaymentsBatchPaymentStatusCreated, nil
	case "processing":
		return PaymentsBatchPaymentStatusProcessing, nil
	case "partially_successful":
		return PaymentsBatchPaymentStatusPartiallySuccessful, nil
	case "succeeded":
		return PaymentsBatchPaymentStatusSucceeded, nil
	case "failed":
		return PaymentsBatchPaymentStatusFailed, nil
	}
	var t PaymentsBatchPaymentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p PaymentsBatchPaymentStatus) Ptr() *PaymentsBatchPaymentStatus {
	return &p
}

type SinglePaymentIntent struct {
	Object *PaymentObjectPayable `json:"object" url:"object"`
	// Must be provided if payable's document id is missing.
	PaymentReference *string                  `json:"payment_reference,omitempty" url:"payment_reference,omitempty"`
	Recipient        *PaymentIntentsRecipient `json:"recipient" url:"recipient"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SinglePaymentIntent) GetObject() *PaymentObjectPayable {
	if s == nil {
		return nil
	}
	return s.Object
}

func (s *SinglePaymentIntent) GetPaymentReference() *string {
	if s == nil {
		return nil
	}
	return s.PaymentReference
}

func (s *SinglePaymentIntent) GetRecipient() *PaymentIntentsRecipient {
	if s == nil {
		return nil
	}
	return s.Recipient
}

func (s *SinglePaymentIntent) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SinglePaymentIntent) UnmarshalJSON(data []byte) error {
	type unmarshaler SinglePaymentIntent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SinglePaymentIntent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SinglePaymentIntent) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type SinglePaymentIntentResponse struct {
	Id               string                   `json:"id" url:"id"`
	CreatedAt        time.Time                `json:"created_at" url:"created_at"`
	Amount           int                      `json:"amount" url:"amount"`
	Currency         CurrencyEnum             `json:"currency" url:"currency"`
	Error            map[string]interface{}   `json:"error,omitempty" url:"error,omitempty"`
	Object           *PaymentObjectPayable    `json:"object" url:"object"`
	PaymentReference string                   `json:"payment_reference" url:"payment_reference"`
	Recipient        *PaymentIntentsRecipient `json:"recipient" url:"recipient"`
	Status           string                   `json:"status" url:"status"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SinglePaymentIntentResponse) GetId() string {
	if s == nil {
		return ""
	}
	return s.Id
}

func (s *SinglePaymentIntentResponse) GetCreatedAt() time.Time {
	if s == nil {
		return time.Time{}
	}
	return s.CreatedAt
}

func (s *SinglePaymentIntentResponse) GetAmount() int {
	if s == nil {
		return 0
	}
	return s.Amount
}

func (s *SinglePaymentIntentResponse) GetCurrency() CurrencyEnum {
	if s == nil {
		return ""
	}
	return s.Currency
}

func (s *SinglePaymentIntentResponse) GetError() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.Error
}

func (s *SinglePaymentIntentResponse) GetObject() *PaymentObjectPayable {
	if s == nil {
		return nil
	}
	return s.Object
}

func (s *SinglePaymentIntentResponse) GetPaymentReference() string {
	if s == nil {
		return ""
	}
	return s.PaymentReference
}

func (s *SinglePaymentIntentResponse) GetRecipient() *PaymentIntentsRecipient {
	if s == nil {
		return nil
	}
	return s.Recipient
}

func (s *SinglePaymentIntentResponse) GetStatus() string {
	if s == nil {
		return ""
	}
	return s.Status
}

func (s *SinglePaymentIntentResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SinglePaymentIntentResponse) UnmarshalJSON(data []byte) error {
	type embed SinglePaymentIntentResponse
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SinglePaymentIntentResponse(unmarshaler.embed)
	s.CreatedAt = unmarshaler.CreatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SinglePaymentIntentResponse) MarshalJSON() ([]byte, error) {
	type embed SinglePaymentIntentResponse
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
	}{
		embed:     embed(*s),
		CreatedAt: internal.NewDateTime(s.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (s *SinglePaymentIntentResponse) String() string {
	if len(s.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(s.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
