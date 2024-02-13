package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type External struct {
    Entity
}
// NewExternal instantiates a new External and sets the default values.
func NewExternal()(*External) {
    m := &External{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExternalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateExternalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternal(), nil
}
// GetConnections gets the connections property value. The connections property
// returns a []ExternalConnectionable when successful
func (m *External) GetConnections()([]ExternalConnectionable) {
    val, err := m.GetBackingStore().Get("connections")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExternalConnectionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *External) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExternalConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExternalConnectionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExternalConnectionable)
                }
            }
            m.SetConnections(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *External) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConnections()))
        for i, v := range m.GetConnections() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("connections", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnections sets the connections property value. The connections property
func (m *External) SetConnections(value []ExternalConnectionable)() {
    err := m.GetBackingStore().Set("connections", value)
    if err != nil {
        panic(err)
    }
}
type Externalable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnections()([]ExternalConnectionable)
    SetConnections(value []ExternalConnectionable)()
}
