package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsRequestBuilder provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
type ManagedTenantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsRequestBuilderGetQueryParameters the operations available to interact with the multi-tenant management platform.
type ManagedTenantsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsRequestBuilderGetQueryParameters
}
// ManagedTenantsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AggregatedPolicyCompliances provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*ManagedTenantsAggregatedPolicyCompliancesRequestBuilder) {
    return NewManagedTenantsAggregatedPolicyCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AggregatedPolicyCompliancesById provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AggregatedPolicyCompliancesById(id string)(*ManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["aggregatedPolicyCompliance%2Did"] = id
    }
    return NewManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AuditEvents()(*ManagedTenantsAuditEventsRequestBuilder) {
    return NewManagedTenantsAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AuditEventsById(id string)(*ManagedTenantsAuditEventsAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent%2Did"] = id
    }
    return NewManagedTenantsAuditEventsAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcConnections provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcConnections()(*ManagedTenantsCloudPcConnectionsRequestBuilder) {
    return NewManagedTenantsCloudPcConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcConnectionsById provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcConnectionsById(id string)(*ManagedTenantsCloudPcConnectionsCloudPcConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcConnection%2Did"] = id
    }
    return NewManagedTenantsCloudPcConnectionsCloudPcConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcDevices provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcDevices()(*ManagedTenantsCloudPcDevicesRequestBuilder) {
    return NewManagedTenantsCloudPcDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcDevicesById provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcDevicesById(id string)(*ManagedTenantsCloudPcDevicesCloudPcDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDevice%2Did"] = id
    }
    return NewManagedTenantsCloudPcDevicesCloudPcDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcsOverview provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcsOverview()(*ManagedTenantsCloudPcsOverviewRequestBuilder) {
    return NewManagedTenantsCloudPcsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcsOverviewById provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcsOverviewById(id string)(*ManagedTenantsCloudPcsOverviewCloudPcOverviewTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOverview%2DtenantId"] = id
    }
    return NewManagedTenantsCloudPcsOverviewCloudPcOverviewTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicyCoverages provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPolicyCoveragesById provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoveragesById(id string)(*ManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicyCoverage%2Did"] = id
    }
    return NewManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedTenantsRequestBuilderInternal instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    m := &ManagedTenantsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedTenantsRequestBuilder instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CredentialUserRegistrationsSummaries provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*ManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) {
    return NewManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationsSummariesById provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CredentialUserRegistrationsSummariesById(id string)(*ManagedTenantsCredentialUserRegistrationsSummariesCredentialUserRegistrationsSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationsSummary%2Did"] = id
    }
    return NewManagedTenantsCredentialUserRegistrationsSummariesCredentialUserRegistrationsSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property managedTenants for tenantRelationships
func (m *ManagedTenantsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = id
    }
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// ManagedDeviceCompliances provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ManagedTenantsManagedDeviceCompliancesRequestBuilder) {
    return NewManagedTenantsManagedDeviceCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceCompliancesById provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliancesById(id string)(*ManagedTenantsManagedDeviceCompliancesManagedDeviceComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceCompliance%2Did"] = id
    }
    return NewManagedTenantsManagedDeviceCompliancesManagedDeviceComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceComplianceTrends provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*ManagedTenantsManagedDeviceComplianceTrendsRequestBuilder) {
    return NewManagedTenantsManagedDeviceComplianceTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceComplianceTrendsById provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrendsById(id string)(*ManagedTenantsManagedDeviceComplianceTrendsManagedDeviceComplianceTrendItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceComplianceTrend%2Did"] = id
    }
    return NewManagedTenantsManagedDeviceComplianceTrendsManagedDeviceComplianceTrendItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertLogs provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertLogs()(*ManagedTenantsManagedTenantAlertLogsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertLogsById provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertLogsById(id string)(*ManagedTenantsManagedTenantAlertLogsManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertLogsManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitions provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitions()(*ManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitionsById provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitionsById(id string)(*ManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRuleDefinition%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRules provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRules()(*ManagedTenantsManagedTenantAlertRulesRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRulesById provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRulesById(id string)(*ManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRule%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlerts provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlerts()(*ManagedTenantsManagedTenantAlertsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertsById provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertsById(id string)(*ManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlert%2Did"] = id
    }
    return NewManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantApiNotifications provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantApiNotifications()(*ManagedTenantsManagedTenantApiNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantApiNotificationsById provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantApiNotificationsById(id string)(*ManagedTenantsManagedTenantApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return NewManagedTenantsManagedTenantApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantEmailNotifications provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantEmailNotifications()(*ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantEmailNotificationsById provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantEmailNotificationsById(id string)(*ManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return NewManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantTicketingEndpoints provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantTicketingEndpoints()(*ManagedTenantsManagedTenantTicketingEndpointsRequestBuilder) {
    return NewManagedTenantsManagedTenantTicketingEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantTicketingEndpointsById provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantTicketingEndpointsById(id string)(*ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantTicketingEndpoint%2Did"] = id
    }
    return NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActions provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActions()(*ManagedTenantsManagementActionsRequestBuilder) {
    return NewManagedTenantsManagementActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionsById provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionsById(id string)(*ManagedTenantsManagementActionsManagementActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementAction%2Did"] = id
    }
    return NewManagedTenantsManagementActionsManagementActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatuses provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*ManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) {
    return NewManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatusesById provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatusesById(id string)(*ManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementActionTenantDeploymentStatus%2Did"] = id
    }
    return NewManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementIntents provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementIntents()(*ManagedTenantsManagementIntentsRequestBuilder) {
    return NewManagedTenantsManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementIntentsById provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementIntentsById(id string)(*ManagedTenantsManagementIntentsManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementIntent%2Did"] = id
    }
    return NewManagedTenantsManagementIntentsManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollections provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollections()(*ManagedTenantsManagementTemplateCollectionsRequestBuilder) {
    return NewManagedTenantsManagementTemplateCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionsById provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionsById(id string)(*ManagedTenantsManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollection%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummaries provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummaries()(*ManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return NewManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummariesById provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummariesById(id string)(*ManagedTenantsManagementTemplateCollectionTenantSummariesManagementTemplateCollectionTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollectionTenantSummary%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateCollectionTenantSummariesManagementTemplateCollectionTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplates()(*ManagedTenantsManagementTemplatesRequestBuilder) {
    return NewManagedTenantsManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplatesById provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplatesById(id string)(*ManagedTenantsManagementTemplatesManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplate%2Did"] = id
    }
    return NewManagedTenantsManagementTemplatesManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateSteps provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateSteps()(*ManagedTenantsManagementTemplateStepsRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepsById provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepsById(id string)(*ManagedTenantsManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStep%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepTenantSummaries provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummaries()(*ManagedTenantsManagementTemplateStepTenantSummariesRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepTenantSummariesById provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummariesById(id string)(*ManagedTenantsManagementTemplateStepTenantSummariesManagementTemplateStepTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepTenantSummary%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateStepTenantSummariesManagementTemplateStepTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepVersions provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*ManagedTenantsManagementTemplateStepVersionsRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepVersionsById provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersionsById(id string)(*ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepVersion%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MyRoles provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) MyRoles()(*ManagedTenantsMyRolesRequestBuilder) {
    return NewManagedTenantsMyRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyRolesById provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) MyRolesById(id string)(*ManagedTenantsMyRolesMyRoleTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["myRole%2DtenantId"] = id
    }
    return NewManagedTenantsMyRolesMyRoleTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedTenants in tenantRelationships
func (m *ManagedTenantsRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedTenantsRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// TenantGroups provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantGroups()(*ManagedTenantsTenantGroupsRequestBuilder) {
    return NewManagedTenantsTenantGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantGroupsById provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantGroupsById(id string)(*ManagedTenantsTenantGroupsTenantGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantGroup%2Did"] = id
    }
    return NewManagedTenantsTenantGroupsTenantGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) Tenants()(*ManagedTenantsTenantsRequestBuilder) {
    return NewManagedTenantsTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsById(id string)(*ManagedTenantsTenantsTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenant%2Did"] = id
    }
    return NewManagedTenantsTenantsTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsCustomizedInformation provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*ManagedTenantsTenantsCustomizedInformationRequestBuilder) {
    return NewManagedTenantsTenantsCustomizedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsCustomizedInformationById provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformationById(id string)(*ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantCustomizedInformation%2Did"] = id
    }
    return NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsDetailedInformation provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformation()(*ManagedTenantsTenantsDetailedInformationRequestBuilder) {
    return NewManagedTenantsTenantsDetailedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsDetailedInformationById provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformationById(id string)(*ManagedTenantsTenantsDetailedInformationTenantDetailedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantDetailedInformation%2Did"] = id
    }
    return NewManagedTenantsTenantsDetailedInformationTenantDetailedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantTags provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantTags()(*ManagedTenantsTenantTagsRequestBuilder) {
    return NewManagedTenantsTenantTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantTagsById provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantTagsById(id string)(*ManagedTenantsTenantTagsTenantTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantTag%2Did"] = id
    }
    return NewManagedTenantsTenantTagsTenantTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property managedTenants for tenantRelationships
func (m *ManagedTenantsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedTenants in tenantRelationships
func (m *ManagedTenantsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedTenantsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WindowsDeviceMalwareStates provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*ManagedTenantsWindowsDeviceMalwareStatesRequestBuilder) {
    return NewManagedTenantsWindowsDeviceMalwareStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDeviceMalwareStatesById provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStatesById(id string)(*ManagedTenantsWindowsDeviceMalwareStatesWindowsDeviceMalwareStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState%2Did"] = id
    }
    return NewManagedTenantsWindowsDeviceMalwareStatesWindowsDeviceMalwareStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsProtectionStates provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStates()(*ManagedTenantsWindowsProtectionStatesRequestBuilder) {
    return NewManagedTenantsWindowsProtectionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionStatesById provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStatesById(id string)(*ManagedTenantsWindowsProtectionStatesWindowsProtectionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsProtectionState%2Did"] = id
    }
    return NewManagedTenantsWindowsProtectionStatesWindowsProtectionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
