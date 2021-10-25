package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsSettings struct {
    additionalData map[string]interface{};
    configurationManagerDataConnectorConfigured *bool;
}
func NewUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    m := &UserExperienceAnalyticsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserExperienceAnalyticsSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserExperienceAnalyticsSettings) GetConfigurationManagerDataConnectorConfigured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerDataConnectorConfigured
    }
}
func (m *UserExperienceAnalyticsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configurationManagerDataConnectorConfigured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetConfigurationManagerDataConnectorConfigured(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsSettings) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("configurationManagerDataConnectorConfigured", m.GetConfigurationManagerDataConnectorConfigured())
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
func (m *UserExperienceAnalyticsSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserExperienceAnalyticsSettings) SetConfigurationManagerDataConnectorConfigured(value *bool)() {
    m.configurationManagerDataConnectorConfigured = value
}
