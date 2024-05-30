package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.customerPayment entity.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetQueryParameters get customer from financials
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetQueryParameters
}
// CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderInternal instantiates a new CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) {
    m := &CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/customerPayments/{customerPayment%2Did}/customer{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder instantiates a new CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder and sets the default values.
func NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.customer entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerCurrencyRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) Currency()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerCurrencyRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerCurrencyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property customer for financials
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get customer from financials
// returns a Customerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customer in financials
// returns a Customerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PaymentMethod provides operations to manage the paymentMethod property of the microsoft.graph.customer entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymentmethodPaymentMethodRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) PaymentMethod()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymentmethodPaymentMethodRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymentmethodPaymentMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.customer entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymenttermPaymentTermRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) PaymentTerm()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymenttermPaymentTermRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPaymenttermPaymentTermRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Picture provides operations to manage the picture property of the microsoft.graph.customer entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPictureRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) Picture()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPictureRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerPictureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.customer entity.
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerShipmentmethodShipmentMethodRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) ShipmentMethod()(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerShipmentmethodShipmentMethodRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerShipmentmethodShipmentMethodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property customer for financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get customer from financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property customer in financials
// returns a *RequestInformation when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder when successful
func (m *CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) WithUrl(rawUrl string)(*CompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder) {
    return NewCompaniesItemCustomerpaymentjournalsItemCustomerpaymentsItemCustomerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
