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

// checks if the PricebookResourceObjectResponseForCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PricebookResourceObjectResponseForCustomer{}

// PricebookResourceObjectResponseForCustomer struct for PricebookResourceObjectResponseForCustomer
type PricebookResourceObjectResponseForCustomer struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	Rate *float32 `json:"rate,omitempty"`
	DiscountedRate *float32 `json:"discounted_rate,omitempty"`
	Price *float32 `json:"price,omitempty"`
	ActualPrice *float32 `json:"actual_price,omitempty"`
}

// NewPricebookResourceObjectResponseForCustomer instantiates a new PricebookResourceObjectResponseForCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPricebookResourceObjectResponseForCustomer() *PricebookResourceObjectResponseForCustomer {
	this := PricebookResourceObjectResponseForCustomer{}
	return &this
}

// NewPricebookResourceObjectResponseForCustomerWithDefaults instantiates a new PricebookResourceObjectResponseForCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricebookResourceObjectResponseForCustomerWithDefaults() *PricebookResourceObjectResponseForCustomer {
	this := PricebookResourceObjectResponseForCustomer{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PricebookResourceObjectResponseForCustomer) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PricebookResourceObjectResponseForCustomer) SetName(v string) {
	o.Name = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *PricebookResourceObjectResponseForCustomer) SetAmount(v int32) {
	o.Amount = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *PricebookResourceObjectResponseForCustomer) SetRate(v float32) {
	o.Rate = &v
}

// GetDiscountedRate returns the DiscountedRate field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetDiscountedRate() float32 {
	if o == nil || IsNil(o.DiscountedRate) {
		var ret float32
		return ret
	}
	return *o.DiscountedRate
}

// GetDiscountedRateOk returns a tuple with the DiscountedRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetDiscountedRateOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountedRate) {
		return nil, false
	}
	return o.DiscountedRate, true
}

// HasDiscountedRate returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasDiscountedRate() bool {
	if o != nil && !IsNil(o.DiscountedRate) {
		return true
	}

	return false
}

// SetDiscountedRate gets a reference to the given float32 and assigns it to the DiscountedRate field.
func (o *PricebookResourceObjectResponseForCustomer) SetDiscountedRate(v float32) {
	o.DiscountedRate = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PricebookResourceObjectResponseForCustomer) SetPrice(v float32) {
	o.Price = &v
}

// GetActualPrice returns the ActualPrice field value if set, zero value otherwise.
func (o *PricebookResourceObjectResponseForCustomer) GetActualPrice() float32 {
	if o == nil || IsNil(o.ActualPrice) {
		var ret float32
		return ret
	}
	return *o.ActualPrice
}

// GetActualPriceOk returns a tuple with the ActualPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PricebookResourceObjectResponseForCustomer) GetActualPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ActualPrice) {
		return nil, false
	}
	return o.ActualPrice, true
}

// HasActualPrice returns a boolean if a field has been set.
func (o *PricebookResourceObjectResponseForCustomer) HasActualPrice() bool {
	if o != nil && !IsNil(o.ActualPrice) {
		return true
	}

	return false
}

// SetActualPrice gets a reference to the given float32 and assigns it to the ActualPrice field.
func (o *PricebookResourceObjectResponseForCustomer) SetActualPrice(v float32) {
	o.ActualPrice = &v
}

func (o PricebookResourceObjectResponseForCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PricebookResourceObjectResponseForCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.DiscountedRate) {
		toSerialize["discounted_rate"] = o.DiscountedRate
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.ActualPrice) {
		toSerialize["actual_price"] = o.ActualPrice
	}
	return toSerialize, nil
}

type NullablePricebookResourceObjectResponseForCustomer struct {
	value *PricebookResourceObjectResponseForCustomer
	isSet bool
}

func (v NullablePricebookResourceObjectResponseForCustomer) Get() *PricebookResourceObjectResponseForCustomer {
	return v.value
}

func (v *NullablePricebookResourceObjectResponseForCustomer) Set(val *PricebookResourceObjectResponseForCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullablePricebookResourceObjectResponseForCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullablePricebookResourceObjectResponseForCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePricebookResourceObjectResponseForCustomer(val *PricebookResourceObjectResponseForCustomer) *NullablePricebookResourceObjectResponseForCustomer {
	return &NullablePricebookResourceObjectResponseForCustomer{value: val, isSet: true}
}

func (v NullablePricebookResourceObjectResponseForCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePricebookResourceObjectResponseForCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


