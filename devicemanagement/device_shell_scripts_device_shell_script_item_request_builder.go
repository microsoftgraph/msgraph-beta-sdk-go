package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceShellScriptsDeviceShellScriptItemRequestBuilder provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceShellScriptsDeviceShellScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters the list of device shell scripts associated with the tenant.
type DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) Assign()(*DeviceShellScriptsItemAssignRequestBuilder) {
    return NewDeviceShellScriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) Assignments()(*DeviceShellScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceShellScriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return NewDeviceShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    m := &DeviceShellScriptsDeviceShellScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewDeviceShellScriptsDeviceShellScriptItemRequestBuilder instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceShellScriptsDeviceShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceShellScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceShellScriptsItemDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return NewDeviceShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device shell scripts associated with the tenant.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) GroupAssignments()(*DeviceShellScriptsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceShellScriptsItemGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return NewDeviceShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) RunSummary()(*DeviceShellScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceShellScriptsItemRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the list of device shell scripts associated with the tenant.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceShellScriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
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
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) UserRunStates()(*DeviceShellScriptsItemUserRunStatesRequestBuilder) {
    return NewDeviceShellScriptsItemUserRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserRunStatesById provides operations to manage the userRunStates property of the microsoft.graph.deviceShellScript entity.
func (m *DeviceShellScriptsDeviceShellScriptItemRequestBuilder) UserRunStatesById(id string)(*DeviceShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return NewDeviceShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
