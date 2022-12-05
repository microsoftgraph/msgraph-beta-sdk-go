package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters the list of assigned devices for the profile.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetQueryParameters
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignResourceAccountToDevice provides operations to call the assignResourceAccountToDevice method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignResourceAccountToDevice()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignResourceAccountToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignUserToDevice provides operations to call the assignUserToDevice method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) AssignUserToDevice()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignUserToDeviceRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemAssignUserToDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    m := &DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder instantiates a new WindowsAutopilotDeviceIdentityItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignedDevices for deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of assigned devices for the profile.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property assignedDevices in deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property assignedDevices for deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeploymentProfile provides operations to manage the deploymentProfile property of the microsoft.graph.windowsAutopilotDeviceIdentity entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) DeploymentProfile()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemDeploymentProfileRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of assigned devices for the profile.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// IntendedDeploymentProfile provides operations to manage the intendedDeploymentProfile property of the microsoft.graph.windowsAutopilotDeviceIdentity entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) IntendedDeploymentProfile()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property assignedDevices in deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeviceIdentityable), nil
}
// UnassignResourceAccountFromDevice provides operations to call the unassignResourceAccountFromDevice method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignResourceAccountFromDevice()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUnassignResourceAccountFromDeviceRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUnassignResourceAccountFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnassignUserFromDevice provides operations to call the unassignUserFromDevice method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UnassignUserFromDevice()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUnassignUserFromDeviceRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUnassignUserFromDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateDeviceProperties provides operations to call the updateDeviceProperties method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) UpdateDeviceProperties()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemUpdateDevicePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
