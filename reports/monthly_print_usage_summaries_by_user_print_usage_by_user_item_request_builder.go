package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
type MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetQueryParameters get monthlyPrintUsageSummariesByUser from reports
type MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetQueryParameters
}
// MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal instantiates a new PrintUsageByUserItemRequestBuilder and sets the default values.
func NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    m := &MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/monthlyPrintUsageSummariesByUser/{printUsageByUser%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder instantiates a new PrintUsageByUserItemRequestBuilder and sets the default values.
func NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property monthlyPrintUsageSummariesByUser for reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get monthlyPrintUsageSummariesByUser from reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the navigation property monthlyPrintUsageSummariesByUser in reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete navigation property monthlyPrintUsageSummariesByUser for reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get monthlyPrintUsageSummariesByUser from reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property monthlyPrintUsageSummariesByUser in reports
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintUsageByUserable, requestConfiguration *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The monthlyPrintUsageSummariesByUser navigation property is deprecated and will stop returning data on July 31, 2023. Please use the monthlyPrintUsageByUser navigation property instead of this. as of 2023-06/Tasks_And_Plans
func (m *MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) WithUrl(rawUrl string)(*MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
