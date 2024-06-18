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

// checks if the HistoricalInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricalInstance{}

// HistoricalInstance struct for HistoricalInstance
type HistoricalInstance struct {
	Status *bool `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Instances []HistoricalInstancesFields `json:"instances,omitempty"`
	InstanceCount *int32 `json:"instance_count,omitempty"`
}

// NewHistoricalInstance instantiates a new HistoricalInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalInstance() *HistoricalInstance {
	this := HistoricalInstance{}
	return &this
}

// NewHistoricalInstanceWithDefaults instantiates a new HistoricalInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalInstanceWithDefaults() *HistoricalInstance {
	this := HistoricalInstance{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HistoricalInstance) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalInstance) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HistoricalInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HistoricalInstance) SetStatus(v bool) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HistoricalInstance) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalInstance) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HistoricalInstance) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HistoricalInstance) SetMessage(v string) {
	o.Message = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *HistoricalInstance) GetInstances() []HistoricalInstancesFields {
	if o == nil || IsNil(o.Instances) {
		var ret []HistoricalInstancesFields
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalInstance) GetInstancesOk() ([]HistoricalInstancesFields, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *HistoricalInstance) HasInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []HistoricalInstancesFields and assigns it to the Instances field.
func (o *HistoricalInstance) SetInstances(v []HistoricalInstancesFields) {
	o.Instances = v
}

// GetInstanceCount returns the InstanceCount field value if set, zero value otherwise.
func (o *HistoricalInstance) GetInstanceCount() int32 {
	if o == nil || IsNil(o.InstanceCount) {
		var ret int32
		return ret
	}
	return *o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalInstance) GetInstanceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.InstanceCount) {
		return nil, false
	}
	return o.InstanceCount, true
}

// HasInstanceCount returns a boolean if a field has been set.
func (o *HistoricalInstance) HasInstanceCount() bool {
	if o != nil && !IsNil(o.InstanceCount) {
		return true
	}

	return false
}

// SetInstanceCount gets a reference to the given int32 and assigns it to the InstanceCount field.
func (o *HistoricalInstance) SetInstanceCount(v int32) {
	o.InstanceCount = &v
}

func (o HistoricalInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !IsNil(o.InstanceCount) {
		toSerialize["instance_count"] = o.InstanceCount
	}
	return toSerialize, nil
}

type NullableHistoricalInstance struct {
	value *HistoricalInstance
	isSet bool
}

func (v NullableHistoricalInstance) Get() *HistoricalInstance {
	return v.value
}

func (v *NullableHistoricalInstance) Set(val *HistoricalInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalInstance(val *HistoricalInstance) *NullableHistoricalInstance {
	return &NullableHistoricalInstance{value: val, isSet: true}
}

func (v NullableHistoricalInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


