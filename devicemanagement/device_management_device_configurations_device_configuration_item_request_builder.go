package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters the device configurations.
type DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) Assign()(*DeviceManagementDeviceConfigurationsItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignedAccessMultiModeProfiles provides operations to call the assignedAccessMultiModeProfiles method.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) AssignedAccessMultiModeProfiles()(*DeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemAssignedAccessMultiModeProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) Assignments()(*DeviceManagementDeviceConfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceConfigurationsItemAssignmentsDeviceConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsItemAssignmentsDeviceConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    m := &DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the device configurations.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceManagementDeviceConfigurationsItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*DeviceManagementDeviceConfigurationsItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatuses()(*DeviceManagementDeviceConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*DeviceManagementDeviceConfigurationsItemDeviceStatusesDeviceConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationDeviceStatus%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsItemDeviceStatusesDeviceConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatusOverview()(*DeviceManagementDeviceConfigurationsItemDeviceStatusOverviewRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device configurations.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// GetOmaSettingPlainTextValueWithSecretReferenceValueId provides operations to call the getOmaSettingPlainTextValue method.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) GetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*DeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId);
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) GroupAssignments()(*DeviceManagementDeviceConfigurationsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceManagementDeviceConfigurationsItemGroupAssignmentsDeviceConfigurationGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationGroupAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsItemGroupAssignmentsDeviceConfigurationGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatuses()(*DeviceManagementDeviceConfigurationsItemUserStatusesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatusesById(id string)(*DeviceManagementDeviceConfigurationsItemUserStatusesDeviceConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationUserStatus%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsItemUserStatusesDeviceConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatusOverview()(*DeviceManagementDeviceConfigurationsItemUserStatusOverviewRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsPrivacyAccessControls provides operations to call the windowsPrivacyAccessControls method.
func (m *DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) WindowsPrivacyAccessControls()(*DeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsItemWindowsPrivacyAccessControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
