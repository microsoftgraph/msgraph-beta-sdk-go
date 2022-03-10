package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PhoneAuthenticationMethodable 
type PhoneAuthenticationMethodable interface {
    AuthenticationMethodable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPhoneNumber()(*string)
    GetPhoneType()(*AuthenticationPhoneType)
    GetSmsSignInState()(*AuthenticationMethodSignInState)
    SetPhoneNumber(value *string)()
    SetPhoneType(value *AuthenticationPhoneType)()
    SetSmsSignInState(value *AuthenticationMethodSignInState)()
}
