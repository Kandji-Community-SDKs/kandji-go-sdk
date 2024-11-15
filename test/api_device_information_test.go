/*
Kandji API

Testing DeviceInformationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package kandji_sdk

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Kandji-Community-SDKs/kandji-go-sdk"
)

func Test_kandji_sdk_DeviceInformationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceInformationAPIService CancelLostMode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceInformationAPI.CancelLostMode(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceActivity(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceApps(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceDetails(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceLibraryItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceLibraryItems(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceLostModeDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceLostModeDetails(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceParameters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceParameters(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService GetDeviceStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceInformationAPI.GetDeviceStatus(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceInformationAPIService ListDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeviceInformationAPI.ListDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
