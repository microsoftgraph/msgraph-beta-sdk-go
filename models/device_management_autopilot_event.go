package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotEvent 
type DeviceManagementAutopilotEvent struct {
    Entity
    // Time spent in user ESP.
    accountSetupDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    accountSetupStatus *WindowsAutopilotDeploymentState
    // Autopilot deployment duration including enrollment.
    deploymentDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Deployment end time.
    deploymentEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Deployment start time.
    deploymentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    deploymentState *WindowsAutopilotDeploymentState
    // Total deployment duration from enrollment to Desktop screen.
    deploymentTotalDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Device id associated with the object
    deviceId *string
    // Time spent in device enrollment.
    devicePreparationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Device registration date.
    deviceRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device serial number.
    deviceSerialNumber *string
    // Time spent in device ESP.
    deviceSetupDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
    deviceSetupStatus *WindowsAutopilotDeploymentState
    // Enrollment failure details.
    enrollmentFailureDetails *string
    // Device enrollment start date.
    enrollmentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
    enrollmentState *EnrollmentState
    // Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
    enrollmentType *WindowsAutopilotEnrollmentType
    // Time when the event occurred .
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Managed device name.
    managedDeviceName *string
    // Device operating system version.
    osVersion *string
    // Policy and application status details for this device.
    policyStatusDetails []DeviceManagementAutopilotPolicyStatusDetailable
    // Count of applications targeted.
    targetedAppCount *int32
    // Count of policies targeted.
    targetedPolicyCount *int32
    // User principal name used to enroll the device.
    userPrincipalName *string
    // Enrollment Status Page profile name
    windows10EnrollmentCompletionPageConfigurationDisplayName *string
    // Enrollment Status Page profile ID
    windows10EnrollmentCompletionPageConfigurationId *string
    // Autopilot profile name.
    windowsAutopilotDeploymentProfileDisplayName *string
}
// NewDeviceManagementAutopilotEvent instantiates a new deviceManagementAutopilotEvent and sets the default values.
func NewDeviceManagementAutopilotEvent()(*DeviceManagementAutopilotEvent) {
    m := &DeviceManagementAutopilotEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementAutopilotEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAutopilotEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAutopilotEvent(), nil
}
// GetAccountSetupDuration gets the accountSetupDuration property value. Time spent in user ESP.
func (m *DeviceManagementAutopilotEvent) GetAccountSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.accountSetupDuration
    }
}
// GetAccountSetupStatus gets the accountSetupStatus property value. Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetAccountSetupStatus()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.accountSetupStatus
    }
}
// GetDeploymentDuration gets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) GetDeploymentDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.deploymentDuration
    }
}
// GetDeploymentEndDateTime gets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deploymentEndDateTime
    }
}
// GetDeploymentStartDateTime gets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deploymentStartDateTime
    }
}
// GetDeploymentState gets the deploymentState property value. Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetDeploymentState()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.deploymentState
    }
}
// GetDeploymentTotalDuration gets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) GetDeploymentTotalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.deploymentTotalDuration
    }
}
// GetDeviceId gets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDevicePreparationDuration gets the devicePreparationDuration property value. Time spent in device enrollment.
func (m *DeviceManagementAutopilotEvent) GetDevicePreparationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.devicePreparationDuration
    }
}
// GetDeviceRegisteredDateTime gets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deviceRegisteredDateTime
    }
}
// GetDeviceSerialNumber gets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) GetDeviceSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceSerialNumber
    }
}
// GetDeviceSetupDuration gets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.deviceSetupDuration
    }
}
// GetDeviceSetupStatus gets the deviceSetupStatus property value. Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.deviceSetupStatus
    }
}
// GetEnrollmentFailureDetails gets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentFailureDetails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentFailureDetails
    }
}
// GetEnrollmentStartDateTime gets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentStartDateTime
    }
}
// GetEnrollmentState gets the enrollmentState property value. Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// GetEnrollmentType gets the enrollmentType property value. Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentType()(*WindowsAutopilotEnrollmentType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAutopilotEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountSetupDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountSetupDuration(val)
        }
        return nil
    }
    res["accountSetupStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountSetupStatus(val.(*WindowsAutopilotDeploymentState))
        }
        return nil
    }
    res["deploymentDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentDuration(val)
        }
        return nil
    }
    res["deploymentEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentEndDateTime(val)
        }
        return nil
    }
    res["deploymentStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStartDateTime(val)
        }
        return nil
    }
    res["deploymentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentState(val.(*WindowsAutopilotDeploymentState))
        }
        return nil
    }
    res["deploymentTotalDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentTotalDuration(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["devicePreparationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePreparationDuration(val)
        }
        return nil
    }
    res["deviceRegisteredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegisteredDateTime(val)
        }
        return nil
    }
    res["deviceSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceSerialNumber(val)
        }
        return nil
    }
    res["deviceSetupDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceSetupDuration(val)
        }
        return nil
    }
    res["deviceSetupStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeploymentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceSetupStatus(val.(*WindowsAutopilotDeploymentState))
        }
        return nil
    }
    res["enrollmentFailureDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentFailureDetails(val)
        }
        return nil
    }
    res["enrollmentStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentStartDateTime(val)
        }
        return nil
    }
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*WindowsAutopilotEnrollmentType))
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
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
    res["policyStatusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementAutopilotPolicyStatusDetailable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementAutopilotPolicyStatusDetailable)
            }
            m.SetPolicyStatusDetails(res)
        }
        return nil
    }
    res["targetedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedAppCount(val)
        }
        return nil
    }
    res["targetedPolicyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedPolicyCount(val)
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
    res["windows10EnrollmentCompletionPageConfigurationDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10EnrollmentCompletionPageConfigurationDisplayName(val)
        }
        return nil
    }
    res["windows10EnrollmentCompletionPageConfigurationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10EnrollmentCompletionPageConfigurationId(val)
        }
        return nil
    }
    res["windowsAutopilotDeploymentProfileDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsAutopilotDeploymentProfileDisplayName(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceName gets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) GetManagedDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceName
    }
}
// GetOsVersion gets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetPolicyStatusDetails gets the policyStatusDetails property value. Policy and application status details for this device.
func (m *DeviceManagementAutopilotEvent) GetPolicyStatusDetails()([]DeviceManagementAutopilotPolicyStatusDetailable) {
    if m == nil {
        return nil
    } else {
        return m.policyStatusDetails
    }
}
// GetTargetedAppCount gets the targetedAppCount property value. Count of applications targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedAppCount
    }
}
// GetTargetedPolicyCount gets the targetedPolicyCount property value. Count of policies targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedPolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.targetedPolicyCount
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetWindows10EnrollmentCompletionPageConfigurationDisplayName gets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windows10EnrollmentCompletionPageConfigurationDisplayName
    }
}
// GetWindows10EnrollmentCompletionPageConfigurationId gets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windows10EnrollmentCompletionPageConfigurationId
    }
}
// GetWindowsAutopilotDeploymentProfileDisplayName gets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) GetWindowsAutopilotDeploymentProfileDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeploymentProfileDisplayName
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementAutopilotEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("accountSetupDuration", m.GetAccountSetupDuration())
        if err != nil {
            return err
        }
    }
    if m.GetAccountSetupStatus() != nil {
        cast := (*m.GetAccountSetupStatus()).String()
        err = writer.WriteStringValue("accountSetupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("deploymentDuration", m.GetDeploymentDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentEndDateTime", m.GetDeploymentEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentStartDateTime", m.GetDeploymentStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentState() != nil {
        cast := (*m.GetDeploymentState()).String()
        err = writer.WriteStringValue("deploymentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("deploymentTotalDuration", m.GetDeploymentTotalDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("devicePreparationDuration", m.GetDevicePreparationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deviceRegisteredDateTime", m.GetDeviceRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceSerialNumber", m.GetDeviceSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("deviceSetupDuration", m.GetDeviceSetupDuration())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSetupStatus() != nil {
        cast := (*m.GetDeviceSetupStatus()).String()
        err = writer.WriteStringValue("deviceSetupStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentFailureDetails", m.GetEnrollmentFailureDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrollmentStartDateTime", m.GetEnrollmentStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err = writer.WriteStringValue("enrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
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
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyStatusDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicyStatusDetails()))
        for i, v := range m.GetPolicyStatusDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("policyStatusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("targetedAppCount", m.GetTargetedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("targetedPolicyCount", m.GetTargetedPolicyCount())
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
    {
        err = writer.WriteStringValue("windows10EnrollmentCompletionPageConfigurationDisplayName", m.GetWindows10EnrollmentCompletionPageConfigurationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windows10EnrollmentCompletionPageConfigurationId", m.GetWindows10EnrollmentCompletionPageConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("windowsAutopilotDeploymentProfileDisplayName", m.GetWindowsAutopilotDeploymentProfileDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountSetupDuration sets the accountSetupDuration property value. Time spent in user ESP.
func (m *DeviceManagementAutopilotEvent) SetAccountSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.accountSetupDuration = value
    }
}
// SetAccountSetupStatus sets the accountSetupStatus property value. Deployment status for the enrollment status page account setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)() {
    if m != nil {
        m.accountSetupStatus = value
    }
}
// SetDeploymentDuration sets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) SetDeploymentDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.deploymentDuration = value
    }
}
// SetDeploymentEndDateTime sets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deploymentEndDateTime = value
    }
}
// SetDeploymentStartDateTime sets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deploymentStartDateTime = value
    }
}
// SetDeploymentState sets the deploymentState property value. Deployment state like Success, Failure, InProgress, SuccessWithTimeout. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) SetDeploymentState(value *WindowsAutopilotDeploymentState)() {
    if m != nil {
        m.deploymentState = value
    }
}
// SetDeploymentTotalDuration sets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) SetDeploymentTotalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.deploymentTotalDuration = value
    }
}
// SetDeviceId sets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDevicePreparationDuration sets the devicePreparationDuration property value. Time spent in device enrollment.
func (m *DeviceManagementAutopilotEvent) SetDevicePreparationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.devicePreparationDuration = value
    }
}
// SetDeviceRegisteredDateTime sets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deviceRegisteredDateTime = value
    }
}
// SetDeviceSerialNumber sets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) SetDeviceSerialNumber(value *string)() {
    if m != nil {
        m.deviceSerialNumber = value
    }
}
// SetDeviceSetupDuration sets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.deviceSetupDuration = value
    }
}
// SetDeviceSetupStatus sets the deviceSetupStatus property value. Deployment status for the enrollment status page device setup phase. Possible values are: unknown, success, inProgress, failure, successWithTimeout, notAttempted, disabled.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)() {
    if m != nil {
        m.deviceSetupStatus = value
    }
}
// SetEnrollmentFailureDetails sets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentFailureDetails(value *string)() {
    if m != nil {
        m.enrollmentFailureDetails = value
    }
}
// SetEnrollmentStartDateTime sets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.enrollmentStartDateTime = value
    }
}
// SetEnrollmentState sets the enrollmentState property value. Enrollment state like Enrolled, Failed. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentState(value *EnrollmentState)() {
    if m != nil {
        m.enrollmentState = value
    }
}
// SetEnrollmentType sets the enrollmentType property value. Enrollment type. Possible values are: unknown, azureADJoinedWithAutopilotProfile, offlineDomainJoined, azureADJoinedUsingDeviceAuthWithAutopilotProfile, azureADJoinedUsingDeviceAuthWithoutAutopilotProfile, azureADJoinedWithOfflineAutopilotProfile, azureADJoinedWithWhiteGlove, offlineDomainJoinedWithWhiteGlove, offlineDomainJoinedWithOfflineAutopilotProfile.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentType(value *WindowsAutopilotEnrollmentType)() {
    if m != nil {
        m.enrollmentType = value
    }
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventDateTime = value
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) SetManagedDeviceName(value *string)() {
    if m != nil {
        m.managedDeviceName = value
    }
}
// SetOsVersion sets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetPolicyStatusDetails sets the policyStatusDetails property value. Policy and application status details for this device.
func (m *DeviceManagementAutopilotEvent) SetPolicyStatusDetails(value []DeviceManagementAutopilotPolicyStatusDetailable)() {
    if m != nil {
        m.policyStatusDetails = value
    }
}
// SetTargetedAppCount sets the targetedAppCount property value. Count of applications targeted.
func (m *DeviceManagementAutopilotEvent) SetTargetedAppCount(value *int32)() {
    if m != nil {
        m.targetedAppCount = value
    }
}
// SetTargetedPolicyCount sets the targetedPolicyCount property value. Count of policies targeted.
func (m *DeviceManagementAutopilotEvent) SetTargetedPolicyCount(value *int32)() {
    if m != nil {
        m.targetedPolicyCount = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetWindows10EnrollmentCompletionPageConfigurationDisplayName sets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)() {
    if m != nil {
        m.windows10EnrollmentCompletionPageConfigurationDisplayName = value
    }
}
// SetWindows10EnrollmentCompletionPageConfigurationId sets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationId(value *string)() {
    if m != nil {
        m.windows10EnrollmentCompletionPageConfigurationId = value
    }
}
// SetWindowsAutopilotDeploymentProfileDisplayName sets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) SetWindowsAutopilotDeploymentProfileDisplayName(value *string)() {
    if m != nil {
        m.windowsAutopilotDeploymentProfileDisplayName = value
    }
}
