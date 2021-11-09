package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConfigurationManagerClientInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Configuration Manager Client Id from SCCM
    clientIdentifier *string;
    // Configuration Manager Client blocked status from SCCM
    isBlocked *bool;
}
// Instantiates a new configurationManagerClientInformation and sets the default values.
func NewConfigurationManagerClientInformation()(*ConfigurationManagerClientInformation) {
    m := &ConfigurationManagerClientInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the clientIdentifier property value. Configuration Manager Client Id from SCCM
func (m *ConfigurationManagerClientInformation) GetClientIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientIdentifier
    }
}
// Gets the isBlocked property value. Configuration Manager Client blocked status from SCCM
func (m *ConfigurationManagerClientInformation) GetIsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBlocked
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ConfigurationManagerClientInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the clientIdentifier property value. Configuration Manager Client Id from SCCM
// Parameters:
//  - value : Value to set for the clientIdentifier property.
func (m *ConfigurationManagerClientInformation) SetClientIdentifier(value *string)() {
    m.clientIdentifier = value
}
// Sets the isBlocked property value. Configuration Manager Client blocked status from SCCM
// Parameters:
//  - value : Value to set for the isBlocked property.
func (m *ConfigurationManagerClientInformation) SetIsBlocked(value *bool)() {
    m.isBlocked = value
}
