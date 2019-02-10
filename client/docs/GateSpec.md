# GateSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the gate | [optional] [default to null]
**Name** | **string** | Gate name, as it would appear in a policy document | [optional] [default to null]
**State** | **string** | State of the gate and transitively all triggers it contains if not &#39;active&#39; | [optional] [default to null]
**SupercededBy** | **string** | The name of another trigger that supercedes this on functionally if this is deprecated | [optional] [default to null]
**Triggers** | [**[]TriggerSpec**](TriggerSpec.md) | List of the triggers that can fire for this Gate | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


