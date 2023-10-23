# \RawApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmazonGetOrder**](RawApi.md#AmazonGetOrder) | **Get** /2023-03-20/raw/amazon/orders/{order_id} | Get an Order
[**AmazonIdentity**](RawApi.md#AmazonIdentity) | **Get** /2023-03-20/raw/amazon/identity | Identity
[**AmazonOrders**](RawApi.md#AmazonOrders) | **Get** /2023-03-20/raw/amazon/orders | Orders



## AmazonGetOrder

> Order20230320Response AmazonGetOrder(ctx, orderId).Execute()

Get an Order



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderId := "orderId_example" // string | The id of the order.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawApi.AmazonGetOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawApi.AmazonGetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmazonGetOrder`: Order20230320Response
    fmt.Fprintf(os.Stdout, "Response from `RawApi.AmazonGetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The id of the order. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAmazonGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Order20230320Response**](Order20230320Response.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AmazonIdentity

> IdentityResponse AmazonIdentity(ctx).Execute()

Identity



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawApi.AmazonIdentity(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawApi.AmazonIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmazonIdentity`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `RawApi.AmazonIdentity`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAmazonIdentityRequest struct via the builder pattern


### Return type

[**IdentityResponse**](IdentityResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AmazonOrders

> []map[string]interface{} AmazonOrders(ctx).StartDate(startDate).EndDate(endDate).Execute()

Orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    startDate := time.Now() // string | The start date of the search period
    endDate := time.Now() // string | The end date of the search period

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RawApi.AmazonOrders(context.Background()).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RawApi.AmazonOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AmazonOrders`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RawApi.AmazonOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAmazonOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | The start date of the search period | 
 **endDate** | **string** | The end date of the search period | 

### Return type

**[]map[string]interface{}**

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

