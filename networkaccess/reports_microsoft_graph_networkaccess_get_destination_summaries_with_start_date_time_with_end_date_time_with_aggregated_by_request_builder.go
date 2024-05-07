package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder provides operations to call the getDestinationSummaries method.
type ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters get counts of the visits to the top destination aggregations.
type ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters struct {
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
// ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal instantiates a new ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, aggregatedBy *string, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    m := &ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/reports/microsoft.graph.networkaccess.getDestinationSummaries(startDateTime={startDateTime},endDateTime={endDateTime},aggregatedBy='{aggregatedBy}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if aggregatedBy != nil {
        m.BaseRequestBuilder.PathParameters["aggregatedBy"] = *aggregatedBy
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder instantiates a new ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get get counts of the visits to the top destination aggregations.
// Deprecated: This method is obsolete. Use GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse instead.
// returns a ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable), nil
}
// GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse get counts of the visits to the top destination aggregations.
// returns a ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) GetAsGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable), nil
}
// ToGetRequestInformation get counts of the visits to the top destination aggregations.
// returns a *RequestInformation when successful
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder when successful
func (m *ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) WithUrl(rawUrl string)(*ReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    return NewReportsMicrosoftGraphNetworkaccessGetDestinationSummariesWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
