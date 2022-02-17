package exportdeviceandappmanagementdatawithskipwithtop

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder builds and executes requests for operations under \users\{user-id}\microsoft.graph.exportDeviceAndAppManagementData(skip={skip},top={top})
type ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetOptions options for Get
type ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExportDeviceAndAppManagementDataWithSkipWithTopResponse union type wrapper for classes deviceAndAppManagementData
type ExportDeviceAndAppManagementDataWithSkipWithTopResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceAndAppManagementData
    deviceAndAppManagementData *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData;
}
// NewExportDeviceAndAppManagementDataWithSkipWithTopResponse instantiates a new exportDeviceAndAppManagementDataWithSkipWithTopResponse and sets the default values.
func NewExportDeviceAndAppManagementDataWithSkipWithTopResponse()(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse) {
    m := &ExportDeviceAndAppManagementDataWithSkipWithTopResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementData gets the deviceAndAppManagementData property value. Union type representation for type deviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetDeviceAndAppManagementData()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementData sets the deviceAndAppManagementData property value. Union type representation for type deviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) SetDeviceAndAppManagementData(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData)() {
    if m != nil {
        m.deviceAndAppManagementData = value
    }
}
// NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal instantiates a new ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, skip *int32, top *int32)(*ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    m := &ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/microsoft.graph.exportDeviceAndAppManagementData(skip={skip},top={top})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if skip != nil {
        urlTplParams["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if top != nil {
        urlTplParams["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder instantiates a new ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function exportDeviceAndAppManagementData
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) CreateGetRequestInformation(options *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) Get(options *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetOptions)(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExportDeviceAndAppManagementDataWithSkipWithTopResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse), nil
}
