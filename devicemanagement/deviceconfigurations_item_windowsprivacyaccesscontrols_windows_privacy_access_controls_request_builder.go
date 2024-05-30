package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder provides operations to call the windowsPrivacyAccessControls method.
type DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderInternal instantiates a new DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) {
    m := &DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/windowsPrivacyAccessControls", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder instantiates a new DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action windowsPrivacyAccessControls
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) Post(ctx context.Context, body DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsPostRequestBodyable, requestConfiguration *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(error) {
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
func (m *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) ToPostRequestInformation(ctx context.Context, body DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsPostRequestBodyable, requestConfiguration *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder when successful
func (m *DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder) {
    return NewDeviceconfigurationsItemWindowsprivacyaccesscontrolsWindowsPrivacyAccessControlsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
