package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PasswordValidationInformation provides operations to call the validatePassword method.
type PasswordValidationInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
    isValid *bool;
    // The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
    validationResults []ValidationResultable;
}
// NewPasswordValidationInformation instantiates a new passwordValidationInformation and sets the default values.
func NewPasswordValidationInformation()(*PasswordValidationInformation) {
    m := &PasswordValidationInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePasswordValidationInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordValidationInformationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPasswordValidationInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordValidationInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordValidationInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isValid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["validationResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateValidationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ValidationResultable, len(val))
            for i, v := range val {
                res[i] = v.(ValidationResultable)
            }
            m.SetValidationResults(res)
        }
        return nil
    }
    return res
}
// GetIsValid gets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetIsValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValid
    }
}
// GetValidationResults gets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetValidationResults()([]ValidationResultable) {
    if m == nil {
        return nil
    } else {
        return m.validationResults
    }
}
func (m *PasswordValidationInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PasswordValidationInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    if m.GetValidationResults() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValidationResults()))
        for i, v := range m.GetValidationResults() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetIsValid sets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
func (m *PasswordValidationInformation) SetIsValid(value *bool)() {
    if m != nil {
        m.isValid = value
    }
}
// SetValidationResults sets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
func (m *PasswordValidationInformation) SetValidationResults(value []ValidationResultable)() {
    if m != nil {
        m.validationResults = value
    }
}
