package print

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
type ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    m := &ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports/microsoft.graph.getSkypeForBusinessParticipantActivityCounts(period='{period}')";
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
// NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder instantiates a new GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSkypeForBusinessParticipantActivityCounts
func (m *ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getSkypeForBusinessParticipantActivityCounts
func (m *ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphGetSkypeForBusinessParticipantActivityCountsWithPeriodGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
