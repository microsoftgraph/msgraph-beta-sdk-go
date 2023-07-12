package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdvancedThreatProtectionOnboardingStateSummary windows defender advanced threat protection onboarding state summary across the account.
type AdvancedThreatProtectionOnboardingStateSummary struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewAdvancedThreatProtectionOnboardingStateSummary instantiates a new advancedThreatProtectionOnboardingStateSummary and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    m := &AdvancedThreatProtectionOnboardingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedThreatProtectionOnboardingStateSummary(), nil
}
// GetAdvancedThreatProtectionOnboardingDeviceSettingStates gets the advancedThreatProtectionOnboardingDeviceSettingStates property value. Not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingStateable) {
    val, err := m.GetBackingStore().Get("advancedThreatProtectionOnboardingDeviceSettingStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AdvancedThreatProtectionOnboardingDeviceSettingStateable)
    }
    return nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetCompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetConflictDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("conflictDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of error devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetErrorDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("errorDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedThreatProtectionOnboardingDeviceSettingStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdvancedThreatProtectionOnboardingDeviceSettingStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AdvancedThreatProtectionOnboardingDeviceSettingStateable)
                }
            }
            m.SetAdvancedThreatProtectionOnboardingDeviceSettingStates(res)
        }
        return nil
    }
    res["compliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["conflictDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictDeviceCount(val)
        }
        return nil
    }
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
    res["nonCompliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantDeviceCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["notAssignedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotAssignedDeviceCount(val)
        }
        return nil
    }
    res["remediatedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedDeviceCount(val)
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
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("nonCompliantDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("notApplicableDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetNotAssignedDeviceCount gets the notAssignedDeviceCount property value. Number of not assigned devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotAssignedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("notAssignedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetRemediatedDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("remediatedDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetUnknownDeviceCount()(*int32) {
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
func (m *AdvancedThreatProtectionOnboardingStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates()))
        for i, v := range m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("advancedThreatProtectionOnboardingDeviceSettingStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictDeviceCount", m.GetConflictDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notAssignedDeviceCount", m.GetNotAssignedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
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
// SetAdvancedThreatProtectionOnboardingDeviceSettingStates sets the advancedThreatProtectionOnboardingDeviceSettingStates property value. Not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetAdvancedThreatProtectionOnboardingDeviceSettingStates(value []AdvancedThreatProtectionOnboardingDeviceSettingStateable)() {
    err := m.GetBackingStore().Set("advancedThreatProtectionOnboardingDeviceSettingStates", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetCompliantDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("compliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetConflictDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("conflictDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of error devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetErrorDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("errorDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("nonCompliantDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("notApplicableDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetNotAssignedDeviceCount sets the notAssignedDeviceCount property value. Number of not assigned devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotAssignedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("notAssignedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("remediatedDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetUnknownDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("unknownDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// AdvancedThreatProtectionOnboardingStateSummaryable 
type AdvancedThreatProtectionOnboardingStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingStateable)
    GetCompliantDeviceCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotAssignedDeviceCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetAdvancedThreatProtectionOnboardingDeviceSettingStates(value []AdvancedThreatProtectionOnboardingDeviceSettingStateable)()
    SetCompliantDeviceCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotAssignedDeviceCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}
