package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder provides operations to call the getComplianceSettingNonComplianceReport method.
type ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderInternal instantiates a new GetComplianceSettingNonComplianceReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder) {
    m := &ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getComplianceSettingNonComplianceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder instantiates a new GetComplianceSettingNonComplianceReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getComplianceSettingNonComplianceReport
func (m *ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder) Post(ctx context.Context, body ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getComplianceSettingNonComplianceReport
func (m *ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetComplianceSettingNonComplianceReportGetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
