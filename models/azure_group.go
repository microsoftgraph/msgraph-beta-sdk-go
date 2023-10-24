package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AzureGroup 
type AzureGroup struct {
    AzureIdentity
}
// NewAzureGroup instantiates a new azureGroup and sets the default values.
func NewAzureGroup()(*AzureGroup) {
    m := &AzureGroup{
        AzureIdentity: *NewAzureIdentity(),
    }
    odataTypeValue := "#microsoft.graph.azureGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAzureGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AzureGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzureIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzureIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// AzureGroupable 
type AzureGroupable interface {
    AzureIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
