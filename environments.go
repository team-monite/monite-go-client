// This file was auto-generated by Fern from our API Definition.

package fluidstack

// Environments defines all of the API environments.
// These values can be used with the WithBaseURL
// RequestOption to override the client's default environment,
// if any.
var Environments = struct {
	Sandbox      string
	EuProduction string
	NaProduction string
}{
	Sandbox:      "https://api.sandbox.monite.com/v1",
	EuProduction: "https://api.monite.com/v1",
	NaProduction: "https://us.api.monite.com/v1",
}
