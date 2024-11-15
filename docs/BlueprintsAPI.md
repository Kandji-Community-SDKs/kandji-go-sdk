# \BlueprintsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignLibraryItem**](BlueprintsAPI.md#AssignLibraryItem) | **Post** /api/v1/blueprints/{blueprint_id}/assign-library-item | Assign Library Item
[**CreateBlueprint**](BlueprintsAPI.md#CreateBlueprint) | **Post** /api/v1/blueprints | Create Blueprint
[**DeleteBlueprint**](BlueprintsAPI.md#DeleteBlueprint) | **Delete** /api/v1/blueprints/{blueprint_id} | Delete Blueprint
[**GetBlueprint**](BlueprintsAPI.md#GetBlueprint) | **Get** /api/v1/blueprints/{blueprint_id} | Get Blueprint
[**GetBlueprintTemplates**](BlueprintsAPI.md#GetBlueprintTemplates) | **Get** /api/v1/blueprints/templates/ | Get Blueprint Templates
[**GetManualEnrollmentProfile**](BlueprintsAPI.md#GetManualEnrollmentProfile) | **Get** /api/v1/blueprints/{blueprint_id}/ota-enrollment-profile | Get Manual Enrollment Profile
[**ListBlueprints**](BlueprintsAPI.md#ListBlueprints) | **Get** /api/v1/blueprints | List Blueprints
[**ListLibraryItems**](BlueprintsAPI.md#ListLibraryItems) | **Get** /api/v1/blueprints/{blueprint_id}/list-library-items | List Library Items
[**RemoveLibraryItem**](BlueprintsAPI.md#RemoveLibraryItem) | **Post** /api/v1/blueprints/{blueprint_id}/remove-library-item | Remove Library Item
[**UpdateBlueprint**](BlueprintsAPI.md#UpdateBlueprint) | **Patch** /api/v1/blueprints/{blueprint_id} | Update Blueprint



## AssignLibraryItem

> map[string]interface{} AssignLibraryItem(ctx, blueprintId).Body(body).Execute()

Assign Library Item



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
	blueprintId := "blueprintId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.AssignLibraryItem(context.Background(), blueprintId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.AssignLibraryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignLibraryItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.AssignLibraryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignLibraryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBlueprint

> map[string]interface{} CreateBlueprint(ctx).Name(name).EnrollmentCodeIsActive(enrollmentCodeIsActive).EnrollmentCodeCode(enrollmentCodeCode).SourceType(sourceType).SourceId(sourceId).Type_(type_).Execute()

Create Blueprint



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
	name := "name_example" // string | (required) Set the name of the Blueprint. The name provided must be unique.
	enrollmentCodeIsActive := "enrollmentCodeIsActive_example" // string | (required) Enable or Disable the Blueprint for manual device enrollment from the enrollment portal
	enrollmentCodeCode := "enrollmentCodeCode_example" // string | Optionally, set the enrollment code of the Blueprint. This key is not required. If an enrollment code is not supplied in the payload body, it will be randomly generated. The enrollment code will be returned in the response and visible in the Web app.
	sourceType := "sourceType_example" // string | Set the source to create the blueprint from. Possible options: <code>template</code> and <code>blueprint</code>.
	sourceId := "sourceId_example" // string | Set either the source template ID, or the source Blueprint ID to clone an existing template or blueprint.
	type_ := "type__example" // string | Choose the type of blueprint to create. Options: <code>classic</code> or <code>map</code>

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.CreateBlueprint(context.Background()).Name(name).EnrollmentCodeIsActive(enrollmentCodeIsActive).EnrollmentCodeCode(enrollmentCodeCode).SourceType(sourceType).SourceId(sourceId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.CreateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlueprint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.CreateBlueprint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (required) Set the name of the Blueprint. The name provided must be unique. | 
 **enrollmentCodeIsActive** | **string** | (required) Enable or Disable the Blueprint for manual device enrollment from the enrollment portal | 
 **enrollmentCodeCode** | **string** | Optionally, set the enrollment code of the Blueprint. This key is not required. If an enrollment code is not supplied in the payload body, it will be randomly generated. The enrollment code will be returned in the response and visible in the Web app. | 
 **sourceType** | **string** | Set the source to create the blueprint from. Possible options: &lt;code&gt;template&lt;/code&gt; and &lt;code&gt;blueprint&lt;/code&gt;. | 
 **sourceId** | **string** | Set either the source template ID, or the source Blueprint ID to clone an existing template or blueprint. | 
 **type_** | **string** | Choose the type of blueprint to create. Options: &lt;code&gt;classic&lt;/code&gt; or &lt;code&gt;map&lt;/code&gt; | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlueprint

> DeleteBlueprint(ctx, blueprintId).Execute()

Delete Blueprint



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
	blueprintId := "blueprintId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.DeleteBlueprint(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.DeleteBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprint

> GetBlueprint(ctx, blueprintId).Execute()

Get Blueprint



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
	blueprintId := "blueprintId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.GetBlueprint(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.GetBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlueprintTemplates

> GetBlueprintTemplates(ctx).Limit(limit).Offset(offset).Execute()

Get Blueprint Templates

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
	limit := "100" // string | Number of results to return per page. (optional)
	offset := "400" // string | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlueprintsAPI.GetBlueprintTemplates(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.GetBlueprintTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlueprintTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Number of results to return per page. | 
 **offset** | **string** | The initial index from which to return the results. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualEnrollmentProfile

> string GetManualEnrollmentProfile(ctx, blueprintId).Sso(sso).Execute()

Get Manual Enrollment Profile



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
	blueprintId := "blueprintId_example" // string | 
	sso := "true" // string | Use the <code>sso</code> query parameter, set to <code>true</code>, to return a URL instead of the manual enrollment profile. This parameter should only be used for blueprints in which &quot;Require Authentication&quot; is configured for Manual Enrollment. The returned URL must be used to authenticate via SSO to receive an enrollment profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.GetManualEnrollmentProfile(context.Background(), blueprintId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.GetManualEnrollmentProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManualEnrollmentProfile`: string
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.GetManualEnrollmentProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualEnrollmentProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sso** | **string** | Use the &lt;code&gt;sso&lt;/code&gt; query parameter, set to &lt;code&gt;true&lt;/code&gt;, to return a URL instead of the manual enrollment profile. This parameter should only be used for blueprints in which &amp;quot;Require Authentication&amp;quot; is configured for Manual Enrollment. The returned URL must be used to authenticate via SSO to receive an enrollment profile. | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-apple-aspen-config

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlueprints

> map[string]interface{} ListBlueprints(ctx).Id(id).IdIn(idIn).Name(name).Limit(limit).Offset(offset).Execute()

List Blueprints



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
	id := "97e4e175-1631-43f6-a02b-33fd1c748ab8" // string | Look up a specific Blueprint by its ID (optional)
	idIn := "11f4eb9a-10ed-4c3d-a7c1-fb87f95743fb,6391086e-85a1-4820-813c-f9c75025fff4" // string | Specify a list of Blueprint IDs to limit the results to.  Multiple values may be separated by commas. There is a double underscore (<code>__</code>) between id and in (optional)
	name := "testing_blueprint" // string | Return Blueprint names &quot;containing&quot; the specified search string. (optional)
	limit := "100" // string | Number of results to return per page. (optional)
	offset := "400" // string | The initial index from which to return the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.ListBlueprints(context.Background()).Id(id).IdIn(idIn).Name(name).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.ListBlueprints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlueprints`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.ListBlueprints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlueprintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Look up a specific Blueprint by its ID | 
 **idIn** | **string** | Specify a list of Blueprint IDs to limit the results to.  Multiple values may be separated by commas. There is a double underscore (&lt;code&gt;__&lt;/code&gt;) between id and in | 
 **name** | **string** | Return Blueprint names &amp;quot;containing&amp;quot; the specified search string. | 
 **limit** | **string** | Number of results to return per page. | 
 **offset** | **string** | The initial index from which to return the results. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLibraryItems

> map[string]interface{} ListLibraryItems(ctx, blueprintId).Execute()

List Library Items



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
	blueprintId := "blueprintId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.ListLibraryItems(context.Background(), blueprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.ListLibraryItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLibraryItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.ListLibraryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLibraryItemsRequest struct via the builder pattern


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


## RemoveLibraryItem

> map[string]interface{} RemoveLibraryItem(ctx, blueprintId).Body(body).Execute()

Remove Library Item



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
	blueprintId := "blueprintId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.RemoveLibraryItem(context.Background(), blueprintId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.RemoveLibraryItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLibraryItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.RemoveLibraryItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLibraryItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlueprint

> map[string]interface{} UpdateBlueprint(ctx, blueprintId).Name(name).Description(description).EnrollmentCodeCode(enrollmentCodeCode).EnrollmentCodeIsActive(enrollmentCodeIsActive).Execute()

Update Blueprint



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
	blueprintId := "blueprintId_example" // string | 
	name := "name_example" // string | Update the name of the Blueprint
	description := "description_example" // string | Update the description of the Blueprint 
	enrollmentCodeCode := "enrollmentCodeCode_example" // string | Update the enrollment code of the Blueprint 
	enrollmentCodeIsActive := "enrollmentCodeIsActive_example" // string | Disable the Blueprint for manual device enrollment from the enrollment portal.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlueprintsAPI.UpdateBlueprint(context.Background(), blueprintId).Name(name).Description(description).EnrollmentCodeCode(enrollmentCodeCode).EnrollmentCodeIsActive(enrollmentCodeIsActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlueprintsAPI.UpdateBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBlueprint`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BlueprintsAPI.UpdateBlueprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blueprintId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlueprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Update the name of the Blueprint | 
 **description** | **string** | Update the description of the Blueprint  | 
 **enrollmentCodeCode** | **string** | Update the enrollment code of the Blueprint  | 
 **enrollmentCodeIsActive** | **string** | Disable the Blueprint for manual device enrollment from the enrollment portal. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

