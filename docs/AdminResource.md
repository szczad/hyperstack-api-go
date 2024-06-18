# AdminResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**AdminInstanceResources**](AdminInstanceResources.md) |  | [optional] 
**Volume** | Pointer to [**AdminVolumeResource**](AdminVolumeResource.md) |  | [optional] 
**Container** | Pointer to [**AdminContainerResource**](AdminContainerResource.md) |  | [optional] 
**Cluster** | Pointer to [**AdminClusterResource**](AdminClusterResource.md) |  | [optional] 

## Methods

### NewAdminResource

`func NewAdminResource() *AdminResource`

NewAdminResource instantiates a new AdminResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminResourceWithDefaults

`func NewAdminResourceWithDefaults() *AdminResource`

NewAdminResourceWithDefaults instantiates a new AdminResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AdminResource) GetInstance() AdminInstanceResources`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AdminResource) GetInstanceOk() (*AdminInstanceResources, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AdminResource) SetInstance(v AdminInstanceResources)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AdminResource) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetVolume

`func (o *AdminResource) GetVolume() AdminVolumeResource`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *AdminResource) GetVolumeOk() (*AdminVolumeResource, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *AdminResource) SetVolume(v AdminVolumeResource)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *AdminResource) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetContainer

`func (o *AdminResource) GetContainer() AdminContainerResource`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AdminResource) GetContainerOk() (*AdminContainerResource, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AdminResource) SetContainer(v AdminContainerResource)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *AdminResource) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetCluster

`func (o *AdminResource) GetCluster() AdminClusterResource`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AdminResource) GetClusterOk() (*AdminClusterResource, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AdminResource) SetCluster(v AdminClusterResource)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *AdminResource) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


