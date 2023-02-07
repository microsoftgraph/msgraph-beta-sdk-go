package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder provides operations to call the getFormsUserActivityUserCounts method.
type ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal instantiates a new MicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder) {
    m := &ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.getFormsUserActivityUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder instantiates a new MicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getFormsUserActivityUserCounts
func (m *ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// ToGetRequestInformation invoke function getFormsUserActivityUserCounts
func (m *ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetFormsUserActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
