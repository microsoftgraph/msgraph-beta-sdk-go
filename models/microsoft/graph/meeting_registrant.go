package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MeetingRegistrant struct {
    Entity
    // The registrant's answer to custom questions.
    customQuestionAnswers []CustomQuestionAnswer;
    // The email address of the registrant.
    email *string;
    // The first name of the registrant.
    firstName *string;
    // A unique web URL for the registrant to join the meeting. Read-only.
    joinWebUrl *string;
    // The last name of the registrant.
    lastName *string;
    // Time in UTC when the registrant registers for the meeting. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The registration status of the registrant. Read-only.
    status *MeetingRegistrantStatus;
}
// Instantiates a new meetingRegistrant and sets the default values.
func NewMeetingRegistrant()(*MeetingRegistrant) {
    m := &MeetingRegistrant{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the customQuestionAnswers property value. The registrant's answer to custom questions.
func (m *MeetingRegistrant) GetCustomQuestionAnswers()([]CustomQuestionAnswer) {
    if m == nil {
        return nil
    } else {
        return m.customQuestionAnswers
    }
}
// Gets the email property value. The email address of the registrant.
func (m *MeetingRegistrant) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the firstName property value. The first name of the registrant.
func (m *MeetingRegistrant) GetFirstName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstName
    }
}
// Gets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrant) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// Gets the lastName property value. The last name of the registrant.
func (m *MeetingRegistrant) GetLastName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastName
    }
}
// Gets the registrationDateTime property value. Time in UTC when the registrant registers for the meeting. Read-only.
func (m *MeetingRegistrant) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registrationDateTime
    }
}
// Gets the status property value. The registration status of the registrant. Read-only.
func (m *MeetingRegistrant) GetStatus()(*MeetingRegistrantStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *MeetingRegistrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customQuestionAnswers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomQuestionAnswer() })
        if err != nil {
            return err
        }
        res := make([]CustomQuestionAnswer, len(val))
        for i, v := range val {
            res[i] = *(v.(*CustomQuestionAnswer))
        }
        m.SetCustomQuestionAnswers(res)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["firstName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFirstName(val)
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoinWebUrl(val)
        return nil
    }
    res["lastName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastName(val)
        return nil
    }
    res["registrationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRegistrationDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingRegistrantStatus)
        if err != nil {
            return err
        }
        cast := val.(MeetingRegistrantStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *MeetingRegistrant) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MeetingRegistrant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomQuestionAnswers()))
        for i, v := range m.GetCustomQuestionAnswers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestionAnswers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the customQuestionAnswers property value. The registrant's answer to custom questions.
// Parameters:
//  - value : Value to set for the customQuestionAnswers property.
func (m *MeetingRegistrant) SetCustomQuestionAnswers(value []CustomQuestionAnswer)() {
    m.customQuestionAnswers = value
}
// Sets the email property value. The email address of the registrant.
// Parameters:
//  - value : Value to set for the email property.
func (m *MeetingRegistrant) SetEmail(value *string)() {
    m.email = value
}
// Sets the firstName property value. The first name of the registrant.
// Parameters:
//  - value : Value to set for the firstName property.
func (m *MeetingRegistrant) SetFirstName(value *string)() {
    m.firstName = value
}
// Sets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
// Parameters:
//  - value : Value to set for the joinWebUrl property.
func (m *MeetingRegistrant) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// Sets the lastName property value. The last name of the registrant.
// Parameters:
//  - value : Value to set for the lastName property.
func (m *MeetingRegistrant) SetLastName(value *string)() {
    m.lastName = value
}
// Sets the registrationDateTime property value. Time in UTC when the registrant registers for the meeting. Read-only.
// Parameters:
//  - value : Value to set for the registrationDateTime property.
func (m *MeetingRegistrant) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// Sets the status property value. The registration status of the registrant. Read-only.
// Parameters:
//  - value : Value to set for the status property.
func (m *MeetingRegistrant) SetStatus(value *MeetingRegistrantStatus)() {
    m.status = value
}
