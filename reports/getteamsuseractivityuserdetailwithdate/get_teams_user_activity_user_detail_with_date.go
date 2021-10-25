package getteamsuseractivityuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetTeamsUserActivityUserDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    adHocMeetingsAttendedCount *int64;
    adHocMeetingsOrganizedCount *int64;
    assignedProducts []string;
    audioDuration *string;
    callCount *int64;
    deletedDate *string;
    hasOtherAction *bool;
    isDeleted *bool;
    isLicensed *bool;
    lastActivityDate *string;
    meetingCount *int64;
    meetingsAttendedCount *int64;
    meetingsOrganizedCount *int64;
    privateChatMessageCount *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    scheduledOneTimeMeetingsAttendedCount *int64;
    scheduledOneTimeMeetingsOrganizedCount *int64;
    scheduledRecurringMeetingsAttendedCount *int64;
    scheduledRecurringMeetingsOrganizedCount *int64;
    screenShareDuration *string;
    teamChatMessageCount *int64;
    userPrincipalName *string;
    videoDuration *string;
}
func NewGetTeamsUserActivityUserDetailWithDate()(*GetTeamsUserActivityUserDetailWithDate) {
    m := &GetTeamsUserActivityUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetAdHocMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.adHocMeetingsAttendedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetAdHocMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.adHocMeetingsOrganizedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetAudioDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.audioDuration
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetCallCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.callCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetHasOtherAction()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOtherAction
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetIsLicensed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLicensed
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetMeetingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingsAttendedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingsOrganizedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetPrivateChatMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.privateChatMessageCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetScheduledOneTimeMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledOneTimeMeetingsAttendedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetScheduledOneTimeMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledOneTimeMeetingsOrganizedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetScheduledRecurringMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledRecurringMeetingsAttendedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetScheduledRecurringMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledRecurringMeetingsOrganizedCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetScreenShareDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.screenShareDuration
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetTeamChatMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamChatMessageCount
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetVideoDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.videoDuration
    }
}
func (m *GetTeamsUserActivityUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adHocMeetingsAttendedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAdHocMeetingsAttendedCount(val)
        return nil
    }
    res["adHocMeetingsOrganizedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAdHocMeetingsOrganizedCount(val)
        return nil
    }
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignedProducts(res)
        return nil
    }
    res["audioDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAudioDuration(val)
        return nil
    }
    res["callCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCallCount(val)
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
        return nil
    }
    res["hasOtherAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasOtherAction(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["isLicensed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLicensed(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["meetingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMeetingCount(val)
        return nil
    }
    res["meetingsAttendedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMeetingsAttendedCount(val)
        return nil
    }
    res["meetingsOrganizedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMeetingsOrganizedCount(val)
        return nil
    }
    res["privateChatMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPrivateChatMessageCount(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["scheduledOneTimeMeetingsAttendedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetScheduledOneTimeMeetingsAttendedCount(val)
        return nil
    }
    res["scheduledOneTimeMeetingsOrganizedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetScheduledOneTimeMeetingsOrganizedCount(val)
        return nil
    }
    res["scheduledRecurringMeetingsAttendedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetScheduledRecurringMeetingsAttendedCount(val)
        return nil
    }
    res["scheduledRecurringMeetingsOrganizedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetScheduledRecurringMeetingsOrganizedCount(val)
        return nil
    }
    res["screenShareDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScreenShareDuration(val)
        return nil
    }
    res["teamChatMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTeamChatMessageCount(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["videoDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVideoDuration(val)
        return nil
    }
    return res
}
func (m *GetTeamsUserActivityUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
func (m *GetTeamsUserActivityUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("adHocMeetingsAttendedCount", m.GetAdHocMeetingsAttendedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("adHocMeetingsOrganizedCount", m.GetAdHocMeetingsOrganizedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("audioDuration", m.GetAudioDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("callCount", m.GetCallCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasOtherAction", m.GetHasOtherAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLicensed", m.GetIsLicensed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("meetingCount", m.GetMeetingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("meetingsAttendedCount", m.GetMeetingsAttendedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("meetingsOrganizedCount", m.GetMeetingsOrganizedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("privateChatMessageCount", m.GetPrivateChatMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("scheduledOneTimeMeetingsAttendedCount", m.GetScheduledOneTimeMeetingsAttendedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("scheduledOneTimeMeetingsOrganizedCount", m.GetScheduledOneTimeMeetingsOrganizedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("scheduledRecurringMeetingsAttendedCount", m.GetScheduledRecurringMeetingsAttendedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("scheduledRecurringMeetingsOrganizedCount", m.GetScheduledRecurringMeetingsOrganizedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("screenShareDuration", m.GetScreenShareDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamChatMessageCount", m.GetTeamChatMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoDuration", m.GetVideoDuration())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetAdHocMeetingsAttendedCount(value *int64)() {
    m.adHocMeetingsAttendedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetAdHocMeetingsOrganizedCount(value *int64)() {
    m.adHocMeetingsOrganizedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetAudioDuration(value *string)() {
    m.audioDuration = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetCallCount(value *int64)() {
    m.callCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetHasOtherAction(value *bool)() {
    m.hasOtherAction = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetIsLicensed(value *bool)() {
    m.isLicensed = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetMeetingCount(value *int64)() {
    m.meetingCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetMeetingsAttendedCount(value *int64)() {
    m.meetingsAttendedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetMeetingsOrganizedCount(value *int64)() {
    m.meetingsOrganizedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetPrivateChatMessageCount(value *int64)() {
    m.privateChatMessageCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetScheduledOneTimeMeetingsAttendedCount(value *int64)() {
    m.scheduledOneTimeMeetingsAttendedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetScheduledOneTimeMeetingsOrganizedCount(value *int64)() {
    m.scheduledOneTimeMeetingsOrganizedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetScheduledRecurringMeetingsAttendedCount(value *int64)() {
    m.scheduledRecurringMeetingsAttendedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetScheduledRecurringMeetingsOrganizedCount(value *int64)() {
    m.scheduledRecurringMeetingsOrganizedCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetScreenShareDuration(value *string)() {
    m.screenShareDuration = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetTeamChatMessageCount(value *int64)() {
    m.teamChatMessageCount = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *GetTeamsUserActivityUserDetailWithDate) SetVideoDuration(value *string)() {
    m.videoDuration = value
}
