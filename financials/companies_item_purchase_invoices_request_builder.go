package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemPurchaseInvoicesRequestBuilder provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
type CompaniesItemPurchaseInvoicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemPurchaseInvoicesRequestBuilderGetQueryParameters get purchaseInvoices from financials
type CompaniesItemPurchaseInvoicesRequestBuilderGetQueryParameters struct {
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
// CompaniesItemPurchaseInvoicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemPurchaseInvoicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemPurchaseInvoicesRequestBuilderGetQueryParameters
}
// ByPurchaseInvoiceId provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) ByPurchaseInvoiceId(purchaseInvoiceId string)(*CompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if purchaseInvoiceId != "" {
        urlTplParams["purchaseInvoice%2Did"] = purchaseInvoiceId
    }
    return NewCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByPurchaseInvoiceIdGuid provides operations to manage the purchaseInvoices property of the microsoft.graph.company entity.
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) ByPurchaseInvoiceIdGuid(purchaseInvoiceId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["purchaseInvoice%2Did"] = purchaseInvoiceId.String()
    return NewCompaniesItemPurchaseInvoicesPurchaseInvoiceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemPurchaseInvoicesRequestBuilderInternal instantiates a new PurchaseInvoicesRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoicesRequestBuilder) {
    m := &CompaniesItemPurchaseInvoicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/purchaseInvoices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemPurchaseInvoicesRequestBuilder instantiates a new PurchaseInvoicesRequestBuilder and sets the default values.
func NewCompaniesItemPurchaseInvoicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemPurchaseInvoicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemPurchaseInvoicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) Count()(*CompaniesItemPurchaseInvoicesCountRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get purchaseInvoices from financials
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PurchaseInvoiceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemPurchaseInvoicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompaniesItemPurchaseInvoicesRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemPurchaseInvoicesRequestBuilder) {
    return NewCompaniesItemPurchaseInvoicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
