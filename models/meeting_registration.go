package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistration 
type MeetingRegistration struct {
    MeetingRegistrationBase
}
// NewMeetingRegistration instantiates a new meetingRegistration and sets the default values.
func NewMeetingRegistration()(*MeetingRegistration) {
    m := &MeetingRegistration{
        MeetingRegistrationBase: *NewMeetingRegistrationBase(),
    }
    odataTypeValue := "#microsoft.graph.meetingRegistration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMeetingRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingRegistration(), nil
}
// GetCustomQuestions gets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) GetCustomQuestions()([]MeetingRegistrationQuestionable) {
    val, err := m.GetBackingStore().Get("customQuestions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingRegistrationQuestionable)
    }
    return nil
}
// GetDescription gets the description property value. The description of the meeting.
func (m *MeetingRegistration) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrationBase.GetFieldDeserializers()
    res["customQuestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingRegistrationQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingRegistrationQuestionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeetingRegistrationQuestionable)
                }
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["registrationPageViewCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageViewCount(val)
        }
        return nil
    }
    res["registrationPageWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageWebUrl(val)
        }
        return nil
    }
    res["speakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingSpeakerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingSpeakerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeetingSpeakerable)
                }
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingRegistration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegistrationPageViewCount gets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) GetRegistrationPageViewCount()(*int32) {
    val, err := m.GetBackingStore().Get("registrationPageViewCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRegistrationPageWebUrl gets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) GetRegistrationPageWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("registrationPageWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSpeakers gets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) GetSpeakers()([]MeetingSpeakerable) {
    val, err := m.GetBackingStore().Get("speakers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingSpeakerable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSubject gets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MeetingRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrationBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
func (m *MeetingRegistration) SetCustomQuestions(value []MeetingRegistrationQuestionable)() {
    err := m.GetBackingStore().Set("customQuestions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the meeting.
func (m *MeetingRegistration) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingRegistration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationPageViewCount sets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) SetRegistrationPageViewCount(value *int32)() {
    err := m.GetBackingStore().Set("registrationPageViewCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationPageWebUrl sets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) SetRegistrationPageWebUrl(value *string)() {
    err := m.GetBackingStore().Set("registrationPageWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeakers sets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) SetSpeakers(value []MeetingSpeakerable)() {
    err := m.GetBackingStore().Set("speakers", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// MeetingRegistrationable 
type MeetingRegistrationable interface {
    MeetingRegistrationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomQuestions()([]MeetingRegistrationQuestionable)
    GetDescription()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetRegistrationPageViewCount()(*int32)
    GetRegistrationPageWebUrl()(*string)
    GetSpeakers()([]MeetingSpeakerable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    SetCustomQuestions(value []MeetingRegistrationQuestionable)()
    SetDescription(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetRegistrationPageViewCount(value *int32)()
    SetRegistrationPageWebUrl(value *string)()
    SetSpeakers(value []MeetingSpeakerable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
}
