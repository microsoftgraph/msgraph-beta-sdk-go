package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagement struct {
    Entity
    accountMoveCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    adminConsent *AdminConsent;
    advancedThreatProtectionOnboardingStateSummary *AdvancedThreatProtectionOnboardingStateSummary;
    androidDeviceOwnerEnrollmentProfiles []AndroidDeviceOwnerEnrollmentProfile;
    androidForWorkAppConfigurationSchemas []AndroidForWorkAppConfigurationSchema;
    androidForWorkEnrollmentProfiles []AndroidForWorkEnrollmentProfile;
    androidForWorkSettings *AndroidForWorkSettings;
    androidManagedStoreAccountEnterpriseSettings *AndroidManagedStoreAccountEnterpriseSettings;
    androidManagedStoreAppConfigurationSchemas []AndroidManagedStoreAppConfigurationSchema;
    applePushNotificationCertificate *ApplePushNotificationCertificate;
    appleUserInitiatedEnrollmentProfiles []AppleUserInitiatedEnrollmentProfile;
    assignmentFilters []DeviceAndAppManagementAssignmentFilter;
    auditEvents []AuditEvent;
    autopilotEvents []DeviceManagementAutopilotEvent;
    cartToClassAssociations []CartToClassAssociation;
    categories []DeviceManagementSettingCategory;
    chromeOSOnboardingSettings []ChromeOSOnboardingSettings;
    cloudPCConnectivityIssues []CloudPCConnectivityIssue;
    comanagedDevices []ManagedDevice;
    comanagementEligibleDevices []ComanagementEligibleDevice;
    complianceManagementPartners []ComplianceManagementPartner;
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    configManagerCollections []ConfigManagerCollection;
    configurationCategories []DeviceManagementConfigurationCategory;
    configurationPolicies []DeviceManagementConfigurationPolicy;
    configurationPolicyTemplates []DeviceManagementConfigurationPolicyTemplate;
    configurationSettings []DeviceManagementConfigurationSettingDefinition;
    dataSharingConsents []DataSharingConsent;
    depOnboardingSettings []DepOnboardingSetting;
    derivedCredentials []DeviceManagementDerivedCredentialSettings;
    detectedApps []DetectedApp;
    deviceCategories []DeviceCategory;
    deviceCompliancePolicies []DeviceCompliancePolicy;
    deviceCompliancePolicyDeviceStateSummary *DeviceCompliancePolicyDeviceStateSummary;
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummary;
    deviceComplianceReportSummarizationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deviceComplianceScripts []DeviceComplianceScript;
    deviceConfigurationConflictSummary []DeviceConfigurationConflictSummary;
    deviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary;
    deviceConfigurationRestrictedAppsViolations []RestrictedAppsViolation;
    deviceConfigurations []DeviceConfiguration;
    deviceConfigurationsAllManagedDeviceCertificateStates []ManagedAllDeviceCertificateState;
    deviceConfigurationUserStateSummaries *DeviceConfigurationUserStateSummary;
    deviceCustomAttributeShellScripts []DeviceCustomAttributeShellScript;
    deviceEnrollmentConfigurations []DeviceEnrollmentConfiguration;
    deviceHealthScripts []DeviceHealthScript;
    deviceManagementPartners []DeviceManagementPartner;
    deviceManagementScripts []DeviceManagementScript;
    deviceProtectionOverview *DeviceProtectionOverview;
    deviceShellScripts []DeviceShellScript;
    domainJoinConnectors []DeviceManagementDomainJoinConnector;
    embeddedSIMActivationCodePools []EmbeddedSIMActivationCodePool;
    exchangeConnectors []DeviceManagementExchangeConnector;
    exchangeOnPremisesPolicies []DeviceManagementExchangeOnPremisesPolicy;
    exchangeOnPremisesPolicy *DeviceManagementExchangeOnPremisesPolicy;
    groupPolicyCategories []GroupPolicyCategory;
    groupPolicyConfigurations []GroupPolicyConfiguration;
    groupPolicyDefinitionFiles []GroupPolicyDefinitionFile;
    groupPolicyDefinitions []GroupPolicyDefinition;
    groupPolicyMigrationReports []GroupPolicyMigrationReport;
    groupPolicyObjectFiles []GroupPolicyObjectFile;
    groupPolicyUploadedDefinitionFiles []GroupPolicyUploadedDefinitionFile;
    importedDeviceIdentities []ImportedDeviceIdentity;
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentity;
    intents []DeviceManagementIntent;
    intuneAccountId *string;
    intuneBrand *IntuneBrand;
    intuneBrandingProfiles []IntuneBrandingProfile;
    iosUpdateStatuses []IosUpdateDeviceStatus;
    lastReportAggregationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    legacyPcManangementEnabled *bool;
    macOSSoftwareUpdateAccountSummaries []MacOSSoftwareUpdateAccountSummary;
    managedDeviceCleanupSettings *ManagedDeviceCleanupSettings;
    managedDeviceEncryptionStates []ManagedDeviceEncryptionState;
    managedDeviceOverview *ManagedDeviceOverview;
    managedDevices []ManagedDevice;
    managementConditions []ManagementCondition;
    managementConditionStatements []ManagementConditionStatement;
    maximumDepTokens *int32;
    microsoftTunnelConfigurations []MicrosoftTunnelConfiguration;
    microsoftTunnelHealthThresholds []MicrosoftTunnelHealthThreshold;
    microsoftTunnelServerLogCollectionResponses []MicrosoftTunnelServerLogCollectionResponse;
    microsoftTunnelSites []MicrosoftTunnelSite;
    mobileAppTroubleshootingEvents []MobileAppTroubleshootingEvent;
    mobileThreatDefenseConnectors []MobileThreatDefenseConnector;
    ndesConnectors []NdesConnector;
    notificationMessageTemplates []NotificationMessageTemplate;
    remoteActionAudits []RemoteActionAudit;
    remoteAssistancePartners []RemoteAssistancePartner;
    remoteAssistanceSettings *RemoteAssistanceSettings;
    reports *DeviceManagementReports;
    resourceAccessProfiles []DeviceManagementResourceAccessProfileBase;
    resourceOperations []ResourceOperation;
    reusablePolicySettings []DeviceManagementReusablePolicySetting;
    reusableSettings []DeviceManagementConfigurationSettingDefinition;
    roleAssignments []DeviceAndAppManagementRoleAssignment;
    roleDefinitions []RoleDefinition;
    roleScopeTags []RoleScopeTag;
    settingDefinitions []DeviceManagementSettingDefinition;
    settings *DeviceManagementSettings;
    softwareUpdateStatusSummary *SoftwareUpdateStatusSummary;
    subscriptions *DeviceManagementSubscriptions;
    subscriptionState *DeviceManagementSubscriptionState;
    telecomExpenseManagementPartners []TelecomExpenseManagementPartner;
    templates []DeviceManagementTemplate;
    templateSettings []DeviceManagementConfigurationSettingTemplate;
    termsAndConditions []TermsAndConditions;
    troubleshootingEvents []DeviceManagementTroubleshootingEvent;
    unlicensedAdminstratorsEnabled *bool;
    userExperienceAnalyticsAppHealthApplicationPerformance []UserExperienceAnalyticsAppHealthApplicationPerformance;
    userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion;
    userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion;
    userExperienceAnalyticsAppHealthDeviceModelPerformance []UserExperienceAnalyticsAppHealthDeviceModelPerformance;
    userExperienceAnalyticsAppHealthDevicePerformance []UserExperienceAnalyticsAppHealthDevicePerformance;
    userExperienceAnalyticsAppHealthDevicePerformanceDetails []UserExperienceAnalyticsAppHealthDevicePerformanceDetails;
    userExperienceAnalyticsAppHealthOSVersionPerformance []UserExperienceAnalyticsAppHealthOSVersionPerformance;
    userExperienceAnalyticsAppHealthOverview *UserExperienceAnalyticsCategory;
    userExperienceAnalyticsBaselines []UserExperienceAnalyticsBaseline;
    userExperienceAnalyticsCategories []UserExperienceAnalyticsCategory;
    userExperienceAnalyticsDeviceMetricHistory []UserExperienceAnalyticsMetricHistory;
    userExperienceAnalyticsDevicePerformance []UserExperienceAnalyticsDevicePerformance;
    userExperienceAnalyticsDeviceScores []UserExperienceAnalyticsDeviceScores;
    userExperienceAnalyticsDeviceStartupHistory []UserExperienceAnalyticsDeviceStartupHistory;
    userExperienceAnalyticsDeviceStartupProcesses []UserExperienceAnalyticsDeviceStartupProcess;
    userExperienceAnalyticsDeviceStartupProcessPerformance []UserExperienceAnalyticsDeviceStartupProcessPerformance;
    userExperienceAnalyticsDevicesWithoutCloudIdentity []UserExperienceAnalyticsDeviceWithoutCloudIdentity;
    userExperienceAnalyticsImpactingProcess []UserExperienceAnalyticsImpactingProcess;
    userExperienceAnalyticsMetricHistory []UserExperienceAnalyticsMetricHistory;
    userExperienceAnalyticsNotAutopilotReadyDevice []UserExperienceAnalyticsNotAutopilotReadyDevice;
    userExperienceAnalyticsOverview *UserExperienceAnalyticsOverview;
    userExperienceAnalyticsRegressionSummary *UserExperienceAnalyticsRegressionSummary;
    userExperienceAnalyticsRemoteConnection []UserExperienceAnalyticsRemoteConnection;
    userExperienceAnalyticsResourcePerformance []UserExperienceAnalyticsResourcePerformance;
    userExperienceAnalyticsScoreHistory []UserExperienceAnalyticsScoreHistory;
    userExperienceAnalyticsSettings *UserExperienceAnalyticsSettings;
    userExperienceAnalyticsWorkFromAnywhereMetrics []UserExperienceAnalyticsWorkFromAnywhereMetric;
    userPfxCertificates []UserPFXCertificate;
    virtualEndpoint *VirtualEndpoint;
    windowsAutopilotDeploymentProfiles []WindowsAutopilotDeploymentProfile;
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentity;
    windowsAutopilotSettings *WindowsAutopilotSettings;
    windowsFeatureUpdateProfiles []WindowsFeatureUpdateProfile;
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummary;
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummary;
    windowsMalwareInformation []WindowsMalwareInformation;
    windowsMalwareOverview *WindowsMalwareOverview;
    windowsQualityUpdateProfiles []WindowsQualityUpdateProfile;
    windowsUpdateCatalogItems []WindowsUpdateCatalogItem;
}
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagement) GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.accountMoveCompletionDateTime
    }
}
func (m *DeviceManagement) GetAdminConsent()(*AdminConsent) {
    if m == nil {
        return nil
    } else {
        return m.adminConsent
    }
}
func (m *DeviceManagement) GetAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionOnboardingStateSummary
    }
}
func (m *DeviceManagement) GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.androidDeviceOwnerEnrollmentProfiles
    }
}
func (m *DeviceManagement) GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchema) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkAppConfigurationSchemas
    }
}
func (m *DeviceManagement) GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkEnrollmentProfiles
    }
}
func (m *DeviceManagement) GetAndroidForWorkSettings()(*AndroidForWorkSettings) {
    if m == nil {
        return nil
    } else {
        return m.androidForWorkSettings
    }
}
func (m *DeviceManagement) GetAndroidManagedStoreAccountEnterpriseSettings()(*AndroidManagedStoreAccountEnterpriseSettings) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedStoreAccountEnterpriseSettings
    }
}
func (m *DeviceManagement) GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchema) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedStoreAppConfigurationSchemas
    }
}
func (m *DeviceManagement) GetApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    if m == nil {
        return nil
    } else {
        return m.applePushNotificationCertificate
    }
}
func (m *DeviceManagement) GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfile) {
    if m == nil {
        return nil
    } else {
        return m.appleUserInitiatedEnrollmentProfiles
    }
}
func (m *DeviceManagement) GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilter) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilters
    }
}
func (m *DeviceManagement) GetAuditEvents()([]AuditEvent) {
    if m == nil {
        return nil
    } else {
        return m.auditEvents
    }
}
func (m *DeviceManagement) GetAutopilotEvents()([]DeviceManagementAutopilotEvent) {
    if m == nil {
        return nil
    } else {
        return m.autopilotEvents
    }
}
func (m *DeviceManagement) GetCartToClassAssociations()([]CartToClassAssociation) {
    if m == nil {
        return nil
    } else {
        return m.cartToClassAssociations
    }
}
func (m *DeviceManagement) GetCategories()([]DeviceManagementSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *DeviceManagement) GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettings) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSOnboardingSettings
    }
}
func (m *DeviceManagement) GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssue) {
    if m == nil {
        return nil
    } else {
        return m.cloudPCConnectivityIssues
    }
}
func (m *DeviceManagement) GetComanagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.comanagedDevices
    }
}
func (m *DeviceManagement) GetComanagementEligibleDevices()([]ComanagementEligibleDevice) {
    if m == nil {
        return nil
    } else {
        return m.comanagementEligibleDevices
    }
}
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.complianceManagementPartners
    }
}
func (m *DeviceManagement) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
func (m *DeviceManagement) GetConfigManagerCollections()([]ConfigManagerCollection) {
    if m == nil {
        return nil
    } else {
        return m.configManagerCollections
    }
}
func (m *DeviceManagement) GetConfigurationCategories()([]DeviceManagementConfigurationCategory) {
    if m == nil {
        return nil
    } else {
        return m.configurationCategories
    }
}
func (m *DeviceManagement) GetConfigurationPolicies()([]DeviceManagementConfigurationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.configurationPolicies
    }
}
func (m *DeviceManagement) GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplate) {
    if m == nil {
        return nil
    } else {
        return m.configurationPolicyTemplates
    }
}
func (m *DeviceManagement) GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.configurationSettings
    }
}
func (m *DeviceManagement) GetDataSharingConsents()([]DataSharingConsent) {
    if m == nil {
        return nil
    } else {
        return m.dataSharingConsents
    }
}
func (m *DeviceManagement) GetDepOnboardingSettings()([]DepOnboardingSetting) {
    if m == nil {
        return nil
    } else {
        return m.depOnboardingSettings
    }
}
func (m *DeviceManagement) GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettings) {
    if m == nil {
        return nil
    } else {
        return m.derivedCredentials
    }
}
func (m *DeviceManagement) GetDetectedApps()([]DetectedApp) {
    if m == nil {
        return nil
    } else {
        return m.detectedApps
    }
}
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategories
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicies
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyDeviceStateSummary
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
func (m *DeviceManagement) GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceReportSummarizationDateTime
    }
}
func (m *DeviceManagement) GetDeviceComplianceScripts()([]DeviceComplianceScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScripts
    }
}
func (m *DeviceManagement) GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationConflictSummary
    }
}
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationDeviceStateSummaries
    }
}
func (m *DeviceManagement) GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolation) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationRestrictedAppsViolations
    }
}
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurations
    }
}
func (m *DeviceManagement) GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateState) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationsAllManagedDeviceCertificateStates
    }
}
func (m *DeviceManagement) GetDeviceConfigurationUserStateSummaries()(*DeviceConfigurationUserStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationUserStateSummaries
    }
}
func (m *DeviceManagement) GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceCustomAttributeShellScripts
    }
}
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentConfigurations
    }
}
func (m *DeviceManagement) GetDeviceHealthScripts()([]DeviceHealthScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthScripts
    }
}
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementPartners
    }
}
func (m *DeviceManagement) GetDeviceManagementScripts()([]DeviceManagementScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementScripts
    }
}
func (m *DeviceManagement) GetDeviceProtectionOverview()(*DeviceProtectionOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceProtectionOverview
    }
}
func (m *DeviceManagement) GetDeviceShellScripts()([]DeviceShellScript) {
    if m == nil {
        return nil
    } else {
        return m.deviceShellScripts
    }
}
func (m *DeviceManagement) GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnector) {
    if m == nil {
        return nil
    } else {
        return m.domainJoinConnectors
    }
}
func (m *DeviceManagement) GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePool) {
    if m == nil {
        return nil
    } else {
        return m.embeddedSIMActivationCodePools
    }
}
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnector) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectors
    }
}
func (m *DeviceManagement) GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicy) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOnPremisesPolicies
    }
}
func (m *DeviceManagement) GetExchangeOnPremisesPolicy()(*DeviceManagementExchangeOnPremisesPolicy) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOnPremisesPolicy
    }
}
func (m *DeviceManagement) GetGroupPolicyCategories()([]GroupPolicyCategory) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCategories
    }
}
func (m *DeviceManagement) GetGroupPolicyConfigurations()([]GroupPolicyConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyConfigurations
    }
}
func (m *DeviceManagement) GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyDefinitionFiles
    }
}
func (m *DeviceManagement) GetGroupPolicyDefinitions()([]GroupPolicyDefinition) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyDefinitions
    }
}
func (m *DeviceManagement) GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReport) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyMigrationReports
    }
}
func (m *DeviceManagement) GetGroupPolicyObjectFiles()([]GroupPolicyObjectFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectFiles
    }
}
func (m *DeviceManagement) GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFile) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyUploadedDefinitionFiles
    }
}
func (m *DeviceManagement) GetImportedDeviceIdentities()([]ImportedDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentities
    }
}
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedWindowsAutopilotDeviceIdentities
    }
}
func (m *DeviceManagement) GetIntents()([]DeviceManagementIntent) {
    if m == nil {
        return nil
    } else {
        return m.intents
    }
}
func (m *DeviceManagement) GetIntuneAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneAccountId
    }
}
func (m *DeviceManagement) GetIntuneBrand()(*IntuneBrand) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrand
    }
}
func (m *DeviceManagement) GetIntuneBrandingProfiles()([]IntuneBrandingProfile) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrandingProfiles
    }
}
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.iosUpdateStatuses
    }
}
func (m *DeviceManagement) GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportAggregationDateTime
    }
}
func (m *DeviceManagement) GetLegacyPcManangementEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.legacyPcManangementEnabled
    }
}
func (m *DeviceManagement) GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummary) {
    if m == nil {
        return nil
    } else {
        return m.macOSSoftwareUpdateAccountSummaries
    }
}
func (m *DeviceManagement) GetManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCleanupSettings
    }
}
func (m *DeviceManagement) GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionState) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceEncryptionStates
    }
}
func (m *DeviceManagement) GetManagedDeviceOverview()(*ManagedDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOverview
    }
}
func (m *DeviceManagement) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
func (m *DeviceManagement) GetManagementConditions()([]ManagementCondition) {
    if m == nil {
        return nil
    } else {
        return m.managementConditions
    }
}
func (m *DeviceManagement) GetManagementConditionStatements()([]ManagementConditionStatement) {
    if m == nil {
        return nil
    } else {
        return m.managementConditionStatements
    }
}
func (m *DeviceManagement) GetMaximumDepTokens()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumDepTokens
    }
}
func (m *DeviceManagement) GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfigurations
    }
}
func (m *DeviceManagement) GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThreshold) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelHealthThresholds
    }
}
func (m *DeviceManagement) GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponse) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServerLogCollectionResponses
    }
}
func (m *DeviceManagement) GetMicrosoftTunnelSites()([]MicrosoftTunnelSite) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelSites
    }
}
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppTroubleshootingEvents
    }
}
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnector) {
    if m == nil {
        return nil
    } else {
        return m.mobileThreatDefenseConnectors
    }
}
func (m *DeviceManagement) GetNdesConnectors()([]NdesConnector) {
    if m == nil {
        return nil
    } else {
        return m.ndesConnectors
    }
}
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplate) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageTemplates
    }
}
func (m *DeviceManagement) GetRemoteActionAudits()([]RemoteActionAudit) {
    if m == nil {
        return nil
    } else {
        return m.remoteActionAudits
    }
}
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartner) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistancePartners
    }
}
func (m *DeviceManagement) GetRemoteAssistanceSettings()(*RemoteAssistanceSettings) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSettings
    }
}
func (m *DeviceManagement) GetReports()(*DeviceManagementReports) {
    if m == nil {
        return nil
    } else {
        return m.reports
    }
}
func (m *DeviceManagement) GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBase) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccessProfiles
    }
}
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperation) {
    if m == nil {
        return nil
    } else {
        return m.resourceOperations
    }
}
func (m *DeviceManagement) GetReusablePolicySettings()([]DeviceManagementReusablePolicySetting) {
    if m == nil {
        return nil
    } else {
        return m.reusablePolicySettings
    }
}
func (m *DeviceManagement) GetReusableSettings()([]DeviceManagementConfigurationSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.reusableSettings
    }
}
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
func (m *DeviceManagement) GetRoleScopeTags()([]RoleScopeTag) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTags
    }
}
func (m *DeviceManagement) GetSettingDefinitions()([]DeviceManagementSettingDefinition) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
func (m *DeviceManagement) GetSettings()(*DeviceManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateStatusSummary
    }
}
func (m *DeviceManagement) GetSubscriptions()(*DeviceManagementSubscriptions) {
    if m == nil {
        return nil
    } else {
        return m.subscriptions
    }
}
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionState
    }
}
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.telecomExpenseManagementPartners
    }
}
func (m *DeviceManagement) GetTemplates()([]DeviceManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templates
    }
}
func (m *DeviceManagement) GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplate) {
    if m == nil {
        return nil
    } else {
        return m.templateSettings
    }
}
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingEvents
    }
}
func (m *DeviceManagement) GetUnlicensedAdminstratorsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.unlicensedAdminstratorsEnabled
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDeviceModelPerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDevicePerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthDevicePerformanceDetails
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthOSVersionPerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(*UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsAppHealthOverview
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaseline) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsBaselines
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsCategories
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceMetricHistory
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDevicePerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScores) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceScores
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupHistory
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcess) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupProcesses
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDeviceStartupProcessPerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsDevicesWithoutCloudIdentity
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcess) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsImpactingProcess
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsMetricHistory
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDevice) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsNotAutopilotReadyDevice
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(*UserExperienceAnalyticsOverview) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsOverview
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsRegressionSummary()(*UserExperienceAnalyticsRegressionSummary) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsRegressionSummary
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnection) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsRemoteConnection
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformance) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsResourcePerformance
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistory) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsScoreHistory
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsSettings
    }
}
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetric) {
    if m == nil {
        return nil
    } else {
        return m.userExperienceAnalyticsWorkFromAnywhereMetrics
    }
}
func (m *DeviceManagement) GetUserPfxCertificates()([]UserPFXCertificate) {
    if m == nil {
        return nil
    } else {
        return m.userPfxCertificates
    }
}
func (m *DeviceManagement) GetVirtualEndpoint()(*VirtualEndpoint) {
    if m == nil {
        return nil
    } else {
        return m.virtualEndpoint
    }
}
func (m *DeviceManagement) GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeploymentProfiles
    }
}
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeviceIdentities
    }
}
func (m *DeviceManagement) GetWindowsAutopilotSettings()(*WindowsAutopilotSettings) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotSettings
    }
}
func (m *DeviceManagement) GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsFeatureUpdateProfiles
    }
}
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionAppLearningSummaries
    }
}
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionNetworkLearningSummaries
    }
}
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformation) {
    if m == nil {
        return nil
    } else {
        return m.windowsMalwareInformation
    }
}
func (m *DeviceManagement) GetWindowsMalwareOverview()(*WindowsMalwareOverview) {
    if m == nil {
        return nil
    } else {
        return m.windowsMalwareOverview
    }
}
func (m *DeviceManagement) GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfile) {
    if m == nil {
        return nil
    } else {
        return m.windowsQualityUpdateProfiles
    }
}
func (m *DeviceManagement) GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItem) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateCatalogItems
    }
}
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
func (m *DeviceManagement) SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.accountMoveCompletionDateTime = value
}
func (m *DeviceManagement) SetAdminConsent(value *AdminConsent)() {
    m.adminConsent = value
}
func (m *DeviceManagement) SetAdvancedThreatProtectionOnboardingStateSummary(value *AdvancedThreatProtectionOnboardingStateSummary)() {
    m.advancedThreatProtectionOnboardingStateSummary = value
}
func (m *DeviceManagement) SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfile)() {
    m.androidDeviceOwnerEnrollmentProfiles = value
}
func (m *DeviceManagement) SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchema)() {
    m.androidForWorkAppConfigurationSchemas = value
}
func (m *DeviceManagement) SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfile)() {
    m.androidForWorkEnrollmentProfiles = value
}
func (m *DeviceManagement) SetAndroidForWorkSettings(value *AndroidForWorkSettings)() {
    m.androidForWorkSettings = value
}
func (m *DeviceManagement) SetAndroidManagedStoreAccountEnterpriseSettings(value *AndroidManagedStoreAccountEnterpriseSettings)() {
    m.androidManagedStoreAccountEnterpriseSettings = value
}
func (m *DeviceManagement) SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchema)() {
    m.androidManagedStoreAppConfigurationSchemas = value
}
func (m *DeviceManagement) SetApplePushNotificationCertificate(value *ApplePushNotificationCertificate)() {
    m.applePushNotificationCertificate = value
}
func (m *DeviceManagement) SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfile)() {
    m.appleUserInitiatedEnrollmentProfiles = value
}
func (m *DeviceManagement) SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilter)() {
    m.assignmentFilters = value
}
func (m *DeviceManagement) SetAuditEvents(value []AuditEvent)() {
    m.auditEvents = value
}
func (m *DeviceManagement) SetAutopilotEvents(value []DeviceManagementAutopilotEvent)() {
    m.autopilotEvents = value
}
func (m *DeviceManagement) SetCartToClassAssociations(value []CartToClassAssociation)() {
    m.cartToClassAssociations = value
}
func (m *DeviceManagement) SetCategories(value []DeviceManagementSettingCategory)() {
    m.categories = value
}
func (m *DeviceManagement) SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettings)() {
    m.chromeOSOnboardingSettings = value
}
func (m *DeviceManagement) SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssue)() {
    m.cloudPCConnectivityIssues = value
}
func (m *DeviceManagement) SetComanagedDevices(value []ManagedDevice)() {
    m.comanagedDevices = value
}
func (m *DeviceManagement) SetComanagementEligibleDevices(value []ComanagementEligibleDevice)() {
    m.comanagementEligibleDevices = value
}
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartner)() {
    m.complianceManagementPartners = value
}
func (m *DeviceManagement) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    m.conditionalAccessSettings = value
}
func (m *DeviceManagement) SetConfigManagerCollections(value []ConfigManagerCollection)() {
    m.configManagerCollections = value
}
func (m *DeviceManagement) SetConfigurationCategories(value []DeviceManagementConfigurationCategory)() {
    m.configurationCategories = value
}
func (m *DeviceManagement) SetConfigurationPolicies(value []DeviceManagementConfigurationPolicy)() {
    m.configurationPolicies = value
}
func (m *DeviceManagement) SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplate)() {
    m.configurationPolicyTemplates = value
}
func (m *DeviceManagement) SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinition)() {
    m.configurationSettings = value
}
func (m *DeviceManagement) SetDataSharingConsents(value []DataSharingConsent)() {
    m.dataSharingConsents = value
}
func (m *DeviceManagement) SetDepOnboardingSettings(value []DepOnboardingSetting)() {
    m.depOnboardingSettings = value
}
func (m *DeviceManagement) SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettings)() {
    m.derivedCredentials = value
}
func (m *DeviceManagement) SetDetectedApps(value []DetectedApp)() {
    m.detectedApps = value
}
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategory)() {
    m.deviceCategories = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicy)() {
    m.deviceCompliancePolicies = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value *DeviceCompliancePolicyDeviceStateSummary)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
func (m *DeviceManagement) SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceComplianceReportSummarizationDateTime = value
}
func (m *DeviceManagement) SetDeviceComplianceScripts(value []DeviceComplianceScript)() {
    m.deviceComplianceScripts = value
}
func (m *DeviceManagement) SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummary)() {
    m.deviceConfigurationConflictSummary = value
}
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value *DeviceConfigurationDeviceStateSummary)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
func (m *DeviceManagement) SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolation)() {
    m.deviceConfigurationRestrictedAppsViolations = value
}
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfiguration)() {
    m.deviceConfigurations = value
}
func (m *DeviceManagement) SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateState)() {
    m.deviceConfigurationsAllManagedDeviceCertificateStates = value
}
func (m *DeviceManagement) SetDeviceConfigurationUserStateSummaries(value *DeviceConfigurationUserStateSummary)() {
    m.deviceConfigurationUserStateSummaries = value
}
func (m *DeviceManagement) SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScript)() {
    m.deviceCustomAttributeShellScripts = value
}
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfiguration)() {
    m.deviceEnrollmentConfigurations = value
}
func (m *DeviceManagement) SetDeviceHealthScripts(value []DeviceHealthScript)() {
    m.deviceHealthScripts = value
}
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartner)() {
    m.deviceManagementPartners = value
}
func (m *DeviceManagement) SetDeviceManagementScripts(value []DeviceManagementScript)() {
    m.deviceManagementScripts = value
}
func (m *DeviceManagement) SetDeviceProtectionOverview(value *DeviceProtectionOverview)() {
    m.deviceProtectionOverview = value
}
func (m *DeviceManagement) SetDeviceShellScripts(value []DeviceShellScript)() {
    m.deviceShellScripts = value
}
func (m *DeviceManagement) SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnector)() {
    m.domainJoinConnectors = value
}
func (m *DeviceManagement) SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePool)() {
    m.embeddedSIMActivationCodePools = value
}
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnector)() {
    m.exchangeConnectors = value
}
func (m *DeviceManagement) SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicy)() {
    m.exchangeOnPremisesPolicies = value
}
func (m *DeviceManagement) SetExchangeOnPremisesPolicy(value *DeviceManagementExchangeOnPremisesPolicy)() {
    m.exchangeOnPremisesPolicy = value
}
func (m *DeviceManagement) SetGroupPolicyCategories(value []GroupPolicyCategory)() {
    m.groupPolicyCategories = value
}
func (m *DeviceManagement) SetGroupPolicyConfigurations(value []GroupPolicyConfiguration)() {
    m.groupPolicyConfigurations = value
}
func (m *DeviceManagement) SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFile)() {
    m.groupPolicyDefinitionFiles = value
}
func (m *DeviceManagement) SetGroupPolicyDefinitions(value []GroupPolicyDefinition)() {
    m.groupPolicyDefinitions = value
}
func (m *DeviceManagement) SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReport)() {
    m.groupPolicyMigrationReports = value
}
func (m *DeviceManagement) SetGroupPolicyObjectFiles(value []GroupPolicyObjectFile)() {
    m.groupPolicyObjectFiles = value
}
func (m *DeviceManagement) SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFile)() {
    m.groupPolicyUploadedDefinitionFiles = value
}
func (m *DeviceManagement) SetImportedDeviceIdentities(value []ImportedDeviceIdentity)() {
    m.importedDeviceIdentities = value
}
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentity)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
func (m *DeviceManagement) SetIntents(value []DeviceManagementIntent)() {
    m.intents = value
}
func (m *DeviceManagement) SetIntuneAccountId(value *string)() {
    m.intuneAccountId = value
}
func (m *DeviceManagement) SetIntuneBrand(value *IntuneBrand)() {
    m.intuneBrand = value
}
func (m *DeviceManagement) SetIntuneBrandingProfiles(value []IntuneBrandingProfile)() {
    m.intuneBrandingProfiles = value
}
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatus)() {
    m.iosUpdateStatuses = value
}
func (m *DeviceManagement) SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportAggregationDateTime = value
}
func (m *DeviceManagement) SetLegacyPcManangementEnabled(value *bool)() {
    m.legacyPcManangementEnabled = value
}
func (m *DeviceManagement) SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummary)() {
    m.macOSSoftwareUpdateAccountSummaries = value
}
func (m *DeviceManagement) SetManagedDeviceCleanupSettings(value *ManagedDeviceCleanupSettings)() {
    m.managedDeviceCleanupSettings = value
}
func (m *DeviceManagement) SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionState)() {
    m.managedDeviceEncryptionStates = value
}
func (m *DeviceManagement) SetManagedDeviceOverview(value *ManagedDeviceOverview)() {
    m.managedDeviceOverview = value
}
func (m *DeviceManagement) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
func (m *DeviceManagement) SetManagementConditions(value []ManagementCondition)() {
    m.managementConditions = value
}
func (m *DeviceManagement) SetManagementConditionStatements(value []ManagementConditionStatement)() {
    m.managementConditionStatements = value
}
func (m *DeviceManagement) SetMaximumDepTokens(value *int32)() {
    m.maximumDepTokens = value
}
func (m *DeviceManagement) SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfiguration)() {
    m.microsoftTunnelConfigurations = value
}
func (m *DeviceManagement) SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThreshold)() {
    m.microsoftTunnelHealthThresholds = value
}
func (m *DeviceManagement) SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponse)() {
    m.microsoftTunnelServerLogCollectionResponses = value
}
func (m *DeviceManagement) SetMicrosoftTunnelSites(value []MicrosoftTunnelSite)() {
    m.microsoftTunnelSites = value
}
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEvent)() {
    m.mobileAppTroubleshootingEvents = value
}
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnector)() {
    m.mobileThreatDefenseConnectors = value
}
func (m *DeviceManagement) SetNdesConnectors(value []NdesConnector)() {
    m.ndesConnectors = value
}
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplate)() {
    m.notificationMessageTemplates = value
}
func (m *DeviceManagement) SetRemoteActionAudits(value []RemoteActionAudit)() {
    m.remoteActionAudits = value
}
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartner)() {
    m.remoteAssistancePartners = value
}
func (m *DeviceManagement) SetRemoteAssistanceSettings(value *RemoteAssistanceSettings)() {
    m.remoteAssistanceSettings = value
}
func (m *DeviceManagement) SetReports(value *DeviceManagementReports)() {
    m.reports = value
}
func (m *DeviceManagement) SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBase)() {
    m.resourceAccessProfiles = value
}
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperation)() {
    m.resourceOperations = value
}
func (m *DeviceManagement) SetReusablePolicySettings(value []DeviceManagementReusablePolicySetting)() {
    m.reusablePolicySettings = value
}
func (m *DeviceManagement) SetReusableSettings(value []DeviceManagementConfigurationSettingDefinition)() {
    m.reusableSettings = value
}
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignment)() {
    m.roleAssignments = value
}
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinition)() {
    m.roleDefinitions = value
}
func (m *DeviceManagement) SetRoleScopeTags(value []RoleScopeTag)() {
    m.roleScopeTags = value
}
func (m *DeviceManagement) SetSettingDefinitions(value []DeviceManagementSettingDefinition)() {
    m.settingDefinitions = value
}
func (m *DeviceManagement) SetSettings(value *DeviceManagementSettings)() {
    m.settings = value
}
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value *SoftwareUpdateStatusSummary)() {
    m.softwareUpdateStatusSummary = value
}
func (m *DeviceManagement) SetSubscriptions(value *DeviceManagementSubscriptions)() {
    m.subscriptions = value
}
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartner)() {
    m.telecomExpenseManagementPartners = value
}
func (m *DeviceManagement) SetTemplates(value []DeviceManagementTemplate)() {
    m.templates = value
}
func (m *DeviceManagement) SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplate)() {
    m.templateSettings = value
}
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditions)() {
    m.termsAndConditions = value
}
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    m.troubleshootingEvents = value
}
func (m *DeviceManagement) SetUnlicensedAdminstratorsEnabled(value *bool)() {
    m.unlicensedAdminstratorsEnabled = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformance)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion)() {
    m.userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformance)() {
    m.userExperienceAnalyticsAppHealthDeviceModelPerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformance)() {
    m.userExperienceAnalyticsAppHealthDevicePerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetails)() {
    m.userExperienceAnalyticsAppHealthDevicePerformanceDetails = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformance)() {
    m.userExperienceAnalyticsAppHealthOSVersionPerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value *UserExperienceAnalyticsCategory)() {
    m.userExperienceAnalyticsAppHealthOverview = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaseline)() {
    m.userExperienceAnalyticsBaselines = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategory)() {
    m.userExperienceAnalyticsCategories = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistory)() {
    m.userExperienceAnalyticsDeviceMetricHistory = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformance)() {
    m.userExperienceAnalyticsDevicePerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScores)() {
    m.userExperienceAnalyticsDeviceScores = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistory)() {
    m.userExperienceAnalyticsDeviceStartupHistory = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcess)() {
    m.userExperienceAnalyticsDeviceStartupProcesses = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformance)() {
    m.userExperienceAnalyticsDeviceStartupProcessPerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentity)() {
    m.userExperienceAnalyticsDevicesWithoutCloudIdentity = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcess)() {
    m.userExperienceAnalyticsImpactingProcess = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistory)() {
    m.userExperienceAnalyticsMetricHistory = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDevice)() {
    m.userExperienceAnalyticsNotAutopilotReadyDevice = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value *UserExperienceAnalyticsOverview)() {
    m.userExperienceAnalyticsOverview = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsRegressionSummary(value *UserExperienceAnalyticsRegressionSummary)() {
    m.userExperienceAnalyticsRegressionSummary = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnection)() {
    m.userExperienceAnalyticsRemoteConnection = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformance)() {
    m.userExperienceAnalyticsResourcePerformance = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistory)() {
    m.userExperienceAnalyticsScoreHistory = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value *UserExperienceAnalyticsSettings)() {
    m.userExperienceAnalyticsSettings = value
}
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetric)() {
    m.userExperienceAnalyticsWorkFromAnywhereMetrics = value
}
func (m *DeviceManagement) SetUserPfxCertificates(value []UserPFXCertificate)() {
    m.userPfxCertificates = value
}
func (m *DeviceManagement) SetVirtualEndpoint(value *VirtualEndpoint)() {
    m.virtualEndpoint = value
}
func (m *DeviceManagement) SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfile)() {
    m.windowsAutopilotDeploymentProfiles = value
}
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentity)() {
    m.windowsAutopilotDeviceIdentities = value
}
func (m *DeviceManagement) SetWindowsAutopilotSettings(value *WindowsAutopilotSettings)() {
    m.windowsAutopilotSettings = value
}
func (m *DeviceManagement) SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfile)() {
    m.windowsFeatureUpdateProfiles = value
}
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummary)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummary)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformation)() {
    m.windowsMalwareInformation = value
}
func (m *DeviceManagement) SetWindowsMalwareOverview(value *WindowsMalwareOverview)() {
    m.windowsMalwareOverview = value
}
func (m *DeviceManagement) SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfile)() {
    m.windowsQualityUpdateProfiles = value
}
func (m *DeviceManagement) SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItem)() {
    m.windowsUpdateCatalogItems = value
}
