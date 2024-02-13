package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagedAppRegistration represents the synchronization details of a Windows app, with management capabilities, for a specific user.
type WindowsManagedAppRegistration struct {
    ManagedAppRegistration
}
// NewWindowsManagedAppRegistration instantiates a new WindowsManagedAppRegistration and sets the default values.
func NewWindowsManagedAppRegistration()(*WindowsManagedAppRegistration) {
    m := &WindowsManagedAppRegistration{
        ManagedAppRegistration: *NewManagedAppRegistration(),
    }
    odataTypeValue := "#microsoft.graph.windowsManagedAppRegistration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsManagedAppRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsManagedAppRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagedAppRegistration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsManagedAppRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppRegistration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WindowsManagedAppRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppRegistration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type WindowsManagedAppRegistrationable interface {
    ManagedAppRegistrationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
