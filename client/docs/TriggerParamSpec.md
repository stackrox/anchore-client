# TriggerParamSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] [default to null]
**Example** | **string** | An example value for the parameter (encoded as a string if the parameter is an object or list type) | [optional] [default to null]
**Name** | **string** | Parameter name as it appears in policy document | [optional] [default to null]
**Required** | **bool** | Is this a required parameter or optional | [optional] [default to null]
**State** | **string** | State of the trigger parameter | [optional] [default to null]
**SupercededBy** | **string** | The name of another trigger that supercedes this on functionally if this is deprecated | [optional] [default to null]
**Validator** | [***interface{}**](interface{}.md) | If present, a definition for validation of input. Typically a jsonschema object that can be used to validate an input against. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


