package getyammergroupsactivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetYammerGroupsActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    liked *int64;
    posted *int64;
    read *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
}
func NewGetYammerGroupsActivityCountsWithPeriod()(*GetYammerGroupsActivityCountsWithPeriod) {
    m := &GetYammerGroupsActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.liked
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.posted
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.read
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetYammerGroupsActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["liked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLiked(val)
        return nil
    }
    res["posted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPosted(val)
        return nil
    }
    res["read"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRead(val)
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
    return res
}
func (m *GetYammerGroupsActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetYammerGroupsActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("liked", m.GetLiked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("posted", m.GetPosted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("read", m.GetRead())
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
    return nil
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetLiked(value *int64)() {
    m.liked = value
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetPosted(value *int64)() {
    m.posted = value
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetRead(value *int64)() {
    m.read = value
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetYammerGroupsActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
