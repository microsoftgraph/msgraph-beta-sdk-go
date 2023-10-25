package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedAwsRoleFinding 
type OverprovisionedAwsRoleFinding struct {
    IdentityFinding
}
// NewOverprovisionedAwsRoleFinding instantiates a new overprovisionedAwsRoleFinding and sets the default values.
func NewOverprovisionedAwsRoleFinding()(*OverprovisionedAwsRoleFinding) {
    m := &OverprovisionedAwsRoleFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedAwsRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedAwsRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedAwsRoleFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedAwsRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedAwsRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedAwsRoleFindingable 
type OverprovisionedAwsRoleFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
