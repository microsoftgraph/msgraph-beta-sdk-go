package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters the list of device shell scripts associated with the tenant.
type DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) Assign()(*DeviceManagementDeviceShellScriptsItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) Assignments()(*DeviceManagementDeviceShellScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    m := &DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device shell scripts associated with the tenant.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceManagementDeviceShellScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceManagementDeviceShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return NewDeviceManagementDeviceShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device shell scripts associated with the tenant.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) GroupAssignments()(*DeviceManagementDeviceShellScriptsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceManagementDeviceShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) RunSummary()(*DeviceManagementDeviceShellScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) UserRunStates()(*DeviceManagementDeviceShellScriptsItemUserRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsItemUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById provides operations to manage the userRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) UserRunStatesById(id string)(*DeviceManagementDeviceShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return NewDeviceManagementDeviceShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
