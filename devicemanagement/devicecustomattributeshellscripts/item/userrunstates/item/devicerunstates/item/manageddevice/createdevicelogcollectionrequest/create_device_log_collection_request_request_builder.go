package createdevicelogcollectionrequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CreateDeviceLogCollectionRequestRequestBuilder builds and executes requests for operations under \deviceManagement\deviceCustomAttributeShellScripts\{deviceCustomAttributeShellScript-id}\userRunStates\{deviceManagementScriptUserState-id}\deviceRunStates\{deviceManagementScriptDeviceState-id}\managedDevice\microsoft.graph.createDeviceLogCollectionRequest
type CreateDeviceLogCollectionRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreateDeviceLogCollectionRequestRequestBuilderPostOptions options for Post
type CreateDeviceLogCollectionRequestRequestBuilderPostOptions struct {
    // 
    Body *CreateDeviceLogCollectionRequestRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CreateDeviceLogCollectionRequestResponse union type wrapper for classes deviceLogCollectionResponse
type CreateDeviceLogCollectionRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceLogCollectionResponse
    deviceLogCollectionResponse *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionResponse;
}
// NewCreateDeviceLogCollectionRequestResponse instantiates a new createDeviceLogCollectionRequestResponse and sets the default values.
func NewCreateDeviceLogCollectionRequestResponse()(*CreateDeviceLogCollectionRequestResponse) {
    m := &CreateDeviceLogCollectionRequestResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateDeviceLogCollectionRequestResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceLogCollectionResponse gets the deviceLogCollectionResponse property value. Union type representation for type deviceLogCollectionResponse
func (m *CreateDeviceLogCollectionRequestResponse) GetDeviceLogCollectionResponse()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionResponse) {
    if m == nil {
        return nil
    } else {
        return m.deviceLogCollectionResponse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateDeviceLogCollectionRequestResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceLogCollectionResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceLogCollectionResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLogCollectionResponse(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionResponse))
        }
        return nil
    }
    return res
}
func (m *CreateDeviceLogCollectionRequestResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateDeviceLogCollectionRequestResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceLogCollectionResponse", m.GetDeviceLogCollectionResponse())
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
func (m *CreateDeviceLogCollectionRequestResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceLogCollectionResponse sets the deviceLogCollectionResponse property value. Union type representation for type deviceLogCollectionResponse
func (m *CreateDeviceLogCollectionRequestResponse) SetDeviceLogCollectionResponse(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceLogCollectionResponse)() {
    if m != nil {
        m.deviceLogCollectionResponse = value
    }
}
// NewCreateDeviceLogCollectionRequestRequestBuilderInternal instantiates a new CreateDeviceLogCollectionRequestRequestBuilder and sets the default values.
func NewCreateDeviceLogCollectionRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateDeviceLogCollectionRequestRequestBuilder) {
    m := &CreateDeviceLogCollectionRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}/userRunStates/{deviceManagementScriptUserState_id}/deviceRunStates/{deviceManagementScriptDeviceState_id}/managedDevice/microsoft.graph.createDeviceLogCollectionRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreateDeviceLogCollectionRequestRequestBuilder instantiates a new CreateDeviceLogCollectionRequestRequestBuilder and sets the default values.
func NewCreateDeviceLogCollectionRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreateDeviceLogCollectionRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreateDeviceLogCollectionRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action createDeviceLogCollectionRequest
func (m *CreateDeviceLogCollectionRequestRequestBuilder) CreatePostRequestInformation(options *CreateDeviceLogCollectionRequestRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action createDeviceLogCollectionRequest
func (m *CreateDeviceLogCollectionRequestRequestBuilder) Post(options *CreateDeviceLogCollectionRequestRequestBuilderPostOptions)(*CreateDeviceLogCollectionRequestResponse, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCreateDeviceLogCollectionRequestResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*CreateDeviceLogCollectionRequestResponse), nil
}
