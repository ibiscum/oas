# FlatObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**String** | Pointer to **string** |  | [optional] 
**Boolean** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **float32** |  | [optional] 

## Methods

### NewFlatObject

`func NewFlatObject() *FlatObject`

NewFlatObject instantiates a new FlatObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatObjectWithDefaults

`func NewFlatObjectWithDefaults() *FlatObject`

NewFlatObjectWithDefaults instantiates a new FlatObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetString

`func (o *FlatObject) GetString() string`

GetString returns the String field if non-nil, zero value otherwise.

### GetStringOk

`func (o *FlatObject) GetStringOk() (*string, bool)`

GetStringOk returns a tuple with the String field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetString

`func (o *FlatObject) SetString(v string)`

SetString sets String field to given value.

### HasString

`func (o *FlatObject) HasString() bool`

HasString returns a boolean if a field has been set.

### GetBoolean

`func (o *FlatObject) GetBoolean() bool`

GetBoolean returns the Boolean field if non-nil, zero value otherwise.

### GetBooleanOk

`func (o *FlatObject) GetBooleanOk() (*bool, bool)`

GetBooleanOk returns a tuple with the Boolean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolean

`func (o *FlatObject) SetBoolean(v bool)`

SetBoolean sets Boolean field to given value.

### HasBoolean

`func (o *FlatObject) HasBoolean() bool`

HasBoolean returns a boolean if a field has been set.

### GetNumber

`func (o *FlatObject) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FlatObject) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FlatObject) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FlatObject) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


