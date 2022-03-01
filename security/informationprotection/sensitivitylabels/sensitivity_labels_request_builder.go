package sensitivitylabels

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/extractcontentlabel"
    i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateclassificationresults"
    i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateapplication"
    ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e "github.com/microsoftgraph/msgraph-beta-sdk-go/security/informationprotection/sensitivitylabels/evaluateremoval"
)

// SensitivityLabelsRequestBuilder builds and executes requests for operations under \security\informationProtection\sensitivityLabels
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
// SensitivityLabelsRequestBuilderGetQueryParameters get sensitivityLabels from security
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
    m.urlTemplate = "{+baseurl}/security/informationProtection/sensitivityLabels{?top,skip,search,filter,count,orderby,select,expand}";
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
// CreateGetRequestInformation get sensitivityLabels from security
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
// CreatePostRequestInformation create new navigation property to sensitivityLabels for security
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
func (m *SensitivityLabelsRequestBuilder) EvaluateApplication()(*i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a.EvaluateApplicationRequestBuilder) {
    return i957182eebe9c6b496009074d85aad428d49f0594ec0d6b6159474b3f4773707a.NewEvaluateApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateClassificationResults()(*i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d.EvaluateClassificationResultsRequestBuilder) {
    return i29fcfaae13ef83638781424dc061005bcd763ce8206749ffada5d6a890ec642d.NewEvaluateClassificationResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) EvaluateRemoval()(*ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e.EvaluateRemovalRequestBuilder) {
    return ibcbf9e301f23628c55f07a925c00eab304609586bd89ac974cf3dc82de40422e.NewEvaluateRemovalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SensitivityLabelsRequestBuilder) ExtractContentLabel()(*i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35.ExtractContentLabelRequestBuilder) {
    return i2166bc7b786fde2109d924f08f6296c65e77de9464021480f8892fbdf99b3f35.NewExtractContentLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get sensitivityLabels from security
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
// Post create new navigation property to sensitivityLabels for security
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
