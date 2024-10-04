package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder provides operations to manage the salesOrderLines property of the microsoft.graph.salesOrder entity.
type CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetQueryParameters get salesOrderLines from financials
type CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesOrderLine entity.
// returns a *CompaniesItemSalesOrdersItemSalesOrderLinesItemAccountRequestBuilder when successful
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) Account()(*CompaniesItemSalesOrdersItemSalesOrderLinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal instantiates a new CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    m := &CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}/salesOrderLines/{salesOrderLine%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder instantiates a new CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get salesOrderLines from financials
// returns a SalesOrderLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, error) {
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
// returns a *CompaniesItemSalesOrdersItemSalesOrderLinesItemItem_EscapedRequestBuilder when successful
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) Item()(*CompaniesItemSalesOrdersItemSalesOrderLinesItemItem_EscapedRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesItemItem_EscapedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesOrderLines in financials
// returns a SalesOrderLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, requestConfiguration *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, error) {
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
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderLineable, requestConfiguration *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder when successful
func (m *CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder) {
    return NewCompaniesItemSalesOrdersItemSalesOrderLinesSalesOrderLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
