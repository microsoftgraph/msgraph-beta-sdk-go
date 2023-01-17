package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterValidationResult represents result of Validation API.
type AssignmentFilterValidationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicator to valid or invalid rule.
    isValidRule *bool
    // The OdataType property
    odataType *string
}
// NewAssignmentFilterValidationResult instantiates a new assignmentFilterValidationResult and sets the default values.
func NewAssignmentFilterValidationResult()(*AssignmentFilterValidationResult) {
    m := &AssignmentFilterValidationResult{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAssignmentFilterValidationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterValidationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterValidationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterValidationResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterValidationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isValidRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidRule(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIsValidRule gets the isValidRule property value. Indicator to valid or invalid rule.
func (m *AssignmentFilterValidationResult) GetIsValidRule()(*bool) {
    return m.isValidRule
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterValidationResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AssignmentFilterValidationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValidRule", m.GetIsValidRule())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterValidationResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsValidRule sets the isValidRule property value. Indicator to valid or invalid rule.
func (m *AssignmentFilterValidationResult) SetIsValidRule(value *bool)() {
    m.isValidRule = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterValidationResult) SetOdataType(value *string)() {
    m.odataType = value
}
