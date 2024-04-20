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

// checks if the Pet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pet{}

// Pet struct for Pet
type Pet struct {
	PetType string `json:"pet_type"`
	AdditionalProperties map[string]interface{}
}

type _Pet Pet

// NewPet instantiates a new Pet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPet(petType string) *Pet {
	this := Pet{}
	this.PetType = petType
	return &this
}

// NewPetWithDefaults instantiates a new Pet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPetWithDefaults() *Pet {
	this := Pet{}
	return &this
}

// GetPetType returns the PetType field value
func (o *Pet) GetPetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PetType
}

// GetPetTypeOk returns a tuple with the PetType field value
// and a boolean to check if the value has been set.
func (o *Pet) GetPetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PetType, true
}

// SetPetType sets field value
func (o *Pet) SetPetType(v string) {
	o.PetType = v
}

func (o Pet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pet_type"] = o.PetType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Pet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pet_type",
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

	varPet := _Pet{}

	err = json.Unmarshal(data, &varPet)

	if err != nil {
		return err
	}

	*o = Pet(varPet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pet_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePet struct {
	value *Pet
	isSet bool
}

func (v NullablePet) Get() *Pet {
	return v.value
}

func (v *NullablePet) Set(val *Pet) {
	v.value = val
	v.isSet = true
}

func (v NullablePet) IsSet() bool {
	return v.isSet
}

func (v *NullablePet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePet(val *Pet) *NullablePet {
	return &NullablePet{value: val, isSet: true}
}

func (v NullablePet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


