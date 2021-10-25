package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterValidationResult struct {
    additionalData map[string]interface{};
    isValidRule *bool;
}
func NewAssignmentFilterValidationResult()(*AssignmentFilterValidationResult) {
    m := &AssignmentFilterValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentFilterValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentFilterValidationResult) GetIsValidRule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidRule
    }
}
func (m *AssignmentFilterValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isValidRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValidRule(val)
        return nil
    }
    return res
}
func (m *AssignmentFilterValidationResult) IsNil()(bool) {
    return m == nil
}
func (m *AssignmentFilterValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValidRule", m.GetIsValidRule())
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
func (m *AssignmentFilterValidationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentFilterValidationResult) SetIsValidRule(value *bool)() {
    m.isValidRule = value
}
