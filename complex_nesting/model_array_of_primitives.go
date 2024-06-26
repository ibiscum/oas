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

// checks if the ArrayOfPrimitives type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArrayOfPrimitives{}

// ArrayOfPrimitives Array of primitives
type ArrayOfPrimitives struct {
}

// NewArrayOfPrimitives instantiates a new ArrayOfPrimitives object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArrayOfPrimitives() *ArrayOfPrimitives {
	this := ArrayOfPrimitives{}
	return &this
}

// NewArrayOfPrimitivesWithDefaults instantiates a new ArrayOfPrimitives object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArrayOfPrimitivesWithDefaults() *ArrayOfPrimitives {
	this := ArrayOfPrimitives{}
	return &this
}

func (o ArrayOfPrimitives) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArrayOfPrimitives) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return toSerialize, nil
}

func (o *ArrayOfPrimitives) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableArrayOfPrimitives struct {
	value ArrayOfPrimitives
	isSet bool
}

func (v NullableArrayOfPrimitives) Get() ArrayOfPrimitives {
	return v.value
}

func (v *NullableArrayOfPrimitives) Set(val ArrayOfPrimitives) {
	v.value = val
	v.isSet = true
}

func (v NullableArrayOfPrimitives) IsSet() bool {
	return v.isSet
}

func (v *NullableArrayOfPrimitives) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArrayOfPrimitives(val ArrayOfPrimitives) *NullableArrayOfPrimitives {
	return &NullableArrayOfPrimitives{value: val, isSet: true}
}

func (v NullableArrayOfPrimitives) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArrayOfPrimitives) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


