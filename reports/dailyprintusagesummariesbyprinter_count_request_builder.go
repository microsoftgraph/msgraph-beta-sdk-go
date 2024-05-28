package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DailyprintusagesummariesbyprinterCountRequestBuilder provides operations to count the resources in the collection.
type DailyprintusagesummariesbyprinterCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyprintusagesummariesbyprinterCountRequestBuilderGetQueryParameters get the number of the resource
type DailyprintusagesummariesbyprinterCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DailyprintusagesummariesbyprinterCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagesummariesbyprinterCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyprintusagesummariesbyprinterCountRequestBuilderGetQueryParameters
}
// NewDailyprintusagesummariesbyprinterCountRequestBuilderInternal instantiates a new DailyprintusagesummariesbyprinterCountRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyprinterCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyprinterCountRequestBuilder) {
    m := &DailyprintusagesummariesbyprinterCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageSummariesByPrinter/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewDailyprintusagesummariesbyprinterCountRequestBuilder instantiates a new DailyprintusagesummariesbyprinterCountRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyprinterCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyprinterCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyprintusagesummariesbyprinterCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagesummariesbyprinterCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyprinterCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *DailyprintusagesummariesbyprinterCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyprinterCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *DailyprintusagesummariesbyprinterCountRequestBuilder when successful
func (m *DailyprintusagesummariesbyprinterCountRequestBuilder) WithUrl(rawUrl string)(*DailyprintusagesummariesbyprinterCountRequestBuilder) {
    return NewDailyprintusagesummariesbyprinterCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
