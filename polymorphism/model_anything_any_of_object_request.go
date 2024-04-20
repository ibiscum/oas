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

// AnythingAnyOfObjectRequest struct for AnythingAnyOfObjectRequest
type AnythingAnyOfObjectRequest struct {
	FirstTypeOfObject *FirstTypeOfObject
	SecondTypeOfObject *SecondTypeOfObject
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnythingAnyOfObjectRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FirstTypeOfObject
	err = json.Unmarshal(data, &dst.FirstTypeOfObject);
	if err == nil {
		jsonFirstTypeOfObject, _ := json.Marshal(dst.FirstTypeOfObject)
		if string(jsonFirstTypeOfObject) == "{}" { // empty struct
			dst.FirstTypeOfObject = nil
		} else {
			return nil // data stored in dst.FirstTypeOfObject, return on the first match
		}
	} else {
		dst.FirstTypeOfObject = nil
	}

	// try to unmarshal JSON data into SecondTypeOfObject
	err = json.Unmarshal(data, &dst.SecondTypeOfObject);
	if err == nil {
		jsonSecondTypeOfObject, _ := json.Marshal(dst.SecondTypeOfObject)
		if string(jsonSecondTypeOfObject) == "{}" { // empty struct
			dst.SecondTypeOfObject = nil
		} else {
			return nil // data stored in dst.SecondTypeOfObject, return on the first match
		}
	} else {
		dst.SecondTypeOfObject = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AnythingAnyOfObjectRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnythingAnyOfObjectRequest) MarshalJSON() ([]byte, error) {
	if src.FirstTypeOfObject != nil {
		return json.Marshal(&src.FirstTypeOfObject)
	}

	if src.SecondTypeOfObject != nil {
		return json.Marshal(&src.SecondTypeOfObject)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnythingAnyOfObjectRequest struct {
	value *AnythingAnyOfObjectRequest
	isSet bool
}

func (v NullableAnythingAnyOfObjectRequest) Get() *AnythingAnyOfObjectRequest {
	return v.value
}

func (v *NullableAnythingAnyOfObjectRequest) Set(val *AnythingAnyOfObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnythingAnyOfObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnythingAnyOfObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnythingAnyOfObjectRequest(val *AnythingAnyOfObjectRequest) *NullableAnythingAnyOfObjectRequest {
	return &NullableAnythingAnyOfObjectRequest{value: val, isSet: true}
}

func (v NullableAnythingAnyOfObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnythingAnyOfObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

