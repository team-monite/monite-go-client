// This file was auto-generated by Fern from our API Definition.

package taxrates

import (
	context "context"
	monitegoclient "github.com/team-monite/monite-go-client"
	accounting "github.com/team-monite/monite-go-client/accounting"
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

// Get all tax rate accounts
func (c *Client) Get(
	ctx context.Context,
	request *accounting.TaxRatesGetRequest,
	opts ...option.RequestOption,
) (*monitegoclient.AccountingTaxRateListResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := baseURL + "/accounting_tax_rates"
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

	var response *monitegoclient.AccountingTaxRateListResponse
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

// Get tax rate account by id
func (c *Client) GetById(
	ctx context.Context,
	taxRateId string,
	opts ...option.RequestOption,
) (*monitegoclient.AccountingTaxRateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://api.sandbox.monite.com/v1",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/accounting_tax_rates/%v",
		taxRateId,
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

	var response *monitegoclient.AccountingTaxRateResponse
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
