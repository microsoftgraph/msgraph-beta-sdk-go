package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365ServicesUserCounts 
type Office365ServicesUserCounts struct {
    Entity
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
    reportRefreshDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
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
// NewOffice365ServicesUserCounts instantiates a new Office365ServicesUserCounts and sets the default values.
func NewOffice365ServicesUserCounts()(*Office365ServicesUserCounts) {
    m := &Office365ServicesUserCounts{
        Entity: *NewEntity(),
    }
    return m
}
// GetExchangeActive gets the exchangeActive property value. The number of active users on Exchange. Any user who can read and send email is considered an active user.
func (m *Office365ServicesUserCounts) GetExchangeActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeActive
    }
}
// GetExchangeInactive gets the exchangeInactive property value. The number of inactive users on Exchange.
func (m *Office365ServicesUserCounts) GetExchangeInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeInactive
    }
}
// GetOffice365Active gets the office365Active property value. The number of active users on Microsoft 365.
func (m *Office365ServicesUserCounts) GetOffice365Active()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Active
    }
}
// GetOffice365Inactive gets the office365Inactive property value. The number of inactive users on Microsoft 365.
func (m *Office365ServicesUserCounts) GetOffice365Inactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Inactive
    }
}
// GetOneDriveActive gets the oneDriveActive property value. The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *Office365ServicesUserCounts) GetOneDriveActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveActive
    }
}
// GetOneDriveInactive gets the oneDriveInactive property value. The number of inactive users on OneDrive.
func (m *Office365ServicesUserCounts) GetOneDriveInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveInactive
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365ServicesUserCounts) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ServicesUserCounts) GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharePointActive gets the sharePointActive property value. The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *Office365ServicesUserCounts) GetSharePointActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActive
    }
}
// GetSharePointInactive gets the sharePointInactive property value. The number of inactive users on SharePoint.
func (m *Office365ServicesUserCounts) GetSharePointInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointInactive
    }
}
// GetSkypeForBusinessActive gets the skypeForBusinessActive property value. The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *Office365ServicesUserCounts) GetSkypeForBusinessActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessActive
    }
}
// GetSkypeForBusinessInactive gets the skypeForBusinessInactive property value. The number of inactive users on Skype For Business.
func (m *Office365ServicesUserCounts) GetSkypeForBusinessInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessInactive
    }
}
// GetTeamsActive gets the teamsActive property value. The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *Office365ServicesUserCounts) GetTeamsActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsActive
    }
}
// GetTeamsInactive gets the teamsInactive property value. The number of inactive users on Microsoft Teams.
func (m *Office365ServicesUserCounts) GetTeamsInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsInactive
    }
}
// GetYammerActive gets the yammerActive property value. The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *Office365ServicesUserCounts) GetYammerActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerActive
    }
}
// GetYammerInactive gets the yammerInactive property value. The number of inactive users on Yammer.
func (m *Office365ServicesUserCounts) GetYammerInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerInactive
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365ServicesUserCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeActive(val)
        }
        return nil
    }
    res["exchangeInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeInactive(val)
        }
        return nil
    }
    res["office365Active"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffice365Active(val)
        }
        return nil
    }
    res["office365Inactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffice365Inactive(val)
        }
        return nil
    }
    res["oneDriveActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveActive(val)
        }
        return nil
    }
    res["oneDriveInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveInactive(val)
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
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointActive(val)
        }
        return nil
    }
    res["sharePointInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointInactive(val)
        }
        return nil
    }
    res["skypeForBusinessActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessActive(val)
        }
        return nil
    }
    res["skypeForBusinessInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeForBusinessInactive(val)
        }
        return nil
    }
    res["teamsActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsActive(val)
        }
        return nil
    }
    res["teamsInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsInactive(val)
        }
        return nil
    }
    res["yammerActive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerActive(val)
        }
        return nil
    }
    res["yammerInactive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerInactive(val)
        }
        return nil
    }
    return res
}
func (m *Office365ServicesUserCounts) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Office365ServicesUserCounts) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
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
// SetExchangeActive sets the exchangeActive property value. The number of active users on Exchange. Any user who can read and send email is considered an active user.
func (m *Office365ServicesUserCounts) SetExchangeActive(value *int64)() {
    if m != nil {
        m.exchangeActive = value
    }
}
// SetExchangeInactive sets the exchangeInactive property value. The number of inactive users on Exchange.
func (m *Office365ServicesUserCounts) SetExchangeInactive(value *int64)() {
    if m != nil {
        m.exchangeInactive = value
    }
}
// SetOffice365Active sets the office365Active property value. The number of active users on Microsoft 365.
func (m *Office365ServicesUserCounts) SetOffice365Active(value *int64)() {
    if m != nil {
        m.office365Active = value
    }
}
// SetOffice365Inactive sets the office365Inactive property value. The number of inactive users on Microsoft 365.
func (m *Office365ServicesUserCounts) SetOffice365Inactive(value *int64)() {
    if m != nil {
        m.office365Inactive = value
    }
}
// SetOneDriveActive sets the oneDriveActive property value. The number of active users on OneDrive. Any user who viewed or edited files, shared files internally or externally, or synced files is considered an active user.
func (m *Office365ServicesUserCounts) SetOneDriveActive(value *int64)() {
    if m != nil {
        m.oneDriveActive = value
    }
}
// SetOneDriveInactive sets the oneDriveInactive property value. The number of inactive users on OneDrive.
func (m *Office365ServicesUserCounts) SetOneDriveInactive(value *int64)() {
    if m != nil {
        m.oneDriveInactive = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365ServicesUserCounts) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365ServicesUserCounts) SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharePointActive sets the sharePointActive property value. The number of active users on SharePoint. Any user who viewed or edited files, shared files internally or externally, synced files, or viewed SharePoint pages is considered an active user.
func (m *Office365ServicesUserCounts) SetSharePointActive(value *int64)() {
    if m != nil {
        m.sharePointActive = value
    }
}
// SetSharePointInactive sets the sharePointInactive property value. The number of inactive users on SharePoint.
func (m *Office365ServicesUserCounts) SetSharePointInactive(value *int64)() {
    if m != nil {
        m.sharePointInactive = value
    }
}
// SetSkypeForBusinessActive sets the skypeForBusinessActive property value. The number of active users on Skype For Business. Any user who organized or participated in conferences, or joined peer-to-peer sessions is considered an active user.
func (m *Office365ServicesUserCounts) SetSkypeForBusinessActive(value *int64)() {
    if m != nil {
        m.skypeForBusinessActive = value
    }
}
// SetSkypeForBusinessInactive sets the skypeForBusinessInactive property value. The number of inactive users on Skype For Business.
func (m *Office365ServicesUserCounts) SetSkypeForBusinessInactive(value *int64)() {
    if m != nil {
        m.skypeForBusinessInactive = value
    }
}
// SetTeamsActive sets the teamsActive property value. The number of active users on Microsoft Teams. Any user who posted messages in team channels, sent messages in private chat sessions, or participated in meetings or calls is considered an active user.
func (m *Office365ServicesUserCounts) SetTeamsActive(value *int64)() {
    if m != nil {
        m.teamsActive = value
    }
}
// SetTeamsInactive sets the teamsInactive property value. The number of inactive users on Microsoft Teams.
func (m *Office365ServicesUserCounts) SetTeamsInactive(value *int64)() {
    if m != nil {
        m.teamsInactive = value
    }
}
// SetYammerActive sets the yammerActive property value. The number of active users on Yammer. Any user who can post, read, or like messages is considered an active user.
func (m *Office365ServicesUserCounts) SetYammerActive(value *int64)() {
    if m != nil {
        m.yammerActive = value
    }
}
// SetYammerInactive sets the yammerInactive property value. The number of inactive users on Yammer.
func (m *Office365ServicesUserCounts) SetYammerInactive(value *int64)() {
    if m != nil {
        m.yammerInactive = value
    }
}
