package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters the device compliance policies.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetQueryParameters
}
// DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assign()(*DeviceCompliancePoliciesItemAssignRequestBuilder) {
    return NewDeviceCompliancePoliciesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Assignments()(*DeviceCompliancePoliciesItemAssignmentsRequestBuilder) {
    return NewDeviceCompliancePoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*DeviceCompliancePoliciesItemAssignmentsDeviceCompliancePolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicyAssignment%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemAssignmentsDeviceCompliancePolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    m := &DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder instantiates a new DeviceCompliancePolicyItemRequestBuilder and sets the default values.
func NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceSettingStateSummaries provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummaries()(*DeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceSettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceSettingStateSummariesById provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceSettingStateSummariesById(id string)(*DeviceCompliancePoliciesItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["settingStateDeviceSummary%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemDeviceSettingStateSummariesSettingStateDeviceSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatuses()(*DeviceCompliancePoliciesItemDeviceStatusesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceStatusesById provides operations to manage the deviceStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusesById(id string)(*DeviceCompliancePoliciesItemDeviceStatusesDeviceComplianceDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceDeviceStatus%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemDeviceStatusesDeviceComplianceDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// DeviceStatusOverview provides operations to manage the deviceStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) DeviceStatusOverview()(*DeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilder) {
    return NewDeviceCompliancePoliciesItemDeviceStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get the device compliance policies.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// Patch update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable), nil
}
// ScheduleActionsForRules provides operations to call the scheduleActionsForRules method.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduleActionsForRules()(*DeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduleActionsForRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScheduledActionsForRule provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*DeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ScheduledActionsForRuleById provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*DeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScheduledActionForRule%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemScheduledActionsForRuleDeviceComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// ToDeleteRequestInformation delete navigation property deviceCompliancePolicies for deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the device compliance policies.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceCompliancePolicies in deviceManagement
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceCompliancePolicyable, requestConfiguration *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatuses()(*DeviceCompliancePoliciesItemUserStatusesRequestBuilder) {
    return NewDeviceCompliancePoliciesItemUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserStatusesById provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusesById(id string)(*DeviceCompliancePoliciesItemUserStatusesDeviceComplianceUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceUserStatus%2Did"] = id
    }
    return NewDeviceCompliancePoliciesItemUserStatusesDeviceComplianceUserStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// UserStatusOverview provides operations to manage the userStatusOverview property of the microsoft.graph.deviceCompliancePolicy entity.
func (m *DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) UserStatusOverview()(*DeviceCompliancePoliciesItemUserStatusOverviewRequestBuilder) {
    return NewDeviceCompliancePoliciesItemUserStatusOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
