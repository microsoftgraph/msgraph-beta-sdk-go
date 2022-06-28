package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenIdConnectProvider 
type OpenIdConnectProvider struct {
    IdentityProvider
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
    claimsMapping ClaimsMappingable
    // The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
    domainHint *string
    // The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
    metadataUrl *string
    // The response mode defines the method that should be used to send the data back from the custom identity provider to Azure AD B2C. The following response modes can be used: form_post, query. query response mode means the code or token is returned as a query parameter. form_post response mode is recommended for the best security. The response is transmitted via the HTTP POST method, with the code or token being encoded in the body using the application/x-www-form-urlencoded format. It is a required property.
    responseMode *OpenIdConnectResponseMode
    // response type describes what kind of information is sent back in the initial call to the authorization_endpoint of the custom identity provider. The following response types can be used: code , id_token , token. It is a required property.
    responseType *OpenIdConnectResponseTypes
    // Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
    scope *string
}
// NewOpenIdConnectProvider instantiates a new OpenIdConnectProvider and sets the default values.
func NewOpenIdConnectProvider()(*OpenIdConnectProvider) {
    m := &OpenIdConnectProvider{
        IdentityProvider: *NewIdentityProvider(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOpenIdConnectProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOpenIdConnectProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOpenIdConnectProvider(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OpenIdConnectProvider) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClaimsMapping gets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
func (m *OpenIdConnectProvider) GetClaimsMapping()(ClaimsMappingable) {
    if m == nil {
        return nil
    } else {
        return m.claimsMapping
    }
}
// GetDomainHint gets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectProvider) GetDomainHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainHint
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OpenIdConnectProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProvider.GetFieldDeserializers()
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
// GetMetadataUrl gets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
func (m *OpenIdConnectProvider) GetMetadataUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metadataUrl
    }
}
// GetResponseMode gets the responseMode property value. The response mode defines the method that should be used to send the data back from the custom identity provider to Azure AD B2C. The following response modes can be used: form_post, query. query response mode means the code or token is returned as a query parameter. form_post response mode is recommended for the best security. The response is transmitted via the HTTP POST method, with the code or token being encoded in the body using the application/x-www-form-urlencoded format. It is a required property.
func (m *OpenIdConnectProvider) GetResponseMode()(*OpenIdConnectResponseMode) {
    if m == nil {
        return nil
    } else {
        return m.responseMode
    }
}
// GetResponseType gets the responseType property value. response type describes what kind of information is sent back in the initial call to the authorization_endpoint of the custom identity provider. The following response types can be used: code , id_token , token. It is a required property.
func (m *OpenIdConnectProvider) GetResponseType()(*OpenIdConnectResponseTypes) {
    if m == nil {
        return nil
    } else {
        return m.responseType
    }
}
// GetScope gets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
func (m *OpenIdConnectProvider) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Serialize serializes information the current object
func (m *OpenIdConnectProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProvider.Serialize(writer)
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OpenIdConnectProvider) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClaimsMapping sets the claimsMapping property value. After the OIDC provider sends an ID token back to Azure AD, Azure AD needs to be able to map the claims from the received token to the claims that Azure AD recognizes and uses. This complex type captures that mapping. It is a required property.
func (m *OpenIdConnectProvider) SetClaimsMapping(value ClaimsMappingable)() {
    if m != nil {
        m.claimsMapping = value
    }
}
// SetDomainHint sets the domainHint property value. The domain hint can be used to skip directly to the sign-in page of the specified identity provider, instead of having the user make a selection among the list of available identity providers.
func (m *OpenIdConnectProvider) SetDomainHint(value *string)() {
    if m != nil {
        m.domainHint = value
    }
}
// SetMetadataUrl sets the metadataUrl property value. The URL for the metadata document of the OpenID Connect identity provider. Every OpenID Connect identity provider describes a metadata document that contains most of the information required to perform sign-in. This includes information such as the URLs to use and the location of the service's public signing keys. The OpenID Connect metadata document is always located at an endpoint that ends in .well-known/openid-configuration . For the OpenID Connect identity provider you are looking to add, you will need to provide the metadata URL. It is a required property and is read only after creation.
func (m *OpenIdConnectProvider) SetMetadataUrl(value *string)() {
    if m != nil {
        m.metadataUrl = value
    }
}
// SetResponseMode sets the responseMode property value. The response mode defines the method that should be used to send the data back from the custom identity provider to Azure AD B2C. The following response modes can be used: form_post, query. query response mode means the code or token is returned as a query parameter. form_post response mode is recommended for the best security. The response is transmitted via the HTTP POST method, with the code or token being encoded in the body using the application/x-www-form-urlencoded format. It is a required property.
func (m *OpenIdConnectProvider) SetResponseMode(value *OpenIdConnectResponseMode)() {
    if m != nil {
        m.responseMode = value
    }
}
// SetResponseType sets the responseType property value. response type describes what kind of information is sent back in the initial call to the authorization_endpoint of the custom identity provider. The following response types can be used: code , id_token , token. It is a required property.
func (m *OpenIdConnectProvider) SetResponseType(value *OpenIdConnectResponseTypes)() {
    if m != nil {
        m.responseType = value
    }
}
// SetScope sets the scope property value. Scope defines the information and permissions you are looking to gather from your custom identity provider. OpenID Connect requests must contain the openid scope value in order to receive the ID token from the identity provider. Without the ID token, users are not able to sign in to Azure AD B2C using the custom identity provider. Other scopes can be appended separated by space. For more details about the scope limitations see RFC6749 Section 3.3. It is a required property.
func (m *OpenIdConnectProvider) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
