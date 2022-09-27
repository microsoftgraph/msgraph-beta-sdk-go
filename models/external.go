package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// External 
type External struct {
    Entity
    // The connections property
    connections []ExternalConnectionable
}
// NewExternal instantiates a new External and sets the default values.
func NewExternal()(*External) {
    m := &External{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.external";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternal(), nil
}
// GetConnections gets the connections property value. The connections property
func (m *External) GetConnections()([]ExternalConnectionable) {
    return m.connections
}
// GetFieldDeserializers the deserialization information for the current model
func (m *External) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExternalConnectionFromDiscriminatorValue , m.SetConnections)
    return res
}
// Serialize serializes information the current object
func (m *External) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConnections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetConnections())
        err = writer.WriteCollectionOfObjectValues("connections", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnections sets the connections property value. The connections property
func (m *External) SetConnections(value []ExternalConnectionable)() {
    m.connections = value
}
