package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GcpGroup struct {
    GcpIdentity
}
// NewGcpGroup instantiates a new GcpGroup and sets the default values.
func NewGcpGroup()(*GcpGroup) {
    m := &GcpGroup{
        GcpIdentity: *NewGcpIdentity(),
    }
    odataTypeValue := "#microsoft.graph.gcpGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateGcpGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGcpGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGcpGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GcpGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.GcpIdentity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *GcpGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.GcpIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type GcpGroupable interface {
    GcpIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
