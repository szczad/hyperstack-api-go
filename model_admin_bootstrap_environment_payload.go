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

// checks if the AdminBootstrapEnvironmentPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminBootstrapEnvironmentPayload{}

// AdminBootstrapEnvironmentPayload struct for AdminBootstrapEnvironmentPayload
type AdminBootstrapEnvironmentPayload struct {
	OrgId int32 `json:"org_id"`
}

type _AdminBootstrapEnvironmentPayload AdminBootstrapEnvironmentPayload

// NewAdminBootstrapEnvironmentPayload instantiates a new AdminBootstrapEnvironmentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminBootstrapEnvironmentPayload(orgId int32) *AdminBootstrapEnvironmentPayload {
	this := AdminBootstrapEnvironmentPayload{}
	this.OrgId = orgId
	return &this
}

// NewAdminBootstrapEnvironmentPayloadWithDefaults instantiates a new AdminBootstrapEnvironmentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminBootstrapEnvironmentPayloadWithDefaults() *AdminBootstrapEnvironmentPayload {
	this := AdminBootstrapEnvironmentPayload{}
	return &this
}

// GetOrgId returns the OrgId field value
func (o *AdminBootstrapEnvironmentPayload) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *AdminBootstrapEnvironmentPayload) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *AdminBootstrapEnvironmentPayload) SetOrgId(v int32) {
	o.OrgId = v
}

func (o AdminBootstrapEnvironmentPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminBootstrapEnvironmentPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["org_id"] = o.OrgId
	return toSerialize, nil
}

func (o *AdminBootstrapEnvironmentPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"org_id",
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

	varAdminBootstrapEnvironmentPayload := _AdminBootstrapEnvironmentPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdminBootstrapEnvironmentPayload)

	if err != nil {
		return err
	}

	*o = AdminBootstrapEnvironmentPayload(varAdminBootstrapEnvironmentPayload)

	return err
}

type NullableAdminBootstrapEnvironmentPayload struct {
	value *AdminBootstrapEnvironmentPayload
	isSet bool
}

func (v NullableAdminBootstrapEnvironmentPayload) Get() *AdminBootstrapEnvironmentPayload {
	return v.value
}

func (v *NullableAdminBootstrapEnvironmentPayload) Set(val *AdminBootstrapEnvironmentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminBootstrapEnvironmentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminBootstrapEnvironmentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminBootstrapEnvironmentPayload(val *AdminBootstrapEnvironmentPayload) *NullableAdminBootstrapEnvironmentPayload {
	return &NullableAdminBootstrapEnvironmentPayload{value: val, isSet: true}
}

func (v NullableAdminBootstrapEnvironmentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminBootstrapEnvironmentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


