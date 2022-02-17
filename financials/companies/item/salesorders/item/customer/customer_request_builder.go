package customer

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/paymentmethod"
    i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/paymentterm"
    i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/picture"
    i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/currency"
    ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesorders/item/customer/shipmentmethod"
)

// CustomerRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesOrders\{salesOrder-id}\customer
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
// CustomerRequestBuilderGetQueryParameters get customer from financials
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
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesOrders/{salesOrder_id}/customer{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property customer for financials
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
// CreateGetRequestInformation get customer from financials
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
// CreatePatchRequestInformation update the navigation property customer in financials
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
func (m *CustomerRequestBuilder) Currency()(*i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7.CurrencyRequestBuilder) {
    return i94cb70f6e35ad5abaabad528a25dfd0606df89e3979609cfd5627caee2679bd7.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property customer for financials
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
// Get get customer from financials
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
// Patch update the navigation property customer in financials
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
func (m *CustomerRequestBuilder) PaymentMethod()(*i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e.PaymentMethodRequestBuilder) {
    return i468b9765555f0250a72696562f7c7e2a23bfe30de9cc145889ae3b989a32fd0e.NewPaymentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerRequestBuilder) PaymentTerm()(*i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1.PaymentTermRequestBuilder) {
    return i7c670bde21fab22d4289501fa0857f039be7212c1aa38eff2d9faf690afb7cd1.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CustomerRequestBuilder) Picture()(*i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.PictureRequestBuilder) {
    return i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.NewPictureRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PictureById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.financials.companies.item.salesOrders.item.customer.picture.item collection
func (m *CustomerRequestBuilder) PictureById(id string)(*i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.PictureRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["picture_id"] = id
    }
    return i8ae3f31ac08678308bf7791f9e4b50a6f00c010ce9883d3ac272bd0e2b35b37a.NewPictureRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CustomerRequestBuilder) ShipmentMethod()(*ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632.ShipmentMethodRequestBuilder) {
    return ifffabf61bbe8bc6a7cce1d58fdc3ccf29bc030985b947f5584d45e89c8fc0632.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
