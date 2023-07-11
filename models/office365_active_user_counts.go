package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365ActiveUserCounts 
type Office365ActiveUserCounts struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewOffice365ActiveUserCounts instantiates a new office365ActiveUserCounts and sets the default values.
func NewOffice365ActiveUserCounts()(*Office365ActiveUserCounts) {
    m := &Office365ActiveUserCounts{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365ActiveUserCountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365ActiveUserCountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365ActiveUserCounts(), nil
}
// GetExchange gets the exchange property value. The number of active users in Exchange. Any user who can read and send email is considered an active user.
func (m *Office365ActiveUserCounts) GetExchange()(*int64) {
    val, err := m.GetBackingStore().Get("exchange")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365ActiveUserCounts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchange(val)
        }
        return nil
    }
    res["office365"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffice365(val)
        }
        return nil
    }
    res["oneDrive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDrive(val)
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
    res["sharePoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePoint(val)
        }
        return nil
    }
    res["skypeForBusiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusiness(val)
        }
        return nil
    }
    res["teams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeams(val)
        }
        return nil
    }
    res["yammer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammer(val)
        }
        return nil
    }
    return res
}
// GetOffice365 gets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
func (m *Office365ActiveUserCounts) GetOffice365()(*int64) {
    val, err := m.GetBackingStore().Get("office365")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetOneDrive gets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *Office365ActiveUserCounts) GetOneDrive()(*int64) {
    val, err := m.GetBackingStore().Get("oneDrive")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetReportDate gets the reportDate property value. The date on which a number of users were active.
func (m *Office365ActiveUserCounts) GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
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
func (m *Office365ActiveUserCounts) GetReportPeriod()(*string) {
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
func (m *Office365ActiveUserCounts) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("reportRefreshDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSharePoint gets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *Office365ActiveUserCounts) GetSharePoint()(*int64) {
    val, err := m.GetBackingStore().Get("sharePoint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSkypeForBusiness gets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *Office365ActiveUserCounts) GetSkypeForBusiness()(*int64) {
    val, err := m.GetBackingStore().Get("skypeForBusiness")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTeams gets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *Office365ActiveUserCounts) GetTeams()(*int64) {
    val, err := m.GetBackingStore().Get("teams")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammer gets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *Office365ActiveUserCounts) GetYammer()(*int64) {
    val, err := m.GetBackingStore().Get("yammer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Office365ActiveUserCounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchange", m.GetExchange())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("office365", m.GetOffice365())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("oneDrive", m.GetOneDrive())
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
        err = writer.WriteInt64Value("sharePoint", m.GetSharePoint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("skypeForBusiness", m.GetSkypeForBusiness())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teams", m.GetTeams())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammer", m.GetYammer())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchange sets the exchange property value. The number of active users in Exchange. Any user who can read and send email is considered an active user.
func (m *Office365ActiveUserCounts) SetExchange(value *int64)() {
    err := m.GetBackingStore().Set("exchange", value)
    if err != nil {
        panic(err)
    }
}
// SetOffice365 sets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
func (m *Office365ActiveUserCounts) SetOffice365(value *int64)() {
    err := m.GetBackingStore().Set("office365", value)
    if err != nil {
        panic(err)
    }
}
// SetOneDrive sets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *Office365ActiveUserCounts) SetOneDrive(value *int64)() {
    err := m.GetBackingStore().Set("oneDrive", value)
    if err != nil {
        panic(err)
    }
}
// SetReportDate sets the reportDate property value. The date on which a number of users were active.
func (m *Office365ActiveUserCounts) SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportDate", value)
    if err != nil {
        panic(err)
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365ActiveUserCounts) SetReportPeriod(value *string)() {
    err := m.GetBackingStore().Set("reportPeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ActiveUserCounts) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportRefreshDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePoint sets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *Office365ActiveUserCounts) SetSharePoint(value *int64)() {
    err := m.GetBackingStore().Set("sharePoint", value)
    if err != nil {
        panic(err)
    }
}
// SetSkypeForBusiness sets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *Office365ActiveUserCounts) SetSkypeForBusiness(value *int64)() {
    err := m.GetBackingStore().Set("skypeForBusiness", value)
    if err != nil {
        panic(err)
    }
}
// SetTeams sets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *Office365ActiveUserCounts) SetTeams(value *int64)() {
    err := m.GetBackingStore().Set("teams", value)
    if err != nil {
        panic(err)
    }
}
// SetYammer sets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *Office365ActiveUserCounts) SetYammer(value *int64)() {
    err := m.GetBackingStore().Set("yammer", value)
    if err != nil {
        panic(err)
    }
}
// Office365ActiveUserCountsable 
type Office365ActiveUserCountsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchange()(*int64)
    GetOffice365()(*int64)
    GetOneDrive()(*int64)
    GetReportDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePoint()(*int64)
    GetSkypeForBusiness()(*int64)
    GetTeams()(*int64)
    GetYammer()(*int64)
    SetExchange(value *int64)()
    SetOffice365(value *int64)()
    SetOneDrive(value *int64)()
    SetReportDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePoint(value *int64)()
    SetSkypeForBusiness(value *int64)()
    SetTeams(value *int64)()
    SetYammer(value *int64)()
}
