package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementIntentDeviceSettingStateSummary struct {
    Entity
    // Number of compliant devices
    compliantCount *int32;
    // Number of devices in conflict
    conflictCount *int32;
    // Number of error devices
    errorCount *int32;
    // Number of non compliant devices
    nonCompliantCount *int32;
    // Number of not applicable devices
    notApplicableCount *int32;
    // Number of remediated devices
    remediatedCount *int32;
    // Name of a setting
    settingName *string;
}
// Instantiates a new deviceManagementIntentDeviceSettingStateSummary and sets the default values.
func NewDeviceManagementIntentDeviceSettingStateSummary()(*DeviceManagementIntentDeviceSettingStateSummary) {
    m := &DeviceManagementIntentDeviceSettingStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the compliantCount property value. Number of compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetCompliantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliantCount
    }
}
// Gets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
// Gets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// Gets the nonCompliantCount property value. Number of non compliant devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNonCompliantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.nonCompliantCount
    }
}
// Gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// Gets the remediatedCount property value. Number of remediated devices
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetRemediatedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remediatedCount
    }
}
// Gets the settingName property value. Name of a setting
func (m *DeviceManagementIntentDeviceSettingStateSummary) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the compliantCount property value. Number of compliant devices
// Parameters:
//  - value : Value to set for the compliantCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetCompliantCount(value *int32)() {
    m.compliantCount = value
}
// Sets the conflictCount property value. Number of devices in conflict
// Parameters:
//  - value : Value to set for the conflictCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
// Sets the errorCount property value. Number of error devices
// Parameters:
//  - value : Value to set for the errorCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// Sets the nonCompliantCount property value. Number of non compliant devices
// Parameters:
//  - value : Value to set for the nonCompliantCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNonCompliantCount(value *int32)() {
    m.nonCompliantCount = value
}
// Sets the notApplicableCount property value. Number of not applicable devices
// Parameters:
//  - value : Value to set for the notApplicableCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// Sets the remediatedCount property value. Number of remediated devices
// Parameters:
//  - value : Value to set for the remediatedCount property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetRemediatedCount(value *int32)() {
    m.remediatedCount = value
}
// Sets the settingName property value. Name of a setting
// Parameters:
//  - value : Value to set for the settingName property.
func (m *DeviceManagementIntentDeviceSettingStateSummary) SetSettingName(value *string)() {
    m.settingName = value
}
