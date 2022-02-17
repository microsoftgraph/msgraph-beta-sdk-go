package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ie2e5bef67a7e8106830897b4790a818964af4e79807a777a83c0395718133484 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines/item/account"
    ie391dc3382801307bb68740c42b4f1301ab41dd9b59ec3276fbdd441fda52bfa "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/purchaseinvoicelines/item/item"
)

// PurchaseInvoiceLineRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\purchaseInvoices\{purchaseInvoice-id}\purchaseInvoiceLines\{purchaseInvoiceLine-id}
type PurchaseInvoiceLineRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PurchaseInvoiceLineRequestBuilderDeleteOptions options for Delete
type PurchaseInvoiceLineRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PurchaseInvoiceLineRequestBuilderGetOptions options for Get
type PurchaseInvoiceLineRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PurchaseInvoiceLineRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PurchaseInvoiceLineRequestBuilderGetQueryParameters get purchaseInvoiceLines from financials
type PurchaseInvoiceLineRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PurchaseInvoiceLineRequestBuilderPatchOptions options for Patch
type PurchaseInvoiceLineRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoiceLine;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PurchaseInvoiceLineRequestBuilder) Account()(*ie2e5bef67a7e8106830897b4790a818964af4e79807a777a83c0395718133484.AccountRequestBuilder) {
    return ie2e5bef67a7e8106830897b4790a818964af4e79807a777a83c0395718133484.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPurchaseInvoiceLineRequestBuilderInternal instantiates a new PurchaseInvoiceLineRequestBuilder and sets the default values.
func NewPurchaseInvoiceLineRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PurchaseInvoiceLineRequestBuilder) {
    m := &PurchaseInvoiceLineRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/purchaseInvoices/{purchaseInvoice_id}/purchaseInvoiceLines/{purchaseInvoiceLine_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPurchaseInvoiceLineRequestBuilder instantiates a new PurchaseInvoiceLineRequestBuilder and sets the default values.
func NewPurchaseInvoiceLineRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PurchaseInvoiceLineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPurchaseInvoiceLineRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property purchaseInvoiceLines for financials
func (m *PurchaseInvoiceLineRequestBuilder) CreateDeleteRequestInformation(options *PurchaseInvoiceLineRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get purchaseInvoiceLines from financials
func (m *PurchaseInvoiceLineRequestBuilder) CreateGetRequestInformation(options *PurchaseInvoiceLineRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property purchaseInvoiceLines in financials
func (m *PurchaseInvoiceLineRequestBuilder) CreatePatchRequestInformation(options *PurchaseInvoiceLineRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property purchaseInvoiceLines for financials
func (m *PurchaseInvoiceLineRequestBuilder) Delete(options *PurchaseInvoiceLineRequestBuilderDeleteOptions)(error) {
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
// Get get purchaseInvoiceLines from financials
func (m *PurchaseInvoiceLineRequestBuilder) Get(options *PurchaseInvoiceLineRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoiceLine, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPurchaseInvoiceLine() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PurchaseInvoiceLine), nil
}
func (m *PurchaseInvoiceLineRequestBuilder) Item()(*ie391dc3382801307bb68740c42b4f1301ab41dd9b59ec3276fbdd441fda52bfa.ItemRequestBuilder) {
    return ie391dc3382801307bb68740c42b4f1301ab41dd9b59ec3276fbdd441fda52bfa.NewItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property purchaseInvoiceLines in financials
func (m *PurchaseInvoiceLineRequestBuilder) Patch(options *PurchaseInvoiceLineRequestBuilderPatchOptions)(error) {
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
