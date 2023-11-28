package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder provides operations to manage the purchaseInvoiceLines property of the microsoft.graph.purchaseInvoice entity.
type CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters get purchaseInvoiceLines from financials
type CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.purchaseInvoiceLine entity.
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Account()(*CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesItemAccountRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal instantiates a new PurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    m := &CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices/{purchaseInvoice%2Did}/purchaseInvoiceLines/{purchaseInvoiceLine%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder instantiates a new PurchaseInvoiceLineItemRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get purchaseInvoiceLines from financials
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Item()(*CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesItemItemRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesItemItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property purchaseInvoiceLines in financials
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceLineable, requestConfiguration *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesItemPurchaseInvoiceLinesPurchaseInvoiceLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
