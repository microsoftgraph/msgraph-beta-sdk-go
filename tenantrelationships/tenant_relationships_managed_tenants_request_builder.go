package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// TenantRelationshipsManagedTenantsRequestBuilder provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
type TenantRelationshipsManagedTenantsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TenantRelationshipsManagedTenantsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TenantRelationshipsManagedTenantsRequestBuilderGetQueryParameters the operations available to interact with the multi-tenant management platform.
type TenantRelationshipsManagedTenantsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsManagedTenantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsManagedTenantsRequestBuilderGetQueryParameters
}
// TenantRelationshipsManagedTenantsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsManagedTenantsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AggregatedPolicyCompliances provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) AggregatedPolicyCompliances()(*TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AggregatedPolicyCompliancesById provides operations to manage the aggregatedPolicyCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) AggregatedPolicyCompliancesById(id string)(*TenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["aggregatedPolicyCompliance%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsAggregatedPolicyCompliancesAggregatedPolicyComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) AuditEvents()(*TenantRelationshipsManagedTenantsAuditEventsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) AuditEventsById(id string)(*TenantRelationshipsManagedTenantsAuditEventsAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsAuditEventsAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcConnections provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcConnections()(*TenantRelationshipsManagedTenantsCloudPcConnectionsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsCloudPcConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcConnectionsById provides operations to manage the cloudPcConnections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcConnectionsById(id string)(*TenantRelationshipsManagedTenantsCloudPcConnectionsCloudPcConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcConnection%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsCloudPcConnectionsCloudPcConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcDevices provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcDevices()(*TenantRelationshipsManagedTenantsCloudPcDevicesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsCloudPcDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcDevicesById provides operations to manage the cloudPcDevices property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcDevicesById(id string)(*TenantRelationshipsManagedTenantsCloudPcDevicesCloudPcDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcDevice%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsCloudPcDevicesCloudPcDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPcsOverview provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcsOverview()(*TenantRelationshipsManagedTenantsCloudPcsOverviewRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsCloudPcsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPcsOverviewById provides operations to manage the cloudPcsOverview property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CloudPcsOverviewById(id string)(*TenantRelationshipsManagedTenantsCloudPcsOverviewCloudPcOverviewTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPcOverview%2DtenantId"] = id
    }
    return NewTenantRelationshipsManagedTenantsCloudPcsOverviewCloudPcOverviewTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessPolicyCoverages provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ConditionalAccessPolicyCoverages()(*TenantRelationshipsManagedTenantsConditionalAccessPolicyCoveragesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsConditionalAccessPolicyCoveragesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConditionalAccessPolicyCoveragesById provides operations to manage the conditionalAccessPolicyCoverages property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ConditionalAccessPolicyCoveragesById(id string)(*TenantRelationshipsManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicyCoverage%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsConditionalAccessPolicyCoveragesConditionalAccessPolicyCoverageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTenantRelationshipsManagedTenantsRequestBuilderInternal instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsRequestBuilder) {
    m := &TenantRelationshipsManagedTenantsRequestBuilder{
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
// NewTenantRelationshipsManagedTenantsRequestBuilder instantiates a new ManagedTenantsRequestBuilder and sets the default values.
func NewTenantRelationshipsManagedTenantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsManagedTenantsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsManagedTenantsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property managedTenants for tenantRelationships
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the operations available to interact with the multi-tenant management platform.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property managedTenants in tenantRelationships
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CredentialUserRegistrationsSummaries provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CredentialUserRegistrationsSummaries()(*TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationsSummariesById provides operations to manage the credentialUserRegistrationsSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) CredentialUserRegistrationsSummariesById(id string)(*TenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesCredentialUserRegistrationsSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationsSummary%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsCredentialUserRegistrationsSummariesCredentialUserRegistrationsSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property managedTenants for tenantRelationships
func (m *TenantRelationshipsManagedTenantsRequestBuilder) Delete(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*TenantRelationshipsManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*TenantRelationshipsManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the operations available to interact with the multi-tenant management platform.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
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
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedDeviceCompliances()(*TenantRelationshipsManagedTenantsManagedDeviceCompliancesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedDeviceCompliancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceCompliancesById provides operations to manage the managedDeviceCompliances property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedDeviceCompliancesById(id string)(*TenantRelationshipsManagedTenantsManagedDeviceCompliancesManagedDeviceComplianceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceCompliance%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedDeviceCompliancesManagedDeviceComplianceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceComplianceTrends provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedDeviceComplianceTrends()(*TenantRelationshipsManagedTenantsManagedDeviceComplianceTrendsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedDeviceComplianceTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceComplianceTrendsById provides operations to manage the managedDeviceComplianceTrends property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedDeviceComplianceTrendsById(id string)(*TenantRelationshipsManagedTenantsManagedDeviceComplianceTrendsManagedDeviceComplianceTrendItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceComplianceTrend%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedDeviceComplianceTrendsManagedDeviceComplianceTrendItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertLogs provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertLogs()(*TenantRelationshipsManagedTenantsManagedTenantAlertLogsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertLogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertLogsById provides operations to manage the managedTenantAlertLogs property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertLogsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertLogsManagedTenantAlertLogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertLog%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertLogsManagedTenantAlertLogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitions provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitions()(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRuleDefinitionsById provides operations to manage the managedTenantAlertRuleDefinitions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertRuleDefinitionsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRuleDefinition%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRuleDefinitionsManagedTenantAlertRuleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlertRules provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertRules()(*TenantRelationshipsManagedTenantsManagedTenantAlertRulesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertRulesById provides operations to manage the managedTenantAlertRules property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertRulesById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlertRule%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertRulesManagedTenantAlertRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantAlerts provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlerts()(*TenantRelationshipsManagedTenantsManagedTenantAlertsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantAlertsById provides operations to manage the managedTenantAlerts property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantAlertsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantAlert%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantAlertsManagedTenantAlertItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantApiNotifications provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantApiNotifications()(*TenantRelationshipsManagedTenantsManagedTenantApiNotificationsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantApiNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantApiNotificationsById provides operations to manage the managedTenantApiNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantApiNotificationsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantApiNotificationsManagedTenantApiNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantApiNotification%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantApiNotificationsManagedTenantApiNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantEmailNotifications provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantEmailNotifications()(*TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantEmailNotificationsById provides operations to manage the managedTenantEmailNotifications property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantEmailNotificationsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantEmailNotification%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantEmailNotificationsManagedTenantEmailNotificationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedTenantTicketingEndpoints provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantTicketingEndpoints()(*TenantRelationshipsManagedTenantsManagedTenantTicketingEndpointsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagedTenantTicketingEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedTenantTicketingEndpointsById provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagedTenantTicketingEndpointsById(id string)(*TenantRelationshipsManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedTenantTicketingEndpoint%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActions provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementActions()(*TenantRelationshipsManagedTenantsManagementActionsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionsById provides operations to manage the managementActions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementActionsById(id string)(*TenantRelationshipsManagedTenantsManagementActionsManagementActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementAction%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementActionsManagementActionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatuses provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatuses()(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementActionTenantDeploymentStatusesById provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementActionTenantDeploymentStatusesById(id string)(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementActionTenantDeploymentStatus%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesManagementActionTenantDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementIntents provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementIntents()(*TenantRelationshipsManagedTenantsManagementIntentsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementIntentsById provides operations to manage the managementIntents property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementIntentsById(id string)(*TenantRelationshipsManagedTenantsManagementIntentsManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementIntent%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementIntentsManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollections provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateCollections()(*TenantRelationshipsManagedTenantsManagementTemplateCollectionsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionsById provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateCollectionsById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollection%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummaries provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummaries()(*TenantRelationshipsManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateCollectionTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateCollectionTenantSummariesById provides operations to manage the managementTemplateCollectionTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateCollectionTenantSummariesById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateCollectionTenantSummariesManagementTemplateCollectionTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateCollectionTenantSummary%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateCollectionTenantSummariesManagementTemplateCollectionTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplates()(*TenantRelationshipsManagedTenantsManagementTemplatesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplatesById provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplatesById(id string)(*TenantRelationshipsManagedTenantsManagementTemplatesManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplate%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplatesManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateSteps provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateSteps()(*TenantRelationshipsManagedTenantsManagementTemplateStepsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepsById provides operations to manage the managementTemplateSteps property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateStepsById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateStepsManagementTemplateStepItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStep%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepsManagementTemplateStepItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepTenantSummaries provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummaries()(*TenantRelationshipsManagedTenantsManagementTemplateStepTenantSummariesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepTenantSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepTenantSummariesById provides operations to manage the managementTemplateStepTenantSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateStepTenantSummariesById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateStepTenantSummariesManagementTemplateStepTenantSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepTenantSummary%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepTenantSummariesManagementTemplateStepTenantSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagementTemplateStepVersions provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateStepVersions()(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagementTemplateStepVersionsById provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) ManagementTemplateStepVersionsById(id string)(*TenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepVersion%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MyRoles provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) MyRoles()(*TenantRelationshipsManagedTenantsMyRolesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsMyRolesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MyRolesById provides operations to manage the myRoles property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) MyRolesById(id string)(*TenantRelationshipsManagedTenantsMyRolesMyRoleTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["myRole%2DtenantId"] = id
    }
    return NewTenantRelationshipsManagedTenantsMyRolesMyRoleTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property managedTenants in tenantRelationships
func (m *TenantRelationshipsManagedTenantsRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, requestConfiguration *TenantRelationshipsManagedTenantsRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
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
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantGroups()(*TenantRelationshipsManagedTenantsTenantGroupsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsTenantGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantGroupsById provides operations to manage the tenantGroups property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantGroupsById(id string)(*TenantRelationshipsManagedTenantsTenantGroupsTenantGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantGroup%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsTenantGroupsTenantGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) Tenants()(*TenantRelationshipsManagedTenantsTenantsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsById provides operations to manage the tenants property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantsById(id string)(*TenantRelationshipsManagedTenantsTenantsTenantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenant%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsTenantsTenantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsCustomizedInformation provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantsCustomizedInformation()(*TenantRelationshipsManagedTenantsTenantsCustomizedInformationRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsCustomizedInformationById provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantsCustomizedInformationById(id string)(*TenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantCustomizedInformation%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantsDetailedInformation provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantsDetailedInformation()(*TenantRelationshipsManagedTenantsTenantsDetailedInformationRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsTenantsDetailedInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantsDetailedInformationById provides operations to manage the tenantsDetailedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantsDetailedInformationById(id string)(*TenantRelationshipsManagedTenantsTenantsDetailedInformationTenantDetailedInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantDetailedInformation%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsTenantsDetailedInformationTenantDetailedInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantTags provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantTags()(*TenantRelationshipsManagedTenantsTenantTagsRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsTenantTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TenantTagsById provides operations to manage the tenantTags property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) TenantTagsById(id string)(*TenantRelationshipsManagedTenantsTenantTagsTenantTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["tenantTag%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsTenantTagsTenantTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsDeviceMalwareStates provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) WindowsDeviceMalwareStates()(*TenantRelationshipsManagedTenantsWindowsDeviceMalwareStatesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsWindowsDeviceMalwareStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDeviceMalwareStatesById provides operations to manage the windowsDeviceMalwareStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) WindowsDeviceMalwareStatesById(id string)(*TenantRelationshipsManagedTenantsWindowsDeviceMalwareStatesWindowsDeviceMalwareStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDeviceMalwareState%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsWindowsDeviceMalwareStatesWindowsDeviceMalwareStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsProtectionStates provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) WindowsProtectionStates()(*TenantRelationshipsManagedTenantsWindowsProtectionStatesRequestBuilder) {
    return NewTenantRelationshipsManagedTenantsWindowsProtectionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsProtectionStatesById provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *TenantRelationshipsManagedTenantsRequestBuilder) WindowsProtectionStatesById(id string)(*TenantRelationshipsManagedTenantsWindowsProtectionStatesWindowsProtectionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsProtectionState%2Did"] = id
    }
    return NewTenantRelationshipsManagedTenantsWindowsProtectionStatesWindowsProtectionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
