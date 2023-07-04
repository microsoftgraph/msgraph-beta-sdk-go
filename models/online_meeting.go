package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeeting 
type OnlineMeeting struct {
    Entity
}
// NewOnlineMeeting instantiates a new onlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnlineMeetingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.virtualEventSession":
                        return NewVirtualEventSession(), nil
                }
            }
        }
    }
    return NewOnlineMeeting(), nil
}
// GetAllowAttendeeToEnableCamera gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    val, err := m.GetBackingStore().Get("allowAttendeeToEnableCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowAttendeeToEnableMic gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    val, err := m.GetBackingStore().Get("allowAttendeeToEnableMic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowedPresenters gets the allowedPresenters property value. Specifies who can be a presenter in a meeting.
func (m *OnlineMeeting) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    val, err := m.GetBackingStore().Get("allowedPresenters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OnlineMeetingPresenters)
    }
    return nil
}
// GetAllowMeetingChat gets the allowMeetingChat property value. Specifies the mode of meeting chat.
func (m *OnlineMeeting) GetAllowMeetingChat()(*MeetingChatMode) {
    val, err := m.GetBackingStore().Get("allowMeetingChat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingChatMode)
    }
    return nil
}
// GetAllowParticipantsToChangeName gets the allowParticipantsToChangeName property value. Specifies if participants are allowed to rename themselves in an instance of the meeting.
func (m *OnlineMeeting) GetAllowParticipantsToChangeName()(*bool) {
    val, err := m.GetBackingStore().Get("allowParticipantsToChangeName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowRecording gets the allowRecording property value. Indicates whether recording is enabled for the meeting.
func (m *OnlineMeeting) GetAllowRecording()(*bool) {
    val, err := m.GetBackingStore().Get("allowRecording")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowTeamworkReactions gets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    val, err := m.GetBackingStore().Get("allowTeamworkReactions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowTranscription gets the allowTranscription property value. Indicates whether transcription is enabled for the meeting.
func (m *OnlineMeeting) GetAllowTranscription()(*bool) {
    val, err := m.GetBackingStore().Get("allowTranscription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAlternativeRecording gets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
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
// GetAnonymizeIdentityForRoles gets the anonymizeIdentityForRoles property value. Specifies whose identity will be anonymized in the meeting. Possible values are: attendee. The attendee value cannot be removed through a PATCH operation once added.
func (m *OnlineMeeting) GetAnonymizeIdentityForRoles()([]OnlineMeetingRole) {
    val, err := m.GetBackingStore().Get("anonymizeIdentityForRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OnlineMeetingRole)
    }
    return nil
}
// GetAttendanceReports gets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) GetAttendanceReports()([]MeetingAttendanceReportable) {
    val, err := m.GetBackingStore().Get("attendanceReports")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MeetingAttendanceReportable)
    }
    return nil
}
// GetAttendeeReport gets the attendeeReport property value. The content stream of the attendee report of a Teams live event. Read-only.
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
// GetAudioConferencing gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) GetAudioConferencing()(AudioConferencingable) {
    val, err := m.GetBackingStore().Get("audioConferencing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AudioConferencingable)
    }
    return nil
}
// GetBroadcastRecording gets the broadcastRecording property value. The broadcastRecording property
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
// GetBroadcastSettings gets the broadcastSettings property value. Settings related to a live event.
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
// GetChatInfo gets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) GetChatInfo()(ChatInfoable) {
    val, err := m.GetBackingStore().Get("chatInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ChatInfoable)
    }
    return nil
}
// GetCreationDateTime gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
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
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowAttendeeToEnableCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableCamera(val)
        }
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableMic(val)
        }
        return nil
    }
    res["allowedPresenters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedPresenters(val.(*OnlineMeetingPresenters))
        }
        return nil
    }
    res["allowMeetingChat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMeetingChat(val.(*MeetingChatMode))
        }
        return nil
    }
    res["allowParticipantsToChangeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowParticipantsToChangeName(val)
        }
        return nil
    }
    res["allowRecording"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowRecording(val)
        }
        return nil
    }
    res["allowTeamworkReactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamworkReactions(val)
        }
        return nil
    }
    res["allowTranscription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTranscription(val)
        }
        return nil
    }
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
    res["anonymizeIdentityForRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnlineMeetingRole)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingRole, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*OnlineMeetingRole))
                }
            }
            m.SetAnonymizeIdentityForRoles(res)
        }
        return nil
    }
    res["attendanceReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingAttendanceReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingAttendanceReportable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MeetingAttendanceReportable)
                }
            }
            m.SetAttendanceReports(res)
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
    res["audioConferencing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudioConferencingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioConferencing(val.(AudioConferencingable))
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
    res["chatInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(ChatInfoable))
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
    res["isEntryExitAnnounced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEntryExitAnnounced(val)
        }
        return nil
    }
    res["joinInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinInformation(val.(ItemBodyable))
        }
        return nil
    }
    res["joinMeetingIdSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJoinMeetingIdSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingIdSettings(val.(JoinMeetingIdSettingsable))
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
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["lobbyBypassSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(LobbyBypassSettingsable))
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
    res["recordAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordAutomatically(val)
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
    res["shareMeetingChatHistoryDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatHistoryDefaultMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareMeetingChatHistoryDefault(val.(*MeetingChatHistoryDefaultMode))
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
    res["videoTeleconferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoTeleconferenceId(val)
        }
        return nil
    }
    res["virtualAppointment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualAppointmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualAppointment(val.(VirtualAppointmentable))
        }
        return nil
    }
    res["watermarkProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWatermarkProtectionValuesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatermarkProtection(val.(WatermarkProtectionValuesable))
        }
        return nil
    }
    return res
}
// GetIsBroadcast gets the isBroadcast property value. Indicates whether this is a Teams live event.
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
// GetIsEntryExitAnnounced gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    val, err := m.GetBackingStore().Get("isEntryExitAnnounced")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJoinInformation gets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
func (m *OnlineMeeting) GetJoinInformation()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("joinInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetJoinMeetingIdSettings gets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings cannot be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
func (m *OnlineMeeting) GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable) {
    val, err := m.GetBackingStore().Get("joinMeetingIdSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(JoinMeetingIdSettingsable)
    }
    return nil
}
// GetJoinUrl gets the joinUrl property value. The joinUrl property
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
// GetJoinWebUrl gets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) GetJoinWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("joinWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
func (m *OnlineMeeting) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    val, err := m.GetBackingStore().Get("lobbyBypassSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LobbyBypassSettingsable)
    }
    return nil
}
// GetMeetingAttendanceReport gets the meetingAttendanceReport property value. The meetingAttendanceReport property
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
// GetParticipants gets the participants property value. The participants associated with the online meeting. This includes the organizer and the attendees.
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
// GetRecordAutomatically gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) GetRecordAutomatically()(*bool) {
    val, err := m.GetBackingStore().Get("recordAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRecording gets the recording property value. The content stream of the recording of a Teams live event. Read-only.
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
// GetRecordings gets the recordings property value. The recordings property
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
// GetRegistration gets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
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
// GetShareMeetingChatHistoryDefault gets the shareMeetingChatHistoryDefault property value. Specifies whether meeting chat history is shared with participants.  Possible values are: all, none, unknownFutureValue.
func (m *OnlineMeeting) GetShareMeetingChatHistoryDefault()(*MeetingChatHistoryDefaultMode) {
    val, err := m.GetBackingStore().Get("shareMeetingChatHistoryDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MeetingChatHistoryDefaultMode)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
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
// GetSubject gets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTranscripts gets the transcripts property value. The transcripts of an online meeting. Read-only.
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
// GetVideoTeleconferenceId gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    val, err := m.GetBackingStore().Get("videoTeleconferenceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualAppointment gets the virtualAppointment property value. The virtualAppointment property
func (m *OnlineMeeting) GetVirtualAppointment()(VirtualAppointmentable) {
    val, err := m.GetBackingStore().Get("virtualAppointment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualAppointmentable)
    }
    return nil
}
// GetWatermarkProtection gets the watermarkProtection property value. Specifies whether a watermark should be applied to a content type by the client application.
func (m *OnlineMeeting) GetWatermarkProtection()(WatermarkProtectionValuesable) {
    val, err := m.GetBackingStore().Get("watermarkProtection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WatermarkProtectionValuesable)
    }
    return nil
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
    if m.GetAllowMeetingChat() != nil {
        cast := (*m.GetAllowMeetingChat()).String()
        err = writer.WriteStringValue("allowMeetingChat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowParticipantsToChangeName", m.GetAllowParticipantsToChangeName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowRecording", m.GetAllowRecording())
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
        err = writer.WriteBoolValue("allowTranscription", m.GetAllowTranscription())
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendanceReports()))
        for i, v := range m.GetAttendanceReports() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    if m.GetShareMeetingChatHistoryDefault() != nil {
        cast := (*m.GetShareMeetingChatHistoryDefault()).String()
        err = writer.WriteStringValue("shareMeetingChatHistoryDefault", &cast)
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
    {
        err = writer.WriteObjectValue("watermarkProtection", m.GetWatermarkProtection())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAttendeeToEnableCamera sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    err := m.GetBackingStore().Set("allowAttendeeToEnableCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowAttendeeToEnableMic sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    err := m.GetBackingStore().Set("allowAttendeeToEnableMic", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowedPresenters sets the allowedPresenters property value. Specifies who can be a presenter in a meeting.
func (m *OnlineMeeting) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    err := m.GetBackingStore().Set("allowedPresenters", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowMeetingChat sets the allowMeetingChat property value. Specifies the mode of meeting chat.
func (m *OnlineMeeting) SetAllowMeetingChat(value *MeetingChatMode)() {
    err := m.GetBackingStore().Set("allowMeetingChat", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowParticipantsToChangeName sets the allowParticipantsToChangeName property value. Specifies if participants are allowed to rename themselves in an instance of the meeting.
func (m *OnlineMeeting) SetAllowParticipantsToChangeName(value *bool)() {
    err := m.GetBackingStore().Set("allowParticipantsToChangeName", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowRecording sets the allowRecording property value. Indicates whether recording is enabled for the meeting.
func (m *OnlineMeeting) SetAllowRecording(value *bool)() {
    err := m.GetBackingStore().Set("allowRecording", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowTeamworkReactions sets the allowTeamworkReactions property value. Indicates if Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    err := m.GetBackingStore().Set("allowTeamworkReactions", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowTranscription sets the allowTranscription property value. Indicates whether transcription is enabled for the meeting.
func (m *OnlineMeeting) SetAllowTranscription(value *bool)() {
    err := m.GetBackingStore().Set("allowTranscription", value)
    if err != nil {
        panic(err)
    }
}
// SetAlternativeRecording sets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) SetAlternativeRecording(value []byte)() {
    err := m.GetBackingStore().Set("alternativeRecording", value)
    if err != nil {
        panic(err)
    }
}
// SetAnonymizeIdentityForRoles sets the anonymizeIdentityForRoles property value. Specifies whose identity will be anonymized in the meeting. Possible values are: attendee. The attendee value cannot be removed through a PATCH operation once added.
func (m *OnlineMeeting) SetAnonymizeIdentityForRoles(value []OnlineMeetingRole)() {
    err := m.GetBackingStore().Set("anonymizeIdentityForRoles", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendanceReports sets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) SetAttendanceReports(value []MeetingAttendanceReportable)() {
    err := m.GetBackingStore().Set("attendanceReports", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendeeReport sets the attendeeReport property value. The content stream of the attendee report of a Teams live event. Read-only.
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    err := m.GetBackingStore().Set("attendeeReport", value)
    if err != nil {
        panic(err)
    }
}
// SetAudioConferencing sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) SetAudioConferencing(value AudioConferencingable)() {
    err := m.GetBackingStore().Set("audioConferencing", value)
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
// SetBroadcastSettings sets the broadcastSettings property value. Settings related to a live event.
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
// SetChatInfo sets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) SetChatInfo(value ChatInfoable)() {
    err := m.GetBackingStore().Set("chatInfo", value)
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
// SetIsBroadcast sets the isBroadcast property value. Indicates whether this is a Teams live event.
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    err := m.GetBackingStore().Set("isBroadcast", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEntryExitAnnounced sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    err := m.GetBackingStore().Set("isEntryExitAnnounced", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinInformation sets the joinInformation property value. The join information in the language and locale variant specified in 'Accept-Language' request HTTP header. Read-only.
func (m *OnlineMeeting) SetJoinInformation(value ItemBodyable)() {
    err := m.GetBackingStore().Set("joinInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetJoinMeetingIdSettings sets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings cannot be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
func (m *OnlineMeeting) SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)() {
    err := m.GetBackingStore().Set("joinMeetingIdSettings", value)
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
// SetJoinWebUrl sets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) SetJoinWebUrl(value *string)() {
    err := m.GetBackingStore().Set("joinWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting lobby.
func (m *OnlineMeeting) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    err := m.GetBackingStore().Set("lobbyBypassSettings", value)
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
// SetParticipants sets the participants property value. The participants associated with the online meeting. This includes the organizer and the attendees.
func (m *OnlineMeeting) SetParticipants(value MeetingParticipantsable)() {
    err := m.GetBackingStore().Set("participants", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordAutomatically sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) SetRecordAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("recordAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetRecording sets the recording property value. The content stream of the recording of a Teams live event. Read-only.
func (m *OnlineMeeting) SetRecording(value []byte)() {
    err := m.GetBackingStore().Set("recording", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordings sets the recordings property value. The recordings property
func (m *OnlineMeeting) SetRecordings(value []CallRecordingable)() {
    err := m.GetBackingStore().Set("recordings", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistration sets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) SetRegistration(value MeetingRegistrationable)() {
    err := m.GetBackingStore().Set("registration", value)
    if err != nil {
        panic(err)
    }
}
// SetShareMeetingChatHistoryDefault sets the shareMeetingChatHistoryDefault property value. Specifies whether meeting chat history is shared with participants.  Possible values are: all, none, unknownFutureValue.
func (m *OnlineMeeting) SetShareMeetingChatHistoryDefault(value *MeetingChatHistoryDefaultMode)() {
    err := m.GetBackingStore().Set("shareMeetingChatHistoryDefault", value)
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
// SetSubject sets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
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
// SetVideoTeleconferenceId sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    err := m.GetBackingStore().Set("videoTeleconferenceId", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualAppointment sets the virtualAppointment property value. The virtualAppointment property
func (m *OnlineMeeting) SetVirtualAppointment(value VirtualAppointmentable)() {
    err := m.GetBackingStore().Set("virtualAppointment", value)
    if err != nil {
        panic(err)
    }
}
// SetWatermarkProtection sets the watermarkProtection property value. Specifies whether a watermark should be applied to a content type by the client application.
func (m *OnlineMeeting) SetWatermarkProtection(value WatermarkProtectionValuesable)() {
    err := m.GetBackingStore().Set("watermarkProtection", value)
    if err != nil {
        panic(err)
    }
}
// OnlineMeetingable 
type OnlineMeetingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAttendeeToEnableCamera()(*bool)
    GetAllowAttendeeToEnableMic()(*bool)
    GetAllowedPresenters()(*OnlineMeetingPresenters)
    GetAllowMeetingChat()(*MeetingChatMode)
    GetAllowParticipantsToChangeName()(*bool)
    GetAllowRecording()(*bool)
    GetAllowTeamworkReactions()(*bool)
    GetAllowTranscription()(*bool)
    GetAlternativeRecording()([]byte)
    GetAnonymizeIdentityForRoles()([]OnlineMeetingRole)
    GetAttendanceReports()([]MeetingAttendanceReportable)
    GetAttendeeReport()([]byte)
    GetAudioConferencing()(AudioConferencingable)
    GetBroadcastRecording()([]byte)
    GetBroadcastSettings()(BroadcastMeetingSettingsable)
    GetCapabilities()([]MeetingCapabilities)
    GetChatInfo()(ChatInfoable)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetIsBroadcast()(*bool)
    GetIsEntryExitAnnounced()(*bool)
    GetJoinInformation()(ItemBodyable)
    GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable)
    GetJoinUrl()(*string)
    GetJoinWebUrl()(*string)
    GetLobbyBypassSettings()(LobbyBypassSettingsable)
    GetMeetingAttendanceReport()(MeetingAttendanceReportable)
    GetParticipants()(MeetingParticipantsable)
    GetRecordAutomatically()(*bool)
    GetRecording()([]byte)
    GetRecordings()([]CallRecordingable)
    GetRegistration()(MeetingRegistrationable)
    GetShareMeetingChatHistoryDefault()(*MeetingChatHistoryDefaultMode)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    GetTranscripts()([]CallTranscriptable)
    GetVideoTeleconferenceId()(*string)
    GetVirtualAppointment()(VirtualAppointmentable)
    GetWatermarkProtection()(WatermarkProtectionValuesable)
    SetAllowAttendeeToEnableCamera(value *bool)()
    SetAllowAttendeeToEnableMic(value *bool)()
    SetAllowedPresenters(value *OnlineMeetingPresenters)()
    SetAllowMeetingChat(value *MeetingChatMode)()
    SetAllowParticipantsToChangeName(value *bool)()
    SetAllowRecording(value *bool)()
    SetAllowTeamworkReactions(value *bool)()
    SetAllowTranscription(value *bool)()
    SetAlternativeRecording(value []byte)()
    SetAnonymizeIdentityForRoles(value []OnlineMeetingRole)()
    SetAttendanceReports(value []MeetingAttendanceReportable)()
    SetAttendeeReport(value []byte)()
    SetAudioConferencing(value AudioConferencingable)()
    SetBroadcastRecording(value []byte)()
    SetBroadcastSettings(value BroadcastMeetingSettingsable)()
    SetCapabilities(value []MeetingCapabilities)()
    SetChatInfo(value ChatInfoable)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetIsBroadcast(value *bool)()
    SetIsEntryExitAnnounced(value *bool)()
    SetJoinInformation(value ItemBodyable)()
    SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)()
    SetJoinUrl(value *string)()
    SetJoinWebUrl(value *string)()
    SetLobbyBypassSettings(value LobbyBypassSettingsable)()
    SetMeetingAttendanceReport(value MeetingAttendanceReportable)()
    SetParticipants(value MeetingParticipantsable)()
    SetRecordAutomatically(value *bool)()
    SetRecording(value []byte)()
    SetRecordings(value []CallRecordingable)()
    SetRegistration(value MeetingRegistrationable)()
    SetShareMeetingChatHistoryDefault(value *MeetingChatHistoryDefaultMode)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
    SetTranscripts(value []CallTranscriptable)()
    SetVideoTeleconferenceId(value *string)()
    SetVirtualAppointment(value VirtualAppointmentable)()
    SetWatermarkProtection(value WatermarkProtectionValuesable)()
}
