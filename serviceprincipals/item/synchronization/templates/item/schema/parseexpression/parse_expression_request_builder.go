package parseexpression

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ParseExpressionRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\synchronization\templates\{synchronizationTemplate-id}\schema\microsoft.graph.parseExpression
type ParseExpressionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParseExpressionRequestBuilderPostOptions options for Post
type ParseExpressionRequestBuilderPostOptions struct {
    // 
    Body *ParseExpressionRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParseExpressionResponse union type wrapper for classes parseExpressionResponse
type ParseExpressionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type parseExpressionResponse
    parseExpressionResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExpressionResponse;
}
// NewParseExpressionResponse instantiates a new parseExpressionResponse and sets the default values.
func NewParseExpressionResponse()(*ParseExpressionResponse) {
    m := &ParseExpressionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ParseExpressionResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetParseExpressionResponse gets the parseExpressionResponse property value. Union type representation for type parseExpressionResponse
func (m *ParseExpressionResponse) GetParseExpressionResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExpressionResponse) {
    if m == nil {
        return nil
    } else {
        return m.parseExpressionResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParseExpressionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["parseExpressionResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewParseExpressionResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParseExpressionResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExpressionResponse))
        }
        return nil
    }
    return res
}
func (m *ParseExpressionResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ParseExpressionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("parseExpressionResponse", m.GetParseExpressionResponse())
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
func (m *ParseExpressionResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetParseExpressionResponse sets the parseExpressionResponse property value. Union type representation for type parseExpressionResponse
func (m *ParseExpressionResponse) SetParseExpressionResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseExpressionResponse)() {
    if m != nil {
        m.parseExpressionResponse = value
    }
}
// NewParseExpressionRequestBuilderInternal instantiates a new ParseExpressionRequestBuilder and sets the default values.
func NewParseExpressionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParseExpressionRequestBuilder) {
    m := &ParseExpressionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/synchronization/templates/{synchronizationTemplate_id}/schema/microsoft.graph.parseExpression";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParseExpressionRequestBuilder instantiates a new ParseExpressionRequestBuilder and sets the default values.
func NewParseExpressionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParseExpressionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParseExpressionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action parseExpression
func (m *ParseExpressionRequestBuilder) CreatePostRequestInformation(options *ParseExpressionRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action parseExpression
func (m *ParseExpressionRequestBuilder) Post(options *ParseExpressionRequestBuilderPostOptions)(*ParseExpressionResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParseExpressionResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ParseExpressionResponse), nil
}
