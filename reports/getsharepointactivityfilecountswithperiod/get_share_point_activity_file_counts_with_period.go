package getsharepointactivityfilecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSharePointActivityFileCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    sharedExternally *int64;
    sharedInternally *int64;
    synced *int64;
    viewedOrEdited *int64;
}
func NewGetSharePointActivityFileCountsWithPeriod()(*GetSharePointActivityFileCountsWithPeriod) {
    m := &GetSharePointActivityFileCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetSharedExternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedExternally
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetSharedInternally()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharedInternally
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetSynced()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.synced
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetViewedOrEdited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.viewedOrEdited
    }
}
func (m *GetSharePointActivityFileCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["sharedExternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedExternally(val)
        return nil
    }
    res["sharedInternally"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharedInternally(val)
        return nil
    }
    res["synced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSynced(val)
        return nil
    }
    res["viewedOrEdited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetViewedOrEdited(val)
        return nil
    }
    return res
}
func (m *GetSharePointActivityFileCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSharePointActivityFileCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteInt64Value("sharedExternally", m.GetSharedExternally())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharedInternally", m.GetSharedInternally())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("synced", m.GetSynced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("viewedOrEdited", m.GetViewedOrEdited())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetSharedExternally(value *int64)() {
    m.sharedExternally = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetSharedInternally(value *int64)() {
    m.sharedInternally = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetSynced(value *int64)() {
    m.synced = value
}
func (m *GetSharePointActivityFileCountsWithPeriod) SetViewedOrEdited(value *int64)() {
    m.viewedOrEdited = value
}
