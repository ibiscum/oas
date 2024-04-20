# ErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**Inner** | Pointer to [**ErrorMessage**](ErrorMessage.md) |  | [optional] 
**CanBeRetried** | Pointer to **string** |  | [optional] 
**DetailedErrorCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewErrorMessage

`func NewErrorMessage() *ErrorMessage`

NewErrorMessage instantiates a new ErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageWithDefaults

`func NewErrorMessageWithDefaults() *ErrorMessage`

NewErrorMessageWithDefaults instantiates a new ErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *ErrorMessage) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ErrorMessage) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ErrorMessage) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ErrorMessage) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *ErrorMessage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorMessage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorMessage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorMessage) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ErrorMessage) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ErrorMessage) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetInner

`func (o *ErrorMessage) GetInner() ErrorMessage`

GetInner returns the Inner field if non-nil, zero value otherwise.

### GetInnerOk

`func (o *ErrorMessage) GetInnerOk() (*ErrorMessage, bool)`

GetInnerOk returns a tuple with the Inner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInner

`func (o *ErrorMessage) SetInner(v ErrorMessage)`

SetInner sets Inner field to given value.

### HasInner

`func (o *ErrorMessage) HasInner() bool`

HasInner returns a boolean if a field has been set.

### GetCanBeRetried

`func (o *ErrorMessage) GetCanBeRetried() string`

GetCanBeRetried returns the CanBeRetried field if non-nil, zero value otherwise.

### GetCanBeRetriedOk

`func (o *ErrorMessage) GetCanBeRetriedOk() (*string, bool)`

GetCanBeRetriedOk returns a tuple with the CanBeRetried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeRetried

`func (o *ErrorMessage) SetCanBeRetried(v string)`

SetCanBeRetried sets CanBeRetried field to given value.

### HasCanBeRetried

`func (o *ErrorMessage) HasCanBeRetried() bool`

HasCanBeRetried returns a boolean if a field has been set.

### GetDetailedErrorCode

`func (o *ErrorMessage) GetDetailedErrorCode() int32`

GetDetailedErrorCode returns the DetailedErrorCode field if non-nil, zero value otherwise.

### GetDetailedErrorCodeOk

`func (o *ErrorMessage) GetDetailedErrorCodeOk() (*int32, bool)`

GetDetailedErrorCodeOk returns a tuple with the DetailedErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedErrorCode

`func (o *ErrorMessage) SetDetailedErrorCode(v int32)`

SetDetailedErrorCode sets DetailedErrorCode field to given value.

### HasDetailedErrorCode

`func (o *ErrorMessage) HasDetailedErrorCode() bool`

HasDetailedErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


