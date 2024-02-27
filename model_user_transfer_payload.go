/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserTransferPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserTransferPayload{}

// UserTransferPayload struct for UserTransferPayload
type UserTransferPayload struct {
	OrgId int32 `json:"org_id"`
	Role string `json:"role"`
}

type _UserTransferPayload UserTransferPayload

// NewUserTransferPayload instantiates a new UserTransferPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTransferPayload(orgId int32, role string) *UserTransferPayload {
	this := UserTransferPayload{}
	this.OrgId = orgId
	this.Role = role
	return &this
}

// NewUserTransferPayloadWithDefaults instantiates a new UserTransferPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTransferPayloadWithDefaults() *UserTransferPayload {
	this := UserTransferPayload{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *UserTransferPayload) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UserTransferPayload) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *UserTransferPayload) SetOrgId(v int32) {
	o.OrgId = v
}

// GetRole returns the Role field value
func (o *UserTransferPayload) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserTransferPayload) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserTransferPayload) SetRole(v string) {
	o.Role = v
}

func (o UserTransferPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserTransferPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *UserTransferPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
		"role",
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

	varUserTransferPayload := _UserTransferPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserTransferPayload)

	if err != nil {
		return err
	}

	*o = UserTransferPayload(varUserTransferPayload)

	return err
}

type NullableUserTransferPayload struct {
	value *UserTransferPayload
	isSet bool
}

func (v NullableUserTransferPayload) Get() *UserTransferPayload {
	return v.value
}

func (v *NullableUserTransferPayload) Set(val *UserTransferPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTransferPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTransferPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTransferPayload(val *UserTransferPayload) *NullableUserTransferPayload {
	return &NullableUserTransferPayload{value: val, isSet: true}
}

func (v NullableUserTransferPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTransferPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

