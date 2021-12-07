package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateLabelJobResultGroup 
type EvaluateLabelJobResultGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    automatic *EvaluateLabelJobResult;
    // 
    recommended *EvaluateLabelJobResult;
}
// NewEvaluateLabelJobResultGroup instantiates a new evaluateLabelJobResultGroup and sets the default values.
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutomatic gets the automatic property value. 
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.automatic
    }
}
// GetRecommended gets the recommended property value. 
func (m *EvaluateLabelJobResultGroup) GetRecommended()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.recommended
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutomatic sets the automatic property value. 
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value *EvaluateLabelJobResult)() {
    if m != nil {
        m.automatic = value
    }
}
// SetRecommended sets the recommended property value. 
func (m *EvaluateLabelJobResultGroup) SetRecommended(value *EvaluateLabelJobResult)() {
    if m != nil {
        m.recommended = value
    }
}
