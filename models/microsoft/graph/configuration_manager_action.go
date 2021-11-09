package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConfigurationManagerAction struct {
    // The action type to trigger on Configuration Manager client. Possible values are: refreshMachinePolicy, refreshUserPolicy, wakeUpClient, appEvaluation.
    action *ConfigurationManagerActionType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// Instantiates a new configurationManagerAction and sets the default values.
func NewConfigurationManagerAction()(*ConfigurationManagerAction) {
    m := &ConfigurationManagerAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the action property value. The action type to trigger on Configuration Manager client. Possible values are: refreshMachinePolicy, refreshUserPolicy, wakeUpClient, appEvaluation.
func (m *ConfigurationManagerAction) GetAction()(*ConfigurationManagerActionType) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// The deserialization information for the current model
func (m *ConfigurationManagerAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerActionType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ConfigurationManagerActionType)
            m.SetAction(&cast)
        }
        return nil
    }
    return res
}
func (m *ConfigurationManagerAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the action property value. The action type to trigger on Configuration Manager client. Possible values are: refreshMachinePolicy, refreshUserPolicy, wakeUpClient, appEvaluation.
// Parameters:
//  - value : Value to set for the action property.
func (m *ConfigurationManagerAction) SetAction(value *ConfigurationManagerActionType)() {
    m.action = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ConfigurationManagerAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
