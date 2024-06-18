# Excludebillingpostpayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **int32** | The ID of the resource which is being excluded from billing. | 
**Exclude** | **bool** | &#x60;true&#x60; excludes the resource from billing while &#x60;false&#x60; does not. | 

## Methods

### NewExcludebillingpostpayload

`func NewExcludebillingpostpayload(resourceId int32, exclude bool, ) *Excludebillingpostpayload`

NewExcludebillingpostpayload instantiates a new Excludebillingpostpayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludebillingpostpayloadWithDefaults

`func NewExcludebillingpostpayloadWithDefaults() *Excludebillingpostpayload`

NewExcludebillingpostpayloadWithDefaults instantiates a new Excludebillingpostpayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *Excludebillingpostpayload) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Excludebillingpostpayload) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Excludebillingpostpayload) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetExclude

`func (o *Excludebillingpostpayload) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Excludebillingpostpayload) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Excludebillingpostpayload) SetExclude(v bool)`

SetExclude sets Exclude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


