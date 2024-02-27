/*
Infrahub-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateSecurityRulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSecurityRulePayload{}

// CreateSecurityRulePayload struct for CreateSecurityRulePayload
type CreateSecurityRulePayload struct {
	// The direction of traffic that the firewall rule applies to.
	Direction string `json:"direction"`
	// The network protocol associated with the rule. Call the [`GET /core/sg-rules-protocols`](https://infrahub-api-doc.nexgencloud.com/#get-/core/sg-rules-protocols) endpoint to retrieve a list of permitted network protocols.
	Protocol string `json:"protocol"`
	// The Ethernet type associated with the rule.
	Ethertype string `json:"ethertype"`
	// The IP address range that is allowed to access the specified port. Use \"0.0.0.0/0\" to allow any IP address.
	RemoteIpPrefix string `json:"remote_ip_prefix"`
}

type _CreateSecurityRulePayload CreateSecurityRulePayload

// NewCreateSecurityRulePayload instantiates a new CreateSecurityRulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityRulePayload(direction string, protocol string, ethertype string, remoteIpPrefix string) *CreateSecurityRulePayload {
	this := CreateSecurityRulePayload{}
	this.Direction = direction
	this.Protocol = protocol
	this.Ethertype = ethertype
	this.RemoteIpPrefix = remoteIpPrefix
	return &this
}

// NewCreateSecurityRulePayloadWithDefaults instantiates a new CreateSecurityRulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityRulePayloadWithDefaults() *CreateSecurityRulePayload {
	this := CreateSecurityRulePayload{}
	return &this
}

// GetDirection returns the Direction field value
func (o *CreateSecurityRulePayload) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityRulePayload) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *CreateSecurityRulePayload) SetDirection(v string) {
	o.Direction = v
}

// GetProtocol returns the Protocol field value
func (o *CreateSecurityRulePayload) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityRulePayload) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateSecurityRulePayload) SetProtocol(v string) {
	o.Protocol = v
}

// GetEthertype returns the Ethertype field value
func (o *CreateSecurityRulePayload) GetEthertype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ethertype
}

// GetEthertypeOk returns a tuple with the Ethertype field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityRulePayload) GetEthertypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ethertype, true
}

// SetEthertype sets field value
func (o *CreateSecurityRulePayload) SetEthertype(v string) {
	o.Ethertype = v
}

// GetRemoteIpPrefix returns the RemoteIpPrefix field value
func (o *CreateSecurityRulePayload) GetRemoteIpPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteIpPrefix
}

// GetRemoteIpPrefixOk returns a tuple with the RemoteIpPrefix field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityRulePayload) GetRemoteIpPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteIpPrefix, true
}

// SetRemoteIpPrefix sets field value
func (o *CreateSecurityRulePayload) SetRemoteIpPrefix(v string) {
	o.RemoteIpPrefix = v
}

func (o CreateSecurityRulePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSecurityRulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direction"] = o.Direction
	toSerialize["protocol"] = o.Protocol
	toSerialize["ethertype"] = o.Ethertype
	toSerialize["remote_ip_prefix"] = o.RemoteIpPrefix
	return toSerialize, nil
}

func (o *CreateSecurityRulePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"direction",
		"protocol",
		"ethertype",
		"remote_ip_prefix",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateSecurityRulePayload := _CreateSecurityRulePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSecurityRulePayload)

	if err != nil {
		return err
	}

	*o = CreateSecurityRulePayload(varCreateSecurityRulePayload)

	return err
}

type NullableCreateSecurityRulePayload struct {
	value *CreateSecurityRulePayload
	isSet bool
}

func (v NullableCreateSecurityRulePayload) Get() *CreateSecurityRulePayload {
	return v.value
}

func (v *NullableCreateSecurityRulePayload) Set(val *CreateSecurityRulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityRulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityRulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityRulePayload(val *CreateSecurityRulePayload) *NullableCreateSecurityRulePayload {
	return &NullableCreateSecurityRulePayload{value: val, isSet: true}
}

func (v NullableCreateSecurityRulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityRulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

