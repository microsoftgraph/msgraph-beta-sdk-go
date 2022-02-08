package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistrationBase 
type MeetingRegistrationBase struct {
    Entity
    // Specifies who can register for the meeting.
    allowedRegistrant *MeetingAudience;
    // Registrants of the online meeting.
    registrants []MeetingRegistrantBase;
}
// NewMeetingRegistrationBase instantiates a new meetingRegistrationBase and sets the default values.
func NewMeetingRegistrationBase()(*MeetingRegistrationBase) {
    m := &MeetingRegistrationBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowedRegistrant gets the allowedRegistrant property value. Specifies who can register for the meeting.
func (m *MeetingRegistrationBase) GetAllowedRegistrant()(*MeetingAudience) {
    if m == nil {
        return nil
    } else {
        return m.allowedRegistrant
    }
}
// GetRegistrants gets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) GetRegistrants()([]MeetingRegistrantBase) {
    if m == nil {
        return nil
    } else {
        return m.registrants
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrationBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRegistrant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedRegistrant(val.(*MeetingAudience))
        }
        return nil
    }
    res["registrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingRegistrantBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingRegistrantBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingRegistrantBase))
            }
            m.SetRegistrants(res)
        }
        return nil
    }
    return res
}
func (m *MeetingRegistrationBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingRegistrationBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegistrants()))
        for i, v := range m.GetRegistrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m != nil {
        m.allowedRegistrant = value
    }
}
// SetRegistrants sets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistrationBase) SetRegistrants(value []MeetingRegistrantBase)() {
    if m != nil {
        m.registrants = value
    }
}
