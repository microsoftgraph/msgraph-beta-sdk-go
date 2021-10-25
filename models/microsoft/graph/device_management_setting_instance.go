package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettingInstance struct {
    Entity
    definitionId *string;
    valueJson *string;
}
func NewDeviceManagementSettingInstance()(*DeviceManagementSettingInstance) {
    m := &DeviceManagementSettingInstance{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementSettingInstance) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
func (m *DeviceManagementSettingInstance) GetValueJson()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueJson
    }
}
func (m *DeviceManagementSettingInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefinitionId(val)
        return nil
    }
    res["valueJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValueJson(val)
        return nil
    }
    return res
}
func (m *DeviceManagementSettingInstance) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementSettingInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("definitionId", m.GetDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueJson", m.GetValueJson())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementSettingInstance) SetDefinitionId(value *string)() {
    m.definitionId = value
}
func (m *DeviceManagementSettingInstance) SetValueJson(value *string)() {
    m.valueJson = value
}
