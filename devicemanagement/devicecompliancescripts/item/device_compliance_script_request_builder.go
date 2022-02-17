package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/runsummary"
    i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assign"
    i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assignments"
    if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates"
    i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item"
    ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assignments/item"
)

// DeviceComplianceScriptRequestBuilder builds and executes requests for operations under \deviceManagement\deviceComplianceScripts\{deviceComplianceScript-id}
type DeviceComplianceScriptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceComplianceScriptRequestBuilderDeleteOptions options for Delete
type DeviceComplianceScriptRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceComplianceScriptRequestBuilderGetOptions options for Get
type DeviceComplianceScriptRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceComplianceScriptRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceComplianceScriptRequestBuilderGetQueryParameters the list of device compliance scripts associated with the tenant.
type DeviceComplianceScriptRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceComplianceScriptRequestBuilderPatchOptions options for Patch
type DeviceComplianceScriptRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScript;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceComplianceScriptRequestBuilder) Assign()(*i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee.AssignRequestBuilder) {
    return i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceComplianceScriptRequestBuilder) Assignments()(*i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa.AssignmentsRequestBuilder) {
    return i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceComplianceScripts.item.assignments.item collection
func (m *DeviceComplianceScriptRequestBuilder) AssignmentsById(id string)(*ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4.DeviceHealthScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment_id"] = id
    }
    return ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4.NewDeviceHealthScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceComplianceScriptRequestBuilderInternal instantiates a new DeviceComplianceScriptRequestBuilder and sets the default values.
func NewDeviceComplianceScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceComplianceScriptRequestBuilder) {
    m := &DeviceComplianceScriptRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceComplianceScriptRequestBuilder instantiates a new DeviceComplianceScriptRequestBuilder and sets the default values.
func NewDeviceComplianceScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceComplianceScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) CreateDeleteRequestInformation(options *DeviceComplianceScriptRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) CreateGetRequestInformation(options *DeviceComplianceScriptRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) CreatePatchRequestInformation(options *DeviceComplianceScriptRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) Delete(options *DeviceComplianceScriptRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceComplianceScriptRequestBuilder) DeviceRunStates()(*if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92.DeviceRunStatesRequestBuilder) {
    return if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceComplianceScripts.item.deviceRunStates.item collection
func (m *DeviceComplianceScriptRequestBuilder) DeviceRunStatesById(id string)(*i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346.DeviceComplianceScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScriptDeviceState_id"] = id
    }
    return i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346.NewDeviceComplianceScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) Get(options *DeviceComplianceScriptRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceComplianceScript() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceComplianceScript), nil
}
// Patch the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptRequestBuilder) Patch(options *DeviceComplianceScriptRequestBuilderPatchOptions)(error) {
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
func (m *DeviceComplianceScriptRequestBuilder) RunSummary()(*i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a.RunSummaryRequestBuilder) {
    return i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
