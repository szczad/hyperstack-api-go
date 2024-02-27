# AdminEnvrionmentResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Instances** | Pointer to [**[]AdminInstanceResources**](AdminInstanceResources.md) |  | [optional] 
**Volumes** | Pointer to [**[]AdminVolumeResource**](AdminVolumeResource.md) |  | [optional] 
**Containers** | Pointer to [**[]AdminContainerResource**](AdminContainerResource.md) |  | [optional] 
**Clusters** | Pointer to [**[]AdminClusterResource**](AdminClusterResource.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminEnvrionmentResources

`func NewAdminEnvrionmentResources() *AdminEnvrionmentResources`

NewAdminEnvrionmentResources instantiates a new AdminEnvrionmentResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminEnvrionmentResourcesWithDefaults

`func NewAdminEnvrionmentResourcesWithDefaults() *AdminEnvrionmentResources`

NewAdminEnvrionmentResourcesWithDefaults instantiates a new AdminEnvrionmentResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminEnvrionmentResources) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminEnvrionmentResources) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminEnvrionmentResources) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdminEnvrionmentResources) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminEnvrionmentResources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminEnvrionmentResources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminEnvrionmentResources) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminEnvrionmentResources) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *AdminEnvrionmentResources) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AdminEnvrionmentResources) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AdminEnvrionmentResources) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AdminEnvrionmentResources) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetInstances

`func (o *AdminEnvrionmentResources) GetInstances() []AdminInstanceResources`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *AdminEnvrionmentResources) GetInstancesOk() (*[]AdminInstanceResources, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *AdminEnvrionmentResources) SetInstances(v []AdminInstanceResources)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *AdminEnvrionmentResources) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetVolumes

`func (o *AdminEnvrionmentResources) GetVolumes() []AdminVolumeResource`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AdminEnvrionmentResources) GetVolumesOk() (*[]AdminVolumeResource, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AdminEnvrionmentResources) SetVolumes(v []AdminVolumeResource)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AdminEnvrionmentResources) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetContainers

`func (o *AdminEnvrionmentResources) GetContainers() []AdminContainerResource`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AdminEnvrionmentResources) GetContainersOk() (*[]AdminContainerResource, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AdminEnvrionmentResources) SetContainers(v []AdminContainerResource)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AdminEnvrionmentResources) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetClusters

`func (o *AdminEnvrionmentResources) GetClusters() []AdminClusterResource`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *AdminEnvrionmentResources) GetClustersOk() (*[]AdminClusterResource, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *AdminEnvrionmentResources) SetClusters(v []AdminClusterResource)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *AdminEnvrionmentResources) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminEnvrionmentResources) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminEnvrionmentResources) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminEnvrionmentResources) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminEnvrionmentResources) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


