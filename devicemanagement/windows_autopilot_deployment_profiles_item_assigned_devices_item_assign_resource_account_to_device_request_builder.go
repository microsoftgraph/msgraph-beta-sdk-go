package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder provides operations to call the assignResourceAccountToDevice method.
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderInternal instantiates a new AssignResourceAccountToDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) {
    m := &WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}/assignResourceAccountToDevice", pathParameters),
    }
    return m
}
// NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder instantiates a new AssignResourceAccountToDeviceRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post assigns resource account to Autopilot devices.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) Post(ctx context.Context, body WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyable, requestConfiguration *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration)(error) {
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
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) ToPostRequestInformation(ctx context.Context, body WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDevicePostRequestBodyable, requestConfiguration *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) WithUrl(rawUrl string)(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
