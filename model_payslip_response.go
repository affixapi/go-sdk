/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (xhr mode) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEveGhyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEveGhyL2dyb3VwcyIsIi8yMDIzLTAzLTAxL3hoci9pZGVudGl0eSIsIi8yMDIzLTAzLTAxL3hoci9wYXlydW5zIiwiLzIwMjMtMDMtMDEveGhyL3BheXJ1bnMvOnBheXJ1bl9pZCIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1iYWxhbmNlcyIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1lbnRyaWVzIiwiLzIwMjMtMDMtMDEveGhyL3RpbWVzaGVldHMiLCIvMjAyMy0wMy0wMS94aHIvd29yay1sb2NhdGlvbnMiXSwidG9rZW4iOiIzODIzNTNlMi05N2ZiLTRmMWEtOTYxYy0zZDI5OTViNzYxMTUiLCJpYXQiOjE3MTE4MTA3MTQsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InhociIsImF1ZCI6IjNGREFFREY5LTFEQ0E0RjU0LTg3OTQ5RjZBLTQxMDI3NjQzIn0.zUJPaT6IxcIdr8b9iO6u-Rr5I-ohTHPYTrQGrgOFghbEbovItiwr9Wk479GnJVJc3WR8bxAwUMAE4Ul6Okdk6Q ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEveGhyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEveGhyL2dyb3VwcyIsIi8yMDIzLTAzLTAxL3hoci9pZGVudGl0eSIsIi8yMDIzLTAzLTAxL3hoci9wYXlydW5zIiwiLzIwMjMtMDMtMDEveGhyL3BheXJ1bnMvOnBheXJ1bl9pZCIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1iYWxhbmNlcyIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1lbnRyaWVzIiwiLzIwMjMtMDMtMDEveGhyL3RpbWVzaGVldHMiLCIvMjAyMy0wMy0wMS94aHIvd29yay1sb2NhdGlvbnMiXSwidG9rZW4iOiIzODIzNTNlMi05N2ZiLTRmMWEtOTYxYy0zZDI5OTViNzYxMTUiLCJpYXQiOjE3MTE4MTA3MTQsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InhociIsImF1ZCI6IjNGREFFREY5LTFEQ0E0RjU0LTg3OTQ5RjZBLTQxMDI3NjQzIn0.zUJPaT6IxcIdr8b9iO6u-Rr5I-ohTHPYTrQGrgOFghbEbovItiwr9Wk479GnJVJc3WR8bxAwUMAE4Ul6Okdk6Q' \\   'https://dev.api.affixapi.com/2023-03-01/xhr/employees' ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEveGhyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEveGhyL2dyb3VwcyIsIi8yMDIzLTAzLTAxL3hoci9pZGVudGl0eSIsIi8yMDIzLTAzLTAxL3hoci9wYXlydW5zIiwiLzIwMjMtMDMtMDEveGhyL3BheXJ1bnMvOnBheXJ1bl9pZCIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1iYWxhbmNlcyIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1lbnRyaWVzIiwiLzIwMjMtMDMtMDEveGhyL3RpbWVzaGVldHMiLCIvMjAyMy0wMy0wMS94aHIvd29yay1sb2NhdGlvbnMiXSwidG9rZW4iOiIzYjg4MDc2NC1kMGFmLTQ5ZDAtOGM5OS00YzIwYjE2MTJjOTMiLCJpYXQiOjE3MTE4MTA4NTgsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJ4aHIiLCJhdWQiOiIwOEJCMDgxRS1EOUFCNEQxNC04REY5OTIzMy02NjYxNUNFOSJ9.n3pJmmfegU21Tko_TyUyCHi4ITvfd75T8NFFTHmf1r8AI8yCUYTWdfNjyZZWcZD6z50I3Wsk2rAd8GDWXn4vlg ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEveGhyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS94aHIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEveGhyL2dyb3VwcyIsIi8yMDIzLTAzLTAxL3hoci9pZGVudGl0eSIsIi8yMDIzLTAzLTAxL3hoci9wYXlydW5zIiwiLzIwMjMtMDMtMDEveGhyL3BheXJ1bnMvOnBheXJ1bl9pZCIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1iYWxhbmNlcyIsIi8yMDIzLTAzLTAxL3hoci90aW1lLW9mZi1lbnRyaWVzIiwiLzIwMjMtMDMtMDEveGhyL3RpbWVzaGVldHMiLCIvMjAyMy0wMy0wMS94aHIvd29yay1sb2NhdGlvbnMiXSwidG9rZW4iOiIzYjg4MDc2NC1kMGFmLTQ5ZDAtOGM5OS00YzIwYjE2MTJjOTMiLCJpYXQiOjE3MTE4MTA4NTgsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJ4aHIiLCJhdWQiOiIwOEJCMDgxRS1EOUFCNEQxNC04REY5OTIzMy02NjYxNUNFOSJ9.n3pJmmfegU21Tko_TyUyCHi4ITvfd75T8NFFTHmf1r8AI8yCUYTWdfNjyZZWcZD6z50I3Wsk2rAd8GDWXn4vlg' \\   'https://api.affixapi.com/2023-03-01/xhr/employees' ```  # Compression We support `brotli`, `gzip`, and `deflate` compression algorithms.  To enable, pass the `Accept-Encoding` header with one or all of the values: `br`, `gzip`, `deflate`, or `identity` (no compression)  In the response, you will receive the `Content-Encoding` response header indicating the compression algorithm used in the data payload to enable you to decompress the result. If the `Accept-Encoding: identity` header was passed, no `Content-Encoding` response header is sent back, as no compression algorithm was used.  # Webhooks An exciting feature for HR/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.v20230301.Employees>[     <api.v20230301.Employee>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.v20230301.Employee>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  While verifying the Affix API signature header should be your primary method of confirming validity, you can also whitelist our outbound webhook static IP addresses.  ``` dev:   - 52.210.169.82   - 52.210.38.77   - 3.248.135.204  prod:   - 52.51.160.102   - 54.220.83.244   - 3.254.213.171 ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: developers@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PayslipResponse struct for PayslipResponse
type PayslipResponse struct {
	// The Affix-assigned id of the payslip
	Id NullableString `json:"id"`
	// the remote system-assigned id of the payrun
	RemoteId NullableString `json:"remote_id"`
	EmployeeId string `json:"employee_id"`
	EmployeeRemoteId string `json:"employee_remote_id"`
	PayrunId string `json:"payrun_id"`
	PayrunRemoteId string `json:"payrun_remote_id"`
	PayrunType NullablePayrunTypeResponse `json:"payrun_type"`
	Currency NullableCurrencyNotNullResponse `json:"currency"`
	// if USD/EUR/GBP, in cent
	GrossPay NullableFloat32 `json:"gross_pay"`
	// if USD/EUR/GBP, in cent
	NetPay NullableFloat32 `json:"net_pay"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	PaymentDate string `json:"payment_date"`
	Earnings []PayslipResponseEarnings `json:"earnings"`
	// Items paid by the employer that are not included in gross pay, such as employer-paid portion of private health insurance 
	Contributions []PayslipResponseContributions `json:"contributions"`
	Deductions []PayslipResponseDeductions `json:"deductions"`
	Taxes []PayslipResponseTaxes `json:"taxes"`
}

// NewPayslipResponse instantiates a new PayslipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayslipResponse(id NullableString, remoteId NullableString, employeeId string, employeeRemoteId string, payrunId string, payrunRemoteId string, payrunType NullablePayrunTypeResponse, currency NullableCurrencyNotNullResponse, grossPay NullableFloat32, netPay NullableFloat32, startDate string, endDate string, paymentDate string, earnings []PayslipResponseEarnings, contributions []PayslipResponseContributions, deductions []PayslipResponseDeductions, taxes []PayslipResponseTaxes) *PayslipResponse {
	this := PayslipResponse{}
	this.Id = id
	this.RemoteId = remoteId
	this.EmployeeId = employeeId
	this.EmployeeRemoteId = employeeRemoteId
	this.PayrunId = payrunId
	this.PayrunRemoteId = payrunRemoteId
	this.PayrunType = payrunType
	this.Currency = currency
	this.GrossPay = grossPay
	this.NetPay = netPay
	this.StartDate = startDate
	this.EndDate = endDate
	this.PaymentDate = paymentDate
	this.Earnings = earnings
	this.Contributions = contributions
	this.Deductions = deductions
	this.Taxes = taxes
	return &this
}

// NewPayslipResponseWithDefaults instantiates a new PayslipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayslipResponseWithDefaults() *PayslipResponse {
	this := PayslipResponse{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayslipResponse) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *PayslipResponse) SetId(v string) {
	o.Id.Set(&v)
}

// GetRemoteId returns the RemoteId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PayslipResponse) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}

	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// SetRemoteId sets field value
func (o *PayslipResponse) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}

// GetEmployeeId returns the EmployeeId field value
func (o *PayslipResponse) GetEmployeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEmployeeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeId, true
}

// SetEmployeeId sets field value
func (o *PayslipResponse) SetEmployeeId(v string) {
	o.EmployeeId = v
}

// GetEmployeeRemoteId returns the EmployeeRemoteId field value
func (o *PayslipResponse) GetEmployeeRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeRemoteId
}

// GetEmployeeRemoteIdOk returns a tuple with the EmployeeRemoteId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEmployeeRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeRemoteId, true
}

// SetEmployeeRemoteId sets field value
func (o *PayslipResponse) SetEmployeeRemoteId(v string) {
	o.EmployeeRemoteId = v
}

// GetPayrunId returns the PayrunId field value
func (o *PayslipResponse) GetPayrunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayrunId
}

// GetPayrunIdOk returns a tuple with the PayrunId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetPayrunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayrunId, true
}

// SetPayrunId sets field value
func (o *PayslipResponse) SetPayrunId(v string) {
	o.PayrunId = v
}

// GetPayrunRemoteId returns the PayrunRemoteId field value
func (o *PayslipResponse) GetPayrunRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayrunRemoteId
}

// GetPayrunRemoteIdOk returns a tuple with the PayrunRemoteId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetPayrunRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayrunRemoteId, true
}

// SetPayrunRemoteId sets field value
func (o *PayslipResponse) SetPayrunRemoteId(v string) {
	o.PayrunRemoteId = v
}

// GetPayrunType returns the PayrunType field value
// If the value is explicit nil, the zero value for PayrunTypeResponse will be returned
func (o *PayslipResponse) GetPayrunType() PayrunTypeResponse {
	if o == nil || o.PayrunType.Get() == nil {
		var ret PayrunTypeResponse
		return ret
	}

	return *o.PayrunType.Get()
}

// GetPayrunTypeOk returns a tuple with the PayrunType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetPayrunTypeOk() (*PayrunTypeResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayrunType.Get(), o.PayrunType.IsSet()
}

// SetPayrunType sets field value
func (o *PayslipResponse) SetPayrunType(v PayrunTypeResponse) {
	o.PayrunType.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for CurrencyNotNullResponse will be returned
func (o *PayslipResponse) GetCurrency() CurrencyNotNullResponse {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyNotNullResponse
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetCurrencyOk() (*CurrencyNotNullResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *PayslipResponse) SetCurrency(v CurrencyNotNullResponse) {
	o.Currency.Set(&v)
}

// GetGrossPay returns the GrossPay field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PayslipResponse) GetGrossPay() float32 {
	if o == nil || o.GrossPay.Get() == nil {
		var ret float32
		return ret
	}

	return *o.GrossPay.Get()
}

// GetGrossPayOk returns a tuple with the GrossPay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetGrossPayOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GrossPay.Get(), o.GrossPay.IsSet()
}

// SetGrossPay sets field value
func (o *PayslipResponse) SetGrossPay(v float32) {
	o.GrossPay.Set(&v)
}

// GetNetPay returns the NetPay field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PayslipResponse) GetNetPay() float32 {
	if o == nil || o.NetPay.Get() == nil {
		var ret float32
		return ret
	}

	return *o.NetPay.Get()
}

// GetNetPayOk returns a tuple with the NetPay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetNetPayOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetPay.Get(), o.NetPay.IsSet()
}

// SetNetPay sets field value
func (o *PayslipResponse) SetNetPay(v float32) {
	o.NetPay.Set(&v)
}

// GetStartDate returns the StartDate field value
func (o *PayslipResponse) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *PayslipResponse) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *PayslipResponse) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *PayslipResponse) SetEndDate(v string) {
	o.EndDate = v
}

// GetPaymentDate returns the PaymentDate field value
func (o *PayslipResponse) GetPaymentDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentDate, true
}

// SetPaymentDate sets field value
func (o *PayslipResponse) SetPaymentDate(v string) {
	o.PaymentDate = v
}

// GetEarnings returns the Earnings field value
// If the value is explicit nil, the zero value for []PayslipResponseEarnings will be returned
func (o *PayslipResponse) GetEarnings() []PayslipResponseEarnings {
	if o == nil {
		var ret []PayslipResponseEarnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetEarningsOk() (*[]PayslipResponseEarnings, bool) {
	if o == nil || o.Earnings == nil {
		return nil, false
	}
	return &o.Earnings, true
}

// SetEarnings sets field value
func (o *PayslipResponse) SetEarnings(v []PayslipResponseEarnings) {
	o.Earnings = v
}

// GetContributions returns the Contributions field value
// If the value is explicit nil, the zero value for []PayslipResponseContributions will be returned
func (o *PayslipResponse) GetContributions() []PayslipResponseContributions {
	if o == nil {
		var ret []PayslipResponseContributions
		return ret
	}

	return o.Contributions
}

// GetContributionsOk returns a tuple with the Contributions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetContributionsOk() (*[]PayslipResponseContributions, bool) {
	if o == nil || o.Contributions == nil {
		return nil, false
	}
	return &o.Contributions, true
}

// SetContributions sets field value
func (o *PayslipResponse) SetContributions(v []PayslipResponseContributions) {
	o.Contributions = v
}

// GetDeductions returns the Deductions field value
// If the value is explicit nil, the zero value for []PayslipResponseDeductions will be returned
func (o *PayslipResponse) GetDeductions() []PayslipResponseDeductions {
	if o == nil {
		var ret []PayslipResponseDeductions
		return ret
	}

	return o.Deductions
}

// GetDeductionsOk returns a tuple with the Deductions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetDeductionsOk() (*[]PayslipResponseDeductions, bool) {
	if o == nil || o.Deductions == nil {
		return nil, false
	}
	return &o.Deductions, true
}

// SetDeductions sets field value
func (o *PayslipResponse) SetDeductions(v []PayslipResponseDeductions) {
	o.Deductions = v
}

// GetTaxes returns the Taxes field value
// If the value is explicit nil, the zero value for []PayslipResponseTaxes will be returned
func (o *PayslipResponse) GetTaxes() []PayslipResponseTaxes {
	if o == nil {
		var ret []PayslipResponseTaxes
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetTaxesOk() (*[]PayslipResponseTaxes, bool) {
	if o == nil || o.Taxes == nil {
		return nil, false
	}
	return &o.Taxes, true
}

// SetTaxes sets field value
func (o *PayslipResponse) SetTaxes(v []PayslipResponseTaxes) {
	o.Taxes = v
}

func (o PayslipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if true {
		toSerialize["employee_id"] = o.EmployeeId
	}
	if true {
		toSerialize["employee_remote_id"] = o.EmployeeRemoteId
	}
	if true {
		toSerialize["payrun_id"] = o.PayrunId
	}
	if true {
		toSerialize["payrun_remote_id"] = o.PayrunRemoteId
	}
	if true {
		toSerialize["payrun_type"] = o.PayrunType.Get()
	}
	if true {
		toSerialize["currency"] = o.Currency.Get()
	}
	if true {
		toSerialize["gross_pay"] = o.GrossPay.Get()
	}
	if true {
		toSerialize["net_pay"] = o.NetPay.Get()
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["payment_date"] = o.PaymentDate
	}
	if o.Earnings != nil {
		toSerialize["earnings"] = o.Earnings
	}
	if o.Contributions != nil {
		toSerialize["contributions"] = o.Contributions
	}
	if o.Deductions != nil {
		toSerialize["deductions"] = o.Deductions
	}
	if o.Taxes != nil {
		toSerialize["taxes"] = o.Taxes
	}
	return json.Marshal(toSerialize)
}

type NullablePayslipResponse struct {
	value *PayslipResponse
	isSet bool
}

func (v NullablePayslipResponse) Get() *PayslipResponse {
	return v.value
}

func (v *NullablePayslipResponse) Set(val *PayslipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayslipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayslipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayslipResponse(val *PayslipResponse) *NullablePayslipResponse {
	return &NullablePayslipResponse{value: val, isSet: true}
}

func (v NullablePayslipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayslipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


