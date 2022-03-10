package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses"
    i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/deploysummary"
    i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assign"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments"
    ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/assignments/item"
    if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/wdacsupplementalpolicies/item/devicestatuses/item"
)

// WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder provides operations to manage the wdacSupplementalPolicies property of the microsoft.graph.deviceAppManagement entity.
type WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteOptions options for Delete
type WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetOptions options for Get
type WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters the collection of Windows Defender Application Control Supplemental Policies.
type WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchOptions options for Patch
type WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assign()(*i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.AssignRequestBuilder) {
    return i406c14183761992d0042e2f9ff1bea41245b27611fe8826e75799a28823ca007.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Assignments()(*idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.AssignmentsRequestBuilder) {
    return idc699f5102959b816736faa2d215afc9e66f562ab6af43c03f0699e6ef2fa7f4.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.wdacSupplementalPolicies.item.assignments.item collection
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) AssignmentsById(id string)(*ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.WindowsDefenderApplicationControlSupplementalPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyAssignment_id"] = id
    }
    return ibf0406f487f23a25d59623f151d2327c977ef4cb5bf52e23fcd9992cbf752eb2.NewWindowsDefenderApplicationControlSupplementalPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder instantiates a new WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreateDeleteRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreateGetRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) CreatePatchRequestInformation(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property wdacSupplementalPolicies for deviceAppManagement
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Delete(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeploySummary()(*i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.DeploySummaryRequestBuilder) {
    return i2075ede6b5773dacc4381d01916b488d778b97f02f27c309ab4b1e54d7a9fc93.NewDeploySummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatuses()(*i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.DeviceStatusesRequestBuilder) {
    return i13775af57c3f03fa03a53f3283beac6df287db3f25fa4e4ee1a441cd3996dd0b.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.wdacSupplementalPolicies.item.deviceStatuses.item collection
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) DeviceStatusesById(id string)(*if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus_id"] = id
    }
    return if7f30d88a5e2c08a355aab945ead20a738d45d9716825973a8471d163909ac9b.NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the collection of Windows Defender Application Control Supplemental Policies.
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Get(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WindowsDefenderApplicationControlSupplementalPolicyable), nil
}
// Patch update the navigation property wdacSupplementalPolicies in deviceAppManagement
func (m *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilder) Patch(options *WindowsDefenderApplicationControlSupplementalPolicyItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
