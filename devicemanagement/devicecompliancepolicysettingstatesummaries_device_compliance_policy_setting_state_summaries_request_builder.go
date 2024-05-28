package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters the summary states of compliance policy settings for this account.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters struct {
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
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceCompliancePolicySettingStateSummaryId provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) ByDeviceCompliancePolicySettingStateSummaryId(deviceCompliancePolicySettingStateSummaryId string)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceCompliancePolicySettingStateSummaryId != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = deviceCompliancePolicySettingStateSummaryId
    }
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal instantiates a new DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    m := &DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder instantiates a new DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecompliancepolicysettingstatesummariesCountRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) Count()(*DevicecompliancepolicysettingstatesummariesCountRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the summary states of compliance policy settings for this account.
// returns a DeviceCompliancePolicySettingStateSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicySettingStateSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryCollectionResponseable), nil
}
// Post create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement
// returns a DeviceCompliancePolicySettingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryable), nil
}
// ToGetRequestInformation the summary states of compliance policy settings for this account.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceCompliancePolicySettingStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
