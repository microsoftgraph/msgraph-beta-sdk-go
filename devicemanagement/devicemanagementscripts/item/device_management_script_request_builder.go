package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates"
    i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assignments"
    i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates"
    if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assign"
    if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/runsummary"
    ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/groupassignments"
    i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/groupassignments/item"
    i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/devicerunstates/item"
    i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/assignments/item"
    icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicemanagementscripts/item/userrunstates/item"
)

// DeviceManagementScriptRequestBuilder builds and executes requests for operations under \deviceManagement\deviceManagementScripts\{deviceManagementScript-id}
type DeviceManagementScriptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementScriptRequestBuilderDeleteOptions options for Delete
type DeviceManagementScriptRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementScriptRequestBuilderGetOptions options for Get
type DeviceManagementScriptRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementScriptRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementScriptRequestBuilderGetQueryParameters the list of device management scripts associated with the tenant.
type DeviceManagementScriptRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementScriptRequestBuilderPatchOptions options for Patch
type DeviceManagementScriptRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementScriptRequestBuilder) Assign()(*if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.AssignRequestBuilder) {
    return if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) Assignments()(*i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.AssignmentsRequestBuilder) {
    return i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.assignments.item collection
func (m *DeviceManagementScriptRequestBuilder) AssignmentsById(id string)(*i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.DeviceManagementScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.NewDeviceManagementScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementScriptRequestBuilderInternal instantiates a new DeviceManagementScriptRequestBuilder and sets the default values.
func NewDeviceManagementScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementScriptRequestBuilder) {
    m := &DeviceManagementScriptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementScriptRequestBuilder instantiates a new DeviceManagementScriptRequestBuilder and sets the default values.
func NewDeviceManagementScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementScriptRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) CreateGetRequestInformation(options *DeviceManagementScriptRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementScriptRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) Delete(options *DeviceManagementScriptRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceManagementScriptRequestBuilder) DeviceRunStates()(*i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.DeviceRunStatesRequestBuilder) {
    return i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.deviceRunStates.item collection
func (m *DeviceManagementScriptRequestBuilder) DeviceRunStatesById(id string)(*i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.DeviceManagementScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.NewDeviceManagementScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) Get(options *DeviceManagementScriptRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementScript() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementScript), nil
}
func (m *DeviceManagementScriptRequestBuilder) GroupAssignments()(*ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.GroupAssignmentsRequestBuilder) {
    return ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.groupAssignments.item collection
func (m *DeviceManagementScriptRequestBuilder) GroupAssignmentsById(id string)(*i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.DeviceManagementScriptGroupAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.NewDeviceManagementScriptGroupAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptRequestBuilder) Patch(options *DeviceManagementScriptRequestBuilderPatchOptions)(error) {
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
func (m *DeviceManagementScriptRequestBuilder) RunSummary()(*if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.RunSummaryRequestBuilder) {
    return if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementScriptRequestBuilder) UserRunStates()(*i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.UserRunStatesRequestBuilder) {
    return i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.userRunStates.item collection
func (m *DeviceManagementScriptRequestBuilder) UserRunStatesById(id string)(*icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.DeviceManagementScriptUserStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.NewDeviceManagementScriptUserStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
