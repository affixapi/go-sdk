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

// Employee20230301Response struct for Employee20230301Response
type Employee20230301Response struct {
	// The Affix-assigned id of the individual
	Id string `json:"id"`
	// the remote system-assigned id of the individual
	RemoteId string `json:"remote_id"`
	EmployeeNumber NullableString `json:"employee_number"`
	// the first name of the individual
	FirstName string `json:"first_name"`
	// the last name of the individual
	LastName string `json:"last_name"`
	DisplayFullName NullableString `json:"display_full_name"`
	Nationality NullableString `json:"nationality"`
	JobTitle NullableString `json:"job_title"`
	// the work email of the individual
	WorkEmail NullableString `json:"work_email"`
	// the personal email of the individual
	PersonalEmail NullableString `json:"personal_email"`
	// +1234567890
	MobilePhoneNumber NullableString `json:"mobile_phone_number"`
	TaxId NullableString `json:"tax_id"`
	Gender NullableString `json:"gender"`
	Ethnicity NullableString `json:"ethnicity"`
	MaritalStatus NullableString `json:"marital_status"`
	DateOfBirth NullableString `json:"date_of_birth"`
	EmploymentStatus NullableString `json:"employment_status"`
	EmploymentType NullableString `json:"employment_type"`
	StartDate NullableString `json:"start_date"`
	RemoteCreatedAt NullableString `json:"remote_created_at"`
	TerminationDate NullableString `json:"termination_date"`
	Avatar NullableString `json:"avatar"`
	HomeLocation Employee20230301ResponseHomeLocation `json:"home_location"`
	WorkLocation NullableEmployee20230301ResponseWorkLocation `json:"work_location"`
	Manager NullableEmployee20230301ResponseManager `json:"manager"`
	BankAccount NullableEmployee20230301ResponseBankAccount `json:"bank_account"`
	Employments []EmploymentResponse `json:"employments"`
	CustomFields map[string]interface{} `json:"custom_fields"`
	Groups []Employee20230301ResponseGroups `json:"groups"`
	Company NullableEmployee20230301ResponseCompany `json:"company"`
}

// NewEmployee20230301Response instantiates a new Employee20230301Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee20230301Response(id string, remoteId string, employeeNumber NullableString, firstName string, lastName string, displayFullName NullableString, nationality NullableString, jobTitle NullableString, workEmail NullableString, personalEmail NullableString, mobilePhoneNumber NullableString, taxId NullableString, gender NullableString, ethnicity NullableString, maritalStatus NullableString, dateOfBirth NullableString, employmentStatus NullableString, employmentType NullableString, startDate NullableString, remoteCreatedAt NullableString, terminationDate NullableString, avatar NullableString, homeLocation Employee20230301ResponseHomeLocation, workLocation NullableEmployee20230301ResponseWorkLocation, manager NullableEmployee20230301ResponseManager, bankAccount NullableEmployee20230301ResponseBankAccount, employments []EmploymentResponse, customFields map[string]interface{}, groups []Employee20230301ResponseGroups, company NullableEmployee20230301ResponseCompany) *Employee20230301Response {
	this := Employee20230301Response{}
	this.Id = id
	this.RemoteId = remoteId
	this.EmployeeNumber = employeeNumber
	this.FirstName = firstName
	this.LastName = lastName
	this.DisplayFullName = displayFullName
	this.Nationality = nationality
	this.JobTitle = jobTitle
	this.WorkEmail = workEmail
	this.PersonalEmail = personalEmail
	this.MobilePhoneNumber = mobilePhoneNumber
	this.TaxId = taxId
	this.Gender = gender
	this.Ethnicity = ethnicity
	this.MaritalStatus = maritalStatus
	this.DateOfBirth = dateOfBirth
	this.EmploymentStatus = employmentStatus
	this.EmploymentType = employmentType
	this.StartDate = startDate
	this.RemoteCreatedAt = remoteCreatedAt
	this.TerminationDate = terminationDate
	this.Avatar = avatar
	this.HomeLocation = homeLocation
	this.WorkLocation = workLocation
	this.Manager = manager
	this.BankAccount = bankAccount
	this.Employments = employments
	this.CustomFields = customFields
	this.Groups = groups
	this.Company = company
	return &this
}

// NewEmployee20230301ResponseWithDefaults instantiates a new Employee20230301Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployee20230301ResponseWithDefaults() *Employee20230301Response {
	this := Employee20230301Response{}
	return &this
}

// GetId returns the Id field value
func (o *Employee20230301Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Employee20230301Response) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Employee20230301Response) SetId(v string) {
	o.Id = v
}

// GetRemoteId returns the RemoteId field value
func (o *Employee20230301Response) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *Employee20230301Response) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *Employee20230301Response) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetEmployeeNumber returns the EmployeeNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmployeeNumber.Get()
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetEmployeeNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmployeeNumber.Get(), o.EmployeeNumber.IsSet()
}

// SetEmployeeNumber sets field value
func (o *Employee20230301Response) SetEmployeeNumber(v string) {
	o.EmployeeNumber.Set(&v)
}

// GetFirstName returns the FirstName field value
func (o *Employee20230301Response) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *Employee20230301Response) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *Employee20230301Response) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *Employee20230301Response) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *Employee20230301Response) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *Employee20230301Response) SetLastName(v string) {
	o.LastName = v
}

// GetDisplayFullName returns the DisplayFullName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetDisplayFullName() string {
	if o == nil || o.DisplayFullName.Get() == nil {
		var ret string
		return ret
	}

	return *o.DisplayFullName.Get()
}

// GetDisplayFullNameOk returns a tuple with the DisplayFullName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetDisplayFullNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayFullName.Get(), o.DisplayFullName.IsSet()
}

// SetDisplayFullName sets field value
func (o *Employee20230301Response) SetDisplayFullName(v string) {
	o.DisplayFullName.Set(&v)
}

// GetNationality returns the Nationality field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetNationality() string {
	if o == nil || o.Nationality.Get() == nil {
		var ret string
		return ret
	}

	return *o.Nationality.Get()
}

// GetNationalityOk returns a tuple with the Nationality field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetNationalityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Nationality.Get(), o.Nationality.IsSet()
}

// SetNationality sets field value
func (o *Employee20230301Response) SetNationality(v string) {
	o.Nationality.Set(&v)
}

// GetJobTitle returns the JobTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetJobTitle() string {
	if o == nil || o.JobTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.JobTitle.Get()
}

// GetJobTitleOk returns a tuple with the JobTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetJobTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.JobTitle.Get(), o.JobTitle.IsSet()
}

// SetJobTitle sets field value
func (o *Employee20230301Response) SetJobTitle(v string) {
	o.JobTitle.Set(&v)
}

// GetWorkEmail returns the WorkEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetWorkEmail() string {
	if o == nil || o.WorkEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.WorkEmail.Get()
}

// GetWorkEmailOk returns a tuple with the WorkEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetWorkEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkEmail.Get(), o.WorkEmail.IsSet()
}

// SetWorkEmail sets field value
func (o *Employee20230301Response) SetWorkEmail(v string) {
	o.WorkEmail.Set(&v)
}

// GetPersonalEmail returns the PersonalEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetPersonalEmail() string {
	if o == nil || o.PersonalEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.PersonalEmail.Get()
}

// GetPersonalEmailOk returns a tuple with the PersonalEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetPersonalEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PersonalEmail.Get(), o.PersonalEmail.IsSet()
}

// SetPersonalEmail sets field value
func (o *Employee20230301Response) SetPersonalEmail(v string) {
	o.PersonalEmail.Set(&v)
}

// GetMobilePhoneNumber returns the MobilePhoneNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetMobilePhoneNumber() string {
	if o == nil || o.MobilePhoneNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.MobilePhoneNumber.Get()
}

// GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetMobilePhoneNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MobilePhoneNumber.Get(), o.MobilePhoneNumber.IsSet()
}

// SetMobilePhoneNumber sets field value
func (o *Employee20230301Response) SetMobilePhoneNumber(v string) {
	o.MobilePhoneNumber.Set(&v)
}

// GetTaxId returns the TaxId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetTaxId() string {
	if o == nil || o.TaxId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TaxId.Get()
}

// GetTaxIdOk returns a tuple with the TaxId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetTaxIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxId.Get(), o.TaxId.IsSet()
}

// SetTaxId sets field value
func (o *Employee20230301Response) SetTaxId(v string) {
	o.TaxId.Set(&v)
}

// GetGender returns the Gender field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetGender() string {
	if o == nil || o.Gender.Get() == nil {
		var ret string
		return ret
	}

	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetGenderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// SetGender sets field value
func (o *Employee20230301Response) SetGender(v string) {
	o.Gender.Set(&v)
}

// GetEthnicity returns the Ethnicity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetEthnicity() string {
	if o == nil || o.Ethnicity.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ethnicity.Get()
}

// GetEthnicityOk returns a tuple with the Ethnicity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetEthnicityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ethnicity.Get(), o.Ethnicity.IsSet()
}

// SetEthnicity sets field value
func (o *Employee20230301Response) SetEthnicity(v string) {
	o.Ethnicity.Set(&v)
}

// GetMaritalStatus returns the MaritalStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetMaritalStatus() string {
	if o == nil || o.MaritalStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.MaritalStatus.Get()
}

// GetMaritalStatusOk returns a tuple with the MaritalStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetMaritalStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MaritalStatus.Get(), o.MaritalStatus.IsSet()
}

// SetMaritalStatus sets field value
func (o *Employee20230301Response) SetMaritalStatus(v string) {
	o.MaritalStatus.Set(&v)
}

// GetDateOfBirth returns the DateOfBirth field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateOfBirth.Get()
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetDateOfBirthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateOfBirth.Get(), o.DateOfBirth.IsSet()
}

// SetDateOfBirth sets field value
func (o *Employee20230301Response) SetDateOfBirth(v string) {
	o.DateOfBirth.Set(&v)
}

// GetEmploymentStatus returns the EmploymentStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetEmploymentStatus() string {
	if o == nil || o.EmploymentStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmploymentStatus.Get()
}

// GetEmploymentStatusOk returns a tuple with the EmploymentStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetEmploymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentStatus.Get(), o.EmploymentStatus.IsSet()
}

// SetEmploymentStatus sets field value
func (o *Employee20230301Response) SetEmploymentStatus(v string) {
	o.EmploymentStatus.Set(&v)
}

// GetEmploymentType returns the EmploymentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetEmploymentType() string {
	if o == nil || o.EmploymentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.EmploymentType.Get()
}

// GetEmploymentTypeOk returns a tuple with the EmploymentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetEmploymentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmploymentType.Get(), o.EmploymentType.IsSet()
}

// SetEmploymentType sets field value
func (o *Employee20230301Response) SetEmploymentType(v string) {
	o.EmploymentType.Set(&v)
}

// GetStartDate returns the StartDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetStartDate() string {
	if o == nil || o.StartDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// SetStartDate sets field value
func (o *Employee20230301Response) SetStartDate(v string) {
	o.StartDate.Set(&v)
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetRemoteCreatedAt() string {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetRemoteCreatedAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// SetRemoteCreatedAt sets field value
func (o *Employee20230301Response) SetRemoteCreatedAt(v string) {
	o.RemoteCreatedAt.Set(&v)
}

// GetTerminationDate returns the TerminationDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetTerminationDate() string {
	if o == nil || o.TerminationDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.TerminationDate.Get()
}

// GetTerminationDateOk returns a tuple with the TerminationDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetTerminationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TerminationDate.Get(), o.TerminationDate.IsSet()
}

// SetTerminationDate sets field value
func (o *Employee20230301Response) SetTerminationDate(v string) {
	o.TerminationDate.Set(&v)
}

// GetAvatar returns the Avatar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Employee20230301Response) GetAvatar() string {
	if o == nil || o.Avatar.Get() == nil {
		var ret string
		return ret
	}

	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetAvatarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// SetAvatar sets field value
func (o *Employee20230301Response) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// GetHomeLocation returns the HomeLocation field value
func (o *Employee20230301Response) GetHomeLocation() Employee20230301ResponseHomeLocation {
	if o == nil {
		var ret Employee20230301ResponseHomeLocation
		return ret
	}

	return o.HomeLocation
}

// GetHomeLocationOk returns a tuple with the HomeLocation field value
// and a boolean to check if the value has been set.
func (o *Employee20230301Response) GetHomeLocationOk() (*Employee20230301ResponseHomeLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HomeLocation, true
}

// SetHomeLocation sets field value
func (o *Employee20230301Response) SetHomeLocation(v Employee20230301ResponseHomeLocation) {
	o.HomeLocation = v
}

// GetWorkLocation returns the WorkLocation field value
// If the value is explicit nil, the zero value for Employee20230301ResponseWorkLocation will be returned
func (o *Employee20230301Response) GetWorkLocation() Employee20230301ResponseWorkLocation {
	if o == nil || o.WorkLocation.Get() == nil {
		var ret Employee20230301ResponseWorkLocation
		return ret
	}

	return *o.WorkLocation.Get()
}

// GetWorkLocationOk returns a tuple with the WorkLocation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetWorkLocationOk() (*Employee20230301ResponseWorkLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WorkLocation.Get(), o.WorkLocation.IsSet()
}

// SetWorkLocation sets field value
func (o *Employee20230301Response) SetWorkLocation(v Employee20230301ResponseWorkLocation) {
	o.WorkLocation.Set(&v)
}

// GetManager returns the Manager field value
// If the value is explicit nil, the zero value for Employee20230301ResponseManager will be returned
func (o *Employee20230301Response) GetManager() Employee20230301ResponseManager {
	if o == nil || o.Manager.Get() == nil {
		var ret Employee20230301ResponseManager
		return ret
	}

	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetManagerOk() (*Employee20230301ResponseManager, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// SetManager sets field value
func (o *Employee20230301Response) SetManager(v Employee20230301ResponseManager) {
	o.Manager.Set(&v)
}

// GetBankAccount returns the BankAccount field value
// If the value is explicit nil, the zero value for Employee20230301ResponseBankAccount will be returned
func (o *Employee20230301Response) GetBankAccount() Employee20230301ResponseBankAccount {
	if o == nil || o.BankAccount.Get() == nil {
		var ret Employee20230301ResponseBankAccount
		return ret
	}

	return *o.BankAccount.Get()
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetBankAccountOk() (*Employee20230301ResponseBankAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BankAccount.Get(), o.BankAccount.IsSet()
}

// SetBankAccount sets field value
func (o *Employee20230301Response) SetBankAccount(v Employee20230301ResponseBankAccount) {
	o.BankAccount.Set(&v)
}

// GetEmployments returns the Employments field value
// If the value is explicit nil, the zero value for []EmploymentResponse will be returned
func (o *Employee20230301Response) GetEmployments() []EmploymentResponse {
	if o == nil {
		var ret []EmploymentResponse
		return ret
	}

	return o.Employments
}

// GetEmploymentsOk returns a tuple with the Employments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetEmploymentsOk() (*[]EmploymentResponse, bool) {
	if o == nil || o.Employments == nil {
		return nil, false
	}
	return &o.Employments, true
}

// SetEmployments sets field value
func (o *Employee20230301Response) SetEmployments(v []EmploymentResponse) {
	o.Employments = v
}

// GetCustomFields returns the CustomFields field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *Employee20230301Response) GetCustomFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetCustomFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return &o.CustomFields, true
}

// SetCustomFields sets field value
func (o *Employee20230301Response) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetGroups returns the Groups field value
// If the value is explicit nil, the zero value for []Employee20230301ResponseGroups will be returned
func (o *Employee20230301Response) GetGroups() []Employee20230301ResponseGroups {
	if o == nil {
		var ret []Employee20230301ResponseGroups
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetGroupsOk() (*[]Employee20230301ResponseGroups, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *Employee20230301Response) SetGroups(v []Employee20230301ResponseGroups) {
	o.Groups = v
}

// GetCompany returns the Company field value
// If the value is explicit nil, the zero value for Employee20230301ResponseCompany will be returned
func (o *Employee20230301Response) GetCompany() Employee20230301ResponseCompany {
	if o == nil || o.Company.Get() == nil {
		var ret Employee20230301ResponseCompany
		return ret
	}

	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Employee20230301Response) GetCompanyOk() (*Employee20230301ResponseCompany, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// SetCompany sets field value
func (o *Employee20230301Response) SetCompany(v Employee20230301ResponseCompany) {
	o.Company.Set(&v)
}

func (o Employee20230301Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["remote_id"] = o.RemoteId
	}
	if true {
		toSerialize["employee_number"] = o.EmployeeNumber.Get()
	}
	if true {
		toSerialize["first_name"] = o.FirstName
	}
	if true {
		toSerialize["last_name"] = o.LastName
	}
	if true {
		toSerialize["display_full_name"] = o.DisplayFullName.Get()
	}
	if true {
		toSerialize["nationality"] = o.Nationality.Get()
	}
	if true {
		toSerialize["job_title"] = o.JobTitle.Get()
	}
	if true {
		toSerialize["work_email"] = o.WorkEmail.Get()
	}
	if true {
		toSerialize["personal_email"] = o.PersonalEmail.Get()
	}
	if true {
		toSerialize["mobile_phone_number"] = o.MobilePhoneNumber.Get()
	}
	if true {
		toSerialize["tax_id"] = o.TaxId.Get()
	}
	if true {
		toSerialize["gender"] = o.Gender.Get()
	}
	if true {
		toSerialize["ethnicity"] = o.Ethnicity.Get()
	}
	if true {
		toSerialize["marital_status"] = o.MaritalStatus.Get()
	}
	if true {
		toSerialize["date_of_birth"] = o.DateOfBirth.Get()
	}
	if true {
		toSerialize["employment_status"] = o.EmploymentStatus.Get()
	}
	if true {
		toSerialize["employment_type"] = o.EmploymentType.Get()
	}
	if true {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if true {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if true {
		toSerialize["termination_date"] = o.TerminationDate.Get()
	}
	if true {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if true {
		toSerialize["home_location"] = o.HomeLocation
	}
	if true {
		toSerialize["work_location"] = o.WorkLocation.Get()
	}
	if true {
		toSerialize["manager"] = o.Manager.Get()
	}
	if true {
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
	if true {
		toSerialize["company"] = o.Company.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEmployee20230301Response struct {
	value *Employee20230301Response
	isSet bool
}

func (v NullableEmployee20230301Response) Get() *Employee20230301Response {
	return v.value
}

func (v *NullableEmployee20230301Response) Set(val *Employee20230301Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee20230301Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee20230301Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee20230301Response(val *Employee20230301Response) *NullableEmployee20230301Response {
	return &NullableEmployee20230301Response{value: val, isSet: true}
}

func (v NullableEmployee20230301Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee20230301Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


