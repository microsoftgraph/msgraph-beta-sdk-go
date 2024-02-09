package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdIdentitySource struct {
    PermissionsDefinitionIdentitySource
}
// NewEdIdentitySource instantiates a new EdIdentitySource and sets the default values.
func NewEdIdentitySource()(*EdIdentitySource) {
    m := &EdIdentitySource{
        PermissionsDefinitionIdentitySource: *NewPermissionsDefinitionIdentitySource(),
    }
    odataTypeValue := "#microsoft.graph.edIdentitySource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdIdentitySourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdIdentitySourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdIdentitySource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdIdentitySource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PermissionsDefinitionIdentitySource.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EdIdentitySource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PermissionsDefinitionIdentitySource.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type EdIdentitySourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PermissionsDefinitionIdentitySourceable
}
