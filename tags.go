// This file was auto-generated by Fern from our API Definition.

package fluidstack

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/team-monite/monite-go-client/internal"
)

type TagCreateSchema struct {
	// The tag category.
	Category *TagCategory `json:"category,omitempty" url:"-"`
	// The tag description.
	Description *string `json:"description,omitempty" url:"-"`
	// The tag name. Must be unique.
	Name string `json:"name" url:"-"`
}

type TagsGetRequest struct {
	// Sort order (ascending by default). Typically used together with the `sort` parameter.
	Order *OrderEnum `json:"-" url:"order,omitempty"`
	// The number of items (0 .. 100) to return in a single page of the response. The response may contain fewer items if it is the last or only page.
	Limit *int `json:"-" url:"limit,omitempty"`
	// A pagination token obtained from a previous call to this endpoint. Use it to get the next or previous page of results for your initial query. If `pagination_token` is specified, all other query parameters are ignored and inferred from the initial query.
	//
	// If not specified, the first page of results will be returned.
	PaginationToken *string `json:"-" url:"pagination_token,omitempty"`
	// The field to sort the results by. Typically used together with the `order` parameter.
	Sort                  *TagCursorFields `json:"-" url:"sort,omitempty"`
	CreatedByEntityUserId *string          `json:"-" url:"created_by_entity_user_id,omitempty"`
	NameIn                []*string        `json:"-" url:"name__in,omitempty"`
	IdIn                  []*string        `json:"-" url:"id__in,omitempty"`
}

type TagCursorFields string

const (
	TagCursorFieldsCreatedAt TagCursorFields = "created_at"
	TagCursorFieldsUpdatedAt TagCursorFields = "updated_at"
)

func NewTagCursorFieldsFromString(s string) (TagCursorFields, error) {
	switch s {
	case "created_at":
		return TagCursorFieldsCreatedAt, nil
	case "updated_at":
		return TagCursorFieldsUpdatedAt, nil
	}
	var t TagCursorFields
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TagCursorFields) Ptr() *TagCursorFields {
	return &t
}

// A paginated list of tags.
type TagsPaginationResponse struct {
	Data []*TagReadSchema `json:"data" url:"data"`
	// A token that can be sent in the `pagination_token` query parameter to get the next page of results, or `null` if there is no next page (i.e. you've reached the last page).
	NextPaginationToken *string `json:"next_pagination_token,omitempty" url:"next_pagination_token,omitempty"`
	// A token that can be sent in the `pagination_token` query parameter to get the previous page of results, or `null` if there is no previous page (i.e. you've reached the first page).
	PrevPaginationToken *string `json:"prev_pagination_token,omitempty" url:"prev_pagination_token,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *TagsPaginationResponse) GetData() []*TagReadSchema {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TagsPaginationResponse) GetNextPaginationToken() *string {
	if t == nil {
		return nil
	}
	return t.NextPaginationToken
}

func (t *TagsPaginationResponse) GetPrevPaginationToken() *string {
	if t == nil {
		return nil
	}
	return t.PrevPaginationToken
}

func (t *TagsPaginationResponse) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *TagsPaginationResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler TagsPaginationResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TagsPaginationResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *TagsPaginationResponse) String() string {
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

type TagUpdateSchema struct {
	// The tag category.
	Category *TagCategory `json:"category,omitempty" url:"-"`
	// The tag description.
	Description *string `json:"description,omitempty" url:"-"`
	// The tag name. Must be unique.
	Name *string `json:"name,omitempty" url:"-"`
}
