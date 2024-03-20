/*
 * Affix API
 *
 * The affixapi.com API documentation.  # Introduction Affix API is an OAuth 2.1 application that allows developers to access customer data, without developers needing to manage or maintain integrations; or collect login credentials or API keys from users for these third party systems.  # OAuth 2.1 Affix API follows the [OAuth 2.1 spec](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-08).  As an OAuth application, Affix API handles not only both the collection of sensitive user credentials or API keys, but also builds and maintains the integrations with the providers, so you don't have to.  # How to obtain an access token in order to get started, you must:   - register a `client_id`   - direct your user to the sign in flow  (`https://connect.affixapi.com`     [with the appropriate query     parameters](https://github.com/affixapi/starter-kit/tree/master/connect))   - capture `authorization_code` we will send to your redirect URI after     the sign in flow is complete and exchange that `authorization_code` for     a Bearer token  # Sandbox keys (developer mode) ### dev ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6ImQ1OTZhMmYzLWYzNzktNGE1ZC1hMmRhLTk4OWJmYWViYTg1ZCIsImlhdCI6MTcwMjkyMDkwMywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5kZXYuZW5naW5lZXJpbmcuYWZmaXhhcGkuY29tIiwic3ViIjoiZGV2ZWxvcGVyIiwiYXVkIjoiM0ZEQUVERjktMURDQTRGNTQtODc5NDlGNkEtNDEwMjc2NDMifQ.VLWYjCQvBS0C3ZA6_J3-U-idZj5EYI2IlDdTjAWBxSIHGufp6cqaVodKsF2BeIqcIeB3P0lW-KL9mY3xGd7ckQ' \\   'https://dev.api.affixapi.com/2023-03-01/developer/employees' ```  ### prod ``` eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw ```  #### `employees` endpoint sample: ``` curl --fail \\   -X GET \\   -H 'Authorization: Bearer eyJhbGciOiJFUzI1NiIsImtpZCI6Ims5RmxwSFR1YklmZWNsUU5QRVZzeFcxazFZZ0Zfbk1BWllOSGVuOFQxdGciLCJ0eXAiOiJKV1MifQ.eyJwcm92aWRlciI6InNhbmRib3giLCJzY29wZXMiOlsiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2NvbXBhbnkiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWUiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvZW1wbG95ZWVzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL2lkZW50aXR5IiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3BheXJ1bnMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvcGF5cnVucy86cGF5cnVuX2lkIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWJhbGFuY2VzIiwiLzIwMjMtMDMtMDEvZGV2ZWxvcGVyL3RpbWUtb2ZmLWVudHJpZXMiLCIvMjAyMy0wMy0wMS9kZXZlbG9wZXIvdGltZXNoZWV0cyJdLCJ0b2tlbiI6IjI5YjFjYTg4LWNlNjktNDgyZC1iNGZjLTkzMWMzZmJkYWM4ZSIsImlhdCI6MTcwMjkyMTA4MywiaXNzIjoicHVibGljYXBpLWludGVybWVkaWF0ZS5wcm9kLmVuZ2luZWVyaW5nLmFmZml4YXBpLmNvbSIsInN1YiI6ImRldmVsb3BlciIsImF1ZCI6IjA4QkIwODFFLUQ5QUI0RDE0LThERjk5MjMzLTY2NjE1Q0U5In0.2zdpFAmiyYiYk6MOcbXNUwwR4M1Fextnaac340x54AidiWXCyw-u9KeavbqfYF6q8a9kcDLrxhJ8Wc_3tIzuVw' \\   'https://api.affixapi.com/2023-03-01/developer/employees' ```  # Compression We support `brotli`, `gzip`, and `deflate` compression algorithms.  To enable, pass the `Accept-Encoding` header with one or all of the values: `br`, `gzip`, `deflate`, or `identity` (no compression)  In the response, you will receive the `Content-Encoding` response header indicating the compression algorithm used in the data payload to enable you to decompress the result. If the `Accept-Encoding: identity` header was passed, no `Content-Encoding` response header is sent back, as no compression algorithm was used.  # Webhooks An exciting feature for HR/Payroll modes are webhooks.  If enabled, your `webhook_uri` is set on your `client_id` for the respective environment: `dev | prod`  Webhooks are configured to make live requests to the underlying integration 1x/hr, and if a difference is detected since the last request, we will send a request to your `webhook_uri` with this shape:  ``` {    added: <api.v20230301.Employees>[     <api.v20230301.Employee>{       ...,       date_of_birth: '2010-08-06',       display_full_name: 'Daija Rogahn',       employee_number: '57993',       employment_status: 'pending',       employment_type: 'other',       employments: [         {           currency: 'eur',           effective_date: '2022-02-25',           employment_type: 'other',           job_title: 'Dynamic Implementation Manager',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 96000,         },       ],       first_name: 'Daija',       ...     }   ],   removed: [],   updated: [     <api.v20230301.Employee>{       ...,       date_of_birth: '2009-11-09',       display_full_name: 'Lourdes Stiedemann',       employee_number: '63189',       employment_status: 'leave',       employment_type: 'full_time',       employments: [         {           currency: 'gbp',           effective_date: '2023-01-16',           employment_type: 'full_time',           job_title: 'Forward Brand Planner',           pay_frequency: 'semimonthly',           pay_period: 'YEAR',           pay_rate: 86000,         },       ],       first_name: 'Lourdes',     }   ] } ```  the following headers will be sent with webhook requests:  ``` x-affix-api-signature: ab8474e609db95d5df3adc39ea3add7a7544bd215c5c520a30a650ae93a2fba7  x-affix-api-origin:  webhooks-employees-webhook  user-agent:  affixapi.com ```  Before trusting the payload, you should sign the payload and verify the signature matches the signature sent by the `affixapi.com` service.  This secures that the data sent to your `webhook_uri` is from the `affixapi.com` server.  The signature is created by combining the signing secret (your `client_secret`) with the body of the request sent using a standard HMAC-SHA256 keyed hash.  The signature can be created via:   - create an `HMAC` with your `client_secret`   - update the `HMAC` with the payload   - get the hex digest -> this is the signature  Sample `typescript` code that follows this recipe:  ``` import { createHmac } from 'crypto';  export const computeSignature = ({   str,   signingSecret, }: {   signingSecret: string;   str: string; }): string => {   const hmac = createHmac('sha256', signingSecret);   hmac.update(str);   const signature = hmac.digest('hex');    return signature; }; ```  While verifying the Affix API signature header should be your primary method of confirming validity, you can also whitelist our outbound webhook static IP addresses.  ``` dev:   - 52.210.169.82   - 52.210.38.77   - 3.248.135.204  prod:   - 52.51.160.102   - 54.220.83.244   - 3.254.213.171 ```  ## Rate limits Open endpoints (not gated by an API key) (applied at endpoint level):   - 15 requests every 1 minute (by IP address)   - 25 requests every 5 minutes (by IP address)  Gated endpoints (require an API key) (applied at endpoint level):   - 40 requests every 1 minute (by IP address)   - 40 requests every 5 minutes (by `client_id`)  Things to keep in mind:   - Open endpoints (not gated by an API key) will likely be called by your     users, not you, so rate limits generally would not apply to you.   - As a developer, rate limits are applied at the endpoint granularity.     - For example, say the rate limits below are 10 requests per minute by ip.       from that same ip, within 1 minute, you get:       - 10 requests per minute on `/orders`,       - another 10 requests per minute on `/items`,       - and another 10 requests per minute on `/identity`,       - for a total of 30 requests per minute. 
 *
 * API version: 2023-03-01
 * Contact: developers@affixapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateEmployeeRequest struct for CreateEmployeeRequest
type CreateEmployeeRequest struct {
	EmployeeNumber NullableString `json:"employee_number,omitempty"`
	// the first name of the individual
	FirstName string `json:"first_name"`
	// the last name of the individual
	LastName string `json:"last_name"`
	DisplayFullName NullableString `json:"display_full_name,omitempty"`
	Nationality NullableString `json:"nationality,omitempty"`
	JobTitle NullableString `json:"job_title,omitempty"`
	// the work email of the individual
	WorkEmail NullableString `json:"work_email,omitempty"`
	// the personal email of the individual
	PersonalEmail NullableString `json:"personal_email,omitempty"`
	// +1234567890
	MobilePhoneNumber NullableString `json:"mobile_phone_number,omitempty"`
	TaxId NullableString `json:"tax_id,omitempty"`
	Gender NullableString `json:"gender,omitempty"`
	Ethnicity NullableString `json:"ethnicity,omitempty"`
	MaritalStatus NullableString `json:"marital_status,omitempty"`
	DateOfBirth NullableString `json:"date_of_birth,omitempty"`
	EmploymentStatus NullableString `json:"employment_status,omitempty"`
	EmploymentType NullableString `json:"employment_type,omitempty"`
	StartDate NullableString `json:"start_date,omitempty"`
	TerminationDate NullableString `json:"termination_date,omitempty"`
	Avatar NullableString `json:"avatar,omitempty"`
	HomeLocation NullableAddressNoNonNullRequest `json:"home_location,omitempty"`
	WorkLocation NullableLocationNoNonNullRequest `json:"work_location,omitempty"`
	Manager NullableCreateEmployeeRequestManager `json:"manager,omitempty"`
	BankAccount NullableCreateEmployeeRequestBankAccount `json:"bank_account,omitempty"`
	Employments []EmploymentNoNullEnumRequest `json:"employments,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Groups []GroupNoNullEnumRequest `json:"groups,omitempty"`
}

// NewCreateEmployeeRequest instantiates a new CreateEmployeeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmployeeRequest(firstName string, lastName string) *CreateEmployeeRequest {
	this := CreateEmployeeRequest{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewCreateEmployeeRequestWithDefaults instantiates a new CreateEmployeeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmployeeRequestWithDefaults() *CreateEmployeeRequest {
	this := CreateEmployeeRequest{}
	return &this
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNumber.Get()
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetEmployeeNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNumber.Get(), o.EmployeeNumber.IsSet()
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasEmployeeNumber() bool {
	if o != nil && o.EmployeeNumber.IsSet() {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given NullableString and assigns it to the EmployeeNumber field.
func (o *CreateEmployeeRequest) SetEmployeeNumber(v string) {
	o.EmployeeNumber.Set(&v)
}
// SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil
func (o *CreateEmployeeRequest) SetEmployeeNumberNil() {
	o.EmployeeNumber.Set(nil)
}

// UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetEmployeeNumber() {
	o.EmployeeNumber.Unset()
}

// GetFirstName returns the FirstName field value
func (o *CreateEmployeeRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreateEmployeeRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreateEmployeeRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CreateEmployeeRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreateEmployeeRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreateEmployeeRequest) SetLastName(v string) {
	o.LastName = v
}

// GetDisplayFullName returns the DisplayFullName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetDisplayFullName() string {
	if o == nil || o.DisplayFullName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayFullName.Get()
}

// GetDisplayFullNameOk returns a tuple with the DisplayFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetDisplayFullNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayFullName.Get(), o.DisplayFullName.IsSet()
}

// HasDisplayFullName returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasDisplayFullName() bool {
	if o != nil && o.DisplayFullName.IsSet() {
		return true
	}

	return false
}

// SetDisplayFullName gets a reference to the given NullableString and assigns it to the DisplayFullName field.
func (o *CreateEmployeeRequest) SetDisplayFullName(v string) {
	o.DisplayFullName.Set(&v)
}
// SetDisplayFullNameNil sets the value for DisplayFullName to be an explicit nil
func (o *CreateEmployeeRequest) SetDisplayFullNameNil() {
	o.DisplayFullName.Set(nil)
}

// UnsetDisplayFullName ensures that no value is present for DisplayFullName, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetDisplayFullName() {
	o.DisplayFullName.Unset()
}

// GetNationality returns the Nationality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetNationality() string {
	if o == nil || o.Nationality.Get() == nil {
		var ret string
		return ret
	}
	return *o.Nationality.Get()
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetNationalityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nationality.Get(), o.Nationality.IsSet()
}

// HasNationality returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasNationality() bool {
	if o != nil && o.Nationality.IsSet() {
		return true
	}

	return false
}

// SetNationality gets a reference to the given NullableString and assigns it to the Nationality field.
func (o *CreateEmployeeRequest) SetNationality(v string) {
	o.Nationality.Set(&v)
}
// SetNationalityNil sets the value for Nationality to be an explicit nil
func (o *CreateEmployeeRequest) SetNationalityNil() {
	o.Nationality.Set(nil)
}

// UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetNationality() {
	o.Nationality.Unset()
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetJobTitle() string {
	if o == nil || o.JobTitle.Get() == nil {
		var ret string
		return ret
	}
	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// HasJobTitle returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasJobTitle() bool {
	if o != nil && o.JobTitle.IsSet() {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given NullableString and assigns it to the JobTitle field.
func (o *CreateEmployeeRequest) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}
// SetJobTitleNil sets the value for JobTitle to be an explicit nil
func (o *CreateEmployeeRequest) SetJobTitleNil() {
	o.JobTitle.Set(nil)
}

// UnsetJobTitle ensures that no value is present for JobTitle, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetJobTitle() {
	o.JobTitle.Unset()
}

// GetWorkEmail returns the WorkEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetWorkEmail() string {
	if o == nil || o.WorkEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.WorkEmail.Get()
}

// GetWorkEmailOk returns a tuple with the WorkEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetWorkEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkEmail.Get(), o.WorkEmail.IsSet()
}

// HasWorkEmail returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasWorkEmail() bool {
	if o != nil && o.WorkEmail.IsSet() {
		return true
	}

	return false
}

// SetWorkEmail gets a reference to the given NullableString and assigns it to the WorkEmail field.
func (o *CreateEmployeeRequest) SetWorkEmail(v string) {
	o.WorkEmail.Set(&v)
}
// SetWorkEmailNil sets the value for WorkEmail to be an explicit nil
func (o *CreateEmployeeRequest) SetWorkEmailNil() {
	o.WorkEmail.Set(nil)
}

// UnsetWorkEmail ensures that no value is present for WorkEmail, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetWorkEmail() {
	o.WorkEmail.Unset()
}

// GetPersonalEmail returns the PersonalEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetPersonalEmail() string {
	if o == nil || o.PersonalEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.PersonalEmail.Get()
}

// GetPersonalEmailOk returns a tuple with the PersonalEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetPersonalEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalEmail.Get(), o.PersonalEmail.IsSet()
}

// HasPersonalEmail returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasPersonalEmail() bool {
	if o != nil && o.PersonalEmail.IsSet() {
		return true
	}

	return false
}

// SetPersonalEmail gets a reference to the given NullableString and assigns it to the PersonalEmail field.
func (o *CreateEmployeeRequest) SetPersonalEmail(v string) {
	o.PersonalEmail.Set(&v)
}
// SetPersonalEmailNil sets the value for PersonalEmail to be an explicit nil
func (o *CreateEmployeeRequest) SetPersonalEmailNil() {
	o.PersonalEmail.Set(nil)
}

// UnsetPersonalEmail ensures that no value is present for PersonalEmail, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetPersonalEmail() {
	o.PersonalEmail.Unset()
}

// GetMobilePhoneNumber returns the MobilePhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetMobilePhoneNumber() string {
	if o == nil || o.MobilePhoneNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.MobilePhoneNumber.Get()
}

// GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetMobilePhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MobilePhoneNumber.Get(), o.MobilePhoneNumber.IsSet()
}

// HasMobilePhoneNumber returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasMobilePhoneNumber() bool {
	if o != nil && o.MobilePhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetMobilePhoneNumber gets a reference to the given NullableString and assigns it to the MobilePhoneNumber field.
func (o *CreateEmployeeRequest) SetMobilePhoneNumber(v string) {
	o.MobilePhoneNumber.Set(&v)
}
// SetMobilePhoneNumberNil sets the value for MobilePhoneNumber to be an explicit nil
func (o *CreateEmployeeRequest) SetMobilePhoneNumberNil() {
	o.MobilePhoneNumber.Set(nil)
}

// UnsetMobilePhoneNumber ensures that no value is present for MobilePhoneNumber, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetMobilePhoneNumber() {
	o.MobilePhoneNumber.Unset()
}

// GetTaxId returns the TaxId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetTaxId() string {
	if o == nil || o.TaxId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetTaxIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// HasTaxId returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasTaxId() bool {
	if o != nil && o.TaxId.IsSet() {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given NullableString and assigns it to the TaxId field.
func (o *CreateEmployeeRequest) SetTaxId(v string) {
	o.TaxId.Set(&v)
}
// SetTaxIdNil sets the value for TaxId to be an explicit nil
func (o *CreateEmployeeRequest) SetTaxIdNil() {
	o.TaxId.Set(nil)
}

// UnsetTaxId ensures that no value is present for TaxId, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetTaxId() {
	o.TaxId.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetGender() string {
	if o == nil || o.Gender.Get() == nil {
		var ret string
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetGenderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableString and assigns it to the Gender field.
func (o *CreateEmployeeRequest) SetGender(v string) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *CreateEmployeeRequest) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetGender() {
	o.Gender.Unset()
}

// GetEthnicity returns the Ethnicity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetEthnicity() string {
	if o == nil || o.Ethnicity.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ethnicity.Get()
}

// GetEthnicityOk returns a tuple with the Ethnicity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetEthnicityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ethnicity.Get(), o.Ethnicity.IsSet()
}

// HasEthnicity returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasEthnicity() bool {
	if o != nil && o.Ethnicity.IsSet() {
		return true
	}

	return false
}

// SetEthnicity gets a reference to the given NullableString and assigns it to the Ethnicity field.
func (o *CreateEmployeeRequest) SetEthnicity(v string) {
	o.Ethnicity.Set(&v)
}
// SetEthnicityNil sets the value for Ethnicity to be an explicit nil
func (o *CreateEmployeeRequest) SetEthnicityNil() {
	o.Ethnicity.Set(nil)
}

// UnsetEthnicity ensures that no value is present for Ethnicity, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetEthnicity() {
	o.Ethnicity.Unset()
}

// GetMaritalStatus returns the MaritalStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetMaritalStatus() string {
	if o == nil || o.MaritalStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaritalStatus.Get()
}

// GetMaritalStatusOk returns a tuple with the MaritalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetMaritalStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaritalStatus.Get(), o.MaritalStatus.IsSet()
}

// HasMaritalStatus returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasMaritalStatus() bool {
	if o != nil && o.MaritalStatus.IsSet() {
		return true
	}

	return false
}

// SetMaritalStatus gets a reference to the given NullableString and assigns it to the MaritalStatus field.
func (o *CreateEmployeeRequest) SetMaritalStatus(v string) {
	o.MaritalStatus.Set(&v)
}
// SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil
func (o *CreateEmployeeRequest) SetMaritalStatusNil() {
	o.MaritalStatus.Set(nil)
}

// UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetMaritalStatus() {
	o.MaritalStatus.Unset()
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth.IsSet() {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given NullableString and assigns it to the DateOfBirth field.
func (o *CreateEmployeeRequest) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}
// SetDateOfBirthNil sets the value for DateOfBirth to be an explicit nil
func (o *CreateEmployeeRequest) SetDateOfBirthNil() {
	o.DateOfBirth.Set(nil)
}

// UnsetDateOfBirth ensures that no value is present for DateOfBirth, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetDateOfBirth() {
	o.DateOfBirth.Unset()
}

// GetEmploymentStatus returns the EmploymentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetEmploymentStatus() string {
	if o == nil || o.EmploymentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmploymentStatus.Get()
}

// GetEmploymentStatusOk returns a tuple with the EmploymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetEmploymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentStatus.Get(), o.EmploymentStatus.IsSet()
}

// HasEmploymentStatus returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasEmploymentStatus() bool {
	if o != nil && o.EmploymentStatus.IsSet() {
		return true
	}

	return false
}

// SetEmploymentStatus gets a reference to the given NullableString and assigns it to the EmploymentStatus field.
func (o *CreateEmployeeRequest) SetEmploymentStatus(v string) {
	o.EmploymentStatus.Set(&v)
}
// SetEmploymentStatusNil sets the value for EmploymentStatus to be an explicit nil
func (o *CreateEmployeeRequest) SetEmploymentStatusNil() {
	o.EmploymentStatus.Set(nil)
}

// UnsetEmploymentStatus ensures that no value is present for EmploymentStatus, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetEmploymentStatus() {
	o.EmploymentStatus.Unset()
}

// GetEmploymentType returns the EmploymentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetEmploymentType() string {
	if o == nil || o.EmploymentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmploymentType.Get()
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetEmploymentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentType.Get(), o.EmploymentType.IsSet()
}

// HasEmploymentType returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasEmploymentType() bool {
	if o != nil && o.EmploymentType.IsSet() {
		return true
	}

	return false
}

// SetEmploymentType gets a reference to the given NullableString and assigns it to the EmploymentType field.
func (o *CreateEmployeeRequest) SetEmploymentType(v string) {
	o.EmploymentType.Set(&v)
}
// SetEmploymentTypeNil sets the value for EmploymentType to be an explicit nil
func (o *CreateEmployeeRequest) SetEmploymentTypeNil() {
	o.EmploymentType.Set(nil)
}

// UnsetEmploymentType ensures that no value is present for EmploymentType, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetEmploymentType() {
	o.EmploymentType.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *CreateEmployeeRequest) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *CreateEmployeeRequest) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetTerminationDate returns the TerminationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetTerminationDate() string {
	if o == nil || o.TerminationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetTerminationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// HasTerminationDate returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasTerminationDate() bool {
	if o != nil && o.TerminationDate.IsSet() {
		return true
	}

	return false
}

// SetTerminationDate gets a reference to the given NullableString and assigns it to the TerminationDate field.
func (o *CreateEmployeeRequest) SetTerminationDate(v string) {
	o.TerminationDate.Set(&v)
}
// SetTerminationDateNil sets the value for TerminationDate to be an explicit nil
func (o *CreateEmployeeRequest) SetTerminationDateNil() {
	o.TerminationDate.Set(nil)
}

// UnsetTerminationDate ensures that no value is present for TerminationDate, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetTerminationDate() {
	o.TerminationDate.Unset()
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetAvatarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *CreateEmployeeRequest) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *CreateEmployeeRequest) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetHomeLocation returns the HomeLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetHomeLocation() AddressNoNonNullRequest {
	if o == nil || o.HomeLocation.Get() == nil {
		var ret AddressNoNonNullRequest
		return ret
	}
	return *o.HomeLocation.Get()
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetHomeLocationOk() (*AddressNoNonNullRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HomeLocation.Get(), o.HomeLocation.IsSet()
}

// HasHomeLocation returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasHomeLocation() bool {
	if o != nil && o.HomeLocation.IsSet() {
		return true
	}

	return false
}

// SetHomeLocation gets a reference to the given NullableAddressNoNonNullRequest and assigns it to the HomeLocation field.
func (o *CreateEmployeeRequest) SetHomeLocation(v AddressNoNonNullRequest) {
	o.HomeLocation.Set(&v)
}
// SetHomeLocationNil sets the value for HomeLocation to be an explicit nil
func (o *CreateEmployeeRequest) SetHomeLocationNil() {
	o.HomeLocation.Set(nil)
}

// UnsetHomeLocation ensures that no value is present for HomeLocation, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetHomeLocation() {
	o.HomeLocation.Unset()
}

// GetWorkLocation returns the WorkLocation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetWorkLocation() LocationNoNonNullRequest {
	if o == nil || o.WorkLocation.Get() == nil {
		var ret LocationNoNonNullRequest
		return ret
	}
	return *o.WorkLocation.Get()
}

// GetWorkLocationOk returns a tuple with the WorkLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetWorkLocationOk() (*LocationNoNonNullRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkLocation.Get(), o.WorkLocation.IsSet()
}

// HasWorkLocation returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasWorkLocation() bool {
	if o != nil && o.WorkLocation.IsSet() {
		return true
	}

	return false
}

// SetWorkLocation gets a reference to the given NullableLocationNoNonNullRequest and assigns it to the WorkLocation field.
func (o *CreateEmployeeRequest) SetWorkLocation(v LocationNoNonNullRequest) {
	o.WorkLocation.Set(&v)
}
// SetWorkLocationNil sets the value for WorkLocation to be an explicit nil
func (o *CreateEmployeeRequest) SetWorkLocationNil() {
	o.WorkLocation.Set(nil)
}

// UnsetWorkLocation ensures that no value is present for WorkLocation, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetWorkLocation() {
	o.WorkLocation.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetManager() CreateEmployeeRequestManager {
	if o == nil || o.Manager.Get() == nil {
		var ret CreateEmployeeRequestManager
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetManagerOk() (*CreateEmployeeRequestManager, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableCreateEmployeeRequestManager and assigns it to the Manager field.
func (o *CreateEmployeeRequest) SetManager(v CreateEmployeeRequestManager) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *CreateEmployeeRequest) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetManager() {
	o.Manager.Unset()
}

// GetBankAccount returns the BankAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetBankAccount() CreateEmployeeRequestBankAccount {
	if o == nil || o.BankAccount.Get() == nil {
		var ret CreateEmployeeRequestBankAccount
		return ret
	}
	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetBankAccountOk() (*CreateEmployeeRequestBankAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// HasBankAccount returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasBankAccount() bool {
	if o != nil && o.BankAccount.IsSet() {
		return true
	}

	return false
}

// SetBankAccount gets a reference to the given NullableCreateEmployeeRequestBankAccount and assigns it to the BankAccount field.
func (o *CreateEmployeeRequest) SetBankAccount(v CreateEmployeeRequestBankAccount) {
	o.BankAccount.Set(&v)
}
// SetBankAccountNil sets the value for BankAccount to be an explicit nil
func (o *CreateEmployeeRequest) SetBankAccountNil() {
	o.BankAccount.Set(nil)
}

// UnsetBankAccount ensures that no value is present for BankAccount, not even an explicit nil
func (o *CreateEmployeeRequest) UnsetBankAccount() {
	o.BankAccount.Unset()
}

// GetEmployments returns the Employments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetEmployments() []EmploymentNoNullEnumRequest {
	if o == nil  {
		var ret []EmploymentNoNullEnumRequest
		return ret
	}
	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetEmploymentsOk() (*[]EmploymentNoNullEnumRequest, bool) {
	if o == nil || o.Employments == nil {
		return nil, false
	}
	return &o.Employments, true
}

// HasEmployments returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasEmployments() bool {
	if o != nil && o.Employments != nil {
		return true
	}

	return false
}

// SetEmployments gets a reference to the given []EmploymentNoNullEnumRequest and assigns it to the Employments field.
func (o *CreateEmployeeRequest) SetEmployments(v []EmploymentNoNullEnumRequest) {
	o.Employments = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetCustomFields() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetCustomFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return &o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasCustomFields() bool {
	if o != nil && o.CustomFields != nil {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *CreateEmployeeRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateEmployeeRequest) GetGroups() []GroupNoNullEnumRequest {
	if o == nil  {
		var ret []GroupNoNullEnumRequest
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateEmployeeRequest) GetGroupsOk() (*[]GroupNoNullEnumRequest, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CreateEmployeeRequest) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupNoNullEnumRequest and assigns it to the Groups field.
func (o *CreateEmployeeRequest) SetGroups(v []GroupNoNullEnumRequest) {
	o.Groups = v
}

func (o CreateEmployeeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmployeeNumber.IsSet() {
		toSerialize["employee_number"] = o.EmployeeNumber.Get()
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if o.DisplayFullName.IsSet() {
		toSerialize["display_full_name"] = o.DisplayFullName.Get()
	}
	if o.Nationality.IsSet() {
		toSerialize["nationality"] = o.Nationality.Get()
	}
	if o.JobTitle.IsSet() {
		toSerialize["job_title"] = o.JobTitle.Get()
	}
	if o.WorkEmail.IsSet() {
		toSerialize["work_email"] = o.WorkEmail.Get()
	}
	if o.PersonalEmail.IsSet() {
		toSerialize["personal_email"] = o.PersonalEmail.Get()
	}
	if o.MobilePhoneNumber.IsSet() {
		toSerialize["mobile_phone_number"] = o.MobilePhoneNumber.Get()
	}
	if o.TaxId.IsSet() {
		toSerialize["tax_id"] = o.TaxId.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Ethnicity.IsSet() {
		toSerialize["ethnicity"] = o.Ethnicity.Get()
	}
	if o.MaritalStatus.IsSet() {
		toSerialize["marital_status"] = o.MaritalStatus.Get()
	}
	if o.DateOfBirth.IsSet() {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if o.EmploymentStatus.IsSet() {
		toSerialize["employment_status"] = o.EmploymentStatus.Get()
	}
	if o.EmploymentType.IsSet() {
		toSerialize["employment_type"] = o.EmploymentType.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.TerminationDate.IsSet() {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.HomeLocation.IsSet() {
		toSerialize["home_location"] = o.HomeLocation.Get()
	}
	if o.WorkLocation.IsSet() {
		toSerialize["work_location"] = o.WorkLocation.Get()
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if o.BankAccount.IsSet() {
		toSerialize["bank_account"] = o.BankAccount.Get()
	}
	if o.Employments != nil {
		toSerialize["employments"] = o.Employments
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEmployeeRequest struct {
	value *CreateEmployeeRequest
	isSet bool
}

func (v NullableCreateEmployeeRequest) Get() *CreateEmployeeRequest {
	return v.value
}

func (v *NullableCreateEmployeeRequest) Set(val *CreateEmployeeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmployeeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmployeeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmployeeRequest(val *CreateEmployeeRequest) *NullableCreateEmployeeRequest {
	return &NullableCreateEmployeeRequest{value: val, isSet: true}
}

func (v NullableCreateEmployeeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmployeeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


