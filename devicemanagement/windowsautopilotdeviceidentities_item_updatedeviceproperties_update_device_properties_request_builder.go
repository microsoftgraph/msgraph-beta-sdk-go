package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder provides operations to call the updateDeviceProperties method.
type WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal instantiates a new WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    m := &WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/updateDeviceProperties", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder instantiates a new WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post updates properties on Autopilot devices.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) Post(ctx context.Context, body WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBodyable, requestConfiguration *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation updates properties on Autopilot devices.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesPostRequestBodyable, requestConfiguration *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
