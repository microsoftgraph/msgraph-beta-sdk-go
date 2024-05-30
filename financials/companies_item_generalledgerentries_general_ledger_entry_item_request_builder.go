package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
type CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters get generalLedgerEntries from financials
type CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters
}
// Account provides operations to manage the account property of the microsoft.graph.generalLedgerEntry entity.
// returns a *CompaniesItemGeneralledgerentriesItemAccountRequestBuilder when successful
func (m *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) Account()(*CompaniesItemGeneralledgerentriesItemAccountRequestBuilder) {
    return NewCompaniesItemGeneralledgerentriesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderInternal instantiates a new CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder and sets the default values.
func NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) {
    m := &CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/generalLedgerEntries/{generalLedgerEntry%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder instantiates a new CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder and sets the default values.
func NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get generalLedgerEntries from financials
// returns a GeneralLedgerEntryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GeneralLedgerEntryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGeneralLedgerEntryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GeneralLedgerEntryable), nil
}
// ToGetRequestInformation get generalLedgerEntries from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder when successful
func (m *CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder) {
    return NewCompaniesItemGeneralledgerentriesGeneralLedgerEntryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
