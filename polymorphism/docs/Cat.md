# Cat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetType** | **string** |  | 
**Hunts** | Pointer to **bool** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 

## Methods

### NewCat

`func NewCat(petType string, ) *Cat`

NewCat instantiates a new Cat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatWithDefaults

`func NewCatWithDefaults() *Cat`

NewCatWithDefaults instantiates a new Cat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPetType

`func (o *Cat) GetPetType() string`

GetPetType returns the PetType field if non-nil, zero value otherwise.

### GetPetTypeOk

`func (o *Cat) GetPetTypeOk() (*string, bool)`

GetPetTypeOk returns a tuple with the PetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetType

`func (o *Cat) SetPetType(v string)`

SetPetType sets PetType field to given value.


### GetHunts

`func (o *Cat) GetHunts() bool`

GetHunts returns the Hunts field if non-nil, zero value otherwise.

### GetHuntsOk

`func (o *Cat) GetHuntsOk() (*bool, bool)`

GetHuntsOk returns a tuple with the Hunts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunts

`func (o *Cat) SetHunts(v bool)`

SetHunts sets Hunts field to given value.

### HasHunts

`func (o *Cat) HasHunts() bool`

HasHunts returns a boolean if a field has been set.

### GetAge

`func (o *Cat) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Cat) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Cat) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *Cat) HasAge() bool`

HasAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


