package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementExportJob struct {
    Entity
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    filter *string;
    format *DeviceManagementReportFileFormat;
    localizationType *DeviceManagementExportJobLocalizationType;
    reportName *string;
    requestDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    select_escaped []string;
    snapshotId *string;
    status *DeviceManagementReportStatus;
    url *string;
}
func NewDeviceManagementExportJob()(*DeviceManagementExportJob) {
    m := &DeviceManagementExportJob{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *DeviceManagementExportJob) GetFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filter
    }
}
func (m *DeviceManagementExportJob) GetFormat()(*DeviceManagementReportFileFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *DeviceManagementExportJob) GetLocalizationType()(*DeviceManagementExportJobLocalizationType) {
    if m == nil {
        return nil
    } else {
        return m.localizationType
    }
}
func (m *DeviceManagementExportJob) GetReportName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportName
    }
}
func (m *DeviceManagementExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestDateTime
    }
}
func (m *DeviceManagementExportJob) GetSelect_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.select_escaped
    }
}
func (m *DeviceManagementExportJob) GetSnapshotId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.snapshotId
    }
}
func (m *DeviceManagementExportJob) GetStatus()(*DeviceManagementReportStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *DeviceManagementExportJob) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *DeviceManagementExportJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementReportFileFormat)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementReportFileFormat)
        m.SetFormat(&cast)
        return nil
    }
    res["localizationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExportJobLocalizationType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExportJobLocalizationType)
        m.SetLocalizationType(&cast)
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
    res["requestDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestDateTime(val)
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
    res["snapshotId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSnapshotId(val)
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
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *DeviceManagementExportJob) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementExportJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetFormat() != nil {
        cast := m.GetFormat().String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalizationType() != nil {
        cast := m.GetLocalizationType().String()
        err = writer.WriteStringValue("localizationType", &cast)
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
        err = writer.WriteTimeValue("requestDateTime", m.GetRequestDateTime())
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
        err = writer.WriteStringValue("snapshotId", m.GetSnapshotId())
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
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceManagementExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *DeviceManagementExportJob) SetFilter(value *string)() {
    m.filter = value
}
func (m *DeviceManagementExportJob) SetFormat(value *DeviceManagementReportFileFormat)() {
    m.format = value
}
func (m *DeviceManagementExportJob) SetLocalizationType(value *DeviceManagementExportJobLocalizationType)() {
    m.localizationType = value
}
func (m *DeviceManagementExportJob) SetReportName(value *string)() {
    m.reportName = value
}
func (m *DeviceManagementExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestDateTime = value
}
func (m *DeviceManagementExportJob) SetSelect_escaped(value []string)() {
    m.select_escaped = value
}
func (m *DeviceManagementExportJob) SetSnapshotId(value *string)() {
    m.snapshotId = value
}
func (m *DeviceManagementExportJob) SetStatus(value *DeviceManagementReportStatus)() {
    m.status = value
}
func (m *DeviceManagementExportJob) SetUrl(value *string)() {
    m.url = value
}
