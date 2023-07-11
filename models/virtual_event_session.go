package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventSession 
type VirtualEventSession struct {
    OnlineMeeting
}
// NewVirtualEventSession instantiates a new virtualEventSession and sets the default values.
func NewVirtualEventSession()(*VirtualEventSession) {
    m := &VirtualEventSession{
        OnlineMeeting: *NewOnlineMeeting(),
    }
    return m
}
// CreateVirtualEventSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnlineMeeting.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *VirtualEventSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnlineMeeting.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// VirtualEventSessionable 
type VirtualEventSessionable interface {
    OnlineMeetingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
