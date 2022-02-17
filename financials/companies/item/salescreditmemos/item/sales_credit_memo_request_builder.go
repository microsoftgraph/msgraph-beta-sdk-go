package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/customer"
    i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines"
    i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/currency"
    i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/paymentterm"
    i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item"
)

// SalesCreditMemoRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesCreditMemos\{salesCreditMemo-id}
type SalesCreditMemoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesCreditMemoRequestBuilderDeleteOptions options for Delete
type SalesCreditMemoRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesCreditMemoRequestBuilderGetOptions options for Get
type SalesCreditMemoRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesCreditMemoRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesCreditMemoRequestBuilderGetQueryParameters get salesCreditMemos from financials
type SalesCreditMemoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesCreditMemoRequestBuilderPatchOptions options for Patch
type SalesCreditMemoRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemo;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSalesCreditMemoRequestBuilderInternal instantiates a new SalesCreditMemoRequestBuilder and sets the default values.
func NewSalesCreditMemoRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesCreditMemoRequestBuilder) {
    m := &SalesCreditMemoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesCreditMemos/{salesCreditMemo_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesCreditMemoRequestBuilder instantiates a new SalesCreditMemoRequestBuilder and sets the default values.
func NewSalesCreditMemoRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesCreditMemoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesCreditMemoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesCreditMemos for financials
func (m *SalesCreditMemoRequestBuilder) CreateDeleteRequestInformation(options *SalesCreditMemoRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get salesCreditMemos from financials
func (m *SalesCreditMemoRequestBuilder) CreateGetRequestInformation(options *SalesCreditMemoRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesCreditMemos in financials
func (m *SalesCreditMemoRequestBuilder) CreatePatchRequestInformation(options *SalesCreditMemoRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesCreditMemoRequestBuilder) Currency()(*i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859.CurrencyRequestBuilder) {
    return i298450934317995f84337c22e16db4039037fb67ba1a9500bf1e697c652cc859.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesCreditMemoRequestBuilder) Customer()(*i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189.CustomerRequestBuilder) {
    return i198b60984c102ac71de04a9e9a675074ba97cf40609c40333010889081621189.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesCreditMemos for financials
func (m *SalesCreditMemoRequestBuilder) Delete(options *SalesCreditMemoRequestBuilderDeleteOptions)(error) {
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
// Get get salesCreditMemos from financials
func (m *SalesCreditMemoRequestBuilder) Get(options *SalesCreditMemoRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemo, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesCreditMemo() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemo), nil
}
// Patch update the navigation property salesCreditMemos in financials
func (m *SalesCreditMemoRequestBuilder) Patch(options *SalesCreditMemoRequestBuilderPatchOptions)(error) {
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
func (m *SalesCreditMemoRequestBuilder) PaymentTerm()(*i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412.PaymentTermRequestBuilder) {
    return i7472604d594d7fc8b26a208c86be9abd17982368a2fd0977a2f07f55ddbda412.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesCreditMemoRequestBuilder) SalesCreditMemoLines()(*i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51.SalesCreditMemoLinesRequestBuilder) {
    return i2871eb11fd553a28e1a220512b521cc37c977c7f73da911a23b028656f22cb51.NewSalesCreditMemoLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesCreditMemoLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesCreditMemos.item.salesCreditMemoLines.item collection
func (m *SalesCreditMemoRequestBuilder) SalesCreditMemoLinesById(id string)(*i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045.SalesCreditMemoLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesCreditMemoLine_id"] = id
    }
    return i45ad6020b7b06f1d1193f8182422eaf641a20591f59d12b7f6704bccd0a06045.NewSalesCreditMemoLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
