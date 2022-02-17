package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item/item"
    i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salescreditmemos/item/salescreditmemolines/item/account"
)

// SalesCreditMemoLineRequestBuilder builds and executes requests for operations under \financials\companies\{company-id}\salesCreditMemos\{salesCreditMemo-id}\salesCreditMemoLines\{salesCreditMemoLine-id}
type SalesCreditMemoLineRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SalesCreditMemoLineRequestBuilderDeleteOptions options for Delete
type SalesCreditMemoLineRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesCreditMemoLineRequestBuilderGetOptions options for Get
type SalesCreditMemoLineRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SalesCreditMemoLineRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SalesCreditMemoLineRequestBuilderGetQueryParameters get salesCreditMemoLines from financials
type SalesCreditMemoLineRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SalesCreditMemoLineRequestBuilderPatchOptions options for Patch
type SalesCreditMemoLineRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemoLine;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *SalesCreditMemoLineRequestBuilder) Account()(*i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65.AccountRequestBuilder) {
    return i54ca4d303dc3c2edef6bb7873d8427ed6509db394ba39c0fa8453891e9657a65.NewAccountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSalesCreditMemoLineRequestBuilderInternal instantiates a new SalesCreditMemoLineRequestBuilder and sets the default values.
func NewSalesCreditMemoLineRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesCreditMemoLineRequestBuilder) {
    m := &SalesCreditMemoLineRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/financials/companies/{company_id}/salesCreditMemos/{salesCreditMemo_id}/salesCreditMemoLines/{salesCreditMemoLine_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSalesCreditMemoLineRequestBuilder instantiates a new SalesCreditMemoLineRequestBuilder and sets the default values.
func NewSalesCreditMemoLineRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesCreditMemoLineRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesCreditMemoLineRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property salesCreditMemoLines for financials
func (m *SalesCreditMemoLineRequestBuilder) CreateDeleteRequestInformation(options *SalesCreditMemoLineRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineRequestBuilder) CreateGetRequestInformation(options *SalesCreditMemoLineRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineRequestBuilder) CreatePatchRequestInformation(options *SalesCreditMemoLineRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property salesCreditMemoLines for financials
func (m *SalesCreditMemoLineRequestBuilder) Delete(options *SalesCreditMemoLineRequestBuilderDeleteOptions)(error) {
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
// Get get salesCreditMemoLines from financials
func (m *SalesCreditMemoLineRequestBuilder) Get(options *SalesCreditMemoLineRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemoLine, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesCreditMemoLine() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesCreditMemoLine), nil
}
func (m *SalesCreditMemoLineRequestBuilder) Item()(*i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66.ItemRequestBuilder) {
    return i17418f95bbf4a60c43605413d6584a144f2f400b8d8de167f1729a805c902c66.NewItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property salesCreditMemoLines in financials
func (m *SalesCreditMemoLineRequestBuilder) Patch(options *SalesCreditMemoLineRequestBuilderPatchOptions)(error) {
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
