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

// CompensationHistoryResponse struct for CompensationHistoryResponse
type CompensationHistoryResponse struct {
	PayRate NullableFloat32 `json:"pay_rate"`
	PayPeriod NullableString `json:"pay_period"`
	PayFrequency NullableString `json:"pay_frequency"`
	EmploymentType NullableString `json:"employment_type"`
	Currency NullableCurrencyResponse `json:"currency"`
	EffectiveDate NullableString `json:"effective_date"`
	Notes NullableString `json:"notes"`
}

// NewCompensationHistoryResponse instantiates a new CompensationHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompensationHistoryResponse(payRate NullableFloat32, payPeriod NullableString, payFrequency NullableString, employmentType NullableString, currency NullableCurrencyResponse, effectiveDate NullableString, notes NullableString) *CompensationHistoryResponse {
	this := CompensationHistoryResponse{}
	this.PayRate = payRate
	this.PayPeriod = payPeriod
	this.PayFrequency = payFrequency
	this.EmploymentType = employmentType
	this.Currency = currency
	this.EffectiveDate = effectiveDate
	this.Notes = notes
	return &this
}

// NewCompensationHistoryResponseWithDefaults instantiates a new CompensationHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompensationHistoryResponseWithDefaults() *CompensationHistoryResponse {
	this := CompensationHistoryResponse{}
	return &this
}

// GetPayRate returns the PayRate field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *CompensationHistoryResponse) GetPayRate() float32 {
	if o == nil || o.PayRate.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PayRate.Get()
}

// GetPayRateOk returns a tuple with the PayRate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetPayRateOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayRate.Get(), o.PayRate.IsSet()
}

// SetPayRate sets field value
func (o *CompensationHistoryResponse) SetPayRate(v float32) {
	o.PayRate.Set(&v)
}

// GetPayPeriod returns the PayPeriod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompensationHistoryResponse) GetPayPeriod() string {
	if o == nil || o.PayPeriod.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayPeriod.Get()
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetPayPeriodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayPeriod.Get(), o.PayPeriod.IsSet()
}

// SetPayPeriod sets field value
func (o *CompensationHistoryResponse) SetPayPeriod(v string) {
	o.PayPeriod.Set(&v)
}

// GetPayFrequency returns the PayFrequency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompensationHistoryResponse) GetPayFrequency() string {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetPayFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// SetPayFrequency sets field value
func (o *CompensationHistoryResponse) SetPayFrequency(v string) {
	o.PayFrequency.Set(&v)
}

// GetEmploymentType returns the EmploymentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompensationHistoryResponse) GetEmploymentType() string {
	if o == nil || o.EmploymentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmploymentType.Get()
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetEmploymentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentType.Get(), o.EmploymentType.IsSet()
}

// SetEmploymentType sets field value
func (o *CompensationHistoryResponse) SetEmploymentType(v string) {
	o.EmploymentType.Set(&v)
}

// GetCurrency returns the Currency field value
// If the value is explicit nil, the zero value for CurrencyResponse will be returned
func (o *CompensationHistoryResponse) GetCurrency() CurrencyResponse {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyResponse
		return ret
	}

	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetCurrencyOk() (*CurrencyResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// SetCurrency sets field value
func (o *CompensationHistoryResponse) SetCurrency(v CurrencyResponse) {
	o.Currency.Set(&v)
}

// GetEffectiveDate returns the EffectiveDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompensationHistoryResponse) GetEffectiveDate() string {
	if o == nil || o.EffectiveDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EffectiveDate.Get()
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetEffectiveDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EffectiveDate.Get(), o.EffectiveDate.IsSet()
}

// SetEffectiveDate sets field value
func (o *CompensationHistoryResponse) SetEffectiveDate(v string) {
	o.EffectiveDate.Set(&v)
}

// GetNotes returns the Notes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CompensationHistoryResponse) GetNotes() string {
	if o == nil || o.Notes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompensationHistoryResponse) GetNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// SetNotes sets field value
func (o *CompensationHistoryResponse) SetNotes(v string) {
	o.Notes.Set(&v)
}

func (o CompensationHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pay_rate"] = o.PayRate.Get()
	}
	if true {
		toSerialize["pay_period"] = o.PayPeriod.Get()
	}
	if true {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if true {
		toSerialize["employment_type"] = o.EmploymentType.Get()
	}
	if true {
		toSerialize["currency"] = o.Currency.Get()
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate.Get()
	}
	if true {
		toSerialize["notes"] = o.Notes.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCompensationHistoryResponse struct {
	value *CompensationHistoryResponse
	isSet bool
}

func (v NullableCompensationHistoryResponse) Get() *CompensationHistoryResponse {
	return v.value
}

func (v *NullableCompensationHistoryResponse) Set(val *CompensationHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompensationHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompensationHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompensationHistoryResponse(val *CompensationHistoryResponse) *NullableCompensationHistoryResponse {
	return &NullableCompensationHistoryResponse{value: val, isSet: true}
}

func (v NullableCompensationHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompensationHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


