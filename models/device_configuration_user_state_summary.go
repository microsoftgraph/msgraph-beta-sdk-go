package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DeviceConfigurationUserStateSummary struct {
    Entity
}
// NewDeviceConfigurationUserStateSummary instantiates a new DeviceConfigurationUserStateSummary and sets the default values.
func NewDeviceConfigurationUserStateSummary()(*DeviceConfigurationUserStateSummary) {
    m := &DeviceConfigurationUserStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationUserStateSummary(), nil
}
// GetCompliantUserCount gets the compliantUserCount property value. Number of compliant users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetCompliantUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliantUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConflictUserCount gets the conflictUserCount property value. Number of conflict users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetConflictUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("conflictUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorUserCount gets the errorUserCount property value. Number of error users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetErrorUserCount()(*int32) {
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
func (m *DeviceConfigurationUserStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantUserCount(val)
        }
        return nil
    }
    res["conflictUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictUserCount(val)
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
    res["nonCompliantUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantUserCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["remediatedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedUserCount(val)
        }
        return nil
    }
    res["unknownUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownUserCount(val)
        }
        return nil
    }
    return res
}
// GetNonCompliantUserCount gets the nonCompliantUserCount property value. Number of NonCompliant users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetNonCompliantUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("nonCompliantUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of not applicable users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetNotApplicableUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemediatedUserCount gets the remediatedUserCount property value. Number of remediated users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetRemediatedUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("remediatedUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of unknown users
// returns a *int32 when successful
func (m *DeviceConfigurationUserStateSummary) GetUnknownUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("unknownUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceConfigurationUserStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantUserCount", m.GetCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictUserCount", m.GetConflictUserCount())
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
        err = writer.WriteInt32Value("nonCompliantUserCount", m.GetNonCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedUserCount", m.GetRemediatedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownUserCount", m.GetUnknownUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantUserCount sets the compliantUserCount property value. Number of compliant users
func (m *DeviceConfigurationUserStateSummary) SetCompliantUserCount(value *int32)() {
    err := m.GetBackingStore().Set("compliantUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetConflictUserCount sets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) SetConflictUserCount(value *int32)() {
    err := m.GetBackingStore().Set("conflictUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorUserCount sets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) SetErrorUserCount(value *int32)() {
    err := m.GetBackingStore().Set("errorUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNonCompliantUserCount sets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) SetNonCompliantUserCount(value *int32)() {
    err := m.GetBackingStore().Set("nonCompliantUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) SetNotApplicableUserCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediatedUserCount sets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) SetRemediatedUserCount(value *int32)() {
    err := m.GetBackingStore().Set("remediatedUserCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) SetUnknownUserCount(value *int32)() {
    err := m.GetBackingStore().Set("unknownUserCount", value)
    if err != nil {
        panic(err)
    }
}
type DeviceConfigurationUserStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantUserCount()(*int32)
    GetConflictUserCount()(*int32)
    GetErrorUserCount()(*int32)
    GetNonCompliantUserCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetRemediatedUserCount()(*int32)
    GetUnknownUserCount()(*int32)
    SetCompliantUserCount(value *int32)()
    SetConflictUserCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetNonCompliantUserCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetRemediatedUserCount(value *int32)()
    SetUnknownUserCount(value *int32)()
}
