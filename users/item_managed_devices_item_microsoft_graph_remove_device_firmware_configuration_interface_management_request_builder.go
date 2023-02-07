package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
type ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal instantiates a new MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    m := &ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.removeDeviceFirmwareConfigurationInterfaceManagement";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder instantiates a new MicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove device from Device Firmware Configuration Interface management
func (m *ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation remove device from Device Firmware Configuration Interface management
func (m *ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemMicrosoftGraphRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
