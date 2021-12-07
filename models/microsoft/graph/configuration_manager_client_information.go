package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigurationManagerClientInformation 
type ConfigurationManagerClientInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Configuration Manager Client Id from SCCM
    clientIdentifier *string;
    // Configuration Manager Client blocked status from SCCM
    isBlocked *bool;
}
// NewConfigurationManagerClientInformation instantiates a new configurationManagerClientInformation and sets the default values.
func NewConfigurationManagerClientInformation()(*ConfigurationManagerClientInformation) {
    m := &ConfigurationManagerClientInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientIdentifier gets the clientIdentifier property value. Configuration Manager Client Id from SCCM
func (m *ConfigurationManagerClientInformation) GetClientIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientIdentifier
    }
}
// GetIsBlocked gets the isBlocked property value. Configuration Manager Client blocked status from SCCM
func (m *ConfigurationManagerClientInformation) GetIsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clientIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientIdentifier(val)
        }
        return nil
    }
    res["isBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBlocked(val)
        }
        return nil
    }
    return res
}
func (m *ConfigurationManagerClientInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientIdentifier sets the clientIdentifier property value. Configuration Manager Client Id from SCCM
func (m *ConfigurationManagerClientInformation) SetClientIdentifier(value *string)() {
    if m != nil {
        m.clientIdentifier = value
    }
}
// SetIsBlocked sets the isBlocked property value. Configuration Manager Client blocked status from SCCM
func (m *ConfigurationManagerClientInformation) SetIsBlocked(value *bool)() {
    if m != nil {
        m.isBlocked = value
    }
}
