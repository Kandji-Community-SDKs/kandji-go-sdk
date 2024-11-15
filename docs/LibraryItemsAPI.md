# \LibraryItemsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLibraryItemActivity**](LibraryItemsAPI.md#GetLibraryItemActivity) | **Get** /api/v1/library/library-items/{library_item_id}/activity | Get Library Item Activity
[**GetLibraryItemStatuses**](LibraryItemsAPI.md#GetLibraryItemStatuses) | **Get** /api/v1/library/library-items/{library_item_id}/status | Get Library Item Statuses



## GetLibraryItemActivity

> map[string]interface{} GetLibraryItemActivity(ctx, libraryItemId).ActivityType(activityType).UserId(userId).UserEmail(userEmail).Limit(limit).Offset(offset).Execute()

Get Library Item Activity



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kandji-Community-SDKs/kandji-go-sdk"
)

func main() {
	libraryItemId := "libraryItemId_example" // string | 
	activityType := "library_item_assignment_changed" // string | Filter actions by this activity type. Choices are: library_item_created, library_item_edited, library_item_deleted, library_item_duplicated, library_item_assignment_changed (optional)
	userId := "{user uuid}" // string | Filter actions by this user (id) (optional)
	userEmail := "me@example.com" // string | Filter actions by this user (email) (optional)
	limit := "10" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "100" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryItemsAPI.GetLibraryItemActivity(context.Background(), libraryItemId).ActivityType(activityType).UserId(userId).UserEmail(userEmail).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryItemsAPI.GetLibraryItemActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryItemActivity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LibraryItemsAPI.GetLibraryItemActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryItemActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityType** | **string** | Filter actions by this activity type. Choices are: library_item_created, library_item_edited, library_item_deleted, library_item_duplicated, library_item_assignment_changed | 
 **userId** | **string** | Filter actions by this user (id) | 
 **userEmail** | **string** | Filter actions by this user (email) | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryItemStatuses

> map[string]interface{} GetLibraryItemStatuses(ctx, libraryItemId).ComputerId(computerId).Limit(limit).Offset(offset).Execute()

Get Library Item Statuses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Kandji-Community-SDKs/kandji-go-sdk"
)

func main() {
	libraryItemId := "libraryItemId_example" // string | 
	computerId := "{device_id}" // string | Query for the status of one device. (optional)
	limit := "300" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results. (optional)
	offset := "offset_example" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryItemsAPI.GetLibraryItemStatuses(context.Background(), libraryItemId).ComputerId(computerId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryItemsAPI.GetLibraryItemStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryItemStatuses`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LibraryItemsAPI.GetLibraryItemStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryItemStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **computerId** | **string** | Query for the status of one device. | 
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

