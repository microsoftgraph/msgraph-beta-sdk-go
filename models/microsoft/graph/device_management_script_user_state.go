package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementScriptUserState struct {
    Entity
    deviceRunStates []DeviceManagementScriptDeviceState;
    errorDeviceCount *int32;
    successDeviceCount *int32;
    userPrincipalName *string;
}
func NewDeviceManagementScriptUserState()(*DeviceManagementScriptUserState) {
    m := &DeviceManagementScriptUserState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementScriptUserState) GetDeviceRunStates()([]DeviceManagementScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
func (m *DeviceManagementScriptUserState) GetErrorDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorDeviceCount
    }
}
func (m *DeviceManagementScriptUserState) GetSuccessDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successDeviceCount
    }
}
func (m *DeviceManagementScriptUserState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *DeviceManagementScriptUserState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptDeviceState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptDeviceState))
        }
        m.SetDeviceRunStates(res)
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
    res["successDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuccessDeviceCount(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *DeviceManagementScriptUserState) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementScriptUserState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
        err = writer.WriteInt32Value("successDeviceCount", m.GetSuccessDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementScriptUserState) SetDeviceRunStates(value []DeviceManagementScriptDeviceState)() {
    m.deviceRunStates = value
}
func (m *DeviceManagementScriptUserState) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
func (m *DeviceManagementScriptUserState) SetSuccessDeviceCount(value *int32)() {
    m.successDeviceCount = value
}
func (m *DeviceManagementScriptUserState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
