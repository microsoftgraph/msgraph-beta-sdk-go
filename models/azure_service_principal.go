package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureServicePrincipal 
type AzureServicePrincipal struct {
    AzureIdentity
}
// NewAzureServicePrincipal instantiates a new azureServicePrincipal and sets the default values.
func NewAzureServicePrincipal()(*AzureServicePrincipal) {
    m := &AzureServicePrincipal{
        AzureIdentity: *NewAzureIdentity(),
    }
    return m
}
// CreateAzureServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureServicePrincipal(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzureIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzureIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzureServicePrincipalable 
type AzureServicePrincipalable interface {
    AzureIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
