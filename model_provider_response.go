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
	"fmt"
)

// ProviderResponse the model 'ProviderResponse'
type ProviderResponse string

// List of ProviderResponse
const (
	SANDBOX ProviderResponse = "sandbox"
	TAXACT ProviderResponse = "taxact"
	TURBOTAX ProviderResponse = "turbotax"
	AMAZON ProviderResponse = "amazon"
	AMAZON_COM_BUSINESS ProviderResponse = "amazon.com business"
	AMAZON_CO_UK ProviderResponse = "amazon.co.uk"
	AMAZON_DE ProviderResponse = "amazon.de"
	BAMBOOHR ProviderResponse = "bamboohr"
	BREATHE ProviderResponse = "breathe"
	CEZANNE ProviderResponse = "cezanne"
	CHARLIEHR ProviderResponse = "charliehr"
	DEEL ProviderResponse = "deel"
	DEPUTY ProviderResponse = "deputy"
	HIBOB ProviderResponse = "hibob"
	HUMAANS ProviderResponse = "humaans"
	IRIS_CASCADE ProviderResponse = "iris cascade"
	PERSONIO_DE ProviderResponse = "personio.de"
	SAGEHR ProviderResponse = "sagehr"
	SAPLINGHR ProviderResponse = "saplinghr"
	STAFFOLOGY ProviderResponse = "staffology"
	UKG_PRO ProviderResponse = "ukg pro"
	XERO ProviderResponse = "xero"
	BRAIN_PAYROLL ProviderResponse = "brain payroll"
	BRIGHTPAY_CONNECT ProviderResponse = "brightpay connect"
	EMPLOYMENT_HERO ProviderResponse = "employment hero"
	FACTORIALHR ProviderResponse = "factorialhr"
	FOURTHHR ProviderResponse = "fourthhr"
	ITRENT ProviderResponse = "itrent"
	OYSTERHR ProviderResponse = "oysterhr"
	PAROLLA_IE ProviderResponse = "parolla.ie"
	PAYCIRCLE ProviderResponse = "paycircle"
	PEOPLEHR ProviderResponse = "peoplehr"
	REMOTE_COM ProviderResponse = "remote.com"
	SAGEONE ProviderResponse = "sageone"
	SIMPLEPAY_IE ProviderResponse = "simplepay.ie"
	PENTO_IO ProviderResponse = "pento.io"
)

var allowedProviderResponseEnumValues = []ProviderResponse{
	"sandbox",
	"taxact",
	"turbotax",
	"amazon",
	"amazon.com business",
	"amazon.co.uk",
	"amazon.de",
	"bamboohr",
	"breathe",
	"cezanne",
	"charliehr",
	"deel",
	"deputy",
	"hibob",
	"humaans",
	"iris cascade",
	"personio.de",
	"sagehr",
	"saplinghr",
	"staffology",
	"ukg pro",
	"xero",
	"brain payroll",
	"brightpay connect",
	"employment hero",
	"factorialhr",
	"fourthhr",
	"itrent",
	"oysterhr",
	"parolla.ie",
	"paycircle",
	"peoplehr",
	"remote.com",
	"sageone",
	"simplepay.ie",
	"pento.io",
}

func (v *ProviderResponse) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderResponse(value)
	for _, existing := range allowedProviderResponseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderResponse", value)
}

// NewProviderResponseFromValue returns a pointer to a valid ProviderResponse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderResponseFromValue(v string) (*ProviderResponse, error) {
	ev := ProviderResponse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderResponse: valid values are %v", v, allowedProviderResponseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderResponse) IsValid() bool {
	for _, existing := range allowedProviderResponseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderResponse value
func (v ProviderResponse) Ptr() *ProviderResponse {
	return &v
}

type NullableProviderResponse struct {
	value *ProviderResponse
	isSet bool
}

func (v NullableProviderResponse) Get() *ProviderResponse {
	return v.value
}

func (v *NullableProviderResponse) Set(val *ProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderResponse(val *ProviderResponse) *NullableProviderResponse {
	return &NullableProviderResponse{value: val, isSet: true}
}

func (v NullableProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

