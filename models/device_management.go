package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagement 
type DeviceManagement struct {
    Entity
    // The date & time when tenant data moved between scaleunits.
    accountMoveCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin consent information.
    adminConsent AdminConsentable
    // The summary state of ATP onboarding state for this account.
    advancedThreatProtectionOnboardingStateSummary AdvancedThreatProtectionOnboardingStateSummaryable
    // Android device owner enrollment profile entities.
    androidDeviceOwnerEnrollmentProfiles []AndroidDeviceOwnerEnrollmentProfileable
    // Android for Work app configuration schema entities.
    androidForWorkAppConfigurationSchemas []AndroidForWorkAppConfigurationSchemaable
    // Android for Work enrollment profile entities.
    androidForWorkEnrollmentProfiles []AndroidForWorkEnrollmentProfileable
    // The singleton Android for Work settings entity.
    androidForWorkSettings AndroidForWorkSettingsable
    // The singleton Android managed store account enterprise settings entity.
    androidManagedStoreAccountEnterpriseSettings AndroidManagedStoreAccountEnterpriseSettingsable
    // Android Enterprise app configuration schema entities.
    androidManagedStoreAppConfigurationSchemas []AndroidManagedStoreAppConfigurationSchemaable
    // Apple push notification certificate.
    applePushNotificationCertificate ApplePushNotificationCertificateable
    // Apple user initiated enrollment profiles
    appleUserInitiatedEnrollmentProfiles []AppleUserInitiatedEnrollmentProfileable
    // The list of assignment filters
    assignmentFilters []DeviceAndAppManagementAssignmentFilterable
    // The Audit Events
    auditEvents []AuditEventable
    // The list of autopilot events for the tenant.
    autopilotEvents []DeviceManagementAutopilotEventable
    // The Cart To Class Associations.
    cartToClassAssociations []CartToClassAssociationable
    // The available categories
    categories []DeviceManagementSettingCategoryable
    // Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
    certificateConnectorDetails []CertificateConnectorDetailsable
    // Collection of ChromeOSOnboardingSettings settings associated with account.
    chromeOSOnboardingSettings []ChromeOSOnboardingSettingsable
    // The list of CloudPC Connectivity Issue.
    cloudPCConnectivityIssues []CloudPCConnectivityIssueable
    // The list of co-managed devices report
    comanagedDevices []ManagedDeviceable
    // The list of co-management eligible devices report
    comanagementEligibleDevices []ComanagementEligibleDeviceable
    // List of all compliance categories
    complianceCategories []DeviceManagementConfigurationCategoryable
    // The list of Compliance Management Partners configured by the tenant.
    complianceManagementPartners []ComplianceManagementPartnerable
    // List of all compliance policies
    compliancePolicies []DeviceManagementCompliancePolicyable
    // List of all ComplianceSettings
    complianceSettings []DeviceManagementConfigurationSettingDefinitionable
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings OnPremisesConditionalAccessSettingsable
    // A list of ConfigManagerCollection
    configManagerCollections []ConfigManagerCollectionable
    // List of all Configuration Categories
    configurationCategories []DeviceManagementConfigurationCategoryable
    // List of all Configuration policies
    configurationPolicies []DeviceManagementConfigurationPolicyable
    // List of all templates
    configurationPolicyTemplates []DeviceManagementConfigurationPolicyTemplateable
    // List of all ConfigurationSettings
    configurationSettings []DeviceManagementConfigurationSettingDefinitionable
    // A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
    dataProcessorServiceForWindowsFeaturesOnboarding DataProcessorServiceForWindowsFeaturesOnboardingable
    // Data sharing consents.
    dataSharingConsents []DataSharingConsentable
    // This collections of multiple DEP tokens per-tenant.
    depOnboardingSettings []DepOnboardingSettingable
    // Collection of Derived credential settings associated with account.
    derivedCredentials []DeviceManagementDerivedCredentialSettingsable
    // The list of detected apps associated with a device.
    detectedApps []DetectedAppable
    // The list of device categories with the tenant.
    deviceCategories []DeviceCategoryable
    // The device compliance policies.
    deviceCompliancePolicies []DeviceCompliancePolicyable
    // The device compliance state summary for this account.
    deviceCompliancePolicyDeviceStateSummary DeviceCompliancePolicyDeviceStateSummaryable
    // The summary states of compliance policy settings for this account.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummaryable
    // The last requested time of device compliance reporting for this account. This property is read-only.
    deviceComplianceReportSummarizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The list of device compliance scripts associated with the tenant.
    deviceComplianceScripts []DeviceComplianceScriptable
    // Summary of policies in conflict state for this account.
    deviceConfigurationConflictSummary []DeviceConfigurationConflictSummaryable
    // The device configuration device state summary for this account.
    deviceConfigurationDeviceStateSummaries DeviceConfigurationDeviceStateSummaryable
    // Restricted apps violations for this account.
    deviceConfigurationRestrictedAppsViolations []RestrictedAppsViolationable
    // The device configurations.
    deviceConfigurations []DeviceConfigurationable
    // Summary of all certificates for all devices.
    deviceConfigurationsAllManagedDeviceCertificateStates []ManagedAllDeviceCertificateStateable
    // The device configuration user state summary for this account.
    deviceConfigurationUserStateSummaries DeviceConfigurationUserStateSummaryable
    // The list of device custom attribute shell scripts associated with the tenant.
    deviceCustomAttributeShellScripts []DeviceCustomAttributeShellScriptable
    // The list of device enrollment configurations
    deviceEnrollmentConfigurations []DeviceEnrollmentConfigurationable
    // The list of device health scripts associated with the tenant.
    deviceHealthScripts []DeviceHealthScriptable
    // The list of Device Management Partners configured by the tenant.
    deviceManagementPartners []DeviceManagementPartnerable
    // The list of device management scripts associated with the tenant.
    deviceManagementScripts []DeviceManagementScriptable
    // Device protection overview.
    deviceProtectionOverview DeviceProtectionOverviewable
    // The list of device shell scripts associated with the tenant.
    deviceShellScripts []DeviceShellScriptable
    // A list of connector objects.
    domainJoinConnectors []DeviceManagementDomainJoinConnectorable
    // The embedded SIM activation code pools created by this account.
    embeddedSIMActivationCodePools []EmbeddedSIMActivationCodePoolable
    // The list of Exchange Connectors configured by the tenant.
    exchangeConnectors []DeviceManagementExchangeConnectorable
    // The list of Exchange On Premisis policies configured by the tenant.
    exchangeOnPremisesPolicies []DeviceManagementExchangeOnPremisesPolicyable
    // The policy which controls mobile device access to Exchange On Premises
    exchangeOnPremisesPolicy DeviceManagementExchangeOnPremisesPolicyable
    // The available group policy categories for this account.
    groupPolicyCategories []GroupPolicyCategoryable
    // The group policy configurations created by this account.
    groupPolicyConfigurations []GroupPolicyConfigurationable
    // The available group policy definition files for this account.
    groupPolicyDefinitionFiles []GroupPolicyDefinitionFileable
    // The available group policy definitions for this account.
    groupPolicyDefinitions []GroupPolicyDefinitionable
    // A list of Group Policy migration reports.
    groupPolicyMigrationReports []GroupPolicyMigrationReportable
    // A list of Group Policy Object files uploaded.
    groupPolicyObjectFiles []GroupPolicyObjectFileable
    // The available group policy uploaded definition files for this account.
    groupPolicyUploadedDefinitionFiles []GroupPolicyUploadedDefinitionFileable
    // The imported device identities.
    importedDeviceIdentities []ImportedDeviceIdentityable
    // Collection of imported Windows autopilot devices.
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentityable
    // The device management intents
    intents []DeviceManagementIntentable
    // Intune Account ID for given tenant
    intuneAccountId *string
    // intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
    intuneBrand IntuneBrandable
    // Intune branding profiles targeted to AAD groups
    intuneBrandingProfiles []IntuneBrandingProfileable
    // The IOS software update installation statuses for this account.
    iosUpdateStatuses []IosUpdateDeviceStatusable
    // The last modified time of reporting for this account. This property is read-only.
    lastReportAggregationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
    legacyPcManangementEnabled *bool
    // The MacOS software update account summaries for this account.
    macOSSoftwareUpdateAccountSummaries []MacOSSoftwareUpdateAccountSummaryable
    // Device cleanup rule
    managedDeviceCleanupSettings ManagedDeviceCleanupSettingsable
    // Encryption report for devices in this account
    managedDeviceEncryptionStates []ManagedDeviceEncryptionStateable
    // Device overview
    managedDeviceOverview ManagedDeviceOverviewable
    // The list of managed devices.
    managedDevices []ManagedDeviceable
    // Maximum number of DEP tokens allowed per-tenant.
    maximumDepTokens *int32
    // Collection of MicrosoftTunnelConfiguration settings associated with account.
    microsoftTunnelConfigurations []MicrosoftTunnelConfigurationable
    // Collection of MicrosoftTunnelHealthThreshold settings associated with account.
    microsoftTunnelHealthThresholds []MicrosoftTunnelHealthThresholdable
    // Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
    microsoftTunnelServerLogCollectionResponses []MicrosoftTunnelServerLogCollectionResponseable
    // Collection of MicrosoftTunnelSite settings associated with account.
    microsoftTunnelSites []MicrosoftTunnelSiteable
    // The collection property of MobileAppTroubleshootingEvent.
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEventable
    // The list of Mobile threat Defense connectors configured by the tenant.
    mobileThreatDefenseConnectors []MobileThreatDefenseConnectorable
    // The collection of Ndes connectors for this account.
    ndesConnectors []NdesConnectorable
    // The Notification Message Templates.
    notificationMessageTemplates []NotificationMessageTemplateable
    // List of OEM Warranty Statuses
    oemWarrantyInformationOnboarding []OemWarrantyInformationOnboardingable
    // A list of OrganizationalMessageDetails
    organizationalMessageDetails []OrganizationalMessageDetailable
    // A list of OrganizationalMessageGuidedContents
    organizationalMessageGuidedContents []OrganizationalMessageGuidedContentable
    // The list of device remote action audits with the tenant.
    remoteActionAudits []RemoteActionAuditable
    // The remote assist partners.
    remoteAssistancePartners []RemoteAssistancePartnerable
    // The remote assistance settings singleton
    remoteAssistanceSettings RemoteAssistanceSettingsable
    // Reports singleton
    reports DeviceManagementReportsable
    // Collection of resource access settings associated with account.
    resourceAccessProfiles []DeviceManagementResourceAccessProfileBaseable
    // The Resource Operations.
    resourceOperations []ResourceOperationable
    // List of all reusable settings that can be referred in a policy
    reusablePolicySettings []DeviceManagementReusablePolicySettingable
    // List of all reusable settings
    reusableSettings []DeviceManagementConfigurationSettingDefinitionable
    // The Role Assignments.
    roleAssignments []DeviceAndAppManagementRoleAssignmentable
    // The Role Definitions.
    roleDefinitions []RoleDefinitionable
    // The Role Scope Tags.
    roleScopeTags []RoleScopeTagable
    // The device management intent setting definitions
    settingDefinitions []DeviceManagementSettingDefinitionable
    // Account level settings.
    settings DeviceManagementSettingsable
    // The software update status summary.
    softwareUpdateStatusSummary SoftwareUpdateStatusSummaryable
    // Tenant mobile device management subscriptions.
    subscriptions *DeviceManagementSubscriptions
    // Tenant mobile device management subscription state.
    subscriptionState *DeviceManagementSubscriptionState
    // The telecom expense management partners.
    telecomExpenseManagementPartners []TelecomExpenseManagementPartnerable
    // The available templates
    templates []DeviceManagementTemplateable
    // List of all TemplateSettings
    templateSettings []DeviceManagementConfigurationSettingTemplateable
    // TenantAttach RBAC Enablement
    tenantAttachRBAC TenantAttachRBACable
    // The terms and conditions associated with device management of the company.
    termsAndConditions []TermsAndConditionsable
    // The list of troubleshooting events for the tenant.
    troubleshootingEvents []DeviceManagementTroubleshootingEventable
    // When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
    unlicensedAdminstratorsEnabled *bool
    // The user experience analytics anomaly entity contains anomaly details.
    userExperienceAnalyticsAnomaly []UserExperienceAnalyticsAnomalyable
    // The user experience analytics anomaly entity contains device details.
    userExperienceAnalyticsAnomalyDevice []UserExperienceAnalyticsAnomalyDeviceable
    // The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
    userExperienceAnalyticsAnomalySeverityOverview UserExperienceAnalyticsAnomalySeverityOverviewable
    // User experience analytics appHealth Application Performance
    userExperienceAnalyticsAppHealthApplicationPerformance []UserExperienceAnalyticsAppHealthApplicationPerformanceable
    // User experience analytics appHealth Application Performance by App Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable
    // User experience analytics appHealth Application Performance by App Version details
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable
    // User experience analytics appHealth Application Performance by App Version Device Id
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable
    // User experience analytics appHealth Application Performance by OS Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable
    // User experience analytics appHealth Model Performance
    userExperienceAnalyticsAppHealthDeviceModelPerformance []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable
    // User experience analytics appHealth Device Performance
    userExperienceAnalyticsAppHealthDevicePerformance []UserExperienceAnalyticsAppHealthDevicePerformanceable
    // User experience analytics device performance details
    userExperienceAnalyticsAppHealthDevicePerformanceDetails []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable
    // User experience analytics appHealth OS version Performance
    userExperienceAnalyticsAppHealthOSVersionPerformance []UserExperienceAnalyticsAppHealthOSVersionPerformanceable
    // User experience analytics appHealth overview
    userExperienceAnalyticsAppHealthOverview UserExperienceAnalyticsCategoryable
    // User experience analytics baselines
    userExperienceAnalyticsBaselines []UserExperienceAnalyticsBaselineable
    // User Experience Analytics Battery Health App Impact
    userExperienceAnalyticsBatteryHealthAppImpact []UserExperienceAnalyticsBatteryHealthAppImpactable
    // User Experience Analytics Battery Health Capacity Details
    userExperienceAnalyticsBatteryHealthCapacityDetails UserExperienceAnalyticsBatteryHealthCapacityDetailsable
    // User Experience Analytics Battery Health Device App Impact
    userExperienceAnalyticsBatteryHealthDeviceAppImpact []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable
    // User Experience Analytics Battery Health Device Performance
    userExperienceAnalyticsBatteryHealthDevicePerformance []UserExperienceAnalyticsBatteryHealthDevicePerformanceable
    // User Experience Analytics Battery Health Device Runtime History
    userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable
    // User Experience Analytics Battery Health Model Performance
    userExperienceAnalyticsBatteryHealthModelPerformance []UserExperienceAnalyticsBatteryHealthModelPerformanceable
    // User Experience Analytics Battery Health Os Performance
    userExperienceAnalyticsBatteryHealthOsPerformance []UserExperienceAnalyticsBatteryHealthOsPerformanceable
    // User Experience Analytics Battery Health Runtime Details
    userExperienceAnalyticsBatteryHealthRuntimeDetails UserExperienceAnalyticsBatteryHealthRuntimeDetailsable
    // User experience analytics categories
    userExperienceAnalyticsCategories []UserExperienceAnalyticsCategoryable
    // User experience analytics device metric history
    userExperienceAnalyticsDeviceMetricHistory []UserExperienceAnalyticsMetricHistoryable
    // User experience analytics device performance
    userExperienceAnalyticsDevicePerformance []UserExperienceAnalyticsDevicePerformanceable
    // The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
    userExperienceAnalyticsDeviceScope UserExperienceAnalyticsDeviceScopeable
    // The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
    userExperienceAnalyticsDeviceScopes []UserExperienceAnalyticsDeviceScopeable
    // User experience analytics device scores
    userExperienceAnalyticsDeviceScores []UserExperienceAnalyticsDeviceScoresable
    // User experience analytics device Startup History
    userExperienceAnalyticsDeviceStartupHistory []UserExperienceAnalyticsDeviceStartupHistoryable
    // User experience analytics device Startup Processes
    userExperienceAnalyticsDeviceStartupProcesses []UserExperienceAnalyticsDeviceStartupProcessable
    // User experience analytics device Startup Process Performance
    userExperienceAnalyticsDeviceStartupProcessPerformance []UserExperienceAnalyticsDeviceStartupProcessPerformanceable
    // User experience analytics devices without cloud identity.
    userExperienceAnalyticsDevicesWithoutCloudIdentity []UserExperienceAnalyticsDeviceWithoutCloudIdentityable
    // User experience analytics impacting process
    userExperienceAnalyticsImpactingProcess []UserExperienceAnalyticsImpactingProcessable
    // User experience analytics metric history
    userExperienceAnalyticsMetricHistory []UserExperienceAnalyticsMetricHistoryable
    // User experience analytics model scores
    userExperienceAnalyticsModelScores []UserExperienceAnalyticsModelScoresable
    // User experience analytics devices not Windows Autopilot ready.
    userExperienceAnalyticsNotAutopilotReadyDevice []UserExperienceAnalyticsNotAutopilotReadyDeviceable
    // User experience analytics overview
    userExperienceAnalyticsOverview UserExperienceAnalyticsOverviewable
    // User experience analytics regression summary
    userExperienceAnalyticsRegressionSummary UserExperienceAnalyticsRegressionSummaryable
    // User experience analytics remote connection
    userExperienceAnalyticsRemoteConnection []UserExperienceAnalyticsRemoteConnectionable
    // User experience analytics resource performance
    userExperienceAnalyticsResourcePerformance []UserExperienceAnalyticsResourcePerformanceable
    // User experience analytics device Startup Score History
    userExperienceAnalyticsScoreHistory []UserExperienceAnalyticsScoreHistoryable
    // User experience analytics device settings
    userExperienceAnalyticsSettings UserExperienceAnalyticsSettingsable
    // User experience analytics work from anywhere hardware readiness metrics.
    userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable
    // User experience analytics work from anywhere metrics.
    userExperienceAnalyticsWorkFromAnywhereMetrics []UserExperienceAnalyticsWorkFromAnywhereMetricable
    // The user experience analytics work from anywhere model performance
    userExperienceAnalyticsWorkFromAnywhereModelPerformance []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable
    // Collection of PFX certificates associated with a user.
    userPfxCertificates []UserPFXCertificateable
    // The virtualEndpoint property
    virtualEndpoint VirtualEndpointable
    // Windows auto pilot deployment profiles
    windowsAutopilotDeploymentProfiles []WindowsAutopilotDeploymentProfileable
    // The Windows autopilot device identities contained collection.
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentityable
    // The Windows autopilot account settings.
    windowsAutopilotSettings WindowsAutopilotSettingsable
    // A collection of windows driver update profiles
    windowsDriverUpdateProfiles []WindowsDriverUpdateProfileable
    // A collection of windows feature update profiles
    windowsFeatureUpdateProfiles []WindowsFeatureUpdateProfileable
    // The windows information protection app learning summaries.
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummaryable
    // The windows information protection network learning summaries.
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummaryable
    // The list of affected malware in the tenant.
    windowsMalwareInformation []WindowsMalwareInformationable
    // Malware overview for windows devices.
    windowsMalwareOverview WindowsMalwareOverviewable
    // A collection of windows quality update profiles
    windowsQualityUpdateProfiles []WindowsQualityUpdateProfileable
    // A collection of windows update catalog items (fetaure updates item , quality updates item)
    windowsUpdateCatalogItems []WindowsUpdateCatalogItemable
    // The Collection of ZebraFotaArtifacts.
    zebraFotaArtifacts []ZebraFotaArtifactable
    // The singleton ZebraFotaConnector associated with account.
    zebraFotaConnector ZebraFotaConnectorable
    // Collection of ZebraFotaDeployments associated with account.
    zebraFotaDeployments []ZebraFotaDeploymentable
}
// NewDeviceManagement instantiates a new DeviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagement";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagement(), nil
}
// GetAccountMoveCompletionDateTime gets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.accountMoveCompletionDateTime
}
// GetAdminConsent gets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) GetAdminConsent()(AdminConsentable) {
    return m.adminConsent
}
// GetAdvancedThreatProtectionOnboardingStateSummary gets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) GetAdvancedThreatProtectionOnboardingStateSummary()(AdvancedThreatProtectionOnboardingStateSummaryable) {
    return m.advancedThreatProtectionOnboardingStateSummary
}
// GetAndroidDeviceOwnerEnrollmentProfiles gets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfileable) {
    return m.androidDeviceOwnerEnrollmentProfiles
}
// GetAndroidForWorkAppConfigurationSchemas gets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchemaable) {
    return m.androidForWorkAppConfigurationSchemas
}
// GetAndroidForWorkEnrollmentProfiles gets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfileable) {
    return m.androidForWorkEnrollmentProfiles
}
// GetAndroidForWorkSettings gets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) GetAndroidForWorkSettings()(AndroidForWorkSettingsable) {
    return m.androidForWorkSettings
}
// GetAndroidManagedStoreAccountEnterpriseSettings gets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) GetAndroidManagedStoreAccountEnterpriseSettings()(AndroidManagedStoreAccountEnterpriseSettingsable) {
    return m.androidManagedStoreAccountEnterpriseSettings
}
// GetAndroidManagedStoreAppConfigurationSchemas gets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchemaable) {
    return m.androidManagedStoreAppConfigurationSchemas
}
// GetApplePushNotificationCertificate gets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable) {
    return m.applePushNotificationCertificate
}
// GetAppleUserInitiatedEnrollmentProfiles gets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfileable) {
    return m.appleUserInitiatedEnrollmentProfiles
}
// GetAssignmentFilters gets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilterable) {
    return m.assignmentFilters
}
// GetAuditEvents gets the auditEvents property value. The Audit Events
func (m *DeviceManagement) GetAuditEvents()([]AuditEventable) {
    return m.auditEvents
}
// GetAutopilotEvents gets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) GetAutopilotEvents()([]DeviceManagementAutopilotEventable) {
    return m.autopilotEvents
}
// GetCartToClassAssociations gets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) GetCartToClassAssociations()([]CartToClassAssociationable) {
    return m.cartToClassAssociations
}
// GetCategories gets the categories property value. The available categories
func (m *DeviceManagement) GetCategories()([]DeviceManagementSettingCategoryable) {
    return m.categories
}
// GetCertificateConnectorDetails gets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) GetCertificateConnectorDetails()([]CertificateConnectorDetailsable) {
    return m.certificateConnectorDetails
}
// GetChromeOSOnboardingSettings gets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettingsable) {
    return m.chromeOSOnboardingSettings
}
// GetCloudPCConnectivityIssues gets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssueable) {
    return m.cloudPCConnectivityIssues
}
// GetComanagedDevices gets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) GetComanagedDevices()([]ManagedDeviceable) {
    return m.comanagedDevices
}
// GetComanagementEligibleDevices gets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) GetComanagementEligibleDevices()([]ComanagementEligibleDeviceable) {
    return m.comanagementEligibleDevices
}
// GetComplianceCategories gets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) GetComplianceCategories()([]DeviceManagementConfigurationCategoryable) {
    return m.complianceCategories
}
// GetComplianceManagementPartners gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartnerable) {
    return m.complianceManagementPartners
}
// GetCompliancePolicies gets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) GetCompliancePolicies()([]DeviceManagementCompliancePolicyable) {
    return m.compliancePolicies
}
// GetComplianceSettings gets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.complianceSettings
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable) {
    return m.conditionalAccessSettings
}
// GetConfigManagerCollections gets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) GetConfigManagerCollections()([]ConfigManagerCollectionable) {
    return m.configManagerCollections
}
// GetConfigurationCategories gets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) GetConfigurationCategories()([]DeviceManagementConfigurationCategoryable) {
    return m.configurationCategories
}
// GetConfigurationPolicies gets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) GetConfigurationPolicies()([]DeviceManagementConfigurationPolicyable) {
    return m.configurationPolicies
}
// GetConfigurationPolicyTemplates gets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplateable) {
    return m.configurationPolicyTemplates
}
// GetConfigurationSettings gets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.configurationSettings
}
// GetDataProcessorServiceForWindowsFeaturesOnboarding gets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) GetDataProcessorServiceForWindowsFeaturesOnboarding()(DataProcessorServiceForWindowsFeaturesOnboardingable) {
    return m.dataProcessorServiceForWindowsFeaturesOnboarding
}
// GetDataSharingConsents gets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) GetDataSharingConsents()([]DataSharingConsentable) {
    return m.dataSharingConsents
}
// GetDepOnboardingSettings gets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) GetDepOnboardingSettings()([]DepOnboardingSettingable) {
    return m.depOnboardingSettings
}
// GetDerivedCredentials gets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettingsable) {
    return m.derivedCredentials
}
// GetDetectedApps gets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) GetDetectedApps()([]DetectedAppable) {
    return m.detectedApps
}
// GetDeviceCategories gets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategoryable) {
    return m.deviceCategories
}
// GetDeviceCompliancePolicies gets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable) {
    return m.deviceCompliancePolicies
}
// GetDeviceCompliancePolicyDeviceStateSummary gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable) {
    return m.deviceCompliancePolicyDeviceStateSummary
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    return m.deviceCompliancePolicySettingStateSummaries
}
// GetDeviceComplianceReportSummarizationDateTime gets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deviceComplianceReportSummarizationDateTime
}
// GetDeviceComplianceScripts gets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceComplianceScripts()([]DeviceComplianceScriptable) {
    return m.deviceComplianceScripts
}
// GetDeviceConfigurationConflictSummary gets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummaryable) {
    return m.deviceConfigurationConflictSummary
}
// GetDeviceConfigurationDeviceStateSummaries gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable) {
    return m.deviceConfigurationDeviceStateSummaries
}
// GetDeviceConfigurationRestrictedAppsViolations gets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolationable) {
    return m.deviceConfigurationRestrictedAppsViolations
}
// GetDeviceConfigurations gets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfigurationable) {
    return m.deviceConfigurations
}
// GetDeviceConfigurationsAllManagedDeviceCertificateStates gets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateStateable) {
    return m.deviceConfigurationsAllManagedDeviceCertificateStates
}
// GetDeviceConfigurationUserStateSummaries gets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationUserStateSummaries()(DeviceConfigurationUserStateSummaryable) {
    return m.deviceConfigurationUserStateSummaries
}
// GetDeviceCustomAttributeShellScripts gets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScriptable) {
    return m.deviceCustomAttributeShellScripts
}
// GetDeviceEnrollmentConfigurations gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable) {
    return m.deviceEnrollmentConfigurations
}
// GetDeviceHealthScripts gets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceHealthScripts()([]DeviceHealthScriptable) {
    return m.deviceHealthScripts
}
// GetDeviceManagementPartners gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartnerable) {
    return m.deviceManagementPartners
}
// GetDeviceManagementScripts gets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceManagementScripts()([]DeviceManagementScriptable) {
    return m.deviceManagementScripts
}
// GetDeviceProtectionOverview gets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) GetDeviceProtectionOverview()(DeviceProtectionOverviewable) {
    return m.deviceProtectionOverview
}
// GetDeviceShellScripts gets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceShellScripts()([]DeviceShellScriptable) {
    return m.deviceShellScripts
}
// GetDomainJoinConnectors gets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnectorable) {
    return m.domainJoinConnectors
}
// GetEmbeddedSIMActivationCodePools gets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePoolable) {
    return m.embeddedSIMActivationCodePools
}
// GetExchangeConnectors gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnectorable) {
    return m.exchangeConnectors
}
// GetExchangeOnPremisesPolicies gets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicyable) {
    return m.exchangeOnPremisesPolicies
}
// GetExchangeOnPremisesPolicy gets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) GetExchangeOnPremisesPolicy()(DeviceManagementExchangeOnPremisesPolicyable) {
    return m.exchangeOnPremisesPolicy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountMoveCompletionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetAccountMoveCompletionDateTime)
    res["adminConsent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAdminConsentFromDiscriminatorValue , m.SetAdminConsent)
    res["advancedThreatProtectionOnboardingStateSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue , m.SetAdvancedThreatProtectionOnboardingStateSummary)
    res["androidDeviceOwnerEnrollmentProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue , m.SetAndroidDeviceOwnerEnrollmentProfiles)
    res["androidForWorkAppConfigurationSchemas"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue , m.SetAndroidForWorkAppConfigurationSchemas)
    res["androidForWorkEnrollmentProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue , m.SetAndroidForWorkEnrollmentProfiles)
    res["androidForWorkSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidForWorkSettingsFromDiscriminatorValue , m.SetAndroidForWorkSettings)
    res["androidManagedStoreAccountEnterpriseSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue , m.SetAndroidManagedStoreAccountEnterpriseSettings)
    res["androidManagedStoreAppConfigurationSchemas"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue , m.SetAndroidManagedStoreAppConfigurationSchemas)
    res["applePushNotificationCertificate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateApplePushNotificationCertificateFromDiscriminatorValue , m.SetApplePushNotificationCertificate)
    res["appleUserInitiatedEnrollmentProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue , m.SetAppleUserInitiatedEnrollmentProfiles)
    res["assignmentFilters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue , m.SetAssignmentFilters)
    res["auditEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue , m.SetAuditEvents)
    res["autopilotEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementAutopilotEventFromDiscriminatorValue , m.SetAutopilotEvents)
    res["cartToClassAssociations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCartToClassAssociationFromDiscriminatorValue , m.SetCartToClassAssociations)
    res["categories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementSettingCategoryFromDiscriminatorValue , m.SetCategories)
    res["certificateConnectorDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCertificateConnectorDetailsFromDiscriminatorValue , m.SetCertificateConnectorDetails)
    res["chromeOSOnboardingSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateChromeOSOnboardingSettingsFromDiscriminatorValue , m.SetChromeOSOnboardingSettings)
    res["cloudPCConnectivityIssues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPCConnectivityIssueFromDiscriminatorValue , m.SetCloudPCConnectivityIssues)
    res["comanagedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue , m.SetComanagedDevices)
    res["comanagementEligibleDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateComanagementEligibleDeviceFromDiscriminatorValue , m.SetComanagementEligibleDevices)
    res["complianceCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue , m.SetComplianceCategories)
    res["complianceManagementPartners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateComplianceManagementPartnerFromDiscriminatorValue , m.SetComplianceManagementPartners)
    res["compliancePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementCompliancePolicyFromDiscriminatorValue , m.SetCompliancePolicies)
    res["complianceSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue , m.SetComplianceSettings)
    res["conditionalAccessSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue , m.SetConditionalAccessSettings)
    res["configManagerCollections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConfigManagerCollectionFromDiscriminatorValue , m.SetConfigManagerCollections)
    res["configurationCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue , m.SetConfigurationCategories)
    res["configurationPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue , m.SetConfigurationPolicies)
    res["configurationPolicyTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue , m.SetConfigurationPolicyTemplates)
    res["configurationSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue , m.SetConfigurationSettings)
    res["dataProcessorServiceForWindowsFeaturesOnboarding"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue , m.SetDataProcessorServiceForWindowsFeaturesOnboarding)
    res["dataSharingConsents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDataSharingConsentFromDiscriminatorValue , m.SetDataSharingConsents)
    res["depOnboardingSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDepOnboardingSettingFromDiscriminatorValue , m.SetDepOnboardingSettings)
    res["derivedCredentials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue , m.SetDerivedCredentials)
    res["detectedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue , m.SetDetectedApps)
    res["deviceCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCategoryFromDiscriminatorValue , m.SetDeviceCategories)
    res["deviceCompliancePolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicyFromDiscriminatorValue , m.SetDeviceCompliancePolicies)
    res["deviceCompliancePolicyDeviceStateSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue , m.SetDeviceCompliancePolicyDeviceStateSummary)
    res["deviceCompliancePolicySettingStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue , m.SetDeviceCompliancePolicySettingStateSummaries)
    res["deviceComplianceReportSummarizationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDeviceComplianceReportSummarizationDateTime)
    res["deviceComplianceScripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceComplianceScriptFromDiscriminatorValue , m.SetDeviceComplianceScripts)
    res["deviceConfigurationConflictSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue , m.SetDeviceConfigurationConflictSummary)
    res["deviceConfigurationDeviceStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue , m.SetDeviceConfigurationDeviceStateSummaries)
    res["deviceConfigurationRestrictedAppsViolations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRestrictedAppsViolationFromDiscriminatorValue , m.SetDeviceConfigurationRestrictedAppsViolations)
    res["deviceConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceConfigurationFromDiscriminatorValue , m.SetDeviceConfigurations)
    res["deviceConfigurationsAllManagedDeviceCertificateStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedAllDeviceCertificateStateFromDiscriminatorValue , m.SetDeviceConfigurationsAllManagedDeviceCertificateStates)
    res["deviceConfigurationUserStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue , m.SetDeviceConfigurationUserStateSummaries)
    res["deviceCustomAttributeShellScripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue , m.SetDeviceCustomAttributeShellScripts)
    res["deviceEnrollmentConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue , m.SetDeviceEnrollmentConfigurations)
    res["deviceHealthScripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceHealthScriptFromDiscriminatorValue , m.SetDeviceHealthScripts)
    res["deviceManagementPartners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementPartnerFromDiscriminatorValue , m.SetDeviceManagementPartners)
    res["deviceManagementScripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementScriptFromDiscriminatorValue , m.SetDeviceManagementScripts)
    res["deviceProtectionOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceProtectionOverviewFromDiscriminatorValue , m.SetDeviceProtectionOverview)
    res["deviceShellScripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceShellScriptFromDiscriminatorValue , m.SetDeviceShellScripts)
    res["domainJoinConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue , m.SetDomainJoinConnectors)
    res["embeddedSIMActivationCodePools"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue , m.SetEmbeddedSIMActivationCodePools)
    res["exchangeConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementExchangeConnectorFromDiscriminatorValue , m.SetExchangeConnectors)
    res["exchangeOnPremisesPolicies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue , m.SetExchangeOnPremisesPolicies)
    res["exchangeOnPremisesPolicy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue , m.SetExchangeOnPremisesPolicy)
    res["groupPolicyCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyCategoryFromDiscriminatorValue , m.SetGroupPolicyCategories)
    res["groupPolicyConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyConfigurationFromDiscriminatorValue , m.SetGroupPolicyConfigurations)
    res["groupPolicyDefinitionFiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyDefinitionFileFromDiscriminatorValue , m.SetGroupPolicyDefinitionFiles)
    res["groupPolicyDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue , m.SetGroupPolicyDefinitions)
    res["groupPolicyMigrationReports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyMigrationReportFromDiscriminatorValue , m.SetGroupPolicyMigrationReports)
    res["groupPolicyObjectFiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyObjectFileFromDiscriminatorValue , m.SetGroupPolicyObjectFiles)
    res["groupPolicyUploadedDefinitionFiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue , m.SetGroupPolicyUploadedDefinitionFiles)
    res["importedDeviceIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateImportedDeviceIdentityFromDiscriminatorValue , m.SetImportedDeviceIdentities)
    res["importedWindowsAutopilotDeviceIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue , m.SetImportedWindowsAutopilotDeviceIdentities)
    res["intents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementIntentFromDiscriminatorValue , m.SetIntents)
    res["intuneAccountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIntuneAccountId)
    res["intuneBrand"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIntuneBrandFromDiscriminatorValue , m.SetIntuneBrand)
    res["intuneBrandingProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIntuneBrandingProfileFromDiscriminatorValue , m.SetIntuneBrandingProfiles)
    res["iosUpdateStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIosUpdateDeviceStatusFromDiscriminatorValue , m.SetIosUpdateStatuses)
    res["lastReportAggregationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastReportAggregationDateTime)
    res["legacyPcManangementEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetLegacyPcManangementEnabled)
    res["macOSSoftwareUpdateAccountSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue , m.SetMacOSSoftwareUpdateAccountSummaries)
    res["managedDeviceCleanupSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedDeviceCleanupSettingsFromDiscriminatorValue , m.SetManagedDeviceCleanupSettings)
    res["managedDeviceEncryptionStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceEncryptionStateFromDiscriminatorValue , m.SetManagedDeviceEncryptionStates)
    res["managedDeviceOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagedDeviceOverviewFromDiscriminatorValue , m.SetManagedDeviceOverview)
    res["managedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue , m.SetManagedDevices)
    res["maximumDepTokens"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaximumDepTokens)
    res["microsoftTunnelConfigurations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMicrosoftTunnelConfigurationFromDiscriminatorValue , m.SetMicrosoftTunnelConfigurations)
    res["microsoftTunnelHealthThresholds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue , m.SetMicrosoftTunnelHealthThresholds)
    res["microsoftTunnelServerLogCollectionResponses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue , m.SetMicrosoftTunnelServerLogCollectionResponses)
    res["microsoftTunnelSites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMicrosoftTunnelSiteFromDiscriminatorValue , m.SetMicrosoftTunnelSites)
    res["mobileAppTroubleshootingEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppTroubleshootingEventFromDiscriminatorValue , m.SetMobileAppTroubleshootingEvents)
    res["mobileThreatDefenseConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileThreatDefenseConnectorFromDiscriminatorValue , m.SetMobileThreatDefenseConnectors)
    res["ndesConnectors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNdesConnectorFromDiscriminatorValue , m.SetNdesConnectors)
    res["notificationMessageTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateNotificationMessageTemplateFromDiscriminatorValue , m.SetNotificationMessageTemplates)
    res["oemWarrantyInformationOnboarding"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOemWarrantyInformationOnboardingFromDiscriminatorValue , m.SetOemWarrantyInformationOnboarding)
    res["organizationalMessageDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOrganizationalMessageDetailFromDiscriminatorValue , m.SetOrganizationalMessageDetails)
    res["organizationalMessageGuidedContents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOrganizationalMessageGuidedContentFromDiscriminatorValue , m.SetOrganizationalMessageGuidedContents)
    res["remoteActionAudits"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRemoteActionAuditFromDiscriminatorValue , m.SetRemoteActionAudits)
    res["remoteAssistancePartners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRemoteAssistancePartnerFromDiscriminatorValue , m.SetRemoteAssistancePartners)
    res["remoteAssistanceSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRemoteAssistanceSettingsFromDiscriminatorValue , m.SetRemoteAssistanceSettings)
    res["reports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementReportsFromDiscriminatorValue , m.SetReports)
    res["resourceAccessProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue , m.SetResourceAccessProfiles)
    res["resourceOperations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateResourceOperationFromDiscriminatorValue , m.SetResourceOperations)
    res["reusablePolicySettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue , m.SetReusablePolicySettings)
    res["reusableSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue , m.SetReusableSettings)
    res["roleAssignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue , m.SetRoleAssignments)
    res["roleDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRoleDefinitionFromDiscriminatorValue , m.SetRoleDefinitions)
    res["roleScopeTags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRoleScopeTagFromDiscriminatorValue , m.SetRoleScopeTags)
    res["settingDefinitions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementSettingDefinitionFromDiscriminatorValue , m.SetSettingDefinitions)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementSettingsFromDiscriminatorValue , m.SetSettings)
    res["softwareUpdateStatusSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue , m.SetSoftwareUpdateStatusSummary)
    res["subscriptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementSubscriptions , m.SetSubscriptions)
    res["subscriptionState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementSubscriptionState , m.SetSubscriptionState)
    res["telecomExpenseManagementPartners"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTelecomExpenseManagementPartnerFromDiscriminatorValue , m.SetTelecomExpenseManagementPartners)
    res["templates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementTemplateFromDiscriminatorValue , m.SetTemplates)
    res["templateSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue , m.SetTemplateSettings)
    res["tenantAttachRBAC"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTenantAttachRBACFromDiscriminatorValue , m.SetTenantAttachRBAC)
    res["termsAndConditions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTermsAndConditionsFromDiscriminatorValue , m.SetTermsAndConditions)
    res["troubleshootingEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue , m.SetTroubleshootingEvents)
    res["unlicensedAdminstratorsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetUnlicensedAdminstratorsEnabled)
    res["userExperienceAnalyticsAnomaly"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue , m.SetUserExperienceAnalyticsAnomaly)
    res["userExperienceAnalyticsAnomalyDevice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue , m.SetUserExperienceAnalyticsAnomalyDevice)
    res["userExperienceAnalyticsAnomalySeverityOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue , m.SetUserExperienceAnalyticsAnomalySeverityOverview)
    res["userExperienceAnalyticsAppHealthApplicationPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthApplicationPerformance)
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion)
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails)
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId)
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion)
    res["userExperienceAnalyticsAppHealthDeviceModelPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthDeviceModelPerformance)
    res["userExperienceAnalyticsAppHealthDevicePerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthDevicePerformance)
    res["userExperienceAnalyticsAppHealthDevicePerformanceDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails)
    res["userExperienceAnalyticsAppHealthOSVersionPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthOSVersionPerformance)
    res["userExperienceAnalyticsAppHealthOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsAppHealthOverview)
    res["userExperienceAnalyticsBaselines"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue , m.SetUserExperienceAnalyticsBaselines)
    res["userExperienceAnalyticsBatteryHealthAppImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthAppImpactFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthAppImpact)
    res["userExperienceAnalyticsBatteryHealthCapacityDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsBatteryHealthCapacityDetailsFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthCapacityDetails)
    res["userExperienceAnalyticsBatteryHealthDeviceAppImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact)
    res["userExperienceAnalyticsBatteryHealthDevicePerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthDevicePerformance)
    res["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory)
    res["userExperienceAnalyticsBatteryHealthModelPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthModelPerformance)
    res["userExperienceAnalyticsBatteryHealthOsPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthOsPerformance)
    res["userExperienceAnalyticsBatteryHealthRuntimeDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsBatteryHealthRuntimeDetailsFromDiscriminatorValue , m.SetUserExperienceAnalyticsBatteryHealthRuntimeDetails)
    res["userExperienceAnalyticsCategories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsCategories)
    res["userExperienceAnalyticsDeviceMetricHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceMetricHistory)
    res["userExperienceAnalyticsDevicePerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsDevicePerformance)
    res["userExperienceAnalyticsDeviceScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceScope)
    res["userExperienceAnalyticsDeviceScopes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceScopes)
    res["userExperienceAnalyticsDeviceScores"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceScores)
    res["userExperienceAnalyticsDeviceStartupHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceStartupHistory)
    res["userExperienceAnalyticsDeviceStartupProcesses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceStartupProcesses)
    res["userExperienceAnalyticsDeviceStartupProcessPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsDeviceStartupProcessPerformance)
    res["userExperienceAnalyticsDevicesWithoutCloudIdentity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue , m.SetUserExperienceAnalyticsDevicesWithoutCloudIdentity)
    res["userExperienceAnalyticsImpactingProcess"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue , m.SetUserExperienceAnalyticsImpactingProcess)
    res["userExperienceAnalyticsMetricHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsMetricHistory)
    res["userExperienceAnalyticsModelScores"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue , m.SetUserExperienceAnalyticsModelScores)
    res["userExperienceAnalyticsNotAutopilotReadyDevice"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue , m.SetUserExperienceAnalyticsNotAutopilotReadyDevice)
    res["userExperienceAnalyticsOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue , m.SetUserExperienceAnalyticsOverview)
    res["userExperienceAnalyticsRegressionSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsRegressionSummaryFromDiscriminatorValue , m.SetUserExperienceAnalyticsRegressionSummary)
    res["userExperienceAnalyticsRemoteConnection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue , m.SetUserExperienceAnalyticsRemoteConnection)
    res["userExperienceAnalyticsResourcePerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsResourcePerformance)
    res["userExperienceAnalyticsScoreHistory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue , m.SetUserExperienceAnalyticsScoreHistory)
    res["userExperienceAnalyticsSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue , m.SetUserExperienceAnalyticsSettings)
    res["userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue , m.SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric)
    res["userExperienceAnalyticsWorkFromAnywhereMetrics"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue , m.SetUserExperienceAnalyticsWorkFromAnywhereMetrics)
    res["userExperienceAnalyticsWorkFromAnywhereModelPerformance"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue , m.SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance)
    res["userPfxCertificates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserPFXCertificateFromDiscriminatorValue , m.SetUserPfxCertificates)
    res["virtualEndpoint"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVirtualEndpointFromDiscriminatorValue , m.SetVirtualEndpoint)
    res["windowsAutopilotDeploymentProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue , m.SetWindowsAutopilotDeploymentProfiles)
    res["windowsAutopilotDeviceIdentities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue , m.SetWindowsAutopilotDeviceIdentities)
    res["windowsAutopilotSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsAutopilotSettingsFromDiscriminatorValue , m.SetWindowsAutopilotSettings)
    res["windowsDriverUpdateProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsDriverUpdateProfileFromDiscriminatorValue , m.SetWindowsDriverUpdateProfiles)
    res["windowsFeatureUpdateProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsFeatureUpdateProfileFromDiscriminatorValue , m.SetWindowsFeatureUpdateProfiles)
    res["windowsInformationProtectionAppLearningSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue , m.SetWindowsInformationProtectionAppLearningSummaries)
    res["windowsInformationProtectionNetworkLearningSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue , m.SetWindowsInformationProtectionNetworkLearningSummaries)
    res["windowsMalwareInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsMalwareInformationFromDiscriminatorValue , m.SetWindowsMalwareInformation)
    res["windowsMalwareOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsMalwareOverviewFromDiscriminatorValue , m.SetWindowsMalwareOverview)
    res["windowsQualityUpdateProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileFromDiscriminatorValue , m.SetWindowsQualityUpdateProfiles)
    res["windowsUpdateCatalogItems"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWindowsUpdateCatalogItemFromDiscriminatorValue , m.SetWindowsUpdateCatalogItems)
    res["zebraFotaArtifacts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateZebraFotaArtifactFromDiscriminatorValue , m.SetZebraFotaArtifacts)
    res["zebraFotaConnector"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateZebraFotaConnectorFromDiscriminatorValue , m.SetZebraFotaConnector)
    res["zebraFotaDeployments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateZebraFotaDeploymentFromDiscriminatorValue , m.SetZebraFotaDeployments)
    return res
}
// GetGroupPolicyCategories gets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) GetGroupPolicyCategories()([]GroupPolicyCategoryable) {
    return m.groupPolicyCategories
}
// GetGroupPolicyConfigurations gets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) GetGroupPolicyConfigurations()([]GroupPolicyConfigurationable) {
    return m.groupPolicyConfigurations
}
// GetGroupPolicyDefinitionFiles gets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFileable) {
    return m.groupPolicyDefinitionFiles
}
// GetGroupPolicyDefinitions gets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitions()([]GroupPolicyDefinitionable) {
    return m.groupPolicyDefinitions
}
// GetGroupPolicyMigrationReports gets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReportable) {
    return m.groupPolicyMigrationReports
}
// GetGroupPolicyObjectFiles gets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) GetGroupPolicyObjectFiles()([]GroupPolicyObjectFileable) {
    return m.groupPolicyObjectFiles
}
// GetGroupPolicyUploadedDefinitionFiles gets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFileable) {
    return m.groupPolicyUploadedDefinitionFiles
}
// GetImportedDeviceIdentities gets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) GetImportedDeviceIdentities()([]ImportedDeviceIdentityable) {
    return m.importedDeviceIdentities
}
// GetImportedWindowsAutopilotDeviceIdentities gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable) {
    return m.importedWindowsAutopilotDeviceIdentities
}
// GetIntents gets the intents property value. The device management intents
func (m *DeviceManagement) GetIntents()([]DeviceManagementIntentable) {
    return m.intents
}
// GetIntuneAccountId gets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) GetIntuneAccountId()(*string) {
    return m.intuneAccountId
}
// GetIntuneBrand gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) GetIntuneBrand()(IntuneBrandable) {
    return m.intuneBrand
}
// GetIntuneBrandingProfiles gets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) GetIntuneBrandingProfiles()([]IntuneBrandingProfileable) {
    return m.intuneBrandingProfiles
}
// GetIosUpdateStatuses gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatusable) {
    return m.iosUpdateStatuses
}
// GetLastReportAggregationDateTime gets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportAggregationDateTime
}
// GetLegacyPcManangementEnabled gets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) GetLegacyPcManangementEnabled()(*bool) {
    return m.legacyPcManangementEnabled
}
// GetMacOSSoftwareUpdateAccountSummaries gets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummaryable) {
    return m.macOSSoftwareUpdateAccountSummaries
}
// GetManagedDeviceCleanupSettings gets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) GetManagedDeviceCleanupSettings()(ManagedDeviceCleanupSettingsable) {
    return m.managedDeviceCleanupSettings
}
// GetManagedDeviceEncryptionStates gets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionStateable) {
    return m.managedDeviceEncryptionStates
}
// GetManagedDeviceOverview gets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) GetManagedDeviceOverview()(ManagedDeviceOverviewable) {
    return m.managedDeviceOverview
}
// GetManagedDevices gets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetMaximumDepTokens gets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) GetMaximumDepTokens()(*int32) {
    return m.maximumDepTokens
}
// GetMicrosoftTunnelConfigurations gets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfigurationable) {
    return m.microsoftTunnelConfigurations
}
// GetMicrosoftTunnelHealthThresholds gets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThresholdable) {
    return m.microsoftTunnelHealthThresholds
}
// GetMicrosoftTunnelServerLogCollectionResponses gets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponseable) {
    return m.microsoftTunnelServerLogCollectionResponses
}
// GetMicrosoftTunnelSites gets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelSites()([]MicrosoftTunnelSiteable) {
    return m.microsoftTunnelSites
}
// GetMobileAppTroubleshootingEvents gets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable) {
    return m.mobileAppTroubleshootingEvents
}
// GetMobileThreatDefenseConnectors gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable) {
    return m.mobileThreatDefenseConnectors
}
// GetNdesConnectors gets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) GetNdesConnectors()([]NdesConnectorable) {
    return m.ndesConnectors
}
// GetNotificationMessageTemplates gets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplateable) {
    return m.notificationMessageTemplates
}
// GetOemWarrantyInformationOnboarding gets the oemWarrantyInformationOnboarding property value. List of OEM Warranty Statuses
func (m *DeviceManagement) GetOemWarrantyInformationOnboarding()([]OemWarrantyInformationOnboardingable) {
    return m.oemWarrantyInformationOnboarding
}
// GetOrganizationalMessageDetails gets the organizationalMessageDetails property value. A list of OrganizationalMessageDetails
func (m *DeviceManagement) GetOrganizationalMessageDetails()([]OrganizationalMessageDetailable) {
    return m.organizationalMessageDetails
}
// GetOrganizationalMessageGuidedContents gets the organizationalMessageGuidedContents property value. A list of OrganizationalMessageGuidedContents
func (m *DeviceManagement) GetOrganizationalMessageGuidedContents()([]OrganizationalMessageGuidedContentable) {
    return m.organizationalMessageGuidedContents
}
// GetRemoteActionAudits gets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) GetRemoteActionAudits()([]RemoteActionAuditable) {
    return m.remoteActionAudits
}
// GetRemoteAssistancePartners gets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartnerable) {
    return m.remoteAssistancePartners
}
// GetRemoteAssistanceSettings gets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) GetRemoteAssistanceSettings()(RemoteAssistanceSettingsable) {
    return m.remoteAssistanceSettings
}
// GetReports gets the reports property value. Reports singleton
func (m *DeviceManagement) GetReports()(DeviceManagementReportsable) {
    return m.reports
}
// GetResourceAccessProfiles gets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBaseable) {
    return m.resourceAccessProfiles
}
// GetResourceOperations gets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperationable) {
    return m.resourceOperations
}
// GetReusablePolicySettings gets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) GetReusablePolicySettings()([]DeviceManagementReusablePolicySettingable) {
    return m.reusablePolicySettings
}
// GetReusableSettings gets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) GetReusableSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    return m.reusableSettings
}
// GetRoleAssignments gets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable) {
    return m.roleAssignments
}
// GetRoleDefinitions gets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinitionable) {
    return m.roleDefinitions
}
// GetRoleScopeTags gets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) GetRoleScopeTags()([]RoleScopeTagable) {
    return m.roleScopeTags
}
// GetSettingDefinitions gets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) GetSettingDefinitions()([]DeviceManagementSettingDefinitionable) {
    return m.settingDefinitions
}
// GetSettings gets the settings property value. Account level settings.
func (m *DeviceManagement) GetSettings()(DeviceManagementSettingsable) {
    return m.settings
}
// GetSoftwareUpdateStatusSummary gets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable) {
    return m.softwareUpdateStatusSummary
}
// GetSubscriptions gets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) GetSubscriptions()(*DeviceManagementSubscriptions) {
    return m.subscriptions
}
// GetSubscriptionState gets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    return m.subscriptionState
}
// GetTelecomExpenseManagementPartners gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable) {
    return m.telecomExpenseManagementPartners
}
// GetTemplates gets the templates property value. The available templates
func (m *DeviceManagement) GetTemplates()([]DeviceManagementTemplateable) {
    return m.templates
}
// GetTemplateSettings gets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplateable) {
    return m.templateSettings
}
// GetTenantAttachRBAC gets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) GetTenantAttachRBAC()(TenantAttachRBACable) {
    return m.tenantAttachRBAC
}
// GetTermsAndConditions gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditionsable) {
    return m.termsAndConditions
}
// GetTroubleshootingEvents gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    return m.troubleshootingEvents
}
// GetUnlicensedAdminstratorsEnabled gets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) GetUnlicensedAdminstratorsEnabled()(*bool) {
    return m.unlicensedAdminstratorsEnabled
}
// GetUserExperienceAnalyticsAnomaly gets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomaly()([]UserExperienceAnalyticsAnomalyable) {
    return m.userExperienceAnalyticsAnomaly
}
// GetUserExperienceAnalyticsAnomalyDevice gets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalyDevice()([]UserExperienceAnalyticsAnomalyDeviceable) {
    return m.userExperienceAnalyticsAnomalyDevice
}
// GetUserExperienceAnalyticsAnomalySeverityOverview gets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalySeverityOverview()(UserExperienceAnalyticsAnomalySeverityOverviewable) {
    return m.userExperienceAnalyticsAnomalySeverityOverview
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformance gets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformance
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable) {
    return m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
}
// GetUserExperienceAnalyticsAppHealthDeviceModelPerformance gets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable) {
    return m.userExperienceAnalyticsAppHealthDeviceModelPerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformance gets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformance
}
// GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails gets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable) {
    return m.userExperienceAnalyticsAppHealthDevicePerformanceDetails
}
// GetUserExperienceAnalyticsAppHealthOSVersionPerformance gets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable) {
    return m.userExperienceAnalyticsAppHealthOSVersionPerformance
}
// GetUserExperienceAnalyticsAppHealthOverview gets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsAppHealthOverview
}
// GetUserExperienceAnalyticsBaselines gets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable) {
    return m.userExperienceAnalyticsBaselines
}
// GetUserExperienceAnalyticsBatteryHealthAppImpact gets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpactable) {
    return m.userExperienceAnalyticsBatteryHealthAppImpact
}
// GetUserExperienceAnalyticsBatteryHealthCapacityDetails gets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(UserExperienceAnalyticsBatteryHealthCapacityDetailsable) {
    return m.userExperienceAnalyticsBatteryHealthCapacityDetails
}
// GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact gets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable) {
    return m.userExperienceAnalyticsBatteryHealthDeviceAppImpact
}
// GetUserExperienceAnalyticsBatteryHealthDevicePerformance gets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthDevicePerformance
}
// GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory gets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable) {
    return m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
}
// GetUserExperienceAnalyticsBatteryHealthModelPerformance gets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthModelPerformance
}
// GetUserExperienceAnalyticsBatteryHealthOsPerformance gets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformanceable) {
    return m.userExperienceAnalyticsBatteryHealthOsPerformance
}
// GetUserExperienceAnalyticsBatteryHealthRuntimeDetails gets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable) {
    return m.userExperienceAnalyticsBatteryHealthRuntimeDetails
}
// GetUserExperienceAnalyticsCategories gets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable) {
    return m.userExperienceAnalyticsCategories
}
// GetUserExperienceAnalyticsDeviceMetricHistory gets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    return m.userExperienceAnalyticsDeviceMetricHistory
}
// GetUserExperienceAnalyticsDevicePerformance gets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable) {
    return m.userExperienceAnalyticsDevicePerformance
}
// GetUserExperienceAnalyticsDeviceScope gets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScope()(UserExperienceAnalyticsDeviceScopeable) {
    return m.userExperienceAnalyticsDeviceScope
}
// GetUserExperienceAnalyticsDeviceScopes gets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScopes()([]UserExperienceAnalyticsDeviceScopeable) {
    return m.userExperienceAnalyticsDeviceScopes
}
// GetUserExperienceAnalyticsDeviceScores gets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable) {
    return m.userExperienceAnalyticsDeviceScores
}
// GetUserExperienceAnalyticsDeviceStartupHistory gets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable) {
    return m.userExperienceAnalyticsDeviceStartupHistory
}
// GetUserExperienceAnalyticsDeviceStartupProcesses gets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable) {
    return m.userExperienceAnalyticsDeviceStartupProcesses
}
// GetUserExperienceAnalyticsDeviceStartupProcessPerformance gets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable) {
    return m.userExperienceAnalyticsDeviceStartupProcessPerformance
}
// GetUserExperienceAnalyticsDevicesWithoutCloudIdentity gets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable) {
    return m.userExperienceAnalyticsDevicesWithoutCloudIdentity
}
// GetUserExperienceAnalyticsImpactingProcess gets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcessable) {
    return m.userExperienceAnalyticsImpactingProcess
}
// GetUserExperienceAnalyticsMetricHistory gets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    return m.userExperienceAnalyticsMetricHistory
}
// GetUserExperienceAnalyticsModelScores gets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable) {
    return m.userExperienceAnalyticsModelScores
}
// GetUserExperienceAnalyticsNotAutopilotReadyDevice gets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDeviceable) {
    return m.userExperienceAnalyticsNotAutopilotReadyDevice
}
// GetUserExperienceAnalyticsOverview gets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable) {
    return m.userExperienceAnalyticsOverview
}
// GetUserExperienceAnalyticsRegressionSummary gets the userExperienceAnalyticsRegressionSummary property value. User experience analytics regression summary
func (m *DeviceManagement) GetUserExperienceAnalyticsRegressionSummary()(UserExperienceAnalyticsRegressionSummaryable) {
    return m.userExperienceAnalyticsRegressionSummary
}
// GetUserExperienceAnalyticsRemoteConnection gets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnectionable) {
    return m.userExperienceAnalyticsRemoteConnection
}
// GetUserExperienceAnalyticsResourcePerformance gets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformanceable) {
    return m.userExperienceAnalyticsResourcePerformance
}
// GetUserExperienceAnalyticsScoreHistory gets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable) {
    return m.userExperienceAnalyticsScoreHistory
}
// GetUserExperienceAnalyticsSettings gets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable) {
    return m.userExperienceAnalyticsSettings
}
// GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric gets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
}
// GetUserExperienceAnalyticsWorkFromAnywhereMetrics gets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable) {
    return m.userExperienceAnalyticsWorkFromAnywhereMetrics
}
// GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance gets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable) {
    return m.userExperienceAnalyticsWorkFromAnywhereModelPerformance
}
// GetUserPfxCertificates gets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) GetUserPfxCertificates()([]UserPFXCertificateable) {
    return m.userPfxCertificates
}
// GetVirtualEndpoint gets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) GetVirtualEndpoint()(VirtualEndpointable) {
    return m.virtualEndpoint
}
// GetWindowsAutopilotDeploymentProfiles gets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfileable) {
    return m.windowsAutopilotDeploymentProfiles
}
// GetWindowsAutopilotDeviceIdentities gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable) {
    return m.windowsAutopilotDeviceIdentities
}
// GetWindowsAutopilotSettings gets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) GetWindowsAutopilotSettings()(WindowsAutopilotSettingsable) {
    return m.windowsAutopilotSettings
}
// GetWindowsDriverUpdateProfiles gets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfileable) {
    return m.windowsDriverUpdateProfiles
}
// GetWindowsFeatureUpdateProfiles gets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfileable) {
    return m.windowsFeatureUpdateProfiles
}
// GetWindowsInformationProtectionAppLearningSummaries gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable) {
    return m.windowsInformationProtectionAppLearningSummaries
}
// GetWindowsInformationProtectionNetworkLearningSummaries gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable) {
    return m.windowsInformationProtectionNetworkLearningSummaries
}
// GetWindowsMalwareInformation gets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformationable) {
    return m.windowsMalwareInformation
}
// GetWindowsMalwareOverview gets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) GetWindowsMalwareOverview()(WindowsMalwareOverviewable) {
    return m.windowsMalwareOverview
}
// GetWindowsQualityUpdateProfiles gets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfileable) {
    return m.windowsQualityUpdateProfiles
}
// GetWindowsUpdateCatalogItems gets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItemable) {
    return m.windowsUpdateCatalogItems
}
// GetZebraFotaArtifacts gets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) GetZebraFotaArtifacts()([]ZebraFotaArtifactable) {
    return m.zebraFotaArtifacts
}
// GetZebraFotaConnector gets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) GetZebraFotaConnector()(ZebraFotaConnectorable) {
    return m.zebraFotaConnector
}
// GetZebraFotaDeployments gets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) GetZebraFotaDeployments()([]ZebraFotaDeploymentable) {
    return m.zebraFotaDeployments
}
// Serialize serializes information the current object
func (m *DeviceManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("accountMoveCompletionDateTime", m.GetAccountMoveCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("adminConsent", m.GetAdminConsent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("advancedThreatProtectionOnboardingStateSummary", m.GetAdvancedThreatProtectionOnboardingStateSummary())
        if err != nil {
            return err
        }
    }
    if m.GetAndroidDeviceOwnerEnrollmentProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAndroidDeviceOwnerEnrollmentProfiles())
        err = writer.WriteCollectionOfObjectValues("androidDeviceOwnerEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkAppConfigurationSchemas() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAndroidForWorkAppConfigurationSchemas())
        err = writer.WriteCollectionOfObjectValues("androidForWorkAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkEnrollmentProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAndroidForWorkEnrollmentProfiles())
        err = writer.WriteCollectionOfObjectValues("androidForWorkEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("androidForWorkSettings", m.GetAndroidForWorkSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("androidManagedStoreAccountEnterpriseSettings", m.GetAndroidManagedStoreAccountEnterpriseSettings())
        if err != nil {
            return err
        }
    }
    if m.GetAndroidManagedStoreAppConfigurationSchemas() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAndroidManagedStoreAppConfigurationSchemas())
        err = writer.WriteCollectionOfObjectValues("androidManagedStoreAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("applePushNotificationCertificate", m.GetApplePushNotificationCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetAppleUserInitiatedEnrollmentProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppleUserInitiatedEnrollmentProfiles())
        err = writer.WriteCollectionOfObjectValues("appleUserInitiatedEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignmentFilters())
        err = writer.WriteCollectionOfObjectValues("assignmentFilters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAuditEvents())
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAutopilotEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAutopilotEvents())
        err = writer.WriteCollectionOfObjectValues("autopilotEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCartToClassAssociations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCartToClassAssociations())
        err = writer.WriteCollectionOfObjectValues("cartToClassAssociations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCategories())
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateConnectorDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCertificateConnectorDetails())
        err = writer.WriteCollectionOfObjectValues("certificateConnectorDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChromeOSOnboardingSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChromeOSOnboardingSettings())
        err = writer.WriteCollectionOfObjectValues("chromeOSOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCConnectivityIssues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPCConnectivityIssues())
        err = writer.WriteCollectionOfObjectValues("cloudPCConnectivityIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagedDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComanagedDevices())
        err = writer.WriteCollectionOfObjectValues("comanagedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagementEligibleDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComanagementEligibleDevices())
        err = writer.WriteCollectionOfObjectValues("comanagementEligibleDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComplianceCategories())
        err = writer.WriteCollectionOfObjectValues("complianceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceManagementPartners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComplianceManagementPartners())
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliancePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCompliancePolicies())
        err = writer.WriteCollectionOfObjectValues("compliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetComplianceSettings())
        err = writer.WriteCollectionOfObjectValues("complianceSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    if m.GetConfigManagerCollections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigManagerCollections())
        err = writer.WriteCollectionOfObjectValues("configManagerCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurationCategories())
        err = writer.WriteCollectionOfObjectValues("configurationCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurationPolicies())
        err = writer.WriteCollectionOfObjectValues("configurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicyTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurationPolicyTemplates())
        err = writer.WriteCollectionOfObjectValues("configurationPolicyTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConfigurationSettings())
        err = writer.WriteCollectionOfObjectValues("configurationSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataProcessorServiceForWindowsFeaturesOnboarding", m.GetDataProcessorServiceForWindowsFeaturesOnboarding())
        if err != nil {
            return err
        }
    }
    if m.GetDataSharingConsents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDataSharingConsents())
        err = writer.WriteCollectionOfObjectValues("dataSharingConsents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDepOnboardingSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDepOnboardingSettings())
        err = writer.WriteCollectionOfObjectValues("depOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentials() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDerivedCredentials())
        err = writer.WriteCollectionOfObjectValues("derivedCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDetectedApps())
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCategories())
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCompliancePolicies())
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicyDeviceStateSummary", m.GetDeviceCompliancePolicyDeviceStateSummary())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCompliancePolicySettingStateSummaries())
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceComplianceScripts())
        err = writer.WriteCollectionOfObjectValues("deviceComplianceScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationConflictSummary() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurationConflictSummary())
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationConflictSummary", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationDeviceStateSummaries", m.GetDeviceConfigurationDeviceStateSummaries())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationRestrictedAppsViolations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurationRestrictedAppsViolations())
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationRestrictedAppsViolations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurations())
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurationsAllManagedDeviceCertificateStates())
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationsAllManagedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationUserStateSummaries", m.GetDeviceConfigurationUserStateSummaries())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCustomAttributeShellScripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCustomAttributeShellScripts())
        err = writer.WriteCollectionOfObjectValues("deviceCustomAttributeShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceEnrollmentConfigurations())
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceHealthScripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceHealthScripts())
        err = writer.WriteCollectionOfObjectValues("deviceHealthScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementPartners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceManagementPartners())
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceManagementScripts())
        err = writer.WriteCollectionOfObjectValues("deviceManagementScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceProtectionOverview", m.GetDeviceProtectionOverview())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceShellScripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceShellScripts())
        err = writer.WriteCollectionOfObjectValues("deviceShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainJoinConnectors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDomainJoinConnectors())
        err = writer.WriteCollectionOfObjectValues("domainJoinConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmbeddedSIMActivationCodePools() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmbeddedSIMActivationCodePools())
        err = writer.WriteCollectionOfObjectValues("embeddedSIMActivationCodePools", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExchangeConnectors())
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeOnPremisesPolicies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExchangeOnPremisesPolicies())
        err = writer.WriteCollectionOfObjectValues("exchangeOnPremisesPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("exchangeOnPremisesPolicy", m.GetExchangeOnPremisesPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyCategories())
        err = writer.WriteCollectionOfObjectValues("groupPolicyCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyConfigurations())
        err = writer.WriteCollectionOfObjectValues("groupPolicyConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitionFiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyDefinitionFiles())
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyDefinitions())
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyMigrationReports() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyMigrationReports())
        err = writer.WriteCollectionOfObjectValues("groupPolicyMigrationReports", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyObjectFiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyObjectFiles())
        err = writer.WriteCollectionOfObjectValues("groupPolicyObjectFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyUploadedDefinitionFiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetGroupPolicyUploadedDefinitionFiles())
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetImportedDeviceIdentities())
        err = writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedWindowsAutopilotDeviceIdentities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetImportedWindowsAutopilotDeviceIdentities())
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIntents())
        err = writer.WriteCollectionOfObjectValues("intents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneAccountId", m.GetIntuneAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intuneBrand", m.GetIntuneBrand())
        if err != nil {
            return err
        }
    }
    if m.GetIntuneBrandingProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIntuneBrandingProfiles())
        err = writer.WriteCollectionOfObjectValues("intuneBrandingProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosUpdateStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetIosUpdateStatuses())
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMacOSSoftwareUpdateAccountSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMacOSSoftwareUpdateAccountSummaries())
        err = writer.WriteCollectionOfObjectValues("macOSSoftwareUpdateAccountSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceCleanupSettings", m.GetManagedDeviceCleanupSettings())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceEncryptionStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceEncryptionStates())
        err = writer.WriteCollectionOfObjectValues("managedDeviceEncryptionStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceOverview", m.GetManagedDeviceOverview())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDevices() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDevices())
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumDepTokens", m.GetMaximumDepTokens())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelConfigurations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMicrosoftTunnelConfigurations())
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelHealthThresholds() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMicrosoftTunnelHealthThresholds())
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelHealthThresholds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelServerLogCollectionResponses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMicrosoftTunnelServerLogCollectionResponses())
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServerLogCollectionResponses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelSites() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMicrosoftTunnelSites())
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelSites", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppTroubleshootingEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileAppTroubleshootingEvents())
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseConnectors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMobileThreatDefenseConnectors())
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNdesConnectors() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNdesConnectors())
        err = writer.WriteCollectionOfObjectValues("ndesConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNotificationMessageTemplates())
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOemWarrantyInformationOnboarding() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOemWarrantyInformationOnboarding())
        err = writer.WriteCollectionOfObjectValues("oemWarrantyInformationOnboarding", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationalMessageDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOrganizationalMessageDetails())
        err = writer.WriteCollectionOfObjectValues("organizationalMessageDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOrganizationalMessageGuidedContents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetOrganizationalMessageGuidedContents())
        err = writer.WriteCollectionOfObjectValues("organizationalMessageGuidedContents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteActionAudits() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRemoteActionAudits())
        err = writer.WriteCollectionOfObjectValues("remoteActionAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistancePartners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRemoteAssistancePartners())
        err = writer.WriteCollectionOfObjectValues("remoteAssistancePartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("remoteAssistanceSettings", m.GetRemoteAssistanceSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    if m.GetResourceAccessProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResourceAccessProfiles())
        err = writer.WriteCollectionOfObjectValues("resourceAccessProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceOperations() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResourceOperations())
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusablePolicySettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReusablePolicySettings())
        err = writer.WriteCollectionOfObjectValues("reusablePolicySettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusableSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReusableSettings())
        err = writer.WriteCollectionOfObjectValues("reusableSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleAssignments())
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleDefinitions())
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTags() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRoleScopeTags())
        err = writer.WriteCollectionOfObjectValues("roleScopeTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingDefinitions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSettingDefinitions())
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateStatusSummary", m.GetSoftwareUpdateStatusSummary())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := (*m.GetSubscriptions()).String()
        err = writer.WriteStringValue("subscriptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionState() != nil {
        cast := (*m.GetSubscriptionState()).String()
        err = writer.WriteStringValue("subscriptionState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTelecomExpenseManagementPartners() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTelecomExpenseManagementPartners())
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTemplates())
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTemplateSettings())
        err = writer.WriteCollectionOfObjectValues("templateSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantAttachRBAC", m.GetTenantAttachRBAC())
        if err != nil {
            return err
        }
    }
    if m.GetTermsAndConditions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTermsAndConditions())
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTroubleshootingEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTroubleshootingEvents())
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomaly() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAnomaly())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomaly", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomalyDevice() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAnomalyDevice())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomalyDevice", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsAnomalySeverityOverview", m.GetUserExperienceAnalyticsAnomalySeverityOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthApplicationPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDeviceModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthDevicePerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformanceDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthOSVersionPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsAppHealthOverview", m.GetUserExperienceAnalyticsAppHealthOverview())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBaselines() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBaselines())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBaselines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthAppImpact() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthAppImpact())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthAppImpact", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsBatteryHealthCapacityDetails", m.GetUserExperienceAnalyticsBatteryHealthCapacityDetails())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceAppImpact", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthModelPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsBatteryHealthOsPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthOsPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsBatteryHealthRuntimeDetails", m.GetUserExperienceAnalyticsBatteryHealthRuntimeDetails())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsCategories())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceMetricHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceMetricHistory())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicePerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDevicePerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsDeviceScope", m.GetUserExperienceAnalyticsDeviceScope())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScopes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceScopes())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScores() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceScores())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceStartupHistory())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcesses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceStartupProcesses())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcessPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicesWithoutCloudIdentity", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsImpactingProcess() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsImpactingProcess())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsImpactingProcess", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsMetricHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsMetricHistory())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsModelScores() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsModelScores())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsModelScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsNotAutopilotReadyDevice())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsNotAutopilotReadyDevice", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsOverview", m.GetUserExperienceAnalyticsOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsRegressionSummary", m.GetUserExperienceAnalyticsRegressionSummary())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsRemoteConnection() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsRemoteConnection())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsRemoteConnection", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsResourcePerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsResourcePerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsResourcePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsScoreHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsScoreHistory())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsScoreHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsSettings", m.GetUserExperienceAnalyticsSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric", m.GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric())
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereMetrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance())
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserPfxCertificates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserPfxCertificates())
        err = writer.WriteCollectionOfObjectValues("userPfxCertificates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("virtualEndpoint", m.GetVirtualEndpoint())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeploymentProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsAutopilotDeploymentProfiles())
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeploymentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeviceIdentities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsAutopilotDeviceIdentities())
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsAutopilotSettings", m.GetWindowsAutopilotSettings())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsDriverUpdateProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsDriverUpdateProfiles())
        err = writer.WriteCollectionOfObjectValues("windowsDriverUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsFeatureUpdateProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsFeatureUpdateProfiles())
        err = writer.WriteCollectionOfObjectValues("windowsFeatureUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionAppLearningSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionAppLearningSummaries())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionNetworkLearningSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsInformationProtectionNetworkLearningSummaries())
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsMalwareInformation() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsMalwareInformation())
        err = writer.WriteCollectionOfObjectValues("windowsMalwareInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsMalwareOverview", m.GetWindowsMalwareOverview())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsQualityUpdateProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsQualityUpdateProfiles())
        err = writer.WriteCollectionOfObjectValues("windowsQualityUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsUpdateCatalogItems() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWindowsUpdateCatalogItems())
        err = writer.WriteCollectionOfObjectValues("windowsUpdateCatalogItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetZebraFotaArtifacts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetZebraFotaArtifacts())
        err = writer.WriteCollectionOfObjectValues("zebraFotaArtifacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("zebraFotaConnector", m.GetZebraFotaConnector())
        if err != nil {
            return err
        }
    }
    if m.GetZebraFotaDeployments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetZebraFotaDeployments())
        err = writer.WriteCollectionOfObjectValues("zebraFotaDeployments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountMoveCompletionDateTime sets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.accountMoveCompletionDateTime = value
}
// SetAdminConsent sets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) SetAdminConsent(value AdminConsentable)() {
    m.adminConsent = value
}
// SetAdvancedThreatProtectionOnboardingStateSummary sets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) SetAdvancedThreatProtectionOnboardingStateSummary(value AdvancedThreatProtectionOnboardingStateSummaryable)() {
    m.advancedThreatProtectionOnboardingStateSummary = value
}
// SetAndroidDeviceOwnerEnrollmentProfiles sets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfileable)() {
    m.androidDeviceOwnerEnrollmentProfiles = value
}
// SetAndroidForWorkAppConfigurationSchemas sets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchemaable)() {
    m.androidForWorkAppConfigurationSchemas = value
}
// SetAndroidForWorkEnrollmentProfiles sets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfileable)() {
    m.androidForWorkEnrollmentProfiles = value
}
// SetAndroidForWorkSettings sets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) SetAndroidForWorkSettings(value AndroidForWorkSettingsable)() {
    m.androidForWorkSettings = value
}
// SetAndroidManagedStoreAccountEnterpriseSettings sets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) SetAndroidManagedStoreAccountEnterpriseSettings(value AndroidManagedStoreAccountEnterpriseSettingsable)() {
    m.androidManagedStoreAccountEnterpriseSettings = value
}
// SetAndroidManagedStoreAppConfigurationSchemas sets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchemaable)() {
    m.androidManagedStoreAppConfigurationSchemas = value
}
// SetApplePushNotificationCertificate sets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)() {
    m.applePushNotificationCertificate = value
}
// SetAppleUserInitiatedEnrollmentProfiles sets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfileable)() {
    m.appleUserInitiatedEnrollmentProfiles = value
}
// SetAssignmentFilters sets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilterable)() {
    m.assignmentFilters = value
}
// SetAuditEvents sets the auditEvents property value. The Audit Events
func (m *DeviceManagement) SetAuditEvents(value []AuditEventable)() {
    m.auditEvents = value
}
// SetAutopilotEvents sets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) SetAutopilotEvents(value []DeviceManagementAutopilotEventable)() {
    m.autopilotEvents = value
}
// SetCartToClassAssociations sets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) SetCartToClassAssociations(value []CartToClassAssociationable)() {
    m.cartToClassAssociations = value
}
// SetCategories sets the categories property value. The available categories
func (m *DeviceManagement) SetCategories(value []DeviceManagementSettingCategoryable)() {
    m.categories = value
}
// SetCertificateConnectorDetails sets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) SetCertificateConnectorDetails(value []CertificateConnectorDetailsable)() {
    m.certificateConnectorDetails = value
}
// SetChromeOSOnboardingSettings sets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettingsable)() {
    m.chromeOSOnboardingSettings = value
}
// SetCloudPCConnectivityIssues sets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssueable)() {
    m.cloudPCConnectivityIssues = value
}
// SetComanagedDevices sets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) SetComanagedDevices(value []ManagedDeviceable)() {
    m.comanagedDevices = value
}
// SetComanagementEligibleDevices sets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) SetComanagementEligibleDevices(value []ComanagementEligibleDeviceable)() {
    m.comanagementEligibleDevices = value
}
// SetComplianceCategories sets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) SetComplianceCategories(value []DeviceManagementConfigurationCategoryable)() {
    m.complianceCategories = value
}
// SetComplianceManagementPartners sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartnerable)() {
    m.complianceManagementPartners = value
}
// SetCompliancePolicies sets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) SetCompliancePolicies(value []DeviceManagementCompliancePolicyable)() {
    m.compliancePolicies = value
}
// SetComplianceSettings sets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.complianceSettings = value
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)() {
    m.conditionalAccessSettings = value
}
// SetConfigManagerCollections sets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) SetConfigManagerCollections(value []ConfigManagerCollectionable)() {
    m.configManagerCollections = value
}
// SetConfigurationCategories sets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) SetConfigurationCategories(value []DeviceManagementConfigurationCategoryable)() {
    m.configurationCategories = value
}
// SetConfigurationPolicies sets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) SetConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)() {
    m.configurationPolicies = value
}
// SetConfigurationPolicyTemplates sets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplateable)() {
    m.configurationPolicyTemplates = value
}
// SetConfigurationSettings sets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.configurationSettings = value
}
// SetDataProcessorServiceForWindowsFeaturesOnboarding sets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) SetDataProcessorServiceForWindowsFeaturesOnboarding(value DataProcessorServiceForWindowsFeaturesOnboardingable)() {
    m.dataProcessorServiceForWindowsFeaturesOnboarding = value
}
// SetDataSharingConsents sets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) SetDataSharingConsents(value []DataSharingConsentable)() {
    m.dataSharingConsents = value
}
// SetDepOnboardingSettings sets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) SetDepOnboardingSettings(value []DepOnboardingSettingable)() {
    m.depOnboardingSettings = value
}
// SetDerivedCredentials sets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettingsable)() {
    m.derivedCredentials = value
}
// SetDetectedApps sets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) SetDetectedApps(value []DetectedAppable)() {
    m.detectedApps = value
}
// SetDeviceCategories sets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategoryable)() {
    m.deviceCategories = value
}
// SetDeviceCompliancePolicies sets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)() {
    m.deviceCompliancePolicies = value
}
// SetDeviceCompliancePolicyDeviceStateSummary sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// SetDeviceComplianceReportSummarizationDateTime sets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceComplianceReportSummarizationDateTime = value
}
// SetDeviceComplianceScripts sets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceComplianceScripts(value []DeviceComplianceScriptable)() {
    m.deviceComplianceScripts = value
}
// SetDeviceConfigurationConflictSummary sets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummaryable)() {
    m.deviceConfigurationConflictSummary = value
}
// SetDeviceConfigurationDeviceStateSummaries sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
// SetDeviceConfigurationRestrictedAppsViolations sets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolationable)() {
    m.deviceConfigurationRestrictedAppsViolations = value
}
// SetDeviceConfigurations sets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfigurationable)() {
    m.deviceConfigurations = value
}
// SetDeviceConfigurationsAllManagedDeviceCertificateStates sets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateStateable)() {
    m.deviceConfigurationsAllManagedDeviceCertificateStates = value
}
// SetDeviceConfigurationUserStateSummaries sets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationUserStateSummaries(value DeviceConfigurationUserStateSummaryable)() {
    m.deviceConfigurationUserStateSummaries = value
}
// SetDeviceCustomAttributeShellScripts sets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScriptable)() {
    m.deviceCustomAttributeShellScripts = value
}
// SetDeviceEnrollmentConfigurations sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)() {
    m.deviceEnrollmentConfigurations = value
}
// SetDeviceHealthScripts sets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceHealthScripts(value []DeviceHealthScriptable)() {
    m.deviceHealthScripts = value
}
// SetDeviceManagementPartners sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartnerable)() {
    m.deviceManagementPartners = value
}
// SetDeviceManagementScripts sets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceManagementScripts(value []DeviceManagementScriptable)() {
    m.deviceManagementScripts = value
}
// SetDeviceProtectionOverview sets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) SetDeviceProtectionOverview(value DeviceProtectionOverviewable)() {
    m.deviceProtectionOverview = value
}
// SetDeviceShellScripts sets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceShellScripts(value []DeviceShellScriptable)() {
    m.deviceShellScripts = value
}
// SetDomainJoinConnectors sets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnectorable)() {
    m.domainJoinConnectors = value
}
// SetEmbeddedSIMActivationCodePools sets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePoolable)() {
    m.embeddedSIMActivationCodePools = value
}
// SetExchangeConnectors sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)() {
    m.exchangeConnectors = value
}
// SetExchangeOnPremisesPolicies sets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicyable)() {
    m.exchangeOnPremisesPolicies = value
}
// SetExchangeOnPremisesPolicy sets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) SetExchangeOnPremisesPolicy(value DeviceManagementExchangeOnPremisesPolicyable)() {
    m.exchangeOnPremisesPolicy = value
}
// SetGroupPolicyCategories sets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) SetGroupPolicyCategories(value []GroupPolicyCategoryable)() {
    m.groupPolicyCategories = value
}
// SetGroupPolicyConfigurations sets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) SetGroupPolicyConfigurations(value []GroupPolicyConfigurationable)() {
    m.groupPolicyConfigurations = value
}
// SetGroupPolicyDefinitionFiles sets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFileable)() {
    m.groupPolicyDefinitionFiles = value
}
// SetGroupPolicyDefinitions sets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitions(value []GroupPolicyDefinitionable)() {
    m.groupPolicyDefinitions = value
}
// SetGroupPolicyMigrationReports sets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReportable)() {
    m.groupPolicyMigrationReports = value
}
// SetGroupPolicyObjectFiles sets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) SetGroupPolicyObjectFiles(value []GroupPolicyObjectFileable)() {
    m.groupPolicyObjectFiles = value
}
// SetGroupPolicyUploadedDefinitionFiles sets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFileable)() {
    m.groupPolicyUploadedDefinitionFiles = value
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)() {
    m.importedDeviceIdentities = value
}
// SetImportedWindowsAutopilotDeviceIdentities sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
// SetIntents sets the intents property value. The device management intents
func (m *DeviceManagement) SetIntents(value []DeviceManagementIntentable)() {
    m.intents = value
}
// SetIntuneAccountId sets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) SetIntuneAccountId(value *string)() {
    m.intuneAccountId = value
}
// SetIntuneBrand sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) SetIntuneBrand(value IntuneBrandable)() {
    m.intuneBrand = value
}
// SetIntuneBrandingProfiles sets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) SetIntuneBrandingProfiles(value []IntuneBrandingProfileable)() {
    m.intuneBrandingProfiles = value
}
// SetIosUpdateStatuses sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)() {
    m.iosUpdateStatuses = value
}
// SetLastReportAggregationDateTime sets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportAggregationDateTime = value
}
// SetLegacyPcManangementEnabled sets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) SetLegacyPcManangementEnabled(value *bool)() {
    m.legacyPcManangementEnabled = value
}
// SetMacOSSoftwareUpdateAccountSummaries sets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummaryable)() {
    m.macOSSoftwareUpdateAccountSummaries = value
}
// SetManagedDeviceCleanupSettings sets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) SetManagedDeviceCleanupSettings(value ManagedDeviceCleanupSettingsable)() {
    m.managedDeviceCleanupSettings = value
}
// SetManagedDeviceEncryptionStates sets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionStateable)() {
    m.managedDeviceEncryptionStates = value
}
// SetManagedDeviceOverview sets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) SetManagedDeviceOverview(value ManagedDeviceOverviewable)() {
    m.managedDeviceOverview = value
}
// SetManagedDevices sets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetMaximumDepTokens sets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) SetMaximumDepTokens(value *int32)() {
    m.maximumDepTokens = value
}
// SetMicrosoftTunnelConfigurations sets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfigurationable)() {
    m.microsoftTunnelConfigurations = value
}
// SetMicrosoftTunnelHealthThresholds sets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThresholdable)() {
    m.microsoftTunnelHealthThresholds = value
}
// SetMicrosoftTunnelServerLogCollectionResponses sets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponseable)() {
    m.microsoftTunnelServerLogCollectionResponses = value
}
// SetMicrosoftTunnelSites sets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelSites(value []MicrosoftTunnelSiteable)() {
    m.microsoftTunnelSites = value
}
// SetMobileAppTroubleshootingEvents sets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)() {
    m.mobileAppTroubleshootingEvents = value
}
// SetMobileThreatDefenseConnectors sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)() {
    m.mobileThreatDefenseConnectors = value
}
// SetNdesConnectors sets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) SetNdesConnectors(value []NdesConnectorable)() {
    m.ndesConnectors = value
}
// SetNotificationMessageTemplates sets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplateable)() {
    m.notificationMessageTemplates = value
}
// SetOemWarrantyInformationOnboarding sets the oemWarrantyInformationOnboarding property value. List of OEM Warranty Statuses
func (m *DeviceManagement) SetOemWarrantyInformationOnboarding(value []OemWarrantyInformationOnboardingable)() {
    m.oemWarrantyInformationOnboarding = value
}
// SetOrganizationalMessageDetails sets the organizationalMessageDetails property value. A list of OrganizationalMessageDetails
func (m *DeviceManagement) SetOrganizationalMessageDetails(value []OrganizationalMessageDetailable)() {
    m.organizationalMessageDetails = value
}
// SetOrganizationalMessageGuidedContents sets the organizationalMessageGuidedContents property value. A list of OrganizationalMessageGuidedContents
func (m *DeviceManagement) SetOrganizationalMessageGuidedContents(value []OrganizationalMessageGuidedContentable)() {
    m.organizationalMessageGuidedContents = value
}
// SetRemoteActionAudits sets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) SetRemoteActionAudits(value []RemoteActionAuditable)() {
    m.remoteActionAudits = value
}
// SetRemoteAssistancePartners sets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)() {
    m.remoteAssistancePartners = value
}
// SetRemoteAssistanceSettings sets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) SetRemoteAssistanceSettings(value RemoteAssistanceSettingsable)() {
    m.remoteAssistanceSettings = value
}
// SetReports sets the reports property value. Reports singleton
func (m *DeviceManagement) SetReports(value DeviceManagementReportsable)() {
    m.reports = value
}
// SetResourceAccessProfiles sets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBaseable)() {
    m.resourceAccessProfiles = value
}
// SetResourceOperations sets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperationable)() {
    m.resourceOperations = value
}
// SetReusablePolicySettings sets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) SetReusablePolicySettings(value []DeviceManagementReusablePolicySettingable)() {
    m.reusablePolicySettings = value
}
// SetReusableSettings sets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) SetReusableSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    m.reusableSettings = value
}
// SetRoleAssignments sets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)() {
    m.roleAssignments = value
}
// SetRoleDefinitions sets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinitionable)() {
    m.roleDefinitions = value
}
// SetRoleScopeTags sets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) SetRoleScopeTags(value []RoleScopeTagable)() {
    m.roleScopeTags = value
}
// SetSettingDefinitions sets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)() {
    m.settingDefinitions = value
}
// SetSettings sets the settings property value. Account level settings.
func (m *DeviceManagement) SetSettings(value DeviceManagementSettingsable)() {
    m.settings = value
}
// SetSoftwareUpdateStatusSummary sets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)() {
    m.softwareUpdateStatusSummary = value
}
// SetSubscriptions sets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) SetSubscriptions(value *DeviceManagementSubscriptions)() {
    m.subscriptions = value
}
// SetSubscriptionState sets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
// SetTelecomExpenseManagementPartners sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)() {
    m.telecomExpenseManagementPartners = value
}
// SetTemplates sets the templates property value. The available templates
func (m *DeviceManagement) SetTemplates(value []DeviceManagementTemplateable)() {
    m.templates = value
}
// SetTemplateSettings sets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplateable)() {
    m.templateSettings = value
}
// SetTenantAttachRBAC sets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) SetTenantAttachRBAC(value TenantAttachRBACable)() {
    m.tenantAttachRBAC = value
}
// SetTermsAndConditions sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditionsable)() {
    m.termsAndConditions = value
}
// SetTroubleshootingEvents sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    m.troubleshootingEvents = value
}
// SetUnlicensedAdminstratorsEnabled sets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) SetUnlicensedAdminstratorsEnabled(value *bool)() {
    m.unlicensedAdminstratorsEnabled = value
}
// SetUserExperienceAnalyticsAnomaly sets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomaly(value []UserExperienceAnalyticsAnomalyable)() {
    m.userExperienceAnalyticsAnomaly = value
}
// SetUserExperienceAnalyticsAnomalyDevice sets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalyDevice(value []UserExperienceAnalyticsAnomalyDeviceable)() {
    m.userExperienceAnalyticsAnomalyDevice = value
}
// SetUserExperienceAnalyticsAnomalySeverityOverview sets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalySeverityOverview(value UserExperienceAnalyticsAnomalySeverityOverviewable)() {
    m.userExperienceAnalyticsAnomalySeverityOverview = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformance sets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformance = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = value
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = value
}
// SetUserExperienceAnalyticsAppHealthDeviceModelPerformance sets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)() {
    m.userExperienceAnalyticsAppHealthDeviceModelPerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformance sets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformance = value
}
// SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails sets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)() {
    m.userExperienceAnalyticsAppHealthDevicePerformanceDetails = value
}
// SetUserExperienceAnalyticsAppHealthOSVersionPerformance sets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)() {
    m.userExperienceAnalyticsAppHealthOSVersionPerformance = value
}
// SetUserExperienceAnalyticsAppHealthOverview sets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsAppHealthOverview = value
}
// SetUserExperienceAnalyticsBaselines sets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)() {
    m.userExperienceAnalyticsBaselines = value
}
// SetUserExperienceAnalyticsBatteryHealthAppImpact sets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpactable)() {
    m.userExperienceAnalyticsBatteryHealthAppImpact = value
}
// SetUserExperienceAnalyticsBatteryHealthCapacityDetails sets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value UserExperienceAnalyticsBatteryHealthCapacityDetailsable)() {
    m.userExperienceAnalyticsBatteryHealthCapacityDetails = value
}
// SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact sets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)() {
    m.userExperienceAnalyticsBatteryHealthDeviceAppImpact = value
}
// SetUserExperienceAnalyticsBatteryHealthDevicePerformance sets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthDevicePerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory sets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)() {
    m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory = value
}
// SetUserExperienceAnalyticsBatteryHealthModelPerformance sets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthModelPerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthOsPerformance sets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformanceable)() {
    m.userExperienceAnalyticsBatteryHealthOsPerformance = value
}
// SetUserExperienceAnalyticsBatteryHealthRuntimeDetails sets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)() {
    m.userExperienceAnalyticsBatteryHealthRuntimeDetails = value
}
// SetUserExperienceAnalyticsCategories sets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)() {
    m.userExperienceAnalyticsCategories = value
}
// SetUserExperienceAnalyticsDeviceMetricHistory sets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    m.userExperienceAnalyticsDeviceMetricHistory = value
}
// SetUserExperienceAnalyticsDevicePerformance sets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)() {
    m.userExperienceAnalyticsDevicePerformance = value
}
// SetUserExperienceAnalyticsDeviceScope sets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScope(value UserExperienceAnalyticsDeviceScopeable)() {
    m.userExperienceAnalyticsDeviceScope = value
}
// SetUserExperienceAnalyticsDeviceScopes sets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScopes(value []UserExperienceAnalyticsDeviceScopeable)() {
    m.userExperienceAnalyticsDeviceScopes = value
}
// SetUserExperienceAnalyticsDeviceScores sets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)() {
    m.userExperienceAnalyticsDeviceScores = value
}
// SetUserExperienceAnalyticsDeviceStartupHistory sets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)() {
    m.userExperienceAnalyticsDeviceStartupHistory = value
}
// SetUserExperienceAnalyticsDeviceStartupProcesses sets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)() {
    m.userExperienceAnalyticsDeviceStartupProcesses = value
}
// SetUserExperienceAnalyticsDeviceStartupProcessPerformance sets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)() {
    m.userExperienceAnalyticsDeviceStartupProcessPerformance = value
}
// SetUserExperienceAnalyticsDevicesWithoutCloudIdentity sets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentityable)() {
    m.userExperienceAnalyticsDevicesWithoutCloudIdentity = value
}
// SetUserExperienceAnalyticsImpactingProcess sets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcessable)() {
    m.userExperienceAnalyticsImpactingProcess = value
}
// SetUserExperienceAnalyticsMetricHistory sets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    m.userExperienceAnalyticsMetricHistory = value
}
// SetUserExperienceAnalyticsModelScores sets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)() {
    m.userExperienceAnalyticsModelScores = value
}
// SetUserExperienceAnalyticsNotAutopilotReadyDevice sets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDeviceable)() {
    m.userExperienceAnalyticsNotAutopilotReadyDevice = value
}
// SetUserExperienceAnalyticsOverview sets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)() {
    m.userExperienceAnalyticsOverview = value
}
// SetUserExperienceAnalyticsRegressionSummary sets the userExperienceAnalyticsRegressionSummary property value. User experience analytics regression summary
func (m *DeviceManagement) SetUserExperienceAnalyticsRegressionSummary(value UserExperienceAnalyticsRegressionSummaryable)() {
    m.userExperienceAnalyticsRegressionSummary = value
}
// SetUserExperienceAnalyticsRemoteConnection sets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnectionable)() {
    m.userExperienceAnalyticsRemoteConnection = value
}
// SetUserExperienceAnalyticsResourcePerformance sets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformanceable)() {
    m.userExperienceAnalyticsResourcePerformance = value
}
// SetUserExperienceAnalyticsScoreHistory sets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)() {
    m.userExperienceAnalyticsScoreHistory = value
}
// SetUserExperienceAnalyticsSettings sets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)() {
    m.userExperienceAnalyticsSettings = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric sets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereMetrics sets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)() {
    m.userExperienceAnalyticsWorkFromAnywhereMetrics = value
}
// SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance sets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)() {
    m.userExperienceAnalyticsWorkFromAnywhereModelPerformance = value
}
// SetUserPfxCertificates sets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) SetUserPfxCertificates(value []UserPFXCertificateable)() {
    m.userPfxCertificates = value
}
// SetVirtualEndpoint sets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) SetVirtualEndpoint(value VirtualEndpointable)() {
    m.virtualEndpoint = value
}
// SetWindowsAutopilotDeploymentProfiles sets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfileable)() {
    m.windowsAutopilotDeploymentProfiles = value
}
// SetWindowsAutopilotDeviceIdentities sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)() {
    m.windowsAutopilotDeviceIdentities = value
}
// SetWindowsAutopilotSettings sets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) SetWindowsAutopilotSettings(value WindowsAutopilotSettingsable)() {
    m.windowsAutopilotSettings = value
}
// SetWindowsDriverUpdateProfiles sets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfileable)() {
    m.windowsDriverUpdateProfiles = value
}
// SetWindowsFeatureUpdateProfiles sets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfileable)() {
    m.windowsFeatureUpdateProfiles = value
}
// SetWindowsInformationProtectionAppLearningSummaries sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
// SetWindowsInformationProtectionNetworkLearningSummaries sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
// SetWindowsMalwareInformation sets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformationable)() {
    m.windowsMalwareInformation = value
}
// SetWindowsMalwareOverview sets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) SetWindowsMalwareOverview(value WindowsMalwareOverviewable)() {
    m.windowsMalwareOverview = value
}
// SetWindowsQualityUpdateProfiles sets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfileable)() {
    m.windowsQualityUpdateProfiles = value
}
// SetWindowsUpdateCatalogItems sets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItemable)() {
    m.windowsUpdateCatalogItems = value
}
// SetZebraFotaArtifacts sets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) SetZebraFotaArtifacts(value []ZebraFotaArtifactable)() {
    m.zebraFotaArtifacts = value
}
// SetZebraFotaConnector sets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) SetZebraFotaConnector(value ZebraFotaConnectorable)() {
    m.zebraFotaConnector = value
}
// SetZebraFotaDeployments sets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) SetZebraFotaDeployments(value []ZebraFotaDeploymentable)() {
    m.zebraFotaDeployments = value
}
