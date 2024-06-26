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
	"reflect"
	"strings"
)

// checks if the Cat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cat{}

// Cat struct for Cat
type Cat struct {
	Pet
	PetType string `json:"pet_type"`
	Hunts *bool `json:"hunts,omitempty"`
	Age *int32 `json:"age,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Cat Cat

// NewCat instantiates a new Cat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCat(petType string) *Cat {
	this := Cat{}
	this.PetType = petType
	return &this
}

// NewCatWithDefaults instantiates a new Cat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatWithDefaults() *Cat {
	this := Cat{}
	return &this
}

// GetPetType returns the PetType field value
func (o *Cat) GetPetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PetType
}

// GetPetTypeOk returns a tuple with the PetType field value
// and a boolean to check if the value has been set.
func (o *Cat) GetPetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PetType, true
}

// SetPetType sets field value
func (o *Cat) SetPetType(v string) {
	o.PetType = v
}

// GetHunts returns the Hunts field value if set, zero value otherwise.
func (o *Cat) GetHunts() bool {
	if o == nil || IsNil(o.Hunts) {
		var ret bool
		return ret
	}
	return *o.Hunts
}

// GetHuntsOk returns a tuple with the Hunts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cat) GetHuntsOk() (*bool, bool) {
	if o == nil || IsNil(o.Hunts) {
		return nil, false
	}
	return o.Hunts, true
}

// HasHunts returns a boolean if a field has been set.
func (o *Cat) HasHunts() bool {
	if o != nil && !IsNil(o.Hunts) {
		return true
	}

	return false
}

// SetHunts gets a reference to the given bool and assigns it to the Hunts field.
func (o *Cat) SetHunts(v bool) {
	o.Hunts = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *Cat) GetAge() int32 {
	if o == nil || IsNil(o.Age) {
		var ret int32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cat) GetAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *Cat) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given int32 and assigns it to the Age field.
func (o *Cat) SetAge(v int32) {
	o.Age = &v
}

func (o Cat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPet, errPet := json.Marshal(o.Pet)
	if errPet != nil {
		return map[string]interface{}{}, errPet
	}
	errPet = json.Unmarshal([]byte(serializedPet), &toSerialize)
	if errPet != nil {
		return map[string]interface{}{}, errPet
	}
	toSerialize["pet_type"] = o.PetType
	if !IsNil(o.Hunts) {
		toSerialize["hunts"] = o.Hunts
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cat) UnmarshalJSON(data []byte) (err error) {
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

	type CatWithoutEmbeddedStruct struct {
		PetType string `json:"pet_type"`
		Hunts *bool `json:"hunts,omitempty"`
		Age *int32 `json:"age,omitempty"`
	}

	varCatWithoutEmbeddedStruct := CatWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varCatWithoutEmbeddedStruct)
	if err == nil {
		varCat := _Cat{}
		varCat.PetType = varCatWithoutEmbeddedStruct.PetType
		varCat.Hunts = varCatWithoutEmbeddedStruct.Hunts
		varCat.Age = varCatWithoutEmbeddedStruct.Age
		*o = Cat(varCat)
	} else {
		return err
	}

	varCat := _Cat{}

	err = json.Unmarshal(data, &varCat)
	if err == nil {
		o.Pet = varCat.Pet
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pet_type")
		delete(additionalProperties, "hunts")
		delete(additionalProperties, "age")

		// remove fields from embedded structs
		reflectPet := reflect.ValueOf(o.Pet)
		for i := 0; i < reflectPet.Type().NumField(); i++ {
			t := reflectPet.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCat struct {
	value *Cat
	isSet bool
}

func (v NullableCat) Get() *Cat {
	return v.value
}

func (v *NullableCat) Set(val *Cat) {
	v.value = val
	v.isSet = true
}

func (v NullableCat) IsSet() bool {
	return v.isSet
}

func (v *NullableCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCat(val *Cat) *NullableCat {
	return &NullableCat{value: val, isSet: true}
}

func (v NullableCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


