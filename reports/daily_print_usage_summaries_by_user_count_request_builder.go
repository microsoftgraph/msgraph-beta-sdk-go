package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DailyPrintUsageSummariesByUserCountRequestBuilder provides operations to count the resources in the collection.
type DailyPrintUsageSummariesByUserCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyPrintUsageSummariesByUserCountRequestBuilderGetQueryParameters get the number of the resource
type DailyPrintUsageSummariesByUserCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DailyPrintUsageSummariesByUserCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyPrintUsageSummariesByUserCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyPrintUsageSummariesByUserCountRequestBuilderGetQueryParameters
}
// NewDailyPrintUsageSummariesByUserCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewDailyPrintUsageSummariesByUserCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyPrintUsageSummariesByUserCountRequestBuilder) {
    m := &DailyPrintUsageSummariesByUserCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageSummariesByUser/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewDailyPrintUsageSummariesByUserCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewDailyPrintUsageSummariesByUserCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyPrintUsageSummariesByUserCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyPrintUsageSummariesByUserCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *DailyPrintUsageSummariesByUserCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyPrintUsageSummariesByUserCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *DailyPrintUsageSummariesByUserCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyPrintUsageSummariesByUserCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans on 2023-06-13 and will be removed 2023-07-31
func (m *DailyPrintUsageSummariesByUserCountRequestBuilder) WithUrl(rawUrl string)(*DailyPrintUsageSummariesByUserCountRequestBuilder) {
    return NewDailyPrintUsageSummariesByUserCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
