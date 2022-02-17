package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2e31a6bfa29e65b54c4f094e1b3d885ad298ea0bdc488715087e71090fc8b390 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors/item/paymentmethod"
    i4f2a79af1f797f4101be13d93f559cd58600612aba723188107968599c252dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors/item/currency"
    i5b37bdd8a81f4c60eea30be142234860667502d1f43a06699b1c5170e7d5cb26 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors/item/picture"
    i6aa6ac55925af7fd5c33a37eeb3aae34a47657aed51f36c14cacd2c0beb65b54 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/vendors/item/paymentterm"
)

// VendorRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\vendors\{vendor-id}
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
// VendorRequestBuilderGetQueryParameters get vendors from financials
type VendorRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// VendorRequestBuilderPatchOptions options for Patch
type VendorRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escaped;
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
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/vendors/{vendor_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVendorRequestBuilder instantiates a new VendorRequestBuilder and sets the default values.
func NewVendorRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*VendorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVendorRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property vendors for financials
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
// CreateGetRequestInformation get vendors from financials
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
// CreatePatchRequestInformation update the navigation property vendors in financials
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
func (m *VendorRequestBuilder) Currency()(*i4f2a79af1f797f4101be13d93f559cd58600612aba723188107968599c252dcb.CurrencyRequestBuilder) {
    return i4f2a79af1f797f4101be13d93f559cd58600612aba723188107968599c252dcb.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property vendors for financials
func (m *VendorRequestBuilder) Delete(options *VendorRequestBuilderDeleteOptions)(error) {
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
// Get get vendors from financials
func (m *VendorRequestBuilder) Get(options *VendorRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escaped, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewVendor_escaped() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Vendor_escaped), nil
}
// Patch update the navigation property vendors in financials
func (m *VendorRequestBuilder) Patch(options *VendorRequestBuilderPatchOptions)(error) {
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
func (m *VendorRequestBuilder) PaymentMethod()(*i2e31a6bfa29e65b54c4f094e1b3d885ad298ea0bdc488715087e71090fc8b390.PaymentMethodRequestBuilder) {
    return i2e31a6bfa29e65b54c4f094e1b3d885ad298ea0bdc488715087e71090fc8b390.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VendorRequestBuilder) PaymentTerm()(*i6aa6ac55925af7fd5c33a37eeb3aae34a47657aed51f36c14cacd2c0beb65b54.PaymentTermRequestBuilder) {
    return i6aa6ac55925af7fd5c33a37eeb3aae34a47657aed51f36c14cacd2c0beb65b54.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *VendorRequestBuilder) Picture()(*i5b37bdd8a81f4c60eea30be142234860667502d1f43a06699b1c5170e7d5cb26.PictureRequestBuilder) {
    return i5b37bdd8a81f4c60eea30be142234860667502d1f43a06699b1c5170e7d5cb26.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.vendors.item.picture.item collection
func (m *VendorRequestBuilder) PictureById(id string)(*i5b37bdd8a81f4c60eea30be142234860667502d1f43a06699b1c5170e7d5cb26.PictureRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return i5b37bdd8a81f4c60eea30be142234860667502d1f43a06699b1c5170e7d5cb26.NewPictureRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
