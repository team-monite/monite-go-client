// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type ApprovalPolicyCreate struct {
	// The date and time (in the ISO 8601 format) when the approval policy becomes active. Only payables submitted for approval during the policy's active period will trigger this policy. If omitted or `null`, the policy is effective immediately. The value will be converted to UTC.
	StartsAt *time.Time `json:"starts_at,omitempty" url:"-"`
	// The date and time (in the ISO 8601 format) when the approval policy stops being active and stops triggering approval workflows.If `ends_at` is provided in the request, then `starts_at` must also be provided and `ends_at` must be later than `starts_at`. The value will be converted to UTC.
	EndsAt *time.Time `json:"ends_at,omitempty" url:"-"`
	// The name of the approval policy.
	Name string `json:"name" url:"-"`
	// A brief description of the approval policy.
	Description string `json:"description" url:"-"`
	// A list of JSON objects that represents the approval policy script. The script contains the logic that determines whether an action should be sent to approval. This field is required, and it should contain at least one script object.
	Script []*ApprovalPolicyCreateScriptItem `json:"script,omitempty" url:"-"`
	// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
	Trigger *ApprovalPolicyCreateTrigger `json:"trigger,omitempty" url:"-"`
}

func (a *ApprovalPolicyCreate) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalPolicyCreate
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*a = ApprovalPolicyCreate(body)
	return nil
}

func (a *ApprovalPolicyCreate) MarshalJSON() ([]byte, error) {
	type embed ApprovalPolicyCreate
	var marshaler = struct {
		embed
		StartsAt *internal.DateTime `json:"starts_at,omitempty"`
		EndsAt   *internal.DateTime `json:"ends_at,omitempty"`
	}{
		embed:    embed(*a),
		StartsAt: internal.NewOptionalDateTime(a.StartsAt),
		EndsAt:   internal.NewOptionalDateTime(a.EndsAt),
	}
	return json.Marshal(marshaler)
}

type ApprovalPoliciesGetRequest struct {
	ProcessId *string `json:"-" url:"process_id,omitempty"`
	// Order by
	Order *OrderEnum `json:"-" url:"order,omitempty"`
	// Max is 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// A token, obtained from previous page. Prior over other filters
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// Allowed sort fields
	Sort          *ApprovalPolicyCursorFields               `json:"-" url:"sort,omitempty"`
	IdIn          []*string                                 `json:"-" url:"id__in,omitempty"`
	Status        *ApprovalPoliciesGetRequestStatus         `json:"-" url:"status,omitempty"`
	StatusIn      []*ApprovalPoliciesGetRequestStatusInItem `json:"-" url:"status__in,omitempty"`
	Name          *string                                   `json:"-" url:"name,omitempty"`
	NameContains  *string                                   `json:"-" url:"name__contains,omitempty"`
	NameNcontains *string                                   `json:"-" url:"name__ncontains,omitempty"`
	CreatedBy     *string                                   `json:"-" url:"created_by,omitempty"`
	CreatedAtGt   *time.Time                                `json:"-" url:"created_at__gt,omitempty"`
	CreatedAtLt   *time.Time                                `json:"-" url:"created_at__lt,omitempty"`
	CreatedAtGte  *time.Time                                `json:"-" url:"created_at__gte,omitempty"`
	CreatedAtLte  *time.Time                                `json:"-" url:"created_at__lte,omitempty"`
	UpdatedAtGt   *time.Time                                `json:"-" url:"updated_at__gt,omitempty"`
	UpdatedAtLt   *time.Time                                `json:"-" url:"updated_at__lt,omitempty"`
	UpdatedAtGte  *time.Time                                `json:"-" url:"updated_at__gte,omitempty"`
	UpdatedAtLte  *time.Time                                `json:"-" url:"updated_at__lte,omitempty"`
}

type ApprovalPolicyCursorFields string

const (
	ApprovalPolicyCursorFieldsCreatedAt ApprovalPolicyCursorFields = "created_at"
	ApprovalPolicyCursorFieldsUpdatedAt ApprovalPolicyCursorFields = "updated_at"
)

func NewApprovalPolicyCursorFieldsFromString(s string) (ApprovalPolicyCursorFields, error) {
	switch s {
	case "created_at":
		return ApprovalPolicyCursorFieldsCreatedAt, nil
	case "updated_at":
		return ApprovalPolicyCursorFieldsUpdatedAt, nil
	}
	var t ApprovalPolicyCursorFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalPolicyCursorFields) Ptr() *ApprovalPolicyCursorFields {
	return &a
}

type ApprovalPolicyResource struct {
	// The date and time (in the ISO 8601 format) when the approval policy becomes active. Only payables submitted for approval during the policy's active period will trigger this policy. If omitted or `null`, the policy is effective immediately. The value will be converted to UTC.
	StartsAt *time.Time `json:"starts_at,omitempty" url:"starts_at,omitempty"`
	// The date and time (in the ISO 8601 format) when the approval policy stops being active and stops triggering approval workflows.If `ends_at` is provided in the request, then `starts_at` must also be provided and `ends_at` must be later than `starts_at`. The value will be converted to UTC.
	EndsAt *time.Time `json:"ends_at,omitempty" url:"ends_at,omitempty"`
	// The name of the approval policy.
	Name string `json:"name" url:"name"`
	// A brief description of the approval policy.
	Description string `json:"description" url:"description"`
	// A list of JSON objects that represents the approval policy script. The script contains the logic that determines whether an action should be sent to approval. This field is required, and it should contain at least one script object.
	Script []*ApprovalPolicyResourceScriptItem `json:"script" url:"script"`
	// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
	Trigger *ApprovalPolicyResourceTrigger `json:"trigger,omitempty" url:"trigger,omitempty"`
	Id      string                         `json:"id" url:"id"`
	// The current status of the approval policy.
	Status    ApprovalPolicyResourceStatus `json:"status" url:"status"`
	CreatedAt time.Time                    `json:"created_at" url:"created_at"`
	UpdatedAt time.Time                    `json:"updated_at" url:"updated_at"`
	CreatedBy string                       `json:"created_by" url:"created_by"`
	UpdatedBy *string                      `json:"updated_by,omitempty" url:"updated_by,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalPolicyResource) GetStartsAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.StartsAt
}

func (a *ApprovalPolicyResource) GetEndsAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.EndsAt
}

func (a *ApprovalPolicyResource) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *ApprovalPolicyResource) GetDescription() string {
	if a == nil {
		return ""
	}
	return a.Description
}

func (a *ApprovalPolicyResource) GetScript() []*ApprovalPolicyResourceScriptItem {
	if a == nil {
		return nil
	}
	return a.Script
}

func (a *ApprovalPolicyResource) GetTrigger() *ApprovalPolicyResourceTrigger {
	if a == nil {
		return nil
	}
	return a.Trigger
}

func (a *ApprovalPolicyResource) GetId() string {
	if a == nil {
		return ""
	}
	return a.Id
}

func (a *ApprovalPolicyResource) GetStatus() ApprovalPolicyResourceStatus {
	if a == nil {
		return ""
	}
	return a.Status
}

func (a *ApprovalPolicyResource) GetCreatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.CreatedAt
}

func (a *ApprovalPolicyResource) GetUpdatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.UpdatedAt
}

func (a *ApprovalPolicyResource) GetCreatedBy() string {
	if a == nil {
		return ""
	}
	return a.CreatedBy
}

func (a *ApprovalPolicyResource) GetUpdatedBy() *string {
	if a == nil {
		return nil
	}
	return a.UpdatedBy
}

func (a *ApprovalPolicyResource) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalPolicyResource) UnmarshalJSON(data []byte) error {
	type embed ApprovalPolicyResource
	var unmarshaler = struct {
		embed
		StartsAt  *internal.DateTime `json:"starts_at,omitempty"`
		EndsAt    *internal.DateTime `json:"ends_at,omitempty"`
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = ApprovalPolicyResource(unmarshaler.embed)
	a.StartsAt = unmarshaler.StartsAt.TimePtr()
	a.EndsAt = unmarshaler.EndsAt.TimePtr()
	a.CreatedAt = unmarshaler.CreatedAt.Time()
	a.UpdatedAt = unmarshaler.UpdatedAt.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalPolicyResource) MarshalJSON() ([]byte, error) {
	type embed ApprovalPolicyResource
	var marshaler = struct {
		embed
		StartsAt  *internal.DateTime `json:"starts_at,omitempty"`
		EndsAt    *internal.DateTime `json:"ends_at,omitempty"`
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*a),
		StartsAt:  internal.NewOptionalDateTime(a.StartsAt),
		EndsAt:    internal.NewOptionalDateTime(a.EndsAt),
		CreatedAt: internal.NewDateTime(a.CreatedAt),
		UpdatedAt: internal.NewDateTime(a.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (a *ApprovalPolicyResource) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApprovalPolicyResourceList struct {
	Data []*ApprovalPolicyResource `json:"data" url:"data"`
	// A token that can be sent in the `pagination_token` query parameter to get the previous page of results, or `null` if there is no previous page (i.e. you've reached the first page).
	PrevPaginationToken *string `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`
	// A token that can be sent in the `pagination_token` query parameter to get the next page of results, or `null` if there is no next page (i.e. you've reached the last page).
	NextPaginationToken *string `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalPolicyResourceList) GetData() []*ApprovalPolicyResource {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApprovalPolicyResourceList) GetPrevPaginationToken() *string {
	if a == nil {
		return nil
	}
	return a.PrevPaginationToken
}

func (a *ApprovalPolicyResourceList) GetNextPaginationToken() *string {
	if a == nil {
		return nil
	}
	return a.NextPaginationToken
}

func (a *ApprovalPolicyResourceList) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalPolicyResourceList) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalPolicyResourceList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalPolicyResourceList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalPolicyResourceList) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApprovalPolicyResourceScriptItem struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyResourceScriptItem) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyResourceScriptItem) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyResourceScriptItem) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyResourceScriptItem) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyResourceScriptItem) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyResourceScriptItem) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyResourceScriptItem) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyResourceScriptItemVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyResourceScriptItem) Accept(visitor ApprovalPolicyResourceScriptItemVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

// The current status of the approval policy.
type ApprovalPolicyResourceStatus string

const (
	ApprovalPolicyResourceStatusActive  ApprovalPolicyResourceStatus = "active"
	ApprovalPolicyResourceStatusPending ApprovalPolicyResourceStatus = "pending"
)

func NewApprovalPolicyResourceStatusFromString(s string) (ApprovalPolicyResourceStatus, error) {
	switch s {
	case "active":
		return ApprovalPolicyResourceStatusActive, nil
	case "pending":
		return ApprovalPolicyResourceStatusPending, nil
	}
	var t ApprovalPolicyResourceStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalPolicyResourceStatus) Ptr() *ApprovalPolicyResourceStatus {
	return &a
}

// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
type ApprovalPolicyResourceTrigger struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyResourceTrigger) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyResourceTrigger) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyResourceTrigger) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyResourceTrigger) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyResourceTrigger) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyResourceTrigger) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyResourceTrigger) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyResourceTriggerVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyResourceTrigger) Accept(visitor ApprovalPolicyResourceTriggerVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyStatus string

const (
	ApprovalPolicyStatusActive  ApprovalPolicyStatus = "active"
	ApprovalPolicyStatusDeleted ApprovalPolicyStatus = "deleted"
	ApprovalPolicyStatusPending ApprovalPolicyStatus = "pending"
)

func NewApprovalPolicyStatusFromString(s string) (ApprovalPolicyStatus, error) {
	switch s {
	case "active":
		return ApprovalPolicyStatusActive, nil
	case "deleted":
		return ApprovalPolicyStatusDeleted, nil
	case "pending":
		return ApprovalPolicyStatusPending, nil
	}
	var t ApprovalPolicyStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalPolicyStatus) Ptr() *ApprovalPolicyStatus {
	return &a
}

type ApprovalPoliciesGetRequestStatus string

const (
	ApprovalPoliciesGetRequestStatusActive  ApprovalPoliciesGetRequestStatus = "active"
	ApprovalPoliciesGetRequestStatusPending ApprovalPoliciesGetRequestStatus = "pending"
)

func NewApprovalPoliciesGetRequestStatusFromString(s string) (ApprovalPoliciesGetRequestStatus, error) {
	switch s {
	case "active":
		return ApprovalPoliciesGetRequestStatusActive, nil
	case "pending":
		return ApprovalPoliciesGetRequestStatusPending, nil
	}
	var t ApprovalPoliciesGetRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalPoliciesGetRequestStatus) Ptr() *ApprovalPoliciesGetRequestStatus {
	return &a
}

type ApprovalPoliciesGetRequestStatusInItem string

const (
	ApprovalPoliciesGetRequestStatusInItemActive  ApprovalPoliciesGetRequestStatusInItem = "active"
	ApprovalPoliciesGetRequestStatusInItemPending ApprovalPoliciesGetRequestStatusInItem = "pending"
)

func NewApprovalPoliciesGetRequestStatusInItemFromString(s string) (ApprovalPoliciesGetRequestStatusInItem, error) {
	switch s {
	case "active":
		return ApprovalPoliciesGetRequestStatusInItemActive, nil
	case "pending":
		return ApprovalPoliciesGetRequestStatusInItemPending, nil
	}
	var t ApprovalPoliciesGetRequestStatusInItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalPoliciesGetRequestStatusInItem) Ptr() *ApprovalPoliciesGetRequestStatusInItem {
	return &a
}

type ApprovalPolicyCreateScriptItem struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyCreateScriptItem) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyCreateScriptItem) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyCreateScriptItem) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyCreateScriptItem) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyCreateScriptItem) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyCreateScriptItem) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyCreateScriptItem) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyCreateScriptItemVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyCreateScriptItem) Accept(visitor ApprovalPolicyCreateScriptItemVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
type ApprovalPolicyCreateTrigger struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyCreateTrigger) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyCreateTrigger) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyCreateTrigger) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyCreateTrigger) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyCreateTrigger) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyCreateTrigger) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyCreateTrigger) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyCreateTriggerVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyCreateTrigger) Accept(visitor ApprovalPolicyCreateTriggerVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyUpdateScriptItem struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyUpdateScriptItem) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyUpdateScriptItem) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyUpdateScriptItem) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyUpdateScriptItem) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyUpdateScriptItem) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyUpdateScriptItem) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyUpdateScriptItem) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyUpdateScriptItemVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyUpdateScriptItem) Accept(visitor ApprovalPolicyUpdateScriptItemVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
type ApprovalPolicyUpdateTrigger struct {
	Boolean          bool
	Double           float64
	String           string
	UnknownList      []interface{}
	StringUnknownMap map[string]interface{}

	typ string
}

func (a *ApprovalPolicyUpdateTrigger) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *ApprovalPolicyUpdateTrigger) GetDouble() float64 {
	if a == nil {
		return 0
	}
	return a.Double
}

func (a *ApprovalPolicyUpdateTrigger) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *ApprovalPolicyUpdateTrigger) GetUnknownList() []interface{} {
	if a == nil {
		return nil
	}
	return a.UnknownList
}

func (a *ApprovalPolicyUpdateTrigger) GetStringUnknownMap() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.StringUnknownMap
}

func (a *ApprovalPolicyUpdateTrigger) UnmarshalJSON(data []byte) error {
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		a.typ = "Double"
		a.Double = valueDouble
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		a.typ = "UnknownList"
		a.UnknownList = valueUnknownList
		return nil
	}
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		a.typ = "StringUnknownMap"
		a.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalPolicyUpdateTrigger) MarshalJSON() ([]byte, error) {
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return json.Marshal(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return json.Marshal(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return json.Marshal(a.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyUpdateTriggerVisitor interface {
	VisitBoolean(bool) error
	VisitDouble(float64) error
	VisitString(string) error
	VisitUnknownList([]interface{}) error
	VisitStringUnknownMap(map[string]interface{}) error
}

func (a *ApprovalPolicyUpdateTrigger) Accept(visitor ApprovalPolicyUpdateTriggerVisitor) error {
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	if a.typ == "Double" || a.Double != 0 {
		return visitor.VisitDouble(a.Double)
	}
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "UnknownList" || a.UnknownList != nil {
		return visitor.VisitUnknownList(a.UnknownList)
	}
	if a.typ == "StringUnknownMap" || a.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(a.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalPolicyUpdate struct {
	// The date and time (in the ISO 8601 format) when the approval policy becomes active. Only payables submitted for approval during the policy's active period will trigger this policy. If omitted or `null`, the policy is effective immediately. The value will be converted to UTC.
	StartsAt *time.Time `json:"starts_at,omitempty" url:"-"`
	// The date and time (in the ISO 8601 format) when the approval policy stops being active and stops triggering approval workflows.If `ends_at` is provided in the request, then `starts_at` must also be provided and `ends_at` must be later than `starts_at`. The value will be converted to UTC.
	EndsAt *time.Time `json:"ends_at,omitempty" url:"-"`
	// The name of the approval policy.
	Name *string `json:"name,omitempty" url:"-"`
	// A brief description of the approval policy.
	Description *string `json:"description,omitempty" url:"-"`
	// A list of JSON objects that represents the approval policy script. The script contains the logic that determines whether an action should be sent to approval. This field is required, and it should contain at least one script object.
	Script []*ApprovalPolicyUpdateScriptItem `json:"script,omitempty" url:"-"`
	// A JSON object that represents the trigger for the approval policy. The trigger specifies the event that will trigger the policy to be evaluated.
	Trigger *ApprovalPolicyUpdateTrigger `json:"trigger,omitempty" url:"-"`
	// A string that represents the current status of the approval policy.
	Status *ApprovalPolicyStatus `json:"status,omitempty" url:"-"`
}

func (a *ApprovalPolicyUpdate) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalPolicyUpdate
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*a = ApprovalPolicyUpdate(body)
	return nil
}

func (a *ApprovalPolicyUpdate) MarshalJSON() ([]byte, error) {
	type embed ApprovalPolicyUpdate
	var marshaler = struct {
		embed
		StartsAt *internal.DateTime `json:"starts_at,omitempty"`
		EndsAt   *internal.DateTime `json:"ends_at,omitempty"`
	}{
		embed:    embed(*a),
		StartsAt: internal.NewOptionalDateTime(a.StartsAt),
		EndsAt:   internal.NewOptionalDateTime(a.EndsAt),
	}
	return json.Marshal(marshaler)
}
