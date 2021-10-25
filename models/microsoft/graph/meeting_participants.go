package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingParticipants struct {
    additionalData map[string]interface{};
    attendees []MeetingParticipantInfo;
    contributors []MeetingParticipantInfo;
    organizer *MeetingParticipantInfo;
    producers []MeetingParticipantInfo;
}
func NewMeetingParticipants()(*MeetingParticipants) {
    m := &MeetingParticipants{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MeetingParticipants) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MeetingParticipants) GetAttendees()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
func (m *MeetingParticipants) GetContributors()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.contributors
    }
}
func (m *MeetingParticipants) GetOrganizer()(*MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
func (m *MeetingParticipants) GetProducers()([]MeetingParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.producers
    }
}
func (m *MeetingParticipants) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]MeetingParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingParticipantInfo))
        }
        m.SetAttendees(res)
        return nil
    }
    res["contributors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]MeetingParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingParticipantInfo))
        }
        m.SetContributors(res)
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        m.SetOrganizer(val.(*MeetingParticipantInfo))
        return nil
    }
    res["producers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]MeetingParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingParticipantInfo))
        }
        m.SetProducers(res)
        return nil
    }
    return res
}
func (m *MeetingParticipants) IsNil()(bool) {
    return m == nil
}
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
func (m *MeetingParticipants) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MeetingParticipants) SetAttendees(value []MeetingParticipantInfo)() {
    m.attendees = value
}
func (m *MeetingParticipants) SetContributors(value []MeetingParticipantInfo)() {
    m.contributors = value
}
func (m *MeetingParticipants) SetOrganizer(value *MeetingParticipantInfo)() {
    m.organizer = value
}
func (m *MeetingParticipants) SetProducers(value []MeetingParticipantInfo)() {
    m.producers = value
}
