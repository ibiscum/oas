# Dog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetType** | **string** |  | 
**Bark** | Pointer to **bool** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 

## Methods

### NewDog

`func NewDog(petType string, ) *Dog`

NewDog instantiates a new Dog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDogWithDefaults

`func NewDogWithDefaults() *Dog`

NewDogWithDefaults instantiates a new Dog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPetType

`func (o *Dog) GetPetType() string`

GetPetType returns the PetType field if non-nil, zero value otherwise.

### GetPetTypeOk

`func (o *Dog) GetPetTypeOk() (*string, bool)`

GetPetTypeOk returns a tuple with the PetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetType

`func (o *Dog) SetPetType(v string)`

SetPetType sets PetType field to given value.


### GetBark

`func (o *Dog) GetBark() bool`

GetBark returns the Bark field if non-nil, zero value otherwise.

### GetBarkOk

`func (o *Dog) GetBarkOk() (*bool, bool)`

GetBarkOk returns a tuple with the Bark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBark

`func (o *Dog) SetBark(v bool)`

SetBark sets Bark field to given value.

### HasBark

`func (o *Dog) HasBark() bool`

HasBark returns a boolean if a field has been set.

### GetBreed

`func (o *Dog) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *Dog) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *Dog) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *Dog) HasBreed() bool`

HasBreed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


