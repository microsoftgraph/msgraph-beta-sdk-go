package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentDeviceStateSummary 
type DeviceManagementIntentDeviceStateSummary struct {
    Entity
    // Number of devices in conflict
    conflictCount *int32;
    // Number of error devices
    errorCount *int32;
    // Number of failed devices
    failedCount *int32;
    // Number of not applicable devices
    notApplicableCount *int32;
    // Number of not applicable devices due to mismatch platform and policy
    notApplicablePlatformCount *int32;
    // Number of succeeded devices
    successCount *int32;
}
// NewDeviceManagementIntentDeviceStateSummary instantiates a new deviceManagementIntentDeviceStateSummary and sets the default values.
func NewDeviceManagementIntentDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    m := &DeviceManagementIntentDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceStateSummary) GetConflictCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.conflictCount
    }
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceStateSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// GetFailedCount gets the failedCount property value. Number of failed devices
func (m *DeviceManagementIntentDeviceStateSummary) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// GetNotApplicablePlatformCount gets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicablePlatformCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicablePlatformCount
    }
}
// GetSuccessCount gets the successCount property value. Number of succeeded devices
func (m *DeviceManagementIntentDeviceStateSummary) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["notApplicablePlatformCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicablePlatformCount(val)
        }
        return nil
    }
    res["successCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceManagementIntentDeviceStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceStateSummary) SetConflictCount(value *int32)() {
    if m != nil {
        m.conflictCount = value
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceStateSummary) SetErrorCount(value *int32)() {
    if m != nil {
        m.errorCount = value
    }
}
// SetFailedCount sets the failedCount property value. Number of failed devices
func (m *DeviceManagementIntentDeviceStateSummary) SetFailedCount(value *int32)() {
    if m != nil {
        m.failedCount = value
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicableCount(value *int32)() {
    if m != nil {
        m.notApplicableCount = value
    }
}
// SetNotApplicablePlatformCount sets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicablePlatformCount(value *int32)() {
    if m != nil {
        m.notApplicablePlatformCount = value
    }
}
// SetSuccessCount sets the successCount property value. Number of succeeded devices
func (m *DeviceManagementIntentDeviceStateSummary) SetSuccessCount(value *int32)() {
    if m != nil {
        m.successCount = value
    }
}
