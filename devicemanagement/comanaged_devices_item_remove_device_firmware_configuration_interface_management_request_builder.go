package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder provides operations to call the removeDeviceFirmwareConfigurationInterfaceManagement method.
type ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal instantiates a new ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    m := &ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/{managedDevice%2Did}/removeDeviceFirmwareConfigurationInterfaceManagement", pathParameters),
    }
    return m
}
// NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder instantiates a new ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder and sets the default values.
func NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove device from Device Firmware Configuration Interface management
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) Post(ctx context.Context, requestConfiguration *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(error) {
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
// returns a *RequestInformation when successful
func (m *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder when successful
func (m *ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder) {
    return NewComanagedDevicesItemRemoveDeviceFirmwareConfigurationInterfaceManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
