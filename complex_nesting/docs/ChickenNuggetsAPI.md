# \ChickenNuggetsAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAnything**](ChickenNuggetsAPI.md#PostAnything) | **Post** /anything/top-level-array | Simple Array
[**PostAnythingSimple**](ChickenNuggetsAPI.md#PostAnythingSimple) | **Post** /anything/top-level-object/simple | Simple Object
[**PostMultiSchemaOfEverything**](ChickenNuggetsAPI.md#PostMultiSchemaOfEverything) | **Post** /multischema/of-everything | Multischema of Everything
[**PostObjectOfEverything**](ChickenNuggetsAPI.md#PostObjectOfEverything) | **Post** /top-level-object/of-everything | Object of
[**PutAnything**](ChickenNuggetsAPI.md#PutAnything) | **Put** /anything/top-level-array | Array of Everything



## PostAnything

> ArrayOfFlatObjects PostAnything(ctx).ArrayOfFlatObjects(arrayOfFlatObjects).Execute()

Simple Array



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/complex_nesting"
)

func main() {
	arrayOfFlatObjects := *openapiclient.NewArrayOfFlatObjects() // ArrayOfFlatObjects |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PostAnything(context.Background()).ArrayOfFlatObjects(arrayOfFlatObjects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PostAnything``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAnything`: ArrayOfFlatObjects
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PostAnything`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAnythingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **arrayOfFlatObjects** | [**ArrayOfFlatObjects**](ArrayOfFlatObjects.md) |  | 

### Return type

[**ArrayOfFlatObjects**](ArrayOfFlatObjects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAnythingSimple

> PostAnythingSimple200Response PostAnythingSimple(ctx).PostAnythingSimpleRequest(postAnythingSimpleRequest).Execute()

Simple Object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/complex_nesting"
)

func main() {
	postAnythingSimpleRequest := *openapiclient.NewPostAnythingSimpleRequest() // PostAnythingSimpleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PostAnythingSimple(context.Background()).PostAnythingSimpleRequest(postAnythingSimpleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PostAnythingSimple``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAnythingSimple`: PostAnythingSimple200Response
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PostAnythingSimple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAnythingSimpleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postAnythingSimpleRequest** | [**PostAnythingSimpleRequest**](PostAnythingSimpleRequest.md) |  | 

### Return type

[**PostAnythingSimple200Response**](PostAnythingSimple200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMultiSchemaOfEverything

> MultischemaOfEverything PostMultiSchemaOfEverything(ctx).MultischemaOfEverything(multischemaOfEverything).Execute()

Multischema of Everything



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/complex_nesting"
)

func main() {
	multischemaOfEverything := openapiclient.MultischemaOfEverything{ArrayOfFlatObjects: openapiclient.NewArrayOfFlatObjects()} // MultischemaOfEverything |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PostMultiSchemaOfEverything(context.Background()).MultischemaOfEverything(multischemaOfEverything).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PostMultiSchemaOfEverything``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMultiSchemaOfEverything`: MultischemaOfEverything
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PostMultiSchemaOfEverything`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMultiSchemaOfEverythingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multischemaOfEverything** | [**MultischemaOfEverything**](MultischemaOfEverything.md) |  | 

### Return type

[**MultischemaOfEverything**](MultischemaOfEverything.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectOfEverything

> ObjectOfEverything PostObjectOfEverything(ctx).ObjectOfEverything(objectOfEverything).Execute()

Object of



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/complex_nesting"
)

func main() {
	objectOfEverything := *openapiclient.NewObjectOfEverything() // ObjectOfEverything |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PostObjectOfEverything(context.Background()).ObjectOfEverything(objectOfEverything).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PostObjectOfEverything``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectOfEverything`: ObjectOfEverything
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PostObjectOfEverything`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectOfEverythingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectOfEverything** | [**ObjectOfEverything**](ObjectOfEverything.md) |  | 

### Return type

[**ObjectOfEverything**](ObjectOfEverything.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAnything

> []ObjectOfEverything PutAnything(ctx).ObjectOfEverything(objectOfEverything).Execute()

Array of Everything



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/complex_nesting"
)

func main() {
	objectOfEverything := []openapiclient.ObjectOfEverything{*openapiclient.NewObjectOfEverything()} // []ObjectOfEverything |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChickenNuggetsAPI.PutAnything(context.Background()).ObjectOfEverything(objectOfEverything).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChickenNuggetsAPI.PutAnything``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAnything`: []ObjectOfEverything
	fmt.Fprintf(os.Stdout, "Response from `ChickenNuggetsAPI.PutAnything`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAnythingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectOfEverything** | [**[]ObjectOfEverything**](ObjectOfEverything.md) |  | 

### Return type

[**[]ObjectOfEverything**](ObjectOfEverything.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

