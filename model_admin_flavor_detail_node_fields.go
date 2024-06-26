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

// checks if the AdminFlavorDetailNodeFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminFlavorDetailNodeFields{}

// AdminFlavorDetailNodeFields struct for AdminFlavorDetailNodeFields
type AdminFlavorDetailNodeFields struct {
	Id *int32 `json:"id,omitempty"`
	OpenstackId *string `json:"openstack_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Available *int32 `json:"available,omitempty"`
	Status *string `json:"status,omitempty"`
	ProvisionDate *time.Time `json:"provision_date,omitempty"`
	Projects []string `json:"projects,omitempty"`
}

// NewAdminFlavorDetailNodeFields instantiates a new AdminFlavorDetailNodeFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminFlavorDetailNodeFields() *AdminFlavorDetailNodeFields {
	this := AdminFlavorDetailNodeFields{}
	return &this
}

// NewAdminFlavorDetailNodeFieldsWithDefaults instantiates a new AdminFlavorDetailNodeFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminFlavorDetailNodeFieldsWithDefaults() *AdminFlavorDetailNodeFields {
	this := AdminFlavorDetailNodeFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AdminFlavorDetailNodeFields) SetId(v int32) {
	o.Id = &v
}

// GetOpenstackId returns the OpenstackId field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetOpenstackId() string {
	if o == nil || IsNil(o.OpenstackId) {
		var ret string
		return ret
	}
	return *o.OpenstackId
}

// GetOpenstackIdOk returns a tuple with the OpenstackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetOpenstackIdOk() (*string, bool) {
	if o == nil || IsNil(o.OpenstackId) {
		return nil, false
	}
	return o.OpenstackId, true
}

// HasOpenstackId returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasOpenstackId() bool {
	if o != nil && !IsNil(o.OpenstackId) {
		return true
	}

	return false
}

// SetOpenstackId gets a reference to the given string and assigns it to the OpenstackId field.
func (o *AdminFlavorDetailNodeFields) SetOpenstackId(v string) {
	o.OpenstackId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdminFlavorDetailNodeFields) SetName(v string) {
	o.Name = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetAvailable() int32 {
	if o == nil || IsNil(o.Available) {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetAvailableOk() (*int32, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *AdminFlavorDetailNodeFields) SetAvailable(v int32) {
	o.Available = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AdminFlavorDetailNodeFields) SetStatus(v string) {
	o.Status = &v
}

// GetProvisionDate returns the ProvisionDate field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetProvisionDate() time.Time {
	if o == nil || IsNil(o.ProvisionDate) {
		var ret time.Time
		return ret
	}
	return *o.ProvisionDate
}

// GetProvisionDateOk returns a tuple with the ProvisionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetProvisionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProvisionDate) {
		return nil, false
	}
	return o.ProvisionDate, true
}

// HasProvisionDate returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasProvisionDate() bool {
	if o != nil && !IsNil(o.ProvisionDate) {
		return true
	}

	return false
}

// SetProvisionDate gets a reference to the given time.Time and assigns it to the ProvisionDate field.
func (o *AdminFlavorDetailNodeFields) SetProvisionDate(v time.Time) {
	o.ProvisionDate = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *AdminFlavorDetailNodeFields) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminFlavorDetailNodeFields) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *AdminFlavorDetailNodeFields) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *AdminFlavorDetailNodeFields) SetProjects(v []string) {
	o.Projects = v
}

func (o AdminFlavorDetailNodeFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminFlavorDetailNodeFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ProvisionDate) {
		toSerialize["provision_date"] = o.ProvisionDate
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableAdminFlavorDetailNodeFields struct {
	value *AdminFlavorDetailNodeFields
	isSet bool
}

func (v NullableAdminFlavorDetailNodeFields) Get() *AdminFlavorDetailNodeFields {
	return v.value
}

func (v *NullableAdminFlavorDetailNodeFields) Set(val *AdminFlavorDetailNodeFields) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminFlavorDetailNodeFields) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminFlavorDetailNodeFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminFlavorDetailNodeFields(val *AdminFlavorDetailNodeFields) *NullableAdminFlavorDetailNodeFields {
	return &NullableAdminFlavorDetailNodeFields{value: val, isSet: true}
}

func (v NullableAdminFlavorDetailNodeFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminFlavorDetailNodeFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


