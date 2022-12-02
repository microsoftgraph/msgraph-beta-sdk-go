package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters the device compliance policies.
type DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assign()(*DeviceManagementDeviceCompliancePoliciesItemAssignRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assignments()(*DeviceManagementDeviceCompliancePoliciesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementDeviceCompliancePoliciesItemAssignmentsDeviceCompliancePolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyAssignment%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesItemAssignmentsDeviceCompliancePolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    m := &DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the device compliance policies.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceManagementDeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*DeviceManagementDeviceCompliancePoliciesItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatuses()(*DeviceManagementDeviceCompliancePoliciesItemDeviceStatusesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusesById(id string)(*DeviceManagementDeviceCompliancePoliciesItemDeviceStatusesDeviceComplianceDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceDeviceStatus%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesItemDeviceStatusesDeviceComplianceDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*DeviceManagementDeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the device compliance policies.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// Patch update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// ScheduleActionsForRules provides operations to call the scheduleActionsForRules method.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*DeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRule provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScheduledActionForRule%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatuses()(*DeviceManagementDeviceCompliancePoliciesItemUserStatusesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusesById(id string)(*DeviceManagementDeviceCompliancePoliciesItemUserStatusesDeviceComplianceUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceUserStatus%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesItemUserStatusesDeviceComplianceUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*DeviceManagementDeviceCompliancePoliciesItemUserStatusOverviewRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesItemUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
