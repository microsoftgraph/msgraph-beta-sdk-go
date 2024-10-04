package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder provides operations to manage the salesCreditMemoLines property of the microsoft.graph.salesCreditMemo entity.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters get salesCreditMemoLines from financials
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetQueryParameters
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.salesCreditMemoLine entity.
// returns a *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Account()(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal instantiates a new CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    m := &CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder instantiates a new CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get salesCreditMemoLines from financials
// returns a SalesCreditMemoLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable), nil
}
// Item provides operations to manage the item property of the microsoft.graph.salesCreditMemoLine entity.
// returns a *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItem_EscapedRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Item()(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItem_EscapedRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItem_EscapedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property salesCreditMemoLines in financials
// returns a SalesCreditMemoLineable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesCreditMemoLineFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable), nil
}
// ToGetRequestInformation get salesCreditMemoLines from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property salesCreditMemoLines in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesCreditMemoLineable, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesSalesCreditMemoLineItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
