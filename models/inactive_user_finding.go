package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InactiveUserFinding struct {
    IdentityFinding
}
// NewInactiveUserFinding instantiates a new InactiveUserFinding and sets the default values.
func NewInactiveUserFinding()(*InactiveUserFinding) {
    m := &InactiveUserFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateInactiveUserFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInactiveUserFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveUserFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InactiveUserFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InactiveUserFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type InactiveUserFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
