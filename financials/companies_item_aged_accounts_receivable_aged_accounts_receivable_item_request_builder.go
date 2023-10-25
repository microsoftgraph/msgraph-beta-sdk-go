package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder provides operations to manage the agedAccountsReceivable property of the microsoft.graph.company entity.
type CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters get agedAccountsReceivable from financials
type CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetQueryParameters
}
// NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal instantiates a new AgedAccountsReceivableItemRequestBuilder and sets the default values.
func NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    m := &CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/agedAccountsReceivable/{agedAccountsReceivable%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder instantiates a new AgedAccountsReceivableItemRequestBuilder and sets the default values.
func NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get agedAccountsReceivable from financials
func (m *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsReceivableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgedAccountsReceivableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsReceivableable), nil
}
// ToGetRequestInformation get agedAccountsReceivable from financials
func (m *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder) {
    return NewCompaniesItemAgedAccountsReceivableAgedAccountsReceivableItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
