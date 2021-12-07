package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ValidationResult 
type ValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The string containing the reason for why the rule passed or not. Read-only. Not nullable.
    message *string;
    // The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
    ruleName *string;
    // Whether the password passed or failed the validation rule. Read-only. Not nullable.
    validationPassed *bool;
}
// NewValidationResult instantiates a new validationResult and sets the default values.
func NewValidationResult()(*ValidationResult) {
    m := &ValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMessage gets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
func (m *ValidationResult) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetRuleName gets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
func (m *ValidationResult) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
// GetValidationPassed gets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
func (m *ValidationResult) GetValidationPassed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.validationPassed
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    res["validationPassed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidationPassed(val)
        }
        return nil
    }
    return res
}
func (m *ValidationResult) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ValidationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleName", m.GetRuleName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("validationPassed", m.GetValidationPassed())
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
func (m *ValidationResult) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessage sets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
func (m *ValidationResult) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetRuleName sets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
func (m *ValidationResult) SetRuleName(value *string)() {
    if m != nil {
        m.ruleName = value
    }
}
// SetValidationPassed sets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
func (m *ValidationResult) SetValidationPassed(value *bool)() {
    if m != nil {
        m.validationPassed = value
    }
}
