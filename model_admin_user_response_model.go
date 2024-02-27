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

// checks if the AdminUserResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminUserResponseModel{}

// AdminUserResponseModel struct for AdminUserResponseModel
type AdminUserResponseModel struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	User *AdminUserFields `json:"user,omitempty"`
}

// NewAdminUserResponseModel instantiates a new AdminUserResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminUserResponseModel() *AdminUserResponseModel {
	this := AdminUserResponseModel{}
	return &this
}

// NewAdminUserResponseModelWithDefaults instantiates a new AdminUserResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminUserResponseModelWithDefaults() *AdminUserResponseModel {
	this := AdminUserResponseModel{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminUserResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminUserResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AdminUserResponseModel) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AdminUserResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AdminUserResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AdminUserResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AdminUserResponseModel) GetUser() AdminUserFields {
	if o == nil || IsNil(o.User) {
		var ret AdminUserFields
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserResponseModel) GetUserOk() (*AdminUserFields, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AdminUserResponseModel) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given AdminUserFields and assigns it to the User field.
func (o *AdminUserResponseModel) SetUser(v AdminUserFields) {
	o.User = &v
}

func (o AdminUserResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminUserResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableAdminUserResponseModel struct {
	value *AdminUserResponseModel
	isSet bool
}

func (v NullableAdminUserResponseModel) Get() *AdminUserResponseModel {
	return v.value
}

func (v *NullableAdminUserResponseModel) Set(val *AdminUserResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminUserResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminUserResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminUserResponseModel(val *AdminUserResponseModel) *NullableAdminUserResponseModel {
	return &NullableAdminUserResponseModel{value: val, isSet: true}
}

func (v NullableAdminUserResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminUserResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

