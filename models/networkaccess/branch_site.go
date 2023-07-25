package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// BranchSite 
type BranchSite struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewBranchSite instantiates a new branchSite and sets the default values.
func NewBranchSite()(*BranchSite) {
    m := &BranchSite{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateBranchSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBranchSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBranchSite(), nil
}
// GetBandwidthCapacity gets the bandwidthCapacity property value. Determines the maximum allowed Mbps (megabits per second) bandwidth from a branch site. The possible values are:250,500,750,1000.
func (m *BranchSite) GetBandwidthCapacity()(*int64) {
    val, err := m.GetBackingStore().Get("bandwidthCapacity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetConnectivityState gets the connectivityState property value. Determines the branch site status. The possible values are: pending, connected, inactive, error.
func (m *BranchSite) GetConnectivityState()(*ConnectivityState) {
    val, err := m.GetBackingStore().Get("connectivityState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ConnectivityState)
    }
    return nil
}
// GetCountry gets the country property value. The branch site is created in the specified country.
func (m *BranchSite) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceLinks gets the deviceLinks property value. Each unique CPE device associated with a branch is specified. Supports $expand.
func (m *BranchSite) GetDeviceLinks()([]DeviceLinkable) {
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
func (m *BranchSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetForwardingProfiles gets the forwardingProfiles property value. Each forwarding profile associated with a branch site is specified. Supports $expand.
func (m *BranchSite) GetForwardingProfiles()([]ForwardingProfileable) {
    val, err := m.GetBackingStore().Get("forwardingProfiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ForwardingProfileable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. last modified time.
func (m *BranchSite) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *BranchSite) GetName()(*string) {
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
func (m *BranchSite) GetRegion()(*Region) {
    val, err := m.GetBackingStore().Get("region")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Region)
    }
    return nil
}
// GetVersion gets the version property value. The branch version.
func (m *BranchSite) GetVersion()(*string) {
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
func (m *BranchSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetBandwidthCapacity sets the bandwidthCapacity property value. Determines the maximum allowed Mbps (megabits per second) bandwidth from a branch site. The possible values are:250,500,750,1000.
func (m *BranchSite) SetBandwidthCapacity(value *int64)() {
    err := m.GetBackingStore().Set("bandwidthCapacity", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectivityState sets the connectivityState property value. Determines the branch site status. The possible values are: pending, connected, inactive, error.
func (m *BranchSite) SetConnectivityState(value *ConnectivityState)() {
    err := m.GetBackingStore().Set("connectivityState", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. The branch site is created in the specified country.
func (m *BranchSite) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLinks sets the deviceLinks property value. Each unique CPE device associated with a branch is specified. Supports $expand.
func (m *BranchSite) SetDeviceLinks(value []DeviceLinkable)() {
    err := m.GetBackingStore().Set("deviceLinks", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardingProfiles sets the forwardingProfiles property value. Each forwarding profile associated with a branch site is specified. Supports $expand.
func (m *BranchSite) SetForwardingProfiles(value []ForwardingProfileable)() {
    err := m.GetBackingStore().Set("forwardingProfiles", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. last modified time.
func (m *BranchSite) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name.
func (m *BranchSite) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetRegion sets the region property value. The region property
func (m *BranchSite) SetRegion(value *Region)() {
    err := m.GetBackingStore().Set("region", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. The branch version.
func (m *BranchSite) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// BranchSiteable 
type BranchSiteable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthCapacity()(*int64)
    GetConnectivityState()(*ConnectivityState)
    GetCountry()(*string)
    GetDeviceLinks()([]DeviceLinkable)
    GetForwardingProfiles()([]ForwardingProfileable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetRegion()(*Region)
    GetVersion()(*string)
    SetBandwidthCapacity(value *int64)()
    SetConnectivityState(value *ConnectivityState)()
    SetCountry(value *string)()
    SetDeviceLinks(value []DeviceLinkable)()
    SetForwardingProfiles(value []ForwardingProfileable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetRegion(value *Region)()
    SetVersion(value *string)()
}
