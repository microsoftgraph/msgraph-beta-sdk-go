package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder provides operations to call the getHealthMetrics method.
type MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderInternal instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) {
    m := &MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}/getHealthMetrics", pathParameters),
    }
    return m
}
// NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder instantiates a new MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder and sets the default values.
func NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getHealthMetrics
// Deprecated: This method is obsolete. Use PostAsGetHealthMetricsPostResponse instead.
// returns a MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) Post(ctx context.Context, body MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderPostRequestConfiguration)(MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsResponseable), nil
}
// PostAsGetHealthMetricsPostResponse invoke action getHealthMetrics
// returns a MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) PostAsGetHealthMetricsPostResponse(ctx context.Context, body MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderPostRequestConfiguration)(MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostResponseable), nil
}
// ToPostRequestInformation invoke action getHealthMetrics
// returns a *RequestInformation when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder when successful
func (m *MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) WithUrl(rawUrl string)(*MicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder) {
    return NewMicrosofttunnelsitesItemMicrosofttunnelserversItemGethealthmetricsGetHealthMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
