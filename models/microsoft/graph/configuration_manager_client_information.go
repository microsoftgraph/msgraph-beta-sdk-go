package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConfigurationManagerClientInformation struct {
    additionalData map[string]interface{};
    clientIdentifier *string;
    isBlocked *bool;
}
func NewConfigurationManagerClientInformation()(*ConfigurationManagerClientInformation) {
    m := &ConfigurationManagerClientInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConfigurationManagerClientInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConfigurationManagerClientInformation) GetClientIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientIdentifier
    }
}
func (m *ConfigurationManagerClientInformation) GetIsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBlocked
    }
}
func (m *ConfigurationManagerClientInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientIdentifier(val)
        return nil
    }
    res["isBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBlocked(val)
        return nil
    }
    return res
}
func (m *ConfigurationManagerClientInformation) IsNil()(bool) {
    return m == nil
}
func (m *ConfigurationManagerClientInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientIdentifier", m.GetClientIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBlocked", m.GetIsBlocked())
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
func (m *ConfigurationManagerClientInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConfigurationManagerClientInformation) SetClientIdentifier(value *string)() {
    m.clientIdentifier = value
}
func (m *ConfigurationManagerClientInformation) SetIsBlocked(value *bool)() {
    m.isBlocked = value
}
