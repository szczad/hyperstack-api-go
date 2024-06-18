# \VirtualMachineAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFirewallRuleToVirtualMachine**](VirtualMachineAPI.md#AddFirewallRuleToVirtualMachine) | **Post** /core/virtual-machines/{id}/sg-rules | Add firewall rule to virtual machine
[**AttachFirewallsToAVM**](VirtualMachineAPI.md#AttachFirewallsToAVM) | **Post** /core/virtual-machines/{vm_id}/attach-firewalls | Attach Firewalls to a VM
[**CreateVirtualMachine**](VirtualMachineAPI.md#CreateVirtualMachine) | **Post** /core/virtual-machines | Create virtual machine
[**DeleteFirewallRuleFromVirtualMachine**](VirtualMachineAPI.md#DeleteFirewallRuleFromVirtualMachine) | **Delete** /core/virtual-machines/{virtual_machine_id}/sg-rules/{sg_rule_id} | Delete firewall rule from virtual machine
[**DeleteVirtualMachine**](VirtualMachineAPI.md#DeleteVirtualMachine) | **Delete** /core/virtual-machines/{id} | Delete virtual machine
[**EditLabelsOfAnExistingVM**](VirtualMachineAPI.md#EditLabelsOfAnExistingVM) | **Put** /core/virtual-machines/{virtual_machine_id}/label | Edit labels of an existing VM
[**HardRebootVirtualMachine**](VirtualMachineAPI.md#HardRebootVirtualMachine) | **Get** /core/virtual-machines/{id}/hard-reboot | Hard reboot virtual machine
[**HibernateVirtualMachine**](VirtualMachineAPI.md#HibernateVirtualMachine) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate | Hibernate virtual machine
[**ListVirtualMachines**](VirtualMachineAPI.md#ListVirtualMachines) | **Get** /core/virtual-machines | List virtual machines
[**ResizeVirtualMachine**](VirtualMachineAPI.md#ResizeVirtualMachine) | **Post** /core/virtual-machines/{virtual_machine_id}/resize | Resize virtual machine
[**RestoreVirtualMachineFromHibernation**](VirtualMachineAPI.md#RestoreVirtualMachineFromHibernation) | **Get** /core/virtual-machines/{virtual_machine_id}/hibernate-restore | Restore virtual machine from hibernation
[**RetrieveVirtualMachineDetails**](VirtualMachineAPI.md#RetrieveVirtualMachineDetails) | **Get** /core/virtual-machines/{id} | Retrieve virtual machine details
[**RetrieveVirtualMachinePerformanceMetrics**](VirtualMachineAPI.md#RetrieveVirtualMachinePerformanceMetrics) | **Get** /core/virtual-machines/{virtual_machine_id}/metrics | Retrieve virtual machine performance metrics
[**RetrieveVirtualMachinesAssociatedWithAContract**](VirtualMachineAPI.md#RetrieveVirtualMachinesAssociatedWithAContract) | **Get** /core/virtual-machines/contract/{contract_id}/virtual-machines | Retrieve virtual machines associated with a contract
[**StartVirtualMachine**](VirtualMachineAPI.md#StartVirtualMachine) | **Get** /core/virtual-machines/{id}/start | Start virtual machine
[**StopVirtualMachine**](VirtualMachineAPI.md#StopVirtualMachine) | **Get** /core/virtual-machines/{id}/stop | Stop virtual machine



## AddFirewallRuleToVirtualMachine

> SecurityGroupRule AddFirewallRuleToVirtualMachine(ctx, id).Payload(payload).Execute()

Add firewall rule to virtual machine



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
	payload := *openapiclient.NewCreateSecurityRulePayload("Direction_example", "any", "Ethertype_example", "RemoteIpPrefix_example") // CreateSecurityRulePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.AddFirewallRuleToVirtualMachine(context.Background(), id).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.AddFirewallRuleToVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFirewallRuleToVirtualMachine`: SecurityGroupRule
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.AddFirewallRuleToVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFirewallRuleToVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**CreateSecurityRulePayload**](CreateSecurityRulePayload.md) |  | 

### Return type

[**SecurityGroupRule**](SecurityGroupRule.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachFirewallsToAVM

> ResponseModel AttachFirewallsToAVM(ctx, vmId).Payload(payload).Execute()

Attach Firewalls to a VM



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
	vmId := int32(56) // int32 | 
	payload := *openapiclient.NewAttachFirewallsToVMPayload([]int32{int32(123)}) // AttachFirewallsToVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.AttachFirewallsToAVM(context.Background(), vmId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.AttachFirewallsToAVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachFirewallsToAVM`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.AttachFirewallsToAVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachFirewallsToAVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**AttachFirewallsToVMPayload**](AttachFirewallsToVMPayload.md) |  | 

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


## CreateVirtualMachine

> CreateInstancesResponse CreateVirtualMachine(ctx).Payload(payload).Execute()

Create virtual machine



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
	payload := *openapiclient.NewCreateInstancesPayload("Name_example", "EnvironmentName_example", "FlavorName_example", "KeyName_example", int32(123)) // CreateInstancesPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.CreateVirtualMachine(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.CreateVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVirtualMachine`: CreateInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.CreateVirtualMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**CreateInstancesPayload**](CreateInstancesPayload.md) |  | 

### Return type

[**CreateInstancesResponse**](CreateInstancesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRuleFromVirtualMachine

> ResponseModel DeleteFirewallRuleFromVirtualMachine(ctx, virtualMachineId, sgRuleId).Execute()

Delete firewall rule from virtual machine



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
	sgRuleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine(context.Background(), virtualMachineId, sgRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallRuleFromVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteFirewallRuleFromVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 
**sgRuleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleFromVirtualMachineRequest struct via the builder pattern


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


## DeleteVirtualMachine

> ResponseModel DeleteVirtualMachine(ctx, id).Execute()

Delete virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.DeleteVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.DeleteVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.DeleteVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualMachineRequest struct via the builder pattern


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


## EditLabelsOfAnExistingVM

> ResponseModel EditLabelsOfAnExistingVM(ctx, virtualMachineId).Payload(payload).Execute()

Edit labels of an existing VM

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
	payload := *openapiclient.NewEditlabelofanexistingVMPayload() // EditlabelofanexistingVMPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.EditLabelsOfAnExistingVM(context.Background(), virtualMachineId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.EditLabelsOfAnExistingVM``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditLabelsOfAnExistingVM`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.EditLabelsOfAnExistingVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditLabelsOfAnExistingVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**EditlabelofanexistingVMPayload**](EditlabelofanexistingVMPayload.md) |  | 

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


## HardRebootVirtualMachine

> ResponseModel HardRebootVirtualMachine(ctx, id).Execute()

Hard reboot virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.HardRebootVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.HardRebootVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HardRebootVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.HardRebootVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardRebootVirtualMachineRequest struct via the builder pattern


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


## HibernateVirtualMachine

> ResponseModel HibernateVirtualMachine(ctx, virtualMachineId).Execute()

Hibernate virtual machine



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.HibernateVirtualMachine(context.Background(), virtualMachineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.HibernateVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HibernateVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.HibernateVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiHibernateVirtualMachineRequest struct via the builder pattern


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


## ListVirtualMachines

> Instances ListVirtualMachines(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

List virtual machines



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
	page := TODO // interface{} | Page Number (optional)
	pageSize := TODO // interface{} | Data Per Page (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.ListVirtualMachines(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.ListVirtualMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachines`: Instances
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.ListVirtualMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | [**interface{}**](interface{}.md) | Page Number | 
 **pageSize** | [**interface{}**](interface{}.md) | Data Per Page | 
 **search** | **string** |  | 

### Return type

[**Instances**](Instances.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVirtualMachine

> ResponseModel ResizeVirtualMachine(ctx, virtualMachineId).Payload(payload).Execute()

Resize virtual machine



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
	payload := *openapiclient.NewInstanceResizePayload() // InstanceResizePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.ResizeVirtualMachine(context.Background(), virtualMachineId).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.ResizeVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResizeVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.ResizeVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | [**InstanceResizePayload**](InstanceResizePayload.md) |  | 

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


## RestoreVirtualMachineFromHibernation

> ResponseModel RestoreVirtualMachineFromHibernation(ctx, virtualMachineId).Execute()

Restore virtual machine from hibernation



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.RestoreVirtualMachineFromHibernation(context.Background(), virtualMachineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RestoreVirtualMachineFromHibernation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreVirtualMachineFromHibernation`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RestoreVirtualMachineFromHibernation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreVirtualMachineFromHibernationRequest struct via the builder pattern


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


## RetrieveVirtualMachineDetails

> Instance RetrieveVirtualMachineDetails(ctx, id).Execute()

Retrieve virtual machine details



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
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachineDetails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachineDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachineDetails`: Instance
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachineDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVirtualMachineDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Instance**](Instance.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVirtualMachinePerformanceMetrics

> MetricsFields RetrieveVirtualMachinePerformanceMetrics(ctx, virtualMachineId).Duration(duration).Execute()

Retrieve virtual machine performance metrics



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
	duration := "duration_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics(context.Background(), virtualMachineId).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachinePerformanceMetrics`: MetricsFields
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachinePerformanceMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualMachineId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVirtualMachinePerformanceMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **string** |  | 

### Return type

[**MetricsFields**](MetricsFields.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveVirtualMachinesAssociatedWithAContract

> ContractInstancesResponse RetrieveVirtualMachinesAssociatedWithAContract(ctx, contractId).Execute()

Retrieve virtual machines associated with a contract

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
	contractId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract(context.Background(), contractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveVirtualMachinesAssociatedWithAContract`: ContractInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.RetrieveVirtualMachinesAssociatedWithAContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveVirtualMachinesAssociatedWithAContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContractInstancesResponse**](ContractInstancesResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVirtualMachine

> ResponseModel StartVirtualMachine(ctx, id).Execute()

Start virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.StartVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.StartVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.StartVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVirtualMachineRequest struct via the builder pattern


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


## StopVirtualMachine

> ResponseModel StopVirtualMachine(ctx, id).Execute()

Stop virtual machine



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
	resp, r, err := apiClient.VirtualMachineAPI.StopVirtualMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineAPI.StopVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopVirtualMachine`: ResponseModel
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineAPI.StopVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVirtualMachineRequest struct via the builder pattern


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

