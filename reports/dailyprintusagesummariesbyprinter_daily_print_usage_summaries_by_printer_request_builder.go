package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
type DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters get dailyPrintUsageSummariesByPrinter from reports
type DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters struct {
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
// DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetQueryParameters
}
// DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrintUsageByPrinterId provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *DailyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder when successful
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) ByPrintUsageByPrinterId(printUsageByPrinterId string)(*DailyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if printUsageByPrinterId != "" {
        urlTplParams["printUsageByPrinter%2Did"] = printUsageByPrinterId
    }
    return NewDailyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderInternal instantiates a new DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) {
    m := &DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/dailyPrintUsageSummariesByPrinter{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder instantiates a new DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder and sets the default values.
func NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DailyprintusagesummariesbyprinterCountRequestBuilder when successful
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) Count()(*DailyprintusagesummariesbyprinterCountRequestBuilder) {
    return NewDailyprintusagesummariesbyprinterCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get dailyPrintUsageSummariesByPrinter from reports
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByPrinterCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) Get(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintUsageByPrinterCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterCollectionResponseable), nil
}
// Post create new navigation property to dailyPrintUsageSummariesByPrinter for reports
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByPrinterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, requestConfiguration *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintUsageByPrinterFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable), nil
}
// ToGetRequestInformation get dailyPrintUsageSummariesByPrinter from reports
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to dailyPrintUsageSummariesByPrinter for reports
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, requestConfiguration *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The dailyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the dailyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder when successful
func (m *DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) WithUrl(rawUrl string)(*DailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewDailyprintusagesummariesbyprinterDailyPrintUsageSummariesByPrinterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
