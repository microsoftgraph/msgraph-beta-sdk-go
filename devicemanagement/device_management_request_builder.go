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
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidDeviceOwnerEnrollmentProfiles provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfiles()(*AndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidDeviceOwnerEnrollmentProfilesById provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfilesById(id string)(*AndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAndroidDeviceOwnerEnrollmentProfilesAndroidDeviceOwnerEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AndroidForWorkAppConfigurationSchemas provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemas()(*AndroidForWorkAppConfigurationSchemasRequestBuilder) {
    return NewAndroidForWorkAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidForWorkAppConfigurationSchemasById provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemasById(id string)(*AndroidForWorkAppConfigurationSchemasAndroidForWorkAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAndroidForWorkAppConfigurationSchemasAndroidForWorkAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AndroidForWorkEnrollmentProfiles provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfiles()(*AndroidForWorkEnrollmentProfilesRequestBuilder) {
    return NewAndroidForWorkEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidForWorkEnrollmentProfilesById provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfilesById(id string)(*AndroidForWorkEnrollmentProfilesAndroidForWorkEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAndroidForWorkEnrollmentProfilesAndroidForWorkEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AndroidForWorkSettings provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidForWorkSettings()(*AndroidForWorkSettingsRequestBuilder) {
    return NewAndroidForWorkSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidManagedStoreAccountEnterpriseSettings provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidManagedStoreAppConfigurationSchemas provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemas()(*AndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    return NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AndroidManagedStoreAppConfigurationSchemasById provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemasById(id string)(*AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*ApplePushNotificationCertificateRequestBuilder) {
    return NewApplePushNotificationCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppleUserInitiatedEnrollmentProfiles provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfiles()(*AppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AppleUserInitiatedEnrollmentProfilesById provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfilesById(id string)(*AppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AssignmentFilters provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFilters()(*AssignmentFiltersRequestBuilder) {
    return NewAssignmentFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentFiltersById provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AssignmentFiltersById(id string)(*AssignmentFiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAssignmentFiltersDeviceAndAppManagementAssignmentFilterItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEvents()(*AuditEventsRequestBuilder) {
    return NewAuditEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AuditEventsById provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AuditEventsById(id string)(*AuditEventsAuditEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAuditEventsAuditEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// AutopilotEvents provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEvents()(*AutopilotEventsRequestBuilder) {
    return NewAutopilotEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AutopilotEventsById provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) AutopilotEventsById(id string)(*AutopilotEventsDeviceManagementAutopilotEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewAutopilotEventsDeviceManagementAutopilotEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// CartToClassAssociations provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociations()(*CartToClassAssociationsRequestBuilder) {
    return NewCartToClassAssociationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CartToClassAssociationsById provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CartToClassAssociationsById(id string)(*CartToClassAssociationsCartToClassAssociationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCartToClassAssociationsCartToClassAssociationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Categories()(*CategoriesRequestBuilder) {
    return NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CategoriesById(id string)(*CategoriesDeviceManagementSettingCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCategoriesDeviceManagementSettingCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// CertificateConnectorDetails provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetails()(*CertificateConnectorDetailsRequestBuilder) {
    return NewCertificateConnectorDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CertificateConnectorDetailsById provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetailsById(id string)(*CertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCertificateConnectorDetailsCertificateConnectorDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ChromeOSOnboardingSettings provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettings()(*ChromeOSOnboardingSettingsRequestBuilder) {
    return NewChromeOSOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ChromeOSOnboardingSettingsById provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettingsById(id string)(*ChromeOSOnboardingSettingsChromeOSOnboardingSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewChromeOSOnboardingSettingsChromeOSOnboardingSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// CloudPCConnectivityIssues provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssues()(*CloudPCConnectivityIssuesRequestBuilder) {
    return NewCloudPCConnectivityIssuesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CloudPCConnectivityIssuesById provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssuesById(id string)(*CloudPCConnectivityIssuesCloudPCConnectivityIssueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCloudPCConnectivityIssuesCloudPCConnectivityIssueItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ComanagedDevices provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevices()(*ComanagedDevicesRequestBuilder) {
    return NewComanagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ComanagedDevicesById provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagedDevicesById(id string)(*ComanagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewComanagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ComanagementEligibleDevices provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevices()(*ComanagementEligibleDevicesRequestBuilder) {
    return NewComanagementEligibleDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ComanagementEligibleDevicesById provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevicesById(id string)(*ComanagementEligibleDevicesComanagementEligibleDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewComanagementEligibleDevicesComanagementEligibleDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ComplianceCategories provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategories()(*ComplianceCategoriesRequestBuilder) {
    return NewComplianceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ComplianceCategoriesById provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceCategoriesById(id string)(*ComplianceCategoriesDeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewComplianceCategoriesDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*ComplianceManagementPartnersRequestBuilder) {
    return NewComplianceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ComplianceManagementPartnersById provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartnersById(id string)(*ComplianceManagementPartnersComplianceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewComplianceManagementPartnersComplianceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// CompliancePolicies provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePolicies()(*CompliancePoliciesRequestBuilder) {
    return NewCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// CompliancePoliciesById provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) CompliancePoliciesById(id string)(*CompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewCompliancePoliciesDeviceManagementCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ComplianceSettings provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettings()(*ComplianceSettingsRequestBuilder) {
    return NewComplianceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ComplianceSettingsById provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ComplianceSettingsById(id string)(*ComplianceSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewComplianceSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*ConditionalAccessSettingsRequestBuilder) {
    return NewConditionalAccessSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigManagerCollections provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollections()(*ConfigManagerCollectionsRequestBuilder) {
    return NewConfigManagerCollectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigManagerCollectionsById provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigManagerCollectionsById(id string)(*ConfigManagerCollectionsConfigManagerCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewConfigManagerCollectionsConfigManagerCollectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ConfigurationCategories provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategories()(*ConfigurationCategoriesRequestBuilder) {
    return NewConfigurationCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigurationCategoriesById provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationCategoriesById(id string)(*ConfigurationCategoriesDeviceManagementConfigurationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewConfigurationCategoriesDeviceManagementConfigurationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ConfigurationPolicies provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicies()(*ConfigurationPoliciesRequestBuilder) {
    return NewConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigurationPoliciesById provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPoliciesById(id string)(*ConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ConfigurationPolicyTemplates provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplates()(*ConfigurationPolicyTemplatesRequestBuilder) {
    return NewConfigurationPolicyTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigurationPolicyTemplatesById provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplatesById(id string)(*ConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewConfigurationPolicyTemplatesDeviceManagementConfigurationPolicyTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ConfigurationSettings provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettings()(*ConfigurationSettingsRequestBuilder) {
    return NewConfigurationSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ConfigurationSettingsById provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ConfigurationSettingsById(id string)(*ConfigurationSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewConfigurationSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
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
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DataSharingConsents provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsents()(*DataSharingConsentsRequestBuilder) {
    return NewDataSharingConsentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DataSharingConsentsById provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DataSharingConsentsById(id string)(*DataSharingConsentsDataSharingConsentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDataSharingConsentsDataSharingConsentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DepOnboardingSettings provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettings()(*DepOnboardingSettingsRequestBuilder) {
    return NewDepOnboardingSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DepOnboardingSettingsById provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DepOnboardingSettingsById(id string)(*DepOnboardingSettingsDepOnboardingSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDepOnboardingSettingsDepOnboardingSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DerivedCredentials provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentials()(*DerivedCredentialsRequestBuilder) {
    return NewDerivedCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DerivedCredentialsById provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DerivedCredentialsById(id string)(*DerivedCredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDerivedCredentialsDeviceManagementDerivedCredentialSettingsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DetectedAppsRequestBuilder) {
    return NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DetectedAppsById provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DetectedAppsById(id string)(*DetectedAppsDetectedAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDetectedAppsDetectedAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DeviceCategoriesRequestBuilder) {
    return NewDeviceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCategoriesById provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCategoriesById(id string)(*DeviceCategoriesDeviceCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceCategoriesDeviceCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DeviceCompliancePoliciesRequestBuilder) {
    return NewDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePoliciesById provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePoliciesById(id string)(*DeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceCompliancePoliciesDeviceCompliancePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCompliancePolicySettingStateSummariesById provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*DeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceComplianceScripts provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScripts()(*DeviceComplianceScriptsRequestBuilder) {
    return NewDeviceComplianceScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceComplianceScriptsById provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceComplianceScriptsById(id string)(*DeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceComplianceScriptsDeviceComplianceScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceConfigurationConflictSummary provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceConfigurationConflictSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationConflictSummaryById provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummaryById(id string)(*DeviceConfigurationConflictSummaryDeviceConfigurationConflictSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationConflictSummaryDeviceConfigurationConflictSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationRestrictedAppsViolations provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolations()(*DeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationRestrictedAppsViolationsById provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolationsById(id string)(*DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceConfigurationsRequestBuilder) {
    return NewDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationsAllManagedDeviceCertificateStates provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStates()(*DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesById provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStatesById(id string)(*DeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesManagedAllDeviceCertificateStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceConfigurationsById provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsById(id string)(*DeviceConfigurationsDeviceConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceConfigurationsDeviceConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceConfigurationUserStateSummaries provides operations to manage the deviceConfigurationUserStateSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceConfigurationUserStateSummaries()(*DeviceConfigurationUserStateSummariesRequestBuilder) {
    return NewDeviceConfigurationUserStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCustomAttributeShellScripts provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScripts()(*DeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceCustomAttributeShellScriptsById provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScriptsById(id string)(*DeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceCustomAttributeShellScriptsDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceEnrollmentConfigurationsById provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*DeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceHealthScripts provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScripts()(*DeviceHealthScriptsRequestBuilder) {
    return NewDeviceHealthScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceHealthScriptsById provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceHealthScriptsById(id string)(*DeviceHealthScriptsDeviceHealthScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceHealthScriptsDeviceHealthScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DeviceManagementPartnersRequestBuilder) {
    return NewDeviceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceManagementPartnersById provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementPartnersById(id string)(*DeviceManagementPartnersDeviceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceManagementPartnersDeviceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceManagementScripts provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScripts()(*DeviceManagementScriptsRequestBuilder) {
    return NewDeviceManagementScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceManagementScriptsById provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceManagementScriptsById(id string)(*DeviceManagementScriptsDeviceManagementScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceManagementScriptsDeviceManagementScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DeviceShellScripts provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScripts()(*DeviceShellScriptsRequestBuilder) {
    return NewDeviceShellScriptsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DeviceShellScriptsById provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DeviceShellScriptsById(id string)(*DeviceShellScriptsDeviceShellScriptItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDeviceShellScriptsDeviceShellScriptItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DomainJoinConnectors provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectors()(*DomainJoinConnectorsRequestBuilder) {
    return NewDomainJoinConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DomainJoinConnectorsById provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) DomainJoinConnectorsById(id string)(*DomainJoinConnectorsDeviceManagementDomainJoinConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewDomainJoinConnectorsDeviceManagementDomainJoinConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// EmbeddedSIMActivationCodePools provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePools()(*EmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// EmbeddedSIMActivationCodePoolsById provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePoolsById(id string)(*EmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewEmbeddedSIMActivationCodePoolsEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ExchangeConnectorsRequestBuilder) {
    return NewExchangeConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExchangeConnectorsById provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeConnectorsById(id string)(*ExchangeConnectorsDeviceManagementExchangeConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewExchangeConnectorsDeviceManagementExchangeConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ExchangeOnPremisesPolicies provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicies()(*ExchangeOnPremisesPoliciesRequestBuilder) {
    return NewExchangeOnPremisesPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ExchangeOnPremisesPoliciesById provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPoliciesById(id string)(*ExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewExchangeOnPremisesPoliciesDeviceManagementExchangeOnPremisesPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ExchangeOnPremisesPolicy provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicy()(*ExchangeOnPremisesPolicyRequestBuilder) {
    return NewExchangeOnPremisesPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get deviceManagement
func (m *DeviceManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable), nil
}
// GroupPolicyCategories provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategories()(*GroupPolicyCategoriesRequestBuilder) {
    return NewGroupPolicyCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyCategoriesById provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyCategoriesById(id string)(*GroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyCategoriesGroupPolicyCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyConfigurations provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurations()(*GroupPolicyConfigurationsRequestBuilder) {
    return NewGroupPolicyConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyConfigurationsById provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurationsById(id string)(*GroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyDefinitionFiles provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFiles()(*GroupPolicyDefinitionFilesRequestBuilder) {
    return NewGroupPolicyDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyDefinitionFilesById provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFilesById(id string)(*GroupPolicyDefinitionFilesGroupPolicyDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyDefinitionFilesGroupPolicyDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyDefinitions provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitions()(*GroupPolicyDefinitionsRequestBuilder) {
    return NewGroupPolicyDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyDefinitionsById provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionsById(id string)(*GroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyMigrationReports provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReports()(*GroupPolicyMigrationReportsRequestBuilder) {
    return NewGroupPolicyMigrationReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyMigrationReportsById provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReportsById(id string)(*GroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyMigrationReportsGroupPolicyMigrationReportItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyObjectFiles provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFiles()(*GroupPolicyObjectFilesRequestBuilder) {
    return NewGroupPolicyObjectFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyObjectFilesById provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFilesById(id string)(*GroupPolicyObjectFilesGroupPolicyObjectFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyObjectFilesGroupPolicyObjectFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// GroupPolicyUploadedDefinitionFiles provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFiles()(*GroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// GroupPolicyUploadedDefinitionFilesById provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFilesById(id string)(*GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ImportedDeviceIdentities provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentities()(*ImportedDeviceIdentitiesRequestBuilder) {
    return NewImportedDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ImportedDeviceIdentitiesById provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentitiesById(id string)(*ImportedDeviceIdentitiesImportedDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewImportedDeviceIdentitiesImportedDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ImportedWindowsAutopilotDeviceIdentitiesById provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentitiesById(id string)(*ImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewImportedWindowsAutopilotDeviceIdentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Intents provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Intents()(*IntentsRequestBuilder) {
    return NewIntentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IntentsById provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntentsById(id string)(*IntentsDeviceManagementIntentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntentsDeviceManagementIntentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// IntuneBrandingProfiles provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfiles()(*IntuneBrandingProfilesRequestBuilder) {
    return NewIntuneBrandingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IntuneBrandingProfilesById provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfilesById(id string)(*IntuneBrandingProfilesIntuneBrandingProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIntuneBrandingProfilesIntuneBrandingProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*IosUpdateStatusesRequestBuilder) {
    return NewIosUpdateStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// IosUpdateStatusesById provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) IosUpdateStatusesById(id string)(*IosUpdateStatusesIosUpdateDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewIosUpdateStatusesIosUpdateDeviceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MacOSSoftwareUpdateAccountSummaries provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummaries()(*MacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    return NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MacOSSoftwareUpdateAccountSummariesById provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummariesById(id string)(*MacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMacOSSoftwareUpdateAccountSummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedDeviceEncryptionStates provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStates()(*ManagedDeviceEncryptionStatesRequestBuilder) {
    return NewManagedDeviceEncryptionStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDeviceEncryptionStatesById provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStatesById(id string)(*ManagedDeviceEncryptionStatesManagedDeviceEncryptionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedDeviceEncryptionStatesManagedDeviceEncryptionStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*ManagedDeviceOverviewRequestBuilder) {
    return NewManagedDeviceOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*ManagedDevicesRequestBuilder) {
    return NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ManagedDevicesById provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ManagedDevicesById(id string)(*ManagedDevicesManagedDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewManagedDevicesManagedDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftGraphEnableAndroidDeviceAdministratorEnrollment provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphEnableAndroidDeviceAdministratorEnrollment()(*MicrosoftGraphEnableAndroidDeviceAdministratorEnrollmentEnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    return NewMicrosoftGraphEnableAndroidDeviceAdministratorEnrollmentEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableLegacyPcManagement provides operations to call the enableLegacyPcManagement method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphEnableLegacyPcManagement()(*MicrosoftGraphEnableLegacyPcManagementEnableLegacyPcManagementRequestBuilder) {
    return NewMicrosoftGraphEnableLegacyPcManagementEnableLegacyPcManagementRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEnableUnlicensedAdminstrators provides operations to call the enableUnlicensedAdminstrators method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphEnableUnlicensedAdminstrators()(*MicrosoftGraphEnableUnlicensedAdminstratorsEnableUnlicensedAdminstratorsRequestBuilder) {
    return NewMicrosoftGraphEnableUnlicensedAdminstratorsEnableUnlicensedAdminstratorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEvaluateAssignmentFilter provides operations to call the evaluateAssignmentFilter method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphEvaluateAssignmentFilter()(*MicrosoftGraphEvaluateAssignmentFilterEvaluateAssignmentFilterRequestBuilder) {
    return NewMicrosoftGraphEvaluateAssignmentFilterEvaluateAssignmentFilterRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAssignedRoleDetails provides operations to call the getAssignedRoleDetails method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetAssignedRoleDetails()(*MicrosoftGraphGetAssignedRoleDetailsGetAssignedRoleDetailsRequestBuilder) {
    return NewMicrosoftGraphGetAssignedRoleDetailsGetAssignedRoleDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetAssignmentFiltersStatusDetails provides operations to call the getAssignmentFiltersStatusDetails method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetAssignmentFiltersStatusDetails()(*MicrosoftGraphGetAssignmentFiltersStatusDetailsGetAssignmentFiltersStatusDetailsRequestBuilder) {
    return NewMicrosoftGraphGetAssignmentFiltersStatusDetailsGetAssignmentFiltersStatusDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComanagedDevicesSummary provides operations to call the getComanagedDevicesSummary method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetComanagedDevicesSummary()(*MicrosoftGraphGetComanagedDevicesSummaryGetComanagedDevicesSummaryRequestBuilder) {
    return NewMicrosoftGraphGetComanagedDevicesSummaryGetComanagedDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetComanagementEligibleDevicesSummary provides operations to call the getComanagementEligibleDevicesSummary method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetComanagementEligibleDevicesSummary()(*MicrosoftGraphGetComanagementEligibleDevicesSummaryGetComanagementEligibleDevicesSummaryRequestBuilder) {
    return NewMicrosoftGraphGetComanagementEligibleDevicesSummaryGetComanagementEligibleDevicesSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEffectivePermissions provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetEffectivePermissions()(*MicrosoftGraphGetEffectivePermissionsGetEffectivePermissionsRequestBuilder) {
    return NewMicrosoftGraphGetEffectivePermissionsGetEffectivePermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetEffectivePermissionsWithScope(scope *string)(*MicrosoftGraphGetEffectivePermissionsWithScopeGetEffectivePermissionsWithScopeRequestBuilder) {
    return NewMicrosoftGraphGetEffectivePermissionsWithScopeGetEffectivePermissionsWithScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter, scope)
}
// MicrosoftGraphGetRoleScopeTagsByIdsWithIds provides operations to call the getRoleScopeTagsByIds method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetRoleScopeTagsByIdsWithIds(ids *string)(*MicrosoftGraphGetRoleScopeTagsByIdsWithIdsGetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewMicrosoftGraphGetRoleScopeTagsByIdsWithIdsGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter, ids)
}
// MicrosoftGraphGetRoleScopeTagsByResourceWithResource provides operations to call the getRoleScopeTagsByResource method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetRoleScopeTagsByResourceWithResource(resource *string)(*MicrosoftGraphGetRoleScopeTagsByResourceWithResourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewMicrosoftGraphGetRoleScopeTagsByResourceWithResourceGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource)
}
// MicrosoftGraphGetSuggestedEnrollmentLimitWithEnrollmentType provides operations to call the getSuggestedEnrollmentLimit method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphGetSuggestedEnrollmentLimitWithEnrollmentType(enrollmentType *string)(*MicrosoftGraphGetSuggestedEnrollmentLimitWithEnrollmentTypeGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    return NewMicrosoftGraphGetSuggestedEnrollmentLimitWithEnrollmentTypeGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter, enrollmentType)
}
// MicrosoftGraphScopedForResourceWithResource provides operations to call the scopedForResource method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphScopedForResourceWithResource(resource *string)(*MicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilder) {
    return NewMicrosoftGraphScopedForResourceWithResourceScopedForResourceWithResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter, resource)
}
// MicrosoftGraphSendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphSendCustomNotificationToCompanyPortal()(*MicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewMicrosoftGraphSendCustomNotificationToCompanyPortalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUserExperienceAnalyticsSummarizedDeviceScopes provides operations to call the userExperienceAnalyticsSummarizedDeviceScopes method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphUserExperienceAnalyticsSummarizedDeviceScopes()(*MicrosoftGraphUserExperienceAnalyticsSummarizedDeviceScopesUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder) {
    return NewMicrosoftGraphUserExperienceAnalyticsSummarizedDeviceScopesUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphUserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphUserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*MicrosoftGraphUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewMicrosoftGraphUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
func (m *DeviceManagementRequestBuilder) MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*MicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewMicrosoftGraphVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, domainName)
}
// MicrosoftTunnelConfigurations provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurations()(*MicrosoftTunnelConfigurationsRequestBuilder) {
    return NewMicrosoftTunnelConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftTunnelConfigurationsById provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurationsById(id string)(*MicrosoftTunnelConfigurationsMicrosoftTunnelConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMicrosoftTunnelConfigurationsMicrosoftTunnelConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftTunnelHealthThresholds provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholds()(*MicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewMicrosoftTunnelHealthThresholdsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftTunnelHealthThresholdsById provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholdsById(id string)(*MicrosoftTunnelHealthThresholdsMicrosoftTunnelHealthThresholdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMicrosoftTunnelHealthThresholdsMicrosoftTunnelHealthThresholdItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftTunnelServerLogCollectionResponses provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponses()(*MicrosoftTunnelServerLogCollectionResponsesRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftTunnelServerLogCollectionResponsesById provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponsesById(id string)(*MicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMicrosoftTunnelServerLogCollectionResponsesMicrosoftTunnelServerLogCollectionResponseItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MicrosoftTunnelSites provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSites()(*MicrosoftTunnelSitesRequestBuilder) {
    return NewMicrosoftTunnelSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftTunnelSitesById provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSitesById(id string)(*MicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*MobileAppTroubleshootingEventsRequestBuilder) {
    return NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileAppTroubleshootingEventsById provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEventsById(id string)(*MobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMobileAppTroubleshootingEventsMobileAppTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*MobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobileThreatDefenseConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MobileThreatDefenseConnectorsById provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectorsById(id string)(*MobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewMobileThreatDefenseConnectorsMobileThreatDefenseConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Monitoring provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Monitoring()(*MonitoringRequestBuilder) {
    return NewMonitoringRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NdesConnectors provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectors()(*NdesConnectorsRequestBuilder) {
    return NewNdesConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NdesConnectorsById provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NdesConnectorsById(id string)(*NdesConnectorsNdesConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewNdesConnectorsNdesConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*NotificationMessageTemplatesRequestBuilder) {
    return NewNotificationMessageTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NotificationMessageTemplatesById provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplatesById(id string)(*NotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewNotificationMessageTemplatesNotificationMessageTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// OemWarrantyInformationOnboarding provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboardingRequestBuilder) {
    return NewOemWarrantyInformationOnboardingRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OemWarrantyInformationOnboardingById provides operations to manage the oemWarrantyInformationOnboarding property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) OemWarrantyInformationOnboardingById(id string)(*OemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewOemWarrantyInformationOnboardingOemWarrantyInformationOnboardingItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Patch update deviceManagement
func (m *DeviceManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
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
    return NewRemoteActionAuditsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoteActionAuditsById provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteActionAuditsById(id string)(*RemoteActionAuditsRemoteActionAuditItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewRemoteActionAuditsRemoteActionAuditItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*RemoteAssistancePartnersRequestBuilder) {
    return NewRemoteAssistancePartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RemoteAssistancePartnersById provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartnersById(id string)(*RemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewRemoteAssistancePartnersRemoteAssistancePartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// RemoteAssistanceSettings provides operations to manage the remoteAssistanceSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RemoteAssistanceSettings()(*RemoteAssistanceSettingsRequestBuilder) {
    return NewRemoteAssistanceSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourceAccessProfiles provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfiles()(*ResourceAccessProfilesRequestBuilder) {
    return NewResourceAccessProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourceAccessProfilesById provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceAccessProfilesById(id string)(*ResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*ResourceOperationsRequestBuilder) {
    return NewResourceOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ResourceOperationsById provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ResourceOperationsById(id string)(*ResourceOperationsResourceOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewResourceOperationsResourceOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ReusablePolicySettings provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettings()(*ReusablePolicySettingsRequestBuilder) {
    return NewReusablePolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReusablePolicySettingsById provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusablePolicySettingsById(id string)(*ReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ReusableSettings provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettings()(*ReusableSettingsRequestBuilder) {
    return NewReusableSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ReusableSettingsById provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ReusableSettingsById(id string)(*ReusableSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewReusableSettingsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*RoleAssignmentsRequestBuilder) {
    return NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleAssignmentsById provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*RoleAssignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewRoleAssignmentsDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*RoleDefinitionsRequestBuilder) {
    return NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleDefinitionsById provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*RoleDefinitionsRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewRoleDefinitionsRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTags()(*RoleScopeTagsRequestBuilder) {
    return NewRoleScopeTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// RoleScopeTagsById provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) RoleScopeTagsById(id string)(*RoleScopeTagsRoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewRoleScopeTagsRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ServiceNowConnections provides operations to manage the serviceNowConnections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ServiceNowConnections()(*ServiceNowConnectionsRequestBuilder) {
    return NewServiceNowConnectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ServiceNowConnectionsById provides operations to manage the serviceNowConnections property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ServiceNowConnectionsById(id string)(*ServiceNowConnectionsServiceNowConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewServiceNowConnectionsServiceNowConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitions()(*SettingDefinitionsRequestBuilder) {
    return NewSettingDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// SettingDefinitionsById provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SettingDefinitionsById(id string)(*SettingDefinitionsDeviceManagementSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewSettingDefinitionsDeviceManagementSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareUpdateStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*TelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomExpenseManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TelecomExpenseManagementPartnersById provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartnersById(id string)(*TelecomExpenseManagementPartnersTelecomExpenseManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTelecomExpenseManagementPartnersTelecomExpenseManagementPartnerItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Templates provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemplatesById provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplatesById(id string)(*TemplatesDeviceManagementTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTemplatesDeviceManagementTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// TemplateSettings provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettings()(*TemplateSettingsRequestBuilder) {
    return NewTemplateSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TemplateSettingsById provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TemplateSettingsById(id string)(*TemplateSettingsDeviceManagementConfigurationSettingTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTemplateSettingsDeviceManagementConfigurationSettingTemplateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// TenantAttachRBAC provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TenantAttachRBAC()(*TenantAttachRBACRequestBuilder) {
    return NewTenantAttachRBACRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*TermsAndConditionsRequestBuilder) {
    return NewTermsAndConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TermsAndConditionsById provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TermsAndConditionsById(id string)(*TermsAndConditionsTermsAndConditionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTermsAndConditionsTermsAndConditionsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ToGetRequestInformation get deviceManagement
func (m *DeviceManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update deviceManagement
func (m *DeviceManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TroubleshootingEvents provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*TroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// TroubleshootingEventsById provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) TroubleshootingEventsById(id string)(*TroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewTroubleshootingEventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAnomaly provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomaly()(*UserExperienceAnalyticsAnomalyRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAnomalyById provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyById(id string)(*UserExperienceAnalyticsAnomalyUserExperienceAnalyticsAnomalyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAnomalyUserExperienceAnalyticsAnomalyItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAnomalyDevice provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDevice()(*UserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAnomalyDeviceById provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDeviceById(id string)(*UserExperienceAnalyticsAnomalyDeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAnomalyDeviceUserExperienceAnalyticsAnomalyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceUserExperienceAnalyticsAppHealthApplicationPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionById(id string)(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformanceById(id string)(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*UserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceById(id string)(*UserExperienceAnalyticsAppHealthDevicePerformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceUserExperienceAnalyticsAppHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetailsById(id string)(*UserExperienceAnalyticsAppHealthDevicePerformanceDetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsAppHealthOSVersionPerformanceById provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformanceById(id string)(*UserExperienceAnalyticsAppHealthOSVersionPerformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*UserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*UserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBaselinesById provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselinesById(id string)(*UserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBaselinesUserExperienceAnalyticsBaselineItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpact()(*UserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpactById(id string)(*UserExperienceAnalyticsBatteryHealthAppImpactUserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthAppImpactUserExperienceAnalyticsBatteryHealthAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthCapacityDetails provides operations to manage the userExperienceAnalyticsBatteryHealthCapacityDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpactById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpactById(id string)(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthDevicePerformance provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDevicePerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthDevicePerformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryById(id string)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthModelPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthModelPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthModelPerformanceUserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthModelPerformanceUserExperienceAnalyticsBatteryHealthModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthOsPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsBatteryHealthOsPerformanceById provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformanceById(id string)(*UserExperienceAnalyticsBatteryHealthOsPerformanceUserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsBatteryHealthOsPerformanceUserExperienceAnalyticsBatteryHealthOsPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsBatteryHealthRuntimeDetails provides operations to manage the userExperienceAnalyticsBatteryHealthRuntimeDetails property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*UserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsCategoriesById provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategoriesById(id string)(*UserExperienceAnalyticsCategoriesUserExperienceAnalyticsCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsCategoriesUserExperienceAnalyticsCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceMetricHistory provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistory()(*UserExperienceAnalyticsDeviceMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceMetricHistoryById provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistoryById(id string)(*UserExperienceAnalyticsDeviceMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDevicePerformanceById provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformanceById(id string)(*UserExperienceAnalyticsDevicePerformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDevicePerformanceUserExperienceAnalyticsDevicePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceScope provides operations to manage the userExperienceAnalyticsDeviceScope property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScope()(*UserExperienceAnalyticsDeviceScopeRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceScopes provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopes()(*UserExperienceAnalyticsDeviceScopesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceScopesById provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopesById(id string)(*UserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceScopesUserExperienceAnalyticsDeviceScopeItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*UserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceScoresById provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScoresById(id string)(*UserExperienceAnalyticsDeviceScoresUserExperienceAnalyticsDeviceScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceScoresUserExperienceAnalyticsDeviceScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceStartupHistoryById provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistoryById(id string)(*UserExperienceAnalyticsDeviceStartupHistoryUserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceStartupHistoryUserExperienceAnalyticsDeviceStartupHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*UserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcessesById provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessesById(id string)(*UserExperienceAnalyticsDeviceStartupProcessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceStartupProcessesUserExperienceAnalyticsDeviceStartupProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcessPerformanceById provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformanceById(id string)(*UserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentity provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentity()(*UserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentityById provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentityById(id string)(*UserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDevicesWithoutCloudIdentityUserExperienceAnalyticsDeviceWithoutCloudIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsDeviceTimelineEvent provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceTimelineEvent()(*UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsDeviceTimelineEventById provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceTimelineEventById(id string)(*UserExperienceAnalyticsDeviceTimelineEventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsDeviceTimelineEventUserExperienceAnalyticsDeviceTimelineEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsImpactingProcess provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsImpactingProcessById provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcessById(id string)(*UserExperienceAnalyticsImpactingProcessUserExperienceAnalyticsImpactingProcessItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsImpactingProcessUserExperienceAnalyticsImpactingProcessItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsMetricHistoryById provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistoryById(id string)(*UserExperienceAnalyticsMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsMetricHistoryUserExperienceAnalyticsMetricHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsModelScoresById provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScoresById(id string)(*UserExperienceAnalyticsModelScoresUserExperienceAnalyticsModelScoresItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsModelScoresUserExperienceAnalyticsModelScoresItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsNotAutopilotReadyDevice provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceById provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDeviceById(id string)(*UserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceUserExperienceAnalyticsNotAutopilotReadyDeviceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsRemoteConnection provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsRemoteConnectionById provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnectionById(id string)(*UserExperienceAnalyticsRemoteConnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsRemoteConnectionUserExperienceAnalyticsRemoteConnectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsResourcePerformance provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsResourcePerformanceById provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformanceById(id string)(*UserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsResourcePerformanceUserExperienceAnalyticsResourcePerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*UserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsScoreHistoryById provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistoryById(id string)(*UserExperienceAnalyticsScoreHistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsScoreHistoryUserExperienceAnalyticsScoreHistoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereMetricsById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetricsById(id string)(*UserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformanceById(id string)(*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// UserPfxCertificates provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificates()(*UserPfxCertificatesRequestBuilder) {
    return NewUserPfxCertificatesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// UserPfxCertificatesById provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) UserPfxCertificatesById(id string)(*UserPfxCertificatesUserPFXCertificateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewUserPfxCertificatesUserPFXCertificateItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*VirtualEndpointRequestBuilder) {
    return NewVirtualEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsAutopilotDeploymentProfiles provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfiles()(*WindowsAutopilotDeploymentProfilesRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsAutopilotDeploymentProfilesById provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfilesById(id string)(*WindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsAutopilotDeploymentProfilesWindowsAutopilotDeploymentProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*WindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsAutopilotDeviceIdentitiesById provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentitiesById(id string)(*WindowsAutopilotDeviceIdentitiesWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsAutopilotDeviceIdentitiesWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsAutopilotSettings provides operations to manage the windowsAutopilotSettings property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsAutopilotSettings()(*WindowsAutopilotSettingsRequestBuilder) {
    return NewWindowsAutopilotSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsDriverUpdateProfiles provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfiles()(*WindowsDriverUpdateProfilesRequestBuilder) {
    return NewWindowsDriverUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsDriverUpdateProfilesById provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfilesById(id string)(*WindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsDriverUpdateProfilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsFeatureUpdateProfiles provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfiles()(*WindowsFeatureUpdateProfilesRequestBuilder) {
    return NewWindowsFeatureUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsFeatureUpdateProfilesById provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfilesById(id string)(*WindowsFeatureUpdateProfilesWindowsFeatureUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsFeatureUpdateProfilesWindowsFeatureUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*WindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsInformationProtectionAppLearningSummariesById provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummariesById(id string)(*WindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsInformationProtectionAppLearningSummariesWindowsInformationProtectionAppLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*WindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsInformationProtectionNetworkLearningSummariesById provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummariesById(id string)(*WindowsInformationProtectionNetworkLearningSummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsInformationProtectionNetworkLearningSummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*WindowsMalwareInformationRequestBuilder) {
    return NewWindowsMalwareInformationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsMalwareInformationById provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformationById(id string)(*WindowsMalwareInformationWindowsMalwareInformationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsMalwareInformationWindowsMalwareInformationItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsQualityUpdateProfiles provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfiles()(*WindowsQualityUpdateProfilesRequestBuilder) {
    return NewWindowsQualityUpdateProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsQualityUpdateProfilesById provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfilesById(id string)(*WindowsQualityUpdateProfilesWindowsQualityUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsQualityUpdateProfilesWindowsQualityUpdateProfileItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// WindowsUpdateCatalogItems provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItems()(*WindowsUpdateCatalogItemsRequestBuilder) {
    return NewWindowsUpdateCatalogItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// WindowsUpdateCatalogItemsById provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItemsById(id string)(*WindowsUpdateCatalogItemsWindowsUpdateCatalogItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewWindowsUpdateCatalogItemsWindowsUpdateCatalogItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ZebraFotaArtifacts provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifacts()(*ZebraFotaArtifactsRequestBuilder) {
    return NewZebraFotaArtifactsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ZebraFotaArtifactsById provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifactsById(id string)(*ZebraFotaArtifactsZebraFotaArtifactItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewZebraFotaArtifactsZebraFotaArtifactItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// ZebraFotaConnector provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaConnector()(*ZebraFotaConnectorRequestBuilder) {
    return NewZebraFotaConnectorRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ZebraFotaDeployments provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeployments()(*ZebraFotaDeploymentsRequestBuilder) {
    return NewZebraFotaDeploymentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ZebraFotaDeploymentsById provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
func (m *DeviceManagementRequestBuilder) ZebraFotaDeploymentsById(id string)(*ZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewZebraFotaDeploymentsZebraFotaDeploymentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
