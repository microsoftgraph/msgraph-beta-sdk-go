package sensitivitylabels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateremoval"
    i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateapplication"
    icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/security/informationprotection/sensitivitylabels/extractcontentlabel"
)

// SensitivityLabelsRequestBuilder builds and executes requests for operations under \me\security\informationProtection\sensitivityLabels
type SensitivityLabelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SensitivityLabelsRequestBuilderGetOptions options for Get
type SensitivityLabelsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SensitivityLabelsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SensitivityLabelsRequestBuilderGetQueryParameters get sensitivityLabels from me
type SensitivityLabelsRequestBuilderGetQueryParameters struct {
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
// SensitivityLabelsRequestBuilderPostOptions options for Post
type SensitivityLabelsRequestBuilderPostOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitivityLabel;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSensitivityLabelsRequestBuilderInternal instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    m := &SensitivityLabelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/security/informationProtection/sensitivityLabels{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSensitivityLabelsRequestBuilder instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get sensitivityLabels from me
func (m *SensitivityLabelsRequestBuilder) CreateGetRequestInformation(options *SensitivityLabelsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to sensitivityLabels for me
func (m *SensitivityLabelsRequestBuilder) CreatePostRequestInformation(options *SensitivityLabelsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25.EvaluateApplicationRequestBuilder) {
    return i88759b56304a762698295808a5bfb06f29418867560c74fec3dc6fd888ff3f25.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c.EvaluateClassificationResultsRequestBuilder) {
    return icc10ec374205c05b4532d9c89d0c36886f133046713bc767e65df773ff64ed0c.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b.EvaluateRemovalRequestBuilder) {
    return i5cad7b98159f40ed5a5c934595269dacf1b687ffeb3281ebb743bd111e0a847b.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0.ExtractContentLabelRequestBuilder) {
    return ie0b8e402d3912c7159a7b1635da1d709d914855376e51d3fa858ac845bf853a0.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get sensitivityLabels from me
func (m *SensitivityLabelsRequestBuilder) Get(options *SensitivityLabelsRequestBuilderGetOptions)(*SensitivityLabelsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSensitivityLabelsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*SensitivityLabelsResponse), nil
}
// Post create new navigation property to sensitivityLabels for me
func (m *SensitivityLabelsRequestBuilder) Post(options *SensitivityLabelsRequestBuilderPostOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitivityLabel, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSensitivityLabel() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SensitivityLabel), nil
}
