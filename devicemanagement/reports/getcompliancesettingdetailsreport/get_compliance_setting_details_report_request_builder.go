package getcompliancesettingdetailsreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetComplianceSettingDetailsReportRequestBuilder provides operations to call the getComplianceSettingDetailsReport method.
type GetComplianceSettingDetailsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetComplianceSettingDetailsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetComplianceSettingDetailsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetComplianceSettingDetailsReportRequestBuilderInternal instantiates a new GetComplianceSettingDetailsReportRequestBuilder and sets the default values.
func NewGetComplianceSettingDetailsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetComplianceSettingDetailsReportRequestBuilder) {
    m := &GetComplianceSettingDetailsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getComplianceSettingDetailsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetComplianceSettingDetailsReportRequestBuilder instantiates a new GetComplianceSettingDetailsReportRequestBuilder and sets the default values.
func NewGetComplianceSettingDetailsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetComplianceSettingDetailsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetComplianceSettingDetailsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) CreatePostRequestInformation(body GetComplianceSettingDetailsReportPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetComplianceSettingDetailsReportPostRequestBodyable, requestConfiguration *GetComplianceSettingDetailsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) Post(body GetComplianceSettingDetailsReportPostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetComplianceSettingDetailsReportPostRequestBodyable, requestConfiguration *GetComplianceSettingDetailsReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
