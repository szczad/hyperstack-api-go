# \PaymentAPI

All URIs are relative to *https://infrahub-api.nexgencloud.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDetails**](PaymentAPI.md#GetDetails) | **Get** /billing/payment/payment-details | GET: View payment details
[**PostPayment**](PaymentAPI.md#PostPayment) | **Post** /billing/payment/payment-initiate | POST: Initiate payment



## GetDetails

> PaymentDetailsResponse GetDetails(ctx).Execute()

GET: View payment details

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
	resp, r, err := apiClient.PaymentAPI.GetDetails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.GetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetails`: PaymentDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.GetDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsRequest struct via the builder pattern


### Return type

[**PaymentDetailsResponse**](PaymentDetailsResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPayment

> PaymentInitiateResponse PostPayment(ctx).Payload(payload).Execute()

POST: Initiate payment

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
	payload := *openapiclient.NewPaymentInitiatePayload() // PaymentInitiatePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentAPI.PostPayment(context.Background()).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentAPI.PostPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPayment`: PaymentInitiateResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentAPI.PostPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payload** | [**PaymentInitiatePayload**](PaymentInitiatePayload.md) |  | 

### Return type

[**PaymentInitiateResponse**](PaymentInitiateResponse.md)

### Authorization

[apiKey](../README.md#apiKey), [accessToken](../README.md#accessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

