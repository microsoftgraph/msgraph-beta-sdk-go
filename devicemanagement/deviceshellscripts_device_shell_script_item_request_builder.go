package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceshellscriptsDeviceShellScriptItemRequestBuilder provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceshellscriptsDeviceShellScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceshellscriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceshellscriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetQueryParameters the list of device shell scripts associated with the tenant.
type DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceshellscriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceshellscriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DeviceshellscriptsItemAssignRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) Assign()(*DeviceshellscriptsItemAssignRequestBuilder) {
    return NewDeviceshellscriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceShellScript entity.
// returns a *DeviceshellscriptsItemAssignmentsRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) Assignments()(*DeviceshellscriptsItemAssignmentsRequestBuilder) {
    return NewDeviceshellscriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceshellscriptsDeviceShellScriptItemRequestBuilderInternal instantiates a new DeviceshellscriptsDeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceshellscriptsDeviceShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceshellscriptsDeviceShellScriptItemRequestBuilder) {
    m := &DeviceshellscriptsDeviceShellScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceshellscriptsDeviceShellScriptItemRequestBuilder instantiates a new DeviceshellscriptsDeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceshellscriptsDeviceShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceshellscriptsDeviceShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceshellscriptsDeviceShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceShellScripts for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceShellScript entity.
// returns a *DeviceshellscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceshellscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDeviceshellscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device shell scripts associated with the tenant.
// returns a DeviceShellScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *DeviceshellscriptsItemGroupassignmentsGroupAssignmentsRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) GroupAssignments()(*DeviceshellscriptsItemGroupassignmentsGroupAssignmentsRequestBuilder) {
    return NewDeviceshellscriptsItemGroupassignmentsGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceShellScripts in deviceManagement
// returns a DeviceShellScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// returns a *DeviceshellscriptsItemRunsummaryRunSummaryRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) RunSummary()(*DeviceshellscriptsItemRunsummaryRunSummaryRequestBuilder) {
    return NewDeviceshellscriptsItemRunsummaryRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceShellScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of device shell scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceShellScripts in deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, requestConfiguration *DeviceshellscriptsDeviceShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceShellScript entity.
// returns a *DeviceshellscriptsItemUserrunstatesUserRunStatesRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) UserRunStates()(*DeviceshellscriptsItemUserrunstatesUserRunStatesRequestBuilder) {
    return NewDeviceshellscriptsItemUserrunstatesUserRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceshellscriptsDeviceShellScriptItemRequestBuilder when successful
func (m *DeviceshellscriptsDeviceShellScriptItemRequestBuilder) WithUrl(rawUrl string)(*DeviceshellscriptsDeviceShellScriptItemRequestBuilder) {
    return NewDeviceshellscriptsDeviceShellScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
