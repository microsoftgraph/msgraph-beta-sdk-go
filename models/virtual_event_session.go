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
    res["registrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationable)
                }
            }
            m.SetRegistrations(res)
        }
        return nil
    }
    return res
}
// GetRegistrations gets the registrations property value. Registration records of this virtual event session.
func (m *VirtualEventSession) GetRegistrations()([]VirtualEventRegistrationable) {
    val, err := m.GetBackingStore().Get("registrations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventRegistrationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnlineMeeting.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrations()))
        for i, v := range m.GetRegistrations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegistrations sets the registrations property value. Registration records of this virtual event session.
func (m *VirtualEventSession) SetRegistrations(value []VirtualEventRegistrationable)() {
    err := m.GetBackingStore().Set("registrations", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventSessionable 
type VirtualEventSessionable interface {
    OnlineMeetingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRegistrations()([]VirtualEventRegistrationable)
    SetRegistrations(value []VirtualEventRegistrationable)()
}
