package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    iaedea421f986c2290ff5e85cf9cd350ce7d46ab761c4a52f3368aa7a04a81155 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines/item/account"
    ief6a29ce7506934c46b05f2c7309d4828ce1140b27abd5e2f67e0ff81bd7f9b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines/item/item"
)

// SalesQuoteLineRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesQuotes\{salesQuote-id}\salesQuoteLines\{salesQuoteLine-id}
type SalesQuoteLineRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesQuoteLineRequestBuilderDeleteOptions options for Delete
type SalesQuoteLineRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesQuoteLineRequestBuilderGetOptions options for Get
type SalesQuoteLineRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesQuoteLineRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesQuoteLineRequestBuilderGetQueryParameters get salesQuoteLines from financials
type SalesQuoteLineRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesQuoteLineRequestBuilderPatchOptions options for Patch
type SalesQuoteLineRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteLine;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SalesQuoteLineRequestBuilder) Account()(*iaedea421f986c2290ff5e85cf9cd350ce7d46ab761c4a52f3368aa7a04a81155.AccountRequestBuilder) {
    return iaedea421f986c2290ff5e85cf9cd350ce7d46ab761c4a52f3368aa7a04a81155.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSalesQuoteLineRequestBuilderInternal instantiates a new SalesQuoteLineRequestBuilder and sets the default values.
func NewSalesQuoteLineRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteLineRequestBuilder) {
    m := &SalesQuoteLineRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesQuotes/{salesQuote_id}/salesQuoteLines/{salesQuoteLine_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesQuoteLineRequestBuilder instantiates a new SalesQuoteLineRequestBuilder and sets the default values.
func NewSalesQuoteLineRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteLineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesQuoteLineRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesQuoteLines for financials
func (m *SalesQuoteLineRequestBuilder) CreateDeleteRequestInformation(options *SalesQuoteLineRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get salesQuoteLines from financials
func (m *SalesQuoteLineRequestBuilder) CreateGetRequestInformation(options *SalesQuoteLineRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesQuoteLines in financials
func (m *SalesQuoteLineRequestBuilder) CreatePatchRequestInformation(options *SalesQuoteLineRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property salesQuoteLines for financials
func (m *SalesQuoteLineRequestBuilder) Delete(options *SalesQuoteLineRequestBuilderDeleteOptions)(error) {
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
// Get get salesQuoteLines from financials
func (m *SalesQuoteLineRequestBuilder) Get(options *SalesQuoteLineRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteLine, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesQuoteLine() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuoteLine), nil
}
func (m *SalesQuoteLineRequestBuilder) Item()(*ief6a29ce7506934c46b05f2c7309d4828ce1140b27abd5e2f67e0ff81bd7f9b2.ItemRequestBuilder) {
    return ief6a29ce7506934c46b05f2c7309d4828ce1140b27abd5e2f67e0ff81bd7f9b2.NewItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesQuoteLines in financials
func (m *SalesQuoteLineRequestBuilder) Patch(options *SalesQuoteLineRequestBuilderPatchOptions)(error) {
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
