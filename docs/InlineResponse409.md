# InlineResponse409

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The authorization code insert that caused a conflict  | [optional] 
**Token** | Pointer to **string** | The access token insert that caused a conflict  | [optional] 
**Error** | **string** | The description of the type of error, ie conflict of two requests of the the same authorization code  | 

## Methods

### NewInlineResponse409

`func NewInlineResponse409(error_ string, ) *InlineResponse409`

NewInlineResponse409 instantiates a new InlineResponse409 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse409WithDefaults

`func NewInlineResponse409WithDefaults() *InlineResponse409`

NewInlineResponse409WithDefaults instantiates a new InlineResponse409 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineResponse409) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse409) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse409) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse409) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetToken

`func (o *InlineResponse409) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InlineResponse409) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InlineResponse409) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InlineResponse409) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse409) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse409) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse409) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


