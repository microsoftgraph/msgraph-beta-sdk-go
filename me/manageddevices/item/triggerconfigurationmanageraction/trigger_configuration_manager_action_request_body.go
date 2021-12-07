package triggerconfigurationmanageraction

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// TriggerConfigurationManagerActionRequestBody 
type TriggerConfigurationManagerActionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Parameter for action triggerConfigurationManagerAction
    configurationManagerAction *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConfigurationManagerAction;
}
// NewTriggerConfigurationManagerActionRequestBody instantiates a new triggerConfigurationManagerActionRequestBody and sets the default values.
func NewTriggerConfigurationManagerActionRequestBody()(*TriggerConfigurationManagerActionRequestBody) {
    m := &TriggerConfigurationManagerActionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TriggerConfigurationManagerActionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfigurationManagerAction gets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *TriggerConfigurationManagerActionRequestBody) GetConfigurationManagerAction()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConfigurationManagerAction) {
    if m == nil {
        return nil
    } else {
        return m.configurationManagerAction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggerConfigurationManagerActionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configurationManagerAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewConfigurationManagerAction() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerAction(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConfigurationManagerAction))
        }
        return nil
    }
    return res
}
func (m *TriggerConfigurationManagerActionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TriggerConfigurationManagerActionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("configurationManagerAction", m.GetConfigurationManagerAction())
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
func (m *TriggerConfigurationManagerActionRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfigurationManagerAction sets the configurationManagerAction property value. Parameter for action triggerConfigurationManagerAction
func (m *TriggerConfigurationManagerActionRequestBody) SetConfigurationManagerAction(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ConfigurationManagerAction)() {
    if m != nil {
        m.configurationManagerAction = value
    }
}
