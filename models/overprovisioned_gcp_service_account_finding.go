package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedGcpServiceAccountFinding 
type OverprovisionedGcpServiceAccountFinding struct {
    IdentityFinding
}
// NewOverprovisionedGcpServiceAccountFinding instantiates a new overprovisionedGcpServiceAccountFinding and sets the default values.
func NewOverprovisionedGcpServiceAccountFinding()(*OverprovisionedGcpServiceAccountFinding) {
    m := &OverprovisionedGcpServiceAccountFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedGcpServiceAccountFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedGcpServiceAccountFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedGcpServiceAccountFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedGcpServiceAccountFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedGcpServiceAccountFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedGcpServiceAccountFindingable 
type OverprovisionedGcpServiceAccountFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
