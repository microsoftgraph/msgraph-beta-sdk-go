package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementReportSchedule 
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
    select []string;
    // Time that the delivery of the scheduled reports starts
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Subject of the scheduled reports that are delivered
    subject *string;
    // The Id of the User who created the report
    userId *string;
}
// NewDeviceManagementReportSchedule instantiates a new deviceManagementReportSchedule and sets the default values.
func NewDeviceManagementReportSchedule()(*DeviceManagementReportSchedule) {
    m := &DeviceManagementReportSchedule{
        Entity: *NewEntity(),
    }
    return m
}
// GetEmails gets the emails property value. Emails to which the scheduled reports are delivered
func (m *DeviceManagementReportSchedule) GetEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
// GetEndDateTime gets the endDateTime property value. Time that the delivery of the scheduled reports ends
func (m *DeviceManagementReportSchedule) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFilter gets the filter property value. Filters applied on the report
func (m *DeviceManagementReportSchedule) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
// GetFormat gets the format property value. Format of the scheduled report. Possible values are: csv, pdf.
func (m *DeviceManagementReportSchedule) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetOrderBy gets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementReportSchedule) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
// GetRecurrence gets the recurrence property value. Frequency of scheduled report delivery. Possible values are: none, daily, weekly, monthly.
func (m *DeviceManagementReportSchedule) GetRecurrence()(*DeviceManagementScheduledReportRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetReportName gets the reportName property value. Name of the report
func (m *DeviceManagementReportSchedule) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
// GetReportScheduleName gets the reportScheduleName property value. Name of the schedule
func (m *DeviceManagementReportSchedule) GetReportScheduleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportScheduleName
    }
}
// GetSelect gets the select property value. Columns selected from the report
func (m *DeviceManagementReportSchedule) GetSelect()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select
    }
}
// GetStartDateTime gets the startDateTime property value. Time that the delivery of the scheduled reports starts
func (m *DeviceManagementReportSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetSubject gets the subject property value. Subject of the scheduled reports that are delivered
func (m *DeviceManagementReportSchedule) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetUserId gets the userId property value. The Id of the User who created the report
func (m *DeviceManagementReportSchedule) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementReportSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["filter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementReportFileFormat)
            m.SetFormat(&cast)
        }
        return nil
    }
    res["orderBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrderBy(res)
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementScheduledReportRecurrence)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementScheduledReportRecurrence)
            m.SetRecurrence(&cast)
        }
        return nil
    }
    res["reportName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val)
        }
        return nil
    }
    res["reportScheduleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportScheduleName(val)
        }
        return nil
    }
    res["select"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSelect(res)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementReportSchedule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err = writer.WriteCollectionOfStringValues("select", m.GetSelect())
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
// SetEmails sets the emails property value. Emails to which the scheduled reports are delivered
func (m *DeviceManagementReportSchedule) SetEmails(value []string)() {
    if m != nil {
        m.emails = value
    }
}
// SetEndDateTime sets the endDateTime property value. Time that the delivery of the scheduled reports ends
func (m *DeviceManagementReportSchedule) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetFilter sets the filter property value. Filters applied on the report
func (m *DeviceManagementReportSchedule) SetFilter(value *string)() {
    if m != nil {
        m.filter = value
    }
}
// SetFormat sets the format property value. Format of the scheduled report. Possible values are: csv, pdf.
func (m *DeviceManagementReportSchedule) SetFormat(value *DeviceManagementReportFileFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetOrderBy sets the orderBy property value. Ordering of columns in the report
func (m *DeviceManagementReportSchedule) SetOrderBy(value []string)() {
    if m != nil {
        m.orderBy = value
    }
}
// SetRecurrence sets the recurrence property value. Frequency of scheduled report delivery. Possible values are: none, daily, weekly, monthly.
func (m *DeviceManagementReportSchedule) SetRecurrence(value *DeviceManagementScheduledReportRecurrence)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetReportName sets the reportName property value. Name of the report
func (m *DeviceManagementReportSchedule) SetReportName(value *string)() {
    if m != nil {
        m.reportName = value
    }
}
// SetReportScheduleName sets the reportScheduleName property value. Name of the schedule
func (m *DeviceManagementReportSchedule) SetReportScheduleName(value *string)() {
    if m != nil {
        m.reportScheduleName = value
    }
}
// SetSelect sets the select property value. Columns selected from the report
func (m *DeviceManagementReportSchedule) SetSelect(value []string)() {
    if m != nil {
        m.select = value
    }
}
// SetStartDateTime sets the startDateTime property value. Time that the delivery of the scheduled reports starts
func (m *DeviceManagementReportSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetSubject sets the subject property value. Subject of the scheduled reports that are delivered
func (m *DeviceManagementReportSchedule) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetUserId sets the userId property value. The Id of the User who created the report
func (m *DeviceManagementReportSchedule) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
