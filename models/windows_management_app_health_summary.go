package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementAppHealthSummary contains properties for the health summary of the Windows management app.
type WindowsManagementAppHealthSummary struct {
    Entity
}
// NewWindowsManagementAppHealthSummary instantiates a new windowsManagementAppHealthSummary and sets the default values.
func NewWindowsManagementAppHealthSummary()(*WindowsManagementAppHealthSummary) {
    m := &WindowsManagementAppHealthSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsManagementAppHealthSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagementAppHealthSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagementAppHealthSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagementAppHealthSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthyDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthyDeviceCount(val)
        }
        return nil
    }
    res["unhealthyDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnhealthyDeviceCount(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetHealthyDeviceCount gets the healthyDeviceCount property value. Healthy device count.
func (m *WindowsManagementAppHealthSummary) GetHealthyDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("healthyDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnhealthyDeviceCount gets the unhealthyDeviceCount property value. Unhealthy device count.
func (m *WindowsManagementAppHealthSummary) GetUnhealthyDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("unhealthyDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Unknown device count.
func (m *WindowsManagementAppHealthSummary) GetUnknownDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("unknownDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsManagementAppHealthSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("healthyDeviceCount", m.GetHealthyDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unhealthyDeviceCount", m.GetUnhealthyDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthyDeviceCount sets the healthyDeviceCount property value. Healthy device count.
func (m *WindowsManagementAppHealthSummary) SetHealthyDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("healthyDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnhealthyDeviceCount sets the unhealthyDeviceCount property value. Unhealthy device count.
func (m *WindowsManagementAppHealthSummary) SetUnhealthyDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("unhealthyDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Unknown device count.
func (m *WindowsManagementAppHealthSummary) SetUnknownDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("unknownDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// WindowsManagementAppHealthSummaryable 
type WindowsManagementAppHealthSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHealthyDeviceCount()(*int32)
    GetUnhealthyDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetHealthyDeviceCount(value *int32)()
    SetUnhealthyDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}
