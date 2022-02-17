package exportdeviceandappmanagementdata

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExportDeviceAndAppManagementDataRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.exportDeviceAndAppManagementData()
type ExportDeviceAndAppManagementDataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExportDeviceAndAppManagementDataRequestBuilderGetOptions options for Get
type ExportDeviceAndAppManagementDataRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExportDeviceAndAppManagementDataResponse union type wrapper for classes deviceAndAppManagementData
type ExportDeviceAndAppManagementDataResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceAndAppManagementData
    deviceAndAppManagementData *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData;
}
// NewExportDeviceAndAppManagementDataResponse instantiates a new exportDeviceAndAppManagementDataResponse and sets the default values.
func NewExportDeviceAndAppManagementDataResponse()(*ExportDeviceAndAppManagementDataResponse) {
    m := &ExportDeviceAndAppManagementDataResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExportDeviceAndAppManagementDataResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementData gets the deviceAndAppManagementData property value. Union type representation for type deviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataResponse) GetDeviceAndAppManagementData()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExportDeviceAndAppManagementDataResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceAndAppManagementData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceAndAppManagementData() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementData(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData))
        }
        return nil
    }
    return res
}
func (m *ExportDeviceAndAppManagementDataResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExportDeviceAndAppManagementDataResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementData", m.GetDeviceAndAppManagementData())
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
func (m *ExportDeviceAndAppManagementDataResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementData sets the deviceAndAppManagementData property value. Union type representation for type deviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataResponse) SetDeviceAndAppManagementData(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData)() {
    if m != nil {
        m.deviceAndAppManagementData = value
    }
}
// NewExportDeviceAndAppManagementDataRequestBuilderInternal instantiates a new ExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewExportDeviceAndAppManagementDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExportDeviceAndAppManagementDataRequestBuilder) {
    m := &ExportDeviceAndAppManagementDataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.exportDeviceAndAppManagementData()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExportDeviceAndAppManagementDataRequestBuilder instantiates a new ExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewExportDeviceAndAppManagementDataRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExportDeviceAndAppManagementDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportDeviceAndAppManagementDataRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function exportDeviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataRequestBuilder) CreateGetRequestInformation(options *ExportDeviceAndAppManagementDataRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function exportDeviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataRequestBuilder) Get(options *ExportDeviceAndAppManagementDataRequestBuilderGetOptions)(*ExportDeviceAndAppManagementDataResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExportDeviceAndAppManagementDataResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExportDeviceAndAppManagementDataResponse), nil
}
