package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OidcClientSecretAuthentication struct {
    OidcClientAuthentication
}
// NewOidcClientSecretAuthentication instantiates a new OidcClientSecretAuthentication and sets the default values.
func NewOidcClientSecretAuthentication()(*OidcClientSecretAuthentication) {
    m := &OidcClientSecretAuthentication{
        OidcClientAuthentication: *NewOidcClientAuthentication(),
    }
    odataTypeValue := "#microsoft.graph.oidcClientSecretAuthentication"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOidcClientSecretAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOidcClientSecretAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOidcClientSecretAuthentication(), nil
}
// GetClientSecret gets the clientSecret property value. The client secret obtained from configuring the client application on the external OpenID Connect identity provider. The property includes the client secret and enables the identity provider to use either the clientsecretpost authentication method.
// returns a *string when successful
func (m *OidcClientSecretAuthentication) GetClientSecret()(*string) {
    val, err := m.GetBackingStore().Get("clientSecret")
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
func (m *OidcClientSecretAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OidcClientAuthentication.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *OidcClientSecretAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OidcClientAuthentication.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientSecret sets the clientSecret property value. The client secret obtained from configuring the client application on the external OpenID Connect identity provider. The property includes the client secret and enables the identity provider to use either the clientsecretpost authentication method.
func (m *OidcClientSecretAuthentication) SetClientSecret(value *string)() {
    err := m.GetBackingStore().Set("clientSecret", value)
    if err != nil {
        panic(err)
    }
}
type OidcClientSecretAuthenticationable interface {
    OidcClientAuthenticationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientSecret()(*string)
    SetClientSecret(value *string)()
}
