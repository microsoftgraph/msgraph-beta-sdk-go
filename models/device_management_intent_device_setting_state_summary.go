package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceSettingStateSummary entity that represents device setting state summary for an intent
type DeviceManagementIntentDeviceSettingStateSummary struct {
    Entity
    // Number of compliant devices
    compliantCount *int32
    // Number of devices in conflict
    conflictCount *int32
    // Number of error devices
    errorCount *int32
    // Number of non compliant devices
    nonCompliantCount *int32
    // Number of not applicable devices
    notApplicableCount *int32
    // Number of remediated devices
    remediatedCount *int32
    // Name of a setting
    settingName *string
}
// NewDeviceManagementIntentDeviceSettingStateSummary instantiates a new deviceManagementIntentDeviceSettingStateSummary and sets the default values.
func NewDeviceManagementIntentDeviceSettingStateSummary()(*DeviceManagementIntentDeviceSettingStateSummary) {
    m := &DeviceManagementIntentDeviceSettingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentDeviceSettingStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentDeviceSettingStateSummary(), nil
}
// GetCompliantCount gets the compliantCount property value. Number of compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetCompliantCount()(*int32) {
    return m.compliantCount
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetConflictCount()(*int32) {
    return m.conflictCount
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetErrorCount()(*int32) {
    return m.errorCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompliantCount)
    res["conflictCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConflictCount)
    res["errorCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorCount)
    res["nonCompliantCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNonCompliantCount)
    res["notApplicableCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNotApplicableCount)
    res["remediatedCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemediatedCount)
    res["settingName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingName)
    return res
}
// GetNonCompliantCount gets the nonCompliantCount property value. Number of non compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNonCompliantCount()(*int32) {
    return m.nonCompliantCount
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNotApplicableCount()(*int32) {
    return m.notApplicableCount
}
// GetRemediatedCount gets the remediatedCount property value. Number of remediated devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetRemediatedCount()(*int32) {
    return m.remediatedCount
}
// GetSettingName gets the settingName property value. Name of a setting
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetSettingName()(*string) {
    return m.settingName
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentDeviceSettingStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantCount", m.GetCompliantCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictCount", m.GetConflictCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantCount", m.GetNonCompliantCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedCount", m.GetRemediatedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantCount sets the compliantCount property value. Number of compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetCompliantCount(value *int32)() {
    m.compliantCount = value
}
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// SetNonCompliantCount sets the nonCompliantCount property value. Number of non compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNonCompliantCount(value *int32)() {
    m.nonCompliantCount = value
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// SetRemediatedCount sets the remediatedCount property value. Number of remediated devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetRemediatedCount(value *int32)() {
    m.remediatedCount = value
}
// SetSettingName sets the settingName property value. Name of a setting
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetSettingName(value *string)() {
    m.settingName = value
}
