/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (payroll) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiI1ZGZmMWRlMy1kZTk2LTQzZWYtYTgwMi0wYWZlNWQ5NzU4MzkiLCJpYXQiOjE2OTM4Mzk1NjIsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InBheXJvbGwiLCJhdWQiOiIzRkRBRURGOS0xRENBNEY1NC04Nzk0OUY2QS00MTAyNzY0MyJ9.D8i_TwguGZ9vok7RCUuk5FoUJGWEGnI-4VzReI1DAg0ytZaQ8XFmVxC4RxGYBxah1ZJpcrbAsuhssGEDXffIgQ ```  ``` curl --fail \\   \"https://dev.api.affixapi.com/2023-03-01/payroll/employees\" \\   -H \"Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiI1ZGZmMWRlMy1kZTk2LTQzZWYtYTgwMi0wYWZlNWQ5NzU4MzkiLCJpYXQiOjE2OTM4Mzk1NjIsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6InBheXJvbGwiLCJhdWQiOiIzRkRBRURGOS0xRENBNEY1NC04Nzk0OUY2QS00MTAyNzY0MyJ9.D8i_TwguGZ9vok7RCUuk5FoUJGWEGnI-4VzReI1DAg0ytZaQ8XFmVxC4RxGYBxah1ZJpcrbAsuhssGEDXffIgQ\" \\   -X GET ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiJhODdhNGE0My1jMWM2LTQ5ZTktOTUxZS0zNzAzNzRjODhiNjciLCJpYXQiOjE2OTM4Mzk2NjMsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJwYXlyb2xsIiwiYXVkIjoiMDhCQjA4MUUtRDlBQjREMTQtOERGOTkyMzMtNjY2MTVDRTkifQ.AJPINjuEzd9JH1fpf6l1L6L4XdHNIsVcaDsicFe8Giz69zqK1z6za3usbUz71kKJNmm6Eo8HfLnowCi8bcAnow ```  ``` curl --fail \\   \"https://api.affixapi.com/2023-03-01/payroll/employees\" \\   -H \"Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvcGF5cm9sbC9lbXBsb3llZXMiLCIvMjAyMy0wMy0wMS9wYXlyb2xsL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zIiwiLzIwMjMtMDMtMDEvcGF5cm9sbC9wYXlydW5zLzpwYXlydW5faWQiXSwidG9rZW4iOiJhODdhNGE0My1jMWM2LTQ5ZTktOTUxZS0zNzAzNzRjODhiNjciLCJpYXQiOjE2OTM4Mzk2NjMsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJwYXlyb2xsIiwiYXVkIjoiMDhCQjA4MUUtRDlBQjREMTQtOERGOTkyMzMtNjY2MTVDRTkifQ.AJPINjuEzd9JH1fpf6l1L6L4XdHNIsVcaDsicFe8Giz69zqK1z6za3usbUz71kKJNmm6Eo8HfLnowCi8bcAnow\" \\   -X GET ```  # Webhooks An exciting feature for HRIS/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.Employees20230301Response>[     <api.Employee20230301>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.Employee20230301>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: john@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// LocationResponse struct for LocationResponse
type LocationResponse struct {
	StreetAddress NullableString `json:"street_address"`
	Locality NullableString `json:"locality"`
	// The administrative area of the address. If US or CA, the two-letter state abbreviation 
	AdministrativeArea NullableString `json:"administrative_area"`
	// The two-letter abbreviation of the country 
	Country NullableString `json:"country"`
	PostCode NullableString `json:"post_code"`
}

// NewLocationResponse instantiates a new LocationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationResponse(streetAddress NullableString, locality NullableString, administrativeArea NullableString, country NullableString, postCode NullableString) *LocationResponse {
	this := LocationResponse{}
	this.StreetAddress = streetAddress
	this.Locality = locality
	this.AdministrativeArea = administrativeArea
	this.Country = country
	this.PostCode = postCode
	return &this
}

// NewLocationResponseWithDefaults instantiates a new LocationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationResponseWithDefaults() *LocationResponse {
	this := LocationResponse{}
	return &this
}

// GetStreetAddress returns the StreetAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationResponse) GetStreetAddress() string {
	if o == nil || o.StreetAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.StreetAddress.Get()
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationResponse) GetStreetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StreetAddress.Get(), o.StreetAddress.IsSet()
}

// SetStreetAddress sets field value
func (o *LocationResponse) SetStreetAddress(v string) {
	o.StreetAddress.Set(&v)
}

// GetLocality returns the Locality field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationResponse) GetLocality() string {
	if o == nil || o.Locality.Get() == nil {
		var ret string
		return ret
	}

	return *o.Locality.Get()
}

// GetLocalityOk returns a tuple with the Locality field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationResponse) GetLocalityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locality.Get(), o.Locality.IsSet()
}

// SetLocality sets field value
func (o *LocationResponse) SetLocality(v string) {
	o.Locality.Set(&v)
}

// GetAdministrativeArea returns the AdministrativeArea field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationResponse) GetAdministrativeArea() string {
	if o == nil || o.AdministrativeArea.Get() == nil {
		var ret string
		return ret
	}

	return *o.AdministrativeArea.Get()
}

// GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationResponse) GetAdministrativeAreaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdministrativeArea.Get(), o.AdministrativeArea.IsSet()
}

// SetAdministrativeArea sets field value
func (o *LocationResponse) SetAdministrativeArea(v string) {
	o.AdministrativeArea.Set(&v)
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationResponse) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}

	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationResponse) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value
func (o *LocationResponse) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetPostCode returns the PostCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LocationResponse) GetPostCode() string {
	if o == nil || o.PostCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.PostCode.Get()
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationResponse) GetPostCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostCode.Get(), o.PostCode.IsSet()
}

// SetPostCode sets field value
func (o *LocationResponse) SetPostCode(v string) {
	o.PostCode.Set(&v)
}

func (o LocationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["street_address"] = o.StreetAddress.Get()
	}
	if true {
		toSerialize["locality"] = o.Locality.Get()
	}
	if true {
		toSerialize["administrative_area"] = o.AdministrativeArea.Get()
	}
	if true {
		toSerialize["country"] = o.Country.Get()
	}
	if true {
		toSerialize["post_code"] = o.PostCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLocationResponse struct {
	value *LocationResponse
	isSet bool
}

func (v NullableLocationResponse) Get() *LocationResponse {
	return v.value
}

func (v *NullableLocationResponse) Set(val *LocationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationResponse(val *LocationResponse) *NullableLocationResponse {
	return &NullableLocationResponse{value: val, isSet: true}
}

func (v NullableLocationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


