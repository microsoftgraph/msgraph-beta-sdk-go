package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettingDependency struct {
    additionalData map[string]interface{};
    constraints []DeviceManagementConstraint;
    definitionId *string;
}
func NewDeviceManagementSettingDependency()(*DeviceManagementSettingDependency) {
    m := &DeviceManagementSettingDependency{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementSettingDependency) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementSettingDependency) GetConstraints()([]DeviceManagementConstraint) {
    if m == nil {
        return nil
    } else {
        return m.constraints
    }
}
func (m *DeviceManagementSettingDependency) GetDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.definitionId
    }
}
func (m *DeviceManagementSettingDependency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["constraints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConstraint() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConstraint, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConstraint))
        }
        m.SetConstraints(res)
        return nil
    }
    res["definitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefinitionId(val)
        return nil
    }
    return res
}
func (m *DeviceManagementSettingDependency) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementSettingDependency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("definitionId", m.GetDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementSettingDependency) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementSettingDependency) SetConstraints(value []DeviceManagementConstraint)() {
    m.constraints = value
}
func (m *DeviceManagementSettingDependency) SetDefinitionId(value *string)() {
    m.definitionId = value
}
