package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LoginPageTextVisibilitySettings 
type LoginPageTextVisibilitySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hideCannotAccessYourAccount *bool;
    // 
    hideForgotMyPassword *bool;
    // 
    hidePrivacyAndCookies *bool;
    // 
    hideResetItNow *bool;
    // 
    hideTermsOfUse *bool;
}
// NewLoginPageTextVisibilitySettings instantiates a new loginPageTextVisibilitySettings and sets the default values.
func NewLoginPageTextVisibilitySettings()(*LoginPageTextVisibilitySettings) {
    m := &LoginPageTextVisibilitySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageTextVisibilitySettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHideCannotAccessYourAccount gets the hideCannotAccessYourAccount property value. 
func (m *LoginPageTextVisibilitySettings) GetHideCannotAccessYourAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideCannotAccessYourAccount
    }
}
// GetHideForgotMyPassword gets the hideForgotMyPassword property value. 
func (m *LoginPageTextVisibilitySettings) GetHideForgotMyPassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideForgotMyPassword
    }
}
// GetHidePrivacyAndCookies gets the hidePrivacyAndCookies property value. 
func (m *LoginPageTextVisibilitySettings) GetHidePrivacyAndCookies()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacyAndCookies
    }
}
// GetHideResetItNow gets the hideResetItNow property value. 
func (m *LoginPageTextVisibilitySettings) GetHideResetItNow()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideResetItNow
    }
}
// GetHideTermsOfUse gets the hideTermsOfUse property value. 
func (m *LoginPageTextVisibilitySettings) GetHideTermsOfUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideTermsOfUse
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LoginPageTextVisibilitySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hideCannotAccessYourAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideCannotAccessYourAccount(val)
        }
        return nil
    }
    res["hideForgotMyPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideForgotMyPassword(val)
        }
        return nil
    }
    res["hidePrivacyAndCookies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePrivacyAndCookies(val)
        }
        return nil
    }
    res["hideResetItNow"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideResetItNow(val)
        }
        return nil
    }
    res["hideTermsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *LoginPageTextVisibilitySettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LoginPageTextVisibilitySettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageTextVisibilitySettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHideCannotAccessYourAccount sets the hideCannotAccessYourAccount property value. 
func (m *LoginPageTextVisibilitySettings) SetHideCannotAccessYourAccount(value *bool)() {
    m.hideCannotAccessYourAccount = value
}
// SetHideForgotMyPassword sets the hideForgotMyPassword property value. 
func (m *LoginPageTextVisibilitySettings) SetHideForgotMyPassword(value *bool)() {
    m.hideForgotMyPassword = value
}
// SetHidePrivacyAndCookies sets the hidePrivacyAndCookies property value. 
func (m *LoginPageTextVisibilitySettings) SetHidePrivacyAndCookies(value *bool)() {
    m.hidePrivacyAndCookies = value
}
// SetHideResetItNow sets the hideResetItNow property value. 
func (m *LoginPageTextVisibilitySettings) SetHideResetItNow(value *bool)() {
    m.hideResetItNow = value
}
// SetHideTermsOfUse sets the hideTermsOfUse property value. 
func (m *LoginPageTextVisibilitySettings) SetHideTermsOfUse(value *bool)() {
    m.hideTermsOfUse = value
}
