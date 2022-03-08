package vendor_escaped

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i7d6d8b8d7190a5ae78eca79894d25c85fe55630db9a3a082ecf50fb8b91212f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/paymentterm"
    i92a917a0186915b670b0fd1c9b2a2186e591519b280fcf76975baec3ed95f9a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/paymentmethod"
    ia526393eb7b43a4b816611314931e6467a22bc3e148195dde146c51385fc71de "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/picture"
    id3530cd4beaef287f14ef57e5ce261093ee364b54bcd985342b7f6aba98c12a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/currency"
    icca345d7b061ef3cc7f41feb32c4b0a236d5a873a6bf9291ba53dd28a25ba727 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/purchaseinvoices/item/vendor_escaped/picture/item"
)

// VendorRequestBuilder provides operations to manage the vendor property of the microsoft.graph.purchaseInvoice entity.
type VendorRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// VendorRequestBuilderDeleteOptions options for Delete
type VendorRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// VendorRequestBuilderGetOptions options for Get
type VendorRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *VendorRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// VendorRequestBuilderGetQueryParameters get vendor from financials
type VendorRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// VendorRequestBuilderPatchOptions options for Patch
type VendorRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escapedable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewVendorRequestBuilderInternal instantiates a new VendorRequestBuilder and sets the default values.
func NewVendorRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VendorRequestBuilder) {
    m := &VendorRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/purchaseInvoices/{purchaseInvoice_id}/vendor{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVendorRequestBuilder instantiates a new VendorRequestBuilder and sets the default values.
func NewVendorRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VendorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVendorRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property vendor for financials
func (m *VendorRequestBuilder) CreateDeleteRequestInformation(options *VendorRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get vendor from financials
func (m *VendorRequestBuilder) CreateGetRequestInformation(options *VendorRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property vendor in financials
func (m *VendorRequestBuilder) CreatePatchRequestInformation(options *VendorRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *VendorRequestBuilder) Currency()(*id3530cd4beaef287f14ef57e5ce261093ee364b54bcd985342b7f6aba98c12a9.CurrencyRequestBuilder) {
    return id3530cd4beaef287f14ef57e5ce261093ee364b54bcd985342b7f6aba98c12a9.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property vendor for financials
func (m *VendorRequestBuilder) Delete(options *VendorRequestBuilderDeleteOptions)(error) {
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
// Get get vendor from financials
func (m *VendorRequestBuilder) Get(options *VendorRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escapedable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateVendor_escapedFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escapedable), nil
}
// Patch update the navigation property vendor in financials
func (m *VendorRequestBuilder) Patch(options *VendorRequestBuilderPatchOptions)(error) {
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
func (m *VendorRequestBuilder) PaymentMethod()(*i92a917a0186915b670b0fd1c9b2a2186e591519b280fcf76975baec3ed95f9a0.PaymentMethodRequestBuilder) {
    return i92a917a0186915b670b0fd1c9b2a2186e591519b280fcf76975baec3ed95f9a0.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VendorRequestBuilder) PaymentTerm()(*i7d6d8b8d7190a5ae78eca79894d25c85fe55630db9a3a082ecf50fb8b91212f5.PaymentTermRequestBuilder) {
    return i7d6d8b8d7190a5ae78eca79894d25c85fe55630db9a3a082ecf50fb8b91212f5.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VendorRequestBuilder) Picture()(*ia526393eb7b43a4b816611314931e6467a22bc3e148195dde146c51385fc71de.PictureRequestBuilder) {
    return ia526393eb7b43a4b816611314931e6467a22bc3e148195dde146c51385fc71de.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.purchaseInvoices.item.vendor.picture.item collection
func (m *VendorRequestBuilder) PictureById(id string)(*icca345d7b061ef3cc7f41feb32c4b0a236d5a873a6bf9291ba53dd28a25ba727.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return icca345d7b061ef3cc7f41feb32c4b0a236d5a873a6bf9291ba53dd28a25ba727.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
