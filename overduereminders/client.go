// This file was auto-generated by Fern from our API Definition.

package overduereminders

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

func (c *Client) Get(
	ctx context.Context,
	opts ...option.RequestOption,
) (*monitegoclient.AllOverdueRemindersResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/overdue_reminders"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.AllOverdueRemindersResponse
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

func (c *Client) Create(
	ctx context.Context,
	request *monitegoclient.OverdueReminderRequest,
	opts ...option.RequestOption,
) (*monitegoclient.OverdueReminderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/overdue_reminders"
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

	var response *monitegoclient.OverdueReminderResponse
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
	overdueReminderId string,
	opts ...option.RequestOption,
) (*monitegoclient.OverdueReminderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/overdue_reminders/%v",
		overdueReminderId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
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

	var response *monitegoclient.OverdueReminderResponse
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

func (c *Client) DeleteById(
	ctx context.Context,
	overdueReminderId string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/overdue_reminders/%v",
		overdueReminderId,
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

func (c *Client) UpdateById(
	ctx context.Context,
	overdueReminderId string,
	request *monitegoclient.OverdueReminderUpdateRequest,
	opts ...option.RequestOption,
) (*monitegoclient.OverdueReminderResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/overdue_reminders/%v",
		overdueReminderId,
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

	var response *monitegoclient.OverdueReminderResponse
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
