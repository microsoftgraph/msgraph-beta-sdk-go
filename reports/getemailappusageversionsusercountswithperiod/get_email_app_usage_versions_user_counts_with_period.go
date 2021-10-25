package getemailappusageversionsusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetEmailAppUsageVersionsUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    outlook2007 *int64;
    outlook2010 *int64;
    outlook2013 *int64;
    outlook2016 *int64;
    outlook2019 *int64;
    outlookM365 *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    undetermined *int64;
}
func NewGetEmailAppUsageVersionsUserCountsWithPeriod()(*GetEmailAppUsageVersionsUserCountsWithPeriod) {
    m := &GetEmailAppUsageVersionsUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlook2007()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlook2007
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlook2010()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlook2010
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlook2013()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlook2013
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlook2016()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlook2016
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlook2019()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlook2019
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetOutlookM365()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookM365
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetUndetermined()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.undetermined
    }
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["outlook2007"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlook2007(val)
        return nil
    }
    res["outlook2010"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlook2010(val)
        return nil
    }
    res["outlook2013"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlook2013(val)
        return nil
    }
    res["outlook2016"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlook2016(val)
        return nil
    }
    res["outlook2019"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlook2019(val)
        return nil
    }
    res["outlookM365"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlookM365(val)
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
    res["undetermined"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUndetermined(val)
        return nil
    }
    return res
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("outlook2007", m.GetOutlook2007())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlook2010", m.GetOutlook2010())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlook2013", m.GetOutlook2013())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlook2016", m.GetOutlook2016())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlook2019", m.GetOutlook2019())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlookM365", m.GetOutlookM365())
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
        err = writer.WriteInt64Value("undetermined", m.GetUndetermined())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlook2007(value *int64)() {
    m.outlook2007 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlook2010(value *int64)() {
    m.outlook2010 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlook2013(value *int64)() {
    m.outlook2013 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlook2016(value *int64)() {
    m.outlook2016 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlook2019(value *int64)() {
    m.outlook2019 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetOutlookM365(value *int64)() {
    m.outlookM365 = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetEmailAppUsageVersionsUserCountsWithPeriod) SetUndetermined(value *int64)() {
    m.undetermined = value
}
