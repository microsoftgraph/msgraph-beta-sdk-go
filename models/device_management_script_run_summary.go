package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptRunSummary contains properties for the run summary of a device management script.
type DeviceManagementScriptRunSummary struct {
    Entity
}
// NewDeviceManagementScriptRunSummary instantiates a new DeviceManagementScriptRunSummary and sets the default values.
func NewDeviceManagementScriptRunSummary()(*DeviceManagementScriptRunSummary) {
    m := &DeviceManagementScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementScriptRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptRunSummary(), nil
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count.
// returns a *int32 when successful
func (m *DeviceManagementScriptRunSummary) GetErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorUserCount gets the errorUserCount property value. Error user count.
// returns a *int32 when successful
func (m *DeviceManagementScriptRunSummary) GetErrorUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementScriptRunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["errorUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
        }
        return nil
    }
    res["successDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessDeviceCount(val)
        }
        return nil
    }
    res["successUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessUserCount(val)
        }
        return nil
    }
    return res
}
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count.
// returns a *int32 when successful
func (m *DeviceManagementScriptRunSummary) GetSuccessDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("successDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSuccessUserCount gets the successUserCount property value. Success user count.
// returns a *int32 when successful
func (m *DeviceManagementScriptRunSummary) GetSuccessUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("successUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptRunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successUserCount", m.GetSuccessUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) SetErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("errorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorUserCount sets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) SetErrorUserCount(value *int32)() {
    err := m.GetBackingStore().Set("errorUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) SetSuccessDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("successDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessUserCount sets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) SetSuccessUserCount(value *int32)() {
    err := m.GetBackingStore().Set("successUserCount", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementScriptRunSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorDeviceCount()(*int32)
    GetErrorUserCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetSuccessUserCount()(*int32)
    SetErrorDeviceCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetSuccessUserCount(value *int32)()
}
