# Chicago

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **int32** |  | 
**RegulatoryBody** | **string** |  | 
**Registration** | [**ChicagoRegistration**](ChicagoRegistration.md) |  | 

## Methods

### NewChicago

`func NewChicago(listingId int32, regulatoryBody string, registration ChicagoRegistration, ) *Chicago`

NewChicago instantiates a new Chicago object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChicagoWithDefaults

`func NewChicagoWithDefaults() *Chicago`

NewChicagoWithDefaults instantiates a new Chicago object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *Chicago) GetListingId() int32`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *Chicago) GetListingIdOk() (*int32, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *Chicago) SetListingId(v int32)`

SetListingId sets ListingId field to given value.


### GetRegulatoryBody

`func (o *Chicago) GetRegulatoryBody() string`

GetRegulatoryBody returns the RegulatoryBody field if non-nil, zero value otherwise.

### GetRegulatoryBodyOk

`func (o *Chicago) GetRegulatoryBodyOk() (*string, bool)`

GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryBody

`func (o *Chicago) SetRegulatoryBody(v string)`

SetRegulatoryBody sets RegulatoryBody field to given value.


### GetRegistration

`func (o *Chicago) GetRegistration() ChicagoRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Chicago) GetRegistrationOk() (*ChicagoRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Chicago) SetRegistration(v ChicagoRegistration)`

SetRegistration sets Registration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


