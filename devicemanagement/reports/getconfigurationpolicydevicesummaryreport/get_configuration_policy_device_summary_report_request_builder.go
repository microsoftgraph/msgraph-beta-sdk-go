package getconfigurationpolicydevicesummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigurationPolicyDeviceSummaryReportRequestBuilder provides operations to call the getConfigurationPolicyDeviceSummaryReport method.
type GetConfigurationPolicyDeviceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal instantiates a new GetConfigurationPolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    m := &GetConfigurationPolicyDeviceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationPolicyDeviceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationPolicyDeviceSummaryReportRequestBuilder instantiates a new GetConfigurationPolicyDeviceSummaryReportRequestBuilder and sets the default values.
func NewGetConfigurationPolicyDeviceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationPolicyDeviceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationPolicyDeviceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationPolicyDeviceSummaryReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationPolicyDeviceSummaryReportRequestBodyable, requestConfiguration *GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) PostWithResponseHandler(body GetConfigurationPolicyDeviceSummaryReportRequestBodyable, requestConfiguration *GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration)(GetConfigurationPolicyDeviceSummaryReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getConfigurationPolicyDeviceSummaryReport
func (m *GetConfigurationPolicyDeviceSummaryReportRequestBuilder) PostWithResponseHandler(body GetConfigurationPolicyDeviceSummaryReportRequestBodyable, requestConfiguration *GetConfigurationPolicyDeviceSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetConfigurationPolicyDeviceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationPolicyDeviceSummaryReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationPolicyDeviceSummaryReportResponseable), nil
}
