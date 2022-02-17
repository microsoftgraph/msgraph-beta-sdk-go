package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/shipmentmethod"
    i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/currency"
    i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines"
    i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancel"
    i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/post"
    i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/cancelandsend"
    i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/send"
    i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/customer"
    i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/paymentterm"
    iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/postandsend"
    if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesinvoices/item/salesinvoicelines/item"
)

// SalesInvoiceRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesInvoices\{salesInvoice-id}
type SalesInvoiceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesInvoiceRequestBuilderDeleteOptions options for Delete
type SalesInvoiceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesInvoiceRequestBuilderGetOptions options for Get
type SalesInvoiceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesInvoiceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesInvoiceRequestBuilderGetQueryParameters get salesInvoices from financials
type SalesInvoiceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesInvoiceRequestBuilderPatchOptions options for Patch
type SalesInvoiceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SalesInvoiceRequestBuilder) Cancel()(*i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.CancelRequestBuilder) {
    return i36d06f0f74d1b6fcb1f7c52c368d55e3ae1c78fea244620a809d74b4050c21f2.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) CancelAndSend()(*i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.CancelAndSendRequestBuilder) {
    return i4029970bec644a3ffd8071814f20a9d8eb5762118c76c4b80b23161fb86c2fee.NewCancelAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSalesInvoiceRequestBuilderInternal instantiates a new SalesInvoiceRequestBuilder and sets the default values.
func NewSalesInvoiceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesInvoiceRequestBuilder) {
    m := &SalesInvoiceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesInvoices/{salesInvoice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesInvoiceRequestBuilder instantiates a new SalesInvoiceRequestBuilder and sets the default values.
func NewSalesInvoiceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesInvoiceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesInvoiceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesInvoices for financials
func (m *SalesInvoiceRequestBuilder) CreateDeleteRequestInformation(options *SalesInvoiceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get salesInvoices from financials
func (m *SalesInvoiceRequestBuilder) CreateGetRequestInformation(options *SalesInvoiceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesInvoices in financials
func (m *SalesInvoiceRequestBuilder) CreatePatchRequestInformation(options *SalesInvoiceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesInvoiceRequestBuilder) Currency()(*i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.CurrencyRequestBuilder) {
    return i0650e7911918a65a74f67fe1c197bdc18b90494b3361ef73af6e6fed5d0e8e47.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Customer()(*i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.CustomerRequestBuilder) {
    return i93777e17bc2372af2aae81772eb24d80da0156546c0cdfca56a28e9618f8d889.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property salesInvoices for financials
func (m *SalesInvoiceRequestBuilder) Delete(options *SalesInvoiceRequestBuilderDeleteOptions)(error) {
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
// Get get salesInvoices from financials
func (m *SalesInvoiceRequestBuilder) Get(options *SalesInvoiceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesInvoice() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesInvoice), nil
}
// Patch update the navigation property salesInvoices in financials
func (m *SalesInvoiceRequestBuilder) Patch(options *SalesInvoiceRequestBuilderPatchOptions)(error) {
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
func (m *SalesInvoiceRequestBuilder) PaymentTerm()(*i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.PaymentTermRequestBuilder) {
    return i973a00f44707ce3fe81e75ce782fefb6c04714e28ca331b622c712a4b29cb6cb.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Post()(*i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.PostRequestBuilder) {
    return i37e5adacee5361adfd45947751eeada85d4fb063399b1012befbf77338d1551b.NewPostRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) PostAndSend()(*iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.PostAndSendRequestBuilder) {
    return iff094cb38c7f93c080b831ba288f86fafe7cb758d1a92b96280ccf5a97b91609.NewPostAndSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) SalesInvoiceLines()(*i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.SalesInvoiceLinesRequestBuilder) {
    return i34ab7468987d894218eb3c0a57bf71468ae3b457fdc7c9a1e1e06dac5693aad7.NewSalesInvoiceLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SalesInvoiceLinesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesInvoices.item.salesInvoiceLines.item collection
func (m *SalesInvoiceRequestBuilder) SalesInvoiceLinesById(id string)(*if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.SalesInvoiceLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["salesInvoiceLine_id"] = id
    }
    return if8797c053079d45780da07f4c01fb650fc8015394e895a3fdf2d677f2013e9ae.NewSalesInvoiceLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) Send()(*i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.SendRequestBuilder) {
    return i73fdc6873e10630216ab19990248e2068cf89ebe3b996094b3b49c3edfb216d3.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesInvoiceRequestBuilder) ShipmentMethod()(*i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.ShipmentMethodRequestBuilder) {
    return i02155747beea3539ff57236a7d1a9ced210439349101163afadd0ca4dddc11ff.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
