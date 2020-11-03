// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type dataFlowsClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *dataFlowsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Rename - Renames a dataflow.
func (client *dataFlowsClient) Rename(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowsRenameOptions) (*azcore.Response, error) {
	req, err := client.RenameCreateRequest(ctx, dataFlowName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.RenameHandleError(resp)
	}
	return resp, nil
}

// RenameCreateRequest creates the Rename request.
func (client *dataFlowsClient) RenameCreateRequest(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowsRenameOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// RenameHandleError handles the Rename error response.
func (client *dataFlowsClient) RenameHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}