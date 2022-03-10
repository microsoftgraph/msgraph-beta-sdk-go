package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CredentialUserRegistrationDetailsable 
type CredentialUserRegistrationDetailsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAuthMethods()([]RegistrationAuthMethod)
    GetIsCapable()(*bool)
    GetIsEnabled()(*bool)
    GetIsMfaRegistered()(*bool)
    GetIsRegistered()(*bool)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetAuthMethods(value []RegistrationAuthMethod)()
    SetIsCapable(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsMfaRegistered(value *bool)()
    SetIsRegistered(value *bool)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
