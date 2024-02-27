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

// checks if the NewConfigurationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewConfigurationsResponse{}

// NewConfigurationsResponse struct for NewConfigurationsResponse
type NewConfigurationsResponse struct {
	Var1x *int32 `json:"1x,omitempty"`
	Var2x *int32 `json:"2x,omitempty"`
	Var4x *int32 `json:"4x,omitempty"`
	Var8x *int32 `json:"8x,omitempty"`
	Var10x *int32 `json:"10x,omitempty"`
}

// NewNewConfigurationsResponse instantiates a new NewConfigurationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewConfigurationsResponse() *NewConfigurationsResponse {
	this := NewConfigurationsResponse{}
	return &this
}

// NewNewConfigurationsResponseWithDefaults instantiates a new NewConfigurationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewConfigurationsResponseWithDefaults() *NewConfigurationsResponse {
	this := NewConfigurationsResponse{}
	return &this
}

// GetVar1x returns the Var1x field value if set, zero value otherwise.
func (o *NewConfigurationsResponse) GetVar1x() int32 {
	if o == nil || IsNil(o.Var1x) {
		var ret int32
		return ret
	}
	return *o.Var1x
}

// GetVar1xOk returns a tuple with the Var1x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfigurationsResponse) GetVar1xOk() (*int32, bool) {
	if o == nil || IsNil(o.Var1x) {
		return nil, false
	}
	return o.Var1x, true
}

// HasVar1x returns a boolean if a field has been set.
func (o *NewConfigurationsResponse) HasVar1x() bool {
	if o != nil && !IsNil(o.Var1x) {
		return true
	}

	return false
}

// SetVar1x gets a reference to the given int32 and assigns it to the Var1x field.
func (o *NewConfigurationsResponse) SetVar1x(v int32) {
	o.Var1x = &v
}

// GetVar2x returns the Var2x field value if set, zero value otherwise.
func (o *NewConfigurationsResponse) GetVar2x() int32 {
	if o == nil || IsNil(o.Var2x) {
		var ret int32
		return ret
	}
	return *o.Var2x
}

// GetVar2xOk returns a tuple with the Var2x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfigurationsResponse) GetVar2xOk() (*int32, bool) {
	if o == nil || IsNil(o.Var2x) {
		return nil, false
	}
	return o.Var2x, true
}

// HasVar2x returns a boolean if a field has been set.
func (o *NewConfigurationsResponse) HasVar2x() bool {
	if o != nil && !IsNil(o.Var2x) {
		return true
	}

	return false
}

// SetVar2x gets a reference to the given int32 and assigns it to the Var2x field.
func (o *NewConfigurationsResponse) SetVar2x(v int32) {
	o.Var2x = &v
}

// GetVar4x returns the Var4x field value if set, zero value otherwise.
func (o *NewConfigurationsResponse) GetVar4x() int32 {
	if o == nil || IsNil(o.Var4x) {
		var ret int32
		return ret
	}
	return *o.Var4x
}

// GetVar4xOk returns a tuple with the Var4x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfigurationsResponse) GetVar4xOk() (*int32, bool) {
	if o == nil || IsNil(o.Var4x) {
		return nil, false
	}
	return o.Var4x, true
}

// HasVar4x returns a boolean if a field has been set.
func (o *NewConfigurationsResponse) HasVar4x() bool {
	if o != nil && !IsNil(o.Var4x) {
		return true
	}

	return false
}

// SetVar4x gets a reference to the given int32 and assigns it to the Var4x field.
func (o *NewConfigurationsResponse) SetVar4x(v int32) {
	o.Var4x = &v
}

// GetVar8x returns the Var8x field value if set, zero value otherwise.
func (o *NewConfigurationsResponse) GetVar8x() int32 {
	if o == nil || IsNil(o.Var8x) {
		var ret int32
		return ret
	}
	return *o.Var8x
}

// GetVar8xOk returns a tuple with the Var8x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfigurationsResponse) GetVar8xOk() (*int32, bool) {
	if o == nil || IsNil(o.Var8x) {
		return nil, false
	}
	return o.Var8x, true
}

// HasVar8x returns a boolean if a field has been set.
func (o *NewConfigurationsResponse) HasVar8x() bool {
	if o != nil && !IsNil(o.Var8x) {
		return true
	}

	return false
}

// SetVar8x gets a reference to the given int32 and assigns it to the Var8x field.
func (o *NewConfigurationsResponse) SetVar8x(v int32) {
	o.Var8x = &v
}

// GetVar10x returns the Var10x field value if set, zero value otherwise.
func (o *NewConfigurationsResponse) GetVar10x() int32 {
	if o == nil || IsNil(o.Var10x) {
		var ret int32
		return ret
	}
	return *o.Var10x
}

// GetVar10xOk returns a tuple with the Var10x field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewConfigurationsResponse) GetVar10xOk() (*int32, bool) {
	if o == nil || IsNil(o.Var10x) {
		return nil, false
	}
	return o.Var10x, true
}

// HasVar10x returns a boolean if a field has been set.
func (o *NewConfigurationsResponse) HasVar10x() bool {
	if o != nil && !IsNil(o.Var10x) {
		return true
	}

	return false
}

// SetVar10x gets a reference to the given int32 and assigns it to the Var10x field.
func (o *NewConfigurationsResponse) SetVar10x(v int32) {
	o.Var10x = &v
}

func (o NewConfigurationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewConfigurationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var1x) {
		toSerialize["1x"] = o.Var1x
	}
	if !IsNil(o.Var2x) {
		toSerialize["2x"] = o.Var2x
	}
	if !IsNil(o.Var4x) {
		toSerialize["4x"] = o.Var4x
	}
	if !IsNil(o.Var8x) {
		toSerialize["8x"] = o.Var8x
	}
	if !IsNil(o.Var10x) {
		toSerialize["10x"] = o.Var10x
	}
	return toSerialize, nil
}

type NullableNewConfigurationsResponse struct {
	value *NewConfigurationsResponse
	isSet bool
}

func (v NullableNewConfigurationsResponse) Get() *NewConfigurationsResponse {
	return v.value
}

func (v *NullableNewConfigurationsResponse) Set(val *NewConfigurationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfigurationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfigurationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfigurationsResponse(val *NewConfigurationsResponse) *NullableNewConfigurationsResponse {
	return &NullableNewConfigurationsResponse{value: val, isSet: true}
}

func (v NullableNewConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfigurationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


