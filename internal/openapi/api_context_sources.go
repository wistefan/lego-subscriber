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


// ContextSourcesApiService ContextSourcesApi service
type ContextSourcesApiService service

type ApiCreateCSourceSubscriptionRequest struct {
	ctx context.Context
	ApiService *ContextSourcesApiService
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
func (a *ContextSourcesApiService) CreateCSourceSubscription(ctx context.Context) ApiCreateCSourceSubscriptionRequest {
	return ApiCreateCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ContextSourcesApiService) CreateCSourceSubscriptionExecute(r ApiCreateCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.CreateCSourceSubscription")
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

type ApiQueryCsourcesRequest struct {
	ctx context.Context
	ApiService *ContextSourcesApiService
	id *string
	idPattern *string
	type_ *string
	attrs *string
	q *string
	georel *string
	geometry *string
	coordinates *string
	geoproperty *string
	limit *int32
}

// Comma separated list of URIs to be retrieved
func (r ApiQueryCsourcesRequest) Id(id string) ApiQueryCsourcesRequest {
	r.id = &id
	return r
}

// Regular expression that must be matched by Entity ids
func (r ApiQueryCsourcesRequest) IdPattern(idPattern string) ApiQueryCsourcesRequest {
	r.idPattern = &idPattern
	return r
}

// Comma separated list of Entity type names to be retrieved
func (r ApiQueryCsourcesRequest) Type_(type_ string) ApiQueryCsourcesRequest {
	r.type_ = &type_
	return r
}

// Comma separated list of attribute names (properties or relationships) to be retrieved
func (r ApiQueryCsourcesRequest) Attrs(attrs string) ApiQueryCsourcesRequest {
	r.attrs = &attrs
	return r
}

// Query
func (r ApiQueryCsourcesRequest) Q(q string) ApiQueryCsourcesRequest {
	r.q = &q
	return r
}

// Geo-relationship
func (r ApiQueryCsourcesRequest) Georel(georel string) ApiQueryCsourcesRequest {
	r.georel = &georel
	return r
}

// Geometry
func (r ApiQueryCsourcesRequest) Geometry(geometry string) ApiQueryCsourcesRequest {
	r.geometry = &geometry
	return r
}

// Coordinates serialized as a string
func (r ApiQueryCsourcesRequest) Coordinates(coordinates string) ApiQueryCsourcesRequest {
	r.coordinates = &coordinates
	return r
}

// The name of the property that contains the geo-spatial data that will be used to resolve the geoquery
func (r ApiQueryCsourcesRequest) Geoproperty(geoproperty string) ApiQueryCsourcesRequest {
	r.geoproperty = &geoproperty
	return r
}

// Pagination limit
func (r ApiQueryCsourcesRequest) Limit(limit int32) ApiQueryCsourcesRequest {
	r.limit = &limit
	return r
}

func (r ApiQueryCsourcesRequest) Execute() ([]ContextSourceRegistration, *http.Response, error) {
	return r.ApiService.QueryCsourcesExecute(r)
}

/*
QueryCsources Method for QueryCsources

Retrieve a set of context sources which matches a specific query from an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryCsourcesRequest
*/
func (a *ContextSourcesApiService) QueryCsources(ctx context.Context) ApiQueryCsourcesRequest {
	return ApiQueryCsourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ContextSourceRegistration
func (a *ContextSourcesApiService) QueryCsourcesExecute(r ApiQueryCsourcesRequest) ([]ContextSourceRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ContextSourceRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.QueryCsources")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceRegistrations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.idPattern != nil {
		localVarQueryParams.Add("idPattern", parameterToString(*r.idPattern, ""))
	}
	if r.type_ != nil {
		localVarQueryParams.Add("type", parameterToString(*r.type_, ""))
	}
	if r.attrs != nil {
		localVarQueryParams.Add("attrs", parameterToString(*r.attrs, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	if r.georel != nil {
		localVarQueryParams.Add("georel", parameterToString(*r.georel, ""))
	}
	if r.geometry != nil {
		localVarQueryParams.Add("geometry", parameterToString(*r.geometry, ""))
	}
	if r.coordinates != nil {
		localVarQueryParams.Add("coordinates", parameterToString(*r.coordinates, ""))
	}
	if r.geoproperty != nil {
		localVarQueryParams.Add("geoproperty", parameterToString(*r.geoproperty, ""))
	}
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

type ApiRegisterCsourceRequest struct {
	ctx context.Context
	ApiService *ContextSourcesApiService
	contextSourceRegistration *ContextSourceRegistration
}

func (r ApiRegisterCsourceRequest) ContextSourceRegistration(contextSourceRegistration ContextSourceRegistration) ApiRegisterCsourceRequest {
	r.contextSourceRegistration = &contextSourceRegistration
	return r
}

func (r ApiRegisterCsourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RegisterCsourceExecute(r)
}

/*
RegisterCsource Method for RegisterCsource

Registers a new context source within an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterCsourceRequest
*/
func (a *ContextSourcesApiService) RegisterCsource(ctx context.Context) ApiRegisterCsourceRequest {
	return ApiRegisterCsourceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ContextSourcesApiService) RegisterCsourceExecute(r ApiRegisterCsourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RegisterCsource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceRegistrations/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contextSourceRegistration == nil {
		return nil, reportError("contextSourceRegistration is required and must be specified")
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
	localVarPostBody = r.contextSourceRegistration
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
	ApiService *ContextSourcesApiService
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
func (a *ContextSourcesApiService) RemoveCSourceSubscription(ctx context.Context, subscriptionId string) ApiRemoveCSourceSubscriptionRequest {
	return ApiRemoveCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *ContextSourcesApiService) RemoveCSourceSubscriptionExecute(r ApiRemoveCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RemoveCSourceSubscription")
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

type ApiRemoveCsourceRequest struct {
	ctx context.Context
	ApiService *ContextSourcesApiService
	registrationId string
}

func (r ApiRemoveCsourceRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveCsourceExecute(r)
}

/*
RemoveCsource Method for RemoveCsource

Removes an specific context source registration within an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrationId Registration Id
 @return ApiRemoveCsourceRequest
*/
func (a *ContextSourcesApiService) RemoveCsource(ctx context.Context, registrationId string) ApiRemoveCsourceRequest {
	return ApiRemoveCsourceRequest{
		ApiService: a,
		ctx: ctx,
		registrationId: registrationId,
	}
}

// Execute executes the request
func (a *ContextSourcesApiService) RemoveCsourceExecute(r ApiRemoveCsourceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RemoveCsource")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceRegistrations/{registrationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registrationId"+"}", url.PathEscape(parameterToString(r.registrationId, "")), -1)

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
	ApiService *ContextSourcesApiService
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
func (a *ContextSourcesApiService) RetrieveCSourceSubscriptions(ctx context.Context) ApiRetrieveCSourceSubscriptionsRequest {
	return ApiRetrieveCSourceSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Subscription
func (a *ContextSourcesApiService) RetrieveCSourceSubscriptionsExecute(r ApiRetrieveCSourceSubscriptionsRequest) ([]Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RetrieveCSourceSubscriptions")
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
	ApiService *ContextSourcesApiService
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
func (a *ContextSourcesApiService) RetrieveCSourceSubscriptionsById(ctx context.Context, subscriptionId string) ApiRetrieveCSourceSubscriptionsByIdRequest {
	return ApiRetrieveCSourceSubscriptionsByIdRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return Subscription
func (a *ContextSourcesApiService) RetrieveCSourceSubscriptionsByIdExecute(r ApiRetrieveCSourceSubscriptionsByIdRequest) (*Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RetrieveCSourceSubscriptionsById")
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

type ApiRetrieveCsourceRequest struct {
	ctx context.Context
	ApiService *ContextSourcesApiService
	registrationId string
}

func (r ApiRetrieveCsourceRequest) Execute() (*ContextSourceRegistration, *http.Response, error) {
	return r.ApiService.RetrieveCsourceExecute(r)
}

/*
RetrieveCsource Method for RetrieveCsource

Retrieves a specific context source registration from an NGSI-LD system

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param registrationId Registration Id
 @return ApiRetrieveCsourceRequest
*/
func (a *ContextSourcesApiService) RetrieveCsource(ctx context.Context, registrationId string) ApiRetrieveCsourceRequest {
	return ApiRetrieveCsourceRequest{
		ApiService: a,
		ctx: ctx,
		registrationId: registrationId,
	}
}

// Execute executes the request
//  @return ContextSourceRegistration
func (a *ContextSourcesApiService) RetrieveCsourceExecute(r ApiRetrieveCsourceRequest) (*ContextSourceRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContextSourceRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.RetrieveCsource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/csourceRegistrations/{registrationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"registrationId"+"}", url.PathEscape(parameterToString(r.registrationId, "")), -1)

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
	ApiService *ContextSourcesApiService
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
func (a *ContextSourcesApiService) UpdateCSourceSubscription(ctx context.Context, subscriptionId string) ApiUpdateCSourceSubscriptionRequest {
	return ApiUpdateCSourceSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
func (a *ContextSourcesApiService) UpdateCSourceSubscriptionExecute(r ApiUpdateCSourceSubscriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContextSourcesApiService.UpdateCSourceSubscription")
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
