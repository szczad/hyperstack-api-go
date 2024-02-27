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

// checks if the InstancesSummaryFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstancesSummaryFields{}

// InstancesSummaryFields struct for InstancesSummaryFields
type InstancesSummaryFields struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Environment *string `json:"environment,omitempty"`
	OrgId *int32 `json:"org_id,omitempty"`
	EnvironmentId *int32 `json:"environment_id,omitempty"`
	Image *string `json:"image,omitempty"`
	ImageId *int32 `json:"image_id,omitempty"`
	Flavor *string `json:"flavor,omitempty"`
	FlavorId *int32 `json:"flavor_id,omitempty"`
	Keypair *string `json:"keypair,omitempty"`
	KeypairId *int32 `json:"keypair_id,omitempty"`
	FixedIp *string `json:"fixed_ip,omitempty"`
	FloatingIp *string `json:"floating_ip,omitempty"`
	FloatingIpStatus *string `json:"floating_ip_status,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewInstancesSummaryFields instantiates a new InstancesSummaryFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstancesSummaryFields() *InstancesSummaryFields {
	this := InstancesSummaryFields{}
	return &this
}

// NewInstancesSummaryFieldsWithDefaults instantiates a new InstancesSummaryFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstancesSummaryFieldsWithDefaults() *InstancesSummaryFields {
	this := InstancesSummaryFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InstancesSummaryFields) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstancesSummaryFields) SetName(v string) {
	o.Name = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *InstancesSummaryFields) SetEnvironment(v string) {
	o.Environment = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *InstancesSummaryFields) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetEnvironmentId() int32 {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret int32
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetEnvironmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given int32 and assigns it to the EnvironmentId field.
func (o *InstancesSummaryFields) SetEnvironmentId(v int32) {
	o.EnvironmentId = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *InstancesSummaryFields) SetImage(v string) {
	o.Image = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetImageId() int32 {
	if o == nil || IsNil(o.ImageId) {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetImageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *InstancesSummaryFields) SetImageId(v int32) {
	o.ImageId = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetFlavor() string {
	if o == nil || IsNil(o.Flavor) {
		var ret string
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetFlavorOk() (*string, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given string and assigns it to the Flavor field.
func (o *InstancesSummaryFields) SetFlavor(v string) {
	o.Flavor = &v
}

// GetFlavorId returns the FlavorId field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetFlavorId() int32 {
	if o == nil || IsNil(o.FlavorId) {
		var ret int32
		return ret
	}
	return *o.FlavorId
}

// GetFlavorIdOk returns a tuple with the FlavorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetFlavorIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FlavorId) {
		return nil, false
	}
	return o.FlavorId, true
}

// HasFlavorId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasFlavorId() bool {
	if o != nil && !IsNil(o.FlavorId) {
		return true
	}

	return false
}

// SetFlavorId gets a reference to the given int32 and assigns it to the FlavorId field.
func (o *InstancesSummaryFields) SetFlavorId(v int32) {
	o.FlavorId = &v
}

// GetKeypair returns the Keypair field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetKeypair() string {
	if o == nil || IsNil(o.Keypair) {
		var ret string
		return ret
	}
	return *o.Keypair
}

// GetKeypairOk returns a tuple with the Keypair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetKeypairOk() (*string, bool) {
	if o == nil || IsNil(o.Keypair) {
		return nil, false
	}
	return o.Keypair, true
}

// HasKeypair returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasKeypair() bool {
	if o != nil && !IsNil(o.Keypair) {
		return true
	}

	return false
}

// SetKeypair gets a reference to the given string and assigns it to the Keypair field.
func (o *InstancesSummaryFields) SetKeypair(v string) {
	o.Keypair = &v
}

// GetKeypairId returns the KeypairId field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetKeypairId() int32 {
	if o == nil || IsNil(o.KeypairId) {
		var ret int32
		return ret
	}
	return *o.KeypairId
}

// GetKeypairIdOk returns a tuple with the KeypairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetKeypairIdOk() (*int32, bool) {
	if o == nil || IsNil(o.KeypairId) {
		return nil, false
	}
	return o.KeypairId, true
}

// HasKeypairId returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasKeypairId() bool {
	if o != nil && !IsNil(o.KeypairId) {
		return true
	}

	return false
}

// SetKeypairId gets a reference to the given int32 and assigns it to the KeypairId field.
func (o *InstancesSummaryFields) SetKeypairId(v int32) {
	o.KeypairId = &v
}

// GetFixedIp returns the FixedIp field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetFixedIp() string {
	if o == nil || IsNil(o.FixedIp) {
		var ret string
		return ret
	}
	return *o.FixedIp
}

// GetFixedIpOk returns a tuple with the FixedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetFixedIpOk() (*string, bool) {
	if o == nil || IsNil(o.FixedIp) {
		return nil, false
	}
	return o.FixedIp, true
}

// HasFixedIp returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasFixedIp() bool {
	if o != nil && !IsNil(o.FixedIp) {
		return true
	}

	return false
}

// SetFixedIp gets a reference to the given string and assigns it to the FixedIp field.
func (o *InstancesSummaryFields) SetFixedIp(v string) {
	o.FixedIp = &v
}

// GetFloatingIp returns the FloatingIp field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetFloatingIp() string {
	if o == nil || IsNil(o.FloatingIp) {
		var ret string
		return ret
	}
	return *o.FloatingIp
}

// GetFloatingIpOk returns a tuple with the FloatingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetFloatingIpOk() (*string, bool) {
	if o == nil || IsNil(o.FloatingIp) {
		return nil, false
	}
	return o.FloatingIp, true
}

// HasFloatingIp returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasFloatingIp() bool {
	if o != nil && !IsNil(o.FloatingIp) {
		return true
	}

	return false
}

// SetFloatingIp gets a reference to the given string and assigns it to the FloatingIp field.
func (o *InstancesSummaryFields) SetFloatingIp(v string) {
	o.FloatingIp = &v
}

// GetFloatingIpStatus returns the FloatingIpStatus field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetFloatingIpStatus() string {
	if o == nil || IsNil(o.FloatingIpStatus) {
		var ret string
		return ret
	}
	return *o.FloatingIpStatus
}

// GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetFloatingIpStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FloatingIpStatus) {
		return nil, false
	}
	return o.FloatingIpStatus, true
}

// HasFloatingIpStatus returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasFloatingIpStatus() bool {
	if o != nil && !IsNil(o.FloatingIpStatus) {
		return true
	}

	return false
}

// SetFloatingIpStatus gets a reference to the given string and assigns it to the FloatingIpStatus field.
func (o *InstancesSummaryFields) SetFloatingIpStatus(v string) {
	o.FloatingIpStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InstancesSummaryFields) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InstancesSummaryFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InstancesSummaryFields) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstancesSummaryFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InstancesSummaryFields) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InstancesSummaryFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InstancesSummaryFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstancesSummaryFields) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImageId) {
		toSerialize["image_id"] = o.ImageId
	}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	if !IsNil(o.FlavorId) {
		toSerialize["flavor_id"] = o.FlavorId
	}
	if !IsNil(o.Keypair) {
		toSerialize["keypair"] = o.Keypair
	}
	if !IsNil(o.KeypairId) {
		toSerialize["keypair_id"] = o.KeypairId
	}
	if !IsNil(o.FixedIp) {
		toSerialize["fixed_ip"] = o.FixedIp
	}
	if !IsNil(o.FloatingIp) {
		toSerialize["floating_ip"] = o.FloatingIp
	}
	if !IsNil(o.FloatingIpStatus) {
		toSerialize["floating_ip_status"] = o.FloatingIpStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableInstancesSummaryFields struct {
	value *InstancesSummaryFields
	isSet bool
}

func (v NullableInstancesSummaryFields) Get() *InstancesSummaryFields {
	return v.value
}

func (v *NullableInstancesSummaryFields) Set(val *InstancesSummaryFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInstancesSummaryFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInstancesSummaryFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstancesSummaryFields(val *InstancesSummaryFields) *NullableInstancesSummaryFields {
	return &NullableInstancesSummaryFields{value: val, isSet: true}
}

func (v NullableInstancesSummaryFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstancesSummaryFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

