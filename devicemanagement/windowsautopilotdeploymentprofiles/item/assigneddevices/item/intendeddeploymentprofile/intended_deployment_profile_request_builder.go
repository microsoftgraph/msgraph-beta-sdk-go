package intendeddeploymentprofile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i47a36770f76accda94a606e1fd41230502b1841824ab3419425eeec60fc4bf1e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/intendeddeploymentprofile/assign"
    i7b5f3b9718c5165180676dec0c1f7d07e3b6bd686b6cb5b56aeaec2de1656116 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item/intendeddeploymentprofile/ref"
)

// IntendedDeploymentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeploymentProfiles\{windowsAutopilotDeploymentProfile-id}\assignedDevices\{windowsAutopilotDeviceIdentity-id}\intendedDeploymentProfile
type IntendedDeploymentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IntendedDeploymentProfileRequestBuilderGetOptions options for Get
type IntendedDeploymentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IntendedDeploymentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IntendedDeploymentProfileRequestBuilderGetQueryParameters deployment profile intended to be assigned to the Windows autopilot device.
type IntendedDeploymentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *IntendedDeploymentProfileRequestBuilder) Assign()(*i47a36770f76accda94a606e1fd41230502b1841824ab3419425eeec60fc4bf1e.AssignRequestBuilder) {
    return i47a36770f76accda94a606e1fd41230502b1841824ab3419425eeec60fc4bf1e.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIntendedDeploymentProfileRequestBuilderInternal instantiates a new IntendedDeploymentProfileRequestBuilder and sets the default values.
func NewIntendedDeploymentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IntendedDeploymentProfileRequestBuilder) {
    m := &IntendedDeploymentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile_id}/assignedDevices/{windowsAutopilotDeviceIdentity_id}/intendedDeploymentProfile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIntendedDeploymentProfileRequestBuilder instantiates a new IntendedDeploymentProfileRequestBuilder and sets the default values.
func NewIntendedDeploymentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IntendedDeploymentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntendedDeploymentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation deployment profile intended to be assigned to the Windows autopilot device.
func (m *IntendedDeploymentProfileRequestBuilder) CreateGetRequestInformation(options *IntendedDeploymentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// Get deployment profile intended to be assigned to the Windows autopilot device.
func (m *IntendedDeploymentProfileRequestBuilder) Get(options *IntendedDeploymentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeploymentProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsAutopilotDeploymentProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeploymentProfile), nil
}
func (m *IntendedDeploymentProfileRequestBuilder) Ref()(*i7b5f3b9718c5165180676dec0c1f7d07e3b6bd686b6cb5b56aeaec2de1656116.RefRequestBuilder) {
    return i7b5f3b9718c5165180676dec0c1f7d07e3b6bd686b6cb5b56aeaec2de1656116.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
