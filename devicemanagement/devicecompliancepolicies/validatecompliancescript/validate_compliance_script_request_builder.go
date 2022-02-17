package validatecompliancescript

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ValidateComplianceScriptRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCompliancePolicies\microsoft.graph.validateComplianceScript
type ValidateComplianceScriptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ValidateComplianceScriptRequestBuilderPostOptions options for Post
type ValidateComplianceScriptRequestBuilderPostOptions struct {
    // 
    Body *ValidateComplianceScriptRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ValidateComplianceScriptResponse union type wrapper for classes deviceComplianceScriptValidationResult
type ValidateComplianceScriptResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceComplianceScriptValidationResult
    deviceComplianceScriptValidationResult *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScriptValidationResult;
}
// NewValidateComplianceScriptResponse instantiates a new validateComplianceScriptResponse and sets the default values.
func NewValidateComplianceScriptResponse()(*ValidateComplianceScriptResponse) {
    m := &ValidateComplianceScriptResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateComplianceScriptResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceComplianceScriptValidationResult gets the deviceComplianceScriptValidationResult property value. Union type representation for type deviceComplianceScriptValidationResult
func (m *ValidateComplianceScriptResponse) GetDeviceComplianceScriptValidationResult()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScriptValidationResult) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptValidationResult
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateComplianceScriptResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceComplianceScriptValidationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceComplianceScriptValidationResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptValidationResult(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScriptValidationResult))
        }
        return nil
    }
    return res
}
func (m *ValidateComplianceScriptResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidateComplianceScriptResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceComplianceScriptValidationResult", m.GetDeviceComplianceScriptValidationResult())
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
func (m *ValidateComplianceScriptResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceComplianceScriptValidationResult sets the deviceComplianceScriptValidationResult property value. Union type representation for type deviceComplianceScriptValidationResult
func (m *ValidateComplianceScriptResponse) SetDeviceComplianceScriptValidationResult(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScriptValidationResult)() {
    if m != nil {
        m.deviceComplianceScriptValidationResult = value
    }
}
// NewValidateComplianceScriptRequestBuilderInternal instantiates a new ValidateComplianceScriptRequestBuilder and sets the default values.
func NewValidateComplianceScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidateComplianceScriptRequestBuilder) {
    m := &ValidateComplianceScriptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/microsoft.graph.validateComplianceScript";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewValidateComplianceScriptRequestBuilder instantiates a new ValidateComplianceScriptRequestBuilder and sets the default values.
func NewValidateComplianceScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ValidateComplianceScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewValidateComplianceScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action validateComplianceScript
func (m *ValidateComplianceScriptRequestBuilder) CreatePostRequestInformation(options *ValidateComplianceScriptRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action validateComplianceScript
func (m *ValidateComplianceScriptRequestBuilder) Post(options *ValidateComplianceScriptRequestBuilderPostOptions)(*ValidateComplianceScriptResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewValidateComplianceScriptResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ValidateComplianceScriptResponse), nil
}
