package getquiettimepolicyusersreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetQuietTimePolicyUsersReportRequestBuilder provides operations to call the getQuietTimePolicyUsersReport method.
type GetQuietTimePolicyUsersReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetQuietTimePolicyUsersReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetQuietTimePolicyUsersReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetQuietTimePolicyUsersReportRequestBuilderInternal instantiates a new GetQuietTimePolicyUsersReportRequestBuilder and sets the default values.
func NewGetQuietTimePolicyUsersReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetQuietTimePolicyUsersReportRequestBuilder) {
    m := &GetQuietTimePolicyUsersReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getQuietTimePolicyUsersReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetQuietTimePolicyUsersReportRequestBuilder instantiates a new GetQuietTimePolicyUsersReportRequestBuilder and sets the default values.
func NewGetQuietTimePolicyUsersReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetQuietTimePolicyUsersReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetQuietTimePolicyUsersReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getQuietTimePolicyUsersReport
func (m *GetQuietTimePolicyUsersReportRequestBuilder) CreatePostRequestInformation(body GetQuietTimePolicyUsersReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getQuietTimePolicyUsersReport
func (m *GetQuietTimePolicyUsersReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetQuietTimePolicyUsersReportPostRequestBodyable, requestConfiguration *GetQuietTimePolicyUsersReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getQuietTimePolicyUsersReport
func (m *GetQuietTimePolicyUsersReportRequestBuilder) Post(body GetQuietTimePolicyUsersReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getQuietTimePolicyUsersReport
func (m *GetQuietTimePolicyUsersReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetQuietTimePolicyUsersReportPostRequestBodyable, requestConfiguration *GetQuietTimePolicyUsersReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
