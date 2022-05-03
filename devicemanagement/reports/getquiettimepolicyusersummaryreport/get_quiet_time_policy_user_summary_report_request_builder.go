package getquiettimepolicyusersummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetQuietTimePolicyUserSummaryReportRequestBuilder provides operations to call the getQuietTimePolicyUserSummaryReport method.
type GetQuietTimePolicyUserSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetQuietTimePolicyUserSummaryReportRequestBuilderInternal instantiates a new GetQuietTimePolicyUserSummaryReportRequestBuilder and sets the default values.
func NewGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetQuietTimePolicyUserSummaryReportRequestBuilder) {
    m := &GetQuietTimePolicyUserSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getQuietTimePolicyUserSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetQuietTimePolicyUserSummaryReportRequestBuilder instantiates a new GetQuietTimePolicyUserSummaryReportRequestBuilder and sets the default values.
func NewGetQuietTimePolicyUserSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetQuietTimePolicyUserSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetQuietTimePolicyUserSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetQuietTimePolicyUserSummaryReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetQuietTimePolicyUserSummaryReportRequestBodyable, requestConfiguration *GetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) PostWithResponseHandler(body GetQuietTimePolicyUserSummaryReportRequestBodyable, requestConfiguration *GetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration)(GetQuietTimePolicyUserSummaryReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) PostWithResponseHandler(body GetQuietTimePolicyUserSummaryReportRequestBodyable, requestConfiguration *GetQuietTimePolicyUserSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetQuietTimePolicyUserSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetQuietTimePolicyUserSummaryReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetQuietTimePolicyUserSummaryReportResponseable), nil
}
