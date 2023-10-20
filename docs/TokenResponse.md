# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The issued access_token | [readonly] 
**Mode** | [**ModeResponse**](ModeResponse.md) |  | 
**Provider** | [**ProviderResponse**](ProviderResponse.md) |  | 
**Scopes** | [**[]ScopesResponse**](ScopesResponse.md) | One or more scope values indicating which parts of the user&#39;s account you wish to access.  Note, slight deviation from the OAuth 2.1 spec in that the param is scopes (plural) is used vs scope (singular)  | [readonly] 
**TokenType** | **string** | The token type to pass in the &#x60;Authorization&#x60; header | [readonly] 

## Methods

### NewTokenResponse

`func NewTokenResponse(accessToken string, mode ModeResponse, provider ProviderResponse, scopes []ScopesResponse, tokenType string, ) *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetMode

`func (o *TokenResponse) GetMode() ModeResponse`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TokenResponse) GetModeOk() (*ModeResponse, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TokenResponse) SetMode(v ModeResponse)`

SetMode sets Mode field to given value.


### GetProvider

`func (o *TokenResponse) GetProvider() ProviderResponse`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *TokenResponse) GetProviderOk() (*ProviderResponse, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *TokenResponse) SetProvider(v ProviderResponse)`

SetProvider sets Provider field to given value.


### GetScopes

`func (o *TokenResponse) GetScopes() []ScopesResponse`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenResponse) GetScopesOk() (*[]ScopesResponse, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenResponse) SetScopes(v []ScopesResponse)`

SetScopes sets Scopes field to given value.


### GetTokenType

`func (o *TokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


