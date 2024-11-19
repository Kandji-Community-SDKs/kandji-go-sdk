# \CustomScriptsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomScript**](CustomScriptsAPI.md#CreateCustomScript) | **Post** /api/v1/library/custom-scripts | Create Custom Script
[**DeleteCustomScript**](CustomScriptsAPI.md#DeleteCustomScript) | **Delete** /api/v1/library/custom-scripts/{library_item_id} | Delete Custom Script
[**GetCustomScript**](CustomScriptsAPI.md#GetCustomScript) | **Get** /api/v1/library/custom-scripts/{library_item_id} | Get Custom Script
[**ListCustomScripts**](CustomScriptsAPI.md#ListCustomScripts) | **Get** /api/v1/library/custom-scripts | List Custom Scripts
[**UpdateCustomScript**](CustomScriptsAPI.md#UpdateCustomScript) | **Patch** /api/v1/library/custom-scripts/{library_item_id} | Update Custom Script



## CreateCustomScript

> map[string]interface{} CreateCustomScript(ctx).Body(body).Execute()

Create Custom Script



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
	resp, r, err := apiClient.CustomScriptsAPI.CreateCustomScript(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.CreateCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.CreateCustomScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomScript

> DeleteCustomScript(ctx, libraryItemId).Execute()

Delete Custom Script



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
	r, err := apiClient.CustomScriptsAPI.DeleteCustomScript(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.DeleteCustomScript``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCustomScriptRequest struct via the builder pattern


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


## GetCustomScript

> map[string]interface{} GetCustomScript(ctx, libraryItemId).Execute()

Get Custom Script



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
	resp, r, err := apiClient.CustomScriptsAPI.GetCustomScript(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.GetCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.GetCustomScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomScriptRequest struct via the builder pattern


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


## ListCustomScripts

> map[string]interface{} ListCustomScripts(ctx).Page(page).Execute()

List Custom Scripts



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
	page := "1" // string | Optional page number. Used when results exceed pagination threshold. A hard upper limit is set at 300 device records returned per request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomScriptsAPI.ListCustomScripts(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.ListCustomScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomScripts`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.ListCustomScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number. Used when results exceed pagination threshold. A hard upper limit is set at 300 device records returned per request. | 

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


## UpdateCustomScript

> map[string]interface{} UpdateCustomScript(ctx, libraryItemId).Body(body).Execute()

Update Custom Script



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomScriptsAPI.UpdateCustomScript(context.Background(), libraryItemId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomScriptsAPI.UpdateCustomScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomScriptsAPI.UpdateCustomScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

