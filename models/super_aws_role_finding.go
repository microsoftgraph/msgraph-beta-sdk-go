package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SuperAwsRoleFinding 
type SuperAwsRoleFinding struct {
    IdentityFinding
}
// NewSuperAwsRoleFinding instantiates a new superAwsRoleFinding and sets the default values.
func NewSuperAwsRoleFinding()(*SuperAwsRoleFinding) {
    m := &SuperAwsRoleFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateSuperAwsRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSuperAwsRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSuperAwsRoleFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SuperAwsRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SuperAwsRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SuperAwsRoleFindingable 
type SuperAwsRoleFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
