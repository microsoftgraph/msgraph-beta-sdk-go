package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityCounts 
type Office365GroupsActivityCounts struct {
    Entity
}
// NewOffice365GroupsActivityCounts instantiates a new office365GroupsActivityCounts and sets the default values.
func NewOffice365GroupsActivityCounts()(*Office365GroupsActivityCounts) {
    m := &Office365GroupsActivityCounts{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityCounts(), nil
}
// GetExchangeEmailsReceived gets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
func (m *Office365GroupsActivityCounts) GetExchangeEmailsReceived()(*int64) {
    val, err := m.GetBackingStore().Get("exchangeEmailsReceived")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeEmailsReceived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeEmailsReceived(val)
        }
        return nil
    }
    res["reportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["teamsChannelMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsChannelMessages(val)
        }
        return nil
    }
    res["teamsMeetingsOrganized"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsMeetingsOrganized(val)
        }
        return nil
    }
    res["yammerMessagesLiked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesLiked(val)
        }
        return nil
    }
    res["yammerMessagesPosted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesPosted(val)
        }
        return nil
    }
    res["yammerMessagesRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesRead(val)
        }
        return nil
    }
    return res
}
// GetReportDate gets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
func (m *Office365GroupsActivityCounts) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("reportDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityCounts) GetReportPeriod()(*string) {
    val, err := m.GetBackingStore().Get("reportPeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityCounts) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("reportRefreshDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetTeamsChannelMessages gets the teamsChannelMessages property value. The number of channel messages in Teams team.
func (m *Office365GroupsActivityCounts) GetTeamsChannelMessages()(*int64) {
    val, err := m.GetBackingStore().Get("teamsChannelMessages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTeamsMeetingsOrganized gets the teamsMeetingsOrganized property value. The number of meetings organized in Teams team.
func (m *Office365GroupsActivityCounts) GetTeamsMeetingsOrganized()(*int64) {
    val, err := m.GetBackingStore().Get("teamsMeetingsOrganized")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerMessagesLiked gets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesLiked()(*int64) {
    val, err := m.GetBackingStore().Get("yammerMessagesLiked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerMessagesPosted gets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesPosted()(*int64) {
    val, err := m.GetBackingStore().Get("yammerMessagesPosted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerMessagesRead gets the yammerMessagesRead property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityCounts) GetYammerMessagesRead()(*int64) {
    val, err := m.GetBackingStore().Get("yammerMessagesRead")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeEmailsReceived", m.GetExchangeEmailsReceived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportDate", m.GetReportDate())
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
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsChannelMessages", m.GetTeamsChannelMessages())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsMeetingsOrganized", m.GetTeamsMeetingsOrganized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesLiked", m.GetYammerMessagesLiked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesPosted", m.GetYammerMessagesPosted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesRead", m.GetYammerMessagesRead())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchangeEmailsReceived sets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
func (m *Office365GroupsActivityCounts) SetExchangeEmailsReceived(value *int64)() {
    err := m.GetBackingStore().Set("exchangeEmailsReceived", value)
    if err != nil {
        panic(err)
    }
}
// SetReportDate sets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
func (m *Office365GroupsActivityCounts) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportDate", value)
    if err != nil {
        panic(err)
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityCounts) SetReportPeriod(value *string)() {
    err := m.GetBackingStore().Set("reportPeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityCounts) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportRefreshDate", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsChannelMessages sets the teamsChannelMessages property value. The number of channel messages in Teams team.
func (m *Office365GroupsActivityCounts) SetTeamsChannelMessages(value *int64)() {
    err := m.GetBackingStore().Set("teamsChannelMessages", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsMeetingsOrganized sets the teamsMeetingsOrganized property value. The number of meetings organized in Teams team.
func (m *Office365GroupsActivityCounts) SetTeamsMeetingsOrganized(value *int64)() {
    err := m.GetBackingStore().Set("teamsMeetingsOrganized", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerMessagesLiked sets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesLiked(value *int64)() {
    err := m.GetBackingStore().Set("yammerMessagesLiked", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerMessagesPosted sets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesPosted(value *int64)() {
    err := m.GetBackingStore().Set("yammerMessagesPosted", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerMessagesRead sets the yammerMessagesRead property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityCounts) SetYammerMessagesRead(value *int64)() {
    err := m.GetBackingStore().Set("yammerMessagesRead", value)
    if err != nil {
        panic(err)
    }
}
// Office365GroupsActivityCountsable 
type Office365GroupsActivityCountsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchangeEmailsReceived()(*int64)
    GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetTeamsChannelMessages()(*int64)
    GetTeamsMeetingsOrganized()(*int64)
    GetYammerMessagesLiked()(*int64)
    GetYammerMessagesPosted()(*int64)
    GetYammerMessagesRead()(*int64)
    SetExchangeEmailsReceived(value *int64)()
    SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetTeamsChannelMessages(value *int64)()
    SetTeamsMeetingsOrganized(value *int64)()
    SetYammerMessagesLiked(value *int64)()
    SetYammerMessagesPosted(value *int64)()
    SetYammerMessagesRead(value *int64)()
}
