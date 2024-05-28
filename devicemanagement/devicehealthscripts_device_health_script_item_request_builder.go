package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicehealthscriptsDeviceHealthScriptItemRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DevicehealthscriptsDeviceHealthScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicehealthscriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetQueryParameters
}
// DevicehealthscriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicehealthscriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DevicehealthscriptsItemAssignRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) Assign()(*DevicehealthscriptsItemAssignRequestBuilder) {
    return NewDevicehealthscriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceHealthScript entity.
// returns a *DevicehealthscriptsItemAssignmentsRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) Assignments()(*DevicehealthscriptsItemAssignmentsRequestBuilder) {
    return NewDevicehealthscriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilderInternal instantiates a new DevicehealthscriptsDeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) {
    m := &DevicehealthscriptsDeviceHealthScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilder instantiates a new DevicehealthscriptsDeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceHealthScripts for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceHealthScript entity.
// returns a *DevicehealthscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) DeviceRunStates()(*DevicehealthscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDevicehealthscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device health scripts associated with the tenant.
// returns a DeviceHealthScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// GetGlobalScriptHighestAvailableVersion provides operations to call the getGlobalScriptHighestAvailableVersion method.
// returns a *DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) GetGlobalScriptHighestAvailableVersion()(*DevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return NewDevicehealthscriptsItemGetglobalscripthighestavailableversionGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetRemediationHistory provides operations to call the getRemediationHistory method.
// returns a *DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) GetRemediationHistory()(*DevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilder) {
    return NewDevicehealthscriptsItemGetremediationhistoryGetRemediationHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceHealthScripts in deviceManagement
// returns a DeviceHealthScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceHealthScript entity.
// returns a *DevicehealthscriptsItemRunsummaryRunSummaryRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) RunSummary()(*DevicehealthscriptsItemRunsummaryRunSummaryRequestBuilder) {
    return NewDevicehealthscriptsItemRunsummaryRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceHealthScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of device health scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property deviceHealthScripts in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, requestConfiguration *DevicehealthscriptsDeviceHealthScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// UpdateGlobalScript provides operations to call the updateGlobalScript method.
// returns a *DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) UpdateGlobalScript()(*DevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilder) {
    return NewDevicehealthscriptsItemUpdateglobalscriptUpdateGlobalScriptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder when successful
func (m *DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) WithUrl(rawUrl string)(*DevicehealthscriptsDeviceHealthScriptItemRequestBuilder) {
    return NewDevicehealthscriptsDeviceHealthScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
