package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// B2cAuthenticationMethodsPolicy 
type B2cAuthenticationMethodsPolicy struct {
    Entity
    // The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
    isEmailPasswordAuthenticationEnabled *bool;
    // The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
    isPhoneOneTimePasswordAuthenticationEnabled *bool;
    // The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
    isUserNameAuthenticationEnabled *bool;
}
// NewB2cAuthenticationMethodsPolicy instantiates a new b2cAuthenticationMethodsPolicy and sets the default values.
func NewB2cAuthenticationMethodsPolicy()(*B2cAuthenticationMethodsPolicy) {
    m := &B2cAuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetIsEmailPasswordAuthenticationEnabled gets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsEmailPasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEmailPasswordAuthenticationEnabled
    }
}
// GetIsPhoneOneTimePasswordAuthenticationEnabled gets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsPhoneOneTimePasswordAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPhoneOneTimePasswordAuthenticationEnabled
    }
}
// GetIsUserNameAuthenticationEnabled gets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) GetIsUserNameAuthenticationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isUserNameAuthenticationEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cAuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isEmailPasswordAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEmailPasswordAuthenticationEnabled(val)
        }
        return nil
    }
    res["isPhoneOneTimePasswordAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPhoneOneTimePasswordAuthenticationEnabled(val)
        }
        return nil
    }
    res["isUserNameAuthenticationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUserNameAuthenticationEnabled(val)
        }
        return nil
    }
    return res
}
func (m *B2cAuthenticationMethodsPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetIsEmailPasswordAuthenticationEnabled sets the isEmailPasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using email if the email and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsEmailPasswordAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isEmailPasswordAuthenticationEnabled = value
    }
}
// SetIsPhoneOneTimePasswordAuthenticationEnabled sets the isPhoneOneTimePasswordAuthenticationEnabled property value. The tenant admin can configure local accounts using phone number if the phone number and one-time password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsPhoneOneTimePasswordAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isPhoneOneTimePasswordAuthenticationEnabled = value
    }
}
// SetIsUserNameAuthenticationEnabled sets the isUserNameAuthenticationEnabled property value. The tenant admin can configure local accounts using username if the username and password authentication method is enabled.
func (m *B2cAuthenticationMethodsPolicy) SetIsUserNameAuthenticationEnabled(value *bool)() {
    if m != nil {
        m.isUserNameAuthenticationEnabled = value
    }
}
