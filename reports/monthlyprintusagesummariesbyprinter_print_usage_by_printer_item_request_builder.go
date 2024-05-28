package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
type MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetQueryParameters get monthlyPrintUsageSummariesByPrinter from reports
type MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetQueryParameters
}
// MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderInternal instantiates a new MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder and sets the default values.
func NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) {
    m := &MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/monthlyPrintUsageSummariesByPrinter/{printUsageByPrinter%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder instantiates a new MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder and sets the default values.
func NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property monthlyPrintUsageSummariesByPrinter for reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get monthlyPrintUsageSummariesByPrinter from reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByPrinterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property monthlyPrintUsageSummariesByPrinter in reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a PrintUsageByPrinterable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property monthlyPrintUsageSummariesByPrinter for reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get monthlyPrintUsageSummariesByPrinter from reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property monthlyPrintUsageSummariesByPrinter in reports
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *RequestInformation when successful
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByPrinterable, requestConfiguration *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Deprecated: The monthlyPrintUsageSummariesByPrinter navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByPrinter navigation property instead of this. as of 2023-06/Tasks_And_Plans
// returns a *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder when successful
func (m *MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) WithUrl(rawUrl string)(*MonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder) {
    return NewMonthlyprintusagesummariesbyprinterPrintUsageByPrinterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
