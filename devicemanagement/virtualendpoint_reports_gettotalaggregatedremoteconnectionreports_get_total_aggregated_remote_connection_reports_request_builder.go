package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder provides operations to call the getTotalAggregatedRemoteConnectionReports method.
type VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal instantiates a new VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    m := &VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/reports/getTotalAggregatedRemoteConnectionReports", pathParameters),
    }
    return m
}
// NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder instantiates a new VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the total aggregated remote connection usage of a Cloud PC during a given time span.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpcreports-gettotalaggregatedremoteconnectionreports?view=graph-rest-beta
func (m *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToPostRequestInformation get the total aggregated remote connection usage of a Cloud PC during a given time span.
// returns a *RequestInformation when successful
func (m *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder when successful
func (m *VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    return NewVirtualendpointReportsGettotalaggregatedremoteconnectionreportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
