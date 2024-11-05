package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CloudPcExportJob struct {
    Entity
}
// NewCloudPcExportJob instantiates a new CloudPcExportJob and sets the default values.
func NewCloudPcExportJob()(*CloudPcExportJob) {
    m := &CloudPcExportJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcExportJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudPcExportJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcExportJob(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time when the export job expires.
// returns a *Time when successful
func (m *CloudPcExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExportJobStatus gets the exportJobStatus property value. The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue. Read-only.
// returns a *CloudPcExportJobStatus when successful
func (m *CloudPcExportJob) GetExportJobStatus()(*CloudPcExportJobStatus) {
    val, err := m.GetBackingStore().Get("exportJobStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcExportJobStatus)
    }
    return nil
}
// GetExportUrl gets the exportUrl property value. The storage account URL of the exported report. It can be used to download the file.
// returns a *string when successful
func (m *CloudPcExportJob) GetExportUrl()(*string) {
    val, err := m.GetBackingStore().Get("exportUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudPcExportJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["exportJobStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcExportJobStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportJobStatus(val.(*CloudPcExportJobStatus))
        }
        return nil
    }
    res["exportUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportUrl(val)
        }
        return nil
    }
    res["filter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilter(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["reportName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcReportName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportName(val.(*CloudPcReportName))
        }
        return nil
    }
    res["requestDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestDateTime(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSelectEscaped(res)
        }
        return nil
    }
    return res
}
// GetFilter gets the filter property value. The filter applied on the report.
// returns a *string when successful
func (m *CloudPcExportJob) GetFilter()(*string) {
    val, err := m.GetBackingStore().Get("filter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFormat gets the format property value. The format of the exported report.
// returns a *string when successful
func (m *CloudPcExportJob) GetFormat()(*string) {
    val, err := m.GetBackingStore().Get("format")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReportName gets the reportName property value. The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports, totalAggregatedRemoteConnectionReports, sharedUseLicenseUsageReport, sharedUseLicenseUsageRealTimeReport, unknownFutureValue,  noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport, frontlineLicenseUsageRealTimeReport,  remoteConnectionQualityReports, inaccessibleCloudPcReports, actionStatusReport, rawRemoteConnectionReports, cloudPcUsageCategoryReports, crossRegionDisasterRecoveryReport, regionalConnectionQualityTrendReport, regionalConnectionQualityInsightsReport, remoteConnectionQualityReport, bulkActionStatusReport. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport, frontlineLicenseUsageRealTimeReport, remoteConnectionQualityReports, inaccessibleCloudPcReports, rawRemoteConnectionReports, cloudPcUsageCategoryReports, crossRegionDisasterRecoveryReport.
// returns a *CloudPcReportName when successful
func (m *CloudPcExportJob) GetReportName()(*CloudPcReportName) {
    val, err := m.GetBackingStore().Get("reportName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcReportName)
    }
    return nil
}
// GetRequestDateTime gets the requestDateTime property value. The date and time when the export job was requested.
// returns a *Time when successful
func (m *CloudPcExportJob) GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSelectEscaped gets the select property value. The selected columns of the report.
// returns a []string when successful
func (m *CloudPcExportJob) GetSelectEscaped()([]string) {
    val, err := m.GetBackingStore().Get("selectEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcExportJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetExportJobStatus() != nil {
        cast := (*m.GetExportJobStatus()).String()
        err = writer.WriteStringValue("exportJobStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exportUrl", m.GetExportUrl())
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
        err = writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    if m.GetReportName() != nil {
        cast := (*m.GetReportName()).String()
        err = writer.WriteStringValue("reportName", &cast)
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
    if m.GetSelectEscaped() != nil {
        err = writer.WriteCollectionOfStringValues("select", m.GetSelectEscaped())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time when the export job expires.
func (m *CloudPcExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExportJobStatus sets the exportJobStatus property value. The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue. Read-only.
func (m *CloudPcExportJob) SetExportJobStatus(value *CloudPcExportJobStatus)() {
    err := m.GetBackingStore().Set("exportJobStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetExportUrl sets the exportUrl property value. The storage account URL of the exported report. It can be used to download the file.
func (m *CloudPcExportJob) SetExportUrl(value *string)() {
    err := m.GetBackingStore().Set("exportUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetFilter sets the filter property value. The filter applied on the report.
func (m *CloudPcExportJob) SetFilter(value *string)() {
    err := m.GetBackingStore().Set("filter", value)
    if err != nil {
        panic(err)
    }
}
// SetFormat sets the format property value. The format of the exported report.
func (m *CloudPcExportJob) SetFormat(value *string)() {
    err := m.GetBackingStore().Set("format", value)
    if err != nil {
        panic(err)
    }
}
// SetReportName sets the reportName property value. The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports, totalAggregatedRemoteConnectionReports, sharedUseLicenseUsageReport, sharedUseLicenseUsageRealTimeReport, unknownFutureValue,  noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport, frontlineLicenseUsageRealTimeReport,  remoteConnectionQualityReports, inaccessibleCloudPcReports, actionStatusReport, rawRemoteConnectionReports, cloudPcUsageCategoryReports, crossRegionDisasterRecoveryReport, regionalConnectionQualityTrendReport, regionalConnectionQualityInsightsReport, remoteConnectionQualityReport, bulkActionStatusReport. You must use the Prefer: include-unknown-enum-members request header to get the following values in this evolvable enum: noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport, frontlineLicenseUsageRealTimeReport, remoteConnectionQualityReports, inaccessibleCloudPcReports, rawRemoteConnectionReports, cloudPcUsageCategoryReports, crossRegionDisasterRecoveryReport.
func (m *CloudPcExportJob) SetReportName(value *CloudPcReportName)() {
    err := m.GetBackingStore().Set("reportName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestDateTime sets the requestDateTime property value. The date and time when the export job was requested.
func (m *CloudPcExportJob) SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSelectEscaped sets the select property value. The selected columns of the report.
func (m *CloudPcExportJob) SetSelectEscaped(value []string)() {
    err := m.GetBackingStore().Set("selectEscaped", value)
    if err != nil {
        panic(err)
    }
}
type CloudPcExportJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExportJobStatus()(*CloudPcExportJobStatus)
    GetExportUrl()(*string)
    GetFilter()(*string)
    GetFormat()(*string)
    GetReportName()(*CloudPcReportName)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSelectEscaped()([]string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExportJobStatus(value *CloudPcExportJobStatus)()
    SetExportUrl(value *string)()
    SetFilter(value *string)()
    SetFormat(value *string)()
    SetReportName(value *CloudPcReportName)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSelectEscaped(value []string)()
}
