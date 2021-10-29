package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type B2cAuthenticationMethodsPolicy struct {
    Entity
    // The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
    isEmailPasswordAuthenticationEnabled *bool;
    // The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
    isPhoneOneTimePasswordAuthenticationEnabled *bool;
    // The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
    isUserNameAuthenticationEnabled *bool;
}
// Instantiates a new b2cAuthenticationMethodsPolicy and sets the default values.
func NewB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    m := &B2cAuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsEmailPasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEmailPasswordAuthenticationEnabled
    }
}
// Gets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPhoneOneTimePasswordAuthenticationEnabled
    }
}
// Gets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsUserNameAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUserNameAuthenticationEnabled
    }
}
// The deserialization information for the current model
func (m *B2cAuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEmailPasswordAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEmailPasswordAuthenticationEnabled(val)
        return nil
    }
    res["isPhoneOneTimePasswordAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPhoneOneTimePasswordAuthenticationEnabled(val)
        return nil
    }
    res["isUserNameAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsUserNameAuthenticationEnabled(val)
        return nil
    }
    return res
}
func (m *B2cAuthenticationMethodsPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *B2cAuthenticationMethodsPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEmailPasswordAuthenticationEnabled", m.GetIsEmailPasswordAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPhoneOneTimePasswordAuthenticationEnabled", m.GetIsPhoneOneTimePasswordAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isUserNameAuthenticationEnabled", m.GetIsUserNameAuthenticationEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
// Parameters:
//  - value : Value to set for the isEmailPasswordAuthenticationEnabled property.
func (m *B2cAuthenticationMethodsPolicy) SetIsEmailPasswordAuthenticationEnabled(value *bool)() {
    m.isEmailPasswordAuthenticationEnabled = value
}
// Sets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
// Parameters:
//  - value : Value to set for the isPhoneOneTimePasswordAuthenticationEnabled property.
func (m *B2cAuthenticationMethodsPolicy) SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)() {
    m.isPhoneOneTimePasswordAuthenticationEnabled = value
}
// Sets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
// Parameters:
//  - value : Value to set for the isUserNameAuthenticationEnabled property.
func (m *B2cAuthenticationMethodsPolicy) SetIsUserNameAuthenticationEnabled(value *bool)() {
    m.isUserNameAuthenticationEnabled = value
}
