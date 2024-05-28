package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder provides operations to manage the account property of the microsoft.graph.customerPaymentJournal entity.
type CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetQueryParameters get account from financials
type CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetQueryParameters
}
// NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderInternal instantiates a new CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) {
    m := &CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/account{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder instantiates a new CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get account from financials
// returns a Accountable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Accountable, error) {
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
func (m *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemAccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
