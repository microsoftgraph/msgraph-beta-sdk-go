package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines"
    i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/currency"
    iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/paymentterm"
    ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer"
    ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines/item"
)

// SalesOrderItemRequestBuilder provides operations to manage the salesOrders property of the microsoft.graph.company entity.
type SalesOrderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SalesOrderItemRequestBuilderGetQueryParameters get salesOrders from financials
type SalesOrderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesOrderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesOrderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesOrderItemRequestBuilderGetQueryParameters
}
// SalesOrderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesOrderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSalesOrderItemRequestBuilderInternal instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewSalesOrderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesOrderItemRequestBuilder) {
    m := &SalesOrderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company%2Did}/salesOrders/{salesOrder%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesOrderItemRequestBuilder instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewSalesOrderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SalesOrderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesOrderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SalesOrderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, requestConfiguration *SalesOrderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Currency the currency property
func (m *SalesOrderItemRequestBuilder) Currency()(*i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.CurrencyRequestBuilder) {
    return i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Customer the customer property
func (m *SalesOrderItemRequestBuilder) Customer()(*ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.CustomerRequestBuilder) {
    return ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SalesOrderItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable), nil
}
// Patch update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, requestConfiguration *SalesOrderItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// PaymentTerm the paymentTerm property
func (m *SalesOrderItemRequestBuilder) PaymentTerm()(*iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.PaymentTermRequestBuilder) {
    return iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLines the salesOrderLines property
func (m *SalesOrderItemRequestBuilder) SalesOrderLines()(*i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d.SalesOrderLinesRequestBuilder) {
    return i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d.NewSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrders.item.salesOrderLines.item collection
func (m *SalesOrderItemRequestBuilder) SalesOrderLinesById(id string)(*ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d.SalesOrderLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine%2Did"] = id
    }
    return ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d.NewSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
