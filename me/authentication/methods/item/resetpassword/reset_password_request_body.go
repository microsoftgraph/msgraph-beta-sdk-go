package resetpassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResetPasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    newPassword *string;
    // 
    requireChangeOnNextSignIn *bool;
}
// Instantiates a new resetPasswordRequestBody and sets the default values.
func NewResetPasswordRequestBody()(*ResetPasswordRequestBody) {
    m := &ResetPasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResetPasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the newPassword property value. 
func (m *ResetPasswordRequestBody) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
// Gets the requireChangeOnNextSignIn property value. 
func (m *ResetPasswordRequestBody) GetRequireChangeOnNextSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requireChangeOnNextSignIn
    }
}
// The deserialization information for the current model
func (m *ResetPasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    res["requireChangeOnNextSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireChangeOnNextSignIn(val)
        }
        return nil
    }
    return res
}
func (m *ResetPasswordRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResetPasswordRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("requireChangeOnNextSignIn", m.GetRequireChangeOnNextSignIn())
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
func (m *ResetPasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the newPassword property value. 
// Parameters:
//  - value : Value to set for the newPassword property.
func (m *ResetPasswordRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
// Sets the requireChangeOnNextSignIn property value. 
// Parameters:
//  - value : Value to set for the requireChangeOnNextSignIn property.
func (m *ResetPasswordRequestBody) SetRequireChangeOnNextSignIn(value *bool)() {
    m.requireChangeOnNextSignIn = value
}
