package getoffice365servicesusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365ServicesUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    exchangeActive *int64;
    exchangeInactive *int64;
    office365Active *int64;
    office365Inactive *int64;
    oneDriveActive *int64;
    oneDriveInactive *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    sharePointActive *int64;
    sharePointInactive *int64;
    skypeForBusinessActive *int64;
    skypeForBusinessInactive *int64;
    teamsActive *int64;
    teamsInactive *int64;
    yammerActive *int64;
    yammerInactive *int64;
}
func NewGetOffice365ServicesUserCountsWithPeriod()(*GetOffice365ServicesUserCountsWithPeriod) {
    m := &GetOffice365ServicesUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetExchangeActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetExchangeInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeInactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOffice365Active()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Active
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOffice365Inactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365Inactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOneDriveActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetOneDriveInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDriveInactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSharePointActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSharePointInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointInactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSkypeForBusinessActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetSkypeForBusinessInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusinessInactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetTeamsActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetTeamsInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teamsInactive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetYammerActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerActive
    }
}
func (m *GetOffice365ServicesUserCountsWithPeriod) GetYammerInactive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerInactive
    }
}
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
func (m *GetOffice365ServicesUserCountsWithPeriod) SetExchangeActive(value *int64)() {
    m.exchangeActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetExchangeInactive(value *int64)() {
    m.exchangeInactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOffice365Active(value *int64)() {
    m.office365Active = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOffice365Inactive(value *int64)() {
    m.office365Inactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOneDriveActive(value *int64)() {
    m.oneDriveActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetOneDriveInactive(value *int64)() {
    m.oneDriveInactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSharePointActive(value *int64)() {
    m.sharePointActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSharePointInactive(value *int64)() {
    m.sharePointInactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSkypeForBusinessActive(value *int64)() {
    m.skypeForBusinessActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetSkypeForBusinessInactive(value *int64)() {
    m.skypeForBusinessInactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetTeamsActive(value *int64)() {
    m.teamsActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetTeamsInactive(value *int64)() {
    m.teamsInactive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetYammerActive(value *int64)() {
    m.yammerActive = value
}
func (m *GetOffice365ServicesUserCountsWithPeriod) SetYammerInactive(value *int64)() {
    m.yammerInactive = value
}
