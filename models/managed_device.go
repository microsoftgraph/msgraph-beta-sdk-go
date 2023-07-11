package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
    Entity
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
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
    val, err := m.GetBackingStore().Get("aadRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    val, err := m.GetBackingStore().Get("activationLockBypassCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    val, err := m.GetBackingStore().Get("androidSecurityPatchLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAssignmentFilterEvaluationStatusDetails gets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetailsable) {
    val, err := m.GetBackingStore().Get("assignmentFilterEvaluationStatusDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignmentFilterEvaluationStatusDetailsable)
    }
    return nil
}
// GetAutopilotEnrolled gets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) GetAutopilotEnrolled()(*bool) {
    val, err := m.GetBackingStore().Get("autopilotEnrolled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureActiveDirectoryDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureActiveDirectoryDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureADDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("azureADRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBootstrapTokenEscrowed gets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) GetBootstrapTokenEscrowed()(*bool) {
    val, err := m.GetBackingStore().Get("bootstrapTokenEscrowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetChassisType gets the chassisType property value. Chassis type.
func (m *ManagedDevice) GetChassisType()(*ChassisType) {
    val, err := m.GetBackingStore().Get("chassisType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ChassisType)
    }
    return nil
}
// GetChromeOSDeviceInfo gets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device. Default is an empty list. To retrieve actual values GET call needs to be made, with device id and included in select parameter.
func (m *ManagedDevice) GetChromeOSDeviceInfo()([]ChromeOSDevicePropertyable) {
    val, err := m.GetBackingStore().Get("chromeOSDeviceInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ChromeOSDevicePropertyable)
    }
    return nil
}
// GetCloudPcRemoteActionResults gets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResultable) {
    val, err := m.GetBackingStore().Get("cloudPcRemoteActionResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudPcRemoteActionResultable)
    }
    return nil
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("complianceGracePeriodExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetComplianceState gets the complianceState property value. Compliance state.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    val, err := m.GetBackingStore().Get("complianceState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComplianceState)
    }
    return nil
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    val, err := m.GetBackingStore().Get("configurationManagerClientEnabledFeatures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConfigurationManagerClientEnabledFeaturesable)
    }
    return nil
}
// GetConfigurationManagerClientHealthState gets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientHealthState()(ConfigurationManagerClientHealthStateable) {
    val, err := m.GetBackingStore().Get("configurationManagerClientHealthState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConfigurationManagerClientHealthStateable)
    }
    return nil
}
// GetConfigurationManagerClientInformation gets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientInformation()(ConfigurationManagerClientInformationable) {
    val, err := m.GetBackingStore().Get("configurationManagerClientInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ConfigurationManagerClientInformationable)
    }
    return nil
}
// GetDetectedApps gets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) GetDetectedApps()([]DetectedAppable) {
    val, err := m.GetBackingStore().Get("detectedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DetectedAppable)
    }
    return nil
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    val, err := m.GetBackingStore().Get("deviceActionResults")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceActionResultable)
    }
    return nil
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    val, err := m.GetBackingStore().Get("deviceCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceCategoryable)
    }
    return nil
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("deviceCategoryDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    val, err := m.GetBackingStore().Get("deviceCompliancePolicyStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceCompliancePolicyStateable)
    }
    return nil
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    val, err := m.GetBackingStore().Get("deviceConfigurationStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceConfigurationStateable)
    }
    return nil
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    val, err := m.GetBackingStore().Get("deviceEnrollmentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceEnrollmentType)
    }
    return nil
}
// GetDeviceFirmwareConfigurationInterfaceManaged gets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) GetDeviceFirmwareConfigurationInterfaceManaged()(*bool) {
    val, err := m.GetBackingStore().Get("deviceFirmwareConfigurationInterfaceManaged")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    val, err := m.GetBackingStore().Get("deviceHealthAttestationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceHealthAttestationStateable)
    }
    return nil
}
// GetDeviceHealthScriptStates gets the deviceHealthScriptStates property value. Results of device health scripts that ran for this device. Default is empty list. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthScriptStates()([]DeviceHealthScriptPolicyStateable) {
    val, err := m.GetBackingStore().Get("deviceHealthScriptStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceHealthScriptPolicyStateable)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    val, err := m.GetBackingStore().Get("deviceRegistrationState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceRegistrationState)
    }
    return nil
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ManagedDevice) GetDeviceType()(*DeviceType) {
    val, err := m.GetBackingStore().Get("deviceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceType)
    }
    return nil
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    val, err := m.GetBackingStore().Get("easActivated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("easActivationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("easDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("emailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("enrolledDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) GetEnrollmentProfileName()(*string) {
    val, err := m.GetBackingStore().Get("enrollmentProfileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    val, err := m.GetBackingStore().Get("ethernetMacAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExchangeAccessState gets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    val, err := m.GetBackingStore().Get("exchangeAccessState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementExchangeAccessState)
    }
    return nil
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    val, err := m.GetBackingStore().Get("exchangeAccessStateReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementExchangeAccessStateReason)
    }
    return nil
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("exchangeLastSuccessfulSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadRegistered(val)
        }
        return nil
    }
    res["activationLockBypassCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["assignmentFilterEvaluationStatusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationStatusDetailsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignmentFilterEvaluationStatusDetailsable)
                }
            }
            m.SetAssignmentFilterEvaluationStatusDetails(res)
        }
        return nil
    }
    res["autopilotEnrolled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopilotEnrolled(val)
        }
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureActiveDirectoryDeviceId(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["bootstrapTokenEscrowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootstrapTokenEscrowed(val)
        }
        return nil
    }
    res["chassisType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChassisType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChassisType(val.(*ChassisType))
        }
        return nil
    }
    res["chromeOSDeviceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChromeOSDevicePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChromeOSDevicePropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ChromeOSDevicePropertyable)
                }
            }
            m.SetChromeOSDeviceInfo(res)
        }
        return nil
    }
    res["cloudPcRemoteActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcRemoteActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcRemoteActionResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudPcRemoteActionResultable)
                }
            }
            m.SetCloudPcRemoteActionResults(res)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceState))
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(ConfigurationManagerClientEnabledFeaturesable))
        }
        return nil
    }
    res["configurationManagerClientHealthState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientHealthStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientHealthState(val.(ConfigurationManagerClientHealthStateable))
        }
        return nil
    }
    res["configurationManagerClientInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientInformation(val.(ConfigurationManagerClientInformationable))
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
    res["deviceActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceActionResultable)
                }
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(DeviceCategoryable))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceCompliancePolicyStateable)
                }
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceConfigurationStateable)
                }
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["deviceFirmwareConfigurationInterfaceManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFirmwareConfigurationInterfaceManaged(val)
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(DeviceHealthAttestationStateable))
        }
        return nil
    }
    res["deviceHealthScriptStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceHealthScriptPolicyStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceHealthScriptPolicyStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceHealthScriptPolicyStateable)
                }
            }
            m.SetDeviceHealthScriptStates(res)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationState(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["easActivated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessState(val.(*DeviceManagementExchangeAccessState))
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessStateReason(val.(*DeviceManagementExchangeAccessStateReason))
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["hardwareInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHardwareInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareInformation(val.(HardwareInformationable))
        }
        return nil
    }
    res["iccid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["joinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinType(val.(*JoinType))
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["logCollectionRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLogCollectionResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLogCollectionResponseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceLogCollectionResponseable)
                }
            }
            m.SetLogCollectionRequests(res)
        }
        return nil
    }
    res["lostModeState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLostModeState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLostModeState(val.(*LostModeState))
        }
        return nil
    }
    res["managedDeviceMobileAppConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ManagedDeviceMobileAppConfigurationStateable)
                }
            }
            m.SetManagedDeviceMobileAppConfigurationStates(res)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["managementAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementCertificateExpirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementCertificateExpirationDate(val)
        }
        return nil
    }
    res["managementFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceManagementFeatures)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementFeatures(val.(*ManagedDeviceManagementFeatures))
        }
        return nil
    }
    res["managementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerReportedThreatState(val.(*ManagedDevicePartnerReportedHealthState))
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["preferMdmOverGroupPolicyAppliedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferMdmOverGroupPolicyAppliedDateTime(val)
        }
        return nil
    }
    res["processorArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorArchitecture(val.(*ManagedDeviceArchitecture))
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["requireUserEnrollmentApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserEnrollmentApproval(val)
        }
        return nil
    }
    res["retireAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetireAfterDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["securityBaselineStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecurityBaselineStateable)
                }
            }
            m.SetSecurityBaselineStates(res)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["skuFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuFamily(val)
        }
        return nil
    }
    res["skuNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["specificationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecificationVersion(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Userable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    res["usersLoggedOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLoggedOnUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LoggedOnUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LoggedOnUserable)
                }
            }
            m.SetUsersLoggedOn(res)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiMacAddress(val)
        }
        return nil
    }
    res["windowsActiveMalwareCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsActiveMalwareCount(val)
        }
        return nil
    }
    res["windowsProtectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsProtectionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsProtectionState(val.(WindowsProtectionStateable))
        }
        return nil
    }
    res["windowsRemediatedMalwareCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("freeStorageSpaceInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetHardwareInformation gets the hardwareInformation property value. The hardward details for the device. Includes information such as storage space, manufacturer, serial number, etc. By default most property of this type are set to null/0/false and enum defaults for associated types. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetHardwareInformation()(HardwareInformationable) {
    val, err := m.GetBackingStore().Get("hardwareInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(HardwareInformationable)
    }
    return nil
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    val, err := m.GetBackingStore().Get("iccid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    val, err := m.GetBackingStore().Get("imei")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    val, err := m.GetBackingStore().Get("isEncrypted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    val, err := m.GetBackingStore().Get("isSupervised")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJailBroken gets the jailBroken property value. Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    val, err := m.GetBackingStore().Get("jailBroken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetJoinType gets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) GetJoinType()(*JoinType) {
    val, err := m.GetBackingStore().Get("joinType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*JoinType)
    }
    return nil
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLogCollectionRequests gets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) GetLogCollectionRequests()([]DeviceLogCollectionResponseable) {
    val, err := m.GetBackingStore().Get("logCollectionRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceLogCollectionResponseable)
    }
    return nil
}
// GetLostModeState gets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) GetLostModeState()(*LostModeState) {
    val, err := m.GetBackingStore().Get("lostModeState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LostModeState)
    }
    return nil
}
// GetManagedDeviceMobileAppConfigurationStates gets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceMobileAppConfigurationStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceMobileAppConfigurationStateable)
    }
    return nil
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    val, err := m.GetBackingStore().Get("managedDeviceOwnerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceOwnerType)
    }
    return nil
}
// GetManagementAgent gets the managementAgent property value. Management agent type.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    val, err := m.GetBackingStore().Get("managementAgent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementAgentType)
    }
    return nil
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("managementCertificateExpirationDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetManagementFeatures gets the managementFeatures property value. Device management features.
func (m *ManagedDevice) GetManagementFeatures()(*ManagedDeviceManagementFeatures) {
    val, err := m.GetBackingStore().Get("managementFeatures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceManagementFeatures)
    }
    return nil
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) GetManagementState()(*ManagementState) {
    val, err := m.GetBackingStore().Get("managementState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementState)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    val, err := m.GetBackingStore().Get("meid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
func (m *ManagedDevice) GetNotes()(*string) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagedDevice) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *ManagedDevice) GetOwnerType()(*OwnerType) {
    val, err := m.GetBackingStore().Get("ownerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OwnerType)
    }
    return nil
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    val, err := m.GetBackingStore().Get("partnerReportedThreatState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDevicePartnerReportedHealthState)
    }
    return nil
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. Read-only. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("physicalMemoryInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetPreferMdmOverGroupPolicyAppliedDateTime gets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("preferMdmOverGroupPolicyAppliedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetProcessorArchitecture gets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) GetProcessorArchitecture()(*ManagedDeviceArchitecture) {
    val, err := m.GetBackingStore().Get("processorArchitecture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceArchitecture)
    }
    return nil
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    val, err := m.GetBackingStore().Get("remoteAssistanceSessionErrorDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    val, err := m.GetBackingStore().Get("remoteAssistanceSessionUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    val, err := m.GetBackingStore().Get("requireUserEnrollmentApproval")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("retireAfterDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSecurityBaselineStates gets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) GetSecurityBaselineStates()([]SecurityBaselineStateable) {
    val, err := m.GetBackingStore().Get("securityBaselineStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecurityBaselineStateable)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkuFamily gets the skuFamily property value. Device sku family
func (m *ManagedDevice) GetSkuFamily()(*string) {
    val, err := m.GetBackingStore().Get("skuFamily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkuNumber gets the skuNumber property value. Device sku number, see also: https://learn.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) GetSkuNumber()(*int32) {
    val, err := m.GetBackingStore().Get("skuNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSpecificationVersion gets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) GetSpecificationVersion()(*string) {
    val, err := m.GetBackingStore().Get("specificationVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    val, err := m.GetBackingStore().Get("subscriberCarrier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("totalStorageSpaceInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    val, err := m.GetBackingStore().Get("udid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("userDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsers gets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) GetUsers()([]Userable) {
    val, err := m.GetBackingStore().Get("users")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Userable)
    }
    return nil
}
// GetUsersLoggedOn gets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) GetUsersLoggedOn()([]LoggedOnUserable) {
    val, err := m.GetBackingStore().Get("usersLoggedOn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LoggedOnUserable)
    }
    return nil
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    val, err := m.GetBackingStore().Get("wiFiMacAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWindowsActiveMalwareCount gets the windowsActiveMalwareCount property value. Count of active malware for this windows device. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) GetWindowsActiveMalwareCount()(*int32) {
    val, err := m.GetBackingStore().Get("windowsActiveMalwareCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWindowsProtectionState gets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) GetWindowsProtectionState()(WindowsProtectionStateable) {
    val, err := m.GetBackingStore().Get("windowsProtectionState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsProtectionStateable)
    }
    return nil
}
// GetWindowsRemediatedMalwareCount gets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) GetWindowsRemediatedMalwareCount()(*int32) {
    val, err := m.GetBackingStore().Get("windowsRemediatedMalwareCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentFilterEvaluationStatusDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilterEvaluationStatusDetails()))
        for i, v := range m.GetAssignmentFilterEvaluationStatusDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChromeOSDeviceInfo()))
        for i, v := range m.GetChromeOSDeviceInfo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSDeviceInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcRemoteActionResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcRemoteActionResults()))
        for i, v := range m.GetCloudPcRemoteActionResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteBoolValue("deviceFirmwareConfigurationInterfaceManaged", m.GetDeviceFirmwareConfigurationInterfaceManaged())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceHealthScriptStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceHealthScriptStates()))
        for i, v := range m.GetDeviceHealthScriptStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceHealthScriptStates", cast)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogCollectionRequests()))
        for i, v := range m.GetLogCollectionRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceMobileAppConfigurationStates()))
        for i, v := range m.GetManagedDeviceMobileAppConfigurationStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecurityBaselineStates()))
        for i, v := range m.GetSecurityBaselineStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("aadRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    err := m.GetBackingStore().Set("activationLockBypassCode", value)
    if err != nil {
        panic(err)
    }
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    err := m.GetBackingStore().Set("androidSecurityPatchLevel", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentFilterEvaluationStatusDetails sets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetailsable)() {
    err := m.GetBackingStore().Set("assignmentFilterEvaluationStatusDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetAutopilotEnrolled sets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) SetAutopilotEnrolled(value *bool)() {
    err := m.GetBackingStore().Set("autopilotEnrolled", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureActiveDirectoryDeviceId(value *string)() {
    err := m.GetBackingStore().Set("azureActiveDirectoryDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    err := m.GetBackingStore().Set("azureADDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    err := m.GetBackingStore().Set("azureADRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetBootstrapTokenEscrowed sets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) SetBootstrapTokenEscrowed(value *bool)() {
    err := m.GetBackingStore().Set("bootstrapTokenEscrowed", value)
    if err != nil {
        panic(err)
    }
}
// SetChassisType sets the chassisType property value. Chassis type.
func (m *ManagedDevice) SetChassisType(value *ChassisType)() {
    err := m.GetBackingStore().Set("chassisType", value)
    if err != nil {
        panic(err)
    }
}
// SetChromeOSDeviceInfo sets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device. Default is an empty list. To retrieve actual values GET call needs to be made, with device id and included in select parameter.
func (m *ManagedDevice) SetChromeOSDeviceInfo(value []ChromeOSDevicePropertyable)() {
    err := m.GetBackingStore().Set("chromeOSDeviceInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcRemoteActionResults sets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResultable)() {
    err := m.GetBackingStore().Set("cloudPcRemoteActionResults", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("complianceGracePeriodExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceState sets the complianceState property value. Compliance state.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    err := m.GetBackingStore().Set("complianceState", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    err := m.GetBackingStore().Set("configurationManagerClientEnabledFeatures", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationManagerClientHealthState sets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientHealthState(value ConfigurationManagerClientHealthStateable)() {
    err := m.GetBackingStore().Set("configurationManagerClientHealthState", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationManagerClientInformation sets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientInformation(value ConfigurationManagerClientInformationable)() {
    err := m.GetBackingStore().Set("configurationManagerClientInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectedApps sets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) SetDetectedApps(value []DetectedAppable)() {
    err := m.GetBackingStore().Set("detectedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    err := m.GetBackingStore().Set("deviceActionResults", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    err := m.GetBackingStore().Set("deviceCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    err := m.GetBackingStore().Set("deviceCategoryDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    err := m.GetBackingStore().Set("deviceCompliancePolicyStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    err := m.GetBackingStore().Set("deviceConfigurationStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    err := m.GetBackingStore().Set("deviceEnrollmentType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceFirmwareConfigurationInterfaceManaged sets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) SetDeviceFirmwareConfigurationInterfaceManaged(value *bool)() {
    err := m.GetBackingStore().Set("deviceFirmwareConfigurationInterfaceManaged", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    err := m.GetBackingStore().Set("deviceHealthAttestationState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceHealthScriptStates sets the deviceHealthScriptStates property value. Results of device health scripts that ran for this device. Default is empty list. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthScriptStates(value []DeviceHealthScriptPolicyStateable)() {
    err := m.GetBackingStore().Set("deviceHealthScriptStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    err := m.GetBackingStore().Set("deviceRegistrationState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ManagedDevice) SetDeviceType(value *DeviceType)() {
    err := m.GetBackingStore().Set("deviceType", value)
    if err != nil {
        panic(err)
    }
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    err := m.GetBackingStore().Set("easActivated", value)
    if err != nil {
        panic(err)
    }
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("easActivationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    err := m.GetBackingStore().Set("easDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("emailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("enrolledDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) SetEnrollmentProfileName(value *string)() {
    err := m.GetBackingStore().Set("enrollmentProfileName", value)
    if err != nil {
        panic(err)
    }
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Indicates Ethernet MAC Address of the device. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    err := m.GetBackingStore().Set("ethernetMacAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeAccessState sets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    err := m.GetBackingStore().Set("exchangeAccessState", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    err := m.GetBackingStore().Set("exchangeAccessStateReason", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("exchangeLastSuccessfulSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    err := m.GetBackingStore().Set("freeStorageSpaceInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareInformation sets the hardwareInformation property value. The hardward details for the device. Includes information such as storage space, manufacturer, serial number, etc. By default most property of this type are set to null/0/false and enum defaults for associated types. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetHardwareInformation(value HardwareInformationable)() {
    err := m.GetBackingStore().Set("hardwareInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    err := m.GetBackingStore().Set("iccid", value)
    if err != nil {
        panic(err)
    }
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    err := m.GetBackingStore().Set("imei", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    err := m.GetBackingStore().Set("isEncrypted", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    err := m.GetBackingStore().Set("isSupervised", value)
    if err != nil {
        panic(err)
    }
}
// SetJailBroken sets the jailBroken property value. Whether the device is jail broken or rooted. Default is an empty string. Supports $filter operator 'eq' and 'or'. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    err := m.GetBackingStore().Set("jailBroken", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinType sets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) SetJoinType(value *JoinType)() {
    err := m.GetBackingStore().Set("joinType", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. Supports $filter operator 'lt' and 'gt'. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLogCollectionRequests sets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) SetLogCollectionRequests(value []DeviceLogCollectionResponseable)() {
    err := m.GetBackingStore().Set("logCollectionRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetLostModeState sets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) SetLostModeState(value *LostModeState)() {
    err := m.GetBackingStore().Set("lostModeState", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceMobileAppConfigurationStates sets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationStateable)() {
    err := m.GetBackingStore().Set("managedDeviceMobileAppConfigurationStates", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    err := m.GetBackingStore().Set("managedDeviceOwnerType", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementAgent sets the managementAgent property value. Management agent type.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    err := m.GetBackingStore().Set("managementAgent", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("managementCertificateExpirationDate", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementFeatures sets the managementFeatures property value. Device management features.
func (m *ManagedDevice) SetManagementFeatures(value *ManagedDeviceManagementFeatures)() {
    err := m.GetBackingStore().Set("managementFeatures", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) SetManagementState(value *ManagementState)() {
    err := m.GetBackingStore().Set("managementState", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    err := m.GetBackingStore().Set("meid", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin. Default is null. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported.
func (m *ManagedDevice) SetNotes(value *string)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDevice) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    err := m.GetBackingStore().Set("operatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ManagedDevice) SetOwnerType(value *OwnerType)() {
    err := m.GetBackingStore().Set("ownerType", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    err := m.GetBackingStore().Set("partnerReportedThreatState", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. Read-only. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    err := m.GetBackingStore().Set("physicalMemoryInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetPreferMdmOverGroupPolicyAppliedDateTime sets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("preferMdmOverGroupPolicyAppliedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorArchitecture sets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) SetProcessorArchitecture(value *ManagedDeviceArchitecture)() {
    err := m.GetBackingStore().Set("processorArchitecture", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    err := m.GetBackingStore().Set("remoteAssistanceSessionErrorDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    err := m.GetBackingStore().Set("remoteAssistanceSessionUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    err := m.GetBackingStore().Set("requireUserEnrollmentApproval", value)
    if err != nil {
        panic(err)
    }
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("retireAfterDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityBaselineStates sets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) SetSecurityBaselineStates(value []SecurityBaselineStateable)() {
    err := m.GetBackingStore().Set("securityBaselineStates", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuFamily sets the skuFamily property value. Device sku family
func (m *ManagedDevice) SetSkuFamily(value *string)() {
    err := m.GetBackingStore().Set("skuFamily", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuNumber sets the skuNumber property value. Device sku number, see also: https://learn.microsoft.com/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) SetSkuNumber(value *int32)() {
    err := m.GetBackingStore().Set("skuNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSpecificationVersion sets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) SetSpecificationVersion(value *string)() {
    err := m.GetBackingStore().Set("specificationVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    err := m.GetBackingStore().Set("subscriberCarrier", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    err := m.GetBackingStore().Set("totalStorageSpaceInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. Default is an empty string. To retrieve actual values GET call needs to be made, with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    err := m.GetBackingStore().Set("udid", value)
    if err != nil {
        panic(err)
    }
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    err := m.GetBackingStore().Set("userDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []Userable)() {
    err := m.GetBackingStore().Set("users", value)
    if err != nil {
        panic(err)
    }
}
// SetUsersLoggedOn sets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) SetUsersLoggedOn(value []LoggedOnUserable)() {
    err := m.GetBackingStore().Set("usersLoggedOn", value)
    if err != nil {
        panic(err)
    }
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    err := m.GetBackingStore().Set("wiFiMacAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsActiveMalwareCount sets the windowsActiveMalwareCount property value. Count of active malware for this windows device. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) SetWindowsActiveMalwareCount(value *int32)() {
    err := m.GetBackingStore().Set("windowsActiveMalwareCount", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsProtectionState sets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) SetWindowsProtectionState(value WindowsProtectionStateable)() {
    err := m.GetBackingStore().Set("windowsProtectionState", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsRemediatedMalwareCount sets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. Default is 0. To retrieve actual values GET call needs to be made, with device id and included in select parameter. This property is read-only.
func (m *ManagedDevice) SetWindowsRemediatedMalwareCount(value *int32)() {
    err := m.GetBackingStore().Set("windowsRemediatedMalwareCount", value)
    if err != nil {
        panic(err)
    }
}
// ManagedDeviceable 
type ManagedDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAadRegistered()(*bool)
    GetActivationLockBypassCode()(*string)
    GetAndroidSecurityPatchLevel()(*string)
    GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetailsable)
    GetAutopilotEnrolled()(*bool)
    GetAzureActiveDirectoryDeviceId()(*string)
    GetAzureADDeviceId()(*string)
    GetAzureADRegistered()(*bool)
    GetBootstrapTokenEscrowed()(*bool)
    GetChassisType()(*ChassisType)
    GetChromeOSDeviceInfo()([]ChromeOSDevicePropertyable)
    GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResultable)
    GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetComplianceState()(*ComplianceState)
    GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable)
    GetConfigurationManagerClientHealthState()(ConfigurationManagerClientHealthStateable)
    GetConfigurationManagerClientInformation()(ConfigurationManagerClientInformationable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceActionResults()([]DeviceActionResultable)
    GetDeviceCategory()(DeviceCategoryable)
    GetDeviceCategoryDisplayName()(*string)
    GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable)
    GetDeviceConfigurationStates()([]DeviceConfigurationStateable)
    GetDeviceEnrollmentType()(*DeviceEnrollmentType)
    GetDeviceFirmwareConfigurationInterfaceManaged()(*bool)
    GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable)
    GetDeviceHealthScriptStates()([]DeviceHealthScriptPolicyStateable)
    GetDeviceName()(*string)
    GetDeviceRegistrationState()(*DeviceRegistrationState)
    GetDeviceType()(*DeviceType)
    GetEasActivated()(*bool)
    GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasDeviceId()(*string)
    GetEmailAddress()(*string)
    GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEnrollmentProfileName()(*string)
    GetEthernetMacAddress()(*string)
    GetExchangeAccessState()(*DeviceManagementExchangeAccessState)
    GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason)
    GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFreeStorageSpaceInBytes()(*int64)
    GetHardwareInformation()(HardwareInformationable)
    GetIccid()(*string)
    GetImei()(*string)
    GetIsEncrypted()(*bool)
    GetIsSupervised()(*bool)
    GetJailBroken()(*string)
    GetJoinType()(*JoinType)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogCollectionRequests()([]DeviceLogCollectionResponseable)
    GetLostModeState()(*LostModeState)
    GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationStateable)
    GetManagedDeviceName()(*string)
    GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType)
    GetManagementAgent()(*ManagementAgentType)
    GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementFeatures()(*ManagedDeviceManagementFeatures)
    GetManagementState()(*ManagementState)
    GetManufacturer()(*string)
    GetMeid()(*string)
    GetModel()(*string)
    GetNotes()(*string)
    GetOdataType()(*string)
    GetOperatingSystem()(*string)
    GetOsVersion()(*string)
    GetOwnerType()(*OwnerType)
    GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState)
    GetPhoneNumber()(*string)
    GetPhysicalMemoryInBytes()(*int64)
    GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessorArchitecture()(*ManagedDeviceArchitecture)
    GetRemoteAssistanceSessionErrorDetails()(*string)
    GetRemoteAssistanceSessionUrl()(*string)
    GetRequireUserEnrollmentApproval()(*bool)
    GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetSecurityBaselineStates()([]SecurityBaselineStateable)
    GetSerialNumber()(*string)
    GetSkuFamily()(*string)
    GetSkuNumber()(*int32)
    GetSpecificationVersion()(*string)
    GetSubscriberCarrier()(*string)
    GetTotalStorageSpaceInBytes()(*int64)
    GetUdid()(*string)
    GetUserDisplayName()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetUsers()([]Userable)
    GetUsersLoggedOn()([]LoggedOnUserable)
    GetWiFiMacAddress()(*string)
    GetWindowsActiveMalwareCount()(*int32)
    GetWindowsProtectionState()(WindowsProtectionStateable)
    GetWindowsRemediatedMalwareCount()(*int32)
    SetAadRegistered(value *bool)()
    SetActivationLockBypassCode(value *string)()
    SetAndroidSecurityPatchLevel(value *string)()
    SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetailsable)()
    SetAutopilotEnrolled(value *bool)()
    SetAzureActiveDirectoryDeviceId(value *string)()
    SetAzureADDeviceId(value *string)()
    SetAzureADRegistered(value *bool)()
    SetBootstrapTokenEscrowed(value *bool)()
    SetChassisType(value *ChassisType)()
    SetChromeOSDeviceInfo(value []ChromeOSDevicePropertyable)()
    SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResultable)()
    SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetComplianceState(value *ComplianceState)()
    SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)()
    SetConfigurationManagerClientHealthState(value ConfigurationManagerClientHealthStateable)()
    SetConfigurationManagerClientInformation(value ConfigurationManagerClientInformationable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceActionResults(value []DeviceActionResultable)()
    SetDeviceCategory(value DeviceCategoryable)()
    SetDeviceCategoryDisplayName(value *string)()
    SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)()
    SetDeviceConfigurationStates(value []DeviceConfigurationStateable)()
    SetDeviceEnrollmentType(value *DeviceEnrollmentType)()
    SetDeviceFirmwareConfigurationInterfaceManaged(value *bool)()
    SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)()
    SetDeviceHealthScriptStates(value []DeviceHealthScriptPolicyStateable)()
    SetDeviceName(value *string)()
    SetDeviceRegistrationState(value *DeviceRegistrationState)()
    SetDeviceType(value *DeviceType)()
    SetEasActivated(value *bool)()
    SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasDeviceId(value *string)()
    SetEmailAddress(value *string)()
    SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEnrollmentProfileName(value *string)()
    SetEthernetMacAddress(value *string)()
    SetExchangeAccessState(value *DeviceManagementExchangeAccessState)()
    SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)()
    SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFreeStorageSpaceInBytes(value *int64)()
    SetHardwareInformation(value HardwareInformationable)()
    SetIccid(value *string)()
    SetImei(value *string)()
    SetIsEncrypted(value *bool)()
    SetIsSupervised(value *bool)()
    SetJailBroken(value *string)()
    SetJoinType(value *JoinType)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogCollectionRequests(value []DeviceLogCollectionResponseable)()
    SetLostModeState(value *LostModeState)()
    SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationStateable)()
    SetManagedDeviceName(value *string)()
    SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)()
    SetManagementAgent(value *ManagementAgentType)()
    SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementFeatures(value *ManagedDeviceManagementFeatures)()
    SetManagementState(value *ManagementState)()
    SetManufacturer(value *string)()
    SetMeid(value *string)()
    SetModel(value *string)()
    SetNotes(value *string)()
    SetOdataType(value *string)()
    SetOperatingSystem(value *string)()
    SetOsVersion(value *string)()
    SetOwnerType(value *OwnerType)()
    SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)()
    SetPhoneNumber(value *string)()
    SetPhysicalMemoryInBytes(value *int64)()
    SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessorArchitecture(value *ManagedDeviceArchitecture)()
    SetRemoteAssistanceSessionErrorDetails(value *string)()
    SetRemoteAssistanceSessionUrl(value *string)()
    SetRequireUserEnrollmentApproval(value *bool)()
    SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetSecurityBaselineStates(value []SecurityBaselineStateable)()
    SetSerialNumber(value *string)()
    SetSkuFamily(value *string)()
    SetSkuNumber(value *int32)()
    SetSpecificationVersion(value *string)()
    SetSubscriberCarrier(value *string)()
    SetTotalStorageSpaceInBytes(value *int64)()
    SetUdid(value *string)()
    SetUserDisplayName(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetUsers(value []Userable)()
    SetUsersLoggedOn(value []LoggedOnUserable)()
    SetWiFiMacAddress(value *string)()
    SetWindowsActiveMalwareCount(value *int32)()
    SetWindowsProtectionState(value WindowsProtectionStateable)()
    SetWindowsRemediatedMalwareCount(value *int32)()
}
