package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
type CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetQueryParameters get salesCreditMemos from financials
type CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderInternal instantiates a new CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) {
    m := &CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder instantiates a new CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalescreditmemosItemCurrencyRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) Currency()(*CompaniesItemSalescreditmemosItemCurrencyRequestBuilder) {
    return NewCompaniesItemSalescreditmemosItemCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalescreditmemosItemCustomerRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) Customer()(*CompaniesItemSalescreditmemosItemCustomerRequestBuilder) {
    return NewCompaniesItemSalescreditmemosItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesCreditMemos from financials
// returns a SalesCreditMemoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
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
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, error) {
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
// returns a *CompaniesItemSalescreditmemosItemPaymenttermPaymentTermRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) PaymentTerm()(*CompaniesItemSalescreditmemosItemPaymenttermPaymentTermRequestBuilder) {
    return NewCompaniesItemSalescreditmemosItemPaymenttermPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SalesCreditMemoLines provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
// returns a *CompaniesItemSalescreditmemosItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) SalesCreditMemoLines()(*CompaniesItemSalescreditmemosItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilder) {
    return NewCompaniesItemSalescreditmemosItemSalescreditmemolinesSalesCreditMemoLinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get salesCreditMemos from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoable, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) {
    return NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
