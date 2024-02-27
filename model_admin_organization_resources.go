/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
)

// checks if the AdminOrganizationResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminOrganizationResources{}

// AdminOrganizationResources struct for AdminOrganizationResources
type AdminOrganizationResources struct {
	OrganizationId *int32 `json:"organization_id,omitempty"`
	Environments []AdminEnvrionmentResources `json:"environments,omitempty"`
}

// NewAdminOrganizationResources instantiates a new AdminOrganizationResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminOrganizationResources() *AdminOrganizationResources {
	this := AdminOrganizationResources{}
	return &this
}

// NewAdminOrganizationResourcesWithDefaults instantiates a new AdminOrganizationResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminOrganizationResourcesWithDefaults() *AdminOrganizationResources {
	this := AdminOrganizationResources{}
	return &this
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *AdminOrganizationResources) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminOrganizationResources) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AdminOrganizationResources) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *AdminOrganizationResources) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *AdminOrganizationResources) GetEnvironments() []AdminEnvrionmentResources {
	if o == nil || IsNil(o.Environments) {
		var ret []AdminEnvrionmentResources
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminOrganizationResources) GetEnvironmentsOk() ([]AdminEnvrionmentResources, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *AdminOrganizationResources) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []AdminEnvrionmentResources and assigns it to the Environments field.
func (o *AdminOrganizationResources) SetEnvironments(v []AdminEnvrionmentResources) {
	o.Environments = v
}

func (o AdminOrganizationResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminOrganizationResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	return toSerialize, nil
}

type NullableAdminOrganizationResources struct {
	value *AdminOrganizationResources
	isSet bool
}

func (v NullableAdminOrganizationResources) Get() *AdminOrganizationResources {
	return v.value
}

func (v *NullableAdminOrganizationResources) Set(val *AdminOrganizationResources) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminOrganizationResources) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminOrganizationResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminOrganizationResources(val *AdminOrganizationResources) *NullableAdminOrganizationResources {
	return &NullableAdminOrganizationResources{value: val, isSet: true}
}

func (v NullableAdminOrganizationResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminOrganizationResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


