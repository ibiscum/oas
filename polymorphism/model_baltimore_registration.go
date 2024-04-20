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

// checks if the BaltimoreRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaltimoreRegistration{}

// BaltimoreRegistration struct for BaltimoreRegistration
type BaltimoreRegistration struct {
	PermitNumber string `json:"permit_number"`
	ExpirationDate string `json:"expiration_date"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	ListingAddress string `json:"listing_address"`
	AttestationExistingRegistration bool `json:"attestation_existing_registration"`
	AdditionalProperties map[string]interface{}
}

type _BaltimoreRegistration BaltimoreRegistration

// NewBaltimoreRegistration instantiates a new BaltimoreRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaltimoreRegistration(permitNumber string, expirationDate string, fullName string, email string, listingAddress string, attestationExistingRegistration bool) *BaltimoreRegistration {
	this := BaltimoreRegistration{}
	this.PermitNumber = permitNumber
	this.ExpirationDate = expirationDate
	this.FullName = fullName
	this.Email = email
	this.ListingAddress = listingAddress
	this.AttestationExistingRegistration = attestationExistingRegistration
	return &this
}

// NewBaltimoreRegistrationWithDefaults instantiates a new BaltimoreRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaltimoreRegistrationWithDefaults() *BaltimoreRegistration {
	this := BaltimoreRegistration{}
	return &this
}

// GetPermitNumber returns the PermitNumber field value
func (o *BaltimoreRegistration) GetPermitNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermitNumber
}

// GetPermitNumberOk returns a tuple with the PermitNumber field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetPermitNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermitNumber, true
}

// SetPermitNumber sets field value
func (o *BaltimoreRegistration) SetPermitNumber(v string) {
	o.PermitNumber = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *BaltimoreRegistration) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *BaltimoreRegistration) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetFullName returns the FullName field value
func (o *BaltimoreRegistration) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *BaltimoreRegistration) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *BaltimoreRegistration) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BaltimoreRegistration) SetEmail(v string) {
	o.Email = v
}

// GetListingAddress returns the ListingAddress field value
func (o *BaltimoreRegistration) GetListingAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListingAddress
}

// GetListingAddressOk returns a tuple with the ListingAddress field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetListingAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListingAddress, true
}

// SetListingAddress sets field value
func (o *BaltimoreRegistration) SetListingAddress(v string) {
	o.ListingAddress = v
}

// GetAttestationExistingRegistration returns the AttestationExistingRegistration field value
func (o *BaltimoreRegistration) GetAttestationExistingRegistration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AttestationExistingRegistration
}

// GetAttestationExistingRegistrationOk returns a tuple with the AttestationExistingRegistration field value
// and a boolean to check if the value has been set.
func (o *BaltimoreRegistration) GetAttestationExistingRegistrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttestationExistingRegistration, true
}

// SetAttestationExistingRegistration sets field value
func (o *BaltimoreRegistration) SetAttestationExistingRegistration(v bool) {
	o.AttestationExistingRegistration = v
}

func (o BaltimoreRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaltimoreRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permit_number"] = o.PermitNumber
	toSerialize["expiration_date"] = o.ExpirationDate
	toSerialize["full_name"] = o.FullName
	toSerialize["email"] = o.Email
	toSerialize["listing_address"] = o.ListingAddress
	toSerialize["attestation_existing_registration"] = o.AttestationExistingRegistration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaltimoreRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permit_number",
		"expiration_date",
		"full_name",
		"email",
		"listing_address",
		"attestation_existing_registration",
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

	varBaltimoreRegistration := _BaltimoreRegistration{}

	err = json.Unmarshal(data, &varBaltimoreRegistration)

	if err != nil {
		return err
	}

	*o = BaltimoreRegistration(varBaltimoreRegistration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "permit_number")
		delete(additionalProperties, "expiration_date")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "listing_address")
		delete(additionalProperties, "attestation_existing_registration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaltimoreRegistration struct {
	value *BaltimoreRegistration
	isSet bool
}

func (v NullableBaltimoreRegistration) Get() *BaltimoreRegistration {
	return v.value
}

func (v *NullableBaltimoreRegistration) Set(val *BaltimoreRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableBaltimoreRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableBaltimoreRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaltimoreRegistration(val *BaltimoreRegistration) *NullableBaltimoreRegistration {
	return &NullableBaltimoreRegistration{value: val, isSet: true}
}

func (v NullableBaltimoreRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaltimoreRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


