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

// TimeOffEntryResponse struct for TimeOffEntryResponse
type TimeOffEntryResponse struct {
	// The Affix-assigned id of the time off entry
	Id string `json:"id"`
	// the remote system-assigned id of the time off entry
	RemoteId string `json:"remote_id"`
	// the Affix-assigned id of the individual
	EmployeeId string `json:"employee_id"`
	// the remote system-assigned id of the individual
	EmployeeRemoteId string `json:"employee_remote_id"`
	StartDate NullableString `json:"start_date"`
	EndDate NullableString `json:"end_date"`
	Amount float32 `json:"amount"`
	Unit string `json:"unit"`
	Status NullableString `json:"status"`
	EmployeeNote NullableString `json:"employee_note"`
	// The Affix-assigned id of the policy
	PolicyId NullableString `json:"policy_id"`
	// The remote system-assigned id of the policy
	PolicyRemoteId NullableString `json:"policy_remote_id"`
	// The name of the policy, as assigned by the remote system
	PolicyName NullableString `json:"policy_name"`
	PolicyType NullablePolicyTypeResponse `json:"policy_type"`
	RemoteCreatedAt NullableString `json:"remote_created_at"`
	RemoteModifiedAt NullableString `json:"remote_modified_at"`
}

// NewTimeOffEntryResponse instantiates a new TimeOffEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeOffEntryResponse(id string, remoteId string, employeeId string, employeeRemoteId string, startDate NullableString, endDate NullableString, amount float32, unit string, status NullableString, employeeNote NullableString, policyId NullableString, policyRemoteId NullableString, policyName NullableString, policyType NullablePolicyTypeResponse, remoteCreatedAt NullableString, remoteModifiedAt NullableString) *TimeOffEntryResponse {
	this := TimeOffEntryResponse{}
	this.Id = id
	this.RemoteId = remoteId
	this.EmployeeId = employeeId
	this.EmployeeRemoteId = employeeRemoteId
	this.StartDate = startDate
	this.EndDate = endDate
	this.Amount = amount
	this.Unit = unit
	this.Status = status
	this.EmployeeNote = employeeNote
	this.PolicyId = policyId
	this.PolicyRemoteId = policyRemoteId
	this.PolicyName = policyName
	this.PolicyType = policyType
	this.RemoteCreatedAt = remoteCreatedAt
	this.RemoteModifiedAt = remoteModifiedAt
	return &this
}

// NewTimeOffEntryResponseWithDefaults instantiates a new TimeOffEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeOffEntryResponseWithDefaults() *TimeOffEntryResponse {
	this := TimeOffEntryResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TimeOffEntryResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TimeOffEntryResponse) SetId(v string) {
	o.Id = v
}

// GetRemoteId returns the RemoteId field value
func (o *TimeOffEntryResponse) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *TimeOffEntryResponse) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetEmployeeId returns the EmployeeId field value
func (o *TimeOffEntryResponse) GetEmployeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetEmployeeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeId, true
}

// SetEmployeeId sets field value
func (o *TimeOffEntryResponse) SetEmployeeId(v string) {
	o.EmployeeId = v
}

// GetEmployeeRemoteId returns the EmployeeRemoteId field value
func (o *TimeOffEntryResponse) GetEmployeeRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeRemoteId
}

// GetEmployeeRemoteIdOk returns a tuple with the EmployeeRemoteId field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetEmployeeRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeRemoteId, true
}

// SetEmployeeRemoteId sets field value
func (o *TimeOffEntryResponse) SetEmployeeRemoteId(v string) {
	o.EmployeeRemoteId = v
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *TimeOffEntryResponse) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// GetEndDate returns the EndDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// SetEndDate sets field value
func (o *TimeOffEntryResponse) SetEndDate(v string) {
	o.EndDate.Set(&v)
}

// GetAmount returns the Amount field value
func (o *TimeOffEntryResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TimeOffEntryResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetUnit returns the Unit field value
func (o *TimeOffEntryResponse) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *TimeOffEntryResponse) GetUnitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *TimeOffEntryResponse) SetUnit(v string) {
	o.Unit = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *TimeOffEntryResponse) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetEmployeeNote returns the EmployeeNote field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetEmployeeNote() string {
	if o == nil || o.EmployeeNote.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployeeNote.Get()
}

// GetEmployeeNoteOk returns a tuple with the EmployeeNote field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetEmployeeNoteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNote.Get(), o.EmployeeNote.IsSet()
}

// SetEmployeeNote sets field value
func (o *TimeOffEntryResponse) SetEmployeeNote(v string) {
	o.EmployeeNote.Set(&v)
}

// GetPolicyId returns the PolicyId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetPolicyId() string {
	if o == nil || o.PolicyId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PolicyId.Get()
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetPolicyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyId.Get(), o.PolicyId.IsSet()
}

// SetPolicyId sets field value
func (o *TimeOffEntryResponse) SetPolicyId(v string) {
	o.PolicyId.Set(&v)
}

// GetPolicyRemoteId returns the PolicyRemoteId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetPolicyRemoteId() string {
	if o == nil || o.PolicyRemoteId.Get() == nil {
		var ret string
		return ret
	}

	return *o.PolicyRemoteId.Get()
}

// GetPolicyRemoteIdOk returns a tuple with the PolicyRemoteId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetPolicyRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyRemoteId.Get(), o.PolicyRemoteId.IsSet()
}

// SetPolicyRemoteId sets field value
func (o *TimeOffEntryResponse) SetPolicyRemoteId(v string) {
	o.PolicyRemoteId.Set(&v)
}

// GetPolicyName returns the PolicyName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetPolicyName() string {
	if o == nil || o.PolicyName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PolicyName.Get()
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetPolicyNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyName.Get(), o.PolicyName.IsSet()
}

// SetPolicyName sets field value
func (o *TimeOffEntryResponse) SetPolicyName(v string) {
	o.PolicyName.Set(&v)
}

// GetPolicyType returns the PolicyType field value
// If the value is explicit nil, the zero value for PolicyTypeResponse will be returned
func (o *TimeOffEntryResponse) GetPolicyType() PolicyTypeResponse {
	if o == nil || o.PolicyType.Get() == nil {
		var ret PolicyTypeResponse
		return ret
	}

	return *o.PolicyType.Get()
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetPolicyTypeOk() (*PolicyTypeResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PolicyType.Get(), o.PolicyType.IsSet()
}

// SetPolicyType sets field value
func (o *TimeOffEntryResponse) SetPolicyType(v PolicyTypeResponse) {
	o.PolicyType.Set(&v)
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetRemoteCreatedAt() string {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetRemoteCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// SetRemoteCreatedAt sets field value
func (o *TimeOffEntryResponse) SetRemoteCreatedAt(v string) {
	o.RemoteCreatedAt.Set(&v)
}

// GetRemoteModifiedAt returns the RemoteModifiedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TimeOffEntryResponse) GetRemoteModifiedAt() string {
	if o == nil || o.RemoteModifiedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.RemoteModifiedAt.Get()
}

// GetRemoteModifiedAtOk returns a tuple with the RemoteModifiedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeOffEntryResponse) GetRemoteModifiedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteModifiedAt.Get(), o.RemoteModifiedAt.IsSet()
}

// SetRemoteModifiedAt sets field value
func (o *TimeOffEntryResponse) SetRemoteModifiedAt(v string) {
	o.RemoteModifiedAt.Set(&v)
}

func (o TimeOffEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["remote_id"] = o.RemoteId
	}
	if true {
		toSerialize["employee_id"] = o.EmployeeId
	}
	if true {
		toSerialize["employee_remote_id"] = o.EmployeeRemoteId
	}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["employee_note"] = o.EmployeeNote.Get()
	}
	if true {
		toSerialize["policy_id"] = o.PolicyId.Get()
	}
	if true {
		toSerialize["policy_remote_id"] = o.PolicyRemoteId.Get()
	}
	if true {
		toSerialize["policy_name"] = o.PolicyName.Get()
	}
	if true {
		toSerialize["policy_type"] = o.PolicyType.Get()
	}
	if true {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if true {
		toSerialize["remote_modified_at"] = o.RemoteModifiedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTimeOffEntryResponse struct {
	value *TimeOffEntryResponse
	isSet bool
}

func (v NullableTimeOffEntryResponse) Get() *TimeOffEntryResponse {
	return v.value
}

func (v *NullableTimeOffEntryResponse) Set(val *TimeOffEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeOffEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeOffEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeOffEntryResponse(val *TimeOffEntryResponse) *NullableTimeOffEntryResponse {
	return &NullableTimeOffEntryResponse{value: val, isSet: true}
}

func (v NullableTimeOffEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeOffEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


