# \OfficialApi

All URIs are relative to *https://api.affixapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficialCompanies20230301**](OfficialApi.md#OfficialCompanies20230301) | **Get** /2023-03-01/official/company | Company
[**OfficialCreateEmployee20230301**](OfficialApi.md#OfficialCreateEmployee20230301) | **Post** /2023-03-01/official/employee | Create Employee
[**OfficialEmployees20230301**](OfficialApi.md#OfficialEmployees20230301) | **Get** /2023-03-01/official/employees | Employees
[**OfficialGroups20230301**](OfficialApi.md#OfficialGroups20230301) | **Get** /2023-03-01/official/groups | Groups
[**OfficialTimeOffBalances20230301**](OfficialApi.md#OfficialTimeOffBalances20230301) | **Get** /2023-03-01/official/time-off-balances | Time off balances
[**OfficialTimeOffEntries20230301**](OfficialApi.md#OfficialTimeOffEntries20230301) | **Get** /2023-03-01/official/time-off-entries | Time off entries
[**OfficialTimesheets20230301**](OfficialApi.md#OfficialTimesheets20230301) | **Get** /2023-03-01/official/timesheets | Timesheets
[**OfficialWorkLocations20230301**](OfficialApi.md#OfficialWorkLocations20230301) | **Get** /2023-03-01/official/work-locations | Work Locations
[**Officialdentity20230301**](OfficialApi.md#Officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity



## OfficialCompanies20230301

> []CompanyResponse OfficialCompanies20230301(ctx).Execute()

Company



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
    resp, r, err := api_client.OfficialApi.OfficialCompanies20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialCompanies20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialCompanies20230301`: []CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialCompanies20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialCompanies20230301Request struct via the builder pattern


### Return type

[**[]CompanyResponse**](CompanyResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialCreateEmployee20230301

> EmployeeResponse OfficialCreateEmployee20230301(ctx).CreateEmployeeRequest(createEmployeeRequest).Execute()

Create Employee



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
    createEmployeeRequest := *openapiclient.NewCreateEmployeeRequest("Greg", "Hirsch") // CreateEmployeeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OfficialApi.OfficialCreateEmployee20230301(context.Background()).CreateEmployeeRequest(createEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialCreateEmployee20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialCreateEmployee20230301`: EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialCreateEmployee20230301`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfficialCreateEmployee20230301Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmployeeRequest** | [**CreateEmployeeRequest**](CreateEmployeeRequest.md) |  | 

### Return type

[**EmployeeResponse**](EmployeeResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialEmployees20230301

> []EmployeeResponse OfficialEmployees20230301(ctx).Execute()

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
    // response from `OfficialEmployees20230301`: []EmployeeResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialEmployees20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialEmployees20230301Request struct via the builder pattern


### Return type

[**[]EmployeeResponse**](EmployeeResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialGroups20230301

> []GroupResponse OfficialGroups20230301(ctx).Execute()

Groups



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
    resp, r, err := api_client.OfficialApi.OfficialGroups20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialGroups20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialGroups20230301`: []GroupResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialGroups20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialGroups20230301Request struct via the builder pattern


### Return type

[**[]GroupResponse**](GroupResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialTimeOffBalances20230301

> []TimeOffBalanceResponse OfficialTimeOffBalances20230301(ctx).Execute()

Time off balances



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
    resp, r, err := api_client.OfficialApi.OfficialTimeOffBalances20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialTimeOffBalances20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialTimeOffBalances20230301`: []TimeOffBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialTimeOffBalances20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialTimeOffBalances20230301Request struct via the builder pattern


### Return type

[**[]TimeOffBalanceResponse**](TimeOffBalanceResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialTimeOffEntries20230301

> []TimeOffEntryResponse OfficialTimeOffEntries20230301(ctx).Execute()

Time off entries



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
    resp, r, err := api_client.OfficialApi.OfficialTimeOffEntries20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialTimeOffEntries20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialTimeOffEntries20230301`: []TimeOffEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialTimeOffEntries20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialTimeOffEntries20230301Request struct via the builder pattern


### Return type

[**[]TimeOffEntryResponse**](TimeOffEntryResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialTimesheets20230301

> []TimesheetResponse OfficialTimesheets20230301(ctx).Execute()

Timesheets



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
    resp, r, err := api_client.OfficialApi.OfficialTimesheets20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialTimesheets20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialTimesheets20230301`: []TimesheetResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialTimesheets20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialTimesheets20230301Request struct via the builder pattern


### Return type

[**[]TimesheetResponse**](TimesheetResponse.md)

### Authorization

[access-token](../README.md#access-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficialWorkLocations20230301

> []LocationResponse OfficialWorkLocations20230301(ctx).Execute()

Work Locations



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
    resp, r, err := api_client.OfficialApi.OfficialWorkLocations20230301(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OfficialApi.OfficialWorkLocations20230301``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficialWorkLocations20230301`: []LocationResponse
    fmt.Fprintf(os.Stdout, "Response from `OfficialApi.OfficialWorkLocations20230301`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOfficialWorkLocations20230301Request struct via the builder pattern


### Return type

[**[]LocationResponse**](LocationResponse.md)

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

