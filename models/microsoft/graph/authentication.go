package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Authentication 
type Authentication struct {
    Entity
    // 
    emailMethods []EmailAuthenticationMethod;
    // 
    fido2Methods []Fido2AuthenticationMethod;
    // 
    methods []AuthenticationMethod;
    // 
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethod;
    // 
    operations []LongRunningOperation;
    // 
    passwordlessMicrosoftAuthenticatorMethods []PasswordlessMicrosoftAuthenticatorAuthenticationMethod;
    // 
    passwordMethods []PasswordAuthenticationMethod;
    // 
    phoneMethods []PhoneAuthenticationMethod;
    // 
    softwareOathMethods []SoftwareOathAuthenticationMethod;
    // 
    temporaryAccessPassMethods []TemporaryAccessPassAuthenticationMethod;
    // 
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethod;
}
// NewAuthentication instantiates a new authentication and sets the default values.
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
// GetEmailMethods gets the emailMethods property value. 
func (m *Authentication) GetEmailMethods()([]EmailAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.emailMethods
    }
}
// GetFido2Methods gets the fido2Methods property value. 
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.fido2Methods
    }
}
// GetMethods gets the methods property value. 
func (m *Authentication) GetMethods()([]AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.methods
    }
}
// GetMicrosoftAuthenticatorMethods gets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.microsoftAuthenticatorMethods
    }
}
// GetOperations gets the operations property value. 
func (m *Authentication) GetOperations()([]LongRunningOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPasswordlessMicrosoftAuthenticatorMethods gets the passwordlessMicrosoftAuthenticatorMethods property value. 
func (m *Authentication) GetPasswordlessMicrosoftAuthenticatorMethods()([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.passwordlessMicrosoftAuthenticatorMethods
    }
}
// GetPasswordMethods gets the passwordMethods property value. 
func (m *Authentication) GetPasswordMethods()([]PasswordAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.passwordMethods
    }
}
// GetPhoneMethods gets the phoneMethods property value. 
func (m *Authentication) GetPhoneMethods()([]PhoneAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.phoneMethods
    }
}
// GetSoftwareOathMethods gets the softwareOathMethods property value. 
func (m *Authentication) GetSoftwareOathMethods()([]SoftwareOathAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.softwareOathMethods
    }
}
// GetTemporaryAccessPassMethods gets the temporaryAccessPassMethods property value. 
func (m *Authentication) GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.temporaryAccessPassMethods
    }
}
// GetWindowsHelloForBusinessMethods gets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessMethods
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authentication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*EmailAuthenticationMethod))
            }
            m.SetEmailMethods(res)
        }
        return nil
    }
    res["fido2Methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFido2AuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*Fido2AuthenticationMethod))
            }
            m.SetFido2Methods(res)
        }
        return nil
    }
    res["methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationMethod))
            }
            m.SetMethods(res)
        }
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftAuthenticatorAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*MicrosoftAuthenticatorAuthenticationMethod))
            }
            m.SetMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLongRunningOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LongRunningOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*LongRunningOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["passwordlessMicrosoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*PasswordlessMicrosoftAuthenticatorAuthenticationMethod))
            }
            m.SetPasswordlessMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["passwordMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*PasswordAuthenticationMethod))
            }
            m.SetPasswordMethods(res)
        }
        return nil
    }
    res["phoneMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhoneAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhoneAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*PhoneAuthenticationMethod))
            }
            m.SetPhoneMethods(res)
        }
        return nil
    }
    res["softwareOathMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSoftwareOathAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SoftwareOathAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*SoftwareOathAuthenticationMethod))
            }
            m.SetSoftwareOathMethods(res)
        }
        return nil
    }
    res["temporaryAccessPassMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemporaryAccessPassAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemporaryAccessPassAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*TemporaryAccessPassAuthenticationMethod))
            }
            m.SetTemporaryAccessPassMethods(res)
        }
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsHelloForBusinessAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsHelloForBusinessAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsHelloForBusinessAuthenticationMethod))
            }
            m.SetWindowsHelloForBusinessMethods(res)
        }
        return nil
    }
    return res
}
func (m *Authentication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Authentication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmailMethods()))
        for i, v := range m.GetEmailMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("emailMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFido2Methods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFido2Methods()))
        for i, v := range m.GetFido2Methods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fido2Methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMethods()))
        for i, v := range m.GetMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetMicrosoftAuthenticatorMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordlessMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordlessMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetPasswordlessMicrosoftAuthenticatorMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("passwordlessMicrosoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordMethods()))
        for i, v := range m.GetPasswordMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("passwordMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhoneMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhoneMethods()))
        for i, v := range m.GetPhoneMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phoneMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSoftwareOathMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSoftwareOathMethods()))
        for i, v := range m.GetSoftwareOathMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("softwareOathMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemporaryAccessPassMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemporaryAccessPassMethods()))
        for i, v := range m.GetTemporaryAccessPassMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("temporaryAccessPassMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsHelloForBusinessMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsHelloForBusinessMethods()))
        for i, v := range m.GetWindowsHelloForBusinessMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsHelloForBusinessMethods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailMethods sets the emailMethods property value. 
func (m *Authentication) SetEmailMethods(value []EmailAuthenticationMethod)() {
    if m != nil {
        m.emailMethods = value
    }
}
// SetFido2Methods sets the fido2Methods property value. 
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethod)() {
    if m != nil {
        m.fido2Methods = value
    }
}
// SetMethods sets the methods property value. 
func (m *Authentication) SetMethods(value []AuthenticationMethod)() {
    if m != nil {
        m.methods = value
    }
}
// SetMicrosoftAuthenticatorMethods sets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethod)() {
    if m != nil {
        m.microsoftAuthenticatorMethods = value
    }
}
// SetOperations sets the operations property value. 
func (m *Authentication) SetOperations(value []LongRunningOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetPasswordlessMicrosoftAuthenticatorMethods sets the passwordlessMicrosoftAuthenticatorMethods property value. 
func (m *Authentication) SetPasswordlessMicrosoftAuthenticatorMethods(value []PasswordlessMicrosoftAuthenticatorAuthenticationMethod)() {
    if m != nil {
        m.passwordlessMicrosoftAuthenticatorMethods = value
    }
}
// SetPasswordMethods sets the passwordMethods property value. 
func (m *Authentication) SetPasswordMethods(value []PasswordAuthenticationMethod)() {
    if m != nil {
        m.passwordMethods = value
    }
}
// SetPhoneMethods sets the phoneMethods property value. 
func (m *Authentication) SetPhoneMethods(value []PhoneAuthenticationMethod)() {
    if m != nil {
        m.phoneMethods = value
    }
}
// SetSoftwareOathMethods sets the softwareOathMethods property value. 
func (m *Authentication) SetSoftwareOathMethods(value []SoftwareOathAuthenticationMethod)() {
    if m != nil {
        m.softwareOathMethods = value
    }
}
// SetTemporaryAccessPassMethods sets the temporaryAccessPassMethods property value. 
func (m *Authentication) SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethod)() {
    if m != nil {
        m.temporaryAccessPassMethods = value
    }
}
// SetWindowsHelloForBusinessMethods sets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethod)() {
    if m != nil {
        m.windowsHelloForBusinessMethods = value
    }
}
