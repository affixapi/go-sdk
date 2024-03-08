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

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	// The issued access_token
	AccessToken string `json:"access_token"`
	Mode ModeResponse `json:"mode"`
	Provider ProviderResponse `json:"provider"`
	// One or more scope values indicating which parts of the user's account you wish to access.  Note, slight deviation from the OAuth 2.1 spec in that the param is scopes (plural) is used vs scope (singular) 
	Scopes []ScopesResponse `json:"scopes"`
	// The token type to pass in the `Authorization` header
	TokenType string `json:"token_type"`
}

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse(accessToken string, mode ModeResponse, provider ProviderResponse, scopes []ScopesResponse, tokenType string) *TokenResponse {
	this := TokenResponse{}
	this.AccessToken = accessToken
	this.Mode = mode
	this.Provider = provider
	this.Scopes = scopes
	this.TokenType = tokenType
	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *TokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *TokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetMode returns the Mode field value
func (o *TokenResponse) GetMode() ModeResponse {
	if o == nil {
		var ret ModeResponse
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetModeOk() (*ModeResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TokenResponse) SetMode(v ModeResponse) {
	o.Mode = v
}

// GetProvider returns the Provider field value
func (o *TokenResponse) GetProvider() ProviderResponse {
	if o == nil {
		var ret ProviderResponse
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetProviderOk() (*ProviderResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *TokenResponse) SetProvider(v ProviderResponse) {
	o.Provider = v
}

// GetScopes returns the Scopes field value
func (o *TokenResponse) GetScopes() []ScopesResponse {
	if o == nil {
		var ret []ScopesResponse
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetScopesOk() (*[]ScopesResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scopes, true
}

// SetScopes sets field value
func (o *TokenResponse) SetScopes(v []ScopesResponse) {
	o.Scopes = v
}

// GetTokenType returns the TokenType field value
func (o *TokenResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *TokenResponse) SetTokenType(v string) {
	o.TokenType = v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["scopes"] = o.Scopes
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


