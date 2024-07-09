package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder provides operations to call the getHealthMetrics method.
type MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderInternal instantiates a new MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) {
    m := &MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}/microsoftTunnelServers/{microsoftTunnelServer%2Did}/getHealthMetrics", pathParameters),
    }
    return m
}
// NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder instantiates a new MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder and sets the default values.
func NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getHealthMetrics
// Deprecated: This method is obsolete. Use PostAsGetHealthMetricsPostResponse instead.
// returns a MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) Post(ctx context.Context, body MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsResponseable), nil
}
// PostAsGetHealthMetricsPostResponse invoke action getHealthMetrics
// returns a MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) PostAsGetHealthMetricsPostResponse(ctx context.Context, body MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostResponseable), nil
}
// ToPostRequestInformation invoke action getHealthMetrics
// returns a *RequestInformation when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) ToPostRequestInformation(ctx context.Context, body MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsPostRequestBodyable, requestConfiguration *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder when successful
func (m *MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) WithUrl(rawUrl string)(*MicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder) {
    return NewMicrosoftTunnelSitesItemMicrosoftTunnelServersItemGetHealthMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
