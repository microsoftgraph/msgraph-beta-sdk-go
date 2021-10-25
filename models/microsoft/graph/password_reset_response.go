package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PasswordResetResponse struct {
    additionalData map[string]interface{};
    newPassword *string;
}
func NewPasswordResetResponse()(*PasswordResetResponse) {
    m := &PasswordResetResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PasswordResetResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PasswordResetResponse) GetNewPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newPassword
    }
}
func (m *PasswordResetResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNewPassword(val)
        return nil
    }
    return res
}
func (m *PasswordResetResponse) IsNil()(bool) {
    return m == nil
}
func (m *PasswordResetResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
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
func (m *PasswordResetResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PasswordResetResponse) SetNewPassword(value *string)() {
    m.newPassword = value
}
