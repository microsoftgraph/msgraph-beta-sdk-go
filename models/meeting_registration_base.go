package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationBase 
type MeetingRegistrationBase struct {
    Entity
}
// NewMeetingRegistrationBase instantiates a new meetingRegistrationBase and sets the default values.
func NewMeetingRegistrationBase()(*MeetingRegistrationBase) {
    m := &MeetingRegistrationBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMeetingRegistrationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.externalMeetingRegistration":
                        return NewExternalMeetingRegistration(), nil
                    case "#microsoft.graph.meetingRegistration":
                        return NewMeetingRegistration(), nil
                }
            }
        }
    }
    return NewMeetingRegistrationBase(), nil
}
// GetAllowedRegistrant gets the allowedRegistrant property value. Specifies who can register for the meeting.
func (m *MeetingRegistrationBase) GetAllowedRegistrant()(*MeetingAudience) {
    val, err := m.GetBackingStore().Get("allowedRegistrant")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingAudience)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRegistrant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedRegistrant(val.(*MeetingAudience))
        }
        return nil
    }
    res["registrants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingRegistrantBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingRegistrantBaseable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingRegistrantBaseable)
            }
            m.SetRegistrants(res)
        }
        return nil
    }
    return res
}
// GetRegistrants gets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) GetRegistrants()([]MeetingRegistrantBaseable) {
    val, err := m.GetBackingStore().Get("registrants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingRegistrantBaseable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MeetingRegistrationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedRegistrant() != nil {
        cast := (*m.GetAllowedRegistrant()).String()
        err = writer.WriteStringValue("allowedRegistrant", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegistrants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrants()))
        for i, v := range m.GetRegistrants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("registrants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedRegistrant sets the allowedRegistrant property value. Specifies who can register for the meeting.
func (m *MeetingRegistrationBase) SetAllowedRegistrant(value *MeetingAudience)() {
    err := m.GetBackingStore().Set("allowedRegistrant", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrants sets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) SetRegistrants(value []MeetingRegistrantBaseable)() {
    err := m.GetBackingStore().Set("registrants", value)
    if err != nil {
        panic(err)
    }
}
// MeetingRegistrationBaseable 
type MeetingRegistrationBaseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedRegistrant()(*MeetingAudience)
    GetRegistrants()([]MeetingRegistrantBaseable)
    SetAllowedRegistrant(value *MeetingAudience)()
    SetRegistrants(value []MeetingRegistrantBaseable)()
}
