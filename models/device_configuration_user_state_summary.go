package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationUserStateSummary 
type DeviceConfigurationUserStateSummary struct {
    Entity
    // Number of compliant users
    compliantUserCount *int32
    // Number of conflict users
    conflictUserCount *int32
    // Number of error users
    errorUserCount *int32
    // Number of NonCompliant users
    nonCompliantUserCount *int32
    // Number of not applicable users
    notApplicableUserCount *int32
    // Number of remediated users
    remediatedUserCount *int32
    // Number of unknown users
    unknownUserCount *int32
}
// NewDeviceConfigurationUserStateSummary instantiates a new deviceConfigurationUserStateSummary and sets the default values.
func NewDeviceConfigurationUserStateSummary()(*DeviceConfigurationUserStateSummary) {
    m := &DeviceConfigurationUserStateSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceConfigurationUserStateSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationUserStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationUserStateSummary(), nil
}
// GetCompliantUserCount gets the compliantUserCount property value. Number of compliant users
func (m *DeviceConfigurationUserStateSummary) GetCompliantUserCount()(*int32) {
    return m.compliantUserCount
}
// GetConflictUserCount gets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) GetConflictUserCount()(*int32) {
    return m.conflictUserCount
}
// GetErrorUserCount gets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) GetErrorUserCount()(*int32) {
    return m.errorUserCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationUserStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompliantUserCount)
    res["conflictUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConflictUserCount)
    res["errorUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorUserCount)
    res["nonCompliantUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNonCompliantUserCount)
    res["notApplicableUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNotApplicableUserCount)
    res["remediatedUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemediatedUserCount)
    res["unknownUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnknownUserCount)
    return res
}
// GetNonCompliantUserCount gets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) GetNonCompliantUserCount()(*int32) {
    return m.nonCompliantUserCount
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) GetNotApplicableUserCount()(*int32) {
    return m.notApplicableUserCount
}
// GetRemediatedUserCount gets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) GetRemediatedUserCount()(*int32) {
    return m.remediatedUserCount
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) GetUnknownUserCount()(*int32) {
    return m.unknownUserCount
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
    m.compliantUserCount = value
}
// SetConflictUserCount sets the conflictUserCount property value. Number of conflict users
func (m *DeviceConfigurationUserStateSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
// SetErrorUserCount sets the errorUserCount property value. Number of error users
func (m *DeviceConfigurationUserStateSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
// SetNonCompliantUserCount sets the nonCompliantUserCount property value. Number of NonCompliant users
func (m *DeviceConfigurationUserStateSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of not applicable users
func (m *DeviceConfigurationUserStateSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// SetRemediatedUserCount sets the remediatedUserCount property value. Number of remediated users
func (m *DeviceConfigurationUserStateSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of unknown users
func (m *DeviceConfigurationUserStateSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}
