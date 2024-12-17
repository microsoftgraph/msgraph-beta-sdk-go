package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OidcPrivateJwtKeyClientAuthentication struct {
    OidcClientAuthentication
}
// NewOidcPrivateJwtKeyClientAuthentication instantiates a new OidcPrivateJwtKeyClientAuthentication and sets the default values.
func NewOidcPrivateJwtKeyClientAuthentication()(*OidcPrivateJwtKeyClientAuthentication) {
    m := &OidcPrivateJwtKeyClientAuthentication{
        OidcClientAuthentication: *NewOidcClientAuthentication(),
    }
    odataTypeValue := "#microsoft.graph.oidcPrivateJwtKeyClientAuthentication"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOidcPrivateJwtKeyClientAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOidcPrivateJwtKeyClientAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOidcPrivateJwtKeyClientAuthentication(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OidcPrivateJwtKeyClientAuthentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OidcClientAuthentication.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OidcPrivateJwtKeyClientAuthentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OidcClientAuthentication.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type OidcPrivateJwtKeyClientAuthenticationable interface {
    OidcClientAuthenticationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
