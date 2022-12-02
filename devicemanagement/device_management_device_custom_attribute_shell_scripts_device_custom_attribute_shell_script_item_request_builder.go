package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters the list of device custom attribute shell scripts associated with the tenant.
type DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assign()(*DeviceManagementDeviceCustomAttributeShellScriptsItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assignments()(*DeviceManagementDeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceCustomAttributeShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    m := &DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceManagementDeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceManagementDeviceCustomAttributeShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignments()(*DeviceManagementDeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceManagementDeviceCustomAttributeShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) RunSummary()(*DeviceManagementDeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStates()(*DeviceManagementDeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStatesById(id string)(*DeviceManagementDeviceCustomAttributeShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return NewDeviceManagementDeviceCustomAttributeShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
