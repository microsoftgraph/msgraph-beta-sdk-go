package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistration 
type MeetingRegistration struct {
    MeetingRegistrationBase
    // Custom registration questions.
    customQuestions []MeetingRegistrationQuestion;
    // The description of the meeting.
    description *string;
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The number of times the registration page has been visited. Read-only.
    registrationPageViewCount *int32;
    // The URL of the registration page. Read-only.
    registrationPageWebUrl *string;
    // The meeting speaker's information.
    speakers []MeetingSpeaker;
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The subject of the meeting.
    subject *string;
}
// NewMeetingRegistration instantiates a new meetingRegistration and sets the default values.
func NewMeetingRegistration()(*MeetingRegistration) {
    m := &MeetingRegistration{
        MeetingRegistrationBase: *NewMeetingRegistrationBase(),
    }
    return m
}
// GetCustomQuestions gets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) GetCustomQuestions()([]MeetingRegistrationQuestion) {
    if m == nil {
        return nil
    } else {
        return m.customQuestions
    }
}
// GetDescription gets the description property value. The description of the meeting.
func (m *MeetingRegistration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetRegistrationPageViewCount gets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) GetRegistrationPageViewCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.registrationPageViewCount
    }
}
// GetRegistrationPageWebUrl gets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) GetRegistrationPageWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registrationPageWebUrl
    }
}
// GetSpeakers gets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) GetSpeakers()([]MeetingSpeaker) {
    if m == nil {
        return nil
    } else {
        return m.speakers
    }
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetSubject gets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.MeetingRegistrationBase.GetFieldDeserializers()
    res["customQuestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingRegistrationQuestion() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingRegistrationQuestion, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingRegistrationQuestion))
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["registrationPageViewCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageViewCount(val)
        }
        return nil
    }
    res["registrationPageWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageWebUrl(val)
        }
        return nil
    }
    res["speakers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingSpeaker() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingSpeaker, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingSpeaker))
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
func (m *MeetingRegistration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingRegistration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.MeetingRegistrationBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("registrationPageViewCount", m.GetRegistrationPageViewCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registrationPageWebUrl", m.GetRegistrationPageWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSpeakers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("speakers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomQuestions sets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) SetCustomQuestions(value []MeetingRegistrationQuestion)() {
    if m != nil {
        m.customQuestions = value
    }
}
// SetDescription sets the description property value. The description of the meeting.
func (m *MeetingRegistration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetRegistrationPageViewCount sets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) SetRegistrationPageViewCount(value *int32)() {
    if m != nil {
        m.registrationPageViewCount = value
    }
}
// SetRegistrationPageWebUrl sets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) SetRegistrationPageWebUrl(value *string)() {
    if m != nil {
        m.registrationPageWebUrl = value
    }
}
// SetSpeakers sets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) SetSpeakers(value []MeetingSpeaker)() {
    if m != nil {
        m.speakers = value
    }
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetSubject sets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
