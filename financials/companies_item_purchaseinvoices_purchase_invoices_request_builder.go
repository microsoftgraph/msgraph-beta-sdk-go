package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
type CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetQueryParameters get purchaseInvoices from financials
type CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetQueryParameters
}
// ByPurchaseInvoiceId provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) ByPurchaseInvoiceId(purchaseInvoiceId string)(*CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if purchaseInvoiceId != "" {
        urlTplParams["purchaseInvoice%2Did"] = purchaseInvoiceId
    }
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPurchaseInvoiceIdGuid provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
// returns a *CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) ByPurchaseInvoiceIdGuid(purchaseInvoiceId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["purchaseInvoice%2Did"] = purchaseInvoiceId.String()
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderInternal instantiates a new CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) {
    m := &CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder instantiates a new CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemPurchaseinvoicesCountRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) Count()(*CompaniesItemPurchaseinvoicesCountRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get purchaseInvoices from financials
// returns a PurchaseInvoiceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePurchaseInvoiceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceCollectionResponseable), nil
}
// ToGetRequestInformation get purchaseInvoices from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder when successful
func (m *CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder) {
    return NewCompaniesItemPurchaseinvoicesPurchaseInvoicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
