package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/runsummary"
    i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assign"
    i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assignments"
    if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates"
    i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/devicerunstates/item"
    ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecompliancescripts/item/assignments/item"
)

// DeviceComplianceScriptItemRequestBuilder provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
type DeviceComplianceScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceComplianceScriptItemRequestBuilderGetQueryParameters the list of device compliance scripts associated with the tenant.
type DeviceComplianceScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceComplianceScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceComplianceScriptItemRequestBuilderGetQueryParameters
}
// DeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceComplianceScriptItemRequestBuilder) Assign()(*i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee.AssignRequestBuilder) {
    return i45e12b0171a260bd111ff8eb9860e05c962a02a5d44af5e58aa4e28a348596ee.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceComplianceScriptItemRequestBuilder) Assignments()(*i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa.AssignmentsRequestBuilder) {
    return i4e865afe945f8eacc484b71a81eb991983bd4d613cfa3059de4617383b52a4aa.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceComplianceScripts.item.assignments.item collection
func (m *DeviceComplianceScriptItemRequestBuilder) AssignmentsById(id string)(*ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4.DeviceHealthScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScriptAssignment%2Did"] = id
    }
    return ic7b354fd2efd9673bdefbd282eeff6bb6291327de8753e174e035a4bfada36a4.NewDeviceHealthScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceComplianceScriptItemRequestBuilderInternal instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptItemRequestBuilder) {
    m := &DeviceComplianceScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceComplianceScripts/{deviceComplianceScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceComplianceScriptItemRequestBuilder instantiates a new DeviceComplianceScriptItemRequestBuilder and sets the default values.
func NewDeviceComplianceScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceComplianceScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceComplianceScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceComplianceScripts for deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceComplianceScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates the deviceRunStates property
func (m *DeviceComplianceScriptItemRequestBuilder) DeviceRunStates()(*if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92.DeviceRunStatesRequestBuilder) {
    return if542347d2ac966ddf56d493104cd69d633b13bf4470c216e225aef6a4df09c92.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceComplianceScripts.item.deviceRunStates.item collection
func (m *DeviceComplianceScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346.DeviceComplianceScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScriptDeviceState%2Did"] = id
    }
    return i8390be3634cd13ee17f47bcaaaed3a3064fb2f729417ec9de7372c477682c346.NewDeviceComplianceScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device compliance scripts associated with the tenant.
func (m *DeviceComplianceScriptItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceComplianceScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceScriptFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable), nil
}
// Patch update the navigation property deviceComplianceScripts in deviceManagement
func (m *DeviceComplianceScriptItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceScriptable, requestConfiguration *DeviceComplianceScriptItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RunSummary the runSummary property
func (m *DeviceComplianceScriptItemRequestBuilder) RunSummary()(*i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a.RunSummaryRequestBuilder) {
    return i0393d87a31f08c6eabcc775ffc153f7b6a6405f40f97e2be37b9016212579e4a.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
