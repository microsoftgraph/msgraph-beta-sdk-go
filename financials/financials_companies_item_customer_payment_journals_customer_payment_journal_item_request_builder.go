package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder provides operations to manage the customerPaymentJournals property of the microsoft.graph.company entity.
type FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters get customerPaymentJournals from financials
type FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Account provides operations to manage the account property of the microsoft.graph.customerPaymentJournal entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Account()(*FinancialsCompaniesItemCustomerPaymentJournalsItemAccountRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal instantiates a new CustomerPaymentJournalItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    m := &FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder instantiates a new CustomerPaymentJournalItemRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customerPaymentJournals for financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get customerPaymentJournals from financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property customerPaymentJournals in financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CustomerPayments provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CustomerPayments()(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsRequestBuilder) {
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentsById provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) CustomerPaymentsById(id string)(*FinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment%2Did"] = id
    }
    return NewFinancialsCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property customerPaymentJournals for financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get customerPaymentJournals from financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentJournalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable), nil
}
// Patch update the navigation property customerPaymentJournals in financials
func (m *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, requestConfiguration *FinancialsCompaniesItemCustomerPaymentJournalsCustomerPaymentJournalItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentJournalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentJournalable), nil
}
