package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentterm"
    i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/picture"
    i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/paymentmethod"
    i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/currency"
    i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/shipmentmethod"
    i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/customers/item/picture/item"
)

// CustomerItemRequestBuilder provides operations to manage the customers property of the microsoft.graph.company entity.
type CustomerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustomerItemRequestBuilderDeleteOptions options for Delete
type CustomerItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerItemRequestBuilderGetOptions options for Get
type CustomerItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustomerItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustomerItemRequestBuilderGetQueryParameters get customers from financials
type CustomerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CustomerItemRequestBuilderPatchOptions options for Patch
type CustomerItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customerable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCustomerItemRequestBuilderInternal instantiates a new CustomerItemRequestBuilder and sets the default values.
func NewCustomerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerItemRequestBuilder) {
    m := &CustomerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/customers/{customer_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustomerItemRequestBuilder instantiates a new CustomerItemRequestBuilder and sets the default values.
func NewCustomerItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustomerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustomerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property customers for financials
func (m *CustomerItemRequestBuilder) CreateDeleteRequestInformation(options *CustomerItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CustomerItemRequestBuilder) CreateGetRequestInformation(options *CustomerItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CustomerItemRequestBuilder) CreatePatchRequestInformation(options *CustomerItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CustomerItemRequestBuilder) Currency()(*i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.CurrencyRequestBuilder) {
    return i6bc3953952f810373b911374975be4e9a9c21d03cf9ced2448cbf983453be4bd.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customers for financials
func (m *CustomerItemRequestBuilder) Delete(options *CustomerItemRequestBuilderDeleteOptions)(error) {
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
// Get get customers from financials
func (m *CustomerItemRequestBuilder) Get(options *CustomerItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCustomerFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Customerable), nil
}
// Patch update the navigation property customers in financials
func (m *CustomerItemRequestBuilder) Patch(options *CustomerItemRequestBuilderPatchOptions)(error) {
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
func (m *CustomerItemRequestBuilder) PaymentMethod()(*i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.PaymentMethodRequestBuilder) {
    return i474632449d4a2f3329d90a51c09385ccf74639e95385a607b9e07138278f97c9.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerItemRequestBuilder) PaymentTerm()(*i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.PaymentTermRequestBuilder) {
    return i173cb04239565eb79ec2ce13140615c6eafd7a121e0e73e41f5d0229e3aabad1.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerItemRequestBuilder) Picture()(*i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.PictureRequestBuilder) {
    return i2fe5ce21826b9c98f2c62e7016604584f4b4baf0a38d60738456107a6a7a626f.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.customers.item.picture.item collection
func (m *CustomerItemRequestBuilder) PictureById(id string)(*i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622.PictureItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return i9baa9dabf27258ce59ba51a2bc2f149d740bc0da5f91de489d977ee7ac8d5622.NewPictureItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustomerItemRequestBuilder) ShipmentMethod()(*i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.ShipmentMethodRequestBuilder) {
    return i7bf4724f04aac20b6c44073110ece9327308609c5e49ae7023a43b232ce90622.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
