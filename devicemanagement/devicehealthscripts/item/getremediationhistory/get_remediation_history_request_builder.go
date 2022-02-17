package getremediationhistory

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetRemediationHistoryRequestBuilder builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\microsoft.graph.getRemediationHistory()
type GetRemediationHistoryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetRemediationHistoryRequestBuilderGetOptions options for Get
type GetRemediationHistoryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetRemediationHistoryResponse union type wrapper for classes deviceHealthScriptRemediationHistory
type GetRemediationHistoryResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceHealthScriptRemediationHistory
    deviceHealthScriptRemediationHistory *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptRemediationHistory;
}
// NewGetRemediationHistoryResponse instantiates a new getRemediationHistoryResponse and sets the default values.
func NewGetRemediationHistoryResponse()(*GetRemediationHistoryResponse) {
    m := &GetRemediationHistoryResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetRemediationHistoryResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceHealthScriptRemediationHistory gets the deviceHealthScriptRemediationHistory property value. Union type representation for type deviceHealthScriptRemediationHistory
func (m *GetRemediationHistoryResponse) GetDeviceHealthScriptRemediationHistory()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptRemediationHistory) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthScriptRemediationHistory
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetRemediationHistoryResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceHealthScriptRemediationHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceHealthScriptRemediationHistory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthScriptRemediationHistory(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptRemediationHistory))
        }
        return nil
    }
    return res
}
func (m *GetRemediationHistoryResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetRemediationHistoryResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceHealthScriptRemediationHistory", m.GetDeviceHealthScriptRemediationHistory())
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
func (m *GetRemediationHistoryResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceHealthScriptRemediationHistory sets the deviceHealthScriptRemediationHistory property value. Union type representation for type deviceHealthScriptRemediationHistory
func (m *GetRemediationHistoryResponse) SetDeviceHealthScriptRemediationHistory(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptRemediationHistory)() {
    if m != nil {
        m.deviceHealthScriptRemediationHistory = value
    }
}
// NewGetRemediationHistoryRequestBuilderInternal instantiates a new GetRemediationHistoryRequestBuilder and sets the default values.
func NewGetRemediationHistoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRemediationHistoryRequestBuilder) {
    m := &GetRemediationHistoryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}/microsoft.graph.getRemediationHistory()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetRemediationHistoryRequestBuilder instantiates a new GetRemediationHistoryRequestBuilder and sets the default values.
func NewGetRemediationHistoryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetRemediationHistoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetRemediationHistoryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation function to get the number of remediations by a device health scripts
func (m *GetRemediationHistoryRequestBuilder) CreateGetRequestInformation(options *GetRemediationHistoryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get function to get the number of remediations by a device health scripts
func (m *GetRemediationHistoryRequestBuilder) Get(options *GetRemediationHistoryRequestBuilderGetOptions)(*GetRemediationHistoryResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetRemediationHistoryResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetRemediationHistoryResponse), nil
}
