package communications

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder provides operations to call the getPstnOnlineMeetingDialoutReport method.
type CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters get aggregated report of usage and money spent for the audio conferencing dial-out service over a selected period as a collection of pstnOnlineMeetingDialoutReport entries.The report is aggregated by user, user location, destination context (domestic/international), and currency. The report includes:
type CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetQueryParameters
}
// NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderInternal instantiates a new CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, toDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) {
    m := &CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/callRecords/microsoft.graph.callRecords.getPstnOnlineMeetingDialoutReport(fromDateTime={fromDateTime},toDateTime={toDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if fromDateTime != nil {
        m.BaseRequestBuilder.PathParameters["fromDateTime"] = (*fromDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if toDateTime != nil {
        m.BaseRequestBuilder.PathParameters["toDateTime"] = (*toDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder instantiates a new CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder and sets the default values.
func NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get get aggregated report of usage and money spent for the audio conferencing dial-out service over a selected period as a collection of pstnOnlineMeetingDialoutReport entries.The report is aggregated by user, user location, destination context (domestic/international), and currency. The report includes:
// Deprecated: This method is obsolete. Use GetAsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponse instead.
// returns a CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeResponseable), nil
}
// GetAsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponse get aggregated report of usage and money spent for the audio conferencing dial-out service over a selected period as a collection of pstnOnlineMeetingDialoutReport entries.The report is aggregated by user, user location, destination context (domestic/international), and currency. The report includes:
// returns a CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) GetAsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponse(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeGetResponseable), nil
}
// ToGetRequestInformation get aggregated report of usage and money spent for the audio conferencing dial-out service over a selected period as a collection of pstnOnlineMeetingDialoutReport entries.The report is aggregated by user, user location, destination context (domestic/international), and currency. The report includes:
// returns a *RequestInformation when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder when successful
func (m *CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) WithUrl(rawUrl string)(*CallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder) {
    return NewCallrecordsMicrosoftgraphcallrecordsgetpstnonlinemeetingdialoutreportwithfromdatetimewithtodatetimeMicrosoftGraphCallRecordsGetPstnOnlineMeetingDialoutReportWithFromDateTimeWithToDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
