package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentUserStateSummary 
type DeviceManagementIntentUserStateSummary struct {
    Entity
    // Number of users in conflict
    conflictCount *int32
    // Number of error users
    errorCount *int32
    // Number of failed users
    failedCount *int32
    // Number of not applicable users
    notApplicableCount *int32
    // Number of succeeded users
    successCount *int32
}
// NewDeviceManagementIntentUserStateSummary instantiates a new deviceManagementIntentUserStateSummary and sets the default values.
func NewDeviceManagementIntentUserStateSummary()(*DeviceManagementIntentUserStateSummary) {
    m := &DeviceManagementIntentUserStateSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementIntentUserStateSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementIntentUserStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentUserStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentUserStateSummary(), nil
}
// GetConflictCount gets the conflictCount property value. Number of users in conflict
func (m *DeviceManagementIntentUserStateSummary) GetConflictCount()(*int32) {
    return m.conflictCount
}
// GetErrorCount gets the errorCount property value. Number of error users
func (m *DeviceManagementIntentUserStateSummary) GetErrorCount()(*int32) {
    return m.errorCount
}
// GetFailedCount gets the failedCount property value. Number of failed users
func (m *DeviceManagementIntentUserStateSummary) GetFailedCount()(*int32) {
    return m.failedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentUserStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConflictCount)
    res["errorCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorCount)
    res["failedCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetFailedCount)
    res["notApplicableCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNotApplicableCount)
    res["successCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSuccessCount)
    return res
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable users
func (m *DeviceManagementIntentUserStateSummary) GetNotApplicableCount()(*int32) {
    return m.notApplicableCount
}
// GetSuccessCount gets the successCount property value. Number of succeeded users
func (m *DeviceManagementIntentUserStateSummary) GetSuccessCount()(*int32) {
    return m.successCount
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentUserStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
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
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConflictCount sets the conflictCount property value. Number of users in conflict
func (m *DeviceManagementIntentUserStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
// SetErrorCount sets the errorCount property value. Number of error users
func (m *DeviceManagementIntentUserStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// SetFailedCount sets the failedCount property value. Number of failed users
func (m *DeviceManagementIntentUserStateSummary) SetFailedCount(value *int32)() {
    m.failedCount = value
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable users
func (m *DeviceManagementIntentUserStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// SetSuccessCount sets the successCount property value. Number of succeeded users
func (m *DeviceManagementIntentUserStateSummary) SetSuccessCount(value *int32)() {
    m.successCount = value
}
