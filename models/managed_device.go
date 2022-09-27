package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
    Entity
    // Whether the device is Azure Active Directory registered. This property is read-only.
    aadRegistered *bool
    // Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
    activationLockBypassCode *string
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string
    // Managed device mobile app configuration states for this device.
    assignmentFilterEvaluationStatusDetails []AssignmentFilterEvaluationStatusDetailsable
    // Reports if the managed device is enrolled via auto-pilot. This property is read-only.
    autopilotEnrolled *bool
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureActiveDirectoryDeviceId *string
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool
    // Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
    bootstrapTokenEscrowed *bool
    // Chassis type.
    chassisType *ChassisType
    // List of properties of the ChromeOS Device.
    chromeOSDeviceInfo []ChromeOSDevicePropertyable
    // The cloudPcRemoteActionResults property
    cloudPcRemoteActionResults []CloudPcRemoteActionResultable
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Compliance state.
    complianceState *ComplianceState
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures ConfigurationManagerClientEnabledFeaturesable
    // Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
    configurationManagerClientHealthState ConfigurationManagerClientHealthStateable
    // Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
    configurationManagerClientInformation ConfigurationManagerClientInformationable
    // All applications currently installed on the device
    detectedApps []DetectedAppable
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResultable
    // Device category
    deviceCategory DeviceCategoryable
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyStateable
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationStateable
    // Possible ways of adding a mobile device to management.
    deviceEnrollmentType *DeviceEnrollmentType
    // Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
    deviceFirmwareConfigurationInterfaceManaged *bool
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState DeviceHealthAttestationStateable
    // Name of the device. This property is read-only.
    deviceName *string
    // Device registration status.
    deviceRegistrationState *DeviceRegistrationState
    // Device type.
    deviceType *DeviceType
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
    enrollmentProfileName *string
    // Ethernet MAC. This property is read-only.
    ethernetMacAddress *string
    // Device Exchange Access State.
    exchangeAccessState *DeviceManagementExchangeAccessState
    // Device Exchange Access State Reason.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free Storage in Bytes. This property is read-only.
    freeStorageSpaceInBytes *int64
    // The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
    hardwareInformation HardwareInformationable
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
    iccid *string
    // IMEI. This property is read-only.
    imei *string
    // Device encryption status. This property is read-only.
    isEncrypted *bool
    // Device supervised status. This property is read-only.
    isSupervised *bool
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string
    // Device enrollment join type.
    joinType *JoinType
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of log collection requests
    logCollectionRequests []DeviceLogCollectionResponseable
    // State of lost mode, indicating if lost mode is enabled or disabled
    lostModeState *LostModeState
    // Managed device mobile app configuration states for this device.
    managedDeviceMobileAppConfigurationStates []ManagedDeviceMobileAppConfigurationStateable
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string
    // Owner type of device.
    managedDeviceOwnerType *ManagedDeviceOwnerType
    // Management agent type.
    managementAgent *ManagementAgentType
    // Reports device management certificate expiration date. This property is read-only.
    managementCertificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device management features.
    managementFeatures *ManagedDeviceManagementFeatures
    // Management state of device in Microsoft Intune.
    managementState *ManagementState
    // Manufacturer of the device. This property is read-only.
    manufacturer *string
    // MEID. This property is read-only.
    meid *string
    // Model of the device. This property is read-only.
    model *string
    // Notes on the device created by IT Admin
    notes *string
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string
    // Operating system version of the device. This property is read-only.
    osVersion *string
    // Owner type of device.
    ownerType *OwnerType
    // Available health states for the Device Health API
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState
    // Phone number of the device. This property is read-only.
    phoneNumber *string
    // Total Memory in Bytes. This property is read-only.
    physicalMemoryInBytes *int64
    // Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
    preferMdmOverGroupPolicyAppliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Processor architecture
    processorArchitecture *ManagedDeviceArchitecture
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string
    // Reports if the managed iOS device is user approval enrollment. This property is read-only.
    requireUserEnrollmentApproval *bool
    // Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
    retireAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tag IDs for this Device instance.
    roleScopeTagIds []string
    // Security baseline states for this device.
    securityBaselineStates []SecurityBaselineStateable
    // SerialNumber. This property is read-only.
    serialNumber *string
    // Device sku family
    skuFamily *string
    // Device sku number, see also: https://docs.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
    skuNumber *int32
    // Specification version. This property is read-only.
    specificationVersion *string
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64
    // Unique Device Identifier for iOS and macOS devices. This property is read-only.
    udid *string
    // User display name. This property is read-only.
    userDisplayName *string
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string
    // Device user principal name. This property is read-only.
    userPrincipalName *string
    // The primary users associated with the managed device.
    users []Userable
    // Indicates the last logged on users of a device. This property is read-only.
    usersLoggedOn []LoggedOnUserable
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string
    // Count of active malware for this windows device. This property is read-only.
    windowsActiveMalwareCount *int32
    // The device protection status. This property is read-only.
    windowsProtectionState WindowsProtectionStateable
    // Count of remediated malware for this windows device. This property is read-only.
    windowsRemediatedMalwareCount *int32
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsManagedDevice":
                        return NewWindowsManagedDevice(), nil
                }
            }
        }
    }
    return NewManagedDevice(), nil
}
// GetAadRegistered gets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAadRegistered()(*bool) {
    return m.aadRegistered
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    return m.activationLockBypassCode
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    return m.androidSecurityPatchLevel
}
// GetAssignmentFilterEvaluationStatusDetails gets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetailsable) {
    return m.assignmentFilterEvaluationStatusDetails
}
// GetAutopilotEnrolled gets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) GetAutopilotEnrolled()(*bool) {
    return m.autopilotEnrolled
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureActiveDirectoryDeviceId()(*string) {
    return m.azureActiveDirectoryDeviceId
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    return m.azureADDeviceId
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    return m.azureADRegistered
}
// GetBootstrapTokenEscrowed gets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) GetBootstrapTokenEscrowed()(*bool) {
    return m.bootstrapTokenEscrowed
}
// GetChassisType gets the chassisType property value. Chassis type.
func (m *ManagedDevice) GetChassisType()(*ChassisType) {
    return m.chassisType
}
// GetChromeOSDeviceInfo gets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) GetChromeOSDeviceInfo()([]ChromeOSDevicePropertyable) {
    return m.chromeOSDeviceInfo
}
// GetCloudPcRemoteActionResults gets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResultable) {
    return m.cloudPcRemoteActionResults
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetComplianceState gets the complianceState property value. Compliance state.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    return m.complianceState
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    return m.configurationManagerClientEnabledFeatures
}
// GetConfigurationManagerClientHealthState gets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientHealthState()(ConfigurationManagerClientHealthStateable) {
    return m.configurationManagerClientHealthState
}
// GetConfigurationManagerClientInformation gets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientInformation()(ConfigurationManagerClientInformationable) {
    return m.configurationManagerClientInformation
}
// GetDetectedApps gets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) GetDetectedApps()([]DetectedAppable) {
    return m.detectedApps
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    return m.deviceActionResults
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    return m.deviceCategory
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    return m.deviceCategoryDisplayName
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    return m.deviceCompliancePolicyStates
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    return m.deviceConfigurationStates
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    return m.deviceEnrollmentType
}
// GetDeviceFirmwareConfigurationInterfaceManaged gets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) GetDeviceFirmwareConfigurationInterfaceManaged()(*bool) {
    return m.deviceFirmwareConfigurationInterfaceManaged
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    return m.deviceHealthAttestationState
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    return m.deviceRegistrationState
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ManagedDevice) GetDeviceType()(*DeviceType) {
    return m.deviceType
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    return m.easActivated
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.easActivationDateTime
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    return m.easDeviceId
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) GetEnrollmentProfileName()(*string) {
    return m.enrollmentProfileName
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    return m.ethernetMacAddress
}
// GetExchangeAccessState gets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    return m.exchangeAccessState
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    return m.exchangeAccessStateReason
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.exchangeLastSuccessfulSyncDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAadRegistered)
    res["activationLockBypassCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivationLockBypassCode)
    res["androidSecurityPatchLevel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAndroidSecurityPatchLevel)
    res["assignmentFilterEvaluationStatusDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue , m.SetAssignmentFilterEvaluationStatusDetails)
    res["autopilotEnrolled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAutopilotEnrolled)
    res["azureActiveDirectoryDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureActiveDirectoryDeviceId)
    res["azureADDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureADDeviceId)
    res["azureADRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAzureADRegistered)
    res["bootstrapTokenEscrowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBootstrapTokenEscrowed)
    res["chassisType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseChassisType , m.SetChassisType)
    res["chromeOSDeviceInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateChromeOSDevicePropertyFromDiscriminatorValue , m.SetChromeOSDeviceInfo)
    res["cloudPcRemoteActionResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcRemoteActionResultFromDiscriminatorValue , m.SetCloudPcRemoteActionResults)
    res["complianceGracePeriodExpirationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetComplianceGracePeriodExpirationDateTime)
    res["complianceState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComplianceState , m.SetComplianceState)
    res["configurationManagerClientEnabledFeatures"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue , m.SetConfigurationManagerClientEnabledFeatures)
    res["configurationManagerClientHealthState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConfigurationManagerClientHealthStateFromDiscriminatorValue , m.SetConfigurationManagerClientHealthState)
    res["configurationManagerClientInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateConfigurationManagerClientInformationFromDiscriminatorValue , m.SetConfigurationManagerClientInformation)
    res["detectedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue , m.SetDetectedApps)
    res["deviceActionResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue , m.SetDeviceActionResults)
    res["deviceCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceCategoryFromDiscriminatorValue , m.SetDeviceCategory)
    res["deviceCategoryDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceCategoryDisplayName)
    res["deviceCompliancePolicyStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue , m.SetDeviceCompliancePolicyStates)
    res["deviceConfigurationStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue , m.SetDeviceConfigurationStates)
    res["deviceEnrollmentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceEnrollmentType , m.SetDeviceEnrollmentType)
    res["deviceFirmwareConfigurationInterfaceManaged"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDeviceFirmwareConfigurationInterfaceManaged)
    res["deviceHealthAttestationState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue , m.SetDeviceHealthAttestationState)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["deviceRegistrationState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceRegistrationState , m.SetDeviceRegistrationState)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceType , m.SetDeviceType)
    res["easActivated"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEasActivated)
    res["easActivationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEasActivationDateTime)
    res["easDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEasDeviceId)
    res["emailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmailAddress)
    res["enrolledDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEnrolledDateTime)
    res["enrollmentProfileName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEnrollmentProfileName)
    res["ethernetMacAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEthernetMacAddress)
    res["exchangeAccessState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementExchangeAccessState , m.SetExchangeAccessState)
    res["exchangeAccessStateReason"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementExchangeAccessStateReason , m.SetExchangeAccessStateReason)
    res["exchangeLastSuccessfulSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetExchangeLastSuccessfulSyncDateTime)
    res["freeStorageSpaceInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetFreeStorageSpaceInBytes)
    res["hardwareInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateHardwareInformationFromDiscriminatorValue , m.SetHardwareInformation)
    res["iccid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIccid)
    res["imei"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetImei)
    res["isEncrypted"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEncrypted)
    res["isSupervised"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSupervised)
    res["jailBroken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJailBroken)
    res["joinType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseJoinType , m.SetJoinType)
    res["lastSyncDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastSyncDateTime)
    res["logCollectionRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceLogCollectionResponseFromDiscriminatorValue , m.SetLogCollectionRequests)
    res["lostModeState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLostModeState , m.SetLostModeState)
    res["managedDeviceMobileAppConfigurationStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue , m.SetManagedDeviceMobileAppConfigurationStates)
    res["managedDeviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceName)
    res["managedDeviceOwnerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDeviceOwnerType , m.SetManagedDeviceOwnerType)
    res["managementAgent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementAgentType , m.SetManagementAgent)
    res["managementCertificateExpirationDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetManagementCertificateExpirationDate)
    res["managementFeatures"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDeviceManagementFeatures , m.SetManagementFeatures)
    res["managementState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementState , m.SetManagementState)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["meid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMeid)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["operatingSystem"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOperatingSystem)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["ownerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOwnerType , m.SetOwnerType)
    res["partnerReportedThreatState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDevicePartnerReportedHealthState , m.SetPartnerReportedThreatState)
    res["phoneNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPhoneNumber)
    res["physicalMemoryInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetPhysicalMemoryInBytes)
    res["preferMdmOverGroupPolicyAppliedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetPreferMdmOverGroupPolicyAppliedDateTime)
    res["processorArchitecture"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagedDeviceArchitecture , m.SetProcessorArchitecture)
    res["remoteAssistanceSessionErrorDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteAssistanceSessionErrorDetails)
    res["remoteAssistanceSessionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteAssistanceSessionUrl)
    res["requireUserEnrollmentApproval"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRequireUserEnrollmentApproval)
    res["retireAfterDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRetireAfterDateTime)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["securityBaselineStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSecurityBaselineStateFromDiscriminatorValue , m.SetSecurityBaselineStates)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    res["skuFamily"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSkuFamily)
    res["skuNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSkuNumber)
    res["specificationVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSpecificationVersion)
    res["subscriberCarrier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriberCarrier)
    res["totalStorageSpaceInBytes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTotalStorageSpaceInBytes)
    res["udid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUdid)
    res["userDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserDisplayName)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["users"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserFromDiscriminatorValue , m.SetUsers)
    res["usersLoggedOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLoggedOnUserFromDiscriminatorValue , m.SetUsersLoggedOn)
    res["wiFiMacAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWiFiMacAddress)
    res["windowsActiveMalwareCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindowsActiveMalwareCount)
    res["windowsProtectionState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateWindowsProtectionStateFromDiscriminatorValue , m.SetWindowsProtectionState)
    res["windowsRemediatedMalwareCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindowsRemediatedMalwareCount)
    return res
}
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    return m.freeStorageSpaceInBytes
}
// GetHardwareInformation gets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
func (m *ManagedDevice) GetHardwareInformation()(HardwareInformationable) {
    return m.hardwareInformation
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    return m.iccid
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    return m.imei
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetJailBroken gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    return m.jailBroken
}
// GetJoinType gets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) GetJoinType()(*JoinType) {
    return m.joinType
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetLogCollectionRequests gets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) GetLogCollectionRequests()([]DeviceLogCollectionResponseable) {
    return m.logCollectionRequests
}
// GetLostModeState gets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) GetLostModeState()(*LostModeState) {
    return m.lostModeState
}
// GetManagedDeviceMobileAppConfigurationStates gets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationStateable) {
    return m.managedDeviceMobileAppConfigurationStates
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    return m.managedDeviceOwnerType
}
// GetManagementAgent gets the managementAgent property value. Management agent type.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    return m.managementAgent
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.managementCertificateExpirationDate
}
// GetManagementFeatures gets the managementFeatures property value. Device management features.
func (m *ManagedDevice) GetManagementFeatures()(*ManagedDeviceManagementFeatures) {
    return m.managementFeatures
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) GetManagementState()(*ManagementState) {
    return m.managementState
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    return m.meid
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    return m.model
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) GetNotes()(*string) {
    return m.notes
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *ManagedDevice) GetOwnerType()(*OwnerType) {
    return m.ownerType
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    return m.partnerReportedThreatState
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    return m.physicalMemoryInBytes
}
// GetPreferMdmOverGroupPolicyAppliedDateTime gets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.preferMdmOverGroupPolicyAppliedDateTime
}
// GetProcessorArchitecture gets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) GetProcessorArchitecture()(*ManagedDeviceArchitecture) {
    return m.processorArchitecture
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    return m.remoteAssistanceSessionErrorDetails
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    return m.remoteAssistanceSessionUrl
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    return m.requireUserEnrollmentApproval
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.retireAfterDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSecurityBaselineStates gets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) GetSecurityBaselineStates()([]SecurityBaselineStateable) {
    return m.securityBaselineStates
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSkuFamily gets the skuFamily property value. Device sku family
func (m *ManagedDevice) GetSkuFamily()(*string) {
    return m.skuFamily
}
// GetSkuNumber gets the skuNumber property value. Device sku number, see also: https://docs.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) GetSkuNumber()(*int32) {
    return m.skuNumber
}
// GetSpecificationVersion gets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) GetSpecificationVersion()(*string) {
    return m.specificationVersion
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    return m.subscriberCarrier
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    return m.totalStorageSpaceInBytes
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    return m.udid
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUsers gets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) GetUsers()([]Userable) {
    return m.users
}
// GetUsersLoggedOn gets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) GetUsersLoggedOn()([]LoggedOnUserable) {
    return m.usersLoggedOn
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    return m.wiFiMacAddress
}
// GetWindowsActiveMalwareCount gets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsActiveMalwareCount()(*int32) {
    return m.windowsActiveMalwareCount
}
// GetWindowsProtectionState gets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) GetWindowsProtectionState()(WindowsProtectionStateable) {
    return m.windowsProtectionState
}
// GetWindowsRemediatedMalwareCount gets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsRemediatedMalwareCount()(*int32) {
    return m.windowsRemediatedMalwareCount
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentFilterEvaluationStatusDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignmentFilterEvaluationStatusDetails())
        err = writer.WriteCollectionOfObjectValues("assignmentFilterEvaluationStatusDetails", cast)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChromeOSDeviceInfo())
        err = writer.WriteCollectionOfObjectValues("chromeOSDeviceInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcRemoteActionResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCloudPcRemoteActionResults())
        err = writer.WriteCollectionOfObjectValues("cloudPcRemoteActionResults", cast)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDetectedApps())
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
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
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceCompliancePolicyStates())
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceConfigurationStates())
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
        err = writer.WriteBoolValue("deviceFirmwareConfigurationInterfaceManaged", m.GetDeviceFirmwareConfigurationInterfaceManaged())
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
    if m.GetJoinType() != nil {
        cast := (*m.GetJoinType()).String()
        err = writer.WriteStringValue("joinType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLogCollectionRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLogCollectionRequests())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceMobileAppConfigurationStates())
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
        err = writer.WriteStringValue("notes", m.GetNotes())
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
    if m.GetProcessorArchitecture() != nil {
        cast := (*m.GetProcessorArchitecture()).String()
        err = writer.WriteStringValue("processorArchitecture", &cast)
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSecurityBaselineStates())
        err = writer.WriteCollectionOfObjectValues("securityBaselineStates", cast)
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
    if m.GetUsers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUsers())
        err = writer.WriteCollectionOfObjectValues("users", cast)
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
    return nil
}
// SetAadRegistered sets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAadRegistered(value *bool)() {
    m.aadRegistered = value
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. Code that allows the Activation Lock on a device to be bypassed. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
// SetAssignmentFilterEvaluationStatusDetails sets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetailsable)() {
    m.assignmentFilterEvaluationStatusDetails = value
}
// SetAutopilotEnrolled sets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) SetAutopilotEnrolled(value *bool)() {
    m.autopilotEnrolled = value
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureActiveDirectoryDeviceId(value *string)() {
    m.azureActiveDirectoryDeviceId = value
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
// SetBootstrapTokenEscrowed sets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) SetBootstrapTokenEscrowed(value *bool)() {
    m.bootstrapTokenEscrowed = value
}
// SetChassisType sets the chassisType property value. Chassis type.
func (m *ManagedDevice) SetChassisType(value *ChassisType)() {
    m.chassisType = value
}
// SetChromeOSDeviceInfo sets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) SetChromeOSDeviceInfo(value []ChromeOSDevicePropertyable)() {
    m.chromeOSDeviceInfo = value
}
// SetCloudPcRemoteActionResults sets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResultable)() {
    m.cloudPcRemoteActionResults = value
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetComplianceState sets the complianceState property value. Compliance state.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    m.configurationManagerClientEnabledFeatures = value
}
// SetConfigurationManagerClientHealthState sets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientHealthState(value ConfigurationManagerClientHealthStateable)() {
    m.configurationManagerClientHealthState = value
}
// SetConfigurationManagerClientInformation sets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientInformation(value ConfigurationManagerClientInformationable)() {
    m.configurationManagerClientInformation = value
}
// SetDetectedApps sets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) SetDetectedApps(value []DetectedAppable)() {
    m.detectedApps = value
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    m.deviceActionResults = value
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    m.deviceCategory = value
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    m.deviceCompliancePolicyStates = value
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    m.deviceConfigurationStates = value
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
// SetDeviceFirmwareConfigurationInterfaceManaged sets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) SetDeviceFirmwareConfigurationInterfaceManaged(value *bool)() {
    m.deviceFirmwareConfigurationInterfaceManaged = value
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    m.deviceHealthAttestationState = value
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ManagedDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Ethernet MAC. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
// SetExchangeAccessState sets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
// SetHardwareInformation sets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. This property is read-only.
func (m *ManagedDevice) SetHardwareInformation(value HardwareInformationable)() {
    m.hardwareInformation = value
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetJailBroken sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
// SetJoinType sets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) SetJoinType(value *JoinType)() {
    m.joinType = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetLogCollectionRequests sets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) SetLogCollectionRequests(value []DeviceLogCollectionResponseable)() {
    m.logCollectionRequests = value
}
// SetLostModeState sets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) SetLostModeState(value *LostModeState)() {
    m.lostModeState = value
}
// SetManagedDeviceMobileAppConfigurationStates sets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationStateable)() {
    m.managedDeviceMobileAppConfigurationStates = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
// SetManagementAgent sets the managementAgent property value. Management agent type.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managementCertificateExpirationDate = value
}
// SetManagementFeatures sets the managementFeatures property value. Device management features.
func (m *ManagedDevice) SetManagementFeatures(value *ManagedDeviceManagementFeatures)() {
    m.managementFeatures = value
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ManagedDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
// SetPreferMdmOverGroupPolicyAppliedDateTime sets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.preferMdmOverGroupPolicyAppliedDateTime = value
}
// SetProcessorArchitecture sets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) SetProcessorArchitecture(value *ManagedDeviceArchitecture)() {
    m.processorArchitecture = value
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    m.requireUserEnrollmentApproval = value
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.retireAfterDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSecurityBaselineStates sets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) SetSecurityBaselineStates(value []SecurityBaselineStateable)() {
    m.securityBaselineStates = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSkuFamily sets the skuFamily property value. Device sku family
func (m *ManagedDevice) SetSkuFamily(value *string)() {
    m.skuFamily = value
}
// SetSkuNumber sets the skuNumber property value. Device sku number, see also: https://docs.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) SetSkuNumber(value *int32)() {
    m.skuNumber = value
}
// SetSpecificationVersion sets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) SetSpecificationVersion(value *string)() {
    m.specificationVersion = value
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []Userable)() {
    m.users = value
}
// SetUsersLoggedOn sets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) SetUsersLoggedOn(value []LoggedOnUserable)() {
    m.usersLoggedOn = value
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
// SetWindowsActiveMalwareCount sets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsActiveMalwareCount(value *int32)() {
    m.windowsActiveMalwareCount = value
}
// SetWindowsProtectionState sets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) SetWindowsProtectionState(value WindowsProtectionStateable)() {
    m.windowsProtectionState = value
}
// SetWindowsRemediatedMalwareCount sets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsRemediatedMalwareCount(value *int32)() {
    m.windowsRemediatedMalwareCount = value
}
