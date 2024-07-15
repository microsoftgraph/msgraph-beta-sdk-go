package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder provides operations to call the windowsPrivacyAccessControls method.
type DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderInternal instantiates a new DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) {
    m := &DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/windowsPrivacyAccessControls", pathParameters),
    }
    return m
}
// NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder instantiates a new DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action windowsPrivacyAccessControls
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) Post(ctx context.Context, body DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyable, requestConfiguration *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action windowsPrivacyAccessControls
// returns a *RequestInformation when successful
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceConfigurationsItemWindowsPrivacyAccessControlsPostRequestBodyable, requestConfiguration *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder when successful
func (m *DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) WithUrl(rawUrl string)(*DeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) {
    return NewDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
