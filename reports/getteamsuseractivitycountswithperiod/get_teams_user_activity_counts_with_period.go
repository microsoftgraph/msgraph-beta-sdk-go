package getteamsuseractivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetTeamsUserActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    calls *int64;
    meetings *int64;
    privateChatMessages *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    teamChatMessages *int64;
}
func NewGetTeamsUserActivityCountsWithPeriod()(*GetTeamsUserActivityCountsWithPeriod) {
    m := &GetTeamsUserActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetCalls()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetMeetings()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.meetings
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetPrivateChatMessages()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.privateChatMessages
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetTeamChatMessages()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamChatMessages
    }
}
func (m *GetTeamsUserActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetTeamsUserActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetTeamsUserActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetTeamsUserActivityCountsWithPeriod) SetCalls(value *int64)() {
    m.calls = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetMeetings(value *int64)() {
    m.meetings = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetPrivateChatMessages(value *int64)() {
    m.privateChatMessages = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetTeamsUserActivityCountsWithPeriod) SetTeamChatMessages(value *int64)() {
    m.teamChatMessages = value
}
