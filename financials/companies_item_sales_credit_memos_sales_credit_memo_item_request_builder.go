package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
type CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters get salesCreditMemos from financials
type CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal instantiates a new CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    m := &CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder instantiates a new CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalesCreditMemosItemCurrencyRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Currency()(*CompaniesItemSalesCreditMemosItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalesCreditMemosItemCustomerRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Customer()(*CompaniesItemSalesCreditMemosItemCustomerRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesCreditMemos from financials
// returns a SalesCreditMemoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable), nil
}
// Patch update the navigation property salesCreditMemos in financials
// returns a SalesCreditMemoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable), nil
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalesCreditMemosItemPaymentTermRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) PaymentTerm()(*CompaniesItemSalesCreditMemosItemPaymentTermRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) SalesCreditMemoLines()(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get salesCreditMemos from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property salesCreditMemos in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosSalesCreditMemoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
