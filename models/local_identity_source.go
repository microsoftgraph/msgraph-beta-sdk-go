package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalIdentitySource struct {
    PermissionsDefinitionIdentitySource
}
// NewLocalIdentitySource instantiates a new LocalIdentitySource and sets the default values.
func NewLocalIdentitySource()(*LocalIdentitySource) {
    m := &LocalIdentitySource{
        PermissionsDefinitionIdentitySource: *NewPermissionsDefinitionIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.localIdentitySource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateLocalIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalIdentitySource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionIdentitySource.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LocalIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type LocalIdentitySourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionIdentitySourceable
}
