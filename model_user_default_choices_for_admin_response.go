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

// checks if the UserDefaultChoicesForAdminResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDefaultChoicesForAdminResponse{}

// UserDefaultChoicesForAdminResponse struct for UserDefaultChoicesForAdminResponse
type UserDefaultChoicesForAdminResponse struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	UserDefaultChoices []UserDefaultChoiceForAdminFields `json:"user_default_choices,omitempty"`
}

// NewUserDefaultChoicesForAdminResponse instantiates a new UserDefaultChoicesForAdminResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDefaultChoicesForAdminResponse() *UserDefaultChoicesForAdminResponse {
	this := UserDefaultChoicesForAdminResponse{}
	return &this
}

// NewUserDefaultChoicesForAdminResponseWithDefaults instantiates a new UserDefaultChoicesForAdminResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDefaultChoicesForAdminResponseWithDefaults() *UserDefaultChoicesForAdminResponse {
	this := UserDefaultChoicesForAdminResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserDefaultChoicesForAdminResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoicesForAdminResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserDefaultChoicesForAdminResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *UserDefaultChoicesForAdminResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserDefaultChoicesForAdminResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoicesForAdminResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserDefaultChoicesForAdminResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserDefaultChoicesForAdminResponse) SetMessage(v string) {
	o.Message = &v
}

// GetUserDefaultChoices returns the UserDefaultChoices field value if set, zero value otherwise.
func (o *UserDefaultChoicesForAdminResponse) GetUserDefaultChoices() []UserDefaultChoiceForAdminFields {
	if o == nil || IsNil(o.UserDefaultChoices) {
		var ret []UserDefaultChoiceForAdminFields
		return ret
	}
	return o.UserDefaultChoices
}

// GetUserDefaultChoicesOk returns a tuple with the UserDefaultChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDefaultChoicesForAdminResponse) GetUserDefaultChoicesOk() ([]UserDefaultChoiceForAdminFields, bool) {
	if o == nil || IsNil(o.UserDefaultChoices) {
		return nil, false
	}
	return o.UserDefaultChoices, true
}

// HasUserDefaultChoices returns a boolean if a field has been set.
func (o *UserDefaultChoicesForAdminResponse) HasUserDefaultChoices() bool {
	if o != nil && !IsNil(o.UserDefaultChoices) {
		return true
	}

	return false
}

// SetUserDefaultChoices gets a reference to the given []UserDefaultChoiceForAdminFields and assigns it to the UserDefaultChoices field.
func (o *UserDefaultChoicesForAdminResponse) SetUserDefaultChoices(v []UserDefaultChoiceForAdminFields) {
	o.UserDefaultChoices = v
}

func (o UserDefaultChoicesForAdminResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDefaultChoicesForAdminResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.UserDefaultChoices) {
		toSerialize["user_default_choices"] = o.UserDefaultChoices
	}
	return toSerialize, nil
}

type NullableUserDefaultChoicesForAdminResponse struct {
	value *UserDefaultChoicesForAdminResponse
	isSet bool
}

func (v NullableUserDefaultChoicesForAdminResponse) Get() *UserDefaultChoicesForAdminResponse {
	return v.value
}

func (v *NullableUserDefaultChoicesForAdminResponse) Set(val *UserDefaultChoicesForAdminResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDefaultChoicesForAdminResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDefaultChoicesForAdminResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDefaultChoicesForAdminResponse(val *UserDefaultChoicesForAdminResponse) *NullableUserDefaultChoicesForAdminResponse {
	return &NullableUserDefaultChoicesForAdminResponse{value: val, isSet: true}
}

func (v NullableUserDefaultChoicesForAdminResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDefaultChoicesForAdminResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


