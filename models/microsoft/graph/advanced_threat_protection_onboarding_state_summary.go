package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AdvancedThreatProtectionOnboardingStateSummary struct {
    Entity
    advancedThreatProtectionOnboardingDeviceSettingStates []AdvancedThreatProtectionOnboardingDeviceSettingState;
    compliantDeviceCount *int32;
    conflictDeviceCount *int32;
    errorDeviceCount *int32;
    nonCompliantDeviceCount *int32;
    notApplicableDeviceCount *int32;
    notAssignedDeviceCount *int32;
    remediatedDeviceCount *int32;
    unknownDeviceCount *int32;
}
func NewAdvancedThreatProtectionOnboardingStateSummary()(*AdvancedThreatProtectionOnboardingStateSummary) {
    m := &AdvancedThreatProtectionOnboardingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetAdvancedThreatProtectionOnboardingDeviceSettingStates()([]AdvancedThreatProtectionOnboardingDeviceSettingState) {
    if m == nil {
        return nil
    } else {
        return m.advancedThreatProtectionOnboardingDeviceSettingStates
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetConflictDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNonCompliantDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetNotAssignedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notAssignedDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetRemediatedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedDeviceCount
    }
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
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
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetAdvancedThreatProtectionOnboardingDeviceSettingStates(value []AdvancedThreatProtectionOnboardingDeviceSettingState)() {
    m.advancedThreatProtectionOnboardingDeviceSettingStates = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetNotAssignedDeviceCount(value *int32)() {
    m.notAssignedDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
func (m *AdvancedThreatProtectionOnboardingStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
