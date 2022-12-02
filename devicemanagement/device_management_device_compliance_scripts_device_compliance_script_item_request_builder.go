package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters the list of device compliance scripts associated with the tenant.
type DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Assign()(*DeviceManagementDeviceComplianceScriptsItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceComplianceScriptsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Assignments()(*DeviceManagementDeviceComplianceScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceComplianceScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceComplianceScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceComplianceScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    m := &DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) DeviceRunStates()(*DeviceManagementDeviceComplianceScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceManagementDeviceComplianceScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceManagementDeviceComplianceScriptsItemDeviceRunStatesDeviceComplianceScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScriptDeviceState%2Did"] = id
    }
    return NewDeviceManagementDeviceComplianceScriptsItemDeviceRunStatesDeviceComplianceScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device compliance scripts associated with the tenant.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// Patch update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) RunSummary()(*DeviceManagementDeviceComplianceScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceManagementDeviceComplianceScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
