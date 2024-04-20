# Andalucia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Registration** | Pointer to [**AndaluciaRegistration**](AndaluciaRegistration.md) |  | [optional] 
**Exemption** | Pointer to [**AndaluciaExemption**](AndaluciaExemption.md) |  | [optional] 

## Methods

### NewAndalucia

`func NewAndalucia(listingId int32, regulatoryBody string, ) *Andalucia`

NewAndalucia instantiates a new Andalucia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndaluciaWithDefaults

`func NewAndaluciaWithDefaults() *Andalucia`

NewAndaluciaWithDefaults instantiates a new Andalucia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *Andalucia) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *Andalucia) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *Andalucia) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *Andalucia) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *Andalucia) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *Andalucia) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetRegistration

`func (o *Andalucia) GetRegistration() AndaluciaRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Andalucia) GetRegistrationOk() (*AndaluciaRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Andalucia) SetRegistration(v AndaluciaRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *Andalucia) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetExemption

`func (o *Andalucia) GetExemption() AndaluciaExemption`

GetExemption returns the Exemption field if non-nil, zero value otherwise.

### GetExemptionOk

`func (o *Andalucia) GetExemptionOk() (*AndaluciaExemption, bool)`

GetExemptionOk returns a tuple with the Exemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemption

`func (o *Andalucia) SetExemption(v AndaluciaExemption)`

SetExemption sets Exemption field to given value.

### HasExemption

`func (o *Andalucia) HasExemption() bool`

HasExemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


