package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
type CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters get customerPaymentJournals from financials
type CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters
}
// CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.customerPaymentJournal entity.
// returns a *CompaniesItemCustomerPaymentJournalsItemAccountRequestBuilder when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Account()(*CompaniesItemCustomerPaymentJournalsItemAccountRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemAccountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal instantiates a new CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    m := &CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder instantiates a new CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
// returns a *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsRequestBuilder when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CustomerPayments()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property customerPaymentJournals for financials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get customerPaymentJournals from financials
// returns a CustomerPaymentJournalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentJournalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable), nil
}
// Patch update the navigation property customerPaymentJournals in financials
// returns a CustomerPaymentJournalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentJournalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable), nil
}
// ToDeleteRequestInformation delete navigation property customerPaymentJournals for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get customerPaymentJournals from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customerPaymentJournals in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, requestConfiguration *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder when successful
func (m *CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
