package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder provides operations to call the getMetricsForMfaSignInFailure method.
type ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetQueryParameters invoke function getMetricsForMfaSignInFailure
type ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetQueryParameters struct {
    // Usage: aggregationIntervalInMinutes=@aggregationIntervalInMinutes
    AggregationIntervalInMinutes *int32 `uriparametername:"aggregationIntervalInMinutes"`
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
// ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetQueryParameters
}
// NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderInternal instantiates a new GetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder and sets the default values.
func NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, exclusiveIntervalEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, inclusiveIntervalStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) {
    m := &ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/serviceActivity/getMetricsForMfaSignInFailure(inclusiveIntervalStartDateTime={inclusiveIntervalStartDateTime},exclusiveIntervalEndDateTime={exclusiveIntervalEndDateTime},aggregationIntervalInMinutes=@aggregationIntervalInMinutes){?aggregationIntervalInMinutes*,%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    if exclusiveIntervalEndDateTime != nil {
        m.BaseRequestBuilder.PathParameters["exclusiveIntervalEndDateTime"] = (*exclusiveIntervalEndDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    if inclusiveIntervalStartDateTime != nil {
        m.BaseRequestBuilder.PathParameters["inclusiveIntervalStartDateTime"] = (*inclusiveIntervalStartDateTime).Format(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.RFC3339)
    }
    return m
}
// NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder instantiates a new GetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder and sets the default values.
func NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function getMetricsForMfaSignInFailure
// Deprecated: This method is obsolete. Use GetAsGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse instead.
func (m *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetRequestConfiguration)(ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesResponseable), nil
}
// GetAsGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse invoke function getMetricsForMfaSignInFailure
func (m *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) GetAsGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponse(ctx context.Context, requestConfiguration *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetRequestConfiguration)(ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesGetResponseable), nil
}
// ToGetRequestInformation invoke function getMetricsForMfaSignInFailure
func (m *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) WithUrl(rawUrl string)(*ServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder) {
    return NewServiceActivityGetMetricsForMfaSignInFailureWithInclusiveIntervalStartDateTimeWithExclusiveIntervalEndDateTimeWithAggregationIntervalInMinutesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
