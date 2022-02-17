package getcomanagementeligibledevicessummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetComanagementEligibleDevicesSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.getComanagementEligibleDevicesSummary()
type GetComanagementEligibleDevicesSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetComanagementEligibleDevicesSummaryRequestBuilderGetOptions options for Get
type GetComanagementEligibleDevicesSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetComanagementEligibleDevicesSummaryResponse union type wrapper for classes comanagementEligibleDevicesSummary
type GetComanagementEligibleDevicesSummaryResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type comanagementEligibleDevicesSummary
    comanagementEligibleDevicesSummary *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagementEligibleDevicesSummary;
}
// NewGetComanagementEligibleDevicesSummaryResponse instantiates a new getComanagementEligibleDevicesSummaryResponse and sets the default values.
func NewGetComanagementEligibleDevicesSummaryResponse()(*GetComanagementEligibleDevicesSummaryResponse) {
    m := &GetComanagementEligibleDevicesSummaryResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetComanagementEligibleDevicesSummaryResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComanagementEligibleDevicesSummary gets the comanagementEligibleDevicesSummary property value. Union type representation for type comanagementEligibleDevicesSummary
func (m *GetComanagementEligibleDevicesSummaryResponse) GetComanagementEligibleDevicesSummary()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagementEligibleDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.comanagementEligibleDevicesSummary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetComanagementEligibleDevicesSummaryResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comanagementEligibleDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewComanagementEligibleDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComanagementEligibleDevicesSummary(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagementEligibleDevicesSummary))
        }
        return nil
    }
    return res
}
func (m *GetComanagementEligibleDevicesSummaryResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetComanagementEligibleDevicesSummaryResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comanagementEligibleDevicesSummary", m.GetComanagementEligibleDevicesSummary())
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
func (m *GetComanagementEligibleDevicesSummaryResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComanagementEligibleDevicesSummary sets the comanagementEligibleDevicesSummary property value. Union type representation for type comanagementEligibleDevicesSummary
func (m *GetComanagementEligibleDevicesSummaryResponse) SetComanagementEligibleDevicesSummary(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagementEligibleDevicesSummary)() {
    if m != nil {
        m.comanagementEligibleDevicesSummary = value
    }
}
// NewGetComanagementEligibleDevicesSummaryRequestBuilderInternal instantiates a new GetComanagementEligibleDevicesSummaryRequestBuilder and sets the default values.
func NewGetComanagementEligibleDevicesSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComanagementEligibleDevicesSummaryRequestBuilder) {
    m := &GetComanagementEligibleDevicesSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getComanagementEligibleDevicesSummary()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetComanagementEligibleDevicesSummaryRequestBuilder instantiates a new GetComanagementEligibleDevicesSummaryRequestBuilder and sets the default values.
func NewGetComanagementEligibleDevicesSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComanagementEligibleDevicesSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetComanagementEligibleDevicesSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getComanagementEligibleDevicesSummary
func (m *GetComanagementEligibleDevicesSummaryRequestBuilder) CreateGetRequestInformation(options *GetComanagementEligibleDevicesSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getComanagementEligibleDevicesSummary
func (m *GetComanagementEligibleDevicesSummaryRequestBuilder) Get(options *GetComanagementEligibleDevicesSummaryRequestBuilderGetOptions)(*GetComanagementEligibleDevicesSummaryResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetComanagementEligibleDevicesSummaryResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetComanagementEligibleDevicesSummaryResponse), nil
}
