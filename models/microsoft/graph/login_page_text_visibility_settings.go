package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new loginPageTextVisibilitySettings and sets the default values.
func NewLoginPageTextVisibilitySettings()(*LoginPageTextVisibilitySettings) {
    m := &LoginPageTextVisibilitySettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LoginPageTextVisibilitySettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hideCannotAccessYourAccount property value. 
func (m *LoginPageTextVisibilitySettings) GetHideCannotAccessYourAccount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideCannotAccessYourAccount
    }
}
// Gets the hideForgotMyPassword property value. 
func (m *LoginPageTextVisibilitySettings) GetHideForgotMyPassword()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideForgotMyPassword
    }
}
// Gets the hidePrivacyAndCookies property value. 
func (m *LoginPageTextVisibilitySettings) GetHidePrivacyAndCookies()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidePrivacyAndCookies
    }
}
// Gets the hideResetItNow property value. 
func (m *LoginPageTextVisibilitySettings) GetHideResetItNow()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideResetItNow
    }
}
// Gets the hideTermsOfUse property value. 
func (m *LoginPageTextVisibilitySettings) GetHideTermsOfUse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideTermsOfUse
    }
}
// The deserialization information for the current model
func (m *LoginPageTextVisibilitySettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hideCannotAccessYourAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideCannotAccessYourAccount(val)
        return nil
    }
    res["hideForgotMyPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideForgotMyPassword(val)
        return nil
    }
    res["hidePrivacyAndCookies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidePrivacyAndCookies(val)
        return nil
    }
    res["hideResetItNow"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideResetItNow(val)
        return nil
    }
    res["hideTermsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideTermsOfUse(val)
        return nil
    }
    return res
}
func (m *LoginPageTextVisibilitySettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *LoginPageTextVisibilitySettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hideCannotAccessYourAccount property value. 
// Parameters:
//  - value : Value to set for the hideCannotAccessYourAccount property.
func (m *LoginPageTextVisibilitySettings) SetHideCannotAccessYourAccount(value *bool)() {
    m.hideCannotAccessYourAccount = value
}
// Sets the hideForgotMyPassword property value. 
// Parameters:
//  - value : Value to set for the hideForgotMyPassword property.
func (m *LoginPageTextVisibilitySettings) SetHideForgotMyPassword(value *bool)() {
    m.hideForgotMyPassword = value
}
// Sets the hidePrivacyAndCookies property value. 
// Parameters:
//  - value : Value to set for the hidePrivacyAndCookies property.
func (m *LoginPageTextVisibilitySettings) SetHidePrivacyAndCookies(value *bool)() {
    m.hidePrivacyAndCookies = value
}
// Sets the hideResetItNow property value. 
// Parameters:
//  - value : Value to set for the hideResetItNow property.
func (m *LoginPageTextVisibilitySettings) SetHideResetItNow(value *bool)() {
    m.hideResetItNow = value
}
// Sets the hideTermsOfUse property value. 
// Parameters:
//  - value : Value to set for the hideTermsOfUse property.
func (m *LoginPageTextVisibilitySettings) SetHideTermsOfUse(value *bool)() {
    m.hideTermsOfUse = value
}
