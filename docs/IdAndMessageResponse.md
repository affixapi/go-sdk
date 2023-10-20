# IdAndMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A description of the error | 
**Id** | **string** | The id of the entity that the id applies. For example, the transaction id | 

## Methods

### NewIdAndMessageResponse

`func NewIdAndMessageResponse(message string, id string, ) *IdAndMessageResponse`

NewIdAndMessageResponse instantiates a new IdAndMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdAndMessageResponseWithDefaults

`func NewIdAndMessageResponseWithDefaults() *IdAndMessageResponse`

NewIdAndMessageResponseWithDefaults instantiates a new IdAndMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *IdAndMessageResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IdAndMessageResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IdAndMessageResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetId

`func (o *IdAndMessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdAndMessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdAndMessageResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


