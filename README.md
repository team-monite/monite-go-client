# Monite Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)

The Monite Go library provides convenient access to the Monite API from Go.

## Requirements

This module requires Go version >= 1.18.

## Installation

Run the following command to use the Monite Go library in your module:

```sh
go get github.com/team-monite/monite-go-client
```

## Usage

```go
import moniteclient "github.com/team-monite/monite-go-client/client"

client := moniteclient.NewClient(moniteclient.WithToken("<YOUR_AUTH_TOKEN>"))
```

## Usage

```go
import (
  monite       "github.com/team-monite/monite-go-client"
  moniteclient "github.com/team-monite/monite-go-client/client"
)

client := moniteclient.NewClient(moniteclient.WithToken("<YOUR_AUTH_TOKEN>"))
response, err := moniteclient.Products.Create(
  context.TODO(),
  &monite.ProductServiceRequest{
    Name: "Name",
  },
)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := moniteclient.Products.Create(
  context.TODO(),
  &monite.ProductServiceRequest{
    Name: "Name",
  },
)
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which includes
configuring authorization tokens to be sent on every request, or providing your own instrumented
`*http.Client`. Both of these options are shown below:

```go
client := moniteclient.NewClient(
  moniteclient.WithToken("<YOUR_AUTH_TOKEN>"),
  moniteclient.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
response, err := moniteclient.Products.Create(
  context.TODO(),
  &monite.ProductServiceRequest{
    Name: "Name",
  },
)
if err != nil {
  if badRequestErr, ok := err.(*monite.BadRequestError);
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := moniteclient.Products.Create(
  context.TODO(),
  &monite.ProductServiceRequest{
    Name: "Name",
  },
)
if err != nil {
  var badRequestErr *monite.BadRequestError
  if errors.As(err, badRequestErr) {
    // Do something with the bad request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := moniteclient.Products.Create(
  context.TODO(),
  &monite.ProductServiceRequest{
    Name: "Name",
  },
)
if err != nil {
  return fmt.Errorf("failed to generate response: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. 
Additions made directly to this library would have to be moved over to our generation code, 
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
 a proof of concept, but know that we will not be able to merge it as-is. We suggest opening 
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
