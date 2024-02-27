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

// checks if the AdminAddUpdateImageOrganizationPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminAddUpdateImageOrganizationPayload{}

// AdminAddUpdateImageOrganizationPayload struct for AdminAddUpdateImageOrganizationPayload
type AdminAddUpdateImageOrganizationPayload struct {
	Name string `json:"name"`
	OpenstackId string `json:"openstack_id"`
	RegionId int32 `json:"region_id"`
	Type string `json:"type"`
	Version string `json:"version"`
	Size int32 `json:"size"`
	IsPublic bool `json:"is_public"`
	Description *string `json:"description,omitempty"`
	Labels []string `json:"labels,omitempty"`
	Organizations []int32 `json:"organizations"`
}

type _AdminAddUpdateImageOrganizationPayload AdminAddUpdateImageOrganizationPayload

// NewAdminAddUpdateImageOrganizationPayload instantiates a new AdminAddUpdateImageOrganizationPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminAddUpdateImageOrganizationPayload(name string, openstackId string, regionId int32, type_ string, version string, size int32, isPublic bool, organizations []int32) *AdminAddUpdateImageOrganizationPayload {
	this := AdminAddUpdateImageOrganizationPayload{}
	this.Name = name
	this.OpenstackId = openstackId
	this.RegionId = regionId
	this.Type = type_
	this.Version = version
	this.Size = size
	this.IsPublic = isPublic
	this.Organizations = organizations
	return &this
}

// NewAdminAddUpdateImageOrganizationPayloadWithDefaults instantiates a new AdminAddUpdateImageOrganizationPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminAddUpdateImageOrganizationPayloadWithDefaults() *AdminAddUpdateImageOrganizationPayload {
	this := AdminAddUpdateImageOrganizationPayload{}
	return &this
}

// GetName returns the Name field value
func (o *AdminAddUpdateImageOrganizationPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetName(v string) {
	o.Name = v
}

// GetOpenstackId returns the OpenstackId field value
func (o *AdminAddUpdateImageOrganizationPayload) GetOpenstackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetOpenstackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenstackId, true
}

// SetOpenstackId sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetOpenstackId(v string) {
	o.OpenstackId = v
}

// GetRegionId returns the RegionId field value
func (o *AdminAddUpdateImageOrganizationPayload) GetRegionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetRegionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetRegionId(v int32) {
	o.RegionId = v
}

// GetType returns the Type field value
func (o *AdminAddUpdateImageOrganizationPayload) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *AdminAddUpdateImageOrganizationPayload) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetVersion(v string) {
	o.Version = v
}

// GetSize returns the Size field value
func (o *AdminAddUpdateImageOrganizationPayload) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetSize(v int32) {
	o.Size = v
}

// GetIsPublic returns the IsPublic field value
func (o *AdminAddUpdateImageOrganizationPayload) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdminAddUpdateImageOrganizationPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdminAddUpdateImageOrganizationPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdminAddUpdateImageOrganizationPayload) SetDescription(v string) {
	o.Description = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AdminAddUpdateImageOrganizationPayload) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AdminAddUpdateImageOrganizationPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *AdminAddUpdateImageOrganizationPayload) SetLabels(v []string) {
	o.Labels = v
}

// GetOrganizations returns the Organizations field value
func (o *AdminAddUpdateImageOrganizationPayload) GetOrganizations() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value
// and a boolean to check if the value has been set.
func (o *AdminAddUpdateImageOrganizationPayload) GetOrganizationsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organizations, true
}

// SetOrganizations sets field value
func (o *AdminAddUpdateImageOrganizationPayload) SetOrganizations(v []int32) {
	o.Organizations = v
}

func (o AdminAddUpdateImageOrganizationPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminAddUpdateImageOrganizationPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["openstack_id"] = o.OpenstackId
	toSerialize["region_id"] = o.RegionId
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	toSerialize["size"] = o.Size
	toSerialize["is_public"] = o.IsPublic
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["organizations"] = o.Organizations
	return toSerialize, nil
}

func (o *AdminAddUpdateImageOrganizationPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"openstack_id",
		"region_id",
		"type",
		"version",
		"size",
		"is_public",
		"organizations",
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

	varAdminAddUpdateImageOrganizationPayload := _AdminAddUpdateImageOrganizationPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAdminAddUpdateImageOrganizationPayload)

	if err != nil {
		return err
	}

	*o = AdminAddUpdateImageOrganizationPayload(varAdminAddUpdateImageOrganizationPayload)

	return err
}

type NullableAdminAddUpdateImageOrganizationPayload struct {
	value *AdminAddUpdateImageOrganizationPayload
	isSet bool
}

func (v NullableAdminAddUpdateImageOrganizationPayload) Get() *AdminAddUpdateImageOrganizationPayload {
	return v.value
}

func (v *NullableAdminAddUpdateImageOrganizationPayload) Set(val *AdminAddUpdateImageOrganizationPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminAddUpdateImageOrganizationPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminAddUpdateImageOrganizationPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminAddUpdateImageOrganizationPayload(val *AdminAddUpdateImageOrganizationPayload) *NullableAdminAddUpdateImageOrganizationPayload {
	return &NullableAdminAddUpdateImageOrganizationPayload{value: val, isSet: true}
}

func (v NullableAdminAddUpdateImageOrganizationPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminAddUpdateImageOrganizationPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


