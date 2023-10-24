package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SuperGcpServiceAccountFinding 
type SuperGcpServiceAccountFinding struct {
    IdentityFinding
}
// NewSuperGcpServiceAccountFinding instantiates a new superGcpServiceAccountFinding and sets the default values.
func NewSuperGcpServiceAccountFinding()(*SuperGcpServiceAccountFinding) {
    m := &SuperGcpServiceAccountFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateSuperGcpServiceAccountFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSuperGcpServiceAccountFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSuperGcpServiceAccountFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SuperGcpServiceAccountFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SuperGcpServiceAccountFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SuperGcpServiceAccountFindingable 
type SuperGcpServiceAccountFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
