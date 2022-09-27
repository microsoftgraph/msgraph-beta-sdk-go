package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidationResult 
type ValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The string containing the reason for why the rule passed or not. Read-only. Not nullable.
    message *string
    // The OdataType property
    odataType *string
    // The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
    ruleName *string
    // Whether the password passed or failed the validation rule. Read-only. Not nullable.
    validationPassed *bool
}
// NewValidationResult instantiates a new validationResult and sets the default values.
func NewValidationResult()(*ValidationResult) {
    m := &ValidationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.validationResult";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateValidationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidationResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMessage)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["ruleName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRuleName)
    res["validationPassed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetValidationPassed)
    return res
}
// GetMessage gets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
func (m *ValidationResult) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ValidationResult) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleName gets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
func (m *ValidationResult) GetRuleName()(*string) {
    return m.ruleName
}
// GetValidationPassed gets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
func (m *ValidationResult) GetValidationPassed()(*bool) {
    return m.validationPassed
}
// Serialize serializes information the current object
func (m *ValidationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetMessage sets the message property value. The string containing the reason for why the rule passed or not. Read-only. Not nullable.
func (m *ValidationResult) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ValidationResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleName sets the ruleName property value. The string containing the name of the password validation rule that the action was validated against. Read-only. Not nullable.
func (m *ValidationResult) SetRuleName(value *string)() {
    m.ruleName = value
}
// SetValidationPassed sets the validationPassed property value. Whether the password passed or failed the validation rule. Read-only. Not nullable.
func (m *ValidationResult) SetValidationPassed(value *bool)() {
    m.validationPassed = value
}
