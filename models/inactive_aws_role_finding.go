package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveAwsRoleFinding 
type InactiveAwsRoleFinding struct {
    IdentityFinding
}
// NewInactiveAwsRoleFinding instantiates a new inactiveAwsRoleFinding and sets the default values.
func NewInactiveAwsRoleFinding()(*InactiveAwsRoleFinding) {
    m := &InactiveAwsRoleFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateInactiveAwsRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveAwsRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveAwsRoleFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveAwsRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InactiveAwsRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InactiveAwsRoleFindingable 
type InactiveAwsRoleFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
