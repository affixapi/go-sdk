/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (official mode, employees endpoint) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiJhNmNhODc5YS01OGEzLTQ2MGEtYTZlZC04N2E0NmRlMmMyNzMiLCJpYXQiOjE2OTc5ODUxMzEsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6Im9mZmljaWFsIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.Mw-eYth5VL7jpSVfnh88Tl8Cn-6-bKvjnE4GPtmuUdIS7VAvB5ijQksPjOM3FHkF382oh4bym_FAyQN_UE4mmg ```  ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiJhNmNhODc5YS01OGEzLTQ2MGEtYTZlZC04N2E0NmRlMmMyNzMiLCJpYXQiOjE2OTc5ODUxMzEsImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUuZGV2LmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6Im9mZmljaWFsIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.Mw-eYth5VL7jpSVfnh88Tl8Cn-6-bKvjnE4GPtmuUdIS7VAvB5ijQksPjOM3FHkF382oh4bym_FAyQN_UE4mmg' \\   'https://dev.api.affixapi.com/2023-03-01/official/employees' ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiIwYzU2ZjcwMS0wYmFhLTQxOTQtYmU5Ni01ZThiOTExMzZmZDUiLCJpYXQiOjE2OTc5ODUwMTksImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJvZmZpY2lhbCIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.REb0qtwnn--ql2gWFb32FWilezTtJq8USN3Uj4NXoY8aJgwkjisca5ReRh5xyfprKSz_yOEcD1JwTrOlgkvf-Q ```  ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvb2ZmaWNpYWwvaWRlbnRpdHkiXSwidG9rZW4iOiIwYzU2ZjcwMS0wYmFhLTQxOTQtYmU5Ni01ZThiOTExMzZmZDUiLCJpYXQiOjE2OTc5ODUwMTksImlzcyI6InB1YmxpY2FwaS1pbnRlcm1lZGlhdGUucHJvZC5lbmdpbmVlcmluZy5hZmZpeGFwaS5jb20iLCJzdWIiOiJvZmZpY2lhbCIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.REb0qtwnn--ql2gWFb32FWilezTtJq8USN3Uj4NXoY8aJgwkjisca5ReRh5xyfprKSz_yOEcD1JwTrOlgkvf-Q' \\   'https://api.affixapi.com/2023-03-01/official/employees' ```  # Webhooks An exciting feature for HR/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.Employees20230301Response>[     <api.Employee20230301>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.Employee20230301>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: john@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Employee20230301ResponseManager struct for Employee20230301ResponseManager
type Employee20230301ResponseManager struct {
	FirstName NullableString `json:"first_name"`
	LastName NullableString `json:"last_name"`
	Id string `json:"id"`
	WorkEmail NullableString `json:"work_email"`
	RemoteId NullableString `json:"remote_id"`
}

// NewEmployee20230301ResponseManager instantiates a new Employee20230301ResponseManager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee20230301ResponseManager(firstName NullableString, lastName NullableString, id string, workEmail NullableString, remoteId NullableString) *Employee20230301ResponseManager {
	this := Employee20230301ResponseManager{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Id = id
	this.WorkEmail = workEmail
	this.RemoteId = remoteId
	return &this
}

// NewEmployee20230301ResponseManagerWithDefaults instantiates a new Employee20230301ResponseManager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployee20230301ResponseManagerWithDefaults() *Employee20230301ResponseManager {
	this := Employee20230301ResponseManager{}
	return &this
}

// GetFirstName returns the FirstName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301ResponseManager) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301ResponseManager) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// SetFirstName sets field value
func (o *Employee20230301ResponseManager) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// GetLastName returns the LastName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301ResponseManager) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}

	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301ResponseManager) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// SetLastName sets field value
func (o *Employee20230301ResponseManager) SetLastName(v string) {
	o.LastName.Set(&v)
}

// GetId returns the Id field value
func (o *Employee20230301ResponseManager) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Employee20230301ResponseManager) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Employee20230301ResponseManager) SetId(v string) {
	o.Id = v
}

// GetWorkEmail returns the WorkEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301ResponseManager) GetWorkEmail() string {
	if o == nil || o.WorkEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.WorkEmail.Get()
}

// GetWorkEmailOk returns a tuple with the WorkEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301ResponseManager) GetWorkEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkEmail.Get(), o.WorkEmail.IsSet()
}

// SetWorkEmail sets field value
func (o *Employee20230301ResponseManager) SetWorkEmail(v string) {
	o.WorkEmail.Set(&v)
}

// GetRemoteId returns the RemoteId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301ResponseManager) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}

	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301ResponseManager) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// SetRemoteId sets field value
func (o *Employee20230301ResponseManager) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}

func (o Employee20230301ResponseManager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if true {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["work_email"] = o.WorkEmail.Get()
	}
	if true {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEmployee20230301ResponseManager struct {
	value *Employee20230301ResponseManager
	isSet bool
}

func (v NullableEmployee20230301ResponseManager) Get() *Employee20230301ResponseManager {
	return v.value
}

func (v *NullableEmployee20230301ResponseManager) Set(val *Employee20230301ResponseManager) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee20230301ResponseManager) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee20230301ResponseManager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee20230301ResponseManager(val *Employee20230301ResponseManager) *NullableEmployee20230301ResponseManager {
	return &NullableEmployee20230301ResponseManager{value: val, isSet: true}
}

func (v NullableEmployee20230301ResponseManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee20230301ResponseManager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


