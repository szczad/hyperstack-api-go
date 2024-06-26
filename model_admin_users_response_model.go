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

// checks if the AdminUsersResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminUsersResponseModel{}

// AdminUsersResponseModel struct for AdminUsersResponseModel
type AdminUsersResponseModel struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Users []AdminUserFields `json:"users,omitempty"`
}

// NewAdminUsersResponseModel instantiates a new AdminUsersResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUsersResponseModel() *AdminUsersResponseModel {
	this := AdminUsersResponseModel{}
	return &this
}

// NewAdminUsersResponseModelWithDefaults instantiates a new AdminUsersResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUsersResponseModelWithDefaults() *AdminUsersResponseModel {
	this := AdminUsersResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminUsersResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUsersResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminUsersResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AdminUsersResponseModel) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AdminUsersResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUsersResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AdminUsersResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AdminUsersResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *AdminUsersResponseModel) GetUsers() []AdminUserFields {
	if o == nil || IsNil(o.Users) {
		var ret []AdminUserFields
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUsersResponseModel) GetUsersOk() ([]AdminUserFields, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *AdminUsersResponseModel) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []AdminUserFields and assigns it to the Users field.
func (o *AdminUsersResponseModel) SetUsers(v []AdminUserFields) {
	o.Users = v
}

func (o AdminUsersResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminUsersResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableAdminUsersResponseModel struct {
	value *AdminUsersResponseModel
	isSet bool
}

func (v NullableAdminUsersResponseModel) Get() *AdminUsersResponseModel {
	return v.value
}

func (v *NullableAdminUsersResponseModel) Set(val *AdminUsersResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUsersResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUsersResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUsersResponseModel(val *AdminUsersResponseModel) *NullableAdminUsersResponseModel {
	return &NullableAdminUsersResponseModel{value: val, isSet: true}
}

func (v NullableAdminUsersResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUsersResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


