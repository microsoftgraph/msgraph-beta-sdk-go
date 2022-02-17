package evaluate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// EvaluateRequestBuilder builds and executes requests for operations under \informationProtection\sensitivityLabels\microsoft.graph.evaluate
type EvaluateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EvaluateRequestBuilderPostOptions options for Post
type EvaluateRequestBuilderPostOptions struct {
    // 
    Body *EvaluateRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EvaluateResponse union type wrapper for classes evaluateLabelJobResponse
type EvaluateResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type evaluateLabelJobResponse
    evaluateLabelJobResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateLabelJobResponse;
}
// NewEvaluateResponse instantiates a new evaluateResponse and sets the default values.
func NewEvaluateResponse()(*EvaluateResponse) {
    m := &EvaluateResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEvaluateLabelJobResponse gets the evaluateLabelJobResponse property value. Union type representation for type evaluateLabelJobResponse
func (m *EvaluateResponse) GetEvaluateLabelJobResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateLabelJobResponse) {
    if m == nil {
        return nil
    } else {
        return m.evaluateLabelJobResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["evaluateLabelJobResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvaluateLabelJobResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluateLabelJobResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateLabelJobResponse))
        }
        return nil
    }
    return res
}
func (m *EvaluateResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("evaluateLabelJobResponse", m.GetEvaluateLabelJobResponse())
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
func (m *EvaluateResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEvaluateLabelJobResponse sets the evaluateLabelJobResponse property value. Union type representation for type evaluateLabelJobResponse
func (m *EvaluateResponse) SetEvaluateLabelJobResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateLabelJobResponse)() {
    if m != nil {
        m.evaluateLabelJobResponse = value
    }
}
// NewEvaluateRequestBuilderInternal instantiates a new EvaluateRequestBuilder and sets the default values.
func NewEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EvaluateRequestBuilder) {
    m := &EvaluateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/sensitivityLabels/microsoft.graph.evaluate";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEvaluateRequestBuilder instantiates a new EvaluateRequestBuilder and sets the default values.
func NewEvaluateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action evaluate
func (m *EvaluateRequestBuilder) CreatePostRequestInformation(options *EvaluateRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action evaluate
func (m *EvaluateRequestBuilder) Post(options *EvaluateRequestBuilderPostOptions)(*EvaluateResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*EvaluateResponse), nil
}
