package userexperienceanalyticssummarizeworkfromanywheredevices

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder builds and executes requests for operations under \deviceManagement\microsoft.graph.userExperienceAnalyticsSummarizeWorkFromAnywhereDevices()
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse union type wrapper for classes userExperienceAnalyticsWorkFromAnywhereDevicesSummary
type UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type userExperienceAnalyticsWorkFromAnywhereDevicesSummary
    userExperienceAnalyticsWorkFromAnywhereDevicesSummary *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummary;
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse instantiates a new userExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse()(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) {
    m := &UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary gets the userExperienceAnalyticsWorkFromAnywhereDevicesSummary property value. Union type representation for type userExperienceAnalyticsWorkFromAnywhereDevicesSummary
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) GetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsWorkFromAnywhereDevicesSummary
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["userExperienceAnalyticsWorkFromAnywhereDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummary))
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("userExperienceAnalyticsWorkFromAnywhereDevicesSummary", m.GetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary())
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
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary sets the userExperienceAnalyticsWorkFromAnywhereDevicesSummary property value. Union type representation for type userExperienceAnalyticsWorkFromAnywhereDevicesSummary
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse) SetUserExperienceAnalyticsWorkFromAnywhereDevicesSummary(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDevicesSummary)() {
    if m != nil {
        m.userExperienceAnalyticsWorkFromAnywhereDevicesSummary = value
    }
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    m := &UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoft.graph.userExperienceAnalyticsSummarizeWorkFromAnywhereDevices()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder instantiates a new UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function userExperienceAnalyticsSummarizeWorkFromAnywhereDevices
func (m *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) Get(options *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderGetOptions)(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesResponse), nil
}
