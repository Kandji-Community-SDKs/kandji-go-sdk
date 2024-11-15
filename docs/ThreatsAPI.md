# \ThreatsAPI

All URIs are relative to *https://&lt;sub_domain&gt;.api.kandji.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetThreatDetails**](ThreatsAPI.md#GetThreatDetails) | **Get** /api/v1/threat-details | Get Threat Details



## GetThreatDetails

> map[string]interface{} GetThreatDetails(ctx).Classification(classification).DateRange(dateRange).DeviceId(deviceId).Status(status).SortBy(sortBy).Term(term).Limit(limit).Offset(offset).Execute()

Get Threat Details



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
	classification := "malware" // string | Return all records matching a specified classification. The following classification options are available: <code>malware</code> and <code>pup</code>. Leave this parameter empty to return all classification types. (optional)
	dateRange := "7" // string | Return all records within a specified number of days. Any positive number of days may be specified. Examples: <code>7</code>, <code>30</code>, <code>60</code>, <code>90</code>, <code>180</code>, or <code>365</code>. (optional)
	deviceId := "15fcec08-xxxx-xxxx-xxxx-7c2f950910eb" // string |  (optional)
	status := "quarantined" // string | Return all records matching a specified status. The following status options are available: <code>quarantined</code>, <code>not_quarantined</code>, or <code>released</code>. Leave this parameter empty to return all status types. (optional)
	sortBy := "status" // string | <p>Results can be sorted with the following options: </p> <ul> <li>threat_name</li> <li>classification</li> <li>device_name</li> <li>process_name</li> <li>process_owner</li> <li>detection_date</li> <li>status</li> </ul> <p>Prepending a dash (-) to the parameter value will reverse the order of the returned results.</p> <p><code>?sort_by=-device_name</code> will order the response by device_name in descending order.</p> (optional)
	term := "Chrome" // string | <p>Search term to filter threat results.</p> <p>The response will include anything matching the following fields: <code>device_name</code>, <code>file_hash</code>, and <code>file_path</code>.</p> <p>So if you search for <code>bad file</code>, the results will include anywhere <code>bad file</code> exists in the three fields above.</p> (optional)
	limit := "1000" // string | <p>A hard upper <code>limit</code> is set at 1000 records returned per request. If more records are expected, pagination should be used using the <code>limit</code> and <code>offset</code> parameters. </p> <p>Additionally, parameter queries can be added to a request to limit the results.</p> (optional)
	offset := "1" // string | Specify the starting record to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThreatsAPI.GetThreatDetails(context.Background()).Classification(classification).DateRange(dateRange).DeviceId(deviceId).Status(status).SortBy(sortBy).Term(term).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatsAPI.GetThreatDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThreatDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ThreatsAPI.GetThreatDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetThreatDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classification** | **string** | Return all records matching a specified classification. The following classification options are available: &lt;code&gt;malware&lt;/code&gt; and &lt;code&gt;pup&lt;/code&gt;. Leave this parameter empty to return all classification types. | 
 **dateRange** | **string** | Return all records within a specified number of days. Any positive number of days may be specified. Examples: &lt;code&gt;7&lt;/code&gt;, &lt;code&gt;30&lt;/code&gt;, &lt;code&gt;60&lt;/code&gt;, &lt;code&gt;90&lt;/code&gt;, &lt;code&gt;180&lt;/code&gt;, or &lt;code&gt;365&lt;/code&gt;. | 
 **deviceId** | **string** |  | 
 **status** | **string** | Return all records matching a specified status. The following status options are available: &lt;code&gt;quarantined&lt;/code&gt;, &lt;code&gt;not_quarantined&lt;/code&gt;, or &lt;code&gt;released&lt;/code&gt;. Leave this parameter empty to return all status types. | 
 **sortBy** | **string** | &lt;p&gt;Results can be sorted with the following options: &lt;/p&gt; &lt;ul&gt; &lt;li&gt;threat_name&lt;/li&gt; &lt;li&gt;classification&lt;/li&gt; &lt;li&gt;device_name&lt;/li&gt; &lt;li&gt;process_name&lt;/li&gt; &lt;li&gt;process_owner&lt;/li&gt; &lt;li&gt;detection_date&lt;/li&gt; &lt;li&gt;status&lt;/li&gt; &lt;/ul&gt; &lt;p&gt;Prepending a dash (-) to the parameter value will reverse the order of the returned results.&lt;/p&gt; &lt;p&gt;&lt;code&gt;?sort_by&#x3D;-device_name&lt;/code&gt; will order the response by device_name in descending order.&lt;/p&gt; | 
 **term** | **string** | &lt;p&gt;Search term to filter threat results.&lt;/p&gt; &lt;p&gt;The response will include anything matching the following fields: &lt;code&gt;device_name&lt;/code&gt;, &lt;code&gt;file_hash&lt;/code&gt;, and &lt;code&gt;file_path&lt;/code&gt;.&lt;/p&gt; &lt;p&gt;So if you search for &lt;code&gt;bad file&lt;/code&gt;, the results will include anywhere &lt;code&gt;bad file&lt;/code&gt; exists in the three fields above.&lt;/p&gt; | 
 **limit** | **string** | &lt;p&gt;A hard upper &lt;code&gt;limit&lt;/code&gt; is set at 1000 records returned per request. If more records are expected, pagination should be used using the &lt;code&gt;limit&lt;/code&gt; and &lt;code&gt;offset&lt;/code&gt; parameters. &lt;/p&gt; &lt;p&gt;Additionally, parameter queries can be added to a request to limit the results.&lt;/p&gt; | 
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

