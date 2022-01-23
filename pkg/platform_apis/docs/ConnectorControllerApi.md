# \ConnectorControllerApi

All URIs are relative to *https://host.docker.internal:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeCatalogConnectorJarsUsingPOST**](ConnectorControllerApi.md#AnalyzeCatalogConnectorJarsUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/catalog-connectors:analyze-jars | analyzeCatalogConnectorJars
[**CreateCatalogConnectorUsingPOST**](ConnectorControllerApi.md#CreateCatalogConnectorUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/catalog-connectors | createCatalogConnector
[**CreateConnectorUsingPOST**](ConnectorControllerApi.md#CreateConnectorUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/connectors | createConnector
[**CreateFormatUsingPOST**](ConnectorControllerApi.md#CreateFormatUsingPOST) | **Post** /sql/v1beta1/namespaces/{ns}/formats | createFormat
[**DeleteCatalogConnectorUsingDELETE**](ConnectorControllerApi.md#DeleteCatalogConnectorUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/catalog-connectors/{name} | deleteCatalogConnector
[**DeleteConnectorUsingDELETE**](ConnectorControllerApi.md#DeleteConnectorUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/connectors/{name} | deleteConnector
[**DeleteFormatUsingDELETE**](ConnectorControllerApi.md#DeleteFormatUsingDELETE) | **Delete** /sql/v1beta1/namespaces/{ns}/formats/{name} | deleteFormat
[**GetCatalogConnectorUsingGET**](ConnectorControllerApi.md#GetCatalogConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/catalog-connectors/{catalogConnectorResourceId} | getCatalogConnector
[**GetConnectorUsingGET**](ConnectorControllerApi.md#GetConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors/{connectorResourceId} | getConnector
[**GetFormatUsingGET**](ConnectorControllerApi.md#GetFormatUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats/{formatResourceId} | getFormat
[**ListCatalogConnectorsUsingGET**](ConnectorControllerApi.md#ListCatalogConnectorsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/catalog-connectors | listCatalogConnectors
[**ListCatalogsReferencingCatalogConnectorUsingGET**](ConnectorControllerApi.md#ListCatalogsReferencingCatalogConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/catalog-connectors/{name}:list-catalogs | listCatalogsReferencingCatalogConnector
[**ListConnectorsUsingGET**](ConnectorControllerApi.md#ListConnectorsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors | listConnectors
[**ListFormatsUsingGET**](ConnectorControllerApi.md#ListFormatsUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats | listFormats
[**ListTablesReferencingConnectorUsingGET**](ConnectorControllerApi.md#ListTablesReferencingConnectorUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/connectors/{name}:list-tables | listTablesReferencingConnector
[**ListTablesReferencingFormatUsingGET**](ConnectorControllerApi.md#ListTablesReferencingFormatUsingGET) | **Get** /sql/v1beta1/namespaces/{ns}/formats/{name}:list-tables | listTablesReferencingFormat
[**UpdateCatalogConnectorUsingPUT**](ConnectorControllerApi.md#UpdateCatalogConnectorUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/catalog-connectors/{name} | updateCatalogConnector
[**UpdateConnectorUsingPUT**](ConnectorControllerApi.md#UpdateConnectorUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/connectors/{name} | updateConnector
[**UpdateFormatUsingPUT**](ConnectorControllerApi.md#UpdateFormatUsingPUT) | **Put** /sql/v1beta1/namespaces/{ns}/formats/{name} | updateFormat


# **AnalyzeCatalogConnectorJarsUsingPOST**
> AnalyzeCatalogConnectorJarsResponse AnalyzeCatalogConnectorJarsUsingPOST(ctx, jarUris, ns)
analyzeCatalogConnectorJars

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jarUris** | [**JarUris**](JarUris.md)| jarUris | 
  **ns** | **string**| ns | 

### Return type

[**AnalyzeCatalogConnectorJarsResponse**](AnalyzeCatalogConnectorJarsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCatalogConnectorUsingPOST**
> CreateCatalogConnectorResponse CreateCatalogConnectorUsingPOST(ctx, catalogConnector, ns)
createCatalogConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalogConnector** | [**CatalogConnector**](CatalogConnector.md)| catalogConnector | 
  **ns** | **string**| ns | 

### Return type

[**CreateCatalogConnectorResponse**](CreateCatalogConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnectorUsingPOST**
> CreateConnectorResponse CreateConnectorUsingPOST(ctx, connector, ns)
createConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | [**Connector**](Connector.md)| connector | 
  **ns** | **string**| ns | 

### Return type

[**CreateConnectorResponse**](CreateConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFormatUsingPOST**
> CreateFormatResponse CreateFormatUsingPOST(ctx, format, ns)
createFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **format** | [**Format**](Format.md)| format | 
  **ns** | **string**| ns | 

### Return type

[**CreateFormatResponse**](CreateFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCatalogConnectorUsingDELETE**
> DeleteCatalogConnectorResponse DeleteCatalogConnectorUsingDELETE(ctx, name, ns)
deleteCatalogConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**DeleteCatalogConnectorResponse**](DeleteCatalogConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnectorUsingDELETE**
> DeleteConnectorResponse DeleteConnectorUsingDELETE(ctx, name, ns)
deleteConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**DeleteConnectorResponse**](DeleteConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFormatUsingDELETE**
> DeleteFormatResponse DeleteFormatUsingDELETE(ctx, name, ns)
deleteFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**DeleteFormatResponse**](DeleteFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCatalogConnectorUsingGET**
> GetCatalogConnectorResponse GetCatalogConnectorUsingGET(ctx, catalogConnectorResourceId, ns)
getCatalogConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalogConnectorResourceId** | **string**| catalogConnectorResourceId | 
  **ns** | **string**| ns | 

### Return type

[**GetCatalogConnectorResponse**](GetCatalogConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorUsingGET**
> GetConnectorResponse GetConnectorUsingGET(ctx, connectorResourceId, ns)
getConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorResourceId** | **string**| connectorResourceId | 
  **ns** | **string**| ns | 

### Return type

[**GetConnectorResponse**](GetConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFormatUsingGET**
> GetFormatResponse GetFormatUsingGET(ctx, formatResourceId, ns)
getFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **formatResourceId** | **string**| formatResourceId | 
  **ns** | **string**| ns | 

### Return type

[**GetFormatResponse**](GetFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogConnectorsUsingGET**
> ListCatalogConnectorsResponse ListCatalogConnectorsUsingGET(ctx, ns)
listCatalogConnectors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListCatalogConnectorsResponse**](ListCatalogConnectorsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalogsReferencingCatalogConnectorUsingGET**
> ListCatalogsReferencingCatalogConnectorResponse ListCatalogsReferencingCatalogConnectorUsingGET(ctx, name, ns)
listCatalogsReferencingCatalogConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**ListCatalogsReferencingCatalogConnectorResponse**](ListCatalogsReferencingCatalogConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListConnectorsUsingGET**
> ListConnectorsResponse ListConnectorsUsingGET(ctx, ns)
listConnectors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListConnectorsResponse**](ListConnectorsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFormatsUsingGET**
> ListFormatsResponse ListFormatsUsingGET(ctx, ns)
listFormats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ns** | **string**| ns | 

### Return type

[**ListFormatsResponse**](ListFormatsResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesReferencingConnectorUsingGET**
> ListTablesReferencingConnectorResponse ListTablesReferencingConnectorUsingGET(ctx, name, ns)
listTablesReferencingConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**ListTablesReferencingConnectorResponse**](ListTablesReferencingConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTablesReferencingFormatUsingGET**
> ListTablesReferencingFormatResponse ListTablesReferencingFormatUsingGET(ctx, name, ns)
listTablesReferencingFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**ListTablesReferencingFormatResponse**](ListTablesReferencingFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCatalogConnectorUsingPUT**
> UpdateCatalogConnectorResponse UpdateCatalogConnectorUsingPUT(ctx, catalogConnector, name, ns)
updateCatalogConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalogConnector** | [**CatalogConnector**](CatalogConnector.md)| catalogConnector | 
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**UpdateCatalogConnectorResponse**](UpdateCatalogConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectorUsingPUT**
> UpdateConnectorResponse UpdateConnectorUsingPUT(ctx, connector, name, ns)
updateConnector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connector** | [**Connector**](Connector.md)| connector | 
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**UpdateConnectorResponse**](UpdateConnectorResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFormatUsingPUT**
> UpdateFormatResponse UpdateFormatUsingPUT(ctx, format, name, ns)
updateFormat

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **format** | [**Format**](Format.md)| format | 
  **name** | **string**| name | 
  **ns** | **string**| ns | 

### Return type

[**UpdateFormatResponse**](UpdateFormatResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

