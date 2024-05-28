package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the Windows autopilot device identities contained collection.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllowNextEnrollment provides operations to call the allowNextEnrollment method.
// returns a *WindowsautopilotdeviceidentitiesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) AllowNextEnrollment()(*WindowsautopilotdeviceidentitiesItemAllownextenrollmentAllowNextEnrollmentRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemAllownextenrollmentAllowNextEnrollmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignResourceAccountToDevice provides operations to call the assignResourceAccountToDevice method.
// returns a *WindowsautopilotdeviceidentitiesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*WindowsautopilotdeviceidentitiesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemAssignresourceaccounttodeviceAssignResourceAccountToDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignUserToDevice provides operations to call the assignUserToDevice method.
// returns a *WindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*WindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemAssignusertodeviceAssignUserToDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *WindowsautopilotdeviceidentitiesItemDeploymentprofileDeploymentProfileRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*WindowsautopilotdeviceidentitiesItemDeploymentprofileDeploymentProfileRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemDeploymentprofileDeploymentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Windows autopilot device identities contained collection.
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
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
// returns a *WindowsautopilotdeviceidentitiesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*WindowsautopilotdeviceidentitiesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemIntendeddeploymentprofileIntendedDeploymentProfileRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
// returns a WindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
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
// ToDeleteRequestInformation delete navigation property windowsAutopilotDeviceIdentities for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Windows autopilot device identities contained collection.
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsAutopilotDeviceIdentities in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*WindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUnassignresourceaccountfromdeviceUnassignResourceAccountFromDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UnassignUserFromDevice provides operations to call the unassignUserFromDevice method.
// returns a *WindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*WindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUnassignuserfromdeviceUnassignUserFromDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UpdateDeviceProperties provides operations to call the updateDeviceProperties method.
// returns a *WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*WindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesItemUpdatedevicepropertiesUpdateDevicePropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) WithUrl(rawUrl string)(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
