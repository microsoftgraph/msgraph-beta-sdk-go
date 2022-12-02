package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder provides operations to manage the intendedDeploymentProfile property of the microsoft.graph.windowsAutopilotDeviceIdentity entity.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetQueryParameters deployment profile intended to be assigned to the Windows autopilot device.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetQueryParameters
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderInternal instantiates a new IntendedDeploymentProfileRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder) {
    m := &DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}/assignedDevices/{windowsAutopilotDeviceIdentity%2Did}/intendedDeploymentProfile{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder instantiates a new IntendedDeploymentProfileRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation deployment profile intended to be assigned to the Windows autopilot device.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get deployment profile intended to be assigned to the Windows autopilot device.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesItemIntendedDeploymentProfileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable), nil
}
