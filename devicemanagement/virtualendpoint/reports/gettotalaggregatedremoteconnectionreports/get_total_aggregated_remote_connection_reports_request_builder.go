package gettotalaggregatedremoteconnectionreports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetTotalAggregatedRemoteConnectionReportsRequestBuilder provides operations to call the getTotalAggregatedRemoteConnectionReports method.
type GetTotalAggregatedRemoteConnectionReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal instantiates a new GetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    m := &GetTotalAggregatedRemoteConnectionReportsRequestBuilder{
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
// NewGetTotalAggregatedRemoteConnectionReportsRequestBuilder instantiates a new GetTotalAggregatedRemoteConnectionReportsRequestBuilder and sets the default values.
func NewGetTotalAggregatedRemoteConnectionReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTotalAggregatedRemoteConnectionReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTotalAggregatedRemoteConnectionReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation get the total aggregated remote connection usage of a Cloud PC during a given time span.
func (m *GetTotalAggregatedRemoteConnectionReportsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body GetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *GetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post get the total aggregated remote connection usage of a Cloud PC during a given time span.
func (m *GetTotalAggregatedRemoteConnectionReportsRequestBuilder) Post(ctx context.Context, body GetTotalAggregatedRemoteConnectionReportsPostRequestBodyable, requestConfiguration *GetTotalAggregatedRemoteConnectionReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
