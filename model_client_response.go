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

// ClientResponse struct for ClientResponse
type ClientResponse struct {
	// The client ID you received when you first created the application
	ClientId string `json:"client_id"`
	// The client secret(s). It is an array datatype to allow for rotation of secrets without downtime for your customers 
	ClientSecret []string `json:"client_secret"`
	// Indicates the URI to return the user to after authorization is complete, which is the endpoint on your server to receive the authorization_code.  Must be identical to the redirect URI provided in the original link.  Please email me after signup and I will set both your client secret and redirect_uri (required) when you reach out. 
	RedirectUris []string `json:"redirect_uris"`
	// Name of your app that shows up in the Affix Connect application 
	Name string `json:"name"`
	// If enabled, webhooks will be sent to this endpoint 
	WebhookUri NullableString `json:"webhook_uri,omitempty"`
}

// NewClientResponse instantiates a new ClientResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientResponse(clientId string, clientSecret []string, redirectUris []string, name string) *ClientResponse {
	this := ClientResponse{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.RedirectUris = redirectUris
	this.Name = name
	return &this
}

// NewClientResponseWithDefaults instantiates a new ClientResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientResponseWithDefaults() *ClientResponse {
	this := ClientResponse{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *ClientResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ClientResponse) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ClientResponse) GetClientSecret() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetClientSecretOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ClientResponse) SetClientSecret(v []string) {
	o.ClientSecret = v
}

// GetRedirectUris returns the RedirectUris field value
func (o *ClientResponse) GetRedirectUris() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetRedirectUrisOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *ClientResponse) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetName returns the Name field value
func (o *ClientResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientResponse) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientResponse) SetName(v string) {
	o.Name = v
}

// GetWebhookUri returns the WebhookUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientResponse) GetWebhookUri() string {
	if o == nil || o.WebhookUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebhookUri.Get()
}

// GetWebhookUriOk returns a tuple with the WebhookUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientResponse) GetWebhookUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebhookUri.Get(), o.WebhookUri.IsSet()
}

// HasWebhookUri returns a boolean if a field has been set.
func (o *ClientResponse) HasWebhookUri() bool {
	if o != nil && o.WebhookUri.IsSet() {
		return true
	}

	return false
}

// SetWebhookUri gets a reference to the given NullableString and assigns it to the WebhookUri field.
func (o *ClientResponse) SetWebhookUri(v string) {
	o.WebhookUri.Set(&v)
}
// SetWebhookUriNil sets the value for WebhookUri to be an explicit nil
func (o *ClientResponse) SetWebhookUriNil() {
	o.WebhookUri.Set(nil)
}

// UnsetWebhookUri ensures that no value is present for WebhookUri, not even an explicit nil
func (o *ClientResponse) UnsetWebhookUri() {
	o.WebhookUri.Unset()
}

func (o ClientResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.WebhookUri.IsSet() {
		toSerialize["webhook_uri"] = o.WebhookUri.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableClientResponse struct {
	value *ClientResponse
	isSet bool
}

func (v NullableClientResponse) Get() *ClientResponse {
	return v.value
}

func (v *NullableClientResponse) Set(val *ClientResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientResponse(val *ClientResponse) *NullableClientResponse {
	return &NullableClientResponse{value: val, isSet: true}
}

func (v NullableClientResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


