package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeviceLink 
type DeviceLink struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDeviceLink instantiates a new deviceLink and sets the default values.
func NewDeviceLink()(*DeviceLink) {
    m := &DeviceLink{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDeviceLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceLink(), nil
}
// GetBandwidthCapacityInMbps gets the bandwidthCapacityInMbps property value. Determines the maximum allowed Mbps (megabits per second) bandwidth from a branch site. The possible values are:250,500,750,1000.
func (m *DeviceLink) GetBandwidthCapacityInMbps()(*BandwidthCapacityInMbps) {
    val, err := m.GetBackingStore().Get("bandwidthCapacityInMbps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*BandwidthCapacityInMbps)
    }
    return nil
}
// GetBgpConfiguration gets the bgpConfiguration property value. The bgpConfiguration property
func (m *DeviceLink) GetBgpConfiguration()(BgpConfigurationable) {
    val, err := m.GetBackingStore().Get("bgpConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(BgpConfigurationable)
    }
    return nil
}
// GetDeviceVendor gets the deviceVendor property value. The deviceVendor property
func (m *DeviceLink) GetDeviceVendor()(*DeviceVendor) {
    val, err := m.GetBackingStore().Get("deviceVendor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceVendor)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bandwidthCapacityInMbps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBandwidthCapacityInMbps)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthCapacityInMbps(val.(*BandwidthCapacityInMbps))
        }
        return nil
    }
    res["bgpConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBgpConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBgpConfiguration(val.(BgpConfigurationable))
        }
        return nil
    }
    res["deviceVendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceVendor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceVendor(val.(*DeviceVendor))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
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
    res["redundancyConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRedundancyConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedundancyConfiguration(val.(RedundancyConfigurationable))
        }
        return nil
    }
    res["tunnelConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTunnelConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTunnelConfiguration(val.(TunnelConfigurationable))
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. Specifies the client IPv4 of the link
func (m *DeviceLink) GetIpAddress()(*string) {
    val, err := m.GetBackingStore().Get("ipAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. last modified time.
func (m *DeviceLink) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetName gets the name property value. Name.
func (m *DeviceLink) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRedundancyConfiguration gets the redundancyConfiguration property value. The redundancyConfiguration property
func (m *DeviceLink) GetRedundancyConfiguration()(RedundancyConfigurationable) {
    val, err := m.GetBackingStore().Get("redundancyConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RedundancyConfigurationable)
    }
    return nil
}
// GetTunnelConfiguration gets the tunnelConfiguration property value. The tunnelConfiguration property
func (m *DeviceLink) GetTunnelConfiguration()(TunnelConfigurationable) {
    val, err := m.GetBackingStore().Get("tunnelConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TunnelConfigurationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBandwidthCapacityInMbps() != nil {
        cast := (*m.GetBandwidthCapacityInMbps()).String()
        err = writer.WriteStringValue("bandwidthCapacityInMbps", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bgpConfiguration", m.GetBgpConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceVendor() != nil {
        cast := (*m.GetDeviceVendor()).String()
        err = writer.WriteStringValue("deviceVendor", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
    {
        err = writer.WriteObjectValue("redundancyConfiguration", m.GetRedundancyConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tunnelConfiguration", m.GetTunnelConfiguration())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBandwidthCapacityInMbps sets the bandwidthCapacityInMbps property value. Determines the maximum allowed Mbps (megabits per second) bandwidth from a branch site. The possible values are:250,500,750,1000.
func (m *DeviceLink) SetBandwidthCapacityInMbps(value *BandwidthCapacityInMbps)() {
    err := m.GetBackingStore().Set("bandwidthCapacityInMbps", value)
    if err != nil {
        panic(err)
    }
}
// SetBgpConfiguration sets the bgpConfiguration property value. The bgpConfiguration property
func (m *DeviceLink) SetBgpConfiguration(value BgpConfigurationable)() {
    err := m.GetBackingStore().Set("bgpConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceVendor sets the deviceVendor property value. The deviceVendor property
func (m *DeviceLink) SetDeviceVendor(value *DeviceVendor)() {
    err := m.GetBackingStore().Set("deviceVendor", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddress sets the ipAddress property value. Specifies the client IPv4 of the link
func (m *DeviceLink) SetIpAddress(value *string)() {
    err := m.GetBackingStore().Set("ipAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. last modified time.
func (m *DeviceLink) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name.
func (m *DeviceLink) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetRedundancyConfiguration sets the redundancyConfiguration property value. The redundancyConfiguration property
func (m *DeviceLink) SetRedundancyConfiguration(value RedundancyConfigurationable)() {
    err := m.GetBackingStore().Set("redundancyConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetTunnelConfiguration sets the tunnelConfiguration property value. The tunnelConfiguration property
func (m *DeviceLink) SetTunnelConfiguration(value TunnelConfigurationable)() {
    err := m.GetBackingStore().Set("tunnelConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// DeviceLinkable 
type DeviceLinkable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthCapacityInMbps()(*BandwidthCapacityInMbps)
    GetBgpConfiguration()(BgpConfigurationable)
    GetDeviceVendor()(*DeviceVendor)
    GetIpAddress()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetRedundancyConfiguration()(RedundancyConfigurationable)
    GetTunnelConfiguration()(TunnelConfigurationable)
    SetBandwidthCapacityInMbps(value *BandwidthCapacityInMbps)()
    SetBgpConfiguration(value BgpConfigurationable)()
    SetDeviceVendor(value *DeviceVendor)()
    SetIpAddress(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetRedundancyConfiguration(value RedundancyConfigurationable)()
    SetTunnelConfiguration(value TunnelConfigurationable)()
}
