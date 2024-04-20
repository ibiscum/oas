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
	"fmt"
)

// ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata - struct for ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata
type ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata struct {
	FlatObject *FlatObject
	ObjectWithArray *ObjectWithArray
	MapmapOfStringAny *map[string]interface{}
}

// FlatObjectAsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata is a convenience function that returns FlatObject wrapped in ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata
func FlatObjectAsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata(v *FlatObject) ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata {
	return ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata{
		FlatObject: v,
	}
}

// ObjectWithArrayAsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata is a convenience function that returns ObjectWithArray wrapped in ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata
func ObjectWithArrayAsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata(v *ObjectWithArray) ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata {
	return ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata{
		ObjectWithArray: v,
	}
}

// map[string]interface{}AsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata is a convenience function that returns map[string]interface{} wrapped in ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata
func MapmapOfStringAnyAsObjectOfAdditionalPropertiesObjectPolymorphism1Metadata(v *map[string]interface{}) ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata {
	return ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	match := 0
	// try to unmarshal data into FlatObject
	err = newStrictDecoder(data).Decode(&dst.FlatObject)
	if err == nil {
		jsonFlatObject, _ := json.Marshal(dst.FlatObject)
		if string(jsonFlatObject) == "{}" { // empty struct
			dst.FlatObject = nil
		} else {
			match++
		}
	} else {
		dst.FlatObject = nil
	}

	// try to unmarshal data into ObjectWithArray
	err = newStrictDecoder(data).Decode(&dst.ObjectWithArray)
	if err == nil {
		jsonObjectWithArray, _ := json.Marshal(dst.ObjectWithArray)
		if string(jsonObjectWithArray) == "{}" { // empty struct
			dst.ObjectWithArray = nil
		} else {
			match++
		}
	} else {
		dst.ObjectWithArray = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			match++
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FlatObject = nil
		dst.ObjectWithArray = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) MarshalJSON() ([]byte, error) {
	if src.FlatObject != nil {
		return json.Marshal(&src.FlatObject)
	}

	if src.ObjectWithArray != nil {
		return json.Marshal(&src.ObjectWithArray)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FlatObject != nil {
		return obj.FlatObject
	}

	if obj.ObjectWithArray != nil {
		return obj.ObjectWithArray
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata struct {
	value *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata
	isSet bool
}

func (v NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) Get() *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata {
	return v.value
}

func (v *NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) Set(val *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata(val *ObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) *NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata {
	return &NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata{value: val, isSet: true}
}

func (v NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectOfAdditionalPropertiesObjectPolymorphism1Metadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


