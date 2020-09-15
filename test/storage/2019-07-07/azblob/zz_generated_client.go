// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const scope = "https://storage.azure.com/.default"
const telemetryInfo = "azsdk-go-azblob/<version>"

// clientOptions contains configuration settings for the default client's pipeline.
type clientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// defaultClientOptions creates a clientOptions type initialized with default values.
func defaultClientOptions() clientOptions {
	return clientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *clientOptions) telemetryOptions() azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return to
}

type client struct {
	u string
	p azcore.Pipeline
}

// newClient creates an instance of the client type with the specified endpoint.
func newClient(endpoint string, cred azcore.Credential, options *clientOptions) *client {
	if options == nil {
		o := defaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{scope}}}),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return newClientWithPipeline(endpoint, p)
}

// newClientWithPipeline creates an instance of the client type with the specified endpoint and pipeline.
func newClientWithPipeline(endpoint string, p azcore.Pipeline) *client {
	return &client{u: endpoint, p: p}
}