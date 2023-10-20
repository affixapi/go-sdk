# IntrospectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client ID you received when you first created the application | [readonly] 
**Mode** | [**ModeResponse**](ModeResponse.md) |  | 
**Provider** | [**ProviderResponse**](ProviderResponse.md) |  | 
**Scopes** | [**[]ScopesResponse**](ScopesResponse.md) | One or more scope values indicating which parts of the user&#39;s account you wish to access.  Note, slight deviation from the OAuth 2.1 spec in that the param is scopes (plural) is used vs scope (singular)  | [readonly] 

## Methods

### NewIntrospectResponse

`func NewIntrospectResponse(clientId string, mode ModeResponse, provider ProviderResponse, scopes []ScopesResponse, ) *IntrospectResponse`

NewIntrospectResponse instantiates a new IntrospectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntrospectResponseWithDefaults

`func NewIntrospectResponseWithDefaults() *IntrospectResponse`

NewIntrospectResponseWithDefaults instantiates a new IntrospectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IntrospectResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IntrospectResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IntrospectResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetMode

`func (o *IntrospectResponse) GetMode() ModeResponse`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *IntrospectResponse) GetModeOk() (*ModeResponse, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *IntrospectResponse) SetMode(v ModeResponse)`

SetMode sets Mode field to given value.


### GetProvider

`func (o *IntrospectResponse) GetProvider() ProviderResponse`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IntrospectResponse) GetProviderOk() (*ProviderResponse, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IntrospectResponse) SetProvider(v ProviderResponse)`

SetProvider sets Provider field to given value.


### GetScopes

`func (o *IntrospectResponse) GetScopes() []ScopesResponse`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IntrospectResponse) GetScopesOk() (*[]ScopesResponse, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IntrospectResponse) SetScopes(v []ScopesResponse)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


