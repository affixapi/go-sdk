# ClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **[]string** | The client secret(s). It is an array datatype to allow for rotation of secrets without downtime for your customers  | 
**RedirectUris** | **[]string** | Indicates the URI to return the user to after authorization is complete, which is the endpoint on your server to receive the authorization_code.  Must be identical to the redirect URI provided in the original link.  Please email me after signup and I will set both your client secret and redirect_uri (required) when you reach out.  | 
**Name** | **string** | Name of your app that shows up in the Affix Connect application  | 
**WebhookUri** | Pointer to **NullableString** | If enabled, webhooks will be sent to this endpoint  | [optional] 

## Methods

### NewClientRequest

`func NewClientRequest(clientSecret []string, redirectUris []string, name string, ) *ClientRequest`

NewClientRequest instantiates a new ClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRequestWithDefaults

`func NewClientRequestWithDefaults() *ClientRequest`

NewClientRequestWithDefaults instantiates a new ClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *ClientRequest) GetClientSecret() []string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientRequest) GetClientSecretOk() (*[]string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientRequest) SetClientSecret(v []string)`

SetClientSecret sets ClientSecret field to given value.


### GetRedirectUris

`func (o *ClientRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ClientRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ClientRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetName

`func (o *ClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWebhookUri

`func (o *ClientRequest) GetWebhookUri() string`

GetWebhookUri returns the WebhookUri field if non-nil, zero value otherwise.

### GetWebhookUriOk

`func (o *ClientRequest) GetWebhookUriOk() (*string, bool)`

GetWebhookUriOk returns a tuple with the WebhookUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUri

`func (o *ClientRequest) SetWebhookUri(v string)`

SetWebhookUri sets WebhookUri field to given value.

### HasWebhookUri

`func (o *ClientRequest) HasWebhookUri() bool`

HasWebhookUri returns a boolean if a field has been set.

### SetWebhookUriNil

`func (o *ClientRequest) SetWebhookUriNil(b bool)`

 SetWebhookUriNil sets the value for WebhookUri to be an explicit nil

### UnsetWebhookUri
`func (o *ClientRequest) UnsetWebhookUri()`

UnsetWebhookUri ensures that no value is present for WebhookUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


