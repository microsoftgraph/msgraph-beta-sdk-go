package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder provides operations to call the getTotalAggregatedRemoteConnectionReports method.
type VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports/microsoft.graph.getTotalAggregatedRemoteConnectionReports";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder instantiates a new GetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post get the total aggregated remote connection usage of a Cloud PC during a given time span.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpcreports-gettotalaggregatedremoteconnectionreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation get the total aggregated remote connection usage of a Cloud PC during a given time span.
func (m *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilder) ToPostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
