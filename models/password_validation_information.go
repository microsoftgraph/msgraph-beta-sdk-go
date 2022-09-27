package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordValidationInformation 
type PasswordValidationInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
    isValid *bool
    // The OdataType property
    odataType *string
    // The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
    validationResults []ValidationResultable
}
// NewPasswordValidationInformation instantiates a new passwordValidationInformation and sets the default values.
func NewPasswordValidationInformation()(*PasswordValidationInformation) {
    m := &PasswordValidationInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.passwordValidationInformation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePasswordValidationInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordValidationInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordValidationInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordValidationInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordValidationInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isValid"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsValid)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["validationResults"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateValidationResultFromDiscriminatorValue , m.SetValidationResults)
    return res
}
// GetIsValid gets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetIsValid()(*bool) {
    return m.isValid
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordValidationInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetValidationResults gets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetValidationResults()([]ValidationResultable) {
    return m.validationResults
}
// Serialize serializes information the current object
func (m *PasswordValidationInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValid", m.GetIsValid())
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
    if m.GetValidationResults() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValidationResults())
        err := writer.WriteCollectionOfObjectValues("validationResults", cast)
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
func (m *PasswordValidationInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsValid sets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
func (m *PasswordValidationInformation) SetIsValid(value *bool)() {
    m.isValid = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordValidationInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValidationResults sets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
func (m *PasswordValidationInformation) SetValidationResults(value []ValidationResultable)() {
    m.validationResults = value
}
