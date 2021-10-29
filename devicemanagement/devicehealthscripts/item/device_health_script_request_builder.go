package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i05314609e57e0d9e5aa7e76bdcbb8ca16d49864aea0242fdfdced18942bc70c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates"
    i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/assign"
    i3b70df7db94f29e1c1c07c52b95071aa9e28aad20a4834af6f304eedc10d5e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/getglobalscripthighestavailableversion"
    i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/runsummary"
    i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/updateglobalscript"
    i8502e1c0748c67731dee41442535b5e69e56e1c7ab9e6762145c06069143454d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/getremediationhistory"
    id52be8ef6fb72e4097e092f65d25d26c2a99a6a4791081ca180c2a6ce4a48e57 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/assignments"
    i216848b587cfb32e95d29eb2e1cf76aba1de54cac3b4b002cdfe641e20ee3e0d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/assignments/item"
    i88afe480603ea409059f63721a157eb2e70f28e5bf25444bfb05d2baa0ea62fb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicehealthscripts/item/devicerunstates/item"
)

// Builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}
type DeviceHealthScriptRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type DeviceHealthScriptRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type DeviceHealthScriptRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceHealthScriptRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The list of device health scripts associated with the tenant.
type DeviceHealthScriptRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type DeviceHealthScriptRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceHealthScriptRequestBuilder) Assign()(*i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.AssignRequestBuilder) {
    return i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceHealthScriptRequestBuilder) Assignments()(*id52be8ef6fb72e4097e092f65d25d26c2a99a6a4791081ca180c2a6ce4a48e57.AssignmentsRequestBuilder) {
    return id52be8ef6fb72e4097e092f65d25d26c2a99a6a4791081ca180c2a6ce4a48e57.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.deviceHealthScripts.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceHealthScriptRequestBuilder) AssignmentsById(id string)(*i216848b587cfb32e95d29eb2e1cf76aba1de54cac3b4b002cdfe641e20ee3e0d.DeviceHealthScriptAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment_id"] = id
    }
    return i216848b587cfb32e95d29eb2e1cf76aba1de54cac3b4b002cdfe641e20ee3e0d.NewDeviceHealthScriptAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DeviceHealthScriptRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceHealthScriptRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptRequestBuilder) {
    m := &DeviceHealthScriptRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeviceHealthScriptRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceHealthScriptRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptRequestBuilderInternal(urlParams, requestAdapter)
}
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) CreateDeleteRequestInformation(options *DeviceHealthScriptRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) CreateGetRequestInformation(options *DeviceHealthScriptRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) CreatePatchRequestInformation(options *DeviceHealthScriptRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) Delete(options *DeviceHealthScriptRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceHealthScriptRequestBuilder) DeviceRunStates()(*i05314609e57e0d9e5aa7e76bdcbb8ca16d49864aea0242fdfdced18942bc70c7.DeviceRunStatesRequestBuilder) {
    return i05314609e57e0d9e5aa7e76bdcbb8ca16d49864aea0242fdfdced18942bc70c7.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.deviceHealthScripts.item.deviceRunStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceHealthScriptRequestBuilder) DeviceRunStatesById(id string)(*i88afe480603ea409059f63721a157eb2e70f28e5bf25444bfb05d2baa0ea62fb.DeviceHealthScriptDeviceStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptDeviceState_id"] = id
    }
    return i88afe480603ea409059f63721a157eb2e70f28e5bf25444bfb05d2baa0ea62fb.NewDeviceHealthScriptDeviceStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) Get(options *DeviceHealthScriptRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceHealthScript() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScript), nil
}
func (m *DeviceHealthScriptRequestBuilder) GetGlobalScriptHighestAvailableVersion()(*i3b70df7db94f29e1c1c07c52b95071aa9e28aad20a4834af6f304eedc10d5e07.GetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return i3b70df7db94f29e1c1c07c52b95071aa9e28aad20a4834af6f304eedc10d5e07.NewGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \deviceManagement\deviceHealthScripts\{deviceHealthScript-id}\microsoft.graph.getRemediationHistory()
func (m *DeviceHealthScriptRequestBuilder) GetRemediationHistory()(*i8502e1c0748c67731dee41442535b5e69e56e1c7ab9e6762145c06069143454d.GetRemediationHistoryRequestBuilder) {
    return i8502e1c0748c67731dee41442535b5e69e56e1c7ab9e6762145c06069143454d.NewGetRemediationHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The list of device health scripts associated with the tenant.
// Parameters:
//  - options : Options for the request
func (m *DeviceHealthScriptRequestBuilder) Patch(options *DeviceHealthScriptRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceHealthScriptRequestBuilder) RunSummary()(*i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.RunSummaryRequestBuilder) {
    return i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceHealthScriptRequestBuilder) UpdateGlobalScript()(*i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.UpdateGlobalScriptRequestBuilder) {
    return i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.NewUpdateGlobalScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
