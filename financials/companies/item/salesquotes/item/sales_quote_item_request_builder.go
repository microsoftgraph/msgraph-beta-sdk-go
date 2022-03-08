package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/currency"
    i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/paymentterm"
    i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/customer"
    i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/makeinvoice"
    if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/send"
    iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/shipmentmethod"
    iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines"
    id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines/item"
)

// SalesQuoteItemRequestBuilder provides operations to manage the salesQuotes property of the microsoft.graph.company entity.
type SalesQuoteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesQuoteItemRequestBuilderDeleteOptions options for Delete
type SalesQuoteItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesQuoteItemRequestBuilderGetOptions options for Get
type SalesQuoteItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesQuoteItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesQuoteItemRequestBuilderGetQueryParameters get salesQuotes from financials
type SalesQuoteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesQuoteItemRequestBuilderPatchOptions options for Patch
type SalesQuoteItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSalesQuoteItemRequestBuilderInternal instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewSalesQuoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteItemRequestBuilder) {
    m := &SalesQuoteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesQuotes/{salesQuote_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesQuoteItemRequestBuilder instantiates a new SalesQuoteItemRequestBuilder and sets the default values.
func NewSalesQuoteItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesQuoteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesQuotes for financials
func (m *SalesQuoteItemRequestBuilder) CreateDeleteRequestInformation(options *SalesQuoteItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get salesQuotes from financials
func (m *SalesQuoteItemRequestBuilder) CreateGetRequestInformation(options *SalesQuoteItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesQuotes in financials
func (m *SalesQuoteItemRequestBuilder) CreatePatchRequestInformation(options *SalesQuoteItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesQuoteItemRequestBuilder) Currency()(*i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.CurrencyRequestBuilder) {
    return i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteItemRequestBuilder) Customer()(*i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.CustomerRequestBuilder) {
    return i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesQuotes for financials
func (m *SalesQuoteItemRequestBuilder) Delete(options *SalesQuoteItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get salesQuotes from financials
func (m *SalesQuoteItemRequestBuilder) Get(options *SalesQuoteItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateSalesQuoteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteable), nil
}
func (m *SalesQuoteItemRequestBuilder) MakeInvoice()(*i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.MakeInvoiceRequestBuilder) {
    return i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.NewMakeInvoiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesQuotes in financials
func (m *SalesQuoteItemRequestBuilder) Patch(options *SalesQuoteItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SalesQuoteItemRequestBuilder) PaymentTerm()(*i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.PaymentTermRequestBuilder) {
    return i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteItemRequestBuilder) SalesQuoteLines()(*iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.SalesQuoteLinesRequestBuilder) {
    return iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.NewSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesQuoteLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesQuotes.item.salesQuoteLines.item collection
func (m *SalesQuoteItemRequestBuilder) SalesQuoteLinesById(id string)(*id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.SalesQuoteLineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesQuoteLine_id"] = id
    }
    return id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.NewSalesQuoteLineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SalesQuoteItemRequestBuilder) Send()(*if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.SendRequestBuilder) {
    return if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteItemRequestBuilder) ShipmentMethod()(*iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.ShipmentMethodRequestBuilder) {
    return iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
