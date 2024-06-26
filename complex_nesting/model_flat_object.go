/*
Responses with various schema formats

https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.3.md#schema-object

API version: 1.0
Contact: a-team@goarmy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package complex_nesting

import (
	"encoding/json"
)

// checks if the FlatObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlatObject{}

// FlatObject Simple flat object
type FlatObject struct {
	String *string `json:"string,omitempty"`
	Boolean *bool `json:"boolean,omitempty"`
	Number *float32 `json:"number,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FlatObject FlatObject

// NewFlatObject instantiates a new FlatObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlatObject() *FlatObject {
	this := FlatObject{}
	return &this
}

// NewFlatObjectWithDefaults instantiates a new FlatObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlatObjectWithDefaults() *FlatObject {
	this := FlatObject{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *FlatObject) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatObject) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *FlatObject) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *FlatObject) SetString(v string) {
	o.String = &v
}

// GetBoolean returns the Boolean field value if set, zero value otherwise.
func (o *FlatObject) GetBoolean() bool {
	if o == nil || IsNil(o.Boolean) {
		var ret bool
		return ret
	}
	return *o.Boolean
}

// GetBooleanOk returns a tuple with the Boolean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatObject) GetBooleanOk() (*bool, bool) {
	if o == nil || IsNil(o.Boolean) {
		return nil, false
	}
	return o.Boolean, true
}

// HasBoolean returns a boolean if a field has been set.
func (o *FlatObject) HasBoolean() bool {
	if o != nil && !IsNil(o.Boolean) {
		return true
	}

	return false
}

// SetBoolean gets a reference to the given bool and assigns it to the Boolean field.
func (o *FlatObject) SetBoolean(v bool) {
	o.Boolean = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *FlatObject) GetNumber() float32 {
	if o == nil || IsNil(o.Number) {
		var ret float32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatObject) GetNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *FlatObject) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given float32 and assigns it to the Number field.
func (o *FlatObject) SetNumber(v float32) {
	o.Number = &v
}

func (o FlatObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlatObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !IsNil(o.Boolean) {
		toSerialize["boolean"] = o.Boolean
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FlatObject) UnmarshalJSON(data []byte) (err error) {
	varFlatObject := _FlatObject{}

	err = json.Unmarshal(data, &varFlatObject)

	if err != nil {
		return err
	}

	*o = FlatObject(varFlatObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "string")
		delete(additionalProperties, "boolean")
		delete(additionalProperties, "number")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFlatObject struct {
	value *FlatObject
	isSet bool
}

func (v NullableFlatObject) Get() *FlatObject {
	return v.value
}

func (v *NullableFlatObject) Set(val *FlatObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFlatObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFlatObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlatObject(val *FlatObject) *NullableFlatObject {
	return &NullableFlatObject{value: val, isSet: true}
}

func (v NullableFlatObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlatObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


