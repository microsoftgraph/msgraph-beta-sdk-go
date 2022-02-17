package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i158ee248967a0aa764abf30309015ccd5fc2a08a2aa5fc76ff38213f29e2782a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpaymentjournals/item/customerpayments"
    icf60ac8c8bdf6d2598e70cbc64335f3349b52f0906822aa33183ce7856795204 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpaymentjournals/item/account"
    ic37aafe0737f76d3603471c2e9ee70925625ac5cf40414bbe1b8b7acfcb5b74a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customerpaymentjournals/item/customerpayments/item"
)

// CustomerPaymentJournalRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\customerPaymentJournals\{customerPaymentJournal-id}
type CustomerPaymentJournalRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustomerPaymentJournalRequestBuilderDeleteOptions options for Delete
type CustomerPaymentJournalRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerPaymentJournalRequestBuilderGetOptions options for Get
type CustomerPaymentJournalRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustomerPaymentJournalRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerPaymentJournalRequestBuilderGetQueryParameters get customerPaymentJournals from financials
type CustomerPaymentJournalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CustomerPaymentJournalRequestBuilderPatchOptions options for Patch
type CustomerPaymentJournalRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomerPaymentJournal;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CustomerPaymentJournalRequestBuilder) Account()(*icf60ac8c8bdf6d2598e70cbc64335f3349b52f0906822aa33183ce7856795204.AccountRequestBuilder) {
    return icf60ac8c8bdf6d2598e70cbc64335f3349b52f0906822aa33183ce7856795204.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCustomerPaymentJournalRequestBuilderInternal instantiates a new CustomerPaymentJournalRequestBuilder and sets the default values.
func NewCustomerPaymentJournalRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerPaymentJournalRequestBuilder) {
    m := &CustomerPaymentJournalRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/customerPaymentJournals/{customerPaymentJournal_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomerPaymentJournalRequestBuilder instantiates a new CustomerPaymentJournalRequestBuilder and sets the default values.
func NewCustomerPaymentJournalRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerPaymentJournalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomerPaymentJournalRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customerPaymentJournals for financials
func (m *CustomerPaymentJournalRequestBuilder) CreateDeleteRequestInformation(options *CustomerPaymentJournalRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get customerPaymentJournals from financials
func (m *CustomerPaymentJournalRequestBuilder) CreateGetRequestInformation(options *CustomerPaymentJournalRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customerPaymentJournals in financials
func (m *CustomerPaymentJournalRequestBuilder) CreatePatchRequestInformation(options *CustomerPaymentJournalRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CustomerPaymentJournalRequestBuilder) CustomerPayments()(*i158ee248967a0aa764abf30309015ccd5fc2a08a2aa5fc76ff38213f29e2782a.CustomerPaymentsRequestBuilder) {
    return i158ee248967a0aa764abf30309015ccd5fc2a08a2aa5fc76ff38213f29e2782a.NewCustomerPaymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomerPaymentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customerPaymentJournals.item.customerPayments.item collection
func (m *CustomerPaymentJournalRequestBuilder) CustomerPaymentsById(id string)(*ic37aafe0737f76d3603471c2e9ee70925625ac5cf40414bbe1b8b7acfcb5b74a.CustomerPaymentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customerPayment_id"] = id
    }
    return ic37aafe0737f76d3603471c2e9ee70925625ac5cf40414bbe1b8b7acfcb5b74a.NewCustomerPaymentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property customerPaymentJournals for financials
func (m *CustomerPaymentJournalRequestBuilder) Delete(options *CustomerPaymentJournalRequestBuilderDeleteOptions)(error) {
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
// Get get customerPaymentJournals from financials
func (m *CustomerPaymentJournalRequestBuilder) Get(options *CustomerPaymentJournalRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomerPaymentJournal, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCustomerPaymentJournal() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CustomerPaymentJournal), nil
}
// Patch update the navigation property customerPaymentJournals in financials
func (m *CustomerPaymentJournalRequestBuilder) Patch(options *CustomerPaymentJournalRequestBuilderPatchOptions)(error) {
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
