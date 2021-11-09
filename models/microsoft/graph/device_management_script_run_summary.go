package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceManagementScriptRunSummary and sets the default values.
func NewDeviceManagementScriptRunSummary()(*DeviceManagementScriptRunSummary) {
    m := &DeviceManagementScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the errorDeviceCount property value. Error device count.
func (m *DeviceManagementScriptRunSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
// Gets the errorUserCount property value. Error user count.
func (m *DeviceManagementScriptRunSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
// Gets the successDeviceCount property value. Success device count.
func (m *DeviceManagementScriptRunSummary) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
// Gets the successUserCount property value. Success user count.
func (m *DeviceManagementScriptRunSummary) GetSuccessUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successUserCount
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the errorDeviceCount property value. Error device count.
// Parameters:
//  - value : Value to set for the errorDeviceCount property.
func (m *DeviceManagementScriptRunSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// Sets the errorUserCount property value. Error user count.
// Parameters:
//  - value : Value to set for the errorUserCount property.
func (m *DeviceManagementScriptRunSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
// Sets the successDeviceCount property value. Success device count.
// Parameters:
//  - value : Value to set for the successDeviceCount property.
func (m *DeviceManagementScriptRunSummary) SetSuccessDeviceCount(value *int32)() {
    m.successDeviceCount = value
}
// Sets the successUserCount property value. Success user count.
// Parameters:
//  - value : Value to set for the successUserCount property.
func (m *DeviceManagementScriptRunSummary) SetSuccessUserCount(value *int32)() {
    m.successUserCount = value
}
