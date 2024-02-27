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

// checks if the PermissionFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionFields{}

// PermissionFields struct for PermissionFields
type PermissionFields struct {
	Id *int32 `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Permission *string `json:"permission,omitempty"`
	Method *string `json:"method,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewPermissionFields instantiates a new PermissionFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionFields() *PermissionFields {
	this := PermissionFields{}
	return &this
}

// NewPermissionFieldsWithDefaults instantiates a new PermissionFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionFieldsWithDefaults() *PermissionFields {
	this := PermissionFields{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PermissionFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PermissionFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PermissionFields) SetId(v int32) {
	o.Id = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PermissionFields) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PermissionFields) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *PermissionFields) SetResource(v string) {
	o.Resource = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *PermissionFields) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *PermissionFields) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *PermissionFields) SetPermission(v string) {
	o.Permission = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PermissionFields) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PermissionFields) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PermissionFields) SetMethod(v string) {
	o.Method = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PermissionFields) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PermissionFields) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *PermissionFields) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PermissionFields) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PermissionFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PermissionFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o PermissionFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullablePermissionFields struct {
	value *PermissionFields
	isSet bool
}

func (v NullablePermissionFields) Get() *PermissionFields {
	return v.value
}

func (v *NullablePermissionFields) Set(val *PermissionFields) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionFields) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionFields(val *PermissionFields) *NullablePermissionFields {
	return &NullablePermissionFields{value: val, isSet: true}
}

func (v NullablePermissionFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

