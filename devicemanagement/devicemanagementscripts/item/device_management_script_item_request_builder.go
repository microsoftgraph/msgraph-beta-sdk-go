package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// DeviceManagementScriptItemRequestBuilder provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
type DeviceManagementScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementScriptItemRequestBuilderGetQueryParameters the list of device management scripts associated with the tenant.
type DeviceManagementScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementScriptItemRequestBuilderGetQueryParameters
}
// DeviceManagementScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceManagementScriptItemRequestBuilder) Assign()(*if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.AssignRequestBuilder) {
    return if33d8d67c34ee54fe873013b66ebc2b6e3dc6e392d8ab242161653be770f5d4b.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceManagementScriptItemRequestBuilder) Assignments()(*i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.AssignmentsRequestBuilder) {
    return i6e0733508da7ff0356e2ca316d29bfbd6cda93ef5db9a39817c2f424b503dec6.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.assignments.item collection
func (m *DeviceManagementScriptItemRequestBuilder) AssignmentsById(id string)(*i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.DeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return i84a45e8ade4a4e0a4b92c01abedac4601754d6c53b15136864f2af744f8c2b02.NewDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementScriptItemRequestBuilderInternal instantiates a new DeviceManagementScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementScriptItemRequestBuilder) {
    m := &DeviceManagementScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceManagementScripts/{deviceManagementScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementScriptItemRequestBuilder instantiates a new DeviceManagementScriptItemRequestBuilder and sets the default values.
func NewDeviceManagementScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceManagementScripts for deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceManagementScripts for deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceManagementScripts in deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceManagementScripts in deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, requestConfiguration *DeviceManagementScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property deviceManagementScripts for deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property deviceManagementScripts for deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceManagementScriptItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates the deviceRunStates property
func (m *DeviceManagementScriptItemRequestBuilder) DeviceRunStates()(*i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.DeviceRunStatesRequestBuilder) {
    return i06231952f18c3df2ed3c4845512be6748846cddc42430ea22f0110bc844935bb.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.deviceRunStates.item collection
func (m *DeviceManagementScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.DeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return i47cd5e12b88aa45ee308ba79574764b3fc2f141f8bbf112dc9af34ef38267657.NewDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceManagementScriptItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the list of device management scripts associated with the tenant.
func (m *DeviceManagementScriptItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceManagementScriptItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementScriptFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable), nil
}
// GroupAssignments the groupAssignments property
func (m *DeviceManagementScriptItemRequestBuilder) GroupAssignments()(*ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.GroupAssignmentsRequestBuilder) {
    return ifbbe46308813fcfd21a07bcac76e413b36786c8c9c9bcea463584d058d80cdaa.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.groupAssignments.item collection
func (m *DeviceManagementScriptItemRequestBuilder) GroupAssignmentsById(id string)(*i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.DeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return i3e402295bee0361812ae67ce9531f7061435900ac0c50c895a12d292d4bc0ec5.NewDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property deviceManagementScripts in deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, requestConfiguration *DeviceManagementScriptItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property deviceManagementScripts in deviceManagement
func (m *DeviceManagementScriptItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementScriptable, requestConfiguration *DeviceManagementScriptItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RunSummary the runSummary property
func (m *DeviceManagementScriptItemRequestBuilder) RunSummary()(*if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.RunSummaryRequestBuilder) {
    return if54f4ed088e5af17a948f9194c1f6e7aa23d72ff5ebeb5ae40c4ac2c1510ee2a.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates the userRunStates property
func (m *DeviceManagementScriptItemRequestBuilder) UserRunStates()(*i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.UserRunStatesRequestBuilder) {
    return i9d58942b600e533779502c28f94ba3cb6ca7a519ae97eef6c53e20d6be174561.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceManagementScripts.item.userRunStates.item collection
func (m *DeviceManagementScriptItemRequestBuilder) UserRunStatesById(id string)(*icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.DeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return icf8e74a74bc8538b6c9fa9f2bec4d9d624de373f52d2137852ef78d9754c2779.NewDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
