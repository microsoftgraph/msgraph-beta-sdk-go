package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EvaluateLabelJobResultGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    automatic *EvaluateLabelJobResult;
    // 
    recommended *EvaluateLabelJobResult;
}
// Instantiates a new evaluateLabelJobResultGroup and sets the default values.
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the automatic property value. 
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.automatic
    }
}
// Gets the recommended property value. 
func (m *EvaluateLabelJobResultGroup) GetRecommended()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.recommended
    }
}
// The deserialization information for the current model
func (m *EvaluateLabelJobResultGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automatic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateLabelJobResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomatic(val.(*EvaluateLabelJobResult))
        }
        return nil
    }
    res["recommended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateLabelJobResult() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommended(val.(*EvaluateLabelJobResult))
        }
        return nil
    }
    return res
}
func (m *EvaluateLabelJobResultGroup) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EvaluateLabelJobResultGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automatic", m.GetAutomatic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recommended", m.GetRecommended())
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
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the automatic property value. 
// Parameters:
//  - value : Value to set for the automatic property.
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value *EvaluateLabelJobResult)() {
    m.automatic = value
}
// Sets the recommended property value. 
// Parameters:
//  - value : Value to set for the recommended property.
func (m *EvaluateLabelJobResultGroup) SetRecommended(value *EvaluateLabelJobResult)() {
    m.recommended = value
}
