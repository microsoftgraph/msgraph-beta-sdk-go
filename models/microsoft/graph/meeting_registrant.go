package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingRegistrant 
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
// NewMeetingRegistrant instantiates a new meetingRegistrant and sets the default values.
func NewMeetingRegistrant()(*MeetingRegistrant) {
    m := &MeetingRegistrant{
        Entity: *NewEntity(),
    }
    return m
}
// GetCustomQuestionAnswers gets the customQuestionAnswers property value. The registrant's answer to custom questions.
func (m *MeetingRegistrant) GetCustomQuestionAnswers()([]CustomQuestionAnswer) {
    if m == nil {
        return nil
    } else {
        return m.customQuestionAnswers
    }
}
// GetEmail gets the email property value. The email address of the registrant.
func (m *MeetingRegistrant) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetFirstName gets the firstName property value. The first name of the registrant.
func (m *MeetingRegistrant) GetFirstName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firstName
    }
}
// GetJoinWebUrl gets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrant) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// GetLastName gets the lastName property value. The last name of the registrant.
func (m *MeetingRegistrant) GetLastName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastName
    }
}
// GetRegistrationDateTime gets the registrationDateTime property value. Time in UTC when the registrant registers for the meeting. Read-only.
func (m *MeetingRegistrant) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registrationDateTime
    }
}
// GetStatus gets the status property value. The registration status of the registrant. Read-only.
func (m *MeetingRegistrant) GetStatus()(*MeetingRegistrantStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customQuestionAnswers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomQuestionAnswer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomQuestionAnswer, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomQuestionAnswer))
            }
            m.SetCustomQuestionAnswers(res)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
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
    res["lastName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["registrationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingRegistrantStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MeetingRegistrantStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *MeetingRegistrant) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCustomQuestionAnswers sets the customQuestionAnswers property value. The registrant's answer to custom questions.
func (m *MeetingRegistrant) SetCustomQuestionAnswers(value []CustomQuestionAnswer)() {
    m.customQuestionAnswers = value
}
// SetEmail sets the email property value. The email address of the registrant.
func (m *MeetingRegistrant) SetEmail(value *string)() {
    m.email = value
}
// SetFirstName sets the firstName property value. The first name of the registrant.
func (m *MeetingRegistrant) SetFirstName(value *string)() {
    m.firstName = value
}
// SetJoinWebUrl sets the joinWebUrl property value. A unique web URL for the registrant to join the meeting. Read-only.
func (m *MeetingRegistrant) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetLastName sets the lastName property value. The last name of the registrant.
func (m *MeetingRegistrant) SetLastName(value *string)() {
    m.lastName = value
}
// SetRegistrationDateTime sets the registrationDateTime property value. Time in UTC when the registrant registers for the meeting. Read-only.
func (m *MeetingRegistrant) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// SetStatus sets the status property value. The registration status of the registrant. Read-only.
func (m *MeetingRegistrant) SetStatus(value *MeetingRegistrantStatus)() {
    m.status = value
}
