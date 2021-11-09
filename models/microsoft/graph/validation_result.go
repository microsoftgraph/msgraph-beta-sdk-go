package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new validationResult and sets the default values.
func NewValidationResult()(*ValidationResult) {
    m := &ValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
func (m *ValidationResult) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
func (m *ValidationResult) GetRuleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ruleName
    }
}
// Gets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
func (m *ValidationResult) GetValidationPassed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.validationPassed
    }
}
// The deserialization information for the current model
func (m *ValidationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["ruleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRuleName(val)
        return nil
    }
    res["validationPassed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetValidationPassed(val)
        return nil
    }
    return res
}
func (m *ValidationResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ValidationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
// Parameters:
//  - value : Value to set for the message property.
func (m *ValidationResult) SetMessage(value *string)() {
    m.message = value
}
// Sets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
// Parameters:
//  - value : Value to set for the ruleName property.
func (m *ValidationResult) SetRuleName(value *string)() {
    m.ruleName = value
}
// Sets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
// Parameters:
//  - value : Value to set for the validationPassed property.
func (m *ValidationResult) SetValidationPassed(value *bool)() {
    m.validationPassed = value
}
