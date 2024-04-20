# StreamsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**UserData** | Pointer to **string** |  | [optional] 

## Methods

### NewStreamsPostRequest

`func NewStreamsPostRequest() *StreamsPostRequest`

NewStreamsPostRequest instantiates a new StreamsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsPostRequestWithDefaults

`func NewStreamsPostRequestWithDefaults() *StreamsPostRequest`

NewStreamsPostRequestWithDefaults instantiates a new StreamsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *StreamsPostRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamsPostRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamsPostRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StreamsPostRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUserData

`func (o *StreamsPostRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *StreamsPostRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *StreamsPostRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *StreamsPostRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


