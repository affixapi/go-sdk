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

// InlineResponse409 struct for InlineResponse409
type InlineResponse409 struct {
	// The authorization code insert that caused a conflict 
	Code *string `json:"code,omitempty"`
	// The access token insert that caused a conflict 
	Token *string `json:"token,omitempty"`
	// The description of the type of error, ie conflict of two requests of the the same authorization code 
	Error string `json:"error"`
}

// NewInlineResponse409 instantiates a new InlineResponse409 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse409(error_ string) *InlineResponse409 {
	this := InlineResponse409{}
	this.Error = error_
	return &this
}

// NewInlineResponse409WithDefaults instantiates a new InlineResponse409 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse409WithDefaults() *InlineResponse409 {
	this := InlineResponse409{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse409) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse409) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse409) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse409) SetCode(v string) {
	o.Code = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *InlineResponse409) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse409) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *InlineResponse409) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *InlineResponse409) SetToken(v string) {
	o.Token = &v
}

// GetError returns the Error field value
func (o *InlineResponse409) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *InlineResponse409) GetErrorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *InlineResponse409) SetError(v string) {
	o.Error = v
}

func (o InlineResponse409) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse409 struct {
	value *InlineResponse409
	isSet bool
}

func (v NullableInlineResponse409) Get() *InlineResponse409 {
	return v.value
}

func (v *NullableInlineResponse409) Set(val *InlineResponse409) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse409) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse409(val *InlineResponse409) *NullableInlineResponse409 {
	return &NullableInlineResponse409{value: val, isSet: true}
}

func (v NullableInlineResponse409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


