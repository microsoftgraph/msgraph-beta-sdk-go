package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelSite entity that represents a Microsoft Tunnel site
type MicrosoftTunnelSite struct {
    Entity
}
// NewMicrosoftTunnelSite instantiates a new microsoftTunnelSite and sets the default values.
func NewMicrosoftTunnelSite()(*MicrosoftTunnelSite) {
    m := &MicrosoftTunnelSite{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelSite(), nil
}
// GetDescription gets the description property value. The site's description (optional)
func (m *MicrosoftTunnelSite) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the site. This property is required when a site is created.
func (m *MicrosoftTunnelSite) GetDisplayName()(*string) {
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
func (m *MicrosoftTunnelSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["internalNetworkProbeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalNetworkProbeUrl(val)
        }
        return nil
    }
    res["microsoftTunnelConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftTunnelConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelConfiguration(val.(MicrosoftTunnelConfigurationable))
        }
        return nil
    }
    res["microsoftTunnelServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelServerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MicrosoftTunnelServerable)
                }
            }
            m.SetMicrosoftTunnelServers(res)
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
    res["publicAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicAddress(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["upgradeAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAutomatically(val)
        }
        return nil
    }
    res["upgradeAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    res["upgradeWindowEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowEndTime(val)
        }
        return nil
    }
    res["upgradeWindowStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowStartTime(val)
        }
        return nil
    }
    res["upgradeWindowUtcOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowUtcOffsetInMinutes(val)
        }
        return nil
    }
    return res
}
// GetInternalNetworkProbeUrl gets the internalNetworkProbeUrl property value. The site's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) GetInternalNetworkProbeUrl()(*string) {
    val, err := m.GetBackingStore().Get("internalNetworkProbeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMicrosoftTunnelConfiguration gets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelConfiguration()(MicrosoftTunnelConfigurationable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MicrosoftTunnelConfigurationable)
    }
    return nil
}
// GetMicrosoftTunnelServers gets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelServers()([]MicrosoftTunnelServerable) {
    val, err := m.GetBackingStore().Get("microsoftTunnelServers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MicrosoftTunnelServerable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftTunnelSite) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublicAddress gets the publicAddress property value. The site's public domain name or IP address
func (m *MicrosoftTunnelSite) GetPublicAddress()(*string) {
    val, err := m.GetBackingStore().Get("publicAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelSite) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetUpgradeAutomatically gets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
func (m *MicrosoftTunnelSite) GetUpgradeAutomatically()(*bool) {
    val, err := m.GetBackingStore().Get("upgradeAutomatically")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUpgradeAvailable gets the upgradeAvailable property value. The site provides the state of when an upgrade is available
func (m *MicrosoftTunnelSite) GetUpgradeAvailable()(*bool) {
    val, err := m.GetBackingStore().Get("upgradeAvailable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUpgradeWindowEndTime gets the upgradeWindowEndTime property value. The site's upgrade window end time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("upgradeWindowEndTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetUpgradeWindowStartTime gets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    val, err := m.GetBackingStore().Get("upgradeWindowStartTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    }
    return nil
}
// GetUpgradeWindowUtcOffsetInMinutes gets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
func (m *MicrosoftTunnelSite) GetUpgradeWindowUtcOffsetInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("upgradeWindowUtcOffsetInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("internalNetworkProbeUrl", m.GetInternalNetworkProbeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("microsoftTunnelConfiguration", m.GetMicrosoftTunnelConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftTunnelServers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelServers()))
        for i, v := range m.GetMicrosoftTunnelServers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicAddress", m.GetPublicAddress())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("upgradeAutomatically", m.GetUpgradeAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("upgradeAvailable", m.GetUpgradeAvailable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("upgradeWindowEndTime", m.GetUpgradeWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("upgradeWindowStartTime", m.GetUpgradeWindowStartTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upgradeWindowUtcOffsetInMinutes", m.GetUpgradeWindowUtcOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The site's description (optional)
func (m *MicrosoftTunnelSite) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the site. This property is required when a site is created.
func (m *MicrosoftTunnelSite) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetInternalNetworkProbeUrl sets the internalNetworkProbeUrl property value. The site's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) SetInternalNetworkProbeUrl(value *string)() {
    err := m.GetBackingStore().Set("internalNetworkProbeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelConfiguration sets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelConfiguration(value MicrosoftTunnelConfigurationable)() {
    err := m.GetBackingStore().Set("microsoftTunnelConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrosoftTunnelServers sets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelServers(value []MicrosoftTunnelServerable)() {
    err := m.GetBackingStore().Set("microsoftTunnelServers", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftTunnelSite) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPublicAddress sets the publicAddress property value. The site's public domain name or IP address
func (m *MicrosoftTunnelSite) SetPublicAddress(value *string)() {
    err := m.GetBackingStore().Set("publicAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance
func (m *MicrosoftTunnelSite) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeAutomatically sets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
func (m *MicrosoftTunnelSite) SetUpgradeAutomatically(value *bool)() {
    err := m.GetBackingStore().Set("upgradeAutomatically", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeAvailable sets the upgradeAvailable property value. The site provides the state of when an upgrade is available
func (m *MicrosoftTunnelSite) SetUpgradeAvailable(value *bool)() {
    err := m.GetBackingStore().Set("upgradeAvailable", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeWindowEndTime sets the upgradeWindowEndTime property value. The site's upgrade window end time of day
func (m *MicrosoftTunnelSite) SetUpgradeWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("upgradeWindowEndTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeWindowStartTime sets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) SetUpgradeWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    err := m.GetBackingStore().Set("upgradeWindowStartTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeWindowUtcOffsetInMinutes sets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
func (m *MicrosoftTunnelSite) SetUpgradeWindowUtcOffsetInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("upgradeWindowUtcOffsetInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// MicrosoftTunnelSiteable 
type MicrosoftTunnelSiteable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInternalNetworkProbeUrl()(*string)
    GetMicrosoftTunnelConfiguration()(MicrosoftTunnelConfigurationable)
    GetMicrosoftTunnelServers()([]MicrosoftTunnelServerable)
    GetOdataType()(*string)
    GetPublicAddress()(*string)
    GetRoleScopeTagIds()([]string)
    GetUpgradeAutomatically()(*bool)
    GetUpgradeAvailable()(*bool)
    GetUpgradeWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetUpgradeWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetUpgradeWindowUtcOffsetInMinutes()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInternalNetworkProbeUrl(value *string)()
    SetMicrosoftTunnelConfiguration(value MicrosoftTunnelConfigurationable)()
    SetMicrosoftTunnelServers(value []MicrosoftTunnelServerable)()
    SetOdataType(value *string)()
    SetPublicAddress(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetUpgradeAutomatically(value *bool)()
    SetUpgradeAvailable(value *bool)()
    SetUpgradeWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetUpgradeWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetUpgradeWindowUtcOffsetInMinutes(value *int32)()
}
