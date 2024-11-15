# \AutomatedDeviceEnrollmentIntegrationsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#CreateAdeIntegration) | **Post** /api/v1/integrations/apple/ade/ | Create ADE integration
[**DeleteAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#DeleteAdeIntegration) | **Delete** /api/v1/integrations/apple/ade/{ade_token_id} | Delete ADE integration
[**DownloadAdePublicKey**](AutomatedDeviceEnrollmentIntegrationsAPI.md#DownloadAdePublicKey) | **Get** /api/v1/integrations/apple/ade/public_key/ | Download ADE public key
[**GetAdeDevice**](AutomatedDeviceEnrollmentIntegrationsAPI.md#GetAdeDevice) | **Get** /api/v1/integrations/apple/ade/devices/{device_id} | Get ADE device
[**GetAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#GetAdeIntegration) | **Get** /api/v1/integrations/apple/ade/{ade_token_id} | Get ADE integration
[**ListAdeDevices**](AutomatedDeviceEnrollmentIntegrationsAPI.md#ListAdeDevices) | **Get** /api/v1/integrations/apple/ade/devices | List ADE devices
[**ListAdeIntegrations**](AutomatedDeviceEnrollmentIntegrationsAPI.md#ListAdeIntegrations) | **Get** /api/v1/integrations/apple/ade | List ADE integrations
[**ListDevicesAssociatedToAdeToken**](AutomatedDeviceEnrollmentIntegrationsAPI.md#ListDevicesAssociatedToAdeToken) | **Get** /api/v1/integrations/apple/ade/{ade_token_id}/devices | List devices associated to ADE token
[**RenewAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#RenewAdeIntegration) | **Post** /api/v1/integrations/apple/ade/{ade_token_id}/renew | Renew ADE integration
[**UpdateAdeDevice**](AutomatedDeviceEnrollmentIntegrationsAPI.md#UpdateAdeDevice) | **Patch** /api/v1/integrations/apple/ade/devices/{device_id} | Update ADE device
[**UpdateAdeIntegration**](AutomatedDeviceEnrollmentIntegrationsAPI.md#UpdateAdeIntegration) | **Patch** /api/v1/integrations/apple/ade/{ade_token_id} | Update ADE integration



## CreateAdeIntegration

> map[string]interface{} CreateAdeIntegration(ctx).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()

Create ADE integration



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
	phone := "phone_example" // string | 
	email := "email_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.CreateAdeIntegration(context.Background()).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.CreateAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdeIntegration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.CreateAdeIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintId** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **file** | ***os.File** | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdeIntegration

> DeleteAdeIntegration(ctx, adeTokenId).Execute()

Delete ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.DeleteAdeIntegration(context.Background(), adeTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.DeleteAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAdePublicKey

> string DownloadAdePublicKey(ctx).Execute()

Download ADE public key



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.DownloadAdePublicKey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.DownloadAdePublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadAdePublicKey`: string
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.DownloadAdePublicKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAdePublicKeyRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-x509-ca-cert

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdeDevice

> map[string]interface{} GetAdeDevice(ctx, deviceId).Execute()

Get ADE device



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
	deviceId := "deviceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdeDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdeDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAdeIntegration

> GetAdeIntegration(ctx, adeTokenId).Execute()

Get ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeIntegration(context.Background(), adeTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.GetAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdeIntegrationRequest struct via the builder pattern


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


## ListAdeDevices

> map[string]interface{} ListAdeDevices(ctx).BlueprintId(blueprintId).UserId(userId).DepAccount(depAccount).DeviceFamily(deviceFamily).Model(model).Os(os).ProfileStatus(profileStatus).SerialNumber(serialNumber).Page(page).Execute()

List ADE devices



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
	blueprintId := "fce0cc58-caa5-40d2-a0d7-a0b257127ec5" // string | Return results &quot;containing&quot; the specified blueprint id (optional)
	userId := "8136" // string | &quot;exact&quot; match on kandji user ID number (optional)
	depAccount := "depAccount_example" // string | The ADE token UUID (optional)
	deviceFamily := "deviceFamily_example" // string | Mac, iPhone, iPad, AppleTV, iPod (optional)
	model := "MacBook Air" // string | Return model results &quot;containing&quot; the specified model string. - &quot;iPad (8th Generation)&quot;, &quot;MacBook Air&quot; (optional)
	os := "os_example" // string | OSX, iOS, tvOS (optional)
	profileStatus := "profileStatus_example" // string | The automated device enrollment profile assignment status - assigned, empty, pushed, removed (optional)
	serialNumber := "serialNumber_example" // string | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. (optional)
	page := "1" // string | Use the <code>page</code> parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the <code>next</code> and <code>previous</code> key attributes in the response can be used to request the next page of results or return to the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeDevices(context.Background()).BlueprintId(blueprintId).UserId(userId).DepAccount(depAccount).DeviceFamily(deviceFamily).Model(model).Os(os).ProfileStatus(profileStatus).SerialNumber(serialNumber).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAdeDevices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAdeDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blueprintId** | **string** | Return results &amp;quot;containing&amp;quot; the specified blueprint id | 
 **userId** | **string** | &amp;quot;exact&amp;quot; match on kandji user ID number | 
 **depAccount** | **string** | The ADE token UUID | 
 **deviceFamily** | **string** | Mac, iPhone, iPad, AppleTV, iPod | 
 **model** | **string** | Return model results &amp;quot;containing&amp;quot; the specified model string. - &amp;quot;iPad (8th Generation)&amp;quot;, &amp;quot;MacBook Air&amp;quot; | 
 **os** | **string** | OSX, iOS, tvOS | 
 **profileStatus** | **string** | The automated device enrollment profile assignment status - assigned, empty, pushed, removed | 
 **serialNumber** | **string** | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. | 
 **page** | **string** | Use the &lt;code&gt;page&lt;/code&gt; parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the &lt;code&gt;next&lt;/code&gt; and &lt;code&gt;previous&lt;/code&gt; key attributes in the response can be used to request the next page of results or return to the previous page. | 

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


## ListAdeIntegrations

> ListAdeIntegrations(ctx).Execute()

List ADE integrations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.ListAdeIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAdeIntegrationsRequest struct via the builder pattern


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


## ListDevicesAssociatedToAdeToken

> map[string]interface{} ListDevicesAssociatedToAdeToken(ctx, adeTokenId).Page(page).Execute()

List devices associated to ADE token



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
	adeTokenId := "adeTokenId_example" // string | 
	page := "1" // string | Use the <code>page</code> parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the <code>next</code> and <code>previous</code> key attributes in the response can be used to request the next page of results or return to the previous page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.ListDevicesAssociatedToAdeToken(context.Background(), adeTokenId).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.ListDevicesAssociatedToAdeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevicesAssociatedToAdeToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.ListDevicesAssociatedToAdeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesAssociatedToAdeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Use the &lt;code&gt;page&lt;/code&gt; parameter to page through results or to request a specific page. By default, if a page is not specified, page 1 is returned. Note: 300 device records are returned per page of results. Alternatively, the &lt;code&gt;next&lt;/code&gt; and &lt;code&gt;previous&lt;/code&gt; key attributes in the response can be used to request the next page of results or return to the previous page. | 

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


## RenewAdeIntegration

> RenewAdeIntegration(ctx, adeTokenId).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()

Renew ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 
	blueprintId := "blueprintId_example" // string | 
	phone := "phone_example" // string | 
	email := "email_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.RenewAdeIntegration(context.Background(), adeTokenId).BlueprintId(blueprintId).Phone(phone).Email(email).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.RenewAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewAdeIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blueprintId** | **string** |  | 
 **phone** | **string** |  | 
 **email** | **string** |  | 
 **file** | ***os.File** | This is the MDM server token file(.p7m) download from ABM. Once downloaded from ABM, the file can be uploaded via API. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdeDevice

> map[string]interface{} UpdateAdeDevice(ctx, deviceId).Body(body).Execute()

Update ADE device



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
	deviceId := "deviceId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeDevice(context.Background(), deviceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdeDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdeDeviceRequest struct via the builder pattern


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


## UpdateAdeIntegration

> UpdateAdeIntegration(ctx, adeTokenId).Body(body).Execute()

Update ADE integration



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
	adeTokenId := "adeTokenId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeIntegration(context.Background(), adeTokenId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomatedDeviceEnrollmentIntegrationsAPI.UpdateAdeIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adeTokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdeIntegrationRequest struct via the builder pattern


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

