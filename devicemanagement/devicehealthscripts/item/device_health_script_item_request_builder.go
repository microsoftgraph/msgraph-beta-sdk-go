package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceHealthScriptItemRequestBuilderDeleteOptions options for Delete
type DeviceHealthScriptItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceHealthScriptItemRequestBuilderGetOptions options for Get
type DeviceHealthScriptItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceHealthScriptItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *DeviceHealthScriptItemRequestBuilder) Assign()(*i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.AssignRequestBuilder) {
    return i104d50d0755116efee9a104857696ef73520f3d25a36981d161874cc755ab31d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
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
func NewDeviceHealthScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptItemRequestBuilder) {
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
func NewDeviceHealthScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceHealthScripts for deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceHealthScriptItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of device health scripts associated with the tenant.
func (m *DeviceHealthScriptItemRequestBuilder) CreateGetRequestInformation(options *DeviceHealthScriptItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceHealthScripts in deviceManagement
func (m *DeviceHealthScriptItemRequestBuilder) CreatePatchRequestInformation(options *DeviceHealthScriptItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
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
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates the deviceRunStates property
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
func (m *DeviceHealthScriptItemRequestBuilder) Get(options *DeviceHealthScriptItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceHealthScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceHealthScriptable), nil
}
// GetGlobalScriptHighestAvailableVersion the getGlobalScriptHighestAvailableVersion property
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
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RunSummary the runSummary property
func (m *DeviceHealthScriptItemRequestBuilder) RunSummary()(*i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.RunSummaryRequestBuilder) {
    return i444c364745747926b04609764de913a9054cabea4bd19dc8dae994f4e3b2581f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateGlobalScript the updateGlobalScript property
func (m *DeviceHealthScriptItemRequestBuilder) UpdateGlobalScript()(*i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.UpdateGlobalScriptRequestBuilder) {
    return i6e62c4306f99236ef789e61f87c4f907a053a18784c8fd89157032e7932b396e.NewUpdateGlobalScriptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
