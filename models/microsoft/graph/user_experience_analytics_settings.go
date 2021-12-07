package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsSettings 
type UserExperienceAnalyticsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
    configurationManagerDataConnectorConfigured *bool;
}
// NewUserExperienceAnalyticsSettings instantiates a new userExperienceAnalyticsSettings and sets the default values.
func NewUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    m := &UserExperienceAnalyticsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfigurationManagerDataConnectorConfigured gets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
func (m *UserExperienceAnalyticsSettings) GetConfigurationManagerDataConnectorConfigured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerDataConnectorConfigured
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configurationManagerDataConnectorConfigured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerDataConnectorConfigured(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfigurationManagerDataConnectorConfigured sets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
func (m *UserExperienceAnalyticsSettings) SetConfigurationManagerDataConnectorConfigured(value *bool)() {
    if m != nil {
        m.configurationManagerDataConnectorConfigured = value
    }
}
