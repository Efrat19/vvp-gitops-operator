# \CatalogControllerApi

All URIs are relative to *https://host.docker.internal:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogUsingGET**](CatalogControllerApi.md#GetCatalogUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat} | getCatalog
[**GetDatabaseUsingGET**](CatalogControllerApi.md#GetDatabaseUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getDatabase | getDatabase
[**GetFunctionUsingGET**](CatalogControllerApi.md#GetFunctionUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getFunction | getFunction
[**GetTableUsingGET**](CatalogControllerApi.md#GetTableUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getTable | getTable
[**GetViewUsingGET**](CatalogControllerApi.md#GetViewUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:getView | getView
[**ListCatalogsUsingGET**](CatalogControllerApi.md#ListCatalogsUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs | listCatalogs
[**ListDatabasesUsingGET**](CatalogControllerApi.md#ListDatabasesUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listDatabases | listDatabases
[**ListFunctionsUsingGET**](CatalogControllerApi.md#ListFunctionsUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listFunctions | listFunctions
[**ListTablesUsingGET**](CatalogControllerApi.md#ListTablesUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listTables | listTables
[**ListViewsUsingGET**](CatalogControllerApi.md#ListViewsUsingGET) | **Get** /catalog/v1beta2/namespaces/{ns}/catalogs/{cat}:listViews | listViews


# **GetCatalogUsingGET**
> GetCatalogResponse GetCatalogUsingGET(ctx, cat, ns)
getCatalog

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 

### Return type

[**GetCatalogResponse**](GetCatalogResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatabaseUsingGET**
> GetDatabaseResponse GetDatabaseUsingGET(ctx, cat, database, ns)
getDatabase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **database** | **string**| database | 
  **ns** | **string**| ns | 

### Return type

[**GetDatabaseResponse**](GetDatabaseResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFunctionUsingGET**
> GetFunctionResponse GetFunctionUsingGET(ctx, cat, function, ns, optional)
getFunction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **function** | **string**| function | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogControllerApiGetFunctionUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiGetFunctionUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **database** | **optional.String**| database | 

### Return type

[**GetFunctionResponse**](GetFunctionResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTableUsingGET**
> GetTableResponse GetTableUsingGET(ctx, cat, ns, table, optional)
getTable

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
  **table** | **string**| table | 
 **optional** | ***CatalogControllerApiGetTableUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiGetTableUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **database** | **optional.String**| database | 

### Return type

[**GetTableResponse**](GetTableResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetViewUsingGET**
> GetViewResponse GetViewUsingGET(ctx, cat, ns, view, optional)
getView

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
  **view** | **string**| view | 
 **optional** | ***CatalogControllerApiGetViewUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiGetViewUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **database** | **optional.String**| database | 

### Return type

[**GetViewResponse**](GetViewResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogsUsingGET**
> ListCatalogsResponse ListCatalogsUsingGET(ctx, ns)
listCatalogs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListCatalogsResponse**](ListCatalogsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDatabasesUsingGET**
> ListDatabasesResponse ListDatabasesUsingGET(ctx, cat, ns)
listDatabases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 

### Return type

[**ListDatabasesResponse**](ListDatabasesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFunctionsUsingGET**
> ListFunctionsResponse ListFunctionsUsingGET(ctx, cat, ns, optional)
listFunctions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogControllerApiListFunctionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiListFunctionsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **database** | **optional.String**| database | 

### Return type

[**ListFunctionsResponse**](ListFunctionsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesUsingGET**
> ListTablesResponse ListTablesUsingGET(ctx, cat, ns, optional)
listTables

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogControllerApiListTablesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiListTablesUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **database** | **optional.String**| database | 

### Return type

[**ListTablesResponse**](ListTablesResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListViewsUsingGET**
> ListViewsResponse ListViewsUsingGET(ctx, cat, ns, optional)
listViews

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cat** | **string**| cat | 
  **ns** | **string**| ns | 
 **optional** | ***CatalogControllerApiListViewsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogControllerApiListViewsUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **database** | **optional.String**| database | 

### Return type

[**ListViewsResponse**](ListViewsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

