# \CustomAppsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomApp**](CustomAppsAPI.md#CreateCustomApp) | **Post** /api/v1/library/custom-apps | Create Custom App
[**DeleteCustomApp**](CustomAppsAPI.md#DeleteCustomApp) | **Delete** /api/v1/library/custom-apps/{library_item_id} | Delete Custom App
[**GetCustomApp**](CustomAppsAPI.md#GetCustomApp) | **Get** /api/v1/library/custom-apps/{library_item_id} | Get Custom App
[**ListCustomApps**](CustomAppsAPI.md#ListCustomApps) | **Get** /api/v1/library/custom-apps | List Custom Apps
[**UpdateCustomApp**](CustomAppsAPI.md#UpdateCustomApp) | **Patch** /api/v1/library/custom-apps/{library_item_id} | Update Custom App
[**UploadCustomApp**](CustomAppsAPI.md#UploadCustomApp) | **Post** /api/v1/library/custom-apps/upload | Upload Custom App



## CreateCustomApp

> map[string]interface{} CreateCustomApp(ctx).Name(name).FileKey(fileKey).InstallType(installType).InstallEnforcement(installEnforcement).ShowInSelfService(showInSelfService).SelfServiceCategoryId(selfServiceCategoryId).SelfServiceRecommended(selfServiceRecommended).Execute()

Create Custom App



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
	name := "name_example" // string | (Required) The name for this Custom App
	fileKey := "fileKey_example" // string | (Required) The S3 key from the <code>Upload Custom App</code> endpont used to upload the custom app file.
	installType := "installType_example" // string | (Required) Options are package, zip, image
	installEnforcement := "installEnforcement_example" // string | (Required) Options are install_once, continuously_enforce, no_enforcement
	showInSelfService := "showInSelfService_example" // string | (Optional, default=false) Displays this app in Self Service
	selfServiceCategoryId := "selfServiceCategoryId_example" // string | (Required for show_in_self_service=true) Self Service Category (by ID) to display app in
	selfServiceRecommended := "selfServiceRecommended_example" // string | (Optional, default=false) Adds recommended flag to app in Self Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.CreateCustomApp(context.Background()).Name(name).FileKey(fileKey).InstallType(installType).InstallEnforcement(installEnforcement).ShowInSelfService(showInSelfService).SelfServiceCategoryId(selfServiceCategoryId).SelfServiceRecommended(selfServiceRecommended).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.CreateCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.CreateCustomApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (Required) The name for this Custom App | 
 **fileKey** | **string** | (Required) The S3 key from the &lt;code&gt;Upload Custom App&lt;/code&gt; endpont used to upload the custom app file. | 
 **installType** | **string** | (Required) Options are package, zip, image | 
 **installEnforcement** | **string** | (Required) Options are install_once, continuously_enforce, no_enforcement | 
 **showInSelfService** | **string** | (Optional, default&#x3D;false) Displays this app in Self Service | 
 **selfServiceCategoryId** | **string** | (Required for show_in_self_service&#x3D;true) Self Service Category (by ID) to display app in | 
 **selfServiceRecommended** | **string** | (Optional, default&#x3D;false) Adds recommended flag to app in Self Service | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomApp

> DeleteCustomApp(ctx, libraryItemId).Execute()

Delete Custom App



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomAppsAPI.DeleteCustomApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.DeleteCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomApp

> map[string]interface{} GetCustomApp(ctx, libraryItemId).Execute()

Get Custom App



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.GetCustomApp(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.GetCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.GetCustomApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListCustomApps

> map[string]interface{} ListCustomApps(ctx).Page(page).Execute()

List Custom Apps



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
	page := "1" // string | Optional page number. Used when results exceed pagination threshold. A hard upper <code>limit</code> is set at 300 device records returned per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.ListCustomApps(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.ListCustomApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.ListCustomApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number. Used when results exceed pagination threshold. A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. | 

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


## UpdateCustomApp

> map[string]interface{} UpdateCustomApp(ctx, libraryItemId).Name(name).Active(active).Execute()

Update Custom App



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
	name := "name_example" // string | Renaming a Custom App
	active := "active_example" // string | (Optional, default=true) Whether this Custom App is active and installable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAppsAPI.UpdateCustomApp(context.Background(), libraryItemId).Name(name).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.UpdateCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomApp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomAppsAPI.UpdateCustomApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Renaming a Custom App | 
 **active** | **string** | (Optional, default&#x3D;true) Whether this Custom App is active and installable | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCustomApp

> UploadCustomApp(ctx).Body(body).Execute()

Upload Custom App



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomAppsAPI.UploadCustomApp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAppsAPI.UploadCustomApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCustomAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

