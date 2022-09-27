package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdvancedThreatProtectionOnboardingStateSummary 
type AdvancedThreatProtectionOnboardingStateSummary struct {
    Entity
    // Not yet documented
    advancedThreatProtectionOnboardingDeviceSettingStates []AdvancedThreatProtectionOnboardingDeviceSettingStateable
    // Number of compliant devices
    compliantDeviceCount *int32
    // Number of conflict devices
    conflictDeviceCount *int32
    // Number of error devices
    errorDeviceCount *int32
    // Number of NonCompliant devices
    nonCompliantDeviceCount *int32
    // Number of not applicable devices
    notApplicableDeviceCount *int32
    // Number of not assigned devices
    notAssignedDeviceCount *int32
    // Number of remediated devices
    remediatedDeviceCount *int32
    // Number of unknown devices
    unknownDeviceCount *int32
}
// NewAdvancedThreatProtectionOnboardingStateSummary instantiates a new advancedThreatProtectionOnboardingStateSummary and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    m := &AdvancedThreatProtectionOnboardingStateSummary{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.advancedThreatProtectionOnboardingStateSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdvancedThreatProtectionOnboardingStateSummary(), nil
}
// GetAdvancedThreatProtectionOnboardingDeviceSettingStates gets the advancedThreatProtectionOnboardingDeviceSettingStates property value. Not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingStateable) {
    return m.advancedThreatProtectionOnboardingDeviceSettingStates
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetConflictDeviceCount()(*int32) {
    return m.conflictDeviceCount
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of error devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedThreatProtectionOnboardingDeviceSettingStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue , m.SetAdvancedThreatProtectionOnboardingDeviceSettingStates)
    res["compliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompliantDeviceCount)
    res["conflictDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetConflictDeviceCount)
    res["errorDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetErrorDeviceCount)
    res["nonCompliantDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNonCompliantDeviceCount)
    res["notApplicableDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNotApplicableDeviceCount)
    res["notAssignedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetNotAssignedDeviceCount)
    res["remediatedDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemediatedDeviceCount)
    res["unknownDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnknownDeviceCount)
    return res
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    return m.nonCompliantDeviceCount
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    return m.notApplicableDeviceCount
}
// GetNotAssignedDeviceCount gets the notAssignedDeviceCount property value. Number of not assigned devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotAssignedDeviceCount()(*int32) {
    return m.notAssignedDeviceCount
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetRemediatedDeviceCount()(*int32) {
    return m.remediatedDeviceCount
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
}
// Serialize serializes information the current object
func (m *AdvancedThreatProtectionOnboardingStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates())
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
    m.advancedThreatProtectionOnboardingDeviceSettingStates = value
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of error devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// SetNotAssignedDeviceCount sets the notAssignedDeviceCount property value. Number of not assigned devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotAssignedDeviceCount(value *int32)() {
    m.notAssignedDeviceCount = value
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
