# Order20230320Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The order identifier, if any, and from the provider, related to the transaction | 
**Date** | **string** | The date the transaction was placed | 
**BillingTo** | **NullableString** | The name of who paid for the order | 
**BillingStreetAddress** | **NullableString** | The street address of who paid for the order | 
**BillingLocality** | **NullableString** | The town/city of who paid for the order | 
**BillingAdministrativeArea** | **NullableString** | The administrative area of who paid for the order. If US, the two-digit state abbreviation  | 
**BillingPostCode** | **NullableString** | The post code of who paid for the order. If US, The 5+4 postal code  | 
**BillingCountry** | **NullableString** | The country of who paid for the order | 
**PaymentInstrument** | **string** | The redacted information of the payment instrument used | 
**Currency** | [**CurrencyResponse**](CurrencyResponse.md) |  | 
**Subtotal** | **int32** | The subtotal of the order; if USD, in cents, ie $9.88 &#x3D;&gt; 988. if GBP, in pence, ie £513.87 &#x3D;&gt; 51397 if EUR, in cent, ie €15,90 &#x3D;&gt; 1590, ie €15.90 &#x3D;&gt; 1590, ie 15,90€ &#x3D;&gt; 1590  | 
**ShippingFee** | **NullableInt32** | The shipping and handling fees of the order; if USD, in cents, ie $9.88 &#x3D;&gt; 988. if GBP, in pence, ie £513.87 &#x3D;&gt; 51397 if EUR, in cent, ie €15,90 &#x3D;&gt; 1590, ie €15.90 &#x3D;&gt; 1590, ie 15,90€ &#x3D;&gt; 1590  | 
**Discounts** | **NullableInt32** | Promotions, if any | 
**OtherFees** | **NullableInt32** | Other fees, for example, California Bottle Deposit Fee; in cents/pence/cent. | 
**Tax** | **NullableInt32** | The sales tax or VAT of the order; in cents/pence/cent. | 
**GiftCardApplied** | **NullableInt32** | If a gift card was used, how much of the order was paid for by gift card  | 
**Total** | **NullableInt32** | The total price of the order; if USD, in cents. ie $9.88 &#x3D;&gt; 988. If order is not complete, this field may be null  | 
**Shipments** | **[]map[string]interface{}** |  | 

## Methods

### NewOrder20230320Response

`func NewOrder20230320Response(id NullableString, date string, billingTo NullableString, billingStreetAddress NullableString, billingLocality NullableString, billingAdministrativeArea NullableString, billingPostCode NullableString, billingCountry NullableString, paymentInstrument string, currency CurrencyResponse, subtotal int32, shippingFee NullableInt32, discounts NullableInt32, otherFees NullableInt32, tax NullableInt32, giftCardApplied NullableInt32, total NullableInt32, shipments []map[string]interface{}, ) *Order20230320Response`

NewOrder20230320Response instantiates a new Order20230320Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrder20230320ResponseWithDefaults

`func NewOrder20230320ResponseWithDefaults() *Order20230320Response`

NewOrder20230320ResponseWithDefaults instantiates a new Order20230320Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order20230320Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order20230320Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order20230320Response) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Order20230320Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Order20230320Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDate

`func (o *Order20230320Response) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Order20230320Response) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Order20230320Response) SetDate(v string)`

SetDate sets Date field to given value.


### GetBillingTo

`func (o *Order20230320Response) GetBillingTo() string`

GetBillingTo returns the BillingTo field if non-nil, zero value otherwise.

### GetBillingToOk

`func (o *Order20230320Response) GetBillingToOk() (*string, bool)`

GetBillingToOk returns a tuple with the BillingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTo

`func (o *Order20230320Response) SetBillingTo(v string)`

SetBillingTo sets BillingTo field to given value.


### SetBillingToNil

`func (o *Order20230320Response) SetBillingToNil(b bool)`

 SetBillingToNil sets the value for BillingTo to be an explicit nil

### UnsetBillingTo
`func (o *Order20230320Response) UnsetBillingTo()`

UnsetBillingTo ensures that no value is present for BillingTo, not even an explicit nil
### GetBillingStreetAddress

`func (o *Order20230320Response) GetBillingStreetAddress() string`

GetBillingStreetAddress returns the BillingStreetAddress field if non-nil, zero value otherwise.

### GetBillingStreetAddressOk

`func (o *Order20230320Response) GetBillingStreetAddressOk() (*string, bool)`

GetBillingStreetAddressOk returns a tuple with the BillingStreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStreetAddress

`func (o *Order20230320Response) SetBillingStreetAddress(v string)`

SetBillingStreetAddress sets BillingStreetAddress field to given value.


### SetBillingStreetAddressNil

`func (o *Order20230320Response) SetBillingStreetAddressNil(b bool)`

 SetBillingStreetAddressNil sets the value for BillingStreetAddress to be an explicit nil

### UnsetBillingStreetAddress
`func (o *Order20230320Response) UnsetBillingStreetAddress()`

UnsetBillingStreetAddress ensures that no value is present for BillingStreetAddress, not even an explicit nil
### GetBillingLocality

`func (o *Order20230320Response) GetBillingLocality() string`

GetBillingLocality returns the BillingLocality field if non-nil, zero value otherwise.

### GetBillingLocalityOk

`func (o *Order20230320Response) GetBillingLocalityOk() (*string, bool)`

GetBillingLocalityOk returns a tuple with the BillingLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLocality

`func (o *Order20230320Response) SetBillingLocality(v string)`

SetBillingLocality sets BillingLocality field to given value.


### SetBillingLocalityNil

`func (o *Order20230320Response) SetBillingLocalityNil(b bool)`

 SetBillingLocalityNil sets the value for BillingLocality to be an explicit nil

### UnsetBillingLocality
`func (o *Order20230320Response) UnsetBillingLocality()`

UnsetBillingLocality ensures that no value is present for BillingLocality, not even an explicit nil
### GetBillingAdministrativeArea

`func (o *Order20230320Response) GetBillingAdministrativeArea() string`

GetBillingAdministrativeArea returns the BillingAdministrativeArea field if non-nil, zero value otherwise.

### GetBillingAdministrativeAreaOk

`func (o *Order20230320Response) GetBillingAdministrativeAreaOk() (*string, bool)`

GetBillingAdministrativeAreaOk returns a tuple with the BillingAdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAdministrativeArea

`func (o *Order20230320Response) SetBillingAdministrativeArea(v string)`

SetBillingAdministrativeArea sets BillingAdministrativeArea field to given value.


### SetBillingAdministrativeAreaNil

`func (o *Order20230320Response) SetBillingAdministrativeAreaNil(b bool)`

 SetBillingAdministrativeAreaNil sets the value for BillingAdministrativeArea to be an explicit nil

### UnsetBillingAdministrativeArea
`func (o *Order20230320Response) UnsetBillingAdministrativeArea()`

UnsetBillingAdministrativeArea ensures that no value is present for BillingAdministrativeArea, not even an explicit nil
### GetBillingPostCode

`func (o *Order20230320Response) GetBillingPostCode() string`

GetBillingPostCode returns the BillingPostCode field if non-nil, zero value otherwise.

### GetBillingPostCodeOk

`func (o *Order20230320Response) GetBillingPostCodeOk() (*string, bool)`

GetBillingPostCodeOk returns a tuple with the BillingPostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPostCode

`func (o *Order20230320Response) SetBillingPostCode(v string)`

SetBillingPostCode sets BillingPostCode field to given value.


### SetBillingPostCodeNil

`func (o *Order20230320Response) SetBillingPostCodeNil(b bool)`

 SetBillingPostCodeNil sets the value for BillingPostCode to be an explicit nil

### UnsetBillingPostCode
`func (o *Order20230320Response) UnsetBillingPostCode()`

UnsetBillingPostCode ensures that no value is present for BillingPostCode, not even an explicit nil
### GetBillingCountry

`func (o *Order20230320Response) GetBillingCountry() string`

GetBillingCountry returns the BillingCountry field if non-nil, zero value otherwise.

### GetBillingCountryOk

`func (o *Order20230320Response) GetBillingCountryOk() (*string, bool)`

GetBillingCountryOk returns a tuple with the BillingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCountry

`func (o *Order20230320Response) SetBillingCountry(v string)`

SetBillingCountry sets BillingCountry field to given value.


### SetBillingCountryNil

`func (o *Order20230320Response) SetBillingCountryNil(b bool)`

 SetBillingCountryNil sets the value for BillingCountry to be an explicit nil

### UnsetBillingCountry
`func (o *Order20230320Response) UnsetBillingCountry()`

UnsetBillingCountry ensures that no value is present for BillingCountry, not even an explicit nil
### GetPaymentInstrument

`func (o *Order20230320Response) GetPaymentInstrument() string`

GetPaymentInstrument returns the PaymentInstrument field if non-nil, zero value otherwise.

### GetPaymentInstrumentOk

`func (o *Order20230320Response) GetPaymentInstrumentOk() (*string, bool)`

GetPaymentInstrumentOk returns a tuple with the PaymentInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstrument

`func (o *Order20230320Response) SetPaymentInstrument(v string)`

SetPaymentInstrument sets PaymentInstrument field to given value.


### GetCurrency

`func (o *Order20230320Response) GetCurrency() CurrencyResponse`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order20230320Response) GetCurrencyOk() (*CurrencyResponse, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order20230320Response) SetCurrency(v CurrencyResponse)`

SetCurrency sets Currency field to given value.


### GetSubtotal

`func (o *Order20230320Response) GetSubtotal() int32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *Order20230320Response) GetSubtotalOk() (*int32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *Order20230320Response) SetSubtotal(v int32)`

SetSubtotal sets Subtotal field to given value.


### GetShippingFee

`func (o *Order20230320Response) GetShippingFee() int32`

GetShippingFee returns the ShippingFee field if non-nil, zero value otherwise.

### GetShippingFeeOk

`func (o *Order20230320Response) GetShippingFeeOk() (*int32, bool)`

GetShippingFeeOk returns a tuple with the ShippingFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingFee

`func (o *Order20230320Response) SetShippingFee(v int32)`

SetShippingFee sets ShippingFee field to given value.


### SetShippingFeeNil

`func (o *Order20230320Response) SetShippingFeeNil(b bool)`

 SetShippingFeeNil sets the value for ShippingFee to be an explicit nil

### UnsetShippingFee
`func (o *Order20230320Response) UnsetShippingFee()`

UnsetShippingFee ensures that no value is present for ShippingFee, not even an explicit nil
### GetDiscounts

`func (o *Order20230320Response) GetDiscounts() int32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *Order20230320Response) GetDiscountsOk() (*int32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *Order20230320Response) SetDiscounts(v int32)`

SetDiscounts sets Discounts field to given value.


### SetDiscountsNil

`func (o *Order20230320Response) SetDiscountsNil(b bool)`

 SetDiscountsNil sets the value for Discounts to be an explicit nil

### UnsetDiscounts
`func (o *Order20230320Response) UnsetDiscounts()`

UnsetDiscounts ensures that no value is present for Discounts, not even an explicit nil
### GetOtherFees

`func (o *Order20230320Response) GetOtherFees() int32`

GetOtherFees returns the OtherFees field if non-nil, zero value otherwise.

### GetOtherFeesOk

`func (o *Order20230320Response) GetOtherFeesOk() (*int32, bool)`

GetOtherFeesOk returns a tuple with the OtherFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFees

`func (o *Order20230320Response) SetOtherFees(v int32)`

SetOtherFees sets OtherFees field to given value.


### SetOtherFeesNil

`func (o *Order20230320Response) SetOtherFeesNil(b bool)`

 SetOtherFeesNil sets the value for OtherFees to be an explicit nil

### UnsetOtherFees
`func (o *Order20230320Response) UnsetOtherFees()`

UnsetOtherFees ensures that no value is present for OtherFees, not even an explicit nil
### GetTax

`func (o *Order20230320Response) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *Order20230320Response) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *Order20230320Response) SetTax(v int32)`

SetTax sets Tax field to given value.


### SetTaxNil

`func (o *Order20230320Response) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *Order20230320Response) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetGiftCardApplied

`func (o *Order20230320Response) GetGiftCardApplied() int32`

GetGiftCardApplied returns the GiftCardApplied field if non-nil, zero value otherwise.

### GetGiftCardAppliedOk

`func (o *Order20230320Response) GetGiftCardAppliedOk() (*int32, bool)`

GetGiftCardAppliedOk returns a tuple with the GiftCardApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardApplied

`func (o *Order20230320Response) SetGiftCardApplied(v int32)`

SetGiftCardApplied sets GiftCardApplied field to given value.


### SetGiftCardAppliedNil

`func (o *Order20230320Response) SetGiftCardAppliedNil(b bool)`

 SetGiftCardAppliedNil sets the value for GiftCardApplied to be an explicit nil

### UnsetGiftCardApplied
`func (o *Order20230320Response) UnsetGiftCardApplied()`

UnsetGiftCardApplied ensures that no value is present for GiftCardApplied, not even an explicit nil
### GetTotal

`func (o *Order20230320Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Order20230320Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Order20230320Response) SetTotal(v int32)`

SetTotal sets Total field to given value.


### SetTotalNil

`func (o *Order20230320Response) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *Order20230320Response) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetShipments

`func (o *Order20230320Response) GetShipments() []map[string]interface{}`

GetShipments returns the Shipments field if non-nil, zero value otherwise.

### GetShipmentsOk

`func (o *Order20230320Response) GetShipmentsOk() (*[]map[string]interface{}, bool)`

GetShipmentsOk returns a tuple with the Shipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipments

`func (o *Order20230320Response) SetShipments(v []map[string]interface{})`

SetShipments sets Shipments field to given value.


### SetShipmentsNil

`func (o *Order20230320Response) SetShipmentsNil(b bool)`

 SetShipmentsNil sets the value for Shipments to be an explicit nil

### UnsetShipments
`func (o *Order20230320Response) UnsetShipments()`

UnsetShipments ensures that no value is present for Shipments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


