package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingParticipants 
type MeetingParticipants struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Information of the meeting attendees.
    attendees []MeetingParticipantInfo;
    // 
    contributors []MeetingParticipantInfo;
    // Information of the meeting organizer.
    organizer *MeetingParticipantInfo;
    // 
    producers []MeetingParticipantInfo;
}
// NewMeetingParticipants instantiates a new meetingParticipants and sets the default values.
func NewMeetingParticipants()(*MeetingParticipants) {
    m := &MeetingParticipants{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipants) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttendees gets the attendees property value. Information of the meeting attendees.
func (m *MeetingParticipants) GetAttendees()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// GetContributors gets the contributors property value. 
func (m *MeetingParticipants) GetContributors()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.contributors
    }
}
// GetOrganizer gets the organizer property value. Information of the meeting organizer.
func (m *MeetingParticipants) GetOrganizer()(*MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// GetProducers gets the producers property value. 
func (m *MeetingParticipants) GetProducers()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.producers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipants) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingParticipantInfo))
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["contributors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingParticipantInfo))
            }
            m.SetContributors(res)
        }
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(*MeetingParticipantInfo))
        }
        return nil
    }
    res["producers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingParticipantInfo))
            }
            m.SetProducers(res)
        }
        return nil
    }
    return res
}
func (m *MeetingParticipants) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingParticipants) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContributors()))
        for i, v := range m.GetContributors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("contributors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProducers()))
        for i, v := range m.GetProducers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("producers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipants) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttendees sets the attendees property value. Information of the meeting attendees.
func (m *MeetingParticipants) SetAttendees(value []MeetingParticipantInfo)() {
    if m != nil {
        m.attendees = value
    }
}
// SetContributors sets the contributors property value. 
func (m *MeetingParticipants) SetContributors(value []MeetingParticipantInfo)() {
    if m != nil {
        m.contributors = value
    }
}
// SetOrganizer sets the organizer property value. Information of the meeting organizer.
func (m *MeetingParticipants) SetOrganizer(value *MeetingParticipantInfo)() {
    if m != nil {
        m.organizer = value
    }
}
// SetProducers sets the producers property value. 
func (m *MeetingParticipants) SetProducers(value []MeetingParticipantInfo)() {
    if m != nil {
        m.producers = value
    }
}
