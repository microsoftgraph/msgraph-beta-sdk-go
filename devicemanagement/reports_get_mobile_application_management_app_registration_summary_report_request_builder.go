package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder provides operations to call the getMobileApplicationManagementAppRegistrationSummaryReport method.
type ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal instantiates a new GetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder and sets the default values.
func NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    m := &ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getMobileApplicationManagementAppRegistrationSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder instantiates a new GetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder and sets the default values.
func NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getMobileApplicationManagementAppRegistrationSummaryReport
func (m *ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ReportsGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getMobileApplicationManagementAppRegistrationSummaryReport
func (m *ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilder) Post(ctx context.Context, body ReportsGetMobileApplicationManagementAppRegistrationSummaryReportPostRequestBodyable, requestConfiguration *ReportsGetMobileApplicationManagementAppRegistrationSummaryReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
