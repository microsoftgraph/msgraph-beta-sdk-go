package decryptbuffer

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// DecryptBufferRequestBuilder builds and executes requests for operations under \informationProtection\microsoft.graph.decryptBuffer
type DecryptBufferRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DecryptBufferRequestBuilderPostOptions options for Post
type DecryptBufferRequestBuilderPostOptions struct {
    // 
    Body *DecryptBufferRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DecryptBufferResponse union type wrapper for classes bufferDecryptionResult
type DecryptBufferResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type bufferDecryptionResult
    bufferDecryptionResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferDecryptionResult;
}
// NewDecryptBufferResponse instantiates a new decryptBufferResponse and sets the default values.
func NewDecryptBufferResponse()(*DecryptBufferResponse) {
    m := &DecryptBufferResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DecryptBufferResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBufferDecryptionResult gets the bufferDecryptionResult property value. Union type representation for type bufferDecryptionResult
func (m *DecryptBufferResponse) GetBufferDecryptionResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferDecryptionResult) {
    if m == nil {
        return nil
    } else {
        return m.bufferDecryptionResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DecryptBufferResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bufferDecryptionResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewBufferDecryptionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBufferDecryptionResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferDecryptionResult))
        }
        return nil
    }
    return res
}
func (m *DecryptBufferResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DecryptBufferResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bufferDecryptionResult", m.GetBufferDecryptionResult())
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
func (m *DecryptBufferResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBufferDecryptionResult sets the bufferDecryptionResult property value. Union type representation for type bufferDecryptionResult
func (m *DecryptBufferResponse) SetBufferDecryptionResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.BufferDecryptionResult)() {
    if m != nil {
        m.bufferDecryptionResult = value
    }
}
// NewDecryptBufferRequestBuilderInternal instantiates a new DecryptBufferRequestBuilder and sets the default values.
func NewDecryptBufferRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DecryptBufferRequestBuilder) {
    m := &DecryptBufferRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/microsoft.graph.decryptBuffer";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDecryptBufferRequestBuilder instantiates a new DecryptBufferRequestBuilder and sets the default values.
func NewDecryptBufferRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DecryptBufferRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDecryptBufferRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action decryptBuffer
func (m *DecryptBufferRequestBuilder) CreatePostRequestInformation(options *DecryptBufferRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action decryptBuffer
func (m *DecryptBufferRequestBuilder) Post(options *DecryptBufferRequestBuilderPostOptions)(*DecryptBufferResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDecryptBufferResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DecryptBufferResponse), nil
}
