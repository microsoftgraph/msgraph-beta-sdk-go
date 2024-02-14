package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DiscoveredApplicationSegmentReport struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDiscoveredApplicationSegmentReport instantiates a new DiscoveredApplicationSegmentReport and sets the default values.
func NewDiscoveredApplicationSegmentReport()(*DiscoveredApplicationSegmentReport) {
    m := &DiscoveredApplicationSegmentReport{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDiscoveredApplicationSegmentReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDiscoveredApplicationSegmentReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDiscoveredApplicationSegmentReport(), nil
}
// GetAccessType gets the accessType property value. The accessType property
// returns a *AccessType when successful
func (m *DiscoveredApplicationSegmentReport) GetAccessType()(*AccessType) {
    val, err := m.GetBackingStore().Get("accessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AccessType)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DiscoveredApplicationSegmentReport) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DiscoveredApplicationSegmentReport) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceCount gets the deviceCount property value. The deviceCount property
// returns a *int32 when successful
func (m *DiscoveredApplicationSegmentReport) GetDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDiscoveredApplicationSegmentId gets the discoveredApplicationSegmentId property value. The discoveredApplicationSegmentId property
// returns a *string when successful
func (m *DiscoveredApplicationSegmentReport) GetDiscoveredApplicationSegmentId()(*string) {
    val, err := m.GetBackingStore().Get("discoveredApplicationSegmentId")
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
func (m *DiscoveredApplicationSegmentReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*AccessType))
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["discoveredApplicationSegmentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveredApplicationSegmentId(val)
        }
        return nil
    }
    res["firstAccessDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstAccessDateTime(val)
        }
        return nil
    }
    res["fqdn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFqdn(val)
        }
        return nil
    }
    res["ip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIp(val)
        }
        return nil
    }
    res["lastAccessDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAccessDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["totalBytesReceived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBytesReceived(val)
        }
        return nil
    }
    res["totalBytesSent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBytesSent(val)
        }
        return nil
    }
    res["transactionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionCount(val)
        }
        return nil
    }
    res["transportProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransportProtocol(val.(*NetworkingProtocol))
        }
        return nil
    }
    res["userCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCount(val)
        }
        return nil
    }
    return res
}
// GetFirstAccessDateTime gets the firstAccessDateTime property value. The firstAccessDateTime property
// returns a *Time when successful
func (m *DiscoveredApplicationSegmentReport) GetFirstAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstAccessDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFqdn gets the fqdn property value. The fqdn property
// returns a *string when successful
func (m *DiscoveredApplicationSegmentReport) GetFqdn()(*string) {
    val, err := m.GetBackingStore().Get("fqdn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIp gets the ip property value. The ip property
// returns a *string when successful
func (m *DiscoveredApplicationSegmentReport) GetIp()(*string) {
    val, err := m.GetBackingStore().Get("ip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastAccessDateTime gets the lastAccessDateTime property value. The lastAccessDateTime property
// returns a *Time when successful
func (m *DiscoveredApplicationSegmentReport) GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastAccessDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DiscoveredApplicationSegmentReport) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPort gets the port property value. The port property
// returns a *int32 when successful
func (m *DiscoveredApplicationSegmentReport) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalBytesReceived gets the totalBytesReceived property value. The totalBytesReceived property
// returns a *int64 when successful
func (m *DiscoveredApplicationSegmentReport) GetTotalBytesReceived()(*int64) {
    val, err := m.GetBackingStore().Get("totalBytesReceived")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalBytesSent gets the totalBytesSent property value. The totalBytesSent property
// returns a *int64 when successful
func (m *DiscoveredApplicationSegmentReport) GetTotalBytesSent()(*int64) {
    val, err := m.GetBackingStore().Get("totalBytesSent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTransactionCount gets the transactionCount property value. The transactionCount property
// returns a *int32 when successful
func (m *DiscoveredApplicationSegmentReport) GetTransactionCount()(*int32) {
    val, err := m.GetBackingStore().Get("transactionCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTransportProtocol gets the transportProtocol property value. The transportProtocol property
// returns a *NetworkingProtocol when successful
func (m *DiscoveredApplicationSegmentReport) GetTransportProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("transportProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// GetUserCount gets the userCount property value. The userCount property
// returns a *int32 when successful
func (m *DiscoveredApplicationSegmentReport) GetUserCount()(*int32) {
    val, err := m.GetBackingStore().Get("userCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DiscoveredApplicationSegmentReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("discoveredApplicationSegmentId", m.GetDiscoveredApplicationSegmentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("firstAccessDateTime", m.GetFirstAccessDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fqdn", m.GetFqdn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ip", m.GetIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastAccessDateTime", m.GetLastAccessDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalBytesReceived", m.GetTotalBytesReceived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalBytesSent", m.GetTotalBytesSent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("transactionCount", m.GetTransactionCount())
        if err != nil {
            return err
        }
    }
    if m.GetTransportProtocol() != nil {
        cast := (*m.GetTransportProtocol()).String()
        err := writer.WriteStringValue("transportProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("userCount", m.GetUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessType sets the accessType property value. The accessType property
func (m *DiscoveredApplicationSegmentReport) SetAccessType(value *AccessType)() {
    err := m.GetBackingStore().Set("accessType", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DiscoveredApplicationSegmentReport) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DiscoveredApplicationSegmentReport) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceCount sets the deviceCount property value. The deviceCount property
func (m *DiscoveredApplicationSegmentReport) SetDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscoveredApplicationSegmentId sets the discoveredApplicationSegmentId property value. The discoveredApplicationSegmentId property
func (m *DiscoveredApplicationSegmentReport) SetDiscoveredApplicationSegmentId(value *string)() {
    err := m.GetBackingStore().Set("discoveredApplicationSegmentId", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstAccessDateTime sets the firstAccessDateTime property value. The firstAccessDateTime property
func (m *DiscoveredApplicationSegmentReport) SetFirstAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstAccessDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFqdn sets the fqdn property value. The fqdn property
func (m *DiscoveredApplicationSegmentReport) SetFqdn(value *string)() {
    err := m.GetBackingStore().Set("fqdn", value)
    if err != nil {
        panic(err)
    }
}
// SetIp sets the ip property value. The ip property
func (m *DiscoveredApplicationSegmentReport) SetIp(value *string)() {
    err := m.GetBackingStore().Set("ip", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAccessDateTime sets the lastAccessDateTime property value. The lastAccessDateTime property
func (m *DiscoveredApplicationSegmentReport) SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAccessDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DiscoveredApplicationSegmentReport) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. The port property
func (m *DiscoveredApplicationSegmentReport) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBytesReceived sets the totalBytesReceived property value. The totalBytesReceived property
func (m *DiscoveredApplicationSegmentReport) SetTotalBytesReceived(value *int64)() {
    err := m.GetBackingStore().Set("totalBytesReceived", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBytesSent sets the totalBytesSent property value. The totalBytesSent property
func (m *DiscoveredApplicationSegmentReport) SetTotalBytesSent(value *int64)() {
    err := m.GetBackingStore().Set("totalBytesSent", value)
    if err != nil {
        panic(err)
    }
}
// SetTransactionCount sets the transactionCount property value. The transactionCount property
func (m *DiscoveredApplicationSegmentReport) SetTransactionCount(value *int32)() {
    err := m.GetBackingStore().Set("transactionCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTransportProtocol sets the transportProtocol property value. The transportProtocol property
func (m *DiscoveredApplicationSegmentReport) SetTransportProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("transportProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. The userCount property
func (m *DiscoveredApplicationSegmentReport) SetUserCount(value *int32)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
type DiscoveredApplicationSegmentReportable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessType()(*AccessType)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceCount()(*int32)
    GetDiscoveredApplicationSegmentId()(*string)
    GetFirstAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFqdn()(*string)
    GetIp()(*string)
    GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetPort()(*int32)
    GetTotalBytesReceived()(*int64)
    GetTotalBytesSent()(*int64)
    GetTransactionCount()(*int32)
    GetTransportProtocol()(*NetworkingProtocol)
    GetUserCount()(*int32)
    SetAccessType(value *AccessType)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceCount(value *int32)()
    SetDiscoveredApplicationSegmentId(value *string)()
    SetFirstAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFqdn(value *string)()
    SetIp(value *string)()
    SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetPort(value *int32)()
    SetTotalBytesReceived(value *int64)()
    SetTotalBytesSent(value *int64)()
    SetTransactionCount(value *int32)()
    SetTransportProtocol(value *NetworkingProtocol)()
    SetUserCount(value *int32)()
}
