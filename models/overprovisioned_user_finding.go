package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedUserFinding 
type OverprovisionedUserFinding struct {
    IdentityFinding
}
// NewOverprovisionedUserFinding instantiates a new overprovisionedUserFinding and sets the default values.
func NewOverprovisionedUserFinding()(*OverprovisionedUserFinding) {
    m := &OverprovisionedUserFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedUserFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedUserFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedUserFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedUserFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedUserFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedUserFindingable 
type OverprovisionedUserFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
