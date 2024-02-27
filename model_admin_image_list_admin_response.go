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

// checks if the AdminImageListAdminResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminImageListAdminResponse{}

// AdminImageListAdminResponse struct for AdminImageListAdminResponse
type AdminImageListAdminResponse struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Data []AdminImageAdminFields `json:"data,omitempty"`
}

// NewAdminImageListAdminResponse instantiates a new AdminImageListAdminResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminImageListAdminResponse() *AdminImageListAdminResponse {
	this := AdminImageListAdminResponse{}
	return &this
}

// NewAdminImageListAdminResponseWithDefaults instantiates a new AdminImageListAdminResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminImageListAdminResponseWithDefaults() *AdminImageListAdminResponse {
	this := AdminImageListAdminResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AdminImageListAdminResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminImageListAdminResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AdminImageListAdminResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *AdminImageListAdminResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AdminImageListAdminResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminImageListAdminResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AdminImageListAdminResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AdminImageListAdminResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AdminImageListAdminResponse) GetData() []AdminImageAdminFields {
	if o == nil || IsNil(o.Data) {
		var ret []AdminImageAdminFields
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminImageListAdminResponse) GetDataOk() ([]AdminImageAdminFields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AdminImageListAdminResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AdminImageAdminFields and assigns it to the Data field.
func (o *AdminImageListAdminResponse) SetData(v []AdminImageAdminFields) {
	o.Data = v
}

func (o AdminImageListAdminResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminImageListAdminResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAdminImageListAdminResponse struct {
	value *AdminImageListAdminResponse
	isSet bool
}

func (v NullableAdminImageListAdminResponse) Get() *AdminImageListAdminResponse {
	return v.value
}

func (v *NullableAdminImageListAdminResponse) Set(val *AdminImageListAdminResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminImageListAdminResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminImageListAdminResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminImageListAdminResponse(val *AdminImageListAdminResponse) *NullableAdminImageListAdminResponse {
	return &NullableAdminImageListAdminResponse{value: val, isSet: true}
}

func (v NullableAdminImageListAdminResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminImageListAdminResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


