package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementIntentDeviceStateSummary struct {
    Entity
    conflictCount *int32;
    errorCount *int32;
    failedCount *int32;
    notApplicableCount *int32;
    notApplicablePlatformCount *int32;
    successCount *int32;
}
func NewDeviceManagementIntentDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    m := &DeviceManagementIntentDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementIntentDeviceStateSummary) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicablePlatformCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicablePlatformCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
func (m *DeviceManagementIntentDeviceStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["failedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedCount(val)
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
    res["notApplicablePlatformCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotApplicablePlatformCount(val)
        return nil
    }
    res["successCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuccessCount(val)
        return nil
    }
    return res
}
func (m *DeviceManagementIntentDeviceStateSummary) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementIntentDeviceStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *DeviceManagementIntentDeviceStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
func (m *DeviceManagementIntentDeviceStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
func (m *DeviceManagementIntentDeviceStateSummary) SetFailedCount(value *int32)() {
    m.failedCount = value
}
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicablePlatformCount(value *int32)() {
    m.notApplicablePlatformCount = value
}
func (m *DeviceManagementIntentDeviceStateSummary) SetSuccessCount(value *int32)() {
    m.successCount = value
}
