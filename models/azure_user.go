package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AzureUser struct {
    AzureIdentity
}
// NewAzureUser instantiates a new AzureUser and sets the default values.
func NewAzureUser()(*AzureUser) {
    m := &AzureUser{
        AzureIdentity: *NewAzureIdentity(),
    }
    odataTypeValue := "#microsoft.graph.azureUser"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAzureUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAzureUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAzureUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AzureUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AzureIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AzureUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AzureIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type AzureUserable interface {
    AzureIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
