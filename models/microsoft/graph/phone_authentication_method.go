package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PhoneAuthenticationMethod 
type PhoneAuthenticationMethod struct {
    AuthenticationMethod
    // The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
    phoneNumber *string;
    // The type of this phone. Possible values are: mobile, alternateMobile, or office.
    phoneType *AuthenticationPhoneType;
    // Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
    smsSignInState *AuthenticationMethodSignInState;
}
// NewPhoneAuthenticationMethod instantiates a new phoneAuthenticationMethod and sets the default values.
func NewPhoneAuthenticationMethod()(*PhoneAuthenticationMethod) {
    m := &PhoneAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// GetPhoneNumber gets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
func (m *PhoneAuthenticationMethod) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetPhoneType gets the phoneType property value. The type of this phone. Possible values are: mobile, alternateMobile, or office.
func (m *PhoneAuthenticationMethod) GetPhoneType()(*AuthenticationPhoneType) {
    if m == nil {
        return nil
    } else {
        return m.phoneType
    }
}
// GetSmsSignInState gets the smsSignInState property value. Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
func (m *PhoneAuthenticationMethod) GetSmsSignInState()(*AuthenticationMethodSignInState) {
    if m == nil {
        return nil
    } else {
        return m.smsSignInState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhoneAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["phoneType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationPhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneType(val.(*AuthenticationPhoneType))
        }
        return nil
    }
    res["smsSignInState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodSignInState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsSignInState(val.(*AuthenticationMethodSignInState))
        }
        return nil
    }
    return res
}
func (m *PhoneAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PhoneAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPhoneType() != nil {
        cast := (*m.GetPhoneType()).String()
        err = writer.WriteStringValue("phoneType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSmsSignInState() != nil {
        cast := (*m.GetSmsSignInState()).String()
        err = writer.WriteStringValue("smsSignInState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPhoneNumber sets the phoneNumber property value. The phone number to text or call for authentication. Phone numbers use the format '+<country code> <number>x<extension>', with extension optional. For example, +1 5555551234 or +1 5555551234x123 are valid. Numbers are rejected when creating/updating if they do not match the required format.
func (m *PhoneAuthenticationMethod) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetPhoneType sets the phoneType property value. The type of this phone. Possible values are: mobile, alternateMobile, or office.
func (m *PhoneAuthenticationMethod) SetPhoneType(value *AuthenticationPhoneType)() {
    if m != nil {
        m.phoneType = value
    }
}
// SetSmsSignInState sets the smsSignInState property value. Whether a phone is ready to be used for SMS sign-in or not. Possible values are: notSupported, notAllowedByPolicy, notEnabled, phoneNumberNotUnique, ready, or notConfigured, unknownFutureValue.
func (m *PhoneAuthenticationMethod) SetSmsSignInState(value *AuthenticationMethodSignInState)() {
    if m != nil {
        m.smsSignInState = value
    }
}
