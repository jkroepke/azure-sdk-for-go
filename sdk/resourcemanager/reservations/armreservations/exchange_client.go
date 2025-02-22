//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armreservations

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ExchangeClient contains the methods for the Exchange group.
// Don't use this type directly, use NewExchangeClient() instead.
type ExchangeClient struct {
	host string
	pl   runtime.Pipeline
}

// NewExchangeClient creates a new instance of ExchangeClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExchangeClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ExchangeClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ExchangeClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginPost - Returns one or more Reservations in exchange for one or more Reservation purchases.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - body - Request containing the refunds and purchases that need to be executed.
//   - options - ExchangeClientBeginPostOptions contains the optional parameters for the ExchangeClient.BeginPost method.
func (client *ExchangeClient) BeginPost(ctx context.Context, body ExchangeRequest, options *ExchangeClientBeginPostOptions) (*runtime.Poller[ExchangeClientPostResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.post(ctx, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ExchangeClientPostResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ExchangeClientPostResponse](options.ResumeToken, client.pl, nil)
	}
}

// Post - Returns one or more Reservations in exchange for one or more Reservation purchases.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *ExchangeClient) post(ctx context.Context, body ExchangeRequest, options *ExchangeClientBeginPostOptions) (*http.Response, error) {
	req, err := client.postCreateRequest(ctx, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// postCreateRequest creates the Post request.
func (client *ExchangeClient) postCreateRequest(ctx context.Context, body ExchangeRequest, options *ExchangeClientBeginPostOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Capacity/exchange"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
