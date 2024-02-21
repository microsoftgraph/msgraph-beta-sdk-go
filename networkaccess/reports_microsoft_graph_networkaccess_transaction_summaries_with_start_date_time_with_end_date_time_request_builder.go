package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder provides operations to call the transactionSummaries method.
type ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters invoke function transactionSummaries
type ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters struct {
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
// ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderInternal instantiates a new ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    m := &ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/reports/microsoft.graph.networkaccess.transactionSummaries(startDateTime={startDateTime},endDateTime={endDateTime}){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = (*endDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = (*startDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder instantiates a new ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function transactionSummaries
// Deprecated: This method is obsolete. Use GetAsTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse instead.
// returns a ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeResponseable), nil
}
// GetAsTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse invoke function transactionSummaries
// returns a ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) GetAsTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponse(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeTransactionSummariesWithStartDateTimeWithEndDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function transactionSummaries
// returns a *RequestInformation when successful
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder when successful
func (m *ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) WithUrl(rawUrl string)(*ReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewReportsMicrosoftGraphNetworkaccessTransactionSummariesWithStartDateTimeWithEndDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
