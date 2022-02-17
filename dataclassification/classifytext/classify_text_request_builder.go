package classifytext

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ClassifyTextRequestBuilder builds and executes requests for operations under \dataClassification\microsoft.graph.classifyText
type ClassifyTextRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ClassifyTextRequestBuilderPostOptions options for Post
type ClassifyTextRequestBuilderPostOptions struct {
    // 
    Body *ClassifyTextRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ClassifyTextResponse union type wrapper for classes classificationJobResponse
type ClassifyTextResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type classificationJobResponse
    classificationJobResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponse;
}
// NewClassifyTextResponse instantiates a new classifyTextResponse and sets the default values.
func NewClassifyTextResponse()(*ClassifyTextResponse) {
    m := &ClassifyTextResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassifyTextResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClassificationJobResponse gets the classificationJobResponse property value. Union type representation for type classificationJobResponse
func (m *ClassifyTextResponse) GetClassificationJobResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponse) {
    if m == nil {
        return nil
    } else {
        return m.classificationJobResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassifyTextResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["classificationJobResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewClassificationJobResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassificationJobResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponse))
        }
        return nil
    }
    return res
}
func (m *ClassifyTextResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ClassifyTextResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("classificationJobResponse", m.GetClassificationJobResponse())
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
func (m *ClassifyTextResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClassificationJobResponse sets the classificationJobResponse property value. Union type representation for type classificationJobResponse
func (m *ClassifyTextResponse) SetClassificationJobResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponse)() {
    if m != nil {
        m.classificationJobResponse = value
    }
}
// NewClassifyTextRequestBuilderInternal instantiates a new ClassifyTextRequestBuilder and sets the default values.
func NewClassifyTextRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassifyTextRequestBuilder) {
    m := &ClassifyTextRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/dataClassification/microsoft.graph.classifyText";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewClassifyTextRequestBuilder instantiates a new ClassifyTextRequestBuilder and sets the default values.
func NewClassifyTextRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ClassifyTextRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassifyTextRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action classifyText
func (m *ClassifyTextRequestBuilder) CreatePostRequestInformation(options *ClassifyTextRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action classifyText
func (m *ClassifyTextRequestBuilder) Post(options *ClassifyTextRequestBuilderPostOptions)(*ClassifyTextResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewClassifyTextResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ClassifyTextResponse), nil
}
