package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterValidationResult 
type AssignmentFilterValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicator to valid or invalid rule.
    isValidRule *bool;
}
// NewAssignmentFilterValidationResult instantiates a new assignmentFilterValidationResult and sets the default values.
func NewAssignmentFilterValidationResult()(*AssignmentFilterValidationResult) {
    m := &AssignmentFilterValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsValidRule gets the isValidRule property value. Indicator to valid or invalid rule.
func (m *AssignmentFilterValidationResult) GetIsValidRule()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidRule
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isValidRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidRule(val)
        }
        return nil
    }
    return res
}
func (m *AssignmentFilterValidationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterValidationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsValidRule sets the isValidRule property value. Indicator to valid or invalid rule.
func (m *AssignmentFilterValidationResult) SetIsValidRule(value *bool)() {
    if m != nil {
        m.isValidRule = value
    }
}
