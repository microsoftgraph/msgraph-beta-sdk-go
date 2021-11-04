package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementCachedReportConfiguration struct {
    Entity
    // Time that the cached report expires
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Filters applied on report creation.
    filter *string;
    // Time that the cached report was last refreshed
    lastRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Caller-managed metadata associated with the report
    metadata *string;
    // Ordering of columns in the report
    orderBy []string;
    // Name of the report
    reportName *string;
    // Columns selected from the report
    select_escaped []string;
    // Status of the cached report. Possible values are: unknown, notStarted, inProgress, completed, failed.
    status *DeviceManagementReportStatus;
}
// Instantiates a new deviceManagementCachedReportConfiguration and sets the default values.
func NewDeviceManagementCachedReportConfiguration()(*DeviceManagementCachedReportConfiguration) {
    m := &DeviceManagementCachedReportConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the expirationDateTime property value. Time that the cached report expires
func (m *DeviceManagementCachedReportConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the filter property value. Filters applied on report creation.
func (m *DeviceManagementCachedReportConfiguration) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// Gets the lastRefreshDateTime property value. Time that the cached report was last refreshed
func (m *DeviceManagementCachedReportConfiguration) GetLastRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshDateTime
    }
}
// Gets the metadata property value. Caller-managed metadata associated with the report
func (m *DeviceManagementCachedReportConfiguration) GetMetadata()(*string) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementCachedReportConfiguration) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// Gets the reportName property value. Name of the report
func (m *DeviceManagementCachedReportConfiguration) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
// Gets the select_escaped property value. Columns selected from the report
func (m *DeviceManagementCachedReportConfiguration) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// Gets the status property value. Status of the cached report. Possible values are: unknown, notStarted, inProgress, completed, failed.
func (m *DeviceManagementCachedReportConfiguration) GetStatus()(*DeviceManagementReportStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *DeviceManagementCachedReportConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilter(val)
        return nil
    }
    res["lastRefreshDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshDateTime(val)
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMetadata(val)
        return nil
    }
    res["orderBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetOrderBy(res)
        return nil
    }
    res["reportName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportName(val)
        return nil
    }
    res["select_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSelect_escaped(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportStatus)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementReportStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementCachedReportConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementCachedReportConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("filter", m.GetFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshDateTime", m.GetLastRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadata", m.GetMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("orderBy", m.GetOrderBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportName", m.GetReportName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("select_escaped", m.GetSelect_escaped())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the expirationDateTime property value. Time that the cached report expires
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *DeviceManagementCachedReportConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the filter property value. Filters applied on report creation.
// Parameters:
//  - value : Value to set for the filter property.
func (m *DeviceManagementCachedReportConfiguration) SetFilter(value *string)() {
    m.filter = value
}
// Sets the lastRefreshDateTime property value. Time that the cached report was last refreshed
// Parameters:
//  - value : Value to set for the lastRefreshDateTime property.
func (m *DeviceManagementCachedReportConfiguration) SetLastRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshDateTime = value
}
// Sets the metadata property value. Caller-managed metadata associated with the report
// Parameters:
//  - value : Value to set for the metadata property.
func (m *DeviceManagementCachedReportConfiguration) SetMetadata(value *string)() {
    m.metadata = value
}
// Sets the orderBy property value. Ordering of columns in the report
// Parameters:
//  - value : Value to set for the orderBy property.
func (m *DeviceManagementCachedReportConfiguration) SetOrderBy(value []string)() {
    m.orderBy = value
}
// Sets the reportName property value. Name of the report
// Parameters:
//  - value : Value to set for the reportName property.
func (m *DeviceManagementCachedReportConfiguration) SetReportName(value *string)() {
    m.reportName = value
}
// Sets the select_escaped property value. Columns selected from the report
// Parameters:
//  - value : Value to set for the select_escaped property.
func (m *DeviceManagementCachedReportConfiguration) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
// Sets the status property value. Status of the cached report. Possible values are: unknown, notStarted, inProgress, completed, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *DeviceManagementCachedReportConfiguration) SetStatus(value *DeviceManagementReportStatus)() {
    m.status = value
}
