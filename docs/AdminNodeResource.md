# AdminNodeResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ProvisionDate** | Pointer to **time.Time** |  | [optional] 
**Organizations** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewAdminNodeResource

`func NewAdminNodeResource() *AdminNodeResource`

NewAdminNodeResource instantiates a new AdminNodeResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminNodeResourceWithDefaults

`func NewAdminNodeResourceWithDefaults() *AdminNodeResource`

NewAdminNodeResourceWithDefaults instantiates a new AdminNodeResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminNodeResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminNodeResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminNodeResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminNodeResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpenstackId

`func (o *AdminNodeResource) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *AdminNodeResource) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *AdminNodeResource) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *AdminNodeResource) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetName

`func (o *AdminNodeResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminNodeResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminNodeResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminNodeResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvailable

`func (o *AdminNodeResource) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AdminNodeResource) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AdminNodeResource) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AdminNodeResource) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetStatus

`func (o *AdminNodeResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminNodeResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminNodeResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminNodeResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvisionDate

`func (o *AdminNodeResource) GetProvisionDate() time.Time`

GetProvisionDate returns the ProvisionDate field if non-nil, zero value otherwise.

### GetProvisionDateOk

`func (o *AdminNodeResource) GetProvisionDateOk() (*time.Time, bool)`

GetProvisionDateOk returns a tuple with the ProvisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionDate

`func (o *AdminNodeResource) SetProvisionDate(v time.Time)`

SetProvisionDate sets ProvisionDate field to given value.

### HasProvisionDate

`func (o *AdminNodeResource) HasProvisionDate() bool`

HasProvisionDate returns a boolean if a field has been set.

### GetOrganizations

`func (o *AdminNodeResource) GetOrganizations() []int32`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *AdminNodeResource) GetOrganizationsOk() (*[]int32, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *AdminNodeResource) SetOrganizations(v []int32)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *AdminNodeResource) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


