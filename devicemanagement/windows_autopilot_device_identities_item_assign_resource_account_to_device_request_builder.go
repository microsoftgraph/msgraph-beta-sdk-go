package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder provides operations to call the assignResourceAccountToDevice method.
type WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderInternal instantiates a new AssignResourceAccountToDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) {
    m := &WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}/assignResourceAccountToDevice", pathParameters),
    }
    return m
}
// NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder instantiates a new AssignResourceAccountToDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post assigns resource account to Autopilot devices.
func (m *WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) Post(ctx context.Context, body WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDevicePostRequestBodyable, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation assigns resource account to Autopilot devices.
func (m *WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDevicePostRequestBodyable, requestConfiguration *WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) WithUrl(rawUrl string)(*WindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder) {
    return NewWindowsAutopilotDeviceIdentitiesItemAssignResourceAccountToDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
