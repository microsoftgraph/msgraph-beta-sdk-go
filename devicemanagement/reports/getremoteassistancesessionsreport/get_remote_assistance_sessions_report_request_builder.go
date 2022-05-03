package getremoteassistancesessionsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetRemoteAssistanceSessionsReportRequestBuilder provides operations to call the getRemoteAssistanceSessionsReport method.
type GetRemoteAssistanceSessionsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetRemoteAssistanceSessionsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetRemoteAssistanceSessionsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetRemoteAssistanceSessionsReportRequestBuilderInternal instantiates a new GetRemoteAssistanceSessionsReportRequestBuilder and sets the default values.
func NewGetRemoteAssistanceSessionsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRemoteAssistanceSessionsReportRequestBuilder) {
    m := &GetRemoteAssistanceSessionsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getRemoteAssistanceSessionsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRemoteAssistanceSessionsReportRequestBuilder instantiates a new GetRemoteAssistanceSessionsReportRequestBuilder and sets the default values.
func NewGetRemoteAssistanceSessionsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetRemoteAssistanceSessionsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRemoteAssistanceSessionsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getRemoteAssistanceSessionsReport
func (m *GetRemoteAssistanceSessionsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetRemoteAssistanceSessionsReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getRemoteAssistanceSessionsReport
func (m *GetRemoteAssistanceSessionsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetRemoteAssistanceSessionsReportRequestBodyable, requestConfiguration *GetRemoteAssistanceSessionsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// PostWithResponseHandler invoke action getRemoteAssistanceSessionsReport
func (m *GetRemoteAssistanceSessionsReportRequestBuilder) PostWithResponseHandler(body GetRemoteAssistanceSessionsReportRequestBodyable, requestConfiguration *GetRemoteAssistanceSessionsReportRequestBuilderPostRequestConfiguration)(GetRemoteAssistanceSessionsReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getRemoteAssistanceSessionsReport
func (m *GetRemoteAssistanceSessionsReportRequestBuilder) PostWithResponseHandler(body GetRemoteAssistanceSessionsReportRequestBodyable, requestConfiguration *GetRemoteAssistanceSessionsReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetRemoteAssistanceSessionsReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetRemoteAssistanceSessionsReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetRemoteAssistanceSessionsReportResponseable), nil
}
