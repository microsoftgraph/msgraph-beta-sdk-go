package createinstance

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateInstanceRequestBuilder builds and executes requests for operations under \deviceManagement\templates\{deviceManagementTemplate-id}\microsoft.graph.createInstance
type CreateInstanceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateInstanceRequestBuilderPostOptions options for Post
type CreateInstanceRequestBuilderPostOptions struct {
    // 
    Body *CreateInstanceRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreateInstanceResponse union type wrapper for classes deviceManagementIntent
type CreateInstanceResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceManagementIntent
    deviceManagementIntent *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent;
}
// NewCreateInstanceResponse instantiates a new createInstanceResponse and sets the default values.
func NewCreateInstanceResponse()(*CreateInstanceResponse) {
    m := &CreateInstanceResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateInstanceResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceManagementIntent gets the deviceManagementIntent property value. Union type representation for type deviceManagementIntent
func (m *CreateInstanceResponse) GetDeviceManagementIntent()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementIntent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateInstanceResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceManagementIntent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementIntent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementIntent(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent))
        }
        return nil
    }
    return res
}
func (m *CreateInstanceResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateInstanceResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceManagementIntent", m.GetDeviceManagementIntent())
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
func (m *CreateInstanceResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceManagementIntent sets the deviceManagementIntent property value. Union type representation for type deviceManagementIntent
func (m *CreateInstanceResponse) SetDeviceManagementIntent(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementIntent)() {
    if m != nil {
        m.deviceManagementIntent = value
    }
}
// NewCreateInstanceRequestBuilderInternal instantiates a new CreateInstanceRequestBuilder and sets the default values.
func NewCreateInstanceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateInstanceRequestBuilder) {
    m := &CreateInstanceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate_id}/microsoft.graph.createInstance";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateInstanceRequestBuilder instantiates a new CreateInstanceRequestBuilder and sets the default values.
func NewCreateInstanceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateInstanceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateInstanceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createInstance
func (m *CreateInstanceRequestBuilder) CreatePostRequestInformation(options *CreateInstanceRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action createInstance
func (m *CreateInstanceRequestBuilder) Post(options *CreateInstanceRequestBuilderPostOptions)(*CreateInstanceResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateInstanceResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreateInstanceResponse), nil
}
