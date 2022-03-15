package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Authentication provides operations to manage the deviceManagement singleton.
type Authentication struct {
    Entity
    // 
    emailMethods []EmailAuthenticationMethodable;
    // 
    fido2Methods []Fido2AuthenticationMethodable;
    // 
    methods []AuthenticationMethodable;
    // 
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethodable;
    // 
    operations []LongRunningOperationable;
    // 
    passwordlessMicrosoftAuthenticatorMethods []PasswordlessMicrosoftAuthenticatorAuthenticationMethodable;
    // 
    passwordMethods []PasswordAuthenticationMethodable;
    // 
    phoneMethods []PhoneAuthenticationMethodable;
    // 
    softwareOathMethods []SoftwareOathAuthenticationMethodable;
    // 
    temporaryAccessPassMethods []TemporaryAccessPassAuthenticationMethodable;
    // 
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethodable;
}
// NewAuthentication instantiates a new authentication and sets the default values.
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAuthentication(), nil
}
// GetEmailMethods gets the emailMethods property value. 
func (m *Authentication) GetEmailMethods()([]EmailAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.emailMethods
    }
}
// GetFido2Methods gets the fido2Methods property value. 
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.fido2Methods
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authentication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(EmailAuthenticationMethodable)
            }
            m.SetEmailMethods(res)
        }
        return nil
    }
    res["fido2Methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFido2AuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(Fido2AuthenticationMethodable)
            }
            m.SetFido2Methods(res)
        }
        return nil
    }
    res["methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodable)
            }
            m.SetMethods(res)
        }
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftAuthenticatorAuthenticationMethodable)
            }
            m.SetMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLongRunningOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LongRunningOperationable, len(val))
            for i, v := range val {
                res[i] = v.(LongRunningOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["passwordlessMicrosoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordlessMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordlessMicrosoftAuthenticatorAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordlessMicrosoftAuthenticatorAuthenticationMethodable)
            }
            m.SetPasswordlessMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["passwordMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordAuthenticationMethodable)
            }
            m.SetPasswordMethods(res)
        }
        return nil
    }
    res["phoneMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhoneAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PhoneAuthenticationMethodable)
            }
            m.SetPhoneMethods(res)
        }
        return nil
    }
    res["softwareOathMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SoftwareOathAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(SoftwareOathAuthenticationMethodable)
            }
            m.SetSoftwareOathMethods(res)
        }
        return nil
    }
    res["temporaryAccessPassMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemporaryAccessPassAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(TemporaryAccessPassAuthenticationMethodable)
            }
            m.SetTemporaryAccessPassMethods(res)
        }
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsHelloForBusinessAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsHelloForBusinessAuthenticationMethodable)
            }
            m.SetWindowsHelloForBusinessMethods(res)
        }
        return nil
    }
    return res
}
// GetMethods gets the methods property value. 
func (m *Authentication) GetMethods()([]AuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.methods
    }
}
// GetMicrosoftAuthenticatorMethods gets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.microsoftAuthenticatorMethods
    }
}
// GetOperations gets the operations property value. 
func (m *Authentication) GetOperations()([]LongRunningOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPasswordlessMicrosoftAuthenticatorMethods gets the passwordlessMicrosoftAuthenticatorMethods property value. 
func (m *Authentication) GetPasswordlessMicrosoftAuthenticatorMethods()([]PasswordlessMicrosoftAuthenticatorAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.passwordlessMicrosoftAuthenticatorMethods
    }
}
// GetPasswordMethods gets the passwordMethods property value. 
func (m *Authentication) GetPasswordMethods()([]PasswordAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.passwordMethods
    }
}
// GetPhoneMethods gets the phoneMethods property value. 
func (m *Authentication) GetPhoneMethods()([]PhoneAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.phoneMethods
    }
}
// GetSoftwareOathMethods gets the softwareOathMethods property value. 
func (m *Authentication) GetSoftwareOathMethods()([]SoftwareOathAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.softwareOathMethods
    }
}
// GetTemporaryAccessPassMethods gets the temporaryAccessPassMethods property value. 
func (m *Authentication) GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.temporaryAccessPassMethods
    }
}
// GetWindowsHelloForBusinessMethods gets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessMethods
    }
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
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFido2Methods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFido2Methods()))
        for i, v := range m.GetFido2Methods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fido2Methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMethods()))
        for i, v := range m.GetMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetMicrosoftAuthenticatorMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordlessMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordlessMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetPasswordlessMicrosoftAuthenticatorMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("passwordlessMicrosoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordMethods()))
        for i, v := range m.GetPasswordMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("passwordMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhoneMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhoneMethods()))
        for i, v := range m.GetPhoneMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phoneMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSoftwareOathMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSoftwareOathMethods()))
        for i, v := range m.GetSoftwareOathMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("softwareOathMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemporaryAccessPassMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemporaryAccessPassMethods()))
        for i, v := range m.GetTemporaryAccessPassMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("temporaryAccessPassMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsHelloForBusinessMethods() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsHelloForBusinessMethods()))
        for i, v := range m.GetWindowsHelloForBusinessMethods() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsHelloForBusinessMethods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailMethods sets the emailMethods property value. 
func (m *Authentication) SetEmailMethods(value []EmailAuthenticationMethodable)() {
    if m != nil {
        m.emailMethods = value
    }
}
// SetFido2Methods sets the fido2Methods property value. 
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethodable)() {
    if m != nil {
        m.fido2Methods = value
    }
}
// SetMethods sets the methods property value. 
func (m *Authentication) SetMethods(value []AuthenticationMethodable)() {
    if m != nil {
        m.methods = value
    }
}
// SetMicrosoftAuthenticatorMethods sets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)() {
    if m != nil {
        m.microsoftAuthenticatorMethods = value
    }
}
// SetOperations sets the operations property value. 
func (m *Authentication) SetOperations(value []LongRunningOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetPasswordlessMicrosoftAuthenticatorMethods sets the passwordlessMicrosoftAuthenticatorMethods property value. 
func (m *Authentication) SetPasswordlessMicrosoftAuthenticatorMethods(value []PasswordlessMicrosoftAuthenticatorAuthenticationMethodable)() {
    if m != nil {
        m.passwordlessMicrosoftAuthenticatorMethods = value
    }
}
// SetPasswordMethods sets the passwordMethods property value. 
func (m *Authentication) SetPasswordMethods(value []PasswordAuthenticationMethodable)() {
    if m != nil {
        m.passwordMethods = value
    }
}
// SetPhoneMethods sets the phoneMethods property value. 
func (m *Authentication) SetPhoneMethods(value []PhoneAuthenticationMethodable)() {
    if m != nil {
        m.phoneMethods = value
    }
}
// SetSoftwareOathMethods sets the softwareOathMethods property value. 
func (m *Authentication) SetSoftwareOathMethods(value []SoftwareOathAuthenticationMethodable)() {
    if m != nil {
        m.softwareOathMethods = value
    }
}
// SetTemporaryAccessPassMethods sets the temporaryAccessPassMethods property value. 
func (m *Authentication) SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethodable)() {
    if m != nil {
        m.temporaryAccessPassMethods = value
    }
}
// SetWindowsHelloForBusinessMethods sets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)() {
    if m != nil {
        m.windowsHelloForBusinessMethods = value
    }
}
