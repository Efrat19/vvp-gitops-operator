# \SqlControllerApi

All URIs are relative to *https://host.docker.internal:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSqlScriptUsingPOST**](SqlControllerApi.md#CreateSqlScriptUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/sqlscripts | createSqlScript
[**DeleteSqlScriptUsingDELETE**](SqlControllerApi.md#DeleteSqlScriptUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/sqlscripts/{sqlScriptName} | deleteSqlScript
[**ExecuteStatementUsingPOST**](SqlControllerApi.md#ExecuteStatementUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/sqlscripts:execute | executeStatement
[**ExecuteStatementsUsingPOST**](SqlControllerApi.md#ExecuteStatementsUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/sqlscripts:execute-multi | executeStatements
[**GetSqlScriptUsingGET**](SqlControllerApi.md#GetSqlScriptUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/sqlscripts/{sqlScriptName} | getSqlScript
[**ListSqlScriptsUsingGET**](SqlControllerApi.md#ListSqlScriptsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/sqlscripts | listSqlScripts
[**SuggestSqlScriptCompletionsUsingPOST**](SqlControllerApi.md#SuggestSqlScriptCompletionsUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/sqlscripts:suggest | suggestSqlScriptCompletions
[**UpdateSqlScriptUsingPUT**](SqlControllerApi.md#UpdateSqlScriptUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/sqlscripts/{sqlScriptName} | updateSqlScript
[**ValidateStatementUsingPOST**](SqlControllerApi.md#ValidateStatementUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/sqlscripts:validate | validateStatement


# **CreateSqlScriptUsingPOST**
> CreateSqlScriptResponse CreateSqlScriptUsingPOST(ctx, ns, sqlScript)
createSqlScript

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **sqlScript** | [**SqlScript**](SqlScript.md)| sqlScript | 

### Return type

[**CreateSqlScriptResponse**](CreateSqlScriptResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSqlScriptUsingDELETE**
> DeleteSqlScriptResponse DeleteSqlScriptUsingDELETE(ctx, ns, sqlScriptName)
deleteSqlScript

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **sqlScriptName** | **string**| sqlScriptName | 

### Return type

[**DeleteSqlScriptResponse**](DeleteSqlScriptResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteStatementUsingPOST**
> ExecuteStatementResponse ExecuteStatementUsingPOST(ctx, ns, statement, optional)
executeStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **statement** | [**Statement**](Statement.md)| statement | 
 **optional** | ***SqlControllerApiExecuteStatementUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SqlControllerApiExecuteStatementUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **catalog** | **optional.String**| catalog | 
 **database** | **optional.String**| database | 

### Return type

[**ExecuteStatementResponse**](ExecuteStatementResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecuteStatementsUsingPOST**
> ExecuteStatementsResponse ExecuteStatementsUsingPOST(ctx, ns, statements, optional)
executeStatements

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **statements** | **[]string**| statements | 
 **optional** | ***SqlControllerApiExecuteStatementsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SqlControllerApiExecuteStatementsUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **catalog** | **optional.String**| catalog | 
 **database** | **optional.String**| database | 
 **stopOnError** | **optional.Bool**| stopOnError | [default to true]

### Return type

[**ExecuteStatementsResponse**](ExecuteStatementsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSqlScriptUsingGET**
> GetSqlScriptResponse GetSqlScriptUsingGET(ctx, ns, sqlScriptName)
getSqlScript

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **sqlScriptName** | **string**| sqlScriptName | 

### Return type

[**GetSqlScriptResponse**](GetSqlScriptResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSqlScriptsUsingGET**
> ListSqlScriptsResponse ListSqlScriptsUsingGET(ctx, ns)
listSqlScripts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListSqlScriptsResponse**](ListSqlScriptsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SuggestSqlScriptCompletionsUsingPOST**
> SuggestSqlScriptCompletionsResponse SuggestSqlScriptCompletionsUsingPOST(ctx, details, ns)
suggestSqlScriptCompletions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **details** | [**SuggestSqlScriptCompletionsDetails**](SuggestSqlScriptCompletionsDetails.md)| details | 
  **ns** | **string**| ns | 

### Return type

[**SuggestSqlScriptCompletionsResponse**](SuggestSqlScriptCompletionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSqlScriptUsingPUT**
> UpdateSqlScriptResponse UpdateSqlScriptUsingPUT(ctx, ns, sqlScript, sqlScriptName)
updateSqlScript

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **sqlScript** | [**SqlScript**](SqlScript.md)| sqlScript | 
  **sqlScriptName** | **string**| sqlScriptName | 

### Return type

[**UpdateSqlScriptResponse**](UpdateSqlScriptResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateStatementUsingPOST**
> ValidateStatementResponse ValidateStatementUsingPOST(ctx, ns, statement, optional)
validateStatement

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 
  **statement** | [**Statement**](Statement.md)| statement | 
 **optional** | ***SqlControllerApiValidateStatementUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SqlControllerApiValidateStatementUsingPOSTOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **catalog** | **optional.String**| catalog | 
 **database** | **optional.String**| database | 

### Return type

[**ValidateStatementResponse**](ValidateStatementResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

