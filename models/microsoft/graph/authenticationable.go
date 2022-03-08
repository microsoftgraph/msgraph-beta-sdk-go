package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Authenticationable 
type Authenticationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEmailMethods()([]EmailAuthenticationMethodable)
    GetFido2Methods()([]Fido2AuthenticationMethodable)
    GetMethods()([]AuthenticationMethodable)
    GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable)
    GetOperations()([]LongRunningOperationable)
    GetPasswordlessMicrosoftAuthenticatorMethods()([]PasswordlessMicrosoftAuthenticatorAuthenticationMethodable)
    GetPasswordMethods()([]PasswordAuthenticationMethodable)
    GetPhoneMethods()([]PhoneAuthenticationMethodable)
    GetSoftwareOathMethods()([]SoftwareOathAuthenticationMethodable)
    GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethodable)
    GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable)
    SetEmailMethods(value []EmailAuthenticationMethodable)()
    SetFido2Methods(value []Fido2AuthenticationMethodable)()
    SetMethods(value []AuthenticationMethodable)()
    SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)()
    SetOperations(value []LongRunningOperationable)()
    SetPasswordlessMicrosoftAuthenticatorMethods(value []PasswordlessMicrosoftAuthenticatorAuthenticationMethodable)()
    SetPasswordMethods(value []PasswordAuthenticationMethodable)()
    SetPhoneMethods(value []PhoneAuthenticationMethodable)()
    SetSoftwareOathMethods(value []SoftwareOathAuthenticationMethodable)()
    SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethodable)()
    SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)()
}
