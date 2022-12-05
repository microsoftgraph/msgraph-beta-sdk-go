package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder provides operations to manage the customer property of the microsoft.graph.salesOrder entity.
type FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetQueryParameters get customer from financials
type FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetQueryParameters
}
// FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderInternal instantiates a new CustomerRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) {
    m := &FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}/customer{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewFinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder instantiates a new CustomerRequestBuilder and sets the default values.
func NewFinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customer for financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get customer from financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customer in financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Currency provides operations to manage the currency property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) Currency()(*FinancialsCompaniesItemSalesOrdersItemCustomerCurrencyRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customer for financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) Delete(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get customer from financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) Get(ctx context.Context, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// Patch update the navigation property customer in financials
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, requestConfiguration *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Customerable), nil
}
// PaymentMethod provides operations to manage the paymentMethod property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) PaymentMethod()(*FinancialsCompaniesItemSalesOrdersItemCustomerPaymentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PaymentTerm provides operations to manage the paymentTerm property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) PaymentTerm()(*FinancialsCompaniesItemSalesOrdersItemCustomerPaymentTermRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Picture provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) Picture()(*FinancialsCompaniesItemSalesOrdersItemCustomerPictureRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById provides operations to manage the picture property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) PictureById(id string)(*FinancialsCompaniesItemSalesOrdersItemCustomerPicturePictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture%2Did"] = id
    }
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerPicturePictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ShipmentMethod provides operations to manage the shipmentMethod property of the microsoft.graph.customer entity.
func (m *FinancialsCompaniesItemSalesOrdersItemCustomerRequestBuilder) ShipmentMethod()(*FinancialsCompaniesItemSalesOrdersItemCustomerShipmentMethodRequestBuilder) {
    return NewFinancialsCompaniesItemSalesOrdersItemCustomerShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
