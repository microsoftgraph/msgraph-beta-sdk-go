package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventWebinar 
type VirtualEventWebinar struct {
    VirtualEvent
}
// NewVirtualEventWebinar instantiates a new VirtualEventWebinar and sets the default values.
func NewVirtualEventWebinar()(*VirtualEventWebinar) {
    m := &VirtualEventWebinar{
        VirtualEvent: *NewVirtualEvent(),
    }
    return m
}
// CreateVirtualEventWebinarFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventWebinarFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventWebinar(), nil
}
// GetAudience gets the audience property value. To whom the webinar is visible.
func (m *VirtualEventWebinar) GetAudience()(*MeetingAudience) {
    val, err := m.GetBackingStore().Get("audience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingAudience)
    }
    return nil
}
// GetCoOrganizers gets the coOrganizers property value. Identity information of co-organizers of the webinar.
func (m *VirtualEventWebinar) GetCoOrganizers()([]CommunicationsUserIdentityable) {
    val, err := m.GetBackingStore().Get("coOrganizers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CommunicationsUserIdentityable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventWebinar) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VirtualEvent.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(*MeetingAudience))
        }
        return nil
    }
    res["coOrganizers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommunicationsUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommunicationsUserIdentityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CommunicationsUserIdentityable)
                }
            }
            m.SetCoOrganizers(res)
        }
        return nil
    }
    res["registration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistration(val.(VirtualEventRegistrationable))
        }
        return nil
    }
    return res
}
// GetRegistration gets the registration property value. Registration configuration of the webinar.
func (m *VirtualEventWebinar) GetRegistration()(VirtualEventRegistrationable) {
    val, err := m.GetBackingStore().Get("registration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualEventRegistrationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventWebinar) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VirtualEvent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudience() != nil {
        cast := (*m.GetAudience()).String()
        err = writer.WriteStringValue("audience", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCoOrganizers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCoOrganizers()))
        for i, v := range m.GetCoOrganizers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("coOrganizers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registration", m.GetRegistration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. To whom the webinar is visible.
func (m *VirtualEventWebinar) SetAudience(value *MeetingAudience)() {
    err := m.GetBackingStore().Set("audience", value)
    if err != nil {
        panic(err)
    }
}
// SetCoOrganizers sets the coOrganizers property value. Identity information of co-organizers of the webinar.
func (m *VirtualEventWebinar) SetCoOrganizers(value []CommunicationsUserIdentityable)() {
    err := m.GetBackingStore().Set("coOrganizers", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistration sets the registration property value. Registration configuration of the webinar.
func (m *VirtualEventWebinar) SetRegistration(value VirtualEventRegistrationable)() {
    err := m.GetBackingStore().Set("registration", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventWebinarable 
type VirtualEventWebinarable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventable
    GetAudience()(*MeetingAudience)
    GetCoOrganizers()([]CommunicationsUserIdentityable)
    GetRegistration()(VirtualEventRegistrationable)
    SetAudience(value *MeetingAudience)()
    SetCoOrganizers(value []CommunicationsUserIdentityable)()
    SetRegistration(value VirtualEventRegistrationable)()
}
