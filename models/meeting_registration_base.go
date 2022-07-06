package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationBase provides operations to manage the collection of activityStatistics entities.
type MeetingRegistrationBase struct {
    Entity
    // Specifies who can register for the meeting.
    allowedRegistrant *MeetingAudience
    // Registrants of the online meeting.
    registrants []MeetingRegistrantBaseable
    // The type property
    type_escaped *string
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
                mappingStr := *mappingValue
                switch mappingStr {
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
    if m == nil {
        return nil
    } else {
        return m.allowedRegistrant
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetRegistrants gets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) GetRegistrants()([]MeetingRegistrantBaseable) {
    if m == nil {
        return nil
    } else {
        return m.registrants
    }
}
// GetType gets the type property value. The type property
func (m *MeetingRegistrationBase) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
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
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedRegistrant sets the allowedRegistrant property value. Specifies who can register for the meeting.
func (m *MeetingRegistrationBase) SetAllowedRegistrant(value *MeetingAudience)() {
    if m != nil {
        m.allowedRegistrant = value
    }
}
// SetRegistrants sets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) SetRegistrants(value []MeetingRegistrantBaseable)() {
    if m != nil {
        m.registrants = value
    }
}
// SetType sets the type property value. The type property
func (m *MeetingRegistrationBase) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
