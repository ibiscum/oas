# Catalonia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Registration** | Pointer to [**CataloniaRegistration**](CataloniaRegistration.md) |  | [optional] 
**Exemption** | Pointer to [**CataloniaExemption**](CataloniaExemption.md) |  | [optional] 

## Methods

### NewCatalonia

`func NewCatalonia(listingId int32, regulatoryBody string, ) *Catalonia`

NewCatalonia instantiates a new Catalonia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCataloniaWithDefaults

`func NewCataloniaWithDefaults() *Catalonia`

NewCataloniaWithDefaults instantiates a new Catalonia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *Catalonia) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *Catalonia) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *Catalonia) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *Catalonia) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *Catalonia) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *Catalonia) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetRegistration

`func (o *Catalonia) GetRegistration() CataloniaRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Catalonia) GetRegistrationOk() (*CataloniaRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Catalonia) SetRegistration(v CataloniaRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *Catalonia) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetExemption

`func (o *Catalonia) GetExemption() CataloniaExemption`

GetExemption returns the Exemption field if non-nil, zero value otherwise.

### GetExemptionOk

`func (o *Catalonia) GetExemptionOk() (*CataloniaExemption, bool)`

GetExemptionOk returns a tuple with the Exemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemption

`func (o *Catalonia) SetExemption(v CataloniaExemption)`

SetExemption sets Exemption field to given value.

### HasExemption

`func (o *Catalonia) HasExemption() bool`

HasExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


