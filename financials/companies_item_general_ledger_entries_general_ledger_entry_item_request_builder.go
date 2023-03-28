package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder provides operations to manage the generalLedgerEntries property of the microsoft.graph.company entity.
type CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters get generalLedgerEntries from financials
type CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetQueryParameters
}
// Account provides operations to manage the account property of the microsoft.graph.generalLedgerEntry entity.
func (m *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) Account()(*CompaniesItemGeneralLedgerEntriesItemAccountRequestBuilder) {
    return NewCompaniesItemGeneralLedgerEntriesItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderInternal instantiates a new GeneralLedgerEntryItemRequestBuilder and sets the default values.
func NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) {
    m := &CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/generalLedgerEntries/{generalLedgerEntry%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder instantiates a new GeneralLedgerEntryItemRequestBuilder and sets the default values.
func NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get generalLedgerEntries from financials
func (m *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GeneralLedgerEntryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
func (m *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemGeneralLedgerEntriesGeneralLedgerEntryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
