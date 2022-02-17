package encryptbuffer

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// EncryptBufferRequestBuilder builds and executes requests for operations under \informationProtection\microsoft.graph.encryptBuffer
type EncryptBufferRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EncryptBufferRequestBuilderPostOptions options for Post
type EncryptBufferRequestBuilderPostOptions struct {
    // 
    Body *EncryptBufferRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EncryptBufferResponse union type wrapper for classes bufferEncryptionResult
type EncryptBufferResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type bufferEncryptionResult
    bufferEncryptionResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferEncryptionResult;
}
// NewEncryptBufferResponse instantiates a new encryptBufferResponse and sets the default values.
func NewEncryptBufferResponse()(*EncryptBufferResponse) {
    m := &EncryptBufferResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptBufferResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBufferEncryptionResult gets the bufferEncryptionResult property value. Union type representation for type bufferEncryptionResult
func (m *EncryptBufferResponse) GetBufferEncryptionResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferEncryptionResult) {
    if m == nil {
        return nil
    } else {
        return m.bufferEncryptionResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptBufferResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bufferEncryptionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBufferEncryptionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBufferEncryptionResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferEncryptionResult))
        }
        return nil
    }
    return res
}
func (m *EncryptBufferResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EncryptBufferResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bufferEncryptionResult", m.GetBufferEncryptionResult())
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
func (m *EncryptBufferResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBufferEncryptionResult sets the bufferEncryptionResult property value. Union type representation for type bufferEncryptionResult
func (m *EncryptBufferResponse) SetBufferEncryptionResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferEncryptionResult)() {
    if m != nil {
        m.bufferEncryptionResult = value
    }
}
// NewEncryptBufferRequestBuilderInternal instantiates a new EncryptBufferRequestBuilder and sets the default values.
func NewEncryptBufferRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EncryptBufferRequestBuilder) {
    m := &EncryptBufferRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/microsoft.graph.encryptBuffer";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEncryptBufferRequestBuilder instantiates a new EncryptBufferRequestBuilder and sets the default values.
func NewEncryptBufferRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EncryptBufferRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEncryptBufferRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action encryptBuffer
func (m *EncryptBufferRequestBuilder) CreatePostRequestInformation(options *EncryptBufferRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action encryptBuffer
func (m *EncryptBufferRequestBuilder) Post(options *EncryptBufferRequestBuilderPostOptions)(*EncryptBufferResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEncryptBufferResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*EncryptBufferResponse), nil
}
