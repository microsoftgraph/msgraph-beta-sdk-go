package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses"
    i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/deploysummary"
    i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assign"
    idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments"
    ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments/item"
    if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses/item"
)

// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder builds and executes requests for operations under \deviceAppManagement\wdacSupplementalPolicies\{windowsDefenderApplicationControlSupplementalPolicy-id}
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderDeleteOptions options for Delete
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetOptions options for Get
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters the collection of Windows Defender Application Control Supplemental Policies.
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderPatchOptions options for Patch
type WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assign()(*i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.AssignRequestBuilder) {
    return i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Assignments()(*idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.AssignmentsRequestBuilder) {
    return idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.wdacSupplementalPolicies.item.assignments.item collection
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) AssignmentsById(id string)(*ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.WindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyAssignment_id"] = id
    }
    return ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDefenderApplicationControlSupplementalPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreateDeleteRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreateGetRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) CreatePatchRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Delete(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderDeleteOptions)(error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeploySummary()(*i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.DeploySummaryRequestBuilder) {
    return i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.NewDeploySummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeviceStatuses()(*i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.DeviceStatusesRequestBuilder) {
    return i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.wdacSupplementalPolicies.item.deviceStatuses.item collection
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) DeviceStatusesById(id string)(*if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus_id"] = id
    }
    return if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Get(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWindowsDefenderApplicationControlSupplementalPolicy() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicy), nil
}
// Patch the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilder) Patch(options *WindowsDefenderApplicationControlSupplementalPolicyRequestBuilderPatchOptions)(error) {
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
