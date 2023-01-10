package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
type CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters list of all compliance policies
type CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetQueryParameters
}
// CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Assign()(*CompliancePoliciesItemAssignRequestBuilder) {
    return NewCompliancePoliciesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Assignments()(*CompliancePoliciesItemAssignmentsRequestBuilder) {
    return NewCompliancePoliciesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) AssignmentsById(id string)(*CompliancePoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyAssignment%2Did"] = id
    }
    return NewCompliancePoliciesItemAssignmentsDeviceManagementConfigurationPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderInternal instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) {
    m := &CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder instantiates a new DeviceManagementCompliancePolicyItemRequestBuilder and sets the default values.
func NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property compliancePolicies for deviceManagement
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get list of all compliance policies
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable), nil
}
// Patch update the navigation property compliancePolicies in deviceManagement
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementCompliancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable), nil
}
// ScheduledActionsForRule provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRule()(*CompliancePoliciesItemScheduledActionsForRuleRequestBuilder) {
    return NewCompliancePoliciesItemScheduledActionsForRuleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ScheduledActionsForRuleById provides operations to manage the scheduledActionsForRule property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) ScheduledActionsForRuleById(id string)(*CompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementComplianceScheduledActionForRule%2Did"] = id
    }
    return NewCompliancePoliciesItemScheduledActionsForRuleDeviceManagementComplianceScheduledActionForRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SetScheduledActions provides operations to call the setScheduledActions method.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) SetScheduledActions()(*CompliancePoliciesItemSetScheduledActionsRequestBuilder) {
    return NewCompliancePoliciesItemSetScheduledActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) Settings()(*CompliancePoliciesItemSettingsRequestBuilder) {
    return NewCompliancePoliciesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.deviceManagementCompliancePolicy entity.
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) SettingsById(id string)(*CompliancePoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSetting%2Did"] = id
    }
    return NewCompliancePoliciesItemSettingsDeviceManagementConfigurationSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property compliancePolicies for deviceManagement
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation list of all compliance policies
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property compliancePolicies in deviceManagement
func (m *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementCompliancePolicyable, requestConfiguration *CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
