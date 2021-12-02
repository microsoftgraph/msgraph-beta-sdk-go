package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementScriptRunSummary 
type DeviceManagementScriptRunSummary struct {
    Entity
    // Error device count.
    errorDeviceCount *int32;
    // Error user count.
    errorUserCount *int32;
    // Success device count.
    successDeviceCount *int32;
    // Success user count.
    successUserCount *int32;
}
// NewDeviceManagementScriptRunSummary instantiates a new deviceManagementScriptRunSummary and sets the default values.
func NewDeviceManagementScriptRunSummary()(*DeviceManagementScriptRunSummary) {
    m := &DeviceManagementScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// GetErrorUserCount gets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// GetSuccessDeviceCount gets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
// GetSuccessUserCount gets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) GetSuccessUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
        }
        return nil
    }
    res["successDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessDeviceCount(val)
        }
        return nil
    }
    res["successUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessUserCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementScriptRunSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptRunSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successUserCount", m.GetSuccessUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) SetErrorDeviceCount(value *int32)() {
    if m != nil {
        m.errorDeviceCount = value
    }
}
// SetErrorUserCount sets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) SetErrorUserCount(value *int32)() {
    if m != nil {
        m.errorUserCount = value
    }
}
// SetSuccessDeviceCount sets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) SetSuccessDeviceCount(value *int32)() {
    if m != nil {
        m.successDeviceCount = value
    }
}
// SetSuccessUserCount sets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) SetSuccessUserCount(value *int32)() {
    if m != nil {
        m.successUserCount = value
    }
}
