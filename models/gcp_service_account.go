package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GcpServiceAccount 
type GcpServiceAccount struct {
    GcpIdentity
}
// NewGcpServiceAccount instantiates a new gcpServiceAccount and sets the default values.
func NewGcpServiceAccount()(*GcpServiceAccount) {
    m := &GcpServiceAccount{
        GcpIdentity: *NewGcpIdentity(),
    }
    odataTypeValue := "#microsoft.graph.gcpServiceAccount"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpServiceAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGcpServiceAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpServiceAccount(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GcpServiceAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GcpServiceAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// GcpServiceAccountable 
type GcpServiceAccountable interface {
    GcpIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
