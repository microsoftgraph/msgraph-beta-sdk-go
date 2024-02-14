package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *DeviceCustomAttributeShellScriptsItemAssignRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assign()(*DeviceCustomAttributeShellScriptsItemAssignRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Assignments()(*DeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder instantiates a new DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCustomAttributeShellScripts for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStates()(*DeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device custom attribute shell scripts associated with the tenant.
// returns a DeviceCustomAttributeShellScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignments()(*DeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceCustomAttributeShellScripts in deviceManagement
// returns a DeviceCustomAttributeShellScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) RunSummary()(*DeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceCustomAttributeShellScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceCustomAttributeShellScripts in deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}", m.BaseRequestBuilder.PathParameters)
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
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.deviceCustomAttributeShellScript entity.
// returns a *DeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStates()(*DeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsItemUserRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder when successful
func (m *DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) WithUrl(rawUrl string)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
