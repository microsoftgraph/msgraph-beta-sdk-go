package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrant 
type MeetingRegistrant struct {
    MeetingRegistrantBase
    // The registrant's answer to custom questions.
    customQuestionAnswers []CustomQuestionAnswerable
    // The email address of the registrant.
    email *string
    // The first name of the registrant.
    firstName *string
    // The last name of the registrant.
    lastName *string
    // Time in UTC when the registrant registers for the meeting. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The registration status of the registrant. Read-only.
    status *MeetingRegistrantStatus
}
// NewMeetingRegistrant instantiates a new MeetingRegistrant and sets the default values.
func NewMeetingRegistrant()(*MeetingRegistrant) {
    m := &MeetingRegistrant{
        MeetingRegistrantBase: *NewMeetingRegistrantBase(),
    }
    odataTypeValue := "#microsoft.graph.meetingRegistrant";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMeetingRegistrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingRegistrant(), nil
}
// GetCustomQuestionAnswers gets the customQuestionAnswers property value. The registrant's answer to custom questions.
func (m *MeetingRegistrant) GetCustomQuestionAnswers()([]CustomQuestionAnswerable) {
    return m.customQuestionAnswers
}
// GetEmail gets the email property value. The email address of the registrant.
func (m *MeetingRegistrant) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrantBase.GetFieldDeserializers()
    res["customQuestionAnswers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCustomQuestionAnswerFromDiscriminatorValue , m.SetCustomQuestionAnswers)
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["firstName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFirstName)
    res["lastName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastName)
    res["registrationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRegistrationDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMeetingRegistrantStatus , m.SetStatus)
    return res
}
// GetFirstName gets the firstName property value. The first name of the registrant.
func (m *MeetingRegistrant) GetFirstName()(*string) {
    return m.firstName
}
// GetLastName gets the lastName property value. The last name of the registrant.
func (m *MeetingRegistrant) GetLastName()(*string) {
    return m.lastName
}
// GetRegistrationDateTime gets the registrationDateTime property value. Time in UTC when the registrant registers for the meeting. Read-only.
func (m *MeetingRegistrant) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.registrationDateTime
}
// GetStatus gets the status property value. The registration status of the registrant. Read-only.
func (m *MeetingRegistrant) GetStatus()(*MeetingRegistrantStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *MeetingRegistrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrantBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomQuestionAnswers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomQuestionAnswers())
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
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomQuestionAnswers sets the customQuestionAnswers property value. The registrant's answer to custom questions.
func (m *MeetingRegistrant) SetCustomQuestionAnswers(value []CustomQuestionAnswerable)() {
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
