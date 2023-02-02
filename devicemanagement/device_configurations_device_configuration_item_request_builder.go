package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationsDeviceConfigurationItemRequestBuilder provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationsDeviceConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters the device configurations.
type DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) Assignments()(*DeviceConfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) AssignmentsById(id string)(*DeviceConfigurationsItemAssignmentsDeviceConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsItemAssignmentsDeviceConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceConfigurationId *string)(*DeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    m := &DeviceConfigurationsDeviceConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceConfigurationId != nil {
        urlTplParams["deviceConfiguration%2Did"] = *deviceConfigurationId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceConfigurationsDeviceConfigurationItemRequestBuilder instantiates a new DeviceConfigurationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationsDeviceConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceConfigurationsItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceConfigurationsItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*DeviceConfigurationsItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatuses()(*DeviceConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewDeviceConfigurationsItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatusesById(id string)(*DeviceConfigurationsItemDeviceStatusesDeviceConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsItemDeviceStatusesDeviceConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) DeviceStatusOverview()(*DeviceConfigurationsItemDeviceStatusOverviewRequestBuilder) {
    return NewDeviceConfigurationsItemDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the device configurations.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) GroupAssignments()(*DeviceConfigurationsItemGroupAssignmentsRequestBuilder) {
    return NewDeviceConfigurationsItemGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupAssignmentsById provides operations to manage the groupAssignments property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) GroupAssignmentsById(id string)(*DeviceConfigurationsItemGroupAssignmentsDeviceConfigurationGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsItemGroupAssignmentsDeviceConfigurationGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) MicrosoftGraphAssign()(*DeviceConfigurationsItemMicrosoftGraphAssignAssignRequestBuilder) {
    return NewDeviceConfigurationsItemMicrosoftGraphAssignAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphAssignedAccessMultiModeProfiles provides operations to call the assignedAccessMultiModeProfiles method.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) MicrosoftGraphAssignedAccessMultiModeProfiles()(*DeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesRequestBuilder) {
    return NewDeviceConfigurationsItemMicrosoftGraphAssignedAccessMultiModeProfilesAssignedAccessMultiModeProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueId provides operations to call the getOmaSettingPlainTextValue method.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) MicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueId(secretReferenceValueId *string)(*DeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return NewDeviceConfigurationsItemMicrosoftGraphGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, secretReferenceValueId)
}
// MicrosoftGraphWindowsPrivacyAccessControls provides operations to call the windowsPrivacyAccessControls method.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) MicrosoftGraphWindowsPrivacyAccessControls()(*DeviceConfigurationsItemMicrosoftGraphWindowsPrivacyAccessControlsWindowsPrivacyAccessControlsRequestBuilder) {
    return NewDeviceConfigurationsItemMicrosoftGraphWindowsPrivacyAccessControlsWindowsPrivacyAccessControlsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property deviceConfigurations for deviceManagement
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the device configurations.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property deviceConfigurations in deviceManagement
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, requestConfiguration *DeviceConfigurationsDeviceConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatuses()(*DeviceConfigurationsItemUserStatusesRequestBuilder) {
    return NewDeviceConfigurationsItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatusesById(id string)(*DeviceConfigurationsItemUserStatusesDeviceConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsItemUserStatusesDeviceConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceConfiguration entity.
func (m *DeviceConfigurationsDeviceConfigurationItemRequestBuilder) UserStatusOverview()(*DeviceConfigurationsItemUserStatusOverviewRequestBuilder) {
    return NewDeviceConfigurationsItemUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
