package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i425668063c4800a10861bb7a90edb930982863c2aa08069ec607be067c55a1d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices"
    i67a9ff80d5845b32459d23f5d061077fd129b768b2766975860591a869ba4db0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assign"
    ib588024c40e95bb44ff654bf2fca2e70fc4f9b254f3553765344e6d0c0ec4a32 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assignments"
    i1d904618d59e89a54c3c9f45822a76448ad3eb5947fa5abe537c63487f9edad9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assignments/item"
    ia6017df9f924cc3a6ee560c37d576949bd563ee7a9fae4b16ea59bb6d899a893 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/windowsautopilotdeploymentprofiles/item/assigneddevices/item"
)

// WindowsAutopilotDeploymentProfileRequestBuilder builds and executes requests for operations under \deviceManagement\windowsAutopilotDeploymentProfiles\{windowsAutopilotDeploymentProfile-id}
type WindowsAutopilotDeploymentProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsAutopilotDeploymentProfileRequestBuilderDeleteOptions options for Delete
type WindowsAutopilotDeploymentProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeploymentProfileRequestBuilderGetOptions options for Get
type WindowsAutopilotDeploymentProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsAutopilotDeploymentProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsAutopilotDeploymentProfileRequestBuilderGetQueryParameters windows auto pilot deployment profiles
type WindowsAutopilotDeploymentProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsAutopilotDeploymentProfileRequestBuilderPatchOptions options for Patch
type WindowsAutopilotDeploymentProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeploymentProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) Assign()(*i67a9ff80d5845b32459d23f5d061077fd129b768b2766975860591a869ba4db0.AssignRequestBuilder) {
    return i67a9ff80d5845b32459d23f5d061077fd129b768b2766975860591a869ba4db0.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) AssignedDevices()(*i425668063c4800a10861bb7a90edb930982863c2aa08069ec607be067c55a1d7.AssignedDevicesRequestBuilder) {
    return i425668063c4800a10861bb7a90edb930982863c2aa08069ec607be067c55a1d7.NewAssignedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignedDevicesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsAutopilotDeploymentProfiles.item.assignedDevices.item collection
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) AssignedDevicesById(id string)(*ia6017df9f924cc3a6ee560c37d576949bd563ee7a9fae4b16ea59bb6d899a893.WindowsAutopilotDeviceIdentityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity_id"] = id
    }
    return ia6017df9f924cc3a6ee560c37d576949bd563ee7a9fae4b16ea59bb6d899a893.NewWindowsAutopilotDeviceIdentityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) Assignments()(*ib588024c40e95bb44ff654bf2fca2e70fc4f9b254f3553765344e6d0c0ec4a32.AssignmentsRequestBuilder) {
    return ib588024c40e95bb44ff654bf2fca2e70fc4f9b254f3553765344e6d0c0ec4a32.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.windowsAutopilotDeploymentProfiles.item.assignments.item collection
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) AssignmentsById(id string)(*i1d904618d59e89a54c3c9f45822a76448ad3eb5947fa5abe537c63487f9edad9.WindowsAutopilotDeploymentProfileAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeploymentProfileAssignment_id"] = id
    }
    return i1d904618d59e89a54c3c9f45822a76448ad3eb5947fa5abe537c63487f9edad9.NewWindowsAutopilotDeploymentProfileAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsAutopilotDeploymentProfileRequestBuilderInternal instantiates a new WindowsAutopilotDeploymentProfileRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeploymentProfileRequestBuilder) {
    m := &WindowsAutopilotDeploymentProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsAutopilotDeploymentProfiles/{windowsAutopilotDeploymentProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsAutopilotDeploymentProfileRequestBuilder instantiates a new WindowsAutopilotDeploymentProfileRequestBuilder and sets the default values.
func NewWindowsAutopilotDeploymentProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsAutopilotDeploymentProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsAutopilotDeploymentProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) CreateDeleteRequestInformation(options *WindowsAutopilotDeploymentProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) CreateGetRequestInformation(options *WindowsAutopilotDeploymentProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) CreatePatchRequestInformation(options *WindowsAutopilotDeploymentProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) Delete(options *WindowsAutopilotDeploymentProfileRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) Get(options *WindowsAutopilotDeploymentProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsAutopilotDeploymentProfile, error) {
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
// Patch windows auto pilot deployment profiles
func (m *WindowsAutopilotDeploymentProfileRequestBuilder) Patch(options *WindowsAutopilotDeploymentProfileRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
