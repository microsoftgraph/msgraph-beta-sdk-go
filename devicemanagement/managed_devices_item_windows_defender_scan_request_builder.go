package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemWindowsDefenderScanRequestBuilder provides operations to call the windowsDefenderScan method.
type ManagedDevicesItemWindowsDefenderScanRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemWindowsDefenderScanRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemWindowsDefenderScanRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemWindowsDefenderScanRequestBuilderInternal instantiates a new WindowsDefenderScanRequestBuilder and sets the default values.
func NewManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    m := &ManagedDevicesItemWindowsDefenderScanRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/windowsDefenderScan", pathParameters),
    }
    return m
}
// NewManagedDevicesItemWindowsDefenderScanRequestBuilder instantiates a new WindowsDefenderScanRequestBuilder and sets the default values.
func NewManagedDevicesItemWindowsDefenderScanRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemWindowsDefenderScanRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action windowsDefenderScan
func (m *ManagedDevicesItemWindowsDefenderScanRequestBuilder) Post(ctx context.Context, body ManagedDevicesItemWindowsDefenderScanPostRequestBodyable, requestConfiguration *ManagedDevicesItemWindowsDefenderScanRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action windowsDefenderScan
func (m *ManagedDevicesItemWindowsDefenderScanRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedDevicesItemWindowsDefenderScanPostRequestBodyable, requestConfiguration *ManagedDevicesItemWindowsDefenderScanRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedDevicesItemWindowsDefenderScanRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemWindowsDefenderScanRequestBuilder) {
    return NewManagedDevicesItemWindowsDefenderScanRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
