package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EvaluateLabelJobResult 
type EvaluateLabelJobResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    responsiblePolicy *ResponsiblePolicy;
    // 
    responsibleSensitiveTypes []ResponsibleSensitiveType;
    // 
    sensitivityLabel *MatchingLabel;
}
// NewEvaluateLabelJobResult instantiates a new evaluateLabelJobResult and sets the default values.
func NewEvaluateLabelJobResult()(*EvaluateLabelJobResult) {
    m := &EvaluateLabelJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetResponsiblePolicy gets the responsiblePolicy property value. 
func (m *EvaluateLabelJobResult) GetResponsiblePolicy()(*ResponsiblePolicy) {
    if m == nil {
        return nil
    } else {
        return m.responsiblePolicy
    }
}
// GetResponsibleSensitiveTypes gets the responsibleSensitiveTypes property value. 
func (m *EvaluateLabelJobResult) GetResponsibleSensitiveTypes()([]ResponsibleSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.responsibleSensitiveTypes
    }
}
// GetSensitivityLabel gets the sensitivityLabel property value. 
func (m *EvaluateLabelJobResult) GetSensitivityLabel()(*MatchingLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabel
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EvaluateLabelJobResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["responsiblePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResponsiblePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiblePolicy(val.(*ResponsiblePolicy))
        }
        return nil
    }
    res["responsibleSensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResponsibleSensitiveType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResponsibleSensitiveType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ResponsibleSensitiveType))
            }
            m.SetResponsibleSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMatchingLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivityLabel(val.(*MatchingLabel))
        }
        return nil
    }
    return res
}
func (m *EvaluateLabelJobResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EvaluateLabelJobResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("responsiblePolicy", m.GetResponsiblePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetResponsibleSensitiveTypes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResponsibleSensitiveTypes()))
        for i, v := range m.GetResponsibleSensitiveTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("responsibleSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sensitivityLabel", m.GetSensitivityLabel())
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
func (m *EvaluateLabelJobResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetResponsiblePolicy sets the responsiblePolicy property value. 
func (m *EvaluateLabelJobResult) SetResponsiblePolicy(value *ResponsiblePolicy)() {
    if m != nil {
        m.responsiblePolicy = value
    }
}
// SetResponsibleSensitiveTypes sets the responsibleSensitiveTypes property value. 
func (m *EvaluateLabelJobResult) SetResponsibleSensitiveTypes(value []ResponsibleSensitiveType)() {
    if m != nil {
        m.responsibleSensitiveTypes = value
    }
}
// SetSensitivityLabel sets the sensitivityLabel property value. 
func (m *EvaluateLabelJobResult) SetSensitivityLabel(value *MatchingLabel)() {
    if m != nil {
        m.sensitivityLabel = value
    }
}
