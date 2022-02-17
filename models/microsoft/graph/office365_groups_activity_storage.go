package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityStorage 
type Office365GroupsActivityStorage struct {
    Entity
    // The storage used in group mailbox.
    mailboxStorageUsedInBytes *int64;
    // The snapshot date for Exchange and SharePoint used storage.
    reportDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The storage used in SharePoint document library.
    siteStorageUsedInBytes *int64;
}
// NewOffice365GroupsActivityStorage instantiates a new Office365GroupsActivityStorage and sets the default values.
func NewOffice365GroupsActivityStorage()(*Office365GroupsActivityStorage) {
    m := &Office365GroupsActivityStorage{
        Entity: *NewEntity(),
    }
    return m
}
// GetMailboxStorageUsedInBytes gets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
func (m *Office365GroupsActivityStorage) GetMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mailboxStorageUsedInBytes
    }
}
// GetReportDate gets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
func (m *Office365GroupsActivityStorage) GetReportDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityStorage) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityStorage) GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSiteStorageUsedInBytes gets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
func (m *Office365GroupsActivityStorage) GetSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.siteStorageUsedInBytes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityStorage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["mailboxStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxStorageUsedInBytes(val)
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
    res["siteStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteStorageUsedInBytes(val)
        }
        return nil
    }
    return res
}
func (m *Office365GroupsActivityStorage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityStorage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("siteStorageUsedInBytes", m.GetSiteStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMailboxStorageUsedInBytes sets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
func (m *Office365GroupsActivityStorage) SetMailboxStorageUsedInBytes(value *int64)() {
    if m != nil {
        m.mailboxStorageUsedInBytes = value
    }
}
// SetReportDate sets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
func (m *Office365GroupsActivityStorage) SetReportDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityStorage) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityStorage) SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSiteStorageUsedInBytes sets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
func (m *Office365GroupsActivityStorage) SetSiteStorageUsedInBytes(value *int64)() {
    if m != nil {
        m.siteStorageUsedInBytes = value
    }
}
