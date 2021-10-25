package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementIntentDeviceSettingStateSummary struct {
    Entity
    compliantCount *int32;
    conflictCount *int32;
    errorCount *int32;
    nonCompliantCount *int32;
    notApplicableCount *int32;
    remediatedCount *int32;
    settingName *string;
}
func NewDeviceManagementIntentDeviceSettingStateSummary()(*DeviceManagementIntentDeviceSettingStateSummary) {
    m := &DeviceManagementIntentDeviceSettingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetCompliantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNonCompliantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetRemediatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedCount
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliantCount(val)
        return nil
    }
    res["conflictCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConflictCount(val)
        return nil
    }
    res["errorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorCount(val)
        return nil
    }
    res["nonCompliantCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNonCompliantCount(val)
        return nil
    }
    res["notApplicableCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicableCount(val)
        return nil
    }
    res["remediatedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemediatedCount(val)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    return res
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetCompliantCount(value *int32)() {
    m.compliantCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNonCompliantCount(value *int32)() {
    m.nonCompliantCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetRemediatedCount(value *int32)() {
    m.remediatedCount = value
}
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetSettingName(value *string)() {
    m.settingName = value
}
