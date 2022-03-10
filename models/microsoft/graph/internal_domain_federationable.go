package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InternalDomainFederationable 
type InternalDomainFederationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    SamlOrWsFedProviderable
    GetActiveSignInUri()(*string)
    GetFederatedIdpMfaBehavior()(*FederatedIdpMfaBehavior)
    GetIsSignedAuthenticationRequestRequired()(*bool)
    GetNextSigningCertificate()(*string)
    GetPromptLoginBehavior()(*PromptLoginBehavior)
    GetSigningCertificateUpdateStatus()(SigningCertificateUpdateStatusable)
    GetSignOutUri()(*string)
    SetActiveSignInUri(value *string)()
    SetFederatedIdpMfaBehavior(value *FederatedIdpMfaBehavior)()
    SetIsSignedAuthenticationRequestRequired(value *bool)()
    SetNextSigningCertificate(value *string)()
    SetPromptLoginBehavior(value *PromptLoginBehavior)()
    SetSigningCertificateUpdateStatus(value SigningCertificateUpdateStatusable)()
    SetSignOutUri(value *string)()
}
