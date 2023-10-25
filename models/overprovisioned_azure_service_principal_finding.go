package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OverprovisionedAzureServicePrincipalFinding 
type OverprovisionedAzureServicePrincipalFinding struct {
    IdentityFinding
}
// NewOverprovisionedAzureServicePrincipalFinding instantiates a new overprovisionedAzureServicePrincipalFinding and sets the default values.
func NewOverprovisionedAzureServicePrincipalFinding()(*OverprovisionedAzureServicePrincipalFinding) {
    m := &OverprovisionedAzureServicePrincipalFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateOverprovisionedAzureServicePrincipalFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOverprovisionedAzureServicePrincipalFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOverprovisionedAzureServicePrincipalFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OverprovisionedAzureServicePrincipalFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OverprovisionedAzureServicePrincipalFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OverprovisionedAzureServicePrincipalFindingable 
type OverprovisionedAzureServicePrincipalFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
