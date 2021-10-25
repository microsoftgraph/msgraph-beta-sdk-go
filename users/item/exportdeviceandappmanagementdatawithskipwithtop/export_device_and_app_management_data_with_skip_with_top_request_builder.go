package exportdeviceandappmanagementdatawithskipwithtop

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ExportDeviceAndAppManagementDataWithSkipWithTopResponse struct {
    additionalData map[string]interface{};
    deviceAndAppManagementData *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData;
}
func NewExportDeviceAndAppManagementDataWithSkipWithTopResponse()(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse) {
    m := &ExportDeviceAndAppManagementDataWithSkipWithTopResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetDeviceAndAppManagementData()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementData
    }
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceAndAppManagementData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceAndAppManagementData() })
        if err != nil {
            return err
        }
        m.SetDeviceAndAppManagementData(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData))
        return nil
    }
    return res
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) IsNil()(bool) {
    return m == nil
}
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
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopResponse) SetDeviceAndAppManagementData(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementData)() {
    m.deviceAndAppManagementData = value
}
func NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, skip *int32, top *int32)(*ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    m := &ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/microsoft.graph.exportDeviceAndAppManagementData(skip={skip},top={top})";
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
func NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) CreateGetRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) Get(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExportDeviceAndAppManagementDataWithSkipWithTopResponse() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*ExportDeviceAndAppManagementDataWithSkipWithTopResponse), nil
}
