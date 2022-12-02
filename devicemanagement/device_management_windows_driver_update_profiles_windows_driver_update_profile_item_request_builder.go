package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters a collection of windows driver update profiles
type DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetQueryParameters
}
// DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Assign()(*DeviceManagementWindowsDriverUpdateProfilesItemAssignRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Assignments()(*DeviceManagementWindowsDriverUpdateProfilesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementWindowsDriverUpdateProfilesItemAssignmentsWindowsDriverUpdateProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateProfileAssignment%2Did"] = id
    }
    return NewDeviceManagementWindowsDriverUpdateProfilesItemAssignmentsWindowsDriverUpdateProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    m := &DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles/{windowsDriverUpdateProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder instantiates a new WindowsDriverUpdateProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of windows driver update profiles
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property windowsDriverUpdateProfiles for deviceManagement
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// DriverInventories provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) DriverInventories()(*DeviceManagementWindowsDriverUpdateProfilesItemDriverInventoriesRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesItemDriverInventoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DriverInventoriesById provides operations to manage the driverInventories property of the microsoft.graph.windowsDriverUpdateProfile entity.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) DriverInventoriesById(id string)(*DeviceManagementWindowsDriverUpdateProfilesItemDriverInventoriesWindowsDriverUpdateInventoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateInventory%2Did"] = id
    }
    return NewDeviceManagementWindowsDriverUpdateProfilesItemDriverInventoriesWindowsDriverUpdateInventoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExecuteAction provides operations to call the executeAction method.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) ExecuteAction()(*DeviceManagementWindowsDriverUpdateProfilesItemExecuteActionRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesItemExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get a collection of windows driver update profiles
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// Patch update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// SyncInventory provides operations to call the syncInventory method.
func (m *DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) SyncInventory()(*DeviceManagementWindowsDriverUpdateProfilesItemSyncInventoryRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesItemSyncInventoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
