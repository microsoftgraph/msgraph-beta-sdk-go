package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.purchaseInvoice entity.
type CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters get purchaseInvoiceLines from financials
type CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.purchaseInvoiceLine entity.
// returns a *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemAccountRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) Account()(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemAccountRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderInternal instantiates a new CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) {
    m := &CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}/purchaseInvoiceLines/{purchaseInvoiceLine%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder instantiates a new CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get purchaseInvoiceLines from financials
// returns a PurchaseInvoiceLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.purchaseInvoiceLine entity.
// returns a *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemItemRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) Item()(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemItemRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesItemItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property purchaseInvoiceLines in financials
// returns a PurchaseInvoiceLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable), nil
}
// ToGetRequestInformation get purchaseInvoiceLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property purchaseInvoiceLines in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesItemPurchaseinvoicelinesPurchaseInvoiceLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
