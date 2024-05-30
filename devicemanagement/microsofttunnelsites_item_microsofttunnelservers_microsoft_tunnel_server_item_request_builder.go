package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
type MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetQueryParameters a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
type MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetQueryParameters
}
// MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderInternal instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) {
    m := &MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateServerLogCollectionRequest provides operations to call the createServerLogCollectionRequest method.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemCreateserverlogcollectionrequestCreateServerLogCollectionRequestRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) CreateServerLogCollectionRequest()(*MicrosofttunnelsitesItemMicrosofttunnelserversItemCreateserverlogcollectionrequestCreateServerLogCollectionRequestRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemCreateserverlogcollectionrequestCreateServerLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property microsoftTunnelServers for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// GenerateServerLogCollectionRequest provides operations to call the generateServerLogCollectionRequest method.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) GenerateServerLogCollectionRequest()(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGenerateserverlogcollectionrequestGenerateServerLogCollectionRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
// returns a MicrosoftTunnelServerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
// GetHealthMetrics provides operations to call the getHealthMetrics method.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) GetHealthMetrics()(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetHealthMetricTimeSeries provides operations to call the getHealthMetricTimeSeries method.
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) GetHealthMetricTimeSeries()(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetrictimeseriesGetHealthMetricTimeSeriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property microsoftTunnelServers in deviceManagement
// returns a MicrosoftTunnelServerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelServerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable), nil
}
// ToDeleteRequestInformation delete navigation property microsoftTunnelServers for deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property microsoftTunnelServers in deviceManagement
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelServerable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversMicrosoftTunnelServerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
