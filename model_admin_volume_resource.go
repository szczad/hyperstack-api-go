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

// checks if the AdminVolumeResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminVolumeResource{}

// AdminVolumeResource struct for AdminVolumeResource
type AdminVolumeResource struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OpenstackId *string `json:"openstack_id,omitempty"`
	Type *string `json:"type,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Status *string `json:"status,omitempty"`
	Bootable *bool `json:"bootable,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewAdminVolumeResource instantiates a new AdminVolumeResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminVolumeResource() *AdminVolumeResource {
	this := AdminVolumeResource{}
	return &this
}

// NewAdminVolumeResourceWithDefaults instantiates a new AdminVolumeResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminVolumeResourceWithDefaults() *AdminVolumeResource {
	this := AdminVolumeResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AdminVolumeResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdminVolumeResource) SetName(v string) {
	o.Name = &v
}

// GetOpenstackId returns the OpenstackId field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetOpenstackId() string {
	if o == nil || IsNil(o.OpenstackId) {
		var ret string
		return ret
	}
	return *o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetOpenstackIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenstackId) {
		return nil, false
	}
	return o.OpenstackId, true
}

// HasOpenstackId returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasOpenstackId() bool {
	if o != nil && !IsNil(o.OpenstackId) {
		return true
	}

	return false
}

// SetOpenstackId gets a reference to the given string and assigns it to the OpenstackId field.
func (o *AdminVolumeResource) SetOpenstackId(v string) {
	o.OpenstackId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AdminVolumeResource) SetType(v string) {
	o.Type = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *AdminVolumeResource) SetSize(v int32) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AdminVolumeResource) SetStatus(v string) {
	o.Status = &v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetBootable() bool {
	if o == nil || IsNil(o.Bootable) {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetBootableOk() (*bool, bool) {
	if o == nil || IsNil(o.Bootable) {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasBootable() bool {
	if o != nil && !IsNil(o.Bootable) {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *AdminVolumeResource) SetBootable(v bool) {
	o.Bootable = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AdminVolumeResource) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminVolumeResource) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AdminVolumeResource) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AdminVolumeResource) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o AdminVolumeResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminVolumeResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OpenstackId) {
		toSerialize["openstack_id"] = o.OpenstackId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAdminVolumeResource struct {
	value *AdminVolumeResource
	isSet bool
}

func (v NullableAdminVolumeResource) Get() *AdminVolumeResource {
	return v.value
}

func (v *NullableAdminVolumeResource) Set(val *AdminVolumeResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminVolumeResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminVolumeResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminVolumeResource(val *AdminVolumeResource) *NullableAdminVolumeResource {
	return &NullableAdminVolumeResource{value: val, isSet: true}
}

func (v NullableAdminVolumeResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminVolumeResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


