package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityFileCounts 
type Office365GroupsActivityFileCounts struct {
    Entity
    // The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
    active *int64;
    // The date on which a number of files were active in the group's SharePoint site.
    reportDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The total number of files in the group's SharePoint document library.
    total *int64;
}
// NewOffice365GroupsActivityFileCounts instantiates a new Office365GroupsActivityFileCounts and sets the default values.
func NewOffice365GroupsActivityFileCounts()(*Office365GroupsActivityFileCounts) {
    m := &Office365GroupsActivityFileCounts{
        Entity: *NewEntity(),
    }
    return m
}
// GetActive gets the active property value. The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) GetActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.active
    }
}
// GetReportDate gets the reportDate property value. The date on which a number of files were active in the group's SharePoint site.
func (m *Office365GroupsActivityFileCounts) GetReportDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityFileCounts) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityFileCounts) GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetTotal gets the total property value. The total number of files in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityFileCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetDateOnlyValue()
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
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
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
func (m *Office365GroupsActivityFileCounts) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityFileCounts) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("reportDate", m.GetReportDate())
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
        err = writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. The number of files that were viewed, edited, shared, or synced in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) SetActive(value *int64)() {
    if m != nil {
        m.active = value
    }
}
// SetReportDate sets the reportDate property value. The date on which a number of files were active in the group's SharePoint site.
func (m *Office365GroupsActivityFileCounts) SetReportDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityFileCounts) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityFileCounts) SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetTotal sets the total property value. The total number of files in the group's SharePoint document library.
func (m *Office365GroupsActivityFileCounts) SetTotal(value *int64)() {
    if m != nil {
        m.total = value
    }
}
