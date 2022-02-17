package deploymentprofile

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1165f0f943f657842b100195e60d9a7d27a5df0291bbcc71b0d2106a7e46b5da "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/deploymentprofile/assign"
    i492f6b1b31cd8e9a26dcf348bfaa01652c021973c5ca8d79d380b33d544b80c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item/deploymentprofile/ref"
)

// DeploymentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeviceIdentities\{windowsAutopilotDeviceIdentity-id}\deploymentProfile
type DeploymentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeploymentProfileRequestBuilderGetOptions options for Get
type DeploymentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeploymentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeploymentProfileRequestBuilderGetQueryParameters deployment profile currently assigned to the Windows autopilot device.
type DeploymentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
func (m *DeploymentProfileRequestBuilder) Assign()(*i1165f0f943f657842b100195e60d9a7d27a5df0291bbcc71b0d2106a7e46b5da.AssignRequestBuilder) {
    return i1165f0f943f657842b100195e60d9a7d27a5df0291bbcc71b0d2106a7e46b5da.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeploymentProfileRequestBuilderInternal instantiates a new DeploymentProfileRequestBuilder and sets the default values.
func NewDeploymentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentProfileRequestBuilder) {
    m := &DeploymentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeviceIdentities/{windowsAutopilotDeviceIdentity_id}/deploymentProfile{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeploymentProfileRequestBuilder instantiates a new DeploymentProfileRequestBuilder and sets the default values.
func NewDeploymentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation deployment profile currently assigned to the Windows autopilot device.
func (m *DeploymentProfileRequestBuilder) CreateGetRequestInformation(options *DeploymentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get deployment profile currently assigned to the Windows autopilot device.
func (m *DeploymentProfileRequestBuilder) Get(options *DeploymentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeploymentProfile, error) {
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
func (m *DeploymentProfileRequestBuilder) Ref()(*i492f6b1b31cd8e9a26dcf348bfaa01652c021973c5ca8d79d380b33d544b80c6.RefRequestBuilder) {
    return i492f6b1b31cd8e9a26dcf348bfaa01652c021973c5ca8d79d380b33d544b80c6.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
