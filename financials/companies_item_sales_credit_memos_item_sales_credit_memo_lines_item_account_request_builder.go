// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder provides operations to manage the account property of the microsoft.graph.salesCreditMemoLine entity.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetQueryParameters get account from financials
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetQueryParameters
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderInternal instantiates a new CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) {
    m := &CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}/account{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder instantiates a new CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get account from financials
// returns a Accountable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Accountable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAccountFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Accountable), nil
}
// ToGetRequestInformation get account from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder when successful
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder) {
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemAccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
