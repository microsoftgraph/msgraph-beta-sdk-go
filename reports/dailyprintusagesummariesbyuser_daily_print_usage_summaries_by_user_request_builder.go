package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
type DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetQueryParameters get dailyPrintUsageSummariesByUser from reports
type DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetQueryParameters
}
// DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrintUsageByUserId provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *DailyprintusagesummariesbyuserPrintUsageByUserItemRequestBuilder when successful
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) ByPrintUsageByUserId(printUsageByUserId string)(*DailyprintusagesummariesbyuserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if printUsageByUserId != "" {
        urlTplParams["printUsageByUser%2Did"] = printUsageByUserId
    }
    return NewDailyprintusagesummariesbyuserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderInternal instantiates a new DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) {
    m := &DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageSummariesByUser{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder instantiates a new DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DailyprintusagesummariesbyuserCountRequestBuilder when successful
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) Count()(*DailyprintusagesummariesbyuserCountRequestBuilder) {
    return NewDailyprintusagesummariesbyuserCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dailyPrintUsageSummariesByUser from reports
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByUserCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintUsageByUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserCollectionResponseable), nil
}
// Post create new navigation property to dailyPrintUsageSummariesByUser for reports
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByUserable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, requestConfiguration *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintUsageByUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable), nil
}
// ToGetRequestInformation get dailyPrintUsageSummariesByUser from reports
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dailyPrintUsageSummariesByUser for reports
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, requestConfiguration *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The dailyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsage navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder when successful
func (m *DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) WithUrl(rawUrl string)(*DailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder) {
    return NewDailyprintusagesummariesbyuserDailyPrintUsageSummariesByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
