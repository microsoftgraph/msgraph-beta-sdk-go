package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type RemoteNetworkHealthEvent struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewRemoteNetworkHealthEvent instantiates a new RemoteNetworkHealthEvent and sets the default values.
func NewRemoteNetworkHealthEvent()(*RemoteNetworkHealthEvent) {
    m := &RemoteNetworkHealthEvent{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateRemoteNetworkHealthEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemoteNetworkHealthEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteNetworkHealthEvent(), nil
}
// GetBgpRoutesAdvertisedCount gets the bgpRoutesAdvertisedCount property value. The number of BGP routes advertised through tunnel.
// returns a *int32 when successful
func (m *RemoteNetworkHealthEvent) GetBgpRoutesAdvertisedCount()(*int32) {
    val, err := m.GetBackingStore().Get("bgpRoutesAdvertisedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time of the original event generation in UTC. Supports $filter (ge, le) and $orderby.
// returns a *Time when successful
func (m *RemoteNetworkHealthEvent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description of the event.
// returns a *string when successful
func (m *RemoteNetworkHealthEvent) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDestinationIp gets the destinationIp property value. The IP address of the destination.
// returns a *string when successful
func (m *RemoteNetworkHealthEvent) GetDestinationIp()(*string) {
    val, err := m.GetBackingStore().Get("destinationIp")
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
func (m *RemoteNetworkHealthEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bgpRoutesAdvertisedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBgpRoutesAdvertisedCount(val)
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
    res["destinationIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationIp(val)
        }
        return nil
    }
    res["receivedBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedBytes(val)
        }
        return nil
    }
    res["remoteNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteNetworkId(val)
        }
        return nil
    }
    res["sentBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentBytes(val)
        }
        return nil
    }
    res["sourceIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIp(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteNetworkStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RemoteNetworkStatus))
        }
        return nil
    }
    return res
}
// GetReceivedBytes gets the receivedBytes property value. The number of bytes sent from the destination to the source.
// returns a *int64 when successful
func (m *RemoteNetworkHealthEvent) GetReceivedBytes()(*int64) {
    val, err := m.GetBackingStore().Get("receivedBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRemoteNetworkId gets the remoteNetworkId property value. A unique identifier for each remoteNetwork site. Supports $filter (eq).
// returns a *string when successful
func (m *RemoteNetworkHealthEvent) GetRemoteNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("remoteNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSentBytes gets the sentBytes property value. The number of bytes sent from the source to the destination for the connection or session.
// returns a *int64 when successful
func (m *RemoteNetworkHealthEvent) GetSentBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sentBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSourceIp gets the sourceIp property value. The public IP address.
// returns a *string when successful
func (m *RemoteNetworkHealthEvent) GetSourceIp()(*string) {
    val, err := m.GetBackingStore().Get("sourceIp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *RemoteNetworkStatus when successful
func (m *RemoteNetworkHealthEvent) GetStatus()(*RemoteNetworkStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RemoteNetworkStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RemoteNetworkHealthEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("bgpRoutesAdvertisedCount", m.GetBgpRoutesAdvertisedCount())
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
        err = writer.WriteStringValue("destinationIp", m.GetDestinationIp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("receivedBytes", m.GetReceivedBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remoteNetworkId", m.GetRemoteNetworkId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sentBytes", m.GetSentBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceIp", m.GetSourceIp())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBgpRoutesAdvertisedCount sets the bgpRoutesAdvertisedCount property value. The number of BGP routes advertised through tunnel.
func (m *RemoteNetworkHealthEvent) SetBgpRoutesAdvertisedCount(value *int32)() {
    err := m.GetBackingStore().Set("bgpRoutesAdvertisedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time of the original event generation in UTC. Supports $filter (ge, le) and $orderby.
func (m *RemoteNetworkHealthEvent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the event.
func (m *RemoteNetworkHealthEvent) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDestinationIp sets the destinationIp property value. The IP address of the destination.
func (m *RemoteNetworkHealthEvent) SetDestinationIp(value *string)() {
    err := m.GetBackingStore().Set("destinationIp", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedBytes sets the receivedBytes property value. The number of bytes sent from the destination to the source.
func (m *RemoteNetworkHealthEvent) SetReceivedBytes(value *int64)() {
    err := m.GetBackingStore().Set("receivedBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteNetworkId sets the remoteNetworkId property value. A unique identifier for each remoteNetwork site. Supports $filter (eq).
func (m *RemoteNetworkHealthEvent) SetRemoteNetworkId(value *string)() {
    err := m.GetBackingStore().Set("remoteNetworkId", value)
    if err != nil {
        panic(err)
    }
}
// SetSentBytes sets the sentBytes property value. The number of bytes sent from the source to the destination for the connection or session.
func (m *RemoteNetworkHealthEvent) SetSentBytes(value *int64)() {
    err := m.GetBackingStore().Set("sentBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceIp sets the sourceIp property value. The public IP address.
func (m *RemoteNetworkHealthEvent) SetSourceIp(value *string)() {
    err := m.GetBackingStore().Set("sourceIp", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *RemoteNetworkHealthEvent) SetStatus(value *RemoteNetworkStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type RemoteNetworkHealthEventable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBgpRoutesAdvertisedCount()(*int32)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDestinationIp()(*string)
    GetReceivedBytes()(*int64)
    GetRemoteNetworkId()(*string)
    GetSentBytes()(*int64)
    GetSourceIp()(*string)
    GetStatus()(*RemoteNetworkStatus)
    SetBgpRoutesAdvertisedCount(value *int32)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDestinationIp(value *string)()
    SetReceivedBytes(value *int64)()
    SetRemoteNetworkId(value *string)()
    SetSentBytes(value *int64)()
    SetSourceIp(value *string)()
    SetStatus(value *RemoteNetworkStatus)()
}
