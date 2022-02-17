package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/post"
    i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped"
    i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/currency"
    ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines"
    i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines/item"
)

// PurchaseInvoiceRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\purchaseInvoices\{purchaseInvoice-id}
type PurchaseInvoiceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PurchaseInvoiceRequestBuilderDeleteOptions options for Delete
type PurchaseInvoiceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PurchaseInvoiceRequestBuilderGetOptions options for Get
type PurchaseInvoiceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PurchaseInvoiceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PurchaseInvoiceRequestBuilderGetQueryParameters get purchaseInvoices from financials
type PurchaseInvoiceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PurchaseInvoiceRequestBuilderPatchOptions options for Patch
type PurchaseInvoiceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPurchaseInvoiceRequestBuilderInternal instantiates a new PurchaseInvoiceRequestBuilder and sets the default values.
func NewPurchaseInvoiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PurchaseInvoiceRequestBuilder) {
    m := &PurchaseInvoiceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/purchaseInvoices/{purchaseInvoice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPurchaseInvoiceRequestBuilder instantiates a new PurchaseInvoiceRequestBuilder and sets the default values.
func NewPurchaseInvoiceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PurchaseInvoiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPurchaseInvoiceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property purchaseInvoices for financials
func (m *PurchaseInvoiceRequestBuilder) CreateDeleteRequestInformation(options *PurchaseInvoiceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get purchaseInvoices from financials
func (m *PurchaseInvoiceRequestBuilder) CreateGetRequestInformation(options *PurchaseInvoiceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property purchaseInvoices in financials
func (m *PurchaseInvoiceRequestBuilder) CreatePatchRequestInformation(options *PurchaseInvoiceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PurchaseInvoiceRequestBuilder) Currency()(*i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272.CurrencyRequestBuilder) {
    return i66ab53c3b2ddae75121633bbd823599d4e4e48fe876a2f914d5a57f2e241b272.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property purchaseInvoices for financials
func (m *PurchaseInvoiceRequestBuilder) Delete(options *PurchaseInvoiceRequestBuilderDeleteOptions)(error) {
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
// Get get purchaseInvoices from financials
func (m *PurchaseInvoiceRequestBuilder) Get(options *PurchaseInvoiceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPurchaseInvoice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoice), nil
}
// Patch update the navigation property purchaseInvoices in financials
func (m *PurchaseInvoiceRequestBuilder) Patch(options *PurchaseInvoiceRequestBuilderPatchOptions)(error) {
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
func (m *PurchaseInvoiceRequestBuilder) Post()(*i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90.PostRequestBuilder) {
    return i3d86c6bd7b31ece7f9e0d0acb1882ec4a41b7e93cf9d45845efb984eab9dfe90.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PurchaseInvoiceRequestBuilder) PurchaseInvoiceLines()(*ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a.PurchaseInvoiceLinesRequestBuilder) {
    return ic84ce80e57a95f249407c29d4b5fec4a8c0714292046ec3ba7c175bab69f542a.NewPurchaseInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PurchaseInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.purchaseInvoices.item.purchaseInvoiceLines.item collection
func (m *PurchaseInvoiceRequestBuilder) PurchaseInvoiceLinesById(id string)(*i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb.PurchaseInvoiceLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["purchaseInvoiceLine_id"] = id
    }
    return i25bfa50bef1bcf1c425633cb974de803465fa55f263668c4ae6bb8fa47aaf7cb.NewPurchaseInvoiceLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PurchaseInvoiceRequestBuilder) Vendor_escaped()(*i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3.VendorRequestBuilder) {
    return i48c88f94a8a21b35ebe6d7459f6ba6cd96a3d5e75749e3dd372526fbcb96f6d3.NewVendorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
