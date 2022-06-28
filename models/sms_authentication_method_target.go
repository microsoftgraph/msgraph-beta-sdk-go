package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SmsAuthenticationMethodTarget 
type SmsAuthenticationMethodTarget struct {
    AuthenticationMethodTarget
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Determines if the users or groups can use this authentication method to sign in to Azure AD. The value is always true.
    isUsableForSignIn *bool
}
// NewSmsAuthenticationMethodTarget instantiates a new SmsAuthenticationMethodTarget and sets the default values.
func NewSmsAuthenticationMethodTarget()(*SmsAuthenticationMethodTarget) {
    m := &SmsAuthenticationMethodTarget{
        AuthenticationMethodTarget: *NewAuthenticationMethodTarget(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSmsAuthenticationMethodTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSmsAuthenticationMethodTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSmsAuthenticationMethodTarget(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SmsAuthenticationMethodTarget) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SmsAuthenticationMethodTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodTarget.GetFieldDeserializers()
    res["isUsableForSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsableForSignIn(val)
        }
        return nil
    }
    return res
}
// GetIsUsableForSignIn gets the isUsableForSignIn property value. Determines if the users or groups can use this authentication method to sign in to Azure AD. The value is always true.
func (m *SmsAuthenticationMethodTarget) GetIsUsableForSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUsableForSignIn
    }
}
// Serialize serializes information the current object
func (m *SmsAuthenticationMethodTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isUsableForSignIn", m.GetIsUsableForSignIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SmsAuthenticationMethodTarget) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsUsableForSignIn sets the isUsableForSignIn property value. Determines if the users or groups can use this authentication method to sign in to Azure AD. The value is always true.
func (m *SmsAuthenticationMethodTarget) SetIsUsableForSignIn(value *bool)() {
    if m != nil {
        m.isUsableForSignIn = value
    }
}
