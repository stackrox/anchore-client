# \SystemApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteService**](SystemApi.md#DeleteService) | **Delete** /system/services/{servicename}/{hostid} | Delete the service config
[**DescribePolicy**](SystemApi.md#DescribePolicy) | **Get** /system/policy_spec | Describe the policy language spec implemented by this service.
[**GetServiceDetail**](SystemApi.md#GetServiceDetail) | **Get** /system | System status
[**GetServicesByName**](SystemApi.md#GetServicesByName) | **Get** /system/services/{servicename} | Get a service configuration and state
[**GetServicesByNameAndHost**](SystemApi.md#GetServicesByNameAndHost) | **Get** /system/services/{servicename}/{hostid} | Get service config for a specific host
[**GetStatus**](SystemApi.md#GetStatus) | **Get** /status | Service status
[**GetSystemFeeds**](SystemApi.md#GetSystemFeeds) | **Get** /system/feeds | list feeds operations and information
[**ListServices**](SystemApi.md#ListServices) | **Get** /system/services | List system services
[**PostSystemFeeds**](SystemApi.md#PostSystemFeeds) | **Post** /system/feeds | trigger feeds operations


# **DeleteService**
> DeleteService(ctx, servicename, hostid)
Delete the service config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **servicename** | **string**|  | 
  **hostid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DescribePolicy**
> []GateSpec DescribePolicy(ctx, )
Describe the policy language spec implemented by this service.

Get the policy language spec for this service

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GateSpec**](GateSpec.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDetail**
> SystemStatusResponse GetServiceDetail(ctx, )
System status

Get the system status including queue lengths

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemStatusResponse**](SystemStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicesByName**
> ServiceList GetServicesByName(ctx, servicename)
Get a service configuration and state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **servicename** | **string**|  | 

### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServicesByNameAndHost**
> ServiceList GetServicesByNameAndHost(ctx, servicename, hostid)
Get service config for a specific host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **servicename** | **string**|  | 
  **hostid** | **string**|  | 

### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatus**
> StatusResponse GetStatus(ctx, )
Service status

Get the API service status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatusResponse**](StatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemFeeds**
> []FeedMetadata GetSystemFeeds(ctx, )
list feeds operations and information



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]FeedMetadata**](FeedMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServices**
> ServiceList ListServices(ctx, )
List system services

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceList**](ServiceList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostSystemFeeds**
> PostSystemFeeds(ctx, optional)
trigger feeds operations



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostSystemFeedsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostSystemFeedsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flush** | **optional.Bool**| instruct system to flush existing data feeds records from anchore-engine | 
 **sync** | **optional.Bool**| instruct system to re-sync data feeds | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

