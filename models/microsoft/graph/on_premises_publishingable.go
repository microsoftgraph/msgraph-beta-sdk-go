package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesPublishingable 
type OnPremisesPublishingable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAlternateUrl()(*string)
    GetApplicationServerTimeout()(*string)
    GetApplicationType()(*string)
    GetExternalAuthenticationType()(*ExternalAuthenticationType)
    GetExternalUrl()(*string)
    GetInternalUrl()(*string)
    GetIsBackendCertificateValidationEnabled()(*bool)
    GetIsHttpOnlyCookieEnabled()(*bool)
    GetIsOnPremPublishingEnabled()(*bool)
    GetIsPersistentCookieEnabled()(*bool)
    GetIsSecureCookieEnabled()(*bool)
    GetIsTranslateHostHeaderEnabled()(*bool)
    GetIsTranslateLinksInBodyEnabled()(*bool)
    GetSingleSignOnSettings()(OnPremisesPublishingSingleSignOnable)
    GetUseAlternateUrlForTranslationAndRedirect()(*bool)
    GetVerifiedCustomDomainCertificatesMetadata()(VerifiedCustomDomainCertificatesMetadataable)
    GetVerifiedCustomDomainKeyCredential()(KeyCredentialable)
    GetVerifiedCustomDomainPasswordCredential()(PasswordCredentialable)
    SetAlternateUrl(value *string)()
    SetApplicationServerTimeout(value *string)()
    SetApplicationType(value *string)()
    SetExternalAuthenticationType(value *ExternalAuthenticationType)()
    SetExternalUrl(value *string)()
    SetInternalUrl(value *string)()
    SetIsBackendCertificateValidationEnabled(value *bool)()
    SetIsHttpOnlyCookieEnabled(value *bool)()
    SetIsOnPremPublishingEnabled(value *bool)()
    SetIsPersistentCookieEnabled(value *bool)()
    SetIsSecureCookieEnabled(value *bool)()
    SetIsTranslateHostHeaderEnabled(value *bool)()
    SetIsTranslateLinksInBodyEnabled(value *bool)()
    SetSingleSignOnSettings(value OnPremisesPublishingSingleSignOnable)()
    SetUseAlternateUrlForTranslationAndRedirect(value *bool)()
    SetVerifiedCustomDomainCertificatesMetadata(value VerifiedCustomDomainCertificatesMetadataable)()
    SetVerifiedCustomDomainKeyCredential(value KeyCredentialable)()
    SetVerifiedCustomDomainPasswordCredential(value PasswordCredentialable)()
}
