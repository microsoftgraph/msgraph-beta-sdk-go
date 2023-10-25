package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder provides operations to call the getDailyAggregatedRemoteConnectionReports method.
type VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal instantiates a new GetDailyAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    m := &VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getDailyAggregatedRemoteConnectionReports", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder instantiates a new GetDailyAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the daily aggregated remote connection reports, such as round trip time, available bandwidth, and so on, in a given period. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-getdailyaggregatedremoteconnectionreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation get the daily aggregated remote connection reports, such as round trip time, available bandwidth, and so on, in a given period. This API is available in the following national cloud deployments.
func (m *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetDailyAggregatedRemoteConnectionReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
