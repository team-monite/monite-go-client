// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
	time "time"
)

type AuditLogsGetRequest struct {
	PaginationToken *string        `json:"-" url:"pagination_token,omitempty"`
	EntityUserId    *string        `json:"-" url:"entity_user_id,omitempty"`
	PathContains    *string        `json:"-" url:"path__contains,omitempty"`
	Type            *LogTypeEnum   `json:"-" url:"type,omitempty"`
	Method          *LogMethodEnum `json:"-" url:"method,omitempty"`
	StatusCode      *int           `json:"-" url:"status_code,omitempty"`
	TimestampGt     *time.Time     `json:"-" url:"timestamp__gt,omitempty"`
	TimestampLt     *time.Time     `json:"-" url:"timestamp__lt,omitempty"`
	TimestampGte    *time.Time     `json:"-" url:"timestamp__gte,omitempty"`
	TimestampLte    *time.Time     `json:"-" url:"timestamp__lte,omitempty"`
	PageSize        *int           `json:"-" url:"page_size,omitempty"`
	PageNum         *int           `json:"-" url:"page_num,omitempty"`
}

type LogMethodEnum string

const (
	LogMethodEnumGet    LogMethodEnum = "GET"
	LogMethodEnumPost   LogMethodEnum = "POST"
	LogMethodEnumPut    LogMethodEnum = "PUT"
	LogMethodEnumPatch  LogMethodEnum = "PATCH"
	LogMethodEnumDelete LogMethodEnum = "DELETE"
)

func NewLogMethodEnumFromString(s string) (LogMethodEnum, error) {
	switch s {
	case "GET":
		return LogMethodEnumGet, nil
	case "POST":
		return LogMethodEnumPost, nil
	case "PUT":
		return LogMethodEnumPut, nil
	case "PATCH":
		return LogMethodEnumPatch, nil
	case "DELETE":
		return LogMethodEnumDelete, nil
	}
	var t LogMethodEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LogMethodEnum) Ptr() *LogMethodEnum {
	return &l
}

type LogResponse struct {
	Id            string                 `json:"id" url:"id"`
	Body          *LogResponseBody       `json:"body,omitempty" url:"body,omitempty"`
	ContentType   string                 `json:"content_type" url:"content_type"`
	EntityId      string                 `json:"entity_id" url:"entity_id"`
	EntityUserId  *string                `json:"entity_user_id,omitempty" url:"entity_user_id,omitempty"`
	Headers       map[string]interface{} `json:"headers,omitempty" url:"headers,omitempty"`
	Method        *string                `json:"method,omitempty" url:"method,omitempty"`
	Params        *string                `json:"params,omitempty" url:"params,omitempty"`
	ParentLogId   *string                `json:"parent_log_id,omitempty" url:"parent_log_id,omitempty"`
	PartnerId     string                 `json:"partner_id" url:"partner_id"`
	Path          *string                `json:"path,omitempty" url:"path,omitempty"`
	StatusCode    *int                   `json:"status_code,omitempty" url:"status_code,omitempty"`
	TargetService string                 `json:"target_service" url:"target_service"`
	Timestamp     time.Time              `json:"timestamp" url:"timestamp"`
	Type          string                 `json:"type" url:"type"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LogResponse) GetId() string {
	if l == nil {
		return ""
	}
	return l.Id
}

func (l *LogResponse) GetBody() *LogResponseBody {
	if l == nil {
		return nil
	}
	return l.Body
}

func (l *LogResponse) GetContentType() string {
	if l == nil {
		return ""
	}
	return l.ContentType
}

func (l *LogResponse) GetEntityId() string {
	if l == nil {
		return ""
	}
	return l.EntityId
}

func (l *LogResponse) GetEntityUserId() *string {
	if l == nil {
		return nil
	}
	return l.EntityUserId
}

func (l *LogResponse) GetHeaders() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.Headers
}

func (l *LogResponse) GetMethod() *string {
	if l == nil {
		return nil
	}
	return l.Method
}

func (l *LogResponse) GetParams() *string {
	if l == nil {
		return nil
	}
	return l.Params
}

func (l *LogResponse) GetParentLogId() *string {
	if l == nil {
		return nil
	}
	return l.ParentLogId
}

func (l *LogResponse) GetPartnerId() string {
	if l == nil {
		return ""
	}
	return l.PartnerId
}

func (l *LogResponse) GetPath() *string {
	if l == nil {
		return nil
	}
	return l.Path
}

func (l *LogResponse) GetStatusCode() *int {
	if l == nil {
		return nil
	}
	return l.StatusCode
}

func (l *LogResponse) GetTargetService() string {
	if l == nil {
		return ""
	}
	return l.TargetService
}

func (l *LogResponse) GetTimestamp() time.Time {
	if l == nil {
		return time.Time{}
	}
	return l.Timestamp
}

func (l *LogResponse) GetType() string {
	if l == nil {
		return ""
	}
	return l.Type
}

func (l *LogResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LogResponse) UnmarshalJSON(data []byte) error {
	type embed LogResponse
	var unmarshaler = struct {
		embed
		Timestamp *internal.DateTime `json:"timestamp"`
	}{
		embed: embed(*l),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*l = LogResponse(unmarshaler.embed)
	l.Timestamp = unmarshaler.Timestamp.Time()
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LogResponse) MarshalJSON() ([]byte, error) {
	type embed LogResponse
	var marshaler = struct {
		embed
		Timestamp *internal.DateTime `json:"timestamp"`
	}{
		embed:     embed(*l),
		Timestamp: internal.NewDateTime(l.Timestamp),
	}
	return json.Marshal(marshaler)
}

func (l *LogResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LogResponseBody struct {
	StringUnknownMap map[string]interface{}
	UnknownList      []interface{}

	typ string
}

func (l *LogResponseBody) GetStringUnknownMap() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.StringUnknownMap
}

func (l *LogResponseBody) GetUnknownList() []interface{} {
	if l == nil {
		return nil
	}
	return l.UnknownList
}

func (l *LogResponseBody) UnmarshalJSON(data []byte) error {
	var valueStringUnknownMap map[string]interface{}
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		l.typ = "StringUnknownMap"
		l.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	var valueUnknownList []interface{}
	if err := json.Unmarshal(data, &valueUnknownList); err == nil {
		l.typ = "UnknownList"
		l.UnknownList = valueUnknownList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LogResponseBody) MarshalJSON() ([]byte, error) {
	if l.typ == "StringUnknownMap" || l.StringUnknownMap != nil {
		return json.Marshal(l.StringUnknownMap)
	}
	if l.typ == "UnknownList" || l.UnknownList != nil {
		return json.Marshal(l.UnknownList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LogResponseBodyVisitor interface {
	VisitStringUnknownMap(map[string]interface{}) error
	VisitUnknownList([]interface{}) error
}

func (l *LogResponseBody) Accept(visitor LogResponseBodyVisitor) error {
	if l.typ == "StringUnknownMap" || l.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(l.StringUnknownMap)
	}
	if l.typ == "UnknownList" || l.UnknownList != nil {
		return visitor.VisitUnknownList(l.UnknownList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LogTypeEnum string

const (
	LogTypeEnumRequest  LogTypeEnum = "request"
	LogTypeEnumResponse LogTypeEnum = "response"
)

func NewLogTypeEnumFromString(s string) (LogTypeEnum, error) {
	switch s {
	case "request":
		return LogTypeEnumRequest, nil
	case "response":
		return LogTypeEnumResponse, nil
	}
	var t LogTypeEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LogTypeEnum) Ptr() *LogTypeEnum {
	return &l
}

type LogsResponse struct {
	Data                []*LogResponse `json:"data" url:"data"`
	NextPaginationToken *string        `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`
	PrevPaginationToken *string        `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`
	TotalLogs           int            `json:"total_logs" url:"total_logs"`
	TotalPages          int            `json:"total_pages" url:"total_pages"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (l *LogsResponse) GetData() []*LogResponse {
	if l == nil {
		return nil
	}
	return l.Data
}

func (l *LogsResponse) GetNextPaginationToken() *string {
	if l == nil {
		return nil
	}
	return l.NextPaginationToken
}

func (l *LogsResponse) GetPrevPaginationToken() *string {
	if l == nil {
		return nil
	}
	return l.PrevPaginationToken
}

func (l *LogsResponse) GetTotalLogs() int {
	if l == nil {
		return 0
	}
	return l.TotalLogs
}

func (l *LogsResponse) GetTotalPages() int {
	if l == nil {
		return 0
	}
	return l.TotalPages
}

func (l *LogsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *LogsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler LogsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LogsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	l.rawJSON = json.RawMessage(data)
	return nil
}

func (l *LogsResponse) String() string {
	if len(l.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(l.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}
