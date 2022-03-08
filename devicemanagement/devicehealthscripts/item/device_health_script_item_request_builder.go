package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// DeviceHealthScriptItemRequestBuilder provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
type DeviceHealthScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceHealthScriptItemRequestBuilderDeleteOptions options for Delete
type DeviceHealthScriptItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceHealthScriptItemRequestBuilderGetOptions options for Get
type DeviceHealthScriptItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceHealthScriptItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceHealthScriptItemRequestBuilderGetQueryParameters the list of device health scripts associated with the tenant.
type DeviceHealthScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceHealthScriptItemRequestBuilderPatchOptions options for Patch
type DeviceHealthScriptItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceHealthScriptItemRequestBuilder) Assign()(*i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.AssignRequestBuilder) {
    return i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceHealthScriptItemRequestBuilder) Assignments()(*id52be8ef6fb72e4097e092f65d25d26c2a99a6a4791081ca180c2a6ce4a48e57.AssignmentsRequestBuilder) {
    return id52be8ef6fb72e4097e092f65d25d26c2a99a6a4791081ca180c2a6ce4a48e57.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceHealthScripts.item.assignments.item collection
func (m *DeviceHealthScriptItemRequestBuilder) AssignmentsById(id string)(*i216848b587cfb32e95d29eb2e1cf76aba1de54cac3b4b002cdfe641e20ee3e0d.DeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment_id"] = id
    }
    return i216848b587cfb32e95d29eb2e1cf76aba1de54cac3b4b002cdfe641e20ee3e0d.NewDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceHealthScriptItemRequestBuilderInternal instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptItemRequestBuilder) {
    m := &DeviceHealthScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceHealthScriptItemRequestBuilder instantiates a new DeviceHealthScriptItemRequestBuilder and sets the default values.
func NewDeviceHealthScriptItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceHealthScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceHealthScriptItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptItemRequestBuilder) CreateGetRequestInformation(options *DeviceHealthScriptItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) CreatePatchRequestInformation(options *DeviceHealthScriptItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) Delete(options *DeviceHealthScriptItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceHealthScriptItemRequestBuilder) DeviceRunStates()(*i05314609e57e0d9e5aa7e76bdcbb8ca16d49864aea0242fdfdced18942bc70c7.DeviceRunStatesRequestBuilder) {
    return i05314609e57e0d9e5aa7e76bdcbb8ca16d49864aea0242fdfdced18942bc70c7.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceHealthScripts.item.deviceRunStates.item collection
func (m *DeviceHealthScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i88afe480603ea409059f63721a157eb2e70f28e5bf25444bfb05d2baa0ea62fb.DeviceHealthScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptDeviceState_id"] = id
    }
    return i88afe480603ea409059f63721a157eb2e70f28e5bf25444bfb05d2baa0ea62fb.NewDeviceHealthScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptItemRequestBuilder) Get(options *DeviceHealthScriptItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceHealthScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceHealthScriptable), nil
}
func (m *DeviceHealthScriptItemRequestBuilder) GetGlobalScriptHighestAvailableVersion()(*i3b70df7db94f29e1c1c07c52b95071aa9e28aad20a4834af6f304eedc10d5e07.GetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return i3b70df7db94f29e1c1c07c52b95071aa9e28aad20a4834af6f304eedc10d5e07.NewGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetRemediationHistory provides operations to call the getRemediationHistory method.
func (m *DeviceHealthScriptItemRequestBuilder) GetRemediationHistory()(*i8502e1c0748c67731dee41442535b5e69e56e1c7ab9e6762145c06069143454d.GetRemediationHistoryRequestBuilder) {
    return i8502e1c0748c67731dee41442535b5e69e56e1c7ab9e6762145c06069143454d.NewGetRemediationHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) Patch(options *DeviceHealthScriptItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceHealthScriptItemRequestBuilder) RunSummary()(*i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.RunSummaryRequestBuilder) {
    return i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceHealthScriptItemRequestBuilder) UpdateGlobalScript()(*i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.UpdateGlobalScriptRequestBuilder) {
    return i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.NewUpdateGlobalScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
