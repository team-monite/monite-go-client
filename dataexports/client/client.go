// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	monitegoclient "github.com/team-monite/monite-go-client"
	core "github.com/team-monite/monite-go-client/core"
	extradata "github.com/team-monite/monite-go-client/dataexports/extradata"
	internal "github.com/team-monite/monite-go-client/internal"
	option "github.com/team-monite/monite-go-client/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header

	ExtraData *extradata.Client
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
		header:    options.ToHeader(),
		ExtraData: extradata.NewClient(opts...),
	}
}

func (c *Client) Get(
	ctx context.Context,
	request *monitegoclient.DataExportsGetRequest,
	opts ...option.RequestOption,
) (*monitegoclient.AllDocumentExportResponseSchema, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/data_exports"
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

	var response *monitegoclient.AllDocumentExportResponseSchema
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

// Request the export of payable and receivable documents with the specified statuses.
func (c *Client) Create(
	ctx context.Context,
	request *monitegoclient.ExportPayloadSchema,
	opts ...option.RequestOption,
) (*monitegoclient.CreateExportTaskResponseSchema, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/data_exports"
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
		404: func(apiError *core.APIError) error {
			return &monitegoclient.NotFoundError{
				APIError: apiError,
			}
		},
		409: func(apiError *core.APIError) error {
			return &monitegoclient.ConflictError{
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

	var response *monitegoclient.CreateExportTaskResponseSchema
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

func (c *Client) GetSupportedFormats(
	ctx context.Context,
	opts ...option.RequestOption,
) ([]*monitegoclient.SupportedFormatSchema, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/data_exports/supported_formats"
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

	var response []*monitegoclient.SupportedFormatSchema
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

func (c *Client) GetById(
	ctx context.Context,
	documentExportId string,
	opts ...option.RequestOption,
) (*monitegoclient.DocumentExportResponseSchema, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/data_exports/%v",
		documentExportId,
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
		404: func(apiError *core.APIError) error {
			return &monitegoclient.NotFoundError{
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

	var response *monitegoclient.DocumentExportResponseSchema
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
