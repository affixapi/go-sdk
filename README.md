# Go API client for openapi

The affixapi.com API documentation.

# Introduction
Affix API is an OAuth 2.1 application that allows developers to access
customer data, without developers needing to manage or maintain
integrations; or collect login credentials or API keys from users for these
third party systems.

# OAuth 2.1
Affix API follows the [OAuth 2.1
spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).

As an OAuth application, Affix API handles not only both the collection of
sensitive user credentials or API keys, but also builds and maintains the
integrations with the providers, so you don't have to.

# How to obtain an access token
in order to get started, you must:
  - register a `client_id`
  - direct your user to the sign in flow  (`https://connect.affixapi.com`
    [with the appropriate query
    parameters](https://github.com/affixapi/starter-kit/tree/master/connect))
  - capture `authorization_code` we will send to your redirect URI after
    the sign in flow is complete and exchange that `authorization_code` for
    a Bearer token

# Sandbox keys (official mode, employees endpoint)
### dev
```
eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiJhNmNhODc5YS01OGEzLTQ2MGEtYTZlZC04N2E0NmRlMmMyNzMiLCJpYXQiOjE2OTc5ODUxMzEsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6Im9mZmljaWFsIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.Mw-eYth5VL7jpSVfnh88Tl8Cn-6-bKvjnE4GPtmuUdIS7VAvB5ijQksPjOM3FHkF382oh4bym_FAyQN_UE4mmg
```

```
curl --fail \\
  -X GET \\
  -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiJhNmNhODc5YS01OGEzLTQ2MGEtYTZlZC04N2E0NmRlMmMyNzMiLCJpYXQiOjE2OTc5ODUxMzEsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6Im9mZmljaWFsIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.Mw-eYth5VL7jpSVfnh88Tl8Cn-6-bKvjnE4GPtmuUdIS7VAvB5ijQksPjOM3FHkF382oh4bym_FAyQN_UE4mmg' \\
  'https://dev.api.affixapi.com/2023-03-01/official/employees'
```

### prod
```
eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiIwYzU2ZjcwMS0wYmFhLTQxOTQtYmU5Ni01ZThiOTExMzZmZDUiLCJpYXQiOjE2OTc5ODUwMTksImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJvZmZpY2lhbCIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.REb0qtwnn--ql2gWFb32FWilezTtJq8USN3Uj4NXoY8aJgwkjisca5ReRh5xyfprKSz_yOEcD1JwTrOlgkvf-Q
```

```
curl --fail \\
  -X GET \\
  -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiIwYzU2ZjcwMS0wYmFhLTQxOTQtYmU5Ni01ZThiOTExMzZmZDUiLCJpYXQiOjE2OTc5ODUwMTksImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJvZmZpY2lhbCIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.REb0qtwnn--ql2gWFb32FWilezTtJq8USN3Uj4NXoY8aJgwkjisca5ReRh5xyfprKSz_yOEcD1JwTrOlgkvf-Q' \\
  'https://api.affixapi.com/2023-03-01/official/employees'
```

# Webhooks
An exciting feature for HR/Payroll modes are webhooks.

If enabled, your `webhook_uri` is set on your `client_id` for the
respective environment: `dev | prod`

Webhooks are configured to make live requests to the underlying integration
1x/hr, and if a difference is detected since the last request, we will send a
request to your `webhook_uri` with this shape:

```
{

  added: <api.Employees20230301Response>[
    <api.Employee20230301>{
      ...,
      date_of_birth: '2010-08-06',
      display_full_name: 'Daija Rogahn',
      employee_number: '57993',
      employment_status: 'pending',
      employment_type: 'other',
      employments: [
        {
          currency: 'eur',
          effective_date: '2022-02-25',
          employment_type: 'other',
          job_title: 'Dynamic Implementation Manager',
          pay_frequency: 'semimonthly',
          pay_period: 'YEAR',
          pay_rate: 96000,
        },
      ],
      first_name: 'Daija',
      ...
    }
  ],
  removed: [],
  updated: [
    <api.Employee20230301>{
      ...,
      date_of_birth: '2009-11-09',
      display_full_name: 'Lourdes Stiedemann',
      employee_number: '63189',
      employment_status: 'leave',
      employment_type: 'full_time',
      employments: [
        {
          currency: 'gbp',
          effective_date: '2023-01-16',
          employment_type: 'full_time',
          job_title: 'Forward Brand Planner',
          pay_frequency: 'semimonthly',
          pay_period: 'YEAR',
          pay_rate: 86000,
        },
      ],
      first_name: 'Lourdes',
    }
  ]
}
```

the following headers will be sent with webhook requests:

```
x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7

x-affix-api-origin:  webhooks-employees-webhook

user-agent:  affixapi.com
```

Before trusting the payload, you should sign the payload and verify the
signature matches the signature sent by the `affixapi.com` service.

This secures that the data sent to your `webhook_uri` is from the
`affixapi.com` server.

The signature is created by combining the signing secret (your
`client_secret`) with the body of the request sent using a standard
HMAC-SHA256 keyed hash.

The signature can be created via:
  - create an `HMAC` with your `client_secret`
  - update the `HMAC` with the payload
  - get the hex digest -> this is the signature

Sample `typescript` code that follows this recipe:

```
import { createHmac } from 'crypto';

export const computeSignature = ({
  str,
  signingSecret,
}: {
  signingSecret: string;
  str: string;
}): string => {
  const hmac = createHmac('sha256', signingSecret);
  hmac.update(str);
  const signature = hmac.digest('hex');

  return signature;
};
```

## Rate limits
Open endpoints (not gated by an API key) (applied at endpoint level):
  - 15 requests every 1 minute (by IP address)
  - 25 requests every 5 minutes (by IP address)

Gated endpoints (require an API key) (applied at endpoint level):
  - 40 requests every 1 minute (by IP address)
  - 40 requests every 5 minutes (by `client_id`)

Things to keep in mind:
  - Open endpoints (not gated by an API key) will likely be called by your
    users, not you, so rate limits generally would not apply to you.
  - As a developer, rate limits are applied at the endpoint granularity.
    - For example, say the rate limits below are 10 requests per minute by ip.
      from that same ip, within 1 minute, you get:
      - 10 requests per minute on `/orders`,
      - another 10 requests per minute on `/items`,
      - and another 10 requests per minute on `/identity`,
      - for a total of 30 requests per minute.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2023-03-01
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.affixapi.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Class20230301Api* | [**DeveloperEmployees20230301**](docs/Class20230301Api.md#developeremployees20230301) | **Get** /2023-03-01/developer/employees | Employees
*Class20230301Api* | [**DeveloperIdentity20230301**](docs/Class20230301Api.md#developeridentity20230301) | **Get** /2023-03-01/developer/identity | Identity
*Class20230301Api* | [**DeveloperPayruns20230301**](docs/Class20230301Api.md#developerpayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
*Class20230301Api* | [**DeveloperPayslips20230301**](docs/Class20230301Api.md#developerpayslips20230301) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Get payslips of a payrun (the payslips/pay stubs/check stubs + detail)
*Class20230301Api* | [**OfficialEmployees20230301**](docs/Class20230301Api.md#officialemployees20230301) | **Get** /2023-03-01/official/employees | Employees
*Class20230301Api* | [**Officialdentity20230301**](docs/Class20230301Api.md#officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity
*Class20230320Api* | [**GetOrder20230320**](docs/Class20230320Api.md#getorder20230320) | **Get** /2023-03-20/retail/orders/{order_id} | Get an Order
*Class20230320Api* | [**Identity20230320**](docs/Class20230320Api.md#identity20230320) | **Get** /2023-03-20/retail/identity | Identity
*Class20230320Api* | [**Orders20230320**](docs/Class20230320Api.md#orders20230320) | **Get** /2023-03-20/retail/orders | Orders
*CoreApi* | [**Providers**](docs/CoreApi.md#providers) | **Get** /providers | List of providers
*DeveloperApi* | [**DeveloperEmployees20230301**](docs/DeveloperApi.md#developeremployees20230301) | **Get** /2023-03-01/developer/employees | Employees
*DeveloperApi* | [**DeveloperIdentity20230301**](docs/DeveloperApi.md#developeridentity20230301) | **Get** /2023-03-01/developer/identity | Identity
*DeveloperApi* | [**DeveloperPayruns20230301**](docs/DeveloperApi.md#developerpayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
*DeveloperApi* | [**DeveloperPayslips20230301**](docs/DeveloperApi.md#developerpayslips20230301) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Get payslips of a payrun (the payslips/pay stubs/check stubs + detail)
*ManagementApi* | [**Client**](docs/ManagementApi.md#client) | **Get** /2023-03-01/management/client | View client
*ManagementApi* | [**Disconnect**](docs/ManagementApi.md#disconnect) | **Post** /2023-03-01/management/disconnect | Disconnect token
*ManagementApi* | [**Introspect**](docs/ManagementApi.md#introspect) | **Get** /2023-03-01/management/introspect | Inspect token
*ManagementApi* | [**Token**](docs/ManagementApi.md#token) | **Post** /2023-03-01/management/token | Exchange &#x60;authorization_code&#x60; for &#x60;access_token&#x60;
*ManagementApi* | [**Tokens**](docs/ManagementApi.md#tokens) | **Get** /2023-03-01/management/tokens | View tokens
*ManagementApi* | [**UpdateClient**](docs/ManagementApi.md#updateclient) | **Post** /2023-03-01/management/client | Update client
*OfficialApi* | [**OfficialEmployees20230301**](docs/OfficialApi.md#officialemployees20230301) | **Get** /2023-03-01/official/employees | Employees
*OfficialApi* | [**Officialdentity20230301**](docs/OfficialApi.md#officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity
*RawApi* | [**AmazonGetOrder**](docs/RawApi.md#amazongetorder) | **Get** /2023-03-20/raw/amazon/orders/{order_id} | Get an Order
*RawApi* | [**AmazonIdentity**](docs/RawApi.md#amazonidentity) | **Get** /2023-03-20/raw/amazon/identity | Identity
*RawApi* | [**AmazonOrders**](docs/RawApi.md#amazonorders) | **Get** /2023-03-20/raw/amazon/orders | Orders
*RetailApi* | [**GetOrder20230320**](docs/RetailApi.md#getorder20230320) | **Get** /2023-03-20/retail/orders/{order_id} | Get an Order
*RetailApi* | [**Identity20230320**](docs/RetailApi.md#identity20230320) | **Get** /2023-03-20/retail/identity | Identity
*RetailApi* | [**Orders20230320**](docs/RetailApi.md#orders20230320) | **Get** /2023-03-20/retail/orders | Orders


## Documentation For Models

 - [ClientRequest](docs/ClientRequest.md)
 - [ClientResponse](docs/ClientResponse.md)
 - [CurrencyResponse](docs/CurrencyResponse.md)
 - [DisconnectResponse](docs/DisconnectResponse.md)
 - [Employee20230301Response](docs/Employee20230301Response.md)
 - [Employee20230301ResponseBankAccount](docs/Employee20230301ResponseBankAccount.md)
 - [Employee20230301ResponseCompany](docs/Employee20230301ResponseCompany.md)
 - [Employee20230301ResponseCompanyLocation](docs/Employee20230301ResponseCompanyLocation.md)
 - [Employee20230301ResponseGroups](docs/Employee20230301ResponseGroups.md)
 - [Employee20230301ResponseHomeLocation](docs/Employee20230301ResponseHomeLocation.md)
 - [Employee20230301ResponseManager](docs/Employee20230301ResponseManager.md)
 - [Employee20230301ResponseWorkLocation](docs/Employee20230301ResponseWorkLocation.md)
 - [EmploymentResponse](docs/EmploymentResponse.md)
 - [IdAndMessageResponse](docs/IdAndMessageResponse.md)
 - [IdentityResponse](docs/IdentityResponse.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [IntrospectResponse](docs/IntrospectResponse.md)
 - [LocationResponse](docs/LocationResponse.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [ModeRequest](docs/ModeRequest.md)
 - [ModeResponse](docs/ModeResponse.md)
 - [Order20230320Response](docs/Order20230320Response.md)
 - [PayrunResponse](docs/PayrunResponse.md)
 - [PayslipResponse](docs/PayslipResponse.md)
 - [PayslipResponseDeductions](docs/PayslipResponseDeductions.md)
 - [PayslipResponseEarnings](docs/PayslipResponseEarnings.md)
 - [PayslipResponseTaxes](docs/PayslipResponseTaxes.md)
 - [ProviderRequest](docs/ProviderRequest.md)
 - [ProviderResponse](docs/ProviderResponse.md)
 - [ScopesRequest](docs/ScopesRequest.md)
 - [ScopesResponse](docs/ScopesResponse.md)
 - [TokenRequest](docs/TokenRequest.md)
 - [TokenResponse](docs/TokenResponse.md)


## Documentation For Authorization



### access-token

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


### basic

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

john@affixapi.com

