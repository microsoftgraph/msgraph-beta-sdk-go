package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionWipeAction represents wipe requests issued by tenant admin for Bring-Your-Own-Device(BYOD) Windows devices.
type WindowsInformationProtectionWipeAction struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewWindowsInformationProtectionWipeAction instantiates a new windowsInformationProtectionWipeAction and sets the default values.
func NewWindowsInformationProtectionWipeAction()(*WindowsInformationProtectionWipeAction) {
    m := &WindowsInformationProtectionWipeAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionWipeActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionWipeAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionWipeAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastCheckInDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckInDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ActionState))
        }
        return nil
    }
    res["targetedDeviceMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceMacAddress(val)
        }
        return nil
    }
    res["targetedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceName(val)
        }
        return nil
    }
    res["targetedDeviceRegistrationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedDeviceRegistrationId(val)
        }
        return nil
    }
    res["targetedUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedUserId(val)
        }
        return nil
    }
    return res
}
// GetLastCheckInDateTime gets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckInDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *WindowsInformationProtectionWipeAction) GetStatus()(*ActionState) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ActionState)
    }
    return nil
}
// GetTargetedDeviceMacAddress gets the targetedDeviceMacAddress property value. Targeted device Mac address.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceMacAddress()(*string) {
    val, err := m.GetBackingStore().Get("targetedDeviceMacAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetedDeviceName gets the targetedDeviceName property value. Targeted device name.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("targetedDeviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetedDeviceRegistrationId gets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedDeviceRegistrationId()(*string) {
    val, err := m.GetBackingStore().Get("targetedDeviceRegistrationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetedUserId gets the targetedUserId property value. The UserId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) GetTargetedUserId()(*string) {
    val, err := m.GetBackingStore().Get("targetedUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionWipeAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastCheckInDateTime", m.GetLastCheckInDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceMacAddress", m.GetTargetedDeviceMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceName", m.GetTargetedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedDeviceRegistrationId", m.GetTargetedDeviceRegistrationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetedUserId", m.GetTargetedUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastCheckInDateTime sets the lastCheckInDateTime property value. Last checkin time of the device that was targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckInDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *WindowsInformationProtectionWipeAction) SetStatus(value *ActionState)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedDeviceMacAddress sets the targetedDeviceMacAddress property value. Targeted device Mac address.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceMacAddress(value *string)() {
    err := m.GetBackingStore().Set("targetedDeviceMacAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedDeviceName sets the targetedDeviceName property value. Targeted device name.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceName(value *string)() {
    err := m.GetBackingStore().Set("targetedDeviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedDeviceRegistrationId sets the targetedDeviceRegistrationId property value. The DeviceRegistrationId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetTargetedDeviceRegistrationId(value *string)() {
    err := m.GetBackingStore().Set("targetedDeviceRegistrationId", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedUserId sets the targetedUserId property value. The UserId being targeted by this wipe action.
func (m *WindowsInformationProtectionWipeAction) SetTargetedUserId(value *string)() {
    err := m.GetBackingStore().Set("targetedUserId", value)
    if err != nil {
        panic(err)
    }
}
// WindowsInformationProtectionWipeActionable 
type WindowsInformationProtectionWipeActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*ActionState)
    GetTargetedDeviceMacAddress()(*string)
    GetTargetedDeviceName()(*string)
    GetTargetedDeviceRegistrationId()(*string)
    GetTargetedUserId()(*string)
    SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *ActionState)()
    SetTargetedDeviceMacAddress(value *string)()
    SetTargetedDeviceName(value *string)()
    SetTargetedDeviceRegistrationId(value *string)()
    SetTargetedUserId(value *string)()
}
