// This file was auto-generated by Fern from our API Definition.

package roles

import (
	context "context"
	monitegoclient "github.com/team-monite/monite-go-client"
	core "github.com/team-monite/monite-go-client/core"
	internal "github.com/team-monite/monite-go-client/internal"
	option "github.com/team-monite/monite-go-client/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

// Find all roles that match the search criteria.
func (c *Client) Get(
	ctx context.Context,
	request *monitegoclient.RolesGetRequest,
	opts ...option.RequestOption,
) (*monitegoclient.RolePaginationResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/roles"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &monitegoclient.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &monitegoclient.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &monitegoclient.ForbiddenError{
				APIError: apiError,
			}
		},
		406: func(apiError *core.APIError) error {
			return &monitegoclient.NotAcceptableError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &monitegoclient.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &monitegoclient.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *monitegoclient.RolePaginationResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Create a new role from the specified values.
func (c *Client) Create(
	ctx context.Context,
	request *monitegoclient.CreateRoleRequest,
	opts ...option.RequestOption,
) (*monitegoclient.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/roles"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &monitegoclient.BadRequestError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &monitegoclient.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &monitegoclient.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *monitegoclient.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetById(
	ctx context.Context,
	roleId string,
	opts ...option.RequestOption,
) (*monitegoclient.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/roles/%v",
		roleId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &monitegoclient.BadRequestError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &monitegoclient.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &monitegoclient.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *monitegoclient.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Change the specified fields with the provided values.
func (c *Client) UpdateById(
	ctx context.Context,
	roleId string,
	request *monitegoclient.UpdateRoleRequest,
	opts ...option.RequestOption,
) (*monitegoclient.RoleResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/roles/%v",
		roleId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &monitegoclient.BadRequestError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &monitegoclient.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &monitegoclient.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *monitegoclient.RoleResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
