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

// PayslipResponse struct for PayslipResponse
type PayslipResponse struct {
	// The Affix-assigned id of the payslip
	Id string `json:"id"`
	// the remote system-assigned id of the payrun
	RemoteId string `json:"remote_id"`
	EmployeeId string `json:"employee_id"`
	PayrunId string `json:"payrun_id"`
	Currency string `json:"currency"`
	// if USD/EUR/GBP, in cent
	GrossPay float32 `json:"gross_pay"`
	// if USD/EUR/GBP, in cent
	NetPay float32 `json:"net_pay"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	CheckDate string `json:"check_date"`
	Earnings []PayslipResponseEarnings `json:"earnings"`
	Deductions []PayslipResponseDeductions `json:"deductions"`
	Taxes []PayslipResponseTaxes `json:"taxes"`
	FieldMappings map[string]interface{} `json:"field_mappings"`
}

// NewPayslipResponse instantiates a new PayslipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayslipResponse(id string, remoteId string, employeeId string, payrunId string, currency string, grossPay float32, netPay float32, startDate string, endDate string, checkDate string, earnings []PayslipResponseEarnings, deductions []PayslipResponseDeductions, taxes []PayslipResponseTaxes, fieldMappings map[string]interface{}) *PayslipResponse {
	this := PayslipResponse{}
	this.Id = id
	this.RemoteId = remoteId
	this.EmployeeId = employeeId
	this.PayrunId = payrunId
	this.Currency = currency
	this.GrossPay = grossPay
	this.NetPay = netPay
	this.StartDate = startDate
	this.EndDate = endDate
	this.CheckDate = checkDate
	this.Earnings = earnings
	this.Deductions = deductions
	this.Taxes = taxes
	this.FieldMappings = fieldMappings
	return &this
}

// NewPayslipResponseWithDefaults instantiates a new PayslipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayslipResponseWithDefaults() *PayslipResponse {
	this := PayslipResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PayslipResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PayslipResponse) SetId(v string) {
	o.Id = v
}

// GetRemoteId returns the RemoteId field value
func (o *PayslipResponse) GetRemoteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemoteId, true
}

// SetRemoteId sets field value
func (o *PayslipResponse) SetRemoteId(v string) {
	o.RemoteId = v
}

// GetEmployeeId returns the EmployeeId field value
func (o *PayslipResponse) GetEmployeeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEmployeeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EmployeeId, true
}

// SetEmployeeId sets field value
func (o *PayslipResponse) SetEmployeeId(v string) {
	o.EmployeeId = v
}

// GetPayrunId returns the PayrunId field value
func (o *PayslipResponse) GetPayrunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayrunId
}

// GetPayrunIdOk returns a tuple with the PayrunId field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetPayrunIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PayrunId, true
}

// SetPayrunId sets field value
func (o *PayslipResponse) SetPayrunId(v string) {
	o.PayrunId = v
}

// GetCurrency returns the Currency field value
func (o *PayslipResponse) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PayslipResponse) SetCurrency(v string) {
	o.Currency = v
}

// GetGrossPay returns the GrossPay field value
func (o *PayslipResponse) GetGrossPay() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.GrossPay
}

// GetGrossPayOk returns a tuple with the GrossPay field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetGrossPayOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GrossPay, true
}

// SetGrossPay sets field value
func (o *PayslipResponse) SetGrossPay(v float32) {
	o.GrossPay = v
}

// GetNetPay returns the NetPay field value
func (o *PayslipResponse) GetNetPay() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NetPay
}

// GetNetPayOk returns a tuple with the NetPay field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetNetPayOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetPay, true
}

// SetNetPay sets field value
func (o *PayslipResponse) SetNetPay(v float32) {
	o.NetPay = v
}

// GetStartDate returns the StartDate field value
func (o *PayslipResponse) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *PayslipResponse) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *PayslipResponse) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *PayslipResponse) SetEndDate(v string) {
	o.EndDate = v
}

// GetCheckDate returns the CheckDate field value
func (o *PayslipResponse) GetCheckDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckDate
}

// GetCheckDateOk returns a tuple with the CheckDate field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetCheckDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CheckDate, true
}

// SetCheckDate sets field value
func (o *PayslipResponse) SetCheckDate(v string) {
	o.CheckDate = v
}

// GetEarnings returns the Earnings field value
func (o *PayslipResponse) GetEarnings() []PayslipResponseEarnings {
	if o == nil {
		var ret []PayslipResponseEarnings
		return ret
	}

	return o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetEarningsOk() (*[]PayslipResponseEarnings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Earnings, true
}

// SetEarnings sets field value
func (o *PayslipResponse) SetEarnings(v []PayslipResponseEarnings) {
	o.Earnings = v
}

// GetDeductions returns the Deductions field value
func (o *PayslipResponse) GetDeductions() []PayslipResponseDeductions {
	if o == nil {
		var ret []PayslipResponseDeductions
		return ret
	}

	return o.Deductions
}

// GetDeductionsOk returns a tuple with the Deductions field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetDeductionsOk() (*[]PayslipResponseDeductions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Deductions, true
}

// SetDeductions sets field value
func (o *PayslipResponse) SetDeductions(v []PayslipResponseDeductions) {
	o.Deductions = v
}

// GetTaxes returns the Taxes field value
func (o *PayslipResponse) GetTaxes() []PayslipResponseTaxes {
	if o == nil {
		var ret []PayslipResponseTaxes
		return ret
	}

	return o.Taxes
}

// GetTaxesOk returns a tuple with the Taxes field value
// and a boolean to check if the value has been set.
func (o *PayslipResponse) GetTaxesOk() (*[]PayslipResponseTaxes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Taxes, true
}

// SetTaxes sets field value
func (o *PayslipResponse) SetTaxes(v []PayslipResponseTaxes) {
	o.Taxes = v
}

// GetFieldMappings returns the FieldMappings field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *PayslipResponse) GetFieldMappings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PayslipResponse) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// SetFieldMappings sets field value
func (o *PayslipResponse) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

func (o PayslipResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["payrun_id"] = o.PayrunId
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["gross_pay"] = o.GrossPay
	}
	if true {
		toSerialize["net_pay"] = o.NetPay
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if true {
		toSerialize["check_date"] = o.CheckDate
	}
	if true {
		toSerialize["earnings"] = o.Earnings
	}
	if true {
		toSerialize["deductions"] = o.Deductions
	}
	if true {
		toSerialize["taxes"] = o.Taxes
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	return json.Marshal(toSerialize)
}

type NullablePayslipResponse struct {
	value *PayslipResponse
	isSet bool
}

func (v NullablePayslipResponse) Get() *PayslipResponse {
	return v.value
}

func (v *NullablePayslipResponse) Set(val *PayslipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePayslipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePayslipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayslipResponse(val *PayslipResponse) *NullablePayslipResponse {
	return &NullablePayslipResponse{value: val, isSet: true}
}

func (v NullablePayslipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayslipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


