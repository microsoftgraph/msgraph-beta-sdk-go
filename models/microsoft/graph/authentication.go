package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Authentication struct {
    Entity
    emailMethods []EmailAuthenticationMethod;
    fido2Methods []Fido2AuthenticationMethod;
    methods []AuthenticationMethod;
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethod;
    operations []LongRunningOperation;
    passwordlessMicrosoftAuthenticatorMethods []PasswordlessMicrosoftAuthenticatorAuthenticationMethod;
    passwordMethods []PasswordAuthenticationMethod;
    phoneMethods []PhoneAuthenticationMethod;
    temporaryAccessPassMethods []TemporaryAccessPassAuthenticationMethod;
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethod;
}
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Authentication) GetEmailMethods()([]EmailAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.emailMethods
    }
}
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.fido2Methods
    }
}
func (m *Authentication) GetMethods()([]AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.methods
    }
}
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.microsoftAuthenticatorMethods
    }
}
func (m *Authentication) GetOperations()([]LongRunningOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Authentication) GetPasswordlessMicrosoftAuthenticatorMethods()([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.passwordlessMicrosoftAuthenticatorMethods
    }
}
func (m *Authentication) GetPasswordMethods()([]PasswordAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.passwordMethods
    }
}
func (m *Authentication) GetPhoneMethods()([]PhoneAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.phoneMethods
    }
}
func (m *Authentication) GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.temporaryAccessPassMethods
    }
}
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessMethods
    }
}
func (m *Authentication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]EmailAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*EmailAuthenticationMethod))
        }
        m.SetEmailMethods(res)
        return nil
    }
    res["fido2Methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFido2AuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]Fido2AuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*Fido2AuthenticationMethod))
        }
        m.SetFido2Methods(res)
        return nil
    }
    res["methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]AuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthenticationMethod))
        }
        m.SetMethods(res)
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftAuthenticatorAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftAuthenticatorAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftAuthenticatorAuthenticationMethod))
        }
        m.SetMicrosoftAuthenticatorMethods(res)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLongRunningOperation() })
        if err != nil {
            return err
        }
        res := make([]LongRunningOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*LongRunningOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["passwordlessMicrosoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*PasswordlessMicrosoftAuthenticatorAuthenticationMethod))
        }
        m.SetPasswordlessMicrosoftAuthenticatorMethods(res)
        return nil
    }
    res["passwordMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]PasswordAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*PasswordAuthenticationMethod))
        }
        m.SetPasswordMethods(res)
        return nil
    }
    res["phoneMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhoneAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]PhoneAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*PhoneAuthenticationMethod))
        }
        m.SetPhoneMethods(res)
        return nil
    }
    res["temporaryAccessPassMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTemporaryAccessPassAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]TemporaryAccessPassAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*TemporaryAccessPassAuthenticationMethod))
        }
        m.SetTemporaryAccessPassMethods(res)
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsHelloForBusinessAuthenticationMethod() })
        if err != nil {
            return err
        }
        res := make([]WindowsHelloForBusinessAuthenticationMethod, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsHelloForBusinessAuthenticationMethod))
        }
        m.SetWindowsHelloForBusinessMethods(res)
        return nil
    }
    return res
}
func (m *Authentication) IsNil()(bool) {
    return m == nil
}
func (m *Authentication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
func (m *Authentication) SetEmailMethods(value []EmailAuthenticationMethod)() {
    m.emailMethods = value
}
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethod)() {
    m.fido2Methods = value
}
func (m *Authentication) SetMethods(value []AuthenticationMethod)() {
    m.methods = value
}
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethod)() {
    m.microsoftAuthenticatorMethods = value
}
func (m *Authentication) SetOperations(value []LongRunningOperation)() {
    m.operations = value
}
func (m *Authentication) SetPasswordlessMicrosoftAuthenticatorMethods(value []PasswordlessMicrosoftAuthenticatorAuthenticationMethod)() {
    m.passwordlessMicrosoftAuthenticatorMethods = value
}
func (m *Authentication) SetPasswordMethods(value []PasswordAuthenticationMethod)() {
    m.passwordMethods = value
}
func (m *Authentication) SetPhoneMethods(value []PhoneAuthenticationMethod)() {
    m.phoneMethods = value
}
func (m *Authentication) SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethod)() {
    m.temporaryAccessPassMethods = value
}
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethod)() {
    m.windowsHelloForBusinessMethods = value
}
