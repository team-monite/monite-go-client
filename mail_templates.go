// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type AddCustomTemplateSchema struct {
	// Jinja2 compatible string with email body
	BodyTemplate string `json:"body_template" url:"-"`
	// Is default template
	IsDefault *bool `json:"is_default,omitempty" url:"-"`
	// Lowercase ISO code of language
	Language *LanguageCodeEnum `json:"language,omitempty" url:"-"`
	// Custom template name
	Name string `json:"name" url:"-"`
	// Jinja2 compatible string with email subject
	SubjectTemplate string `json:"subject_template" url:"-"`
	// Document type of content
	Type DocumentObjectTypeRequestEnum `json:"type" url:"-"`
}

type MailTemplatesGetRequest struct {
	// Order by
	Order *OrderEnum `json:"-" url:"order,omitempty"`
	// Max is 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// A token, obtained from previous page. Prior over other filters
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// Allowed sort fields
	Sort          *CustomTemplatesCursorFields     `json:"-" url:"sort,omitempty"`
	Type          *DocumentObjectTypeRequestEnum   `json:"-" url:"type,omitempty"`
	TypeIn        []*DocumentObjectTypeRequestEnum `json:"-" url:"type__in,omitempty"`
	TypeNotIn     []*DocumentObjectTypeRequestEnum `json:"-" url:"type__not_in,omitempty"`
	IsDefault     *bool                            `json:"-" url:"is_default,omitempty"`
	Name          *string                          `json:"-" url:"name,omitempty"`
	NameIexact    *string                          `json:"-" url:"name__iexact,omitempty"`
	NameContains  *string                          `json:"-" url:"name__contains,omitempty"`
	NameIcontains *string                          `json:"-" url:"name__icontains,omitempty"`
}

type PreviewTemplateRequest struct {
	// Body text of the template
	Body string `json:"body" url:"-"`
	// Document type of content
	DocumentType DocumentObjectTypeRequestEnum `json:"document_type" url:"-"`
	// Lowercase ISO code of language
	LanguageCode LanguageCodeEnum `json:"language_code" url:"-"`
	// Subject text of the template
	Subject string `json:"subject" url:"-"`
}

type CustomTemplateDataSchema struct {
	// ID of email template
	Id string `json:"id" url:"id"`
	// Template created date and time
	CreatedAt time.Time `json:"created_at" url:"created_at"`
	// Template updated date and time
	UpdatedAt time.Time `json:"updated_at" url:"updated_at"`
	// Jinja2 compatible email body template
	BodyTemplate string `json:"body_template" url:"body_template"`
	// Is default template
	IsDefault bool `json:"is_default" url:"is_default"`
	// Lowercase ISO code of language
	Language string `json:"language" url:"language"`
	// Name of the template
	Name string `json:"name" url:"name"`
	// Jinja2 compatible email subject template
	SubjectTemplate string `json:"subject_template" url:"subject_template"`
	// Document type of content
	Type string `json:"type" url:"type"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CustomTemplateDataSchema) GetId() string {
	if c == nil {
		return ""
	}
	return c.Id
}

func (c *CustomTemplateDataSchema) GetCreatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.CreatedAt
}

func (c *CustomTemplateDataSchema) GetUpdatedAt() time.Time {
	if c == nil {
		return time.Time{}
	}
	return c.UpdatedAt
}

func (c *CustomTemplateDataSchema) GetBodyTemplate() string {
	if c == nil {
		return ""
	}
	return c.BodyTemplate
}

func (c *CustomTemplateDataSchema) GetIsDefault() bool {
	if c == nil {
		return false
	}
	return c.IsDefault
}

func (c *CustomTemplateDataSchema) GetLanguage() string {
	if c == nil {
		return ""
	}
	return c.Language
}

func (c *CustomTemplateDataSchema) GetName() string {
	if c == nil {
		return ""
	}
	return c.Name
}

func (c *CustomTemplateDataSchema) GetSubjectTemplate() string {
	if c == nil {
		return ""
	}
	return c.SubjectTemplate
}

func (c *CustomTemplateDataSchema) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *CustomTemplateDataSchema) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CustomTemplateDataSchema) UnmarshalJSON(data []byte) error {
	type embed CustomTemplateDataSchema
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*c),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*c = CustomTemplateDataSchema(unmarshaler.embed)
	c.CreatedAt = unmarshaler.CreatedAt.Time()
	c.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CustomTemplateDataSchema) MarshalJSON() ([]byte, error) {
	type embed CustomTemplateDataSchema
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*c),
		CreatedAt: internal.NewDateTime(c.CreatedAt),
		UpdatedAt: internal.NewDateTime(c.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (c *CustomTemplateDataSchema) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type CustomTemplatesCursorFields string

const (
	CustomTemplatesCursorFieldsType CustomTemplatesCursorFields = "type"
	CustomTemplatesCursorFieldsName CustomTemplatesCursorFields = "name"
)

func NewCustomTemplatesCursorFieldsFromString(s string) (CustomTemplatesCursorFields, error) {
	switch s {
	case "type":
		return CustomTemplatesCursorFieldsType, nil
	case "name":
		return CustomTemplatesCursorFieldsName, nil
	}
	var t CustomTemplatesCursorFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c CustomTemplatesCursorFields) Ptr() *CustomTemplatesCursorFields {
	return &c
}

type CustomTemplatesPaginationResponse struct {
	// All user-defined email templates
	Data []*CustomTemplateDataSchema `json:"data" url:"data"`
	// A token that can be sent in the `pagination_token` query parameter to get the next page of results, or `null` if there is no next page (i.e. you've reached the last page).
	NextPaginationToken *string `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`
	// A token that can be sent in the `pagination_token` query parameter to get the previous page of results, or `null` if there is no previous page (i.e. you've reached the first page).
	PrevPaginationToken *string `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (c *CustomTemplatesPaginationResponse) GetData() []*CustomTemplateDataSchema {
	if c == nil {
		return nil
	}
	return c.Data
}

func (c *CustomTemplatesPaginationResponse) GetNextPaginationToken() *string {
	if c == nil {
		return nil
	}
	return c.NextPaginationToken
}

func (c *CustomTemplatesPaginationResponse) GetPrevPaginationToken() *string {
	if c == nil {
		return nil
	}
	return c.PrevPaginationToken
}

func (c *CustomTemplatesPaginationResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CustomTemplatesPaginationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CustomTemplatesPaginationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CustomTemplatesPaginationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	c.rawJSON = json.RawMessage(data)
	return nil
}

func (c *CustomTemplatesPaginationResponse) String() string {
	if len(c.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(c.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type DocumentObjectTypeRequestEnum string

const (
	DocumentObjectTypeRequestEnumReceivablesQuote            DocumentObjectTypeRequestEnum = "receivables_quote"
	DocumentObjectTypeRequestEnumReceivablesInvoice          DocumentObjectTypeRequestEnum = "receivables_invoice"
	DocumentObjectTypeRequestEnumReceivablesPaidInvoice      DocumentObjectTypeRequestEnum = "receivables_paid_invoice"
	DocumentObjectTypeRequestEnumReceivablesCreditNote       DocumentObjectTypeRequestEnum = "receivables_credit_note"
	DocumentObjectTypeRequestEnumReceivablesDiscountReminder DocumentObjectTypeRequestEnum = "receivables_discount_reminder"
	DocumentObjectTypeRequestEnumReceivablesFinalReminder    DocumentObjectTypeRequestEnum = "receivables_final_reminder"
	DocumentObjectTypeRequestEnumPayablesPurchaseOrder       DocumentObjectTypeRequestEnum = "payables_purchase_order"
	DocumentObjectTypeRequestEnumPayablesNotifyApprover      DocumentObjectTypeRequestEnum = "payables_notify_approver"
	DocumentObjectTypeRequestEnumPayablesNotifyPayer         DocumentObjectTypeRequestEnum = "payables_notify_payer"
)

func NewDocumentObjectTypeRequestEnumFromString(s string) (DocumentObjectTypeRequestEnum, error) {
	switch s {
	case "receivables_quote":
		return DocumentObjectTypeRequestEnumReceivablesQuote, nil
	case "receivables_invoice":
		return DocumentObjectTypeRequestEnumReceivablesInvoice, nil
	case "receivables_paid_invoice":
		return DocumentObjectTypeRequestEnumReceivablesPaidInvoice, nil
	case "receivables_credit_note":
		return DocumentObjectTypeRequestEnumReceivablesCreditNote, nil
	case "receivables_discount_reminder":
		return DocumentObjectTypeRequestEnumReceivablesDiscountReminder, nil
	case "receivables_final_reminder":
		return DocumentObjectTypeRequestEnumReceivablesFinalReminder, nil
	case "payables_purchase_order":
		return DocumentObjectTypeRequestEnumPayablesPurchaseOrder, nil
	case "payables_notify_approver":
		return DocumentObjectTypeRequestEnumPayablesNotifyApprover, nil
	case "payables_notify_payer":
		return DocumentObjectTypeRequestEnumPayablesNotifyPayer, nil
	}
	var t DocumentObjectTypeRequestEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (d DocumentObjectTypeRequestEnum) Ptr() *DocumentObjectTypeRequestEnum {
	return &d
}

type PreviewTemplateResponse struct {
	BodyPreview    string `json:"body_preview" url:"body_preview"`
	SubjectPreview string `json:"subject_preview" url:"subject_preview"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PreviewTemplateResponse) GetBodyPreview() string {
	if p == nil {
		return ""
	}
	return p.BodyPreview
}

func (p *PreviewTemplateResponse) GetSubjectPreview() string {
	if p == nil {
		return ""
	}
	return p.SubjectPreview
}

func (p *PreviewTemplateResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PreviewTemplateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PreviewTemplateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PreviewTemplateResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PreviewTemplateResponse) String() string {
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

type SystemTemplateDataSchema struct {
	// Array of templates
	AvailableTemplates []*TemplateDataSchema `json:"available_templates" url:"available_templates"`
	// Name of the template
	TemplateName string `json:"template_name" url:"template_name"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SystemTemplateDataSchema) GetAvailableTemplates() []*TemplateDataSchema {
	if s == nil {
		return nil
	}
	return s.AvailableTemplates
}

func (s *SystemTemplateDataSchema) GetTemplateName() string {
	if s == nil {
		return ""
	}
	return s.TemplateName
}

func (s *SystemTemplateDataSchema) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SystemTemplateDataSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler SystemTemplateDataSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SystemTemplateDataSchema(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SystemTemplateDataSchema) String() string {
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

type SystemTemplates struct {
	// All pre-defined email templates
	Data []*SystemTemplateDataSchema `json:"data" url:"data"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (s *SystemTemplates) GetData() []*SystemTemplateDataSchema {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SystemTemplates) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SystemTemplates) UnmarshalJSON(data []byte) error {
	type unmarshaler SystemTemplates
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SystemTemplates(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	s.rawJSON = json.RawMessage(data)
	return nil
}

func (s *SystemTemplates) String() string {
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

type TemplateDataSchema struct {
	// Jinja2 compatible email body template
	BodyTemplate string `json:"body_template" url:"body_template"`
	// Lowercase ISO code of language
	Language string `json:"language" url:"language"`
	// Jinja2 compatible email subject template
	SubjectTemplate string `json:"subject_template" url:"subject_template"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TemplateDataSchema) GetBodyTemplate() string {
	if t == nil {
		return ""
	}
	return t.BodyTemplate
}

func (t *TemplateDataSchema) GetLanguage() string {
	if t == nil {
		return ""
	}
	return t.Language
}

func (t *TemplateDataSchema) GetSubjectTemplate() string {
	if t == nil {
		return ""
	}
	return t.SubjectTemplate
}

func (t *TemplateDataSchema) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TemplateDataSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler TemplateDataSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TemplateDataSchema(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TemplateDataSchema) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type UpdateCustomTemplateSchemaRequest struct {
	// Jinja2 compatible string with email body
	BodyTemplate *string `json:"body_template,omitempty" url:"-"`
	// Lowercase ISO code of language
	Language *LanguageCodeEnum `json:"language,omitempty" url:"-"`
	// Custom template name
	Name *string `json:"name,omitempty" url:"-"`
	// Jinja2 compatible string with email subject
	SubjectTemplate *string `json:"subject_template,omitempty" url:"-"`
}
