package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
type CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetQueryParameters get salesCreditMemos from financials
type CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetQueryParameters struct {
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
// CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetQueryParameters
}
// BySalesCreditMemoId provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) BySalesCreditMemoId(salesCreditMemoId string)(*CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if salesCreditMemoId != "" {
        urlTplParams["salesCreditMemo%2Did"] = salesCreditMemoId
    }
    return NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// BySalesCreditMemoIdGuid provides operations to manage the salesCreditMemos property of the microsoft.graph.company entity.
// returns a *CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) BySalesCreditMemoIdGuid(salesCreditMemoId i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)(*CompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["salesCreditMemo%2Did"] = salesCreditMemoId.String()
    return NewCompaniesItemSalescreditmemosSalesCreditMemoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderInternal instantiates a new CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) {
    m := &CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder instantiates a new CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder and sets the default values.
func NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompaniesItemSalescreditmemosCountRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) Count()(*CompaniesItemSalescreditmemosCountRequestBuilder) {
    return NewCompaniesItemSalescreditmemosCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get salesCreditMemos from financials
// returns a SalesCreditMemoCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoCollectionResponseable), nil
}
// ToGetRequestInformation get salesCreditMemos from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder when successful
func (m *CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder) {
    return NewCompaniesItemSalescreditmemosSalesCreditMemosRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
