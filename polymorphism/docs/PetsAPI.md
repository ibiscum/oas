# \PetsAPI

All URIs are relative to *http://localhost:8001*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnyOfWithType**](PetsAPI.md#AnyOfWithType) | **Post** /any-of-with-type | anyOf With Type
[**AnythingAllOfObject**](PetsAPI.md#AnythingAllOfObject) | **Post** /anything/all-of-object | allOf with listed objects
[**AnythingAnyOfObject**](PetsAPI.md#AnythingAnyOfObject) | **Post** /anything/any-of-object | anyOf object
[**AnythingAnyOfPrimitive**](PetsAPI.md#AnythingAnyOfPrimitive) | **Post** /anything/any-of-primitive | anyOf primitive
[**AnythingNestedOneOfObjectRef**](PetsAPI.md#AnythingNestedOneOfObjectRef) | **Post** /anything/nested-one-of-object-ref | nested oneOf object with $ref pointers
[**AnythingNestedOneOfObjectWithNestedOneOf**](PetsAPI.md#AnythingNestedOneOfObjectWithNestedOneOf) | **Post** /anything/nested-one-of-object-with-nested-one-of | 
[**AnythingNestedOneOfRef**](PetsAPI.md#AnythingNestedOneOfRef) | **Post** /anything/nested-one-of-ref | 
[**AnythingOneOfObject**](PetsAPI.md#AnythingOneOfObject) | **Post** /anything/one-of-object | oneOf object
[**AnythingOneOfObjectRef**](PetsAPI.md#AnythingOneOfObjectRef) | **Post** /anything/one-of-object-ref | oneOf object with $ref pointers
[**AnythingOneOfPrimitive**](PetsAPI.md#AnythingOneOfPrimitive) | **Post** /anything/one-of-primitive | oneOf primitive
[**GetPets**](PetsAPI.md#GetPets) | **Patch** /pets | oneOf request with a nested allOf
[**OneOfComplex**](PetsAPI.md#OneOfComplex) | **Post** /one-of-complex | oneOf object with a complex schema
[**OneOfWithType**](PetsAPI.md#OneOfWithType) | **Post** /one-of-with-type | One Of With Type



## AnyOfWithType

> AnyOfWithType(ctx).AnyOfWithTypeRequest(anyOfWithTypeRequest).Execute()

anyOf With Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anyOfWithTypeRequest := *openapiclient.NewAnyOfWithTypeRequest() // AnyOfWithTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnyOfWithType(context.Background()).AnyOfWithTypeRequest(anyOfWithTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnyOfWithType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnyOfWithTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anyOfWithTypeRequest** | [**AnyOfWithTypeRequest**](AnyOfWithTypeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingAllOfObject

> AnythingAllOfObject(ctx).AnythingAllOfObjectRequest(anythingAllOfObjectRequest).Execute()

allOf with listed objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingAllOfObjectRequest := *openapiclient.NewAnythingAllOfObjectRequest() // AnythingAllOfObjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingAllOfObject(context.Background()).AnythingAllOfObjectRequest(anythingAllOfObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingAllOfObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingAllOfObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingAllOfObjectRequest** | [**AnythingAllOfObjectRequest**](AnythingAllOfObjectRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingAnyOfObject

> AnythingAnyOfObject(ctx).AnythingAnyOfObjectRequest(anythingAnyOfObjectRequest).Execute()

anyOf object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingAnyOfObjectRequest := *openapiclient.NewAnythingAnyOfObjectRequest() // AnythingAnyOfObjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingAnyOfObject(context.Background()).AnythingAnyOfObjectRequest(anythingAnyOfObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingAnyOfObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingAnyOfObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingAnyOfObjectRequest** | [**AnythingAnyOfObjectRequest**](AnythingAnyOfObjectRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingAnyOfPrimitive

> AnythingAnyOfPrimitive(ctx).AnythingAnyOfPrimitiveRequest(anythingAnyOfPrimitiveRequest).Execute()

anyOf primitive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingAnyOfPrimitiveRequest := *openapiclient.NewAnythingAnyOfPrimitiveRequest() // AnythingAnyOfPrimitiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingAnyOfPrimitive(context.Background()).AnythingAnyOfPrimitiveRequest(anythingAnyOfPrimitiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingAnyOfPrimitive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingAnyOfPrimitiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingAnyOfPrimitiveRequest** | [**AnythingAnyOfPrimitiveRequest**](AnythingAnyOfPrimitiveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingNestedOneOfObjectRef

> AnythingNestedOneOfObjectRef(ctx).AnythingNestedOneOfObjectRefRequest(anythingNestedOneOfObjectRefRequest).Execute()

nested oneOf object with $ref pointers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingNestedOneOfObjectRefRequest := *openapiclient.NewAnythingNestedOneOfObjectRefRequest() // AnythingNestedOneOfObjectRefRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingNestedOneOfObjectRef(context.Background()).AnythingNestedOneOfObjectRefRequest(anythingNestedOneOfObjectRefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingNestedOneOfObjectRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingNestedOneOfObjectRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingNestedOneOfObjectRefRequest** | [**AnythingNestedOneOfObjectRefRequest**](AnythingNestedOneOfObjectRefRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingNestedOneOfObjectWithNestedOneOf

> AnythingNestedOneOfObjectWithNestedOneOf(ctx).AnythingNestedOneOfObjectWithNestedOneOfRequest(anythingNestedOneOfObjectWithNestedOneOfRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingNestedOneOfObjectWithNestedOneOfRequest := *openapiclient.NewAnythingNestedOneOfObjectWithNestedOneOfRequest() // AnythingNestedOneOfObjectWithNestedOneOfRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingNestedOneOfObjectWithNestedOneOf(context.Background()).AnythingNestedOneOfObjectWithNestedOneOfRequest(anythingNestedOneOfObjectWithNestedOneOfRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingNestedOneOfObjectWithNestedOneOf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingNestedOneOfObjectWithNestedOneOfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingNestedOneOfObjectWithNestedOneOfRequest** | [**AnythingNestedOneOfObjectWithNestedOneOfRequest**](AnythingNestedOneOfObjectWithNestedOneOfRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingNestedOneOfRef

> AnythingNestedOneOfRef(ctx).AnythingNestedOneOfRefRequest(anythingNestedOneOfRefRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingNestedOneOfRefRequest := *openapiclient.NewAnythingNestedOneOfRefRequest() // AnythingNestedOneOfRefRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingNestedOneOfRef(context.Background()).AnythingNestedOneOfRefRequest(anythingNestedOneOfRefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingNestedOneOfRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingNestedOneOfRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingNestedOneOfRefRequest** | [**AnythingNestedOneOfRefRequest**](AnythingNestedOneOfRefRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingOneOfObject

> AnythingOneOfObject(ctx).AnythingOneOfObjectRequest(anythingOneOfObjectRequest).Execute()

oneOf object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingOneOfObjectRequest := openapiclient.anything_one_of_object_request{FirstTypeOfObject: openapiclient.NewFirstTypeOfObject()} // AnythingOneOfObjectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingOneOfObject(context.Background()).AnythingOneOfObjectRequest(anythingOneOfObjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingOneOfObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingOneOfObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingOneOfObjectRequest** | [**AnythingOneOfObjectRequest**](AnythingOneOfObjectRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingOneOfObjectRef

> AnythingOneOfObjectRef(ctx).AnythingOneOfObjectRefRequest(anythingOneOfObjectRefRequest).Execute()

oneOf object with $ref pointers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingOneOfObjectRefRequest := openapiclient.anything_one_of_object_ref_request{Object1: openapiclient.NewObject1()} // AnythingOneOfObjectRefRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingOneOfObjectRef(context.Background()).AnythingOneOfObjectRefRequest(anythingOneOfObjectRefRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingOneOfObjectRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingOneOfObjectRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingOneOfObjectRefRequest** | [**AnythingOneOfObjectRefRequest**](AnythingOneOfObjectRefRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnythingOneOfPrimitive

> AnythingOneOfPrimitive(ctx).AnythingOneOfPrimitiveRequest(anythingOneOfPrimitiveRequest).Execute()

oneOf primitive



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	anythingOneOfPrimitiveRequest := openapiclient.anything_one_of_primitive_request{Int32: new(int32)} // AnythingOneOfPrimitiveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.AnythingOneOfPrimitive(context.Background()).AnythingOneOfPrimitiveRequest(anythingOneOfPrimitiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.AnythingOneOfPrimitive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnythingOneOfPrimitiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anythingOneOfPrimitiveRequest** | [**AnythingOneOfPrimitiveRequest**](AnythingOneOfPrimitiveRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPets

> GetPets(ctx).GetPetsRequest(getPetsRequest).Execute()

oneOf request with a nested allOf



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	getPetsRequest := openapiclient.get_pets_request{Cat: openapiclient.NewCat("PetType_example")} // GetPetsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.GetPets(context.Background()).GetPetsRequest(getPetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.GetPets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getPetsRequest** | [**GetPetsRequest**](GetPetsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OneOfComplex

> OneOfComplex(ctx).OneOfComplexRequest(oneOfComplexRequest).Execute()

oneOf object with a complex schema



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	oneOfComplexRequest := openapiclient.one_of_complex_request{Andalucia: openapiclient.NewAndalucia(int32(123), "RegulatoryBody_example")} // OneOfComplexRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.OneOfComplex(context.Background()).OneOfComplexRequest(oneOfComplexRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.OneOfComplex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOneOfComplexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oneOfComplexRequest** | [**OneOfComplexRequest**](OneOfComplexRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OneOfWithType

> OneOfWithType(ctx).OneOfWithTypeRequest(oneOfWithTypeRequest).Execute()

One Of With Type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/polymorphism"
)

func main() {
	oneOfWithTypeRequest := *openapiclient.NewOneOfWithTypeRequest() // OneOfWithTypeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PetsAPI.OneOfWithType(context.Background()).OneOfWithTypeRequest(oneOfWithTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PetsAPI.OneOfWithType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOneOfWithTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oneOfWithTypeRequest** | [**OneOfWithTypeRequest**](OneOfWithTypeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

