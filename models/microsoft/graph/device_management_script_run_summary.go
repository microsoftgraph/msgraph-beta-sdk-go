package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementScriptRunSummary struct {
    Entity
    errorDeviceCount *int32;
    errorUserCount *int32;
    successDeviceCount *int32;
    successUserCount *int32;
}
func NewDeviceManagementScriptRunSummary()(*DeviceManagementScriptRunSummary) {
    m := &DeviceManagementScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementScriptRunSummary) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
func (m *DeviceManagementScriptRunSummary) GetErrorUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorUserCount
    }
}
func (m *DeviceManagementScriptRunSummary) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
func (m *DeviceManagementScriptRunSummary) GetSuccessUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successUserCount
    }
}
func (m *DeviceManagementScriptRunSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["errorDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorDeviceCount(val)
        return nil
    }
    res["errorUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorUserCount(val)
        return nil
    }
    res["successDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuccessDeviceCount(val)
        return nil
    }
    res["successUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuccessUserCount(val)
        return nil
    }
    return res
}
func (m *DeviceManagementScriptRunSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceManagementScriptRunSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
func (m *DeviceManagementScriptRunSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
func (m *DeviceManagementScriptRunSummary) SetSuccessDeviceCount(value *int32)() {
    m.successDeviceCount = value
}
func (m *DeviceManagementScriptRunSummary) SetSuccessUserCount(value *int32)() {
    m.successUserCount = value
}
