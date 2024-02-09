package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UnknownSource struct {
    AuthorizationSystemIdentitySource
}
// NewUnknownSource instantiates a new UnknownSource and sets the default values.
func NewUnknownSource()(*UnknownSource) {
    m := &UnknownSource{
        AuthorizationSystemIdentitySource: *NewAuthorizationSystemIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.unknownSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnknownSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUnknownSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnknownSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UnknownSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthorizationSystemIdentitySource.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UnknownSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthorizationSystemIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type UnknownSourceable interface {
    AuthorizationSystemIdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
