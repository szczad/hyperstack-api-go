# \FirewallAttachmentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachFirewallToVMs**](FirewallAttachmentAPI.md#AttachFirewallToVMs) | **Post** /core/firewalls/{firewall_id}/update-attachments | Attach Firewalls to VMs



## AttachFirewallToVMs

> ResponseModel AttachFirewallToVMs(ctx, firewallId).Payload(payload).Execute()

Attach Firewalls to VMs

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
	firewallId := int32(56) // int32 | 
	payload := *openapiclient.NewAttachFirewallWithVM([]int32{int32(123)}) // AttachFirewallWithVM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAttachmentAPI.AttachFirewallToVMs(context.Background(), firewallId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAttachmentAPI.AttachFirewallToVMs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachFirewallToVMs`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FirewallAttachmentAPI.AttachFirewallToVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachFirewallToVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachFirewallWithVM**](AttachFirewallWithVM.md) |  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

