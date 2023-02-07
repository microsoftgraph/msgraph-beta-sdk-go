package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters the list of device compliance scripts associated with the tenant.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetQueryParameters
}
// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Assignments()(*DeviceComplianceScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceComplianceScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceComplianceScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment%2Did"] = id
    }
    return NewDeviceComplianceScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    m := &DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) DeviceRunStates()(*DeviceComplianceScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceComplianceScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceComplianceScriptsItemDeviceRunStatesDeviceComplianceScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScriptDeviceState%2Did"] = id
    }
    return NewDeviceComplianceScriptsItemDeviceRunStatesDeviceComplianceScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) MicrosoftGraphAssign()(*DeviceComplianceScriptsItemMicrosoftGraphAssignRequestBuilder) {
    return NewDeviceComplianceScriptsItemMicrosoftGraphAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceComplianceScript entity.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) RunSummary()(*DeviceComplianceScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceComplianceScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
