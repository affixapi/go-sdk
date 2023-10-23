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

// Order20230320Response struct for Order20230320Response
type Order20230320Response struct {
	// The order identifier, if any, and from the provider, related to the transaction
	Id NullableString `json:"id"`
	// The date the transaction was placed
	Date string `json:"date"`
	// The name of who paid for the order
	BillingTo NullableString `json:"billing_to"`
	// The street address of who paid for the order
	BillingStreetAddress NullableString `json:"billing_street_address"`
	// The town/city of who paid for the order
	BillingLocality NullableString `json:"billing_locality"`
	// The administrative area of who paid for the order. If US, the two-digit state abbreviation 
	BillingAdministrativeArea NullableString `json:"billing_administrative_area"`
	// The post code of who paid for the order. If US, The 5+4 postal code 
	BillingPostCode NullableString `json:"billing_post_code"`
	// The country of who paid for the order
	BillingCountry NullableString `json:"billing_country"`
	// The redacted information of the payment instrument used
	PaymentInstrument string `json:"payment_instrument"`
	Currency CurrencyResponse `json:"currency"`
	// The subtotal of the order; if USD, in cents, ie $9.88 => 988. if GBP, in pence, ie £513.87 => 51397 if EUR, in cent, ie €15,90 => 1590, ie €15.90 => 1590, ie 15,90€ => 1590 
	Subtotal int32 `json:"subtotal"`
	// The shipping and handling fees of the order; if USD, in cents, ie $9.88 => 988. if GBP, in pence, ie £513.87 => 51397 if EUR, in cent, ie €15,90 => 1590, ie €15.90 => 1590, ie 15,90€ => 1590 
	ShippingFee NullableInt32 `json:"shipping_fee"`
	// Promotions, if any
	Discounts NullableInt32 `json:"discounts"`
	// Other fees, for example, California Bottle Deposit Fee; in cents/pence/cent.
	OtherFees NullableInt32 `json:"other_fees"`
	// The sales tax or VAT of the order; in cents/pence/cent.
	Tax NullableInt32 `json:"tax"`
	// If a gift card was used, how much of the order was paid for by gift card 
	GiftCardApplied NullableInt32 `json:"gift_card_applied"`
	// The total price of the order; if USD, in cents. ie $9.88 => 988. If order is not complete, this field may be null 
	Total NullableInt32 `json:"total"`
	Shipments []map[string]interface{} `json:"shipments"`
}

// NewOrder20230320Response instantiates a new Order20230320Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder20230320Response(id NullableString, date string, billingTo NullableString, billingStreetAddress NullableString, billingLocality NullableString, billingAdministrativeArea NullableString, billingPostCode NullableString, billingCountry NullableString, paymentInstrument string, currency CurrencyResponse, subtotal int32, shippingFee NullableInt32, discounts NullableInt32, otherFees NullableInt32, tax NullableInt32, giftCardApplied NullableInt32, total NullableInt32, shipments []map[string]interface{}) *Order20230320Response {
	this := Order20230320Response{}
	this.Id = id
	this.Date = date
	this.BillingTo = billingTo
	this.BillingStreetAddress = billingStreetAddress
	this.BillingLocality = billingLocality
	this.BillingAdministrativeArea = billingAdministrativeArea
	this.BillingPostCode = billingPostCode
	this.BillingCountry = billingCountry
	this.PaymentInstrument = paymentInstrument
	this.Currency = currency
	this.Subtotal = subtotal
	this.ShippingFee = shippingFee
	this.Discounts = discounts
	this.OtherFees = otherFees
	this.Tax = tax
	this.GiftCardApplied = giftCardApplied
	this.Total = total
	this.Shipments = shipments
	return &this
}

// NewOrder20230320ResponseWithDefaults instantiates a new Order20230320Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrder20230320ResponseWithDefaults() *Order20230320Response {
	this := Order20230320Response{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *Order20230320Response) SetId(v string) {
	o.Id.Set(&v)
}

// GetDate returns the Date field value
func (o *Order20230320Response) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Order20230320Response) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Order20230320Response) SetDate(v string) {
	o.Date = v
}

// GetBillingTo returns the BillingTo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingTo() string {
	if o == nil || o.BillingTo.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingTo.Get()
}

// GetBillingToOk returns a tuple with the BillingTo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingTo.Get(), o.BillingTo.IsSet()
}

// SetBillingTo sets field value
func (o *Order20230320Response) SetBillingTo(v string) {
	o.BillingTo.Set(&v)
}

// GetBillingStreetAddress returns the BillingStreetAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingStreetAddress() string {
	if o == nil || o.BillingStreetAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingStreetAddress.Get()
}

// GetBillingStreetAddressOk returns a tuple with the BillingStreetAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingStreetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingStreetAddress.Get(), o.BillingStreetAddress.IsSet()
}

// SetBillingStreetAddress sets field value
func (o *Order20230320Response) SetBillingStreetAddress(v string) {
	o.BillingStreetAddress.Set(&v)
}

// GetBillingLocality returns the BillingLocality field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingLocality() string {
	if o == nil || o.BillingLocality.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingLocality.Get()
}

// GetBillingLocalityOk returns a tuple with the BillingLocality field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingLocalityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingLocality.Get(), o.BillingLocality.IsSet()
}

// SetBillingLocality sets field value
func (o *Order20230320Response) SetBillingLocality(v string) {
	o.BillingLocality.Set(&v)
}

// GetBillingAdministrativeArea returns the BillingAdministrativeArea field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingAdministrativeArea() string {
	if o == nil || o.BillingAdministrativeArea.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingAdministrativeArea.Get()
}

// GetBillingAdministrativeAreaOk returns a tuple with the BillingAdministrativeArea field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingAdministrativeAreaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingAdministrativeArea.Get(), o.BillingAdministrativeArea.IsSet()
}

// SetBillingAdministrativeArea sets field value
func (o *Order20230320Response) SetBillingAdministrativeArea(v string) {
	o.BillingAdministrativeArea.Set(&v)
}

// GetBillingPostCode returns the BillingPostCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingPostCode() string {
	if o == nil || o.BillingPostCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingPostCode.Get()
}

// GetBillingPostCodeOk returns a tuple with the BillingPostCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingPostCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingPostCode.Get(), o.BillingPostCode.IsSet()
}

// SetBillingPostCode sets field value
func (o *Order20230320Response) SetBillingPostCode(v string) {
	o.BillingPostCode.Set(&v)
}

// GetBillingCountry returns the BillingCountry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Order20230320Response) GetBillingCountry() string {
	if o == nil || o.BillingCountry.Get() == nil {
		var ret string
		return ret
	}

	return *o.BillingCountry.Get()
}

// GetBillingCountryOk returns a tuple with the BillingCountry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetBillingCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BillingCountry.Get(), o.BillingCountry.IsSet()
}

// SetBillingCountry sets field value
func (o *Order20230320Response) SetBillingCountry(v string) {
	o.BillingCountry.Set(&v)
}

// GetPaymentInstrument returns the PaymentInstrument field value
func (o *Order20230320Response) GetPaymentInstrument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentInstrument
}

// GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field value
// and a boolean to check if the value has been set.
func (o *Order20230320Response) GetPaymentInstrumentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentInstrument, true
}

// SetPaymentInstrument sets field value
func (o *Order20230320Response) SetPaymentInstrument(v string) {
	o.PaymentInstrument = v
}

// GetCurrency returns the Currency field value
func (o *Order20230320Response) GetCurrency() CurrencyResponse {
	if o == nil {
		var ret CurrencyResponse
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Order20230320Response) GetCurrencyOk() (*CurrencyResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Order20230320Response) SetCurrency(v CurrencyResponse) {
	o.Currency = v
}

// GetSubtotal returns the Subtotal field value
func (o *Order20230320Response) GetSubtotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *Order20230320Response) GetSubtotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *Order20230320Response) SetSubtotal(v int32) {
	o.Subtotal = v
}

// GetShippingFee returns the ShippingFee field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetShippingFee() int32 {
	if o == nil || o.ShippingFee.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ShippingFee.Get()
}

// GetShippingFeeOk returns a tuple with the ShippingFee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetShippingFeeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShippingFee.Get(), o.ShippingFee.IsSet()
}

// SetShippingFee sets field value
func (o *Order20230320Response) SetShippingFee(v int32) {
	o.ShippingFee.Set(&v)
}

// GetDiscounts returns the Discounts field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetDiscounts() int32 {
	if o == nil || o.Discounts.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Discounts.Get()
}

// GetDiscountsOk returns a tuple with the Discounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetDiscountsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Discounts.Get(), o.Discounts.IsSet()
}

// SetDiscounts sets field value
func (o *Order20230320Response) SetDiscounts(v int32) {
	o.Discounts.Set(&v)
}

// GetOtherFees returns the OtherFees field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetOtherFees() int32 {
	if o == nil || o.OtherFees.Get() == nil {
		var ret int32
		return ret
	}

	return *o.OtherFees.Get()
}

// GetOtherFeesOk returns a tuple with the OtherFees field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetOtherFeesOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OtherFees.Get(), o.OtherFees.IsSet()
}

// SetOtherFees sets field value
func (o *Order20230320Response) SetOtherFees(v int32) {
	o.OtherFees.Set(&v)
}

// GetTax returns the Tax field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetTax() int32 {
	if o == nil || o.Tax.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Tax.Get()
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetTaxOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Tax.Get(), o.Tax.IsSet()
}

// SetTax sets field value
func (o *Order20230320Response) SetTax(v int32) {
	o.Tax.Set(&v)
}

// GetGiftCardApplied returns the GiftCardApplied field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetGiftCardApplied() int32 {
	if o == nil || o.GiftCardApplied.Get() == nil {
		var ret int32
		return ret
	}

	return *o.GiftCardApplied.Get()
}

// GetGiftCardAppliedOk returns a tuple with the GiftCardApplied field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetGiftCardAppliedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GiftCardApplied.Get(), o.GiftCardApplied.IsSet()
}

// SetGiftCardApplied sets field value
func (o *Order20230320Response) SetGiftCardApplied(v int32) {
	o.GiftCardApplied.Set(&v)
}

// GetTotal returns the Total field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Order20230320Response) GetTotal() int32 {
	if o == nil || o.Total.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// SetTotal sets field value
func (o *Order20230320Response) SetTotal(v int32) {
	o.Total.Set(&v)
}

// GetShipments returns the Shipments field value
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned
func (o *Order20230320Response) GetShipments() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order20230320Response) GetShipmentsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Shipments == nil {
		return nil, false
	}
	return &o.Shipments, true
}

// SetShipments sets field value
func (o *Order20230320Response) SetShipments(v []map[string]interface{}) {
	o.Shipments = v
}

func (o Order20230320Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id.Get()
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["billing_to"] = o.BillingTo.Get()
	}
	if true {
		toSerialize["billing_street_address"] = o.BillingStreetAddress.Get()
	}
	if true {
		toSerialize["billing_locality"] = o.BillingLocality.Get()
	}
	if true {
		toSerialize["billing_administrative_area"] = o.BillingAdministrativeArea.Get()
	}
	if true {
		toSerialize["billing_post_code"] = o.BillingPostCode.Get()
	}
	if true {
		toSerialize["billing_country"] = o.BillingCountry.Get()
	}
	if true {
		toSerialize["payment_instrument"] = o.PaymentInstrument
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["subtotal"] = o.Subtotal
	}
	if true {
		toSerialize["shipping_fee"] = o.ShippingFee.Get()
	}
	if true {
		toSerialize["discounts"] = o.Discounts.Get()
	}
	if true {
		toSerialize["other_fees"] = o.OtherFees.Get()
	}
	if true {
		toSerialize["tax"] = o.Tax.Get()
	}
	if true {
		toSerialize["gift_card_applied"] = o.GiftCardApplied.Get()
	}
	if true {
		toSerialize["total"] = o.Total.Get()
	}
	if o.Shipments != nil {
		toSerialize["shipments"] = o.Shipments
	}
	return json.Marshal(toSerialize)
}

type NullableOrder20230320Response struct {
	value *Order20230320Response
	isSet bool
}

func (v NullableOrder20230320Response) Get() *Order20230320Response {
	return v.value
}

func (v *NullableOrder20230320Response) Set(val *Order20230320Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder20230320Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder20230320Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder20230320Response(val *Order20230320Response) *NullableOrder20230320Response {
	return &NullableOrder20230320Response{value: val, isSet: true}
}

func (v NullableOrder20230320Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder20230320Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


