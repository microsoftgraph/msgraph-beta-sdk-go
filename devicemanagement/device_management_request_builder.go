package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementRequestBuilder provides operations to manage the deviceManagement singleton.
type DeviceManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
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
// returns a *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidDeviceOwnerEnrollmentProfiles provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidDeviceOwnerEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfiles()(*AndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkAppConfigurationSchemas provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidForWorkAppConfigurationSchemasRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemas()(*AndroidForWorkAppConfigurationSchemasRequestBuilder) {
    return NewAndroidForWorkAppConfigurationSchemasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkEnrollmentProfiles provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidForWorkEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfiles()(*AndroidForWorkEnrollmentProfilesRequestBuilder) {
    return NewAndroidForWorkEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkSettings provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidForWorkSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkSettings()(*AndroidForWorkSettingsRequestBuilder) {
    return NewAndroidForWorkSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidManagedStoreAccountEnterpriseSettings provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidManagedStoreAppConfigurationSchemas provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidManagedStoreAppConfigurationSchemasRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemas()(*AndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    return NewAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
// returns a *ApplePushNotificationCertificateRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*ApplePushNotificationCertificateRequestBuilder) {
    return NewApplePushNotificationCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppleUserInitiatedEnrollmentProfiles provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AppleUserInitiatedEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfiles()(*AppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilters provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
// returns a *AssignmentFiltersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AssignmentFilters()(*AssignmentFiltersRequestBuilder) {
    return NewAssignmentFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AuditEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AuditEvents()(*AuditEventsRequestBuilder) {
    return NewAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AutopilotEvents provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AutopilotEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AutopilotEvents()(*AutopilotEventsRequestBuilder) {
    return NewAutopilotEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CartToClassAssociations provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
// returns a *CartToClassAssociationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CartToClassAssociations()(*CartToClassAssociationsRequestBuilder) {
    return NewCartToClassAssociationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
// returns a *CategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Categories()(*CategoriesRequestBuilder) {
    return NewCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CertificateConnectorDetails provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
// returns a *CertificateConnectorDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetails()(*CertificateConnectorDetailsRequestBuilder) {
    return NewCertificateConnectorDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChromeOSOnboardingSettings provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ChromeOSOnboardingSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettings()(*ChromeOSOnboardingSettingsRequestBuilder) {
    return NewChromeOSOnboardingSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPCConnectivityIssues provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
// returns a *CloudPCConnectivityIssuesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssues()(*CloudPCConnectivityIssuesRequestBuilder) {
    return NewCloudPCConnectivityIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComanagedDevices provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ComanagedDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComanagedDevices()(*ComanagedDevicesRequestBuilder) {
    return NewComanagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComanagementEligibleDevices provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ComanagementEligibleDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevices()(*ComanagementEligibleDevicesRequestBuilder) {
    return NewComanagementEligibleDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceCategories provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
// returns a *ComplianceCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceCategories()(*ComplianceCategoriesRequestBuilder) {
    return NewComplianceCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *ComplianceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*ComplianceManagementPartnersRequestBuilder) {
    return NewComplianceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompliancePolicies provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CompliancePolicies()(*CompliancePoliciesRequestBuilder) {
    return NewCompliancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceSettings provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ComplianceSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceSettings()(*ComplianceSettingsRequestBuilder) {
    return NewComplianceSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ConditionalAccessSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*ConditionalAccessSettingsRequestBuilder) {
    return NewConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigManagerCollections provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigManagerCollectionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigManagerCollections()(*ConfigManagerCollectionsRequestBuilder) {
    return NewConfigManagerCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationCategories provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationCategories()(*ConfigurationCategoriesRequestBuilder) {
    return NewConfigurationCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationPolicies provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationPolicies()(*ConfigurationPoliciesRequestBuilder) {
    return NewConfigurationPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationPolicyTemplates provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationPolicyTemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplates()(*ConfigurationPolicyTemplatesRequestBuilder) {
    return NewConfigurationPolicyTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationSettings provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationSettings()(*ConfigurationSettingsRequestBuilder) {
    return NewConfigurationSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DataSharingConsents provides operations to manage the dataSharingConsents property of the microsoft.graph.deviceManagement entity.
// returns a *DataSharingConsentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DataSharingConsents()(*DataSharingConsentsRequestBuilder) {
    return NewDataSharingConsentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DepOnboardingSettings provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
// returns a *DepOnboardingSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DepOnboardingSettings()(*DepOnboardingSettingsRequestBuilder) {
    return NewDepOnboardingSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DerivedCredentials provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
// returns a *DerivedCredentialsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DerivedCredentials()(*DerivedCredentialsRequestBuilder) {
    return NewDerivedCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
// returns a *DetectedAppsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DetectedAppsRequestBuilder) {
    return NewDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DeviceCategoriesRequestBuilder) {
    return NewDeviceCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceCompliancePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DeviceCompliancePoliciesRequestBuilder) {
    return NewDeviceCompliancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceCompliancePolicyDeviceStateSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceCompliancePolicySettingStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceComplianceScripts provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceComplianceScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceComplianceScripts()(*DeviceComplianceScriptsRequestBuilder) {
    return NewDeviceComplianceScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationConflictSummary provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationConflictSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummary()(*DeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceConfigurationConflictSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationDeviceStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationProfiles provides operations to manage the deviceConfigurationProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationProfiles()(*DeviceConfigurationProfilesRequestBuilder) {
    return NewDeviceConfigurationProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationRestrictedAppsViolations provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationRestrictedAppsViolationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolations()(*DeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceConfigurationsRequestBuilder) {
    return NewDeviceConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationsAllManagedDeviceCertificateStates provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStates()(*DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationUserStateSummaries provides operations to manage the deviceConfigurationUserStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceConfigurationUserStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationUserStateSummaries()(*DeviceConfigurationUserStateSummariesRequestBuilder) {
    return NewDeviceConfigurationUserStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCustomAttributeShellScripts provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceCustomAttributeShellScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScripts()(*DeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDeviceCustomAttributeShellScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScripts provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceHealthScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceHealthScripts()(*DeviceHealthScriptsRequestBuilder) {
    return NewDeviceHealthScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DeviceManagementPartnersRequestBuilder) {
    return NewDeviceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementScripts provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceManagementScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceManagementScripts()(*DeviceManagementScriptsRequestBuilder) {
    return NewDeviceManagementScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceShellScripts provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceShellScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceShellScripts()(*DeviceShellScriptsRequestBuilder) {
    return NewDeviceShellScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DomainJoinConnectors provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *DomainJoinConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DomainJoinConnectors()(*DomainJoinConnectorsRequestBuilder) {
    return NewDomainJoinConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ElevationRequests provides operations to manage the elevationRequests property of the microsoft.graph.deviceManagement entity.
// returns a *ElevationRequestsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ElevationRequests()(*ElevationRequestsRequestBuilder) {
    return NewElevationRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmbeddedSIMActivationCodePools provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
// returns a *EmbeddedSIMActivationCodePoolsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePools()(*EmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewEmbeddedSIMActivationCodePoolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableAndroidDeviceAdministratorEnrollment provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
// returns a *EnableAndroidDeviceAdministratorEnrollmentRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableAndroidDeviceAdministratorEnrollment()(*EnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    return NewEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableEndpointPrivilegeManagement provides operations to call the enableEndpointPrivilegeManagement method.
// returns a *EnableEndpointPrivilegeManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableEndpointPrivilegeManagement()(*EnableEndpointPrivilegeManagementRequestBuilder) {
    return NewEnableEndpointPrivilegeManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLegacyPcManagement provides operations to call the enableLegacyPcManagement method.
// returns a *EnableLegacyPcManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableLegacyPcManagement()(*EnableLegacyPcManagementRequestBuilder) {
    return NewEnableLegacyPcManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableUnlicensedAdminstrators provides operations to call the enableUnlicensedAdminstrators method.
// returns a *EnableUnlicensedAdminstratorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableUnlicensedAdminstrators()(*EnableUnlicensedAdminstratorsRequestBuilder) {
    return NewEnableUnlicensedAdminstratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EndpointPrivilegeManagementProvisioningStatus provides operations to manage the endpointPrivilegeManagementProvisioningStatus property of the microsoft.graph.deviceManagement entity.
// returns a *EndpointPrivilegeManagementProvisioningStatusRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EndpointPrivilegeManagementProvisioningStatus()(*EndpointPrivilegeManagementProvisioningStatusRequestBuilder) {
    return NewEndpointPrivilegeManagementProvisioningStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateAssignmentFilter provides operations to call the evaluateAssignmentFilter method.
// returns a *EvaluateAssignmentFilterRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EvaluateAssignmentFilter()(*EvaluateAssignmentFilterRequestBuilder) {
    return NewEvaluateAssignmentFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ExchangeConnectorsRequestBuilder) {
    return NewExchangeConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeOnPremisesPolicies provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeOnPremisesPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicies()(*ExchangeOnPremisesPoliciesRequestBuilder) {
    return NewExchangeOnPremisesPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeOnPremisesPolicy provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeOnPremisesPolicyRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicy()(*ExchangeOnPremisesPolicyRequestBuilder) {
    return NewExchangeOnPremisesPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get deviceManagement
// returns a DeviceManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable), nil
}
// GetAssignedRoleDetails provides operations to call the getAssignedRoleDetails method.
// returns a *GetAssignedRoleDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetAssignedRoleDetails()(*GetAssignedRoleDetailsRequestBuilder) {
    return NewGetAssignedRoleDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAssignmentFiltersStatusDetails provides operations to call the getAssignmentFiltersStatusDetails method.
// returns a *GetAssignmentFiltersStatusDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetAssignmentFiltersStatusDetails()(*GetAssignmentFiltersStatusDetailsRequestBuilder) {
    return NewGetAssignmentFiltersStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComanagedDevicesSummary provides operations to call the getComanagedDevicesSummary method.
// returns a *GetComanagedDevicesSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetComanagedDevicesSummary()(*GetComanagedDevicesSummaryRequestBuilder) {
    return NewGetComanagedDevicesSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComanagementEligibleDevicesSummary provides operations to call the getComanagementEligibleDevicesSummary method.
// returns a *GetComanagementEligibleDevicesSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetComanagementEligibleDevicesSummary()(*GetComanagementEligibleDevicesSummaryRequestBuilder) {
    return NewGetComanagementEligibleDevicesSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
// returns a *GetEffectivePermissionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetEffectivePermissions()(*GetEffectivePermissionsRequestBuilder) {
    return NewGetEffectivePermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
// returns a *GetEffectivePermissionsWithScopeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*GetEffectivePermissionsWithScopeRequestBuilder) {
    return NewGetEffectivePermissionsWithScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, scope)
}
// GetRoleScopeTagsByIdsWithIds provides operations to call the getRoleScopeTagsByIds method.
// returns a *GetRoleScopeTagsByIdsWithIdsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByIdsWithIds(ids *string)(*GetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, ids)
}
// GetRoleScopeTagsByResourceWithResource provides operations to call the getRoleScopeTagsByResource method.
// returns a *GetRoleScopeTagsByResourceWithResourceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByResourceWithResource(resource *string)(*GetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, resource)
}
// GetSuggestedEnrollmentLimitWithEnrollmentType provides operations to call the getSuggestedEnrollmentLimit method.
// returns a *GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetSuggestedEnrollmentLimitWithEnrollmentType(enrollmentType *string)(*GetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    return NewGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, enrollmentType)
}
// GroupPolicyCategories provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyCategories()(*GroupPolicyCategoriesRequestBuilder) {
    return NewGroupPolicyCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyConfigurations provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurations()(*GroupPolicyConfigurationsRequestBuilder) {
    return NewGroupPolicyConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyDefinitionFiles provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyDefinitionFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFiles()(*GroupPolicyDefinitionFilesRequestBuilder) {
    return NewGroupPolicyDefinitionFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyDefinitions provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitions()(*GroupPolicyDefinitionsRequestBuilder) {
    return NewGroupPolicyDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyMigrationReports provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyMigrationReportsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReports()(*GroupPolicyMigrationReportsRequestBuilder) {
    return NewGroupPolicyMigrationReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyObjectFiles provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyObjectFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFiles()(*GroupPolicyObjectFilesRequestBuilder) {
    return NewGroupPolicyObjectFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyUploadedDefinitionFiles provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GroupPolicyUploadedDefinitionFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFiles()(*GroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HardwareConfigurations provides operations to manage the hardwareConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *HardwareConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) HardwareConfigurations()(*HardwareConfigurationsRequestBuilder) {
    return NewHardwareConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HardwarePasswordDetails provides operations to manage the hardwarePasswordDetails property of the microsoft.graph.deviceManagement entity.
// returns a *HardwarePasswordDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) HardwarePasswordDetails()(*HardwarePasswordDetailsRequestBuilder) {
    return NewHardwarePasswordDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HardwarePasswordInfo provides operations to manage the hardwarePasswordInfo property of the microsoft.graph.deviceManagement entity.
// returns a *HardwarePasswordInfoRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) HardwarePasswordInfo()(*HardwarePasswordInfoRequestBuilder) {
    return NewHardwarePasswordInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedDeviceIdentities provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentities()(*ImportedDeviceIdentitiesRequestBuilder) {
    return NewImportedDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Intents provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
// returns a *IntentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Intents()(*IntentsRequestBuilder) {
    return NewIntentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntuneBrandingProfiles provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *IntuneBrandingProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfiles()(*IntuneBrandingProfilesRequestBuilder) {
    return NewIntuneBrandingProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
// returns a *IosUpdateStatusesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*IosUpdateStatusesRequestBuilder) {
    return NewIosUpdateStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MacOSSoftwareUpdateAccountSummaries provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *MacOSSoftwareUpdateAccountSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummaries()(*MacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    return NewMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceCleanupRules provides operations to manage the managedDeviceCleanupRules property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDeviceCleanupRulesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceCleanupRules()(*ManagedDeviceCleanupRulesRequestBuilder) {
    return NewManagedDeviceCleanupRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceEncryptionStates provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDeviceEncryptionStatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStates()(*ManagedDeviceEncryptionStatesRequestBuilder) {
    return NewManagedDeviceEncryptionStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDeviceOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*ManagedDeviceOverviewRequestBuilder) {
    return NewManagedDeviceOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*ManagedDevicesRequestBuilder) {
    return NewManagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceWindowsOSImages provides operations to manage the managedDeviceWindowsOSImages property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDeviceWindowsOSImagesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceWindowsOSImages()(*ManagedDeviceWindowsOSImagesRequestBuilder) {
    return NewManagedDeviceWindowsOSImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelConfigurations provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosoftTunnelConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurations()(*MicrosoftTunnelConfigurationsRequestBuilder) {
    return NewMicrosoftTunnelConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelHealthThresholds provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosoftTunnelHealthThresholdsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholds()(*MicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewMicrosoftTunnelHealthThresholdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelServerLogCollectionResponses provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosoftTunnelServerLogCollectionResponsesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponses()(*MicrosoftTunnelServerLogCollectionResponsesRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelSites provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosoftTunnelSitesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSites()(*MicrosoftTunnelSitesRequestBuilder) {
    return NewMicrosoftTunnelSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *MobileAppTroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*MobileAppTroubleshootingEventsRequestBuilder) {
    return NewMobileAppTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *MobileThreatDefenseConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*MobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobileThreatDefenseConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Monitoring provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
// returns a *MonitoringRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Monitoring()(*MonitoringRequestBuilder) {
    return NewMonitoringRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NdesConnectors provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *NdesConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) NdesConnectors()(*NdesConnectorsRequestBuilder) {
    return NewNdesConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *NotificationMessageTemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*NotificationMessageTemplatesRequestBuilder) {
    return NewNotificationMessageTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OperationApprovalPolicies provides operations to manage the operationApprovalPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *OperationApprovalPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) OperationApprovalPolicies()(*OperationApprovalPoliciesRequestBuilder) {
    return NewOperationApprovalPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OperationApprovalRequests provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
// returns a *OperationApprovalRequestsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) OperationApprovalRequests()(*OperationApprovalRequestsRequestBuilder) {
    return NewOperationApprovalRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update deviceManagement
// returns a DeviceManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceManagementRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable), nil
}
// PrivilegeManagementElevations provides operations to manage the privilegeManagementElevations property of the microsoft.graph.deviceManagement entity.
// returns a *PrivilegeManagementElevationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) PrivilegeManagementElevations()(*PrivilegeManagementElevationsRequestBuilder) {
    return NewPrivilegeManagementElevationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteActionAudits provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteActionAuditsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteActionAudits()(*RemoteActionAuditsRequestBuilder) {
    return NewRemoteActionAuditsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteAssistancePartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*RemoteAssistancePartnersRequestBuilder) {
    return NewRemoteAssistancePartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteAssistanceSettings provides operations to manage the remoteAssistanceSettings property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteAssistanceSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteAssistanceSettings()(*RemoteAssistanceSettingsRequestBuilder) {
    return NewRemoteAssistanceSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
// returns a *ReportsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceAccessProfiles provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceAccessProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ResourceAccessProfiles()(*ResourceAccessProfilesRequestBuilder) {
    return NewResourceAccessProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceOperationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*ResourceOperationsRequestBuilder) {
    return NewResourceOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReusablePolicySettings provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
// returns a *ReusablePolicySettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ReusablePolicySettings()(*ReusablePolicySettingsRequestBuilder) {
    return NewReusablePolicySettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReusableSettings provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ReusableSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ReusableSettings()(*ReusableSettingsRequestBuilder) {
    return NewReusableSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
// returns a *RoleAssignmentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*RoleAssignmentsRequestBuilder) {
    return NewRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *RoleDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*RoleDefinitionsRequestBuilder) {
    return NewRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
// returns a *RoleScopeTagsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleScopeTags()(*RoleScopeTagsRequestBuilder) {
    return NewRoleScopeTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScopedForResourceWithResource provides operations to call the scopedForResource method.
// returns a *ScopedForResourceWithResourceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ScopedForResourceWithResource(resource *string)(*ScopedForResourceWithResourceRequestBuilder) {
    return NewScopedForResourceWithResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, resource)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *SendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SendCustomNotificationToCompanyPortal()(*SendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceNowConnections provides operations to manage the serviceNowConnections property of the microsoft.graph.deviceManagement entity.
// returns a *ServiceNowConnectionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ServiceNowConnections()(*ServiceNowConnectionsRequestBuilder) {
    return NewServiceNowConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *SettingDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SettingDefinitions()(*SettingDefinitionsRequestBuilder) {
    return NewSettingDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
// returns a *SoftwareUpdateStatusSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareUpdateStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *TelecomExpenseManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*TelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomExpenseManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateInsights provides operations to manage the templateInsights property of the microsoft.graph.deviceManagement entity.
// returns a *TemplateInsightsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TemplateInsights()(*TemplateInsightsRequestBuilder) {
    return NewTemplateInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Templates provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
// returns a *TemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateSettings provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
// returns a *TemplateSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TemplateSettings()(*TemplateSettingsRequestBuilder) {
    return NewTemplateSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantAttachRBAC provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
// returns a *TenantAttachRBACRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TenantAttachRBAC()(*TenantAttachRBACRequestBuilder) {
    return NewTenantAttachRBACRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
// returns a *TermsAndConditionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*TermsAndConditionsRequestBuilder) {
    return NewTermsAndConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TroubleshootingEvents provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *TroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*TroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomaly provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAnomalyRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomaly()(*UserExperienceAnalyticsAnomalyRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverview provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyCorrelationGroupOverview()(*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomalyDevice provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAnomalyDeviceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDevice()(*UserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*UserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsAppHealthOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*UserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBaselinesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*UserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpact()(*UserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthCapacityDetails provides operations to manage the userExperienceAnalyticsBatteryHealthCapacityDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDevicePerformance provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthModelPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformance()(*UserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthOsPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformance()(*UserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthRuntimeDetails provides operations to manage the userExperienceAnalyticsBatteryHealthRuntimeDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder) {
    return NewUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*UserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceMetricHistory provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceMetricHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistory()(*UserExperienceAnalyticsDeviceMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceMetricHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScope provides operations to manage the userExperienceAnalyticsDeviceScope property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceScopeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScope()(*UserExperienceAnalyticsDeviceScopeRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScopes provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceScopesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopes()(*UserExperienceAnalyticsDeviceScopesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*UserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*UserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceStartupProcessesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*UserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentity provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentity()(*UserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceTimelineEvent provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsDeviceTimelineEventRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceTimelineEvent()(*UserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    return NewUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsImpactingProcess provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsImpactingProcessRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsMetricHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsModelScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsNotAutopilotReadyDevice provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserExperienceAnalyticsOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsRemoteConnection provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsRemoteConnectionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsResourcePerformance provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsResourcePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsScoreHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*UserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsSummarizedDeviceScopes provides operations to call the userExperienceAnalyticsSummarizedDeviceScopes method.
// returns a *UserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizedDeviceScopes()(*UserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder) {
    return NewUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
// returns a *UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*UserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*UserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserPfxCertificates provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
// returns a *UserPfxCertificatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserPfxCertificates()(*UserPfxCertificatesRequestBuilder) {
    return NewUserPfxCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
// returns a *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, domainName)
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
// returns a *VirtualEndpointRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*VirtualEndpointRequestBuilder) {
    return NewVirtualEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotDeploymentProfiles provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsAutopilotDeploymentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfiles()(*WindowsAutopilotDeploymentProfilesRequestBuilder) {
    return NewWindowsAutopilotDeploymentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*WindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotSettings provides operations to manage the windowsAutopilotSettings property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsAutopilotSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotSettings()(*WindowsAutopilotSettingsRequestBuilder) {
    return NewWindowsAutopilotSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDriverUpdateProfiles provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsDriverUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfiles()(*WindowsDriverUpdateProfilesRequestBuilder) {
    return NewWindowsDriverUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsFeatureUpdateProfiles provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsFeatureUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfiles()(*WindowsFeatureUpdateProfilesRequestBuilder) {
    return NewWindowsFeatureUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsInformationProtectionAppLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*WindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsInformationProtectionNetworkLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*WindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsMalwareInformationRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*WindowsMalwareInformationRequestBuilder) {
    return NewWindowsMalwareInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsQualityUpdatePolicies provides operations to manage the windowsQualityUpdatePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsQualityUpdatePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdatePolicies()(*WindowsQualityUpdatePoliciesRequestBuilder) {
    return NewWindowsQualityUpdatePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsQualityUpdateProfiles provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsQualityUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfiles()(*WindowsQualityUpdateProfilesRequestBuilder) {
    return NewWindowsQualityUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsUpdateCatalogItems provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsUpdateCatalogItemsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItems()(*WindowsUpdateCatalogItemsRequestBuilder) {
    return NewWindowsUpdateCatalogItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRequestBuilder) {
    return NewDeviceManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// ZebraFotaArtifacts provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
// returns a *ZebraFotaArtifactsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifacts()(*ZebraFotaArtifactsRequestBuilder) {
    return NewZebraFotaArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ZebraFotaConnector provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
// returns a *ZebraFotaConnectorRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaConnector()(*ZebraFotaConnectorRequestBuilder) {
    return NewZebraFotaConnectorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ZebraFotaDeployments provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
// returns a *ZebraFotaDeploymentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaDeployments()(*ZebraFotaDeploymentsRequestBuilder) {
    return NewZebraFotaDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
