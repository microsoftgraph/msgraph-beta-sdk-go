package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
type DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters list of DeviceComplianceDeviceStatus.
type DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters struct {
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
// DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceComplianceDeviceStatusId provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
// returns a *DevicecompliancepoliciesItemDevicestatusesDeviceComplianceDeviceStatusItemRequestBuilder when successful
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ByDeviceComplianceDeviceStatusId(deviceComplianceDeviceStatusId string)(*DevicecompliancepoliciesItemDevicestatusesDeviceComplianceDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceComplianceDeviceStatusId != "" {
        urlTplParams["deviceComplianceDeviceStatus%2Did"] = deviceComplianceDeviceStatusId
    }
    return NewDevicecompliancepoliciesItemDevicestatusesDeviceComplianceDeviceStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal instantiates a new DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    m := &DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/deviceStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder instantiates a new DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecompliancepoliciesItemDevicestatusesCountRequestBuilder when successful
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Count()(*DevicecompliancepoliciesItemDevicestatusesCountRequestBuilder) {
    return NewDevicecompliancepoliciesItemDevicestatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of DeviceComplianceDeviceStatus.
// returns a DeviceComplianceDeviceStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceDeviceStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusCollectionResponseable), nil
}
// Post create new navigation property to deviceStatuses for deviceManagement
// returns a DeviceComplianceDeviceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusable, requestConfiguration *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceDeviceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusable), nil
}
// ToGetRequestInformation list of DeviceComplianceDeviceStatus.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceStatuses for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceDeviceStatusable, requestConfiguration *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewDevicecompliancepoliciesItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
