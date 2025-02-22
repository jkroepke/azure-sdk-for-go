//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationRuntimeObjectMetadataClient contains the methods for the IntegrationRuntimeObjectMetadata group.
// Don't use this type directly, use NewIntegrationRuntimeObjectMetadataClient() instead.
type IntegrationRuntimeObjectMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationRuntimeObjectMetadataClient creates a new instance of IntegrationRuntimeObjectMetadataClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationRuntimeObjectMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeObjectMetadataClient, error) {
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
	client := &IntegrationRuntimeObjectMetadataClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a SSIS integration runtime object metadata by specified path. The return is pageable metadata list.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - integrationRuntimeName - The integration runtime name.
//   - options - IntegrationRuntimeObjectMetadataClientGetOptions contains the optional parameters for the IntegrationRuntimeObjectMetadataClient.Get
//     method.
func (client *IntegrationRuntimeObjectMetadataClient) Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientGetOptions) (IntegrationRuntimeObjectMetadataClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeObjectMetadataClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeObjectMetadataClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/getObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.GetMetadataRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.GetMetadataRequest)
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeObjectMetadataClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeObjectMetadataClientGetResponse, error) {
	result := IntegrationRuntimeObjectMetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SsisObjectMetadataListResponse); err != nil {
		return IntegrationRuntimeObjectMetadataClientGetResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refresh a SSIS integration runtime object metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - integrationRuntimeName - The integration runtime name.
//   - options - IntegrationRuntimeObjectMetadataClientBeginRefreshOptions contains the optional parameters for the IntegrationRuntimeObjectMetadataClient.BeginRefresh
//     method.
func (client *IntegrationRuntimeObjectMetadataClient) BeginRefresh(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (*runtime.Poller[IntegrationRuntimeObjectMetadataClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[IntegrationRuntimeObjectMetadataClientRefreshResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[IntegrationRuntimeObjectMetadataClientRefreshResponse](options.ResumeToken, client.pl, nil)
	}
}

// Refresh - Refresh a SSIS integration runtime object metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
func (client *IntegrationRuntimeObjectMetadataClient) refresh(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
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

// refreshCreateRequest creates the Refresh request.
func (client *IntegrationRuntimeObjectMetadataClient) refreshCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/refreshObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
