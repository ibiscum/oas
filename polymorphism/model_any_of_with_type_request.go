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

// checks if the AnyOfWithTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnyOfWithTypeRequest{}

// AnyOfWithTypeRequest struct for AnyOfWithTypeRequest
type AnyOfWithTypeRequest struct {
	Event *AnyOfWithTypeRequestEvent `json:"event,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AnyOfWithTypeRequest AnyOfWithTypeRequest

// NewAnyOfWithTypeRequest instantiates a new AnyOfWithTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnyOfWithTypeRequest() *AnyOfWithTypeRequest {
	this := AnyOfWithTypeRequest{}
	return &this
}

// NewAnyOfWithTypeRequestWithDefaults instantiates a new AnyOfWithTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyOfWithTypeRequestWithDefaults() *AnyOfWithTypeRequest {
	this := AnyOfWithTypeRequest{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AnyOfWithTypeRequest) GetEvent() AnyOfWithTypeRequestEvent {
	if o == nil || IsNil(o.Event) {
		var ret AnyOfWithTypeRequestEvent
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyOfWithTypeRequest) GetEventOk() (*AnyOfWithTypeRequestEvent, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AnyOfWithTypeRequest) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given AnyOfWithTypeRequestEvent and assigns it to the Event field.
func (o *AnyOfWithTypeRequest) SetEvent(v AnyOfWithTypeRequestEvent) {
	o.Event = &v
}

func (o AnyOfWithTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnyOfWithTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AnyOfWithTypeRequest) UnmarshalJSON(data []byte) (err error) {
	varAnyOfWithTypeRequest := _AnyOfWithTypeRequest{}

	err = json.Unmarshal(data, &varAnyOfWithTypeRequest)

	if err != nil {
		return err
	}

	*o = AnyOfWithTypeRequest(varAnyOfWithTypeRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAnyOfWithTypeRequest struct {
	value *AnyOfWithTypeRequest
	isSet bool
}

func (v NullableAnyOfWithTypeRequest) Get() *AnyOfWithTypeRequest {
	return v.value
}

func (v *NullableAnyOfWithTypeRequest) Set(val *AnyOfWithTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyOfWithTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyOfWithTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyOfWithTypeRequest(val *AnyOfWithTypeRequest) *NullableAnyOfWithTypeRequest {
	return &NullableAnyOfWithTypeRequest{value: val, isSet: true}
}

func (v NullableAnyOfWithTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnyOfWithTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


