package getconfigurationsettingsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetConfigurationSettingsReportRequestBuilder provides operations to call the getConfigurationSettingsReport method.
type GetConfigurationSettingsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetConfigurationSettingsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetConfigurationSettingsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetConfigurationSettingsReportRequestBuilderInternal instantiates a new GetConfigurationSettingsReportRequestBuilder and sets the default values.
func NewGetConfigurationSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationSettingsReportRequestBuilder) {
    m := &GetConfigurationSettingsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationSettingsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetConfigurationSettingsReportRequestBuilder instantiates a new GetConfigurationSettingsReportRequestBuilder and sets the default values.
func NewGetConfigurationSettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetConfigurationSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetConfigurationSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationSettingsReport
func (m *GetConfigurationSettingsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationSettingsReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getConfigurationSettingsReport
func (m *GetConfigurationSettingsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetConfigurationSettingsReportRequestBodyable, requestConfiguration *GetConfigurationSettingsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getConfigurationSettingsReport
func (m *GetConfigurationSettingsReportRequestBuilder) PostWithResponseHandler(body GetConfigurationSettingsReportRequestBodyable, requestConfiguration *GetConfigurationSettingsReportRequestBuilderPostRequestConfiguration)(GetConfigurationSettingsReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getConfigurationSettingsReport
func (m *GetConfigurationSettingsReportRequestBuilder) PostWithResponseHandler(body GetConfigurationSettingsReportRequestBodyable, requestConfiguration *GetConfigurationSettingsReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetConfigurationSettingsReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetConfigurationSettingsReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetConfigurationSettingsReportResponseable), nil
}
