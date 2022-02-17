package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentterm"
    i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/picture"
    i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentmethod"
    i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/currency"
    i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/shipmentmethod"
)

// CustomerRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\customers\{customer-id}
type CustomerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustomerRequestBuilderDeleteOptions options for Delete
type CustomerRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerRequestBuilderGetOptions options for Get
type CustomerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustomerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerRequestBuilderGetQueryParameters get customers from financials
type CustomerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CustomerRequestBuilderPatchOptions options for Patch
type CustomerRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCustomerRequestBuilderInternal instantiates a new CustomerRequestBuilder and sets the default values.
func NewCustomerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerRequestBuilder) {
    m := &CustomerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/customers/{customer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomerRequestBuilder instantiates a new CustomerRequestBuilder and sets the default values.
func NewCustomerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customers for financials
func (m *CustomerRequestBuilder) CreateDeleteRequestInformation(options *CustomerRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get customers from financials
func (m *CustomerRequestBuilder) CreateGetRequestInformation(options *CustomerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property customers in financials
func (m *CustomerRequestBuilder) CreatePatchRequestInformation(options *CustomerRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CustomerRequestBuilder) Currency()(*i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.CurrencyRequestBuilder) {
    return i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customers for financials
func (m *CustomerRequestBuilder) Delete(options *CustomerRequestBuilderDeleteOptions)(error) {
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
// Get get customers from financials
func (m *CustomerRequestBuilder) Get(options *CustomerRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCustomer() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customer), nil
}
// Patch update the navigation property customers in financials
func (m *CustomerRequestBuilder) Patch(options *CustomerRequestBuilderPatchOptions)(error) {
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
func (m *CustomerRequestBuilder) PaymentMethod()(*i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.PaymentMethodRequestBuilder) {
    return i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerRequestBuilder) PaymentTerm()(*i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.PaymentTermRequestBuilder) {
    return i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerRequestBuilder) Picture()(*i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.PictureRequestBuilder) {
    return i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customers.item.picture.item collection
func (m *CustomerRequestBuilder) PictureById(id string)(*i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.PictureRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.NewPictureRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustomerRequestBuilder) ShipmentMethod()(*i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.ShipmentMethodRequestBuilder) {
    return i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
