package getoffice365activeusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365ActiveUserCountsWithPeriod 
type GetOffice365ActiveUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of active users in Exchange. Any user who can read and send email is considered an active user.
    exchange *int64;
    // The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
    office365 *int64;
    // The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
    oneDrive *int64;
    // The date on which a number of users were active.
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
    sharePoint *int64;
    // The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
    skypeForBusiness *int64;
    // The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
    teams *int64;
    // The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
    yammer *int64;
}
// NewGetOffice365ActiveUserCountsWithPeriod instantiates a new getOffice365ActiveUserCountsWithPeriod and sets the default values.
func NewGetOffice365ActiveUserCountsWithPeriod()(*GetOffice365ActiveUserCountsWithPeriod) {
    m := &GetOffice365ActiveUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetExchange gets the exchange property value. The number of active users in Exchange. Any user who can read and send email is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetExchange()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchange
    }
}
// GetOffice365 gets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOffice365()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365
    }
}
// GetOneDrive gets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOneDrive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDrive
    }
}
// GetReportDate gets the reportDate property value. The date on which a number of users were active.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharePoint gets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSharePoint()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePoint
    }
}
// GetSkypeForBusiness gets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSkypeForBusiness()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusiness
    }
}
// GetTeams gets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetTeams()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teams
    }
}
// GetYammer gets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetYammer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammer
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365ActiveUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchange(val)
        }
        return nil
    }
    res["office365"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffice365(val)
        }
        return nil
    }
    res["oneDrive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDrive(val)
        }
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePoint(val)
        }
        return nil
    }
    res["skypeForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusiness(val)
        }
        return nil
    }
    res["teams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeams(val)
        }
        return nil
    }
    res["yammer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *GetOffice365ActiveUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOffice365ActiveUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetOffice365ActiveUserCountsWithPeriod) SetExchange(value *int64)() {
    if m != nil {
        m.exchange = value
    }
}
// SetOffice365 sets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOffice365(value *int64)() {
    if m != nil {
        m.office365 = value
    }
}
// SetOneDrive sets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOneDrive(value *int64)() {
    if m != nil {
        m.oneDrive = value
    }
}
// SetReportDate sets the reportDate property value. The date on which a number of users were active.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportDate(value *string)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharePoint sets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSharePoint(value *int64)() {
    if m != nil {
        m.sharePoint = value
    }
}
// SetSkypeForBusiness sets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSkypeForBusiness(value *int64)() {
    if m != nil {
        m.skypeForBusiness = value
    }
}
// SetTeams sets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetTeams(value *int64)() {
    if m != nil {
        m.teams = value
    }
}
// SetYammer sets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetYammer(value *int64)() {
    if m != nil {
        m.yammer = value
    }
}
