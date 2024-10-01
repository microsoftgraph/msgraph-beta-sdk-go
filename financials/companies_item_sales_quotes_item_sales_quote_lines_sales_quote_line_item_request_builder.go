package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder provides operations to manage the salesQuoteLines property of the microsoft.graph.salesQuote entity.
type CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetQueryParameters get salesQuoteLines from financials
type CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesQuoteLine entity.
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) Account()(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal instantiates a new CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    m := &CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesQuotes/{salesQuote%2Did}/salesQuoteLines/{salesQuoteLine%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder instantiates a new CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get salesQuoteLines from financials
// returns a SalesQuoteLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesQuoteLine entity.
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) Item()(*CompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesItemItem_EscapedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesQuoteLines in financials
// returns a SalesQuoteLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesQuoteLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable), nil
}
// ToGetRequestInformation get salesQuoteLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property salesQuoteLines in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesQuoteLineable, requestConfiguration *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder when successful
func (m *CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder) {
    return NewCompaniesItemSalesQuotesItemSalesQuoteLinesSalesQuoteLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
