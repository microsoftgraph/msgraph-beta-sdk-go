package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder provides operations to call the getTotalAggregatedRemoteConnectionReports method.
type VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal instantiates a new GetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    m := &VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getTotalAggregatedRemoteConnectionReports", pathParameters),
    }
    return m
}
// NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder instantiates a new GetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the total aggregated remote connection usage of a Cloud PC during a given time span. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-gettotalaggregatedremoteconnectionreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation get the total aggregated remote connection usage of a Cloud PC during a given time span. This API is available in the following national cloud deployments.
func (m *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
