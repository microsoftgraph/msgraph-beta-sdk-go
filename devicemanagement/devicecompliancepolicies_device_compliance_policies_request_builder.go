package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetQueryParameters the device compliance policies.
type DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceCompliancePolicyId provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) ByDeviceCompliancePolicyId(deviceCompliancePolicyId string)(*DevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceCompliancePolicyId != "" {
        urlTplParams["deviceCompliancePolicy%2Did"] = deviceCompliancePolicyId
    }
    return NewDevicecompliancepoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderInternal instantiates a new DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) {
    m := &DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder instantiates a new DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecompliancepoliciesCountRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) Count()(*DevicecompliancepoliciesCountRequestBuilder) {
    return NewDevicecompliancepoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the device compliance policies.
// returns a DeviceCompliancePolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyCollectionResponseable), nil
}
// GetDevicesScheduledToRetire provides operations to call the getDevicesScheduledToRetire method.
// returns a *DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) GetDevicesScheduledToRetire()(*DevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilder) {
    return NewDevicecompliancepoliciesGetdevicesscheduledtoretireGetDevicesScheduledToRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetNoncompliantDevicesToRetire provides operations to call the getNoncompliantDevicesToRetire method.
// returns a *DevicecompliancepoliciesGetnoncompliantdevicestoretireGetNoncompliantDevicesToRetireRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) GetNoncompliantDevicesToRetire()(*DevicecompliancepoliciesGetnoncompliantdevicestoretireGetNoncompliantDevicesToRetireRequestBuilder) {
    return NewDevicecompliancepoliciesGetnoncompliantdevicestoretireGetNoncompliantDevicesToRetireRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
// returns a *DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) HasPayloadLinks()(*DevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewDevicecompliancepoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to deviceCompliancePolicies for deviceManagement
// returns a DeviceCompliancePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// RefreshDeviceComplianceReportSummarization provides operations to call the refreshDeviceComplianceReportSummarization method.
// returns a *DevicecompliancepoliciesRefreshdevicecompliancereportsummarizationRefreshDeviceComplianceReportSummarizationRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) RefreshDeviceComplianceReportSummarization()(*DevicecompliancepoliciesRefreshdevicecompliancereportsummarizationRefreshDeviceComplianceReportSummarizationRequestBuilder) {
    return NewDevicecompliancepoliciesRefreshdevicecompliancereportsummarizationRefreshDeviceComplianceReportSummarizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetScheduledRetireState provides operations to call the setScheduledRetireState method.
// returns a *DevicecompliancepoliciesSetscheduledretirestateSetScheduledRetireStateRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) SetScheduledRetireState()(*DevicecompliancepoliciesSetscheduledretirestateSetScheduledRetireStateRequestBuilder) {
    return NewDevicecompliancepoliciesSetscheduledretirestateSetScheduledRetireStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the device compliance policies.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceCompliancePolicies for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// ValidateComplianceScript provides operations to call the validateComplianceScript method.
// returns a *DevicecompliancepoliciesValidatecompliancescriptValidateComplianceScriptRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) ValidateComplianceScript()(*DevicecompliancepoliciesValidatecompliancescriptValidateComplianceScriptRequestBuilder) {
    return NewDevicecompliancepoliciesValidatecompliancescriptValidateComplianceScriptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder when successful
func (m *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) {
    return NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
