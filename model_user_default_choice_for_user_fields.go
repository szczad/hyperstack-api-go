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

// checks if the UserDefaultChoiceForUserFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefaultChoiceForUserFields{}

// UserDefaultChoiceForUserFields struct for UserDefaultChoiceForUserFields
type UserDefaultChoiceForUserFields struct {
	RegionId *int32 `json:"region_id,omitempty"`
	EnvironmentId *int32 `json:"environment_id,omitempty"`
	KeypairId *int32 `json:"keypair_id,omitempty"`
	FlavorId *int32 `json:"flavor_id,omitempty"`
	ImageId *int32 `json:"image_id,omitempty"`
}

// NewUserDefaultChoiceForUserFields instantiates a new UserDefaultChoiceForUserFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefaultChoiceForUserFields() *UserDefaultChoiceForUserFields {
	this := UserDefaultChoiceForUserFields{}
	return &this
}

// NewUserDefaultChoiceForUserFieldsWithDefaults instantiates a new UserDefaultChoiceForUserFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefaultChoiceForUserFieldsWithDefaults() *UserDefaultChoiceForUserFields {
	this := UserDefaultChoiceForUserFields{}
	return &this
}

// GetRegionId returns the RegionId field value if set, zero value otherwise.
func (o *UserDefaultChoiceForUserFields) GetRegionId() int32 {
	if o == nil || IsNil(o.RegionId) {
		var ret int32
		return ret
	}
	return *o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoiceForUserFields) GetRegionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RegionId) {
		return nil, false
	}
	return o.RegionId, true
}

// HasRegionId returns a boolean if a field has been set.
func (o *UserDefaultChoiceForUserFields) HasRegionId() bool {
	if o != nil && !IsNil(o.RegionId) {
		return true
	}

	return false
}

// SetRegionId gets a reference to the given int32 and assigns it to the RegionId field.
func (o *UserDefaultChoiceForUserFields) SetRegionId(v int32) {
	o.RegionId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *UserDefaultChoiceForUserFields) GetEnvironmentId() int32 {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret int32
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoiceForUserFields) GetEnvironmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *UserDefaultChoiceForUserFields) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given int32 and assigns it to the EnvironmentId field.
func (o *UserDefaultChoiceForUserFields) SetEnvironmentId(v int32) {
	o.EnvironmentId = &v
}

// GetKeypairId returns the KeypairId field value if set, zero value otherwise.
func (o *UserDefaultChoiceForUserFields) GetKeypairId() int32 {
	if o == nil || IsNil(o.KeypairId) {
		var ret int32
		return ret
	}
	return *o.KeypairId
}

// GetKeypairIdOk returns a tuple with the KeypairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoiceForUserFields) GetKeypairIdOk() (*int32, bool) {
	if o == nil || IsNil(o.KeypairId) {
		return nil, false
	}
	return o.KeypairId, true
}

// HasKeypairId returns a boolean if a field has been set.
func (o *UserDefaultChoiceForUserFields) HasKeypairId() bool {
	if o != nil && !IsNil(o.KeypairId) {
		return true
	}

	return false
}

// SetKeypairId gets a reference to the given int32 and assigns it to the KeypairId field.
func (o *UserDefaultChoiceForUserFields) SetKeypairId(v int32) {
	o.KeypairId = &v
}

// GetFlavorId returns the FlavorId field value if set, zero value otherwise.
func (o *UserDefaultChoiceForUserFields) GetFlavorId() int32 {
	if o == nil || IsNil(o.FlavorId) {
		var ret int32
		return ret
	}
	return *o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoiceForUserFields) GetFlavorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FlavorId) {
		return nil, false
	}
	return o.FlavorId, true
}

// HasFlavorId returns a boolean if a field has been set.
func (o *UserDefaultChoiceForUserFields) HasFlavorId() bool {
	if o != nil && !IsNil(o.FlavorId) {
		return true
	}

	return false
}

// SetFlavorId gets a reference to the given int32 and assigns it to the FlavorId field.
func (o *UserDefaultChoiceForUserFields) SetFlavorId(v int32) {
	o.FlavorId = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *UserDefaultChoiceForUserFields) GetImageId() int32 {
	if o == nil || IsNil(o.ImageId) {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoiceForUserFields) GetImageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *UserDefaultChoiceForUserFields) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *UserDefaultChoiceForUserFields) SetImageId(v int32) {
	o.ImageId = &v
}

func (o UserDefaultChoiceForUserFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefaultChoiceForUserFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RegionId) {
		toSerialize["region_id"] = o.RegionId
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.KeypairId) {
		toSerialize["keypair_id"] = o.KeypairId
	}
	if !IsNil(o.FlavorId) {
		toSerialize["flavor_id"] = o.FlavorId
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	return toSerialize, nil
}

type NullableUserDefaultChoiceForUserFields struct {
	value *UserDefaultChoiceForUserFields
	isSet bool
}

func (v NullableUserDefaultChoiceForUserFields) Get() *UserDefaultChoiceForUserFields {
	return v.value
}

func (v *NullableUserDefaultChoiceForUserFields) Set(val *UserDefaultChoiceForUserFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefaultChoiceForUserFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefaultChoiceForUserFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefaultChoiceForUserFields(val *UserDefaultChoiceForUserFields) *NullableUserDefaultChoiceForUserFields {
	return &NullableUserDefaultChoiceForUserFields{value: val, isSet: true}
}

func (v NullableUserDefaultChoiceForUserFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefaultChoiceForUserFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

