package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDevice 
type ManagedDevice struct {
    Entity
    // Whether the device is Azure Active Directory registered. This property is read-only.
    aadRegistered *bool;
    // Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
    activationLockBypassCode *string;
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string;
    // Managed device mobile app configuration states for this device.
    assignmentFilterEvaluationStatusDetails []AssignmentFilterEvaluationStatusDetails;
    // Reports if the managed device is enrolled via auto-pilot. This property is read-only.
    autopilotEnrolled *bool;
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureActiveDirectoryDeviceId *string;
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string;
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool;
    // Chassis type of the device. This property is read-only. Possible values are: unknown, desktop, laptop, worksWorkstation, enterpriseServer, phone, tablet, mobileOther, mobileUnknown.
    chassisType *ChassisType;
    // List of properties of the ChromeOS Device.
    chromeOSDeviceInfo []ChromeOSDeviceProperty;
    // 
    cloudPcRemoteActionResults []CloudPcRemoteActionResult;
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
    complianceState *ComplianceState;
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures *ConfigurationManagerClientEnabledFeatures;
    // Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
    configurationManagerClientHealthState *ConfigurationManagerClientHealthState;
    // Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
    configurationManagerClientInformation *ConfigurationManagerClientInformation;
    // All applications currently installed on the device
    detectedApps []DetectedApp;
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResult;
    // Device category
    deviceCategory *DeviceCategory;
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string;
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyState;
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationState;
    // Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
    deviceEnrollmentType *DeviceEnrollmentType;
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState *DeviceHealthAttestationState;
    // Name of the device. This property is read-only.
    deviceName *string;
    // Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
    deviceRegistrationState *DeviceRegistrationState;
    // Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
    deviceType *DeviceType;
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool;
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string;
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string;
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
    enrollmentProfileName *string;
    // Ethernet MAC. This property is read-only.
    ethernetMacAddress *string;
    // The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
    exchangeAccessState *DeviceManagementExchangeAccessState;
    // The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason;
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Free Storage in Bytes. This property is read-only.
    freeStorageSpaceInBytes *int64;
    // The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
    hardwareInformation *HardwareInformation;
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
    iccid *string;
    // IMEI. This property is read-only.
    imei *string;
    // Device encryption status. This property is read-only.
    isEncrypted *bool;
    // Device supervised status. This property is read-only.
    isSupervised *bool;
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string;
    // Device join type. Possible values are: unknown, azureADJoined, azureADRegistered, hybridAzureADJoined.
    joinType *JoinType;
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of log collection requests
    logCollectionRequests []DeviceLogCollectionResponse;
    // Indicates if Lost mode is enabled or disabled. This property is read-only. Possible values are: disabled, enabled.
    lostModeState *LostModeState;
    // Managed device mobile app configuration states for this device.
    managedDeviceMobileAppConfigurationStates []ManagedDeviceMobileAppConfigurationState;
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string;
    // Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
    managedDeviceOwnerType *ManagedDeviceOwnerType;
    // Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementAgent *ManagementAgentType;
    // Reports device management certificate expiration date. This property is read-only.
    managementCertificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device management features. Possible values are: none, microsoftManagedDesktop.
    managementFeatures *ManagedDeviceManagementFeatures;
    // Management state of the device. This property is read-only. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
    managementState *ManagementState;
    // Manufacturer of the device. This property is read-only.
    manufacturer *string;
    // MEID. This property is read-only.
    meid *string;
    // Model of the device. This property is read-only.
    model *string;
    // Notes on the device created by IT Admin
    notes *string;
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string;
    // Operating system version of the device. This property is read-only.
    osVersion *string;
    // Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
    ownerType *OwnerType;
    // Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState;
    // Phone number of the device. This property is read-only.
    phoneNumber *string;
    // Total Memory in Bytes. This property is read-only.
    physicalMemoryInBytes *int64;
    // Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
    preferMdmOverGroupPolicyAppliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Processor architecture. This property is read-only. Possible values are: unknown, x86, x64, arm, arM64.
    processorArchitecture *ManagedDeviceArchitecture;
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string;
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string;
    // Reports if the managed iOS device is user approval enrollment. This property is read-only.
    requireUserEnrollmentApproval *bool;
    // Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
    retireAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tag IDs for this Device instance.
    roleScopeTagIds []string;
    // Security baseline states for this device.
    securityBaselineStates []SecurityBaselineState;
    // SerialNumber. This property is read-only.
    serialNumber *string;
    // Device sku family
    skuFamily *string;
    // Device sku number, see also: GetProductInfo function. Valid values 0 to 2147483647. This property is read-only.
    skuNumber *int32;
    // Specification version. This property is read-only.
    specificationVersion *string;
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string;
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64;
    // Unique Device Identifier for iOS and macOS devices. This property is read-only.
    udid *string;
    // User display name. This property is read-only.
    userDisplayName *string;
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string;
    // Device user principal name. This property is read-only.
    userPrincipalName *string;
    // The primary users associated with the managed device.
    users []User;
    // Indicates the last logged on users of a device. This property is read-only.
    usersLoggedOn []LoggedOnUser;
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string;
    // Count of active malware for this windows device. This property is read-only.
    windowsActiveMalwareCount *int32;
    // The device protection status. This property is read-only.
    windowsProtectionState *WindowsProtectionState;
    // Count of remediated malware for this windows device. This property is read-only.
    windowsRemediatedMalwareCount *int32;
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// GetAadRegistered gets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAadRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.aadRegistered
    }
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activationLockBypassCode
    }
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.androidSecurityPatchLevel
    }
}
// GetAssignmentFilterEvaluationStatusDetails gets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterEvaluationStatusDetails
    }
}
// GetAutopilotEnrolled gets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) GetAutopilotEnrolled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autopilotEnrolled
    }
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureActiveDirectoryDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureActiveDirectoryDeviceId
    }
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADDeviceId
    }
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureADRegistered
    }
}
// GetChassisType gets the chassisType property value. Chassis type of the device. This property is read-only. Possible values are: unknown, desktop, laptop, worksWorkstation, enterpriseServer, phone, tablet, mobileOther, mobileUnknown.
func (m *ManagedDevice) GetChassisType()(*ChassisType) {
    if m == nil {
        return nil
    } else {
        return m.chassisType
    }
}
// GetChromeOSDeviceInfo gets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) GetChromeOSDeviceInfo()([]ChromeOSDeviceProperty) {
    if m == nil {
        return nil
    } else {
        return m.chromeOSDeviceInfo
    }
}
// GetCloudPcRemoteActionResults gets the cloudPcRemoteActionResults property value. 
func (m *ManagedDevice) GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResult) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcRemoteActionResults
    }
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
// GetComplianceState gets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientEnabledFeatures
    }
}
// GetConfigurationManagerClientHealthState gets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientHealthState()(*ConfigurationManagerClientHealthState) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientHealthState
    }
}
// GetConfigurationManagerClientInformation gets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientInformation()(*ConfigurationManagerClientInformation) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerClientInformation
    }
}
// GetDetectedApps gets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) GetDetectedApps()([]DetectedApp) {
    if m == nil {
        return nil
    } else {
        return m.detectedApps
    }
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResult) {
    if m == nil {
        return nil
    } else {
        return m.deviceActionResults
    }
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(*DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategoryDisplayName
    }
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyState) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyStates
    }
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationStates
    }
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentType
    }
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationState
    }
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegistrationState
    }
}
// GetDeviceType gets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *ManagedDevice) GetDeviceType()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.easActivated
    }
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.easActivationDateTime
    }
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.easDeviceId
    }
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDateTime
    }
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) GetEnrollmentProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfileName
    }
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ethernetMacAddress
    }
}
// GetExchangeAccessState gets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessState
    }
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAccessStateReason
    }
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.exchangeLastSuccessfulSyncDateTime
    }
}
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpaceInBytes
    }
}
// GetHardwareInformation gets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
func (m *ManagedDevice) GetHardwareInformation()(*HardwareInformation) {
    if m == nil {
        return nil
    } else {
        return m.hardwareInformation
    }
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iccid
    }
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// GetJailBroken gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jailBroken
    }
}
// GetJoinType gets the joinType property value. Device join type. Possible values are: unknown, azureADJoined, azureADRegistered, hybridAzureADJoined.
func (m *ManagedDevice) GetJoinType()(*JoinType) {
    if m == nil {
        return nil
    } else {
        return m.joinType
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetLogCollectionRequests gets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) GetLogCollectionRequests()([]DeviceLogCollectionResponse) {
    if m == nil {
        return nil
    } else {
        return m.logCollectionRequests
    }
}
// GetLostModeState gets the lostModeState property value. Indicates if Lost mode is enabled or disabled. This property is read-only. Possible values are: disabled, enabled.
func (m *ManagedDevice) GetLostModeState()(*LostModeState) {
    if m == nil {
        return nil
    } else {
        return m.lostModeState
    }
}
// GetManagedDeviceMobileAppConfigurationStates gets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationState) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceMobileAppConfigurationStates
    }
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOwnerType
    }
}
// GetManagementAgent gets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgent
    }
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.managementCertificateExpirationDate
    }
}
// GetManagementFeatures gets the managementFeatures property value. Device management features. Possible values are: none, microsoftManagedDesktop.
func (m *ManagedDevice) GetManagementFeatures()(*ManagedDeviceManagementFeatures) {
    if m == nil {
        return nil
    } else {
        return m.managementFeatures
    }
}
// GetManagementState gets the managementState property value. Management state of the device. This property is read-only. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
func (m *ManagedDevice) GetManagementState()(*ManagementState) {
    if m == nil {
        return nil
    } else {
        return m.managementState
    }
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetOwnerType gets the ownerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    if m == nil {
        return nil
    } else {
        return m.partnerReportedThreatState
    }
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.physicalMemoryInBytes
    }
}
// GetPreferMdmOverGroupPolicyAppliedDateTime gets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.preferMdmOverGroupPolicyAppliedDateTime
    }
}
// GetProcessorArchitecture gets the processorArchitecture property value. Processor architecture. This property is read-only. Possible values are: unknown, x86, x64, arm, arM64.
func (m *ManagedDevice) GetProcessorArchitecture()(*ManagedDeviceArchitecture) {
    if m == nil {
        return nil
    } else {
        return m.processorArchitecture
    }
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionErrorDetails
    }
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistanceSessionUrl
    }
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireUserEnrollmentApproval
    }
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.retireAfterDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetSecurityBaselineStates gets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) GetSecurityBaselineStates()([]SecurityBaselineState) {
    if m == nil {
        return nil
    } else {
        return m.securityBaselineStates
    }
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetSkuFamily gets the skuFamily property value. Device sku family
func (m *ManagedDevice) GetSkuFamily()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuFamily
    }
}
// GetSkuNumber gets the skuNumber property value. Device sku number, see also: GetProductInfo function. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) GetSkuNumber()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skuNumber
    }
}
// GetSpecificationVersion gets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) GetSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.specificationVersion
    }
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpaceInBytes
    }
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.udid
    }
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUsers gets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) GetUsers()([]User) {
    if m == nil {
        return nil
    } else {
        return m.users
    }
}
// GetUsersLoggedOn gets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) GetUsersLoggedOn()([]LoggedOnUser) {
    if m == nil {
        return nil
    } else {
        return m.usersLoggedOn
    }
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wiFiMacAddress
    }
}
// GetWindowsActiveMalwareCount gets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsActiveMalwareCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsActiveMalwareCount
    }
}
// GetWindowsProtectionState gets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) GetWindowsProtectionState()(*WindowsProtectionState) {
    if m == nil {
        return nil
    } else {
        return m.windowsProtectionState
    }
}
// GetWindowsRemediatedMalwareCount gets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsRemediatedMalwareCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsRemediatedMalwareCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadRegistered(val)
        }
        return nil
    }
    res["activationLockBypassCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["assignmentFilterEvaluationStatusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentFilterEvaluationStatusDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationStatusDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignmentFilterEvaluationStatusDetails))
            }
            m.SetAssignmentFilterEvaluationStatusDetails(res)
        }
        return nil
    }
    res["autopilotEnrolled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopilotEnrolled(val)
        }
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureActiveDirectoryDeviceId(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["chassisType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChassisType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChassisType(val.(*ChassisType))
        }
        return nil
    }
    res["chromeOSDeviceInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChromeOSDeviceProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChromeOSDeviceProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChromeOSDeviceProperty))
            }
            m.SetChromeOSDeviceInfo(res)
        }
        return nil
    }
    res["cloudPcRemoteActionResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcRemoteActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcRemoteActionResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcRemoteActionResult))
            }
            m.SetCloudPcRemoteActionResults(res)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceState))
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigurationManagerClientEnabledFeatures() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(*ConfigurationManagerClientEnabledFeatures))
        }
        return nil
    }
    res["configurationManagerClientHealthState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigurationManagerClientHealthState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientHealthState(val.(*ConfigurationManagerClientHealthState))
        }
        return nil
    }
    res["configurationManagerClientInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConfigurationManagerClientInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientInformation(val.(*ConfigurationManagerClientInformation))
        }
        return nil
    }
    res["detectedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*DetectedApp))
            }
            m.SetDetectedApps(res)
        }
        return nil
    }
    res["deviceActionResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceActionResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceActionResult))
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(*DeviceCategory))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceCompliancePolicyState))
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationState))
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceHealthAttestationState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(*DeviceHealthAttestationState))
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationState(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["easActivated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessState(val.(*DeviceManagementExchangeAccessState))
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessStateReason(val.(*DeviceManagementExchangeAccessStateReason))
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["hardwareInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHardwareInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareInformation(val.(*HardwareInformation))
        }
        return nil
    }
    res["iccid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["joinType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinType(val.(*JoinType))
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["logCollectionRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceLogCollectionResponse() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLogCollectionResponse, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceLogCollectionResponse))
            }
            m.SetLogCollectionRequests(res)
        }
        return nil
    }
    res["lostModeState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLostModeState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLostModeState(val.(*LostModeState))
        }
        return nil
    }
    res["managedDeviceMobileAppConfigurationStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationState, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceMobileAppConfigurationState))
            }
            m.SetManagedDeviceMobileAppConfigurationStates(res)
        }
        return nil
    }
    res["managedDeviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["managementAgent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementCertificateExpirationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementCertificateExpirationDate(val)
        }
        return nil
    }
    res["managementFeatures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceManagementFeatures)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementFeatures(val.(*ManagedDeviceManagementFeatures))
        }
        return nil
    }
    res["managementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerReportedThreatState(val.(*ManagedDevicePartnerReportedHealthState))
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["preferMdmOverGroupPolicyAppliedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferMdmOverGroupPolicyAppliedDateTime(val)
        }
        return nil
    }
    res["processorArchitecture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorArchitecture(val.(*ManagedDeviceArchitecture))
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["requireUserEnrollmentApproval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserEnrollmentApproval(val)
        }
        return nil
    }
    res["retireAfterDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetireAfterDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["securityBaselineStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityBaselineState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineState, len(val))
            for i, v := range val {
                res[i] = *(v.(*SecurityBaselineState))
            }
            m.SetSecurityBaselineStates(res)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["skuFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuFamily(val)
        }
        return nil
    }
    res["skuNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["specificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecificationVersion(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["users"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]User, len(val))
            for i, v := range val {
                res[i] = *(v.(*User))
            }
            m.SetUsers(res)
        }
        return nil
    }
    res["usersLoggedOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLoggedOnUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LoggedOnUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*LoggedOnUser))
            }
            m.SetUsersLoggedOn(res)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiMacAddress(val)
        }
        return nil
    }
    res["windowsActiveMalwareCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsActiveMalwareCount(val)
        }
        return nil
    }
    res["windowsProtectionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsProtectionState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsProtectionState(val.(*WindowsProtectionState))
        }
        return nil
    }
    res["windowsRemediatedMalwareCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsRemediatedMalwareCount(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDevice) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("aadRegistered", m.GetAadRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("activationLockBypassCode", m.GetActivationLockBypassCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("androidSecurityPatchLevel", m.GetAndroidSecurityPatchLevel())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterEvaluationStatusDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentFilterEvaluationStatusDetails()))
        for i, v := range m.GetAssignmentFilterEvaluationStatusDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentFilterEvaluationStatusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autopilotEnrolled", m.GetAutopilotEnrolled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureActiveDirectoryDeviceId", m.GetAzureActiveDirectoryDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureADDeviceId", m.GetAzureADDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureADRegistered", m.GetAzureADRegistered())
        if err != nil {
            return err
        }
    }
    if m.GetChassisType() != nil {
        cast := (*m.GetChassisType()).String()
        err = writer.WriteStringValue("chassisType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChromeOSDeviceInfo() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChromeOSDeviceInfo()))
        for i, v := range m.GetChromeOSDeviceInfo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSDeviceInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcRemoteActionResults() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCloudPcRemoteActionResults()))
        for i, v := range m.GetCloudPcRemoteActionResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcRemoteActionResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err = writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientEnabledFeatures", m.GetConfigurationManagerClientEnabledFeatures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientHealthState", m.GetConfigurationManagerClientHealthState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientInformation", m.GetConfigurationManagerClientInformation())
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
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
    if m.GetDeviceActionResults() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceActionResults()))
        for i, v := range m.GetDeviceActionResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceActionResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceCategoryDisplayName", m.GetDeviceCategoryDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := (*m.GetDeviceEnrollmentType()).String()
        err = writer.WriteStringValue("deviceEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceHealthAttestationState", m.GetDeviceHealthAttestationState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRegistrationState() != nil {
        cast := (*m.GetDeviceRegistrationState()).String()
        err = writer.WriteStringValue("deviceRegistrationState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("easActivated", m.GetEasActivated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("easActivationDateTime", m.GetEasActivationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("easDeviceId", m.GetEasDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrolledDateTime", m.GetEnrolledDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentProfileName", m.GetEnrollmentProfileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ethernetMacAddress", m.GetEthernetMacAddress())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessState() != nil {
        cast := (*m.GetExchangeAccessState()).String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := (*m.GetExchangeAccessStateReason()).String()
        err = writer.WriteStringValue("exchangeAccessStateReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("exchangeLastSuccessfulSyncDateTime", m.GetExchangeLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("freeStorageSpaceInBytes", m.GetFreeStorageSpaceInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hardwareInformation", m.GetHardwareInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("iccid", m.GetIccid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imei", m.GetImei())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jailBroken", m.GetJailBroken())
        if err != nil {
            return err
        }
    }
    if m.GetJoinType() != nil {
        cast := (*m.GetJoinType()).String()
        err = writer.WriteStringValue("joinType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLogCollectionRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLogCollectionRequests()))
        for i, v := range m.GetLogCollectionRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("logCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLostModeState() != nil {
        cast := (*m.GetLostModeState()).String()
        err = writer.WriteStringValue("lostModeState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceMobileAppConfigurationStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDeviceMobileAppConfigurationStates()))
        for i, v := range m.GetManagedDeviceMobileAppConfigurationStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceMobileAppConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceOwnerType() != nil {
        cast := (*m.GetManagedDeviceOwnerType()).String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
        err = writer.WriteStringValue("managementAgent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("managementCertificateExpirationDate", m.GetManagementCertificateExpirationDate())
        if err != nil {
            return err
        }
    }
    if m.GetManagementFeatures() != nil {
        cast := (*m.GetManagementFeatures()).String()
        err = writer.WriteStringValue("managementFeatures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementState() != nil {
        cast := (*m.GetManagementState()).String()
        err = writer.WriteStringValue("managementState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meid", m.GetMeid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err = writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerReportedThreatState() != nil {
        cast := (*m.GetPartnerReportedThreatState()).String()
        err = writer.WriteStringValue("partnerReportedThreatState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("physicalMemoryInBytes", m.GetPhysicalMemoryInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("preferMdmOverGroupPolicyAppliedDateTime", m.GetPreferMdmOverGroupPolicyAppliedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetProcessorArchitecture() != nil {
        cast := (*m.GetProcessorArchitecture()).String()
        err = writer.WriteStringValue("processorArchitecture", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteAssistanceSessionErrorDetails", m.GetRemoteAssistanceSessionErrorDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteAssistanceSessionUrl", m.GetRemoteAssistanceSessionUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requireUserEnrollmentApproval", m.GetRequireUserEnrollmentApproval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("retireAfterDateTime", m.GetRetireAfterDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityBaselineStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecurityBaselineStates()))
        for i, v := range m.GetSecurityBaselineStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("securityBaselineStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuFamily", m.GetSkuFamily())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("skuNumber", m.GetSkuNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("specificationVersion", m.GetSpecificationVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriberCarrier", m.GetSubscriberCarrier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalStorageSpaceInBytes", m.GetTotalStorageSpaceInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("udid", m.GetUdid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUsersLoggedOn() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsersLoggedOn()))
        for i, v := range m.GetUsersLoggedOn() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("usersLoggedOn", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wiFiMacAddress", m.GetWiFiMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("windowsActiveMalwareCount", m.GetWindowsActiveMalwareCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsProtectionState", m.GetWindowsProtectionState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("windowsRemediatedMalwareCount", m.GetWindowsRemediatedMalwareCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadRegistered sets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAadRegistered(value *bool)() {
    if m != nil {
        m.aadRegistered = value
    }
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    if m != nil {
        m.activationLockBypassCode = value
    }
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    if m != nil {
        m.androidSecurityPatchLevel = value
    }
}
// SetAssignmentFilterEvaluationStatusDetails sets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetails)() {
    if m != nil {
        m.assignmentFilterEvaluationStatusDetails = value
    }
}
// SetAutopilotEnrolled sets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) SetAutopilotEnrolled(value *bool)() {
    if m != nil {
        m.autopilotEnrolled = value
    }
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureActiveDirectoryDeviceId(value *string)() {
    if m != nil {
        m.azureActiveDirectoryDeviceId = value
    }
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    if m != nil {
        m.azureADDeviceId = value
    }
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    if m != nil {
        m.azureADRegistered = value
    }
}
// SetChassisType sets the chassisType property value. Chassis type of the device. This property is read-only. Possible values are: unknown, desktop, laptop, worksWorkstation, enterpriseServer, phone, tablet, mobileOther, mobileUnknown.
func (m *ManagedDevice) SetChassisType(value *ChassisType)() {
    if m != nil {
        m.chassisType = value
    }
}
// SetChromeOSDeviceInfo sets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) SetChromeOSDeviceInfo(value []ChromeOSDeviceProperty)() {
    if m != nil {
        m.chromeOSDeviceInfo = value
    }
}
// SetCloudPcRemoteActionResults sets the cloudPcRemoteActionResults property value. 
func (m *ManagedDevice) SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResult)() {
    if m != nil {
        m.cloudPcRemoteActionResults = value
    }
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.complianceGracePeriodExpirationDateTime = value
    }
}
// SetComplianceState sets the complianceState property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    if m != nil {
        m.complianceState = value
    }
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value *ConfigurationManagerClientEnabledFeatures)() {
    if m != nil {
        m.configurationManagerClientEnabledFeatures = value
    }
}
// SetConfigurationManagerClientHealthState sets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientHealthState(value *ConfigurationManagerClientHealthState)() {
    if m != nil {
        m.configurationManagerClientHealthState = value
    }
}
// SetConfigurationManagerClientInformation sets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientInformation(value *ConfigurationManagerClientInformation)() {
    if m != nil {
        m.configurationManagerClientInformation = value
    }
}
// SetDetectedApps sets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) SetDetectedApps(value []DetectedApp)() {
    if m != nil {
        m.detectedApps = value
    }
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResult)() {
    if m != nil {
        m.deviceActionResults = value
    }
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value *DeviceCategory)() {
    if m != nil {
        m.deviceCategory = value
    }
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    if m != nil {
        m.deviceCategoryDisplayName = value
    }
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyState)() {
    if m != nil {
        m.deviceCompliancePolicyStates = value
    }
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationState)() {
    if m != nil {
        m.deviceConfigurationStates = value
    }
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Enrollment type of the device. This property is read-only. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement, windowsAzureADJoinUsingDeviceAuth, appleUserEnrollment, appleUserEnrollmentWithServiceAccount.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    if m != nil {
        m.deviceEnrollmentType = value
    }
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value *DeviceHealthAttestationState)() {
    if m != nil {
        m.deviceHealthAttestationState = value
    }
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration state. This property is read-only. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    if m != nil {
        m.deviceRegistrationState = value
    }
}
// SetDeviceType sets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *ManagedDevice) SetDeviceType(value *DeviceType)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    if m != nil {
        m.easActivated = value
    }
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.easActivationDateTime = value
    }
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    if m != nil {
        m.easDeviceId = value
    }
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.enrolledDateTime = value
    }
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) SetEnrollmentProfileName(value *string)() {
    if m != nil {
        m.enrollmentProfileName = value
    }
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    if m != nil {
        m.ethernetMacAddress = value
    }
}
// SetExchangeAccessState sets the exchangeAccessState property value. The Access State of the device in Exchange. This property is read-only. Possible values are: none, unknown, allowed, blocked, quarantined.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    if m != nil {
        m.exchangeAccessState = value
    }
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. The reason for the device's access state in Exchange. This property is read-only. Possible values are: none, unknown, exchangeGlobalRule, exchangeIndividualRule, exchangeDeviceRule, exchangeUpgrade, exchangeMailboxPolicy, other, compliant, notCompliant, notEnrolled, unknownLocation, mfaRequired, azureADBlockDueToAccessPolicy, compromisedPassword, deviceNotKnownWithManagedApp.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    if m != nil {
        m.exchangeAccessStateReason = value
    }
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.exchangeLastSuccessfulSyncDateTime = value
    }
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    if m != nil {
        m.freeStorageSpaceInBytes = value
    }
}
// SetHardwareInformation sets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
func (m *ManagedDevice) SetHardwareInformation(value *HardwareInformation)() {
    if m != nil {
        m.hardwareInformation = value
    }
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    if m != nil {
        m.iccid = value
    }
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    if m != nil {
        m.imei = value
    }
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    if m != nil {
        m.isEncrypted = value
    }
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    if m != nil {
        m.isSupervised = value
    }
}
// SetJailBroken sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    if m != nil {
        m.jailBroken = value
    }
}
// SetJoinType sets the joinType property value. Device join type. Possible values are: unknown, azureADJoined, azureADRegistered, hybridAzureADJoined.
func (m *ManagedDevice) SetJoinType(value *JoinType)() {
    if m != nil {
        m.joinType = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetLogCollectionRequests sets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) SetLogCollectionRequests(value []DeviceLogCollectionResponse)() {
    if m != nil {
        m.logCollectionRequests = value
    }
}
// SetLostModeState sets the lostModeState property value. Indicates if Lost mode is enabled or disabled. This property is read-only. Possible values are: disabled, enabled.
func (m *ManagedDevice) SetLostModeState(value *LostModeState)() {
    if m != nil {
        m.lostModeState = value
    }
}
// SetManagedDeviceMobileAppConfigurationStates sets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationState)() {
    if m != nil {
        m.managedDeviceMobileAppConfigurationStates = value
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    if m != nil {
        m.managedDeviceName = value
    }
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    if m != nil {
        m.managedDeviceOwnerType = value
    }
}
// SetManagementAgent sets the managementAgent property value. Management channel of the device. Intune, EAS, etc. This property is read-only. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    if m != nil {
        m.managementAgent = value
    }
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.managementCertificateExpirationDate = value
    }
}
// SetManagementFeatures sets the managementFeatures property value. Device management features. Possible values are: none, microsoftManagedDesktop.
func (m *ManagedDevice) SetManagementFeatures(value *ManagedDeviceManagementFeatures)() {
    if m != nil {
        m.managementFeatures = value
    }
}
// SetManagementState sets the managementState property value. Management state of the device. This property is read-only. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
func (m *ManagedDevice) SetManagementState(value *ManagementState)() {
    if m != nil {
        m.managementState = value
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    if m != nil {
        m.meid = value
    }
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    if m != nil {
        m.operatingSystem = value
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetOwnerType sets the ownerType property value. Ownership of the device. Can be 'company' or 'personal'. Possible values are: unknown, company, personal.
func (m *ManagedDevice) SetOwnerType(value *OwnerType)() {
    if m != nil {
        m.ownerType = value
    }
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Indicates the threat state of a device when a Mobile Threat Defense partner is in use by the account and device. Read Only. This property is read-only. Possible values are: unknown, activated, deactivated, secured, lowSeverity, mediumSeverity, highSeverity, unresponsive, compromised, misconfigured.
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    if m != nil {
        m.partnerReportedThreatState = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    if m != nil {
        m.physicalMemoryInBytes = value
    }
}
// SetPreferMdmOverGroupPolicyAppliedDateTime sets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.preferMdmOverGroupPolicyAppliedDateTime = value
    }
}
// SetProcessorArchitecture sets the processorArchitecture property value. Processor architecture. This property is read-only. Possible values are: unknown, x86, x64, arm, arM64.
func (m *ManagedDevice) SetProcessorArchitecture(value *ManagedDeviceArchitecture)() {
    if m != nil {
        m.processorArchitecture = value
    }
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    if m != nil {
        m.remoteAssistanceSessionErrorDetails = value
    }
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    if m != nil {
        m.remoteAssistanceSessionUrl = value
    }
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    if m != nil {
        m.requireUserEnrollmentApproval = value
    }
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.retireAfterDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetSecurityBaselineStates sets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) SetSecurityBaselineStates(value []SecurityBaselineState)() {
    if m != nil {
        m.securityBaselineStates = value
    }
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetSkuFamily sets the skuFamily property value. Device sku family
func (m *ManagedDevice) SetSkuFamily(value *string)() {
    if m != nil {
        m.skuFamily = value
    }
}
// SetSkuNumber sets the skuNumber property value. Device sku number, see also: GetProductInfo function. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) SetSkuNumber(value *int32)() {
    if m != nil {
        m.skuNumber = value
    }
}
// SetSpecificationVersion sets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) SetSpecificationVersion(value *string)() {
    if m != nil {
        m.specificationVersion = value
    }
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    if m != nil {
        m.subscriberCarrier = value
    }
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    if m != nil {
        m.totalStorageSpaceInBytes = value
    }
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    if m != nil {
        m.udid = value
    }
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    if m != nil {
        m.userDisplayName = value
    }
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []User)() {
    if m != nil {
        m.users = value
    }
}
// SetUsersLoggedOn sets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) SetUsersLoggedOn(value []LoggedOnUser)() {
    if m != nil {
        m.usersLoggedOn = value
    }
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    if m != nil {
        m.wiFiMacAddress = value
    }
}
// SetWindowsActiveMalwareCount sets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsActiveMalwareCount(value *int32)() {
    if m != nil {
        m.windowsActiveMalwareCount = value
    }
}
// SetWindowsProtectionState sets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) SetWindowsProtectionState(value *WindowsProtectionState)() {
    if m != nil {
        m.windowsProtectionState = value
    }
}
// SetWindowsRemediatedMalwareCount sets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsRemediatedMalwareCount(value *int32)() {
    if m != nil {
        m.windowsRemediatedMalwareCount = value
    }
}
