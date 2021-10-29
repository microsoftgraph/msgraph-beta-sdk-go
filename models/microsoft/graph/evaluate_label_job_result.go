package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new evaluateLabelJobResult and sets the default values.
func NewEvaluateLabelJobResult()(*EvaluateLabelJobResult) {
    m := &EvaluateLabelJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EvaluateLabelJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the responsiblePolicy property value. 
func (m *EvaluateLabelJobResult) GetResponsiblePolicy()(*ResponsiblePolicy) {
    if m == nil {
        return nil
    } else {
        return m.responsiblePolicy
    }
}
// Gets the responsibleSensitiveTypes property value. 
func (m *EvaluateLabelJobResult) GetResponsibleSensitiveTypes()([]ResponsibleSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.responsibleSensitiveTypes
    }
}
// Gets the sensitivityLabel property value. 
func (m *EvaluateLabelJobResult) GetSensitivityLabel()(*MatchingLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabel
    }
}
// The deserialization information for the current model
func (m *EvaluateLabelJobResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["responsiblePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResponsiblePolicy() })
        if err != nil {
            return err
        }
        m.SetResponsiblePolicy(val.(*ResponsiblePolicy))
        return nil
    }
    res["responsibleSensitiveTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResponsibleSensitiveType() })
        if err != nil {
            return err
        }
        res := make([]ResponsibleSensitiveType, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResponsibleSensitiveType))
        }
        m.SetResponsibleSensitiveTypes(res)
        return nil
    }
    res["sensitivityLabel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMatchingLabel() })
        if err != nil {
            return err
        }
        m.SetSensitivityLabel(val.(*MatchingLabel))
        return nil
    }
    return res
}
func (m *EvaluateLabelJobResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EvaluateLabelJobResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("responsiblePolicy", m.GetResponsiblePolicy())
        if err != nil {
            return err
        }
    }
    {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EvaluateLabelJobResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the responsiblePolicy property value. 
// Parameters:
//  - value : Value to set for the responsiblePolicy property.
func (m *EvaluateLabelJobResult) SetResponsiblePolicy(value *ResponsiblePolicy)() {
    m.responsiblePolicy = value
}
// Sets the responsibleSensitiveTypes property value. 
// Parameters:
//  - value : Value to set for the responsibleSensitiveTypes property.
func (m *EvaluateLabelJobResult) SetResponsibleSensitiveTypes(value []ResponsibleSensitiveType)() {
    m.responsibleSensitiveTypes = value
}
// Sets the sensitivityLabel property value. 
// Parameters:
//  - value : Value to set for the sensitivityLabel property.
func (m *EvaluateLabelJobResult) SetSensitivityLabel(value *MatchingLabel)() {
    m.sensitivityLabel = value
}
