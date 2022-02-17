package evaluatedynamicmembership

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// EvaluateDynamicMembershipRequestBuilder builds and executes requests for operations under \groups\{group-id}\microsoft.graph.evaluateDynamicMembership
type EvaluateDynamicMembershipRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EvaluateDynamicMembershipRequestBuilderPostOptions options for Post
type EvaluateDynamicMembershipRequestBuilderPostOptions struct {
    // 
    Body *EvaluateDynamicMembershipRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EvaluateDynamicMembershipResponse union type wrapper for classes evaluateDynamicMembershipResult
type EvaluateDynamicMembershipResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type evaluateDynamicMembershipResult
    evaluateDynamicMembershipResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateDynamicMembershipResult;
}
// NewEvaluateDynamicMembershipResponse instantiates a new evaluateDynamicMembershipResponse and sets the default values.
func NewEvaluateDynamicMembershipResponse()(*EvaluateDynamicMembershipResponse) {
    m := &EvaluateDynamicMembershipResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateDynamicMembershipResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEvaluateDynamicMembershipResult gets the evaluateDynamicMembershipResult property value. Union type representation for type evaluateDynamicMembershipResult
func (m *EvaluateDynamicMembershipResponse) GetEvaluateDynamicMembershipResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateDynamicMembershipResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluateDynamicMembershipResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateDynamicMembershipResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["evaluateDynamicMembershipResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvaluateDynamicMembershipResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluateDynamicMembershipResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateDynamicMembershipResult))
        }
        return nil
    }
    return res
}
func (m *EvaluateDynamicMembershipResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateDynamicMembershipResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("evaluateDynamicMembershipResult", m.GetEvaluateDynamicMembershipResult())
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
func (m *EvaluateDynamicMembershipResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEvaluateDynamicMembershipResult sets the evaluateDynamicMembershipResult property value. Union type representation for type evaluateDynamicMembershipResult
func (m *EvaluateDynamicMembershipResponse) SetEvaluateDynamicMembershipResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EvaluateDynamicMembershipResult)() {
    if m != nil {
        m.evaluateDynamicMembershipResult = value
    }
}
// NewEvaluateDynamicMembershipRequestBuilderInternal instantiates a new EvaluateDynamicMembershipRequestBuilder and sets the default values.
func NewEvaluateDynamicMembershipRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EvaluateDynamicMembershipRequestBuilder) {
    m := &EvaluateDynamicMembershipRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/microsoft.graph.evaluateDynamicMembership";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEvaluateDynamicMembershipRequestBuilder instantiates a new EvaluateDynamicMembershipRequestBuilder and sets the default values.
func NewEvaluateDynamicMembershipRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EvaluateDynamicMembershipRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEvaluateDynamicMembershipRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action evaluateDynamicMembership
func (m *EvaluateDynamicMembershipRequestBuilder) CreatePostRequestInformation(options *EvaluateDynamicMembershipRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action evaluateDynamicMembership
func (m *EvaluateDynamicMembershipRequestBuilder) Post(options *EvaluateDynamicMembershipRequestBuilderPostOptions)(*EvaluateDynamicMembershipResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateDynamicMembershipResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*EvaluateDynamicMembershipResponse), nil
}
