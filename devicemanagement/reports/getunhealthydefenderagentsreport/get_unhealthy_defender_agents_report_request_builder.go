package getunhealthydefenderagentsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetUnhealthyDefenderAgentsReportRequestBuilder provides operations to call the getUnhealthyDefenderAgentsReport method.
type GetUnhealthyDefenderAgentsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetUnhealthyDefenderAgentsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetUnhealthyDefenderAgentsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetUnhealthyDefenderAgentsReportRequestBuilderInternal instantiates a new GetUnhealthyDefenderAgentsReportRequestBuilder and sets the default values.
func NewGetUnhealthyDefenderAgentsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetUnhealthyDefenderAgentsReportRequestBuilder) {
    m := &GetUnhealthyDefenderAgentsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getUnhealthyDefenderAgentsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetUnhealthyDefenderAgentsReportRequestBuilder instantiates a new GetUnhealthyDefenderAgentsReportRequestBuilder and sets the default values.
func NewGetUnhealthyDefenderAgentsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetUnhealthyDefenderAgentsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetUnhealthyDefenderAgentsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getUnhealthyDefenderAgentsReport
func (m *GetUnhealthyDefenderAgentsReportRequestBuilder) CreatePostRequestInformation(body GetUnhealthyDefenderAgentsReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getUnhealthyDefenderAgentsReport
func (m *GetUnhealthyDefenderAgentsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetUnhealthyDefenderAgentsReportPostRequestBodyable, requestConfiguration *GetUnhealthyDefenderAgentsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getUnhealthyDefenderAgentsReport
func (m *GetUnhealthyDefenderAgentsReportRequestBuilder) Post(body GetUnhealthyDefenderAgentsReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getUnhealthyDefenderAgentsReport
func (m *GetUnhealthyDefenderAgentsReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetUnhealthyDefenderAgentsReportPostRequestBodyable, requestConfiguration *GetUnhealthyDefenderAgentsReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
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
