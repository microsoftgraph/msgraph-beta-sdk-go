package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareconfigurationsHardwareConfigurationItemRequestBuilder provides operations to manage the hardwareConfigurations property of the microsoft.graph.deviceManagement entity.
type HardwareconfigurationsHardwareConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareconfigurationsHardwareConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareconfigurationsHardwareConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetQueryParameters the hardware configurations for this account.
type HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetQueryParameters
}
// HardwareconfigurationsHardwareConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareconfigurationsHardwareConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *HardwareconfigurationsItemAssignRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) Assign()(*HardwareconfigurationsItemAssignRequestBuilder) {
    return NewHardwareconfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareconfigurationsItemAssignmentsRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) Assignments()(*HardwareconfigurationsItemAssignmentsRequestBuilder) {
    return NewHardwareconfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewHardwareconfigurationsHardwareConfigurationItemRequestBuilderInternal instantiates a new HardwareconfigurationsHardwareConfigurationItemRequestBuilder and sets the default values.
func NewHardwareconfigurationsHardwareConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareconfigurationsHardwareConfigurationItemRequestBuilder) {
    m := &HardwareconfigurationsHardwareConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHardwareconfigurationsHardwareConfigurationItemRequestBuilder instantiates a new HardwareconfigurationsHardwareConfigurationItemRequestBuilder and sets the default values.
func NewHardwareconfigurationsHardwareConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareconfigurationsHardwareConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareconfigurationsHardwareConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property hardwareConfigurations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceRunStates provides operations to manage the deviceRunStates property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) DeviceRunStates()(*HardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilder) {
    return NewHardwareconfigurationsItemDevicerunstatesDeviceRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the hardware configurations for this account.
// returns a HardwareConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable), nil
}
// Patch update the navigation property hardwareConfigurations in deviceManagement
// returns a HardwareConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable), nil
}
// RunSummary provides operations to manage the runSummary property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareconfigurationsItemRunsummaryRunSummaryRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) RunSummary()(*HardwareconfigurationsItemRunsummaryRunSummaryRequestBuilder) {
    return NewHardwareconfigurationsItemRunsummaryRunSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property hardwareConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the hardware configurations for this account.
// returns a *RequestInformation when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hardwareConfigurations in deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationable, requestConfiguration *HardwareconfigurationsHardwareConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserRunStates provides operations to manage the userRunStates property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareconfigurationsItemUserrunstatesUserRunStatesRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) UserRunStates()(*HardwareconfigurationsItemUserrunstatesUserRunStatesRequestBuilder) {
    return NewHardwareconfigurationsItemUserrunstatesUserRunStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HardwareconfigurationsHardwareConfigurationItemRequestBuilder when successful
func (m *HardwareconfigurationsHardwareConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*HardwareconfigurationsHardwareConfigurationItemRequestBuilder) {
    return NewHardwareconfigurationsHardwareConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
