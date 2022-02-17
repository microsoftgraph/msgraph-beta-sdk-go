package generatekey

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GenerateKeyRequestBuilder builds and executes requests for operations under \trustFramework\keySets\{trustFrameworkKeySet-id}\microsoft.graph.generateKey
type GenerateKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GenerateKeyRequestBuilderPostOptions options for Post
type GenerateKeyRequestBuilderPostOptions struct {
    // 
    Body *GenerateKeyRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GenerateKeyResponse union type wrapper for classes trustFrameworkKey
type GenerateKeyResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type trustFrameworkKey
    trustFrameworkKey *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKey;
}
// NewGenerateKeyResponse instantiates a new generateKeyResponse and sets the default values.
func NewGenerateKeyResponse()(*GenerateKeyResponse) {
    m := &GenerateKeyResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateKeyResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTrustFrameworkKey gets the trustFrameworkKey property value. Union type representation for type trustFrameworkKey
func (m *GenerateKeyResponse) GetTrustFrameworkKey()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKey) {
    if m == nil {
        return nil
    } else {
        return m.trustFrameworkKey
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GenerateKeyResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["trustFrameworkKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTrustFrameworkKey() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustFrameworkKey(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKey))
        }
        return nil
    }
    return res
}
func (m *GenerateKeyResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GenerateKeyResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("trustFrameworkKey", m.GetTrustFrameworkKey())
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
func (m *GenerateKeyResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTrustFrameworkKey sets the trustFrameworkKey property value. Union type representation for type trustFrameworkKey
func (m *GenerateKeyResponse) SetTrustFrameworkKey(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKey)() {
    if m != nil {
        m.trustFrameworkKey = value
    }
}
// NewGenerateKeyRequestBuilderInternal instantiates a new GenerateKeyRequestBuilder and sets the default values.
func NewGenerateKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateKeyRequestBuilder) {
    m := &GenerateKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet_id}/microsoft.graph.generateKey";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGenerateKeyRequestBuilder instantiates a new GenerateKeyRequestBuilder and sets the default values.
func NewGenerateKeyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GenerateKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGenerateKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action generateKey
func (m *GenerateKeyRequestBuilder) CreatePostRequestInformation(options *GenerateKeyRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action generateKey
func (m *GenerateKeyRequestBuilder) Post(options *GenerateKeyRequestBuilderPostOptions)(*GenerateKeyResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGenerateKeyResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GenerateKeyResponse), nil
}
