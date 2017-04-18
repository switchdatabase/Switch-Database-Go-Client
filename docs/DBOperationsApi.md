# \DBOperationsApi

All URIs are relative to *http://tr02.switchapi.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPost**](DBOperationsApi.md#AddPost) | **Post** /Add | Add is used for adding a data object to the list created at Switch DB.
[**ListPost**](DBOperationsApi.md#ListPost) | **Post** /List | It&#39;s used for listing a data added before.
[**SetDelete**](DBOperationsApi.md#SetDelete) | **Delete** /Set | It&#39;s used for deleting a data added before at Switch DB.
[**SetPost**](DBOperationsApi.md#SetPost) | **Post** /Set | It&#39;s used for updating a data object that is already added to Switch DB.


# **AddPost**
> AddResponse AddPost($aPIKey, $accessToken, $list, $body)

Add is used for adding a data object to the list created at Switch DB.

You can choose the list that will be added tha data set to with List parameter that will be sent to Header. It's equal to INSERT query at the relational databases model. The data set that will be added to the database is transmited at request body. For versions upper than v1.2.1, if the data sent is an array, all items in the array are added one by one, so they are added as a whole. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIKey** | **string**| Your Switch API Key. | 
 **accessToken** | **string**| Your Access Token. | 
 **list** | **string**| Your data list name. | 
 **body** | **string**| Your new value JSON. | 

### Return type

[**AddResponse**](AddResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPost**
> ListPost($aPIKey, $accessToken, $list, $body)

It's used for listing a data added before.

List parameter sent remarks the list that will be do listing work on at Header. It's equal to SELECT query at relational databases. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIKey** | **string**| Your Switch API Key. | 
 **accessToken** | **string**| Your Access Token. | 
 **list** | **string**| Your data list name. | 
 **body** | [**Body**](Body.md)| Your Switch DB Query. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDelete**
> SetResponse SetDelete($aPIKey, $accessToken, $list, $listItemId)

It's used for deleting a data added before at Switch DB.

List parameter sent remarks the list that will be deleted data from at Header and ListItemId parameter remarks the ID consisted by Switch DB for the data that will be deleted. It's equal to DELETE query at relational databases. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIKey** | **string**| Your Switch API Key. | 
 **accessToken** | **string**| Your Access Token. | 
 **list** | **string**| Your data list name. | 
 **listItemId** | **string**| Your list item id. | 

### Return type

[**SetResponse**](SetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPost**
> SetResponse SetPost($aPIKey, $accessToken, $list, $listItemId, $body)

It's used for updating a data object that is already added to Switch DB.

In order to UPDATE a object, Listname and ListItemId parameters should be sent in the Header of the REQUEST as \"List\", \"ListItemId\", respectively, as shown in the example below. It's equal to UPDATE query at relational databases. The data set that will be edited is transmited to the database at request body. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIKey** | **string**| Your Switch API Key. | 
 **accessToken** | **string**| Your Access Token. | 
 **list** | **string**| Your data list name. | 
 **listItemId** | **string**| Your list item id. | 
 **body** | **string**| Your new value JSON. | 

### Return type

[**SetResponse**](SetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

