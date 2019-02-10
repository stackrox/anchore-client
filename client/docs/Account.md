# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the account was created | [optional] [default to null]
**Email** | **string** | Optional email address associated with the account | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | The timestamp of the last update to the account metadata itself (not users or creds) | [optional] [default to null]
**Name** | **string** | The account identifier, not updatable after creation | [default to null]
**State** | **string** | State of the account. Disabled accounts prevent member users from logging in, deleting accounts are disabled and pending deletion and will be removed once all owned resources are garbage collected by the system | [optional] [default to null]
**Type_** | **string** | The user type (admin vs user). If not specified in a POST request, &#39;user&#39; is default | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


