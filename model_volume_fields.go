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

// checks if the VolumeFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeFields{}

// VolumeFields struct for VolumeFields
type VolumeFields struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Environment *EnvironmentFieldsforVolume `json:"environment,omitempty"`
	Description *string `json:"description,omitempty"`
	VolumeType *string `json:"volume_type,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Status *string `json:"status,omitempty"`
	Bootable *bool `json:"bootable,omitempty"`
	ImageId *int32 `json:"image_id,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewVolumeFields instantiates a new VolumeFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeFields() *VolumeFields {
	this := VolumeFields{}
	return &this
}

// NewVolumeFieldsWithDefaults instantiates a new VolumeFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeFieldsWithDefaults() *VolumeFields {
	this := VolumeFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VolumeFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VolumeFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VolumeFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VolumeFields) SetName(v string) {
	o.Name = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *VolumeFields) GetEnvironment() EnvironmentFieldsforVolume {
	if o == nil || IsNil(o.Environment) {
		var ret EnvironmentFieldsforVolume
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetEnvironmentOk() (*EnvironmentFieldsforVolume, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *VolumeFields) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given EnvironmentFieldsforVolume and assigns it to the Environment field.
func (o *VolumeFields) SetEnvironment(v EnvironmentFieldsforVolume) {
	o.Environment = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VolumeFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VolumeFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VolumeFields) SetDescription(v string) {
	o.Description = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *VolumeFields) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType) {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *VolumeFields) HasVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *VolumeFields) SetVolumeType(v string) {
	o.VolumeType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VolumeFields) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VolumeFields) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *VolumeFields) SetSize(v int64) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VolumeFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VolumeFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VolumeFields) SetStatus(v string) {
	o.Status = &v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *VolumeFields) GetBootable() bool {
	if o == nil || IsNil(o.Bootable) {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetBootableOk() (*bool, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *VolumeFields) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *VolumeFields) SetBootable(v bool) {
	o.Bootable = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *VolumeFields) GetImageId() int32 {
	if o == nil || IsNil(o.ImageId) {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetImageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *VolumeFields) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *VolumeFields) SetImageId(v int32) {
	o.ImageId = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *VolumeFields) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *VolumeFields) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *VolumeFields) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VolumeFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VolumeFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VolumeFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *VolumeFields) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *VolumeFields) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *VolumeFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o VolumeFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volume_type"] = o.VolumeType
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Bootable) {
		toSerialize["bootable"] = o.Bootable
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableVolumeFields struct {
	value *VolumeFields
	isSet bool
}

func (v NullableVolumeFields) Get() *VolumeFields {
	return v.value
}

func (v *NullableVolumeFields) Set(val *VolumeFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeFields(val *VolumeFields) *NullableVolumeFields {
	return &NullableVolumeFields{value: val, isSet: true}
}

func (v NullableVolumeFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


