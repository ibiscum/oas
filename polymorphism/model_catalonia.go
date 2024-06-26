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

// checks if the Catalonia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Catalonia{}

// Catalonia struct for Catalonia
type Catalonia struct {
	ListingId int32 `json:"listing_id"`
	RegulatoryBody string `json:"regulatory_body"`
	Registration *CataloniaRegistration `json:"registration,omitempty"`
	Exemption *CataloniaExemption `json:"exemption,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Catalonia Catalonia

// NewCatalonia instantiates a new Catalonia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalonia(listingId int32, regulatoryBody string) *Catalonia {
	this := Catalonia{}
	this.ListingId = listingId
	this.RegulatoryBody = regulatoryBody
	return &this
}

// NewCataloniaWithDefaults instantiates a new Catalonia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCataloniaWithDefaults() *Catalonia {
	this := Catalonia{}
	return &this
}

// GetListingId returns the ListingId field value
func (o *Catalonia) GetListingId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ListingId
}

// GetListingIdOk returns a tuple with the ListingId field value
// and a boolean to check if the value has been set.
func (o *Catalonia) GetListingIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingId, true
}

// SetListingId sets field value
func (o *Catalonia) SetListingId(v int32) {
	o.ListingId = v
}

// GetRegulatoryBody returns the RegulatoryBody field value
func (o *Catalonia) GetRegulatoryBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegulatoryBody
}

// GetRegulatoryBodyOk returns a tuple with the RegulatoryBody field value
// and a boolean to check if the value has been set.
func (o *Catalonia) GetRegulatoryBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegulatoryBody, true
}

// SetRegulatoryBody sets field value
func (o *Catalonia) SetRegulatoryBody(v string) {
	o.RegulatoryBody = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *Catalonia) GetRegistration() CataloniaRegistration {
	if o == nil || IsNil(o.Registration) {
		var ret CataloniaRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Catalonia) GetRegistrationOk() (*CataloniaRegistration, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *Catalonia) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given CataloniaRegistration and assigns it to the Registration field.
func (o *Catalonia) SetRegistration(v CataloniaRegistration) {
	o.Registration = &v
}

// GetExemption returns the Exemption field value if set, zero value otherwise.
func (o *Catalonia) GetExemption() CataloniaExemption {
	if o == nil || IsNil(o.Exemption) {
		var ret CataloniaExemption
		return ret
	}
	return *o.Exemption
}

// GetExemptionOk returns a tuple with the Exemption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Catalonia) GetExemptionOk() (*CataloniaExemption, bool) {
	if o == nil || IsNil(o.Exemption) {
		return nil, false
	}
	return o.Exemption, true
}

// HasExemption returns a boolean if a field has been set.
func (o *Catalonia) HasExemption() bool {
	if o != nil && !IsNil(o.Exemption) {
		return true
	}

	return false
}

// SetExemption gets a reference to the given CataloniaExemption and assigns it to the Exemption field.
func (o *Catalonia) SetExemption(v CataloniaExemption) {
	o.Exemption = &v
}

func (o Catalonia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Catalonia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["listing_id"] = o.ListingId
	toSerialize["regulatory_body"] = o.RegulatoryBody
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if !IsNil(o.Exemption) {
		toSerialize["exemption"] = o.Exemption
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Catalonia) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"listing_id",
		"regulatory_body",
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

	varCatalonia := _Catalonia{}

	err = json.Unmarshal(data, &varCatalonia)

	if err != nil {
		return err
	}

	*o = Catalonia(varCatalonia)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "listing_id")
		delete(additionalProperties, "regulatory_body")
		delete(additionalProperties, "registration")
		delete(additionalProperties, "exemption")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatalonia struct {
	value *Catalonia
	isSet bool
}

func (v NullableCatalonia) Get() *Catalonia {
	return v.value
}

func (v *NullableCatalonia) Set(val *Catalonia) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalonia) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalonia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalonia(val *Catalonia) *NullableCatalonia {
	return &NullableCatalonia{value: val, isSet: true}
}

func (v NullableCatalonia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalonia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


