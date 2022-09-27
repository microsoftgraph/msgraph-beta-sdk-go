package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeeting provides operations to manage the collection of accessReviewDecision entities.
type OnlineMeeting struct {
    Entity
    // Indicates whether attendees can turn on their camera.
    allowAttendeeToEnableCamera *bool
    // Indicates whether attendees can turn on their microphone.
    allowAttendeeToEnableMic *bool
    // Specifies who can be a presenter in a meeting.
    allowedPresenters *OnlineMeetingPresenters
    // Indicates if Teams reactions are enabled for the meeting.
    allowTeamworkReactions *bool
    // The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
    alternativeRecording []byte
    // The anonymizeIdentityForRoles property
    anonymizeIdentityForRoles []OnlineMeetingRole
    // The attendance reports of an online meeting. Read-only.
    attendanceReports []MeetingAttendanceReportable
    // The content stream of the attendee report of a Teams live event. Read-only.
    attendeeReport []byte
    // The phone access (dial-in) information for an online meeting. Read-only.
    audioConferencing AudioConferencingable
    // Settings related to a live event.
    broadcastSettings BroadcastMeetingSettingsable
    // The capabilities property
    capabilities []MeetingCapabilities
    // The chat information associated with this online meeting.
    chatInfo ChatInfoable
    // The meeting creation time in UTC. Read-only.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The external ID. A custom ID. Optional.
    externalId *string
    // Indicates whether this is a Teams live event.
    isBroadcast *bool
    // Indicates whether to announce when callers join or leave.
    isEntryExitAnnounced *bool
    // The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
    joinInformation ItemBodyable
    // Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode.
    joinMeetingIdSettings JoinMeetingIdSettingsable
    // The joinUrl property
    joinUrl *string
    // The join URL of the online meeting. Read-only.
    joinWebUrl *string
    // Specifies which participants can bypass the meeting lobby.
    lobbyBypassSettings LobbyBypassSettingsable
    // The meetingAttendanceReport property
    meetingAttendanceReport MeetingAttendanceReportable
    // The participants associated with the online meeting. This includes the organizer and the attendees.
    participants MeetingParticipantsable
    // Indicates whether to record the meeting automatically.
    recordAutomatically *bool
    // The content stream of the recording of a Teams live event. Read-only.
    recording []byte
    // The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
    registration MeetingRegistrationable
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject of the online meeting.
    subject *string
    // The transcripts of an online meeting. Read-only.
    transcripts []CallTranscriptable
    // The video teleconferencing ID. Read-only.
    videoTeleconferenceId *string
    // The virtualAppointment property
    virtualAppointment VirtualAppointmentable
}
// NewOnlineMeeting instantiates a new onlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.onlineMeeting";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnlineMeetingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeeting(), nil
}
// GetAllowAttendeeToEnableCamera gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    return m.allowAttendeeToEnableCamera
}
// GetAllowAttendeeToEnableMic gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    return m.allowAttendeeToEnableMic
}
// GetAllowedPresenters gets the allowedPresenters property value. Specifies who can be a presenter in a meeting.
func (m *OnlineMeeting) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    return m.allowedPresenters
}
// GetAllowTeamworkReactions gets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    return m.allowTeamworkReactions
}
// GetAlternativeRecording gets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) GetAlternativeRecording()([]byte) {
    return m.alternativeRecording
}
// GetAnonymizeIdentityForRoles gets the anonymizeIdentityForRoles property value. The anonymizeIdentityForRoles property
func (m *OnlineMeeting) GetAnonymizeIdentityForRoles()([]OnlineMeetingRole) {
    return m.anonymizeIdentityForRoles
}
// GetAttendanceReports gets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) GetAttendanceReports()([]MeetingAttendanceReportable) {
    return m.attendanceReports
}
// GetAttendeeReport gets the attendeeReport property value. The content stream of the attendee report of a Teams live event. Read-only.
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    return m.attendeeReport
}
// GetAudioConferencing gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) GetAudioConferencing()(AudioConferencingable) {
    return m.audioConferencing
}
// GetBroadcastSettings gets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) GetBroadcastSettings()(BroadcastMeetingSettingsable) {
    return m.broadcastSettings
}
// GetCapabilities gets the capabilities property value. The capabilities property
func (m *OnlineMeeting) GetCapabilities()([]MeetingCapabilities) {
    return m.capabilities
}
// GetChatInfo gets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) GetChatInfo()(ChatInfoable) {
    return m.chatInfo
}
// GetCreationDateTime gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowAttendeeToEnableCamera"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowAttendeeToEnableCamera)
    res["allowAttendeeToEnableMic"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowAttendeeToEnableMic)
    res["allowedPresenters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOnlineMeetingPresenters , m.SetAllowedPresenters)
    res["allowTeamworkReactions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowTeamworkReactions)
    res["alternativeRecording"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetAlternativeRecording)
    res["anonymizeIdentityForRoles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseOnlineMeetingRole , m.SetAnonymizeIdentityForRoles)
    res["attendanceReports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMeetingAttendanceReportFromDiscriminatorValue , m.SetAttendanceReports)
    res["attendeeReport"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetAttendeeReport)
    res["audioConferencing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAudioConferencingFromDiscriminatorValue , m.SetAudioConferencing)
    res["broadcastSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBroadcastMeetingSettingsFromDiscriminatorValue , m.SetBroadcastSettings)
    res["capabilities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfEnumValues(ParseMeetingCapabilities , m.SetCapabilities)
    res["chatInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateChatInfoFromDiscriminatorValue , m.SetChatInfo)
    res["creationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreationDateTime)
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["isBroadcast"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBroadcast)
    res["isEntryExitAnnounced"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsEntryExitAnnounced)
    res["joinInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetJoinInformation)
    res["joinMeetingIdSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateJoinMeetingIdSettingsFromDiscriminatorValue , m.SetJoinMeetingIdSettings)
    res["joinUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJoinUrl)
    res["joinWebUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJoinWebUrl)
    res["lobbyBypassSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue , m.SetLobbyBypassSettings)
    res["meetingAttendanceReport"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMeetingAttendanceReportFromDiscriminatorValue , m.SetMeetingAttendanceReport)
    res["participants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMeetingParticipantsFromDiscriminatorValue , m.SetParticipants)
    res["recordAutomatically"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRecordAutomatically)
    res["recording"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetRecording)
    res["registration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMeetingRegistrationFromDiscriminatorValue , m.SetRegistration)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    res["transcripts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCallTranscriptFromDiscriminatorValue , m.SetTranscripts)
    res["videoTeleconferenceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVideoTeleconferenceId)
    res["virtualAppointment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateVirtualAppointmentFromDiscriminatorValue , m.SetVirtualAppointment)
    return res
}
// GetIsBroadcast gets the isBroadcast property value. Indicates whether this is a Teams live event.
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    return m.isBroadcast
}
// GetIsEntryExitAnnounced gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    return m.isEntryExitAnnounced
}
// GetJoinInformation gets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
func (m *OnlineMeeting) GetJoinInformation()(ItemBodyable) {
    return m.joinInformation
}
// GetJoinMeetingIdSettings gets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode.
func (m *OnlineMeeting) GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable) {
    return m.joinMeetingIdSettings
}
// GetJoinUrl gets the joinUrl property value. The joinUrl property
func (m *OnlineMeeting) GetJoinUrl()(*string) {
    return m.joinUrl
}
// GetJoinWebUrl gets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
func (m *OnlineMeeting) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    return m.lobbyBypassSettings
}
// GetMeetingAttendanceReport gets the meetingAttendanceReport property value. The meetingAttendanceReport property
func (m *OnlineMeeting) GetMeetingAttendanceReport()(MeetingAttendanceReportable) {
    return m.meetingAttendanceReport
}
// GetParticipants gets the participants property value. The participants associated with the online meeting. This includes the organizer and the attendees.
func (m *OnlineMeeting) GetParticipants()(MeetingParticipantsable) {
    return m.participants
}
// GetRecordAutomatically gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) GetRecordAutomatically()(*bool) {
    return m.recordAutomatically
}
// GetRecording gets the recording property value. The content stream of the recording of a Teams live event. Read-only.
func (m *OnlineMeeting) GetRecording()([]byte) {
    return m.recording
}
// GetRegistration gets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) GetRegistration()(MeetingRegistrationable) {
    return m.registration
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) GetSubject()(*string) {
    return m.subject
}
// GetTranscripts gets the transcripts property value. The transcripts of an online meeting. Read-only.
func (m *OnlineMeeting) GetTranscripts()([]CallTranscriptable) {
    return m.transcripts
}
// GetVideoTeleconferenceId gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    return m.videoTeleconferenceId
}
// GetVirtualAppointment gets the virtualAppointment property value. The virtualAppointment property
func (m *OnlineMeeting) GetVirtualAppointment()(VirtualAppointmentable) {
    return m.virtualAppointment
}
// Serialize serializes information the current object
func (m *OnlineMeeting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableCamera", m.GetAllowAttendeeToEnableCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableMic", m.GetAllowAttendeeToEnableMic())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedPresenters() != nil {
        cast := (*m.GetAllowedPresenters()).String()
        err = writer.WriteStringValue("allowedPresenters", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTeamworkReactions", m.GetAllowTeamworkReactions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("alternativeRecording", m.GetAlternativeRecording())
        if err != nil {
            return err
        }
    }
    if m.GetAnonymizeIdentityForRoles() != nil {
        err = writer.WriteCollectionOfStringValues("anonymizeIdentityForRoles", SerializeOnlineMeetingRole(m.GetAnonymizeIdentityForRoles()))
        if err != nil {
            return err
        }
    }
    if m.GetAttendanceReports() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttendanceReports())
        err = writer.WriteCollectionOfObjectValues("attendanceReports", cast)
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
        err = writer.WriteObjectValue("audioConferencing", m.GetAudioConferencing())
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
        err = writer.WriteObjectValue("chatInfo", m.GetChatInfo())
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
        err = writer.WriteBoolValue("isEntryExitAnnounced", m.GetIsEntryExitAnnounced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinInformation", m.GetJoinInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinMeetingIdSettings", m.GetJoinMeetingIdSettings())
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
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lobbyBypassSettings", m.GetLobbyBypassSettings())
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
        err = writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordAutomatically", m.GetRecordAutomatically())
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
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetTranscripts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTranscripts())
        err = writer.WriteCollectionOfObjectValues("transcripts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoTeleconferenceId", m.GetVideoTeleconferenceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("virtualAppointment", m.GetVirtualAppointment())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAttendeeToEnableCamera sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    m.allowAttendeeToEnableCamera = value
}
// SetAllowAttendeeToEnableMic sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    m.allowAttendeeToEnableMic = value
}
// SetAllowedPresenters sets the allowedPresenters property value. Specifies who can be a presenter in a meeting.
func (m *OnlineMeeting) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    m.allowedPresenters = value
}
// SetAllowTeamworkReactions sets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    m.allowTeamworkReactions = value
}
// SetAlternativeRecording sets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) SetAlternativeRecording(value []byte)() {
    m.alternativeRecording = value
}
// SetAnonymizeIdentityForRoles sets the anonymizeIdentityForRoles property value. The anonymizeIdentityForRoles property
func (m *OnlineMeeting) SetAnonymizeIdentityForRoles(value []OnlineMeetingRole)() {
    m.anonymizeIdentityForRoles = value
}
// SetAttendanceReports sets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) SetAttendanceReports(value []MeetingAttendanceReportable)() {
    m.attendanceReports = value
}
// SetAttendeeReport sets the attendeeReport property value. The content stream of the attendee report of a Teams live event. Read-only.
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    m.attendeeReport = value
}
// SetAudioConferencing sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) SetAudioConferencing(value AudioConferencingable)() {
    m.audioConferencing = value
}
// SetBroadcastSettings sets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) SetBroadcastSettings(value BroadcastMeetingSettingsable)() {
    m.broadcastSettings = value
}
// SetCapabilities sets the capabilities property value. The capabilities property
func (m *OnlineMeeting) SetCapabilities(value []MeetingCapabilities)() {
    m.capabilities = value
}
// SetChatInfo sets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) SetChatInfo(value ChatInfoable)() {
    m.chatInfo = value
}
// SetCreationDateTime sets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) SetExternalId(value *string)() {
    m.externalId = value
}
// SetIsBroadcast sets the isBroadcast property value. Indicates whether this is a Teams live event.
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    m.isBroadcast = value
}
// SetIsEntryExitAnnounced sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    m.isEntryExitAnnounced = value
}
// SetJoinInformation sets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
func (m *OnlineMeeting) SetJoinInformation(value ItemBodyable)() {
    m.joinInformation = value
}
// SetJoinMeetingIdSettings sets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode.
func (m *OnlineMeeting) SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)() {
    m.joinMeetingIdSettings = value
}
// SetJoinUrl sets the joinUrl property value. The joinUrl property
func (m *OnlineMeeting) SetJoinUrl(value *string)() {
    m.joinUrl = value
}
// SetJoinWebUrl sets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
func (m *OnlineMeeting) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    m.lobbyBypassSettings = value
}
// SetMeetingAttendanceReport sets the meetingAttendanceReport property value. The meetingAttendanceReport property
func (m *OnlineMeeting) SetMeetingAttendanceReport(value MeetingAttendanceReportable)() {
    m.meetingAttendanceReport = value
}
// SetParticipants sets the participants property value. The participants associated with the online meeting. This includes the organizer and the attendees.
func (m *OnlineMeeting) SetParticipants(value MeetingParticipantsable)() {
    m.participants = value
}
// SetRecordAutomatically sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) SetRecordAutomatically(value *bool)() {
    m.recordAutomatically = value
}
// SetRecording sets the recording property value. The content stream of the recording of a Teams live event. Read-only.
func (m *OnlineMeeting) SetRecording(value []byte)() {
    m.recording = value
}
// SetRegistration sets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) SetRegistration(value MeetingRegistrationable)() {
    m.registration = value
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) SetSubject(value *string)() {
    m.subject = value
}
// SetTranscripts sets the transcripts property value. The transcripts of an online meeting. Read-only.
func (m *OnlineMeeting) SetTranscripts(value []CallTranscriptable)() {
    m.transcripts = value
}
// SetVideoTeleconferenceId sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    m.videoTeleconferenceId = value
}
// SetVirtualAppointment sets the virtualAppointment property value. The virtualAppointment property
func (m *OnlineMeeting) SetVirtualAppointment(value VirtualAppointmentable)() {
    m.virtualAppointment = value
}
