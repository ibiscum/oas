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
)

// checks if the SecondTypeOfObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecondTypeOfObject{}

// SecondTypeOfObject struct for SecondTypeOfObject
type SecondTypeOfObject struct {
	C *string `json:"c,omitempty"`
	D *string `json:"d,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecondTypeOfObject SecondTypeOfObject

// NewSecondTypeOfObject instantiates a new SecondTypeOfObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondTypeOfObject() *SecondTypeOfObject {
	this := SecondTypeOfObject{}
	return &this
}

// NewSecondTypeOfObjectWithDefaults instantiates a new SecondTypeOfObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondTypeOfObjectWithDefaults() *SecondTypeOfObject {
	this := SecondTypeOfObject{}
	return &this
}

// GetC returns the C field value if set, zero value otherwise.
func (o *SecondTypeOfObject) GetC() string {
	if o == nil || IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondTypeOfObject) GetCOk() (*string, bool) {
	if o == nil || IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *SecondTypeOfObject) HasC() bool {
	if o != nil && !IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *SecondTypeOfObject) SetC(v string) {
	o.C = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *SecondTypeOfObject) GetD() string {
	if o == nil || IsNil(o.D) {
		var ret string
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondTypeOfObject) GetDOk() (*string, bool) {
	if o == nil || IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *SecondTypeOfObject) HasD() bool {
	if o != nil && !IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *SecondTypeOfObject) SetD(v string) {
	o.D = &v
}

func (o SecondTypeOfObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecondTypeOfObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	if !IsNil(o.D) {
		toSerialize["d"] = o.D
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecondTypeOfObject) UnmarshalJSON(data []byte) (err error) {
	varSecondTypeOfObject := _SecondTypeOfObject{}

	err = json.Unmarshal(data, &varSecondTypeOfObject)

	if err != nil {
		return err
	}

	*o = SecondTypeOfObject(varSecondTypeOfObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "c")
		delete(additionalProperties, "d")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecondTypeOfObject struct {
	value *SecondTypeOfObject
	isSet bool
}

func (v NullableSecondTypeOfObject) Get() *SecondTypeOfObject {
	return v.value
}

func (v *NullableSecondTypeOfObject) Set(val *SecondTypeOfObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondTypeOfObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondTypeOfObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondTypeOfObject(val *SecondTypeOfObject) *NullableSecondTypeOfObject {
	return &NullableSecondTypeOfObject{value: val, isSet: true}
}

func (v NullableSecondTypeOfObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondTypeOfObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


