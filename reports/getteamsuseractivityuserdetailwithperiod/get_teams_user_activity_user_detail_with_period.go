package getteamsuseractivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetTeamsUserActivityUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of ad hoc meetings a user participated in.
    adHocMeetingsAttendedCount *int64;
    // The number of ad hoc meetings a user organized.
    adHocMeetingsOrganizedCount *int64;
    // Products the user assigned with.
    assignedProducts []string;
    // Audio duration the user participated in.
    audioDuration *string;
    // The number of 1:1 calls that the user participated in.
    callCount *int64;
    // The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
    deletedDate *string;
    // The User is active but has performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
    hasOtherAction *bool;
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool;
    // Whether the user has been assigned a Teams license.
    isLicensed *bool;
    // The last date that the user participated in a Microsoft Teams activity.
    lastActivityDate *string;
    // The number of online meetings that the user participated in.
    meetingCount *int64;
    // The sum of the one-time scheduled, recurring, ad hoc and unclassified meetings a user participated in.
    meetingsAttendedCount *int64;
    // The sum of one-time scheduled, Recurring, ad hoc and unclassified meetings a user organized.
    meetingsOrganizedCount *int64;
    // The number of unique messages that the user posted in a private chat.
    privateChatMessageCount *int64;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of the one-time scheduled meetings a user participated in.
    scheduledOneTimeMeetingsAttendedCount *int64;
    // The number of one-time scheduled meetings a user organized.
    scheduledOneTimeMeetingsOrganizedCount *int64;
    // The number of the recurring meetings a user participated in.
    scheduledRecurringMeetingsAttendedCount *int64;
    // The number of recurring meetings a user organized.
    scheduledRecurringMeetingsOrganizedCount *int64;
    // Screen sharing duration the user participated in.
    screenShareDuration *string;
    // The number of unique messages that the user posted in a team chat.
    teamChatMessageCount *int64;
    // The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
    userPrincipalName *string;
    // Video duration the user participated in.
    videoDuration *string;
}
// Instantiates a new getTeamsUserActivityUserDetailWithPeriod and sets the default values.
func NewGetTeamsUserActivityUserDetailWithPeriod()(*GetTeamsUserActivityUserDetailWithPeriod) {
    m := &GetTeamsUserActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the adHocMeetingsAttendedCount property value. The number of ad hoc meetings a user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetAdHocMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.adHocMeetingsAttendedCount
    }
}
// Gets the adHocMeetingsOrganizedCount property value. The number of ad hoc meetings a user organized.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetAdHocMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.adHocMeetingsOrganizedCount
    }
}
// Gets the assignedProducts property value. Products the user assigned with.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
// Gets the audioDuration property value. Audio duration the user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetAudioDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.audioDuration
    }
}
// Gets the callCount property value. The number of 1:1 calls that the user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetCallCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.callCount
    }
}
// Gets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the hasOtherAction property value. The User is active but has performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetHasOtherAction()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOtherAction
    }
}
// Gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the isLicensed property value. Whether the user has been assigned a Teams license.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetIsLicensed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLicensed
    }
}
// Gets the lastActivityDate property value. The last date that the user participated in a Microsoft Teams activity.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the meetingCount property value. The number of online meetings that the user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetMeetingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingCount
    }
}
// Gets the meetingsAttendedCount property value. The sum of the one-time scheduled, recurring, ad hoc and unclassified meetings a user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingsAttendedCount
    }
}
// Gets the meetingsOrganizedCount property value. The sum of one-time scheduled, Recurring, ad hoc and unclassified meetings a user organized.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetingsOrganizedCount
    }
}
// Gets the privateChatMessageCount property value. The number of unique messages that the user posted in a private chat.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetPrivateChatMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.privateChatMessageCount
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the scheduledOneTimeMeetingsAttendedCount property value. The number of the one-time scheduled meetings a user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetScheduledOneTimeMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledOneTimeMeetingsAttendedCount
    }
}
// Gets the scheduledOneTimeMeetingsOrganizedCount property value. The number of one-time scheduled meetings a user organized.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetScheduledOneTimeMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledOneTimeMeetingsOrganizedCount
    }
}
// Gets the scheduledRecurringMeetingsAttendedCount property value. The number of the recurring meetings a user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetScheduledRecurringMeetingsAttendedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledRecurringMeetingsAttendedCount
    }
}
// Gets the scheduledRecurringMeetingsOrganizedCount property value. The number of recurring meetings a user organized.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetScheduledRecurringMeetingsOrganizedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.scheduledRecurringMeetingsOrganizedCount
    }
}
// Gets the screenShareDuration property value. Screen sharing duration the user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetScreenShareDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.screenShareDuration
    }
}
// Gets the teamChatMessageCount property value. The number of unique messages that the user posted in a team chat.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetTeamChatMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamChatMessageCount
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the videoDuration property value. Video duration the user participated in.
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetVideoDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.videoDuration
    }
}
// The deserialization information for the current model
func (m *GetTeamsUserActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
            res[i] = *(v.(*string))
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
func (m *GetTeamsUserActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetTeamsUserActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the adHocMeetingsAttendedCount property value. The number of ad hoc meetings a user participated in.
// Parameters:
//  - value : Value to set for the adHocMeetingsAttendedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetAdHocMeetingsAttendedCount(value *int64)() {
    m.adHocMeetingsAttendedCount = value
}
// Sets the adHocMeetingsOrganizedCount property value. The number of ad hoc meetings a user organized.
// Parameters:
//  - value : Value to set for the adHocMeetingsOrganizedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetAdHocMeetingsOrganizedCount(value *int64)() {
    m.adHocMeetingsOrganizedCount = value
}
// Sets the assignedProducts property value. Products the user assigned with.
// Parameters:
//  - value : Value to set for the assignedProducts property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
// Sets the audioDuration property value. Audio duration the user participated in.
// Parameters:
//  - value : Value to set for the audioDuration property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetAudioDuration(value *string)() {
    m.audioDuration = value
}
// Sets the callCount property value. The number of 1:1 calls that the user participated in.
// Parameters:
//  - value : Value to set for the callCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetCallCount(value *int64)() {
    m.callCount = value
}
// Sets the deletedDate property value. The date when the delete operation happened. Default value is 'null' when the user has not been deleted.
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the hasOtherAction property value. The User is active but has performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
// Parameters:
//  - value : Value to set for the hasOtherAction property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetHasOtherAction(value *bool)() {
    m.hasOtherAction = value
}
// Sets the isDeleted property value. Whether this user has been deleted or soft deleted.
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the isLicensed property value. Whether the user has been assigned a Teams license.
// Parameters:
//  - value : Value to set for the isLicensed property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetIsLicensed(value *bool)() {
    m.isLicensed = value
}
// Sets the lastActivityDate property value. The last date that the user participated in a Microsoft Teams activity.
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the meetingCount property value. The number of online meetings that the user participated in.
// Parameters:
//  - value : Value to set for the meetingCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetMeetingCount(value *int64)() {
    m.meetingCount = value
}
// Sets the meetingsAttendedCount property value. The sum of the one-time scheduled, recurring, ad hoc and unclassified meetings a user participated in.
// Parameters:
//  - value : Value to set for the meetingsAttendedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetMeetingsAttendedCount(value *int64)() {
    m.meetingsAttendedCount = value
}
// Sets the meetingsOrganizedCount property value. The sum of one-time scheduled, Recurring, ad hoc and unclassified meetings a user organized.
// Parameters:
//  - value : Value to set for the meetingsOrganizedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetMeetingsOrganizedCount(value *int64)() {
    m.meetingsOrganizedCount = value
}
// Sets the privateChatMessageCount property value. The number of unique messages that the user posted in a private chat.
// Parameters:
//  - value : Value to set for the privateChatMessageCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetPrivateChatMessageCount(value *int64)() {
    m.privateChatMessageCount = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the scheduledOneTimeMeetingsAttendedCount property value. The number of the one-time scheduled meetings a user participated in.
// Parameters:
//  - value : Value to set for the scheduledOneTimeMeetingsAttendedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetScheduledOneTimeMeetingsAttendedCount(value *int64)() {
    m.scheduledOneTimeMeetingsAttendedCount = value
}
// Sets the scheduledOneTimeMeetingsOrganizedCount property value. The number of one-time scheduled meetings a user organized.
// Parameters:
//  - value : Value to set for the scheduledOneTimeMeetingsOrganizedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetScheduledOneTimeMeetingsOrganizedCount(value *int64)() {
    m.scheduledOneTimeMeetingsOrganizedCount = value
}
// Sets the scheduledRecurringMeetingsAttendedCount property value. The number of the recurring meetings a user participated in.
// Parameters:
//  - value : Value to set for the scheduledRecurringMeetingsAttendedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetScheduledRecurringMeetingsAttendedCount(value *int64)() {
    m.scheduledRecurringMeetingsAttendedCount = value
}
// Sets the scheduledRecurringMeetingsOrganizedCount property value. The number of recurring meetings a user organized.
// Parameters:
//  - value : Value to set for the scheduledRecurringMeetingsOrganizedCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetScheduledRecurringMeetingsOrganizedCount(value *int64)() {
    m.scheduledRecurringMeetingsOrganizedCount = value
}
// Sets the screenShareDuration property value. Screen sharing duration the user participated in.
// Parameters:
//  - value : Value to set for the screenShareDuration property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetScreenShareDuration(value *string)() {
    m.screenShareDuration = value
}
// Sets the teamChatMessageCount property value. The number of unique messages that the user posted in a team chat.
// Parameters:
//  - value : Value to set for the teamChatMessageCount property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetTeamChatMessageCount(value *int64)() {
    m.teamChatMessageCount = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant’s collection of verified domains. This property is required when a user is created.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the videoDuration property value. Video duration the user participated in.
// Parameters:
//  - value : Value to set for the videoDuration property.
func (m *GetTeamsUserActivityUserDetailWithPeriod) SetVideoDuration(value *string)() {
    m.videoDuration = value
}
