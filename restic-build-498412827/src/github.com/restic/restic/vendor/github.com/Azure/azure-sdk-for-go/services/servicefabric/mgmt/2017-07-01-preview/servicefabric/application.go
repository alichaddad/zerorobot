package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ApplicationClient is the azure Service Fabric Resource Provider API Client
type ApplicationClient struct {
	BaseClient
}

// NewApplicationClient creates an instance of the ApplicationClient client.
func NewApplicationClient() ApplicationClient {
	return NewApplicationClientWithBaseURI(DefaultBaseURI)
}

// NewApplicationClientWithBaseURI creates an instance of the ApplicationClient client.
func NewApplicationClientWithBaseURI(baseURI string) ApplicationClient {
	return ApplicationClient{NewWithBaseURI(baseURI)}
}

// Delete deletes an application resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
func (client ApplicationClient) Delete(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (result ApplicationDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationClient) DeletePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) DeleteSender(req *http.Request) (future ApplicationDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns an application resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
func (client ApplicationClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (result ApplicationResource, err error) {
	req, err := client.GetPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationClient) GetPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationClient) GetResponder(resp *http.Response) (result ApplicationResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns all application resources in the specified cluster.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource
func (client ApplicationClient) List(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string) (result ApplicationResourceList, err error) {
	req, err := client.ListPreparer(ctx, subscriptionID, resourceGroupName, clusterName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationClient) ListPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationClient) ListResponder(resp *http.Response) (result ApplicationResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch updates an application resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// parameters is the application resource for patch operations.
func (client ApplicationClient) Patch(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, parameters ApplicationResourceUpdate) (result ApplicationPatchFuture, err error) {
	req, err := client.PatchPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Patch", nil, "Failure preparing request")
		return
	}

	result, err = client.PatchSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Patch", result.Response(), "Failure sending request")
		return
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client ApplicationClient) PatchPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, parameters ApplicationResourceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) PatchSender(req *http.Request) (future ApplicationPatchFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client ApplicationClient) PatchResponder(resp *http.Response) (result ApplicationResourceUpdate, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put creates or updates an application resource with the specified name.
//
// subscriptionID is the customer subscription identifier resourceGroupName is the name of the resource group.
// clusterName is the name of the cluster resource applicationName is the name of the application resource.
// parameters is the application resource.
func (client ApplicationClient) Put(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, parameters ApplicationResource) (result ApplicationPutFuture, err error) {
	req, err := client.PutPreparer(ctx, subscriptionID, resourceGroupName, clusterName, applicationName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Put", nil, "Failure preparing request")
		return
	}

	result, err = client.PutSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ApplicationClient", "Put", result.Response(), "Failure sending request")
		return
	}

	return
}

// PutPreparer prepares the Put request.
func (client ApplicationClient) PutPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, clusterName string, applicationName string, parameters ApplicationResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2017-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationClient) PutSender(req *http.Request) (future ApplicationPutFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client ApplicationClient) PutResponder(resp *http.Response) (result ApplicationResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
