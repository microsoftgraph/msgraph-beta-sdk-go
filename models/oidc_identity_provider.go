package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OidcIdentityProvider struct {
    IdentityProviderBase
}
// NewOidcIdentityProvider instantiates a new OidcIdentityProvider and sets the default values.
func NewOidcIdentityProvider()(*OidcIdentityProvider) {
    m := &OidcIdentityProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    odataTypeValue := "#microsoft.graph.oidcIdentityProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOidcIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOidcIdentityProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOidcIdentityProvider(), nil
}
// GetClientAuthentication gets the clientAuthentication property value. The clientAuthentication property
// returns a OidcClientAuthenticationable when successful
func (m *OidcIdentityProvider) GetClientAuthentication()(OidcClientAuthenticationable) {
    val, err := m.GetBackingStore().Get("clientAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OidcClientAuthenticationable)
    }
    return nil
}
// GetClientId gets the clientId property value. The client ID for the application obtained when registering the application with the identity provider.
// returns a *string when successful
func (m *OidcIdentityProvider) GetClientId()(*string) {
    val, err := m.GetBackingStore().Get("clientId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OidcIdentityProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["clientAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOidcClientAuthenticationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAuthentication(val.(OidcClientAuthenticationable))
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
    res["inboundClaimMapping"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOidcInboundClaimMappingOverrideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundClaimMapping(val.(OidcInboundClaimMappingOverrideable))
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["responseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOidcResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseType(val.(*OidcResponseType))
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
    res["wellKnownEndpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWellKnownEndpoint(val)
        }
        return nil
    }
    return res
}
// GetInboundClaimMapping gets the inboundClaimMapping property value. After the OIDC provider sends an ID token back to Microsoft Entra External ID, Microsoft Entra External ID needs to be able to map the claims from the received token to the claims that Microsoft Entra ID recognizes and uses. This complex type captures that mapping.
// returns a OidcInboundClaimMappingOverrideable when successful
func (m *OidcIdentityProvider) GetInboundClaimMapping()(OidcInboundClaimMappingOverrideable) {
    val, err := m.GetBackingStore().Get("inboundClaimMapping")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OidcInboundClaimMappingOverrideable)
    }
    return nil
}
// GetIssuer gets the issuer property value. The issuer URI. Issuer URI is a case-sensitive URL using https scheme contains scheme, host, and optionally, port number and path components and no query or fragment components. Note: Configuring other Microsoft Entra tenants as an external identity provider is currently not supported. As a result, the microsoftonline.com domain in the issuer URI is not accepted.
// returns a *string when successful
func (m *OidcIdentityProvider) GetIssuer()(*string) {
    val, err := m.GetBackingStore().Get("issuer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResponseType gets the responseType property value. The responseType property
// returns a *OidcResponseType when successful
func (m *OidcIdentityProvider) GetResponseType()(*OidcResponseType) {
    val, err := m.GetBackingStore().Get("responseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OidcResponseType)
    }
    return nil
}
// GetScope gets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider.
// returns a *string when successful
func (m *OidcIdentityProvider) GetScope()(*string) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWellKnownEndpoint gets the wellKnownEndpoint property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Note: The metadata document should, at minimum, contain the following properties: issuer, authorizationendpoint, tokenendpoint, tokenendpointauthmethodssupported, responsetypessupported, subjecttypessupported and jwks_uri. Visit OpenID Connect Discovery specifications for more details.
// returns a *string when successful
func (m *OidcIdentityProvider) GetWellKnownEndpoint()(*string) {
    val, err := m.GetBackingStore().Get("wellKnownEndpoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OidcIdentityProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("clientAuthentication", m.GetClientAuthentication())
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
        err = writer.WriteObjectValue("inboundClaimMapping", m.GetInboundClaimMapping())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
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
    {
        err = writer.WriteStringValue("wellKnownEndpoint", m.GetWellKnownEndpoint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientAuthentication sets the clientAuthentication property value. The clientAuthentication property
func (m *OidcIdentityProvider) SetClientAuthentication(value OidcClientAuthenticationable)() {
    err := m.GetBackingStore().Set("clientAuthentication", value)
    if err != nil {
        panic(err)
    }
}
// SetClientId sets the clientId property value. The client ID for the application obtained when registering the application with the identity provider.
func (m *OidcIdentityProvider) SetClientId(value *string)() {
    err := m.GetBackingStore().Set("clientId", value)
    if err != nil {
        panic(err)
    }
}
// SetInboundClaimMapping sets the inboundClaimMapping property value. After the OIDC provider sends an ID token back to Microsoft Entra External ID, Microsoft Entra External ID needs to be able to map the claims from the received token to the claims that Microsoft Entra ID recognizes and uses. This complex type captures that mapping.
func (m *OidcIdentityProvider) SetInboundClaimMapping(value OidcInboundClaimMappingOverrideable)() {
    err := m.GetBackingStore().Set("inboundClaimMapping", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuer sets the issuer property value. The issuer URI. Issuer URI is a case-sensitive URL using https scheme contains scheme, host, and optionally, port number and path components and no query or fragment components. Note: Configuring other Microsoft Entra tenants as an external identity provider is currently not supported. As a result, the microsoftonline.com domain in the issuer URI is not accepted.
func (m *OidcIdentityProvider) SetIssuer(value *string)() {
    err := m.GetBackingStore().Set("issuer", value)
    if err != nil {
        panic(err)
    }
}
// SetResponseType sets the responseType property value. The responseType property
func (m *OidcIdentityProvider) SetResponseType(value *OidcResponseType)() {
    err := m.GetBackingStore().Set("responseType", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider.
func (m *OidcIdentityProvider) SetScope(value *string)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetWellKnownEndpoint sets the wellKnownEndpoint property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration. Note: The metadata document should, at minimum, contain the following properties: issuer, authorizationendpoint, tokenendpoint, tokenendpointauthmethodssupported, responsetypessupported, subjecttypessupported and jwks_uri. Visit OpenID Connect Discovery specifications for more details.
func (m *OidcIdentityProvider) SetWellKnownEndpoint(value *string)() {
    err := m.GetBackingStore().Set("wellKnownEndpoint", value)
    if err != nil {
        panic(err)
    }
}
type OidcIdentityProviderable interface {
    IdentityProviderBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientAuthentication()(OidcClientAuthenticationable)
    GetClientId()(*string)
    GetInboundClaimMapping()(OidcInboundClaimMappingOverrideable)
    GetIssuer()(*string)
    GetResponseType()(*OidcResponseType)
    GetScope()(*string)
    GetWellKnownEndpoint()(*string)
    SetClientAuthentication(value OidcClientAuthenticationable)()
    SetClientId(value *string)()
    SetInboundClaimMapping(value OidcInboundClaimMappingOverrideable)()
    SetIssuer(value *string)()
    SetResponseType(value *OidcResponseType)()
    SetScope(value *string)()
    SetWellKnownEndpoint(value *string)()
}
