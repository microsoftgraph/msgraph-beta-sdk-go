package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnlineMeeting 
type OnlineMeeting struct {
    Entity
    // 
    accessLevel *AccessLevel;
    // Indicates whether attendees can turn on their camera.
    allowAttendeeToEnableCamera *bool;
    // Indicates whether attendees can turn on their microphone.
    allowAttendeeToEnableMic *bool;
    // Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
    allowedPresenters *OnlineMeetingPresenters;
    // Specifies the mode of meeting chat.
    allowMeetingChat *MeetingChatMode;
    // Indicates whether Teams reactions are enabled for the meeting.
    allowTeamworkReactions *bool;
    // The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
    alternativeRecording []byte;
    // The attendance reports of an online meeting. Read-only.
    attendanceReports []MeetingAttendanceReport;
    // The content stream of the attendee report of a Microsoft Teams live event. Read-only.
    attendeeReport []byte;
    // The phone access (dial-in) information for an online meeting. Read-only.
    audioConferencing *AudioConferencing;
    // Settings related to a live event.
    broadcastSettings *BroadcastMeetingSettings;
    // 
    canceledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    capabilities []MeetingCapabilities;
    // The chat information associated with this online meeting.
    chatInfo *ChatInfo;
    // The meeting creation time in UTC. Read-only.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    entryExitAnnouncement *bool;
    // 
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The external ID. A custom ID. Optional.
    externalId *string;
    // Indicates if this is a Teams live event.
    isBroadcast *bool;
    // 
    isCancelled *bool;
    // Indicates whether to announce when callers join or leave.
    isEntryExitAnnounced *bool;
    // The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
    joinInformation *ItemBody;
    // 
    joinMeetingIdSettings *JoinMeetingIdSettings;
    // 
    joinUrl *string;
    // Specifies which participants can bypass the meeting   lobby.
    lobbyBypassSettings *LobbyBypassSettings;
    // 
    meetingAttendanceReport *MeetingAttendanceReport;
    // The participants associated with the online meeting.  This includes the organizer and the attendees.
    participants *MeetingParticipants;
    // Indicates whether to record the meeting automatically.
    recordAutomatically *bool;
    // The content stream of the recording of a Teams live event. Read-only.
    recording []byte;
    // The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
    registration *MeetingRegistration;
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The subject of the online meeting.
    subject *string;
    // The video teleconferencing ID. Read-only.
    videoTeleconferenceId *string;
}
// NewOnlineMeeting instantiates a new onlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessLevel gets the accessLevel property value. 
func (m *OnlineMeeting) GetAccessLevel()(*AccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
// GetAllowAttendeeToEnableCamera gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableCamera
    }
}
// GetAllowAttendeeToEnableMic gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableMic
    }
}
// GetAllowedPresenters gets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
func (m *OnlineMeeting) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    if m == nil {
        return nil
    } else {
        return m.allowedPresenters
    }
}
// GetAllowMeetingChat gets the allowMeetingChat property value. Specifies the mode of meeting chat.
func (m *OnlineMeeting) GetAllowMeetingChat()(*MeetingChatMode) {
    if m == nil {
        return nil
    } else {
        return m.allowMeetingChat
    }
}
// GetAllowTeamworkReactions gets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamworkReactions
    }
}
// GetAlternativeRecording gets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) GetAlternativeRecording()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.alternativeRecording
    }
}
// GetAttendanceReports gets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) GetAttendanceReports()([]MeetingAttendanceReport) {
    if m == nil {
        return nil
    } else {
        return m.attendanceReports
    }
}
// GetAttendeeReport gets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.attendeeReport
    }
}
// GetAudioConferencing gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) GetAudioConferencing()(*AudioConferencing) {
    if m == nil {
        return nil
    } else {
        return m.audioConferencing
    }
}
// GetBroadcastSettings gets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) GetBroadcastSettings()(*BroadcastMeetingSettings) {
    if m == nil {
        return nil
    } else {
        return m.broadcastSettings
    }
}
// GetCanceledDateTime gets the canceledDateTime property value. 
func (m *OnlineMeeting) GetCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.canceledDateTime
    }
}
// GetCapabilities gets the capabilities property value. 
func (m *OnlineMeeting) GetCapabilities()([]MeetingCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
// GetChatInfo gets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) GetChatInfo()(*ChatInfo) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
// GetCreationDateTime gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetEntryExitAnnouncement gets the entryExitAnnouncement property value. 
func (m *OnlineMeeting) GetEntryExitAnnouncement()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.entryExitAnnouncement
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. 
func (m *OnlineMeeting) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetExternalId gets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetIsBroadcast gets the isBroadcast property value. Indicates if this is a Teams live event.
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBroadcast
    }
}
// GetIsCancelled gets the isCancelled property value. 
func (m *OnlineMeeting) GetIsCancelled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCancelled
    }
}
// GetIsEntryExitAnnounced gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEntryExitAnnounced
    }
}
// GetJoinInformation gets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
func (m *OnlineMeeting) GetJoinInformation()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.joinInformation
    }
}
// GetJoinMeetingIdSettings gets the joinMeetingIdSettings property value. 
func (m *OnlineMeeting) GetJoinMeetingIdSettings()(*JoinMeetingIdSettings) {
    if m == nil {
        return nil
    } else {
        return m.joinMeetingIdSettings
    }
}
// GetJoinUrl gets the joinUrl property value. 
func (m *OnlineMeeting) GetJoinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinUrl
    }
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
func (m *OnlineMeeting) GetLobbyBypassSettings()(*LobbyBypassSettings) {
    if m == nil {
        return nil
    } else {
        return m.lobbyBypassSettings
    }
}
// GetMeetingAttendanceReport gets the meetingAttendanceReport property value. 
func (m *OnlineMeeting) GetMeetingAttendanceReport()(*MeetingAttendanceReport) {
    if m == nil {
        return nil
    } else {
        return m.meetingAttendanceReport
    }
}
// GetParticipants gets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
func (m *OnlineMeeting) GetParticipants()(*MeetingParticipants) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// GetRecordAutomatically gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) GetRecordAutomatically()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recordAutomatically
    }
}
// GetRecording gets the recording property value. The content stream of the recording of a Teams live event. Read-only.
func (m *OnlineMeeting) GetRecording()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.recording
    }
}
// GetRegistration gets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) GetRegistration()(*MeetingRegistration) {
    if m == nil {
        return nil
    } else {
        return m.registration
    }
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetSubject gets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetVideoTeleconferenceId gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.videoTeleconferenceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*AccessLevel))
        }
        return nil
    }
    res["allowAttendeeToEnableCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableCamera(val)
        }
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableMic(val)
        }
        return nil
    }
    res["allowedPresenters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedPresenters(val.(*OnlineMeetingPresenters))
        }
        return nil
    }
    res["allowMeetingChat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMeetingChat(val.(*MeetingChatMode))
        }
        return nil
    }
    res["allowTeamworkReactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamworkReactions(val)
        }
        return nil
    }
    res["alternativeRecording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeRecording(val)
        }
        return nil
    }
    res["attendanceReports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingAttendanceReport() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingAttendanceReport, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingAttendanceReport))
            }
            m.SetAttendanceReports(res)
        }
        return nil
    }
    res["attendeeReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttendeeReport(val)
        }
        return nil
    }
    res["audioConferencing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudioConferencing() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioConferencing(val.(*AudioConferencing))
        }
        return nil
    }
    res["broadcastSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBroadcastMeetingSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBroadcastSettings(val.(*BroadcastMeetingSettings))
        }
        return nil
    }
    res["canceledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanceledDateTime(val)
        }
        return nil
    }
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseMeetingCapabilities)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingCapabilities, len(val))
            for i, v := range val {
                res[i] = *(v.(*MeetingCapabilities))
            }
            m.SetCapabilities(res)
        }
        return nil
    }
    res["chatInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(*ChatInfo))
        }
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
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
    res["entryExitAnnouncement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryExitAnnouncement(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["isBroadcast"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBroadcast(val)
        }
        return nil
    }
    res["isCancelled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCancelled(val)
        }
        return nil
    }
    res["isEntryExitAnnounced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEntryExitAnnounced(val)
        }
        return nil
    }
    res["joinInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinInformation(val.(*ItemBody))
        }
        return nil
    }
    res["joinMeetingIdSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJoinMeetingIdSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingIdSettings(val.(*JoinMeetingIdSettings))
        }
        return nil
    }
    res["joinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinUrl(val)
        }
        return nil
    }
    res["lobbyBypassSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLobbyBypassSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(*LobbyBypassSettings))
        }
        return nil
    }
    res["meetingAttendanceReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingAttendanceReport() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingAttendanceReport(val.(*MeetingAttendanceReport))
        }
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipants() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(*MeetingParticipants))
        }
        return nil
    }
    res["recordAutomatically"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordAutomatically(val)
        }
        return nil
    }
    res["recording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecording(val)
        }
        return nil
    }
    res["registration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingRegistration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistration(val.(*MeetingRegistration))
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
    res["videoTeleconferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoTeleconferenceId(val)
        }
        return nil
    }
    return res
}
func (m *OnlineMeeting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnlineMeeting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err = writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
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
    if m.GetAttendanceReports() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceReports()))
        for i, v := range m.GetAttendanceReports() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteObjectValue("broadcastSettings", m.GetBroadcastSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("canceledDateTime", m.GetCanceledDateTime())
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
        err = writer.WriteBoolValue("entryExitAnnouncement", m.GetEntryExitAnnouncement())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
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
        err = writer.WriteBoolValue("isCancelled", m.GetIsCancelled())
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
    {
        err = writer.WriteStringValue("videoTeleconferenceId", m.GetVideoTeleconferenceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessLevel sets the accessLevel property value. 
func (m *OnlineMeeting) SetAccessLevel(value *AccessLevel)() {
    if m != nil {
        m.accessLevel = value
    }
}
// SetAllowAttendeeToEnableCamera sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    if m != nil {
        m.allowAttendeeToEnableCamera = value
    }
}
// SetAllowAttendeeToEnableMic sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    if m != nil {
        m.allowAttendeeToEnableMic = value
    }
}
// SetAllowedPresenters sets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
func (m *OnlineMeeting) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    if m != nil {
        m.allowedPresenters = value
    }
}
// SetAllowMeetingChat sets the allowMeetingChat property value. Specifies the mode of meeting chat.
func (m *OnlineMeeting) SetAllowMeetingChat(value *MeetingChatMode)() {
    if m != nil {
        m.allowMeetingChat = value
    }
}
// SetAllowTeamworkReactions sets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    if m != nil {
        m.allowTeamworkReactions = value
    }
}
// SetAlternativeRecording sets the alternativeRecording property value. The content stream of the alternative recording of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) SetAlternativeRecording(value []byte)() {
    if m != nil {
        m.alternativeRecording = value
    }
}
// SetAttendanceReports sets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) SetAttendanceReports(value []MeetingAttendanceReport)() {
    if m != nil {
        m.attendanceReports = value
    }
}
// SetAttendeeReport sets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    if m != nil {
        m.attendeeReport = value
    }
}
// SetAudioConferencing sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) SetAudioConferencing(value *AudioConferencing)() {
    if m != nil {
        m.audioConferencing = value
    }
}
// SetBroadcastSettings sets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) SetBroadcastSettings(value *BroadcastMeetingSettings)() {
    if m != nil {
        m.broadcastSettings = value
    }
}
// SetCanceledDateTime sets the canceledDateTime property value. 
func (m *OnlineMeeting) SetCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.canceledDateTime = value
    }
}
// SetCapabilities sets the capabilities property value. 
func (m *OnlineMeeting) SetCapabilities(value []MeetingCapabilities)() {
    if m != nil {
        m.capabilities = value
    }
}
// SetChatInfo sets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) SetChatInfo(value *ChatInfo)() {
    if m != nil {
        m.chatInfo = value
    }
}
// SetCreationDateTime sets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.creationDateTime = value
    }
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetEntryExitAnnouncement sets the entryExitAnnouncement property value. 
func (m *OnlineMeeting) SetEntryExitAnnouncement(value *bool)() {
    if m != nil {
        m.entryExitAnnouncement = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. 
func (m *OnlineMeeting) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetExternalId sets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetIsBroadcast sets the isBroadcast property value. Indicates if this is a Teams live event.
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    if m != nil {
        m.isBroadcast = value
    }
}
// SetIsCancelled sets the isCancelled property value. 
func (m *OnlineMeeting) SetIsCancelled(value *bool)() {
    if m != nil {
        m.isCancelled = value
    }
}
// SetIsEntryExitAnnounced sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    if m != nil {
        m.isEntryExitAnnounced = value
    }
}
// SetJoinInformation sets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
func (m *OnlineMeeting) SetJoinInformation(value *ItemBody)() {
    if m != nil {
        m.joinInformation = value
    }
}
// SetJoinMeetingIdSettings sets the joinMeetingIdSettings property value. 
func (m *OnlineMeeting) SetJoinMeetingIdSettings(value *JoinMeetingIdSettings)() {
    if m != nil {
        m.joinMeetingIdSettings = value
    }
}
// SetJoinUrl sets the joinUrl property value. 
func (m *OnlineMeeting) SetJoinUrl(value *string)() {
    if m != nil {
        m.joinUrl = value
    }
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
func (m *OnlineMeeting) SetLobbyBypassSettings(value *LobbyBypassSettings)() {
    if m != nil {
        m.lobbyBypassSettings = value
    }
}
// SetMeetingAttendanceReport sets the meetingAttendanceReport property value. 
func (m *OnlineMeeting) SetMeetingAttendanceReport(value *MeetingAttendanceReport)() {
    if m != nil {
        m.meetingAttendanceReport = value
    }
}
// SetParticipants sets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
func (m *OnlineMeeting) SetParticipants(value *MeetingParticipants)() {
    if m != nil {
        m.participants = value
    }
}
// SetRecordAutomatically sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) SetRecordAutomatically(value *bool)() {
    if m != nil {
        m.recordAutomatically = value
    }
}
// SetRecording sets the recording property value. The content stream of the recording of a Teams live event. Read-only.
func (m *OnlineMeeting) SetRecording(value []byte)() {
    if m != nil {
        m.recording = value
    }
}
// SetRegistration sets the registration property value. The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
func (m *OnlineMeeting) SetRegistration(value *MeetingRegistration)() {
    if m != nil {
        m.registration = value
    }
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetSubject sets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetVideoTeleconferenceId sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    if m != nil {
        m.videoTeleconferenceId = value
    }
}
