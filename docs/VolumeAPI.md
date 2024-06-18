# \VolumeAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumeAPI.md#CreateVolume) | **Post** /core/volumes | Create volume
[**DeleteVolume**](VolumeAPI.md#DeleteVolume) | **Delete** /core/volumes/{id} | Delete volume
[**ListVolumeTypes**](VolumeAPI.md#ListVolumeTypes) | **Get** /core/volume-types | List volume types
[**ListVolumes**](VolumeAPI.md#ListVolumes) | **Get** /core/volumes | List volumes



## CreateVolume

> Volume CreateVolume(ctx).Payload(payload).Execute()

Create volume



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {
	payload := *openapiclient.NewCreateVolumePayload("Name_example", "EnvironmentName_example", "Description_example", "VolumeType_example", int32(123)) // CreateVolumePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateVolume(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateVolumePayload**](CreateVolumePayload.md) |  | 

### Return type

[**Volume**](Volume.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> ResponseModel DeleteVolume(ctx, id).Execute()

Delete volume



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.DeleteVolume(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolume`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeTypes

> VolumeTypes ListVolumeTypes(ctx).Execute()

List volume types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumeTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTypes`: VolumeTypes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypesRequest struct via the builder pattern


### Return type

[**VolumeTypes**](VolumeTypes.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> Volumes ListVolumes(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

List volumes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/szczad/hyperstack-api-go"
)

func main() {
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumes(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: Volumes
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**Volumes**](Volumes.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

