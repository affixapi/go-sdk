# Go API client for openapi

The affixapi.com API documentation.

# Introduction
Affix API is an OAuth 2.1 application that allows developers to access
customer data, without developers needing to manage or maintain
integrations; or collect login credentials or API keys from users for these
third party systems.

# OAuth 2.1
Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).

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

# Sandbox keys (developer mode)
### dev
```
eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ
```

#### `employees` endpoint sample:
```
curl --fail \\
  -X GET \\
  -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ' \\
  'https://dev.api.affixapi.com/2023-03-01/developer/employees'
```

### prod
```
eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw
```

#### `employees` endpoint sample:
```
curl --fail \\
  -X GET \\
  -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw' \\
  'https://api.affixapi.com/2023-03-01/developer/employees'
```

# Compression
We support `brotli`, `gzip`, and `deflate` compression algorithms.

To enable, pass the `Accept-Encoding` header with one or all of the values:
`br`, `gzip`, `deflate`, or `identity` (no compression)

In the response, you will receive the `Content-Encoding` response header
indicating the compression algorithm used in the data payload to enable you
to decompress the result. If the `Accept-Encoding: identity` header was
passed, no `Content-Encoding` response header is sent back, as no
compression algorithm was used.

# Webhooks
An exciting feature for HR/Payroll modes are webhooks.

If enabled, your `webhook_uri` is set on your `client_id` for the
respective environment: `dev | prod`

Webhooks are configured to make live requests to the underlying integration
1x/hr, and if a difference is detected since the last request, we will send a
request to your `webhook_uri` with this shape:

```
{

  added: <api.v20230301.Employees>[
    <api.v20230301.Employee>{
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
    <api.v20230301.Employee>{
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
*Class20230301Api* | [**DeveloperCompanies20230301**](docs/Class20230301Api.md#developercompanies20230301) | **Get** /2023-03-01/developer/company | Company
*Class20230301Api* | [**DeveloperCreateEmployee20230301**](docs/Class20230301Api.md#developercreateemployee20230301) | **Post** /2023-03-01/developer/employee | Create employee
*Class20230301Api* | [**DeveloperEmployees20230301**](docs/Class20230301Api.md#developeremployees20230301) | **Get** /2023-03-01/developer/employees | Employees
*Class20230301Api* | [**DeveloperGroups20230301**](docs/Class20230301Api.md#developergroups20230301) | **Get** /2023-03-01/developer/groups | Groups
*Class20230301Api* | [**DeveloperIdentity20230301**](docs/Class20230301Api.md#developeridentity20230301) | **Get** /2023-03-01/developer/identity | Identity
*Class20230301Api* | [**DeveloperPayruns20230301**](docs/Class20230301Api.md#developerpayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
*Class20230301Api* | [**DeveloperPayslips20230301**](docs/Class20230301Api.md#developerpayslips20230301) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Payslips
*Class20230301Api* | [**DeveloperTimeOffBalances20230301**](docs/Class20230301Api.md#developertimeoffbalances20230301) | **Get** /2023-03-01/developer/time-off-balances | Time off balances
*Class20230301Api* | [**DeveloperTimeOffEntries20230301**](docs/Class20230301Api.md#developertimeoffentries20230301) | **Get** /2023-03-01/developer/time-off-entries | Time off entries
*Class20230301Api* | [**DeveloperTimesheets20230301**](docs/Class20230301Api.md#developertimesheets20230301) | **Get** /2023-03-01/developer/timesheets | Timesheets
*Class20230301Api* | [**DeveloperWorkLocations20230301**](docs/Class20230301Api.md#developerworklocations20230301) | **Get** /2023-03-01/developer/work-locations | Work locations
*Class20230301Api* | [**OfficialCompanies20230301**](docs/Class20230301Api.md#officialcompanies20230301) | **Get** /2023-03-01/official/company | Company
*Class20230301Api* | [**OfficialCreateEmployee20230301**](docs/Class20230301Api.md#officialcreateemployee20230301) | **Post** /2023-03-01/official/employee | Create Employee
*Class20230301Api* | [**OfficialEmployees20230301**](docs/Class20230301Api.md#officialemployees20230301) | **Get** /2023-03-01/official/employees | Employees
*Class20230301Api* | [**OfficialGroups20230301**](docs/Class20230301Api.md#officialgroups20230301) | **Get** /2023-03-01/official/groups | Groups
*Class20230301Api* | [**OfficialTimeOffBalances20230301**](docs/Class20230301Api.md#officialtimeoffbalances20230301) | **Get** /2023-03-01/official/time-off-balances | Time off balances
*Class20230301Api* | [**OfficialTimeOffEntries20230301**](docs/Class20230301Api.md#officialtimeoffentries20230301) | **Get** /2023-03-01/official/time-off-entries | Time off entries
*Class20230301Api* | [**OfficialTimesheets20230301**](docs/Class20230301Api.md#officialtimesheets20230301) | **Get** /2023-03-01/official/timesheets | Timesheets
*Class20230301Api* | [**OfficialWorkLocations20230301**](docs/Class20230301Api.md#officialworklocations20230301) | **Get** /2023-03-01/official/work-locations | Work locations
*Class20230301Api* | [**Officialdentity20230301**](docs/Class20230301Api.md#officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity
*CoreApi* | [**Providers**](docs/CoreApi.md#providers) | **Get** /providers | Providers
*DeveloperApi* | [**DeveloperCompanies20230301**](docs/DeveloperApi.md#developercompanies20230301) | **Get** /2023-03-01/developer/company | Company
*DeveloperApi* | [**DeveloperCreateEmployee20230301**](docs/DeveloperApi.md#developercreateemployee20230301) | **Post** /2023-03-01/developer/employee | Create employee
*DeveloperApi* | [**DeveloperEmployees20230301**](docs/DeveloperApi.md#developeremployees20230301) | **Get** /2023-03-01/developer/employees | Employees
*DeveloperApi* | [**DeveloperGroups20230301**](docs/DeveloperApi.md#developergroups20230301) | **Get** /2023-03-01/developer/groups | Groups
*DeveloperApi* | [**DeveloperIdentity20230301**](docs/DeveloperApi.md#developeridentity20230301) | **Get** /2023-03-01/developer/identity | Identity
*DeveloperApi* | [**DeveloperPayruns20230301**](docs/DeveloperApi.md#developerpayruns20230301) | **Get** /2023-03-01/developer/payruns | Payruns
*DeveloperApi* | [**DeveloperPayslips20230301**](docs/DeveloperApi.md#developerpayslips20230301) | **Get** /2023-03-01/developer/payruns/{payrun_id} | Payslips
*DeveloperApi* | [**DeveloperTimeOffBalances20230301**](docs/DeveloperApi.md#developertimeoffbalances20230301) | **Get** /2023-03-01/developer/time-off-balances | Time off balances
*DeveloperApi* | [**DeveloperTimeOffEntries20230301**](docs/DeveloperApi.md#developertimeoffentries20230301) | **Get** /2023-03-01/developer/time-off-entries | Time off entries
*DeveloperApi* | [**DeveloperTimesheets20230301**](docs/DeveloperApi.md#developertimesheets20230301) | **Get** /2023-03-01/developer/timesheets | Timesheets
*DeveloperApi* | [**DeveloperWorkLocations20230301**](docs/DeveloperApi.md#developerworklocations20230301) | **Get** /2023-03-01/developer/work-locations | Work locations
*ManagementApi* | [**Client**](docs/ManagementApi.md#client) | **Get** /2023-03-01/management/client | Client
*ManagementApi* | [**Disconnect**](docs/ManagementApi.md#disconnect) | **Post** /2023-03-01/management/disconnect | Disconnect token
*ManagementApi* | [**Introspect**](docs/ManagementApi.md#introspect) | **Get** /2023-03-01/management/introspect | Inspect token
*ManagementApi* | [**Token**](docs/ManagementApi.md#token) | **Post** /2023-03-01/management/token | Create token
*ManagementApi* | [**Tokens**](docs/ManagementApi.md#tokens) | **Get** /2023-03-01/management/tokens | Tokens
*ManagementApi* | [**UpdateClient**](docs/ManagementApi.md#updateclient) | **Post** /2023-03-01/management/client | Update client
*OfficialApi* | [**OfficialCompanies20230301**](docs/OfficialApi.md#officialcompanies20230301) | **Get** /2023-03-01/official/company | Company
*OfficialApi* | [**OfficialCreateEmployee20230301**](docs/OfficialApi.md#officialcreateemployee20230301) | **Post** /2023-03-01/official/employee | Create Employee
*OfficialApi* | [**OfficialEmployees20230301**](docs/OfficialApi.md#officialemployees20230301) | **Get** /2023-03-01/official/employees | Employees
*OfficialApi* | [**OfficialGroups20230301**](docs/OfficialApi.md#officialgroups20230301) | **Get** /2023-03-01/official/groups | Groups
*OfficialApi* | [**OfficialTimeOffBalances20230301**](docs/OfficialApi.md#officialtimeoffbalances20230301) | **Get** /2023-03-01/official/time-off-balances | Time off balances
*OfficialApi* | [**OfficialTimeOffEntries20230301**](docs/OfficialApi.md#officialtimeoffentries20230301) | **Get** /2023-03-01/official/time-off-entries | Time off entries
*OfficialApi* | [**OfficialTimesheets20230301**](docs/OfficialApi.md#officialtimesheets20230301) | **Get** /2023-03-01/official/timesheets | Timesheets
*OfficialApi* | [**OfficialWorkLocations20230301**](docs/OfficialApi.md#officialworklocations20230301) | **Get** /2023-03-01/official/work-locations | Work locations
*OfficialApi* | [**Officialdentity20230301**](docs/OfficialApi.md#officialdentity20230301) | **Get** /2023-03-01/official/identity | Identity


## Documentation For Models

 - [AddressNoNonNullRequest](docs/AddressNoNonNullRequest.md)
 - [AddressResponse](docs/AddressResponse.md)
 - [ClientRequest](docs/ClientRequest.md)
 - [ClientResponse](docs/ClientResponse.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CreateEmployeeRequest](docs/CreateEmployeeRequest.md)
 - [CreateEmployeeRequestBankAccount](docs/CreateEmployeeRequestBankAccount.md)
 - [CreateEmployeeRequestManager](docs/CreateEmployeeRequestManager.md)
 - [CurrencyRequest](docs/CurrencyRequest.md)
 - [CurrencyResponse](docs/CurrencyResponse.md)
 - [DisconnectResponse](docs/DisconnectResponse.md)
 - [EmployeeResponse](docs/EmployeeResponse.md)
 - [EmploymentNoNullEnumRequest](docs/EmploymentNoNullEnumRequest.md)
 - [EmploymentResponse](docs/EmploymentResponse.md)
 - [GroupNoNullEnumRequest](docs/GroupNoNullEnumRequest.md)
 - [GroupResponse](docs/GroupResponse.md)
 - [IdAndMessageResponse](docs/IdAndMessageResponse.md)
 - [IdentityResponse](docs/IdentityResponse.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InlineResponse401](docs/InlineResponse401.md)
 - [InlineResponse409](docs/InlineResponse409.md)
 - [IntrospectResponse](docs/IntrospectResponse.md)
 - [LocationNoNonNullRequest](docs/LocationNoNonNullRequest.md)
 - [LocationResponse](docs/LocationResponse.md)
 - [MessageResponse](docs/MessageResponse.md)
 - [ModeRequest](docs/ModeRequest.md)
 - [ModeResponse](docs/ModeResponse.md)
 - [PayrunResponse](docs/PayrunResponse.md)
 - [PayslipResponse](docs/PayslipResponse.md)
 - [PayslipResponseContributions](docs/PayslipResponseContributions.md)
 - [PayslipResponseDeductions](docs/PayslipResponseDeductions.md)
 - [PayslipResponseEarnings](docs/PayslipResponseEarnings.md)
 - [PayslipResponseTaxes](docs/PayslipResponseTaxes.md)
 - [ProviderRequest](docs/ProviderRequest.md)
 - [ProviderResponse](docs/ProviderResponse.md)
 - [ScopesRequest](docs/ScopesRequest.md)
 - [ScopesResponse](docs/ScopesResponse.md)
 - [TimeOffBalanceResponse](docs/TimeOffBalanceResponse.md)
 - [TimeOffEntryResponse](docs/TimeOffEntryResponse.md)
 - [TimesheetResponse](docs/TimesheetResponse.md)
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

developers@affixapi.com

