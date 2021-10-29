package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AdvancedThreatProtectionOnboardingStateSummary struct {
    Entity
    // Not yet documented
    advancedThreatProtectionOnboardingDeviceSettingStates []AdvancedThreatProtectionOnboardingDeviceSettingState;
    // Number of compliant devices
    compliantDeviceCount *int32;
    // Number of conflict devices
    conflictDeviceCount *int32;
    // Number of error devices
    errorDeviceCount *int32;
    // Number of NonCompliant devices
    nonCompliantDeviceCount *int32;
    // Number of not applicable devices
    notApplicableDeviceCount *int32;
    // Number of not assigned devices
    notAssignedDeviceCount *int32;
    // Number of remediated devices
    remediatedDeviceCount *int32;
    // Number of unknown devices
    unknownDeviceCount *int32;
}
// Instantiates a new advancedThreatProtectionOnboardingStateSummary and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    m := &AdvancedThreatProtectionOnboardingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the advancedThreatProtectionOnboardingDeviceSettingStates property value. Not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingState) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionOnboardingDeviceSettingStates
    }
}
// Gets the compliantDeviceCount property value. Number of compliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
// Gets the conflictDeviceCount property value. Number of conflict devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
// Gets the errorDeviceCount property value. Number of error devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
// Gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
// Gets the notAssignedDeviceCount property value. Number of not assigned devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotAssignedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notAssignedDeviceCount
    }
}
// Gets the remediatedDeviceCount property value. Number of remediated devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
// Gets the unknownDeviceCount property value. Number of unknown devices
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// The deserialization information for the current model
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["advancedThreatProtectionOnboardingDeviceSettingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAdvancedThreatProtectionOnboardingDeviceSettingState() })
        if err != nil {
            return err
        }
        res := make([]AdvancedThreatProtectionOnboardingDeviceSettingState, len(val))
        for i, v := range val {
            res[i] = *(v.(*AdvancedThreatProtectionOnboardingDeviceSettingState))
        }
        m.SetAdvancedThreatProtectionOnboardingDeviceSettingStates(res)
        return nil
    }
    res["compliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantDeviceCount(val)
        return nil
    }
    res["conflictDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConflictDeviceCount(val)
        return nil
    }
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorDeviceCount(val)
        return nil
    }
    res["nonCompliantDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantDeviceCount(val)
        return nil
    }
    res["notApplicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableDeviceCount(val)
        return nil
    }
    res["notAssignedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotAssignedDeviceCount(val)
        return nil
    }
    res["remediatedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedDeviceCount(val)
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownDeviceCount(val)
        return nil
    }
    return res
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AdvancedThreatProtectionOnboardingStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates()))
        for i, v := range m.GetAdvancedThreatProtectionOnboardingDeviceSettingStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the advancedThreatProtectionOnboardingDeviceSettingStates property value. Not yet documented
// Parameters:
//  - value : Value to set for the advancedThreatProtectionOnboardingDeviceSettingStates property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetAdvancedThreatProtectionOnboardingDeviceSettingStates(value []AdvancedThreatProtectionOnboardingDeviceSettingState)() {
    m.advancedThreatProtectionOnboardingDeviceSettingStates = value
}
// Sets the compliantDeviceCount property value. Number of compliant devices
// Parameters:
//  - value : Value to set for the compliantDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// Sets the conflictDeviceCount property value. Number of conflict devices
// Parameters:
//  - value : Value to set for the conflictDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// Sets the errorDeviceCount property value. Number of error devices
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
// Parameters:
//  - value : Value to set for the nonCompliantDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// Sets the notApplicableDeviceCount property value. Number of not applicable devices
// Parameters:
//  - value : Value to set for the notApplicableDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// Sets the notAssignedDeviceCount property value. Number of not assigned devices
// Parameters:
//  - value : Value to set for the notAssignedDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotAssignedDeviceCount(value *int32)() {
    m.notAssignedDeviceCount = value
}
// Sets the remediatedDeviceCount property value. Number of remediated devices
// Parameters:
//  - value : Value to set for the remediatedDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// Sets the unknownDeviceCount property value. Number of unknown devices
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
