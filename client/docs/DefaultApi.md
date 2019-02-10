# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthNoop**](DefaultApi.md#HealthNoop) | **Get** /health | 
[**Ping**](DefaultApi.md#Ping) | **Get** / | 
[**VersionNoop**](DefaultApi.md#VersionNoop) | **Get** /version | 


# **HealthNoop**
> HealthNoop(ctx, )


Health check, returns 200 and no body if service is running

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> string Ping(ctx, )


Simple status check

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VersionNoop**
> ServiceVersion VersionNoop(ctx, )


Returns the version object for the service, including db schema version info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceVersion**](ServiceVersion.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

