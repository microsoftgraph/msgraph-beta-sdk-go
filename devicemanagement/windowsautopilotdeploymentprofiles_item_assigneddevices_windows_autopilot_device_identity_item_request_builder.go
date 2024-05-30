package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the list of assigned devices for the profile.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowNextEnrollment provides operations to call the allowNextEnrollment method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) AllowNextEnrollment()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignResourceAccountToDevice provides operations to call the assignResourceAccountToDevice method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignUserToDevice provides operations to call the assignUserToDevice method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignusertodeviceAssignUserToDeviceRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignusertodeviceAssignUserToDeviceRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemAssignusertodeviceAssignUserToDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignedDevices for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// DeploymentProfile provides operations to manage the deploymentProfile property of the microsoft.graph.windowsAutopilotDeviceIdentity entity.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemDeploymentprofileDeploymentProfileRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemDeploymentprofileDeploymentProfileRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemDeploymentprofileDeploymentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of assigned devices for the profile.
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// IntendedDeploymentProfile provides operations to manage the intendedDeploymentProfile property of the microsoft.graph.windowsAutopilotDeviceIdentity entity.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignedDevices in deviceManagement
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// ToDeleteRequestInformation delete navigation property assignedDevices for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of assigned devices for the profile.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property assignedDevices in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UnassignResourceAccountFromDevice provides operations to call the unassignResourceAccountFromDevice method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnassignUserFromDevice provides operations to call the unassignUserFromDevice method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateDeviceProperties provides operations to call the updateDeviceProperties method.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*WindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesItemAssigneddevicesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
