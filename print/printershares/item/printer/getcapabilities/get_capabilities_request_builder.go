package getcapabilities

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetCapabilitiesRequestBuilder builds and executes requests for operations under \print\printerShares\{printerShare-id}\printer\microsoft.graph.getCapabilities()
type GetCapabilitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetCapabilitiesRequestBuilderGetOptions options for Get
type GetCapabilitiesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetCapabilitiesResponse union type wrapper for classes printerCapabilities
type GetCapabilitiesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type printerCapabilities
    printerCapabilities *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterCapabilities;
}
// NewGetCapabilitiesResponse instantiates a new getCapabilitiesResponse and sets the default values.
func NewGetCapabilitiesResponse()(*GetCapabilitiesResponse) {
    m := &GetCapabilitiesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetCapabilitiesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPrinterCapabilities gets the printerCapabilities property value. Union type representation for type printerCapabilities
func (m *GetCapabilitiesResponse) GetPrinterCapabilities()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.printerCapabilities
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetCapabilitiesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["printerCapabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrinterCapabilities() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterCapabilities(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterCapabilities))
        }
        return nil
    }
    return res
}
func (m *GetCapabilitiesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetCapabilitiesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("printerCapabilities", m.GetPrinterCapabilities())
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
func (m *GetCapabilitiesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPrinterCapabilities sets the printerCapabilities property value. Union type representation for type printerCapabilities
func (m *GetCapabilitiesResponse) SetPrinterCapabilities(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrinterCapabilities)() {
    if m != nil {
        m.printerCapabilities = value
    }
}
// NewGetCapabilitiesRequestBuilderInternal instantiates a new GetCapabilitiesRequestBuilder and sets the default values.
func NewGetCapabilitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetCapabilitiesRequestBuilder) {
    m := &GetCapabilitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/printerShares/{printerShare_id}/printer/microsoft.graph.getCapabilities()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetCapabilitiesRequestBuilder instantiates a new GetCapabilitiesRequestBuilder and sets the default values.
func NewGetCapabilitiesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetCapabilitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCapabilitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getCapabilities
func (m *GetCapabilitiesRequestBuilder) CreateGetRequestInformation(options *GetCapabilitiesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getCapabilities
func (m *GetCapabilitiesRequestBuilder) Get(options *GetCapabilitiesRequestBuilderGetOptions)(*GetCapabilitiesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetCapabilitiesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetCapabilitiesResponse), nil
}
