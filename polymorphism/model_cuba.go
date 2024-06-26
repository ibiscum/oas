/*
Polymorphism support

https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#schema-object

API version: 1.0.0
Contact: a-team@goarmy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package polymorphism

import (
	"encoding/json"
	"fmt"
)

// checks if the Cuba type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cuba{}

// Cuba struct for Cuba
type Cuba struct {
	ListingId int32 `json:"listing_id"`
	RegulatoryBody string `json:"regulatory_body"`
	Affiliation CubaAffiliation `json:"affiliation"`
	AdditionalProperties map[string]interface{}
}

type _Cuba Cuba

// NewCuba instantiates a new Cuba object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCuba(listingId int32, regulatoryBody string, affiliation CubaAffiliation) *Cuba {
	this := Cuba{}
	this.ListingId = listingId
	this.RegulatoryBody = regulatoryBody
	this.Affiliation = affiliation
	return &this
}

// NewCubaWithDefaults instantiates a new Cuba object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCubaWithDefaults() *Cuba {
	this := Cuba{}
	return &this
}

// GetListingId returns the ListingId field value
func (o *Cuba) GetListingId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value
// and a boolean to check if the value has been set.
func (o *Cuba) GetListingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingId, true
}

// SetListingId sets field value
func (o *Cuba) SetListingId(v int32) {
	o.ListingId = v
}

// GetRegulatoryBody returns the RegulatoryBody field value
func (o *Cuba) GetRegulatoryBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulatoryBody
}

// GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field value
// and a boolean to check if the value has been set.
func (o *Cuba) GetRegulatoryBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulatoryBody, true
}

// SetRegulatoryBody sets field value
func (o *Cuba) SetRegulatoryBody(v string) {
	o.RegulatoryBody = v
}

// GetAffiliation returns the Affiliation field value
func (o *Cuba) GetAffiliation() CubaAffiliation {
	if o == nil {
		var ret CubaAffiliation
		return ret
	}

	return o.Affiliation
}

// GetAffiliationOk returns a tuple with the Affiliation field value
// and a boolean to check if the value has been set.
func (o *Cuba) GetAffiliationOk() (*CubaAffiliation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Affiliation, true
}

// SetAffiliation sets field value
func (o *Cuba) SetAffiliation(v CubaAffiliation) {
	o.Affiliation = v
}

func (o Cuba) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cuba) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listing_id"] = o.ListingId
	toSerialize["regulatory_body"] = o.RegulatoryBody
	toSerialize["affiliation"] = o.Affiliation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cuba) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listing_id",
		"regulatory_body",
		"affiliation",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCuba := _Cuba{}

	err = json.Unmarshal(data, &varCuba)

	if err != nil {
		return err
	}

	*o = Cuba(varCuba)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listing_id")
		delete(additionalProperties, "regulatory_body")
		delete(additionalProperties, "affiliation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCuba struct {
	value *Cuba
	isSet bool
}

func (v NullableCuba) Get() *Cuba {
	return v.value
}

func (v *NullableCuba) Set(val *Cuba) {
	v.value = val
	v.isSet = true
}

func (v NullableCuba) IsSet() bool {
	return v.isSet
}

func (v *NullableCuba) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCuba(val *Cuba) *NullableCuba {
	return &NullableCuba{value: val, isSet: true}
}

func (v NullableCuba) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCuba) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


