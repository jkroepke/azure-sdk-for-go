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

// CredentialOperationsClient contains the methods for the CredentialOperations group.
// Don't use this type directly, use NewCredentialOperationsClient() instead.
type CredentialOperationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCredentialOperationsClient creates a new instance of CredentialOperationsClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCredentialOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CredentialOperationsClient, error) {
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
	client := &CredentialOperationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - credentialName - Credential name
//   - credential - Credential resource definition.
//   - options - CredentialOperationsClientCreateOrUpdateOptions contains the optional parameters for the CredentialOperationsClient.CreateOrUpdate
//     method.
func (client *CredentialOperationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, credential ManagedIdentityCredentialResource, options *CredentialOperationsClientCreateOrUpdateOptions) (CredentialOperationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, credentialName, credential, options)
	if err != nil {
		return CredentialOperationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialOperationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CredentialOperationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CredentialOperationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, credential ManagedIdentityCredentialResource, options *CredentialOperationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}"
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
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, credential)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CredentialOperationsClient) createOrUpdateHandleResponse(resp *http.Response) (CredentialOperationsClientCreateOrUpdateResponse, error) {
	result := CredentialOperationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedIdentityCredentialResource); err != nil {
		return CredentialOperationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - credentialName - Credential name
//   - options - CredentialOperationsClientDeleteOptions contains the optional parameters for the CredentialOperationsClient.Delete
//     method.
func (client *CredentialOperationsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *CredentialOperationsClientDeleteOptions) (CredentialOperationsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, credentialName, options)
	if err != nil {
		return CredentialOperationsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialOperationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CredentialOperationsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return CredentialOperationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CredentialOperationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *CredentialOperationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}"
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
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - credentialName - Credential name
//   - options - CredentialOperationsClientGetOptions contains the optional parameters for the CredentialOperationsClient.Get
//     method.
func (client *CredentialOperationsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *CredentialOperationsClientGetOptions) (CredentialOperationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, credentialName, options)
	if err != nil {
		return CredentialOperationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CredentialOperationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return CredentialOperationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CredentialOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, options *CredentialOperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}"
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
	if credentialName == "" {
		return nil, errors.New("parameter credentialName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{credentialName}", url.PathEscape(credentialName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CredentialOperationsClient) getHandleResponse(resp *http.Response) (CredentialOperationsClientGetResponse, error) {
	result := CredentialOperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedIdentityCredentialResource); err != nil {
		return CredentialOperationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - List credentials.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - CredentialOperationsClientListByFactoryOptions contains the optional parameters for the CredentialOperationsClient.NewListByFactoryPager
//     method.
func (client *CredentialOperationsClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *CredentialOperationsClientListByFactoryOptions) *runtime.Pager[CredentialOperationsClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[CredentialOperationsClientListByFactoryResponse]{
		More: func(page CredentialOperationsClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CredentialOperationsClientListByFactoryResponse) (CredentialOperationsClientListByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CredentialOperationsClientListByFactoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CredentialOperationsClientListByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CredentialOperationsClientListByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFactoryHandleResponse(resp)
		},
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *CredentialOperationsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *CredentialOperationsClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *CredentialOperationsClient) listByFactoryHandleResponse(resp *http.Response) (CredentialOperationsClientListByFactoryResponse, error) {
	result := CredentialOperationsClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CredentialListResponse); err != nil {
		return CredentialOperationsClientListByFactoryResponse{}, err
	}
	return result, nil
}
