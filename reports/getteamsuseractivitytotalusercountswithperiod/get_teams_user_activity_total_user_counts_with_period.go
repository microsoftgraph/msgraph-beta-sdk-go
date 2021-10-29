package getteamsuseractivitytotalusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetTeamsUserActivityTotalUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of users who participated in 1:1 calls.
    calls *int64;
    // The number of users who participated in online meetings.
    meetings *int64;
    // The number of users who were active but performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
    otherActions *int64;
    // The number of users who posted message in a private chat.
    privateChatMessages *int64;
    // The date on which the users performed the activities.
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of users who posted message in a team chat.
    teamChatMessages *int64;
}
// Instantiates a new getTeamsUserActivityTotalUserCountsWithPeriod and sets the default values.
func NewGetTeamsUserActivityTotalUserCountsWithPeriod()(*GetTeamsUserActivityTotalUserCountsWithPeriod) {
    m := &GetTeamsUserActivityTotalUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the calls property value. The number of users who participated in 1:1 calls.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetCalls()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
// Gets the meetings property value. The number of users who participated in online meetings.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetMeetings()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetings
    }
}
// Gets the otherActions property value. The number of users who were active but performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetOtherActions()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.otherActions
    }
}
// Gets the privateChatMessages property value. The number of users who posted message in a private chat.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetPrivateChatMessages()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.privateChatMessages
    }
}
// Gets the reportDate property value. The date on which the users performed the activities.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the teamChatMessages property value. The number of users who posted message in a team chat.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetTeamChatMessages()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamChatMessages
    }
}
// The deserialization information for the current model
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCalls(val)
        return nil
    }
    res["meetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMeetings(val)
        return nil
    }
    res["otherActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOtherActions(val)
        return nil
    }
    res["privateChatMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPrivateChatMessages(val)
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
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
    res["teamChatMessages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTeamChatMessages(val)
        return nil
    }
    return res
}
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("calls", m.GetCalls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("meetings", m.GetMeetings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("otherActions", m.GetOtherActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("privateChatMessages", m.GetPrivateChatMessages())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
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
        err = writer.WriteInt64Value("teamChatMessages", m.GetTeamChatMessages())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the calls property value. The number of users who participated in 1:1 calls.
// Parameters:
//  - value : Value to set for the calls property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetCalls(value *int64)() {
    m.calls = value
}
// Sets the meetings property value. The number of users who participated in online meetings.
// Parameters:
//  - value : Value to set for the meetings property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetMeetings(value *int64)() {
    m.meetings = value
}
// Sets the otherActions property value. The number of users who were active but performed other activities than exposed action types offered in the report (sending or replying to channel messages and chat messages, scheduling or participating in 1:1 calls and meetings). Examples actions are when a user changes the Teams status or the Teams status message or opens a Channel Message post but does not reply.
// Parameters:
//  - value : Value to set for the otherActions property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetOtherActions(value *int64)() {
    m.otherActions = value
}
// Sets the privateChatMessages property value. The number of users who posted message in a private chat.
// Parameters:
//  - value : Value to set for the privateChatMessages property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetPrivateChatMessages(value *int64)() {
    m.privateChatMessages = value
}
// Sets the reportDate property value. The date on which the users performed the activities.
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the teamChatMessages property value. The number of users who posted message in a team chat.
// Parameters:
//  - value : Value to set for the teamChatMessages property.
func (m *GetTeamsUserActivityTotalUserCountsWithPeriod) SetTeamChatMessages(value *int64)() {
    m.teamChatMessages = value
}
