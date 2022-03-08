package areglobalscriptsavailable

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// AreGlobalScriptsAvailableRequestBuilder provides operations to call the areGlobalScriptsAvailable method.
type AreGlobalScriptsAvailableRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AreGlobalScriptsAvailableRequestBuilderGetOptions options for Get
type AreGlobalScriptsAvailableRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AreGlobalScriptsAvailableResponse union type wrapper for classes globalDeviceHealthScriptState
type AreGlobalScriptsAvailableResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type globalDeviceHealthScriptState
    globalDeviceHealthScriptState *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GlobalDeviceHealthScriptState;
}
// NewAreGlobalScriptsAvailableResponse instantiates a new areGlobalScriptsAvailableResponse and sets the default values.
func NewAreGlobalScriptsAvailableResponse()(*AreGlobalScriptsAvailableResponse) {
    m := &AreGlobalScriptsAvailableResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateAreGlobalScriptsAvailableResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAreGlobalScriptsAvailableResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AreGlobalScriptsAvailableResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AreGlobalScriptsAvailableResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["globalDeviceHealthScriptState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseGlobalDeviceHealthScriptState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalDeviceHealthScriptState(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GlobalDeviceHealthScriptState))
        }
        return nil
    }
    return res
}
// GetGlobalDeviceHealthScriptState gets the globalDeviceHealthScriptState property value. Union type representation for type globalDeviceHealthScriptState
func (m *AreGlobalScriptsAvailableResponse) GetGlobalDeviceHealthScriptState()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GlobalDeviceHealthScriptState) {
    if m == nil {
        return nil
    } else {
        return m.globalDeviceHealthScriptState
    }
}
func (m *AreGlobalScriptsAvailableResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AreGlobalScriptsAvailableResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetGlobalDeviceHealthScriptState() != nil {
        cast := (*m.GetGlobalDeviceHealthScriptState()).String()
        err := writer.WriteStringValue("globalDeviceHealthScriptState", &cast)
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
func (m *AreGlobalScriptsAvailableResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGlobalDeviceHealthScriptState sets the globalDeviceHealthScriptState property value. Union type representation for type globalDeviceHealthScriptState
func (m *AreGlobalScriptsAvailableResponse) SetGlobalDeviceHealthScriptState(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GlobalDeviceHealthScriptState)() {
    if m != nil {
        m.globalDeviceHealthScriptState = value
    }
}
// NewAreGlobalScriptsAvailableRequestBuilderInternal instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    m := &AreGlobalScriptsAvailableRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/microsoft.graph.areGlobalScriptsAvailable()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAreGlobalScriptsAvailableRequestBuilder instantiates a new AreGlobalScriptsAvailableRequestBuilder and sets the default values.
func NewAreGlobalScriptsAvailableRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AreGlobalScriptsAvailableRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAreGlobalScriptsAvailableRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) CreateGetRequestInformation(options *AreGlobalScriptsAvailableRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function areGlobalScriptsAvailable
func (m *AreGlobalScriptsAvailableRequestBuilder) Get(options *AreGlobalScriptsAvailableRequestBuilderGetOptions)(AreGlobalScriptsAvailableResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAreGlobalScriptsAvailableResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(AreGlobalScriptsAvailableResponseable), nil
}
