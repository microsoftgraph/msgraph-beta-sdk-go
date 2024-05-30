package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder provides operations to manage the agedAccountsPayable property of the microsoft.graph.company entity.
type CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetQueryParameters get agedAccountsPayable from financials
type CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetQueryParameters
}
// NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderInternal instantiates a new CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder and sets the default values.
func NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) {
    m := &CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/agedAccountsPayable/{agedAccountsPayable%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder instantiates a new CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder and sets the default values.
func NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get agedAccountsPayable from financials
// returns a AgedAccountsPayableable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsPayableable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAgedAccountsPayableFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AgedAccountsPayableable), nil
}
// ToGetRequestInformation get agedAccountsPayable from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder when successful
func (m *CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder) {
    return NewCompaniesItemAgedaccountspayableAgedAccountsPayableItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
