# \DeviceInformationAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelLostMode**](DeviceInformationAPI.md#CancelLostMode) | **Delete** /api/v1/devices/{device_id}/details/lostmode | Cancel Lost Mode
[**GetDeviceActivity**](DeviceInformationAPI.md#GetDeviceActivity) | **Get** /api/v1/devices/{device_id}/activity | Get Device Activity
[**GetDeviceApps**](DeviceInformationAPI.md#GetDeviceApps) | **Get** /api/v1/devices/{device_id}/apps | Get Device Apps
[**GetDeviceDetails**](DeviceInformationAPI.md#GetDeviceDetails) | **Get** /api/v1/devices/{device_id}/details | Get Device Details
[**GetDeviceLibraryItems**](DeviceInformationAPI.md#GetDeviceLibraryItems) | **Get** /api/v1/devices/{device_id}/library-items | Get Device Library Items
[**GetDeviceLostModeDetails**](DeviceInformationAPI.md#GetDeviceLostModeDetails) | **Get** /api/v1/devices/{device_id}/details/lostmode | Get Device Lost Mode details
[**GetDeviceParameters**](DeviceInformationAPI.md#GetDeviceParameters) | **Get** /api/v1/devices/{device_id}/parameters | Get Device Parameters
[**GetDeviceStatus**](DeviceInformationAPI.md#GetDeviceStatus) | **Get** /api/v1/devices/{device_id}/status | Get Device Status
[**ListDevices**](DeviceInformationAPI.md#ListDevices) | **Get** /api/v1/devices | List Devices



## CancelLostMode

> CancelLostMode(ctx, deviceId).Execute()

Cancel Lost Mode



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
	r, err := apiClient.DeviceInformationAPI.CancelLostMode(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.CancelLostMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelLostModeRequest struct via the builder pattern


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


## GetDeviceActivity

> map[string]interface{} GetDeviceActivity(ctx, deviceId).Limit(limit).Offset(offset).Execute()

Get Device Activity



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
	limit := "300" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results.
	offset := "0" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceActivity(context.Background(), deviceId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceActivity`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **offset** | **string** | Specify the starting record to return | 

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


## GetDeviceApps

> map[string]interface{} GetDeviceApps(ctx, deviceId).Execute()

Get Device Apps



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceApps(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppsRequest struct via the builder pattern


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


## GetDeviceDetails

> map[string]interface{} GetDeviceDetails(ctx, deviceId).Execute()

Get Device Details



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceDetails(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceDetailsRequest struct via the builder pattern


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


## GetDeviceLibraryItems

> map[string]interface{} GetDeviceLibraryItems(ctx, deviceId).Execute()

Get Device Library Items



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceLibraryItems(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceLibraryItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLibraryItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceLibraryItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLibraryItemsRequest struct via the builder pattern


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


## GetDeviceLostModeDetails

> map[string]interface{} GetDeviceLostModeDetails(ctx, deviceId).Execute()

Get Device Lost Mode details



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceLostModeDetails(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceLostModeDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLostModeDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceLostModeDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLostModeDetailsRequest struct via the builder pattern


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


## GetDeviceParameters

> map[string]interface{} GetDeviceParameters(ctx, deviceId).Execute()

Get Device Parameters



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceParameters(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceParameters`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceParametersRequest struct via the builder pattern


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


## GetDeviceStatus

> map[string]interface{} GetDeviceStatus(ctx, deviceId).Execute()

Get Device Status



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
	resp, r, err := apiClient.DeviceInformationAPI.GetDeviceStatus(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.GetDeviceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.GetDeviceStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceStatusRequest struct via the builder pattern


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


## ListDevices

> map[string]interface{} ListDevices(ctx).Limit(limit).AssetTag(assetTag).BlueprintId(blueprintId).DeviceId(deviceId).DeviceName(deviceName).FilevaultEnabled(filevaultEnabled).MacAddress(macAddress).Model(model).Ordering(ordering).OsVersion(osVersion).Platform(platform).SerialNumber(serialNumber).TagName(tagName).TagNameIn(tagNameIn).TagId(tagId).TagIdIn(tagIdIn).User(user).UserEmail(userEmail).UserId(userId).UserName(userName).Offset(offset).Execute()

List Devices



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
	limit := "300" // string | A hard upper <code>limit</code> is set at 300 device records returned per request. If more device records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. Additionally, parameter queries can be added to a request to limit the results.
	assetTag := "23245" // string |  (optional)
	blueprintId := "91f97957-2353-4f86-a1ab-64d2b044a596" // string | Return results &quot;containing&quot; the specified blueprint id (optional)
	deviceId := "2cfeb3ac-3b5d-423e-bcff-e2676a3a32da" // string |  (optional)
	deviceName := "Johnny's MacBook Pro" // string |  (optional)
	filevaultEnabled := "true" // string | <p>Query for devices that either have FileVault on (true) or off (false). This parameter only applies to macOS. </p> <p>An empty list <code>[]</code> will be returned if no devices are found with the given parameter value.</p> (optional)
	macAddress := "00:0c:29:05:43:b6" // string | Search for a specific device by MAC address (optional)
	model := "MacBook Air (M1, 2020)" // string | Return model results &quot;containing&quot; the specified model string. (optional)
	ordering := "device_id" // string | <p>The <code>ordering</code> parameter can be used to define how the device records are ordered in the response. Prepending a dash (-) to the parameter value will reverse the order of the returned results.</p> <p><code>?ordering=-serial_number</code> will order the response by serial_number in descending order.</p> <p><strong>Possible values</strong></p> <ul> <li><code>asset_tag</code></li> <li><code>blueprint_id</code></li> <li><code>device_id</code></li> <li><code>device_name</code></li> <li><code>last_check_in</code> - agent checkin</li> <li><code>model</code></li> <li><code>platform</code></li> <li><code>os_version</code></li> <li><code>serial_number</code></li> <li><code>user</code></li> </ul> <p>Additionally, multiple values can be combined in a comma separated list to further customize the ordering of the response.</p> <p><code>?ordering=serial_number,platform</code></p> (optional)
	osVersion := "13.2.3" // string | Return all device records with the specified OS version (optional)
	platform := "iPad" // string | Return all records matching a specific platform. Possible values:<code>Mac</code>, <code>iPad</code>, <code>iPhone</code>, <code>AppleTV</code> (optional)
	serialNumber := "VMC5qeJ5pDkp" // string | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. (optional)
	tagName := "accuhive_01" // string | Return results for given tag name. Case sensitive. (optional)
	tagNameIn := "accuhive_01, accuhive_02" // string | Return results for given tag names separate by commas. Case sensitive. (optional)
	tagId := "tagId_example" // string | Search for a tag by its ID. Case sensitive. (optional)
	tagIdIn := "tagIdIn_example" // string | Return results for given tag IDs separated by commas. Case sensitive. (optional)
	user := "Art Vandelay" // string | Return results &quot;containing&quot; the user name (optional)
	userEmail := "someUser@Kandji.io" // string | Return results &quot;containing&quot; search on email address (optional)
	userId := "1" // string | &quot;exact&quot; match on kandji user ID number (optional)
	userName := "Vandelay" // string | Return results &quot;containing&quot; the assigned user Display Name (optional)
	offset := "0" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceInformationAPI.ListDevices(context.Background()).Limit(limit).AssetTag(assetTag).BlueprintId(blueprintId).DeviceId(deviceId).DeviceName(deviceName).FilevaultEnabled(filevaultEnabled).MacAddress(macAddress).Model(model).Ordering(ordering).OsVersion(osVersion).Platform(platform).SerialNumber(serialNumber).TagName(tagName).TagNameIn(tagNameIn).TagId(tagId).TagIdIn(tagIdIn).User(user).UserEmail(userEmail).UserId(userId).UserName(userName).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceInformationAPI.ListDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDevices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DeviceInformationAPI.ListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 300 device records returned per request. If more device records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. Additionally, parameter queries can be added to a request to limit the results. | 
 **assetTag** | **string** |  | 
 **blueprintId** | **string** | Return results &amp;quot;containing&amp;quot; the specified blueprint id | 
 **deviceId** | **string** |  | 
 **deviceName** | **string** |  | 
 **filevaultEnabled** | **string** | &lt;p&gt;Query for devices that either have FileVault on (true) or off (false). This parameter only applies to macOS. &lt;/p&gt; &lt;p&gt;An empty list &lt;code&gt;[]&lt;/code&gt; will be returned if no devices are found with the given parameter value.&lt;/p&gt; | 
 **macAddress** | **string** | Search for a specific device by MAC address | 
 **model** | **string** | Return model results &amp;quot;containing&amp;quot; the specified model string. | 
 **ordering** | **string** | &lt;p&gt;The &lt;code&gt;ordering&lt;/code&gt; parameter can be used to define how the device records are ordered in the response. Prepending a dash (-) to the parameter value will reverse the order of the returned results.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?ordering&#x3D;-serial_number&lt;/code&gt; will order the response by serial_number in descending order.&lt;/p&gt; &lt;p&gt;&lt;strong&gt;Possible values&lt;/strong&gt;&lt;/p&gt; &lt;ul&gt; &lt;li&gt;&lt;code&gt;asset_tag&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;blueprint_id&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;device_id&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;device_name&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;last_check_in&lt;/code&gt; - agent checkin&lt;/li&gt; &lt;li&gt;&lt;code&gt;model&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;platform&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;os_version&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;serial_number&lt;/code&gt;&lt;/li&gt; &lt;li&gt;&lt;code&gt;user&lt;/code&gt;&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Additionally, multiple values can be combined in a comma separated list to further customize the ordering of the response.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?ordering&#x3D;serial_number,platform&lt;/code&gt;&lt;/p&gt; | 
 **osVersion** | **string** | Return all device records with the specified OS version | 
 **platform** | **string** | Return all records matching a specific platform. Possible values:&lt;code&gt;Mac&lt;/code&gt;, &lt;code&gt;iPad&lt;/code&gt;, &lt;code&gt;iPhone&lt;/code&gt;, &lt;code&gt;AppleTV&lt;/code&gt; | 
 **serialNumber** | **string** | Search for a specific device by Serial Number. If partial serial number is provided in the query, all device containing the partial string will be returned. | 
 **tagName** | **string** | Return results for given tag name. Case sensitive. | 
 **tagNameIn** | **string** | Return results for given tag names separate by commas. Case sensitive. | 
 **tagId** | **string** | Search for a tag by its ID. Case sensitive. | 
 **tagIdIn** | **string** | Return results for given tag IDs separated by commas. Case sensitive. | 
 **user** | **string** | Return results &amp;quot;containing&amp;quot; the user name | 
 **userEmail** | **string** | Return results &amp;quot;containing&amp;quot; search on email address | 
 **userId** | **string** | &amp;quot;exact&amp;quot; match on kandji user ID number | 
 **userName** | **string** | Return results &amp;quot;containing&amp;quot; the assigned user Display Name | 
 **offset** | **string** | Specify the starting record to return | 

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

