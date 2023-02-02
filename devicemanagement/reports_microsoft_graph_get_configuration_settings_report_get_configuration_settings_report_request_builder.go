package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder provides operations to call the getConfigurationSettingsReport method.
type ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderInternal instantiates a new GetConfigurationSettingsReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder) {
    m := &ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getConfigurationSettingsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder instantiates a new GetConfigurationSettingsReportRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getConfigurationSettingsReport
func (m *ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder) Post(ctx context.Context, body ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
// ToPostRequestInformation invoke action getConfigurationSettingsReport
func (m *ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetConfigurationSettingsReportGetConfigurationSettingsReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
