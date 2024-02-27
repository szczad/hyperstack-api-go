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

// checks if the GPUFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GPUFields{}

// GPUFields struct for GPUFields
type GPUFields struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Regions []GPURegionFields `json:"regions,omitempty"`
	ExampleMetadata *string `json:"example_metadata,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewGPUFields instantiates a new GPUFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGPUFields() *GPUFields {
	this := GPUFields{}
	return &this
}

// NewGPUFieldsWithDefaults instantiates a new GPUFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGPUFieldsWithDefaults() *GPUFields {
	this := GPUFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GPUFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GPUFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GPUFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GPUFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GPUFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GPUFields) SetName(v string) {
	o.Name = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *GPUFields) GetRegions() []GPURegionFields {
	if o == nil || IsNil(o.Regions) {
		var ret []GPURegionFields
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetRegionsOk() ([]GPURegionFields, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *GPUFields) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []GPURegionFields and assigns it to the Regions field.
func (o *GPUFields) SetRegions(v []GPURegionFields) {
	o.Regions = v
}

// GetExampleMetadata returns the ExampleMetadata field value if set, zero value otherwise.
func (o *GPUFields) GetExampleMetadata() string {
	if o == nil || IsNil(o.ExampleMetadata) {
		var ret string
		return ret
	}
	return *o.ExampleMetadata
}

// GetExampleMetadataOk returns a tuple with the ExampleMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetExampleMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.ExampleMetadata) {
		return nil, false
	}
	return o.ExampleMetadata, true
}

// HasExampleMetadata returns a boolean if a field has been set.
func (o *GPUFields) HasExampleMetadata() bool {
	if o != nil && !IsNil(o.ExampleMetadata) {
		return true
	}

	return false
}

// SetExampleMetadata gets a reference to the given string and assigns it to the ExampleMetadata field.
func (o *GPUFields) SetExampleMetadata(v string) {
	o.ExampleMetadata = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GPUFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GPUFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GPUFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GPUFields) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPUFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GPUFields) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GPUFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GPUFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GPUFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !IsNil(o.ExampleMetadata) {
		toSerialize["example_metadata"] = o.ExampleMetadata
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGPUFields struct {
	value *GPUFields
	isSet bool
}

func (v NullableGPUFields) Get() *GPUFields {
	return v.value
}

func (v *NullableGPUFields) Set(val *GPUFields) {
	v.value = val
	v.isSet = true
}

func (v NullableGPUFields) IsSet() bool {
	return v.isSet
}

func (v *NullableGPUFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGPUFields(val *GPUFields) *NullableGPUFields {
	return &NullableGPUFields{value: val, isSet: true}
}

func (v NullableGPUFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGPUFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

