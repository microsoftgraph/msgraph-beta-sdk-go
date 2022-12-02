package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters windows auto pilot deployment profiles
type DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters
}
// DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Assign()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignedDevices provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignedDevices()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignedDevicesById provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignedDevicesById(id string)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Assignments()(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementWindowsAutopilotDeploymentProfilesItemAssignmentsWindowsAutopilotDeploymentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeploymentProfileAssignment%2Did"] = id
    }
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesItemAssignmentsWindowsAutopilotDeploymentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal instantiates a new WindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    m := &DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder instantiates a new WindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation windows auto pilot deployment profiles
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windows auto pilot deployment profiles
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
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
// Patch update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
func (m *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
