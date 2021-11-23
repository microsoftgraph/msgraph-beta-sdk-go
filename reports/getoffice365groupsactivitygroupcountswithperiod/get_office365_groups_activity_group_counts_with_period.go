package getoffice365groupsactivitygroupcountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getOffice365GroupsActivityGroupCountsWithPeriod 
type GetOffice365GroupsActivityGroupCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
    active *int64;
    // The date on which a number of groups were active.
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The total number of groups.
    total *int64;
}
// NewGetOffice365GroupsActivityGroupCountsWithPeriod instantiates a new getOffice365GroupsActivityGroupCountsWithPeriod and sets the default values.
func NewGetOffice365GroupsActivityGroupCountsWithPeriod()(*GetOffice365GroupsActivityGroupCountsWithPeriod) {
    m := &GetOffice365GroupsActivityGroupCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetActive gets the active property value. The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetActive()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.active
    }
}
// GetReportDate gets the reportDate property value. The date on which a number of groups were active.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetTotal gets the total property value. The total number of groups.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetTotal()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActive sets the active property value. The number of active groups. A group is considered active if any of the following occurred: group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetActive(value *int64)() {
    m.active = value
}
// SetReportDate sets the reportDate property value. The date on which a number of groups were active.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetTotal sets the total property value. The total number of groups.
func (m *GetOffice365GroupsActivityGroupCountsWithPeriod) SetTotal(value *int64)() {
    m.total = value
}
