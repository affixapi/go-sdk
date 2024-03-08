/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (developer mode) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ' \\   'https://dev.api.affixapi.com/2023-03-01/developer/employees' ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw' \\   'https://api.affixapi.com/2023-03-01/developer/employees' ```  # Compression We support `brotli`, `gzip`, and `deflate` compression algorithms.  To enable, pass the `Accept-Encoding` header with one or all of the values: `br`, `gzip`, `deflate`, or `identity` (no compression)  In the response, you will receive the `Content-Encoding` response header indicating the compression algorithm used in the data payload to enable you to decompress the result. If the `Accept-Encoding: identity` header was passed, no `Content-Encoding` response header is sent back, as no compression algorithm was used.  # Webhooks An exciting feature for HR/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.v20230301.Employees>[     <api.v20230301.Employee>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.v20230301.Employee>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: developers@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EmploymentResponse struct for EmploymentResponse
type EmploymentResponse struct {
	JobTitle NullableString `json:"job_title"`
	PayRate NullableFloat32 `json:"pay_rate"`
	PayPeriod NullableString `json:"pay_period"`
	PayFrequency NullableString `json:"pay_frequency"`
	EmploymentType NullableString `json:"employment_type"`
	Currency CurrencyResponse `json:"currency"`
	EffectiveDate NullableString `json:"effective_date"`
}

// NewEmploymentResponse instantiates a new EmploymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmploymentResponse(jobTitle NullableString, payRate NullableFloat32, payPeriod NullableString, payFrequency NullableString, employmentType NullableString, currency CurrencyResponse, effectiveDate NullableString) *EmploymentResponse {
	this := EmploymentResponse{}
	this.JobTitle = jobTitle
	this.PayRate = payRate
	this.PayPeriod = payPeriod
	this.PayFrequency = payFrequency
	this.EmploymentType = employmentType
	this.Currency = currency
	this.EffectiveDate = effectiveDate
	return &this
}

// NewEmploymentResponseWithDefaults instantiates a new EmploymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmploymentResponseWithDefaults() *EmploymentResponse {
	this := EmploymentResponse{}
	return &this
}

// GetJobTitle returns the JobTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmploymentResponse) GetJobTitle() string {
	if o == nil || o.JobTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// SetJobTitle sets field value
func (o *EmploymentResponse) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}

// GetPayRate returns the PayRate field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *EmploymentResponse) GetPayRate() float32 {
	if o == nil || o.PayRate.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PayRate.Get()
}

// GetPayRateOk returns a tuple with the PayRate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetPayRateOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayRate.Get(), o.PayRate.IsSet()
}

// SetPayRate sets field value
func (o *EmploymentResponse) SetPayRate(v float32) {
	o.PayRate.Set(&v)
}

// GetPayPeriod returns the PayPeriod field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmploymentResponse) GetPayPeriod() string {
	if o == nil || o.PayPeriod.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayPeriod.Get()
}

// GetPayPeriodOk returns a tuple with the PayPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetPayPeriodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayPeriod.Get(), o.PayPeriod.IsSet()
}

// SetPayPeriod sets field value
func (o *EmploymentResponse) SetPayPeriod(v string) {
	o.PayPeriod.Set(&v)
}

// GetPayFrequency returns the PayFrequency field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmploymentResponse) GetPayFrequency() string {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret string
		return ret
	}

	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetPayFrequencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// SetPayFrequency sets field value
func (o *EmploymentResponse) SetPayFrequency(v string) {
	o.PayFrequency.Set(&v)
}

// GetEmploymentType returns the EmploymentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmploymentResponse) GetEmploymentType() string {
	if o == nil || o.EmploymentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmploymentType.Get()
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetEmploymentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentType.Get(), o.EmploymentType.IsSet()
}

// SetEmploymentType sets field value
func (o *EmploymentResponse) SetEmploymentType(v string) {
	o.EmploymentType.Set(&v)
}

// GetCurrency returns the Currency field value
func (o *EmploymentResponse) GetCurrency() CurrencyResponse {
	if o == nil {
		var ret CurrencyResponse
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *EmploymentResponse) GetCurrencyOk() (*CurrencyResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *EmploymentResponse) SetCurrency(v CurrencyResponse) {
	o.Currency = v
}

// GetEffectiveDate returns the EffectiveDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EmploymentResponse) GetEffectiveDate() string {
	if o == nil || o.EffectiveDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EffectiveDate.Get()
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmploymentResponse) GetEffectiveDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EffectiveDate.Get(), o.EffectiveDate.IsSet()
}

// SetEffectiveDate sets field value
func (o *EmploymentResponse) SetEffectiveDate(v string) {
	o.EffectiveDate.Set(&v)
}

func (o EmploymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["job_title"] = o.JobTitle.Get()
	}
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
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["effective_date"] = o.EffectiveDate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEmploymentResponse struct {
	value *EmploymentResponse
	isSet bool
}

func (v NullableEmploymentResponse) Get() *EmploymentResponse {
	return v.value
}

func (v *NullableEmploymentResponse) Set(val *EmploymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmploymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmploymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmploymentResponse(val *EmploymentResponse) *NullableEmploymentResponse {
	return &NullableEmploymentResponse{value: val, isSet: true}
}

func (v NullableEmploymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmploymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


