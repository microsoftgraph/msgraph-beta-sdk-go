package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/runsummary"
    i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments"
    i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assign"
    i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments"
    ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates"
    ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates"
    i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments/item"
    i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item"
    i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item"
    ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments/item"
)

// DeviceCustomAttributeShellScriptItemRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceCustomAttributeShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters the list of device custom attribute shell scripts associated with the tenant.
type DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters
}
// DeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Assign()(*i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.AssignRequestBuilder) {
    return i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Assignments()(*i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.AssignmentsRequestBuilder) {
    return i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.assignments.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) AssignmentsById(id string)(*ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.DeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment%2Did"] = id
    }
    return ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.NewDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptItemRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCustomAttributeShellScriptItemRequestBuilder instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStates()(*ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.DeviceRunStatesRequestBuilder) {
    return ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.deviceRunStates.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.DeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState%2Did"] = id
    }
    return i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.NewDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable), nil
}
// GroupAssignments the groupAssignments property
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignments()(*i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.GroupAssignmentsRequestBuilder) {
    return i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.groupAssignments.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.DeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment%2Did"] = id
    }
    return i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.NewDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCustomAttributeShellScriptable, requestConfiguration *DeviceCustomAttributeShellScriptItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) RunSummary()(*i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.RunSummaryRequestBuilder) {
    return i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates the userRunStates property
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStates()(*ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.UserRunStatesRequestBuilder) {
    return ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.userRunStates.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStatesById(id string)(*i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.DeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState%2Did"] = id
    }
    return i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.NewDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
