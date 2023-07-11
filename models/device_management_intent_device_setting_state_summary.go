package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceSettingStateSummary entity that represents device setting state summary for an intent
type DeviceManagementIntentDeviceSettingStateSummary struct {
    Entity
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
    val, err := m.GetBackingStore().Get("compliantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetConflictCount()(*int32) {
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
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetErrorCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantCount(val)
        }
        return nil
    }
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
    res["nonCompliantCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantCount(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["remediatedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedCount(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    return res
}
// GetNonCompliantCount gets the nonCompliantCount property value. Number of non compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNonCompliantCount()(*int32) {
    val, err := m.GetBackingStore().Get("nonCompliantCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNotApplicableCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediatedCount gets the remediatedCount property value. Number of remediated devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetRemediatedCount()(*int32) {
    val, err := m.GetBackingStore().Get("remediatedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSettingName gets the settingName property value. Name of a setting
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetSettingName()(*string) {
    val, err := m.GetBackingStore().Get("settingName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("compliantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetConflictCount(value *int32)() {
    err := m.GetBackingStore().Set("conflictCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetErrorCount(value *int32)() {
    err := m.GetBackingStore().Set("errorCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNonCompliantCount sets the nonCompliantCount property value. Number of non compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNonCompliantCount(value *int32)() {
    err := m.GetBackingStore().Set("nonCompliantCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNotApplicableCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediatedCount sets the remediatedCount property value. Number of remediated devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetRemediatedCount(value *int32)() {
    err := m.GetBackingStore().Set("remediatedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingName sets the settingName property value. Name of a setting
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetSettingName(value *string)() {
    err := m.GetBackingStore().Set("settingName", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementIntentDeviceSettingStateSummaryable 
type DeviceManagementIntentDeviceSettingStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantCount()(*int32)
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetNonCompliantCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetOdataType()(*string)
    GetRemediatedCount()(*int32)
    GetSettingName()(*string)
    SetCompliantCount(value *int32)()
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetNonCompliantCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetOdataType(value *string)()
    SetRemediatedCount(value *int32)()
    SetSettingName(value *string)()
}
