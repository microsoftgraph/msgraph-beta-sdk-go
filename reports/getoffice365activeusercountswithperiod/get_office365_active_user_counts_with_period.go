package getoffice365activeusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365ActiveUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    exchange *int64;
    office365 *int64;
    oneDrive *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    sharePoint *int64;
    skypeForBusiness *int64;
    teams *int64;
    yammer *int64;
}
func NewGetOffice365ActiveUserCountsWithPeriod()(*GetOffice365ActiveUserCountsWithPeriod) {
    m := &GetOffice365ActiveUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetExchange()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchange
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOffice365()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.office365
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetOneDrive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.oneDrive
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSharePoint()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePoint
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetSkypeForBusiness()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.skypeForBusiness
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetTeams()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.teams
    }
}
func (m *GetOffice365ActiveUserCountsWithPeriod) GetYammer()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammer
    }
}
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
func (m *GetOffice365ActiveUserCountsWithPeriod) SetExchange(value *int64)() {
    m.exchange = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOffice365(value *int64)() {
    m.office365 = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetOneDrive(value *int64)() {
    m.oneDrive = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSharePoint(value *int64)() {
    m.sharePoint = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetSkypeForBusiness(value *int64)() {
    m.skypeForBusiness = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetTeams(value *int64)() {
    m.teams = value
}
func (m *GetOffice365ActiveUserCountsWithPeriod) SetYammer(value *int64)() {
    m.yammer = value
}
