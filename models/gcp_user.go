package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpUser 
type GcpUser struct {
    GcpIdentity
}
// NewGcpUser instantiates a new gcpUser and sets the default values.
func NewGcpUser()(*GcpUser) {
    m := &GcpUser{
        GcpIdentity: *NewGcpIdentity(),
    }
    return m
}
// CreateGcpUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpUser(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GcpUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// GcpUserable 
type GcpUserable interface {
    GcpIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
