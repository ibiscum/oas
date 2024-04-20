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

// checks if the PostAnythingSimple200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAnythingSimple200Response{}

// PostAnythingSimple200Response struct for PostAnythingSimple200Response
type PostAnythingSimple200Response struct {
	NestedObject *FlatObject `json:"nestedObject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAnythingSimple200Response PostAnythingSimple200Response

// NewPostAnythingSimple200Response instantiates a new PostAnythingSimple200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAnythingSimple200Response() *PostAnythingSimple200Response {
	this := PostAnythingSimple200Response{}
	return &this
}

// NewPostAnythingSimple200ResponseWithDefaults instantiates a new PostAnythingSimple200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAnythingSimple200ResponseWithDefaults() *PostAnythingSimple200Response {
	this := PostAnythingSimple200Response{}
	return &this
}

// GetNestedObject returns the NestedObject field value if set, zero value otherwise.
func (o *PostAnythingSimple200Response) GetNestedObject() FlatObject {
	if o == nil || IsNil(o.NestedObject) {
		var ret FlatObject
		return ret
	}
	return *o.NestedObject
}

// GetNestedObjectOk returns a tuple with the NestedObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAnythingSimple200Response) GetNestedObjectOk() (*FlatObject, bool) {
	if o == nil || IsNil(o.NestedObject) {
		return nil, false
	}
	return o.NestedObject, true
}

// HasNestedObject returns a boolean if a field has been set.
func (o *PostAnythingSimple200Response) HasNestedObject() bool {
	if o != nil && !IsNil(o.NestedObject) {
		return true
	}

	return false
}

// SetNestedObject gets a reference to the given FlatObject and assigns it to the NestedObject field.
func (o *PostAnythingSimple200Response) SetNestedObject(v FlatObject) {
	o.NestedObject = &v
}

func (o PostAnythingSimple200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAnythingSimple200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NestedObject) {
		toSerialize["nestedObject"] = o.NestedObject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAnythingSimple200Response) UnmarshalJSON(data []byte) (err error) {
	varPostAnythingSimple200Response := _PostAnythingSimple200Response{}

	err = json.Unmarshal(data, &varPostAnythingSimple200Response)

	if err != nil {
		return err
	}

	*o = PostAnythingSimple200Response(varPostAnythingSimple200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "nestedObject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAnythingSimple200Response struct {
	value *PostAnythingSimple200Response
	isSet bool
}

func (v NullablePostAnythingSimple200Response) Get() *PostAnythingSimple200Response {
	return v.value
}

func (v *NullablePostAnythingSimple200Response) Set(val *PostAnythingSimple200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAnythingSimple200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAnythingSimple200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAnythingSimple200Response(val *PostAnythingSimple200Response) *NullablePostAnythingSimple200Response {
	return &NullablePostAnythingSimple200Response{value: val, isSet: true}
}

func (v NullablePostAnythingSimple200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAnythingSimple200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

