# GetPetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PetType** | **string** |  | 
**Bark** | Pointer to **bool** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 

## Methods

### NewGetPetsRequest

`func NewGetPetsRequest(petType string, ) *GetPetsRequest`

NewGetPetsRequest instantiates a new GetPetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPetsRequestWithDefaults

`func NewGetPetsRequestWithDefaults() *GetPetsRequest`

NewGetPetsRequestWithDefaults instantiates a new GetPetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPetType

`func (o *GetPetsRequest) GetPetType() string`

GetPetType returns the PetType field if non-nil, zero value otherwise.

### GetPetTypeOk

`func (o *GetPetsRequest) GetPetTypeOk() (*string, bool)`

GetPetTypeOk returns a tuple with the PetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPetType

`func (o *GetPetsRequest) SetPetType(v string)`

SetPetType sets PetType field to given value.


### GetBark

`func (o *GetPetsRequest) GetBark() bool`

GetBark returns the Bark field if non-nil, zero value otherwise.

### GetBarkOk

`func (o *GetPetsRequest) GetBarkOk() (*bool, bool)`

GetBarkOk returns a tuple with the Bark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBark

`func (o *GetPetsRequest) SetBark(v bool)`

SetBark sets Bark field to given value.

### HasBark

`func (o *GetPetsRequest) HasBark() bool`

HasBark returns a boolean if a field has been set.

### GetBreed

`func (o *GetPetsRequest) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *GetPetsRequest) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *GetPetsRequest) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *GetPetsRequest) HasBreed() bool`

HasBreed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


