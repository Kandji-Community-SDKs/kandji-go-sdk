# \CustomProfilesAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomProfile**](CustomProfilesAPI.md#CreateCustomProfile) | **Post** /api/v1/library/custom-profiles | Create Custom Profile
[**DeleteCustomProfile**](CustomProfilesAPI.md#DeleteCustomProfile) | **Delete** /api/v1/library/custom-profiles/{library_item_id} | Delete Custom Profile
[**GetCustomProfile**](CustomProfilesAPI.md#GetCustomProfile) | **Get** /api/v1/library/custom-profiles/{library_item_id} | Get Custom Profile
[**ListCustomProfiles**](CustomProfilesAPI.md#ListCustomProfiles) | **Get** /api/v1/library/custom-profiles | List Custom Profiles
[**UpdateCustomProfile**](CustomProfilesAPI.md#UpdateCustomProfile) | **Patch** /api/v1/library/custom-profiles/{library_item_id} | Update Custom Profile



## CreateCustomProfile

> map[string]interface{} CreateCustomProfile(ctx).Name(name).File(file).Active(active).Execute()

Create Custom Profile



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
	name := "name_example" // string | (Required) The profile name
	file := os.NewFile(1234, "some_file") // *os.File | (Required) The path to the profile's .mobileconfig file
	active := "active_example" // string | (Optional, default=true) Whether this library item is active

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomProfilesAPI.CreateCustomProfile(context.Background()).Name(name).File(file).Active(active).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.CreateCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.CreateCustomProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (Required) The profile name | 
 **file** | ***os.File** | (Required) The path to the profile&#39;s .mobileconfig file | 
 **active** | **string** | (Optional, default&#x3D;true) Whether this library item is active | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomProfile

> DeleteCustomProfile(ctx, libraryItemId).Execute()

Delete Custom Profile



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
	r, err := apiClient.CustomProfilesAPI.DeleteCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.DeleteCustomProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCustomProfileRequest struct via the builder pattern


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


## GetCustomProfile

> map[string]interface{} GetCustomProfile(ctx, libraryItemId).Execute()

Get Custom Profile



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
	resp, r, err := apiClient.CustomProfilesAPI.GetCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.GetCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.GetCustomProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomProfileRequest struct via the builder pattern


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


## ListCustomProfiles

> map[string]interface{} ListCustomProfiles(ctx).Page(page).Execute()

List Custom Profiles



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
	page := "1" // string | Optional page number (when results exceed pagination threshold) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomProfilesAPI.ListCustomProfiles(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.ListCustomProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.ListCustomProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCustomProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Optional page number (when results exceed pagination threshold) | 

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


## UpdateCustomProfile

> map[string]interface{} UpdateCustomProfile(ctx, libraryItemId).Execute()

Update Custom Profile



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
	resp, r, err := apiClient.CustomProfilesAPI.UpdateCustomProfile(context.Background(), libraryItemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomProfilesAPI.UpdateCustomProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomProfilesAPI.UpdateCustomProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

