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

// checks if the AdminFlavorResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminFlavorResource{}

// AdminFlavorResource struct for AdminFlavorResource
type AdminFlavorResource struct {
	Id *int32 `json:"id,omitempty"`
	OpenstackId *string `json:"openstack_id,omitempty"`
	Name *string `json:"name,omitempty"`
	RegionName *string `json:"region_name,omitempty"`
	Cpu *int32 `json:"cpu,omitempty"`
	Ram *float32 `json:"ram,omitempty"`
	Disk *int32 `json:"disk,omitempty"`
	Gpu *string `json:"gpu,omitempty"`
	Status *string `json:"status,omitempty"`
	GpuCount *int32 `json:"gpu_count,omitempty"`
	StockAvailable *bool `json:"stock_available,omitempty"`
	Nodes []AdminNodeResource `json:"nodes,omitempty"`
	Vms []int32 `json:"vms,omitempty"`
	IsPublic *bool `json:"is_public,omitempty"`
	IsCustom *bool `json:"is_custom,omitempty"`
	Ephemeral *int32 `json:"ephemeral,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewAdminFlavorResource instantiates a new AdminFlavorResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminFlavorResource() *AdminFlavorResource {
	this := AdminFlavorResource{}
	return &this
}

// NewAdminFlavorResourceWithDefaults instantiates a new AdminFlavorResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminFlavorResourceWithDefaults() *AdminFlavorResource {
	this := AdminFlavorResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AdminFlavorResource) SetId(v int32) {
	o.Id = &v
}

// GetOpenstackId returns the OpenstackId field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetOpenstackId() string {
	if o == nil || IsNil(o.OpenstackId) {
		var ret string
		return ret
	}
	return *o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetOpenstackIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenstackId) {
		return nil, false
	}
	return o.OpenstackId, true
}

// HasOpenstackId returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasOpenstackId() bool {
	if o != nil && !IsNil(o.OpenstackId) {
		return true
	}

	return false
}

// SetOpenstackId gets a reference to the given string and assigns it to the OpenstackId field.
func (o *AdminFlavorResource) SetOpenstackId(v string) {
	o.OpenstackId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdminFlavorResource) SetName(v string) {
	o.Name = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *AdminFlavorResource) SetRegionName(v string) {
	o.RegionName = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *AdminFlavorResource) SetCpu(v int32) {
	o.Cpu = &v
}

// GetRam returns the Ram field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetRam() float32 {
	if o == nil || IsNil(o.Ram) {
		var ret float32
		return ret
	}
	return *o.Ram
}

// GetRamOk returns a tuple with the Ram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetRamOk() (*float32, bool) {
	if o == nil || IsNil(o.Ram) {
		return nil, false
	}
	return o.Ram, true
}

// HasRam returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasRam() bool {
	if o != nil && !IsNil(o.Ram) {
		return true
	}

	return false
}

// SetRam gets a reference to the given float32 and assigns it to the Ram field.
func (o *AdminFlavorResource) SetRam(v float32) {
	o.Ram = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetDisk() int32 {
	if o == nil || IsNil(o.Disk) {
		var ret int32
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.Disk) {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasDisk() bool {
	if o != nil && !IsNil(o.Disk) {
		return true
	}

	return false
}

// SetDisk gets a reference to the given int32 and assigns it to the Disk field.
func (o *AdminFlavorResource) SetDisk(v int32) {
	o.Disk = &v
}

// GetGpu returns the Gpu field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetGpu() string {
	if o == nil || IsNil(o.Gpu) {
		var ret string
		return ret
	}
	return *o.Gpu
}

// GetGpuOk returns a tuple with the Gpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetGpuOk() (*string, bool) {
	if o == nil || IsNil(o.Gpu) {
		return nil, false
	}
	return o.Gpu, true
}

// HasGpu returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasGpu() bool {
	if o != nil && !IsNil(o.Gpu) {
		return true
	}

	return false
}

// SetGpu gets a reference to the given string and assigns it to the Gpu field.
func (o *AdminFlavorResource) SetGpu(v string) {
	o.Gpu = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AdminFlavorResource) SetStatus(v string) {
	o.Status = &v
}

// GetGpuCount returns the GpuCount field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetGpuCount() int32 {
	if o == nil || IsNil(o.GpuCount) {
		var ret int32
		return ret
	}
	return *o.GpuCount
}

// GetGpuCountOk returns a tuple with the GpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetGpuCountOk() (*int32, bool) {
	if o == nil || IsNil(o.GpuCount) {
		return nil, false
	}
	return o.GpuCount, true
}

// HasGpuCount returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasGpuCount() bool {
	if o != nil && !IsNil(o.GpuCount) {
		return true
	}

	return false
}

// SetGpuCount gets a reference to the given int32 and assigns it to the GpuCount field.
func (o *AdminFlavorResource) SetGpuCount(v int32) {
	o.GpuCount = &v
}

// GetStockAvailable returns the StockAvailable field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetStockAvailable() bool {
	if o == nil || IsNil(o.StockAvailable) {
		var ret bool
		return ret
	}
	return *o.StockAvailable
}

// GetStockAvailableOk returns a tuple with the StockAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetStockAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.StockAvailable) {
		return nil, false
	}
	return o.StockAvailable, true
}

// HasStockAvailable returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasStockAvailable() bool {
	if o != nil && !IsNil(o.StockAvailable) {
		return true
	}

	return false
}

// SetStockAvailable gets a reference to the given bool and assigns it to the StockAvailable field.
func (o *AdminFlavorResource) SetStockAvailable(v bool) {
	o.StockAvailable = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetNodes() []AdminNodeResource {
	if o == nil || IsNil(o.Nodes) {
		var ret []AdminNodeResource
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetNodesOk() ([]AdminNodeResource, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []AdminNodeResource and assigns it to the Nodes field.
func (o *AdminFlavorResource) SetNodes(v []AdminNodeResource) {
	o.Nodes = v
}

// GetVms returns the Vms field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetVms() []int32 {
	if o == nil || IsNil(o.Vms) {
		var ret []int32
		return ret
	}
	return o.Vms
}

// GetVmsOk returns a tuple with the Vms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetVmsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Vms) {
		return nil, false
	}
	return o.Vms, true
}

// HasVms returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasVms() bool {
	if o != nil && !IsNil(o.Vms) {
		return true
	}

	return false
}

// SetVms gets a reference to the given []int32 and assigns it to the Vms field.
func (o *AdminFlavorResource) SetVms(v []int32) {
	o.Vms = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *AdminFlavorResource) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetIsCustom returns the IsCustom field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetIsCustom() bool {
	if o == nil || IsNil(o.IsCustom) {
		var ret bool
		return ret
	}
	return *o.IsCustom
}

// GetIsCustomOk returns a tuple with the IsCustom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetIsCustomOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustom) {
		return nil, false
	}
	return o.IsCustom, true
}

// HasIsCustom returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasIsCustom() bool {
	if o != nil && !IsNil(o.IsCustom) {
		return true
	}

	return false
}

// SetIsCustom gets a reference to the given bool and assigns it to the IsCustom field.
func (o *AdminFlavorResource) SetIsCustom(v bool) {
	o.IsCustom = &v
}

// GetEphemeral returns the Ephemeral field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetEphemeral() int32 {
	if o == nil || IsNil(o.Ephemeral) {
		var ret int32
		return ret
	}
	return *o.Ephemeral
}

// GetEphemeralOk returns a tuple with the Ephemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetEphemeralOk() (*int32, bool) {
	if o == nil || IsNil(o.Ephemeral) {
		return nil, false
	}
	return o.Ephemeral, true
}

// HasEphemeral returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasEphemeral() bool {
	if o != nil && !IsNil(o.Ephemeral) {
		return true
	}

	return false
}

// SetEphemeral gets a reference to the given int32 and assigns it to the Ephemeral field.
func (o *AdminFlavorResource) SetEphemeral(v int32) {
	o.Ephemeral = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdminFlavorResource) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AdminFlavorResource) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorResource) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AdminFlavorResource) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AdminFlavorResource) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o AdminFlavorResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminFlavorResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OpenstackId) {
		toSerialize["openstack_id"] = o.OpenstackId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RegionName) {
		toSerialize["region_name"] = o.RegionName
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Ram) {
		toSerialize["ram"] = o.Ram
	}
	if !IsNil(o.Disk) {
		toSerialize["disk"] = o.Disk
	}
	if !IsNil(o.Gpu) {
		toSerialize["gpu"] = o.Gpu
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.GpuCount) {
		toSerialize["gpu_count"] = o.GpuCount
	}
	if !IsNil(o.StockAvailable) {
		toSerialize["stock_available"] = o.StockAvailable
	}
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.Vms) {
		toSerialize["vms"] = o.Vms
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.IsCustom) {
		toSerialize["is_custom"] = o.IsCustom
	}
	if !IsNil(o.Ephemeral) {
		toSerialize["ephemeral"] = o.Ephemeral
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableAdminFlavorResource struct {
	value *AdminFlavorResource
	isSet bool
}

func (v NullableAdminFlavorResource) Get() *AdminFlavorResource {
	return v.value
}

func (v *NullableAdminFlavorResource) Set(val *AdminFlavorResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminFlavorResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminFlavorResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminFlavorResource(val *AdminFlavorResource) *NullableAdminFlavorResource {
	return &NullableAdminFlavorResource{value: val, isSet: true}
}

func (v NullableAdminFlavorResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminFlavorResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

