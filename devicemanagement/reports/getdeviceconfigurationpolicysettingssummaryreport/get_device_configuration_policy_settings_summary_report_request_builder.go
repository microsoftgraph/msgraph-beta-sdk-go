package getdeviceconfigurationpolicysettingssummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder provides operations to call the getDeviceConfigurationPolicySettingsSummaryReport method.
type GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal instantiates a new GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    m := &GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getDeviceConfigurationPolicySettingsSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder instantiates a new GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder and sets the default values.
func NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) CreatePostRequestInformation(body GetDeviceConfigurationPolicySettingsSummaryReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetDeviceConfigurationPolicySettingsSummaryReportPostRequestBodyable, requestConfiguration *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) Post(body GetDeviceConfigurationPolicySettingsSummaryReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getDeviceConfigurationPolicySettingsSummaryReport
func (m *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetDeviceConfigurationPolicySettingsSummaryReportPostRequestBodyable, requestConfiguration *GetDeviceConfigurationPolicySettingsSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
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
