package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagedTenantsRequestBuilder provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
type ManagedtenantsManagedTenantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagedTenantsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedTenantsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedtenantsManagedTenantsRequestBuilderGetQueryParameters the operations available to interact with the multi-tenant management platform.
type ManagedtenantsManagedTenantsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsManagedTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagedTenantsRequestBuilderGetQueryParameters
}
// ManagedtenantsManagedTenantsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedTenantsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AggregatedPolicyCompliances provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyCompliancesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*ManagedtenantsAggregatedpolicycompliancesAggregatedPolicyCompliancesRequestBuilder) {
    return NewManagedtenantsAggregatedpolicycompliancesAggregatedPolicyCompliancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppPerformances provides operations to manage the appPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsAppperformancesAppPerformancesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) AppPerformances()(*ManagedtenantsAppperformancesAppPerformancesRequestBuilder) {
    return NewManagedtenantsAppperformancesAppPerformancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsAuditeventsAuditEventsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) AuditEvents()(*ManagedtenantsAuditeventsAuditEventsRequestBuilder) {
    return NewManagedtenantsAuditeventsAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcConnections provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsCloudpcconnectionsCloudPcConnectionsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) CloudPcConnections()(*ManagedtenantsCloudpcconnectionsCloudPcConnectionsRequestBuilder) {
    return NewManagedtenantsCloudpcconnectionsCloudPcConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcDevices provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsCloudpcdevicesCloudPcDevicesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) CloudPcDevices()(*ManagedtenantsCloudpcdevicesCloudPcDevicesRequestBuilder) {
    return NewManagedtenantsCloudpcdevicesCloudPcDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcsOverview provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsCloudpcsoverviewCloudPcsOverviewRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) CloudPcsOverview()(*ManagedtenantsCloudpcsoverviewCloudPcsOverviewRequestBuilder) {
    return NewManagedtenantsCloudpcsoverviewCloudPcsOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccessPolicyCoverages provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*ManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewManagedtenantsConditionalaccesspolicycoveragesConditionalAccessPolicyCoveragesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsManagedTenantsRequestBuilderInternal instantiates a new ManagedtenantsManagedTenantsRequestBuilder and sets the default values.
func NewManagedtenantsManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedTenantsRequestBuilder) {
    m := &ManagedtenantsManagedTenantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagedTenantsRequestBuilder instantiates a new ManagedtenantsManagedTenantsRequestBuilder and sets the default values.
func NewManagedtenantsManagedTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CredentialUserRegistrationsSummaries provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsCredentialuserregistrationssummariesCredentialUserRegistrationsSummariesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*ManagedtenantsCredentialuserregistrationssummariesCredentialUserRegistrationsSummariesRequestBuilder) {
    return NewManagedtenantsCredentialuserregistrationssummariesCredentialUserRegistrationsSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property managedTenants for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedTenantsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceAppPerformances provides operations to manage the deviceAppPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) DeviceAppPerformances()(*ManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilder) {
    return NewManagedtenantsDeviceappperformancesDeviceAppPerformancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*ManagedtenantsDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewManagedtenantsDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthStatuses provides operations to manage the deviceHealthStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) DeviceHealthStatuses()(*ManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilder) {
    return NewManagedtenantsDevicehealthstatusesDeviceHealthStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the operations available to interact with the multi-tenant management platform.
// returns a ManagedTenantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// ManagedDeviceCompliances provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManageddevicecompliancesManagedDeviceCompliancesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ManagedtenantsManageddevicecompliancesManagedDeviceCompliancesRequestBuilder) {
    return NewManagedtenantsManageddevicecompliancesManagedDeviceCompliancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceComplianceTrends provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManageddevicecompliancetrendsManagedDeviceComplianceTrendsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*ManagedtenantsManageddevicecompliancetrendsManagedDeviceComplianceTrendsRequestBuilder) {
    return NewManagedtenantsManageddevicecompliancetrendsManagedDeviceComplianceTrendsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertLogs provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantalertlogsManagedTenantAlertLogsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantAlertLogs()(*ManagedtenantsManagedtenantalertlogsManagedTenantAlertLogsRequestBuilder) {
    return NewManagedtenantsManagedtenantalertlogsManagedTenantAlertLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertRuleDefinitions provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantalertruledefinitionsManagedTenantAlertRuleDefinitionsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitions()(*ManagedtenantsManagedtenantalertruledefinitionsManagedTenantAlertRuleDefinitionsRequestBuilder) {
    return NewManagedtenantsManagedtenantalertruledefinitionsManagedTenantAlertRuleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertRules provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantalertrulesManagedTenantAlertRulesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantAlertRules()(*ManagedtenantsManagedtenantalertrulesManagedTenantAlertRulesRequestBuilder) {
    return NewManagedtenantsManagedtenantalertrulesManagedTenantAlertRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlerts provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantalertsManagedTenantAlertsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantAlerts()(*ManagedtenantsManagedtenantalertsManagedTenantAlertsRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsManagedTenantAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantApiNotifications provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantapinotificationsManagedTenantApiNotificationsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantApiNotifications()(*ManagedtenantsManagedtenantapinotificationsManagedTenantApiNotificationsRequestBuilder) {
    return NewManagedtenantsManagedtenantapinotificationsManagedTenantApiNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantEmailNotifications provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantemailnotificationsManagedTenantEmailNotificationsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantEmailNotifications()(*ManagedtenantsManagedtenantemailnotificationsManagedTenantEmailNotificationsRequestBuilder) {
    return NewManagedtenantsManagedtenantemailnotificationsManagedTenantEmailNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantTicketingEndpoints provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagedtenantticketingendpointsManagedTenantTicketingEndpointsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagedTenantTicketingEndpoints()(*ManagedtenantsManagedtenantticketingendpointsManagedTenantTicketingEndpointsRequestBuilder) {
    return NewManagedtenantsManagedtenantticketingendpointsManagedTenantTicketingEndpointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementActions provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementactionsManagementActionsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementActions()(*ManagedtenantsManagementactionsManagementActionsRequestBuilder) {
    return NewManagedtenantsManagementactionsManagementActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementActionTenantDeploymentStatuses provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementactiontenantdeploymentstatusesManagementActionTenantDeploymentStatusesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*ManagedtenantsManagementactiontenantdeploymentstatusesManagementActionTenantDeploymentStatusesRequestBuilder) {
    return NewManagedtenantsManagementactiontenantdeploymentstatusesManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementIntents provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementintentsManagementIntentsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementIntents()(*ManagedtenantsManagementintentsManagementIntentsRequestBuilder) {
    return NewManagedtenantsManagementintentsManagementIntentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateCollections provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplateCollections()(*ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionsRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateCollectionTenantSummaries provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummaries()(*ManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectiontenantsummariesManagementTemplateCollectionTenantSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplates()(*ManagedtenantsManagementtemplatesManagementTemplatesRequestBuilder) {
    return NewManagedtenantsManagementtemplatesManagementTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateSteps provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatestepsManagementTemplateStepsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplateSteps()(*ManagedtenantsManagementtemplatestepsManagementTemplateStepsRequestBuilder) {
    return NewManagedtenantsManagementtemplatestepsManagementTemplateStepsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateStepTenantSummaries provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummaries()(*ManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilder) {
    return NewManagedtenantsManagementtemplatesteptenantsummariesManagementTemplateStepTenantSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateStepVersions provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsManagementtemplatestepversionsManagementTemplateStepVersionsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*ManagedtenantsManagementtemplatestepversionsManagementTemplateStepVersionsRequestBuilder) {
    return NewManagedtenantsManagementtemplatestepversionsManagementTemplateStepVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MyRoles provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsMyrolesMyRolesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) MyRoles()(*ManagedtenantsMyrolesMyRolesRequestBuilder) {
    return NewManagedtenantsMyrolesMyRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedTenants in tenantRelationships
// returns a ManagedTenantable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedTenantsRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable), nil
}
// TenantGroups provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantgroupsTenantGroupsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) TenantGroups()(*ManagedtenantsTenantgroupsTenantGroupsRequestBuilder) {
    return NewManagedtenantsTenantgroupsTenantGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) Tenants()(*ManagedtenantsTenantsRequestBuilder) {
    return NewManagedtenantsTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantsCustomizedInformation provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*ManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilder) {
    return NewManagedtenantsTenantscustomizedinformationTenantsCustomizedInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantsDetailedInformation provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenantsdetailedinformationTenantsDetailedInformationRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) TenantsDetailedInformation()(*ManagedtenantsTenantsdetailedinformationTenantsDetailedInformationRequestBuilder) {
    return NewManagedtenantsTenantsdetailedinformationTenantsDetailedInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantTags provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsTenanttagsTenantTagsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) TenantTags()(*ManagedtenantsTenanttagsTenantTagsRequestBuilder) {
    return NewManagedtenantsTenanttagsTenantTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedTenants for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the operations available to interact with the multi-tenant management platform.
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property managedTenants in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedtenantsManagedTenantsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WindowsDeviceMalwareStates provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsWindowsdevicemalwarestatesWindowsDeviceMalwareStatesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*ManagedtenantsWindowsdevicemalwarestatesWindowsDeviceMalwareStatesRequestBuilder) {
    return NewManagedtenantsWindowsdevicemalwarestatesWindowsDeviceMalwareStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionStates provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) WindowsProtectionStates()(*ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) {
    return NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedtenantsManagedTenantsRequestBuilder when successful
func (m *ManagedtenantsManagedTenantsRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagedTenantsRequestBuilder) {
    return NewManagedtenantsManagedTenantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
