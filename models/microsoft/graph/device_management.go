package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagement struct {
    Entity
    // The date & time when tenant data moved between scaleunits.
    accountMoveCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin consent information.
    adminConsent *AdminConsent;
    // The summary state of ATP onboarding state for this account.
    advancedThreatProtectionOnboardingStateSummary *AdvancedThreatProtectionOnboardingStateSummary;
    // Android device owner enrollment profile entities.
    androidDeviceOwnerEnrollmentProfiles []AndroidDeviceOwnerEnrollmentProfile;
    // Android for Work app configuration schema entities.
    androidForWorkAppConfigurationSchemas []AndroidForWorkAppConfigurationSchema;
    // Android for Work enrollment profile entities.
    androidForWorkEnrollmentProfiles []AndroidForWorkEnrollmentProfile;
    // The singleton Android for Work settings entity.
    androidForWorkSettings *AndroidForWorkSettings;
    // The singleton Android managed store account enterprise settings entity.
    androidManagedStoreAccountEnterpriseSettings *AndroidManagedStoreAccountEnterpriseSettings;
    // Android Enterprise app configuration schema entities.
    androidManagedStoreAppConfigurationSchemas []AndroidManagedStoreAppConfigurationSchema;
    // Apple push notification certificate.
    applePushNotificationCertificate *ApplePushNotificationCertificate;
    // Apple user initiated enrollment profiles
    appleUserInitiatedEnrollmentProfiles []AppleUserInitiatedEnrollmentProfile;
    // The list of assignment filters
    assignmentFilters []DeviceAndAppManagementAssignmentFilter;
    // The Audit Events
    auditEvents []AuditEvent;
    // The list of autopilot events for the tenant.
    autopilotEvents []DeviceManagementAutopilotEvent;
    // The Cart To Class Associations.
    cartToClassAssociations []CartToClassAssociation;
    // The available categories
    categories []DeviceManagementSettingCategory;
    // Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
    certificateConnectorDetails []CertificateConnectorDetails;
    // Collection of ChromeOSOnboardingSettings settings associated with account.
    chromeOSOnboardingSettings []ChromeOSOnboardingSettings;
    // The list of CloudPC Connectivity Issue.
    cloudPCConnectivityIssues []CloudPCConnectivityIssue;
    // The list of co-managed devices report
    comanagedDevices []ManagedDevice;
    // The list of co-management eligible devices report
    comanagementEligibleDevices []ComanagementEligibleDevice;
    // List of all compliance categories
    complianceCategories []DeviceManagementConfigurationCategory;
    // The list of Compliance Management Partners configured by the tenant.
    complianceManagementPartners []ComplianceManagementPartner;
    // List of all compliance policies
    compliancePolicies []DeviceManagementCompliancePolicy;
    // List of all ComplianceSettings
    complianceSettings []DeviceManagementConfigurationSettingDefinition;
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    // A list of ConfigManagerCollection
    configManagerCollections []ConfigManagerCollection;
    // List of all Configuration Categories
    configurationCategories []DeviceManagementConfigurationCategory;
    // List of all Configuration policies
    configurationPolicies []DeviceManagementConfigurationPolicy;
    // List of all templates
    configurationPolicyTemplates []DeviceManagementConfigurationPolicyTemplate;
    // List of all ConfigurationSettings
    configurationSettings []DeviceManagementConfigurationSettingDefinition;
    // Data sharing consents.
    dataSharingConsents []DataSharingConsent;
    // This collections of multiple DEP tokens per-tenant.
    depOnboardingSettings []DepOnboardingSetting;
    // Collection of Derived credential settings associated with account.
    derivedCredentials []DeviceManagementDerivedCredentialSettings;
    // The list of detected apps associated with a device.
    detectedApps []DetectedApp;
    // The list of device categories with the tenant.
    deviceCategories []DeviceCategory;
    // The device compliance policies.
    deviceCompliancePolicies []DeviceCompliancePolicy;
    // The device compliance state summary for this account.
    deviceCompliancePolicyDeviceStateSummary *DeviceCompliancePolicyDeviceStateSummary;
    // The summary states of compliance policy settings for this account.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummary;
    // The last requested time of device compliance reporting for this account. This property is read-only.
    deviceComplianceReportSummarizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The list of device compliance scripts associated with the tenant.
    deviceComplianceScripts []DeviceComplianceScript;
    // Summary of policies in conflict state for this account.
    deviceConfigurationConflictSummary []DeviceConfigurationConflictSummary;
    // The device configuration device state summary for this account.
    deviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary;
    // Restricted apps violations for this account.
    deviceConfigurationRestrictedAppsViolations []RestrictedAppsViolation;
    // The device configurations.
    deviceConfigurations []DeviceConfiguration;
    // Summary of all certificates for all devices.
    deviceConfigurationsAllManagedDeviceCertificateStates []ManagedAllDeviceCertificateState;
    // The device configuration user state summary for this account.
    deviceConfigurationUserStateSummaries *DeviceConfigurationUserStateSummary;
    // The list of device custom attribute shell scripts associated with the tenant.
    deviceCustomAttributeShellScripts []DeviceCustomAttributeShellScript;
    // The list of device enrollment configurations
    deviceEnrollmentConfigurations []DeviceEnrollmentConfiguration;
    // The list of device health scripts associated with the tenant.
    deviceHealthScripts []DeviceHealthScript;
    // The list of Device Management Partners configured by the tenant.
    deviceManagementPartners []DeviceManagementPartner;
    // The list of device management scripts associated with the tenant.
    deviceManagementScripts []DeviceManagementScript;
    // Device protection overview.
    deviceProtectionOverview *DeviceProtectionOverview;
    // The list of device shell scripts associated with the tenant.
    deviceShellScripts []DeviceShellScript;
    // A list of connector objects.
    domainJoinConnectors []DeviceManagementDomainJoinConnector;
    // The embedded SIM activation code pools created by this account.
    embeddedSIMActivationCodePools []EmbeddedSIMActivationCodePool;
    // The list of Exchange Connectors configured by the tenant.
    exchangeConnectors []DeviceManagementExchangeConnector;
    // The list of Exchange On Premisis policies configured by the tenant.
    exchangeOnPremisesPolicies []DeviceManagementExchangeOnPremisesPolicy;
    // The policy which controls mobile device access to Exchange On Premises
    exchangeOnPremisesPolicy *DeviceManagementExchangeOnPremisesPolicy;
    // The available group policy categories for this account.
    groupPolicyCategories []GroupPolicyCategory;
    // The group policy configurations created by this account.
    groupPolicyConfigurations []GroupPolicyConfiguration;
    // The available group policy definition files for this account.
    groupPolicyDefinitionFiles []GroupPolicyDefinitionFile;
    // The available group policy definitions for this account.
    groupPolicyDefinitions []GroupPolicyDefinition;
    // A list of Group Policy migration reports.
    groupPolicyMigrationReports []GroupPolicyMigrationReport;
    // A list of Group Policy Object files uploaded.
    groupPolicyObjectFiles []GroupPolicyObjectFile;
    // The available group policy uploaded definition files for this account.
    groupPolicyUploadedDefinitionFiles []GroupPolicyUploadedDefinitionFile;
    // The imported device identities.
    importedDeviceIdentities []ImportedDeviceIdentity;
    // Collection of imported Windows autopilot devices.
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentity;
    // The device management intents
    intents []DeviceManagementIntent;
    // Intune Account Id for given tenant
    intuneAccountId *string;
    // intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
    intuneBrand *IntuneBrand;
    // Intune branding profiles targeted to AAD groups
    intuneBrandingProfiles []IntuneBrandingProfile;
    // The IOS software update installation statuses for this account.
    iosUpdateStatuses []IosUpdateDeviceStatus;
    // The last modified time of reporting for this account. This property is read-only.
    lastReportAggregationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
    legacyPcManangementEnabled *bool;
    // The MacOS software update account summaries for this account.
    macOSSoftwareUpdateAccountSummaries []MacOSSoftwareUpdateAccountSummary;
    // Device cleanup rule
    managedDeviceCleanupSettings *ManagedDeviceCleanupSettings;
    // Encryption report for devices in this account
    managedDeviceEncryptionStates []ManagedDeviceEncryptionState;
    // Device overview
    managedDeviceOverview *ManagedDeviceOverview;
    // The list of managed devices.
    managedDevices []ManagedDevice;
    // The management conditions associated with device management of the company.
    managementConditions []ManagementCondition;
    // The management condition statements associated with device management of the company.
    managementConditionStatements []ManagementConditionStatement;
    // Maximum number of DEP tokens allowed per-tenant.
    maximumDepTokens *int32;
    // Collection of MicrosoftTunnelConfiguration settings associated with account.
    microsoftTunnelConfigurations []MicrosoftTunnelConfiguration;
    // Collection of MicrosoftTunnelHealthThreshold settings associated with account.
    microsoftTunnelHealthThresholds []MicrosoftTunnelHealthThreshold;
    // Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
    microsoftTunnelServerLogCollectionResponses []MicrosoftTunnelServerLogCollectionResponse;
    // Collection of MicrosoftTunnelSite settings associated with account.
    microsoftTunnelSites []MicrosoftTunnelSite;
    // The collection property of MobileAppTroubleshootingEvent.
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEvent;
    // The list of Mobile threat Defense connectors configured by the tenant.
    mobileThreatDefenseConnectors []MobileThreatDefenseConnector;
    // The collection of Ndes connectors for this account.
    ndesConnectors []NdesConnector;
    // The Notification Message Templates.
    notificationMessageTemplates []NotificationMessageTemplate;
    // The list of device remote action audits with the tenant.
    remoteActionAudits []RemoteActionAudit;
    // The remote assist partners.
    remoteAssistancePartners []RemoteAssistancePartner;
    // The remote assistance settings singleton
    remoteAssistanceSettings *RemoteAssistanceSettings;
    // Reports singleton
    reports *DeviceManagementReports;
    // Collection of resource access settings associated with account.
    resourceAccessProfiles []DeviceManagementResourceAccessProfileBase;
    // The Resource Operations.
    resourceOperations []ResourceOperation;
    // List of all reusable settings that can be referred in a policy
    reusablePolicySettings []DeviceManagementReusablePolicySetting;
    // List of all reusable settings
    reusableSettings []DeviceManagementConfigurationSettingDefinition;
    // The Role Assignments.
    roleAssignments []DeviceAndAppManagementRoleAssignment;
    // The Role Definitions.
    roleDefinitions []RoleDefinition;
    // The Role Scope Tags.
    roleScopeTags []RoleScopeTag;
    // The device management intent setting definitions
    settingDefinitions []DeviceManagementSettingDefinition;
    // Account level settings.
    settings *DeviceManagementSettings;
    // The software update status summary.
    softwareUpdateStatusSummary *SoftwareUpdateStatusSummary;
    // Tenant's Subscription. Possible values are: none, intune, office365, intunePremium, intune_EDU, intune_SMB.
    subscriptions *DeviceManagementSubscriptions;
    // Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
    subscriptionState *DeviceManagementSubscriptionState;
    // The telecom expense management partners.
    telecomExpenseManagementPartners []TelecomExpenseManagementPartner;
    // The available templates
    templates []DeviceManagementTemplate;
    // List of all TemplateSettings
    templateSettings []DeviceManagementConfigurationSettingTemplate;
    // The terms and conditions associated with device management of the company.
    termsAndConditions []TermsAndConditions;
    // The list of troubleshooting events for the tenant.
    troubleshootingEvents []DeviceManagementTroubleshootingEvent;
    // When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
    unlicensedAdminstratorsEnabled *bool;
    // User experience analytics appHealth Application Performance
    userExperienceAnalyticsAppHealthApplicationPerformance []UserExperienceAnalyticsAppHealthApplicationPerformance;
    // User experience analytics appHealth Application Performance by App Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion;
    // User experience analytics appHealth Application Performance by App Version details
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails;
    // User experience analytics appHealth Application Performance by App Version Device Id
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId;
    // User experience analytics appHealth Application Performance by OS Version
    userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion;
    // User experience analytics appHealth Model Performance
    userExperienceAnalyticsAppHealthDeviceModelPerformance []UserExperienceAnalyticsAppHealthDeviceModelPerformance;
    // User experience analytics appHealth Device Performance
    userExperienceAnalyticsAppHealthDevicePerformance []UserExperienceAnalyticsAppHealthDevicePerformance;
    // User experience analytics device performance details
    userExperienceAnalyticsAppHealthDevicePerformanceDetails []UserExperienceAnalyticsAppHealthDevicePerformanceDetails;
    // User experience analytics appHealth OS version Performance
    userExperienceAnalyticsAppHealthOSVersionPerformance []UserExperienceAnalyticsAppHealthOSVersionPerformance;
    // User experience analytics appHealth overview
    userExperienceAnalyticsAppHealthOverview *UserExperienceAnalyticsCategory;
    // User experience analytics baselines
    userExperienceAnalyticsBaselines []UserExperienceAnalyticsBaseline;
    // User Experience Analytics Battery Health App Impact
    userExperienceAnalyticsBatteryHealthAppImpact []UserExperienceAnalyticsBatteryHealthAppImpact;
    // User Experience Analytics Battery Health Capacity Details
    userExperienceAnalyticsBatteryHealthCapacityDetails *UserExperienceAnalyticsBatteryHealthCapacityDetails;
    // User Experience Analytics Battery Health Device App Impact
    userExperienceAnalyticsBatteryHealthDeviceAppImpact []UserExperienceAnalyticsBatteryHealthDeviceAppImpact;
    // User Experience Analytics Battery Health Device Performance
    userExperienceAnalyticsBatteryHealthDevicePerformance []UserExperienceAnalyticsBatteryHealthDevicePerformance;
    // User Experience Analytics Battery Health Device Runtime History
    userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory;
    // User Experience Analytics Battery Health Model Performance
    userExperienceAnalyticsBatteryHealthModelPerformance []UserExperienceAnalyticsBatteryHealthModelPerformance;
    // User Experience Analytics Battery Health Os Performance
    userExperienceAnalyticsBatteryHealthOsPerformance []UserExperienceAnalyticsBatteryHealthOsPerformance;
    // User Experience Analytics Battery Health Runtime Details
    userExperienceAnalyticsBatteryHealthRuntimeDetails *UserExperienceAnalyticsBatteryHealthRuntimeDetails;
    // User experience analytics categories
    userExperienceAnalyticsCategories []UserExperienceAnalyticsCategory;
    // User experience analytics device metric history
    userExperienceAnalyticsDeviceMetricHistory []UserExperienceAnalyticsMetricHistory;
    // User experience analytics device performance
    userExperienceAnalyticsDevicePerformance []UserExperienceAnalyticsDevicePerformance;
    // User experience analytics device scores
    userExperienceAnalyticsDeviceScores []UserExperienceAnalyticsDeviceScores;
    // User experience analytics device Startup History
    userExperienceAnalyticsDeviceStartupHistory []UserExperienceAnalyticsDeviceStartupHistory;
    // User experience analytics device Startup Processes
    userExperienceAnalyticsDeviceStartupProcesses []UserExperienceAnalyticsDeviceStartupProcess;
    // User experience analytics device Startup Process Performance
    userExperienceAnalyticsDeviceStartupProcessPerformance []UserExperienceAnalyticsDeviceStartupProcessPerformance;
    // User experience analytics devices without cloud identity.
    userExperienceAnalyticsDevicesWithoutCloudIdentity []UserExperienceAnalyticsDeviceWithoutCloudIdentity;
    // User experience analytics impacting process
    userExperienceAnalyticsImpactingProcess []UserExperienceAnalyticsImpactingProcess;
    // User experience analytics metric history
    userExperienceAnalyticsMetricHistory []UserExperienceAnalyticsMetricHistory;
    // User experience analytics model scores
    userExperienceAnalyticsModelScores []UserExperienceAnalyticsModelScores;
    // User experience analytics devices not Windows Autopilot ready.
    userExperienceAnalyticsNotAutopilotReadyDevice []UserExperienceAnalyticsNotAutopilotReadyDevice;
    // User experience analytics overview
    userExperienceAnalyticsOverview *UserExperienceAnalyticsOverview;
    // User experience analytics regression summary
    userExperienceAnalyticsRegressionSummary *UserExperienceAnalyticsRegressionSummary;
    // User experience analytics remote connection
    userExperienceAnalyticsRemoteConnection []UserExperienceAnalyticsRemoteConnection;
    // User experience analytics resource performance
    userExperienceAnalyticsResourcePerformance []UserExperienceAnalyticsResourcePerformance;
    // User experience analytics device Startup Score History
    userExperienceAnalyticsScoreHistory []UserExperienceAnalyticsScoreHistory;
    // User experience analytics device settings
    userExperienceAnalyticsSettings *UserExperienceAnalyticsSettings;
    // User experience analytics work from anywhere hardware readiness metrics.
    userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric;
    // User experience analytics work from anywhere metrics.
    userExperienceAnalyticsWorkFromAnywhereMetrics []UserExperienceAnalyticsWorkFromAnywhereMetric;
    // Collection of PFX certificates associated with a user.
    userPfxCertificates []UserPFXCertificate;
    // 
    virtualEndpoint *VirtualEndpoint;
    // Windows auto pilot deployment profiles
    windowsAutopilotDeploymentProfiles []WindowsAutopilotDeploymentProfile;
    // The Windows autopilot device identities contained collection.
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentity;
    // The Windows autopilot account settings.
    windowsAutopilotSettings *WindowsAutopilotSettings;
    // A collection of windows driver update profiles
    windowsDriverUpdateProfiles []WindowsDriverUpdateProfile;
    // A collection of windows feature update profiles
    windowsFeatureUpdateProfiles []WindowsFeatureUpdateProfile;
    // The windows information protection app learning summaries.
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummary;
    // The windows information protection network learning summaries.
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummary;
    // The list of affected malware in the tenant.
    windowsMalwareInformation []WindowsMalwareInformation;
    // Malware overview for windows devices.
    windowsMalwareOverview *WindowsMalwareOverview;
    // A collection of windows quality update profiles
    windowsQualityUpdateProfiles []WindowsQualityUpdateProfile;
    // A collection of windows update catalog items (fetaure updates item , quality updates item)
    windowsUpdateCatalogItems []WindowsUpdateCatalogItem;
}
// Instantiates a new deviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.accountMoveCompletionDateTime
    }
}
// Gets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) GetAdminConsent()(*AdminConsent) {
    if m == nil {
        return nil
    } else {
        return m.adminConsent
    }
}
// Gets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) GetAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionOnboardingStateSummary
    }
}
// Gets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceOwnerEnrollmentProfiles
    }
}
// Gets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchema) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkAppConfigurationSchemas
    }
}
// Gets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkEnrollmentProfiles
    }
}
// Gets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) GetAndroidForWorkSettings()(*AndroidForWorkSettings) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkSettings
    }
}
// Gets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) GetAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedStoreAccountEnterpriseSettings
    }
}
// Gets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchema) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedStoreAppConfigurationSchemas
    }
}
// Gets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) GetApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    if m == nil {
        return nil
    } else {
        return m.applePushNotificationCertificate
    }
}
// Gets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.appleUserInitiatedEnrollmentProfiles
    }
}
// Gets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilter) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilters
    }
}
// Gets the auditEvents property value. The Audit Events
func (m *DeviceManagement) GetAuditEvents()([]AuditEvent) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
// Gets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) GetAutopilotEvents()([]DeviceManagementAutopilotEvent) {
    if m == nil {
        return nil
    } else {
        return m.autopilotEvents
    }
}
// Gets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) GetCartToClassAssociations()([]CartToClassAssociation) {
    if m == nil {
        return nil
    } else {
        return m.cartToClassAssociations
    }
}
// Gets the categories property value. The available categories
func (m *DeviceManagement) GetCategories()([]DeviceManagementSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) GetCertificateConnectorDetails()([]CertificateConnectorDetails) {
    if m == nil {
        return nil
    } else {
        return m.certificateConnectorDetails
    }
}
// Gets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettings) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSOnboardingSettings
    }
}
// Gets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssue) {
    if m == nil {
        return nil
    } else {
        return m.cloudPCConnectivityIssues
    }
}
// Gets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) GetComanagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.comanagedDevices
    }
}
// Gets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) GetComanagementEligibleDevices()([]ComanagementEligibleDevice) {
    if m == nil {
        return nil
    } else {
        return m.comanagementEligibleDevices
    }
}
// Gets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) GetComplianceCategories()([]DeviceManagementConfigurationCategory) {
    if m == nil {
        return nil
    } else {
        return m.complianceCategories
    }
}
// Gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.complianceManagementPartners
    }
}
// Gets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) GetCompliancePolicies()([]DeviceManagementCompliancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicies
    }
}
// Gets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.complianceSettings
    }
}
// Gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
// Gets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) GetConfigManagerCollections()([]ConfigManagerCollection) {
    if m == nil {
        return nil
    } else {
        return m.configManagerCollections
    }
}
// Gets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) GetConfigurationCategories()([]DeviceManagementConfigurationCategory) {
    if m == nil {
        return nil
    } else {
        return m.configurationCategories
    }
}
// Gets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) GetConfigurationPolicies()([]DeviceManagementConfigurationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.configurationPolicies
    }
}
// Gets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplate) {
    if m == nil {
        return nil
    } else {
        return m.configurationPolicyTemplates
    }
}
// Gets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.configurationSettings
    }
}
// Gets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) GetDataSharingConsents()([]DataSharingConsent) {
    if m == nil {
        return nil
    } else {
        return m.dataSharingConsents
    }
}
// Gets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) GetDepOnboardingSettings()([]DepOnboardingSetting) {
    if m == nil {
        return nil
    } else {
        return m.depOnboardingSettings
    }
}
// Gets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettings) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentials
    }
}
// Gets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) GetDetectedApps()([]DetectedApp) {
    if m == nil {
        return nil
    } else {
        return m.detectedApps
    }
}
// Gets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategories
    }
}
// Gets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicies
    }
}
// Gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyDeviceStateSummary
    }
}
// Gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
// Gets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceReportSummarizationDateTime
    }
}
// Gets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceComplianceScripts()([]DeviceComplianceScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScripts
    }
}
// Gets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationConflictSummary
    }
}
// Gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationDeviceStateSummaries
    }
}
// Gets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolation) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationRestrictedAppsViolations
    }
}
// Gets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurations
    }
}
// Gets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateState) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationsAllManagedDeviceCertificateStates
    }
}
// Gets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationUserStateSummaries()(*DeviceConfigurationUserStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationUserStateSummaries
    }
}
// Gets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceCustomAttributeShellScripts
    }
}
// Gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentConfigurations
    }
}
// Gets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceHealthScripts()([]DeviceHealthScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthScripts
    }
}
// Gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementPartners
    }
}
// Gets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceManagementScripts()([]DeviceManagementScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementScripts
    }
}
// Gets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) GetDeviceProtectionOverview()(*DeviceProtectionOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceProtectionOverview
    }
}
// Gets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceShellScripts()([]DeviceShellScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceShellScripts
    }
}
// Gets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnector) {
    if m == nil {
        return nil
    } else {
        return m.domainJoinConnectors
    }
}
// Gets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePool) {
    if m == nil {
        return nil
    } else {
        return m.embeddedSIMActivationCodePools
    }
}
// Gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnector) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectors
    }
}
// Gets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicy) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOnPremisesPolicies
    }
}
// Gets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) GetExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicy) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOnPremisesPolicy
    }
}
// Gets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) GetGroupPolicyCategories()([]GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCategories
    }
}
// Gets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) GetGroupPolicyConfigurations()([]GroupPolicyConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyConfigurations
    }
}
// Gets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyDefinitionFiles
    }
}
// Gets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyDefinitions
    }
}
// Gets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReport) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyMigrationReports
    }
}
// Gets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) GetGroupPolicyObjectFiles()([]GroupPolicyObjectFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectFiles
    }
}
// Gets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyUploadedDefinitionFiles
    }
}
// Gets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) GetImportedDeviceIdentities()([]ImportedDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentities
    }
}
// Gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedWindowsAutopilotDeviceIdentities
    }
}
// Gets the intents property value. The device management intents
func (m *DeviceManagement) GetIntents()([]DeviceManagementIntent) {
    if m == nil {
        return nil
    } else {
        return m.intents
    }
}
// Gets the intuneAccountId property value. Intune Account Id for given tenant
func (m *DeviceManagement) GetIntuneAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneAccountId
    }
}
// Gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) GetIntuneBrand()(*IntuneBrand) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrand
    }
}
// Gets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) GetIntuneBrandingProfiles()([]IntuneBrandingProfile) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrandingProfiles
    }
}
// Gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.iosUpdateStatuses
    }
}
// Gets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportAggregationDateTime
    }
}
// Gets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) GetLegacyPcManangementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.legacyPcManangementEnabled
    }
}
// Gets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummary) {
    if m == nil {
        return nil
    } else {
        return m.macOSSoftwareUpdateAccountSummaries
    }
}
// Gets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) GetManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCleanupSettings
    }
}
// Gets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionState) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceEncryptionStates
    }
}
// Gets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) GetManagedDeviceOverview()(*ManagedDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOverview
    }
}
// Gets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// Gets the managementConditions property value. The management conditions associated with device management of the company.
func (m *DeviceManagement) GetManagementConditions()([]ManagementCondition) {
    if m == nil {
        return nil
    } else {
        return m.managementConditions
    }
}
// Gets the managementConditionStatements property value. The management condition statements associated with device management of the company.
func (m *DeviceManagement) GetManagementConditionStatements()([]ManagementConditionStatement) {
    if m == nil {
        return nil
    } else {
        return m.managementConditionStatements
    }
}
// Gets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) GetMaximumDepTokens()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumDepTokens
    }
}
// Gets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfigurations
    }
}
// Gets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThreshold) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelHealthThresholds
    }
}
// Gets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponse) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServerLogCollectionResponses
    }
}
// Gets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelSites()([]MicrosoftTunnelSite) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelSites
    }
}
// Gets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppTroubleshootingEvents
    }
}
// Gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnector) {
    if m == nil {
        return nil
    } else {
        return m.mobileThreatDefenseConnectors
    }
}
// Gets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) GetNdesConnectors()([]NdesConnector) {
    if m == nil {
        return nil
    } else {
        return m.ndesConnectors
    }
}
// Gets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplate) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageTemplates
    }
}
// Gets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) GetRemoteActionAudits()([]RemoteActionAudit) {
    if m == nil {
        return nil
    } else {
        return m.remoteActionAudits
    }
}
// Gets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartner) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistancePartners
    }
}
// Gets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) GetRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSettings
    }
}
// Gets the reports property value. Reports singleton
func (m *DeviceManagement) GetReports()(*DeviceManagementReports) {
    if m == nil {
        return nil
    } else {
        return m.reports
    }
}
// Gets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBase) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccessProfiles
    }
}
// Gets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperation) {
    if m == nil {
        return nil
    } else {
        return m.resourceOperations
    }
}
// Gets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) GetReusablePolicySettings()([]DeviceManagementReusablePolicySetting) {
    if m == nil {
        return nil
    } else {
        return m.reusablePolicySettings
    }
}
// Gets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) GetReusableSettings()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.reusableSettings
    }
}
// Gets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Gets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) GetRoleScopeTags()([]RoleScopeTag) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTags
    }
}
// Gets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) GetSettingDefinitions()([]DeviceManagementSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// Gets the settings property value. Account level settings.
func (m *DeviceManagement) GetSettings()(*DeviceManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateStatusSummary
    }
}
// Gets the subscriptions property value. Tenant's Subscription. Possible values are: none, intune, office365, intunePremium, intune_EDU, intune_SMB.
func (m *DeviceManagement) GetSubscriptions()(*DeviceManagementSubscriptions) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
// Gets the subscriptionState property value. Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionState
    }
}
// Gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.telecomExpenseManagementPartners
    }
}
// Gets the templates property value. The available templates
func (m *DeviceManagement) GetTemplates()([]DeviceManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templates
    }
}
// Gets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templateSettings
    }
}
// Gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// Gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingEvents
    }
}
// Gets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) GetUnlicensedAdminstratorsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.unlicensedAdminstratorsEnabled
    }
}
// Gets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformance
    }
}
// Gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion
    }
}
// Gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
    }
}
// Gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
    }
}
// Gets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
    }
}
// Gets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDeviceModelPerformance
    }
}
// Gets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDevicePerformance
    }
}
// Gets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDevicePerformanceDetails
    }
}
// Gets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthOSVersionPerformance
    }
}
// Gets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthOverview
    }
}
// Gets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaseline) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBaselines
    }
}
// Gets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpact) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthAppImpact
    }
}
// Gets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetails) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthCapacityDetails
    }
}
// Gets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpact) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthDeviceAppImpact
    }
}
// Gets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthDevicePerformance
    }
}
// Gets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
    }
}
// Gets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthModelPerformance
    }
}
// Gets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthOsPerformance
    }
}
// Gets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetails) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBatteryHealthRuntimeDetails
    }
}
// Gets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsCategories
    }
}
// Gets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceMetricHistory
    }
}
// Gets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDevicePerformance
    }
}
// Gets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScores) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceScores
    }
}
// Gets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupHistory
    }
}
// Gets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcess) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupProcesses
    }
}
// Gets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupProcessPerformance
    }
}
// Gets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDevicesWithoutCloudIdentity
    }
}
// Gets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcess) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsImpactingProcess
    }
}
// Gets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetricHistory
    }
}
// Gets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScores) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsModelScores
    }
}
// Gets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDevice) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsNotAutopilotReadyDevice
    }
}
// Gets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsOverview
    }
}
// Gets the userExperienceAnalyticsRegressionSummary property value. User experience analytics regression summary
func (m *DeviceManagement) GetUserExperienceAnalyticsRegressionSummary()(*UserExperienceAnalyticsRegressionSummary) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsRegressionSummary
    }
}
// Gets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnection) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsRemoteConnection
    }
}
// Gets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsResourcePerformance
    }
}
// Gets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsScoreHistory
    }
}
// Gets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsSettings
    }
}
// Gets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
    }
}
// Gets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetric) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsWorkFromAnywhereMetrics
    }
}
// Gets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) GetUserPfxCertificates()([]UserPFXCertificate) {
    if m == nil {
        return nil
    } else {
        return m.userPfxCertificates
    }
}
// Gets the virtualEndpoint property value. 
func (m *DeviceManagement) GetVirtualEndpoint()(*VirtualEndpoint) {
    if m == nil {
        return nil
    } else {
        return m.virtualEndpoint
    }
}
// Gets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeploymentProfiles
    }
}
// Gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeviceIdentities
    }
}
// Gets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) GetWindowsAutopilotSettings()(*WindowsAutopilotSettings) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotSettings
    }
}
// Gets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsDriverUpdateProfiles
    }
}
// Gets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsFeatureUpdateProfiles
    }
}
// Gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionAppLearningSummaries
    }
}
// Gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionNetworkLearningSummaries
    }
}
// Gets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformation) {
    if m == nil {
        return nil
    } else {
        return m.windowsMalwareInformation
    }
}
// Gets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) GetWindowsMalwareOverview()(*WindowsMalwareOverview) {
    if m == nil {
        return nil
    } else {
        return m.windowsMalwareOverview
    }
}
// Gets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsQualityUpdateProfiles
    }
}
// Gets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItem) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateCatalogItems
    }
}
// The deserialization information for the current model
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountMoveCompletionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAccountMoveCompletionDateTime(val)
        return nil
    }
    res["adminConsent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdminConsent() })
        if err != nil {
            return err
        }
        m.SetAdminConsent(val.(*AdminConsent))
        return nil
    }
    res["advancedThreatProtectionOnboardingStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdvancedThreatProtectionOnboardingStateSummary() })
        if err != nil {
            return err
        }
        m.SetAdvancedThreatProtectionOnboardingStateSummary(val.(*AdvancedThreatProtectionOnboardingStateSummary))
        return nil
    }
    res["androidDeviceOwnerEnrollmentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidDeviceOwnerEnrollmentProfile() })
        if err != nil {
            return err
        }
        res := make([]AndroidDeviceOwnerEnrollmentProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidDeviceOwnerEnrollmentProfile))
        }
        m.SetAndroidDeviceOwnerEnrollmentProfiles(res)
        return nil
    }
    res["androidForWorkAppConfigurationSchemas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidForWorkAppConfigurationSchema() })
        if err != nil {
            return err
        }
        res := make([]AndroidForWorkAppConfigurationSchema, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidForWorkAppConfigurationSchema))
        }
        m.SetAndroidForWorkAppConfigurationSchemas(res)
        return nil
    }
    res["androidForWorkEnrollmentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidForWorkEnrollmentProfile() })
        if err != nil {
            return err
        }
        res := make([]AndroidForWorkEnrollmentProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidForWorkEnrollmentProfile))
        }
        m.SetAndroidForWorkEnrollmentProfiles(res)
        return nil
    }
    res["androidForWorkSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidForWorkSettings() })
        if err != nil {
            return err
        }
        m.SetAndroidForWorkSettings(val.(*AndroidForWorkSettings))
        return nil
    }
    res["androidManagedStoreAccountEnterpriseSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedStoreAccountEnterpriseSettings() })
        if err != nil {
            return err
        }
        m.SetAndroidManagedStoreAccountEnterpriseSettings(val.(*AndroidManagedStoreAccountEnterpriseSettings))
        return nil
    }
    res["androidManagedStoreAppConfigurationSchemas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedStoreAppConfigurationSchema() })
        if err != nil {
            return err
        }
        res := make([]AndroidManagedStoreAppConfigurationSchema, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidManagedStoreAppConfigurationSchema))
        }
        m.SetAndroidManagedStoreAppConfigurationSchemas(res)
        return nil
    }
    res["applePushNotificationCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplePushNotificationCertificate() })
        if err != nil {
            return err
        }
        m.SetApplePushNotificationCertificate(val.(*ApplePushNotificationCertificate))
        return nil
    }
    res["appleUserInitiatedEnrollmentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppleUserInitiatedEnrollmentProfile() })
        if err != nil {
            return err
        }
        res := make([]AppleUserInitiatedEnrollmentProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppleUserInitiatedEnrollmentProfile))
        }
        m.SetAppleUserInitiatedEnrollmentProfiles(res)
        return nil
    }
    res["assignmentFilters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentFilter() })
        if err != nil {
            return err
        }
        res := make([]DeviceAndAppManagementAssignmentFilter, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceAndAppManagementAssignmentFilter))
        }
        m.SetAssignmentFilters(res)
        return nil
    }
    res["auditEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuditEvent() })
        if err != nil {
            return err
        }
        res := make([]AuditEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuditEvent))
        }
        m.SetAuditEvents(res)
        return nil
    }
    res["autopilotEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementAutopilotEvent() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementAutopilotEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementAutopilotEvent))
        }
        m.SetAutopilotEvents(res)
        return nil
    }
    res["cartToClassAssociations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCartToClassAssociation() })
        if err != nil {
            return err
        }
        res := make([]CartToClassAssociation, len(val))
        for i, v := range val {
            res[i] = *(v.(*CartToClassAssociation))
        }
        m.SetCartToClassAssociations(res)
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingCategory))
        }
        m.SetCategories(res)
        return nil
    }
    res["certificateConnectorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCertificateConnectorDetails() })
        if err != nil {
            return err
        }
        res := make([]CertificateConnectorDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*CertificateConnectorDetails))
        }
        m.SetCertificateConnectorDetails(res)
        return nil
    }
    res["chromeOSOnboardingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChromeOSOnboardingSettings() })
        if err != nil {
            return err
        }
        res := make([]ChromeOSOnboardingSettings, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChromeOSOnboardingSettings))
        }
        m.SetChromeOSOnboardingSettings(res)
        return nil
    }
    res["cloudPCConnectivityIssues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPCConnectivityIssue() })
        if err != nil {
            return err
        }
        res := make([]CloudPCConnectivityIssue, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPCConnectivityIssue))
        }
        m.SetCloudPCConnectivityIssues(res)
        return nil
    }
    res["comanagedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        res := make([]ManagedDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDevice))
        }
        m.SetComanagedDevices(res)
        return nil
    }
    res["comanagementEligibleDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComanagementEligibleDevice() })
        if err != nil {
            return err
        }
        res := make([]ComanagementEligibleDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComanagementEligibleDevice))
        }
        m.SetComanagementEligibleDevices(res)
        return nil
    }
    res["complianceCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationCategory))
        }
        m.SetComplianceCategories(res)
        return nil
    }
    res["complianceManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]ComplianceManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceManagementPartner))
        }
        m.SetComplianceManagementPartners(res)
        return nil
    }
    res["compliancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementCompliancePolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementCompliancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementCompliancePolicy))
        }
        m.SetCompliancePolicies(res)
        return nil
    }
    res["complianceSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationSettingDefinition))
        }
        m.SetComplianceSettings(res)
        return nil
    }
    res["conditionalAccessSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesConditionalAccessSettings() })
        if err != nil {
            return err
        }
        m.SetConditionalAccessSettings(val.(*OnPremisesConditionalAccessSettings))
        return nil
    }
    res["configManagerCollections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigManagerCollection() })
        if err != nil {
            return err
        }
        res := make([]ConfigManagerCollection, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConfigManagerCollection))
        }
        m.SetConfigManagerCollections(res)
        return nil
    }
    res["configurationCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationCategory))
        }
        m.SetConfigurationCategories(res)
        return nil
    }
    res["configurationPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationPolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationPolicy))
        }
        m.SetConfigurationPolicies(res)
        return nil
    }
    res["configurationPolicyTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationPolicyTemplate() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationPolicyTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationPolicyTemplate))
        }
        m.SetConfigurationPolicyTemplates(res)
        return nil
    }
    res["configurationSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationSettingDefinition))
        }
        m.SetConfigurationSettings(res)
        return nil
    }
    res["dataSharingConsents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDataSharingConsent() })
        if err != nil {
            return err
        }
        res := make([]DataSharingConsent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DataSharingConsent))
        }
        m.SetDataSharingConsents(res)
        return nil
    }
    res["depOnboardingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDepOnboardingSetting() })
        if err != nil {
            return err
        }
        res := make([]DepOnboardingSetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*DepOnboardingSetting))
        }
        m.SetDepOnboardingSettings(res)
        return nil
    }
    res["derivedCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementDerivedCredentialSettings() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementDerivedCredentialSettings, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementDerivedCredentialSettings))
        }
        m.SetDerivedCredentials(res)
        return nil
    }
    res["detectedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedApp() })
        if err != nil {
            return err
        }
        res := make([]DetectedApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*DetectedApp))
        }
        m.SetDetectedApps(res)
        return nil
    }
    res["deviceCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCategory))
        }
        m.SetDeviceCategories(res)
        return nil
    }
    res["deviceCompliancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicy))
        }
        m.SetDeviceCompliancePolicies(res)
        return nil
    }
    res["deviceCompliancePolicyDeviceStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyDeviceStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceCompliancePolicyDeviceStateSummary(val.(*DeviceCompliancePolicyDeviceStateSummary))
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicySettingStateSummary() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicySettingStateSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicySettingStateSummary))
        }
        m.SetDeviceCompliancePolicySettingStateSummaries(res)
        return nil
    }
    res["deviceComplianceReportSummarizationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeviceComplianceReportSummarizationDateTime(val)
        return nil
    }
    res["deviceComplianceScripts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceComplianceScript() })
        if err != nil {
            return err
        }
        res := make([]DeviceComplianceScript, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceComplianceScript))
        }
        m.SetDeviceComplianceScripts(res)
        return nil
    }
    res["deviceConfigurationConflictSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationConflictSummary() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationConflictSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationConflictSummary))
        }
        m.SetDeviceConfigurationConflictSummary(res)
        return nil
    }
    res["deviceConfigurationDeviceStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationDeviceStateSummaries(val.(*DeviceConfigurationDeviceStateSummary))
        return nil
    }
    res["deviceConfigurationRestrictedAppsViolations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRestrictedAppsViolation() })
        if err != nil {
            return err
        }
        res := make([]RestrictedAppsViolation, len(val))
        for i, v := range val {
            res[i] = *(v.(*RestrictedAppsViolation))
        }
        m.SetDeviceConfigurationRestrictedAppsViolations(res)
        return nil
    }
    res["deviceConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfiguration() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfiguration))
        }
        m.SetDeviceConfigurations(res)
        return nil
    }
    res["deviceConfigurationsAllManagedDeviceCertificateStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAllDeviceCertificateState() })
        if err != nil {
            return err
        }
        res := make([]ManagedAllDeviceCertificateState, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAllDeviceCertificateState))
        }
        m.SetDeviceConfigurationsAllManagedDeviceCertificateStates(res)
        return nil
    }
    res["deviceConfigurationUserStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationUserStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationUserStateSummaries(val.(*DeviceConfigurationUserStateSummary))
        return nil
    }
    res["deviceCustomAttributeShellScripts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCustomAttributeShellScript() })
        if err != nil {
            return err
        }
        res := make([]DeviceCustomAttributeShellScript, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCustomAttributeShellScript))
        }
        m.SetDeviceCustomAttributeShellScripts(res)
        return nil
    }
    res["deviceEnrollmentConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceEnrollmentConfiguration() })
        if err != nil {
            return err
        }
        res := make([]DeviceEnrollmentConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceEnrollmentConfiguration))
        }
        m.SetDeviceEnrollmentConfigurations(res)
        return nil
    }
    res["deviceHealthScripts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthScript() })
        if err != nil {
            return err
        }
        res := make([]DeviceHealthScript, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceHealthScript))
        }
        m.SetDeviceHealthScripts(res)
        return nil
    }
    res["deviceManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementPartner))
        }
        m.SetDeviceManagementPartners(res)
        return nil
    }
    res["deviceManagementScripts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScript() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScript, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScript))
        }
        m.SetDeviceManagementScripts(res)
        return nil
    }
    res["deviceProtectionOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceProtectionOverview() })
        if err != nil {
            return err
        }
        m.SetDeviceProtectionOverview(val.(*DeviceProtectionOverview))
        return nil
    }
    res["deviceShellScripts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceShellScript() })
        if err != nil {
            return err
        }
        res := make([]DeviceShellScript, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceShellScript))
        }
        m.SetDeviceShellScripts(res)
        return nil
    }
    res["domainJoinConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementDomainJoinConnector() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementDomainJoinConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementDomainJoinConnector))
        }
        m.SetDomainJoinConnectors(res)
        return nil
    }
    res["embeddedSIMActivationCodePools"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmbeddedSIMActivationCodePool() })
        if err != nil {
            return err
        }
        res := make([]EmbeddedSIMActivationCodePool, len(val))
        for i, v := range val {
            res[i] = *(v.(*EmbeddedSIMActivationCodePool))
        }
        m.SetEmbeddedSIMActivationCodePools(res)
        return nil
    }
    res["exchangeConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeConnector() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExchangeConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExchangeConnector))
        }
        m.SetExchangeConnectors(res)
        return nil
    }
    res["exchangeOnPremisesPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeOnPremisesPolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExchangeOnPremisesPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExchangeOnPremisesPolicy))
        }
        m.SetExchangeOnPremisesPolicies(res)
        return nil
    }
    res["exchangeOnPremisesPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeOnPremisesPolicy() })
        if err != nil {
            return err
        }
        m.SetExchangeOnPremisesPolicy(val.(*DeviceManagementExchangeOnPremisesPolicy))
        return nil
    }
    res["groupPolicyCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyCategory() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyCategory))
        }
        m.SetGroupPolicyCategories(res)
        return nil
    }
    res["groupPolicyConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyConfiguration() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyConfiguration))
        }
        m.SetGroupPolicyConfigurations(res)
        return nil
    }
    res["groupPolicyDefinitionFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinitionFile() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyDefinitionFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyDefinitionFile))
        }
        m.SetGroupPolicyDefinitionFiles(res)
        return nil
    }
    res["groupPolicyDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyDefinition() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyDefinition))
        }
        m.SetGroupPolicyDefinitions(res)
        return nil
    }
    res["groupPolicyMigrationReports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyMigrationReport() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyMigrationReport, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyMigrationReport))
        }
        m.SetGroupPolicyMigrationReports(res)
        return nil
    }
    res["groupPolicyObjectFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyObjectFile() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyObjectFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyObjectFile))
        }
        m.SetGroupPolicyObjectFiles(res)
        return nil
    }
    res["groupPolicyUploadedDefinitionFiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicyUploadedDefinitionFile() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicyUploadedDefinitionFile, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicyUploadedDefinitionFile))
        }
        m.SetGroupPolicyUploadedDefinitionFiles(res)
        return nil
    }
    res["importedDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]ImportedDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ImportedDeviceIdentity))
        }
        m.SetImportedDeviceIdentities(res)
        return nil
    }
    res["importedWindowsAutopilotDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedWindowsAutopilotDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]ImportedWindowsAutopilotDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ImportedWindowsAutopilotDeviceIdentity))
        }
        m.SetImportedWindowsAutopilotDeviceIdentities(res)
        return nil
    }
    res["intents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementIntent() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementIntent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementIntent))
        }
        m.SetIntents(res)
        return nil
    }
    res["intuneAccountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIntuneAccountId(val)
        return nil
    }
    res["intuneBrand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntuneBrand() })
        if err != nil {
            return err
        }
        m.SetIntuneBrand(val.(*IntuneBrand))
        return nil
    }
    res["intuneBrandingProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntuneBrandingProfile() })
        if err != nil {
            return err
        }
        res := make([]IntuneBrandingProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*IntuneBrandingProfile))
        }
        m.SetIntuneBrandingProfiles(res)
        return nil
    }
    res["iosUpdateStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosUpdateDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]IosUpdateDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*IosUpdateDeviceStatus))
        }
        m.SetIosUpdateStatuses(res)
        return nil
    }
    res["lastReportAggregationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportAggregationDateTime(val)
        return nil
    }
    res["legacyPcManangementEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetLegacyPcManangementEnabled(val)
        return nil
    }
    res["macOSSoftwareUpdateAccountSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMacOSSoftwareUpdateAccountSummary() })
        if err != nil {
            return err
        }
        res := make([]MacOSSoftwareUpdateAccountSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*MacOSSoftwareUpdateAccountSummary))
        }
        m.SetMacOSSoftwareUpdateAccountSummaries(res)
        return nil
    }
    res["managedDeviceCleanupSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceCleanupSettings() })
        if err != nil {
            return err
        }
        m.SetManagedDeviceCleanupSettings(val.(*ManagedDeviceCleanupSettings))
        return nil
    }
    res["managedDeviceEncryptionStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceEncryptionState() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceEncryptionState, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceEncryptionState))
        }
        m.SetManagedDeviceEncryptionStates(res)
        return nil
    }
    res["managedDeviceOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceOverview() })
        if err != nil {
            return err
        }
        m.SetManagedDeviceOverview(val.(*ManagedDeviceOverview))
        return nil
    }
    res["managedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        res := make([]ManagedDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDevice))
        }
        m.SetManagedDevices(res)
        return nil
    }
    res["managementConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementCondition() })
        if err != nil {
            return err
        }
        res := make([]ManagementCondition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementCondition))
        }
        m.SetManagementConditions(res)
        return nil
    }
    res["managementConditionStatements"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementConditionStatement() })
        if err != nil {
            return err
        }
        res := make([]ManagementConditionStatement, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementConditionStatement))
        }
        m.SetManagementConditionStatements(res)
        return nil
    }
    res["maximumDepTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaximumDepTokens(val)
        return nil
    }
    res["microsoftTunnelConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelConfiguration() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftTunnelConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftTunnelConfiguration))
        }
        m.SetMicrosoftTunnelConfigurations(res)
        return nil
    }
    res["microsoftTunnelHealthThresholds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelHealthThreshold() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftTunnelHealthThreshold, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftTunnelHealthThreshold))
        }
        m.SetMicrosoftTunnelHealthThresholds(res)
        return nil
    }
    res["microsoftTunnelServerLogCollectionResponses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelServerLogCollectionResponse() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftTunnelServerLogCollectionResponse, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftTunnelServerLogCollectionResponse))
        }
        m.SetMicrosoftTunnelServerLogCollectionResponses(res)
        return nil
    }
    res["microsoftTunnelSites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelSite() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftTunnelSite, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftTunnelSite))
        }
        m.SetMicrosoftTunnelSites(res)
        return nil
    }
    res["mobileAppTroubleshootingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppTroubleshootingEvent() })
        if err != nil {
            return err
        }
        res := make([]MobileAppTroubleshootingEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppTroubleshootingEvent))
        }
        m.SetMobileAppTroubleshootingEvents(res)
        return nil
    }
    res["mobileThreatDefenseConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileThreatDefenseConnector() })
        if err != nil {
            return err
        }
        res := make([]MobileThreatDefenseConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileThreatDefenseConnector))
        }
        m.SetMobileThreatDefenseConnectors(res)
        return nil
    }
    res["ndesConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNdesConnector() })
        if err != nil {
            return err
        }
        res := make([]NdesConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*NdesConnector))
        }
        m.SetNdesConnectors(res)
        return nil
    }
    res["notificationMessageTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotificationMessageTemplate() })
        if err != nil {
            return err
        }
        res := make([]NotificationMessageTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*NotificationMessageTemplate))
        }
        m.SetNotificationMessageTemplates(res)
        return nil
    }
    res["remoteActionAudits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteActionAudit() })
        if err != nil {
            return err
        }
        res := make([]RemoteActionAudit, len(val))
        for i, v := range val {
            res[i] = *(v.(*RemoteActionAudit))
        }
        m.SetRemoteActionAudits(res)
        return nil
    }
    res["remoteAssistancePartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteAssistancePartner() })
        if err != nil {
            return err
        }
        res := make([]RemoteAssistancePartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*RemoteAssistancePartner))
        }
        m.SetRemoteAssistancePartners(res)
        return nil
    }
    res["remoteAssistanceSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteAssistanceSettings() })
        if err != nil {
            return err
        }
        m.SetRemoteAssistanceSettings(val.(*RemoteAssistanceSettings))
        return nil
    }
    res["reports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementReports() })
        if err != nil {
            return err
        }
        m.SetReports(val.(*DeviceManagementReports))
        return nil
    }
    res["resourceAccessProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementResourceAccessProfileBase() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementResourceAccessProfileBase, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementResourceAccessProfileBase))
        }
        m.SetResourceAccessProfiles(res)
        return nil
    }
    res["resourceOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceOperation() })
        if err != nil {
            return err
        }
        res := make([]ResourceOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceOperation))
        }
        m.SetResourceOperations(res)
        return nil
    }
    res["reusablePolicySettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementReusablePolicySetting() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementReusablePolicySetting, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementReusablePolicySetting))
        }
        m.SetReusablePolicySettings(res)
        return nil
    }
    res["reusableSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationSettingDefinition))
        }
        m.SetReusableSettings(res)
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceAndAppManagementRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceAndAppManagementRoleAssignment))
        }
        m.SetRoleAssignments(res)
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleDefinition() })
        if err != nil {
            return err
        }
        res := make([]RoleDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*RoleDefinition))
        }
        m.SetRoleDefinitions(res)
        return nil
    }
    res["roleScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTag() })
        if err != nil {
            return err
        }
        res := make([]RoleScopeTag, len(val))
        for i, v := range val {
            res[i] = *(v.(*RoleScopeTag))
        }
        m.SetRoleScopeTags(res)
        return nil
    }
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingDefinition() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementSettingDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementSettingDefinition))
        }
        m.SetSettingDefinitions(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*DeviceManagementSettings))
        return nil
    }
    res["softwareUpdateStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSoftwareUpdateStatusSummary() })
        if err != nil {
            return err
        }
        m.SetSoftwareUpdateStatusSummary(val.(*SoftwareUpdateStatusSummary))
        return nil
    }
    res["subscriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptions)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementSubscriptions)
        m.SetSubscriptions(&cast)
        return nil
    }
    res["subscriptionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptionState)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementSubscriptionState)
        m.SetSubscriptionState(&cast)
        return nil
    }
    res["telecomExpenseManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTelecomExpenseManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]TelecomExpenseManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*TelecomExpenseManagementPartner))
        }
        m.SetTelecomExpenseManagementPartners(res)
        return nil
    }
    res["templates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTemplate() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTemplate))
        }
        m.SetTemplates(res)
        return nil
    }
    res["templateSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingTemplate() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationSettingTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationSettingTemplate))
        }
        m.SetTemplateSettings(res)
        return nil
    }
    res["termsAndConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditions() })
        if err != nil {
            return err
        }
        res := make([]TermsAndConditions, len(val))
        for i, v := range val {
            res[i] = *(v.(*TermsAndConditions))
        }
        m.SetTermsAndConditions(res)
        return nil
    }
    res["troubleshootingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingEvent() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTroubleshootingEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTroubleshootingEvent))
        }
        m.SetTroubleshootingEvents(res)
        return nil
    }
    res["unlicensedAdminstratorsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUnlicensedAdminstratorsEnabled(val)
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthApplicationPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthApplicationPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthApplicationPerformance))
        }
        m.SetUserExperienceAnalyticsAppHealthApplicationPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion))
        }
        m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails))
        }
        m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId))
        }
        m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion))
        }
        m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthDeviceModelPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthDeviceModelPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthDeviceModelPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthDeviceModelPerformance))
        }
        m.SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthDevicePerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthDevicePerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthDevicePerformance))
        }
        m.SetUserExperienceAnalyticsAppHealthDevicePerformance(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformanceDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthDevicePerformanceDetails))
        }
        m.SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthOSVersionPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAppHealthOSVersionPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsAppHealthOSVersionPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsAppHealthOSVersionPerformance))
        }
        m.SetUserExperienceAnalyticsAppHealthOSVersionPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsAppHealthOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsAppHealthOverview(val.(*UserExperienceAnalyticsCategory))
        return nil
    }
    res["userExperienceAnalyticsBaselines"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBaseline() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBaseline, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBaseline))
        }
        m.SetUserExperienceAnalyticsBaselines(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthAppImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthAppImpact() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthAppImpact, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthAppImpact))
        }
        m.SetUserExperienceAnalyticsBatteryHealthAppImpact(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthCapacityDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthCapacityDetails() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsBatteryHealthCapacityDetails(val.(*UserExperienceAnalyticsBatteryHealthCapacityDetails))
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceAppImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthDeviceAppImpact, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact))
        }
        m.SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDevicePerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthDevicePerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthDevicePerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthDevicePerformance))
        }
        m.SetUserExperienceAnalyticsBatteryHealthDevicePerformance(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory))
        }
        m.SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthModelPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthModelPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthModelPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthModelPerformance))
        }
        m.SetUserExperienceAnalyticsBatteryHealthModelPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthOsPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthOsPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsBatteryHealthOsPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsBatteryHealthOsPerformance))
        }
        m.SetUserExperienceAnalyticsBatteryHealthOsPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthRuntimeDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsBatteryHealthRuntimeDetails() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(val.(*UserExperienceAnalyticsBatteryHealthRuntimeDetails))
        return nil
    }
    res["userExperienceAnalyticsCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCategory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsCategory))
        }
        m.SetUserExperienceAnalyticsCategories(res)
        return nil
    }
    res["userExperienceAnalyticsDeviceMetricHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetricHistory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetricHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetricHistory))
        }
        m.SetUserExperienceAnalyticsDeviceMetricHistory(res)
        return nil
    }
    res["userExperienceAnalyticsDevicePerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDevicePerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDevicePerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDevicePerformance))
        }
        m.SetUserExperienceAnalyticsDevicePerformance(res)
        return nil
    }
    res["userExperienceAnalyticsDeviceScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDeviceScores() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDeviceScores, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDeviceScores))
        }
        m.SetUserExperienceAnalyticsDeviceScores(res)
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDeviceStartupHistory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDeviceStartupHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDeviceStartupHistory))
        }
        m.SetUserExperienceAnalyticsDeviceStartupHistory(res)
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcesses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDeviceStartupProcess() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDeviceStartupProcess, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDeviceStartupProcess))
        }
        m.SetUserExperienceAnalyticsDeviceStartupProcesses(res)
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcessPerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDeviceStartupProcessPerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDeviceStartupProcessPerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDeviceStartupProcessPerformance))
        }
        m.SetUserExperienceAnalyticsDeviceStartupProcessPerformance(res)
        return nil
    }
    res["userExperienceAnalyticsDevicesWithoutCloudIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsDeviceWithoutCloudIdentity() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsDeviceWithoutCloudIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsDeviceWithoutCloudIdentity))
        }
        m.SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(res)
        return nil
    }
    res["userExperienceAnalyticsImpactingProcess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsImpactingProcess() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsImpactingProcess, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsImpactingProcess))
        }
        m.SetUserExperienceAnalyticsImpactingProcess(res)
        return nil
    }
    res["userExperienceAnalyticsMetricHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsMetricHistory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsMetricHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsMetricHistory))
        }
        m.SetUserExperienceAnalyticsMetricHistory(res)
        return nil
    }
    res["userExperienceAnalyticsModelScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsModelScores() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsModelScores, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsModelScores))
        }
        m.SetUserExperienceAnalyticsModelScores(res)
        return nil
    }
    res["userExperienceAnalyticsNotAutopilotReadyDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsNotAutopilotReadyDevice() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsNotAutopilotReadyDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsNotAutopilotReadyDevice))
        }
        m.SetUserExperienceAnalyticsNotAutopilotReadyDevice(res)
        return nil
    }
    res["userExperienceAnalyticsOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsOverview() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsOverview(val.(*UserExperienceAnalyticsOverview))
        return nil
    }
    res["userExperienceAnalyticsRegressionSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsRegressionSummary() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsRegressionSummary(val.(*UserExperienceAnalyticsRegressionSummary))
        return nil
    }
    res["userExperienceAnalyticsRemoteConnection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsRemoteConnection() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsRemoteConnection, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsRemoteConnection))
        }
        m.SetUserExperienceAnalyticsRemoteConnection(res)
        return nil
    }
    res["userExperienceAnalyticsResourcePerformance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsResourcePerformance() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsResourcePerformance, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsResourcePerformance))
        }
        m.SetUserExperienceAnalyticsResourcePerformance(res)
        return nil
    }
    res["userExperienceAnalyticsScoreHistory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsScoreHistory() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsScoreHistory, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsScoreHistory))
        }
        m.SetUserExperienceAnalyticsScoreHistory(res)
        return nil
    }
    res["userExperienceAnalyticsSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsSettings() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsSettings(val.(*UserExperienceAnalyticsSettings))
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric() })
        if err != nil {
            return err
        }
        m.SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric))
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereMetrics"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWorkFromAnywhereMetric() })
        if err != nil {
            return err
        }
        res := make([]UserExperienceAnalyticsWorkFromAnywhereMetric, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserExperienceAnalyticsWorkFromAnywhereMetric))
        }
        m.SetUserExperienceAnalyticsWorkFromAnywhereMetrics(res)
        return nil
    }
    res["userPfxCertificates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserPFXCertificate() })
        if err != nil {
            return err
        }
        res := make([]UserPFXCertificate, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserPFXCertificate))
        }
        m.SetUserPfxCertificates(res)
        return nil
    }
    res["virtualEndpoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVirtualEndpoint() })
        if err != nil {
            return err
        }
        m.SetVirtualEndpoint(val.(*VirtualEndpoint))
        return nil
    }
    res["windowsAutopilotDeploymentProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeploymentProfile() })
        if err != nil {
            return err
        }
        res := make([]WindowsAutopilotDeploymentProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsAutopilotDeploymentProfile))
        }
        m.SetWindowsAutopilotDeploymentProfiles(res)
        return nil
    }
    res["windowsAutopilotDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]WindowsAutopilotDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsAutopilotDeviceIdentity))
        }
        m.SetWindowsAutopilotDeviceIdentities(res)
        return nil
    }
    res["windowsAutopilotSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotSettings() })
        if err != nil {
            return err
        }
        m.SetWindowsAutopilotSettings(val.(*WindowsAutopilotSettings))
        return nil
    }
    res["windowsDriverUpdateProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDriverUpdateProfile() })
        if err != nil {
            return err
        }
        res := make([]WindowsDriverUpdateProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsDriverUpdateProfile))
        }
        m.SetWindowsDriverUpdateProfiles(res)
        return nil
    }
    res["windowsFeatureUpdateProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsFeatureUpdateProfile() })
        if err != nil {
            return err
        }
        res := make([]WindowsFeatureUpdateProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsFeatureUpdateProfile))
        }
        m.SetWindowsFeatureUpdateProfiles(res)
        return nil
    }
    res["windowsInformationProtectionAppLearningSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionAppLearningSummary() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionAppLearningSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionAppLearningSummary))
        }
        m.SetWindowsInformationProtectionAppLearningSummaries(res)
        return nil
    }
    res["windowsInformationProtectionNetworkLearningSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionNetworkLearningSummary() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionNetworkLearningSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionNetworkLearningSummary))
        }
        m.SetWindowsInformationProtectionNetworkLearningSummaries(res)
        return nil
    }
    res["windowsMalwareInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsMalwareInformation() })
        if err != nil {
            return err
        }
        res := make([]WindowsMalwareInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsMalwareInformation))
        }
        m.SetWindowsMalwareInformation(res)
        return nil
    }
    res["windowsMalwareOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsMalwareOverview() })
        if err != nil {
            return err
        }
        m.SetWindowsMalwareOverview(val.(*WindowsMalwareOverview))
        return nil
    }
    res["windowsQualityUpdateProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsQualityUpdateProfile() })
        if err != nil {
            return err
        }
        res := make([]WindowsQualityUpdateProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsQualityUpdateProfile))
        }
        m.SetWindowsQualityUpdateProfiles(res)
        return nil
    }
    res["windowsUpdateCatalogItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsUpdateCatalogItem() })
        if err != nil {
            return err
        }
        res := make([]WindowsUpdateCatalogItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsUpdateCatalogItem))
        }
        m.SetWindowsUpdateCatalogItems(res)
        return nil
    }
    return res
}
func (m *DeviceManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidDeviceOwnerEnrollmentProfiles()))
        for i, v := range m.GetAndroidDeviceOwnerEnrollmentProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("androidDeviceOwnerEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidForWorkAppConfigurationSchemas()))
        for i, v := range m.GetAndroidForWorkAppConfigurationSchemas() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("androidForWorkAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidForWorkEnrollmentProfiles()))
        for i, v := range m.GetAndroidForWorkEnrollmentProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidManagedStoreAppConfigurationSchemas()))
        for i, v := range m.GetAndroidManagedStoreAppConfigurationSchemas() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppleUserInitiatedEnrollmentProfiles()))
        for i, v := range m.GetAppleUserInitiatedEnrollmentProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appleUserInitiatedEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentFilters()))
        for i, v := range m.GetAssignmentFilters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentFilters", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAutopilotEvents()))
        for i, v := range m.GetAutopilotEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("autopilotEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCartToClassAssociations()))
        for i, v := range m.GetCartToClassAssociations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cartToClassAssociations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertificateConnectorDetails()))
        for i, v := range m.GetCertificateConnectorDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("certificateConnectorDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChromeOSOnboardingSettings()))
        for i, v := range m.GetChromeOSOnboardingSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPCConnectivityIssues()))
        for i, v := range m.GetCloudPCConnectivityIssues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCConnectivityIssues", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComanagedDevices()))
        for i, v := range m.GetComanagedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("comanagedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComanagementEligibleDevices()))
        for i, v := range m.GetComanagementEligibleDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("comanagementEligibleDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceCategories()))
        for i, v := range m.GetComplianceCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("complianceCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceManagementPartners()))
        for i, v := range m.GetComplianceManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompliancePolicies()))
        for i, v := range m.GetCompliancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("compliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceSettings()))
        for i, v := range m.GetComplianceSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigManagerCollections()))
        for i, v := range m.GetConfigManagerCollections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configManagerCollections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationCategories()))
        for i, v := range m.GetConfigurationCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configurationCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationPolicies()))
        for i, v := range m.GetConfigurationPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationPolicyTemplates()))
        for i, v := range m.GetConfigurationPolicyTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicyTemplates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationSettings()))
        for i, v := range m.GetConfigurationSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configurationSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDataSharingConsents()))
        for i, v := range m.GetDataSharingConsents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dataSharingConsents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDepOnboardingSettings()))
        for i, v := range m.GetDepOnboardingSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("depOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDerivedCredentials()))
        for i, v := range m.GetDerivedCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("derivedCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCategories()))
        for i, v := range m.GetDeviceCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicies()))
        for i, v := range m.GetDeviceCompliancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deviceComplianceReportSummarizationDateTime", m.GetDeviceComplianceReportSummarizationDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceComplianceScripts()))
        for i, v := range m.GetDeviceComplianceScripts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceComplianceScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurationConflictSummary()))
        for i, v := range m.GetDeviceConfigurationConflictSummary() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurationRestrictedAppsViolations()))
        for i, v := range m.GetDeviceConfigurationRestrictedAppsViolations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationRestrictedAppsViolations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurations()))
        for i, v := range m.GetDeviceConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurationsAllManagedDeviceCertificateStates()))
        for i, v := range m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCustomAttributeShellScripts()))
        for i, v := range m.GetDeviceCustomAttributeShellScripts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCustomAttributeShellScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceEnrollmentConfigurations()))
        for i, v := range m.GetDeviceEnrollmentConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceHealthScripts()))
        for i, v := range m.GetDeviceHealthScripts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceHealthScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceManagementPartners()))
        for i, v := range m.GetDeviceManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceManagementScripts()))
        for i, v := range m.GetDeviceManagementScripts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceShellScripts()))
        for i, v := range m.GetDeviceShellScripts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceShellScripts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDomainJoinConnectors()))
        for i, v := range m.GetDomainJoinConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmbeddedSIMActivationCodePools()))
        for i, v := range m.GetEmbeddedSIMActivationCodePools() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("embeddedSIMActivationCodePools", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExchangeConnectors()))
        for i, v := range m.GetExchangeConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExchangeOnPremisesPolicies()))
        for i, v := range m.GetExchangeOnPremisesPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyCategories()))
        for i, v := range m.GetGroupPolicyCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyConfigurations()))
        for i, v := range m.GetGroupPolicyConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyDefinitionFiles()))
        for i, v := range m.GetGroupPolicyDefinitionFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyDefinitions()))
        for i, v := range m.GetGroupPolicyDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyMigrationReports()))
        for i, v := range m.GetGroupPolicyMigrationReports() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyMigrationReports", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyObjectFiles()))
        for i, v := range m.GetGroupPolicyObjectFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyObjectFiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicyUploadedDefinitionFiles()))
        for i, v := range m.GetGroupPolicyUploadedDefinitionFiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedDeviceIdentities()))
        for i, v := range m.GetImportedDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIntents()))
        for i, v := range m.GetIntents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIntuneBrandingProfiles()))
        for i, v := range m.GetIntuneBrandingProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("intuneBrandingProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosUpdateStatuses()))
        for i, v := range m.GetIosUpdateStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportAggregationDateTime", m.GetLastReportAggregationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("legacyPcManangementEnabled", m.GetLegacyPcManangementEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMacOSSoftwareUpdateAccountSummaries()))
        for i, v := range m.GetMacOSSoftwareUpdateAccountSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDeviceEncryptionStates()))
        for i, v := range m.GetManagedDeviceEncryptionStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementConditions()))
        for i, v := range m.GetManagementConditions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementConditions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementConditionStatements()))
        for i, v := range m.GetManagementConditionStatements() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementConditionStatements", cast)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftTunnelConfigurations()))
        for i, v := range m.GetMicrosoftTunnelConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftTunnelHealthThresholds()))
        for i, v := range m.GetMicrosoftTunnelHealthThresholds() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelHealthThresholds", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftTunnelServerLogCollectionResponses()))
        for i, v := range m.GetMicrosoftTunnelServerLogCollectionResponses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServerLogCollectionResponses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftTunnelSites()))
        for i, v := range m.GetMicrosoftTunnelSites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelSites", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppTroubleshootingEvents()))
        for i, v := range m.GetMobileAppTroubleshootingEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileThreatDefenseConnectors()))
        for i, v := range m.GetMobileThreatDefenseConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNdesConnectors()))
        for i, v := range m.GetNdesConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ndesConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotificationMessageTemplates()))
        for i, v := range m.GetNotificationMessageTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRemoteActionAudits()))
        for i, v := range m.GetRemoteActionAudits() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("remoteActionAudits", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRemoteAssistancePartners()))
        for i, v := range m.GetRemoteAssistancePartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceAccessProfiles()))
        for i, v := range m.GetResourceAccessProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resourceAccessProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceOperations()))
        for i, v := range m.GetResourceOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReusablePolicySettings()))
        for i, v := range m.GetReusablePolicySettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reusablePolicySettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReusableSettings()))
        for i, v := range m.GetReusableSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reusableSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleScopeTags()))
        for i, v := range m.GetRoleScopeTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleScopeTags", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
        cast := m.GetSubscriptions().String()
        err = writer.WriteStringValue("subscriptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionState() != nil {
        cast := m.GetSubscriptionState().String()
        err = writer.WriteStringValue("subscriptionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTelecomExpenseManagementPartners()))
        for i, v := range m.GetTelecomExpenseManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemplateSettings()))
        for i, v := range m.GetTemplateSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("templateSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTermsAndConditions()))
        for i, v := range m.GetTermsAndConditions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTroubleshootingEvents()))
        for i, v := range m.GetTroubleshootingEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unlicensedAdminstratorsEnabled", m.GetUnlicensedAdminstratorsEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDeviceModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformanceDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBaselines()))
        for i, v := range m.GetUserExperienceAnalyticsBaselines() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBaselines", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthAppImpact() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceAppImpact", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthOsPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsCategories()))
        for i, v := range m.GetUserExperienceAnalyticsCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDeviceMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceMetricHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDevicePerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDeviceScores()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScores", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcesses()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcesses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcesses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcessPerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()))
        for i, v := range m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicesWithoutCloudIdentity", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsImpactingProcess()))
        for i, v := range m.GetUserExperienceAnalyticsImpactingProcess() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsImpactingProcess", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsMetricHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsModelScores()))
        for i, v := range m.GetUserExperienceAnalyticsModelScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsModelScores", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsNotAutopilotReadyDevice()))
        for i, v := range m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsRemoteConnection()))
        for i, v := range m.GetUserExperienceAnalyticsRemoteConnection() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsRemoteConnection", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsResourcePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsResourcePerformance() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsResourcePerformance", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsScoreHistory()))
        for i, v := range m.GetUserExperienceAnalyticsScoreHistory() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereMetrics", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserPfxCertificates()))
        for i, v := range m.GetUserPfxCertificates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsAutopilotDeploymentProfiles()))
        for i, v := range m.GetWindowsAutopilotDeploymentProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeploymentProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetWindowsAutopilotDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsDriverUpdateProfiles()))
        for i, v := range m.GetWindowsDriverUpdateProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsDriverUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsFeatureUpdateProfiles()))
        for i, v := range m.GetWindowsFeatureUpdateProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsFeatureUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionAppLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionAppLearningSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionNetworkLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionNetworkLearningSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsMalwareInformation()))
        for i, v := range m.GetWindowsMalwareInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsQualityUpdateProfiles()))
        for i, v := range m.GetWindowsQualityUpdateProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsQualityUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsUpdateCatalogItems()))
        for i, v := range m.GetWindowsUpdateCatalogItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsUpdateCatalogItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
// Parameters:
//  - value : Value to set for the accountMoveCompletionDateTime property.
func (m *DeviceManagement) SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.accountMoveCompletionDateTime = value
}
// Sets the adminConsent property value. Admin consent information.
// Parameters:
//  - value : Value to set for the adminConsent property.
func (m *DeviceManagement) SetAdminConsent(value *AdminConsent)() {
    m.adminConsent = value
}
// Sets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
// Parameters:
//  - value : Value to set for the advancedThreatProtectionOnboardingStateSummary property.
func (m *DeviceManagement) SetAdvancedThreatProtectionOnboardingStateSummary(value *AdvancedThreatProtectionOnboardingStateSummary)() {
    m.advancedThreatProtectionOnboardingStateSummary = value
}
// Sets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
// Parameters:
//  - value : Value to set for the androidDeviceOwnerEnrollmentProfiles property.
func (m *DeviceManagement) SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfile)() {
    m.androidDeviceOwnerEnrollmentProfiles = value
}
// Sets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
// Parameters:
//  - value : Value to set for the androidForWorkAppConfigurationSchemas property.
func (m *DeviceManagement) SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchema)() {
    m.androidForWorkAppConfigurationSchemas = value
}
// Sets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
// Parameters:
//  - value : Value to set for the androidForWorkEnrollmentProfiles property.
func (m *DeviceManagement) SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfile)() {
    m.androidForWorkEnrollmentProfiles = value
}
// Sets the androidForWorkSettings property value. The singleton Android for Work settings entity.
// Parameters:
//  - value : Value to set for the androidForWorkSettings property.
func (m *DeviceManagement) SetAndroidForWorkSettings(value *AndroidForWorkSettings)() {
    m.androidForWorkSettings = value
}
// Sets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
// Parameters:
//  - value : Value to set for the androidManagedStoreAccountEnterpriseSettings property.
func (m *DeviceManagement) SetAndroidManagedStoreAccountEnterpriseSettings(value *AndroidManagedStoreAccountEnterpriseSettings)() {
    m.androidManagedStoreAccountEnterpriseSettings = value
}
// Sets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
// Parameters:
//  - value : Value to set for the androidManagedStoreAppConfigurationSchemas property.
func (m *DeviceManagement) SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchema)() {
    m.androidManagedStoreAppConfigurationSchemas = value
}
// Sets the applePushNotificationCertificate property value. Apple push notification certificate.
// Parameters:
//  - value : Value to set for the applePushNotificationCertificate property.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value *ApplePushNotificationCertificate)() {
    m.applePushNotificationCertificate = value
}
// Sets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
// Parameters:
//  - value : Value to set for the appleUserInitiatedEnrollmentProfiles property.
func (m *DeviceManagement) SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfile)() {
    m.appleUserInitiatedEnrollmentProfiles = value
}
// Sets the assignmentFilters property value. The list of assignment filters
// Parameters:
//  - value : Value to set for the assignmentFilters property.
func (m *DeviceManagement) SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilter)() {
    m.assignmentFilters = value
}
// Sets the auditEvents property value. The Audit Events
// Parameters:
//  - value : Value to set for the auditEvents property.
func (m *DeviceManagement) SetAuditEvents(value []AuditEvent)() {
    m.auditEvents = value
}
// Sets the autopilotEvents property value. The list of autopilot events for the tenant.
// Parameters:
//  - value : Value to set for the autopilotEvents property.
func (m *DeviceManagement) SetAutopilotEvents(value []DeviceManagementAutopilotEvent)() {
    m.autopilotEvents = value
}
// Sets the cartToClassAssociations property value. The Cart To Class Associations.
// Parameters:
//  - value : Value to set for the cartToClassAssociations property.
func (m *DeviceManagement) SetCartToClassAssociations(value []CartToClassAssociation)() {
    m.cartToClassAssociations = value
}
// Sets the categories property value. The available categories
// Parameters:
//  - value : Value to set for the categories property.
func (m *DeviceManagement) SetCategories(value []DeviceManagementSettingCategory)() {
    m.categories = value
}
// Sets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
// Parameters:
//  - value : Value to set for the certificateConnectorDetails property.
func (m *DeviceManagement) SetCertificateConnectorDetails(value []CertificateConnectorDetails)() {
    m.certificateConnectorDetails = value
}
// Sets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
// Parameters:
//  - value : Value to set for the chromeOSOnboardingSettings property.
func (m *DeviceManagement) SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettings)() {
    m.chromeOSOnboardingSettings = value
}
// Sets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
// Parameters:
//  - value : Value to set for the cloudPCConnectivityIssues property.
func (m *DeviceManagement) SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssue)() {
    m.cloudPCConnectivityIssues = value
}
// Sets the comanagedDevices property value. The list of co-managed devices report
// Parameters:
//  - value : Value to set for the comanagedDevices property.
func (m *DeviceManagement) SetComanagedDevices(value []ManagedDevice)() {
    m.comanagedDevices = value
}
// Sets the comanagementEligibleDevices property value. The list of co-management eligible devices report
// Parameters:
//  - value : Value to set for the comanagementEligibleDevices property.
func (m *DeviceManagement) SetComanagementEligibleDevices(value []ComanagementEligibleDevice)() {
    m.comanagementEligibleDevices = value
}
// Sets the complianceCategories property value. List of all compliance categories
// Parameters:
//  - value : Value to set for the complianceCategories property.
func (m *DeviceManagement) SetComplianceCategories(value []DeviceManagementConfigurationCategory)() {
    m.complianceCategories = value
}
// Sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
// Parameters:
//  - value : Value to set for the complianceManagementPartners property.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartner)() {
    m.complianceManagementPartners = value
}
// Sets the compliancePolicies property value. List of all compliance policies
// Parameters:
//  - value : Value to set for the compliancePolicies property.
func (m *DeviceManagement) SetCompliancePolicies(value []DeviceManagementCompliancePolicy)() {
    m.compliancePolicies = value
}
// Sets the complianceSettings property value. List of all ComplianceSettings
// Parameters:
//  - value : Value to set for the complianceSettings property.
func (m *DeviceManagement) SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinition)() {
    m.complianceSettings = value
}
// Sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// Parameters:
//  - value : Value to set for the conditionalAccessSettings property.
func (m *DeviceManagement) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    m.conditionalAccessSettings = value
}
// Sets the configManagerCollections property value. A list of ConfigManagerCollection
// Parameters:
//  - value : Value to set for the configManagerCollections property.
func (m *DeviceManagement) SetConfigManagerCollections(value []ConfigManagerCollection)() {
    m.configManagerCollections = value
}
// Sets the configurationCategories property value. List of all Configuration Categories
// Parameters:
//  - value : Value to set for the configurationCategories property.
func (m *DeviceManagement) SetConfigurationCategories(value []DeviceManagementConfigurationCategory)() {
    m.configurationCategories = value
}
// Sets the configurationPolicies property value. List of all Configuration policies
// Parameters:
//  - value : Value to set for the configurationPolicies property.
func (m *DeviceManagement) SetConfigurationPolicies(value []DeviceManagementConfigurationPolicy)() {
    m.configurationPolicies = value
}
// Sets the configurationPolicyTemplates property value. List of all templates
// Parameters:
//  - value : Value to set for the configurationPolicyTemplates property.
func (m *DeviceManagement) SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplate)() {
    m.configurationPolicyTemplates = value
}
// Sets the configurationSettings property value. List of all ConfigurationSettings
// Parameters:
//  - value : Value to set for the configurationSettings property.
func (m *DeviceManagement) SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinition)() {
    m.configurationSettings = value
}
// Sets the dataSharingConsents property value. Data sharing consents.
// Parameters:
//  - value : Value to set for the dataSharingConsents property.
func (m *DeviceManagement) SetDataSharingConsents(value []DataSharingConsent)() {
    m.dataSharingConsents = value
}
// Sets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
// Parameters:
//  - value : Value to set for the depOnboardingSettings property.
func (m *DeviceManagement) SetDepOnboardingSettings(value []DepOnboardingSetting)() {
    m.depOnboardingSettings = value
}
// Sets the derivedCredentials property value. Collection of Derived credential settings associated with account.
// Parameters:
//  - value : Value to set for the derivedCredentials property.
func (m *DeviceManagement) SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettings)() {
    m.derivedCredentials = value
}
// Sets the detectedApps property value. The list of detected apps associated with a device.
// Parameters:
//  - value : Value to set for the detectedApps property.
func (m *DeviceManagement) SetDetectedApps(value []DetectedApp)() {
    m.detectedApps = value
}
// Sets the deviceCategories property value. The list of device categories with the tenant.
// Parameters:
//  - value : Value to set for the deviceCategories property.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategory)() {
    m.deviceCategories = value
}
// Sets the deviceCompliancePolicies property value. The device compliance policies.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicies property.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicy)() {
    m.deviceCompliancePolicies = value
}
// Sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicyDeviceStateSummary property.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value *DeviceCompliancePolicyDeviceStateSummary)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
// Sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicySettingStateSummaries property.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// Sets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
// Parameters:
//  - value : Value to set for the deviceComplianceReportSummarizationDateTime property.
func (m *DeviceManagement) SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceComplianceReportSummarizationDateTime = value
}
// Sets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
// Parameters:
//  - value : Value to set for the deviceComplianceScripts property.
func (m *DeviceManagement) SetDeviceComplianceScripts(value []DeviceComplianceScript)() {
    m.deviceComplianceScripts = value
}
// Sets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
// Parameters:
//  - value : Value to set for the deviceConfigurationConflictSummary property.
func (m *DeviceManagement) SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummary)() {
    m.deviceConfigurationConflictSummary = value
}
// Sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
// Parameters:
//  - value : Value to set for the deviceConfigurationDeviceStateSummaries property.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value *DeviceConfigurationDeviceStateSummary)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
// Sets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
// Parameters:
//  - value : Value to set for the deviceConfigurationRestrictedAppsViolations property.
func (m *DeviceManagement) SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolation)() {
    m.deviceConfigurationRestrictedAppsViolations = value
}
// Sets the deviceConfigurations property value. The device configurations.
// Parameters:
//  - value : Value to set for the deviceConfigurations property.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfiguration)() {
    m.deviceConfigurations = value
}
// Sets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
// Parameters:
//  - value : Value to set for the deviceConfigurationsAllManagedDeviceCertificateStates property.
func (m *DeviceManagement) SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateState)() {
    m.deviceConfigurationsAllManagedDeviceCertificateStates = value
}
// Sets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
// Parameters:
//  - value : Value to set for the deviceConfigurationUserStateSummaries property.
func (m *DeviceManagement) SetDeviceConfigurationUserStateSummaries(value *DeviceConfigurationUserStateSummary)() {
    m.deviceConfigurationUserStateSummaries = value
}
// Sets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
// Parameters:
//  - value : Value to set for the deviceCustomAttributeShellScripts property.
func (m *DeviceManagement) SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScript)() {
    m.deviceCustomAttributeShellScripts = value
}
// Sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
// Parameters:
//  - value : Value to set for the deviceEnrollmentConfigurations property.
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfiguration)() {
    m.deviceEnrollmentConfigurations = value
}
// Sets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
// Parameters:
//  - value : Value to set for the deviceHealthScripts property.
func (m *DeviceManagement) SetDeviceHealthScripts(value []DeviceHealthScript)() {
    m.deviceHealthScripts = value
}
// Sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
// Parameters:
//  - value : Value to set for the deviceManagementPartners property.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartner)() {
    m.deviceManagementPartners = value
}
// Sets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
// Parameters:
//  - value : Value to set for the deviceManagementScripts property.
func (m *DeviceManagement) SetDeviceManagementScripts(value []DeviceManagementScript)() {
    m.deviceManagementScripts = value
}
// Sets the deviceProtectionOverview property value. Device protection overview.
// Parameters:
//  - value : Value to set for the deviceProtectionOverview property.
func (m *DeviceManagement) SetDeviceProtectionOverview(value *DeviceProtectionOverview)() {
    m.deviceProtectionOverview = value
}
// Sets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
// Parameters:
//  - value : Value to set for the deviceShellScripts property.
func (m *DeviceManagement) SetDeviceShellScripts(value []DeviceShellScript)() {
    m.deviceShellScripts = value
}
// Sets the domainJoinConnectors property value. A list of connector objects.
// Parameters:
//  - value : Value to set for the domainJoinConnectors property.
func (m *DeviceManagement) SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnector)() {
    m.domainJoinConnectors = value
}
// Sets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
// Parameters:
//  - value : Value to set for the embeddedSIMActivationCodePools property.
func (m *DeviceManagement) SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePool)() {
    m.embeddedSIMActivationCodePools = value
}
// Sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
// Parameters:
//  - value : Value to set for the exchangeConnectors property.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnector)() {
    m.exchangeConnectors = value
}
// Sets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
// Parameters:
//  - value : Value to set for the exchangeOnPremisesPolicies property.
func (m *DeviceManagement) SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicy)() {
    m.exchangeOnPremisesPolicies = value
}
// Sets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
// Parameters:
//  - value : Value to set for the exchangeOnPremisesPolicy property.
func (m *DeviceManagement) SetExchangeOnPremisesPolicy(value *DeviceManagementExchangeOnPremisesPolicy)() {
    m.exchangeOnPremisesPolicy = value
}
// Sets the groupPolicyCategories property value. The available group policy categories for this account.
// Parameters:
//  - value : Value to set for the groupPolicyCategories property.
func (m *DeviceManagement) SetGroupPolicyCategories(value []GroupPolicyCategory)() {
    m.groupPolicyCategories = value
}
// Sets the groupPolicyConfigurations property value. The group policy configurations created by this account.
// Parameters:
//  - value : Value to set for the groupPolicyConfigurations property.
func (m *DeviceManagement) SetGroupPolicyConfigurations(value []GroupPolicyConfiguration)() {
    m.groupPolicyConfigurations = value
}
// Sets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
// Parameters:
//  - value : Value to set for the groupPolicyDefinitionFiles property.
func (m *DeviceManagement) SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFile)() {
    m.groupPolicyDefinitionFiles = value
}
// Sets the groupPolicyDefinitions property value. The available group policy definitions for this account.
// Parameters:
//  - value : Value to set for the groupPolicyDefinitions property.
func (m *DeviceManagement) SetGroupPolicyDefinitions(value []GroupPolicyDefinition)() {
    m.groupPolicyDefinitions = value
}
// Sets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
// Parameters:
//  - value : Value to set for the groupPolicyMigrationReports property.
func (m *DeviceManagement) SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReport)() {
    m.groupPolicyMigrationReports = value
}
// Sets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
// Parameters:
//  - value : Value to set for the groupPolicyObjectFiles property.
func (m *DeviceManagement) SetGroupPolicyObjectFiles(value []GroupPolicyObjectFile)() {
    m.groupPolicyObjectFiles = value
}
// Sets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
// Parameters:
//  - value : Value to set for the groupPolicyUploadedDefinitionFiles property.
func (m *DeviceManagement) SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFile)() {
    m.groupPolicyUploadedDefinitionFiles = value
}
// Sets the importedDeviceIdentities property value. The imported device identities.
// Parameters:
//  - value : Value to set for the importedDeviceIdentities property.
func (m *DeviceManagement) SetImportedDeviceIdentities(value []ImportedDeviceIdentity)() {
    m.importedDeviceIdentities = value
}
// Sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
// Parameters:
//  - value : Value to set for the importedWindowsAutopilotDeviceIdentities property.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentity)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
// Sets the intents property value. The device management intents
// Parameters:
//  - value : Value to set for the intents property.
func (m *DeviceManagement) SetIntents(value []DeviceManagementIntent)() {
    m.intents = value
}
// Sets the intuneAccountId property value. Intune Account Id for given tenant
// Parameters:
//  - value : Value to set for the intuneAccountId property.
func (m *DeviceManagement) SetIntuneAccountId(value *string)() {
    m.intuneAccountId = value
}
// Sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
// Parameters:
//  - value : Value to set for the intuneBrand property.
func (m *DeviceManagement) SetIntuneBrand(value *IntuneBrand)() {
    m.intuneBrand = value
}
// Sets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
// Parameters:
//  - value : Value to set for the intuneBrandingProfiles property.
func (m *DeviceManagement) SetIntuneBrandingProfiles(value []IntuneBrandingProfile)() {
    m.intuneBrandingProfiles = value
}
// Sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
// Parameters:
//  - value : Value to set for the iosUpdateStatuses property.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatus)() {
    m.iosUpdateStatuses = value
}
// Sets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
// Parameters:
//  - value : Value to set for the lastReportAggregationDateTime property.
func (m *DeviceManagement) SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportAggregationDateTime = value
}
// Sets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
// Parameters:
//  - value : Value to set for the legacyPcManangementEnabled property.
func (m *DeviceManagement) SetLegacyPcManangementEnabled(value *bool)() {
    m.legacyPcManangementEnabled = value
}
// Sets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
// Parameters:
//  - value : Value to set for the macOSSoftwareUpdateAccountSummaries property.
func (m *DeviceManagement) SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummary)() {
    m.macOSSoftwareUpdateAccountSummaries = value
}
// Sets the managedDeviceCleanupSettings property value. Device cleanup rule
// Parameters:
//  - value : Value to set for the managedDeviceCleanupSettings property.
func (m *DeviceManagement) SetManagedDeviceCleanupSettings(value *ManagedDeviceCleanupSettings)() {
    m.managedDeviceCleanupSettings = value
}
// Sets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
// Parameters:
//  - value : Value to set for the managedDeviceEncryptionStates property.
func (m *DeviceManagement) SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionState)() {
    m.managedDeviceEncryptionStates = value
}
// Sets the managedDeviceOverview property value. Device overview
// Parameters:
//  - value : Value to set for the managedDeviceOverview property.
func (m *DeviceManagement) SetManagedDeviceOverview(value *ManagedDeviceOverview)() {
    m.managedDeviceOverview = value
}
// Sets the managedDevices property value. The list of managed devices.
// Parameters:
//  - value : Value to set for the managedDevices property.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
// Sets the managementConditions property value. The management conditions associated with device management of the company.
// Parameters:
//  - value : Value to set for the managementConditions property.
func (m *DeviceManagement) SetManagementConditions(value []ManagementCondition)() {
    m.managementConditions = value
}
// Sets the managementConditionStatements property value. The management condition statements associated with device management of the company.
// Parameters:
//  - value : Value to set for the managementConditionStatements property.
func (m *DeviceManagement) SetManagementConditionStatements(value []ManagementConditionStatement)() {
    m.managementConditionStatements = value
}
// Sets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
// Parameters:
//  - value : Value to set for the maximumDepTokens property.
func (m *DeviceManagement) SetMaximumDepTokens(value *int32)() {
    m.maximumDepTokens = value
}
// Sets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
// Parameters:
//  - value : Value to set for the microsoftTunnelConfigurations property.
func (m *DeviceManagement) SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfiguration)() {
    m.microsoftTunnelConfigurations = value
}
// Sets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
// Parameters:
//  - value : Value to set for the microsoftTunnelHealthThresholds property.
func (m *DeviceManagement) SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThreshold)() {
    m.microsoftTunnelHealthThresholds = value
}
// Sets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
// Parameters:
//  - value : Value to set for the microsoftTunnelServerLogCollectionResponses property.
func (m *DeviceManagement) SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponse)() {
    m.microsoftTunnelServerLogCollectionResponses = value
}
// Sets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
// Parameters:
//  - value : Value to set for the microsoftTunnelSites property.
func (m *DeviceManagement) SetMicrosoftTunnelSites(value []MicrosoftTunnelSite)() {
    m.microsoftTunnelSites = value
}
// Sets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
// Parameters:
//  - value : Value to set for the mobileAppTroubleshootingEvents property.
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEvent)() {
    m.mobileAppTroubleshootingEvents = value
}
// Sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
// Parameters:
//  - value : Value to set for the mobileThreatDefenseConnectors property.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnector)() {
    m.mobileThreatDefenseConnectors = value
}
// Sets the ndesConnectors property value. The collection of Ndes connectors for this account.
// Parameters:
//  - value : Value to set for the ndesConnectors property.
func (m *DeviceManagement) SetNdesConnectors(value []NdesConnector)() {
    m.ndesConnectors = value
}
// Sets the notificationMessageTemplates property value. The Notification Message Templates.
// Parameters:
//  - value : Value to set for the notificationMessageTemplates property.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplate)() {
    m.notificationMessageTemplates = value
}
// Sets the remoteActionAudits property value. The list of device remote action audits with the tenant.
// Parameters:
//  - value : Value to set for the remoteActionAudits property.
func (m *DeviceManagement) SetRemoteActionAudits(value []RemoteActionAudit)() {
    m.remoteActionAudits = value
}
// Sets the remoteAssistancePartners property value. The remote assist partners.
// Parameters:
//  - value : Value to set for the remoteAssistancePartners property.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartner)() {
    m.remoteAssistancePartners = value
}
// Sets the remoteAssistanceSettings property value. The remote assistance settings singleton
// Parameters:
//  - value : Value to set for the remoteAssistanceSettings property.
func (m *DeviceManagement) SetRemoteAssistanceSettings(value *RemoteAssistanceSettings)() {
    m.remoteAssistanceSettings = value
}
// Sets the reports property value. Reports singleton
// Parameters:
//  - value : Value to set for the reports property.
func (m *DeviceManagement) SetReports(value *DeviceManagementReports)() {
    m.reports = value
}
// Sets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
// Parameters:
//  - value : Value to set for the resourceAccessProfiles property.
func (m *DeviceManagement) SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBase)() {
    m.resourceAccessProfiles = value
}
// Sets the resourceOperations property value. The Resource Operations.
// Parameters:
//  - value : Value to set for the resourceOperations property.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperation)() {
    m.resourceOperations = value
}
// Sets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
// Parameters:
//  - value : Value to set for the reusablePolicySettings property.
func (m *DeviceManagement) SetReusablePolicySettings(value []DeviceManagementReusablePolicySetting)() {
    m.reusablePolicySettings = value
}
// Sets the reusableSettings property value. List of all reusable settings
// Parameters:
//  - value : Value to set for the reusableSettings property.
func (m *DeviceManagement) SetReusableSettings(value []DeviceManagementConfigurationSettingDefinition)() {
    m.reusableSettings = value
}
// Sets the roleAssignments property value. The Role Assignments.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignment)() {
    m.roleAssignments = value
}
// Sets the roleDefinitions property value. The Role Definitions.
// Parameters:
//  - value : Value to set for the roleDefinitions property.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinition)() {
    m.roleDefinitions = value
}
// Sets the roleScopeTags property value. The Role Scope Tags.
// Parameters:
//  - value : Value to set for the roleScopeTags property.
func (m *DeviceManagement) SetRoleScopeTags(value []RoleScopeTag)() {
    m.roleScopeTags = value
}
// Sets the settingDefinitions property value. The device management intent setting definitions
// Parameters:
//  - value : Value to set for the settingDefinitions property.
func (m *DeviceManagement) SetSettingDefinitions(value []DeviceManagementSettingDefinition)() {
    m.settingDefinitions = value
}
// Sets the settings property value. Account level settings.
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagement) SetSettings(value *DeviceManagementSettings)() {
    m.settings = value
}
// Sets the softwareUpdateStatusSummary property value. The software update status summary.
// Parameters:
//  - value : Value to set for the softwareUpdateStatusSummary property.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value *SoftwareUpdateStatusSummary)() {
    m.softwareUpdateStatusSummary = value
}
// Sets the subscriptions property value. Tenant's Subscription. Possible values are: none, intune, office365, intunePremium, intune_EDU, intune_SMB.
// Parameters:
//  - value : Value to set for the subscriptions property.
func (m *DeviceManagement) SetSubscriptions(value *DeviceManagementSubscriptions)() {
    m.subscriptions = value
}
// Sets the subscriptionState property value. Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
// Parameters:
//  - value : Value to set for the subscriptionState property.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
// Sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
// Parameters:
//  - value : Value to set for the telecomExpenseManagementPartners property.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartner)() {
    m.telecomExpenseManagementPartners = value
}
// Sets the templates property value. The available templates
// Parameters:
//  - value : Value to set for the templates property.
func (m *DeviceManagement) SetTemplates(value []DeviceManagementTemplate)() {
    m.templates = value
}
// Sets the templateSettings property value. List of all TemplateSettings
// Parameters:
//  - value : Value to set for the templateSettings property.
func (m *DeviceManagement) SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplate)() {
    m.templateSettings = value
}
// Sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
// Parameters:
//  - value : Value to set for the termsAndConditions property.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditions)() {
    m.termsAndConditions = value
}
// Sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
// Parameters:
//  - value : Value to set for the troubleshootingEvents property.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    m.troubleshootingEvents = value
}
// Sets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
// Parameters:
//  - value : Value to set for the unlicensedAdminstratorsEnabled property.
func (m *DeviceManagement) SetUnlicensedAdminstratorsEnabled(value *bool)() {
    m.unlicensedAdminstratorsEnabled = value
}
// Sets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthApplicationPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformance)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformance = value
}
// Sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion = value
}
// Sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = value
}
// Sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = value
}
// Sets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = value
}
// Sets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthDeviceModelPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformance)() {
    m.userExperienceAnalyticsAppHealthDeviceModelPerformance = value
}
// Sets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthDevicePerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformance)() {
    m.userExperienceAnalyticsAppHealthDevicePerformance = value
}
// Sets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthDevicePerformanceDetails property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetails)() {
    m.userExperienceAnalyticsAppHealthDevicePerformanceDetails = value
}
// Sets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthOSVersionPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformance)() {
    m.userExperienceAnalyticsAppHealthOSVersionPerformance = value
}
// Sets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsAppHealthOverview property.
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value *UserExperienceAnalyticsCategory)() {
    m.userExperienceAnalyticsAppHealthOverview = value
}
// Sets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBaselines property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaseline)() {
    m.userExperienceAnalyticsBaselines = value
}
// Sets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthAppImpact property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpact)() {
    m.userExperienceAnalyticsBatteryHealthAppImpact = value
}
// Sets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthCapacityDetails property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value *UserExperienceAnalyticsBatteryHealthCapacityDetails)() {
    m.userExperienceAnalyticsBatteryHealthCapacityDetails = value
}
// Sets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthDeviceAppImpact property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpact)() {
    m.userExperienceAnalyticsBatteryHealthDeviceAppImpact = value
}
// Sets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthDevicePerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformance)() {
    m.userExperienceAnalyticsBatteryHealthDevicePerformance = value
}
// Sets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory)() {
    m.userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory = value
}
// Sets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthModelPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformance)() {
    m.userExperienceAnalyticsBatteryHealthModelPerformance = value
}
// Sets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthOsPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformance)() {
    m.userExperienceAnalyticsBatteryHealthOsPerformance = value
}
// Sets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsBatteryHealthRuntimeDetails property.
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value *UserExperienceAnalyticsBatteryHealthRuntimeDetails)() {
    m.userExperienceAnalyticsBatteryHealthRuntimeDetails = value
}
// Sets the userExperienceAnalyticsCategories property value. User experience analytics categories
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsCategories property.
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategory)() {
    m.userExperienceAnalyticsCategories = value
}
// Sets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDeviceMetricHistory property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistory)() {
    m.userExperienceAnalyticsDeviceMetricHistory = value
}
// Sets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDevicePerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformance)() {
    m.userExperienceAnalyticsDevicePerformance = value
}
// Sets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDeviceScores property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScores)() {
    m.userExperienceAnalyticsDeviceScores = value
}
// Sets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDeviceStartupHistory property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistory)() {
    m.userExperienceAnalyticsDeviceStartupHistory = value
}
// Sets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDeviceStartupProcesses property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcess)() {
    m.userExperienceAnalyticsDeviceStartupProcesses = value
}
// Sets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDeviceStartupProcessPerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformance)() {
    m.userExperienceAnalyticsDeviceStartupProcessPerformance = value
}
// Sets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsDevicesWithoutCloudIdentity property.
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentity)() {
    m.userExperienceAnalyticsDevicesWithoutCloudIdentity = value
}
// Sets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsImpactingProcess property.
func (m *DeviceManagement) SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcess)() {
    m.userExperienceAnalyticsImpactingProcess = value
}
// Sets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsMetricHistory property.
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistory)() {
    m.userExperienceAnalyticsMetricHistory = value
}
// Sets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsModelScores property.
func (m *DeviceManagement) SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScores)() {
    m.userExperienceAnalyticsModelScores = value
}
// Sets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsNotAutopilotReadyDevice property.
func (m *DeviceManagement) SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDevice)() {
    m.userExperienceAnalyticsNotAutopilotReadyDevice = value
}
// Sets the userExperienceAnalyticsOverview property value. User experience analytics overview
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsOverview property.
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value *UserExperienceAnalyticsOverview)() {
    m.userExperienceAnalyticsOverview = value
}
// Sets the userExperienceAnalyticsRegressionSummary property value. User experience analytics regression summary
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsRegressionSummary property.
func (m *DeviceManagement) SetUserExperienceAnalyticsRegressionSummary(value *UserExperienceAnalyticsRegressionSummary)() {
    m.userExperienceAnalyticsRegressionSummary = value
}
// Sets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsRemoteConnection property.
func (m *DeviceManagement) SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnection)() {
    m.userExperienceAnalyticsRemoteConnection = value
}
// Sets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsResourcePerformance property.
func (m *DeviceManagement) SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformance)() {
    m.userExperienceAnalyticsResourcePerformance = value
}
// Sets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsScoreHistory property.
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistory)() {
    m.userExperienceAnalyticsScoreHistory = value
}
// Sets the userExperienceAnalyticsSettings property value. User experience analytics device settings
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsSettings property.
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value *UserExperienceAnalyticsSettings)() {
    m.userExperienceAnalyticsSettings = value
}
// Sets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric)() {
    m.userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = value
}
// Sets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
// Parameters:
//  - value : Value to set for the userExperienceAnalyticsWorkFromAnywhereMetrics property.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetric)() {
    m.userExperienceAnalyticsWorkFromAnywhereMetrics = value
}
// Sets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
// Parameters:
//  - value : Value to set for the userPfxCertificates property.
func (m *DeviceManagement) SetUserPfxCertificates(value []UserPFXCertificate)() {
    m.userPfxCertificates = value
}
// Sets the virtualEndpoint property value. 
// Parameters:
//  - value : Value to set for the virtualEndpoint property.
func (m *DeviceManagement) SetVirtualEndpoint(value *VirtualEndpoint)() {
    m.virtualEndpoint = value
}
// Sets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
// Parameters:
//  - value : Value to set for the windowsAutopilotDeploymentProfiles property.
func (m *DeviceManagement) SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfile)() {
    m.windowsAutopilotDeploymentProfiles = value
}
// Sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
// Parameters:
//  - value : Value to set for the windowsAutopilotDeviceIdentities property.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentity)() {
    m.windowsAutopilotDeviceIdentities = value
}
// Sets the windowsAutopilotSettings property value. The Windows autopilot account settings.
// Parameters:
//  - value : Value to set for the windowsAutopilotSettings property.
func (m *DeviceManagement) SetWindowsAutopilotSettings(value *WindowsAutopilotSettings)() {
    m.windowsAutopilotSettings = value
}
// Sets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
// Parameters:
//  - value : Value to set for the windowsDriverUpdateProfiles property.
func (m *DeviceManagement) SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfile)() {
    m.windowsDriverUpdateProfiles = value
}
// Sets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
// Parameters:
//  - value : Value to set for the windowsFeatureUpdateProfiles property.
func (m *DeviceManagement) SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfile)() {
    m.windowsFeatureUpdateProfiles = value
}
// Sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionAppLearningSummaries property.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummary)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
// Sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionNetworkLearningSummaries property.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummary)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
// Sets the windowsMalwareInformation property value. The list of affected malware in the tenant.
// Parameters:
//  - value : Value to set for the windowsMalwareInformation property.
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformation)() {
    m.windowsMalwareInformation = value
}
// Sets the windowsMalwareOverview property value. Malware overview for windows devices.
// Parameters:
//  - value : Value to set for the windowsMalwareOverview property.
func (m *DeviceManagement) SetWindowsMalwareOverview(value *WindowsMalwareOverview)() {
    m.windowsMalwareOverview = value
}
// Sets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
// Parameters:
//  - value : Value to set for the windowsQualityUpdateProfiles property.
func (m *DeviceManagement) SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfile)() {
    m.windowsQualityUpdateProfiles = value
}
// Sets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
// Parameters:
//  - value : Value to set for the windowsUpdateCatalogItems property.
func (m *DeviceManagement) SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItem)() {
    m.windowsUpdateCatalogItems = value
}
