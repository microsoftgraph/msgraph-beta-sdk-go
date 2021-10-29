package getoffice365activeusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new getOffice365ActiveUserCountsWithPeriod and sets the default values.
func NewGetOffice365ActiveUserCountsWithPeriod()(*GetOffice365ActiveUserCountsWithPeriod) {
    m := &GetOffice365ActiveUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the exchange property value. The number of active users in Exchange. Any user who can read and send email is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetExchange()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchange
    }
}
// Gets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOffice365()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365
    }
}
// Gets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOneDrive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDrive
    }
}
// Gets the reportDate property value. The date on which a number of users were active.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSharePoint()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePoint
    }
}
// Gets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSkypeForBusiness()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusiness
    }
}
// Gets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetTeams()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teams
    }
}
// Gets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *GetOffice365ActiveUserCountsWithPeriod) GetYammer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammer
    }
}
// The deserialization information for the current model
func (m *GetOffice365ActiveUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchange(val)
        return nil
    }
    res["office365"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOffice365(val)
        return nil
    }
    res["oneDrive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOneDrive(val)
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
    res["sharePoint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePoint(val)
        return nil
    }
    res["skypeForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSkypeForBusiness(val)
        return nil
    }
    res["teams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTeams(val)
        return nil
    }
    res["yammer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammer(val)
        return nil
    }
    return res
}
func (m *GetOffice365ActiveUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the exchange property value. The number of active users in Exchange. Any user who can read and send email is considered an active user.
// Parameters:
//  - value : Value to set for the exchange property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetExchange(value *int64)() {
    m.exchange = value
}
// Sets the office365 property value. The number of active users in Microsoft 365. This number includes all the active users in Exchange, OneDrive, SharePoint, Skype For Business, Yammer, and Microsoft Teams. You can find the definition of active user for each product in the respective property description.
// Parameters:
//  - value : Value to set for the office365 property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOffice365(value *int64)() {
    m.office365 = value
}
// Sets the oneDrive property value. The number of active users in OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
// Parameters:
//  - value : Value to set for the oneDrive property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOneDrive(value *int64)() {
    m.oneDrive = value
}
// Sets the reportDate property value. The date on which a number of users were active.
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharePoint property value. The number of active users in SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
// Parameters:
//  - value : Value to set for the sharePoint property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSharePoint(value *int64)() {
    m.sharePoint = value
}
// Sets the skypeForBusiness property value. The number of active users in Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
// Parameters:
//  - value : Value to set for the skypeForBusiness property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSkypeForBusiness(value *int64)() {
    m.skypeForBusiness = value
}
// Sets the teams property value. The number of active users in Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
// Parameters:
//  - value : Value to set for the teams property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetTeams(value *int64)() {
    m.teams = value
}
// Sets the yammer property value. The number of active users in Yammer. Any user who can post, read, or like messages is considered an active user.
// Parameters:
//  - value : Value to set for the yammer property.
func (m *GetOffice365ActiveUserCountsWithPeriod) SetYammer(value *int64)() {
    m.yammer = value
}
