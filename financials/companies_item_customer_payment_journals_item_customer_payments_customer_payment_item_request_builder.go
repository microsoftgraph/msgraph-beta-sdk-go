package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder provides operations to manage the customerPayments property of the microsoft.graph.customerPaymentJournal entity.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetQueryParameters get customerPayments from financials
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetQueryParameters
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal instantiates a new CustomerPaymentItemRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    m := &CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/customerPayments/{customerPayment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder instantiates a new CustomerPaymentItemRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Customer provides operations to manage the customer property of the microsoft.graph.customerPayment entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) Customer()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property customerPayments for financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get customerPayments from financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable), nil
}
// Patch update the navigation property customerPayments in financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerPaymentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable), nil
}
// ToDeleteRequestInformation delete navigation property customerPayments for financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get customerPayments from financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customerPayments in financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CustomerPaymentable, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsCustomerPaymentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
