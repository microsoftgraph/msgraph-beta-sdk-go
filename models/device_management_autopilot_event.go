package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotEvent represents an Autopilot flow event.
type DeviceManagementAutopilotEvent struct {
    Entity
}
// NewDeviceManagementAutopilotEvent instantiates a new DeviceManagementAutopilotEvent and sets the default values.
func NewDeviceManagementAutopilotEvent()(*DeviceManagementAutopilotEvent) {
    m := &DeviceManagementAutopilotEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementAutopilotEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementAutopilotEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAutopilotEvent(), nil
}
// GetAccountSetupDuration gets the accountSetupDuration property value. Time spent in user ESP.
// returns a *ISODuration when successful
func (m *DeviceManagementAutopilotEvent) GetAccountSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("accountSetupDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAccountSetupStatus gets the accountSetupStatus property value. Deployment states for Autopilot devices
// returns a *WindowsAutopilotDeploymentState when successful
func (m *DeviceManagementAutopilotEvent) GetAccountSetupStatus()(*WindowsAutopilotDeploymentState) {
    val, err := m.GetBackingStore().Get("accountSetupStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotDeploymentState)
    }
    return nil
}
// GetDeploymentDuration gets the deploymentDuration property value. Autopilot deployment duration including enrollment.
// returns a *ISODuration when successful
func (m *DeviceManagementAutopilotEvent) GetDeploymentDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("deploymentDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetDeploymentEndDateTime gets the deploymentEndDateTime property value. Deployment end time.
// returns a *Time when successful
func (m *DeviceManagementAutopilotEvent) GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("deploymentEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentStartDateTime gets the deploymentStartDateTime property value. Deployment start time.
// returns a *Time when successful
func (m *DeviceManagementAutopilotEvent) GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("deploymentStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeploymentState gets the deploymentState property value. Deployment states for Autopilot devices
// returns a *WindowsAutopilotDeploymentState when successful
func (m *DeviceManagementAutopilotEvent) GetDeploymentState()(*WindowsAutopilotDeploymentState) {
    val, err := m.GetBackingStore().Get("deploymentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotDeploymentState)
    }
    return nil
}
// GetDeploymentTotalDuration gets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
// returns a *ISODuration when successful
func (m *DeviceManagementAutopilotEvent) GetDeploymentTotalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("deploymentTotalDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. Device id associated with the object
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceRegisteredDateTime gets the deviceRegisteredDateTime property value. Device registration date.
// returns a *Time when successful
func (m *DeviceManagementAutopilotEvent) GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("deviceRegisteredDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceSerialNumber gets the deviceSerialNumber property value. Device serial number.
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetDeviceSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("deviceSerialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceSetupDuration gets the deviceSetupDuration property value. Time spent in device ESP.
// returns a *ISODuration when successful
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("deviceSetupDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetDeviceSetupStatus gets the deviceSetupStatus property value. Deployment states for Autopilot devices
// returns a *WindowsAutopilotDeploymentState when successful
func (m *DeviceManagementAutopilotEvent) GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState) {
    val, err := m.GetBackingStore().Get("deviceSetupStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotDeploymentState)
    }
    return nil
}
// GetEnrollmentFailureDetails gets the enrollmentFailureDetails property value. Enrollment failure details.
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetEnrollmentFailureDetails()(*string) {
    val, err := m.GetBackingStore().Get("enrollmentFailureDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrollmentStartDateTime gets the enrollmentStartDateTime property value. Device enrollment start date.
// returns a *Time when successful
func (m *DeviceManagementAutopilotEvent) GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("enrollmentStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
// returns a *EnrollmentState when successful
func (m *DeviceManagementAutopilotEvent) GetEnrollmentState()(*EnrollmentState) {
    val, err := m.GetBackingStore().Get("enrollmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentState)
    }
    return nil
}
// GetEnrollmentType gets the enrollmentType property value. The enrollmentType property
// returns a *WindowsAutopilotEnrollmentType when successful
func (m *DeviceManagementAutopilotEvent) GetEnrollmentType()(*WindowsAutopilotEnrollmentType) {
    val, err := m.GetBackingStore().Get("enrollmentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsAutopilotEnrollmentType)
    }
    return nil
}
// GetEventDateTime gets the eventDateTime property value. Time when the event occurred .
// returns a *Time when successful
func (m *DeviceManagementAutopilotEvent) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("eventDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetManagedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. Device operating system version.
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. UserId id associated with the object
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name used to enroll the device.
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWindows10EnrollmentCompletionPageConfigurationDisplayName gets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("windows10EnrollmentCompletionPageConfigurationDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWindows10EnrollmentCompletionPageConfigurationId gets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetWindows10EnrollmentCompletionPageConfigurationId()(*string) {
    val, err := m.GetBackingStore().Get("windows10EnrollmentCompletionPageConfigurationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWindowsAutopilotDeploymentProfileDisplayName gets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
// returns a *string when successful
func (m *DeviceManagementAutopilotEvent) GetWindowsAutopilotDeploymentProfileDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("windowsAutopilotDeploymentProfileDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("accountSetupDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetAccountSetupStatus sets the accountSetupStatus property value. Deployment states for Autopilot devices
func (m *DeviceManagementAutopilotEvent) SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)() {
    err := m.GetBackingStore().Set("accountSetupStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentDuration sets the deploymentDuration property value. Autopilot deployment duration including enrollment.
func (m *DeviceManagementAutopilotEvent) SetDeploymentDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("deploymentDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentEndDateTime sets the deploymentEndDateTime property value. Deployment end time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("deploymentEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentStartDateTime sets the deploymentStartDateTime property value. Deployment start time.
func (m *DeviceManagementAutopilotEvent) SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("deploymentStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentState sets the deploymentState property value. Deployment states for Autopilot devices
func (m *DeviceManagementAutopilotEvent) SetDeploymentState(value *WindowsAutopilotDeploymentState)() {
    err := m.GetBackingStore().Set("deploymentState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentTotalDuration sets the deploymentTotalDuration property value. Total deployment duration from enrollment to Desktop screen.
func (m *DeviceManagementAutopilotEvent) SetDeploymentTotalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("deploymentTotalDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. Device id associated with the object
func (m *DeviceManagementAutopilotEvent) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceRegisteredDateTime sets the deviceRegisteredDateTime property value. Device registration date.
func (m *DeviceManagementAutopilotEvent) SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("deviceRegisteredDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceSerialNumber sets the deviceSerialNumber property value. Device serial number.
func (m *DeviceManagementAutopilotEvent) SetDeviceSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("deviceSerialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceSetupDuration sets the deviceSetupDuration property value. Time spent in device ESP.
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("deviceSetupDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceSetupStatus sets the deviceSetupStatus property value. Deployment states for Autopilot devices
func (m *DeviceManagementAutopilotEvent) SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)() {
    err := m.GetBackingStore().Set("deviceSetupStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentFailureDetails sets the enrollmentFailureDetails property value. Enrollment failure details.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentFailureDetails(value *string)() {
    err := m.GetBackingStore().Set("enrollmentFailureDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentStartDateTime sets the enrollmentStartDateTime property value. Device enrollment start date.
func (m *DeviceManagementAutopilotEvent) SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("enrollmentStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *DeviceManagementAutopilotEvent) SetEnrollmentState(value *EnrollmentState)() {
    err := m.GetBackingStore().Set("enrollmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentType sets the enrollmentType property value. The enrollmentType property
func (m *DeviceManagementAutopilotEvent) SetEnrollmentType(value *WindowsAutopilotEnrollmentType)() {
    err := m.GetBackingStore().Set("enrollmentType", value)
    if err != nil {
        panic(err)
    }
}
// SetEventDateTime sets the eventDateTime property value. Time when the event occurred .
func (m *DeviceManagementAutopilotEvent) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("eventDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceName sets the managedDeviceName property value. Managed device name.
func (m *DeviceManagementAutopilotEvent) SetManagedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. Device operating system version.
func (m *DeviceManagementAutopilotEvent) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. UserId id associated with the object
func (m *DeviceManagementAutopilotEvent) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name used to enroll the device.
func (m *DeviceManagementAutopilotEvent) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows10EnrollmentCompletionPageConfigurationDisplayName sets the windows10EnrollmentCompletionPageConfigurationDisplayName property value. Enrollment Status Page profile name
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)() {
    err := m.GetBackingStore().Set("windows10EnrollmentCompletionPageConfigurationDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows10EnrollmentCompletionPageConfigurationId sets the windows10EnrollmentCompletionPageConfigurationId property value. Enrollment Status Page profile ID
func (m *DeviceManagementAutopilotEvent) SetWindows10EnrollmentCompletionPageConfigurationId(value *string)() {
    err := m.GetBackingStore().Set("windows10EnrollmentCompletionPageConfigurationId", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsAutopilotDeploymentProfileDisplayName sets the windowsAutopilotDeploymentProfileDisplayName property value. Autopilot profile name.
func (m *DeviceManagementAutopilotEvent) SetWindowsAutopilotDeploymentProfileDisplayName(value *string)() {
    err := m.GetBackingStore().Set("windowsAutopilotDeploymentProfileDisplayName", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementAutopilotEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAccountSetupStatus()(*WindowsAutopilotDeploymentState)
    GetDeploymentDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeploymentEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentState()(*WindowsAutopilotDeploymentState)
    GetDeploymentTotalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeviceId()(*string)
    GetDeviceRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceSerialNumber()(*string)
    GetDeviceSetupDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetDeviceSetupStatus()(*WindowsAutopilotDeploymentState)
    GetEnrollmentFailureDetails()(*string)
    GetEnrollmentStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEnrollmentState()(*EnrollmentState)
    GetEnrollmentType()(*WindowsAutopilotEnrollmentType)
    GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceName()(*string)
    GetOsVersion()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetWindows10EnrollmentCompletionPageConfigurationDisplayName()(*string)
    GetWindows10EnrollmentCompletionPageConfigurationId()(*string)
    GetWindowsAutopilotDeploymentProfileDisplayName()(*string)
    SetAccountSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAccountSetupStatus(value *WindowsAutopilotDeploymentState)()
    SetDeploymentDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeploymentEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentState(value *WindowsAutopilotDeploymentState)()
    SetDeploymentTotalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeviceId(value *string)()
    SetDeviceRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceSerialNumber(value *string)()
    SetDeviceSetupDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetDeviceSetupStatus(value *WindowsAutopilotDeploymentState)()
    SetEnrollmentFailureDetails(value *string)()
    SetEnrollmentStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEnrollmentState(value *EnrollmentState)()
    SetEnrollmentType(value *WindowsAutopilotEnrollmentType)()
    SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceName(value *string)()
    SetOsVersion(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetWindows10EnrollmentCompletionPageConfigurationDisplayName(value *string)()
    SetWindows10EnrollmentCompletionPageConfigurationId(value *string)()
    SetWindowsAutopilotDeploymentProfileDisplayName(value *string)()
}
