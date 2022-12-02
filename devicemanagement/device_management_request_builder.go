package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementRequestBuilder provides operations to manage the deviceManagement singleton.
type DeviceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementRequestBuilderGetQueryParameters get deviceManagement
type DeviceManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRequestBuilderGetQueryParameters
}
// DeviceManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedThreatProtectionOnboardingStateSummary provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AdvancedThreatProtectionOnboardingStateSummary()(*DeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    return NewDeviceManagementAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidDeviceOwnerEnrollmentProfiles provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfiles()(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidDeviceOwnerEnrollmentProfilesById provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfilesById(id string)(*DeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidDeviceOwnerEnrollmentProfile%2Did"] = id
    }
    return NewDeviceManagementAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkAppConfigurationSchemas provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemas()(*DeviceManagementAndroidForWorkAppConfigurationSchemasRequestBuilder) {
    return NewDeviceManagementAndroidForWorkAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidForWorkAppConfigurationSchemasById provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemasById(id string)(*DeviceManagementAndroidForWorkAppConfigurationSchemasAndroidForWorkAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidForWorkAppConfigurationSchema%2Did"] = id
    }
    return NewDeviceManagementAndroidForWorkAppConfigurationSchemasAndroidForWorkAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkEnrollmentProfiles provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfiles()(*DeviceManagementAndroidForWorkEnrollmentProfilesRequestBuilder) {
    return NewDeviceManagementAndroidForWorkEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidForWorkEnrollmentProfilesById provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfilesById(id string)(*DeviceManagementAndroidForWorkEnrollmentProfilesAndroidForWorkEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidForWorkEnrollmentProfile%2Did"] = id
    }
    return NewDeviceManagementAndroidForWorkEnrollmentProfilesAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkSettings provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkSettings()(*DeviceManagementAndroidForWorkSettingsRequestBuilder) {
    return NewDeviceManagementAndroidForWorkSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAccountEnterpriseSettings provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAccountEnterpriseSettings()(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAppConfigurationSchemas provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemas()(*DeviceManagementAndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    return NewDeviceManagementAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAppConfigurationSchemasById provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemasById(id string)(*DeviceManagementAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedStoreAppConfigurationSchema%2Did"] = id
    }
    return NewDeviceManagementAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*DeviceManagementApplePushNotificationCertificateRequestBuilder) {
    return NewDeviceManagementApplePushNotificationCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppleUserInitiatedEnrollmentProfiles provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfiles()(*DeviceManagementAppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppleUserInitiatedEnrollmentProfilesById provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfilesById(id string)(*DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleUserInitiatedEnrollmentProfile%2Did"] = id
    }
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentFilters provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFilters()(*DeviceManagementAssignmentFiltersRequestBuilder) {
    return NewDeviceManagementAssignmentFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFiltersById provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFiltersById(id string)(*DeviceManagementAssignmentFiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAndAppManagementAssignmentFilter%2Did"] = id
    }
    return NewDeviceManagementAssignmentFiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEvents()(*DeviceManagementAuditEventsRequestBuilder) {
    return NewDeviceManagementAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEventsById(id string)(*DeviceManagementAuditEventsAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent%2Did"] = id
    }
    return NewDeviceManagementAuditEventsAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AutopilotEvents provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEvents()(*DeviceManagementAutopilotEventsRequestBuilder) {
    return NewDeviceManagementAutopilotEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AutopilotEventsById provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEventsById(id string)(*DeviceManagementAutopilotEventsDeviceManagementAutopilotEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementAutopilotEvent%2Did"] = id
    }
    return NewDeviceManagementAutopilotEventsDeviceManagementAutopilotEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CartToClassAssociations provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociations()(*DeviceManagementCartToClassAssociationsRequestBuilder) {
    return NewDeviceManagementCartToClassAssociationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CartToClassAssociationsById provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociationsById(id string)(*DeviceManagementCartToClassAssociationsCartToClassAssociationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cartToClassAssociation%2Did"] = id
    }
    return NewDeviceManagementCartToClassAssociationsCartToClassAssociationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Categories()(*DeviceManagementCategoriesRequestBuilder) {
    return NewDeviceManagementCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CategoriesById(id string)(*DeviceManagementCategoriesDeviceManagementSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingCategory%2Did"] = id
    }
    return NewDeviceManagementCategoriesDeviceManagementSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CertificateConnectorDetails provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetails()(*DeviceManagementCertificateConnectorDetailsRequestBuilder) {
    return NewDeviceManagementCertificateConnectorDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateConnectorDetailsById provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetailsById(id string)(*DeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateConnectorDetails%2Did"] = id
    }
    return NewDeviceManagementCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChromeOSOnboardingSettings provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettings()(*DeviceManagementChromeOSOnboardingSettingsRequestBuilder) {
    return NewDeviceManagementChromeOSOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChromeOSOnboardingSettingsById provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettingsById(id string)(*DeviceManagementChromeOSOnboardingSettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chromeOSOnboardingSettings%2Did"] = id
    }
    return NewDeviceManagementChromeOSOnboardingSettingsChromeOSOnboardingSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPCConnectivityIssues provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssues()(*DeviceManagementCloudPCConnectivityIssuesRequestBuilder) {
    return NewDeviceManagementCloudPCConnectivityIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCConnectivityIssuesById provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssuesById(id string)(*DeviceManagementCloudPCConnectivityIssuesCloudPCConnectivityIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPCConnectivityIssue%2Did"] = id
    }
    return NewDeviceManagementCloudPCConnectivityIssuesCloudPCConnectivityIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComanagedDevices provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevices()(*DeviceManagementComanagedDevicesRequestBuilder) {
    return NewDeviceManagementComanagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComanagedDevicesById provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevicesById(id string)(*DeviceManagementComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewDeviceManagementComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComanagementEligibleDevices provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevices()(*DeviceManagementComanagementEligibleDevicesRequestBuilder) {
    return NewDeviceManagementComanagementEligibleDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComanagementEligibleDevicesById provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevicesById(id string)(*DeviceManagementComanagementEligibleDevicesComanagementEligibleDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["comanagementEligibleDevice%2Did"] = id
    }
    return NewDeviceManagementComanagementEligibleDevicesComanagementEligibleDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceCategories provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategories()(*DeviceManagementComplianceCategoriesRequestBuilder) {
    return NewDeviceManagementComplianceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceCategoriesById provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategoriesById(id string)(*DeviceManagementComplianceCategoriesDeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationCategory%2Did"] = id
    }
    return NewDeviceManagementComplianceCategoriesDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*DeviceManagementComplianceManagementPartnersRequestBuilder) {
    return NewDeviceManagementComplianceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceManagementPartnersById provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartnersById(id string)(*DeviceManagementComplianceManagementPartnersComplianceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["complianceManagementPartner%2Did"] = id
    }
    return NewDeviceManagementComplianceManagementPartnersComplianceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompliancePolicies provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePolicies()(*DeviceManagementCompliancePoliciesRequestBuilder) {
    return NewDeviceManagementCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompliancePoliciesById provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePoliciesById(id string)(*DeviceManagementCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementCompliancePolicy%2Did"] = id
    }
    return NewDeviceManagementCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceSettings provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettings()(*DeviceManagementComplianceSettingsRequestBuilder) {
    return NewDeviceManagementComplianceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceSettingsById provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettingsById(id string)(*DeviceManagementComplianceSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementComplianceSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*DeviceManagementConditionalAccessSettingsRequestBuilder) {
    return NewDeviceManagementConditionalAccessSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigManagerCollections provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollections()(*DeviceManagementConfigManagerCollectionsRequestBuilder) {
    return NewDeviceManagementConfigManagerCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigManagerCollectionsById provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollectionsById(id string)(*DeviceManagementConfigManagerCollectionsConfigManagerCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["configManagerCollection%2Did"] = id
    }
    return NewDeviceManagementConfigManagerCollectionsConfigManagerCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationCategories provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategories()(*DeviceManagementConfigurationCategoriesRequestBuilder) {
    return NewDeviceManagementConfigurationCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationCategoriesById provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategoriesById(id string)(*DeviceManagementConfigurationCategoriesDeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationCategory%2Did"] = id
    }
    return NewDeviceManagementConfigurationCategoriesDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationPolicies provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicies()(*DeviceManagementConfigurationPoliciesRequestBuilder) {
    return NewDeviceManagementConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationPoliciesById provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPoliciesById(id string)(*DeviceManagementConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicy%2Did"] = id
    }
    return NewDeviceManagementConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationPolicyTemplates provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplates()(*DeviceManagementConfigurationPolicyTemplatesRequestBuilder) {
    return NewDeviceManagementConfigurationPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationPolicyTemplatesById provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplatesById(id string)(*DeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyTemplate%2Did"] = id
    }
    return NewDeviceManagementConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationSettings provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettings()(*DeviceManagementConfigurationSettingsRequestBuilder) {
    return NewDeviceManagementConfigurationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationSettingsById provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettingsById(id string)(*DeviceManagementConfigurationSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementConfigurationSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get deviceManagement
func (m *DeviceManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update deviceManagement
func (m *DeviceManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataSharingConsents provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsents()(*DeviceManagementDataSharingConsentsRequestBuilder) {
    return NewDeviceManagementDataSharingConsentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataSharingConsentsById provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsentsById(id string)(*DeviceManagementDataSharingConsentsDataSharingConsentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSharingConsent%2Did"] = id
    }
    return NewDeviceManagementDataSharingConsentsDataSharingConsentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DepOnboardingSettings provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettings()(*DeviceManagementDepOnboardingSettingsRequestBuilder) {
    return NewDeviceManagementDepOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DepOnboardingSettingsById provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettingsById(id string)(*DeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["depOnboardingSetting%2Did"] = id
    }
    return NewDeviceManagementDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DerivedCredentials provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentials()(*DeviceManagementDerivedCredentialsRequestBuilder) {
    return NewDeviceManagementDerivedCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DerivedCredentialsById provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentialsById(id string)(*DeviceManagementDerivedCredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementDerivedCredentialSettings%2Did"] = id
    }
    return NewDeviceManagementDerivedCredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DeviceManagementDetectedAppsRequestBuilder) {
    return NewDeviceManagementDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedAppsById(id string)(*DeviceManagementDetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewDeviceManagementDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DeviceManagementDeviceCategoriesRequestBuilder) {
    return NewDeviceManagementDeviceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCategoriesById provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategoriesById(id string)(*DeviceManagementDeviceCategoriesDeviceCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCategory%2Did"] = id
    }
    return NewDeviceManagementDeviceCategoriesDeviceCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DeviceManagementDeviceCompliancePoliciesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePoliciesById provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePoliciesById(id string)(*DeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicy%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DeviceManagementDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DeviceManagementDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDeviceManagementDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*DeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = id
    }
    return NewDeviceManagementDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceComplianceScripts provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScripts()(*DeviceManagementDeviceComplianceScriptsRequestBuilder) {
    return NewDeviceManagementDeviceComplianceScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceComplianceScriptsById provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScriptsById(id string)(*DeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScript%2Did"] = id
    }
    return NewDeviceManagementDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationConflictSummary provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummary()(*DeviceManagementDeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationConflictSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationConflictSummaryById provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummaryById(id string)(*DeviceManagementDeviceConfigurationConflictSummaryDeviceConfigurationConflictSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationConflictSummary%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationConflictSummaryDeviceConfigurationConflictSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceManagementDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationRestrictedAppsViolations provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolations()(*DeviceManagementDeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationRestrictedAppsViolationsById provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolationsById(id string)(*DeviceManagementDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["restrictedAppsViolation%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceManagementDeviceConfigurationsRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationsAllManagedDeviceCertificateStates provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStates()(*DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesById provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStatesById(id string)(*DeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAllDeviceCertificateState%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationsById provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsById(id string)(*DeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfiguration%2Did"] = id
    }
    return NewDeviceManagementDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationUserStateSummaries provides operations to manage the deviceConfigurationUserStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationUserStateSummaries()(*DeviceManagementDeviceConfigurationUserStateSummariesRequestBuilder) {
    return NewDeviceManagementDeviceConfigurationUserStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCustomAttributeShellScripts provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScripts()(*DeviceManagementDeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDeviceManagementDeviceCustomAttributeShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCustomAttributeShellScriptsById provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScriptsById(id string)(*DeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCustomAttributeShellScript%2Did"] = id
    }
    return NewDeviceManagementDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceManagementDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceManagementDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*DeviceManagementDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return NewDeviceManagementDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceHealthScripts provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScripts()(*DeviceManagementDeviceHealthScriptsRequestBuilder) {
    return NewDeviceManagementDeviceHealthScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceHealthScriptsById provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScriptsById(id string)(*DeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScript%2Did"] = id
    }
    return NewDeviceManagementDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DeviceManagementDeviceManagementPartnersRequestBuilder) {
    return NewDeviceManagementDeviceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementPartnersById provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartnersById(id string)(*DeviceManagementDeviceManagementPartnersDeviceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementPartner%2Did"] = id
    }
    return NewDeviceManagementDeviceManagementPartnersDeviceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementScripts provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScripts()(*DeviceManagementDeviceManagementScriptsRequestBuilder) {
    return NewDeviceManagementDeviceManagementScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementScriptsById provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScriptsById(id string)(*DeviceManagementDeviceManagementScriptsDeviceManagementScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScript%2Did"] = id
    }
    return NewDeviceManagementDeviceManagementScriptsDeviceManagementScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceShellScripts provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScripts()(*DeviceManagementDeviceShellScriptsRequestBuilder) {
    return NewDeviceManagementDeviceShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceShellScriptsById provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScriptsById(id string)(*DeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceShellScript%2Did"] = id
    }
    return NewDeviceManagementDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DomainJoinConnectors provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectors()(*DeviceManagementDomainJoinConnectorsRequestBuilder) {
    return NewDeviceManagementDomainJoinConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainJoinConnectorsById provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectorsById(id string)(*DeviceManagementDomainJoinConnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementDomainJoinConnector%2Did"] = id
    }
    return NewDeviceManagementDomainJoinConnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EmbeddedSIMActivationCodePools provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePools()(*DeviceManagementEmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmbeddedSIMActivationCodePoolsById provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePoolsById(id string)(*DeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePool%2Did"] = id
    }
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EnableAndroidDeviceAdministratorEnrollment provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
func (m *DeviceManagementRequestBuilder) EnableAndroidDeviceAdministratorEnrollment()(*DeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    return NewDeviceManagementEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLegacyPcManagement provides operations to call the enableLegacyPcManagement method.
func (m *DeviceManagementRequestBuilder) EnableLegacyPcManagement()(*DeviceManagementEnableLegacyPcManagementRequestBuilder) {
    return NewDeviceManagementEnableLegacyPcManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableUnlicensedAdminstrators provides operations to call the enableUnlicensedAdminstrators method.
func (m *DeviceManagementRequestBuilder) EnableUnlicensedAdminstrators()(*DeviceManagementEnableUnlicensedAdminstratorsRequestBuilder) {
    return NewDeviceManagementEnableUnlicensedAdminstratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateAssignmentFilter provides operations to call the evaluateAssignmentFilter method.
func (m *DeviceManagementRequestBuilder) EvaluateAssignmentFilter()(*DeviceManagementEvaluateAssignmentFilterRequestBuilder) {
    return NewDeviceManagementEvaluateAssignmentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*DeviceManagementExchangeConnectorsRequestBuilder) {
    return NewDeviceManagementExchangeConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeConnectorsById provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectorsById(id string)(*DeviceManagementExchangeConnectorsDeviceManagementExchangeConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExchangeConnector%2Did"] = id
    }
    return NewDeviceManagementExchangeConnectorsDeviceManagementExchangeConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExchangeOnPremisesPolicies provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicies()(*DeviceManagementExchangeOnPremisesPoliciesRequestBuilder) {
    return NewDeviceManagementExchangeOnPremisesPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeOnPremisesPoliciesById provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPoliciesById(id string)(*DeviceManagementExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExchangeOnPremisesPolicy%2Did"] = id
    }
    return NewDeviceManagementExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExchangeOnPremisesPolicy provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicyRequestBuilder) {
    return NewDeviceManagementExchangeOnPremisesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get deviceManagement
func (m *DeviceManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable), nil
}
// GetAssignedRoleDetails provides operations to call the getAssignedRoleDetails method.
func (m *DeviceManagementRequestBuilder) GetAssignedRoleDetails()(*DeviceManagementGetAssignedRoleDetailsRequestBuilder) {
    return NewDeviceManagementGetAssignedRoleDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAssignmentFiltersStatusDetails provides operations to call the getAssignmentFiltersStatusDetails method.
func (m *DeviceManagementRequestBuilder) GetAssignmentFiltersStatusDetails()(*DeviceManagementGetAssignmentFiltersStatusDetailsRequestBuilder) {
    return NewDeviceManagementGetAssignmentFiltersStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComanagedDevicesSummary provides operations to call the getComanagedDevicesSummary method.
func (m *DeviceManagementRequestBuilder) GetComanagedDevicesSummary()(*DeviceManagementGetComanagedDevicesSummaryRequestBuilder) {
    return NewDeviceManagementGetComanagedDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComanagementEligibleDevicesSummary provides operations to call the getComanagementEligibleDevicesSummary method.
func (m *DeviceManagementRequestBuilder) GetComanagementEligibleDevicesSummary()(*DeviceManagementGetComanagementEligibleDevicesSummaryRequestBuilder) {
    return NewDeviceManagementGetComanagementEligibleDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) GetEffectivePermissions()(*DeviceManagementGetEffectivePermissionsRequestBuilder) {
    return NewDeviceManagementGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*DeviceManagementGetEffectivePermissionsWithScopeRequestBuilder) {
    return NewDeviceManagementGetEffectivePermissionsWithScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter, scope);
}
// GetRoleScopeTagsByIdsWithIds provides operations to call the getRoleScopeTagsByIds method.
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByIdsWithIds(ids *string)(*DeviceManagementGetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewDeviceManagementGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter, ids);
}
// GetRoleScopeTagsByResourceWithResource provides operations to call the getRoleScopeTagsByResource method.
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByResourceWithResource(resource *string)(*DeviceManagementGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewDeviceManagementGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource);
}
// GetSuggestedEnrollmentLimitWithEnrollmentType provides operations to call the getSuggestedEnrollmentLimit method.
func (m *DeviceManagementRequestBuilder) GetSuggestedEnrollmentLimitWithEnrollmentType(enrollmentType *string)(*DeviceManagementGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    return NewDeviceManagementGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter, enrollmentType);
}
// GroupPolicyCategories provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategories()(*DeviceManagementGroupPolicyCategoriesRequestBuilder) {
    return NewDeviceManagementGroupPolicyCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyCategoriesById provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategoriesById(id string)(*DeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyCategory%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyConfigurations provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurations()(*DeviceManagementGroupPolicyConfigurationsRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyConfigurationsById provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurationsById(id string)(*DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfiguration%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyDefinitionFiles provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFiles()(*DeviceManagementGroupPolicyDefinitionFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyDefinitionFilesById provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFilesById(id string)(*DeviceManagementGroupPolicyDefinitionFilesGroupPolicyDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionFile%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyDefinitionFilesGroupPolicyDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyDefinitions provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitions()(*DeviceManagementGroupPolicyDefinitionsRequestBuilder) {
    return NewDeviceManagementGroupPolicyDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyDefinitionsById provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionsById(id string)(*DeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyMigrationReports provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReports()(*DeviceManagementGroupPolicyMigrationReportsRequestBuilder) {
    return NewDeviceManagementGroupPolicyMigrationReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyMigrationReportsById provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReportsById(id string)(*DeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyMigrationReport%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyObjectFiles provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFiles()(*DeviceManagementGroupPolicyObjectFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyObjectFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyObjectFilesById provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFilesById(id string)(*DeviceManagementGroupPolicyObjectFilesGroupPolicyObjectFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyObjectFile%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyObjectFilesGroupPolicyObjectFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyUploadedDefinitionFiles provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFiles()(*DeviceManagementGroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyUploadedDefinitionFilesById provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFilesById(id string)(*DeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyUploadedDefinitionFile%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ImportedDeviceIdentities provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentities()(*DeviceManagementImportedDeviceIdentitiesRequestBuilder) {
    return NewDeviceManagementImportedDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedDeviceIdentitiesById provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentitiesById(id string)(*DeviceManagementImportedDeviceIdentitiesImportedDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedDeviceIdentity%2Did"] = id
    }
    return NewDeviceManagementImportedDeviceIdentitiesImportedDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*DeviceManagementImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedWindowsAutopilotDeviceIdentitiesById provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentitiesById(id string)(*DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedWindowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Intents provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Intents()(*DeviceManagementIntentsRequestBuilder) {
    return NewDeviceManagementIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntentsById provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntentsById(id string)(*DeviceManagementIntentsDeviceManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntent%2Did"] = id
    }
    return NewDeviceManagementIntentsDeviceManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IntuneBrandingProfiles provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfiles()(*DeviceManagementIntuneBrandingProfilesRequestBuilder) {
    return NewDeviceManagementIntuneBrandingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntuneBrandingProfilesById provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfilesById(id string)(*DeviceManagementIntuneBrandingProfilesIntuneBrandingProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["intuneBrandingProfile%2Did"] = id
    }
    return NewDeviceManagementIntuneBrandingProfilesIntuneBrandingProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*DeviceManagementIosUpdateStatusesRequestBuilder) {
    return NewDeviceManagementIosUpdateStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosUpdateStatusesById provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatusesById(id string)(*DeviceManagementIosUpdateStatusesIosUpdateDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosUpdateDeviceStatus%2Did"] = id
    }
    return NewDeviceManagementIosUpdateStatusesIosUpdateDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MacOSSoftwareUpdateAccountSummaries provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummaries()(*DeviceManagementMacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MacOSSoftwareUpdateAccountSummariesById provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummariesById(id string)(*DeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateAccountSummary%2Did"] = id
    }
    return NewDeviceManagementMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceEncryptionStates provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStates()(*DeviceManagementManagedDeviceEncryptionStatesRequestBuilder) {
    return NewDeviceManagementManagedDeviceEncryptionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEncryptionStatesById provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStatesById(id string)(*DeviceManagementManagedDeviceEncryptionStatesManagedDeviceEncryptionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceEncryptionState%2Did"] = id
    }
    return NewDeviceManagementManagedDeviceEncryptionStatesManagedDeviceEncryptionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*DeviceManagementManagedDeviceOverviewRequestBuilder) {
    return NewDeviceManagementManagedDeviceOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*DeviceManagementManagedDevicesRequestBuilder) {
    return NewDeviceManagementManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevicesById(id string)(*DeviceManagementManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewDeviceManagementManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelConfigurations provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurations()(*DeviceManagementMicrosoftTunnelConfigurationsRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelConfigurationsById provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurationsById(id string)(*DeviceManagementMicrosoftTunnelConfigurationsMicrosoftTunnelConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelConfiguration%2Did"] = id
    }
    return NewDeviceManagementMicrosoftTunnelConfigurationsMicrosoftTunnelConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelHealthThresholds provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholds()(*DeviceManagementMicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelHealthThresholdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelHealthThresholdsById provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholdsById(id string)(*DeviceManagementMicrosoftTunnelHealthThresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelHealthThreshold%2Did"] = id
    }
    return NewDeviceManagementMicrosoftTunnelHealthThresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelServerLogCollectionResponses provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponses()(*DeviceManagementMicrosoftTunnelServerLogCollectionResponsesRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServerLogCollectionResponsesById provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponsesById(id string)(*DeviceManagementMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelServerLogCollectionResponse%2Did"] = id
    }
    return NewDeviceManagementMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelSites provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSites()(*DeviceManagementMicrosoftTunnelSitesRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelSitesById provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSitesById(id string)(*DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelSite%2Did"] = id
    }
    return NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*DeviceManagementMobileAppTroubleshootingEventsRequestBuilder) {
    return NewDeviceManagementMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*DeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return NewDeviceManagementMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*DeviceManagementMobileThreatDefenseConnectorsRequestBuilder) {
    return NewDeviceManagementMobileThreatDefenseConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileThreatDefenseConnectorsById provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectorsById(id string)(*DeviceManagementMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileThreatDefenseConnector%2Did"] = id
    }
    return NewDeviceManagementMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Monitoring provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Monitoring()(*DeviceManagementMonitoringRequestBuilder) {
    return NewDeviceManagementMonitoringRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NdesConnectors provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectors()(*DeviceManagementNdesConnectorsRequestBuilder) {
    return NewDeviceManagementNdesConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NdesConnectorsById provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectorsById(id string)(*DeviceManagementNdesConnectorsNdesConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ndesConnector%2Did"] = id
    }
    return NewDeviceManagementNdesConnectorsNdesConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*DeviceManagementNotificationMessageTemplatesRequestBuilder) {
    return NewDeviceManagementNotificationMessageTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationMessageTemplatesById provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplatesById(id string)(*DeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notificationMessageTemplate%2Did"] = id
    }
    return NewDeviceManagementNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OemWarrantyInformationOnboarding provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboarding()(*DeviceManagementOemWarrantyInformationOnboardingRequestBuilder) {
    return NewDeviceManagementOemWarrantyInformationOnboardingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OemWarrantyInformationOnboardingById provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboardingById(id string)(*DeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oemWarrantyInformationOnboarding%2Did"] = id
    }
    return NewDeviceManagementOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update deviceManagement
func (m *DeviceManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable), nil
}
// RemoteActionAudits provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteActionAudits()(*DeviceManagementRemoteActionAuditsRequestBuilder) {
    return NewDeviceManagementRemoteActionAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteActionAuditsById provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteActionAuditsById(id string)(*DeviceManagementRemoteActionAuditsRemoteActionAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["remoteActionAudit%2Did"] = id
    }
    return NewDeviceManagementRemoteActionAuditsRemoteActionAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*DeviceManagementRemoteAssistancePartnersRequestBuilder) {
    return NewDeviceManagementRemoteAssistancePartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteAssistancePartnersById provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartnersById(id string)(*DeviceManagementRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["remoteAssistancePartner%2Did"] = id
    }
    return NewDeviceManagementRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoteAssistanceSettings provides operations to manage the remoteAssistanceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistanceSettings()(*DeviceManagementRemoteAssistanceSettingsRequestBuilder) {
    return NewDeviceManagementRemoteAssistanceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Reports()(*DeviceManagementReportsRequestBuilder) {
    return NewDeviceManagementReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceAccessProfiles provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfiles()(*DeviceManagementResourceAccessProfilesRequestBuilder) {
    return NewDeviceManagementResourceAccessProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceAccessProfilesById provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfilesById(id string)(*DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementResourceAccessProfileBase%2Did"] = id
    }
    return NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*DeviceManagementResourceOperationsRequestBuilder) {
    return NewDeviceManagementResourceOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceOperationsById provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperationsById(id string)(*DeviceManagementResourceOperationsResourceOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceOperation%2Did"] = id
    }
    return NewDeviceManagementResourceOperationsResourceOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReusablePolicySettings provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettings()(*DeviceManagementReusablePolicySettingsRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReusablePolicySettingsById provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettingsById(id string)(*DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementReusablePolicySetting%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReusableSettings provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettings()(*DeviceManagementReusableSettingsRequestBuilder) {
    return NewDeviceManagementReusableSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReusableSettingsById provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettingsById(id string)(*DeviceManagementReusableSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementReusableSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*DeviceManagementRoleAssignmentsRequestBuilder) {
    return NewDeviceManagementRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*DeviceManagementRoleAssignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAndAppManagementRoleAssignment%2Did"] = id
    }
    return NewDeviceManagementRoleAssignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*DeviceManagementRoleDefinitionsRequestBuilder) {
    return NewDeviceManagementRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*DeviceManagementRoleDefinitionsRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleDefinition%2Did"] = id
    }
    return NewDeviceManagementRoleDefinitionsRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTags()(*DeviceManagementRoleScopeTagsRequestBuilder) {
    return NewDeviceManagementRoleScopeTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleScopeTagsById provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTagsById(id string)(*DeviceManagementRoleScopeTagsRoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleScopeTag%2Did"] = id
    }
    return NewDeviceManagementRoleScopeTagsRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ScopedForResourceWithResource provides operations to call the scopedForResource method.
func (m *DeviceManagementRequestBuilder) ScopedForResourceWithResource(resource *string)(*DeviceManagementScopedForResourceWithResourceRequestBuilder) {
    return NewDeviceManagementScopedForResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *DeviceManagementRequestBuilder) SendCustomNotificationToCompanyPortal()(*DeviceManagementSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewDeviceManagementSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitions()(*DeviceManagementSettingDefinitionsRequestBuilder) {
    return NewDeviceManagementSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitionsById provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitionsById(id string)(*DeviceManagementSettingDefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementSettingDefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*DeviceManagementSoftwareUpdateStatusSummaryRequestBuilder) {
    return NewDeviceManagementSoftwareUpdateStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*DeviceManagementTelecomExpenseManagementPartnersRequestBuilder) {
    return NewDeviceManagementTelecomExpenseManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TelecomExpenseManagementPartnersById provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartnersById(id string)(*DeviceManagementTelecomExpenseManagementPartnersTelecomExpenseManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["telecomExpenseManagementPartner%2Did"] = id
    }
    return NewDeviceManagementTelecomExpenseManagementPartnersTelecomExpenseManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Templates provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Templates()(*DeviceManagementTemplatesRequestBuilder) {
    return NewDeviceManagementTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplatesById(id string)(*DeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate%2Did"] = id
    }
    return NewDeviceManagementTemplatesDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemplateSettings provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettings()(*DeviceManagementTemplateSettingsRequestBuilder) {
    return NewDeviceManagementTemplateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateSettingsById provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettingsById(id string)(*DeviceManagementTemplateSettingsDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate%2Did"] = id
    }
    return NewDeviceManagementTemplateSettingsDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantAttachRBAC provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TenantAttachRBAC()(*DeviceManagementTenantAttachRBACRequestBuilder) {
    return NewDeviceManagementTenantAttachRBACRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*DeviceManagementTermsAndConditionsRequestBuilder) {
    return NewDeviceManagementTermsAndConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsAndConditionsById provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditionsById(id string)(*DeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditions%2Did"] = id
    }
    return NewDeviceManagementTermsAndConditionsTermsAndConditionsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TroubleshootingEvents provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*DeviceManagementTroubleshootingEventsRequestBuilder) {
    return NewDeviceManagementTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TroubleshootingEventsById provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEventsById(id string)(*DeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewDeviceManagementTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAnomaly provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomaly()(*DeviceManagementUserExperienceAnalyticsAnomalyRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAnomalyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyById provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyById(id string)(*DeviceManagementUserExperienceAnalyticsAnomalyUserExperienceAnalyticsAnomalyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAnomaly%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAnomalyUserExperienceAnalyticsAnomalyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyDevice provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDevice()(*DeviceManagementUserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyDeviceById provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDeviceById(id string)(*DeviceManagementUserExperienceAnalyticsAnomalyDeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAnomalyDevice%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAnomalyDeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersion%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthApplicationPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByOSVersion%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDeviceModelPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDevicePerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDevicePerformanceDetails%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthDevicePerformanceDetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOSVersionPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthOSVersionPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsAppHealthOSVersionPerformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*DeviceManagementUserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*DeviceManagementUserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBaselinesById provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselinesById(id string)(*DeviceManagementUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBaseline%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpact()(*DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpactById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactUserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthAppImpact%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthAppImpactUserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthCapacityDetails provides operations to manage the userExperienceAnalyticsBatteryHealthCapacityDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthCapacityDetails()(*DeviceManagementUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpactById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDeviceAppImpact%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDevicePerformance provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformance()(*DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDevicePerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDevicePerformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthModelPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformance()(*DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthModelPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceUserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthModelPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthModelPerformanceUserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthOsPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformance()(*DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthOsPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceUserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthOsPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthOsPerformanceUserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthRuntimeDetails provides operations to manage the userExperienceAnalyticsBatteryHealthRuntimeDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthRuntimeDetails()(*DeviceManagementUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*DeviceManagementUserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsCategoriesById provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategoriesById(id string)(*DeviceManagementUserExperienceAnalyticsCategoriesUserExperienceAnalyticsCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsCategory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsCategoriesUserExperienceAnalyticsCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceMetricHistory provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistory()(*DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceMetricHistoryById provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistoryById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetricHistory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*DeviceManagementUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDevicePerformanceById provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsDevicePerformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDevicePerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDevicePerformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScope provides operations to manage the userExperienceAnalyticsDeviceScope property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScope()(*DeviceManagementUserExperienceAnalyticsDeviceScopeRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScopes provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopes()(*DeviceManagementUserExperienceAnalyticsDeviceScopesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScopesById provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopesById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceScope%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*DeviceManagementUserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScoresById provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScoresById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceScoresUserExperienceAnalyticsDeviceScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceScores%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceScoresUserExperienceAnalyticsDeviceScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupHistoryById provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistoryById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceStartupHistoryUserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupHistory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupHistoryUserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessesById provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessesById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcess%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceById provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcessPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentity provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentity()(*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentityById provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentityById(id string)(*DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceWithoutCloudIdentity%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsImpactingProcess provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcess()(*DeviceManagementUserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsImpactingProcessById provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcessById(id string)(*DeviceManagementUserExperienceAnalyticsImpactingProcessUserExperienceAnalyticsImpactingProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsImpactingProcess%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsImpactingProcessUserExperienceAnalyticsImpactingProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*DeviceManagementUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsMetricHistoryById provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistoryById(id string)(*DeviceManagementUserExperienceAnalyticsMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetricHistory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*DeviceManagementUserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsModelScoresById provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScoresById(id string)(*DeviceManagementUserExperienceAnalyticsModelScoresUserExperienceAnalyticsModelScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsModelScores%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsModelScoresUserExperienceAnalyticsModelScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsNotAutopilotReadyDevice provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDevice()(*DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceById provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDeviceById(id string)(*DeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsNotAutopilotReadyDevice%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*DeviceManagementUserExperienceAnalyticsOverviewRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsRemoteConnection provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnection()(*DeviceManagementUserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsRemoteConnectionById provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnectionById(id string)(*DeviceManagementUserExperienceAnalyticsRemoteConnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsRemoteConnection%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsRemoteConnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsResourcePerformance provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformance()(*DeviceManagementUserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsResourcePerformanceById provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsResourcePerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*DeviceManagementUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsScoreHistoryById provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistoryById(id string)(*DeviceManagementUserExperienceAnalyticsScoreHistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsScoreHistory%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsScoreHistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsSummarizedDeviceScopes provides operations to call the userExperienceAnalyticsSummarizedDeviceScopes method.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizedDeviceScopes()(*DeviceManagementUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*DeviceManagementUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetricsById(id string)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereMetric%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById(id string)(*DeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereModelPerformance%2Did"] = id
    }
    return NewDeviceManagementUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserPfxCertificates provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificates()(*DeviceManagementUserPfxCertificatesRequestBuilder) {
    return NewDeviceManagementUserPfxCertificatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserPfxCertificatesById provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificatesById(id string)(*DeviceManagementUserPfxCertificatesUserPFXCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userPFXCertificate%2Did"] = id
    }
    return NewDeviceManagementUserPfxCertificatesUserPFXCertificateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*DeviceManagementVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewDeviceManagementVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, domainName);
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*DeviceManagementVirtualEndpointRequestBuilder) {
    return NewDeviceManagementVirtualEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeploymentProfiles provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfiles()(*DeviceManagementWindowsAutopilotDeploymentProfilesRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeploymentProfilesById provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfilesById(id string)(*DeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeploymentProfile%2Did"] = id
    }
    return NewDeviceManagementWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*DeviceManagementWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeviceIdentitiesById provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentitiesById(id string)(*DeviceManagementWindowsAutopilotDeviceIdentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewDeviceManagementWindowsAutopilotDeviceIdentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsAutopilotSettings provides operations to manage the windowsAutopilotSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotSettings()(*DeviceManagementWindowsAutopilotSettingsRequestBuilder) {
    return NewDeviceManagementWindowsAutopilotSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDriverUpdateProfiles provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfiles()(*DeviceManagementWindowsDriverUpdateProfilesRequestBuilder) {
    return NewDeviceManagementWindowsDriverUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDriverUpdateProfilesById provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfilesById(id string)(*DeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateProfile%2Did"] = id
    }
    return NewDeviceManagementWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsFeatureUpdateProfiles provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfiles()(*DeviceManagementWindowsFeatureUpdateProfilesRequestBuilder) {
    return NewDeviceManagementWindowsFeatureUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsFeatureUpdateProfilesById provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfilesById(id string)(*DeviceManagementWindowsFeatureUpdateProfilesWindowsFeatureUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsFeatureUpdateProfile%2Did"] = id
    }
    return NewDeviceManagementWindowsFeatureUpdateProfilesWindowsFeatureUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*DeviceManagementWindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewDeviceManagementWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionAppLearningSummariesById provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummariesById(id string)(*DeviceManagementWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionAppLearningSummary%2Did"] = id
    }
    return NewDeviceManagementWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*DeviceManagementWindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewDeviceManagementWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionNetworkLearningSummariesById provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummariesById(id string)(*DeviceManagementWindowsInformationProtectionNetworkLearningSummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionNetworkLearningSummary%2Did"] = id
    }
    return NewDeviceManagementWindowsInformationProtectionNetworkLearningSummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*DeviceManagementWindowsMalwareInformationRequestBuilder) {
    return NewDeviceManagementWindowsMalwareInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsMalwareInformationById provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformationById(id string)(*DeviceManagementWindowsMalwareInformationWindowsMalwareInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsMalwareInformation%2Did"] = id
    }
    return NewDeviceManagementWindowsMalwareInformationWindowsMalwareInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsQualityUpdateProfiles provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfiles()(*DeviceManagementWindowsQualityUpdateProfilesRequestBuilder) {
    return NewDeviceManagementWindowsQualityUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsQualityUpdateProfilesById provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfilesById(id string)(*DeviceManagementWindowsQualityUpdateProfilesWindowsQualityUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsQualityUpdateProfile%2Did"] = id
    }
    return NewDeviceManagementWindowsQualityUpdateProfilesWindowsQualityUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsUpdateCatalogItems provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItems()(*DeviceManagementWindowsUpdateCatalogItemsRequestBuilder) {
    return NewDeviceManagementWindowsUpdateCatalogItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsUpdateCatalogItemsById provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItemsById(id string)(*DeviceManagementWindowsUpdateCatalogItemsWindowsUpdateCatalogItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsUpdateCatalogItem%2Did"] = id
    }
    return NewDeviceManagementWindowsUpdateCatalogItemsWindowsUpdateCatalogItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ZebraFotaArtifacts provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifacts()(*DeviceManagementZebraFotaArtifactsRequestBuilder) {
    return NewDeviceManagementZebraFotaArtifactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaArtifactsById provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifactsById(id string)(*DeviceManagementZebraFotaArtifactsZebraFotaArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["zebraFotaArtifact%2Did"] = id
    }
    return NewDeviceManagementZebraFotaArtifactsZebraFotaArtifactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ZebraFotaConnector provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaConnector()(*DeviceManagementZebraFotaConnectorRequestBuilder) {
    return NewDeviceManagementZebraFotaConnectorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaDeployments provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeployments()(*DeviceManagementZebraFotaDeploymentsRequestBuilder) {
    return NewDeviceManagementZebraFotaDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaDeploymentsById provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeploymentsById(id string)(*DeviceManagementZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["zebraFotaDeployment%2Did"] = id
    }
    return NewDeviceManagementZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
