package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotEvent represents an Autopilot flow event.
type DeviceManagementAutopilotEvent struct {
    Entity
    // Time spent in user ESP.
    accountSetupDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The accountSetupStatus property
    accountSetupStatus *WindowsAutopilotDeploymentState
    // Autopilot deployment duration including enrollment.
    deploymentDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Deployment end time.
    deploymentEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Deployment start time.
    deploymentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deploymentState property
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
    // The deviceSetupStatus property
    deviceSetupStatus *WindowsAutopilotDeploymentState
    // Enrollment failure details.
    enrollmentFailureDetails *string
    // Device enrollment start date.
    enrollmentStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The enrollmentState property
    enrollmentState *EnrollmentState
    // The enrollmentType property
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
    return m.accountSetupDuration
}
// GetAccountSetupStatus gets the accountSetupStatus property value. The accountSetupStatus property
func (m *DeviceManagementAutopilotEvent) GetAccountSetupStatus()(*WindowsAutopilotDeploymentState) {
    return m.accountSetupStatus
}
// GetDeploymentDuration gets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) GetDeploymentDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.deploymentDuration
}
// GetDeploymentEndDateTime gets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deploymentEndDateTime
}
// GetDeploymentStartDateTime gets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deploymentStartDateTime
}
// GetDeploymentState gets the deploymentState property value. The deploymentState property
func (m *DeviceManagementAutopilotEvent) GetDeploymentState()(*WindowsAutopilotDeploymentState) {
    return m.deploymentState
}
// GetDeploymentTotalDuration gets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) GetDeploymentTotalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.deploymentTotalDuration
}
// GetDeviceId gets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDevicePreparationDuration gets the devicePreparationDuration property value. Time spent in device enrollment.
func (m *DeviceManagementAutopilotEvent) GetDevicePreparationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.devicePreparationDuration
}
// GetDeviceRegisteredDateTime gets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deviceRegisteredDateTime
}
// GetDeviceSerialNumber gets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) GetDeviceSerialNumber()(*string) {
    return m.deviceSerialNumber
}
// GetDeviceSetupDuration gets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.deviceSetupDuration
}
// GetDeviceSetupStatus gets the deviceSetupStatus property value. The deviceSetupStatus property
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState) {
    return m.deviceSetupStatus
}
// GetEnrollmentFailureDetails gets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentFailureDetails()(*string) {
    return m.enrollmentFailureDetails
}
// GetEnrollmentStartDateTime gets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrollmentStartDateTime
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
func (m *DeviceManagementAutopilotEvent) GetEnrollmentState()(*EnrollmentState) {
    return m.enrollmentState
}
// GetEnrollmentType gets the enrollmentType property value. The enrollmentType property
func (m *DeviceManagementAutopilotEvent) GetEnrollmentType()(*WindowsAutopilotEnrollmentType) {
    return m.enrollmentType
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAutopilotEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountSetupDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetAccountSetupDuration)
    res["accountSetupStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsAutopilotDeploymentState , m.SetAccountSetupStatus)
    res["deploymentDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDeploymentDuration)
    res["deploymentEndDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDeploymentEndDateTime)
    res["deploymentStartDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDeploymentStartDateTime)
    res["deploymentState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsAutopilotDeploymentState , m.SetDeploymentState)
    res["deploymentTotalDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDeploymentTotalDuration)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["devicePreparationDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDevicePreparationDuration)
    res["deviceRegisteredDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDeviceRegisteredDateTime)
    res["deviceSerialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceSerialNumber)
    res["deviceSetupDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetDeviceSetupDuration)
    res["deviceSetupStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsAutopilotDeploymentState , m.SetDeviceSetupStatus)
    res["enrollmentFailureDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEnrollmentFailureDetails)
    res["enrollmentStartDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEnrollmentStartDateTime)
    res["enrollmentState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentState , m.SetEnrollmentState)
    res["enrollmentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsAutopilotEnrollmentType , m.SetEnrollmentType)
    res["eventDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEventDateTime)
    res["managedDeviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceName)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["policyStatusDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue , m.SetPolicyStatusDetails)
    res["targetedAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTargetedAppCount)
    res["targetedPolicyCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTargetedPolicyCount)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["windows10EnrollmentCompletionPageConfigurationDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWindows10EnrollmentCompletionPageConfigurationDisplayName)
    res["windows10EnrollmentCompletionPageConfigurationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWindows10EnrollmentCompletionPageConfigurationId)
    res["windowsAutopilotDeploymentProfileDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWindowsAutopilotDeploymentProfileDisplayName)
    return res
}
// GetManagedDeviceName gets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetOsVersion gets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPolicyStatusDetails gets the policyStatusDetails property value. Policy and application status details for this device.
func (m *DeviceManagementAutopilotEvent) GetPolicyStatusDetails()([]DeviceManagementAutopilotPolicyStatusDetailable) {
    return m.policyStatusDetails
}
// GetTargetedAppCount gets the targetedAppCount property value. Count of applications targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedAppCount()(*int32) {
    return m.targetedAppCount
}
// GetTargetedPolicyCount gets the targetedPolicyCount property value. Count of policies targeted.
func (m *DeviceManagementAutopilotEvent) GetTargetedPolicyCount()(*int32) {
    return m.targetedPolicyCount
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetWindows10EnrollmentCompletionPageConfigurationDisplayName gets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string) {
    return m.windows10EnrollmentCompletionPageConfigurationDisplayName
}
// GetWindows10EnrollmentCompletionPageConfigurationId gets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationId()(*string) {
    return m.windows10EnrollmentCompletionPageConfigurationId
}
// GetWindowsAutopilotDeploymentProfileDisplayName gets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) GetWindowsAutopilotDeploymentProfileDisplayName()(*string) {
    return m.windowsAutopilotDeploymentProfileDisplayName
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPolicyStatusDetails())
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
    m.accountSetupDuration = value
}
// SetAccountSetupStatus sets the accountSetupStatus property value. The accountSetupStatus property
func (m *DeviceManagementAutopilotEvent) SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)() {
    m.accountSetupStatus = value
}
// SetDeploymentDuration sets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) SetDeploymentDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.deploymentDuration = value
}
// SetDeploymentEndDateTime sets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentEndDateTime = value
}
// SetDeploymentStartDateTime sets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentStartDateTime = value
}
// SetDeploymentState sets the deploymentState property value. The deploymentState property
func (m *DeviceManagementAutopilotEvent) SetDeploymentState(value *WindowsAutopilotDeploymentState)() {
    m.deploymentState = value
}
// SetDeploymentTotalDuration sets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) SetDeploymentTotalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.deploymentTotalDuration = value
}
// SetDeviceId sets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDevicePreparationDuration sets the devicePreparationDuration property value. Time spent in device enrollment.
func (m *DeviceManagementAutopilotEvent) SetDevicePreparationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.devicePreparationDuration = value
}
// SetDeviceRegisteredDateTime sets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deviceRegisteredDateTime = value
}
// SetDeviceSerialNumber sets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) SetDeviceSerialNumber(value *string)() {
    m.deviceSerialNumber = value
}
// SetDeviceSetupDuration sets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.deviceSetupDuration = value
}
// SetDeviceSetupStatus sets the deviceSetupStatus property value. The deviceSetupStatus property
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)() {
    m.deviceSetupStatus = value
}
// SetEnrollmentFailureDetails sets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentFailureDetails(value *string)() {
    m.enrollmentFailureDetails = value
}
// SetEnrollmentStartDateTime sets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrollmentStartDateTime = value
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *DeviceManagementAutopilotEvent) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// SetEnrollmentType sets the enrollmentType property value. The enrollmentType property
func (m *DeviceManagementAutopilotEvent) SetEnrollmentType(value *WindowsAutopilotEnrollmentType)() {
    m.enrollmentType = value
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetOsVersion sets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPolicyStatusDetails sets the policyStatusDetails property value. Policy and application status details for this device.
func (m *DeviceManagementAutopilotEvent) SetPolicyStatusDetails(value []DeviceManagementAutopilotPolicyStatusDetailable)() {
    m.policyStatusDetails = value
}
// SetTargetedAppCount sets the targetedAppCount property value. Count of applications targeted.
func (m *DeviceManagementAutopilotEvent) SetTargetedAppCount(value *int32)() {
    m.targetedAppCount = value
}
// SetTargetedPolicyCount sets the targetedPolicyCount property value. Count of policies targeted.
func (m *DeviceManagementAutopilotEvent) SetTargetedPolicyCount(value *int32)() {
    m.targetedPolicyCount = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetWindows10EnrollmentCompletionPageConfigurationDisplayName sets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)() {
    m.windows10EnrollmentCompletionPageConfigurationDisplayName = value
}
// SetWindows10EnrollmentCompletionPageConfigurationId sets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationId(value *string)() {
    m.windows10EnrollmentCompletionPageConfigurationId = value
}
// SetWindowsAutopilotDeploymentProfileDisplayName sets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) SetWindowsAutopilotDeploymentProfileDisplayName(value *string)() {
    m.windowsAutopilotDeploymentProfileDisplayName = value
}
