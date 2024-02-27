package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnlineMeeting struct {
    OnlineMeetingBase
}
// NewOnlineMeeting instantiates a new OnlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        OnlineMeetingBase: *NewOnlineMeetingBase(),
    }
    odataTypeValue := "#microsoft.graph.onlineMeeting"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnlineMeetingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnlineMeetingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeeting(), nil
}
// GetAlternativeRecording gets the alternativeRecording property value. The alternativeRecording property
// returns a []byte when successful
func (m *OnlineMeeting) GetAlternativeRecording()([]byte) {
    val, err := m.GetBackingStore().Get("alternativeRecording")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetAttendeeReport gets the attendeeReport property value. The attendeeReport property
// returns a []byte when successful
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    val, err := m.GetBackingStore().Get("attendeeReport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetBroadcastRecording gets the broadcastRecording property value. The broadcastRecording property
// returns a []byte when successful
func (m *OnlineMeeting) GetBroadcastRecording()([]byte) {
    val, err := m.GetBackingStore().Get("broadcastRecording")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetBroadcastSettings gets the broadcastSettings property value. The broadcastSettings property
// returns a BroadcastMeetingSettingsable when successful
func (m *OnlineMeeting) GetBroadcastSettings()(BroadcastMeetingSettingsable) {
    val, err := m.GetBackingStore().Get("broadcastSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BroadcastMeetingSettingsable)
    }
    return nil
}
// GetCapabilities gets the capabilities property value. The capabilities property
// returns a []MeetingCapabilities when successful
func (m *OnlineMeeting) GetCapabilities()([]MeetingCapabilities) {
    val, err := m.GetBackingStore().Get("capabilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingCapabilities)
    }
    return nil
}
// GetCreationDateTime gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
// returns a *Time when successful
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("creationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
// returns a *Time when successful
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExternalId gets the externalId property value. The external ID. A custom ID. Optional.
// returns a *string when successful
func (m *OnlineMeeting) GetExternalId()(*string) {
    val, err := m.GetBackingStore().Get("externalId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnlineMeetingBase.GetFieldDeserializers()
    res["alternativeRecording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeRecording(val)
        }
        return nil
    }
    res["attendeeReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttendeeReport(val)
        }
        return nil
    }
    res["broadcastRecording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBroadcastRecording(val)
        }
        return nil
    }
    res["broadcastSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBroadcastMeetingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBroadcastSettings(val.(BroadcastMeetingSettingsable))
        }
        return nil
    }
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseMeetingCapabilities)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingCapabilities, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*MeetingCapabilities))
                }
            }
            m.SetCapabilities(res)
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["isBroadcast"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBroadcast(val)
        }
        return nil
    }
    res["joinUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinUrl(val)
        }
        return nil
    }
    res["meetingAttendanceReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingAttendanceReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingAttendanceReport(val.(MeetingAttendanceReportable))
        }
        return nil
    }
    res["meetingTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingTemplateId(val)
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingParticipantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(MeetingParticipantsable))
        }
        return nil
    }
    res["recording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecording(val)
        }
        return nil
    }
    res["recordings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallRecordingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CallRecordingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CallRecordingable)
                }
            }
            m.SetRecordings(res)
        }
        return nil
    }
    res["registration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistration(val.(MeetingRegistrationable))
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
    res["transcripts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallTranscriptFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CallTranscriptable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CallTranscriptable)
                }
            }
            m.SetTranscripts(res)
        }
        return nil
    }
    return res
}
// GetIsBroadcast gets the isBroadcast property value. The isBroadcast property
// returns a *bool when successful
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    val, err := m.GetBackingStore().Get("isBroadcast")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJoinUrl gets the joinUrl property value. The joinUrl property
// returns a *string when successful
func (m *OnlineMeeting) GetJoinUrl()(*string) {
    val, err := m.GetBackingStore().Get("joinUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMeetingAttendanceReport gets the meetingAttendanceReport property value. The meetingAttendanceReport property
// returns a MeetingAttendanceReportable when successful
func (m *OnlineMeeting) GetMeetingAttendanceReport()(MeetingAttendanceReportable) {
    val, err := m.GetBackingStore().Get("meetingAttendanceReport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MeetingAttendanceReportable)
    }
    return nil
}
// GetMeetingTemplateId gets the meetingTemplateId property value. The ID of the meeting template.
// returns a *string when successful
func (m *OnlineMeeting) GetMeetingTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("meetingTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParticipants gets the participants property value. The participants associated with the online meeting, including the organizer and the attendees.
// returns a MeetingParticipantsable when successful
func (m *OnlineMeeting) GetParticipants()(MeetingParticipantsable) {
    val, err := m.GetBackingStore().Get("participants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MeetingParticipantsable)
    }
    return nil
}
// GetRecording gets the recording property value. The recording property
// returns a []byte when successful
func (m *OnlineMeeting) GetRecording()([]byte) {
    val, err := m.GetBackingStore().Get("recording")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetRecordings gets the recordings property value. The recordings of an online meeting. Read-only.
// returns a []CallRecordingable when successful
func (m *OnlineMeeting) GetRecordings()([]CallRecordingable) {
    val, err := m.GetBackingStore().Get("recordings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CallRecordingable)
    }
    return nil
}
// GetRegistration gets the registration property value. The registration that is enabled for an online meeting. One online meeting can only have one registration enabled.
// returns a MeetingRegistrationable when successful
func (m *OnlineMeeting) GetRegistration()(MeetingRegistrationable) {
    val, err := m.GetBackingStore().Get("registration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MeetingRegistrationable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
// returns a *Time when successful
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTranscripts gets the transcripts property value. The transcripts of an online meeting. Read-only.
// returns a []CallTranscriptable when successful
func (m *OnlineMeeting) GetTranscripts()([]CallTranscriptable) {
    val, err := m.GetBackingStore().Get("transcripts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CallTranscriptable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnlineMeeting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnlineMeetingBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("alternativeRecording", m.GetAlternativeRecording())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("attendeeReport", m.GetAttendeeReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("broadcastRecording", m.GetBroadcastRecording())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("broadcastSettings", m.GetBroadcastSettings())
        if err != nil {
            return err
        }
    }
    if m.GetCapabilities() != nil {
        err = writer.WriteCollectionOfStringValues("capabilities", SerializeMeetingCapabilities(m.GetCapabilities()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBroadcast", m.GetIsBroadcast())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinUrl", m.GetJoinUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingAttendanceReport", m.GetMeetingAttendanceReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("meetingTemplateId", m.GetMeetingTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("recording", m.GetRecording())
        if err != nil {
            return err
        }
    }
    if m.GetRecordings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecordings()))
        for i, v := range m.GetRecordings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("recordings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("registration", m.GetRegistration())
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
    if m.GetTranscripts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTranscripts()))
        for i, v := range m.GetTranscripts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("transcripts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlternativeRecording sets the alternativeRecording property value. The alternativeRecording property
func (m *OnlineMeeting) SetAlternativeRecording(value []byte)() {
    err := m.GetBackingStore().Set("alternativeRecording", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendeeReport sets the attendeeReport property value. The attendeeReport property
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    err := m.GetBackingStore().Set("attendeeReport", value)
    if err != nil {
        panic(err)
    }
}
// SetBroadcastRecording sets the broadcastRecording property value. The broadcastRecording property
func (m *OnlineMeeting) SetBroadcastRecording(value []byte)() {
    err := m.GetBackingStore().Set("broadcastRecording", value)
    if err != nil {
        panic(err)
    }
}
// SetBroadcastSettings sets the broadcastSettings property value. The broadcastSettings property
func (m *OnlineMeeting) SetBroadcastSettings(value BroadcastMeetingSettingsable)() {
    err := m.GetBackingStore().Set("broadcastSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetCapabilities sets the capabilities property value. The capabilities property
func (m *OnlineMeeting) SetCapabilities(value []MeetingCapabilities)() {
    err := m.GetBackingStore().Set("capabilities", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationDateTime sets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("creationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalId sets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) SetExternalId(value *string)() {
    err := m.GetBackingStore().Set("externalId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBroadcast sets the isBroadcast property value. The isBroadcast property
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    err := m.GetBackingStore().Set("isBroadcast", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinUrl sets the joinUrl property value. The joinUrl property
func (m *OnlineMeeting) SetJoinUrl(value *string)() {
    err := m.GetBackingStore().Set("joinUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingAttendanceReport sets the meetingAttendanceReport property value. The meetingAttendanceReport property
func (m *OnlineMeeting) SetMeetingAttendanceReport(value MeetingAttendanceReportable)() {
    err := m.GetBackingStore().Set("meetingAttendanceReport", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingTemplateId sets the meetingTemplateId property value. The ID of the meeting template.
func (m *OnlineMeeting) SetMeetingTemplateId(value *string)() {
    err := m.GetBackingStore().Set("meetingTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetParticipants sets the participants property value. The participants associated with the online meeting, including the organizer and the attendees.
func (m *OnlineMeeting) SetParticipants(value MeetingParticipantsable)() {
    err := m.GetBackingStore().Set("participants", value)
    if err != nil {
        panic(err)
    }
}
// SetRecording sets the recording property value. The recording property
func (m *OnlineMeeting) SetRecording(value []byte)() {
    err := m.GetBackingStore().Set("recording", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordings sets the recordings property value. The recordings of an online meeting. Read-only.
func (m *OnlineMeeting) SetRecordings(value []CallRecordingable)() {
    err := m.GetBackingStore().Set("recordings", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistration sets the registration property value. The registration that is enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) SetRegistration(value MeetingRegistrationable)() {
    err := m.GetBackingStore().Set("registration", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTranscripts sets the transcripts property value. The transcripts of an online meeting. Read-only.
func (m *OnlineMeeting) SetTranscripts(value []CallTranscriptable)() {
    err := m.GetBackingStore().Set("transcripts", value)
    if err != nil {
        panic(err)
    }
}
type OnlineMeetingable interface {
    OnlineMeetingBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternativeRecording()([]byte)
    GetAttendeeReport()([]byte)
    GetBroadcastRecording()([]byte)
    GetBroadcastSettings()(BroadcastMeetingSettingsable)
    GetCapabilities()([]MeetingCapabilities)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetIsBroadcast()(*bool)
    GetJoinUrl()(*string)
    GetMeetingAttendanceReport()(MeetingAttendanceReportable)
    GetMeetingTemplateId()(*string)
    GetParticipants()(MeetingParticipantsable)
    GetRecording()([]byte)
    GetRecordings()([]CallRecordingable)
    GetRegistration()(MeetingRegistrationable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTranscripts()([]CallTranscriptable)
    SetAlternativeRecording(value []byte)()
    SetAttendeeReport(value []byte)()
    SetBroadcastRecording(value []byte)()
    SetBroadcastSettings(value BroadcastMeetingSettingsable)()
    SetCapabilities(value []MeetingCapabilities)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetIsBroadcast(value *bool)()
    SetJoinUrl(value *string)()
    SetMeetingAttendanceReport(value MeetingAttendanceReportable)()
    SetMeetingTemplateId(value *string)()
    SetParticipants(value MeetingParticipantsable)()
    SetRecording(value []byte)()
    SetRecordings(value []CallRecordingable)()
    SetRegistration(value MeetingRegistrationable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTranscripts(value []CallTranscriptable)()
}
