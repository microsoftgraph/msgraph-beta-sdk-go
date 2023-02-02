package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters
}
// DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Assignments()(*DeviceHealthScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceHealthScriptsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) AssignmentsById(id string)(*DeviceHealthScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment%2Did"] = id
    }
    return NewDeviceHealthScriptsItemAssignmentsDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    m := &DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) DeviceRunStates()(*DeviceHealthScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceHealthScriptsItemDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceRunStatesById provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) DeviceRunStatesById(id string)(*DeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptDeviceState%2Did"] = id
    }
    return NewDeviceHealthScriptsItemDeviceRunStatesDeviceHealthScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// Get the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) MicrosoftGraphAssign()(*DeviceHealthScriptsItemMicrosoftGraphAssignAssignRequestBuilder) {
    return NewDeviceHealthScriptsItemMicrosoftGraphAssignAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetGlobalScriptHighestAvailableVersion provides operations to call the getGlobalScriptHighestAvailableVersion method.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) MicrosoftGraphGetGlobalScriptHighestAvailableVersion()(*DeviceHealthScriptsItemMicrosoftGraphGetGlobalScriptHighestAvailableVersionGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return NewDeviceHealthScriptsItemMicrosoftGraphGetGlobalScriptHighestAvailableVersionGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetRemediationHistory provides operations to call the getRemediationHistory method.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) MicrosoftGraphGetRemediationHistory()(*DeviceHealthScriptsItemMicrosoftGraphGetRemediationHistoryGetRemediationHistoryRequestBuilder) {
    return NewDeviceHealthScriptsItemMicrosoftGraphGetRemediationHistoryGetRemediationHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUpdateGlobalScript provides operations to call the updateGlobalScript method.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) MicrosoftGraphUpdateGlobalScript()(*DeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptRequestBuilder) {
    return NewDeviceHealthScriptsItemMicrosoftGraphUpdateGlobalScriptUpdateGlobalScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceHealthScript entity.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) RunSummary()(*DeviceHealthScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceHealthScriptsItemRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DeviceHealthScriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
