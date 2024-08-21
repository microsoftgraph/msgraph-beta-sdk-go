package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
type ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal instantiates a new ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    m := &ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/removeDeviceFirmwareConfigurationInterfaceManagement", pathParameters),
    }
    return m
}
// NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder instantiates a new ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove device from Device Firmware Configuration Interface management
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation remove device from Device Firmware Configuration Interface management
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) WithUrl(rawUrl string)(*ItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewItemManagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
