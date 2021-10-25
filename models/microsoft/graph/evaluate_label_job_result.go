package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EvaluateLabelJobResult struct {
    additionalData map[string]interface{};
    responsiblePolicy *ResponsiblePolicy;
    responsibleSensitiveTypes []ResponsibleSensitiveType;
    sensitivityLabel *MatchingLabel;
}
func NewEvaluateLabelJobResult()(*EvaluateLabelJobResult) {
    m := &EvaluateLabelJobResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EvaluateLabelJobResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EvaluateLabelJobResult) GetResponsiblePolicy()(*ResponsiblePolicy) {
    if m == nil {
        return nil
    } else {
        return m.responsiblePolicy
    }
}
func (m *EvaluateLabelJobResult) GetResponsibleSensitiveTypes()([]ResponsibleSensitiveType) {
    if m == nil {
        return nil
    } else {
        return m.responsibleSensitiveTypes
    }
}
func (m *EvaluateLabelJobResult) GetSensitivityLabel()(*MatchingLabel) {
    if m == nil {
        return nil
    } else {
        return m.sensitivityLabel
    }
}
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
func (m *EvaluateLabelJobResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EvaluateLabelJobResult) SetResponsiblePolicy(value *ResponsiblePolicy)() {
    m.responsiblePolicy = value
}
func (m *EvaluateLabelJobResult) SetResponsibleSensitiveTypes(value []ResponsibleSensitiveType)() {
    m.responsibleSensitiveTypes = value
}
func (m *EvaluateLabelJobResult) SetSensitivityLabel(value *MatchingLabel)() {
    m.sensitivityLabel = value
}
