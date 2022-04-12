package item

import (
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
// SalesOrderItemRequestBuilderDeleteOptions options for Delete
type SalesOrderItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SalesOrderItemRequestBuilderGetOptions options for Get
type SalesOrderItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SalesOrderItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// SalesOrderItemRequestBuilderGetQueryParameters get salesOrders from financials
type SalesOrderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SalesOrderItemRequestBuilderPatchOptions options for Patch
type SalesOrderItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
// CreateDeleteRequestInformation delete navigation property salesOrders for financials
func (m *SalesOrderItemRequestBuilder) CreateDeleteRequestInformation(options *SalesOrderItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) CreateGetRequestInformation(options *SalesOrderItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) CreatePatchRequestInformation(options *SalesOrderItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
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
// Delete delete navigation property salesOrders for financials
func (m *SalesOrderItemRequestBuilder) Delete(options *SalesOrderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) Get(options *SalesOrderItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSalesOrderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SalesOrderable), nil
}
// Patch update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) Patch(options *SalesOrderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
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
