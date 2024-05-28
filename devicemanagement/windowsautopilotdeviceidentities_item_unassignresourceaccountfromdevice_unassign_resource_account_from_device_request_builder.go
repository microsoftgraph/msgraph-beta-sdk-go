package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder provides operations to call the unassignResourceAccountFromDevice method.
type WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderInternal instantiates a new WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) {
    m := &WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/unassignResourceAccountFromDevice", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder instantiates a new WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unassigns the resource account from an Autopilot device.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) Post(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unassigns the resource account from an Autopilot device.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
