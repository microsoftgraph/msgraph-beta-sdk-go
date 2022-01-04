package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistrantBase 
type MeetingRegistrantBase struct {
    Entity
    // A unique web URL for the registrant to join the meeting. Read-only.
    joinWebUrl *string;
}
// NewMeetingRegistrantBase instantiates a new meetingRegistrantBase and sets the default values.
func NewMeetingRegistrantBase()(*MeetingRegistrantBase) {
    m := &MeetingRegistrantBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetJoinWebUrl gets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrantBase) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrantBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *MeetingRegistrantBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingRegistrantBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetJoinWebUrl sets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrantBase) SetJoinWebUrl(value *string)() {
    if m != nil {
        m.joinWebUrl = value
    }
}
