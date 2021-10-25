package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EvaluateLabelJobResultGroup struct {
    additionalData map[string]interface{};
    automatic *EvaluateLabelJobResult;
    recommended *EvaluateLabelJobResult;
}
func NewEvaluateLabelJobResultGroup()(*EvaluateLabelJobResultGroup) {
    m := &EvaluateLabelJobResultGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EvaluateLabelJobResultGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EvaluateLabelJobResultGroup) GetAutomatic()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.automatic
    }
}
func (m *EvaluateLabelJobResultGroup) GetRecommended()(*EvaluateLabelJobResult) {
    if m == nil {
        return nil
    } else {
        return m.recommended
    }
}
func (m *EvaluateLabelJobResultGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["automatic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateLabelJobResult() })
        if err != nil {
            return err
        }
        m.SetAutomatic(val.(*EvaluateLabelJobResult))
        return nil
    }
    res["recommended"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvaluateLabelJobResult() })
        if err != nil {
            return err
        }
        m.SetRecommended(val.(*EvaluateLabelJobResult))
        return nil
    }
    return res
}
func (m *EvaluateLabelJobResultGroup) IsNil()(bool) {
    return m == nil
}
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
func (m *EvaluateLabelJobResultGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EvaluateLabelJobResultGroup) SetAutomatic(value *EvaluateLabelJobResult)() {
    m.automatic = value
}
func (m *EvaluateLabelJobResultGroup) SetRecommended(value *EvaluateLabelJobResult)() {
    m.recommended = value
}
