package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters the list of device custom attribute shell scripts associated with the tenant.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assign()(*DeviceCustomAttributeShellScriptsItemAssignRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assignments()(*DeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceCustomAttributeShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return NewDeviceCustomAttributeShellScriptsItemAssignmentsDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder{
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
// NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceCustomAttributeShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return NewDeviceCustomAttributeShellScriptsItemDeviceRunStatesDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
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
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignments()(*DeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceCustomAttributeShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return NewDeviceCustomAttributeShellScriptsItemGroupAssignmentsDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
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
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) RunSummary()(*DeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStates()(*DeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStatesById(id string)(*DeviceCustomAttributeShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return NewDeviceCustomAttributeShellScriptsItemUserRunStatesDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
