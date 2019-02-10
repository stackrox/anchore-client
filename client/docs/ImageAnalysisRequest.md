# ImageAnalysisRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | [***interface{}**](interface{}.md) | Annotations to be associated with the added image in key/value form | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Optional override of the image creation time, only honored when both tag and digest are also supplied  e.g. 2018-10-17T18:14:00Z | [optional] [default to null]
**Digest** | **string** | A full pullable digest reference for an image. e.g. docker.io/nginx@sha256:abc123 | [optional] [default to null]
**Dockerfile** | **string** | Content of the dockerfile for the image, if available | [optional] [default to null]
**ImageType** | **string** | The type of image this is adding, defaults to \&quot;docker\&quot; | [optional] [default to null]
**Tag** | **string** | Full pullable tag reference for image. e.g. docker.io/nginx:latest | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


