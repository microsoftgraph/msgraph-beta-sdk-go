package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConfigurationManagerAction struct {
    action *ConfigurationManagerActionType;
    additionalData map[string]interface{};
}
func NewConfigurationManagerAction()(*ConfigurationManagerAction) {
    m := &ConfigurationManagerAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConfigurationManagerAction) GetAction()(*ConfigurationManagerActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *ConfigurationManagerAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConfigurationManagerAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerActionType)
        if err != nil {
            return err
        }
        cast := val.(ConfigurationManagerActionType)
        m.SetAction(&cast)
        return nil
    }
    return res
}
func (m *ConfigurationManagerAction) IsNil()(bool) {
    return m == nil
}
func (m *ConfigurationManagerAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
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
func (m *ConfigurationManagerAction) SetAction(value *ConfigurationManagerActionType)() {
    m.action = value
}
func (m *ConfigurationManagerAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
