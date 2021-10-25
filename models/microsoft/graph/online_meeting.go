package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnlineMeeting struct {
    Entity
    accessLevel *AccessLevel;
    allowAttendeeToEnableCamera *bool;
    allowAttendeeToEnableMic *bool;
    allowedPresenters *OnlineMeetingPresenters;
    allowMeetingChat *MeetingChatMode;
    allowTeamworkReactions *bool;
    alternativeRecording []byte;
    attendeeReport []byte;
    audioConferencing *AudioConferencing;
    broadcastSettings *BroadcastMeetingSettings;
    canceledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    capabilities []MeetingCapabilities;
    chatInfo *ChatInfo;
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    entryExitAnnouncement *bool;
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    externalId *string;
    isBroadcast *bool;
    isCancelled *bool;
    isEntryExitAnnounced *bool;
    joinInformation *ItemBody;
    joinUrl *string;
    lobbyBypassSettings *LobbyBypassSettings;
    meetingAttendanceReport *MeetingAttendanceReport;
    participants *MeetingParticipants;
    recording []byte;
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    subject *string;
    videoTeleconferenceId *string;
}
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OnlineMeeting) GetAccessLevel()(*AccessLevel) {
    if m == nil {
        return nil
    } else {
        return m.accessLevel
    }
}
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableCamera
    }
}
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableMic
    }
}
func (m *OnlineMeeting) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    if m == nil {
        return nil
    } else {
        return m.allowedPresenters
    }
}
func (m *OnlineMeeting) GetAllowMeetingChat()(*MeetingChatMode) {
    if m == nil {
        return nil
    } else {
        return m.allowMeetingChat
    }
}
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamworkReactions
    }
}
func (m *OnlineMeeting) GetAlternativeRecording()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.alternativeRecording
    }
}
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.attendeeReport
    }
}
func (m *OnlineMeeting) GetAudioConferencing()(*AudioConferencing) {
    if m == nil {
        return nil
    } else {
        return m.audioConferencing
    }
}
func (m *OnlineMeeting) GetBroadcastSettings()(*BroadcastMeetingSettings) {
    if m == nil {
        return nil
    } else {
        return m.broadcastSettings
    }
}
func (m *OnlineMeeting) GetCanceledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.canceledDateTime
    }
}
func (m *OnlineMeeting) GetCapabilities()([]MeetingCapabilities) {
    if m == nil {
        return nil
    } else {
        return m.capabilities
    }
}
func (m *OnlineMeeting) GetChatInfo()(*ChatInfo) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
func (m *OnlineMeeting) GetEntryExitAnnouncement()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.entryExitAnnouncement
    }
}
func (m *OnlineMeeting) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *OnlineMeeting) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBroadcast
    }
}
func (m *OnlineMeeting) GetIsCancelled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCancelled
    }
}
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEntryExitAnnounced
    }
}
func (m *OnlineMeeting) GetJoinInformation()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.joinInformation
    }
}
func (m *OnlineMeeting) GetJoinUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinUrl
    }
}
func (m *OnlineMeeting) GetLobbyBypassSettings()(*LobbyBypassSettings) {
    if m == nil {
        return nil
    } else {
        return m.lobbyBypassSettings
    }
}
func (m *OnlineMeeting) GetMeetingAttendanceReport()(*MeetingAttendanceReport) {
    if m == nil {
        return nil
    } else {
        return m.meetingAttendanceReport
    }
}
func (m *OnlineMeeting) GetParticipants()(*MeetingParticipants) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
func (m *OnlineMeeting) GetRecording()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.recording
    }
}
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
func (m *OnlineMeeting) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.videoTeleconferenceId
    }
}
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessLevel)
        if err != nil {
            return err
        }
        cast := val.(AccessLevel)
        m.SetAccessLevel(&cast)
        return nil
    }
    res["allowAttendeeToEnableCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAttendeeToEnableCamera(val)
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAttendeeToEnableMic(val)
        return nil
    }
    res["allowedPresenters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        cast := val.(OnlineMeetingPresenters)
        m.SetAllowedPresenters(&cast)
        return nil
    }
    res["allowMeetingChat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatMode)
        if err != nil {
            return err
        }
        cast := val.(MeetingChatMode)
        m.SetAllowMeetingChat(&cast)
        return nil
    }
    res["allowTeamworkReactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowTeamworkReactions(val)
        return nil
    }
    res["alternativeRecording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetAlternativeRecording(val)
        return nil
    }
    res["attendeeReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetAttendeeReport(val)
        return nil
    }
    res["audioConferencing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudioConferencing() })
        if err != nil {
            return err
        }
        m.SetAudioConferencing(val.(*AudioConferencing))
        return nil
    }
    res["broadcastSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBroadcastMeetingSettings() })
        if err != nil {
            return err
        }
        m.SetBroadcastSettings(val.(*BroadcastMeetingSettings))
        return nil
    }
    res["canceledDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCanceledDateTime(val)
        return nil
    }
    res["capabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseMeetingCapabilities)
        if err != nil {
            return err
        }
        res := make([]MeetingCapabilities, len(val))
        for i, v := range val {
            res[i] = *(v.(*MeetingCapabilities))
        }
        m.SetCapabilities(res)
        return nil
    }
    res["chatInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatInfo() })
        if err != nil {
            return err
        }
        m.SetChatInfo(val.(*ChatInfo))
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreationDateTime(val)
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
    res["entryExitAnnouncement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEntryExitAnnouncement(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["isBroadcast"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBroadcast(val)
        return nil
    }
    res["isCancelled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCancelled(val)
        return nil
    }
    res["isEntryExitAnnounced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEntryExitAnnounced(val)
        return nil
    }
    res["joinInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetJoinInformation(val.(*ItemBody))
        return nil
    }
    res["joinUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoinUrl(val)
        return nil
    }
    res["lobbyBypassSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLobbyBypassSettings() })
        if err != nil {
            return err
        }
        m.SetLobbyBypassSettings(val.(*LobbyBypassSettings))
        return nil
    }
    res["meetingAttendanceReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingAttendanceReport() })
        if err != nil {
            return err
        }
        m.SetMeetingAttendanceReport(val.(*MeetingAttendanceReport))
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipants() })
        if err != nil {
            return err
        }
        m.SetParticipants(val.(*MeetingParticipants))
        return nil
    }
    res["recording"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetRecording(val)
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
    res["videoTeleconferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVideoTeleconferenceId(val)
        return nil
    }
    return res
}
func (m *OnlineMeeting) IsNil()(bool) {
    return m == nil
}
func (m *OnlineMeeting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessLevel() != nil {
        cast := m.GetAccessLevel().String()
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
        cast := m.GetAllowedPresenters().String()
        err = writer.WriteStringValue("allowedPresenters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowMeetingChat() != nil {
        cast := m.GetAllowMeetingChat().String()
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
    {
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
        err = writer.WriteByteArrayValue("recording", m.GetRecording())
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
func (m *OnlineMeeting) SetAccessLevel(value *AccessLevel)() {
    m.accessLevel = value
}
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    m.allowAttendeeToEnableCamera = value
}
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    m.allowAttendeeToEnableMic = value
}
func (m *OnlineMeeting) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    m.allowedPresenters = value
}
func (m *OnlineMeeting) SetAllowMeetingChat(value *MeetingChatMode)() {
    m.allowMeetingChat = value
}
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    m.allowTeamworkReactions = value
}
func (m *OnlineMeeting) SetAlternativeRecording(value []byte)() {
    m.alternativeRecording = value
}
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    m.attendeeReport = value
}
func (m *OnlineMeeting) SetAudioConferencing(value *AudioConferencing)() {
    m.audioConferencing = value
}
func (m *OnlineMeeting) SetBroadcastSettings(value *BroadcastMeetingSettings)() {
    m.broadcastSettings = value
}
func (m *OnlineMeeting) SetCanceledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.canceledDateTime = value
}
func (m *OnlineMeeting) SetCapabilities(value []MeetingCapabilities)() {
    m.capabilities = value
}
func (m *OnlineMeeting) SetChatInfo(value *ChatInfo)() {
    m.chatInfo = value
}
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
func (m *OnlineMeeting) SetEntryExitAnnouncement(value *bool)() {
    m.entryExitAnnouncement = value
}
func (m *OnlineMeeting) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *OnlineMeeting) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    m.isBroadcast = value
}
func (m *OnlineMeeting) SetIsCancelled(value *bool)() {
    m.isCancelled = value
}
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    m.isEntryExitAnnounced = value
}
func (m *OnlineMeeting) SetJoinInformation(value *ItemBody)() {
    m.joinInformation = value
}
func (m *OnlineMeeting) SetJoinUrl(value *string)() {
    m.joinUrl = value
}
func (m *OnlineMeeting) SetLobbyBypassSettings(value *LobbyBypassSettings)() {
    m.lobbyBypassSettings = value
}
func (m *OnlineMeeting) SetMeetingAttendanceReport(value *MeetingAttendanceReport)() {
    m.meetingAttendanceReport = value
}
func (m *OnlineMeeting) SetParticipants(value *MeetingParticipants)() {
    m.participants = value
}
func (m *OnlineMeeting) SetRecording(value []byte)() {
    m.recording = value
}
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
func (m *OnlineMeeting) SetSubject(value *string)() {
    m.subject = value
}
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    m.videoTeleconferenceId = value
}
