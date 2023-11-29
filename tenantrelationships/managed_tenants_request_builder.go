package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsRequestBuilder provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
type ManagedTenantsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
    return NewManagedTenantsAggregatedPolicyCompliancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppPerformances provides operations to manage the appPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AppPerformances()(*ManagedTenantsAppPerformancesRequestBuilder) {
    return NewManagedTenantsAppPerformancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) AuditEvents()(*ManagedTenantsAuditEventsRequestBuilder) {
    return NewManagedTenantsAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcConnections provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcConnections()(*ManagedTenantsCloudPcConnectionsRequestBuilder) {
    return NewManagedTenantsCloudPcConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcDevices provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcDevices()(*ManagedTenantsCloudPcDevicesRequestBuilder) {
    return NewManagedTenantsCloudPcDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPcsOverview provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) CloudPcsOverview()(*ManagedTenantsCloudPcsOverviewRequestBuilder) {
    return NewManagedTenantsCloudPcsOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccessPolicyCoverages provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*ManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsRequestBuilderInternal instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsRequestBuilder) {
    m := &ManagedTenantsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants{?%24select,%24expand}", pathParameters),
    }
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
    return NewManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceAppPerformances provides operations to manage the deviceAppPerformances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceAppPerformances()(*ManagedTenantsDeviceAppPerformancesRequestBuilder) {
    return NewManagedTenantsDeviceAppPerformancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthStatuses provides operations to manage the deviceHealthStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) DeviceHealthStatuses()(*ManagedTenantsDeviceHealthStatusesRequestBuilder) {
    return NewManagedTenantsDeviceHealthStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*ManagedTenantsManagedDeviceCompliancesRequestBuilder) {
    return NewManagedTenantsManagedDeviceCompliancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceComplianceTrends provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*ManagedTenantsManagedDeviceComplianceTrendsRequestBuilder) {
    return NewManagedTenantsManagedDeviceComplianceTrendsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertLogs provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertLogs()(*ManagedTenantsManagedTenantAlertLogsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertLogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertRuleDefinitions provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitions()(*ManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlertRules provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlertRules()(*ManagedTenantsManagedTenantAlertRulesRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantAlerts provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantAlerts()(*ManagedTenantsManagedTenantAlertsRequestBuilder) {
    return NewManagedTenantsManagedTenantAlertsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantApiNotifications provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantApiNotifications()(*ManagedTenantsManagedTenantApiNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantApiNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantEmailNotifications provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantEmailNotifications()(*ManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    return NewManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedTenantTicketingEndpoints provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagedTenantTicketingEndpoints()(*ManagedTenantsManagedTenantTicketingEndpointsRequestBuilder) {
    return NewManagedTenantsManagedTenantTicketingEndpointsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementActions provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActions()(*ManagedTenantsManagementActionsRequestBuilder) {
    return NewManagedTenantsManagementActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementActionTenantDeploymentStatuses provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*ManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) {
    return NewManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementIntents provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementIntents()(*ManagedTenantsManagementIntentsRequestBuilder) {
    return NewManagedTenantsManagementIntentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateCollections provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollections()(*ManagedTenantsManagementTemplateCollectionsRequestBuilder) {
    return NewManagedTenantsManagementTemplateCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateCollectionTenantSummaries provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummaries()(*ManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return NewManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplates()(*ManagedTenantsManagementTemplatesRequestBuilder) {
    return NewManagedTenantsManagementTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateSteps provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateSteps()(*ManagedTenantsManagementTemplateStepsRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateStepTenantSummaries provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummaries()(*ManagedTenantsManagementTemplateStepTenantSummariesRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepTenantSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagementTemplateStepVersions provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*ManagedTenantsManagementTemplateStepVersionsRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MyRoles provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) MyRoles()(*ManagedTenantsMyRolesRequestBuilder) {
    return NewManagedTenantsMyRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ManagedTenantsRequestBuilder) TenantGroups()(*ManagedTenantsTenantGroupsRequestBuilder) {
    return NewManagedTenantsTenantGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) Tenants()(*ManagedTenantsTenantsRequestBuilder) {
    return NewManagedTenantsTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantsCustomizedInformation provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*ManagedTenantsTenantsCustomizedInformationRequestBuilder) {
    return NewManagedTenantsTenantsCustomizedInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantsDetailedInformation provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantsDetailedInformation()(*ManagedTenantsTenantsDetailedInformationRequestBuilder) {
    return NewManagedTenantsTenantsDetailedInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantTags provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) TenantTags()(*ManagedTenantsTenantTagsRequestBuilder) {
    return NewManagedTenantsTenantTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managedTenants for tenantRelationships
func (m *ManagedTenantsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *ManagedTenantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *ManagedTenantsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*ManagedTenantsWindowsDeviceMalwareStatesRequestBuilder) {
    return NewManagedTenantsWindowsDeviceMalwareStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsProtectionStates provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsRequestBuilder) WindowsProtectionStates()(*ManagedTenantsWindowsProtectionStatesRequestBuilder) {
    return NewManagedTenantsWindowsProtectionStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsRequestBuilder) {
    return NewManagedTenantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
