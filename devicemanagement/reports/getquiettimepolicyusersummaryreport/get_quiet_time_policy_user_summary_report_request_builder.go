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
// GetQuietTimePolicyUserSummaryReportRequestBuilderPostOptions options for Post
type GetQuietTimePolicyUserSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetQuietTimePolicyUserSummaryReportRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreatePostRequestInformation invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetQuietTimePolicyUserSummaryReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action getQuietTimePolicyUserSummaryReport
func (m *GetQuietTimePolicyUserSummaryReportRequestBuilder) Post(options *GetQuietTimePolicyUserSummaryReportRequestBuilderPostOptions)(GetQuietTimePolicyUserSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetQuietTimePolicyUserSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetQuietTimePolicyUserSummaryReportResponseable), nil
}
