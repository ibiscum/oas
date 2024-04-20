# OneOfComplexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Registration** | [**ChicagoRegistration**](ChicagoRegistration.md) |  | 
**Exemption** | Pointer to [**CataloniaExemption**](CataloniaExemption.md) |  | [optional] 
**Affiliation** | [**CubaAffiliation**](CubaAffiliation.md) |  | 
**Categorization** | [**DenmarkCategorization**](DenmarkCategorization.md) |  | 

## Methods

### NewOneOfComplexRequest

`func NewOneOfComplexRequest(listingId int32, regulatoryBody string, registration ChicagoRegistration, affiliation CubaAffiliation, categorization DenmarkCategorization, ) *OneOfComplexRequest`

NewOneOfComplexRequest instantiates a new OneOfComplexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneOfComplexRequestWithDefaults

`func NewOneOfComplexRequestWithDefaults() *OneOfComplexRequest`

NewOneOfComplexRequestWithDefaults instantiates a new OneOfComplexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *OneOfComplexRequest) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *OneOfComplexRequest) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *OneOfComplexRequest) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *OneOfComplexRequest) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *OneOfComplexRequest) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *OneOfComplexRequest) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetRegistration

`func (o *OneOfComplexRequest) GetRegistration() ChicagoRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *OneOfComplexRequest) GetRegistrationOk() (*ChicagoRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *OneOfComplexRequest) SetRegistration(v ChicagoRegistration)`

SetRegistration sets Registration field to given value.


### GetExemption

`func (o *OneOfComplexRequest) GetExemption() CataloniaExemption`

GetExemption returns the Exemption field if non-nil, zero value otherwise.

### GetExemptionOk

`func (o *OneOfComplexRequest) GetExemptionOk() (*CataloniaExemption, bool)`

GetExemptionOk returns a tuple with the Exemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemption

`func (o *OneOfComplexRequest) SetExemption(v CataloniaExemption)`

SetExemption sets Exemption field to given value.

### HasExemption

`func (o *OneOfComplexRequest) HasExemption() bool`

HasExemption returns a boolean if a field has been set.

### GetAffiliation

`func (o *OneOfComplexRequest) GetAffiliation() CubaAffiliation`

GetAffiliation returns the Affiliation field if non-nil, zero value otherwise.

### GetAffiliationOk

`func (o *OneOfComplexRequest) GetAffiliationOk() (*CubaAffiliation, bool)`

GetAffiliationOk returns a tuple with the Affiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliation

`func (o *OneOfComplexRequest) SetAffiliation(v CubaAffiliation)`

SetAffiliation sets Affiliation field to given value.


### GetCategorization

`func (o *OneOfComplexRequest) GetCategorization() DenmarkCategorization`

GetCategorization returns the Categorization field if non-nil, zero value otherwise.

### GetCategorizationOk

`func (o *OneOfComplexRequest) GetCategorizationOk() (*DenmarkCategorization, bool)`

GetCategorizationOk returns a tuple with the Categorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorization

`func (o *OneOfComplexRequest) SetCategorization(v DenmarkCategorization)`

SetCategorization sets Categorization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


