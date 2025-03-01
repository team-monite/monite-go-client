// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type ApprovalRequestsGetRequest struct {
	// Order by
	Order *OrderEnum `json:"-" url:"order,omitempty"`
	// Max is 100
	Limit *int `json:"-" url:"limit,omitempty"`
	// A token, obtained from previous page. Prior over other filters
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// Allowed sort fields
	Sort         *ApprovalRequestCursorFields `json:"-" url:"sort,omitempty"`
	CreatedAtGt  *time.Time                   `json:"-" url:"created_at__gt,omitempty"`
	CreatedAtLt  *time.Time                   `json:"-" url:"created_at__lt,omitempty"`
	CreatedAtGte *time.Time                   `json:"-" url:"created_at__gte,omitempty"`
	CreatedAtLte *time.Time                   `json:"-" url:"created_at__lte,omitempty"`
	UpdatedAtGt  *time.Time                   `json:"-" url:"updated_at__gt,omitempty"`
	UpdatedAtLt  *time.Time                   `json:"-" url:"updated_at__lt,omitempty"`
	UpdatedAtGte *time.Time                   `json:"-" url:"updated_at__gte,omitempty"`
	UpdatedAtLte *time.Time                   `json:"-" url:"updated_at__lte,omitempty"`
	ObjectId     *string                      `json:"-" url:"object_id,omitempty"`
	ObjectIdIn   []*string                    `json:"-" url:"object_id__in,omitempty"`
	Status       *ApprovalRequestStatus       `json:"-" url:"status,omitempty"`
	StatusIn     []*ApprovalRequestStatus     `json:"-" url:"status__in,omitempty"`
	UserId       *string                      `json:"-" url:"user_id,omitempty"`
	RoleId       *string                      `json:"-" url:"role_id,omitempty"`
	ObjectType   *ObjectType                  `json:"-" url:"object_type,omitempty"`
	ObjectTypeIn []*ObjectType                `json:"-" url:"object_type__in,omitempty"`
	CreatedBy    *string                      `json:"-" url:"created_by,omitempty"`
}

type ApprovalRequestCreateByRoleRequest struct {
	ObjectId              string     `json:"object_id" url:"object_id"`
	ObjectType            ObjectType `json:"object_type" url:"object_type"`
	RequiredApprovalCount int        `json:"required_approval_count" url:"required_approval_count"`
	RoleIds               []string   `json:"role_ids" url:"role_ids"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalRequestCreateByRoleRequest) GetObjectId() string {
	if a == nil {
		return ""
	}
	return a.ObjectId
}

func (a *ApprovalRequestCreateByRoleRequest) GetObjectType() ObjectType {
	if a == nil {
		return ""
	}
	return a.ObjectType
}

func (a *ApprovalRequestCreateByRoleRequest) GetRequiredApprovalCount() int {
	if a == nil {
		return 0
	}
	return a.RequiredApprovalCount
}

func (a *ApprovalRequestCreateByRoleRequest) GetRoleIds() []string {
	if a == nil {
		return nil
	}
	return a.RoleIds
}

func (a *ApprovalRequestCreateByRoleRequest) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalRequestCreateByRoleRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalRequestCreateByRoleRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalRequestCreateByRoleRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalRequestCreateByRoleRequest) String() string {
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

type ApprovalRequestCreateByUserRequest struct {
	ObjectId              string     `json:"object_id" url:"object_id"`
	ObjectType            ObjectType `json:"object_type" url:"object_type"`
	RequiredApprovalCount int        `json:"required_approval_count" url:"required_approval_count"`
	UserIds               []string   `json:"user_ids" url:"user_ids"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalRequestCreateByUserRequest) GetObjectId() string {
	if a == nil {
		return ""
	}
	return a.ObjectId
}

func (a *ApprovalRequestCreateByUserRequest) GetObjectType() ObjectType {
	if a == nil {
		return ""
	}
	return a.ObjectType
}

func (a *ApprovalRequestCreateByUserRequest) GetRequiredApprovalCount() int {
	if a == nil {
		return 0
	}
	return a.RequiredApprovalCount
}

func (a *ApprovalRequestCreateByUserRequest) GetUserIds() []string {
	if a == nil {
		return nil
	}
	return a.UserIds
}

func (a *ApprovalRequestCreateByUserRequest) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalRequestCreateByUserRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalRequestCreateByUserRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalRequestCreateByUserRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalRequestCreateByUserRequest) String() string {
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

type ApprovalRequestCreateRequest struct {
	ApprovalRequestCreateByRoleRequest *ApprovalRequestCreateByRoleRequest
	ApprovalRequestCreateByUserRequest *ApprovalRequestCreateByUserRequest

	typ string
}

func (a *ApprovalRequestCreateRequest) GetApprovalRequestCreateByRoleRequest() *ApprovalRequestCreateByRoleRequest {
	if a == nil {
		return nil
	}
	return a.ApprovalRequestCreateByRoleRequest
}

func (a *ApprovalRequestCreateRequest) GetApprovalRequestCreateByUserRequest() *ApprovalRequestCreateByUserRequest {
	if a == nil {
		return nil
	}
	return a.ApprovalRequestCreateByUserRequest
}

func (a *ApprovalRequestCreateRequest) UnmarshalJSON(data []byte) error {
	valueApprovalRequestCreateByRoleRequest := new(ApprovalRequestCreateByRoleRequest)
	if err := json.Unmarshal(data, &valueApprovalRequestCreateByRoleRequest); err == nil {
		a.typ = "ApprovalRequestCreateByRoleRequest"
		a.ApprovalRequestCreateByRoleRequest = valueApprovalRequestCreateByRoleRequest
		return nil
	}
	valueApprovalRequestCreateByUserRequest := new(ApprovalRequestCreateByUserRequest)
	if err := json.Unmarshal(data, &valueApprovalRequestCreateByUserRequest); err == nil {
		a.typ = "ApprovalRequestCreateByUserRequest"
		a.ApprovalRequestCreateByUserRequest = valueApprovalRequestCreateByUserRequest
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ApprovalRequestCreateRequest) MarshalJSON() ([]byte, error) {
	if a.typ == "ApprovalRequestCreateByRoleRequest" || a.ApprovalRequestCreateByRoleRequest != nil {
		return json.Marshal(a.ApprovalRequestCreateByRoleRequest)
	}
	if a.typ == "ApprovalRequestCreateByUserRequest" || a.ApprovalRequestCreateByUserRequest != nil {
		return json.Marshal(a.ApprovalRequestCreateByUserRequest)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalRequestCreateRequestVisitor interface {
	VisitApprovalRequestCreateByRoleRequest(*ApprovalRequestCreateByRoleRequest) error
	VisitApprovalRequestCreateByUserRequest(*ApprovalRequestCreateByUserRequest) error
}

func (a *ApprovalRequestCreateRequest) Accept(visitor ApprovalRequestCreateRequestVisitor) error {
	if a.typ == "ApprovalRequestCreateByRoleRequest" || a.ApprovalRequestCreateByRoleRequest != nil {
		return visitor.VisitApprovalRequestCreateByRoleRequest(a.ApprovalRequestCreateByRoleRequest)
	}
	if a.typ == "ApprovalRequestCreateByUserRequest" || a.ApprovalRequestCreateByUserRequest != nil {
		return visitor.VisitApprovalRequestCreateByUserRequest(a.ApprovalRequestCreateByUserRequest)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ApprovalRequestCursorFields string

const (
	ApprovalRequestCursorFieldsCreatedAt ApprovalRequestCursorFields = "created_at"
	ApprovalRequestCursorFieldsUpdatedAt ApprovalRequestCursorFields = "updated_at"
)

func NewApprovalRequestCursorFieldsFromString(s string) (ApprovalRequestCursorFields, error) {
	switch s {
	case "created_at":
		return ApprovalRequestCursorFieldsCreatedAt, nil
	case "updated_at":
		return ApprovalRequestCursorFieldsUpdatedAt, nil
	}
	var t ApprovalRequestCursorFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalRequestCursorFields) Ptr() *ApprovalRequestCursorFields {
	return &a
}

type ApprovalRequestResourceList struct {
	Data []*ApprovalRequestResourceWithMetadata `json:"data" url:"data"`
	// A token that can be sent in the `pagination_token` query parameter to get the next page of results, or `null` if there is no next page (i.e. you've reached the last page).
	NextPaginationToken *string `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`
	// A token that can be sent in the `pagination_token` query parameter to get the previous page of results, or `null` if there is no previous page (i.e. you've reached the first page).
	PrevPaginationToken *string `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalRequestResourceList) GetData() []*ApprovalRequestResourceWithMetadata {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApprovalRequestResourceList) GetNextPaginationToken() *string {
	if a == nil {
		return nil
	}
	return a.NextPaginationToken
}

func (a *ApprovalRequestResourceList) GetPrevPaginationToken() *string {
	if a == nil {
		return nil
	}
	return a.PrevPaginationToken
}

func (a *ApprovalRequestResourceList) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalRequestResourceList) UnmarshalJSON(data []byte) error {
	type unmarshaler ApprovalRequestResourceList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApprovalRequestResourceList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApprovalRequestResourceList) String() string {
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

type ApprovalRequestResourceWithMetadata struct {
	Id         string    `json:"id" url:"id"`
	CreatedAt  time.Time `json:"created_at" url:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" url:"updated_at"`
	ApprovedBy []string  `json:"approved_by" url:"approved_by"`
	// ID of the user who created the approval request
	CreatedBy             string                `json:"created_by" url:"created_by"`
	ObjectId              string                `json:"object_id" url:"object_id"`
	ObjectType            ObjectType            `json:"object_type" url:"object_type"`
	RejectedBy            *string               `json:"rejected_by,omitempty" url:"rejected_by,omitempty"`
	RequiredApprovalCount int                   `json:"required_approval_count" url:"required_approval_count"`
	RoleIds               []string              `json:"role_ids" url:"role_ids"`
	Status                ApprovalRequestStatus `json:"status" url:"status"`
	UserIds               []string              `json:"user_ids" url:"user_ids"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *ApprovalRequestResourceWithMetadata) GetId() string {
	if a == nil {
		return ""
	}
	return a.Id
}

func (a *ApprovalRequestResourceWithMetadata) GetCreatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.CreatedAt
}

func (a *ApprovalRequestResourceWithMetadata) GetUpdatedAt() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.UpdatedAt
}

func (a *ApprovalRequestResourceWithMetadata) GetApprovedBy() []string {
	if a == nil {
		return nil
	}
	return a.ApprovedBy
}

func (a *ApprovalRequestResourceWithMetadata) GetCreatedBy() string {
	if a == nil {
		return ""
	}
	return a.CreatedBy
}

func (a *ApprovalRequestResourceWithMetadata) GetObjectId() string {
	if a == nil {
		return ""
	}
	return a.ObjectId
}

func (a *ApprovalRequestResourceWithMetadata) GetObjectType() ObjectType {
	if a == nil {
		return ""
	}
	return a.ObjectType
}

func (a *ApprovalRequestResourceWithMetadata) GetRejectedBy() *string {
	if a == nil {
		return nil
	}
	return a.RejectedBy
}

func (a *ApprovalRequestResourceWithMetadata) GetRequiredApprovalCount() int {
	if a == nil {
		return 0
	}
	return a.RequiredApprovalCount
}

func (a *ApprovalRequestResourceWithMetadata) GetRoleIds() []string {
	if a == nil {
		return nil
	}
	return a.RoleIds
}

func (a *ApprovalRequestResourceWithMetadata) GetStatus() ApprovalRequestStatus {
	if a == nil {
		return ""
	}
	return a.Status
}

func (a *ApprovalRequestResourceWithMetadata) GetUserIds() []string {
	if a == nil {
		return nil
	}
	return a.UserIds
}

func (a *ApprovalRequestResourceWithMetadata) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *ApprovalRequestResourceWithMetadata) UnmarshalJSON(data []byte) error {
	type embed ApprovalRequestResourceWithMetadata
	var unmarshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed: embed(*a),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*a = ApprovalRequestResourceWithMetadata(unmarshaler.embed)
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

func (a *ApprovalRequestResourceWithMetadata) MarshalJSON() ([]byte, error) {
	type embed ApprovalRequestResourceWithMetadata
	var marshaler = struct {
		embed
		CreatedAt *internal.DateTime `json:"created_at"`
		UpdatedAt *internal.DateTime `json:"updated_at"`
	}{
		embed:     embed(*a),
		CreatedAt: internal.NewDateTime(a.CreatedAt),
		UpdatedAt: internal.NewDateTime(a.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (a *ApprovalRequestResourceWithMetadata) String() string {
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

type ApprovalRequestStatus string

const (
	ApprovalRequestStatusWaiting  ApprovalRequestStatus = "waiting"
	ApprovalRequestStatusApproved ApprovalRequestStatus = "approved"
	ApprovalRequestStatusRejected ApprovalRequestStatus = "rejected"
	ApprovalRequestStatusCanceled ApprovalRequestStatus = "canceled"
)

func NewApprovalRequestStatusFromString(s string) (ApprovalRequestStatus, error) {
	switch s {
	case "waiting":
		return ApprovalRequestStatusWaiting, nil
	case "approved":
		return ApprovalRequestStatusApproved, nil
	case "rejected":
		return ApprovalRequestStatusRejected, nil
	case "canceled":
		return ApprovalRequestStatusCanceled, nil
	}
	var t ApprovalRequestStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApprovalRequestStatus) Ptr() *ApprovalRequestStatus {
	return &a
}
