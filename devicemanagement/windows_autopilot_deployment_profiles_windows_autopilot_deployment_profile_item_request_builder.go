package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters windows auto pilot deployment profiles
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetQueryParameters
}
// WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignedDevices provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignedDevices()(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignedDevicesById provides operations to manage the assignedDevices property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignedDevicesById(id string)(*WindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewWindowsAutopilotDeploymentProfilesItemAssignedDevicesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Assignments()(*WindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsAutopilotDeploymentProfile entity.
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) AssignmentsById(id string)(*WindowsAutopilotDeploymentProfilesItemAssignmentsWindowsAutopilotDeploymentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeploymentProfileAssignment%2Did"] = id
    }
    return NewWindowsAutopilotDeploymentProfilesItemAssignmentsWindowsAutopilotDeploymentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal instantiates a new WindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    m := &WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder instantiates a new WindowsAutopilotDeploymentProfileItemRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) MicrosoftGraphAssign()(*WindowsAutopilotDeploymentProfilesItemMicrosoftGraphAssignRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable), nil
}
// ToDeleteRequestInformation delete navigation property windowsAutopilotDeploymentProfiles for deviceManagement
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property windowsAutopilotDeploymentProfiles in deviceManagement
func (m *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAutopilotDeploymentProfileable, requestConfiguration *WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
