package getcomanageddevicessummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetComanagedDevicesSummaryRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.getComanagedDevicesSummary()
type GetComanagedDevicesSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetComanagedDevicesSummaryRequestBuilderGetOptions options for Get
type GetComanagedDevicesSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetComanagedDevicesSummaryResponse union type wrapper for classes comanagedDevicesSummary
type GetComanagedDevicesSummaryResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type comanagedDevicesSummary
    comanagedDevicesSummary *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagedDevicesSummary;
}
// NewGetComanagedDevicesSummaryResponse instantiates a new getComanagedDevicesSummaryResponse and sets the default values.
func NewGetComanagedDevicesSummaryResponse()(*GetComanagedDevicesSummaryResponse) {
    m := &GetComanagedDevicesSummaryResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetComanagedDevicesSummaryResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetComanagedDevicesSummary gets the comanagedDevicesSummary property value. Union type representation for type comanagedDevicesSummary
func (m *GetComanagedDevicesSummaryResponse) GetComanagedDevicesSummary()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagedDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.comanagedDevicesSummary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetComanagedDevicesSummaryResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comanagedDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewComanagedDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComanagedDevicesSummary(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagedDevicesSummary))
        }
        return nil
    }
    return res
}
func (m *GetComanagedDevicesSummaryResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetComanagedDevicesSummaryResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("comanagedDevicesSummary", m.GetComanagedDevicesSummary())
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
func (m *GetComanagedDevicesSummaryResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetComanagedDevicesSummary sets the comanagedDevicesSummary property value. Union type representation for type comanagedDevicesSummary
func (m *GetComanagedDevicesSummaryResponse) SetComanagedDevicesSummary(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ComanagedDevicesSummary)() {
    if m != nil {
        m.comanagedDevicesSummary = value
    }
}
// NewGetComanagedDevicesSummaryRequestBuilderInternal instantiates a new GetComanagedDevicesSummaryRequestBuilder and sets the default values.
func NewGetComanagedDevicesSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComanagedDevicesSummaryRequestBuilder) {
    m := &GetComanagedDevicesSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getComanagedDevicesSummary()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetComanagedDevicesSummaryRequestBuilder instantiates a new GetComanagedDevicesSummaryRequestBuilder and sets the default values.
func NewGetComanagedDevicesSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComanagedDevicesSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetComanagedDevicesSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getComanagedDevicesSummary
func (m *GetComanagedDevicesSummaryRequestBuilder) CreateGetRequestInformation(options *GetComanagedDevicesSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getComanagedDevicesSummary
func (m *GetComanagedDevicesSummaryRequestBuilder) Get(options *GetComanagedDevicesSummaryRequestBuilderGetOptions)(*GetComanagedDevicesSummaryResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetComanagedDevicesSummaryResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetComanagedDevicesSummaryResponse), nil
}
