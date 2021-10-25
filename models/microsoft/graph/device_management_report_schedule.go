package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementReportSchedule struct {
    Entity
    emails []string;
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    filter *string;
    format *DeviceManagementReportFileFormat;
    orderBy []string;
    recurrence *DeviceManagementScheduledReportRecurrence;
    reportName *string;
    reportScheduleName *string;
    select_escaped []string;
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    subject *string;
    userId *string;
}
func NewDeviceManagementReportSchedule()(*DeviceManagementReportSchedule) {
    m := &DeviceManagementReportSchedule{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementReportSchedule) GetEmails()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
func (m *DeviceManagementReportSchedule) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
func (m *DeviceManagementReportSchedule) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
func (m *DeviceManagementReportSchedule) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *DeviceManagementReportSchedule) GetOrderBy()([]string) {
    if m == nil {
        return nil
    } else {
        return m.orderBy
    }
}
func (m *DeviceManagementReportSchedule) GetRecurrence()(*DeviceManagementScheduledReportRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
func (m *DeviceManagementReportSchedule) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
func (m *DeviceManagementReportSchedule) GetReportScheduleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportScheduleName
    }
}
func (m *DeviceManagementReportSchedule) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
func (m *DeviceManagementReportSchedule) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
func (m *DeviceManagementReportSchedule) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *DeviceManagementReportSchedule) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *DeviceManagementReportSchedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
func (m *DeviceManagementReportSchedule) SetEmails(value []string)() {
    m.emails = value
}
func (m *DeviceManagementReportSchedule) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
func (m *DeviceManagementReportSchedule) SetFilter(value *string)() {
    m.filter = value
}
func (m *DeviceManagementReportSchedule) SetFormat(value *DeviceManagementReportFileFormat)() {
    m.format = value
}
func (m *DeviceManagementReportSchedule) SetOrderBy(value []string)() {
    m.orderBy = value
}
func (m *DeviceManagementReportSchedule) SetRecurrence(value *DeviceManagementScheduledReportRecurrence)() {
    m.recurrence = value
}
func (m *DeviceManagementReportSchedule) SetReportName(value *string)() {
    m.reportName = value
}
func (m *DeviceManagementReportSchedule) SetReportScheduleName(value *string)() {
    m.reportScheduleName = value
}
func (m *DeviceManagementReportSchedule) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
func (m *DeviceManagementReportSchedule) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
func (m *DeviceManagementReportSchedule) SetSubject(value *string)() {
    m.subject = value
}
func (m *DeviceManagementReportSchedule) SetUserId(value *string)() {
    m.userId = value
}
