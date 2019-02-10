# MappingRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Image** | [***ImageRef**](ImageRef.md) |  | [default to null]
**Name** | **string** |  | [default to null]
**PolicyId** | **string** | Optional single policy to evalute, if set will override any value in policy_ids, for backwards compatibility. Generally, policy_ids should be used even with a array of length 1. | [optional] [default to null]
**PolicyIds** | **[]string** | List of policyIds to evaluate in order, to completion | [optional] [default to null]
**Registry** | **string** |  | [default to null]
**Repository** | **string** |  | [default to null]
**WhitelistIds** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


