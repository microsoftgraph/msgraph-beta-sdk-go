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
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRequestBuilderGetQueryParameters
}
// DeviceManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedThreatProtectionOnboardingStateSummary provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidDeviceOwnerEnrollmentProfiles provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfiles()(*AndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidDeviceOwnerEnrollmentProfilesById provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfilesById(id string)(*AndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidDeviceOwnerEnrollmentProfile%2Did"] = id
    }
    return NewAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkAppConfigurationSchemas provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemas()(*AndroidForWorkAppConfigurationSchemasRequestBuilder) {
    return NewAndroidForWorkAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidForWorkAppConfigurationSchemasById provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemasById(id string)(*AndroidForWorkAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidForWorkAppConfigurationSchema%2Did"] = id
    }
    return NewAndroidForWorkAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkEnrollmentProfiles provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfiles()(*AndroidForWorkEnrollmentProfilesRequestBuilder) {
    return NewAndroidForWorkEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidForWorkEnrollmentProfilesById provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfilesById(id string)(*AndroidForWorkEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidForWorkEnrollmentProfile%2Did"] = id
    }
    return NewAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AndroidForWorkSettings provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkSettings()(*AndroidForWorkSettingsRequestBuilder) {
    return NewAndroidForWorkSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAccountEnterpriseSettings provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAppConfigurationSchemas provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemas()(*AndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    return NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedStoreAppConfigurationSchemasById provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemasById(id string)(*AndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedStoreAppConfigurationSchema%2Did"] = id
    }
    return NewAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*ApplePushNotificationCertificateRequestBuilder) {
    return NewApplePushNotificationCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppleUserInitiatedEnrollmentProfiles provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfiles()(*AppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppleUserInitiatedEnrollmentProfilesById provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfilesById(id string)(*AppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleUserInitiatedEnrollmentProfile%2Did"] = id
    }
    return NewAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentFilters provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFilters()(*AssignmentFiltersRequestBuilder) {
    return NewAssignmentFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentFiltersById provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFiltersById(id string)(*DeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAndAppManagementAssignmentFilter%2Did"] = id
    }
    return NewDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEvents()(*AuditEventsRequestBuilder) {
    return NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEventsById(id string)(*AuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["auditEvent%2Did"] = id
    }
    return NewAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AutopilotEvents provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEvents()(*AutopilotEventsRequestBuilder) {
    return NewAutopilotEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AutopilotEventsById provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEventsById(id string)(*DeviceManagementAutopilotEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementAutopilotEvent%2Did"] = id
    }
    return NewDeviceManagementAutopilotEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CartToClassAssociations provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociations()(*CartToClassAssociationsRequestBuilder) {
    return NewCartToClassAssociationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CartToClassAssociationsById provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociationsById(id string)(*CartToClassAssociationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cartToClassAssociation%2Did"] = id
    }
    return NewCartToClassAssociationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Categories()(*CategoriesRequestBuilder) {
    return NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CategoriesById(id string)(*DeviceManagementSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingCategory%2Did"] = id
    }
    return NewDeviceManagementSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CertificateConnectorDetails provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetails()(*CertificateConnectorDetailsRequestBuilder) {
    return NewCertificateConnectorDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CertificateConnectorDetailsById provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetailsById(id string)(*CertificateConnectorDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["certificateConnectorDetails%2Did"] = id
    }
    return NewCertificateConnectorDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ChromeOSOnboardingSettings provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettings()(*ChromeOSOnboardingSettingsRequestBuilder) {
    return NewChromeOSOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChromeOSOnboardingSettingsById provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettingsById(id string)(*ChromeOSOnboardingSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chromeOSOnboardingSettings%2Did"] = id
    }
    return NewChromeOSOnboardingSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CloudPCConnectivityIssues provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssues()(*CloudPCConnectivityIssuesRequestBuilder) {
    return NewCloudPCConnectivityIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CloudPCConnectivityIssuesById provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssuesById(id string)(*CloudPCConnectivityIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["cloudPCConnectivityIssue%2Did"] = id
    }
    return NewCloudPCConnectivityIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComanagedDevices provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevices()(*ComanagedDevicesRequestBuilder) {
    return NewComanagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComanagedDevicesById provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevicesById(id string)(*ManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComanagementEligibleDevices provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevices()(*ComanagementEligibleDevicesRequestBuilder) {
    return NewComanagementEligibleDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComanagementEligibleDevicesById provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevicesById(id string)(*ComanagementEligibleDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["comanagementEligibleDevice%2Did"] = id
    }
    return NewComanagementEligibleDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceCategories provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategories()(*ComplianceCategoriesRequestBuilder) {
    return NewComplianceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceCategoriesById provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategoriesById(id string)(*DeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationCategory%2Did"] = id
    }
    return NewDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*ComplianceManagementPartnersRequestBuilder) {
    return NewComplianceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceManagementPartnersById provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartnersById(id string)(*ComplianceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["complianceManagementPartner%2Did"] = id
    }
    return NewComplianceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CompliancePolicies provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePolicies()(*CompliancePoliciesRequestBuilder) {
    return NewCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompliancePoliciesById provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePoliciesById(id string)(*DeviceManagementCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementCompliancePolicy%2Did"] = id
    }
    return NewDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ComplianceSettings provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettings()(*ComplianceSettingsRequestBuilder) {
    return NewComplianceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceSettingsById provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettingsById(id string)(*DeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*ConditionalAccessSettingsRequestBuilder) {
    return NewConditionalAccessSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigManagerCollections provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollections()(*ConfigManagerCollectionsRequestBuilder) {
    return NewConfigManagerCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigManagerCollectionsById provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollectionsById(id string)(*ConfigManagerCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["configManagerCollection%2Did"] = id
    }
    return NewConfigManagerCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationCategories provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategories()(*ConfigurationCategoriesRequestBuilder) {
    return NewConfigurationCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationCategoriesById provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategoriesById(id string)(*DeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationCategory%2Did"] = id
    }
    return NewDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationPolicies provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicies()(*ConfigurationPoliciesRequestBuilder) {
    return NewConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationPoliciesById provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPoliciesById(id string)(*DeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicy%2Did"] = id
    }
    return NewDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationPolicyTemplates provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplates()(*ConfigurationPolicyTemplatesRequestBuilder) {
    return NewConfigurationPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationPolicyTemplatesById provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplatesById(id string)(*DeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicyTemplate%2Did"] = id
    }
    return NewDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConfigurationSettings provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettings()(*ConfigurationSettingsRequestBuilder) {
    return NewConfigurationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConfigurationSettingsById provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettingsById(id string)(*DeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// CreatePatchRequestInformation update deviceManagement
func (m *DeviceManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataSharingConsents provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsents()(*DataSharingConsentsRequestBuilder) {
    return NewDataSharingConsentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataSharingConsentsById provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsentsById(id string)(*DataSharingConsentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataSharingConsent%2Did"] = id
    }
    return NewDataSharingConsentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DepOnboardingSettings provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettings()(*DepOnboardingSettingsRequestBuilder) {
    return NewDepOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DepOnboardingSettingsById provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettingsById(id string)(*DepOnboardingSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["depOnboardingSetting%2Did"] = id
    }
    return NewDepOnboardingSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DerivedCredentials provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentials()(*DerivedCredentialsRequestBuilder) {
    return NewDerivedCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DerivedCredentialsById provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentialsById(id string)(*DeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementDerivedCredentialSettings%2Did"] = id
    }
    return NewDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DetectedAppsRequestBuilder) {
    return NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedAppsById(id string)(*DetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp%2Did"] = id
    }
    return NewDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DeviceCategoriesRequestBuilder) {
    return NewDeviceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCategoriesById provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategoriesById(id string)(*DeviceCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCategory%2Did"] = id
    }
    return NewDeviceCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DeviceCompliancePoliciesRequestBuilder) {
    return NewDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePoliciesById provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePoliciesById(id string)(*DeviceCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicy%2Did"] = id
    }
    return NewDeviceCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*DeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = id
    }
    return NewDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceComplianceScripts provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScripts()(*DeviceComplianceScriptsRequestBuilder) {
    return NewDeviceComplianceScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceComplianceScriptsById provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScriptsById(id string)(*DeviceComplianceScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceComplianceScript%2Did"] = id
    }
    return NewDeviceComplianceScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationConflictSummary provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceConfigurationConflictSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationConflictSummaryById provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummaryById(id string)(*DeviceConfigurationConflictSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfigurationConflictSummary%2Did"] = id
    }
    return NewDeviceConfigurationConflictSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationRestrictedAppsViolations provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolations()(*DeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationRestrictedAppsViolationsById provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolationsById(id string)(*RestrictedAppsViolationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["restrictedAppsViolation%2Did"] = id
    }
    return NewRestrictedAppsViolationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceConfigurationsRequestBuilder) {
    return NewDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationsAllManagedDeviceCertificateStates provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStates()(*DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesById provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStatesById(id string)(*ManagedAllDeviceCertificateStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAllDeviceCertificateState%2Did"] = id
    }
    return NewManagedAllDeviceCertificateStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationsById provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsById(id string)(*DeviceConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfiguration%2Did"] = id
    }
    return NewDeviceConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationUserStateSummaries provides operations to manage the deviceConfigurationUserStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationUserStateSummaries()(*DeviceConfigurationUserStateSummariesRequestBuilder) {
    return NewDeviceConfigurationUserStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCustomAttributeShellScripts provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScripts()(*DeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCustomAttributeShellScriptsById provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScriptsById(id string)(*DeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCustomAttributeShellScript%2Did"] = id
    }
    return NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*DeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = id
    }
    return NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceHealthScripts provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScripts()(*DeviceHealthScriptsRequestBuilder) {
    return NewDeviceHealthScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceHealthScriptsById provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScriptsById(id string)(*DeviceHealthScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceHealthScript%2Did"] = id
    }
    return NewDeviceHealthScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DeviceManagementPartnersRequestBuilder) {
    return NewDeviceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementPartnersById provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartnersById(id string)(*DeviceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementPartner%2Did"] = id
    }
    return NewDeviceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceManagementScripts provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScripts()(*DeviceManagementScriptsRequestBuilder) {
    return NewDeviceManagementScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementScriptsById provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScriptsById(id string)(*DeviceManagementScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScript%2Did"] = id
    }
    return NewDeviceManagementScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceShellScripts provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScripts()(*DeviceShellScriptsRequestBuilder) {
    return NewDeviceShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceShellScriptsById provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScriptsById(id string)(*DeviceShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceShellScript%2Did"] = id
    }
    return NewDeviceShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DomainJoinConnectors provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectors()(*DomainJoinConnectorsRequestBuilder) {
    return NewDomainJoinConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DomainJoinConnectorsById provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectorsById(id string)(*DeviceManagementDomainJoinConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementDomainJoinConnector%2Did"] = id
    }
    return NewDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EmbeddedSIMActivationCodePools provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePools()(*EmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EmbeddedSIMActivationCodePoolsById provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePoolsById(id string)(*EmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePool%2Did"] = id
    }
    return NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EnableAndroidDeviceAdministratorEnrollment provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
func (m *DeviceManagementRequestBuilder) EnableAndroidDeviceAdministratorEnrollment()(*EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    return NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableLegacyPcManagement provides operations to call the enableLegacyPcManagement method.
func (m *DeviceManagementRequestBuilder) EnableLegacyPcManagement()(*EnableLegacyPcManagementRequestBuilder) {
    return NewEnableLegacyPcManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EnableUnlicensedAdminstrators provides operations to call the enableUnlicensedAdminstrators method.
func (m *DeviceManagementRequestBuilder) EnableUnlicensedAdminstrators()(*EnableUnlicensedAdminstratorsRequestBuilder) {
    return NewEnableUnlicensedAdminstratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateAssignmentFilter provides operations to call the evaluateAssignmentFilter method.
func (m *DeviceManagementRequestBuilder) EvaluateAssignmentFilter()(*EvaluateAssignmentFilterRequestBuilder) {
    return NewEvaluateAssignmentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ExchangeConnectorsRequestBuilder) {
    return NewExchangeConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeConnectorsById provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectorsById(id string)(*DeviceManagementExchangeConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExchangeConnector%2Did"] = id
    }
    return NewDeviceManagementExchangeConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExchangeOnPremisesPolicies provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicies()(*ExchangeOnPremisesPoliciesRequestBuilder) {
    return NewExchangeOnPremisesPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeOnPremisesPoliciesById provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPoliciesById(id string)(*DeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExchangeOnPremisesPolicy%2Did"] = id
    }
    return NewDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ExchangeOnPremisesPolicy provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicy()(*ExchangeOnPremisesPolicyRequestBuilder) {
    return NewExchangeOnPremisesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *DeviceManagementRequestBuilder) GetAssignedRoleDetails()(*GetAssignedRoleDetailsRequestBuilder) {
    return NewGetAssignedRoleDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAssignmentFiltersStatusDetails provides operations to call the getAssignmentFiltersStatusDetails method.
func (m *DeviceManagementRequestBuilder) GetAssignmentFiltersStatusDetails()(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    return NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComanagedDevicesSummary provides operations to call the getComanagedDevicesSummary method.
func (m *DeviceManagementRequestBuilder) GetComanagedDevicesSummary()(*GetComanagedDevicesSummaryRequestBuilder) {
    return NewGetComanagedDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComanagementEligibleDevicesSummary provides operations to call the getComanagementEligibleDevicesSummary method.
func (m *DeviceManagementRequestBuilder) GetComanagementEligibleDevicesSummary()(*GetComanagementEligibleDevicesSummaryRequestBuilder) {
    return NewGetComanagementEligibleDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) GetEffectivePermissions()(*GetEffectivePermissionsRequestBuilder) {
    return NewGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*GetEffectivePermissionsWithScopeRequestBuilder) {
    return NewGetEffectivePermissionsWithScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter, scope);
}
// GetRoleScopeTagsByIdsWithIds provides operations to call the getRoleScopeTagsByIds method.
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByIdsWithIds(ids *string)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter, ids);
}
// GetRoleScopeTagsByResourceWithResource provides operations to call the getRoleScopeTagsByResource method.
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByResourceWithResource(resource *string)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource);
}
// GetSuggestedEnrollmentLimitWithEnrollmentType provides operations to call the getSuggestedEnrollmentLimit method.
func (m *DeviceManagementRequestBuilder) GetSuggestedEnrollmentLimitWithEnrollmentType(enrollmentType *string)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter, enrollmentType);
}
// GroupPolicyCategories provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategories()(*GroupPolicyCategoriesRequestBuilder) {
    return NewGroupPolicyCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyCategoriesById provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategoriesById(id string)(*GroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyCategory%2Did"] = id
    }
    return NewGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyConfigurations provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurations()(*GroupPolicyConfigurationsRequestBuilder) {
    return NewGroupPolicyConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyConfigurationsById provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurationsById(id string)(*GroupPolicyConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfiguration%2Did"] = id
    }
    return NewGroupPolicyConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyDefinitionFiles provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFiles()(*GroupPolicyDefinitionFilesRequestBuilder) {
    return NewGroupPolicyDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyDefinitionFilesById provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFilesById(id string)(*GroupPolicyDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionFile%2Did"] = id
    }
    return NewGroupPolicyDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyDefinitions provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitions()(*GroupPolicyDefinitionsRequestBuilder) {
    return NewGroupPolicyDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyDefinitionsById provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionsById(id string)(*GroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = id
    }
    return NewGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyMigrationReports provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReports()(*GroupPolicyMigrationReportsRequestBuilder) {
    return NewGroupPolicyMigrationReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyMigrationReportsById provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReportsById(id string)(*GroupPolicyMigrationReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyMigrationReport%2Did"] = id
    }
    return NewGroupPolicyMigrationReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyObjectFiles provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFiles()(*GroupPolicyObjectFilesRequestBuilder) {
    return NewGroupPolicyObjectFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyObjectFilesById provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFilesById(id string)(*GroupPolicyObjectFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyObjectFile%2Did"] = id
    }
    return NewGroupPolicyObjectFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GroupPolicyUploadedDefinitionFiles provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFiles()(*GroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicyUploadedDefinitionFilesById provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFilesById(id string)(*GroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyUploadedDefinitionFile%2Did"] = id
    }
    return NewGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ImportedDeviceIdentities provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentities()(*ImportedDeviceIdentitiesRequestBuilder) {
    return NewImportedDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedDeviceIdentitiesById provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentitiesById(id string)(*ImportedDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedDeviceIdentity%2Did"] = id
    }
    return NewImportedDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedWindowsAutopilotDeviceIdentitiesById provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentitiesById(id string)(*ImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedWindowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Intents provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Intents()(*IntentsRequestBuilder) {
    return NewIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntentsById provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntentsById(id string)(*DeviceManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementIntent%2Did"] = id
    }
    return NewDeviceManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IntuneBrandingProfiles provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfiles()(*IntuneBrandingProfilesRequestBuilder) {
    return NewIntuneBrandingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IntuneBrandingProfilesById provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfilesById(id string)(*IntuneBrandingProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["intuneBrandingProfile%2Did"] = id
    }
    return NewIntuneBrandingProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*IosUpdateStatusesRequestBuilder) {
    return NewIosUpdateStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosUpdateStatusesById provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatusesById(id string)(*IosUpdateDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosUpdateDeviceStatus%2Did"] = id
    }
    return NewIosUpdateDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MacOSSoftwareUpdateAccountSummaries provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummaries()(*MacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    return NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MacOSSoftwareUpdateAccountSummariesById provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummariesById(id string)(*MacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["macOSSoftwareUpdateAccountSummary%2Did"] = id
    }
    return NewMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceEncryptionStates provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStates()(*ManagedDeviceEncryptionStatesRequestBuilder) {
    return NewManagedDeviceEncryptionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEncryptionStatesById provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStatesById(id string)(*ManagedDeviceEncryptionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceEncryptionState%2Did"] = id
    }
    return NewManagedDeviceEncryptionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*ManagedDeviceOverviewRequestBuilder) {
    return NewManagedDeviceOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*ManagedDevicesRequestBuilder) {
    return NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevicesById(id string)(*ManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice%2Did"] = id
    }
    return NewManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelConfigurations provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurations()(*MicrosoftTunnelConfigurationsRequestBuilder) {
    return NewMicrosoftTunnelConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelConfigurationsById provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurationsById(id string)(*MicrosoftTunnelConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelConfiguration%2Did"] = id
    }
    return NewMicrosoftTunnelConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelHealthThresholds provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholds()(*MicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewMicrosoftTunnelHealthThresholdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelHealthThresholdsById provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholdsById(id string)(*MicrosoftTunnelHealthThresholdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelHealthThreshold%2Did"] = id
    }
    return NewMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelServerLogCollectionResponses provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponses()(*MicrosoftTunnelServerLogCollectionResponsesRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServerLogCollectionResponsesById provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponsesById(id string)(*MicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelServerLogCollectionResponse%2Did"] = id
    }
    return NewMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MicrosoftTunnelSites provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSites()(*MicrosoftTunnelSitesRequestBuilder) {
    return NewMicrosoftTunnelSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelSitesById provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSitesById(id string)(*MicrosoftTunnelSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelSite%2Did"] = id
    }
    return NewMicrosoftTunnelSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*MobileAppTroubleshootingEventsRequestBuilder) {
    return NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppTroubleshootingEventsById provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*MobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppTroubleshootingEvent%2Did"] = id
    }
    return NewMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*MobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobileThreatDefenseConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileThreatDefenseConnectorsById provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectorsById(id string)(*MobileThreatDefenseConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileThreatDefenseConnector%2Did"] = id
    }
    return NewMobileThreatDefenseConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Monitoring provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Monitoring()(*MonitoringRequestBuilder) {
    return NewMonitoringRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NdesConnectors provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectors()(*NdesConnectorsRequestBuilder) {
    return NewNdesConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NdesConnectorsById provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectorsById(id string)(*NdesConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ndesConnector%2Did"] = id
    }
    return NewNdesConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*NotificationMessageTemplatesRequestBuilder) {
    return NewNotificationMessageTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationMessageTemplatesById provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplatesById(id string)(*NotificationMessageTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notificationMessageTemplate%2Did"] = id
    }
    return NewNotificationMessageTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OemWarrantyInformationOnboarding provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboardingRequestBuilder) {
    return NewOemWarrantyInformationOnboardingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OemWarrantyInformationOnboardingById provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboardingById(id string)(*OemWarrantyInformationOnboardingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["oemWarrantyInformationOnboarding%2Did"] = id
    }
    return NewOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *DeviceManagementRequestBuilder) RemoteActionAudits()(*RemoteActionAuditsRequestBuilder) {
    return NewRemoteActionAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteActionAuditsById provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteActionAuditsById(id string)(*RemoteActionAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["remoteActionAudit%2Did"] = id
    }
    return NewRemoteActionAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*RemoteAssistancePartnersRequestBuilder) {
    return NewRemoteAssistancePartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteAssistancePartnersById provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartnersById(id string)(*RemoteAssistancePartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["remoteAssistancePartner%2Did"] = id
    }
    return NewRemoteAssistancePartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoteAssistanceSettings provides operations to manage the remoteAssistanceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistanceSettings()(*RemoteAssistanceSettingsRequestBuilder) {
    return NewRemoteAssistanceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceAccessProfiles provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfiles()(*ResourceAccessProfilesRequestBuilder) {
    return NewResourceAccessProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceAccessProfilesById provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfilesById(id string)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementResourceAccessProfileBase%2Did"] = id
    }
    return NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*ResourceOperationsRequestBuilder) {
    return NewResourceOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceOperationsById provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperationsById(id string)(*ResourceOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceOperation%2Did"] = id
    }
    return NewResourceOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReusablePolicySettings provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettings()(*ReusablePolicySettingsRequestBuilder) {
    return NewReusablePolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReusablePolicySettingsById provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettingsById(id string)(*DeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementReusablePolicySetting%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ReusableSettings provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettings()(*ReusableSettingsRequestBuilder) {
    return NewReusableSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReusableSettingsById provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettingsById(id string)(*DeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*RoleAssignmentsRequestBuilder) {
    return NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*DeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAndAppManagementRoleAssignment%2Did"] = id
    }
    return NewDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*RoleDefinitionsRequestBuilder) {
    return NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*RoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleDefinition%2Did"] = id
    }
    return NewRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTags()(*RoleScopeTagsRequestBuilder) {
    return NewRoleScopeTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleScopeTagsById provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTagsById(id string)(*RoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleScopeTag%2Did"] = id
    }
    return NewRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ScopedForResourceWithResource provides operations to call the scopedForResource method.
func (m *DeviceManagementRequestBuilder) ScopedForResourceWithResource(resource *string)(*ScopedForResourceWithResourceRequestBuilder) {
    return NewScopedForResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource);
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *DeviceManagementRequestBuilder) SendCustomNotificationToCompanyPortal()(*SendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitions()(*SettingDefinitionsRequestBuilder) {
    return NewSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingDefinitionsById provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitionsById(id string)(*DeviceManagementSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingDefinition%2Did"] = id
    }
    return NewDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareUpdateStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*TelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomExpenseManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TelecomExpenseManagementPartnersById provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartnersById(id string)(*TelecomExpenseManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["telecomExpenseManagementPartner%2Did"] = id
    }
    return NewTelecomExpenseManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Templates provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplatesById(id string)(*DeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTemplate%2Did"] = id
    }
    return NewDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TemplateSettings provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettings()(*TemplateSettingsRequestBuilder) {
    return NewTemplateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TemplateSettingsById provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettingsById(id string)(*DeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationSettingTemplate%2Did"] = id
    }
    return NewDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TenantAttachRBAC provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TenantAttachRBAC()(*TenantAttachRBACRequestBuilder) {
    return NewTenantAttachRBACRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*TermsAndConditionsRequestBuilder) {
    return NewTermsAndConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsAndConditionsById provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditionsById(id string)(*TermsAndConditionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditions%2Did"] = id
    }
    return NewTermsAndConditionsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TroubleshootingEvents provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*TroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TroubleshootingEventsById provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEventsById(id string)(*DeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = id
    }
    return NewDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAnomaly provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomaly()(*UserExperienceAnalyticsAnomalyRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyById provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyById(id string)(*UserExperienceAnalyticsAnomalyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAnomaly%2Did"] = id
    }
    return NewUserExperienceAnalyticsAnomalyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyDevice provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDevice()(*UserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAnomalyDeviceById provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDeviceById(id string)(*UserExperienceAnalyticsAnomalyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAnomalyDevice%2Did"] = id
    }
    return NewUserExperienceAnalyticsAnomalyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById(id string)(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersion%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById(id string)(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById(id string)(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthApplicationPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById(id string)(*UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthAppPerformanceByOSVersion%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformanceById(id string)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDeviceModelPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*UserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceById(id string)(*UserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDevicePerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById(id string)(*UserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthDevicePerformanceDetails%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOSVersionPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformanceById(id string)(*UserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsAppHealthOSVersionPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*UserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*UserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBaselinesById provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselinesById(id string)(*UserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBaseline%2Did"] = id
    }
    return NewUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpact()(*UserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpactById(id string)(*UserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthAppImpact%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthCapacityDetails provides operations to manage the userExperienceAnalyticsBatteryHealthCapacityDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpactById(id string)(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDeviceAppImpact%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDevicePerformance provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDevicePerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById(id string)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthModelPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthModelPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthModelPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthOsPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthOsPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsBatteryHealthOsPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsBatteryHealthRuntimeDetails provides operations to manage the userExperienceAnalyticsBatteryHealthRuntimeDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*UserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsCategoriesById provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategoriesById(id string)(*UserExperienceAnalyticsCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsCategory%2Did"] = id
    }
    return NewUserExperienceAnalyticsCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceMetricHistory provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistory()(*UserExperienceAnalyticsDeviceMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceMetricHistoryById provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistoryById(id string)(*UserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetricHistory%2Did"] = id
    }
    return NewUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDevicePerformanceById provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformanceById(id string)(*UserExperienceAnalyticsDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDevicePerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScope provides operations to manage the userExperienceAnalyticsDeviceScope property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScope()(*UserExperienceAnalyticsDeviceScopeRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScopes provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopes()(*UserExperienceAnalyticsDeviceScopesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScopesById provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopesById(id string)(*UserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceScope%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*UserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceScoresById provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScoresById(id string)(*UserExperienceAnalyticsDeviceScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceScores%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupHistoryById provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistoryById(id string)(*UserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupHistory%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*UserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessesById provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessesById(id string)(*UserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcess%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceById provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformanceById(id string)(*UserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceStartupProcessPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentity provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentity()(*UserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentityById provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentityById(id string)(*UserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsDeviceWithoutCloudIdentity%2Did"] = id
    }
    return NewUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsImpactingProcess provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsImpactingProcessById provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcessById(id string)(*UserExperienceAnalyticsImpactingProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsImpactingProcess%2Did"] = id
    }
    return NewUserExperienceAnalyticsImpactingProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsMetricHistoryById provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistoryById(id string)(*UserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsMetricHistory%2Did"] = id
    }
    return NewUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsModelScoresById provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScoresById(id string)(*UserExperienceAnalyticsModelScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsModelScores%2Did"] = id
    }
    return NewUserExperienceAnalyticsModelScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsNotAutopilotReadyDevice provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceById provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDeviceById(id string)(*UserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsNotAutopilotReadyDevice%2Did"] = id
    }
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsRemoteConnection provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsRemoteConnectionById provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnectionById(id string)(*UserExperienceAnalyticsRemoteConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsRemoteConnection%2Did"] = id
    }
    return NewUserExperienceAnalyticsRemoteConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsResourcePerformance provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsResourcePerformanceById provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformanceById(id string)(*UserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsResourcePerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*UserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsScoreHistoryById provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistoryById(id string)(*UserExperienceAnalyticsScoreHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsScoreHistory%2Did"] = id
    }
    return NewUserExperienceAnalyticsScoreHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsSummarizedDeviceScopes provides operations to call the userExperienceAnalyticsSummarizedDeviceScopes method.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizedDeviceScopes()(*UserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder) {
    return NewUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetricsById(id string)(*UserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereMetric%2Did"] = id
    }
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById(id string)(*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereModelPerformance%2Did"] = id
    }
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserPfxCertificates provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificates()(*UserPfxCertificatesRequestBuilder) {
    return NewUserPfxCertificatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserPfxCertificatesById provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificatesById(id string)(*UserPFXCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userPFXCertificate%2Did"] = id
    }
    return NewUserPFXCertificateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, domainName);
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*VirtualEndpointRequestBuilder) {
    return NewVirtualEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeploymentProfiles provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfiles()(*WindowsAutopilotDeploymentProfilesRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeploymentProfilesById provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfilesById(id string)(*WindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeploymentProfile%2Did"] = id
    }
    return NewWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*WindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeviceIdentitiesById provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentitiesById(id string)(*WindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity%2Did"] = id
    }
    return NewWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsAutopilotSettings provides operations to manage the windowsAutopilotSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotSettings()(*WindowsAutopilotSettingsRequestBuilder) {
    return NewWindowsAutopilotSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDriverUpdateProfiles provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfiles()(*WindowsDriverUpdateProfilesRequestBuilder) {
    return NewWindowsDriverUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsDriverUpdateProfilesById provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfilesById(id string)(*WindowsDriverUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsDriverUpdateProfile%2Did"] = id
    }
    return NewWindowsDriverUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsFeatureUpdateProfiles provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfiles()(*WindowsFeatureUpdateProfilesRequestBuilder) {
    return NewWindowsFeatureUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsFeatureUpdateProfilesById provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfilesById(id string)(*WindowsFeatureUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsFeatureUpdateProfile%2Did"] = id
    }
    return NewWindowsFeatureUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*WindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionAppLearningSummariesById provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummariesById(id string)(*WindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionAppLearningSummary%2Did"] = id
    }
    return NewWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*WindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionNetworkLearningSummariesById provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummariesById(id string)(*WindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionNetworkLearningSummary%2Did"] = id
    }
    return NewWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*WindowsMalwareInformationRequestBuilder) {
    return NewWindowsMalwareInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsMalwareInformationById provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformationById(id string)(*WindowsMalwareInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsMalwareInformation%2Did"] = id
    }
    return NewWindowsMalwareInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsQualityUpdateProfiles provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfiles()(*WindowsQualityUpdateProfilesRequestBuilder) {
    return NewWindowsQualityUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsQualityUpdateProfilesById provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfilesById(id string)(*WindowsQualityUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsQualityUpdateProfile%2Did"] = id
    }
    return NewWindowsQualityUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsUpdateCatalogItems provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItems()(*WindowsUpdateCatalogItemsRequestBuilder) {
    return NewWindowsUpdateCatalogItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsUpdateCatalogItemsById provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItemsById(id string)(*WindowsUpdateCatalogItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsUpdateCatalogItem%2Did"] = id
    }
    return NewWindowsUpdateCatalogItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ZebraFotaArtifacts provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifacts()(*ZebraFotaArtifactsRequestBuilder) {
    return NewZebraFotaArtifactsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaArtifactsById provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifactsById(id string)(*ZebraFotaArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["zebraFotaArtifact%2Did"] = id
    }
    return NewZebraFotaArtifactItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ZebraFotaConnector provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaConnector()(*ZebraFotaConnectorRequestBuilder) {
    return NewZebraFotaConnectorRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaDeployments provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeployments()(*ZebraFotaDeploymentsRequestBuilder) {
    return NewZebraFotaDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ZebraFotaDeploymentsById provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeploymentsById(id string)(*ZebraFotaDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["zebraFotaDeployment%2Did"] = id
    }
    return NewZebraFotaDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
