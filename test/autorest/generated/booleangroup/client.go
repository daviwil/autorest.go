// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
)

type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

// Client - Test Infrastructure for AutoRest
type Client struct {
	u *url.URL
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(options *ClientOptions) (*Client, error) {
	return NewClient(DefaultEndpoint, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.Telemetry),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{u: u, p: p}, nil
}

// BoolOperations returns the BoolOperations associated with this client.
func (client *Client) BoolOperations() BoolOperations {
	return &boolOperations{Client: client}
}
