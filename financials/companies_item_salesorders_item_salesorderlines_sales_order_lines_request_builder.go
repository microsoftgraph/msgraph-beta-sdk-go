package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetQueryParameters get salesOrderLines from financials
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetQueryParameters
}
// BySalesOrderLineId provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
// returns a *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) BySalesOrderLineId(salesOrderLineId string)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if salesOrderLineId != "" {
        urlTplParams["salesOrderLine%2Did"] = salesOrderLineId
    }
    return NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderInternal instantiates a new CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) {
    m := &CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}/salesOrderLines{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder instantiates a new CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder and sets the default values.
func NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemSalesordersItemSalesorderlinesCountRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) Count()(*CompaniesItemSalesordersItemSalesorderlinesCountRequestBuilder) {
    return NewCompaniesItemSalesordersItemSalesorderlinesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesOrderLines from financials
// returns a SalesOrderLineCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderLineCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineCollectionResponseable), nil
}
// ToGetRequestInformation get salesOrderLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder) {
    return NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLinesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
