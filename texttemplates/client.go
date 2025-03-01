// This file was auto-generated by Fern from our API Definition.

package texttemplates

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

// Get text templates
func (c *Client) Get(
	ctx context.Context,
	request *monitegoclient.TextTemplatesGetRequest,
	opts ...option.RequestOption,
) (*monitegoclient.TextTemplateResponseList, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/text_templates"
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

	var response *monitegoclient.TextTemplateResponseList
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

// Create a text template
func (c *Client) Create(
	ctx context.Context,
	request *monitegoclient.CreateTextTemplatePayload,
	opts ...option.RequestOption,
) (*monitegoclient.TextTemplateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/text_templates"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.TextTemplateResponse
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

// Get all custom contents
func (c *Client) GetById(
	ctx context.Context,
	textTemplateId string,
	opts ...option.RequestOption,
) (*monitegoclient.TextTemplateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/text_templates/%v",
		textTemplateId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.TextTemplateResponse
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

// Delete custom content by ID
func (c *Client) DeleteById(
	ctx context.Context,
	// UUID text_template ID
	textTemplateId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/text_templates/%v",
		textTemplateId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
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

	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return err
	}
	return nil
}

// Update custom content by ID
func (c *Client) UpdateById(
	ctx context.Context,
	// UUID text_template ID
	textTemplateId string,
	request *monitegoclient.UpdateTextTemplatePayload,
	opts ...option.RequestOption,
) (*monitegoclient.TextTemplateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/text_templates/%v",
		textTemplateId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.TextTemplateResponse
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

// Make text template default
func (c *Client) MakeDefaultById(
	ctx context.Context,
	// UUID text_template ID
	textTemplateId string,
	opts ...option.RequestOption,
) (*monitegoclient.TextTemplateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/text_templates/%v/make_default",
		textTemplateId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.TextTemplateResponse
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
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
