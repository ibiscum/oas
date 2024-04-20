# Pet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetType** | **string** |  | 

## Methods

### NewPet

`func NewPet(petType string, ) *Pet`

NewPet instantiates a new Pet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPetWithDefaults

`func NewPetWithDefaults() *Pet`

NewPetWithDefaults instantiates a new Pet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPetType

`func (o *Pet) GetPetType() string`

GetPetType returns the PetType field if non-nil, zero value otherwise.

### GetPetTypeOk

`func (o *Pet) GetPetTypeOk() (*string, bool)`

GetPetTypeOk returns a tuple with the PetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetType

`func (o *Pet) SetPetType(v string)`

SetPetType sets PetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


