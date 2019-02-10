# TagEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectedAt** | [**time.Time**](time.Time.md) | The timestamp at which the Anchore Engine detected this tag was mapped to the image digest. Does not necessarily indicate when the tag was actually pushed to the registry. | [optional] [default to null]
**Pullstring** | **string** | The pullable string for the tag. E.g. \&quot;docker.io/library/node:latest\&quot; | [optional] [default to null]
**Registry** | **string** | The registry hostname:port section of the pull string | [optional] [default to null]
**Repository** | **string** | The repository section of the pull string | [optional] [default to null]
**Tag** | **string** | The tag-only section of the pull string | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


