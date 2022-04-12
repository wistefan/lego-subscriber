/*
ETSI ISG CIM / NGSI-LD API

This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders

API version: latest
Contact: NGSI-LD@etsi.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// CSourceSubscriptionsApiService CSourceSubscriptionsApi service
type CSourceSubscriptionsApiService service

type ApiCreateCSourceSubscriptionRequest struct {
	ctx context.Context
	ApiService *CSourceSubscriptionsApiService
	subscription *Subscription
}

func (r ApiCreateCSourceSubscriptionRequest) Subscription(subscription Subscription) ApiCreateCSourceSubscriptionRequest {
	r.subscription = &subscription
	return r
}

func (r ApiCreateCSourceSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateCSourceSubscriptionExecute(r)
}

/*
CreateCSourceSubscription Method for CreateCSourceSubscription

Creates a context source discovery Subscription within an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCSourceSubscriptionRequest
*/
func (a *CSourceSubscriptionsApiService) CreateCSourceSubscription(ctx context.Context) ApiCreateCSourceSubscriptionRequest {
	return ApiCreateCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CSourceSubscriptionsApiService) CreateCSourceSubscriptionExecute(r ApiCreateCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CSourceSubscriptionsApiService.CreateCSourceSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceSubscriptions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscription == nil {
		return nil, reportError("subscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/ld+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.subscription
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoveCSourceSubscriptionRequest struct {
	ctx context.Context
	ApiService *CSourceSubscriptionsApiService
	subscriptionId string
}

func (r ApiRemoveCSourceSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveCSourceSubscriptionExecute(r)
}

/*
RemoveCSourceSubscription Method for RemoveCSourceSubscription

Removes a specific Context Source Subscription from an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription Id
 @return ApiRemoveCSourceSubscriptionRequest
*/
func (a *CSourceSubscriptionsApiService) RemoveCSourceSubscription(ctx context.Context, subscriptionId string) ApiRemoveCSourceSubscriptionRequest {
	return ApiRemoveCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *CSourceSubscriptionsApiService) RemoveCSourceSubscriptionExecute(r ApiRemoveCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CSourceSubscriptionsApiService.RemoveCSourceSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceSubscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRetrieveCSourceSubscriptionsRequest struct {
	ctx context.Context
	ApiService *CSourceSubscriptionsApiService
	limit *int32
}

// Pagination limit
func (r ApiRetrieveCSourceSubscriptionsRequest) Limit(limit int32) ApiRetrieveCSourceSubscriptionsRequest {
	r.limit = &limit
	return r
}

func (r ApiRetrieveCSourceSubscriptionsRequest) Execute() ([]Subscription, *http.Response, error) {
	return r.ApiService.RetrieveCSourceSubscriptionsExecute(r)
}

/*
RetrieveCSourceSubscriptions Method for RetrieveCSourceSubscriptions

Retrieves the context source discovery subscriptions available in an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRetrieveCSourceSubscriptionsRequest
*/
func (a *CSourceSubscriptionsApiService) RetrieveCSourceSubscriptions(ctx context.Context) ApiRetrieveCSourceSubscriptionsRequest {
	return ApiRetrieveCSourceSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *CSourceSubscriptionsApiService) RetrieveCSourceSubscriptionsExecute(r ApiRetrieveCSourceSubscriptionsRequest) ([]Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CSourceSubscriptionsApiService.RetrieveCSourceSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceSubscriptions/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetrieveCSourceSubscriptionsByIdRequest struct {
	ctx context.Context
	ApiService *CSourceSubscriptionsApiService
	subscriptionId string
}

func (r ApiRetrieveCSourceSubscriptionsByIdRequest) Execute() (*Subscription, *http.Response, error) {
	return r.ApiService.RetrieveCSourceSubscriptionsByIdExecute(r)
}

/*
RetrieveCSourceSubscriptionsById Method for RetrieveCSourceSubscriptionsById

Retrieves a specific Subscription from an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription Id
 @return ApiRetrieveCSourceSubscriptionsByIdRequest
*/
func (a *CSourceSubscriptionsApiService) RetrieveCSourceSubscriptionsById(ctx context.Context, subscriptionId string) ApiRetrieveCSourceSubscriptionsByIdRequest {
	return ApiRetrieveCSourceSubscriptionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return Subscription
func (a *CSourceSubscriptionsApiService) RetrieveCSourceSubscriptionsByIdExecute(r ApiRetrieveCSourceSubscriptionsByIdRequest) (*Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CSourceSubscriptionsApiService.RetrieveCSourceSubscriptionsById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceSubscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCSourceSubscriptionRequest struct {
	ctx context.Context
	ApiService *CSourceSubscriptionsApiService
	subscriptionId string
	subscriptionFragment *SubscriptionFragment
}

func (r ApiUpdateCSourceSubscriptionRequest) SubscriptionFragment(subscriptionFragment SubscriptionFragment) ApiUpdateCSourceSubscriptionRequest {
	r.subscriptionFragment = &subscriptionFragment
	return r
}

func (r ApiUpdateCSourceSubscriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateCSourceSubscriptionExecute(r)
}

/*
UpdateCSourceSubscription Method for UpdateCSourceSubscription

Updates a specific context source discovery Subscription within an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Subscription Id
 @return ApiUpdateCSourceSubscriptionRequest
*/
func (a *CSourceSubscriptionsApiService) UpdateCSourceSubscription(ctx context.Context, subscriptionId string) ApiUpdateCSourceSubscriptionRequest {
	return ApiUpdateCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *CSourceSubscriptionsApiService) UpdateCSourceSubscriptionExecute(r ApiUpdateCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CSourceSubscriptionsApiService.UpdateCSourceSubscription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceSubscriptions/{subscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionFragment == nil {
		return nil, reportError("subscriptionFragment is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/ld+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/ld+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.subscriptionFragment
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
