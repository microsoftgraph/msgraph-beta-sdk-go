package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetQueryParameters get salesOrderLines from financials
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesOrderLine entity.
// returns a *CompaniesItemSalesordersItemSalesorderlinesItemAccountRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) Account()(*CompaniesItemSalesordersItemSalesorderlinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesordersItemSalesorderlinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderInternal instantiates a new CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) {
    m := &CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}/salesOrderLines/{salesOrderLine%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder instantiates a new CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get salesOrderLines from financials
// returns a SalesOrderLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesOrderLine entity.
// returns a *CompaniesItemSalesordersItemSalesorderlinesItemItemRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) Item()(*CompaniesItemSalesordersItemSalesorderlinesItemItemRequestBuilder) {
    return NewCompaniesItemSalesordersItemSalesorderlinesItemItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesOrderLines in financials
// returns a SalesOrderLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable), nil
}
// ToGetRequestInformation get salesOrderLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property salesOrderLines in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, requestConfiguration *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder when successful
func (m *CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder) {
    return NewCompaniesItemSalesordersItemSalesorderlinesSalesOrderLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
