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
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AdvancedThreatProtectionOnboardingStateSummary()(*AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidDeviceOwnerEnrollmentProfiles provides operations to manage the androidDeviceOwnerEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidDeviceOwnerEnrollmentProfiles()(*AndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilder) {
    return NewAndroiddeviceownerenrollmentprofilesAndroidDeviceOwnerEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkAppConfigurationSchemas provides operations to manage the androidForWorkAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidforworkappconfigurationschemasAndroidForWorkAppConfigurationSchemasRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkAppConfigurationSchemas()(*AndroidforworkappconfigurationschemasAndroidForWorkAppConfigurationSchemasRequestBuilder) {
    return NewAndroidforworkappconfigurationschemasAndroidForWorkAppConfigurationSchemasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkEnrollmentProfiles provides operations to manage the androidForWorkEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkEnrollmentProfiles()(*AndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfilesRequestBuilder) {
    return NewAndroidforworkenrollmentprofilesAndroidForWorkEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidForWorkSettings provides operations to manage the androidForWorkSettings property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidforworksettingsAndroidForWorkSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidForWorkSettings()(*AndroidforworksettingsAndroidForWorkSettingsRequestBuilder) {
    return NewAndroidforworksettingsAndroidForWorkSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidManagedStoreAccountEnterpriseSettings provides operations to manage the androidManagedStoreAccountEnterpriseSettings property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAccountEnterpriseSettings()(*AndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilder) {
    return NewAndroidmanagedstoreaccountenterprisesettingsAndroidManagedStoreAccountEnterpriseSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AndroidManagedStoreAppConfigurationSchemas provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
// returns a *AndroidmanagedstoreappconfigurationschemasAndroidManagedStoreAppConfigurationSchemasRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AndroidManagedStoreAppConfigurationSchemas()(*AndroidmanagedstoreappconfigurationschemasAndroidManagedStoreAppConfigurationSchemasRequestBuilder) {
    return NewAndroidmanagedstoreappconfigurationschemasAndroidManagedStoreAppConfigurationSchemasRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
// returns a *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) {
    return NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppleUserInitiatedEnrollmentProfiles provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AppleUserInitiatedEnrollmentProfiles()(*AppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilder) {
    return NewAppleuserinitiatedenrollmentprofilesAppleUserInitiatedEnrollmentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentFilters provides operations to manage the assignmentFilters property of the microsoft.graph.deviceManagement entity.
// returns a *AssignmentfiltersAssignmentFiltersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AssignmentFilters()(*AssignmentfiltersAssignmentFiltersRequestBuilder) {
    return NewAssignmentfiltersAssignmentFiltersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AuditeventsAuditEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AuditEvents()(*AuditeventsAuditEventsRequestBuilder) {
    return NewAuditeventsAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AutopilotEvents provides operations to manage the autopilotEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AutopiloteventsAutopilotEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AutopilotEvents()(*AutopiloteventsAutopilotEventsRequestBuilder) {
    return NewAutopiloteventsAutopilotEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CartToClassAssociations provides operations to manage the cartToClassAssociations property of the microsoft.graph.deviceManagement entity.
// returns a *CarttoclassassociationsCartToClassAssociationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CartToClassAssociations()(*CarttoclassassociationsCartToClassAssociationsRequestBuilder) {
    return NewCarttoclassassociationsCartToClassAssociationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.deviceManagement entity.
// returns a *CategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Categories()(*CategoriesRequestBuilder) {
    return NewCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CertificateConnectorDetails provides operations to manage the certificateConnectorDetails property of the microsoft.graph.deviceManagement entity.
// returns a *CertificateconnectordetailsCertificateConnectorDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CertificateConnectorDetails()(*CertificateconnectordetailsCertificateConnectorDetailsRequestBuilder) {
    return NewCertificateconnectordetailsCertificateConnectorDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ChromeOSOnboardingSettings provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ChromeOSOnboardingSettings()(*ChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilder) {
    return NewChromeosonboardingsettingsChromeOSOnboardingSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CloudPCConnectivityIssues provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
// returns a *CloudpcconnectivityissuesCloudPCConnectivityIssuesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CloudPCConnectivityIssues()(*CloudpcconnectivityissuesCloudPCConnectivityIssuesRequestBuilder) {
    return NewCloudpcconnectivityissuesCloudPCConnectivityIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComanagedDevices provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ComanageddevicesComanagedDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComanagedDevices()(*ComanageddevicesComanagedDevicesRequestBuilder) {
    return NewComanageddevicesComanagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComanagementEligibleDevices provides operations to manage the comanagementEligibleDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ComanagementeligibledevicesComanagementEligibleDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComanagementEligibleDevices()(*ComanagementeligibledevicesComanagementEligibleDevicesRequestBuilder) {
    return NewComanagementeligibledevicesComanagementEligibleDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceCategories provides operations to manage the complianceCategories property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancecategoriesComplianceCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceCategories()(*CompliancecategoriesComplianceCategoriesRequestBuilder) {
    return NewCompliancecategoriesComplianceCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) {
    return NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompliancePolicies provides operations to manage the compliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancepoliciesCompliancePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) CompliancePolicies()(*CompliancepoliciesCompliancePoliciesRequestBuilder) {
    return NewCompliancepoliciesCompliancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceSettings provides operations to manage the complianceSettings property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancesettingsComplianceSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceSettings()(*CompliancesettingsComplianceSettingsRequestBuilder) {
    return NewCompliancesettingsComplianceSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ConditionalaccesssettingsConditionalAccessSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*ConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    return NewConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigManagerCollections provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigManagerCollections()(*ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) {
    return NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationCategories provides operations to manage the configurationCategories property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationcategoriesConfigurationCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationCategories()(*ConfigurationcategoriesConfigurationCategoriesRequestBuilder) {
    return NewConfigurationcategoriesConfigurationCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationPolicies provides operations to manage the configurationPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationpoliciesConfigurationPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationPolicies()(*ConfigurationpoliciesConfigurationPoliciesRequestBuilder) {
    return NewConfigurationpoliciesConfigurationPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationPolicyTemplates provides operations to manage the configurationPolicyTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationPolicyTemplates()(*ConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilder) {
    return NewConfigurationpolicytemplatesConfigurationPolicyTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConfigurationSettings provides operations to manage the configurationSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigurationsettingsConfigurationSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConfigurationSettings()(*ConfigurationsettingsConfigurationSettingsRequestBuilder) {
    return NewConfigurationsettingsConfigurationSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *DatasharingconsentsDataSharingConsentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DataSharingConsents()(*DatasharingconsentsDataSharingConsentsRequestBuilder) {
    return NewDatasharingconsentsDataSharingConsentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DepOnboardingSettings provides operations to manage the depOnboardingSettings property of the microsoft.graph.deviceManagement entity.
// returns a *DeponboardingsettingsDepOnboardingSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DepOnboardingSettings()(*DeponboardingsettingsDepOnboardingSettingsRequestBuilder) {
    return NewDeponboardingsettingsDepOnboardingSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DerivedCredentials provides operations to manage the derivedCredentials property of the microsoft.graph.deviceManagement entity.
// returns a *DerivedcredentialsDerivedCredentialsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DerivedCredentials()(*DerivedcredentialsDerivedCredentialsRequestBuilder) {
    return NewDerivedcredentialsDerivedCredentialsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
// returns a *DetectedappsDetectedAppsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DetectedappsDetectedAppsRequestBuilder) {
    return NewDetectedappsDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecategoriesDeviceCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DevicecategoriesDeviceCategoriesRequestBuilder) {
    return NewDevicecategoriesDeviceCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) {
    return NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceComplianceScripts provides operations to manage the deviceComplianceScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancescriptsDeviceComplianceScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceComplianceScripts()(*DevicecompliancescriptsDeviceComplianceScriptsRequestBuilder) {
    return NewDevicecompliancescriptsDeviceComplianceScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationConflictSummary provides operations to manage the deviceConfigurationConflictSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationConflictSummary()(*DeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilder) {
    return NewDeviceconfigurationconflictsummaryDeviceConfigurationConflictSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationProfiles provides operations to manage the deviceConfigurationProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationprofilesDeviceConfigurationProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationProfiles()(*DeviceconfigurationprofilesDeviceConfigurationProfilesRequestBuilder) {
    return NewDeviceconfigurationprofilesDeviceConfigurationProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationRestrictedAppsViolations provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationRestrictedAppsViolations()(*DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationsDeviceConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceconfigurationsDeviceConfigurationsRequestBuilder) {
    return NewDeviceconfigurationsDeviceConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationsAllManagedDeviceCertificateStates provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsAllManagedDeviceCertificateStates()(*DeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    return NewDeviceconfigurationsallmanageddevicecertificatestatesDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationUserStateSummaries provides operations to manage the deviceConfigurationUserStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationuserstatesummariesDeviceConfigurationUserStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationUserStateSummaries()(*DeviceconfigurationuserstatesummariesDeviceConfigurationUserStateSummariesRequestBuilder) {
    return NewDeviceconfigurationuserstatesummariesDeviceConfigurationUserStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCustomAttributeShellScripts provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCustomAttributeShellScripts()(*DevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilder) {
    return NewDevicecustomattributeshellscriptsDeviceCustomAttributeShellScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceHealthScripts provides operations to manage the deviceHealthScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicehealthscriptsDeviceHealthScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceHealthScripts()(*DevicehealthscriptsDeviceHealthScriptsRequestBuilder) {
    return NewDevicehealthscriptsDeviceHealthScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *DevicemanagementpartnersDeviceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DevicemanagementpartnersDeviceManagementPartnersRequestBuilder) {
    return NewDevicemanagementpartnersDeviceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementScripts provides operations to manage the deviceManagementScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DevicemanagementscriptsDeviceManagementScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceManagementScripts()(*DevicemanagementscriptsDeviceManagementScriptsRequestBuilder) {
    return NewDevicemanagementscriptsDeviceManagementScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceShellScripts provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceshellscriptsDeviceShellScriptsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceShellScripts()(*DeviceshellscriptsDeviceShellScriptsRequestBuilder) {
    return NewDeviceshellscriptsDeviceShellScriptsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DomainJoinConnectors provides operations to manage the domainJoinConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *DomainjoinconnectorsDomainJoinConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DomainJoinConnectors()(*DomainjoinconnectorsDomainJoinConnectorsRequestBuilder) {
    return NewDomainjoinconnectorsDomainJoinConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ElevationRequests provides operations to manage the elevationRequests property of the microsoft.graph.deviceManagement entity.
// returns a *ElevationrequestsElevationRequestsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ElevationRequests()(*ElevationrequestsElevationRequestsRequestBuilder) {
    return NewElevationrequestsElevationRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EmbeddedSIMActivationCodePools provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
// returns a *EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EmbeddedSIMActivationCodePools()(*EmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilder) {
    return NewEmbeddedsimactivationcodepoolsEmbeddedSIMActivationCodePoolsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableAndroidDeviceAdministratorEnrollment provides operations to call the enableAndroidDeviceAdministratorEnrollment method.
// returns a *EnableandroiddeviceadministratorenrollmentEnableAndroidDeviceAdministratorEnrollmentRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableAndroidDeviceAdministratorEnrollment()(*EnableandroiddeviceadministratorenrollmentEnableAndroidDeviceAdministratorEnrollmentRequestBuilder) {
    return NewEnableandroiddeviceadministratorenrollmentEnableAndroidDeviceAdministratorEnrollmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableLegacyPcManagement provides operations to call the enableLegacyPcManagement method.
// returns a *EnablelegacypcmanagementEnableLegacyPcManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableLegacyPcManagement()(*EnablelegacypcmanagementEnableLegacyPcManagementRequestBuilder) {
    return NewEnablelegacypcmanagementEnableLegacyPcManagementRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EnableUnlicensedAdminstrators provides operations to call the enableUnlicensedAdminstrators method.
// returns a *EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EnableUnlicensedAdminstrators()(*EnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilder) {
    return NewEnableunlicensedadminstratorsEnableUnlicensedAdminstratorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EvaluateAssignmentFilter provides operations to call the evaluateAssignmentFilter method.
// returns a *EvaluateassignmentfilterEvaluateAssignmentFilterRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) EvaluateAssignmentFilter()(*EvaluateassignmentfilterEvaluateAssignmentFilterRequestBuilder) {
    return NewEvaluateassignmentfilterEvaluateAssignmentFilterRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeconnectorsExchangeConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ExchangeconnectorsExchangeConnectorsRequestBuilder) {
    return NewExchangeconnectorsExchangeConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeOnPremisesPolicies provides operations to manage the exchangeOnPremisesPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicies()(*ExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilder) {
    return NewExchangeonpremisespoliciesExchangeOnPremisesPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeOnPremisesPolicy provides operations to manage the exchangeOnPremisesPolicy property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeOnPremisesPolicy()(*ExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilder) {
    return NewExchangeonpremisespolicyExchangeOnPremisesPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetAssignedRoleDetails()(*GetassignedroledetailsGetAssignedRoleDetailsRequestBuilder) {
    return NewGetassignedroledetailsGetAssignedRoleDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetAssignmentFiltersStatusDetails provides operations to call the getAssignmentFiltersStatusDetails method.
// returns a *GetassignmentfiltersstatusdetailsGetAssignmentFiltersStatusDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetAssignmentFiltersStatusDetails()(*GetassignmentfiltersstatusdetailsGetAssignmentFiltersStatusDetailsRequestBuilder) {
    return NewGetassignmentfiltersstatusdetailsGetAssignmentFiltersStatusDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComanagedDevicesSummary provides operations to call the getComanagedDevicesSummary method.
// returns a *GetcomanageddevicessummaryGetComanagedDevicesSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetComanagedDevicesSummary()(*GetcomanageddevicessummaryGetComanagedDevicesSummaryRequestBuilder) {
    return NewGetcomanageddevicessummaryGetComanagedDevicesSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetComanagementEligibleDevicesSummary provides operations to call the getComanagementEligibleDevicesSummary method.
// returns a *GetcomanagementeligibledevicessummaryGetComanagementEligibleDevicesSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetComanagementEligibleDevicesSummary()(*GetcomanagementeligibledevicessummaryGetComanagementEligibleDevicesSummaryRequestBuilder) {
    return NewGetcomanagementeligibledevicessummaryGetComanagementEligibleDevicesSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEffectivePermissions provides operations to call the getEffectivePermissions method.
// returns a *GeteffectivepermissionsGetEffectivePermissionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetEffectivePermissions()(*GeteffectivepermissionsGetEffectivePermissionsRequestBuilder) {
    return NewGeteffectivepermissionsGetEffectivePermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
// returns a *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) {
    return NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, scope)
}
// GetRoleScopeTagsByIdsWithIds provides operations to call the getRoleScopeTagsByIds method.
// returns a *GetrolescopetagsbyidswithidsGetRoleScopeTagsByIdsWithIdsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByIdsWithIds(ids *string)(*GetrolescopetagsbyidswithidsGetRoleScopeTagsByIdsWithIdsRequestBuilder) {
    return NewGetrolescopetagsbyidswithidsGetRoleScopeTagsByIdsWithIdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, ids)
}
// GetRoleScopeTagsByResourceWithResource provides operations to call the getRoleScopeTagsByResource method.
// returns a *GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetRoleScopeTagsByResourceWithResource(resource *string)(*GetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilder) {
    return NewGetrolescopetagsbyresourcewithresourceGetRoleScopeTagsByResourceWithResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, resource)
}
// GetSuggestedEnrollmentLimitWithEnrollmentType provides operations to call the getSuggestedEnrollmentLimit method.
// returns a *GetsuggestedenrollmentlimitwithenrollmenttypeGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetSuggestedEnrollmentLimitWithEnrollmentType(enrollmentType *string)(*GetsuggestedenrollmentlimitwithenrollmenttypeGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilder) {
    return NewGetsuggestedenrollmentlimitwithenrollmenttypeGetSuggestedEnrollmentLimitWithEnrollmentTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, enrollmentType)
}
// GroupPolicyCategories provides operations to manage the groupPolicyCategories property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicycategoriesGroupPolicyCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyCategories()(*GrouppolicycategoriesGroupPolicyCategoriesRequestBuilder) {
    return NewGrouppolicycategoriesGroupPolicyCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyConfigurations provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicyconfigurationsGroupPolicyConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyConfigurations()(*GrouppolicyconfigurationsGroupPolicyConfigurationsRequestBuilder) {
    return NewGrouppolicyconfigurationsGroupPolicyConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyDefinitionFiles provides operations to manage the groupPolicyDefinitionFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicydefinitionfilesGroupPolicyDefinitionFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitionFiles()(*GrouppolicydefinitionfilesGroupPolicyDefinitionFilesRequestBuilder) {
    return NewGrouppolicydefinitionfilesGroupPolicyDefinitionFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyDefinitions provides operations to manage the groupPolicyDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicydefinitionsGroupPolicyDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyDefinitions()(*GrouppolicydefinitionsGroupPolicyDefinitionsRequestBuilder) {
    return NewGrouppolicydefinitionsGroupPolicyDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyMigrationReports provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicymigrationreportsGroupPolicyMigrationReportsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyMigrationReports()(*GrouppolicymigrationreportsGroupPolicyMigrationReportsRequestBuilder) {
    return NewGrouppolicymigrationreportsGroupPolicyMigrationReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyObjectFiles provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyObjectFiles()(*GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) {
    return NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyUploadedDefinitionFiles provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GroupPolicyUploadedDefinitionFiles()(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HardwareConfigurations provides operations to manage the hardwareConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *HardwareconfigurationsHardwareConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) HardwareConfigurations()(*HardwareconfigurationsHardwareConfigurationsRequestBuilder) {
    return NewHardwareconfigurationsHardwareConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// HardwarePasswordInfo provides operations to manage the hardwarePasswordInfo property of the microsoft.graph.deviceManagement entity.
// returns a *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) HardwarePasswordInfo()(*HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) {
    return NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedDeviceIdentities provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ImportedDeviceIdentities()(*ImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilder) {
    return NewImporteddeviceidentitiesImportedDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Intents provides operations to manage the intents property of the microsoft.graph.deviceManagement entity.
// returns a *IntentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Intents()(*IntentsRequestBuilder) {
    return NewIntentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IntuneBrandingProfiles provides operations to manage the intuneBrandingProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *IntunebrandingprofilesIntuneBrandingProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) IntuneBrandingProfiles()(*IntunebrandingprofilesIntuneBrandingProfilesRequestBuilder) {
    return NewIntunebrandingprofilesIntuneBrandingProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
// returns a *IosupdatestatusesIosUpdateStatusesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*IosupdatestatusesIosUpdateStatusesRequestBuilder) {
    return NewIosupdatestatusesIosUpdateStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MacOSSoftwareUpdateAccountSummaries provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MacOSSoftwareUpdateAccountSummaries()(*MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummariesRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceCleanupRules provides operations to manage the managedDeviceCleanupRules property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddevicecleanuprulesManagedDeviceCleanupRulesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceCleanupRules()(*ManageddevicecleanuprulesManagedDeviceCleanupRulesRequestBuilder) {
    return NewManageddevicecleanuprulesManagedDeviceCleanupRulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceEncryptionStates provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddeviceencryptionstatesManagedDeviceEncryptionStatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceEncryptionStates()(*ManageddeviceencryptionstatesManagedDeviceEncryptionStatesRequestBuilder) {
    return NewManageddeviceencryptionstatesManagedDeviceEncryptionStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) {
    return NewManageddeviceoverviewManagedDeviceOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddevicesManagedDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*ManageddevicesManagedDevicesRequestBuilder) {
    return NewManageddevicesManagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelConfigurations provides operations to manage the microsoftTunnelConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelConfigurations()(*MicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilder) {
    return NewMicrosofttunnelconfigurationsMicrosoftTunnelConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelHealthThresholds provides operations to manage the microsoftTunnelHealthThresholds property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelHealthThresholds()(*MicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilder) {
    return NewMicrosofttunnelhealththresholdsMicrosoftTunnelHealthThresholdsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelServerLogCollectionResponses provides operations to manage the microsoftTunnelServerLogCollectionResponses property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponsesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelServerLogCollectionResponses()(*MicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponsesRequestBuilder) {
    return NewMicrosofttunnelserverlogcollectionresponsesMicrosoftTunnelServerLogCollectionResponsesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftTunnelSites provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
// returns a *MicrosofttunnelsitesMicrosoftTunnelSitesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MicrosoftTunnelSites()(*MicrosofttunnelsitesMicrosoftTunnelSitesRequestBuilder) {
    return NewMicrosofttunnelsitesMicrosoftTunnelSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *MobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*MobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder) {
    return NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Monitoring provides operations to manage the monitoring property of the microsoft.graph.deviceManagement entity.
// returns a *MonitoringRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Monitoring()(*MonitoringRequestBuilder) {
    return NewMonitoringRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NdesConnectors provides operations to manage the ndesConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *NdesconnectorsNdesConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) NdesConnectors()(*NdesconnectorsNdesConnectorsRequestBuilder) {
    return NewNdesconnectorsNdesConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *NotificationmessagetemplatesNotificationMessageTemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*NotificationmessagetemplatesNotificationMessageTemplatesRequestBuilder) {
    return NewNotificationmessagetemplatesNotificationMessageTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OperationApprovalPolicies provides operations to manage the operationApprovalPolicies property of the microsoft.graph.deviceManagement entity.
// returns a *OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) OperationApprovalPolicies()(*OperationapprovalpoliciesOperationApprovalPoliciesRequestBuilder) {
    return NewOperationapprovalpoliciesOperationApprovalPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OperationApprovalRequests provides operations to manage the operationApprovalRequests property of the microsoft.graph.deviceManagement entity.
// returns a *OperationapprovalrequestsOperationApprovalRequestsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) OperationApprovalRequests()(*OperationapprovalrequestsOperationApprovalRequestsRequestBuilder) {
    return NewOperationapprovalrequestsOperationApprovalRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) PrivilegeManagementElevations()(*PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) {
    return NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteActionAudits provides operations to manage the remoteActionAudits property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteactionauditsRemoteActionAuditsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteActionAudits()(*RemoteactionauditsRemoteActionAuditsRequestBuilder) {
    return NewRemoteactionauditsRemoteActionAuditsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteassistancepartnersRemoteAssistancePartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*RemoteassistancepartnersRemoteAssistancePartnersRequestBuilder) {
    return NewRemoteassistancepartnersRemoteAssistancePartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoteAssistanceSettings provides operations to manage the remoteAssistanceSettings property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteassistancesettingsRemoteAssistanceSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteAssistanceSettings()(*RemoteassistancesettingsRemoteAssistanceSettingsRequestBuilder) {
    return NewRemoteassistancesettingsRemoteAssistanceSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
// returns a *ReportsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceAccessProfiles provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceaccessprofilesResourceAccessProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ResourceAccessProfiles()(*ResourceaccessprofilesResourceAccessProfilesRequestBuilder) {
    return NewResourceaccessprofilesResourceAccessProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceoperationsResourceOperationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*ResourceoperationsResourceOperationsRequestBuilder) {
    return NewResourceoperationsResourceOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReusablePolicySettings provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
// returns a *ReusablepolicysettingsReusablePolicySettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ReusablePolicySettings()(*ReusablepolicysettingsReusablePolicySettingsRequestBuilder) {
    return NewReusablepolicysettingsReusablePolicySettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReusableSettings provides operations to manage the reusableSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ReusablesettingsReusableSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ReusableSettings()(*ReusablesettingsReusableSettingsRequestBuilder) {
    return NewReusablesettingsReusableSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
// returns a *RoleassignmentsRoleAssignmentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*RoleassignmentsRoleAssignmentsRequestBuilder) {
    return NewRoleassignmentsRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *RoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*RoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewRoledefinitionsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleScopeTags provides operations to manage the roleScopeTags property of the microsoft.graph.deviceManagement entity.
// returns a *RolescopetagsRoleScopeTagsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleScopeTags()(*RolescopetagsRoleScopeTagsRequestBuilder) {
    return NewRolescopetagsRoleScopeTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ScopedForResourceWithResource provides operations to call the scopedForResource method.
// returns a *ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ScopedForResourceWithResource(resource *string)(*ScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilder) {
    return NewScopedforresourcewithresourceScopedForResourceWithResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, resource)
}
// SendCustomNotificationToCompanyPortal provides operations to call the sendCustomNotificationToCompanyPortal method.
// returns a *SendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SendCustomNotificationToCompanyPortal()(*SendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilder) {
    return NewSendcustomnotificationtocompanyportalSendCustomNotificationToCompanyPortalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceNowConnections provides operations to manage the serviceNowConnections property of the microsoft.graph.deviceManagement entity.
// returns a *ServicenowconnectionsServiceNowConnectionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ServiceNowConnections()(*ServicenowconnectionsServiceNowConnectionsRequestBuilder) {
    return NewServicenowconnectionsServiceNowConnectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SettingDefinitions provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *SettingdefinitionsSettingDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SettingDefinitions()(*SettingdefinitionsSettingDefinitionsRequestBuilder) {
    return NewSettingdefinitionsSettingDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
// returns a *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateInsights provides operations to manage the templateInsights property of the microsoft.graph.deviceManagement entity.
// returns a *TemplateinsightsTemplateInsightsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TemplateInsights()(*TemplateinsightsTemplateInsightsRequestBuilder) {
    return NewTemplateinsightsTemplateInsightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Templates provides operations to manage the templates property of the microsoft.graph.deviceManagement entity.
// returns a *TemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Templates()(*TemplatesRequestBuilder) {
    return NewTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TemplateSettings provides operations to manage the templateSettings property of the microsoft.graph.deviceManagement entity.
// returns a *TemplatesettingsTemplateSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TemplateSettings()(*TemplatesettingsTemplateSettingsRequestBuilder) {
    return NewTemplatesettingsTemplateSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TenantAttachRBAC provides operations to manage the tenantAttachRBAC property of the microsoft.graph.deviceManagement entity.
// returns a *TenantattachrbacTenantAttachRBACRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TenantAttachRBAC()(*TenantattachrbacTenantAttachRBACRequestBuilder) {
    return NewTenantattachrbacTenantAttachRBACRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
// returns a *TermsandconditionsTermsAndConditionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*TermsandconditionsTermsAndConditionsRequestBuilder) {
    return NewTermsandconditionsTermsAndConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
// returns a *TroubleshootingeventsTroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*TroubleshootingeventsTroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingeventsTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomaly provides operations to manage the userExperienceAnalyticsAnomaly property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsanomalyUserExperienceAnalyticsAnomalyRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomaly()(*UserexperienceanalyticsanomalyUserExperienceAnalyticsAnomalyRequestBuilder) {
    return NewUserexperienceanalyticsanomalyUserExperienceAnalyticsAnomalyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomalyCorrelationGroupOverview provides operations to manage the userExperienceAnalyticsAnomalyCorrelationGroupOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyCorrelationGroupOverview()(*UserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilder) {
    return NewUserexperienceanalyticsanomalycorrelationgroupoverviewUserExperienceAnalyticsAnomalyCorrelationGroupOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAnomalyDevice provides operations to manage the userExperienceAnalyticsAnomalyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAnomalyDevice()(*UserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilder) {
    return NewUserexperienceanalyticsanomalydeviceUserExperienceAnalyticsAnomalyDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversionUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()(*UserexperienceanalyticsapphealthapplicationperformancebyappversionUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversionUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*UserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*UserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewUserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthAppImpact property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthappimpactUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthAppImpact()(*UserexperienceanalyticsbatteryhealthappimpactUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthappimpactUserExperienceAnalyticsBatteryHealthAppImpactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthCapacityDetails provides operations to manage the userExperienceAnalyticsBatteryHealthCapacityDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthcapacitydetailsUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserexperienceanalyticsbatteryhealthcapacitydetailsUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthcapacitydetailsUserExperienceAnalyticsBatteryHealthCapacityDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceAppImpact property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthdeviceappimpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserexperienceanalyticsbatteryhealthdeviceappimpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceappimpactUserExperienceAnalyticsBatteryHealthDeviceAppImpactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDevicePerformance provides operations to manage the userExperienceAnalyticsBatteryHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDevicePerformance()(*UserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceperformanceUserExperienceAnalyticsBatteryHealthDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()(*UserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthdeviceruntimehistoryUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthModelPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthmodelperformanceUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthModelPerformance()(*UserexperienceanalyticsbatteryhealthmodelperformanceUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthmodelperformanceUserExperienceAnalyticsBatteryHealthModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthOsPerformance provides operations to manage the userExperienceAnalyticsBatteryHealthOsPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthosperformanceUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthOsPerformance()(*UserexperienceanalyticsbatteryhealthosperformanceUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthosperformanceUserExperienceAnalyticsBatteryHealthOsPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBatteryHealthRuntimeDetails provides operations to manage the userExperienceAnalyticsBatteryHealthRuntimeDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbatteryhealthruntimedetailsUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserexperienceanalyticsbatteryhealthruntimedetailsUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilder) {
    return NewUserexperienceanalyticsbatteryhealthruntimedetailsUserExperienceAnalyticsBatteryHealthRuntimeDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*UserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewUserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceMetricHistory provides operations to manage the userExperienceAnalyticsDeviceMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicemetrichistoryUserExperienceAnalyticsDeviceMetricHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceMetricHistory()(*UserexperienceanalyticsdevicemetrichistoryUserExperienceAnalyticsDeviceMetricHistoryRequestBuilder) {
    return NewUserexperienceanalyticsdevicemetrichistoryUserExperienceAnalyticsDeviceMetricHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScope provides operations to manage the userExperienceAnalyticsDeviceScope property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicescopeUserExperienceAnalyticsDeviceScopeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScope()(*UserexperienceanalyticsdevicescopeUserExperienceAnalyticsDeviceScopeRequestBuilder) {
    return NewUserexperienceanalyticsdevicescopeUserExperienceAnalyticsDeviceScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScopes provides operations to manage the userExperienceAnalyticsDeviceScopes property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicescopesUserExperienceAnalyticsDeviceScopesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScopes()(*UserexperienceanalyticsdevicescopesUserExperienceAnalyticsDeviceScopesRequestBuilder) {
    return NewUserexperienceanalyticsdevicescopesUserExperienceAnalyticsDeviceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*UserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewUserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*UserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDevicesWithoutCloudIdentity provides operations to manage the userExperienceAnalyticsDevicesWithoutCloudIdentity property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicesWithoutCloudIdentity()(*UserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilder) {
    return NewUserexperienceanalyticsdeviceswithoutcloudidentityUserExperienceAnalyticsDevicesWithoutCloudIdentityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceTimelineEvent provides operations to manage the userExperienceAnalyticsDeviceTimelineEvent property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceTimelineEvent()(*UserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilder) {
    return NewUserexperienceanalyticsdevicetimelineeventUserExperienceAnalyticsDeviceTimelineEventRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsImpactingProcess provides operations to manage the userExperienceAnalyticsImpactingProcess property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsImpactingProcess()(*UserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilder) {
    return NewUserexperienceanalyticsimpactingprocessUserExperienceAnalyticsImpactingProcessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsNotAutopilotReadyDevice provides operations to manage the userExperienceAnalyticsNotAutopilotReadyDevice property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsNotAutopilotReadyDevice()(*UserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilder) {
    return NewUserexperienceanalyticsnotautopilotreadydeviceUserExperienceAnalyticsNotAutopilotReadyDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsRemoteConnection provides operations to manage the userExperienceAnalyticsRemoteConnection property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsRemoteConnection()(*UserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilder) {
    return NewUserexperienceanalyticsremoteconnectionUserExperienceAnalyticsRemoteConnectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsResourcePerformance provides operations to manage the userExperienceAnalyticsResourcePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsResourcePerformance()(*UserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsresourceperformanceUserExperienceAnalyticsResourcePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsSummarizedDeviceScopes provides operations to call the userExperienceAnalyticsSummarizedDeviceScopes method.
// returns a *UserexperienceanalyticssummarizeddevicescopesUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizedDeviceScopes()(*UserexperienceanalyticssummarizeddevicescopesUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilder) {
    return NewUserexperienceanalyticssummarizeddevicescopesUserExperienceAnalyticsSummarizedDeviceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
// returns a *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserPfxCertificates provides operations to manage the userPfxCertificates property of the microsoft.graph.deviceManagement entity.
// returns a *UserpfxcertificatesUserPfxCertificatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserPfxCertificates()(*UserpfxcertificatesUserPfxCertificatesRequestBuilder) {
    return NewUserpfxcertificatesUserPfxCertificatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
// returns a *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, domainName)
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
// returns a *VirtualendpointVirtualEndpointRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*VirtualendpointVirtualEndpointRequestBuilder) {
    return NewVirtualendpointVirtualEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotDeploymentProfiles provides operations to manage the windowsAutopilotDeploymentProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsautopilotdeploymentprofilesWindowsAutopilotDeploymentProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeploymentProfiles()(*WindowsautopilotdeploymentprofilesWindowsAutopilotDeploymentProfilesRequestBuilder) {
    return NewWindowsautopilotdeploymentprofilesWindowsAutopilotDeploymentProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotSettings provides operations to manage the windowsAutopilotSettings property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsautopilotsettingsWindowsAutopilotSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotSettings()(*WindowsautopilotsettingsWindowsAutopilotSettingsRequestBuilder) {
    return NewWindowsautopilotsettingsWindowsAutopilotSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsDriverUpdateProfiles provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsDriverUpdateProfiles()(*WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) {
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsFeatureUpdateProfiles provides operations to manage the windowsFeatureUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsfeatureupdateprofilesWindowsFeatureUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsFeatureUpdateProfiles()(*WindowsfeatureupdateprofilesWindowsFeatureUpdateProfilesRequestBuilder) {
    return NewWindowsfeatureupdateprofilesWindowsFeatureUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsmalwareinformationWindowsMalwareInformationRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*WindowsmalwareinformationWindowsMalwareInformationRequestBuilder) {
    return NewWindowsmalwareinformationWindowsMalwareInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsQualityUpdateProfiles provides operations to manage the windowsQualityUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsqualityupdateprofilesWindowsQualityUpdateProfilesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsQualityUpdateProfiles()(*WindowsqualityupdateprofilesWindowsQualityUpdateProfilesRequestBuilder) {
    return NewWindowsqualityupdateprofilesWindowsQualityUpdateProfilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsUpdateCatalogItems provides operations to manage the windowsUpdateCatalogItems property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsUpdateCatalogItems()(*WindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilder) {
    return NewWindowsupdatecatalogitemsWindowsUpdateCatalogItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRequestBuilder) {
    return NewDeviceManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
// ZebraFotaArtifacts provides operations to manage the zebraFotaArtifacts property of the microsoft.graph.deviceManagement entity.
// returns a *ZebrafotaartifactsZebraFotaArtifactsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaArtifacts()(*ZebrafotaartifactsZebraFotaArtifactsRequestBuilder) {
    return NewZebrafotaartifactsZebraFotaArtifactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ZebraFotaConnector provides operations to manage the zebraFotaConnector property of the microsoft.graph.deviceManagement entity.
// returns a *ZebrafotaconnectorZebraFotaConnectorRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaConnector()(*ZebrafotaconnectorZebraFotaConnectorRequestBuilder) {
    return NewZebrafotaconnectorZebraFotaConnectorRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ZebraFotaDeployments provides operations to manage the zebraFotaDeployments property of the microsoft.graph.deviceManagement entity.
// returns a *ZebrafotadeploymentsZebraFotaDeploymentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ZebraFotaDeployments()(*ZebrafotadeploymentsZebraFotaDeploymentsRequestBuilder) {
    return NewZebrafotadeploymentsZebraFotaDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
