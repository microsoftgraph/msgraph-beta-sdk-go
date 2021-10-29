package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MeetingRegistration struct {
    Entity
    // Specifies who can register for the meeting.
    allowedRegistrant *MeetingAudience;
    // Custom registration questions.
    customQuestions []MeetingRegistrationQuestion;
    // The description of the meeting.
    description *string;
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Registrants of the online meeting.
    registrants []MeetingRegistrant;
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
// Instantiates a new meetingRegistration and sets the default values.
func NewMeetingRegistration()(*MeetingRegistration) {
    m := &MeetingRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowedRegistrant property value. Specifies who can register for the meeting.
func (m *MeetingRegistration) GetAllowedRegistrant()(*MeetingAudience) {
    if m == nil {
        return nil
    } else {
        return m.allowedRegistrant
    }
}
// Gets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) GetCustomQuestions()([]MeetingRegistrationQuestion) {
    if m == nil {
        return nil
    } else {
        return m.customQuestions
    }
}
// Gets the description property value. The description of the meeting.
func (m *MeetingRegistration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the registrants property value. Registrants of the online meeting.
func (m *MeetingRegistration) GetRegistrants()([]MeetingRegistrant) {
    if m == nil {
        return nil
    } else {
        return m.registrants
    }
}
// Gets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) GetRegistrationPageViewCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.registrationPageViewCount
    }
}
// Gets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) GetRegistrationPageWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.registrationPageWebUrl
    }
}
// Gets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) GetSpeakers()([]MeetingSpeaker) {
    if m == nil {
        return nil
    } else {
        return m.speakers
    }
}
// Gets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// The deserialization information for the current model
func (m *MeetingRegistration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRegistrant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingAudience)
        if err != nil {
            return err
        }
        cast := val.(MeetingAudience)
        m.SetAllowedRegistrant(&cast)
        return nil
    }
    res["customQuestions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingRegistrationQuestion() })
        if err != nil {
            return err
        }
        res := make([]MeetingRegistrationQuestion, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingRegistrationQuestion))
        }
        m.SetCustomQuestions(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["registrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingRegistrant() })
        if err != nil {
            return err
        }
        res := make([]MeetingRegistrant, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingRegistrant))
        }
        m.SetRegistrants(res)
        return nil
    }
    res["registrationPageViewCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRegistrationPageViewCount(val)
        return nil
    }
    res["registrationPageWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRegistrationPageWebUrl(val)
        return nil
    }
    res["speakers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingSpeaker() })
        if err != nil {
            return err
        }
        res := make([]MeetingSpeaker, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingSpeaker))
        }
        m.SetSpeakers(res)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    return res
}
func (m *MeetingRegistration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MeetingRegistration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedRegistrant() != nil {
        cast := m.GetAllowedRegistrant().String()
        err = writer.WriteStringValue("allowedRegistrant", &cast)
        if err != nil {
            return err
        }
    }
    {
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
    {
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
// Sets the allowedRegistrant property value. Specifies who can register for the meeting.
// Parameters:
//  - value : Value to set for the allowedRegistrant property.
func (m *MeetingRegistration) SetAllowedRegistrant(value *MeetingAudience)() {
    m.allowedRegistrant = value
}
// Sets the customQuestions property value. Custom registration questions.
// Parameters:
//  - value : Value to set for the customQuestions property.
func (m *MeetingRegistration) SetCustomQuestions(value []MeetingRegistrationQuestion)() {
    m.customQuestions = value
}
// Sets the description property value. The description of the meeting.
// Parameters:
//  - value : Value to set for the description property.
func (m *MeetingRegistration) SetDescription(value *string)() {
    m.description = value
}
// Sets the endDateTime property value. The meeting end time in UTC.
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *MeetingRegistration) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the registrants property value. Registrants of the online meeting.
// Parameters:
//  - value : Value to set for the registrants property.
func (m *MeetingRegistration) SetRegistrants(value []MeetingRegistrant)() {
    m.registrants = value
}
// Sets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
// Parameters:
//  - value : Value to set for the registrationPageViewCount property.
func (m *MeetingRegistration) SetRegistrationPageViewCount(value *int32)() {
    m.registrationPageViewCount = value
}
// Sets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
// Parameters:
//  - value : Value to set for the registrationPageWebUrl property.
func (m *MeetingRegistration) SetRegistrationPageWebUrl(value *string)() {
    m.registrationPageWebUrl = value
}
// Sets the speakers property value. The meeting speaker's information.
// Parameters:
//  - value : Value to set for the speakers property.
func (m *MeetingRegistration) SetSpeakers(value []MeetingSpeaker)() {
    m.speakers = value
}
// Sets the startDateTime property value. The meeting start time in UTC.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *MeetingRegistration) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the subject property value. The subject of the meeting.
// Parameters:
//  - value : Value to set for the subject property.
func (m *MeetingRegistration) SetSubject(value *string)() {
    m.subject = value
}