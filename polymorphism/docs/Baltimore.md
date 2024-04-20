# Baltimore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Registration** | Pointer to [**BaltimoreRegistration**](BaltimoreRegistration.md) |  | [optional] 
**Exemption** | Pointer to [**BaltimoreExemption**](BaltimoreExemption.md) |  | [optional] 

## Methods

### NewBaltimore

`func NewBaltimore(listingId int32, regulatoryBody string, ) *Baltimore`

NewBaltimore instantiates a new Baltimore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaltimoreWithDefaults

`func NewBaltimoreWithDefaults() *Baltimore`

NewBaltimoreWithDefaults instantiates a new Baltimore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *Baltimore) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *Baltimore) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *Baltimore) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *Baltimore) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *Baltimore) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *Baltimore) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetRegistration

`func (o *Baltimore) GetRegistration() BaltimoreRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Baltimore) GetRegistrationOk() (*BaltimoreRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Baltimore) SetRegistration(v BaltimoreRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *Baltimore) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetExemption

`func (o *Baltimore) GetExemption() BaltimoreExemption`

GetExemption returns the Exemption field if non-nil, zero value otherwise.

### GetExemptionOk

`func (o *Baltimore) GetExemptionOk() (*BaltimoreExemption, bool)`

GetExemptionOk returns a tuple with the Exemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemption

`func (o *Baltimore) SetExemption(v BaltimoreExemption)`

SetExemption sets Exemption field to given value.

### HasExemption

`func (o *Baltimore) HasExemption() bool`

HasExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


