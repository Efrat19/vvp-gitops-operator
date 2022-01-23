# \UdfArtifactControllerApi

All URIs are relative to *https://host.docker.internal:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUdfArtifactUsingPOST**](UdfArtifactControllerApi.md#CreateUdfArtifactUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/udfartifacts | createUdfArtifact
[**DeleteUdfArtifactUsingDELETE**](UdfArtifactControllerApi.md#DeleteUdfArtifactUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/udfartifacts/{udfArtifactName} | deleteUdfArtifact
[**GetUdfArtifactUsingGET**](UdfArtifactControllerApi.md#GetUdfArtifactUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/udfartifacts/{udfArtifactName} | getUdfArtifact
[**ListUdfArtifactsUsingGET**](UdfArtifactControllerApi.md#ListUdfArtifactsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/udfartifacts | listUdfArtifacts
[**UpdateUdfArtifactUsingPUT**](UdfArtifactControllerApi.md#UpdateUdfArtifactUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/udfartifacts/{udfArtifactName} | updateUdfArtifact


# **CreateUdfArtifactUsingPOST**
> CreateUdfArtifactResponse CreateUdfArtifactUsingPOST(ctx, ns, udfArtifact)
createUdfArtifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **udfArtifact** | [**UdfArtifact**](UdfArtifact.md)| udfArtifact | 

### Return type

[**CreateUdfArtifactResponse**](CreateUdfArtifactResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUdfArtifactUsingDELETE**
> DeleteUdfArtifactResponse DeleteUdfArtifactUsingDELETE(ctx, ns, udfArtifactName)
deleteUdfArtifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **udfArtifactName** | **string**| udfArtifactName | 

### Return type

[**DeleteUdfArtifactResponse**](DeleteUdfArtifactResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUdfArtifactUsingGET**
> GetUdfArtifactResponse GetUdfArtifactUsingGET(ctx, ns, udfArtifactName, optional)
getUdfArtifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **udfArtifactName** | **string**| udfArtifactName | 
 **optional** | ***UdfArtifactControllerApiGetUdfArtifactUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UdfArtifactControllerApiGetUdfArtifactUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requireFunctionNames** | **optional.Bool**| requireFunctionNames | 

### Return type

[**GetUdfArtifactResponse**](GetUdfArtifactResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUdfArtifactsUsingGET**
> ListUdfArtifactsResponse ListUdfArtifactsUsingGET(ctx, ns, optional)
listUdfArtifacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
 **optional** | ***UdfArtifactControllerApiListUdfArtifactsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UdfArtifactControllerApiListUdfArtifactsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requireFunctionNames** | **optional.Bool**| requireFunctionNames | 

### Return type

[**ListUdfArtifactsResponse**](ListUdfArtifactsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUdfArtifactUsingPUT**
> UpdateUdfArtifactResponse UpdateUdfArtifactUsingPUT(ctx, ns, udfArtifact, udfArtifactName)
updateUdfArtifact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **udfArtifact** | [**UdfArtifact**](UdfArtifact.md)| udfArtifact | 
  **udfArtifactName** | **string**| udfArtifactName | 

### Return type

[**UpdateUdfArtifactResponse**](UpdateUdfArtifactResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

