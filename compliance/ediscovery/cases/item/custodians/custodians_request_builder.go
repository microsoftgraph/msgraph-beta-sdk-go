package custodians

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/removehold"
    ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/applyhold"
)

// CustodiansRequestBuilder builds and executes requests for operations under \compliance\ediscovery\cases\{case-id}\custodians
type CustodiansRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CustodiansRequestBuilderGetOptions options for Get
type CustodiansRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CustodiansRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CustodiansRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type CustodiansRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// CustodiansRequestBuilderPostOptions options for Post
type CustodiansRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CustodiansRequestBuilder) ApplyHold()(*ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6.ApplyHoldRequestBuilder) {
    return ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCustodiansRequestBuilderInternal instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodiansRequestBuilder) {
    m := &CustodiansRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case_id}/custodians{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustodiansRequestBuilder instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) CreateGetRequestInformation(options *CustodiansRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) CreatePostRequestInformation(options *CustodiansRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Get returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) Get(options *CustodiansRequestBuilderGetOptions)(*CustodiansResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustodiansResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CustodiansResponse), nil
}
// Post returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) Post(options *CustodiansRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCustodian() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Custodian), nil
}
func (m *CustodiansRequestBuilder) RemoveHold()(*icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba.RemoveHoldRequestBuilder) {
    return icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
