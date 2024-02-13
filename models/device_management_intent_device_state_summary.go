package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceStateSummary entity that represents device state summary for an intent
type DeviceManagementIntentDeviceStateSummary struct {
    Entity
}
// NewDeviceManagementIntentDeviceStateSummary instantiates a new DeviceManagementIntentDeviceStateSummary and sets the default values.
func NewDeviceManagementIntentDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    m := &DeviceManagementIntentDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentDeviceStateSummary(), nil
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetConflictCount()(*int32) {
    val, err := m.GetBackingStore().Get("conflictCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorCount gets the errorCount property value. Number of error devices
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetErrorCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFailedCount gets the failedCount property value. Number of failed devices
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetFailedCount()(*int32) {
    val, err := m.GetBackingStore().Get("failedCount")
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
func (m *DeviceManagementIntentDeviceStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["notApplicablePlatformCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicablePlatformCount(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicableCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicablePlatformCount gets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicablePlatformCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicablePlatformCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSuccessCount gets the successCount property value. Number of succeeded devices
// returns a *int32 when successful
func (m *DeviceManagementIntentDeviceStateSummary) GetSuccessCount()(*int32) {
    val, err := m.GetBackingStore().Get("successCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentDeviceStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("notApplicablePlatformCount", m.GetNotApplicablePlatformCount())
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
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceStateSummary) SetConflictCount(value *int32)() {
    err := m.GetBackingStore().Set("conflictCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceStateSummary) SetErrorCount(value *int32)() {
    err := m.GetBackingStore().Set("errorCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFailedCount sets the failedCount property value. Number of failed devices
func (m *DeviceManagementIntentDeviceStateSummary) SetFailedCount(value *int32)() {
    err := m.GetBackingStore().Set("failedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicableCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicablePlatformCount sets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicablePlatformCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicablePlatformCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessCount sets the successCount property value. Number of succeeded devices
func (m *DeviceManagementIntentDeviceStateSummary) SetSuccessCount(value *int32)() {
    err := m.GetBackingStore().Set("successCount", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementIntentDeviceStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetFailedCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetNotApplicablePlatformCount()(*int32)
    GetSuccessCount()(*int32)
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetFailedCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetNotApplicablePlatformCount(value *int32)()
    SetSuccessCount(value *int32)()
}
