package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines"
    i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/currency"
    iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/paymentterm"
    ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer"
    ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/salesorderlines/item"
)

// SalesOrderRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesOrders\{salesOrder-id}
type SalesOrderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesOrderRequestBuilderDeleteOptions options for Delete
type SalesOrderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesOrderRequestBuilderGetOptions options for Get
type SalesOrderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesOrderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesOrderRequestBuilderGetQueryParameters get salesOrders from financials
type SalesOrderRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesOrderRequestBuilderPatchOptions options for Patch
type SalesOrderRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSalesOrderRequestBuilderInternal instantiates a new SalesOrderRequestBuilder and sets the default values.
func NewSalesOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesOrderRequestBuilder) {
    m := &SalesOrderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesOrders/{salesOrder_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesOrderRequestBuilder instantiates a new SalesOrderRequestBuilder and sets the default values.
func NewSalesOrderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesOrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesOrders for financials
func (m *SalesOrderRequestBuilder) CreateDeleteRequestInformation(options *SalesOrderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesOrderRequestBuilder) CreateGetRequestInformation(options *SalesOrderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesOrderRequestBuilder) CreatePatchRequestInformation(options *SalesOrderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesOrderRequestBuilder) Currency()(*i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.CurrencyRequestBuilder) {
    return i6b4cd7f670148ff31b09be2a77ce4b56088969f3b75c3798f44df64d7aa566c4.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesOrderRequestBuilder) Customer()(*ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.CustomerRequestBuilder) {
    return ibadb4e5291769990863464c32ca114ea1f6f63e02c97cd3d8b4a41cf346a6191.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesOrders for financials
func (m *SalesOrderRequestBuilder) Delete(options *SalesOrderRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get salesOrders from financials
func (m *SalesOrderRequestBuilder) Get(options *SalesOrderRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesOrder() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesOrder), nil
}
// Patch update the navigation property salesOrders in financials
func (m *SalesOrderRequestBuilder) Patch(options *SalesOrderRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *SalesOrderRequestBuilder) PaymentTerm()(*iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.PaymentTermRequestBuilder) {
    return iacc2dab86f622a8885ab0684a56e478f563750e45e00361ff62aaeb6b0a4da2e.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesOrderRequestBuilder) SalesOrderLines()(*i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d.SalesOrderLinesRequestBuilder) {
    return i209641ff30673fdb09ca02368f6bdc96fd38d1226fdc4e469075752b2475704d.NewSalesOrderLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesOrderLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrders.item.salesOrderLines.item collection
func (m *SalesOrderRequestBuilder) SalesOrderLinesById(id string)(*ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d.SalesOrderLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesOrderLine_id"] = id
    }
    return ic21a121310c89b15c2844b5275d312e77910ddae3d153823d924b5e4b6fa242d.NewSalesOrderLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
