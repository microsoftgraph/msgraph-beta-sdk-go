package getassignedroledetails

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetAssignedRoleDetailsRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.getAssignedRoleDetails()
type GetAssignedRoleDetailsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetAssignedRoleDetailsRequestBuilderGetOptions options for Get
type GetAssignedRoleDetailsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GetAssignedRoleDetailsResponse union type wrapper for classes deviceAndAppManagementAssignedRoleDetails
type GetAssignedRoleDetailsResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type deviceAndAppManagementAssignedRoleDetails
    deviceAndAppManagementAssignedRoleDetails *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignedRoleDetails;
}
// NewGetAssignedRoleDetailsResponse instantiates a new getAssignedRoleDetailsResponse and sets the default values.
func NewGetAssignedRoleDetailsResponse()(*GetAssignedRoleDetailsResponse) {
    m := &GetAssignedRoleDetailsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAssignedRoleDetailsResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceAndAppManagementAssignedRoleDetails gets the deviceAndAppManagementAssignedRoleDetails property value. Union type representation for type deviceAndAppManagementAssignedRoleDetails
func (m *GetAssignedRoleDetailsResponse) GetDeviceAndAppManagementAssignedRoleDetails()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignedRoleDetails) {
    if m == nil {
        return nil
    } else {
        return m.deviceAndAppManagementAssignedRoleDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetAssignedRoleDetailsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceAndAppManagementAssignedRoleDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceAndAppManagementAssignedRoleDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAndAppManagementAssignedRoleDetails(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignedRoleDetails))
        }
        return nil
    }
    return res
}
func (m *GetAssignedRoleDetailsResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetAssignedRoleDetailsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceAndAppManagementAssignedRoleDetails", m.GetDeviceAndAppManagementAssignedRoleDetails())
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
func (m *GetAssignedRoleDetailsResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceAndAppManagementAssignedRoleDetails sets the deviceAndAppManagementAssignedRoleDetails property value. Union type representation for type deviceAndAppManagementAssignedRoleDetails
func (m *GetAssignedRoleDetailsResponse) SetDeviceAndAppManagementAssignedRoleDetails(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceAndAppManagementAssignedRoleDetails)() {
    if m != nil {
        m.deviceAndAppManagementAssignedRoleDetails = value
    }
}
// NewGetAssignedRoleDetailsRequestBuilderInternal instantiates a new GetAssignedRoleDetailsRequestBuilder and sets the default values.
func NewGetAssignedRoleDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAssignedRoleDetailsRequestBuilder) {
    m := &GetAssignedRoleDetailsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.getAssignedRoleDetails()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAssignedRoleDetailsRequestBuilder instantiates a new GetAssignedRoleDetailsRequestBuilder and sets the default values.
func NewGetAssignedRoleDetailsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetAssignedRoleDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAssignedRoleDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation retrieves the assigned role definitions and role assignments of the currently authenticated user.
func (m *GetAssignedRoleDetailsRequestBuilder) CreateGetRequestInformation(options *GetAssignedRoleDetailsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get retrieves the assigned role definitions and role assignments of the currently authenticated user.
func (m *GetAssignedRoleDetailsRequestBuilder) Get(options *GetAssignedRoleDetailsRequestBuilderGetOptions)(*GetAssignedRoleDetailsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGetAssignedRoleDetailsResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*GetAssignedRoleDetailsResponse), nil
}
