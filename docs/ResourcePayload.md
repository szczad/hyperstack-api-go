# ResourcePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **int32** |  | 
**DiscountPercent** | Pointer to **float32** |  | [optional] 
**DiscountType** | **string** |  | 
**DiscountAmount** | Pointer to **float32** |  | [optional] 
**ResourceCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResourcePayload

`func NewResourcePayload(resourceId int32, discountType string, ) *ResourcePayload`

NewResourcePayload instantiates a new ResourcePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePayloadWithDefaults

`func NewResourcePayloadWithDefaults() *ResourcePayload`

NewResourcePayloadWithDefaults instantiates a new ResourcePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourcePayload) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourcePayload) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourcePayload) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetDiscountPercent

`func (o *ResourcePayload) GetDiscountPercent() float32`

GetDiscountPercent returns the DiscountPercent field if non-nil, zero value otherwise.

### GetDiscountPercentOk

`func (o *ResourcePayload) GetDiscountPercentOk() (*float32, bool)`

GetDiscountPercentOk returns a tuple with the DiscountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercent

`func (o *ResourcePayload) SetDiscountPercent(v float32)`

SetDiscountPercent sets DiscountPercent field to given value.

### HasDiscountPercent

`func (o *ResourcePayload) HasDiscountPercent() bool`

HasDiscountPercent returns a boolean if a field has been set.

### GetDiscountType

`func (o *ResourcePayload) GetDiscountType() string`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *ResourcePayload) GetDiscountTypeOk() (*string, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *ResourcePayload) SetDiscountType(v string)`

SetDiscountType sets DiscountType field to given value.


### GetDiscountAmount

`func (o *ResourcePayload) GetDiscountAmount() float32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *ResourcePayload) GetDiscountAmountOk() (*float32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *ResourcePayload) SetDiscountAmount(v float32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *ResourcePayload) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetResourceCount

`func (o *ResourcePayload) GetResourceCount() int32`

GetResourceCount returns the ResourceCount field if non-nil, zero value otherwise.

### GetResourceCountOk

`func (o *ResourcePayload) GetResourceCountOk() (*int32, bool)`

GetResourceCountOk returns a tuple with the ResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCount

`func (o *ResourcePayload) SetResourceCount(v int32)`

SetResourceCount sets ResourceCount field to given value.

### HasResourceCount

`func (o *ResourcePayload) HasResourceCount() bool`

HasResourceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


