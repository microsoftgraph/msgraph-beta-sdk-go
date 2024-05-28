package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters the device compliance state summary for this account.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal instantiates a new DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder and sets the default values.
func NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    m := &DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicyDeviceStateSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder instantiates a new DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder and sets the default values.
func NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the device compliance state summary for this account.
// returns a DeviceCompliancePolicyDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable), nil
}
// Patch update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement
// returns a DeviceCompliancePolicyDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the device compliance state summary for this account.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceCompliancePolicyDeviceStateSummary in deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyDeviceStateSummaryable, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
