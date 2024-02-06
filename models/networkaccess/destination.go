package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// Destination 
type Destination struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDestination instantiates a new destination and sets the default values.
func NewDestination()(*Destination) {
    m := &Destination{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDestinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDestinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDestination(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Destination) GetAdditionalData()(map[string]any) {
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
func (m *Destination) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceCount gets the deviceCount property value. The number of unique devices that were seen.
func (m *Destination) GetDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Destination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["networkingProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkingProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkingProtocol(val.(*NetworkingProtocol))
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
    res["threatCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatCount(val)
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
    res["trafficType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrafficType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficType(val.(*TrafficType))
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
func (m *Destination) GetFirstAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("firstAccessDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFqdn gets the fqdn property value. The fully qualified domain name (FQDN) of the destination.
func (m *Destination) GetFqdn()(*string) {
    val, err := m.GetBackingStore().Get("fqdn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIp gets the ip property value. The internet protocol (IP) used to access the destination.
func (m *Destination) GetIp()(*string) {
    val, err := m.GetBackingStore().Get("ip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastAccessDateTime gets the lastAccessDateTime property value. The most recent access DateTime.
func (m *Destination) GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastAccessDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetNetworkingProtocol gets the networkingProtocol property value. The networkingProtocol property
func (m *Destination) GetNetworkingProtocol()(*NetworkingProtocol) {
    val, err := m.GetBackingStore().Get("networkingProtocol")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*NetworkingProtocol)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Destination) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPort gets the port property value. The numeric identifier that is associated with a specific endpoint in a network.
func (m *Destination) GetPort()(*int32) {
    val, err := m.GetBackingStore().Get("port")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetThreatCount gets the threatCount property value. The threatCount property
func (m *Destination) GetThreatCount()(*int32) {
    val, err := m.GetBackingStore().Get("threatCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalBytesReceived gets the totalBytesReceived property value. The totalBytesReceived property
func (m *Destination) GetTotalBytesReceived()(*int64) {
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
func (m *Destination) GetTotalBytesSent()(*int64) {
    val, err := m.GetBackingStore().Get("totalBytesSent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTrafficType gets the trafficType property value. The trafficType property
func (m *Destination) GetTrafficType()(*TrafficType) {
    val, err := m.GetBackingStore().Get("trafficType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrafficType)
    }
    return nil
}
// GetTransactionCount gets the transactionCount property value. The number of transactions.
func (m *Destination) GetTransactionCount()(*int32) {
    val, err := m.GetBackingStore().Get("transactionCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUserCount gets the userCount property value. The number of unique Microsoft Entra ID users that were seen.
func (m *Destination) GetUserCount()(*int32) {
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
func (m *Destination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
    if m.GetNetworkingProtocol() != nil {
        cast := (*m.GetNetworkingProtocol()).String()
        err := writer.WriteStringValue("networkingProtocol", &cast)
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
        err := writer.WriteInt32Value("threatCount", m.GetThreatCount())
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
    if m.GetTrafficType() != nil {
        cast := (*m.GetTrafficType()).String()
        err := writer.WriteStringValue("trafficType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Destination) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *Destination) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceCount sets the deviceCount property value. The number of unique devices that were seen.
func (m *Destination) SetDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstAccessDateTime sets the firstAccessDateTime property value. The firstAccessDateTime property
func (m *Destination) SetFirstAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("firstAccessDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetFqdn sets the fqdn property value. The fully qualified domain name (FQDN) of the destination.
func (m *Destination) SetFqdn(value *string)() {
    err := m.GetBackingStore().Set("fqdn", value)
    if err != nil {
        panic(err)
    }
}
// SetIp sets the ip property value. The internet protocol (IP) used to access the destination.
func (m *Destination) SetIp(value *string)() {
    err := m.GetBackingStore().Set("ip", value)
    if err != nil {
        panic(err)
    }
}
// SetLastAccessDateTime sets the lastAccessDateTime property value. The most recent access DateTime.
func (m *Destination) SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastAccessDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkingProtocol sets the networkingProtocol property value. The networkingProtocol property
func (m *Destination) SetNetworkingProtocol(value *NetworkingProtocol)() {
    err := m.GetBackingStore().Set("networkingProtocol", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Destination) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPort sets the port property value. The numeric identifier that is associated with a specific endpoint in a network.
func (m *Destination) SetPort(value *int32)() {
    err := m.GetBackingStore().Set("port", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatCount sets the threatCount property value. The threatCount property
func (m *Destination) SetThreatCount(value *int32)() {
    err := m.GetBackingStore().Set("threatCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBytesReceived sets the totalBytesReceived property value. The totalBytesReceived property
func (m *Destination) SetTotalBytesReceived(value *int64)() {
    err := m.GetBackingStore().Set("totalBytesReceived", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBytesSent sets the totalBytesSent property value. The totalBytesSent property
func (m *Destination) SetTotalBytesSent(value *int64)() {
    err := m.GetBackingStore().Set("totalBytesSent", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficType sets the trafficType property value. The trafficType property
func (m *Destination) SetTrafficType(value *TrafficType)() {
    err := m.GetBackingStore().Set("trafficType", value)
    if err != nil {
        panic(err)
    }
}
// SetTransactionCount sets the transactionCount property value. The number of transactions.
func (m *Destination) SetTransactionCount(value *int32)() {
    err := m.GetBackingStore().Set("transactionCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCount sets the userCount property value. The number of unique Microsoft Entra ID users that were seen.
func (m *Destination) SetUserCount(value *int32)() {
    err := m.GetBackingStore().Set("userCount", value)
    if err != nil {
        panic(err)
    }
}
// Destinationable 
type Destinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceCount()(*int32)
    GetFirstAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFqdn()(*string)
    GetIp()(*string)
    GetLastAccessDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNetworkingProtocol()(*NetworkingProtocol)
    GetOdataType()(*string)
    GetPort()(*int32)
    GetThreatCount()(*int32)
    GetTotalBytesReceived()(*int64)
    GetTotalBytesSent()(*int64)
    GetTrafficType()(*TrafficType)
    GetTransactionCount()(*int32)
    GetUserCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceCount(value *int32)()
    SetFirstAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFqdn(value *string)()
    SetIp(value *string)()
    SetLastAccessDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNetworkingProtocol(value *NetworkingProtocol)()
    SetOdataType(value *string)()
    SetPort(value *int32)()
    SetThreatCount(value *int32)()
    SetTotalBytesReceived(value *int64)()
    SetTotalBytesSent(value *int64)()
    SetTrafficType(value *TrafficType)()
    SetTransactionCount(value *int32)()
    SetUserCount(value *int32)()
}
