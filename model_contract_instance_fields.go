/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"time"
)

// checks if the ContractInstanceFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractInstanceFields{}

// ContractInstanceFields struct for ContractInstanceFields
type ContractInstanceFields struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	FlavorName *string `json:"flavor_name,omitempty"`
	GpuCount *int32 `json:"gpu_count,omitempty"`
}

// NewContractInstanceFields instantiates a new ContractInstanceFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInstanceFields() *ContractInstanceFields {
	this := ContractInstanceFields{}
	return &this
}

// NewContractInstanceFieldsWithDefaults instantiates a new ContractInstanceFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInstanceFieldsWithDefaults() *ContractInstanceFields {
	this := ContractInstanceFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContractInstanceFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContractInstanceFields) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ContractInstanceFields) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ContractInstanceFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFlavorName returns the FlavorName field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetFlavorName() string {
	if o == nil || IsNil(o.FlavorName) {
		var ret string
		return ret
	}
	return *o.FlavorName
}

// GetFlavorNameOk returns a tuple with the FlavorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetFlavorNameOk() (*string, bool) {
	if o == nil || IsNil(o.FlavorName) {
		return nil, false
	}
	return o.FlavorName, true
}

// HasFlavorName returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasFlavorName() bool {
	if o != nil && !IsNil(o.FlavorName) {
		return true
	}

	return false
}

// SetFlavorName gets a reference to the given string and assigns it to the FlavorName field.
func (o *ContractInstanceFields) SetFlavorName(v string) {
	o.FlavorName = &v
}

// GetGpuCount returns the GpuCount field value if set, zero value otherwise.
func (o *ContractInstanceFields) GetGpuCount() int32 {
	if o == nil || IsNil(o.GpuCount) {
		var ret int32
		return ret
	}
	return *o.GpuCount
}

// GetGpuCountOk returns a tuple with the GpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractInstanceFields) GetGpuCountOk() (*int32, bool) {
	if o == nil || IsNil(o.GpuCount) {
		return nil, false
	}
	return o.GpuCount, true
}

// HasGpuCount returns a boolean if a field has been set.
func (o *ContractInstanceFields) HasGpuCount() bool {
	if o != nil && !IsNil(o.GpuCount) {
		return true
	}

	return false
}

// SetGpuCount gets a reference to the given int32 and assigns it to the GpuCount field.
func (o *ContractInstanceFields) SetGpuCount(v int32) {
	o.GpuCount = &v
}

func (o ContractInstanceFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInstanceFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FlavorName) {
		toSerialize["flavor_name"] = o.FlavorName
	}
	if !IsNil(o.GpuCount) {
		toSerialize["gpu_count"] = o.GpuCount
	}
	return toSerialize, nil
}

type NullableContractInstanceFields struct {
	value *ContractInstanceFields
	isSet bool
}

func (v NullableContractInstanceFields) Get() *ContractInstanceFields {
	return v.value
}

func (v *NullableContractInstanceFields) Set(val *ContractInstanceFields) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInstanceFields) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInstanceFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInstanceFields(val *ContractInstanceFields) *NullableContractInstanceFields {
	return &NullableContractInstanceFields{value: val, isSet: true}
}

func (v NullableContractInstanceFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInstanceFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


