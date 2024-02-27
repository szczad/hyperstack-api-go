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

// checks if the FlavorAdminResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlavorAdminResponse{}

// FlavorAdminResponse struct for FlavorAdminResponse
type FlavorAdminResponse struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Flavor *FlavorAdminResponseFlavors `json:"flavor,omitempty"`
}

// NewFlavorAdminResponse instantiates a new FlavorAdminResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorAdminResponse() *FlavorAdminResponse {
	this := FlavorAdminResponse{}
	return &this
}

// NewFlavorAdminResponseWithDefaults instantiates a new FlavorAdminResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorAdminResponseWithDefaults() *FlavorAdminResponse {
	this := FlavorAdminResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FlavorAdminResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAdminResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FlavorAdminResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *FlavorAdminResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FlavorAdminResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAdminResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FlavorAdminResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FlavorAdminResponse) SetMessage(v string) {
	o.Message = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *FlavorAdminResponse) GetFlavor() FlavorAdminResponseFlavors {
	if o == nil || IsNil(o.Flavor) {
		var ret FlavorAdminResponseFlavors
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlavorAdminResponse) GetFlavorOk() (*FlavorAdminResponseFlavors, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *FlavorAdminResponse) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given FlavorAdminResponseFlavors and assigns it to the Flavor field.
func (o *FlavorAdminResponse) SetFlavor(v FlavorAdminResponseFlavors) {
	o.Flavor = &v
}

func (o FlavorAdminResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlavorAdminResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	return toSerialize, nil
}

type NullableFlavorAdminResponse struct {
	value *FlavorAdminResponse
	isSet bool
}

func (v NullableFlavorAdminResponse) Get() *FlavorAdminResponse {
	return v.value
}

func (v *NullableFlavorAdminResponse) Set(val *FlavorAdminResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorAdminResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorAdminResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorAdminResponse(val *FlavorAdminResponse) *NullableFlavorAdminResponse {
	return &NullableFlavorAdminResponse{value: val, isSet: true}
}

func (v NullableFlavorAdminResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorAdminResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

