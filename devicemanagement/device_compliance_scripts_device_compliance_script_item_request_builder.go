package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
type DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// Assign provides operations to call the assign method.
// returns a *DeviceComplianceScriptsItemAssignRequestBuilder when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Assign()(*DeviceComplianceScriptsItemAssignRequestBuilder) {
    return NewDeviceComplianceScriptsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceComplianceScript entity.
// returns a *DeviceComplianceScriptsItemAssignmentsRequestBuilder when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Assignments()(*DeviceComplianceScriptsItemAssignmentsRequestBuilder) {
    return NewDeviceComplianceScriptsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal instantiates a new DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    m := &DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder instantiates a new DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceComplianceScripts for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.deviceComplianceScript entity.
// returns a *DeviceComplianceScriptsItemDeviceRunStatesRequestBuilder when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) DeviceRunStates()(*DeviceComplianceScriptsItemDeviceRunStatesRequestBuilder) {
    return NewDeviceComplianceScriptsItemDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device compliance scripts associated with the tenant.
// returns a DeviceComplianceScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// Patch update the navigation property deviceComplianceScripts in deviceManagement
// returns a DeviceComplianceScriptable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.deviceComplianceScript entity.
// returns a *DeviceComplianceScriptsItemRunSummaryRequestBuilder when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) RunSummary()(*DeviceComplianceScriptsItemRunSummaryRequestBuilder) {
    return NewDeviceComplianceScriptsItemRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceComplianceScripts for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of device compliance scripts associated with the tenant.
// returns a *RequestInformation when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceComplianceScripts in deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder when successful
func (m *DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) WithUrl(rawUrl string)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    return NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
