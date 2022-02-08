package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SamlOrWsFedProvider 
type SamlOrWsFedProvider struct {
    IdentityProviderBase
    // Issuer URI of the federation server.
    issuerUri *string;
    // URI of the metadata exchange endpoint used for authentication from rich client applications.
    metadataExchangeUri *string;
    // URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
    passiveSignInUri *string;
    // Preferred authentication protocol. Supported values include saml or wsfed.
    preferredAuthenticationProtocol *AuthenticationProtocol;
    // Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
    signingCertificate *string;
}
// NewSamlOrWsFedProvider instantiates a new samlOrWsFedProvider and sets the default values.
func NewSamlOrWsFedProvider()(*SamlOrWsFedProvider) {
    m := &SamlOrWsFedProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    return m
}
// GetIssuerUri gets the issuerUri property value. Issuer URI of the federation server.
func (m *SamlOrWsFedProvider) GetIssuerUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.issuerUri
    }
}
// GetMetadataExchangeUri gets the metadataExchangeUri property value. URI of the metadata exchange endpoint used for authentication from rich client applications.
func (m *SamlOrWsFedProvider) GetMetadataExchangeUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metadataExchangeUri
    }
}
// GetPassiveSignInUri gets the passiveSignInUri property value. URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
func (m *SamlOrWsFedProvider) GetPassiveSignInUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.passiveSignInUri
    }
}
// GetPreferredAuthenticationProtocol gets the preferredAuthenticationProtocol property value. Preferred authentication protocol. Supported values include saml or wsfed.
func (m *SamlOrWsFedProvider) GetPreferredAuthenticationProtocol()(*AuthenticationProtocol) {
    if m == nil {
        return nil
    } else {
        return m.preferredAuthenticationProtocol
    }
}
// GetSigningCertificate gets the signingCertificate property value. Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
func (m *SamlOrWsFedProvider) GetSigningCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signingCertificate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlOrWsFedProvider) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["issuerUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerUri(val)
        }
        return nil
    }
    res["metadataExchangeUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataExchangeUri(val)
        }
        return nil
    }
    res["passiveSignInUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassiveSignInUri(val)
        }
        return nil
    }
    res["preferredAuthenticationProtocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredAuthenticationProtocol(val.(*AuthenticationProtocol))
        }
        return nil
    }
    res["signingCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificate(val)
        }
        return nil
    }
    return res
}
func (m *SamlOrWsFedProvider) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SamlOrWsFedProvider) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("issuerUri", m.GetIssuerUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadataExchangeUri", m.GetMetadataExchangeUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passiveSignInUri", m.GetPassiveSignInUri())
        if err != nil {
            return err
        }
    }
    if m.GetPreferredAuthenticationProtocol() != nil {
        cast := (*m.GetPreferredAuthenticationProtocol()).String()
        err = writer.WriteStringValue("preferredAuthenticationProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signingCertificate", m.GetSigningCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIssuerUri sets the issuerUri property value. Issuer URI of the federation server.
func (m *SamlOrWsFedProvider) SetIssuerUri(value *string)() {
    if m != nil {
        m.issuerUri = value
    }
}
// SetMetadataExchangeUri sets the metadataExchangeUri property value. URI of the metadata exchange endpoint used for authentication from rich client applications.
func (m *SamlOrWsFedProvider) SetMetadataExchangeUri(value *string)() {
    if m != nil {
        m.metadataExchangeUri = value
    }
}
// SetPassiveSignInUri sets the passiveSignInUri property value. URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
func (m *SamlOrWsFedProvider) SetPassiveSignInUri(value *string)() {
    if m != nil {
        m.passiveSignInUri = value
    }
}
// SetPreferredAuthenticationProtocol sets the preferredAuthenticationProtocol property value. Preferred authentication protocol. Supported values include saml or wsfed.
func (m *SamlOrWsFedProvider) SetPreferredAuthenticationProtocol(value *AuthenticationProtocol)() {
    if m != nil {
        m.preferredAuthenticationProtocol = value
    }
}
// SetSigningCertificate sets the signingCertificate property value. Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
func (m *SamlOrWsFedProvider) SetSigningCertificate(value *string)() {
    if m != nil {
        m.signingCertificate = value
    }
}
