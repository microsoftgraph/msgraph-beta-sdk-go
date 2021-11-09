package getoffice365groupsactivitystoragewithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOffice365GroupsActivityStorageWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The storage used in group mailbox.
    mailboxStorageUsedInBytes *int64;
    // The snapshot date for Exchange and SharePoint used storage.
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The storage used in SharePoint document library.
    siteStorageUsedInBytes *int64;
}
// Instantiates a new getOffice365GroupsActivityStorageWithPeriod and sets the default values.
func NewGetOffice365GroupsActivityStorageWithPeriod()(*GetOffice365GroupsActivityStorageWithPeriod) {
    m := &GetOffice365GroupsActivityStorageWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mailboxStorageUsedInBytes
    }
}
// Gets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.siteStorageUsedInBytes
    }
}
// The deserialization information for the current model
func (m *GetOffice365GroupsActivityStorageWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetOffice365GroupsActivityStorageWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the mailboxStorageUsedInBytes property value. The storage used in group mailbox.
// Parameters:
//  - value : Value to set for the mailboxStorageUsedInBytes property.
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetMailboxStorageUsedInBytes(value *int64)() {
    m.mailboxStorageUsedInBytes = value
}
// Sets the reportDate property value. The snapshot date for Exchange and SharePoint used storage.
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the siteStorageUsedInBytes property value. The storage used in SharePoint document library.
// Parameters:
//  - value : Value to set for the siteStorageUsedInBytes property.
func (m *GetOffice365GroupsActivityStorageWithPeriod) SetSiteStorageUsedInBytes(value *int64)() {
    m.siteStorageUsedInBytes = value
}
