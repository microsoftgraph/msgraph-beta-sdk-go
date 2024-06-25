package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OAuth2ClientCredential struct {
    OAuthClientCredential
}
// NewOAuth2ClientCredential instantiates a new OAuth2ClientCredential and sets the default values.
func NewOAuth2ClientCredential()(*OAuth2ClientCredential) {
    m := &OAuth2ClientCredential{
        OAuthClientCredential: *NewOAuthClientCredential(),
    }
    odataTypeValue := "#microsoft.graph.industryData.oAuth2ClientCredential"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOAuth2ClientCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOAuth2ClientCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOAuth2ClientCredential(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OAuth2ClientCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OAuthClientCredential.GetFieldDeserializers()
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
    res["tokenUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenUrl(val)
        }
        return nil
    }
    return res
}
// GetScope gets the scope property value. The OAuth scope that is provided to the authentication process.
// returns a *string when successful
func (m *OAuth2ClientCredential) GetScope()(*string) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTokenUrl gets the tokenUrl property value. The URL with which to retrieve the token after authentication happens.
// returns a *string when successful
func (m *OAuth2ClientCredential) GetTokenUrl()(*string) {
    val, err := m.GetBackingStore().Get("tokenUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OAuth2ClientCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OAuthClientCredential.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenUrl", m.GetTokenUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScope sets the scope property value. The OAuth scope that is provided to the authentication process.
func (m *OAuth2ClientCredential) SetScope(value *string)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenUrl sets the tokenUrl property value. The URL with which to retrieve the token after authentication happens.
func (m *OAuth2ClientCredential) SetTokenUrl(value *string)() {
    err := m.GetBackingStore().Set("tokenUrl", value)
    if err != nil {
        panic(err)
    }
}
type OAuth2ClientCredentialable interface {
    OAuthClientCredentialable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetScope()(*string)
    GetTokenUrl()(*string)
    SetScope(value *string)()
    SetTokenUrl(value *string)()
}
