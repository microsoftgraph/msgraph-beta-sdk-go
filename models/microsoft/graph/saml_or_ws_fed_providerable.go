package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SamlOrWsFedProviderable 
type SamlOrWsFedProviderable interface {
    IdentityProviderBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIssuerUri()(*string)
    GetMetadataExchangeUri()(*string)
    GetPassiveSignInUri()(*string)
    GetPreferredAuthenticationProtocol()(*AuthenticationProtocol)
    GetSigningCertificate()(*string)
    SetIssuerUri(value *string)()
    SetMetadataExchangeUri(value *string)()
    SetPassiveSignInUri(value *string)()
    SetPreferredAuthenticationProtocol(value *AuthenticationProtocol)()
    SetSigningCertificate(value *string)()
}
