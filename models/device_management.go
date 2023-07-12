package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagement singleton entity that acts as a container for all device management functionality.
type DeviceManagement struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeviceManagement instantiates a new deviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagement(), nil
}
// GetAccountMoveCompletionDateTime gets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("accountMoveCompletionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAdminConsent gets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) GetAdminConsent()(AdminConsentable) {
    val, err := m.GetBackingStore().Get("adminConsent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminConsentable)
    }
    return nil
}
// GetAdvancedThreatProtectionOnboardingStateSummary gets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) GetAdvancedThreatProtectionOnboardingStateSummary()(AdvancedThreatProtectionOnboardingStateSummaryable) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOnboardingStateSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdvancedThreatProtectionOnboardingStateSummaryable)
    }
    return nil
}
// GetAndroidDeviceOwnerEnrollmentProfiles gets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("androidDeviceOwnerEnrollmentProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerEnrollmentProfileable)
    }
    return nil
}
// GetAndroidForWorkAppConfigurationSchemas gets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchemaable) {
    val, err := m.GetBackingStore().Get("androidForWorkAppConfigurationSchemas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidForWorkAppConfigurationSchemaable)
    }
    return nil
}
// GetAndroidForWorkEnrollmentProfiles gets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("androidForWorkEnrollmentProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidForWorkEnrollmentProfileable)
    }
    return nil
}
// GetAndroidForWorkSettings gets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) GetAndroidForWorkSettings()(AndroidForWorkSettingsable) {
    val, err := m.GetBackingStore().Get("androidForWorkSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidForWorkSettingsable)
    }
    return nil
}
// GetAndroidManagedStoreAccountEnterpriseSettings gets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) GetAndroidManagedStoreAccountEnterpriseSettings()(AndroidManagedStoreAccountEnterpriseSettingsable) {
    val, err := m.GetBackingStore().Get("androidManagedStoreAccountEnterpriseSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidManagedStoreAccountEnterpriseSettingsable)
    }
    return nil
}
// GetAndroidManagedStoreAppConfigurationSchemas gets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchemaable) {
    val, err := m.GetBackingStore().Get("androidManagedStoreAppConfigurationSchemas")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidManagedStoreAppConfigurationSchemaable)
    }
    return nil
}
// GetApplePushNotificationCertificate gets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable) {
    val, err := m.GetBackingStore().Get("applePushNotificationCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApplePushNotificationCertificateable)
    }
    return nil
}
// GetAppleUserInitiatedEnrollmentProfiles gets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfileable) {
    val, err := m.GetBackingStore().Get("appleUserInitiatedEnrollmentProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AppleUserInitiatedEnrollmentProfileable)
    }
    return nil
}
// GetAssignmentFilters gets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilterable) {
    val, err := m.GetBackingStore().Get("assignmentFilters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAndAppManagementAssignmentFilterable)
    }
    return nil
}
// GetAuditEvents gets the auditEvents property value. The Audit Events
func (m *DeviceManagement) GetAuditEvents()([]AuditEventable) {
    val, err := m.GetBackingStore().Get("auditEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuditEventable)
    }
    return nil
}
// GetAutopilotEvents gets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) GetAutopilotEvents()([]DeviceManagementAutopilotEventable) {
    val, err := m.GetBackingStore().Get("autopilotEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementAutopilotEventable)
    }
    return nil
}
// GetCartToClassAssociations gets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) GetCartToClassAssociations()([]CartToClassAssociationable) {
    val, err := m.GetBackingStore().Get("cartToClassAssociations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CartToClassAssociationable)
    }
    return nil
}
// GetCategories gets the categories property value. The available categories
func (m *DeviceManagement) GetCategories()([]DeviceManagementSettingCategoryable) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingCategoryable)
    }
    return nil
}
// GetCertificateConnectorDetails gets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) GetCertificateConnectorDetails()([]CertificateConnectorDetailsable) {
    val, err := m.GetBackingStore().Get("certificateConnectorDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateConnectorDetailsable)
    }
    return nil
}
// GetChromeOSOnboardingSettings gets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettingsable) {
    val, err := m.GetBackingStore().Get("chromeOSOnboardingSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ChromeOSOnboardingSettingsable)
    }
    return nil
}
// GetCloudPCConnectivityIssues gets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssueable) {
    val, err := m.GetBackingStore().Get("cloudPCConnectivityIssues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPCConnectivityIssueable)
    }
    return nil
}
// GetComanagedDevices gets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) GetComanagedDevices()([]ManagedDeviceable) {
    val, err := m.GetBackingStore().Get("comanagedDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceable)
    }
    return nil
}
// GetComanagementEligibleDevices gets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) GetComanagementEligibleDevices()([]ComanagementEligibleDeviceable) {
    val, err := m.GetBackingStore().Get("comanagementEligibleDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ComanagementEligibleDeviceable)
    }
    return nil
}
// GetComplianceCategories gets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) GetComplianceCategories()([]DeviceManagementConfigurationCategoryable) {
    val, err := m.GetBackingStore().Get("complianceCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationCategoryable)
    }
    return nil
}
// GetComplianceManagementPartners gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartnerable) {
    val, err := m.GetBackingStore().Get("complianceManagementPartners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ComplianceManagementPartnerable)
    }
    return nil
}
// GetCompliancePolicies gets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) GetCompliancePolicies()([]DeviceManagementCompliancePolicyable) {
    val, err := m.GetBackingStore().Get("compliancePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementCompliancePolicyable)
    }
    return nil
}
// GetComplianceSettings gets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("complianceSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDefinitionable)
    }
    return nil
}
// GetConditionalAccessSettings gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable) {
    val, err := m.GetBackingStore().Get("conditionalAccessSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPremisesConditionalAccessSettingsable)
    }
    return nil
}
// GetConfigManagerCollections gets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) GetConfigManagerCollections()([]ConfigManagerCollectionable) {
    val, err := m.GetBackingStore().Get("configManagerCollections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConfigManagerCollectionable)
    }
    return nil
}
// GetConfigurationCategories gets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) GetConfigurationCategories()([]DeviceManagementConfigurationCategoryable) {
    val, err := m.GetBackingStore().Get("configurationCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationCategoryable)
    }
    return nil
}
// GetConfigurationPolicies gets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) GetConfigurationPolicies()([]DeviceManagementConfigurationPolicyable) {
    val, err := m.GetBackingStore().Get("configurationPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationPolicyable)
    }
    return nil
}
// GetConfigurationPolicyTemplates gets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplateable) {
    val, err := m.GetBackingStore().Get("configurationPolicyTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationPolicyTemplateable)
    }
    return nil
}
// GetConfigurationSettings gets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("configurationSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDefinitionable)
    }
    return nil
}
// GetConnectorStatus gets the connectorStatus property value. The list of connector status for the tenant.
func (m *DeviceManagement) GetConnectorStatus()([]ConnectorStatusDetailsable) {
    val, err := m.GetBackingStore().Get("connectorStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConnectorStatusDetailsable)
    }
    return nil
}
// GetDataProcessorServiceForWindowsFeaturesOnboarding gets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) GetDataProcessorServiceForWindowsFeaturesOnboarding()(DataProcessorServiceForWindowsFeaturesOnboardingable) {
    val, err := m.GetBackingStore().Get("dataProcessorServiceForWindowsFeaturesOnboarding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DataProcessorServiceForWindowsFeaturesOnboardingable)
    }
    return nil
}
// GetDataSharingConsents gets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) GetDataSharingConsents()([]DataSharingConsentable) {
    val, err := m.GetBackingStore().Get("dataSharingConsents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DataSharingConsentable)
    }
    return nil
}
// GetDepOnboardingSettings gets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) GetDepOnboardingSettings()([]DepOnboardingSettingable) {
    val, err := m.GetBackingStore().Get("depOnboardingSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DepOnboardingSettingable)
    }
    return nil
}
// GetDerivedCredentials gets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettingsable) {
    val, err := m.GetBackingStore().Get("derivedCredentials")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementDerivedCredentialSettingsable)
    }
    return nil
}
// GetDetectedApps gets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) GetDetectedApps()([]DetectedAppable) {
    val, err := m.GetBackingStore().Get("detectedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DetectedAppable)
    }
    return nil
}
// GetDeviceCategories gets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategoryable) {
    val, err := m.GetBackingStore().Get("deviceCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCategoryable)
    }
    return nil
}
// GetDeviceCompliancePolicies gets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCompliancePolicyable)
    }
    return nil
}
// GetDeviceCompliancePolicyDeviceStateSummary gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicyDeviceStateSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceCompliancePolicyDeviceStateSummaryable)
    }
    return nil
}
// GetDeviceCompliancePolicySettingStateSummaries gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicySettingStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCompliancePolicySettingStateSummaryable)
    }
    return nil
}
// GetDeviceComplianceReportSummarizationDateTime gets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("deviceComplianceReportSummarizationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceComplianceScripts gets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceComplianceScripts()([]DeviceComplianceScriptable) {
    val, err := m.GetBackingStore().Get("deviceComplianceScripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceComplianceScriptable)
    }
    return nil
}
// GetDeviceConfigurationConflictSummary gets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummaryable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationConflictSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceConfigurationConflictSummaryable)
    }
    return nil
}
// GetDeviceConfigurationDeviceStateSummaries gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationDeviceStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceConfigurationDeviceStateSummaryable)
    }
    return nil
}
// GetDeviceConfigurationRestrictedAppsViolations gets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolationable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationRestrictedAppsViolations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RestrictedAppsViolationable)
    }
    return nil
}
// GetDeviceConfigurations gets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfigurationable) {
    val, err := m.GetBackingStore().Get("deviceConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceConfigurationable)
    }
    return nil
}
// GetDeviceConfigurationsAllManagedDeviceCertificateStates gets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationsAllManagedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedAllDeviceCertificateStateable)
    }
    return nil
}
// GetDeviceConfigurationUserStateSummaries gets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationUserStateSummaries()(DeviceConfigurationUserStateSummaryable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationUserStateSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceConfigurationUserStateSummaryable)
    }
    return nil
}
// GetDeviceCustomAttributeShellScripts gets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScriptable) {
    val, err := m.GetBackingStore().Get("deviceCustomAttributeShellScripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCustomAttributeShellScriptable)
    }
    return nil
}
// GetDeviceEnrollmentConfigurations gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable) {
    val, err := m.GetBackingStore().Get("deviceEnrollmentConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceEnrollmentConfigurationable)
    }
    return nil
}
// GetDeviceHealthScripts gets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceHealthScripts()([]DeviceHealthScriptable) {
    val, err := m.GetBackingStore().Get("deviceHealthScripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptable)
    }
    return nil
}
// GetDeviceManagementPartners gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartnerable) {
    val, err := m.GetBackingStore().Get("deviceManagementPartners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementPartnerable)
    }
    return nil
}
// GetDeviceManagementScripts gets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceManagementScripts()([]DeviceManagementScriptable) {
    val, err := m.GetBackingStore().Get("deviceManagementScripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementScriptable)
    }
    return nil
}
// GetDeviceProtectionOverview gets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) GetDeviceProtectionOverview()(DeviceProtectionOverviewable) {
    val, err := m.GetBackingStore().Get("deviceProtectionOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceProtectionOverviewable)
    }
    return nil
}
// GetDeviceShellScripts gets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) GetDeviceShellScripts()([]DeviceShellScriptable) {
    val, err := m.GetBackingStore().Get("deviceShellScripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceShellScriptable)
    }
    return nil
}
// GetDomainJoinConnectors gets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnectorable) {
    val, err := m.GetBackingStore().Get("domainJoinConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementDomainJoinConnectorable)
    }
    return nil
}
// GetEmbeddedSIMActivationCodePools gets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePoolable) {
    val, err := m.GetBackingStore().Get("embeddedSIMActivationCodePools")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EmbeddedSIMActivationCodePoolable)
    }
    return nil
}
// GetExchangeConnectors gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnectorable) {
    val, err := m.GetBackingStore().Get("exchangeConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementExchangeConnectorable)
    }
    return nil
}
// GetExchangeOnPremisesPolicies gets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicyable) {
    val, err := m.GetBackingStore().Get("exchangeOnPremisesPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementExchangeOnPremisesPolicyable)
    }
    return nil
}
// GetExchangeOnPremisesPolicy gets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) GetExchangeOnPremisesPolicy()(DeviceManagementExchangeOnPremisesPolicyable) {
    val, err := m.GetBackingStore().Get("exchangeOnPremisesPolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementExchangeOnPremisesPolicyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountMoveCompletionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountMoveCompletionDateTime(val)
        }
        return nil
    }
    res["adminConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminConsentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsent(val.(AdminConsentable))
        }
        return nil
    }
    res["advancedThreatProtectionOnboardingStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdvancedThreatProtectionOnboardingStateSummary(val.(AdvancedThreatProtectionOnboardingStateSummaryable))
        }
        return nil
    }
    res["androidDeviceOwnerEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerEnrollmentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerEnrollmentProfileable)
                }
            }
            m.SetAndroidDeviceOwnerEnrollmentProfiles(res)
        }
        return nil
    }
    res["androidForWorkAppConfigurationSchemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkAppConfigurationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkAppConfigurationSchemaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidForWorkAppConfigurationSchemaable)
                }
            }
            m.SetAndroidForWorkAppConfigurationSchemas(res)
        }
        return nil
    }
    res["androidForWorkEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkEnrollmentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidForWorkEnrollmentProfileable)
                }
            }
            m.SetAndroidForWorkEnrollmentProfiles(res)
        }
        return nil
    }
    res["androidForWorkSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidForWorkSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidForWorkSettings(val.(AndroidForWorkSettingsable))
        }
        return nil
    }
    res["androidManagedStoreAccountEnterpriseSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidManagedStoreAccountEnterpriseSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidManagedStoreAccountEnterpriseSettings(val.(AndroidManagedStoreAccountEnterpriseSettingsable))
        }
        return nil
    }
    res["androidManagedStoreAppConfigurationSchemas"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedStoreAppConfigurationSchemaable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidManagedStoreAppConfigurationSchemaable)
                }
            }
            m.SetAndroidManagedStoreAppConfigurationSchemas(res)
        }
        return nil
    }
    res["applePushNotificationCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplePushNotificationCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePushNotificationCertificate(val.(ApplePushNotificationCertificateable))
        }
        return nil
    }
    res["appleUserInitiatedEnrollmentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleUserInitiatedEnrollmentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AppleUserInitiatedEnrollmentProfileable)
                }
            }
            m.SetAppleUserInitiatedEnrollmentProfiles(res)
        }
        return nil
    }
    res["assignmentFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementAssignmentFilterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAndAppManagementAssignmentFilterable)
                }
            }
            m.SetAssignmentFilters(res)
        }
        return nil
    }
    res["auditEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuditEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuditEventable)
                }
            }
            m.SetAuditEvents(res)
        }
        return nil
    }
    res["autopilotEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementAutopilotEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementAutopilotEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementAutopilotEventable)
                }
            }
            m.SetAutopilotEvents(res)
        }
        return nil
    }
    res["cartToClassAssociations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCartToClassAssociationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CartToClassAssociationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CartToClassAssociationable)
                }
            }
            m.SetCartToClassAssociations(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingCategoryable)
                }
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["certificateConnectorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateConnectorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateConnectorDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CertificateConnectorDetailsable)
                }
            }
            m.SetCertificateConnectorDetails(res)
        }
        return nil
    }
    res["chromeOSOnboardingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChromeOSOnboardingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChromeOSOnboardingSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ChromeOSOnboardingSettingsable)
                }
            }
            m.SetChromeOSOnboardingSettings(res)
        }
        return nil
    }
    res["cloudPCConnectivityIssues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPCConnectivityIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPCConnectivityIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPCConnectivityIssueable)
                }
            }
            m.SetCloudPCConnectivityIssues(res)
        }
        return nil
    }
    res["comanagedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceable)
                }
            }
            m.SetComanagedDevices(res)
        }
        return nil
    }
    res["comanagementEligibleDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComanagementEligibleDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComanagementEligibleDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ComanagementEligibleDeviceable)
                }
            }
            m.SetComanagementEligibleDevices(res)
        }
        return nil
    }
    res["complianceCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationCategoryable)
                }
            }
            m.SetComplianceCategories(res)
        }
        return nil
    }
    res["complianceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ComplianceManagementPartnerable)
                }
            }
            m.SetComplianceManagementPartners(res)
        }
        return nil
    }
    res["compliancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementCompliancePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementCompliancePolicyable)
                }
            }
            m.SetCompliancePolicies(res)
        }
        return nil
    }
    res["complianceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
                }
            }
            m.SetComplianceSettings(res)
        }
        return nil
    }
    res["conditionalAccessSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPremisesConditionalAccessSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessSettings(val.(OnPremisesConditionalAccessSettingsable))
        }
        return nil
    }
    res["configManagerCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConfigManagerCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConfigManagerCollectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConfigManagerCollectionable)
                }
            }
            m.SetConfigManagerCollections(res)
        }
        return nil
    }
    res["configurationCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationCategoryable)
                }
            }
            m.SetConfigurationCategories(res)
        }
        return nil
    }
    res["configurationPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationPolicyable)
                }
            }
            m.SetConfigurationPolicies(res)
        }
        return nil
    }
    res["configurationPolicyTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationPolicyTemplateable)
                }
            }
            m.SetConfigurationPolicyTemplates(res)
        }
        return nil
    }
    res["configurationSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
                }
            }
            m.SetConfigurationSettings(res)
        }
        return nil
    }
    res["connectorStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectorStatusDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConnectorStatusDetailsable)
                }
            }
            m.SetConnectorStatus(res)
        }
        return nil
    }
    res["dataProcessorServiceForWindowsFeaturesOnboarding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProcessorServiceForWindowsFeaturesOnboarding(val.(DataProcessorServiceForWindowsFeaturesOnboardingable))
        }
        return nil
    }
    res["dataSharingConsents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDataSharingConsentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DataSharingConsentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DataSharingConsentable)
                }
            }
            m.SetDataSharingConsents(res)
        }
        return nil
    }
    res["depOnboardingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDepOnboardingSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DepOnboardingSettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DepOnboardingSettingable)
                }
            }
            m.SetDepOnboardingSettings(res)
        }
        return nil
    }
    res["derivedCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementDerivedCredentialSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementDerivedCredentialSettingsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementDerivedCredentialSettingsable)
                }
            }
            m.SetDerivedCredentials(res)
        }
        return nil
    }
    res["detectedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedAppable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DetectedAppable)
                }
            }
            m.SetDetectedApps(res)
        }
        return nil
    }
    res["deviceCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCategoryable)
                }
            }
            m.SetDeviceCategories(res)
        }
        return nil
    }
    res["deviceCompliancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicyable)
                }
            }
            m.SetDeviceCompliancePolicies(res)
        }
        return nil
    }
    res["deviceCompliancePolicyDeviceStateSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCompliancePolicyDeviceStateSummary(val.(DeviceCompliancePolicyDeviceStateSummaryable))
        }
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicySettingStateSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicySettingStateSummaryable)
                }
            }
            m.SetDeviceCompliancePolicySettingStateSummaries(res)
        }
        return nil
    }
    res["deviceComplianceReportSummarizationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceReportSummarizationDateTime(val)
        }
        return nil
    }
    res["deviceComplianceScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceComplianceScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceComplianceScriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceComplianceScriptable)
                }
            }
            m.SetDeviceComplianceScripts(res)
        }
        return nil
    }
    res["deviceConfigurationConflictSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationConflictSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationConflictSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationConflictSummaryable)
                }
            }
            m.SetDeviceConfigurationConflictSummary(res)
        }
        return nil
    }
    res["deviceConfigurationDeviceStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationDeviceStateSummaries(val.(DeviceConfigurationDeviceStateSummaryable))
        }
        return nil
    }
    res["deviceConfigurationRestrictedAppsViolations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRestrictedAppsViolationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RestrictedAppsViolationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RestrictedAppsViolationable)
                }
            }
            m.SetDeviceConfigurationRestrictedAppsViolations(res)
        }
        return nil
    }
    res["deviceConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationable)
                }
            }
            m.SetDeviceConfigurations(res)
        }
        return nil
    }
    res["deviceConfigurationsAllManagedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAllDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAllDeviceCertificateStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedAllDeviceCertificateStateable)
                }
            }
            m.SetDeviceConfigurationsAllManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["deviceConfigurationUserStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationUserStateSummaries(val.(DeviceConfigurationUserStateSummaryable))
        }
        return nil
    }
    res["deviceCustomAttributeShellScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCustomAttributeShellScriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCustomAttributeShellScriptable)
                }
            }
            m.SetDeviceCustomAttributeShellScripts(res)
        }
        return nil
    }
    res["deviceEnrollmentConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceEnrollmentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceEnrollmentConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceEnrollmentConfigurationable)
                }
            }
            m.SetDeviceEnrollmentConfigurations(res)
        }
        return nil
    }
    res["deviceHealthScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptable)
                }
            }
            m.SetDeviceHealthScripts(res)
        }
        return nil
    }
    res["deviceManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementPartnerable)
                }
            }
            m.SetDeviceManagementPartners(res)
        }
        return nil
    }
    res["deviceManagementScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementScriptable)
                }
            }
            m.SetDeviceManagementScripts(res)
        }
        return nil
    }
    res["deviceProtectionOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceProtectionOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceProtectionOverview(val.(DeviceProtectionOverviewable))
        }
        return nil
    }
    res["deviceShellScripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceShellScriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceShellScriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceShellScriptable)
                }
            }
            m.SetDeviceShellScripts(res)
        }
        return nil
    }
    res["domainJoinConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementDomainJoinConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementDomainJoinConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementDomainJoinConnectorable)
                }
            }
            m.SetDomainJoinConnectors(res)
        }
        return nil
    }
    res["embeddedSIMActivationCodePools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmbeddedSIMActivationCodePoolable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EmbeddedSIMActivationCodePoolable)
                }
            }
            m.SetEmbeddedSIMActivationCodePools(res)
        }
        return nil
    }
    res["exchangeConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementExchangeConnectorable)
                }
            }
            m.SetExchangeConnectors(res)
        }
        return nil
    }
    res["exchangeOnPremisesPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementExchangeOnPremisesPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementExchangeOnPremisesPolicyable)
                }
            }
            m.SetExchangeOnPremisesPolicies(res)
        }
        return nil
    }
    res["exchangeOnPremisesPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementExchangeOnPremisesPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeOnPremisesPolicy(val.(DeviceManagementExchangeOnPremisesPolicyable))
        }
        return nil
    }
    res["groupPolicyCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyCategoryable)
                }
            }
            m.SetGroupPolicyCategories(res)
        }
        return nil
    }
    res["groupPolicyConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyConfigurationable)
                }
            }
            m.SetGroupPolicyConfigurations(res)
        }
        return nil
    }
    res["groupPolicyDefinitionFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionFileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyDefinitionFileable)
                }
            }
            m.SetGroupPolicyDefinitionFiles(res)
        }
        return nil
    }
    res["groupPolicyDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyDefinitionable)
                }
            }
            m.SetGroupPolicyDefinitions(res)
        }
        return nil
    }
    res["groupPolicyMigrationReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyMigrationReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyMigrationReportable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyMigrationReportable)
                }
            }
            m.SetGroupPolicyMigrationReports(res)
        }
        return nil
    }
    res["groupPolicyObjectFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyObjectFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyObjectFileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyObjectFileable)
                }
            }
            m.SetGroupPolicyObjectFiles(res)
        }
        return nil
    }
    res["groupPolicyUploadedDefinitionFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicyUploadedDefinitionFileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicyUploadedDefinitionFileable)
                }
            }
            m.SetGroupPolicyUploadedDefinitionFiles(res)
        }
        return nil
    }
    res["importedDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImportedDeviceIdentityable)
                }
            }
            m.SetImportedDeviceIdentities(res)
        }
        return nil
    }
    res["importedWindowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ImportedWindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ImportedWindowsAutopilotDeviceIdentityable)
                }
            }
            m.SetImportedWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["intents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementIntentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementIntentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementIntentable)
                }
            }
            m.SetIntents(res)
        }
        return nil
    }
    res["intuneAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneAccountId(val)
        }
        return nil
    }
    res["intuneBrand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIntuneBrandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneBrand(val.(IntuneBrandable))
        }
        return nil
    }
    res["intuneBrandingProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIntuneBrandingProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IntuneBrandingProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IntuneBrandingProfileable)
                }
            }
            m.SetIntuneBrandingProfiles(res)
        }
        return nil
    }
    res["iosUpdateStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosUpdateDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosUpdateDeviceStatusable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosUpdateDeviceStatusable)
                }
            }
            m.SetIosUpdateStatuses(res)
        }
        return nil
    }
    res["lastReportAggregationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportAggregationDateTime(val)
        }
        return nil
    }
    res["legacyPcManangementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegacyPcManangementEnabled(val)
        }
        return nil
    }
    res["macOSSoftwareUpdateAccountSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateAccountSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MacOSSoftwareUpdateAccountSummaryable)
                }
            }
            m.SetMacOSSoftwareUpdateAccountSummaries(res)
        }
        return nil
    }
    res["managedDeviceCleanupSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceCleanupSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceCleanupSettings(val.(ManagedDeviceCleanupSettingsable))
        }
        return nil
    }
    res["managedDeviceEncryptionStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceEncryptionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceEncryptionStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceEncryptionStateable)
                }
            }
            m.SetManagedDeviceEncryptionStates(res)
        }
        return nil
    }
    res["managedDeviceOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOverview(val.(ManagedDeviceOverviewable))
        }
        return nil
    }
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceable)
                }
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["maximumDepTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumDepTokens(val)
        }
        return nil
    }
    res["microsoftTunnelConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTunnelConfigurationable)
                }
            }
            m.SetMicrosoftTunnelConfigurations(res)
        }
        return nil
    }
    res["microsoftTunnelHealthThresholds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelHealthThresholdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTunnelHealthThresholdable)
                }
            }
            m.SetMicrosoftTunnelHealthThresholds(res)
        }
        return nil
    }
    res["microsoftTunnelServerLogCollectionResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelServerLogCollectionResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelServerLogCollectionResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTunnelServerLogCollectionResponseable)
                }
            }
            m.SetMicrosoftTunnelServerLogCollectionResponses(res)
        }
        return nil
    }
    res["microsoftTunnelSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelSiteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTunnelSiteable)
                }
            }
            m.SetMicrosoftTunnelSites(res)
        }
        return nil
    }
    res["mobileAppTroubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppTroubleshootingEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileAppTroubleshootingEventable)
                }
            }
            m.SetMobileAppTroubleshootingEvents(res)
        }
        return nil
    }
    res["mobileThreatDefenseConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileThreatDefenseConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileThreatDefenseConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MobileThreatDefenseConnectorable)
                }
            }
            m.SetMobileThreatDefenseConnectors(res)
        }
        return nil
    }
    res["ndesConnectors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNdesConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NdesConnectorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NdesConnectorable)
                }
            }
            m.SetNdesConnectors(res)
        }
        return nil
    }
    res["notificationMessageTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNotificationMessageTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NotificationMessageTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NotificationMessageTemplateable)
                }
            }
            m.SetNotificationMessageTemplates(res)
        }
        return nil
    }
    res["privilegeManagementElevations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegeManagementElevationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegeManagementElevationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PrivilegeManagementElevationable)
                }
            }
            m.SetPrivilegeManagementElevations(res)
        }
        return nil
    }
    res["remoteActionAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteActionAuditFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteActionAuditable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RemoteActionAuditable)
                }
            }
            m.SetRemoteActionAudits(res)
        }
        return nil
    }
    res["remoteAssistancePartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteAssistancePartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteAssistancePartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RemoteAssistancePartnerable)
                }
            }
            m.SetRemoteAssistancePartners(res)
        }
        return nil
    }
    res["remoteAssistanceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteAssistanceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSettings(val.(RemoteAssistanceSettingsable))
        }
        return nil
    }
    res["reports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementReportsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReports(val.(DeviceManagementReportsable))
        }
        return nil
    }
    res["resourceAccessProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementResourceAccessProfileBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementResourceAccessProfileBaseable)
                }
            }
            m.SetResourceAccessProfiles(res)
        }
        return nil
    }
    res["resourceOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceOperationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ResourceOperationable)
                }
            }
            m.SetResourceOperations(res)
        }
        return nil
    }
    res["reusablePolicySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementReusablePolicySettingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementReusablePolicySettingable)
                }
            }
            m.SetReusablePolicySettings(res)
        }
        return nil
    }
    res["reusableSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
                }
            }
            m.SetReusableSettings(res)
        }
        return nil
    }
    res["roleAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementRoleAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceAndAppManagementRoleAssignmentable)
                }
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["roleDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RoleDefinitionable)
                }
            }
            m.SetRoleDefinitions(res)
        }
        return nil
    }
    res["roleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRoleScopeTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RoleScopeTagable)
                }
            }
            m.SetRoleScopeTags(res)
        }
        return nil
    }
    res["serviceNowConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceNowConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceNowConnectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ServiceNowConnectionable)
                }
            }
            m.SetServiceNowConnections(res)
        }
        return nil
    }
    res["settingDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingDefinitionable)
                }
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(DeviceManagementSettingsable))
        }
        return nil
    }
    res["softwareUpdateStatusSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateStatusSummary(val.(SoftwareUpdateStatusSummaryable))
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptions(val.(*DeviceManagementSubscriptions))
        }
        return nil
    }
    res["subscriptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionState(val.(*DeviceManagementSubscriptionState))
        }
        return nil
    }
    res["telecomExpenseManagementPartners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTelecomExpenseManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TelecomExpenseManagementPartnerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TelecomExpenseManagementPartnerable)
                }
            }
            m.SetTelecomExpenseManagementPartners(res)
        }
        return nil
    }
    res["templateInsights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateInsightsDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateInsightsDefinitionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTemplateInsightsDefinitionable)
                }
            }
            m.SetTemplateInsights(res)
        }
        return nil
    }
    res["templates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTemplateable)
                }
            }
            m.SetTemplates(res)
        }
        return nil
    }
    res["templateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationSettingTemplateable)
                }
            }
            m.SetTemplateSettings(res)
        }
        return nil
    }
    res["tenantAttachRBAC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAttachRBACFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttachRBAC(val.(TenantAttachRBACable))
        }
        return nil
    }
    res["termsAndConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermsAndConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TermsAndConditionsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TermsAndConditionsable)
                }
            }
            m.SetTermsAndConditions(res)
        }
        return nil
    }
    res["troubleshootingEvents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTroubleshootingEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTroubleshootingEventable)
                }
            }
            m.SetTroubleshootingEvents(res)
        }
        return nil
    }
    res["unlicensedAdminstratorsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnlicensedAdminstratorsEnabled(val)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomaly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAnomalyable)
                }
            }
            m.SetUserExperienceAnalyticsAnomaly(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomalyCorrelationGroupOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyCorrelationGroupOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
                }
            }
            m.SetUserExperienceAnalyticsAnomalyCorrelationGroupOverview(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomalyDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAnomalyDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAnomalyDeviceable)
                }
            }
            m.SetUserExperienceAnalyticsAnomalyDevice(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAnomalySeverityOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsAnomalySeverityOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsAnomalySeverityOverview(val.(UserExperienceAnalyticsAnomalySeverityOverviewable))
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthApplicationPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthApplicationPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDeviceModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDeviceModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthDevicePerformanceDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOSVersionPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsAppHealthOSVersionPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsAppHealthOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsAppHealthOverview(val.(UserExperienceAnalyticsCategoryable))
        }
        return nil
    }
    res["userExperienceAnalyticsBaselines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBaselineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBaselineable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBaselineable)
                }
            }
            m.SetUserExperienceAnalyticsBaselines(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthAppImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthAppImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthAppImpactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthAppImpactable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthAppImpact(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthCapacityDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsBatteryHealthCapacityDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsBatteryHealthCapacityDetails(val.(UserExperienceAnalyticsBatteryHealthCapacityDetailsable))
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceAppImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthDevicePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthModelPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthModelPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthModelPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthOsPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsBatteryHealthOsPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsBatteryHealthOsPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsBatteryHealthOsPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsBatteryHealthOsPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsBatteryHealthRuntimeDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsBatteryHealthRuntimeDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(val.(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable))
        }
        return nil
    }
    res["userExperienceAnalyticsCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsCategoryable)
                }
            }
            m.SetUserExperienceAnalyticsCategories(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceMetricHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsMetricHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceMetricHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDevicePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDevicePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDevicePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsDevicePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsDeviceScope(val.(UserExperienceAnalyticsDeviceScopeable))
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceScopeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceScopeable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceScopes(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceScoresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceScoresable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcesses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcesses(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceStartupProcessPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceStartupProcessPerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDevicesWithoutCloudIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceWithoutCloudIdentityable)
                }
            }
            m.SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(res)
        }
        return nil
    }
    res["userExperienceAnalyticsDeviceTimelineEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsDeviceTimelineEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsDeviceTimelineEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsDeviceTimelineEventable)
                }
            }
            m.SetUserExperienceAnalyticsDeviceTimelineEvent(res)
        }
        return nil
    }
    res["userExperienceAnalyticsImpactingProcess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsImpactingProcessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsImpactingProcessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsImpactingProcessable)
                }
            }
            m.SetUserExperienceAnalyticsImpactingProcess(res)
        }
        return nil
    }
    res["userExperienceAnalyticsMetricHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsMetricHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsMetricHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsMetricHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsModelScores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsModelScoresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsModelScoresable)
                }
            }
            m.SetUserExperienceAnalyticsModelScores(res)
        }
        return nil
    }
    res["userExperienceAnalyticsNotAutopilotReadyDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsNotAutopilotReadyDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsNotAutopilotReadyDeviceable)
                }
            }
            m.SetUserExperienceAnalyticsNotAutopilotReadyDevice(res)
        }
        return nil
    }
    res["userExperienceAnalyticsOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsOverview(val.(UserExperienceAnalyticsOverviewable))
        }
        return nil
    }
    res["userExperienceAnalyticsRemoteConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsRemoteConnectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsRemoteConnectionable)
                }
            }
            m.SetUserExperienceAnalyticsRemoteConnection(res)
        }
        return nil
    }
    res["userExperienceAnalyticsResourcePerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsResourcePerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsResourcePerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsResourcePerformance(res)
        }
        return nil
    }
    res["userExperienceAnalyticsScoreHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsScoreHistoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsScoreHistoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsScoreHistoryable)
                }
            }
            m.SetUserExperienceAnalyticsScoreHistory(res)
        }
        return nil
    }
    res["userExperienceAnalyticsSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsSettings(val.(UserExperienceAnalyticsSettingsable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable))
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereMetrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereMetricable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereMetricable)
                }
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereMetrics(res)
        }
        return nil
    }
    res["userExperienceAnalyticsWorkFromAnywhereModelPerformance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
                }
            }
            m.SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(res)
        }
        return nil
    }
    res["userPfxCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserPFXCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserPFXCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UserPFXCertificateable)
                }
            }
            m.SetUserPfxCertificates(res)
        }
        return nil
    }
    res["virtualEndpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualEndpoint(val.(VirtualEndpointable))
        }
        return nil
    }
    res["windowsAutopilotDeploymentProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeploymentProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsAutopilotDeploymentProfileable)
                }
            }
            m.SetWindowsAutopilotDeploymentProfiles(res)
        }
        return nil
    }
    res["windowsAutopilotDeviceIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsAutopilotDeviceIdentityable)
                }
            }
            m.SetWindowsAutopilotDeviceIdentities(res)
        }
        return nil
    }
    res["windowsAutopilotSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsAutopilotSettings(val.(WindowsAutopilotSettingsable))
        }
        return nil
    }
    res["windowsDriverUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDriverUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDriverUpdateProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsDriverUpdateProfileable)
                }
            }
            m.SetWindowsDriverUpdateProfiles(res)
        }
        return nil
    }
    res["windowsFeatureUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsFeatureUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsFeatureUpdateProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsFeatureUpdateProfileable)
                }
            }
            m.SetWindowsFeatureUpdateProfiles(res)
        }
        return nil
    }
    res["windowsInformationProtectionAppLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionAppLearningSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionAppLearningSummaryable)
                }
            }
            m.SetWindowsInformationProtectionAppLearningSummaries(res)
        }
        return nil
    }
    res["windowsInformationProtectionNetworkLearningSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionNetworkLearningSummaryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsInformationProtectionNetworkLearningSummaryable)
                }
            }
            m.SetWindowsInformationProtectionNetworkLearningSummaries(res)
        }
        return nil
    }
    res["windowsMalwareInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsMalwareInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsMalwareInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsMalwareInformationable)
                }
            }
            m.SetWindowsMalwareInformation(res)
        }
        return nil
    }
    res["windowsMalwareOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMalwareOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMalwareOverview(val.(WindowsMalwareOverviewable))
        }
        return nil
    }
    res["windowsQualityUpdateProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsQualityUpdateProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsQualityUpdateProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsQualityUpdateProfileable)
                }
            }
            m.SetWindowsQualityUpdateProfiles(res)
        }
        return nil
    }
    res["windowsUpdateCatalogItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsUpdateCatalogItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsUpdateCatalogItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsUpdateCatalogItemable)
                }
            }
            m.SetWindowsUpdateCatalogItems(res)
        }
        return nil
    }
    res["zebraFotaArtifacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateZebraFotaArtifactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ZebraFotaArtifactable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ZebraFotaArtifactable)
                }
            }
            m.SetZebraFotaArtifacts(res)
        }
        return nil
    }
    res["zebraFotaConnector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateZebraFotaConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZebraFotaConnector(val.(ZebraFotaConnectorable))
        }
        return nil
    }
    res["zebraFotaDeployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateZebraFotaDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ZebraFotaDeploymentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ZebraFotaDeploymentable)
                }
            }
            m.SetZebraFotaDeployments(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyCategories gets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) GetGroupPolicyCategories()([]GroupPolicyCategoryable) {
    val, err := m.GetBackingStore().Get("groupPolicyCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyCategoryable)
    }
    return nil
}
// GetGroupPolicyConfigurations gets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) GetGroupPolicyConfigurations()([]GroupPolicyConfigurationable) {
    val, err := m.GetBackingStore().Get("groupPolicyConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyConfigurationable)
    }
    return nil
}
// GetGroupPolicyDefinitionFiles gets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFileable) {
    val, err := m.GetBackingStore().Get("groupPolicyDefinitionFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyDefinitionFileable)
    }
    return nil
}
// GetGroupPolicyDefinitions gets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) GetGroupPolicyDefinitions()([]GroupPolicyDefinitionable) {
    val, err := m.GetBackingStore().Get("groupPolicyDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyDefinitionable)
    }
    return nil
}
// GetGroupPolicyMigrationReports gets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReportable) {
    val, err := m.GetBackingStore().Get("groupPolicyMigrationReports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyMigrationReportable)
    }
    return nil
}
// GetGroupPolicyObjectFiles gets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) GetGroupPolicyObjectFiles()([]GroupPolicyObjectFileable) {
    val, err := m.GetBackingStore().Get("groupPolicyObjectFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyObjectFileable)
    }
    return nil
}
// GetGroupPolicyUploadedDefinitionFiles gets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFileable) {
    val, err := m.GetBackingStore().Get("groupPolicyUploadedDefinitionFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicyUploadedDefinitionFileable)
    }
    return nil
}
// GetImportedDeviceIdentities gets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) GetImportedDeviceIdentities()([]ImportedDeviceIdentityable) {
    val, err := m.GetBackingStore().Get("importedDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImportedDeviceIdentityable)
    }
    return nil
}
// GetImportedWindowsAutopilotDeviceIdentities gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable) {
    val, err := m.GetBackingStore().Get("importedWindowsAutopilotDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ImportedWindowsAutopilotDeviceIdentityable)
    }
    return nil
}
// GetIntents gets the intents property value. The device management intents
func (m *DeviceManagement) GetIntents()([]DeviceManagementIntentable) {
    val, err := m.GetBackingStore().Get("intents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementIntentable)
    }
    return nil
}
// GetIntuneAccountId gets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("intuneAccountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetIntuneBrand gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) GetIntuneBrand()(IntuneBrandable) {
    val, err := m.GetBackingStore().Get("intuneBrand")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IntuneBrandable)
    }
    return nil
}
// GetIntuneBrandingProfiles gets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) GetIntuneBrandingProfiles()([]IntuneBrandingProfileable) {
    val, err := m.GetBackingStore().Get("intuneBrandingProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IntuneBrandingProfileable)
    }
    return nil
}
// GetIosUpdateStatuses gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatusable) {
    val, err := m.GetBackingStore().Get("iosUpdateStatuses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosUpdateDeviceStatusable)
    }
    return nil
}
// GetLastReportAggregationDateTime gets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastReportAggregationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLegacyPcManangementEnabled gets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) GetLegacyPcManangementEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("legacyPcManangementEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMacOSSoftwareUpdateAccountSummaries gets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummaryable) {
    val, err := m.GetBackingStore().Get("macOSSoftwareUpdateAccountSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSSoftwareUpdateAccountSummaryable)
    }
    return nil
}
// GetManagedDeviceCleanupSettings gets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) GetManagedDeviceCleanupSettings()(ManagedDeviceCleanupSettingsable) {
    val, err := m.GetBackingStore().Get("managedDeviceCleanupSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedDeviceCleanupSettingsable)
    }
    return nil
}
// GetManagedDeviceEncryptionStates gets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceEncryptionStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceEncryptionStateable)
    }
    return nil
}
// GetManagedDeviceOverview gets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) GetManagedDeviceOverview()(ManagedDeviceOverviewable) {
    val, err := m.GetBackingStore().Get("managedDeviceOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ManagedDeviceOverviewable)
    }
    return nil
}
// GetManagedDevices gets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) GetManagedDevices()([]ManagedDeviceable) {
    val, err := m.GetBackingStore().Get("managedDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceable)
    }
    return nil
}
// GetMaximumDepTokens gets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) GetMaximumDepTokens()(*int32) {
    val, err := m.GetBackingStore().Get("maximumDepTokens")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMicrosoftTunnelConfigurations gets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfigurationable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelConfigurations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MicrosoftTunnelConfigurationable)
    }
    return nil
}
// GetMicrosoftTunnelHealthThresholds gets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThresholdable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelHealthThresholds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MicrosoftTunnelHealthThresholdable)
    }
    return nil
}
// GetMicrosoftTunnelServerLogCollectionResponses gets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponseable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelServerLogCollectionResponses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MicrosoftTunnelServerLogCollectionResponseable)
    }
    return nil
}
// GetMicrosoftTunnelSites gets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) GetMicrosoftTunnelSites()([]MicrosoftTunnelSiteable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelSites")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MicrosoftTunnelSiteable)
    }
    return nil
}
// GetMobileAppTroubleshootingEvents gets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable) {
    val, err := m.GetBackingStore().Get("mobileAppTroubleshootingEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileAppTroubleshootingEventable)
    }
    return nil
}
// GetMobileThreatDefenseConnectors gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable) {
    val, err := m.GetBackingStore().Get("mobileThreatDefenseConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MobileThreatDefenseConnectorable)
    }
    return nil
}
// GetNdesConnectors gets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) GetNdesConnectors()([]NdesConnectorable) {
    val, err := m.GetBackingStore().Get("ndesConnectors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NdesConnectorable)
    }
    return nil
}
// GetNotificationMessageTemplates gets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplateable) {
    val, err := m.GetBackingStore().Get("notificationMessageTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NotificationMessageTemplateable)
    }
    return nil
}
// GetPrivilegeManagementElevations gets the privilegeManagementElevations property value. The endpoint privilege management elevation event entity contains elevation details.
func (m *DeviceManagement) GetPrivilegeManagementElevations()([]PrivilegeManagementElevationable) {
    val, err := m.GetBackingStore().Get("privilegeManagementElevations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PrivilegeManagementElevationable)
    }
    return nil
}
// GetRemoteActionAudits gets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) GetRemoteActionAudits()([]RemoteActionAuditable) {
    val, err := m.GetBackingStore().Get("remoteActionAudits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RemoteActionAuditable)
    }
    return nil
}
// GetRemoteAssistancePartners gets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartnerable) {
    val, err := m.GetBackingStore().Get("remoteAssistancePartners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RemoteAssistancePartnerable)
    }
    return nil
}
// GetRemoteAssistanceSettings gets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) GetRemoteAssistanceSettings()(RemoteAssistanceSettingsable) {
    val, err := m.GetBackingStore().Get("remoteAssistanceSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RemoteAssistanceSettingsable)
    }
    return nil
}
// GetReports gets the reports property value. Reports singleton
func (m *DeviceManagement) GetReports()(DeviceManagementReportsable) {
    val, err := m.GetBackingStore().Get("reports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementReportsable)
    }
    return nil
}
// GetResourceAccessProfiles gets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBaseable) {
    val, err := m.GetBackingStore().Get("resourceAccessProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementResourceAccessProfileBaseable)
    }
    return nil
}
// GetResourceOperations gets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperationable) {
    val, err := m.GetBackingStore().Get("resourceOperations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ResourceOperationable)
    }
    return nil
}
// GetReusablePolicySettings gets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) GetReusablePolicySettings()([]DeviceManagementReusablePolicySettingable) {
    val, err := m.GetBackingStore().Get("reusablePolicySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementReusablePolicySettingable)
    }
    return nil
}
// GetReusableSettings gets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) GetReusableSettings()([]DeviceManagementConfigurationSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("reusableSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingDefinitionable)
    }
    return nil
}
// GetRoleAssignments gets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable) {
    val, err := m.GetBackingStore().Get("roleAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAndAppManagementRoleAssignmentable)
    }
    return nil
}
// GetRoleDefinitions gets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinitionable) {
    val, err := m.GetBackingStore().Get("roleDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RoleDefinitionable)
    }
    return nil
}
// GetRoleScopeTags gets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) GetRoleScopeTags()([]RoleScopeTagable) {
    val, err := m.GetBackingStore().Get("roleScopeTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RoleScopeTagable)
    }
    return nil
}
// GetServiceNowConnections gets the serviceNowConnections property value. A list of ServiceNowConnections
func (m *DeviceManagement) GetServiceNowConnections()([]ServiceNowConnectionable) {
    val, err := m.GetBackingStore().Get("serviceNowConnections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ServiceNowConnectionable)
    }
    return nil
}
// GetSettingDefinitions gets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) GetSettingDefinitions()([]DeviceManagementSettingDefinitionable) {
    val, err := m.GetBackingStore().Get("settingDefinitions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingDefinitionable)
    }
    return nil
}
// GetSettings gets the settings property value. Account level settings.
func (m *DeviceManagement) GetSettings()(DeviceManagementSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementSettingsable)
    }
    return nil
}
// GetSoftwareUpdateStatusSummary gets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable) {
    val, err := m.GetBackingStore().Get("softwareUpdateStatusSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SoftwareUpdateStatusSummaryable)
    }
    return nil
}
// GetSubscriptions gets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) GetSubscriptions()(*DeviceManagementSubscriptions) {
    val, err := m.GetBackingStore().Get("subscriptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementSubscriptions)
    }
    return nil
}
// GetSubscriptionState gets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    val, err := m.GetBackingStore().Get("subscriptionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementSubscriptionState)
    }
    return nil
}
// GetTelecomExpenseManagementPartners gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable) {
    val, err := m.GetBackingStore().Get("telecomExpenseManagementPartners")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TelecomExpenseManagementPartnerable)
    }
    return nil
}
// GetTemplateInsights gets the templateInsights property value. List of setting insights in a template
func (m *DeviceManagement) GetTemplateInsights()([]DeviceManagementTemplateInsightsDefinitionable) {
    val, err := m.GetBackingStore().Get("templateInsights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTemplateInsightsDefinitionable)
    }
    return nil
}
// GetTemplates gets the templates property value. The available templates
func (m *DeviceManagement) GetTemplates()([]DeviceManagementTemplateable) {
    val, err := m.GetBackingStore().Get("templates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTemplateable)
    }
    return nil
}
// GetTemplateSettings gets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplateable) {
    val, err := m.GetBackingStore().Get("templateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingTemplateable)
    }
    return nil
}
// GetTenantAttachRBAC gets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) GetTenantAttachRBAC()(TenantAttachRBACable) {
    val, err := m.GetBackingStore().Get("tenantAttachRBAC")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TenantAttachRBACable)
    }
    return nil
}
// GetTermsAndConditions gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditionsable) {
    val, err := m.GetBackingStore().Get("termsAndConditions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TermsAndConditionsable)
    }
    return nil
}
// GetTroubleshootingEvents gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable) {
    val, err := m.GetBackingStore().Get("troubleshootingEvents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTroubleshootingEventable)
    }
    return nil
}
// GetUnlicensedAdminstratorsEnabled gets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) GetUnlicensedAdminstratorsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("unlicensedAdminstratorsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserExperienceAnalyticsAnomaly gets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomaly()([]UserExperienceAnalyticsAnomalyable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAnomaly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyable)
    }
    return nil
}
// GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview gets the userExperienceAnalyticsAnomalyCorrelationGroupOverview property value. The user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview()([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAnomalyCorrelationGroupOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
    }
    return nil
}
// GetUserExperienceAnalyticsAnomalyDevice gets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalyDevice()([]UserExperienceAnalyticsAnomalyDeviceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAnomalyDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAnomalyDeviceable)
    }
    return nil
}
// GetUserExperienceAnalyticsAnomalySeverityOverview gets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) GetUserExperienceAnalyticsAnomalySeverityOverview()(UserExperienceAnalyticsAnomalySeverityOverviewable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAnomalySeverityOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsAnomalySeverityOverviewable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformance gets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthApplicationPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthApplicationPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId gets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion gets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthDeviceModelPerformance gets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthDeviceModelPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthDevicePerformance gets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthDevicePerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthDevicePerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails gets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthDevicePerformanceDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthOSVersionPerformance gets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthOSVersionPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsAppHealthOverview gets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsAppHealthOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsCategoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsBaselines gets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBaselines")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBaselineable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthAppImpact gets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpactable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthAppImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthAppImpactable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthCapacityDetails gets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(UserExperienceAnalyticsBatteryHealthCapacityDetailsable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthCapacityDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsBatteryHealthCapacityDetailsable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact gets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthDeviceAppImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthDevicePerformance gets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthDevicePerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory gets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthModelPerformance gets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthModelPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthModelPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthOsPerformance gets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthOsPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsBatteryHealthOsPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsBatteryHealthRuntimeDetails gets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsBatteryHealthRuntimeDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)
    }
    return nil
}
// GetUserExperienceAnalyticsCategories gets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsCategoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceMetricHistory gets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceMetricHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsMetricHistoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsDevicePerformance gets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDevicePerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDevicePerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceScope gets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScope()(UserExperienceAnalyticsDeviceScopeable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScopeable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceScopes gets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScopes()([]UserExperienceAnalyticsDeviceScopeable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceScopes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceScopeable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceScores gets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceScores")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceScoresable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceStartupHistory gets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceStartupHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceStartupHistoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceStartupProcesses gets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceStartupProcesses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceStartupProcessable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceStartupProcessPerformance gets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceStartupProcessPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsDevicesWithoutCloudIdentity gets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDevicesWithoutCloudIdentity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable)
    }
    return nil
}
// GetUserExperienceAnalyticsDeviceTimelineEvent gets the userExperienceAnalyticsDeviceTimelineEvent property value. The user experience analytics device events entity contains NRT device timeline event details.
func (m *DeviceManagement) GetUserExperienceAnalyticsDeviceTimelineEvent()([]UserExperienceAnalyticsDeviceTimelineEventable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsDeviceTimelineEvent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsDeviceTimelineEventable)
    }
    return nil
}
// GetUserExperienceAnalyticsImpactingProcess gets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcessable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsImpactingProcess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsImpactingProcessable)
    }
    return nil
}
// GetUserExperienceAnalyticsMetricHistory gets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsMetricHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsMetricHistoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsModelScores gets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsModelScores")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsModelScoresable)
    }
    return nil
}
// GetUserExperienceAnalyticsNotAutopilotReadyDevice gets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDeviceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsNotAutopilotReadyDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsNotAutopilotReadyDeviceable)
    }
    return nil
}
// GetUserExperienceAnalyticsOverview gets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsOverviewable)
    }
    return nil
}
// GetUserExperienceAnalyticsRemoteConnection gets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnectionable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsRemoteConnection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsRemoteConnectionable)
    }
    return nil
}
// GetUserExperienceAnalyticsResourcePerformance gets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsResourcePerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsResourcePerformanceable)
    }
    return nil
}
// GetUserExperienceAnalyticsScoreHistory gets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsScoreHistory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsScoreHistoryable)
    }
    return nil
}
// GetUserExperienceAnalyticsSettings gets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsSettingsable)
    }
    return nil
}
// GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric gets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)
    }
    return nil
}
// GetUserExperienceAnalyticsWorkFromAnywhereMetrics gets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsWorkFromAnywhereMetrics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsWorkFromAnywhereMetricable)
    }
    return nil
}
// GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance gets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable) {
    val, err := m.GetBackingStore().Get("userExperienceAnalyticsWorkFromAnywhereModelPerformance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
    }
    return nil
}
// GetUserPfxCertificates gets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) GetUserPfxCertificates()([]UserPFXCertificateable) {
    val, err := m.GetBackingStore().Get("userPfxCertificates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserPFXCertificateable)
    }
    return nil
}
// GetVirtualEndpoint gets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) GetVirtualEndpoint()(VirtualEndpointable) {
    val, err := m.GetBackingStore().Get("virtualEndpoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualEndpointable)
    }
    return nil
}
// GetWindowsAutopilotDeploymentProfiles gets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfileable) {
    val, err := m.GetBackingStore().Get("windowsAutopilotDeploymentProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsAutopilotDeploymentProfileable)
    }
    return nil
}
// GetWindowsAutopilotDeviceIdentities gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable) {
    val, err := m.GetBackingStore().Get("windowsAutopilotDeviceIdentities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsAutopilotDeviceIdentityable)
    }
    return nil
}
// GetWindowsAutopilotSettings gets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) GetWindowsAutopilotSettings()(WindowsAutopilotSettingsable) {
    val, err := m.GetBackingStore().Get("windowsAutopilotSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsAutopilotSettingsable)
    }
    return nil
}
// GetWindowsDriverUpdateProfiles gets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfileable) {
    val, err := m.GetBackingStore().Get("windowsDriverUpdateProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsDriverUpdateProfileable)
    }
    return nil
}
// GetWindowsFeatureUpdateProfiles gets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfileable) {
    val, err := m.GetBackingStore().Get("windowsFeatureUpdateProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsFeatureUpdateProfileable)
    }
    return nil
}
// GetWindowsInformationProtectionAppLearningSummaries gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionAppLearningSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsInformationProtectionAppLearningSummaryable)
    }
    return nil
}
// GetWindowsInformationProtectionNetworkLearningSummaries gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable) {
    val, err := m.GetBackingStore().Get("windowsInformationProtectionNetworkLearningSummaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsInformationProtectionNetworkLearningSummaryable)
    }
    return nil
}
// GetWindowsMalwareInformation gets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) GetWindowsMalwareInformation()([]WindowsMalwareInformationable) {
    val, err := m.GetBackingStore().Get("windowsMalwareInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsMalwareInformationable)
    }
    return nil
}
// GetWindowsMalwareOverview gets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) GetWindowsMalwareOverview()(WindowsMalwareOverviewable) {
    val, err := m.GetBackingStore().Get("windowsMalwareOverview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsMalwareOverviewable)
    }
    return nil
}
// GetWindowsQualityUpdateProfiles gets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfileable) {
    val, err := m.GetBackingStore().Get("windowsQualityUpdateProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsQualityUpdateProfileable)
    }
    return nil
}
// GetWindowsUpdateCatalogItems gets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItemable) {
    val, err := m.GetBackingStore().Get("windowsUpdateCatalogItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsUpdateCatalogItemable)
    }
    return nil
}
// GetZebraFotaArtifacts gets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) GetZebraFotaArtifacts()([]ZebraFotaArtifactable) {
    val, err := m.GetBackingStore().Get("zebraFotaArtifacts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ZebraFotaArtifactable)
    }
    return nil
}
// GetZebraFotaConnector gets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) GetZebraFotaConnector()(ZebraFotaConnectorable) {
    val, err := m.GetBackingStore().Get("zebraFotaConnector")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ZebraFotaConnectorable)
    }
    return nil
}
// GetZebraFotaDeployments gets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) GetZebraFotaDeployments()([]ZebraFotaDeploymentable) {
    val, err := m.GetBackingStore().Get("zebraFotaDeployments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ZebraFotaDeploymentable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidDeviceOwnerEnrollmentProfiles()))
        for i, v := range m.GetAndroidDeviceOwnerEnrollmentProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("androidDeviceOwnerEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkAppConfigurationSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidForWorkAppConfigurationSchemas()))
        for i, v := range m.GetAndroidForWorkAppConfigurationSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("androidForWorkAppConfigurationSchemas", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAndroidForWorkEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidForWorkEnrollmentProfiles()))
        for i, v := range m.GetAndroidForWorkEnrollmentProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetAndroidManagedStoreAppConfigurationSchemas() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidManagedStoreAppConfigurationSchemas()))
        for i, v := range m.GetAndroidManagedStoreAppConfigurationSchemas() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetAppleUserInitiatedEnrollmentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppleUserInitiatedEnrollmentProfiles()))
        for i, v := range m.GetAppleUserInitiatedEnrollmentProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("appleUserInitiatedEnrollmentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilters()))
        for i, v := range m.GetAssignmentFilters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignmentFilters", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAuditEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuditEvents()))
        for i, v := range m.GetAuditEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("auditEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAutopilotEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutopilotEvents()))
        for i, v := range m.GetAutopilotEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("autopilotEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCartToClassAssociations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCartToClassAssociations()))
        for i, v := range m.GetCartToClassAssociations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cartToClassAssociations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertificateConnectorDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateConnectorDetails()))
        for i, v := range m.GetCertificateConnectorDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("certificateConnectorDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChromeOSOnboardingSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChromeOSOnboardingSettings()))
        for i, v := range m.GetChromeOSOnboardingSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPCConnectivityIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPCConnectivityIssues()))
        for i, v := range m.GetCloudPCConnectivityIssues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cloudPCConnectivityIssues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComanagedDevices()))
        for i, v := range m.GetComanagedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("comanagedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComanagementEligibleDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComanagementEligibleDevices()))
        for i, v := range m.GetComanagementEligibleDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("comanagementEligibleDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceCategories()))
        for i, v := range m.GetComplianceCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("complianceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceManagementPartners()))
        for i, v := range m.GetComplianceManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliancePolicies()))
        for i, v := range m.GetCompliancePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("compliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceSettings()))
        for i, v := range m.GetComplianceSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetConfigManagerCollections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigManagerCollections()))
        for i, v := range m.GetConfigManagerCollections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configManagerCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationCategories()))
        for i, v := range m.GetConfigurationCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configurationCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationPolicies()))
        for i, v := range m.GetConfigurationPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationPolicyTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationPolicyTemplates()))
        for i, v := range m.GetConfigurationPolicyTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configurationPolicyTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConfigurationSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConfigurationSettings()))
        for i, v := range m.GetConfigurationSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("configurationSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectorStatus() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnectorStatus()))
        for i, v := range m.GetConnectorStatus() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("connectorStatus", cast)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataSharingConsents()))
        for i, v := range m.GetDataSharingConsents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dataSharingConsents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDepOnboardingSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDepOnboardingSettings()))
        for i, v := range m.GetDepOnboardingSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("depOnboardingSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDerivedCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDerivedCredentials()))
        for i, v := range m.GetDerivedCredentials() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("derivedCredentials", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCategories()))
        for i, v := range m.GetDeviceCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicies()))
        for i, v := range m.GetDeviceCompliancePolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetDeviceCompliancePolicySettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceComplianceScripts()))
        for i, v := range m.GetDeviceComplianceScripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceComplianceScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationConflictSummary() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationConflictSummary()))
        for i, v := range m.GetDeviceConfigurationConflictSummary() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetDeviceConfigurationRestrictedAppsViolations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationRestrictedAppsViolations()))
        for i, v := range m.GetDeviceConfigurationRestrictedAppsViolations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationRestrictedAppsViolations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurations()))
        for i, v := range m.GetDeviceConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationsAllManagedDeviceCertificateStates()))
        for i, v := range m.GetDeviceConfigurationsAllManagedDeviceCertificateStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetDeviceCustomAttributeShellScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCustomAttributeShellScripts()))
        for i, v := range m.GetDeviceCustomAttributeShellScripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCustomAttributeShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceEnrollmentConfigurations()))
        for i, v := range m.GetDeviceEnrollmentConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceHealthScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceHealthScripts()))
        for i, v := range m.GetDeviceHealthScripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceHealthScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementPartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementPartners()))
        for i, v := range m.GetDeviceManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceManagementScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceManagementScripts()))
        for i, v := range m.GetDeviceManagementScripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetDeviceShellScripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceShellScripts()))
        for i, v := range m.GetDeviceShellScripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceShellScripts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDomainJoinConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDomainJoinConnectors()))
        for i, v := range m.GetDomainJoinConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("domainJoinConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmbeddedSIMActivationCodePools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmbeddedSIMActivationCodePools()))
        for i, v := range m.GetEmbeddedSIMActivationCodePools() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("embeddedSIMActivationCodePools", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeConnectors()))
        for i, v := range m.GetExchangeConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeOnPremisesPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExchangeOnPremisesPolicies()))
        for i, v := range m.GetExchangeOnPremisesPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetGroupPolicyCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyCategories()))
        for i, v := range m.GetGroupPolicyCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyConfigurations()))
        for i, v := range m.GetGroupPolicyConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitionFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyDefinitionFiles()))
        for i, v := range m.GetGroupPolicyDefinitionFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyDefinitions()))
        for i, v := range m.GetGroupPolicyDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyMigrationReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyMigrationReports()))
        for i, v := range m.GetGroupPolicyMigrationReports() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyMigrationReports", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyObjectFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyObjectFiles()))
        for i, v := range m.GetGroupPolicyObjectFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyObjectFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicyUploadedDefinitionFiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicyUploadedDefinitionFiles()))
        for i, v := range m.GetGroupPolicyUploadedDefinitionFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicyUploadedDefinitionFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedDeviceIdentities()))
        for i, v := range m.GetImportedDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("importedDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportedWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntents()))
        for i, v := range m.GetIntents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("intuneAccountId", m.GetIntuneAccountId())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIntuneBrandingProfiles()))
        for i, v := range m.GetIntuneBrandingProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("intuneBrandingProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosUpdateStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosUpdateStatuses()))
        for i, v := range m.GetIosUpdateStatuses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMacOSSoftwareUpdateAccountSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMacOSSoftwareUpdateAccountSummaries()))
        for i, v := range m.GetMacOSSoftwareUpdateAccountSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetManagedDeviceEncryptionStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceEncryptionStates()))
        for i, v := range m.GetManagedDeviceEncryptionStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelConfigurations()))
        for i, v := range m.GetMicrosoftTunnelConfigurations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelHealthThresholds() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelHealthThresholds()))
        for i, v := range m.GetMicrosoftTunnelHealthThresholds() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelHealthThresholds", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelServerLogCollectionResponses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelServerLogCollectionResponses()))
        for i, v := range m.GetMicrosoftTunnelServerLogCollectionResponses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServerLogCollectionResponses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelSites()))
        for i, v := range m.GetMicrosoftTunnelSites() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelSites", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppTroubleshootingEvents()))
        for i, v := range m.GetMobileAppTroubleshootingEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppTroubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileThreatDefenseConnectors()))
        for i, v := range m.GetMobileThreatDefenseConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNdesConnectors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNdesConnectors()))
        for i, v := range m.GetNdesConnectors() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("ndesConnectors", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotificationMessageTemplates()))
        for i, v := range m.GetNotificationMessageTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPrivilegeManagementElevations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrivilegeManagementElevations()))
        for i, v := range m.GetPrivilegeManagementElevations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("privilegeManagementElevations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteActionAudits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteActionAudits()))
        for i, v := range m.GetRemoteActionAudits() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("remoteActionAudits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoteAssistancePartners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteAssistancePartners()))
        for i, v := range m.GetRemoteAssistancePartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetResourceAccessProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceAccessProfiles()))
        for i, v := range m.GetResourceAccessProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceAccessProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceOperations()))
        for i, v := range m.GetResourceOperations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusablePolicySettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReusablePolicySettings()))
        for i, v := range m.GetReusablePolicySettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("reusablePolicySettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReusableSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReusableSettings()))
        for i, v := range m.GetReusableSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("reusableSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRoleScopeTags()))
        for i, v := range m.GetRoleScopeTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("roleScopeTags", cast)
        if err != nil {
            return err
        }
    }
    if m.GetServiceNowConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServiceNowConnections()))
        for i, v := range m.GetServiceNowConnections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("serviceNowConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTelecomExpenseManagementPartners()))
        for i, v := range m.GetTelecomExpenseManagementPartners() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateInsights() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateInsights()))
        for i, v := range m.GetTemplateInsights() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("templateInsights", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplates()))
        for i, v := range m.GetTemplates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("templates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateSettings()))
        for i, v := range m.GetTemplateSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTermsAndConditions()))
        for i, v := range m.GetTermsAndConditions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTroubleshootingEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTroubleshootingEvents()))
        for i, v := range m.GetTroubleshootingEvents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomaly() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAnomaly()))
        for i, v := range m.GetUserExperienceAnalyticsAnomaly() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomaly", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview()))
        for i, v := range m.GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAnomalyCorrelationGroupOverview", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAnomalyDevice() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAnomalyDevice()))
        for i, v := range m.GetUserExperienceAnalyticsAnomalyDevice() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDeviceModelPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDeviceModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsAppHealthDevicePerformanceDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsAppHealthOSVersionPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetUserExperienceAnalyticsBaselines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBaselines()))
        for i, v := range m.GetUserExperienceAnalyticsBaselines() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBaselines", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthAppImpact() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthAppImpact() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceAppImpact", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDevicePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDevicePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthModelPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsBatteryHealthModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsBatteryHealthOsPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsBatteryHealthOsPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetUserExperienceAnalyticsCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsCategories()))
        for i, v := range m.GetUserExperienceAnalyticsCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceMetricHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceMetricHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDevicePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDevicePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceScopes()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScopes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceScores()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceScores() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupHistory()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcesses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcesses()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcesses() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcesses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceStartupProcessPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceStartupProcessPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()))
        for i, v := range m.GetUserExperienceAnalyticsDevicesWithoutCloudIdentity() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDevicesWithoutCloudIdentity", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsDeviceTimelineEvent() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsDeviceTimelineEvent()))
        for i, v := range m.GetUserExperienceAnalyticsDeviceTimelineEvent() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsDeviceTimelineEvent", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsImpactingProcess() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsImpactingProcess()))
        for i, v := range m.GetUserExperienceAnalyticsImpactingProcess() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsImpactingProcess", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsMetricHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsMetricHistory()))
        for i, v := range m.GetUserExperienceAnalyticsMetricHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsMetricHistory", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsModelScores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsModelScores()))
        for i, v := range m.GetUserExperienceAnalyticsModelScores() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsModelScores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsNotAutopilotReadyDevice()))
        for i, v := range m.GetUserExperienceAnalyticsNotAutopilotReadyDevice() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetUserExperienceAnalyticsRemoteConnection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsRemoteConnection()))
        for i, v := range m.GetUserExperienceAnalyticsRemoteConnection() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsRemoteConnection", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsResourcePerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsResourcePerformance()))
        for i, v := range m.GetUserExperienceAnalyticsResourcePerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsResourcePerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsScoreHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsScoreHistory()))
        for i, v := range m.GetUserExperienceAnalyticsScoreHistory() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereMetrics() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereMetrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()))
        for i, v := range m.GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("userExperienceAnalyticsWorkFromAnywhereModelPerformance", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserPfxCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserPfxCertificates()))
        for i, v := range m.GetUserPfxCertificates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetWindowsAutopilotDeploymentProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsAutopilotDeploymentProfiles()))
        for i, v := range m.GetWindowsAutopilotDeploymentProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeploymentProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsAutopilotDeviceIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetWindowsAutopilotDeviceIdentities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetWindowsDriverUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsDriverUpdateProfiles()))
        for i, v := range m.GetWindowsDriverUpdateProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsDriverUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsFeatureUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsFeatureUpdateProfiles()))
        for i, v := range m.GetWindowsFeatureUpdateProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsFeatureUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionAppLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionAppLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionAppLearningSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionNetworkLearningSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionNetworkLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionNetworkLearningSummaries() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsMalwareInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsMalwareInformation()))
        for i, v := range m.GetWindowsMalwareInformation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetWindowsQualityUpdateProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsQualityUpdateProfiles()))
        for i, v := range m.GetWindowsQualityUpdateProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsQualityUpdateProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsUpdateCatalogItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsUpdateCatalogItems()))
        for i, v := range m.GetWindowsUpdateCatalogItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("windowsUpdateCatalogItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetZebraFotaArtifacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetZebraFotaArtifacts()))
        for i, v := range m.GetZebraFotaArtifacts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetZebraFotaDeployments()))
        for i, v := range m.GetZebraFotaDeployments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("zebraFotaDeployments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountMoveCompletionDateTime sets the accountMoveCompletionDateTime property value. The date & time when tenant data moved between scaleunits.
func (m *DeviceManagement) SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("accountMoveCompletionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAdminConsent sets the adminConsent property value. Admin consent information.
func (m *DeviceManagement) SetAdminConsent(value AdminConsentable)() {
    err := m.GetBackingStore().Set("adminConsent", value)
    if err != nil {
        panic(err)
    }
}
// SetAdvancedThreatProtectionOnboardingStateSummary sets the advancedThreatProtectionOnboardingStateSummary property value. The summary state of ATP onboarding state for this account.
func (m *DeviceManagement) SetAdvancedThreatProtectionOnboardingStateSummary(value AdvancedThreatProtectionOnboardingStateSummaryable)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOnboardingStateSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidDeviceOwnerEnrollmentProfiles sets the androidDeviceOwnerEnrollmentProfiles property value. Android device owner enrollment profile entities.
func (m *DeviceManagement) SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfileable)() {
    err := m.GetBackingStore().Set("androidDeviceOwnerEnrollmentProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidForWorkAppConfigurationSchemas sets the androidForWorkAppConfigurationSchemas property value. Android for Work app configuration schema entities.
func (m *DeviceManagement) SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchemaable)() {
    err := m.GetBackingStore().Set("androidForWorkAppConfigurationSchemas", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidForWorkEnrollmentProfiles sets the androidForWorkEnrollmentProfiles property value. Android for Work enrollment profile entities.
func (m *DeviceManagement) SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfileable)() {
    err := m.GetBackingStore().Set("androidForWorkEnrollmentProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidForWorkSettings sets the androidForWorkSettings property value. The singleton Android for Work settings entity.
func (m *DeviceManagement) SetAndroidForWorkSettings(value AndroidForWorkSettingsable)() {
    err := m.GetBackingStore().Set("androidForWorkSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidManagedStoreAccountEnterpriseSettings sets the androidManagedStoreAccountEnterpriseSettings property value. The singleton Android managed store account enterprise settings entity.
func (m *DeviceManagement) SetAndroidManagedStoreAccountEnterpriseSettings(value AndroidManagedStoreAccountEnterpriseSettingsable)() {
    err := m.GetBackingStore().Set("androidManagedStoreAccountEnterpriseSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidManagedStoreAppConfigurationSchemas sets the androidManagedStoreAppConfigurationSchemas property value. Android Enterprise app configuration schema entities.
func (m *DeviceManagement) SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchemaable)() {
    err := m.GetBackingStore().Set("androidManagedStoreAppConfigurationSchemas", value)
    if err != nil {
        panic(err)
    }
}
// SetApplePushNotificationCertificate sets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)() {
    err := m.GetBackingStore().Set("applePushNotificationCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetAppleUserInitiatedEnrollmentProfiles sets the appleUserInitiatedEnrollmentProfiles property value. Apple user initiated enrollment profiles
func (m *DeviceManagement) SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfileable)() {
    err := m.GetBackingStore().Set("appleUserInitiatedEnrollmentProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilters sets the assignmentFilters property value. The list of assignment filters
func (m *DeviceManagement) SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilterable)() {
    err := m.GetBackingStore().Set("assignmentFilters", value)
    if err != nil {
        panic(err)
    }
}
// SetAuditEvents sets the auditEvents property value. The Audit Events
func (m *DeviceManagement) SetAuditEvents(value []AuditEventable)() {
    err := m.GetBackingStore().Set("auditEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetAutopilotEvents sets the autopilotEvents property value. The list of autopilot events for the tenant.
func (m *DeviceManagement) SetAutopilotEvents(value []DeviceManagementAutopilotEventable)() {
    err := m.GetBackingStore().Set("autopilotEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetCartToClassAssociations sets the cartToClassAssociations property value. The Cart To Class Associations.
func (m *DeviceManagement) SetCartToClassAssociations(value []CartToClassAssociationable)() {
    err := m.GetBackingStore().Set("cartToClassAssociations", value)
    if err != nil {
        panic(err)
    }
}
// SetCategories sets the categories property value. The available categories
func (m *DeviceManagement) SetCategories(value []DeviceManagementSettingCategoryable)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateConnectorDetails sets the certificateConnectorDetails property value. Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *DeviceManagement) SetCertificateConnectorDetails(value []CertificateConnectorDetailsable)() {
    err := m.GetBackingStore().Set("certificateConnectorDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetChromeOSOnboardingSettings sets the chromeOSOnboardingSettings property value. Collection of ChromeOSOnboardingSettings settings associated with account.
func (m *DeviceManagement) SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettingsable)() {
    err := m.GetBackingStore().Set("chromeOSOnboardingSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPCConnectivityIssues sets the cloudPCConnectivityIssues property value. The list of CloudPC Connectivity Issue.
func (m *DeviceManagement) SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssueable)() {
    err := m.GetBackingStore().Set("cloudPCConnectivityIssues", value)
    if err != nil {
        panic(err)
    }
}
// SetComanagedDevices sets the comanagedDevices property value. The list of co-managed devices report
func (m *DeviceManagement) SetComanagedDevices(value []ManagedDeviceable)() {
    err := m.GetBackingStore().Set("comanagedDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetComanagementEligibleDevices sets the comanagementEligibleDevices property value. The list of co-management eligible devices report
func (m *DeviceManagement) SetComanagementEligibleDevices(value []ComanagementEligibleDeviceable)() {
    err := m.GetBackingStore().Set("comanagementEligibleDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceCategories sets the complianceCategories property value. List of all compliance categories
func (m *DeviceManagement) SetComplianceCategories(value []DeviceManagementConfigurationCategoryable)() {
    err := m.GetBackingStore().Set("complianceCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceManagementPartners sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartnerable)() {
    err := m.GetBackingStore().Set("complianceManagementPartners", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliancePolicies sets the compliancePolicies property value. List of all compliance policies
func (m *DeviceManagement) SetCompliancePolicies(value []DeviceManagementCompliancePolicyable)() {
    err := m.GetBackingStore().Set("compliancePolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceSettings sets the complianceSettings property value. List of all ComplianceSettings
func (m *DeviceManagement) SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    err := m.GetBackingStore().Set("complianceSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessSettings sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)() {
    err := m.GetBackingStore().Set("conditionalAccessSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigManagerCollections sets the configManagerCollections property value. A list of ConfigManagerCollection
func (m *DeviceManagement) SetConfigManagerCollections(value []ConfigManagerCollectionable)() {
    err := m.GetBackingStore().Set("configManagerCollections", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationCategories sets the configurationCategories property value. List of all Configuration Categories
func (m *DeviceManagement) SetConfigurationCategories(value []DeviceManagementConfigurationCategoryable)() {
    err := m.GetBackingStore().Set("configurationCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationPolicies sets the configurationPolicies property value. List of all Configuration policies
func (m *DeviceManagement) SetConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)() {
    err := m.GetBackingStore().Set("configurationPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationPolicyTemplates sets the configurationPolicyTemplates property value. List of all templates
func (m *DeviceManagement) SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplateable)() {
    err := m.GetBackingStore().Set("configurationPolicyTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationSettings sets the configurationSettings property value. List of all ConfigurationSettings
func (m *DeviceManagement) SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    err := m.GetBackingStore().Set("configurationSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorStatus sets the connectorStatus property value. The list of connector status for the tenant.
func (m *DeviceManagement) SetConnectorStatus(value []ConnectorStatusDetailsable)() {
    err := m.GetBackingStore().Set("connectorStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDataProcessorServiceForWindowsFeaturesOnboarding sets the dataProcessorServiceForWindowsFeaturesOnboarding property value. A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
func (m *DeviceManagement) SetDataProcessorServiceForWindowsFeaturesOnboarding(value DataProcessorServiceForWindowsFeaturesOnboardingable)() {
    err := m.GetBackingStore().Set("dataProcessorServiceForWindowsFeaturesOnboarding", value)
    if err != nil {
        panic(err)
    }
}
// SetDataSharingConsents sets the dataSharingConsents property value. Data sharing consents.
func (m *DeviceManagement) SetDataSharingConsents(value []DataSharingConsentable)() {
    err := m.GetBackingStore().Set("dataSharingConsents", value)
    if err != nil {
        panic(err)
    }
}
// SetDepOnboardingSettings sets the depOnboardingSettings property value. This collections of multiple DEP tokens per-tenant.
func (m *DeviceManagement) SetDepOnboardingSettings(value []DepOnboardingSettingable)() {
    err := m.GetBackingStore().Set("depOnboardingSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetDerivedCredentials sets the derivedCredentials property value. Collection of Derived credential settings associated with account.
func (m *DeviceManagement) SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettingsable)() {
    err := m.GetBackingStore().Set("derivedCredentials", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectedApps sets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) SetDetectedApps(value []DetectedAppable)() {
    err := m.GetBackingStore().Set("detectedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCategories sets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategoryable)() {
    err := m.GetBackingStore().Set("deviceCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicies sets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicyDeviceStateSummary sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicyDeviceStateSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicySettingStateSummaries sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicySettingStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceComplianceReportSummarizationDateTime sets the deviceComplianceReportSummarizationDateTime property value. The last requested time of device compliance reporting for this account. This property is read-only.
func (m *DeviceManagement) SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("deviceComplianceReportSummarizationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceComplianceScripts sets the deviceComplianceScripts property value. The list of device compliance scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceComplianceScripts(value []DeviceComplianceScriptable)() {
    err := m.GetBackingStore().Set("deviceComplianceScripts", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationConflictSummary sets the deviceConfigurationConflictSummary property value. Summary of policies in conflict state for this account.
func (m *DeviceManagement) SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummaryable)() {
    err := m.GetBackingStore().Set("deviceConfigurationConflictSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationDeviceStateSummaries sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceConfigurationDeviceStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationRestrictedAppsViolations sets the deviceConfigurationRestrictedAppsViolations property value. Restricted apps violations for this account.
func (m *DeviceManagement) SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolationable)() {
    err := m.GetBackingStore().Set("deviceConfigurationRestrictedAppsViolations", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurations sets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfigurationable)() {
    err := m.GetBackingStore().Set("deviceConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationsAllManagedDeviceCertificateStates sets the deviceConfigurationsAllManagedDeviceCertificateStates property value. Summary of all certificates for all devices.
func (m *DeviceManagement) SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("deviceConfigurationsAllManagedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationUserStateSummaries sets the deviceConfigurationUserStateSummaries property value. The device configuration user state summary for this account.
func (m *DeviceManagement) SetDeviceConfigurationUserStateSummaries(value DeviceConfigurationUserStateSummaryable)() {
    err := m.GetBackingStore().Set("deviceConfigurationUserStateSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCustomAttributeShellScripts sets the deviceCustomAttributeShellScripts property value. The list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScriptable)() {
    err := m.GetBackingStore().Set("deviceCustomAttributeShellScripts", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceEnrollmentConfigurations sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)() {
    err := m.GetBackingStore().Set("deviceEnrollmentConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceHealthScripts sets the deviceHealthScripts property value. The list of device health scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceHealthScripts(value []DeviceHealthScriptable)() {
    err := m.GetBackingStore().Set("deviceHealthScripts", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManagementPartners sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartnerable)() {
    err := m.GetBackingStore().Set("deviceManagementPartners", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManagementScripts sets the deviceManagementScripts property value. The list of device management scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceManagementScripts(value []DeviceManagementScriptable)() {
    err := m.GetBackingStore().Set("deviceManagementScripts", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceProtectionOverview sets the deviceProtectionOverview property value. Device protection overview.
func (m *DeviceManagement) SetDeviceProtectionOverview(value DeviceProtectionOverviewable)() {
    err := m.GetBackingStore().Set("deviceProtectionOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceShellScripts sets the deviceShellScripts property value. The list of device shell scripts associated with the tenant.
func (m *DeviceManagement) SetDeviceShellScripts(value []DeviceShellScriptable)() {
    err := m.GetBackingStore().Set("deviceShellScripts", value)
    if err != nil {
        panic(err)
    }
}
// SetDomainJoinConnectors sets the domainJoinConnectors property value. A list of connector objects.
func (m *DeviceManagement) SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnectorable)() {
    err := m.GetBackingStore().Set("domainJoinConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetEmbeddedSIMActivationCodePools sets the embeddedSIMActivationCodePools property value. The embedded SIM activation code pools created by this account.
func (m *DeviceManagement) SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePoolable)() {
    err := m.GetBackingStore().Set("embeddedSIMActivationCodePools", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeConnectors sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)() {
    err := m.GetBackingStore().Set("exchangeConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeOnPremisesPolicies sets the exchangeOnPremisesPolicies property value. The list of Exchange On Premisis policies configured by the tenant.
func (m *DeviceManagement) SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicyable)() {
    err := m.GetBackingStore().Set("exchangeOnPremisesPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeOnPremisesPolicy sets the exchangeOnPremisesPolicy property value. The policy which controls mobile device access to Exchange On Premises
func (m *DeviceManagement) SetExchangeOnPremisesPolicy(value DeviceManagementExchangeOnPremisesPolicyable)() {
    err := m.GetBackingStore().Set("exchangeOnPremisesPolicy", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyCategories sets the groupPolicyCategories property value. The available group policy categories for this account.
func (m *DeviceManagement) SetGroupPolicyCategories(value []GroupPolicyCategoryable)() {
    err := m.GetBackingStore().Set("groupPolicyCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyConfigurations sets the groupPolicyConfigurations property value. The group policy configurations created by this account.
func (m *DeviceManagement) SetGroupPolicyConfigurations(value []GroupPolicyConfigurationable)() {
    err := m.GetBackingStore().Set("groupPolicyConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyDefinitionFiles sets the groupPolicyDefinitionFiles property value. The available group policy definition files for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFileable)() {
    err := m.GetBackingStore().Set("groupPolicyDefinitionFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyDefinitions sets the groupPolicyDefinitions property value. The available group policy definitions for this account.
func (m *DeviceManagement) SetGroupPolicyDefinitions(value []GroupPolicyDefinitionable)() {
    err := m.GetBackingStore().Set("groupPolicyDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyMigrationReports sets the groupPolicyMigrationReports property value. A list of Group Policy migration reports.
func (m *DeviceManagement) SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReportable)() {
    err := m.GetBackingStore().Set("groupPolicyMigrationReports", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyObjectFiles sets the groupPolicyObjectFiles property value. A list of Group Policy Object files uploaded.
func (m *DeviceManagement) SetGroupPolicyObjectFiles(value []GroupPolicyObjectFileable)() {
    err := m.GetBackingStore().Set("groupPolicyObjectFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyUploadedDefinitionFiles sets the groupPolicyUploadedDefinitionFiles property value. The available group policy uploaded definition files for this account.
func (m *DeviceManagement) SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFileable)() {
    err := m.GetBackingStore().Set("groupPolicyUploadedDefinitionFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetImportedDeviceIdentities sets the importedDeviceIdentities property value. The imported device identities.
func (m *DeviceManagement) SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)() {
    err := m.GetBackingStore().Set("importedDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetImportedWindowsAutopilotDeviceIdentities sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)() {
    err := m.GetBackingStore().Set("importedWindowsAutopilotDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetIntents sets the intents property value. The device management intents
func (m *DeviceManagement) SetIntents(value []DeviceManagementIntentable)() {
    err := m.GetBackingStore().Set("intents", value)
    if err != nil {
        panic(err)
    }
}
// SetIntuneAccountId sets the intuneAccountId property value. Intune Account ID for given tenant
func (m *DeviceManagement) SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("intuneAccountId", value)
    if err != nil {
        panic(err)
    }
}
// SetIntuneBrand sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) SetIntuneBrand(value IntuneBrandable)() {
    err := m.GetBackingStore().Set("intuneBrand", value)
    if err != nil {
        panic(err)
    }
}
// SetIntuneBrandingProfiles sets the intuneBrandingProfiles property value. Intune branding profiles targeted to AAD groups
func (m *DeviceManagement) SetIntuneBrandingProfiles(value []IntuneBrandingProfileable)() {
    err := m.GetBackingStore().Set("intuneBrandingProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetIosUpdateStatuses sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)() {
    err := m.GetBackingStore().Set("iosUpdateStatuses", value)
    if err != nil {
        panic(err)
    }
}
// SetLastReportAggregationDateTime sets the lastReportAggregationDateTime property value. The last modified time of reporting for this account. This property is read-only.
func (m *DeviceManagement) SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastReportAggregationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLegacyPcManangementEnabled sets the legacyPcManangementEnabled property value. The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
func (m *DeviceManagement) SetLegacyPcManangementEnabled(value *bool)() {
    err := m.GetBackingStore().Set("legacyPcManangementEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMacOSSoftwareUpdateAccountSummaries sets the macOSSoftwareUpdateAccountSummaries property value. The MacOS software update account summaries for this account.
func (m *DeviceManagement) SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummaryable)() {
    err := m.GetBackingStore().Set("macOSSoftwareUpdateAccountSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCleanupSettings sets the managedDeviceCleanupSettings property value. Device cleanup rule
func (m *DeviceManagement) SetManagedDeviceCleanupSettings(value ManagedDeviceCleanupSettingsable)() {
    err := m.GetBackingStore().Set("managedDeviceCleanupSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceEncryptionStates sets the managedDeviceEncryptionStates property value. Encryption report for devices in this account
func (m *DeviceManagement) SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionStateable)() {
    err := m.GetBackingStore().Set("managedDeviceEncryptionStates", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceOverview sets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) SetManagedDeviceOverview(value ManagedDeviceOverviewable)() {
    err := m.GetBackingStore().Set("managedDeviceOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDevices sets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDeviceable)() {
    err := m.GetBackingStore().Set("managedDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumDepTokens sets the maximumDepTokens property value. Maximum number of DEP tokens allowed per-tenant.
func (m *DeviceManagement) SetMaximumDepTokens(value *int32)() {
    err := m.GetBackingStore().Set("maximumDepTokens", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelConfigurations sets the microsoftTunnelConfigurations property value. Collection of MicrosoftTunnelConfiguration settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfigurationable)() {
    err := m.GetBackingStore().Set("microsoftTunnelConfigurations", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelHealthThresholds sets the microsoftTunnelHealthThresholds property value. Collection of MicrosoftTunnelHealthThreshold settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThresholdable)() {
    err := m.GetBackingStore().Set("microsoftTunnelHealthThresholds", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelServerLogCollectionResponses sets the microsoftTunnelServerLogCollectionResponses property value. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponseable)() {
    err := m.GetBackingStore().Set("microsoftTunnelServerLogCollectionResponses", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelSites sets the microsoftTunnelSites property value. Collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagement) SetMicrosoftTunnelSites(value []MicrosoftTunnelSiteable)() {
    err := m.GetBackingStore().Set("microsoftTunnelSites", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileAppTroubleshootingEvents sets the mobileAppTroubleshootingEvents property value. The collection property of MobileAppTroubleshootingEvent.
func (m *DeviceManagement) SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)() {
    err := m.GetBackingStore().Set("mobileAppTroubleshootingEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileThreatDefenseConnectors sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)() {
    err := m.GetBackingStore().Set("mobileThreatDefenseConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetNdesConnectors sets the ndesConnectors property value. The collection of Ndes connectors for this account.
func (m *DeviceManagement) SetNdesConnectors(value []NdesConnectorable)() {
    err := m.GetBackingStore().Set("ndesConnectors", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationMessageTemplates sets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplateable)() {
    err := m.GetBackingStore().Set("notificationMessageTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivilegeManagementElevations sets the privilegeManagementElevations property value. The endpoint privilege management elevation event entity contains elevation details.
func (m *DeviceManagement) SetPrivilegeManagementElevations(value []PrivilegeManagementElevationable)() {
    err := m.GetBackingStore().Set("privilegeManagementElevations", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteActionAudits sets the remoteActionAudits property value. The list of device remote action audits with the tenant.
func (m *DeviceManagement) SetRemoteActionAudits(value []RemoteActionAuditable)() {
    err := m.GetBackingStore().Set("remoteActionAudits", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAssistancePartners sets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)() {
    err := m.GetBackingStore().Set("remoteAssistancePartners", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAssistanceSettings sets the remoteAssistanceSettings property value. The remote assistance settings singleton
func (m *DeviceManagement) SetRemoteAssistanceSettings(value RemoteAssistanceSettingsable)() {
    err := m.GetBackingStore().Set("remoteAssistanceSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetReports sets the reports property value. Reports singleton
func (m *DeviceManagement) SetReports(value DeviceManagementReportsable)() {
    err := m.GetBackingStore().Set("reports", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceAccessProfiles sets the resourceAccessProfiles property value. Collection of resource access settings associated with account.
func (m *DeviceManagement) SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBaseable)() {
    err := m.GetBackingStore().Set("resourceAccessProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceOperations sets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperationable)() {
    err := m.GetBackingStore().Set("resourceOperations", value)
    if err != nil {
        panic(err)
    }
}
// SetReusablePolicySettings sets the reusablePolicySettings property value. List of all reusable settings that can be referred in a policy
func (m *DeviceManagement) SetReusablePolicySettings(value []DeviceManagementReusablePolicySettingable)() {
    err := m.GetBackingStore().Set("reusablePolicySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetReusableSettings sets the reusableSettings property value. List of all reusable settings
func (m *DeviceManagement) SetReusableSettings(value []DeviceManagementConfigurationSettingDefinitionable)() {
    err := m.GetBackingStore().Set("reusableSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleAssignments sets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)() {
    err := m.GetBackingStore().Set("roleAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleDefinitions sets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinitionable)() {
    err := m.GetBackingStore().Set("roleDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTags sets the roleScopeTags property value. The Role Scope Tags.
func (m *DeviceManagement) SetRoleScopeTags(value []RoleScopeTagable)() {
    err := m.GetBackingStore().Set("roleScopeTags", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceNowConnections sets the serviceNowConnections property value. A list of ServiceNowConnections
func (m *DeviceManagement) SetServiceNowConnections(value []ServiceNowConnectionable)() {
    err := m.GetBackingStore().Set("serviceNowConnections", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDefinitions sets the settingDefinitions property value. The device management intent setting definitions
func (m *DeviceManagement) SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)() {
    err := m.GetBackingStore().Set("settingDefinitions", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Account level settings.
func (m *DeviceManagement) SetSettings(value DeviceManagementSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareUpdateStatusSummary sets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)() {
    err := m.GetBackingStore().Set("softwareUpdateStatusSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptions sets the subscriptions property value. Tenant mobile device management subscriptions.
func (m *DeviceManagement) SetSubscriptions(value *DeviceManagementSubscriptions)() {
    err := m.GetBackingStore().Set("subscriptions", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionState sets the subscriptionState property value. Tenant mobile device management subscription state.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    err := m.GetBackingStore().Set("subscriptionState", value)
    if err != nil {
        panic(err)
    }
}
// SetTelecomExpenseManagementPartners sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)() {
    err := m.GetBackingStore().Set("telecomExpenseManagementPartners", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateInsights sets the templateInsights property value. List of setting insights in a template
func (m *DeviceManagement) SetTemplateInsights(value []DeviceManagementTemplateInsightsDefinitionable)() {
    err := m.GetBackingStore().Set("templateInsights", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplates sets the templates property value. The available templates
func (m *DeviceManagement) SetTemplates(value []DeviceManagementTemplateable)() {
    err := m.GetBackingStore().Set("templates", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateSettings sets the templateSettings property value. List of all TemplateSettings
func (m *DeviceManagement) SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplateable)() {
    err := m.GetBackingStore().Set("templateSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantAttachRBAC sets the tenantAttachRBAC property value. TenantAttach RBAC Enablement
func (m *DeviceManagement) SetTenantAttachRBAC(value TenantAttachRBACable)() {
    err := m.GetBackingStore().Set("tenantAttachRBAC", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsAndConditions sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditionsable)() {
    err := m.GetBackingStore().Set("termsAndConditions", value)
    if err != nil {
        panic(err)
    }
}
// SetTroubleshootingEvents sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)() {
    err := m.GetBackingStore().Set("troubleshootingEvents", value)
    if err != nil {
        panic(err)
    }
}
// SetUnlicensedAdminstratorsEnabled sets the unlicensedAdminstratorsEnabled property value. When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor does the 350 member limit apply. This property is read-only.
func (m *DeviceManagement) SetUnlicensedAdminstratorsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("unlicensedAdminstratorsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAnomaly sets the userExperienceAnalyticsAnomaly property value. The user experience analytics anomaly entity contains anomaly details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomaly(value []UserExperienceAnalyticsAnomalyable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAnomaly", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAnomalyCorrelationGroupOverview sets the userExperienceAnalyticsAnomalyCorrelationGroupOverview property value. The user experience analytics anomaly correlation group overview entity contains the information for each correlation group of an anomaly.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalyCorrelationGroupOverview(value []UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAnomalyCorrelationGroupOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAnomalyDevice sets the userExperienceAnalyticsAnomalyDevice property value. The user experience analytics anomaly entity contains device details.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalyDevice(value []UserExperienceAnalyticsAnomalyDeviceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAnomalyDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAnomalySeverityOverview sets the userExperienceAnalyticsAnomalySeverityOverview property value. The user experience analytics anomaly severity overview entity contains the count information for each severity of anomaly.
func (m *DeviceManagement) SetUserExperienceAnalyticsAnomalySeverityOverview(value UserExperienceAnalyticsAnomalySeverityOverviewable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAnomalySeverityOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformance sets the userExperienceAnalyticsAppHealthApplicationPerformance property value. User experience analytics appHealth Application Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthApplicationPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion property value. User experience analytics appHealth Application Performance by App Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property value. User experience analytics appHealth Application Performance by App Version details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId sets the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property value. User experience analytics appHealth Application Performance by App Version Device Id
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion sets the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property value. User experience analytics appHealth Application Performance by OS Version
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthDeviceModelPerformance sets the userExperienceAnalyticsAppHealthDeviceModelPerformance property value. User experience analytics appHealth Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthDeviceModelPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthDevicePerformance sets the userExperienceAnalyticsAppHealthDevicePerformance property value. User experience analytics appHealth Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthDevicePerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails sets the userExperienceAnalyticsAppHealthDevicePerformanceDetails property value. User experience analytics device performance details
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthDevicePerformanceDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthOSVersionPerformance sets the userExperienceAnalyticsAppHealthOSVersionPerformance property value. User experience analytics appHealth OS version Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthOSVersionPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsAppHealthOverview sets the userExperienceAnalyticsAppHealthOverview property value. User experience analytics appHealth overview
func (m *DeviceManagement) SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsAppHealthOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBaselines sets the userExperienceAnalyticsBaselines property value. User experience analytics baselines
func (m *DeviceManagement) SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBaselines", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthAppImpact sets the userExperienceAnalyticsBatteryHealthAppImpact property value. User Experience Analytics Battery Health App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpactable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthAppImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthCapacityDetails sets the userExperienceAnalyticsBatteryHealthCapacityDetails property value. User Experience Analytics Battery Health Capacity Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value UserExperienceAnalyticsBatteryHealthCapacityDetailsable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthCapacityDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact sets the userExperienceAnalyticsBatteryHealthDeviceAppImpact property value. User Experience Analytics Battery Health Device App Impact
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthDeviceAppImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthDevicePerformance sets the userExperienceAnalyticsBatteryHealthDevicePerformance property value. User Experience Analytics Battery Health Device Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthDevicePerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory sets the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property value. User Experience Analytics Battery Health Device Runtime History
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthModelPerformance sets the userExperienceAnalyticsBatteryHealthModelPerformance property value. User Experience Analytics Battery Health Model Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthModelPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthOsPerformance sets the userExperienceAnalyticsBatteryHealthOsPerformance property value. User Experience Analytics Battery Health Os Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthOsPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsBatteryHealthRuntimeDetails sets the userExperienceAnalyticsBatteryHealthRuntimeDetails property value. User Experience Analytics Battery Health Runtime Details
func (m *DeviceManagement) SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsBatteryHealthRuntimeDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsCategories sets the userExperienceAnalyticsCategories property value. User experience analytics categories
func (m *DeviceManagement) SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceMetricHistory sets the userExperienceAnalyticsDeviceMetricHistory property value. User experience analytics device metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceMetricHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDevicePerformance sets the userExperienceAnalyticsDevicePerformance property value. User experience analytics device performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDevicePerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceScope sets the userExperienceAnalyticsDeviceScope property value. The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP computing metrics data based on a device scope configuration.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScope(value UserExperienceAnalyticsDeviceScopeable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceScope", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceScopes sets the userExperienceAnalyticsDeviceScopes property value. The user experience analytics device scope entity contains device scope configuration use to apply filtering on the endpoint analytics reports.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScopes(value []UserExperienceAnalyticsDeviceScopeable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceScopes", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceScores sets the userExperienceAnalyticsDeviceScores property value. User experience analytics device scores
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceScores", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceStartupHistory sets the userExperienceAnalyticsDeviceStartupHistory property value. User experience analytics device Startup History
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceStartupHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceStartupProcesses sets the userExperienceAnalyticsDeviceStartupProcesses property value. User experience analytics device Startup Processes
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceStartupProcesses", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceStartupProcessPerformance sets the userExperienceAnalyticsDeviceStartupProcessPerformance property value. User experience analytics device Startup Process Performance
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceStartupProcessPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDevicesWithoutCloudIdentity sets the userExperienceAnalyticsDevicesWithoutCloudIdentity property value. User experience analytics devices without cloud identity.
func (m *DeviceManagement) SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentityable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDevicesWithoutCloudIdentity", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsDeviceTimelineEvent sets the userExperienceAnalyticsDeviceTimelineEvent property value. The user experience analytics device events entity contains NRT device timeline event details.
func (m *DeviceManagement) SetUserExperienceAnalyticsDeviceTimelineEvent(value []UserExperienceAnalyticsDeviceTimelineEventable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsDeviceTimelineEvent", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsImpactingProcess sets the userExperienceAnalyticsImpactingProcess property value. User experience analytics impacting process
func (m *DeviceManagement) SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcessable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsImpactingProcess", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsMetricHistory sets the userExperienceAnalyticsMetricHistory property value. User experience analytics metric history
func (m *DeviceManagement) SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsMetricHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsModelScores sets the userExperienceAnalyticsModelScores property value. User experience analytics model scores
func (m *DeviceManagement) SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsModelScores", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsNotAutopilotReadyDevice sets the userExperienceAnalyticsNotAutopilotReadyDevice property value. User experience analytics devices not Windows Autopilot ready.
func (m *DeviceManagement) SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDeviceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsNotAutopilotReadyDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsOverview sets the userExperienceAnalyticsOverview property value. User experience analytics overview
func (m *DeviceManagement) SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsRemoteConnection sets the userExperienceAnalyticsRemoteConnection property value. User experience analytics remote connection
func (m *DeviceManagement) SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnectionable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsRemoteConnection", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsResourcePerformance sets the userExperienceAnalyticsResourcePerformance property value. User experience analytics resource performance
func (m *DeviceManagement) SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsResourcePerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsScoreHistory sets the userExperienceAnalyticsScoreHistory property value. User experience analytics device Startup Score History
func (m *DeviceManagement) SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsScoreHistory", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsSettings sets the userExperienceAnalyticsSettings property value. User experience analytics device settings
func (m *DeviceManagement) SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric sets the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property value. User experience analytics work from anywhere hardware readiness metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsWorkFromAnywhereMetrics sets the userExperienceAnalyticsWorkFromAnywhereMetrics property value. User experience analytics work from anywhere metrics.
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsWorkFromAnywhereMetrics", value)
    if err != nil {
        panic(err)
    }
}
// SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance sets the userExperienceAnalyticsWorkFromAnywhereModelPerformance property value. The user experience analytics work from anywhere model performance
func (m *DeviceManagement) SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)() {
    err := m.GetBackingStore().Set("userExperienceAnalyticsWorkFromAnywhereModelPerformance", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPfxCertificates sets the userPfxCertificates property value. Collection of PFX certificates associated with a user.
func (m *DeviceManagement) SetUserPfxCertificates(value []UserPFXCertificateable)() {
    err := m.GetBackingStore().Set("userPfxCertificates", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualEndpoint sets the virtualEndpoint property value. The virtualEndpoint property
func (m *DeviceManagement) SetVirtualEndpoint(value VirtualEndpointable)() {
    err := m.GetBackingStore().Set("virtualEndpoint", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsAutopilotDeploymentProfiles sets the windowsAutopilotDeploymentProfiles property value. Windows auto pilot deployment profiles
func (m *DeviceManagement) SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfileable)() {
    err := m.GetBackingStore().Set("windowsAutopilotDeploymentProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsAutopilotDeviceIdentities sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)() {
    err := m.GetBackingStore().Set("windowsAutopilotDeviceIdentities", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsAutopilotSettings sets the windowsAutopilotSettings property value. The Windows autopilot account settings.
func (m *DeviceManagement) SetWindowsAutopilotSettings(value WindowsAutopilotSettingsable)() {
    err := m.GetBackingStore().Set("windowsAutopilotSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsDriverUpdateProfiles sets the windowsDriverUpdateProfiles property value. A collection of windows driver update profiles
func (m *DeviceManagement) SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfileable)() {
    err := m.GetBackingStore().Set("windowsDriverUpdateProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsFeatureUpdateProfiles sets the windowsFeatureUpdateProfiles property value. A collection of windows feature update profiles
func (m *DeviceManagement) SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfileable)() {
    err := m.GetBackingStore().Set("windowsFeatureUpdateProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionAppLearningSummaries sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionAppLearningSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsInformationProtectionNetworkLearningSummaries sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)() {
    err := m.GetBackingStore().Set("windowsInformationProtectionNetworkLearningSummaries", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsMalwareInformation sets the windowsMalwareInformation property value. The list of affected malware in the tenant.
func (m *DeviceManagement) SetWindowsMalwareInformation(value []WindowsMalwareInformationable)() {
    err := m.GetBackingStore().Set("windowsMalwareInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsMalwareOverview sets the windowsMalwareOverview property value. Malware overview for windows devices.
func (m *DeviceManagement) SetWindowsMalwareOverview(value WindowsMalwareOverviewable)() {
    err := m.GetBackingStore().Set("windowsMalwareOverview", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsQualityUpdateProfiles sets the windowsQualityUpdateProfiles property value. A collection of windows quality update profiles
func (m *DeviceManagement) SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfileable)() {
    err := m.GetBackingStore().Set("windowsQualityUpdateProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsUpdateCatalogItems sets the windowsUpdateCatalogItems property value. A collection of windows update catalog items (fetaure updates item , quality updates item)
func (m *DeviceManagement) SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItemable)() {
    err := m.GetBackingStore().Set("windowsUpdateCatalogItems", value)
    if err != nil {
        panic(err)
    }
}
// SetZebraFotaArtifacts sets the zebraFotaArtifacts property value. The Collection of ZebraFotaArtifacts.
func (m *DeviceManagement) SetZebraFotaArtifacts(value []ZebraFotaArtifactable)() {
    err := m.GetBackingStore().Set("zebraFotaArtifacts", value)
    if err != nil {
        panic(err)
    }
}
// SetZebraFotaConnector sets the zebraFotaConnector property value. The singleton ZebraFotaConnector associated with account.
func (m *DeviceManagement) SetZebraFotaConnector(value ZebraFotaConnectorable)() {
    err := m.GetBackingStore().Set("zebraFotaConnector", value)
    if err != nil {
        panic(err)
    }
}
// SetZebraFotaDeployments sets the zebraFotaDeployments property value. Collection of ZebraFotaDeployments associated with account.
func (m *DeviceManagement) SetZebraFotaDeployments(value []ZebraFotaDeploymentable)() {
    err := m.GetBackingStore().Set("zebraFotaDeployments", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementable 
type DeviceManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdminConsent()(AdminConsentable)
    GetAdvancedThreatProtectionOnboardingStateSummary()(AdvancedThreatProtectionOnboardingStateSummaryable)
    GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfileable)
    GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchemaable)
    GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfileable)
    GetAndroidForWorkSettings()(AndroidForWorkSettingsable)
    GetAndroidManagedStoreAccountEnterpriseSettings()(AndroidManagedStoreAccountEnterpriseSettingsable)
    GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchemaable)
    GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable)
    GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfileable)
    GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilterable)
    GetAuditEvents()([]AuditEventable)
    GetAutopilotEvents()([]DeviceManagementAutopilotEventable)
    GetCartToClassAssociations()([]CartToClassAssociationable)
    GetCategories()([]DeviceManagementSettingCategoryable)
    GetCertificateConnectorDetails()([]CertificateConnectorDetailsable)
    GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettingsable)
    GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssueable)
    GetComanagedDevices()([]ManagedDeviceable)
    GetComanagementEligibleDevices()([]ComanagementEligibleDeviceable)
    GetComplianceCategories()([]DeviceManagementConfigurationCategoryable)
    GetComplianceManagementPartners()([]ComplianceManagementPartnerable)
    GetCompliancePolicies()([]DeviceManagementCompliancePolicyable)
    GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetConfigManagerCollections()([]ConfigManagerCollectionable)
    GetConfigurationCategories()([]DeviceManagementConfigurationCategoryable)
    GetConfigurationPolicies()([]DeviceManagementConfigurationPolicyable)
    GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplateable)
    GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetConnectorStatus()([]ConnectorStatusDetailsable)
    GetDataProcessorServiceForWindowsFeaturesOnboarding()(DataProcessorServiceForWindowsFeaturesOnboardingable)
    GetDataSharingConsents()([]DataSharingConsentable)
    GetDepOnboardingSettings()([]DepOnboardingSettingable)
    GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettingsable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceCategories()([]DeviceCategoryable)
    GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable)
    GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceComplianceScripts()([]DeviceComplianceScriptable)
    GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummaryable)
    GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable)
    GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolationable)
    GetDeviceConfigurations()([]DeviceConfigurationable)
    GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateStateable)
    GetDeviceConfigurationUserStateSummaries()(DeviceConfigurationUserStateSummaryable)
    GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScriptable)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceHealthScripts()([]DeviceHealthScriptable)
    GetDeviceManagementPartners()([]DeviceManagementPartnerable)
    GetDeviceManagementScripts()([]DeviceManagementScriptable)
    GetDeviceProtectionOverview()(DeviceProtectionOverviewable)
    GetDeviceShellScripts()([]DeviceShellScriptable)
    GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnectorable)
    GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePoolable)
    GetExchangeConnectors()([]DeviceManagementExchangeConnectorable)
    GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicyable)
    GetExchangeOnPremisesPolicy()(DeviceManagementExchangeOnPremisesPolicyable)
    GetGroupPolicyCategories()([]GroupPolicyCategoryable)
    GetGroupPolicyConfigurations()([]GroupPolicyConfigurationable)
    GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFileable)
    GetGroupPolicyDefinitions()([]GroupPolicyDefinitionable)
    GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReportable)
    GetGroupPolicyObjectFiles()([]GroupPolicyObjectFileable)
    GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFileable)
    GetImportedDeviceIdentities()([]ImportedDeviceIdentityable)
    GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable)
    GetIntents()([]DeviceManagementIntentable)
    GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIntuneBrand()(IntuneBrandable)
    GetIntuneBrandingProfiles()([]IntuneBrandingProfileable)
    GetIosUpdateStatuses()([]IosUpdateDeviceStatusable)
    GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegacyPcManangementEnabled()(*bool)
    GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummaryable)
    GetManagedDeviceCleanupSettings()(ManagedDeviceCleanupSettingsable)
    GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionStateable)
    GetManagedDeviceOverview()(ManagedDeviceOverviewable)
    GetManagedDevices()([]ManagedDeviceable)
    GetMaximumDepTokens()(*int32)
    GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfigurationable)
    GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThresholdable)
    GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponseable)
    GetMicrosoftTunnelSites()([]MicrosoftTunnelSiteable)
    GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable)
    GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable)
    GetNdesConnectors()([]NdesConnectorable)
    GetNotificationMessageTemplates()([]NotificationMessageTemplateable)
    GetPrivilegeManagementElevations()([]PrivilegeManagementElevationable)
    GetRemoteActionAudits()([]RemoteActionAuditable)
    GetRemoteAssistancePartners()([]RemoteAssistancePartnerable)
    GetRemoteAssistanceSettings()(RemoteAssistanceSettingsable)
    GetReports()(DeviceManagementReportsable)
    GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBaseable)
    GetResourceOperations()([]ResourceOperationable)
    GetReusablePolicySettings()([]DeviceManagementReusablePolicySettingable)
    GetReusableSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable)
    GetRoleDefinitions()([]RoleDefinitionable)
    GetRoleScopeTags()([]RoleScopeTagable)
    GetServiceNowConnections()([]ServiceNowConnectionable)
    GetSettingDefinitions()([]DeviceManagementSettingDefinitionable)
    GetSettings()(DeviceManagementSettingsable)
    GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable)
    GetSubscriptions()(*DeviceManagementSubscriptions)
    GetSubscriptionState()(*DeviceManagementSubscriptionState)
    GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable)
    GetTemplateInsights()([]DeviceManagementTemplateInsightsDefinitionable)
    GetTemplates()([]DeviceManagementTemplateable)
    GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplateable)
    GetTenantAttachRBAC()(TenantAttachRBACable)
    GetTermsAndConditions()([]TermsAndConditionsable)
    GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetUnlicensedAdminstratorsEnabled()(*bool)
    GetUserExperienceAnalyticsAnomaly()([]UserExperienceAnalyticsAnomalyable)
    GetUserExperienceAnalyticsAnomalyCorrelationGroupOverview()([]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)
    GetUserExperienceAnalyticsAnomalyDevice()([]UserExperienceAnalyticsAnomalyDeviceable)
    GetUserExperienceAnalyticsAnomalySeverityOverview()(UserExperienceAnalyticsAnomalySeverityOverviewable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
    GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
    GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
    GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable)
    GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpactable)
    GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(UserExperienceAnalyticsBatteryHealthCapacityDetailsable)
    GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)
    GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable)
    GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)
    GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformanceable)
    GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformanceable)
    GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)
    GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistoryable)
    GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable)
    GetUserExperienceAnalyticsDeviceScope()(UserExperienceAnalyticsDeviceScopeable)
    GetUserExperienceAnalyticsDeviceScopes()([]UserExperienceAnalyticsDeviceScopeable)
    GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable)
    GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable)
    GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable)
    GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
    GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable)
    GetUserExperienceAnalyticsDeviceTimelineEvent()([]UserExperienceAnalyticsDeviceTimelineEventable)
    GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcessable)
    GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable)
    GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable)
    GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDeviceable)
    GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable)
    GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnectionable)
    GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformanceable)
    GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable)
    GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable)
    GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
    GetUserPfxCertificates()([]UserPFXCertificateable)
    GetVirtualEndpoint()(VirtualEndpointable)
    GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfileable)
    GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable)
    GetWindowsAutopilotSettings()(WindowsAutopilotSettingsable)
    GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfileable)
    GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfileable)
    GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable)
    GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable)
    GetWindowsMalwareInformation()([]WindowsMalwareInformationable)
    GetWindowsMalwareOverview()(WindowsMalwareOverviewable)
    GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfileable)
    GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItemable)
    GetZebraFotaArtifacts()([]ZebraFotaArtifactable)
    GetZebraFotaConnector()(ZebraFotaConnectorable)
    GetZebraFotaDeployments()([]ZebraFotaDeploymentable)
    SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdminConsent(value AdminConsentable)()
    SetAdvancedThreatProtectionOnboardingStateSummary(value AdvancedThreatProtectionOnboardingStateSummaryable)()
    SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfileable)()
    SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchemaable)()
    SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfileable)()
    SetAndroidForWorkSettings(value AndroidForWorkSettingsable)()
    SetAndroidManagedStoreAccountEnterpriseSettings(value AndroidManagedStoreAccountEnterpriseSettingsable)()
    SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchemaable)()
    SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)()
    SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfileable)()
    SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilterable)()
    SetAuditEvents(value []AuditEventable)()
    SetAutopilotEvents(value []DeviceManagementAutopilotEventable)()
    SetCartToClassAssociations(value []CartToClassAssociationable)()
    SetCategories(value []DeviceManagementSettingCategoryable)()
    SetCertificateConnectorDetails(value []CertificateConnectorDetailsable)()
    SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettingsable)()
    SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssueable)()
    SetComanagedDevices(value []ManagedDeviceable)()
    SetComanagementEligibleDevices(value []ComanagementEligibleDeviceable)()
    SetComplianceCategories(value []DeviceManagementConfigurationCategoryable)()
    SetComplianceManagementPartners(value []ComplianceManagementPartnerable)()
    SetCompliancePolicies(value []DeviceManagementCompliancePolicyable)()
    SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetConfigManagerCollections(value []ConfigManagerCollectionable)()
    SetConfigurationCategories(value []DeviceManagementConfigurationCategoryable)()
    SetConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)()
    SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplateable)()
    SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetConnectorStatus(value []ConnectorStatusDetailsable)()
    SetDataProcessorServiceForWindowsFeaturesOnboarding(value DataProcessorServiceForWindowsFeaturesOnboardingable)()
    SetDataSharingConsents(value []DataSharingConsentable)()
    SetDepOnboardingSettings(value []DepOnboardingSettingable)()
    SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettingsable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceCategories(value []DeviceCategoryable)()
    SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)()
    SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceComplianceScripts(value []DeviceComplianceScriptable)()
    SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummaryable)()
    SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)()
    SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolationable)()
    SetDeviceConfigurations(value []DeviceConfigurationable)()
    SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateStateable)()
    SetDeviceConfigurationUserStateSummaries(value DeviceConfigurationUserStateSummaryable)()
    SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScriptable)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceHealthScripts(value []DeviceHealthScriptable)()
    SetDeviceManagementPartners(value []DeviceManagementPartnerable)()
    SetDeviceManagementScripts(value []DeviceManagementScriptable)()
    SetDeviceProtectionOverview(value DeviceProtectionOverviewable)()
    SetDeviceShellScripts(value []DeviceShellScriptable)()
    SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnectorable)()
    SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePoolable)()
    SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)()
    SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicyable)()
    SetExchangeOnPremisesPolicy(value DeviceManagementExchangeOnPremisesPolicyable)()
    SetGroupPolicyCategories(value []GroupPolicyCategoryable)()
    SetGroupPolicyConfigurations(value []GroupPolicyConfigurationable)()
    SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFileable)()
    SetGroupPolicyDefinitions(value []GroupPolicyDefinitionable)()
    SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReportable)()
    SetGroupPolicyObjectFiles(value []GroupPolicyObjectFileable)()
    SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFileable)()
    SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)()
    SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)()
    SetIntents(value []DeviceManagementIntentable)()
    SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIntuneBrand(value IntuneBrandable)()
    SetIntuneBrandingProfiles(value []IntuneBrandingProfileable)()
    SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)()
    SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegacyPcManangementEnabled(value *bool)()
    SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummaryable)()
    SetManagedDeviceCleanupSettings(value ManagedDeviceCleanupSettingsable)()
    SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionStateable)()
    SetManagedDeviceOverview(value ManagedDeviceOverviewable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetMaximumDepTokens(value *int32)()
    SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfigurationable)()
    SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThresholdable)()
    SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponseable)()
    SetMicrosoftTunnelSites(value []MicrosoftTunnelSiteable)()
    SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)()
    SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)()
    SetNdesConnectors(value []NdesConnectorable)()
    SetNotificationMessageTemplates(value []NotificationMessageTemplateable)()
    SetPrivilegeManagementElevations(value []PrivilegeManagementElevationable)()
    SetRemoteActionAudits(value []RemoteActionAuditable)()
    SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)()
    SetRemoteAssistanceSettings(value RemoteAssistanceSettingsable)()
    SetReports(value DeviceManagementReportsable)()
    SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBaseable)()
    SetResourceOperations(value []ResourceOperationable)()
    SetReusablePolicySettings(value []DeviceManagementReusablePolicySettingable)()
    SetReusableSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)()
    SetRoleDefinitions(value []RoleDefinitionable)()
    SetRoleScopeTags(value []RoleScopeTagable)()
    SetServiceNowConnections(value []ServiceNowConnectionable)()
    SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)()
    SetSettings(value DeviceManagementSettingsable)()
    SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)()
    SetSubscriptions(value *DeviceManagementSubscriptions)()
    SetSubscriptionState(value *DeviceManagementSubscriptionState)()
    SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)()
    SetTemplateInsights(value []DeviceManagementTemplateInsightsDefinitionable)()
    SetTemplates(value []DeviceManagementTemplateable)()
    SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplateable)()
    SetTenantAttachRBAC(value TenantAttachRBACable)()
    SetTermsAndConditions(value []TermsAndConditionsable)()
    SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetUnlicensedAdminstratorsEnabled(value *bool)()
    SetUserExperienceAnalyticsAnomaly(value []UserExperienceAnalyticsAnomalyable)()
    SetUserExperienceAnalyticsAnomalyCorrelationGroupOverview(value []UserExperienceAnalyticsAnomalyCorrelationGroupOverviewable)()
    SetUserExperienceAnalyticsAnomalyDevice(value []UserExperienceAnalyticsAnomalyDeviceable)()
    SetUserExperienceAnalyticsAnomalySeverityOverview(value UserExperienceAnalyticsAnomalySeverityOverviewable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)()
    SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)()
    SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)()
    SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)()
    SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpactable)()
    SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value UserExperienceAnalyticsBatteryHealthCapacityDetailsable)()
    SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)()
    SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)()
    SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)()
    SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)()
    SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)()
    SetUserExperienceAnalyticsDeviceScope(value UserExperienceAnalyticsDeviceScopeable)()
    SetUserExperienceAnalyticsDeviceScopes(value []UserExperienceAnalyticsDeviceScopeable)()
    SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)()
    SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)()
    SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)()
    SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)()
    SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentityable)()
    SetUserExperienceAnalyticsDeviceTimelineEvent(value []UserExperienceAnalyticsDeviceTimelineEventable)()
    SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcessable)()
    SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)()
    SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)()
    SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDeviceable)()
    SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)()
    SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnectionable)()
    SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformanceable)()
    SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)()
    SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)()
    SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)()
    SetUserPfxCertificates(value []UserPFXCertificateable)()
    SetVirtualEndpoint(value VirtualEndpointable)()
    SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfileable)()
    SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)()
    SetWindowsAutopilotSettings(value WindowsAutopilotSettingsable)()
    SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfileable)()
    SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfileable)()
    SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)()
    SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)()
    SetWindowsMalwareInformation(value []WindowsMalwareInformationable)()
    SetWindowsMalwareOverview(value WindowsMalwareOverviewable)()
    SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfileable)()
    SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItemable)()
    SetZebraFotaArtifacts(value []ZebraFotaArtifactable)()
    SetZebraFotaConnector(value ZebraFotaConnectorable)()
    SetZebraFotaDeployments(value []ZebraFotaDeploymentable)()
}
