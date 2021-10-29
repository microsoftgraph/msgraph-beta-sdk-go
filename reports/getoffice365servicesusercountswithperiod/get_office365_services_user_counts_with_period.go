package getoffice365servicesusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOffice365ServicesUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of active users on Exchange. Any user who can read and send email is considered an active user.
    exchangeActive *int64;
    // The number of inactive users on Exchange.
    exchangeInactive *int64;
    // The number of active users on Microsoft 365.
    office365Active *int64;
    // The number of inactive users on Microsoft 365.
    office365Inactive *int64;
    // The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
    oneDriveActive *int64;
    // The number of inactive users on OneDrive.
    oneDriveInactive *int64;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
    sharePointActive *int64;
    // The number of inactive users on SharePoint.
    sharePointInactive *int64;
    // The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
    skypeForBusinessActive *int64;
    // The number of inactive users on Skype For Business.
    skypeForBusinessInactive *int64;
    // The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
    teamsActive *int64;
    // The number of inactive users on Microsoft Teams.
    teamsInactive *int64;
    // The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
    yammerActive *int64;
    // The number of inactive users on Yammer.
    yammerInactive *int64;
}
// Instantiates a new getOffice365ServicesUserCountsWithPeriod and sets the default values.
func NewGetOffice365ServicesUserCountsWithPeriod()(*GetOffice365ServicesUserCountsWithPeriod) {
    m := &GetOffice365ServicesUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the exchangeActive property value. The number of active users on Exchange. Any user who can read and send email is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetExchangeActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeActive
    }
}
// Gets the exchangeInactive property value. The number of inactive users on Exchange.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetExchangeInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeInactive
    }
}
// Gets the office365Active property value. The number of active users on Microsoft 365.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOffice365Active()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Active
    }
}
// Gets the office365Inactive property value. The number of inactive users on Microsoft 365.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOffice365Inactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Inactive
    }
}
// Gets the oneDriveActive property value. The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOneDriveActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveActive
    }
}
// Gets the oneDriveInactive property value. The number of inactive users on OneDrive.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOneDriveInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveInactive
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharePointActive property value. The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSharePointActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActive
    }
}
// Gets the sharePointInactive property value. The number of inactive users on SharePoint.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSharePointInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointInactive
    }
}
// Gets the skypeForBusinessActive property value. The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSkypeForBusinessActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessActive
    }
}
// Gets the skypeForBusinessInactive property value. The number of inactive users on Skype For Business.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSkypeForBusinessInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessInactive
    }
}
// Gets the teamsActive property value. The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetTeamsActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsActive
    }
}
// Gets the teamsInactive property value. The number of inactive users on Microsoft Teams.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetTeamsInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsInactive
    }
}
// Gets the yammerActive property value. The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetYammerActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerActive
    }
}
// Gets the yammerInactive property value. The number of inactive users on Yammer.
func (m *GetOffice365ServicesUserCountsWithPeriod) GetYammerInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerInactive
    }
}
// The deserialization information for the current model
func (m *GetOffice365ServicesUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeActive(val)
        return nil
    }
    res["exchangeInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeInactive(val)
        return nil
    }
    res["office365Active"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOffice365Active(val)
        return nil
    }
    res["office365Inactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOffice365Inactive(val)
        return nil
    }
    res["oneDriveActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOneDriveActive(val)
        return nil
    }
    res["oneDriveInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOneDriveInactive(val)
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
    res["sharePointActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePointActive(val)
        return nil
    }
    res["sharePointInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePointInactive(val)
        return nil
    }
    res["skypeForBusinessActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSkypeForBusinessActive(val)
        return nil
    }
    res["skypeForBusinessInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSkypeForBusinessInactive(val)
        return nil
    }
    res["teamsActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTeamsActive(val)
        return nil
    }
    res["teamsInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTeamsInactive(val)
        return nil
    }
    res["yammerActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerActive(val)
        return nil
    }
    res["yammerInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerInactive(val)
        return nil
    }
    return res
}
func (m *GetOffice365ServicesUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOffice365ServicesUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeActive", m.GetExchangeActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("exchangeInactive", m.GetExchangeInactive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("office365Active", m.GetOffice365Active())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("office365Inactive", m.GetOffice365Inactive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("oneDriveActive", m.GetOneDriveActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("oneDriveInactive", m.GetOneDriveInactive())
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
        err = writer.WriteInt64Value("sharePointActive", m.GetSharePointActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointInactive", m.GetSharePointInactive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("skypeForBusinessActive", m.GetSkypeForBusinessActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("skypeForBusinessInactive", m.GetSkypeForBusinessInactive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsActive", m.GetTeamsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsInactive", m.GetTeamsInactive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerActive", m.GetYammerActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerInactive", m.GetYammerInactive())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the exchangeActive property value. The number of active users on Exchange. Any user who can read and send email is considered an active user.
// Parameters:
//  - value : Value to set for the exchangeActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetExchangeActive(value *int64)() {
    m.exchangeActive = value
}
// Sets the exchangeInactive property value. The number of inactive users on Exchange.
// Parameters:
//  - value : Value to set for the exchangeInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetExchangeInactive(value *int64)() {
    m.exchangeInactive = value
}
// Sets the office365Active property value. The number of active users on Microsoft 365.
// Parameters:
//  - value : Value to set for the office365Active property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOffice365Active(value *int64)() {
    m.office365Active = value
}
// Sets the office365Inactive property value. The number of inactive users on Microsoft 365.
// Parameters:
//  - value : Value to set for the office365Inactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOffice365Inactive(value *int64)() {
    m.office365Inactive = value
}
// Sets the oneDriveActive property value. The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
// Parameters:
//  - value : Value to set for the oneDriveActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOneDriveActive(value *int64)() {
    m.oneDriveActive = value
}
// Sets the oneDriveInactive property value. The number of inactive users on OneDrive.
// Parameters:
//  - value : Value to set for the oneDriveInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOneDriveInactive(value *int64)() {
    m.oneDriveInactive = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharePointActive property value. The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
// Parameters:
//  - value : Value to set for the sharePointActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSharePointActive(value *int64)() {
    m.sharePointActive = value
}
// Sets the sharePointInactive property value. The number of inactive users on SharePoint.
// Parameters:
//  - value : Value to set for the sharePointInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSharePointInactive(value *int64)() {
    m.sharePointInactive = value
}
// Sets the skypeForBusinessActive property value. The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
// Parameters:
//  - value : Value to set for the skypeForBusinessActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSkypeForBusinessActive(value *int64)() {
    m.skypeForBusinessActive = value
}
// Sets the skypeForBusinessInactive property value. The number of inactive users on Skype For Business.
// Parameters:
//  - value : Value to set for the skypeForBusinessInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSkypeForBusinessInactive(value *int64)() {
    m.skypeForBusinessInactive = value
}
// Sets the teamsActive property value. The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
// Parameters:
//  - value : Value to set for the teamsActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetTeamsActive(value *int64)() {
    m.teamsActive = value
}
// Sets the teamsInactive property value. The number of inactive users on Microsoft Teams.
// Parameters:
//  - value : Value to set for the teamsInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetTeamsInactive(value *int64)() {
    m.teamsInactive = value
}
// Sets the yammerActive property value. The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
// Parameters:
//  - value : Value to set for the yammerActive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetYammerActive(value *int64)() {
    m.yammerActive = value
}
// Sets the yammerInactive property value. The number of inactive users on Yammer.
// Parameters:
//  - value : Value to set for the yammerInactive property.
func (m *GetOffice365ServicesUserCountsWithPeriod) SetYammerInactive(value *int64)() {
    m.yammerInactive = value
}
