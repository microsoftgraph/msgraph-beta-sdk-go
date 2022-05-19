package getconfigmanagerdevicepolicystatusreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigManagerDevicePolicyStatusReportRequestBuilder provides operations to call the getConfigManagerDevicePolicyStatusReport method.
type GetConfigManagerDevicePolicyStatusReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetConfigManagerDevicePolicyStatusReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetConfigManagerDevicePolicyStatusReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal instantiates a new GetConfigManagerDevicePolicyStatusReportRequestBuilder and sets the default values.
func NewGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    m := &GetConfigManagerDevicePolicyStatusReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigManagerDevicePolicyStatusReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigManagerDevicePolicyStatusReportRequestBuilder instantiates a new GetConfigManagerDevicePolicyStatusReportRequestBuilder and sets the default values.
func NewGetConfigManagerDevicePolicyStatusReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigManagerDevicePolicyStatusReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigManagerDevicePolicyStatusReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getConfigManagerDevicePolicyStatusReport
func (m *GetConfigManagerDevicePolicyStatusReportRequestBuilder) CreatePostRequestInformation(body GetConfigManagerDevicePolicyStatusReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigManagerDevicePolicyStatusReport
func (m *GetConfigManagerDevicePolicyStatusReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigManagerDevicePolicyStatusReportPostRequestBodyable, requestConfiguration *GetConfigManagerDevicePolicyStatusReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getConfigManagerDevicePolicyStatusReport
func (m *GetConfigManagerDevicePolicyStatusReportRequestBuilder) Post(body GetConfigManagerDevicePolicyStatusReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getConfigManagerDevicePolicyStatusReport
func (m *GetConfigManagerDevicePolicyStatusReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetConfigManagerDevicePolicyStatusReportPostRequestBodyable, requestConfiguration *GetConfigManagerDevicePolicyStatusReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
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
