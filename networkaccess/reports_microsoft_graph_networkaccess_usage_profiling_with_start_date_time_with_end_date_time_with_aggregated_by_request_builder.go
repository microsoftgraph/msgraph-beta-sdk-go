package networkaccess

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder provides operations to call the usageProfiling method.
type ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters invoke function usageProfiling
type ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters struct {
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
// ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetQueryParameters
}
// NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal instantiates a new MicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, aggregatedBy *string, endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    m := &ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/reports/microsoft.graph.networkaccess.usageProfiling(startDateTime={startDateTime},endDateTime={endDateTime},aggregatedBy='{aggregatedBy}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
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
// NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder instantiates a new MicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function usageProfiling
// Deprecated: This method is obsolete. Use GetAsUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse instead.
func (m *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByResponseable), nil
}
// GetAsUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse invoke function usageProfiling
func (m *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) GetAsUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponse(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByGetResponseable), nil
}
// ToGetRequestInformation invoke function usageProfiling
func (m *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) WithUrl(rawUrl string)(*ReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder) {
    return NewReportsMicrosoftGraphNetworkaccessUsageProfilingWithStartDateTimeWithEndDateTimeWithAggregatedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
