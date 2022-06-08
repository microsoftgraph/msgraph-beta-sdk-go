package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoginPageTextVisibilitySettings 
type LoginPageTextVisibilitySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Option to hide the self-service password reset (SSPR) hyperlinks such as 'Can't access your account?', 'Forgot my password' and 'Reset it now' on the sign-in form.
    hideAccountResetCredentials *bool
    // Option to hide the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in form.
    hideCannotAccessYourAccount *bool
    // Option to hide the self-service password reset (SSPR) 'Forgot my password' hyperlink on the sign-in form.
    hideForgotMyPassword *bool
    // Option to hide the 'Privacy & Cookies' hyperlink in the footer.
    hidePrivacyAndCookies *bool
    // Option to hide the self-service password reset (SSPR) 'reset it now' hyperlink on the sign-in form.
    hideResetItNow *bool
    // Option to hide the 'Terms of Use' hyperlink in the footer.
    hideTermsOfUse *bool
}
// NewLoginPageTextVisibilitySettings instantiates a new loginPageTextVisibilitySettings and sets the default values.
func NewLoginPageTextVisibilitySettings()(*LoginPageTextVisibilitySettings) {
    m := &LoginPageTextVisibilitySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLoginPageTextVisibilitySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLoginPageTextVisibilitySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLoginPageTextVisibilitySettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageTextVisibilitySettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoginPageTextVisibilitySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hideAccountResetCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideAccountResetCredentials(val)
        }
        return nil
    }
    res["hideCannotAccessYourAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideCannotAccessYourAccount(val)
        }
        return nil
    }
    res["hideForgotMyPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideForgotMyPassword(val)
        }
        return nil
    }
    res["hidePrivacyAndCookies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePrivacyAndCookies(val)
        }
        return nil
    }
    res["hideResetItNow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideResetItNow(val)
        }
        return nil
    }
    res["hideTermsOfUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideTermsOfUse(val)
        }
        return nil
    }
    return res
}
// GetHideAccountResetCredentials gets the hideAccountResetCredentials property value. Option to hide the self-service password reset (SSPR) hyperlinks such as 'Can't access your account?', 'Forgot my password' and 'Reset it now' on the sign-in form.
func (m *LoginPageTextVisibilitySettings) GetHideAccountResetCredentials()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideAccountResetCredentials
    }
}
// GetHideCannotAccessYourAccount gets the hideCannotAccessYourAccount property value. Option to hide the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) GetHideCannotAccessYourAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideCannotAccessYourAccount
    }
}
// GetHideForgotMyPassword gets the hideForgotMyPassword property value. Option to hide the self-service password reset (SSPR) 'Forgot my password' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) GetHideForgotMyPassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideForgotMyPassword
    }
}
// GetHidePrivacyAndCookies gets the hidePrivacyAndCookies property value. Option to hide the 'Privacy & Cookies' hyperlink in the footer.
func (m *LoginPageTextVisibilitySettings) GetHidePrivacyAndCookies()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacyAndCookies
    }
}
// GetHideResetItNow gets the hideResetItNow property value. Option to hide the self-service password reset (SSPR) 'reset it now' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) GetHideResetItNow()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideResetItNow
    }
}
// GetHideTermsOfUse gets the hideTermsOfUse property value. Option to hide the 'Terms of Use' hyperlink in the footer.
func (m *LoginPageTextVisibilitySettings) GetHideTermsOfUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideTermsOfUse
    }
}
// Serialize serializes information the current object
func (m *LoginPageTextVisibilitySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hideAccountResetCredentials", m.GetHideAccountResetCredentials())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideCannotAccessYourAccount", m.GetHideCannotAccessYourAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideForgotMyPassword", m.GetHideForgotMyPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidePrivacyAndCookies", m.GetHidePrivacyAndCookies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideResetItNow", m.GetHideResetItNow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideTermsOfUse", m.GetHideTermsOfUse())
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
func (m *LoginPageTextVisibilitySettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHideAccountResetCredentials sets the hideAccountResetCredentials property value. Option to hide the self-service password reset (SSPR) hyperlinks such as 'Can't access your account?', 'Forgot my password' and 'Reset it now' on the sign-in form.
func (m *LoginPageTextVisibilitySettings) SetHideAccountResetCredentials(value *bool)() {
    if m != nil {
        m.hideAccountResetCredentials = value
    }
}
// SetHideCannotAccessYourAccount sets the hideCannotAccessYourAccount property value. Option to hide the self-service password reset (SSPR) 'Can't access your account?' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) SetHideCannotAccessYourAccount(value *bool)() {
    if m != nil {
        m.hideCannotAccessYourAccount = value
    }
}
// SetHideForgotMyPassword sets the hideForgotMyPassword property value. Option to hide the self-service password reset (SSPR) 'Forgot my password' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) SetHideForgotMyPassword(value *bool)() {
    if m != nil {
        m.hideForgotMyPassword = value
    }
}
// SetHidePrivacyAndCookies sets the hidePrivacyAndCookies property value. Option to hide the 'Privacy & Cookies' hyperlink in the footer.
func (m *LoginPageTextVisibilitySettings) SetHidePrivacyAndCookies(value *bool)() {
    if m != nil {
        m.hidePrivacyAndCookies = value
    }
}
// SetHideResetItNow sets the hideResetItNow property value. Option to hide the self-service password reset (SSPR) 'reset it now' hyperlink on the sign-in form.
func (m *LoginPageTextVisibilitySettings) SetHideResetItNow(value *bool)() {
    if m != nil {
        m.hideResetItNow = value
    }
}
// SetHideTermsOfUse sets the hideTermsOfUse property value. Option to hide the 'Terms of Use' hyperlink in the footer.
func (m *LoginPageTextVisibilitySettings) SetHideTermsOfUse(value *bool)() {
    if m != nil {
        m.hideTermsOfUse = value
    }
}
