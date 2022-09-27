package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionDeviceRegistration represents device registration records for Bring-Your-Own-Device(BYOD) Windows devices.
type WindowsInformationProtectionDeviceRegistration struct {
    Entity
    // Device Mac address.
    deviceMacAddress *string
    // Device name.
    deviceName *string
    // Device identifier for this device registration record.
    deviceRegistrationId *string
    // Device type, for example, Windows laptop VS Windows phone.
    deviceType *string
    // Last checkin time of the device.
    lastCheckInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // UserId associated with this device registration record.
    userId *string
}
// NewWindowsInformationProtectionDeviceRegistration instantiates a new windowsInformationProtectionDeviceRegistration and sets the default values.
func NewWindowsInformationProtectionDeviceRegistration()(*WindowsInformationProtectionDeviceRegistration) {
    m := &WindowsInformationProtectionDeviceRegistration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsInformationProtectionDeviceRegistration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionDeviceRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionDeviceRegistration(), nil
}
// GetDeviceMacAddress gets the deviceMacAddress property value. Device Mac address.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceMacAddress()(*string) {
    return m.deviceMacAddress
}
// GetDeviceName gets the deviceName property value. Device name.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceRegistrationId gets the deviceRegistrationId property value. Device identifier for this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceRegistrationId()(*string) {
    return m.deviceRegistrationId
}
// GetDeviceType gets the deviceType property value. Device type, for example, Windows laptop VS Windows phone.
func (m *WindowsInformationProtectionDeviceRegistration) GetDeviceType()(*string) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionDeviceRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceMacAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceMacAddress)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["deviceRegistrationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceRegistrationId)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceType)
    res["lastCheckInDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastCheckInDateTime)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetLastCheckInDateTime gets the lastCheckInDateTime property value. Last checkin time of the device.
func (m *WindowsInformationProtectionDeviceRegistration) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCheckInDateTime
}
// GetUserId gets the userId property value. UserId associated with this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionDeviceRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceMacAddress", m.GetDeviceMacAddress())
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
    {
        err = writer.WriteStringValue("deviceRegistrationId", m.GetDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
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
    return nil
}
// SetDeviceMacAddress sets the deviceMacAddress property value. Device Mac address.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceMacAddress(value *string)() {
    m.deviceMacAddress = value
}
// SetDeviceName sets the deviceName property value. Device name.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceRegistrationId sets the deviceRegistrationId property value. Device identifier for this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceRegistrationId(value *string)() {
    m.deviceRegistrationId = value
}
// SetDeviceType sets the deviceType property value. Device type, for example, Windows laptop VS Windows phone.
func (m *WindowsInformationProtectionDeviceRegistration) SetDeviceType(value *string)() {
    m.deviceType = value
}
// SetLastCheckInDateTime sets the lastCheckInDateTime property value. Last checkin time of the device.
func (m *WindowsInformationProtectionDeviceRegistration) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckInDateTime = value
}
// SetUserId sets the userId property value. UserId associated with this device registration record.
func (m *WindowsInformationProtectionDeviceRegistration) SetUserId(value *string)() {
    m.userId = value
}
