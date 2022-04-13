//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SQLPoolWorkloadGroupClient contains the methods for the SQLPoolWorkloadGroup group.
// Don't use this type directly, use NewSQLPoolWorkloadGroupClient() instead.
type SQLPoolWorkloadGroupClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSQLPoolWorkloadGroupClient creates a new instance of SQLPoolWorkloadGroupClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSQLPoolWorkloadGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLPoolWorkloadGroupClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SQLPoolWorkloadGroupClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create Or Update a Sql pool's workload group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// workloadGroupName - The name of the workload group.
// parameters - The requested workload group state.
// options - SQLPoolWorkloadGroupClientBeginCreateOrUpdateOptions contains the optional parameters for the SQLPoolWorkloadGroupClient.BeginCreateOrUpdate
// method.
func (client *SQLPoolWorkloadGroupClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, parameters WorkloadGroup, options *SQLPoolWorkloadGroupClientBeginCreateOrUpdateOptions) (*armruntime.Poller[SQLPoolWorkloadGroupClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SQLPoolWorkloadGroupClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SQLPoolWorkloadGroupClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create Or Update a Sql pool's workload group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLPoolWorkloadGroupClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, parameters WorkloadGroup, options *SQLPoolWorkloadGroupClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLPoolWorkloadGroupClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, parameters WorkloadGroup, options *SQLPoolWorkloadGroupClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Remove Sql pool's workload group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// workloadGroupName - The name of the workload group.
// options - SQLPoolWorkloadGroupClientBeginDeleteOptions contains the optional parameters for the SQLPoolWorkloadGroupClient.BeginDelete
// method.
func (client *SQLPoolWorkloadGroupClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadGroupClientBeginDeleteOptions) (*armruntime.Poller[SQLPoolWorkloadGroupClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[SQLPoolWorkloadGroupClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[SQLPoolWorkloadGroupClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Remove Sql pool's workload group.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SQLPoolWorkloadGroupClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadGroupClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLPoolWorkloadGroupClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadGroupClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a Sql pool's workload group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// workloadGroupName - The name of the workload group.
// options - SQLPoolWorkloadGroupClientGetOptions contains the optional parameters for the SQLPoolWorkloadGroupClient.Get
// method.
func (client *SQLPoolWorkloadGroupClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadGroupClientGetOptions) (SQLPoolWorkloadGroupClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, workloadGroupName, options)
	if err != nil {
		return SQLPoolWorkloadGroupClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SQLPoolWorkloadGroupClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLPoolWorkloadGroupClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLPoolWorkloadGroupClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, workloadGroupName string, options *SQLPoolWorkloadGroupClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups/{workloadGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	if workloadGroupName == "" {
		return nil, errors.New("parameter workloadGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workloadGroupName}", url.PathEscape(workloadGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLPoolWorkloadGroupClient) getHandleResponse(resp *http.Response) (SQLPoolWorkloadGroupClientGetResponse, error) {
	result := SQLPoolWorkloadGroupClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadGroup); err != nil {
		return SQLPoolWorkloadGroupClientGetResponse{}, err
	}
	return result, nil
}

// List - Get list of Sql pool's workload groups.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// sqlPoolName - SQL pool name
// options - SQLPoolWorkloadGroupClientListOptions contains the optional parameters for the SQLPoolWorkloadGroupClient.List
// method.
func (client *SQLPoolWorkloadGroupClient) List(resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolWorkloadGroupClientListOptions) *runtime.Pager[SQLPoolWorkloadGroupClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[SQLPoolWorkloadGroupClientListResponse]{
		More: func(page SQLPoolWorkloadGroupClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLPoolWorkloadGroupClientListResponse) (SQLPoolWorkloadGroupClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, sqlPoolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLPoolWorkloadGroupClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SQLPoolWorkloadGroupClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLPoolWorkloadGroupClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLPoolWorkloadGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sqlPoolName string, options *SQLPoolWorkloadGroupClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/workloadGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sqlPoolName == "" {
		return nil, errors.New("parameter sqlPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlPoolName}", url.PathEscape(sqlPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLPoolWorkloadGroupClient) listHandleResponse(resp *http.Response) (SQLPoolWorkloadGroupClientListResponse, error) {
	result := SQLPoolWorkloadGroupClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadGroupListResult); err != nil {
		return SQLPoolWorkloadGroupClientListResponse{}, err
	}
	return result, nil
}
