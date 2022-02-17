package labels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateclassificationresults"
    i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6 "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateapplication"
    ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/evaluateremoval"
    ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d "github.com/microsoftgraph/msgraph-beta-sdk-go/informationprotection/policy/labels/extractlabel"
)

// LabelsRequestBuilder builds and executes requests for operations under \informationProtection\policy\labels
type LabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// LabelsRequestBuilderGetOptions options for Get
type LabelsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *LabelsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// LabelsRequestBuilderGetQueryParameters get labels from informationProtection
type LabelsRequestBuilderGetQueryParameters struct {
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
// LabelsRequestBuilderPostOptions options for Post
type LabelsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabel;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewLabelsRequestBuilderInternal instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LabelsRequestBuilder) {
    m := &LabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/policy/labels{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLabelsRequestBuilder instantiates a new LabelsRequestBuilder and sets the default values.
func NewLabelsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get labels from informationProtection
func (m *LabelsRequestBuilder) CreateGetRequestInformation(options *LabelsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to labels for informationProtection
func (m *LabelsRequestBuilder) CreatePostRequestInformation(options *LabelsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *LabelsRequestBuilder) EvaluateApplication()(*i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6.EvaluateApplicationRequestBuilder) {
    return i1f4566448b9b69a4c5aa53f28eeb9d10043b88f2acc3d1b58f45685b466b0bc6.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateClassificationResults()(*i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b.EvaluateClassificationResultsRequestBuilder) {
    return i13bad00e070fa1dd33a2d812272260fd137db1c553faa0f03c4be0b28414ed5b.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) EvaluateRemoval()(*ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f.EvaluateRemovalRequestBuilder) {
    return ia970da0fc5eba52aeb4709c9bf149f352a3b6c08a985fcc0e3a4bbeb86d1b14f.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *LabelsRequestBuilder) ExtractLabel()(*ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d.ExtractLabelRequestBuilder) {
    return ibfb726105da3710efbd5bdd0afd208c0efb3e76fca696a1e9b866eee417fbb3d.NewExtractLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get labels from informationProtection
func (m *LabelsRequestBuilder) Get(options *LabelsRequestBuilderGetOptions)(*LabelsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLabelsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*LabelsResponse), nil
}
// Post create new navigation property to labels for informationProtection
func (m *LabelsRequestBuilder) Post(options *LabelsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabel, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewInformationProtectionLabel() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.InformationProtectionLabel), nil
}
