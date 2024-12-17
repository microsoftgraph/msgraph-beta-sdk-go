package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type CloudAppDiscoveryReport struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewCloudAppDiscoveryReport instantiates a new CloudAppDiscoveryReport and sets the default values.
func NewCloudAppDiscoveryReport()(*CloudAppDiscoveryReport) {
    m := &CloudAppDiscoveryReport{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateCloudAppDiscoveryReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudAppDiscoveryReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudAppDiscoveryReport(), nil
}
// GetAnonymizeMachineData gets the anonymizeMachineData property value. Use 1 if the machine information is anonymized; otherwise use 0.
// returns a *bool when successful
func (m *CloudAppDiscoveryReport) GetAnonymizeMachineData()(*bool) {
    val, err := m.GetBackingStore().Get("anonymizeMachineData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAnonymizeUserData gets the anonymizeUserData property value. Use 1 if the user information is anonymized; otherwise use 0.
// returns a *bool when successful
func (m *CloudAppDiscoveryReport) GetAnonymizeUserData()(*bool) {
    val, err := m.GetBackingStore().Get("anonymizeUserData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date in the format specified. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CloudAppDiscoveryReport) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. A comment or description for the report.
// returns a *string when successful
func (m *CloudAppDiscoveryReport) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name of the continuous report.
// returns a *string when successful
func (m *CloudAppDiscoveryReport) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
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
func (m *CloudAppDiscoveryReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anonymizeMachineData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnonymizeMachineData(val)
        }
        return nil
    }
    res["anonymizeUserData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnonymizeUserData(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isSnapshotReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSnapshotReport(val)
        }
        return nil
    }
    res["lastDataReceivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDataReceivedDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["logDataProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLogDataProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogDataProvider(val.(*LogDataProvider))
        }
        return nil
    }
    res["logFileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogFileCount(val)
        }
        return nil
    }
    res["receiverProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReceiverProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceiverProtocol(val.(*ReceiverProtocol))
        }
        return nil
    }
    res["supportedEntityTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EntityType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*EntityType))
                }
            }
            m.SetSupportedEntityTypes(res)
        }
        return nil
    }
    res["supportedTrafficTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseTrafficType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TrafficType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*TrafficType))
                }
            }
            m.SetSupportedTrafficTypes(res)
        }
        return nil
    }
    return res
}
// GetIsSnapshotReport gets the isSnapshotReport property value. Use 1 for a snapshot report; otherwise use 0.
// returns a *bool when successful
func (m *CloudAppDiscoveryReport) GetIsSnapshotReport()(*bool) {
    val, err := m.GetBackingStore().Get("isSnapshotReport")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastDataReceivedDateTime gets the lastDataReceivedDateTime property value. The date when the data was last received. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CloudAppDiscoveryReport) GetLastDataReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastDataReceivedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date when the continuous report was last modified. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *CloudAppDiscoveryReport) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLogDataProvider gets the logDataProvider property value. The applicable log data provider. Possible values are: barracuda, bluecoat, checkpoint, ciscoAsa, ciscoIronportProxy, fortigate, paloAlto, squid, zscaler, mcafeeSwg, ciscoScanSafe, juniperSrx, sophosSg, websenseV75, websenseSiemCef, machineZoneMeraki, squidNative, ciscoFwsm, microsoftIsaW3C, sonicwall, sophosCyberoam, clavister, customParser, juniperSsg, zscalerQradar, juniperSrxSd, juniperSrxWelf, microsoftConditionalAppAccess, ciscoAsaFirepower, genericCef, genericLeef, genericW3C, iFilter, checkpointXml, checkpointSmartViewTracker, barracudaNextGenFw, barracudaNextGenFwWeblog, microsoftDefenderForEndpoint, zscalerCef, sophosXg, iboss, forcepoint, fortios, ciscoIronportWsaIi, paloAltoLeef, forcepointLeef, stormshield, contentkeeper, ciscoIronportWsaIii, checkpointCef, corrata, ciscoFirepowerV6, menloSecurityCef, watchguardXtm, openSystemsSecureWebGateway, wandera, unknownFutureValue.
// returns a *LogDataProvider when successful
func (m *CloudAppDiscoveryReport) GetLogDataProvider()(*LogDataProvider) {
    val, err := m.GetBackingStore().Get("logDataProvider")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LogDataProvider)
    }
    return nil
}
// GetLogFileCount gets the logFileCount property value. The count of log files history.
// returns a *int32 when successful
func (m *CloudAppDiscoveryReport) GetLogFileCount()(*int32) {
    val, err := m.GetBackingStore().Get("logFileCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetReceiverProtocol gets the receiverProtocol property value. The applicable receiver protocol. Possible values are: ftp, ftps, syslogUdp, syslogTcp, syslogTls, unknownFutureValue.
// returns a *ReceiverProtocol when successful
func (m *CloudAppDiscoveryReport) GetReceiverProtocol()(*ReceiverProtocol) {
    val, err := m.GetBackingStore().Get("receiverProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReceiverProtocol)
    }
    return nil
}
// GetSupportedEntityTypes gets the supportedEntityTypes property value. The supported entity type. Possible values are: userName, ipAddress, machineName, other, unknown, unknownFutureValue.
// returns a []EntityType when successful
func (m *CloudAppDiscoveryReport) GetSupportedEntityTypes()([]EntityType) {
    val, err := m.GetBackingStore().Get("supportedEntityTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EntityType)
    }
    return nil
}
// GetSupportedTrafficTypes gets the supportedTrafficTypes property value. The supported traffic type. Possible values are: downloadedBytes, uploadedBytes, unknown, unknownFutureValue.
// returns a []TrafficType when successful
func (m *CloudAppDiscoveryReport) GetSupportedTrafficTypes()([]TrafficType) {
    val, err := m.GetBackingStore().Get("supportedTrafficTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TrafficType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudAppDiscoveryReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("anonymizeMachineData", m.GetAnonymizeMachineData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("anonymizeUserData", m.GetAnonymizeUserData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSnapshotReport", m.GetIsSnapshotReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastDataReceivedDateTime", m.GetLastDataReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLogDataProvider() != nil {
        cast := (*m.GetLogDataProvider()).String()
        err = writer.WriteStringValue("logDataProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("logFileCount", m.GetLogFileCount())
        if err != nil {
            return err
        }
    }
    if m.GetReceiverProtocol() != nil {
        cast := (*m.GetReceiverProtocol()).String()
        err = writer.WriteStringValue("receiverProtocol", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSupportedEntityTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedEntityTypes", SerializeEntityType(m.GetSupportedEntityTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetSupportedTrafficTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedTrafficTypes", SerializeTrafficType(m.GetSupportedTrafficTypes()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnonymizeMachineData sets the anonymizeMachineData property value. Use 1 if the machine information is anonymized; otherwise use 0.
func (m *CloudAppDiscoveryReport) SetAnonymizeMachineData(value *bool)() {
    err := m.GetBackingStore().Set("anonymizeMachineData", value)
    if err != nil {
        panic(err)
    }
}
// SetAnonymizeUserData sets the anonymizeUserData property value. Use 1 if the user information is anonymized; otherwise use 0.
func (m *CloudAppDiscoveryReport) SetAnonymizeUserData(value *bool)() {
    err := m.GetBackingStore().Set("anonymizeUserData", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date in the format specified. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudAppDiscoveryReport) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. A comment or description for the report.
func (m *CloudAppDiscoveryReport) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name of the continuous report.
func (m *CloudAppDiscoveryReport) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSnapshotReport sets the isSnapshotReport property value. Use 1 for a snapshot report; otherwise use 0.
func (m *CloudAppDiscoveryReport) SetIsSnapshotReport(value *bool)() {
    err := m.GetBackingStore().Set("isSnapshotReport", value)
    if err != nil {
        panic(err)
    }
}
// SetLastDataReceivedDateTime sets the lastDataReceivedDateTime property value. The date when the data was last received. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudAppDiscoveryReport) SetLastDataReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastDataReceivedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date when the continuous report was last modified. The Timestamp represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudAppDiscoveryReport) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLogDataProvider sets the logDataProvider property value. The applicable log data provider. Possible values are: barracuda, bluecoat, checkpoint, ciscoAsa, ciscoIronportProxy, fortigate, paloAlto, squid, zscaler, mcafeeSwg, ciscoScanSafe, juniperSrx, sophosSg, websenseV75, websenseSiemCef, machineZoneMeraki, squidNative, ciscoFwsm, microsoftIsaW3C, sonicwall, sophosCyberoam, clavister, customParser, juniperSsg, zscalerQradar, juniperSrxSd, juniperSrxWelf, microsoftConditionalAppAccess, ciscoAsaFirepower, genericCef, genericLeef, genericW3C, iFilter, checkpointXml, checkpointSmartViewTracker, barracudaNextGenFw, barracudaNextGenFwWeblog, microsoftDefenderForEndpoint, zscalerCef, sophosXg, iboss, forcepoint, fortios, ciscoIronportWsaIi, paloAltoLeef, forcepointLeef, stormshield, contentkeeper, ciscoIronportWsaIii, checkpointCef, corrata, ciscoFirepowerV6, menloSecurityCef, watchguardXtm, openSystemsSecureWebGateway, wandera, unknownFutureValue.
func (m *CloudAppDiscoveryReport) SetLogDataProvider(value *LogDataProvider)() {
    err := m.GetBackingStore().Set("logDataProvider", value)
    if err != nil {
        panic(err)
    }
}
// SetLogFileCount sets the logFileCount property value. The count of log files history.
func (m *CloudAppDiscoveryReport) SetLogFileCount(value *int32)() {
    err := m.GetBackingStore().Set("logFileCount", value)
    if err != nil {
        panic(err)
    }
}
// SetReceiverProtocol sets the receiverProtocol property value. The applicable receiver protocol. Possible values are: ftp, ftps, syslogUdp, syslogTcp, syslogTls, unknownFutureValue.
func (m *CloudAppDiscoveryReport) SetReceiverProtocol(value *ReceiverProtocol)() {
    err := m.GetBackingStore().Set("receiverProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedEntityTypes sets the supportedEntityTypes property value. The supported entity type. Possible values are: userName, ipAddress, machineName, other, unknown, unknownFutureValue.
func (m *CloudAppDiscoveryReport) SetSupportedEntityTypes(value []EntityType)() {
    err := m.GetBackingStore().Set("supportedEntityTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedTrafficTypes sets the supportedTrafficTypes property value. The supported traffic type. Possible values are: downloadedBytes, uploadedBytes, unknown, unknownFutureValue.
func (m *CloudAppDiscoveryReport) SetSupportedTrafficTypes(value []TrafficType)() {
    err := m.GetBackingStore().Set("supportedTrafficTypes", value)
    if err != nil {
        panic(err)
    }
}
type CloudAppDiscoveryReportable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnonymizeMachineData()(*bool)
    GetAnonymizeUserData()(*bool)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsSnapshotReport()(*bool)
    GetLastDataReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogDataProvider()(*LogDataProvider)
    GetLogFileCount()(*int32)
    GetReceiverProtocol()(*ReceiverProtocol)
    GetSupportedEntityTypes()([]EntityType)
    GetSupportedTrafficTypes()([]TrafficType)
    SetAnonymizeMachineData(value *bool)()
    SetAnonymizeUserData(value *bool)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsSnapshotReport(value *bool)()
    SetLastDataReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogDataProvider(value *LogDataProvider)()
    SetLogFileCount(value *int32)()
    SetReceiverProtocol(value *ReceiverProtocol)()
    SetSupportedEntityTypes(value []EntityType)()
    SetSupportedTrafficTypes(value []TrafficType)()
}
