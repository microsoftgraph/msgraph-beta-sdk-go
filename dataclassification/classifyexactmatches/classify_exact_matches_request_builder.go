package classifyexactmatches

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ClassifyExactMatchesRequestBuilder builds and executes requests for operations under \dataClassification\microsoft.graph.classifyExactMatches
type ClassifyExactMatchesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ClassifyExactMatchesRequestBuilderPostOptions options for Post
type ClassifyExactMatchesRequestBuilderPostOptions struct {
    // 
    Body *ClassifyExactMatchesRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ClassifyExactMatchesResponse union type wrapper for classes exactMatchClassificationResult
type ClassifyExactMatchesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type exactMatchClassificationResult
    exactMatchClassificationResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchClassificationResult;
}
// NewClassifyExactMatchesResponse instantiates a new classifyExactMatchesResponse and sets the default values.
func NewClassifyExactMatchesResponse()(*ClassifyExactMatchesResponse) {
    m := &ClassifyExactMatchesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExactMatchClassificationResult gets the exactMatchClassificationResult property value. Union type representation for type exactMatchClassificationResult
func (m *ClassifyExactMatchesResponse) GetExactMatchClassificationResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchClassificationResult) {
    if m == nil {
        return nil
    } else {
        return m.exactMatchClassificationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyExactMatchesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["exactMatchClassificationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExactMatchClassificationResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExactMatchClassificationResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchClassificationResult))
        }
        return nil
    }
    return res
}
func (m *ClassifyExactMatchesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassifyExactMatchesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("exactMatchClassificationResult", m.GetExactMatchClassificationResult())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyExactMatchesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExactMatchClassificationResult sets the exactMatchClassificationResult property value. Union type representation for type exactMatchClassificationResult
func (m *ClassifyExactMatchesResponse) SetExactMatchClassificationResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExactMatchClassificationResult)() {
    if m != nil {
        m.exactMatchClassificationResult = value
    }
}
// NewClassifyExactMatchesRequestBuilderInternal instantiates a new ClassifyExactMatchesRequestBuilder and sets the default values.
func NewClassifyExactMatchesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassifyExactMatchesRequestBuilder) {
    m := &ClassifyExactMatchesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/microsoft.graph.classifyExactMatches";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewClassifyExactMatchesRequestBuilder instantiates a new ClassifyExactMatchesRequestBuilder and sets the default values.
func NewClassifyExactMatchesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassifyExactMatchesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassifyExactMatchesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) CreatePostRequestInformation(options *ClassifyExactMatchesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action classifyExactMatches
func (m *ClassifyExactMatchesRequestBuilder) Post(options *ClassifyExactMatchesRequestBuilderPostOptions)(*ClassifyExactMatchesResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassifyExactMatchesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ClassifyExactMatchesResponse), nil
}
