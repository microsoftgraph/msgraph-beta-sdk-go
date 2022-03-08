package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines"
    i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/currency"
    iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/paymentterm"
    ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer"
    ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines/item"
)

// SalesOrderItemRequestBuilder provides operations to manage the salesOrders property of the microsoft.graph.company entity.
type SalesOrderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesOrderItemRequestBuilderDeleteOptions options for Delete
type SalesOrderItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesOrderItemRequestBuilderGetOptions options for Get
type SalesOrderItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesOrderItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesOrderItemRequestBuilderGetQueryParameters get salesOrders from financials
type SalesOrderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesOrderItemRequestBuilderPatchOptions options for Patch
type SalesOrderItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrderable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSalesOrderItemRequestBuilderInternal instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewSalesOrderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesOrderItemRequestBuilder) {
    m := &SalesOrderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesOrders/{salesOrder_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesOrderItemRequestBuilder instantiates a new SalesOrderItemRequestBuilder and sets the default values.
func NewSalesOrderItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesOrderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesOrderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesOrders for financials
func (m *SalesOrderItemRequestBuilder) CreateDeleteRequestInformation(options *SalesOrderItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) CreateGetRequestInformation(options *SalesOrderItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) CreatePatchRequestInformation(options *SalesOrderItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *SalesOrderItemRequestBuilder) Currency()(*i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.CurrencyRequestBuilder) {
    return i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesOrderItemRequestBuilder) Customer()(*ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.CustomerRequestBuilder) {
    return ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesOrders for financials
func (m *SalesOrderItemRequestBuilder) Delete(options *SalesOrderItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get salesOrders from financials
func (m *SalesOrderItemRequestBuilder) Get(options *SalesOrderItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrderable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSalesOrderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrderable), nil
}
// Patch update the navigation property salesOrders in financials
func (m *SalesOrderItemRequestBuilder) Patch(options *SalesOrderItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SalesOrderItemRequestBuilder) PaymentTerm()(*iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.PaymentTermRequestBuilder) {
    return iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["salesOrderLine_id"] = id
    }
    return ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d.NewSalesOrderLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
