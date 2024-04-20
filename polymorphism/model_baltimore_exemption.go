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

// checks if the BaltimoreExemption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaltimoreExemption{}

// BaltimoreExemption struct for BaltimoreExemption
type BaltimoreExemption struct {
	ExemptionReason string `json:"exemption_reason"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	AttestationExemptionClaim bool `json:"attestation_exemption_claim"`
	AdditionalProperties map[string]interface{}
}

type _BaltimoreExemption BaltimoreExemption

// NewBaltimoreExemption instantiates a new BaltimoreExemption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaltimoreExemption(exemptionReason string, fullName string, email string, attestationExemptionClaim bool) *BaltimoreExemption {
	this := BaltimoreExemption{}
	this.ExemptionReason = exemptionReason
	this.FullName = fullName
	this.Email = email
	this.AttestationExemptionClaim = attestationExemptionClaim
	return &this
}

// NewBaltimoreExemptionWithDefaults instantiates a new BaltimoreExemption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaltimoreExemptionWithDefaults() *BaltimoreExemption {
	this := BaltimoreExemption{}
	return &this
}

// GetExemptionReason returns the ExemptionReason field value
func (o *BaltimoreExemption) GetExemptionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExemptionReason
}

// GetExemptionReasonOk returns a tuple with the ExemptionReason field value
// and a boolean to check if the value has been set.
func (o *BaltimoreExemption) GetExemptionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExemptionReason, true
}

// SetExemptionReason sets field value
func (o *BaltimoreExemption) SetExemptionReason(v string) {
	o.ExemptionReason = v
}

// GetFullName returns the FullName field value
func (o *BaltimoreExemption) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *BaltimoreExemption) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *BaltimoreExemption) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *BaltimoreExemption) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BaltimoreExemption) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BaltimoreExemption) SetEmail(v string) {
	o.Email = v
}

// GetAttestationExemptionClaim returns the AttestationExemptionClaim field value
func (o *BaltimoreExemption) GetAttestationExemptionClaim() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AttestationExemptionClaim
}

// GetAttestationExemptionClaimOk returns a tuple with the AttestationExemptionClaim field value
// and a boolean to check if the value has been set.
func (o *BaltimoreExemption) GetAttestationExemptionClaimOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttestationExemptionClaim, true
}

// SetAttestationExemptionClaim sets field value
func (o *BaltimoreExemption) SetAttestationExemptionClaim(v bool) {
	o.AttestationExemptionClaim = v
}

func (o BaltimoreExemption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaltimoreExemption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exemption_reason"] = o.ExemptionReason
	toSerialize["full_name"] = o.FullName
	toSerialize["email"] = o.Email
	toSerialize["attestation_exemption_claim"] = o.AttestationExemptionClaim

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaltimoreExemption) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exemption_reason",
		"full_name",
		"email",
		"attestation_exemption_claim",
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

	varBaltimoreExemption := _BaltimoreExemption{}

	err = json.Unmarshal(data, &varBaltimoreExemption)

	if err != nil {
		return err
	}

	*o = BaltimoreExemption(varBaltimoreExemption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exemption_reason")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "attestation_exemption_claim")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaltimoreExemption struct {
	value *BaltimoreExemption
	isSet bool
}

func (v NullableBaltimoreExemption) Get() *BaltimoreExemption {
	return v.value
}

func (v *NullableBaltimoreExemption) Set(val *BaltimoreExemption) {
	v.value = val
	v.isSet = true
}

func (v NullableBaltimoreExemption) IsSet() bool {
	return v.isSet
}

func (v *NullableBaltimoreExemption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaltimoreExemption(val *BaltimoreExemption) *NullableBaltimoreExemption {
	return &NullableBaltimoreExemption{value: val, isSet: true}
}

func (v NullableBaltimoreExemption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaltimoreExemption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

