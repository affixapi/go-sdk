# DisconnectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disconnected** | **bool** | The access token has been disconnected and is no longer allowed to be used by your application | [readonly] 

## Methods

### NewDisconnectResponse

`func NewDisconnectResponse(disconnected bool, ) *DisconnectResponse`

NewDisconnectResponse instantiates a new DisconnectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisconnectResponseWithDefaults

`func NewDisconnectResponseWithDefaults() *DisconnectResponse`

NewDisconnectResponseWithDefaults instantiates a new DisconnectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnected

`func (o *DisconnectResponse) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *DisconnectResponse) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *DisconnectResponse) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


