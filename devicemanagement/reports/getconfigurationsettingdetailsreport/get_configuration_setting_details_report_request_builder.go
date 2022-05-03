package getconfigurationsettingdetailsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigurationSettingDetailsReportRequestBuilder provides operations to call the getConfigurationSettingDetailsReport method.
type GetConfigurationSettingDetailsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetConfigurationSettingDetailsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetConfigurationSettingDetailsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetConfigurationSettingDetailsReportRequestBuilderInternal instantiates a new GetConfigurationSettingDetailsReportRequestBuilder and sets the default values.
func NewGetConfigurationSettingDetailsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationSettingDetailsReportRequestBuilder) {
    m := &GetConfigurationSettingDetailsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationSettingDetailsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationSettingDetailsReportRequestBuilder instantiates a new GetConfigurationSettingDetailsReportRequestBuilder and sets the default values.
func NewGetConfigurationSettingDetailsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationSettingDetailsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationSettingDetailsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationSettingDetailsReport
func (m *GetConfigurationSettingDetailsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationSettingDetailsReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationSettingDetailsReport
func (m *GetConfigurationSettingDetailsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationSettingDetailsReportRequestBodyable, requestConfiguration *GetConfigurationSettingDetailsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getConfigurationSettingDetailsReport
func (m *GetConfigurationSettingDetailsReportRequestBuilder) PostWithResponseHandler(body GetConfigurationSettingDetailsReportRequestBodyable, requestConfiguration *GetConfigurationSettingDetailsReportRequestBuilderPostRequestConfiguration)(GetConfigurationSettingDetailsReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getConfigurationSettingDetailsReport
func (m *GetConfigurationSettingDetailsReportRequestBuilder) PostWithResponseHandler(body GetConfigurationSettingDetailsReportRequestBodyable, requestConfiguration *GetConfigurationSettingDetailsReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetConfigurationSettingDetailsReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationSettingDetailsReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationSettingDetailsReportResponseable), nil
}
