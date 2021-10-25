package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RegistrationEnforcement struct {
    additionalData map[string]interface{};
    authenticationMethodsRegistrationCampaign *AuthenticationMethodsRegistrationCampaign;
}
func NewRegistrationEnforcement()(*RegistrationEnforcement) {
    m := &RegistrationEnforcement{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RegistrationEnforcement) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RegistrationEnforcement) GetAuthenticationMethodsRegistrationCampaign()(*AuthenticationMethodsRegistrationCampaign) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodsRegistrationCampaign
    }
}
func (m *RegistrationEnforcement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationMethodsRegistrationCampaign"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsRegistrationCampaign() })
        if err != nil {
            return err
        }
        m.SetAuthenticationMethodsRegistrationCampaign(val.(*AuthenticationMethodsRegistrationCampaign))
        return nil
    }
    return res
}
func (m *RegistrationEnforcement) IsNil()(bool) {
    return m == nil
}
func (m *RegistrationEnforcement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationMethodsRegistrationCampaign", m.GetAuthenticationMethodsRegistrationCampaign())
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
func (m *RegistrationEnforcement) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RegistrationEnforcement) SetAuthenticationMethodsRegistrationCampaign(value *AuthenticationMethodsRegistrationCampaign)() {
    m.authenticationMethodsRegistrationCampaign = value
}
