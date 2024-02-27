# \VncUrlAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVNCConsoleLink**](VncUrlAPI.md#GetVNCConsoleLink) | **Get** /core/virtual-machines/{virtual_machine_id}/console/{job_id} | Get VNC Console Link
[**RequestConsole**](VncUrlAPI.md#RequestConsole) | **Get** /core/virtual-machines/{id}/request-console | Request Instance Console



## GetVNCConsoleLink

> VNCURL GetVNCConsoleLink(ctx, virtualMachineId, jobId).Execute()

Get VNC Console Link

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
	virtualMachineId := int32(56) // int32 | 
	jobId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VncUrlAPI.GetVNCConsoleLink(context.Background(), virtualMachineId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.GetVNCConsoleLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVNCConsoleLink`: VNCURL
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.GetVNCConsoleLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 
**jobId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVNCConsoleLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VNCURL**](VNCURL.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestConsole

> RequestConsole RequestConsole(ctx, id).Execute()

Request Instance Console

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
	resp, r, err := apiClient.VncUrlAPI.RequestConsole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VncUrlAPI.RequestConsole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestConsole`: RequestConsole
	fmt.Fprintf(os.Stdout, "Response from `VncUrlAPI.RequestConsole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestConsoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestConsole**](RequestConsole.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

