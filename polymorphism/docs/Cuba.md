# Cuba

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Affiliation** | [**CubaAffiliation**](CubaAffiliation.md) |  | 

## Methods

### NewCuba

`func NewCuba(listingId int32, regulatoryBody string, affiliation CubaAffiliation, ) *Cuba`

NewCuba instantiates a new Cuba object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCubaWithDefaults

`func NewCubaWithDefaults() *Cuba`

NewCubaWithDefaults instantiates a new Cuba object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *Cuba) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *Cuba) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *Cuba) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *Cuba) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *Cuba) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *Cuba) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetAffiliation

`func (o *Cuba) GetAffiliation() CubaAffiliation`

GetAffiliation returns the Affiliation field if non-nil, zero value otherwise.

### GetAffiliationOk

`func (o *Cuba) GetAffiliationOk() (*CubaAffiliation, bool)`

GetAffiliationOk returns a tuple with the Affiliation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliation

`func (o *Cuba) SetAffiliation(v CubaAffiliation)`

SetAffiliation sets Affiliation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


