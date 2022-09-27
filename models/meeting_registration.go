package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistration 
type MeetingRegistration struct {
    MeetingRegistrationBase
    // Custom registration questions.
    customQuestions []MeetingRegistrationQuestionable
    // The description of the meeting.
    description *string
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of times the registration page has been visited. Read-only.
    registrationPageViewCount *int32
    // The URL of the registration page. Read-only.
    registrationPageWebUrl *string
    // The meeting speaker's information.
    speakers []MeetingSpeakerable
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject of the meeting.
    subject *string
}
// NewMeetingRegistration instantiates a new MeetingRegistration and sets the default values.
func NewMeetingRegistration()(*MeetingRegistration) {
    m := &MeetingRegistration{
        MeetingRegistrationBase: *NewMeetingRegistrationBase(),
    }
    odataTypeValue := "#microsoft.graph.meetingRegistration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMeetingRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingRegistration(), nil
}
// GetCustomQuestions gets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) GetCustomQuestions()([]MeetingRegistrationQuestionable) {
    return m.customQuestions
}
// GetDescription gets the description property value. The description of the meeting.
func (m *MeetingRegistration) GetDescription()(*string) {
    return m.description
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrationBase.GetFieldDeserializers()
    res["customQuestions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMeetingRegistrationQuestionFromDiscriminatorValue , m.SetCustomQuestions)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["registrationPageViewCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRegistrationPageViewCount)
    res["registrationPageWebUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistrationPageWebUrl)
    res["speakers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMeetingSpeakerFromDiscriminatorValue , m.SetSpeakers)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    return res
}
// GetRegistrationPageViewCount gets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) GetRegistrationPageViewCount()(*int32) {
    return m.registrationPageViewCount
}
// GetRegistrationPageWebUrl gets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) GetRegistrationPageWebUrl()(*string) {
    return m.registrationPageWebUrl
}
// GetSpeakers gets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) GetSpeakers()([]MeetingSpeakerable) {
    return m.speakers
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *MeetingRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrationBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomQuestions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCustomQuestions())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSpeakers())
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
func (m *MeetingRegistration) SetCustomQuestions(value []MeetingRegistrationQuestionable)() {
    m.customQuestions = value
}
// SetDescription sets the description property value. The description of the meeting.
func (m *MeetingRegistration) SetDescription(value *string)() {
    m.description = value
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetRegistrationPageViewCount sets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) SetRegistrationPageViewCount(value *int32)() {
    m.registrationPageViewCount = value
}
// SetRegistrationPageWebUrl sets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) SetRegistrationPageWebUrl(value *string)() {
    m.registrationPageWebUrl = value
}
// SetSpeakers sets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) SetSpeakers(value []MeetingSpeakerable)() {
    m.speakers = value
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) SetSubject(value *string)() {
    m.subject = value
}
