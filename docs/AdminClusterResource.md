# AdminClusterResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MasterCount** | Pointer to **int32** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodeFlavor** | Pointer to [**AdminFlavorResource**](AdminFlavorResource.md) |  | [optional] 
**EnablePublicIp** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminClusterResource

`func NewAdminClusterResource() *AdminClusterResource`

NewAdminClusterResource instantiates a new AdminClusterResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminClusterResourceWithDefaults

`func NewAdminClusterResourceWithDefaults() *AdminClusterResource`

NewAdminClusterResourceWithDefaults instantiates a new AdminClusterResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminClusterResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminClusterResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminClusterResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminClusterResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminClusterResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminClusterResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminClusterResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminClusterResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *AdminClusterResource) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *AdminClusterResource) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *AdminClusterResource) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *AdminClusterResource) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetStatus

`func (o *AdminClusterResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminClusterResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminClusterResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminClusterResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMasterCount

`func (o *AdminClusterResource) GetMasterCount() int32`

GetMasterCount returns the MasterCount field if non-nil, zero value otherwise.

### GetMasterCountOk

`func (o *AdminClusterResource) GetMasterCountOk() (*int32, bool)`

GetMasterCountOk returns a tuple with the MasterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCount

`func (o *AdminClusterResource) SetMasterCount(v int32)`

SetMasterCount sets MasterCount field to given value.

### HasMasterCount

`func (o *AdminClusterResource) HasMasterCount() bool`

HasMasterCount returns a boolean if a field has been set.

### GetNodeCount

`func (o *AdminClusterResource) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AdminClusterResource) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AdminClusterResource) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AdminClusterResource) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeFlavor

`func (o *AdminClusterResource) GetNodeFlavor() AdminFlavorResource`

GetNodeFlavor returns the NodeFlavor field if non-nil, zero value otherwise.

### GetNodeFlavorOk

`func (o *AdminClusterResource) GetNodeFlavorOk() (*AdminFlavorResource, bool)`

GetNodeFlavorOk returns a tuple with the NodeFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFlavor

`func (o *AdminClusterResource) SetNodeFlavor(v AdminFlavorResource)`

SetNodeFlavor sets NodeFlavor field to given value.

### HasNodeFlavor

`func (o *AdminClusterResource) HasNodeFlavor() bool`

HasNodeFlavor returns a boolean if a field has been set.

### GetEnablePublicIp

`func (o *AdminClusterResource) GetEnablePublicIp() bool`

GetEnablePublicIp returns the EnablePublicIp field if non-nil, zero value otherwise.

### GetEnablePublicIpOk

`func (o *AdminClusterResource) GetEnablePublicIpOk() (*bool, bool)`

GetEnablePublicIpOk returns a tuple with the EnablePublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicIp

`func (o *AdminClusterResource) SetEnablePublicIp(v bool)`

SetEnablePublicIp sets EnablePublicIp field to given value.

### HasEnablePublicIp

`func (o *AdminClusterResource) HasEnablePublicIp() bool`

HasEnablePublicIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminClusterResource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminClusterResource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminClusterResource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminClusterResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


