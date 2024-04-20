# \ChickenNuggetsAPI

All URIs are relative to *https://httpbin.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostStreams**](ChickenNuggetsAPI.md#PostStreams) | **Post** /streams | 



## PostStreams

> PostStreams201Response PostStreams(ctx).CallbackUrl(callbackUrl).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/circular"
)

func main() {
	callbackUrl := "https://tonys-server.com" // string | the location where data will be sent. Must be network accessible by the source server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PostStreams(context.Background()).CallbackUrl(callbackUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PostStreams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStreams`: PostStreams201Response
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PostStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackUrl** | **string** | the location where data will be sent. Must be network accessible by the source server | 

### Return type

[**PostStreams201Response**](PostStreams201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

