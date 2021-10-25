package getoffice365groupsactivitystoragewithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365GroupsActivityStorageWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    mailboxStorageUsedInBytes *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    siteStorageUsedInBytes *int64;
}
func NewGetOffice365GroupsActivityStorageWithPeriod()(*GetOffice365GroupsActivityStorageWithPeriod) {
    m := &GetOffice365GroupsActivityStorageWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mailboxStorageUsedInBytes
    }
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.siteStorageUsedInBytes
    }
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mailboxStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMailboxStorageUsedInBytes(val)
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
    res["siteStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSiteStorageUsedInBytes(val)
        return nil
    }
    return res
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("mailboxStorageUsedInBytes", m.GetMailboxStorageUsedInBytes())
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
        err = writer.WriteInt64Value("siteStorageUsedInBytes", m.GetSiteStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetMailboxStorageUsedInBytes(value *int64)() {
    m.mailboxStorageUsedInBytes = value
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetSiteStorageUsedInBytes(value *int64)() {
    m.siteStorageUsedInBytes = value
}
