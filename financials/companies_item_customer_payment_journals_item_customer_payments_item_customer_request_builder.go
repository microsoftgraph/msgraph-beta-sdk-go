package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.customerPayment entity.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters get customer from financials
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetQueryParameters
}
// CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal instantiates a new CustomerRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) {
    m := &CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/customerPaymentJournals/{customerPaymentJournal%2Did}/customerPayments/{customerPayment%2Did}/customer{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder instantiates a new CustomerRequestBuilder and sets the default values.
func NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// Currency provides operations to manage the currency property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Currency()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerCurrencyRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Delete delete navigation property customer for financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get customer from financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customer in financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PaymentMethod provides operations to manage the paymentMethod property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PaymentMethod()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentMethodRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PaymentTerm()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentTermRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Picture provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) Picture()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPictureRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// PictureById provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) PictureById(id string)(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerPicturePictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.customer entity.
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) ShipmentMethod()(*CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerShipmentMethodRequestBuilder) {
    return NewCompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property customer for financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get customer from financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property customer in financials
func (m *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *CompaniesItemCustomerPaymentJournalsItemCustomerPaymentsItemCustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
