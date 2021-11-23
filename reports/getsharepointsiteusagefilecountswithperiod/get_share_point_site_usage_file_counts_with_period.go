package getsharepointsiteusagefilecountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getSharePointSiteUsageFileCountsWithPeriod 
type GetSharePointSiteUsageFileCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    active *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    siteType *string;
    // 
    total *int64;
}
// NewGetSharePointSiteUsageFileCountsWithPeriod instantiates a new getSharePointSiteUsageFileCountsWithPeriod and sets the default values.
func NewGetSharePointSiteUsageFileCountsWithPeriod()(*GetSharePointSiteUsageFileCountsWithPeriod) {
    m := &GetSharePointSiteUsageFileCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActive gets the active property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.active
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSiteType gets the siteType property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetSiteType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteType
    }
}
// GetTotal gets the total property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSharePointSiteUsageFileCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["active"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
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
    res["siteType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteType(val)
        }
        return nil
    }
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
func (m *GetSharePointSiteUsageFileCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSharePointSiteUsageFileCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("active", m.GetActive())
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
        err = writer.WriteStringValue("siteType", m.GetSiteType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetActive(value *int64)() {
    m.active = value
}
// SetReportDate sets the reportDate property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetSiteType sets the siteType property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetSiteType(value *string)() {
    m.siteType = value
}
// SetTotal sets the total property value. 
func (m *GetSharePointSiteUsageFileCountsWithPeriod) SetTotal(value *int64)() {
    m.total = value
}
