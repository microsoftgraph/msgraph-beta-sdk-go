package getremoteconnectionhistoricalreports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetRemoteConnectionHistoricalReportsRequestBuilder provides operations to call the getRemoteConnectionHistoricalReports method.
type GetRemoteConnectionHistoricalReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetRemoteConnectionHistoricalReportsRequestBuilderInternal instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewGetRemoteConnectionHistoricalReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRemoteConnectionHistoricalReportsRequestBuilder) {
    m := &GetRemoteConnectionHistoricalReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports/microsoft.graph.getRemoteConnectionHistoricalReports";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRemoteConnectionHistoricalReportsRequestBuilder instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewGetRemoteConnectionHistoricalReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRemoteConnectionHistoricalReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRemoteConnectionHistoricalReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation get the remote connection history records of a Cloud PC during a given period.
func (m *GetRemoteConnectionHistoricalReportsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body GetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *GetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post get the remote connection history records of a Cloud PC during a given period.
func (m *GetRemoteConnectionHistoricalReportsRequestBuilder) Post(ctx context.Context, body GetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *GetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
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
