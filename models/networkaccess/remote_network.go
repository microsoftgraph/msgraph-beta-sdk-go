package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type RemoteNetwork struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewRemoteNetwork instantiates a new RemoteNetwork and sets the default values.
func NewRemoteNetwork()(*RemoteNetwork) {
    m := &RemoteNetwork{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateRemoteNetworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRemoteNetworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteNetwork(), nil
}
// GetBandwidthCapacity gets the bandwidthCapacity property value. The bandwidthCapacity property
// returns a *int64 when successful
func (m *RemoteNetwork) GetBandwidthCapacity()(*int64) {
    val, err := m.GetBackingStore().Get("bandwidthCapacity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetConnectivityConfiguration gets the connectivityConfiguration property value. The connectivityConfiguration property
// returns a RemoteNetworkConnectivityConfigurationable when successful
func (m *RemoteNetwork) GetConnectivityConfiguration()(RemoteNetworkConnectivityConfigurationable) {
    val, err := m.GetBackingStore().Get("connectivityConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RemoteNetworkConnectivityConfigurationable)
    }
    return nil
}
// GetConnectivityState gets the connectivityState property value. The connectivityState property
// returns a *ConnectivityState when successful
func (m *RemoteNetwork) GetConnectivityState()(*ConnectivityState) {
    val, err := m.GetBackingStore().Get("connectivityState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConnectivityState)
    }
    return nil
}
// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *RemoteNetwork) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceLinks gets the deviceLinks property value. The deviceLinks property
// returns a []DeviceLinkable when successful
func (m *RemoteNetwork) GetDeviceLinks()([]DeviceLinkable) {
    val, err := m.GetBackingStore().Get("deviceLinks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceLinkable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RemoteNetwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bandwidthCapacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthCapacity(val)
        }
        return nil
    }
    res["connectivityConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRemoteNetworkConnectivityConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectivityConfiguration(val.(RemoteNetworkConnectivityConfigurationable))
        }
        return nil
    }
    res["connectivityState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectivityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectivityState(val.(*ConnectivityState))
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["deviceLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLinkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceLinkable)
                }
            }
            m.SetDeviceLinks(res)
        }
        return nil
    }
    res["forwardingProfiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateForwardingProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ForwardingProfileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ForwardingProfileable)
                }
            }
            m.SetForwardingProfiles(res)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val.(*Region))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetForwardingProfiles gets the forwardingProfiles property value. The forwardingProfiles property
// returns a []ForwardingProfileable when successful
func (m *RemoteNetwork) GetForwardingProfiles()([]ForwardingProfileable) {
    val, err := m.GetBackingStore().Get("forwardingProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ForwardingProfileable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
// returns a *Time when successful
func (m *RemoteNetwork) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RemoteNetwork) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegion gets the region property value. The region property
// returns a *Region when successful
func (m *RemoteNetwork) GetRegion()(*Region) {
    val, err := m.GetBackingStore().Get("region")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Region)
    }
    return nil
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *RemoteNetwork) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RemoteNetwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("bandwidthCapacity", m.GetBandwidthCapacity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectivityConfiguration", m.GetConnectivityConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetConnectivityState() != nil {
        cast := (*m.GetConnectivityState()).String()
        err = writer.WriteStringValue("connectivityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceLinks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceLinks()))
        for i, v := range m.GetDeviceLinks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("deviceLinks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetForwardingProfiles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardingProfiles()))
        for i, v := range m.GetForwardingProfiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("forwardingProfiles", cast)
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
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRegion() != nil {
        cast := (*m.GetRegion()).String()
        err = writer.WriteStringValue("region", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBandwidthCapacity sets the bandwidthCapacity property value. The bandwidthCapacity property
func (m *RemoteNetwork) SetBandwidthCapacity(value *int64)() {
    err := m.GetBackingStore().Set("bandwidthCapacity", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectivityConfiguration sets the connectivityConfiguration property value. The connectivityConfiguration property
func (m *RemoteNetwork) SetConnectivityConfiguration(value RemoteNetworkConnectivityConfigurationable)() {
    err := m.GetBackingStore().Set("connectivityConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectivityState sets the connectivityState property value. The connectivityState property
func (m *RemoteNetwork) SetConnectivityState(value *ConnectivityState)() {
    err := m.GetBackingStore().Set("connectivityState", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The country property
func (m *RemoteNetwork) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLinks sets the deviceLinks property value. The deviceLinks property
func (m *RemoteNetwork) SetDeviceLinks(value []DeviceLinkable)() {
    err := m.GetBackingStore().Set("deviceLinks", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardingProfiles sets the forwardingProfiles property value. The forwardingProfiles property
func (m *RemoteNetwork) SetForwardingProfiles(value []ForwardingProfileable)() {
    err := m.GetBackingStore().Set("forwardingProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *RemoteNetwork) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *RemoteNetwork) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetRegion sets the region property value. The region property
func (m *RemoteNetwork) SetRegion(value *Region)() {
    err := m.GetBackingStore().Set("region", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The version property
func (m *RemoteNetwork) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type RemoteNetworkable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthCapacity()(*int64)
    GetConnectivityConfiguration()(RemoteNetworkConnectivityConfigurationable)
    GetConnectivityState()(*ConnectivityState)
    GetCountry()(*string)
    GetDeviceLinks()([]DeviceLinkable)
    GetForwardingProfiles()([]ForwardingProfileable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetRegion()(*Region)
    GetVersion()(*string)
    SetBandwidthCapacity(value *int64)()
    SetConnectivityConfiguration(value RemoteNetworkConnectivityConfigurationable)()
    SetConnectivityState(value *ConnectivityState)()
    SetCountry(value *string)()
    SetDeviceLinks(value []DeviceLinkable)()
    SetForwardingProfiles(value []ForwardingProfileable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetRegion(value *Region)()
    SetVersion(value *string)()
}
