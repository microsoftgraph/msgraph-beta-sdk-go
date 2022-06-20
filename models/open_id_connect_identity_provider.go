package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenIdConnectIdentityProvider 
type OpenIdConnectIdentityProvider struct {
    IdentityProviderBase
    // After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. Required.
    claimsMapping ClaimsMappingable
    // The client identifier for the application obtained when registering the application with the identity provider. Required.
    clientId *string
    // The client secret for the application obtained when registering the application with the identity provider. The clientSecret has a dependency on responseType. When responseType is code, a secret is required for the auth code exchange.When responseType is id_token the secret is not required because there is no code exchange. The id_token is returned directly from the authorization response. This is write-only. A read operation returns ****.
    clientSecret *string
    // The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
    domainHint *string
    // The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Provide the metadata URL for the OpenID Connect identity provider you add. Read-only. Required.
    metadataUrl *string
    // The response mode defines the method used to send data back from the custom identity provider to Azure AD B2C. Possible values: form_post, query. Required.
    responseMode *OpenIdConnectResponseMode
    // The response type describes the type of information sent back in the initial call to the authorization_endpoint of the custom identity provider. Possible values: code , id_token , token.  Required.
    responseType *OpenIdConnectResponseTypes
    // Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended, separated by a space. For more details about the scope limitations see RFC6749 Section 3.3. Required.
    scope *string
}
// NewOpenIdConnectIdentityProvider instantiates a new OpenIdConnectIdentityProvider and sets the default values.
func NewOpenIdConnectIdentityProvider()(*OpenIdConnectIdentityProvider) {
    m := &OpenIdConnectIdentityProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    return m
}
// CreateOpenIdConnectIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenIdConnectIdentityProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenIdConnectIdentityProvider(), nil
}
// GetClaimsMapping gets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. Required.
func (m *OpenIdConnectIdentityProvider) GetClaimsMapping()(ClaimsMappingable) {
    if m == nil {
        return nil
    } else {
        return m.claimsMapping
    }
}
// GetClientId gets the clientId property value. The client identifier for the application obtained when registering the application with the identity provider. Required.
func (m *OpenIdConnectIdentityProvider) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
// GetClientSecret gets the clientSecret property value. The client secret for the application obtained when registering the application with the identity provider. The clientSecret has a dependency on responseType. When responseType is code, a secret is required for the auth code exchange.When responseType is id_token the secret is not required because there is no code exchange. The id_token is returned directly from the authorization response. This is write-only. A read operation returns ****.
func (m *OpenIdConnectIdentityProvider) GetClientSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientSecret
    }
}
// GetDomainHint gets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectIdentityProvider) GetDomainHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainHint
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenIdConnectIdentityProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["claimsMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClaimsMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClaimsMapping(val.(ClaimsMappingable))
        }
        return nil
    }
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["domainHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainHint(val)
        }
        return nil
    }
    res["metadataUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataUrl(val)
        }
        return nil
    }
    res["responseMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOpenIdConnectResponseMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseMode(val.(*OpenIdConnectResponseMode))
        }
        return nil
    }
    res["responseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOpenIdConnectResponseTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseType(val.(*OpenIdConnectResponseTypes))
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    return res
}
// GetMetadataUrl gets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Provide the metadata URL for the OpenID Connect identity provider you add. Read-only. Required.
func (m *OpenIdConnectIdentityProvider) GetMetadataUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metadataUrl
    }
}
// GetResponseMode gets the responseMode property value. The response mode defines the method used to send data back from the custom identity provider to Azure AD B2C. Possible values: form_post, query. Required.
func (m *OpenIdConnectIdentityProvider) GetResponseMode()(*OpenIdConnectResponseMode) {
    if m == nil {
        return nil
    } else {
        return m.responseMode
    }
}
// GetResponseType gets the responseType property value. The response type describes the type of information sent back in the initial call to the authorization_endpoint of the custom identity provider. Possible values: code , id_token , token.  Required.
func (m *OpenIdConnectIdentityProvider) GetResponseType()(*OpenIdConnectResponseTypes) {
    if m == nil {
        return nil
    } else {
        return m.responseType
    }
}
// GetScope gets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended, separated by a space. For more details about the scope limitations see RFC6749 Section 3.3. Required.
func (m *OpenIdConnectIdentityProvider) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Serialize serializes information the current object
func (m *OpenIdConnectIdentityProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("claimsMapping", m.GetClaimsMapping())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainHint", m.GetDomainHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadataUrl", m.GetMetadataUrl())
        if err != nil {
            return err
        }
    }
    if m.GetResponseMode() != nil {
        cast := (*m.GetResponseMode()).String()
        err = writer.WriteStringValue("responseMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetResponseType() != nil {
        cast := (*m.GetResponseType()).String()
        err = writer.WriteStringValue("responseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClaimsMapping sets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. Required.
func (m *OpenIdConnectIdentityProvider) SetClaimsMapping(value ClaimsMappingable)() {
    if m != nil {
        m.claimsMapping = value
    }
}
// SetClientId sets the clientId property value. The client identifier for the application obtained when registering the application with the identity provider. Required.
func (m *OpenIdConnectIdentityProvider) SetClientId(value *string)() {
    if m != nil {
        m.clientId = value
    }
}
// SetClientSecret sets the clientSecret property value. The client secret for the application obtained when registering the application with the identity provider. The clientSecret has a dependency on responseType. When responseType is code, a secret is required for the auth code exchange.When responseType is id_token the secret is not required because there is no code exchange. The id_token is returned directly from the authorization response. This is write-only. A read operation returns ****.
func (m *OpenIdConnectIdentityProvider) SetClientSecret(value *string)() {
    if m != nil {
        m.clientSecret = value
    }
}
// SetDomainHint sets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectIdentityProvider) SetDomainHint(value *string)() {
    if m != nil {
        m.domainHint = value
    }
}
// SetMetadataUrl sets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Provide the metadata URL for the OpenID Connect identity provider you add. Read-only. Required.
func (m *OpenIdConnectIdentityProvider) SetMetadataUrl(value *string)() {
    if m != nil {
        m.metadataUrl = value
    }
}
// SetResponseMode sets the responseMode property value. The response mode defines the method used to send data back from the custom identity provider to Azure AD B2C. Possible values: form_post, query. Required.
func (m *OpenIdConnectIdentityProvider) SetResponseMode(value *OpenIdConnectResponseMode)() {
    if m != nil {
        m.responseMode = value
    }
}
// SetResponseType sets the responseType property value. The response type describes the type of information sent back in the initial call to the authorization_endpoint of the custom identity provider. Possible values: code , id_token , token.  Required.
func (m *OpenIdConnectIdentityProvider) SetResponseType(value *OpenIdConnectResponseTypes)() {
    if m != nil {
        m.responseType = value
    }
}
// SetScope sets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended, separated by a space. For more details about the scope limitations see RFC6749 Section 3.3. Required.
func (m *OpenIdConnectIdentityProvider) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
