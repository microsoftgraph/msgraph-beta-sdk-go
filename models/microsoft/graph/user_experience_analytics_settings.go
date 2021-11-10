package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
    configurationManagerDataConnectorConfigured *bool;
}
// Instantiates a new userExperienceAnalyticsSettings and sets the default values.
func NewUserExperienceAnalyticsSettings()(*UserExperienceAnalyticsSettings) {
    m := &UserExperienceAnalyticsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
func (m *UserExperienceAnalyticsSettings) GetConfigurationManagerDataConnectorConfigured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerDataConnectorConfigured
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserExperienceAnalyticsSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the configurationManagerDataConnectorConfigured property value. True if Tenant attach is configured. If configured then SCCM tenant attached devices will show up in UXA reporting.
// Parameters:
//  - value : Value to set for the configurationManagerDataConnectorConfigured property.
func (m *UserExperienceAnalyticsSettings) SetConfigurationManagerDataConnectorConfigured(value *bool)() {
    m.configurationManagerDataConnectorConfigured = value
}
