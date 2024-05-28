package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
type DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetQueryParameters the list of device management scripts associated with the tenant.
type DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetQueryParameters
}
// DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *DevicemanagementscriptsItemAssignRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) Assign()(*DevicemanagementscriptsItemAssignRequestBuilder) {
    return NewDevicemanagementscriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemAssignmentsRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) Assignments()(*DevicemanagementscriptsItemAssignmentsRequestBuilder) {
    return NewDevicemanagementscriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilderInternal instantiates a new DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder and sets the default values.
func NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) {
    m := &DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilder instantiates a new DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder and sets the default values.
func NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceManagementScripts for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) DeviceRunStates()(*DevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewDevicemanagementscriptsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device management scripts associated with the tenant.
// returns a DeviceManagementScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemGroupassignmentsGroupAssignmentsRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) GroupAssignments()(*DevicemanagementscriptsItemGroupassignmentsGroupAssignmentsRequestBuilder) {
    return NewDevicemanagementscriptsItemGroupassignmentsGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceManagementScripts in deviceManagement
// returns a DeviceManagementScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemRunsummaryRunSummaryRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) RunSummary()(*DevicemanagementscriptsItemRunsummaryRunSummaryRequestBuilder) {
    return NewDevicemanagementscriptsItemRunsummaryRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceManagementScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of device management scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceManagementScripts in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, requestConfiguration *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceManagementScript entity.
// returns a *DevicemanagementscriptsItemUserrunstatesUserRunStatesRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) UserRunStates()(*DevicemanagementscriptsItemUserrunstatesUserRunStatesRequestBuilder) {
    return NewDevicemanagementscriptsItemUserrunstatesUserRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder when successful
func (m *DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) WithUrl(rawUrl string)(*DevicemanagementscriptsDeviceManagementScriptItemRequestBuilder) {
    return NewDevicemanagementscriptsDeviceManagementScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
