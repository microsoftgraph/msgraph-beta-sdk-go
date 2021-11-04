package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementReportSchedule struct {
    Entity
    // Emails to which the scheduled reports are delivered
    emails []string;
    // Time that the delivery of the scheduled reports ends
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Filters applied on the report
    filter *string;
    // Format of the scheduled report. Possible values are: csv, pdf.
    format *DeviceManagementReportFileFormat;
    // Ordering of columns in the report
    orderBy []string;
    // Frequency of scheduled report delivery. Possible values are: none, daily, weekly, monthly.
    recurrence *DeviceManagementScheduledReportRecurrence;
    // Name of the report
    reportName *string;
    // Name of the schedule
    reportScheduleName *string;
    // Columns selected from the report
    select_escaped []string;
    // Time that the delivery of the scheduled reports starts
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Subject of the scheduled reports that are delivered
    subject *string;
    // The Id of the User who created the report
    userId *string;
}
// Instantiates a new deviceManagementReportSchedule and sets the default values.
func NewDeviceManagementReportSchedule()(*DeviceManagementReportSchedule) {
    m := &DeviceManagementReportSchedule{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the emails property value. Emails to which the scheduled reports are delivered
func (m *DeviceManagementReportSchedule) GetEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
// Gets the endDateTime property value. Time that the delivery of the scheduled reports ends
func (m *DeviceManagementReportSchedule) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the filter property value. Filters applied on the report
func (m *DeviceManagementReportSchedule) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// Gets the format property value. Format of the scheduled report. Possible values are: csv, pdf.
func (m *DeviceManagementReportSchedule) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementReportSchedule) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// Gets the recurrence property value. Frequency of scheduled report delivery. Possible values are: none, daily, weekly, monthly.
func (m *DeviceManagementReportSchedule) GetRecurrence()(*DeviceManagementScheduledReportRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the reportName property value. Name of the report
func (m *DeviceManagementReportSchedule) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
// Gets the reportScheduleName property value. Name of the schedule
func (m *DeviceManagementReportSchedule) GetReportScheduleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportScheduleName
    }
}
// Gets the select_escaped property value. Columns selected from the report
func (m *DeviceManagementReportSchedule) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
// Gets the startDateTime property value. Time that the delivery of the scheduled reports starts
func (m *DeviceManagementReportSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the subject property value. Subject of the scheduled reports that are delivered
func (m *DeviceManagementReportSchedule) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the userId property value. The Id of the User who created the report
func (m *DeviceManagementReportSchedule) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementReportSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetEmails(res)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
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
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementReportFileFormat)
        m.SetFormat(&cast)
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
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementScheduledReportRecurrence)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementScheduledReportRecurrence)
        m.SetRecurrence(&cast)
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
    res["reportScheduleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportScheduleName(val)
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
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *DeviceManagementReportSchedule) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementReportSchedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("emails", m.GetEmails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
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
    if m.GetFormat() != nil {
        cast := m.GetFormat().String()
        err = writer.WriteStringValue("format", &cast)
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
    if m.GetRecurrence() != nil {
        cast := m.GetRecurrence().String()
        err = writer.WriteStringValue("recurrence", &cast)
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
        err = writer.WriteStringValue("reportScheduleName", m.GetReportScheduleName())
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
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the emails property value. Emails to which the scheduled reports are delivered
// Parameters:
//  - value : Value to set for the emails property.
func (m *DeviceManagementReportSchedule) SetEmails(value []string)() {
    m.emails = value
}
// Sets the endDateTime property value. Time that the delivery of the scheduled reports ends
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *DeviceManagementReportSchedule) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the filter property value. Filters applied on the report
// Parameters:
//  - value : Value to set for the filter property.
func (m *DeviceManagementReportSchedule) SetFilter(value *string)() {
    m.filter = value
}
// Sets the format property value. Format of the scheduled report. Possible values are: csv, pdf.
// Parameters:
//  - value : Value to set for the format property.
func (m *DeviceManagementReportSchedule) SetFormat(value *DeviceManagementReportFileFormat)() {
    m.format = value
}
// Sets the orderBy property value. Ordering of columns in the report
// Parameters:
//  - value : Value to set for the orderBy property.
func (m *DeviceManagementReportSchedule) SetOrderBy(value []string)() {
    m.orderBy = value
}
// Sets the recurrence property value. Frequency of scheduled report delivery. Possible values are: none, daily, weekly, monthly.
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *DeviceManagementReportSchedule) SetRecurrence(value *DeviceManagementScheduledReportRecurrence)() {
    m.recurrence = value
}
// Sets the reportName property value. Name of the report
// Parameters:
//  - value : Value to set for the reportName property.
func (m *DeviceManagementReportSchedule) SetReportName(value *string)() {
    m.reportName = value
}
// Sets the reportScheduleName property value. Name of the schedule
// Parameters:
//  - value : Value to set for the reportScheduleName property.
func (m *DeviceManagementReportSchedule) SetReportScheduleName(value *string)() {
    m.reportScheduleName = value
}
// Sets the select_escaped property value. Columns selected from the report
// Parameters:
//  - value : Value to set for the select_escaped property.
func (m *DeviceManagementReportSchedule) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
// Sets the startDateTime property value. Time that the delivery of the scheduled reports starts
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *DeviceManagementReportSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the subject property value. Subject of the scheduled reports that are delivered
// Parameters:
//  - value : Value to set for the subject property.
func (m *DeviceManagementReportSchedule) SetSubject(value *string)() {
    m.subject = value
}
// Sets the userId property value. The Id of the User who created the report
// Parameters:
//  - value : Value to set for the userId property.
func (m *DeviceManagementReportSchedule) SetUserId(value *string)() {
    m.userId = value
}
