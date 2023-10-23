# \OfficialApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficialEmployees20230301**](OfficialApi.md#OfficialEmployees20230301) | **Get** /2023-03-01/official/employees | Employees
[**Officialdentity20230301**](OfficialApi.md#Officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity



## OfficialEmployees20230301

> []Employee20230301Response OfficialEmployees20230301(ctx).Execute()

Employees



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
    resp, r, err := api_client.OfficialApi.OfficialEmployees20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialEmployees20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialEmployees20230301`: []Employee20230301Response
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialEmployees20230301Request struct via the builder pattern


### Return type

[**[]Employee20230301Response**](Employee20230301Response.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Officialdentity20230301

> IdentityResponse Officialdentity20230301(ctx).Execute()

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
    resp, r, err := api_client.OfficialApi.Officialdentity20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.Officialdentity20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Officialdentity20230301`: IdentityResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.Officialdentity20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialdentity20230301Request struct via the builder pattern


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

