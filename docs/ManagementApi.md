# \ManagementApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Client**](ManagementApi.md#Client) | **Get** /2023-03-01/management/client | View client
[**Disconnect**](ManagementApi.md#Disconnect) | **Post** /2023-03-01/management/disconnect | Disconnect token
[**Introspect**](ManagementApi.md#Introspect) | **Get** /2023-03-01/management/introspect | Inspect token
[**Token**](ManagementApi.md#Token) | **Post** /2023-03-01/management/token | Exchange &#x60;authorization_code&#x60; for &#x60;access_token&#x60;
[**Tokens**](ManagementApi.md#Tokens) | **Get** /2023-03-01/management/tokens | View tokens
[**UpdateClient**](ManagementApi.md#UpdateClient) | **Post** /2023-03-01/management/client | Update client



## Client

> ClientResponse Client(ctx).Execute()

View client



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
    resp, r, err := api_client.ManagementApi.Client(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.Client``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Client`: ClientResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.Client`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClientRequest struct via the builder pattern


### Return type

[**ClientResponse**](ClientResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Disconnect

> DisconnectResponse Disconnect(ctx).Execute()

Disconnect token



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
    resp, r, err := api_client.ManagementApi.Disconnect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.Disconnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Disconnect`: DisconnectResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.Disconnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectRequest struct via the builder pattern


### Return type

[**DisconnectResponse**](DisconnectResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Introspect

> IntrospectResponse Introspect(ctx).Execute()

Inspect token



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
    resp, r, err := api_client.ManagementApi.Introspect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.Introspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Introspect`: IntrospectResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.Introspect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntrospectRequest struct via the builder pattern


### Return type

[**IntrospectResponse**](IntrospectResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Token

> TokenResponse Token(ctx).TokenRequest(tokenRequest).Execute()

Exchange `authorization_code` for `access_token`



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
    tokenRequest := *openapiclient.NewTokenRequest("00000000-00000000-00000000-00000000", "ffffffff-ffffffff-ffffffff-ffffffff", "authorization_code", "Y2xpZW50IzkzMTU4MGQwLWYwYjctNGJiOC1iYmZmLWI4MTNlYzMxNTVjYXxjb2RlIzE1MmIwYjk3LTg2ZWMtNDZlNC1hZDUyLWY5ZTAxNzE2MDIwNAo=", "https://example.com") // TokenRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementApi.Token(context.Background()).TokenRequest(tokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.Token``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Token`: TokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tokens

> []map[string]interface{} Tokens(ctx).Execute()

View tokens



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
    resp, r, err := api_client.ManagementApi.Tokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.Tokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Tokens`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.Tokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTokensRequest struct via the builder pattern


### Return type

**[]map[string]interface{}**

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClient

> ClientResponse UpdateClient(ctx).ClientRequest(clientRequest).Execute()

Update client



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
    clientRequest := *openapiclient.NewClientRequest([]string{"ClientSecret_example"}, []string{"RedirectUris_example"}, "Your App") // ClientRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ManagementApi.UpdateClient(context.Background()).ClientRequest(clientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagementApi.UpdateClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClient`: ClientResponse
    fmt.Fprintf(os.Stdout, "Response from `ManagementApi.UpdateClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientRequest** | [**ClientRequest**](ClientRequest.md) |  | 

### Return type

[**ClientResponse**](ClientResponse.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

