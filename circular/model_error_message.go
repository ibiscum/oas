/*
API definition with a circular $ref

API definition with a circular $ref

API version: 1.0.0
Contact: a-team@goarmy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package circular

import (
	"encoding/json"
)

// checks if the ErrorMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessage{}

// ErrorMessage struct for ErrorMessage
type ErrorMessage struct {
	StatusCode *int32 `json:"statusCode,omitempty"`
	Error NullableString `json:"error,omitempty"`
	Inner *ErrorMessage `json:"inner,omitempty"`
	CanBeRetried *string `json:"canBeRetried,omitempty"`
	DetailedErrorCode *int32 `json:"detailedErrorCode,omitempty"`
}

// NewErrorMessage instantiates a new ErrorMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessage() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// NewErrorMessageWithDefaults instantiates a new ErrorMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageWithDefaults() *ErrorMessage {
	this := ErrorMessage{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ErrorMessage) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ErrorMessage) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ErrorMessage) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorMessage) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorMessage) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ErrorMessage) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ErrorMessage) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ErrorMessage) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ErrorMessage) UnsetError() {
	o.Error.Unset()
}

// GetInner returns the Inner field value if set, zero value otherwise.
func (o *ErrorMessage) GetInner() ErrorMessage {
	if o == nil || IsNil(o.Inner) {
		var ret ErrorMessage
		return ret
	}
	return *o.Inner
}

// GetInnerOk returns a tuple with the Inner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetInnerOk() (*ErrorMessage, bool) {
	if o == nil || IsNil(o.Inner) {
		return nil, false
	}
	return o.Inner, true
}

// HasInner returns a boolean if a field has been set.
func (o *ErrorMessage) HasInner() bool {
	if o != nil && !IsNil(o.Inner) {
		return true
	}

	return false
}

// SetInner gets a reference to the given ErrorMessage and assigns it to the Inner field.
func (o *ErrorMessage) SetInner(v ErrorMessage) {
	o.Inner = &v
}

// GetCanBeRetried returns the CanBeRetried field value if set, zero value otherwise.
func (o *ErrorMessage) GetCanBeRetried() string {
	if o == nil || IsNil(o.CanBeRetried) {
		var ret string
		return ret
	}
	return *o.CanBeRetried
}

// GetCanBeRetriedOk returns a tuple with the CanBeRetried field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetCanBeRetriedOk() (*string, bool) {
	if o == nil || IsNil(o.CanBeRetried) {
		return nil, false
	}
	return o.CanBeRetried, true
}

// HasCanBeRetried returns a boolean if a field has been set.
func (o *ErrorMessage) HasCanBeRetried() bool {
	if o != nil && !IsNil(o.CanBeRetried) {
		return true
	}

	return false
}

// SetCanBeRetried gets a reference to the given string and assigns it to the CanBeRetried field.
func (o *ErrorMessage) SetCanBeRetried(v string) {
	o.CanBeRetried = &v
}

// GetDetailedErrorCode returns the DetailedErrorCode field value if set, zero value otherwise.
func (o *ErrorMessage) GetDetailedErrorCode() int32 {
	if o == nil || IsNil(o.DetailedErrorCode) {
		var ret int32
		return ret
	}
	return *o.DetailedErrorCode
}

// GetDetailedErrorCodeOk returns a tuple with the DetailedErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessage) GetDetailedErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.DetailedErrorCode) {
		return nil, false
	}
	return o.DetailedErrorCode, true
}

// HasDetailedErrorCode returns a boolean if a field has been set.
func (o *ErrorMessage) HasDetailedErrorCode() bool {
	if o != nil && !IsNil(o.DetailedErrorCode) {
		return true
	}

	return false
}

// SetDetailedErrorCode gets a reference to the given int32 and assigns it to the DetailedErrorCode field.
func (o *ErrorMessage) SetDetailedErrorCode(v int32) {
	o.DetailedErrorCode = &v
}

func (o ErrorMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !IsNil(o.Inner) {
		toSerialize["inner"] = o.Inner
	}
	if !IsNil(o.CanBeRetried) {
		toSerialize["canBeRetried"] = o.CanBeRetried
	}
	if !IsNil(o.DetailedErrorCode) {
		toSerialize["detailedErrorCode"] = o.DetailedErrorCode
	}
	return toSerialize, nil
}

type NullableErrorMessage struct {
	value *ErrorMessage
	isSet bool
}

func (v NullableErrorMessage) Get() *ErrorMessage {
	return v.value
}

func (v *NullableErrorMessage) Set(val *ErrorMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessage(val *ErrorMessage) *NullableErrorMessage {
	return &NullableErrorMessage{value: val, isSet: true}
}

func (v NullableErrorMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


