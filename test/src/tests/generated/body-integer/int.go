package integergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// IntClient is the test Infrastructure for AutoRest
type IntClient struct {
	BaseClient
}

// NewIntClient creates an instance of the IntClient client.
func NewIntClient() IntClient {
	return NewIntClientWithBaseURI(DefaultBaseURI)
}

// NewIntClientWithBaseURI creates an instance of the IntClient client.
func NewIntClientWithBaseURI(baseURI string) IntClient {
	return IntClient{NewWithBaseURI(baseURI)}
}

// GetInvalid get invalid Int value
func (client IntClient) GetInvalid(ctx context.Context) (result Int32, err error) {
	req, err := client.GetInvalidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client IntClient) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client IntClient) GetInvalidResponder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInvalidUnixTime get invalid Unix time value
func (client IntClient) GetInvalidUnixTime(ctx context.Context) (result UnixTime, err error) {
	req, err := client.GetInvalidUnixTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalidUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalidUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetInvalidUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetInvalidUnixTimePreparer prepares the GetInvalidUnixTime request.
func (client IntClient) GetInvalidUnixTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/invalidunixtime"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidUnixTimeSender sends the GetInvalidUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetInvalidUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetInvalidUnixTimeResponder handles the response to the GetInvalidUnixTime request. The method always
// closes the http.Response Body.
func (client IntClient) GetInvalidUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null Int value
func (client IntClient) GetNull(ctx context.Context) (result Int32, err error) {
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client IntClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetNullSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client IntClient) GetNullResponder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNullUnixTime get null Unix time value
func (client IntClient) GetNullUnixTime(ctx context.Context) (result UnixTime, err error) {
	req, err := client.GetNullUnixTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNullUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNullUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetNullUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetNullUnixTimePreparer prepares the GetNullUnixTime request.
func (client IntClient) GetNullUnixTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/nullunixtime"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullUnixTimeSender sends the GetNullUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetNullUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNullUnixTimeResponder handles the response to the GetNullUnixTime request. The method always
// closes the http.Response Body.
func (client IntClient) GetNullUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowInt32 get overflow Int32 value
func (client IntClient) GetOverflowInt32(ctx context.Context) (result Int32, err error) {
	req, err := client.GetOverflowInt32Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt32", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowInt32Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt32", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowInt32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt32", resp, "Failure responding to request")
	}

	return
}

// GetOverflowInt32Preparer prepares the GetOverflowInt32 request.
func (client IntClient) GetOverflowInt32Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/overflowint32"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOverflowInt32Sender sends the GetOverflowInt32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetOverflowInt32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowInt32Responder handles the response to the GetOverflowInt32 request. The method always
// closes the http.Response Body.
func (client IntClient) GetOverflowInt32Responder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOverflowInt64 get overflow Int64 value
func (client IntClient) GetOverflowInt64(ctx context.Context) (result Int64, err error) {
	req, err := client.GetOverflowInt64Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt64", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOverflowInt64Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt64", resp, "Failure sending request")
		return
	}

	result, err = client.GetOverflowInt64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetOverflowInt64", resp, "Failure responding to request")
	}

	return
}

// GetOverflowInt64Preparer prepares the GetOverflowInt64 request.
func (client IntClient) GetOverflowInt64Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/overflowint64"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOverflowInt64Sender sends the GetOverflowInt64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetOverflowInt64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOverflowInt64Responder handles the response to the GetOverflowInt64 request. The method always
// closes the http.Response Body.
func (client IntClient) GetOverflowInt64Responder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowInt32 get underflow Int32 value
func (client IntClient) GetUnderflowInt32(ctx context.Context) (result Int32, err error) {
	req, err := client.GetUnderflowInt32Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt32", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowInt32Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt32", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowInt32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt32", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowInt32Preparer prepares the GetUnderflowInt32 request.
func (client IntClient) GetUnderflowInt32Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/underflowint32"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUnderflowInt32Sender sends the GetUnderflowInt32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetUnderflowInt32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowInt32Responder handles the response to the GetUnderflowInt32 request. The method always
// closes the http.Response Body.
func (client IntClient) GetUnderflowInt32Responder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnderflowInt64 get underflow Int64 value
func (client IntClient) GetUnderflowInt64(ctx context.Context) (result Int64, err error) {
	req, err := client.GetUnderflowInt64Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt64", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnderflowInt64Sender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt64", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnderflowInt64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnderflowInt64", resp, "Failure responding to request")
	}

	return
}

// GetUnderflowInt64Preparer prepares the GetUnderflowInt64 request.
func (client IntClient) GetUnderflowInt64Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/underflowint64"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUnderflowInt64Sender sends the GetUnderflowInt64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetUnderflowInt64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnderflowInt64Responder handles the response to the GetUnderflowInt64 request. The method always
// closes the http.Response Body.
func (client IntClient) GetUnderflowInt64Responder(resp *http.Response) (result Int64, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetUnixTime get datetime encoded as Unix time value
func (client IntClient) GetUnixTime(ctx context.Context) (result UnixTime, err error) {
	req, err := client.GetUnixTimePreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnixTime", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetUnixTimeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnixTime", resp, "Failure sending request")
		return
	}

	result, err = client.GetUnixTimeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "GetUnixTime", resp, "Failure responding to request")
	}

	return
}

// GetUnixTimePreparer prepares the GetUnixTime request.
func (client IntClient) GetUnixTimePreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/unixtime"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetUnixTimeSender sends the GetUnixTime request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) GetUnixTimeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetUnixTimeResponder handles the response to the GetUnixTime request. The method always
// closes the http.Response Body.
func (client IntClient) GetUnixTimeResponder(resp *http.Response) (result UnixTime, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutMax32 put max int32 value
//
func (client IntClient) PutMax32(ctx context.Context, intBody int32) (result autorest.Response, err error) {
	req, err := client.PutMax32Preparer(ctx, intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax32", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMax32Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax32", resp, "Failure sending request")
		return
	}

	result, err = client.PutMax32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax32", resp, "Failure responding to request")
	}

	return
}

// PutMax32Preparer prepares the PutMax32 request.
func (client IntClient) PutMax32Preparer(ctx context.Context, intBody int32) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/max/32"),
		autorest.WithJSON(intBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMax32Sender sends the PutMax32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) PutMax32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMax32Responder handles the response to the PutMax32 request. The method always
// closes the http.Response Body.
func (client IntClient) PutMax32Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMax64 put max int64 value
//
func (client IntClient) PutMax64(ctx context.Context, intBody int64) (result autorest.Response, err error) {
	req, err := client.PutMax64Preparer(ctx, intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax64", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMax64Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax64", resp, "Failure sending request")
		return
	}

	result, err = client.PutMax64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMax64", resp, "Failure responding to request")
	}

	return
}

// PutMax64Preparer prepares the PutMax64 request.
func (client IntClient) PutMax64Preparer(ctx context.Context, intBody int64) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/max/64"),
		autorest.WithJSON(intBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMax64Sender sends the PutMax64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) PutMax64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMax64Responder handles the response to the PutMax64 request. The method always
// closes the http.Response Body.
func (client IntClient) PutMax64Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMin32 put min int32 value
//
func (client IntClient) PutMin32(ctx context.Context, intBody int32) (result autorest.Response, err error) {
	req, err := client.PutMin32Preparer(ctx, intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin32", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMin32Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin32", resp, "Failure sending request")
		return
	}

	result, err = client.PutMin32Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin32", resp, "Failure responding to request")
	}

	return
}

// PutMin32Preparer prepares the PutMin32 request.
func (client IntClient) PutMin32Preparer(ctx context.Context, intBody int32) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/min/32"),
		autorest.WithJSON(intBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMin32Sender sends the PutMin32 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) PutMin32Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMin32Responder handles the response to the PutMin32 request. The method always
// closes the http.Response Body.
func (client IntClient) PutMin32Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutMin64 put min int64 value
//
func (client IntClient) PutMin64(ctx context.Context, intBody int64) (result autorest.Response, err error) {
	req, err := client.PutMin64Preparer(ctx, intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin64", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutMin64Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin64", resp, "Failure sending request")
		return
	}

	result, err = client.PutMin64Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutMin64", resp, "Failure responding to request")
	}

	return
}

// PutMin64Preparer prepares the PutMin64 request.
func (client IntClient) PutMin64Preparer(ctx context.Context, intBody int64) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/min/64"),
		autorest.WithJSON(intBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutMin64Sender sends the PutMin64 request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) PutMin64Sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutMin64Responder handles the response to the PutMin64 request. The method always
// closes the http.Response Body.
func (client IntClient) PutMin64Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PutUnixTimeDate put datetime encoded as Unix time
//
func (client IntClient) PutUnixTimeDate(ctx context.Context, intBody date.UnixTime) (result autorest.Response, err error) {
	req, err := client.PutUnixTimeDatePreparer(ctx, intBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutUnixTimeDate", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutUnixTimeDateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutUnixTimeDate", resp, "Failure sending request")
		return
	}

	result, err = client.PutUnixTimeDateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integergroup.IntClient", "PutUnixTimeDate", resp, "Failure responding to request")
	}

	return
}

// PutUnixTimeDatePreparer prepares the PutUnixTimeDate request.
func (client IntClient) PutUnixTimeDatePreparer(ctx context.Context, intBody date.UnixTime) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/int/unixtime"),
		autorest.WithJSON(intBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutUnixTimeDateSender sends the PutUnixTimeDate request. The method will close the
// http.Response Body if it receives an error.
func (client IntClient) PutUnixTimeDateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PutUnixTimeDateResponder handles the response to the PutUnixTimeDate request. The method always
// closes the http.Response Body.
func (client IntClient) PutUnixTimeDateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
