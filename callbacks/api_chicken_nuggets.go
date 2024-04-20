/*
Callback Example

API definition with callback

API version: 1.0.0
Contact: a-team@goarmy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package callbacks

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ChickenNuggetsAPIService ChickenNuggetsAPI service
type ChickenNuggetsAPIService service

type ApiPostStreamsRequest struct {
	ctx context.Context
	ApiService *ChickenNuggetsAPIService
	callbackUrl *string
}

// the location where data will be sent. Must be network accessible by the source server
func (r ApiPostStreamsRequest) CallbackUrl(callbackUrl string) ApiPostStreamsRequest {
	r.callbackUrl = &callbackUrl
	return r
}

func (r ApiPostStreamsRequest) Execute() (*PostStreams201Response, *http.Response, error) {
	return r.ApiService.PostStreamsExecute(r)
}

/*
PostStreams Method for PostStreams

subscribes a client to receive out-of-band data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostStreamsRequest
*/
func (a *ChickenNuggetsAPIService) PostStreams(ctx context.Context) ApiPostStreamsRequest {
	return ApiPostStreamsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostStreams201Response
func (a *ChickenNuggetsAPIService) PostStreamsExecute(r ApiPostStreamsRequest) (*PostStreams201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostStreams201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChickenNuggetsAPIService.PostStreams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/streams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.callbackUrl == nil {
		return localVarReturnValue, nil, reportError("callbackUrl is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "callbackUrl", r.callbackUrl, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
