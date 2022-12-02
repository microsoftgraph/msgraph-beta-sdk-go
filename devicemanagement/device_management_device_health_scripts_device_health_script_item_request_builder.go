package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Assign()(*DeviceManagementDeviceHealthScriptsItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Assignments()(*DeviceManagementDeviceHealthScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceHealthScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceHealthScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    m := &DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) DeviceRunStates()(*DeviceManagementDeviceHealthScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceManagementDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptDeviceState%2Did"] = id
    }
    return NewDeviceManagementDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device health scripts associated with the tenant.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// GetGlobalScriptHighestAvailableVersion provides operations to call the getGlobalScriptHighestAvailableVersion method.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) GetGlobalScriptHighestAvailableVersion()(*DeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemediationHistory provides operations to call the getRemediationHistory method.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) GetRemediationHistory()(*DeviceManagementDeviceHealthScriptsItemGetRemediationHistoryRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemGetRemediationHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) RunSummary()(*DeviceManagementDeviceHealthScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateGlobalScript provides operations to call the updateGlobalScript method.
func (m *DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) UpdateGlobalScript()(*DeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsItemUpdateGlobalScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
