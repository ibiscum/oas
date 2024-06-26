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

// AnyOfWithTypeRequestEvent struct for AnyOfWithTypeRequestEvent
type AnyOfWithTypeRequestEvent struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnyOfWithTypeRequestEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AnyOfWithTypeRequestEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnyOfWithTypeRequestEvent) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnyOfWithTypeRequestEvent struct {
	value *AnyOfWithTypeRequestEvent
	isSet bool
}

func (v NullableAnyOfWithTypeRequestEvent) Get() *AnyOfWithTypeRequestEvent {
	return v.value
}

func (v *NullableAnyOfWithTypeRequestEvent) Set(val *AnyOfWithTypeRequestEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyOfWithTypeRequestEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyOfWithTypeRequestEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyOfWithTypeRequestEvent(val *AnyOfWithTypeRequestEvent) *NullableAnyOfWithTypeRequestEvent {
	return &NullableAnyOfWithTypeRequestEvent{value: val, isSet: true}
}

func (v NullableAnyOfWithTypeRequestEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnyOfWithTypeRequestEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


