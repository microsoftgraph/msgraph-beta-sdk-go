package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InactiveAzureServicePrincipalFinding 
type InactiveAzureServicePrincipalFinding struct {
    IdentityFinding
}
// NewInactiveAzureServicePrincipalFinding instantiates a new inactiveAzureServicePrincipalFinding and sets the default values.
func NewInactiveAzureServicePrincipalFinding()(*InactiveAzureServicePrincipalFinding) {
    m := &InactiveAzureServicePrincipalFinding{
        IdentityFinding: *NewIdentityFinding(),
    }
    return m
}
// CreateInactiveAzureServicePrincipalFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInactiveAzureServicePrincipalFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInactiveAzureServicePrincipalFinding(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InactiveAzureServicePrincipalFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityFinding.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *InactiveAzureServicePrincipalFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityFinding.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// InactiveAzureServicePrincipalFindingable 
type InactiveAzureServicePrincipalFindingable interface {
    IdentityFindingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
